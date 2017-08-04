package bootstrapper

import (
	"html/template"
	"path/filepath"

	"bytes"

	"k8s.io/minikube/pkg/util"
)

const etcdTmpl = `
apiVersion: v1
kind: Pod
metadata:
  annotations:
    scheduler.alpha.kubernetes.io/critical-pod: ""
  creationTimestamp: null
  labels:
    component: etcd
    tier: control-plane
  name: etcd
  namespace: kube-system
spec:
  containers:
  - command:
    - etcd
    - --listen-client-urls=http://127.0.0.1:2379
    - --advertise-client-urls=http://127.0.0.1:2379
    - --data-dir=/var/lib/etcd
    image: gcr.io/google_containers/etcd-amd64:3.0.17
    livenessProbe:
      failureThreshold: 8
      httpGet:
        host: 127.0.0.1
        path: /health
        port: 2379
        scheme: HTTP
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: etcd
    resources: {}
    volumeMounts:
    - mountPath: /etc/ssl/certs
      name: certs
    - mountPath: /var/lib/etcd
      name: etcd
    - mountPath: /etc/kubernetes
      name: k8s
      readOnly: true
  hostNetwork: true
  securityContext:
    seLinuxOptions:
      type: spc_t
  volumes:
  - hostPath:
      path: /etc/ssl/certs
    name: certs
  - hostPath:
      path: /var/lib/etcd
    name: etcd
  - hostPath:
      path: /etc/kubernetes
    name: k8s
status: {}
`

const apiserverTmpl = `
apiVersion: v1
kind: Pod
metadata:
  annotations:
    scheduler.alpha.kubernetes.io/critical-pod: ""
  creationTimestamp: null
  labels:
    component: kube-apiserver
    tier: control-plane
  name: kube-apiserver
  namespace: kube-system
spec:
  containers:
  - command:
    - kube-apiserver
    - --allow-privileged={{.AllowPrivileged}}
    - --client-ca-file={{.ClientCAFile}}
    - --tls-cert-file={{.TLSCertFile}}
    - --kubelet-client-certificate={{KubeletClientCert}}
    - --kubelet-client-key={{.KubeletClientKey}}
    - --admission-control={{.AdmissionControl}}
    - --tls-private-key-file={{.TLSPrivateKey}}
    - --service-cluster-ip-range={{.ServiceClusterIPRange}}
    - --secure-port={{.SecurePort}}
    - --insecure-port=0
    - --kubelet-preferred-address-types={{.KubeletPrefferedAddressTypes}}
    - --advertise-address={{.AdvertiseAddress}}
    - --etcd-servers={{.EtcdServers}}
    image: gcr.io/google_containers/kube-apiserver-amd64:v1.7.2
    livenessProbe:
      failureThreshold: 8
      httpGet:
        host: 127.0.0.1
        path: /healthz
        port: 8443
        scheme: HTTPS
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: kube-apiserver
    resources:
      requests:
        cpu: 250m
    volumeMounts:
    - mountPath: /etc/kubernetes
      name: k8s
      readOnly: true
    - mountPath: /etc/ssl/certs
      name: certs
    - mountPath: {{.CertDir}}
      name: certdir
  hostNetwork: true
  volumes:
  - hostPath:
      path: /etc/kubernetes
    name: k8s
  - hostPath:
      path: /etc/ssl/certs
    name: certs
  - hostPath:
      path: {{.CertDir}}
    name: certdir
status: {}
`

func GetAPIServerCommand(k KubernetesConfig) (bytes.Buffer, error) {
	opts := struct {
		CertDir                      string
		AllowPrivileged              bool
		ClientCAFile                 string
		TLSCertFile                  string
		KubeletClientCert            string
		KubeletClientKey             string
		AdmissionControl             string
		TLSPrivateKeyFile            string
		ServiceClusterIPRange        string
		SecurePort                   int
		KubeletPreferredAddressTypes string
		AdvertiseAddress             string
		EtcdServers                  string
	}{
		CertDir:                      util.DefaultCertPath,
		AllowPrivileged:              true,
		ClientCAFile:                 filepath.Join(util.DefaultCertPath, "ca.crt"),
		TLSCertFile:                  filepath.Join(util.DefaultCertPath, "apiserver.crt"),
		KubeletClientCert:            filepath.Join(util.DefaultCertPath, "apiserver-kubelet-client.crt"),
		KubeletClientKey:             filepath.Join(util.DefaultCertPath, "apiserver-kubelet-client.key"),
		AdmissionControl:             "Initializers,NamespaceLifecycle,LimitRanger,ServiceAccount,PersistentVolumeLabel,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,ResourceQuota",
		TLSPrivateKeyFile:            filepath.Join(util.DefaultCertPath, "apiserver.key"),
		ServiceClusterIPRange:        util.DefaultServiceCIDR,
		SecurePort:                   util.APIServerPort,
		AdvertiseAddress:             k.NodeIP,
		EtcdServers:                  "http://127.0.0.1:2379",
		KubeletPreferredAddressTypes: "InternalIP,ExternalIP,Hostname",
	}
	t := template.Must(template.New("apiserverTmpl").Parse(apiserverTmpl))

	b := bytes.Buffer{}
	if err := t.Execute(&b, opts); err != nil {
		return b, err
	}

	return b, nil
}

func GetEtcdCommand(k KubernetesConfig) (bytes.Buffer, error) {
	t := template.Must(template.New("etcdTmpl").Parse(etcdTmpl))
	b := bytes.Buffer{}
	if err := t.Execute(&b, opts); err != nil {
		return b, err
	}

	return b, nil
}
