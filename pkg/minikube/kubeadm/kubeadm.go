package kubeadm

import (
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/pkg/errors"
	"k8s.io/minikube/pkg/minikube/assets"
	"k8s.io/minikube/pkg/minikube/cluster"
	"k8s.io/minikube/pkg/minikube/sshutil"
)

func StartCluster(api libmachine.API, k8s cluster.KubernetesConfig) error {
	//Runs "start localkube""

	return nil
}

func UpdateCluster(d drivers.Driver, k8s cluster.KubernetesConfig) error {
	// "copies localkube into vm"
	// adds minikube addons to directory

	files, err := assets.AddMinikubeDirToAssets("kubeadm", "/kubeadm")
	if err != nil {
		return err
	}

	client, err := sshutil.NewSSHClient(d)
	if err != nil {
		return errors.Wrap(err, "Error creating new ssh client")
	}

	for _, f := range files {
		if err := sshutil.TransferFile(f, client); err != nil {
			return err
		}
	}

	return nil
}
