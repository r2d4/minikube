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

package notify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/blang/semver"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"k8s.io/minikube/pkg/minikube/config"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/version"
)

const updateLinkPrefix = "https://github.com/kubernetes/minikube/releases/tag/v"
const timeout = time.Duration(5 * time.Second)

var (
	timeLayout                = time.RFC1123
	lastUpdateCheckFilePath   = constants.MakeMiniPath("last_update_check")
	githubMinikubeReleasesURL = "https://storage.googleapis.com/minikube/releases.json"
	NotifyMsg                 = make(chan string)
)

func MaybePrintUpdateTextFromGithub() {
	MaybePrintUpdateText(githubMinikubeReleasesURL, lastUpdateCheckFilePath)
	close(NotifyMsg)
}

func MaybePrintUpdateText(url string, lastUpdatePath string) {
	if !shouldCheckURLVersion(lastUpdatePath) {
		return
	}
	latestVersion, err := getLatestVersionFromURL(url)
	if err != nil {
		glog.Errorln(err)
		return
	}
	localVersion, err := version.GetSemverVersion()
	if err != nil {
		glog.Errorln(err)
		return
	}
	if localVersion.Compare(latestVersion) < 0 {
		writeTimeToFile(lastUpdateCheckFilePath, time.Now().UTC())
		NotifyMsg <- fmt.Sprintf(`There is a newer version of minikube available (%s%s).  Download it here:
%s%s
To disable this notification, add WantUpdateNotification: False to the json config file at %s
(you may have to create the file config.json in this folder if you have no previous configuration)
`,
			version.VersionPrefix, latestVersion, updateLinkPrefix, latestVersion, constants.MakeMiniPath("config"))
	}
}

func shouldCheckURLVersion(filePath string) bool {
	if !viper.GetBool(config.WantUpdateNotification) {
		return false
	}
	lastUpdateTime := getTimeFromFileIfExists(filePath)
	if time.Since(lastUpdateTime).Hours() < viper.GetFloat64(config.ReminderWaitPeriodInHours) {
		return false
	}
	return true
}

type release struct {
	Name string
}

type releases []release

func getJson(url string, target *releases) error {
	client := http.Client{
		Timeout: timeout,
	}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getLatestVersionFromURL(url string) (semver.Version, error) {
	var releases releases
	glog.Infof("Checking for updates...")
	if err := getJson(url, &releases); err != nil {
		return semver.Version{}, err
	}
	if len(releases) == 0 {
		return semver.Version{}, fmt.Errorf("There were no json releases at the url specified: %s", url)
	}
	latestVersionString := releases[0].Name
	return semver.Make(strings.TrimPrefix(latestVersionString, version.VersionPrefix))
}

func writeTimeToFile(path string, inputTime time.Time) error {
	err := ioutil.WriteFile(path, []byte(inputTime.Format(timeLayout)), 0644)
	if err != nil {
		return fmt.Errorf("Error writing current update time to file: ", err)
	}
	return nil
}

func getTimeFromFileIfExists(path string) time.Time {
	lastUpdateCheckTime, err := ioutil.ReadFile(path)
	if err != nil {
		return time.Time{}
	}
	timeInFile, err := time.Parse(timeLayout, string(lastUpdateCheckTime))
	if err != nil {
		return time.Time{}
	}
	return timeInFile
}
