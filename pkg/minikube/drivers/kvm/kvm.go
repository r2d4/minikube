package kvm

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io/ioutil"

	"k8s.io/minikube/pkg/minikube/constants"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/ssh"
	"github.com/docker/machine/libmachine/state"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/rgbkrk/libvirt-go"
)

const (
	qemuSystem = "qemu:///system"

	defaultCPU         = 2
	defaultMemory      = 2048
	defaultIsoPath     = ""
	defaultDiskSize    = 20000
	defaultRuntimePort = 2379
)

type Driver struct {
	*drivers.BaseDriver
	CPU         int
	Memory      int
	DiskSize    int
	IsoUrl      string
	DiskPath    string
	MachineName string
}

// Maps libvirt virDomainState to libmachine state
// See
// https://github.com/rgbkrk/libvirt-go/blob/v2.13.0/constants.go#L15
var stateMap = map[int]state.State{
	libvirt.VIR_DOMAIN_NOSTATE:     state.None,
	libvirt.VIR_DOMAIN_RUNNING:     state.Running,
	libvirt.VIR_DOMAIN_BLOCKED:     state.Error, // the domain is blocked on resources
	libvirt.VIR_DOMAIN_PAUSED:      state.Paused,
	libvirt.VIR_DOMAIN_SHUTDOWN:    state.Stopped,
	libvirt.VIR_DOMAIN_CRASHED:     state.Error,
	libvirt.VIR_DOMAIN_PMSUSPENDED: state.Saved,
	libvirt.VIR_DOMAIN_SHUTOFF:     state.Stopped,
}

// New Driver creates a new KVM driver with default settings
func NewDriver() *Driver {
	return &Driver{}
}

// Creates a new KVM virtual machine
// with a qcow2 disk image
// using the driver config
func (d *Driver) Create() error {
	glog.Infoln("Creating KVM VM...")

	//TODO:r2d4 Remove last libmachine/ssh dependency
	glog.Infoln("Creating SSH Key Pair...")
	if err := ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
		return errors.Wrap(err, "Error generating SSH Key Pair")
	}

	glog.Infof("(libvirt) Attempting to establish connection to %s", qemuSystem)
	conn, err := libvirt.NewVirConnection(qemuSystem)
	if err != nil {
		return errors.Wrap(err, "Error connecting to libvirt")
	}
	defer conn.CloseConnection()

	glog.Infoln("Creating qcow2 Disk Image...")
	_, err = d.generateQcow2Image(conn)
	if err != nil {
		return errors.Wrap(err, "Error generating qcow2 image")
	}

	xml, err := d.getDomainXML()
	if err != nil {
		return err
	}

	kvm, err := conn.DomainDefineXML(xml)
	if err != nil {
		return errors.Wrap(err, "Error defining domain")
	}

	err = kvm.Create()
	if err != nil {
		return errors.Wrap(err, "Error creating domain")
	}

	//TODO:r2d4 start the VM here
	return nil
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "kvm-minikube"
}

// TODO not implemented
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	var r []mcnflag.Flag
	return r
}

// GetIP returns an IP or hostname that this host is available at
// TODO not implemented
func (d *Driver) GetIP() (string, error) {
	return "", nil
}

// GetMachineName returns the name of the machine
func (d *Driver) GetMachineName() string {
	return d.MachineName
}

// GetSSHHostname returns the hostname to use for SSH
func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

// GetSSHKeyPath returns key path for use with ssh
func (d *Driver) GetSSHKeyPath() string {
	return constants.MakeMiniPath("certs", "ssh")
}

// GetSSHPort returns port for use with ssh
func (d *Driver) GetSSHPort() (int, error) {
	return 22, nil
}

// GetSSHUsername
func (d *Driver) GetSSHUsername() string {
	if d.SSHUser == "" {
		d.SSHUser = "docker"
	}

	return d.SSHUser
}

// GetURL returns a docker compatible host URL for connecting the the host
func (d *Driver) GetURL() (string, error) {
	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}
	if ip == "" {
		return "", nil
	}
	return fmt.Sprintf("tcp://%s:2379", ip), nil
}

