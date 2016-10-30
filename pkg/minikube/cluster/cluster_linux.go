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

package cluster

import (
	"fmt"
	"path/filepath"

	kvmDriver "github.com/dhiltgen/docker-machine-kvm"
	"github.com/docker/machine/libmachine/drivers"
	"k8s.io/minikube/pkg/minikube/constants"
)

func createKVMHost(config MachineConfig) drivers.Driver {
	d := kvmDriver.NewDriver(constants.MachineName, constants.Minipath)
	o, _ := d.(*kvmDriver.Driver)

	o.Memory = config.Memory
	o.CPU = config.CPUs
	o.Network = config.KvmNetwork
	o.PrivateNetwork = "docker-machines"
	o.Boot2DockerURL = config.GetISOFileURI()
	o.DiskSize = config.DiskSize
	o.DiskPath = filepath.Join(constants.Minipath, "machines", constants.MachineName, fmt.Sprintf("%s.img", constants.MachineName))
	o.ISO = filepath.Join(constants.Minipath, "machines", constants.MachineName, "boot2docker.iso")
	o.CacheMode = "default"
	o.IOMode = "threads"

	return o
}
