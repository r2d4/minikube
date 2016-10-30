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
	"bytes"
	"crypto"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/docker/machine/drivers/virtualbox"
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/auth"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/engine"
	"github.com/docker/machine/libmachine/host"
	"github.com/docker/machine/libmachine/provision"
	"github.com/docker/machine/libmachine/state"
	"github.com/docker/machine/libmachine/swarm"
	"github.com/golang/glog"
	download "github.com/jimmidyson/go-download"
	"github.com/pkg/errors"
	kubeapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"

	"k8s.io/minikube/pkg/minikube/assets"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/minikube/sshutil"
	"k8s.io/minikube/pkg/util"
)

var (
	certs = []string{"ca.crt", "ca.key", "apiserver.crt", "apiserver.key"}
)

const fileScheme = "file"

var certsDir = constants.Minipath

//This init function is used to set the logtostderr variable to false so that INFO level log info does not clutter the CLI
//INFO lvl logging is displayed due to the kubernetes api calling flag.Set("logtostderr", "true") in its init()
//see: https://github.com/kubernetes/kubernetes/blob/master/pkg/util/logs.go#L32-34
func init() {
	flag.Set("logtostderr", "false")
}

// StartHost starts a host VM.
func StartHost(config MachineConfig) (drivers.Driver, error) {
	//TODO:r2d4 Exists

	return createMachine(config)

	driver := Load(config)

	s, err := driver.GetState()
	glog.Infoln("Machine state: ", s)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting state for host")
	}

	if s != state.Running {
		if err := driver.Start(); err != nil {
			return nil, errors.Wrapf(err, "Error starting stopped host")
		}
		//TODO SAVE
	}

	if err := Provision(config); err != nil {
		return nil, errors.Wrap(&util.RetriableError{Err: err}, "Error configuring auth on host")
	}
	return driver, nil
}

func Provision(config MachineConfig) error {
	//filestore := persist.NewFilestore(constants.Minipath, certsDir, certsDir)
	driver := Load(config)
	provisioner := provision.Boot2DockerProvisioner{
		Driver: driver,
	}

	authOptions := auth.Options{
		CertDir:   certsDir,
		StorePath: certsDir,
	}

	engineOptions := engine.Options{
		Env:              config.DockerEnv,
		InsecureRegistry: config.InsecureRegistry,
		RegistryMirror:   config.RegistryMirror,
	}

	swarmOptions := swarm.Options{}

	err := provisioner.Provision(swarmOptions, authOptions, engineOptions)
	if err != nil {
		return errors.Wrap(err, "Error provisioning boot2docker machine")
	}

	return nil
}

// StopHost stops the host VM.
func StopHost(api libmachine.API) error {
	host, err := api.Load(constants.MachineName)
	if err != nil {
		return errors.Wrapf(err, "Error loading host: %s", constants.MachineName)
	}
	if err := host.Stop(); err != nil {
		return errors.Wrapf(err, "Error stopping host: %s", constants.MachineName)
	}
	return nil
}

// DeleteHost deletes the host VM.
func DeleteHost(api libmachine.API) error {
	host, err := api.Load(constants.MachineName)
	if err != nil {
		return errors.Wrapf(err, "Error deleting host: %s", constants.MachineName)
	}
	m := util.MultiError{}
	m.Collect(host.Driver.Remove())
	m.Collect(api.Remove(constants.MachineName))
	return m.ToError()
}

// GetHostStatus gets the status of the host VM.
func GetHostStatus(api libmachine.API) (string, error) {
	dne := "Does Not Exist"
	exists, err := api.Exists(constants.MachineName)
	if err != nil {
		return "", errors.Wrapf(err, "Error checking that api exists for: ", constants.MachineName)
	}
	if !exists {
		return dne, nil
	}

	host, err := api.Load(constants.MachineName)
	if err != nil {
		return "", errors.Wrapf(err, "Error loading api for: ", constants.MachineName)
	}

	s, err := host.Driver.GetState()
	if s.String() == "" {
		return dne, nil
	}
	if err != nil {
		return "", errors.Wrap(err, "Error getting host state")
	}
	return s.String(), nil
}

