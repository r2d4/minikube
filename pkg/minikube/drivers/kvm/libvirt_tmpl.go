package kvm

import (
	"bytes"
	"text/template"
)

var domainTmpl = `
<domain type='kvm'>
  <name>{{.MachineName}}</name> 
  <memory unit='M'>{{.Memory}}</memory>
  <vcpu>{{.CPUs}}</vcpu>
  <os>
    <type>hvm</type>
    <boot dev='cdrom'/>
    <boot dev='hd'/>
    <bootmenu enable='no'/>
  </os>
  <devices>
    <disk type='file' device='cdrom'>
      <source file='{{.IsoPath}}'/>
      <target dev='hdc' bus='ide'/>
      <readonly/>
    </disk>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2' />
      <source file='{{.DiskPath}}'/>
      <target dev='hda' bus='ide'/>
    </disk>
    <interface type='network'>
      <source network='default'/>
    </interface>
  </devices>
</domain>
`

var storageVolTmpl = `
  <volume>
    <name>{{.DiskPath}}</name>
    <capacity unit="MiB">{{.DiskSize}}</capacity>
    <format type="qcow2"/>
  </volume>
`

var storagePoolTmpl = `
<pool type='dir'>
  <name>{{.MachineName}}-pool</name>
  <target>
    <path>
      /home/matt/.minikube
    </path>
  </target>
</pool>
 `

func (d *Driver) getDomainXML() (string, error) {
	tmpl := template.Must(template.New("domain").Parse(domainTmpl))

	var domainXML bytes.Buffer
	err := tmpl.Execute(&domainXML, d)

	return domainXML.String(), err
}

func (d *Driver) getStorageVolXML() (string, error) {
	tmpl := template.Must(template.New("storage_vol").Parse(storageVolTmpl))

	var storageVolTmpl bytes.Buffer
	err := tmpl.Execute(&storageVolTmpl, d)

	return storageVolTmpl.String(), err
}

func (d *Driver) getStoragePoolXML() (string, error) {
	tmpl := template.Must(template.New("storage_pool").Parse(storagePoolTmpl))

	var storagePoolTmpl bytes.Buffer
	err := tmpl.Execute(&storagePoolTmpl, d)

	return storagePoolTmpl.String(), err
}
