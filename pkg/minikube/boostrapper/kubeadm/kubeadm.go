package kubeadm

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"k8s.io/minikube/pkg/minikube/boostrapper"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/minikube/sshutil"
	"k8s.io/minikube/pkg/util"
)

type KubeadmBootstrapper struct {
	c *ssh.Client
}

func NewKubeadmBootstrapper(c *ssh.Client) *KubeadmBootstrapper {
	k := &KubeadmBootstrapper{c: c}
	return k
}

func (k *KubeadmBootstrapper) RestartCluster(k8s bootstrapper.KubernetesConfig) error {
	tmpFile := "/tmp/cert.conf"

	restartTmpl := "sudo /usr/bin/kubeadm alpha phase kubeconfig client-certs"
	restartTmpl += " --cert-dir {{.CertDir}}"
	restartTmpl += " --server {{.IP}}"
	restartTmpl += " --client-name {{.MachineName}}"

	// Output to temp file, since we will have to write this file to a few places.
	restartTmpl += " > " + tmpFile
	t := template.Must(template.New("restartTmpl").Parse(restartTmpl))

	opts := struct {
		CertDir     string
		IP          string
		MachineName string
	}{
		CertDir:     util.DefaultCertPath,
		IP:          k8s.NodeIP,
		MachineName: k8s.NodeName,
	}

	b := bytes.Buffer{}
	if err := t.Execute(&b, opts); err != nil {
		return err
	}

	_, err := sshutil.RunCommandOutput(k.c, b.String())
	if err != nil {
		return err
	}

	dsts := []string{"admin.conf", "controller-manager.conf", "kubelet.conf", "scheduler.conf"}
	for _, dst := range dsts {
		cmd := fmt.Sprintf("sudo cp %s %s", tmpFile, filepath.Join("/etc", "kubernetes", dst))
		_, err := sshutil.RunCommandOutput(k.c, cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k *KubeadmBootstrapper) StartCluster(k8s bootstrapper.KubernetesConfig) error {
	kubeadmTmpl := "sudo /usr/bin/kubeadm init"
	kubeadmTmpl += " --cert-dir {{.CertDir}}"
	kubeadmTmpl += " --service-cidr {{.ServiceCIDR}}"
	kubeadmTmpl += " --apiserver-advertise-address {{.AdvertiseAddress}}"
	kubeadmTmpl += " --apiserver-bind-port {{.APIServerPort}}"
	t := template.Must(template.New("kubeadmTmpl").Parse(kubeadmTmpl))

	opts := struct {
		CertDir          string
		ServiceCIDR      string
		AdvertiseAddress string
		APIServerPort    int
	}{
		CertDir:          util.DefaultCertPath,
		ServiceCIDR:      util.DefaultServiceCIDR,
		AdvertiseAddress: k8s.NodeIP,
		APIServerPort:    util.APIServerPort,
	}

	b := bytes.Buffer{}
	if err := t.Execute(&b, opts); err != nil {
		return err
	}

	_, err := sshutil.RunCommandOutput(k.c, b.String())
	if err != nil {
		return err
	}

	return nil
}

func (k *KubeadmBootstrapper) UpdateCluster(k8s bootstrapper.KubernetesConfig, drivername string) error {
	if err := moveFiles(k.c); err != nil {
		return errors.Wrap(err, "moving kubeadm files into vm")
	}

	_, err := sshutil.RunCommandOutput(k.c, `
sudo cp /kubeadm/kubelet /usr/bin/kubelet &&
sudo cp /kubeadm/kubeadm /usr/bin/kubeadm &&
sudo cp /kubeadm/kubelet.service /lib/systemd/system &&
sudo mkdir -p /etc/systemd/system/kubelet.service.d/ && 
sudo cp /kubeadm/10-kubeadm.conf /etc/systemd/system/kubelet.service.d/10-kubeadm.conf &&
sudo systemctl daemon-reload &&
sudo systemctl enable kubelet &&
sudo systemctl start kubelet
`)
	if err != nil {
		return err
	}

	return nil
}

func moveFiles(c *ssh.Client) error {
	dst := "/kubeadm"
	folders := map[string]string{
		constants.MakeMiniPath("out", "kubeadm", "bin"): "0641",
		constants.MakeMiniPath("deploy"):                "0640",
	}
	for src, perm := range folders {
		if err := sshutil.TransferMinikubeFolderToVM(src, dst, perm, c); err != nil {
			return err
		}
	}

	return nil
}

func (k *KubeadmBootstrapper) GetClusterLogs(follow bool) (string, error) {
	out, err := sshutil.RunCommandOutput(k.c, "journalctl -u kubelet")
	if err != nil {
		return "", errors.Wrap(err, "getting logs")
	}
	return out, nil
}

func (k *KubeadmBootstrapper) GetClusterStatus() (string, error) {
	return "", nil
}
