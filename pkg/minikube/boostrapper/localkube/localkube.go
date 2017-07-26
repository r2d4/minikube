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

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/state"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"k8s.io/minikube/pkg/minikube/assets"
	"k8s.io/minikube/pkg/minikube/cluster"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/minikube/sshutil"
)

// LocalkubeBootstrapper starts and updates a localkube backed cluster.
type LocalkubeBootstrapper struct{}

// StartCluster starts a k8s cluster on the specified Host.
func (*LocalkubeBootstrapper) StartCluster(api libmachine.API, kubernetesConfig cluster.KubernetesConfig) error {
	h, err := cluster.CheckIfApiExistsAndLoad(api)
	if err != nil {
		return errors.Wrap(err, "Error checking that api exists and loading it")
	}

	startCommand, err := GetStartCommand(kubernetesConfig)
	if err != nil {
		return errors.Wrapf(err, "Error generating start command: %s", err)
	}
	glog.Infoln(startCommand)
	output, err := cluster.RunCommand(h, startCommand, true)
	glog.Infoln(output)
	if err != nil {
		return errors.Wrapf(err, "Error running ssh command: %s", startCommand)
	}
	return nil
}

func (*LocalkubeBootstrapper) UpdateCluster(d drivers.Driver, config cluster.KubernetesConfig) error {
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

	if d.DriverName() == "none" {
		// transfer files to correct place on filesystem
		for _, f := range copyableFiles {
			if err := assets.CopyFileLocal(f); err != nil {
				return err
			}
		}
		return nil
	}

	// transfer files to vm via SSH
	client, err := sshutil.NewSSHClient(d)
	if err != nil {
		return errors.Wrap(err, "Error creating new ssh client")
	}

	for _, f := range copyableFiles {
		if err := sshutil.TransferFile(f, client); err != nil {
			return err
		}
	}
	return nil
}

// GetClusterStatus gets the status of localkube from the host VM.
func (*LocalkubeBootstrapper) GetClusterStatus(api libmachine.API) (string, error) {
	h, err := cluster.CheckIfApiExistsAndLoad(api)
	if err != nil {
		return "", err
	}
	s, err := cluster.RunCommand(h, localkubeStatusCommand, false)
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
func (*LocalkubeBootstrapper) GetClusterLogs(api libmachine.API, follow bool) (string, error) {
	h, err := cluster.CheckIfApiExistsAndLoad(api)
	if err != nil {
		return "", errors.Wrap(err, "Error checking that api exists and loading it")
	}
	logsCommand, err := GetLogsCommand(follow)
	if err != nil {
		return "", errors.Wrap(err, "Error getting logs command")
	}
	if follow {
		c, err := h.CreateSSHClient()
		if err != nil {
			return "", errors.Wrap(err, "Error creating ssh client")
		}
		err = c.Shell(logsCommand)
		if err != nil {
			return "", errors.Wrap(err, "error ssh shell")
		}
		return "", err
	}
	s, err := cluster.RunCommand(h, logsCommand, false)

	if err != nil {
		return s, err
	}
	return s, nil
}
