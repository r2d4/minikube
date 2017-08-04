/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package localkube

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"

	"github.com/docker/machine/libmachine/state"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"k8s.io/minikube/pkg/minikube/assets"
	"k8s.io/minikube/pkg/minikube/boostrapper"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/minikube/sshutil"
)

// LocalkubeBootstrapper starts and updates a localkube backed cluster.
type LocalkubeBootstrapper struct {
	c *ssh.Client
	s *ssh.Session
}

func (l *LocalkubeBootstrapper) RunCommand(cmd string) (string, error) {
	var err error
	if l.s == nil {
		l.s, err = l.c.NewSession()
		if err != nil {
			return "", errors.Wrap(err, "getting ssh session")
		}
	}
	o, err := l.s.CombinedOutput(cmd)
	if err != nil {
		return "", errors.Wrapf(err, "running ssh cmd: %s", cmd)
	}
	out := string(o)
	glog.V(7).Infof("SSH Command: %s\n Output: %s", cmd, out)

	return out, nil
}

// StartCluster starts a k8s cluster on the specified Host.
func (l *LocalkubeBootstrapper) StartCluster(kubernetesConfig bootstrapper.KubernetesConfig) error {
	startCommand, err := GetStartCommand(kubernetesConfig)
	if err != nil {
		return errors.Wrapf(err, "Error generating start command: %s", err)
	}

	l.RunCommand(startCommand)
	if err != nil {
		return errors.Wrapf(err, "Error running ssh command: %s", startCommand)
	}
	return nil
}

func (l *LocalkubeBootstrapper) RestartCluster(kubernetesConfig bootstrapper.KubernetesConfig) error {
	return l.StartCluster(kubernetesConfig)
}

//TODO(mrick) get rid of driver dependency here
func (l *LocalkubeBootstrapper) UpdateCluster(config bootstrapper.KubernetesConfig, drivername string) error {
	copyableFiles := []assets.CopyableFile{}
	var localkubeFile assets.CopyableFile
	var err error

	//add url/file/bundled localkube to file list
	if localkubeURIWasSpecified(config) && config.KubernetesVersion != constants.DefaultKubernetesVersion {
		lCacher := localkubeCacher{config}
		localkubeFile, err = lCacher.fetchLocalkubeFromURI()
		if err != nil {
			return errors.Wrap(err, "Error updating localkube from uri")
		}
	} else {
		localkubeFile = assets.NewMemoryAsset("out/localkube", "/usr/local/bin", "localkube", "0777")
	}
	copyableFiles = append(copyableFiles, localkubeFile)

	// add addons to file list
	// custom addons
	assets.AddMinikubeAddonsDirToAssets(&copyableFiles)
	// bundled addons
	for _, addonBundle := range assets.Addons {
		if isEnabled, err := addonBundle.IsEnabled(); err == nil && isEnabled {
			for _, addon := range addonBundle.Assets {
				copyableFiles = append(copyableFiles, addon)
			}
		} else if err != nil {
			return err
		}
	}

	if drivername == "none" {
		// transfer files to correct place on filesystem
		for _, f := range copyableFiles {
			if err := assets.CopyFileLocal(f); err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range copyableFiles {
		if err := sshutil.TransferFileSSHSession(f, l.s); err != nil {
			return err
		}
	}
	return nil
}

// GetClusterStatus gets the status of localkube from the host VM.
func (l *LocalkubeBootstrapper) GetClusterStatus() (string, error) {
	s, err := l.RunCommand(localkubeStatusCommand)
	if err != nil {
		return "", err
	}
	s = strings.TrimSpace(s)
	if state.Running.String() == s {
		return state.Running.String(), nil
	} else if state.Stopped.String() == s {
		return state.Stopped.String(), nil
	} else {
		return "", fmt.Errorf("Error: Unrecognize output from GetLocalkubeStatus: %s", s)
	}
}

// GetClusterLogs gets the localkube logs of the host VM.
// If follow is specified, it will tail the logs
func (l *LocalkubeBootstrapper) GetClusterLogs(follow bool) (string, error) {
	logsCommand, err := GetLogsCommand(follow)
	if err != nil {
		return "", errors.Wrap(err, "Error getting logs command")
	}

	s, err := l.RunCommand(logsCommand)

	if err != nil {
		return s, err
	}
	return s, nil
}
