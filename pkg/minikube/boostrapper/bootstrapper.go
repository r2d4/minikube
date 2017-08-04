package bootstrapper

import "k8s.io/minikube/pkg/util"

type Bootstrapper interface {
	StartCluster(KubernetesConfig) error
	UpdateCluster(cfg KubernetesConfig, drivername string) error
	RestartCluster(cfg KubernetesConfig) error
	GetClusterLogs(follow bool) (string, error)
	GetClusterStatus() (string, error)
}

// KubernetesConfig contains the parameters used to configure the VM Kubernetes.
type KubernetesConfig struct {
	KubernetesVersion string
	NodeIP            string
	NodeName          string
	APIServerName     string
	DNSDomain         string
	ContainerRuntime  string
	NetworkPlugin     string
	FeatureGates      string
	ExtraOptions      util.ExtraOptionSlice
}