// GetLocalkubeStatus gets the status of localkube from the host VM.
func GetLocalkubeStatus(api libmachine.API) (string, error) {
	host, err := CheckIfApiExistsAndLoad(api)
	if err != nil {
		return "", err
	}
	s, err := host.RunSSHCommand(localkubeStatusCommand)
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

type sshAble interface {
	RunSSHCommand(string) (string, error)
}

// MachineConfig contains the parameters used to start a cluster.
type MachineConfig struct {
	MinikubeISO         string
	Memory              int
	CPUs                int
	DiskSize            int
	VMDriver            string
	DockerEnv           []string // Each entry is formatted as KEY=VALUE.
	InsecureRegistry    []string
	RegistryMirror      []string
	HostOnlyCIDR        string // Only used by the virtualbox driver
	HypervVirtualSwitch string
	KvmNetwork          string // Only used by the KVM driver
}

// KubernetesConfig contains the parameters used to configure the VM Kubernetes.
type KubernetesConfig struct {
	KubernetesVersion string
	NodeIP            string
	ContainerRuntime  string
	NetworkPlugin     string
	ExtraOptions      util.ExtraOptionSlice
}

// StartCluster starts a k8s cluster on the specified Host.
func StartCluster(driver drivers.Driver, kubernetesConfig KubernetesConfig) error {
	startCommand, err := GetStartCommand(kubernetesConfig)
	if err != nil {
		return errors.Wrapf(err, "Error generating start command: %s", err)
	}
	commands := []string{stopCommand, startCommand}

	for _, cmd := range commands {
		glog.Infoln(cmd)
		output, err := drivers.RunSSHCommandFromDriver(driver, cmd)
		glog.Infoln(output)
		if err != nil {
			return errors.Wrapf(err, "Error running ssh command: %s", cmd)
		}
	}

	return nil
}

func UpdateCluster(d drivers.Driver, config KubernetesConfig) error {
	client, err := sshutil.NewSSHClient(d)
	if err != nil {
		return errors.Wrap(err, "Error creating new ssh client")
	}

	// transfer localkube from cache/asset to vm
	if localkubeURIWasSpecified(config) {
		lCacher := localkubeCacher{config}
		if err = lCacher.updateLocalkubeFromURI(client); err != nil {
			return errors.Wrap(err, "Error updating localkube from uri")
		}
	} else {
		if err = updateLocalkubeFromAsset(client); err != nil {
			return errors.Wrap(err, "Error updating localkube from asset")
		}
	}
	fileAssets := []assets.CopyableFile{}
	assets.AddMinikubeAddonsDirToAssets(&fileAssets)
	// merge files to copy
	var copyableFiles []assets.CopyableFile
	for _, addonBundle := range assets.Addons {
		if isEnabled, err := addonBundle.IsEnabled(); err == nil && isEnabled {
			for _, addon := range addonBundle.Assets {
				copyableFiles = append(copyableFiles, addon)
			}
		} else if err != nil {
			return err
		}
	}
	copyableFiles = append(copyableFiles, fileAssets...)
	// transfer files to vm
	for _, copyableFile := range copyableFiles {
		if err := sshutil.TransferFile(copyableFile, client); err != nil {
			return err
		}
	}
	return nil
}

func localkubeURIWasSpecified(config KubernetesConfig) bool {
	// see if flag is different than default -> it was passed by user
	return config.KubernetesVersion != constants.DefaultKubernetesVersion
}

// SetupCerts gets the generated credentials required to talk to the APIServer.
func SetupCerts(d drivers.Driver) error {
	localPath := constants.Minipath
	ipStr, err := d.GetIP()
	if err != nil {
		return errors.Wrap(err, "Error getting ip from driver")
	}
	glog.Infoln("Setting up certificates for IP: %s", ipStr)

	ip := net.ParseIP(ipStr)
	caCert := filepath.Join(localPath, "ca.crt")
	caKey := filepath.Join(localPath, "ca.key")
	publicPath := filepath.Join(localPath, "apiserver.crt")
	privatePath := filepath.Join(localPath, "apiserver.key")
	if err := GenerateCerts(caCert, caKey, publicPath, privatePath, ip); err != nil {
		return errors.Wrap(err, "Error generating certs")
	}

	client, err := sshutil.NewSSHClient(d)
	if err != nil {
		return errors.Wrap(err, "Error creating new ssh client")
	}

	for _, cert := range certs {
		p := filepath.Join(localPath, cert)
		data, err := ioutil.ReadFile(p)
		if err != nil {
			return errors.Wrapf(err, "Error reading file: %s", p)
		}
		perms := "0644"
		if strings.HasSuffix(cert, ".key") {
			perms = "0600"
		}
		if err := sshutil.Transfer(bytes.NewReader(data), len(data), util.DefaultCertPath, cert, perms, client); err != nil {
			return errors.Wrapf(err, "Error transferring data: %s", string(data))
		}
	}
	return nil
}

func createVirtualboxHost(config MachineConfig) drivers.Driver {
	d := virtualbox.NewDriver(constants.MachineName, constants.Minipath)
	d.Boot2DockerURL = config.GetISOFileURI()
	d.Memory = config.Memory
	d.CPU = config.CPUs
	d.DiskSize = int(config.DiskSize)
	d.HostOnlyCIDR = config.HostOnlyCIDR
	return d
}

func (m *MachineConfig) CacheMinikubeISOFromURL() error {
	options := download.FileOptions{
		Mkdirs: download.MkdirAll,
		Options: download.Options{
			ProgressBars: &download.ProgressBarOptions{
				MaxWidth: 80,
			},
		},
	}

	// Validate the ISO if it was the default URL, before writing it to disk.
	if m.MinikubeISO == constants.DefaultIsoUrl {
		options.Checksum = constants.DefaultIsoShaUrl
		options.ChecksumHash = crypto.SHA256
	}

	fmt.Println("Downloading Minikube ISO")
	if err := download.ToFile(m.MinikubeISO, m.GetISOCacheFilepath(), options); err != nil {
		return errors.Wrap(err, "Error downloading Minikube ISO")
	}

	return nil
}

func (m *MachineConfig) ShouldCacheMinikubeISO() bool {
	// store the miniube-iso inside the .minikube dir

	urlObj, err := url.Parse(m.MinikubeISO)
	if err != nil {
		return false
	}
	if urlObj.Scheme == fileScheme {
		return false
	}
	if m.IsMinikubeISOCached() {
		return false
	}
	return true
}

func (m *MachineConfig) GetISOCacheFilepath() string {
	return filepath.Join(constants.Minipath, "cache", "iso", filepath.Base(m.MinikubeISO))
}

func (m *MachineConfig) GetISOFileURI() string {
	urlObj, err := url.Parse(m.MinikubeISO)
	if err != nil {
		return m.MinikubeISO
	}
	if urlObj.Scheme == fileScheme {
		return m.MinikubeISO
	}
	isoPath := filepath.Join(constants.Minipath, "cache", "iso", filepath.Base(m.MinikubeISO))
	// As this is a file URL there should be no backslashes regardless of platform running on.
	return "file://" + filepath.ToSlash(isoPath)
}

func (m *MachineConfig) IsMinikubeISOCached() bool {
	if _, err := os.Stat(m.GetISOCacheFilepath()); os.IsNotExist(err) {
		return false
	}
	return true
}

// TODO read this from config file
func Load(config MachineConfig) drivers.Driver {
	switch config.VMDriver {
	case "virtualbox":
		return createVirtualboxHost(config)
	case "vmwarefusion":
		return createVMwareFusionHost(config)
	case "kvm":
		return createKVMHost(config)
	case "xhyve":
		return createXhyveHost(config)
	case "hyperv":
		return createHypervHost(config)
	default:
		glog.Exitf("Unsupported driver: %s\n", config.VMDriver)
		return nil
	}
}

func createMachine(config MachineConfig) (drivers.Driver, error) {
	if config.ShouldCacheMinikubeISO() {
		if err := config.CacheMinikubeISOFromURL(); err != nil {
			return nil, errors.Wrap(err, "Error attempting to cache minikube iso from url")
		}
	}

	d := Load(config)
	glog.Infof("Using machine config: %+v", config)
	if err := d.PreCreateCheck(); err != nil {
		return nil, err
	}
	if err := d.Create(); err != nil {
		return d, errors.Wrap(err, "Error creating machine")
	}

	//TODO SAVE

	err := Provision(config)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// GetHostDockerEnv gets the necessary docker env variables to allow the use of docker through minikube's vm
func GetHostDockerEnv(api libmachine.API) (map[string]string, error) {
	host, err := CheckIfApiExistsAndLoad(api)
	if err != nil {
		return nil, errors.Wrap(err, "Error checking that api exists and loading it")
	}
	ip, err := host.Driver.GetIP()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting ip from host")
	}

	tcpPrefix := "tcp://"
	port := "2376"

	envMap := map[string]string{
		"DOCKER_TLS_VERIFY": "1",
		"DOCKER_HOST":       tcpPrefix + net.JoinHostPort(ip, port),
		"DOCKER_CERT_PATH":  constants.MakeMiniPath("certs"),
	}
	return envMap, nil
}

// GetHostLogs gets the localkube logs of the host VM.
func GetHostLogs(api libmachine.API) (string, error) {
	host, err := CheckIfApiExistsAndLoad(api)
	if err != nil {
		return "", errors.Wrap(err, "Error checking that api exists and loading it")
	}
	s, err := host.RunSSHCommand(logsCommand)
	if err != nil {
		return "", err
	}
	return s, nil
}

func CheckIfApiExistsAndLoad(api libmachine.API) (*host.Host, error) {
	exists, err := api.Exists(constants.MachineName)
	if err != nil {
		return nil, errors.Wrapf(err, "Error checking that api exists for: ", constants.MachineName)
	}
	if !exists {
		return nil, errors.Errorf("Machine does not exist for api.Exists(%s)", constants.MachineName)
	}

	host, err := api.Load(constants.MachineName)
	if err != nil {
		return nil, errors.Wrapf(err, "Error loading api for: ", constants.MachineName)
	}
	return host, nil
}

func CreateSSHShell(api libmachine.API, args []string) error {
	host, err := CheckIfApiExistsAndLoad(api)
	if err != nil {
		return errors.Wrap(err, "Error checking if api exist and loading it")
	}

	currentState, err := host.Driver.GetState()
	if err != nil {
		return errors.Wrap(err, "Error getting state of host")
	}

	if currentState != state.Running {
		return errors.Errorf("Error: Cannot run ssh command: Host %q is not running", constants.MachineName)
	}

	client, err := host.CreateSSHClient()
	if err != nil {
		return errors.Wrap(err, "Error creating ssh client")
	}
	return client.Shell(strings.Join(args, " "))
}

type ipPort struct {
	IP   string
	Port int32
}

func GetServiceURLsForService(api libmachine.API, namespace, service string, t *template.Template) ([]string, error) {
	host, err := CheckIfApiExistsAndLoad(api)
	if err != nil {
		return nil, errors.Wrap(err, "Error checking if api exist and loading it")
	}

	ip, err := host.Driver.GetIP()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting ip from host")
	}

	client, err := GetKubernetesClient()
	if err != nil {
		return nil, err
	}

	return getServiceURLsWithClient(client, ip, namespace, service, t)
}

