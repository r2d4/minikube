package kubeadm

import (
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/pkg/errors"
	"k8s.io/minikube/pkg/minikube/cluster"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/minikube/sshutil"
)

func StartCluster(api libmachine.API, k8s cluster.KubernetesConfig) error {
	//Runs "start localkube""

	return nil
}

func UpdateCluster(d drivers.Driver, k8s cluster.KubernetesConfig) error {
	// "copies localkube into vm"
	// adds minikube addons to directory
	client, err := sshutil.NewSSHClient(d)
	if err != nil {
		return errors.Wrap(err, "Error creating new ssh client")
	}

	folders := map[string]string{
		constants.MakeMiniPath("kubeadm"): "/kubeadm",
		constants.MakeMiniPath("deploy"):  "/kubeadm",
	}

	for src, dst := range folders {
		if err := sshutil.TransferHostFolderToVM(src, dst, client); err != nil {
			return err
		}
	}

	return nil
}