// GetState returns the state that the host is in (running, stopped, etc.)
func (d *Driver) GetState() (state.State, error) {
	conn, err := libvirt.NewVirConnection(qemuSystem)
	if err != nil {
		return state.None, err
	}
	defer conn.CloseConnection()
	domain, err := conn.LookupDomainByName(d.MachineName)
	if err != nil {
		return state.None, errors.Wrap(err, "Error when checking state")
	}
	machineState, err := domain.GetState()
	if err != nil {
		return state.None, err
	}
	glog.Infof("GetState: State: %d, Reason: %d", machineState[0], machineState[1])
	if val, ok := stateMap[machineState[0]]; ok {
		return val, nil
	}

	return state.None, errors.Wrapf(err, "Libvirt returned invalid state %d, reason: %d", machineState[0], machineState[1])
}

// Kill stops a host forcefully
func (d *Driver) Kill() error {
	return nil
}

// PreCreateCheck should check things like
// network, kvm installed, permissions correct, etc.
func (d *Driver) PreCreateCheck() error {
	return nil
}

// Remove deletes a host
func (d *Driver) Remove() error {
	return nil
}

// Restart a host
func (d *Driver) Restart() error {
	return nil
}

// not implementing
func (d *Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {
	//no-op, not implementing
	return nil
}

// Start a host
func (d *Driver) Start() error {
	return nil
}

// Stop
func (d *Driver) Stop() error {
	return nil
}

func (d *Driver) pubKeyPath() string {
	return d.GetSSHKeyPath() + ".pub"
}

func (d *Driver) generateQcow2Image(conn libvirt.VirConnection) (*libvirt.VirStorageVol, error) {

	xml, err := d.getStorageVolXML()
	if err != nil {
		return nil, err
	}
	pool, err := conn.StoragePoolDefineXML(xml, 0)
	if err != nil {
		return nil, err
	}

	err = pool.Create(0)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating storage pool")
	}

	xml, err = d.getStoragePoolXML()
	if err != nil {
		return nil, err
	}

	vol, err := pool.StorageVolCreateXML(xml, 0)
	//TODO:r2d4 nil deference fix this
	return &vol, err
}

// generateUserdataBundle generates a tar that includes
// a magic incantation that tells boot2docker to format the disk
// it also includes the ssh keys
// see https://github.com/kubernetes/minikube/blob/master/deploy/iso/minikube-iso/package/automount/minikube-automount
func (d *Driver) generateUserdataBundle() (*bytes.Buffer, error) {
	glog.Infoln("Generating userdata bundle for b2d...")
	magicString := "boot2docker, please format-me"

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	// The magic string must go first.
	file := &tar.Header{
		Name: magicString,
		Size: int64(len(magicString)),
	}
	if err := tw.WriteHeader(file); err != nil {
		return nil, err
	}
	if _, err := tw.Write([]byte(magicString)); err != nil {
		return nil, err
	}

	// Create .ssh dir
	file = &tar.Header{
		Name:     ".ssh",
		Typeflag: tar.TypeDir,
		Mode:     0700,
	}
	if err := tw.WriteHeader(file); err != nil {
		return nil, err
	}

	pubKey, err := ioutil.ReadFile(d.pubKeyPath())
	if err != nil {
		return nil, err
	}
	file = &tar.Header{
		Name: ".ssh/authorized_keys",
		Size: int64(len(pubKey)),
		Mode: 0644,
	}
	if err := tw.WriteHeader(file); err != nil {
		return nil, err
	}
	if _, err := tw.Write([]byte(pubKey)); err != nil {
		return nil, err
	}
	file = &tar.Header{
		Name: ".ssh/authorizedkeys2",
		Size: int64(len(pubKey)),
		Mode: 0644,
	}

	if err := tw.WriteHeader(file); err != nil {
		return nil, err
	}
	if _, err := tw.Write([]byte(pubKey)); err != nil {
		return nil, err
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf, nil
}
