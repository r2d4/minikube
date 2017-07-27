package bootstrapper

import (
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/drivers"
	"k8s.io/minikube/pkg/minikube/cluster"
)

type Bootstrapper interface {
	StartCluster(api libmachine.API, k8s cluster.KubernetesConfig) error
	UpdateCluster(d drivers.Driver, ck8s cluster.KubernetesConfig) error
	GetClusterLogs(api libmachine.API, follow bool) (string, error)
	GetClusterStatus(api libmachine.API) (string, error)
}
