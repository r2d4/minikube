package kubeadm

import (
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/pkg/errors"
	"k8s.io/minikube/pkg/minikube/cluster"
	"k8s.io/minikube/pkg/minikube/sshutil"
)

type KubeadmBootstrapper struct{}

func (*KubeadmBootstrapper) StartCluster(api libmachine.API, k8s cluster.KubernetesConfig) error {
	//Runs "start localkube""

	return nil
}

func (*KubeadmBootstrapper) UpdateCluster(d drivers.Driver, k8s cluster.KubernetesConfig) error {
	// "copies localkube into vm"
	// adds minikube addons to directory
	if err := moveFiles(d); err != nil {
		return errors.Wrap(err, "moving kubeadm files into vm")
	}

	client, _ := sshutil.NewSSHClient(d)
	sess, _ := client.NewSession()
	sess.Run("sudo cp /kubeadm/kubelet /usr/bin/kubelet")
	sess.Run("sudo cp /kubeadm/kubeadm /usr/bin/kubeadm")
	sess.Run("sudo cp /kubeadm/kubelet.service /lib/systemd/system")
	sess.Run("sudo mkdir -p /etc/systemd/system/kubelet.service.d/ && sudo cp /kubeadm/10-kubeadm.conf /etc/systemd/system/kubelet.service.d/10-kubeadm.conf")
	sess.Run("sudo systemctl daemon-reload")
	sess.Run("sudo systemctl enable kubelet")
	sess.Run("sudo systemctl start kubelet")
	// sess.Run("sudo /usr/bin/kubeadm init")

	//var unitPath = "/etc/systemd/system/kubelet.service"

	return nil
}

func moveFiles(d drivers.Driver) error {
	client, err := sshutil.NewSSHClient(d)
	if err != nil {
		return errors.Wrap(err, "Error creating new ssh client")
	}

	if err := sshutil.TransferMinikubeFolderToVM("bin", "/kubeadm", "0641", client); err != nil {
		return err
	}

	if err := sshutil.TransferMinikubeFolderToVM("deploy", "/kubeadm", "0640", client); err != nil {
		return err
	}

	return nil
}

func (*KubeadmBootstrapper) GetClusterLogs(api libmachine.API, follow bool) (string, error) {
	return "", nil
}

func (*KubeadmBootstrapper) GetClusterStatus(api libmachine.API) (string, error) {
	return "", nil
}