func getServiceURLsWithClient(client *unversioned.Client, ip, namespace, service string, t *template.Template) ([]string, error) {
	if t == nil {
		return nil, errors.New("Error, attempted to generate service url with nil --format template")
	}

	ports, err := getServicePorts(client, namespace, service)
	if err != nil {
		return nil, err
	}
	urls := []string{}
	for _, port := range ports {

		var doc bytes.Buffer
		err = t.Execute(&doc, ipPort{ip, port})
		if err != nil {
			return nil, err
		}

		u, err := url.Parse(doc.String())
		if err != nil {
			return nil, err
		}

		urls = append(urls, u.String())
	}
	return urls, nil
}

type serviceGetter interface {
	Get(name string) (*kubeapi.Service, error)
	List(kubeapi.ListOptions) (*kubeapi.ServiceList, error)
}

func getServicePorts(client *unversioned.Client, namespace, service string) ([]int32, error) {
	services := client.Services(namespace)
	return getServicePortsFromServiceGetter(services, service)
}

type MissingNodePortError struct {
	service *kubeapi.Service
}

func (e MissingNodePortError) Error() string {
	return fmt.Sprintf("Service %s/%s does not have a node port. To have one assigned automatically, the service type must be NodePort or LoadBalancer, but this service is of type %s.", e.service.Namespace, e.service.Name, e.service.Spec.Type)
}

func getServiceFromServiceGetter(services serviceGetter, service string) (*kubeapi.Service, error) {
	svc, err := services.Get(service)
	if err != nil {
		return nil, fmt.Errorf("Error getting %s service: %s", service, err)
	}
	return svc, nil
}

func getServicePortsFromServiceGetter(services serviceGetter, service string) ([]int32, error) {
	svc, err := getServiceFromServiceGetter(services, service)
	if err != nil {
		return nil, err
	}
	var nodePorts []int32
	if len(svc.Spec.Ports) > 0 {
		for _, port := range svc.Spec.Ports {
			if port.NodePort > 0 {
				nodePorts = append(nodePorts, port.NodePort)
			}
		}
	}
	if len(nodePorts) == 0 {
		return nil, MissingNodePortError{svc}
	}
	return nodePorts, nil
}

func GetKubernetesClient() (*unversioned.Client, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("Error creating kubeConfig: %s", err)
	}
	client, err := unversioned.New(config)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating new client from kubeConfig.ClientConfig()")
	}
	return client, nil
}

// EnsureMinikubeRunningOrExit checks that minikube has a status available and that
// that the status is `Running`, otherwise it will exit
func EnsureMinikubeRunningOrExit(api libmachine.API, exitStatus int) {
	s, err := GetHostStatus(api)
	if err != nil {
		glog.Errorln("Error getting machine status:", err)
		os.Exit(1)
	}
	if s != state.Running.String() {
		fmt.Fprintln(os.Stdout, "minikube is not currently running so the service cannot be accessed")
		os.Exit(exitStatus)
	}
}

type ServiceURL struct {
	Namespace string
	Name      string
	URLs      []string
}

type ServiceURLs []ServiceURL

func GetServiceURLs(api libmachine.API, namespace string, t *template.Template) (ServiceURLs, error) {
	host, err := CheckIfApiExistsAndLoad(api)
	if err != nil {
		return nil, err
	}

	ip, err := host.Driver.GetIP()
	if err != nil {
		return nil, err
	}

	client, err := GetKubernetesClient()
	if err != nil {
		return nil, err
	}

	getter := client.Services(namespace)

	svcs, err := getter.List(kubeapi.ListOptions{})
	if err != nil {
		return nil, err
	}

	var serviceURLs []ServiceURL

	for _, svc := range svcs.Items {
		urls, err := getServiceURLsWithClient(client, ip, svc.Namespace, svc.Name, t)
		if err != nil {
			if _, ok := err.(MissingNodePortError); ok {
				serviceURLs = append(serviceURLs, ServiceURL{Namespace: svc.Namespace, Name: svc.Name})
				continue
			}
			return nil, err
		}
		serviceURLs = append(serviceURLs, ServiceURL{Namespace: svc.Namespace, Name: svc.Name, URLs: urls})
	}

	return serviceURLs, nil
}
