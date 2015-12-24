/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

// Package app does all of the work necessary to create a Kubernetes
// APIServer by binding together the API, master and APIServer infrastructure.
// It can be configured and called directly or via the hyperkube framework.
package app

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	systemd "github.com/coreos/go-systemd/daemon"
	"k8s.io/kubernetes/pkg/admission"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/latest"
	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/api/unversioned"
	apiutil "k8s.io/kubernetes/pkg/api/util"
	"k8s.io/kubernetes/pkg/api/validation"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apiserver"
	"k8s.io/kubernetes/pkg/apiserver/authenticator"
	"k8s.io/kubernetes/pkg/capabilities"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/cloudprovider"
	serviceaccountcontroller "k8s.io/kubernetes/pkg/controller/serviceaccount"
	"k8s.io/kubernetes/pkg/genericapiserver"
	kubeletclient "k8s.io/kubernetes/pkg/kubelet/client"
	"k8s.io/kubernetes/pkg/master"
	"k8s.io/kubernetes/pkg/master/ports"
	"k8s.io/kubernetes/pkg/serviceaccount"
	"k8s.io/kubernetes/pkg/storage"
	etcdstorage "k8s.io/kubernetes/pkg/storage/etcd"
	"k8s.io/kubernetes/pkg/util"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	// Maximum duration before timing out read/write requests
	// Set to a value larger than the timeouts in each watch server.
	ReadWriteTimeout = time.Minute * 60
	//TODO: This can be tightened up. It still matches objects named watch or proxy.
	defaultLongRunningRequestRE = "(/|^)((watch|proxy)(/|$)|(logs?|portforward|exec|attach)/?$)"
)

// APIServer runs a kubernetes api server.
type APIServer struct {
	InsecureBindAddress        net.IP
	InsecurePort               int
	BindAddress                net.IP
	AdvertiseAddress           net.IP
	SecurePort                 int
	ExternalHost               string
	TLSCertFile                string
	TLSPrivateKeyFile          string
	CertDirectory              string
	APIPrefix                  string
	APIGroupPrefix             string
	DeprecatedStorageVersion   string
	StorageVersions            string
	CloudProvider              string
	CloudConfigFile            string
	EventTTL                   time.Duration
	BasicAuthFile              string
	ClientCAFile               string
	TokenAuthFile              string
	OIDCIssuerURL              string
	OIDCClientID               string
	OIDCCAFile                 string
	OIDCUsernameClaim          string
	ServiceAccountKeyFile      string
	ServiceAccountLookup       bool
	KeystoneURL                string
	AuthorizationMode          string
	AuthorizationPolicyFile    string
	AdmissionControl           string
	AdmissionControlConfigFile string
	EtcdServerList             []string
	EtcdServersOverrides       []string
	EtcdPathPrefix             string
	CorsAllowedOriginList      []string
	AllowPrivileged            bool
	ServiceClusterIPRange      net.IPNet // TODO: make this a list
	ServiceNodePortRange       util.PortRange
	EnableLogsSupport          bool
	MasterServiceNamespace     string
	MasterCount                int
	RuntimeConfig              util.ConfigurationMap
	KubeletConfig              kubeletclient.KubeletClientConfig
	EnableProfiling            bool
	EnableWatchCache           bool
	MaxRequestsInFlight        int
	MinRequestTimeout          int
	LongRunningRequestRE       string
	SSHUser                    string
	SSHKeyfile                 string
	MaxConnectionBytesPerSec   int64
	KubernetesServiceNodePort  int
}

// NewAPIServer creates a new APIServer object with default parameters
func NewAPIServer() *APIServer {
	s := APIServer{
		InsecurePort:           8080,
		InsecureBindAddress:    net.ParseIP("127.0.0.1"),
		BindAddress:            net.ParseIP("0.0.0.0"),
		SecurePort:             6443,
		APIPrefix:              "/api",
		APIGroupPrefix:         "/apis",
		EventTTL:               1 * time.Hour,
		AuthorizationMode:      "AlwaysAllow",
		AdmissionControl:       "AlwaysAdmit",
		EtcdPathPrefix:         genericapiserver.DefaultEtcdPathPrefix,
		EnableLogsSupport:      true,
		MasterServiceNamespace: api.NamespaceDefault,
		MasterCount:            1,
		CertDirectory:          "/var/run/kubernetes",
		StorageVersions:        latest.AllPreferredGroupVersions(),

		RuntimeConfig: make(util.ConfigurationMap),
		KubeletConfig: kubeletclient.KubeletClientConfig{
			Port:        ports.KubeletPort,
			EnableHttps: true,
			HTTPTimeout: time.Duration(5) * time.Second,
		},
	}

	return &s
}

// NewAPIServerCommand creates a *cobra.Command object with default parameters
func NewAPIServerCommand() *cobra.Command {
	s := NewAPIServer()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use: "kube-apiserver",
		Long: `The Kubernetes API server validates and configures data
for the api objects which include pods, services, replicationcontrollers, and
others. The API Server services REST operations and provides the frontend to the
cluster's shared state through which all other components interact.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet
func (s *APIServer) AddFlags(fs *pflag.FlagSet) {
	// Note: the weird ""+ in below lines seems to be the only way to get gofmt to
	// arrange these text blocks sensibly. Grrr.
	fs.IntVar(&s.InsecurePort, "insecure-port", s.InsecurePort, ""+
		"The port on which to serve unsecured, unauthenticated access. Default 8080. It is assumed "+
		"that firewall rules are set up such that this port is not reachable from outside of "+
		"the cluster and that port 443 on the cluster's public address is proxied to this "+
		"port. This is performed by nginx in the default setup.")
	fs.IntVar(&s.InsecurePort, "port", s.InsecurePort, "DEPRECATED: see --insecure-port instead")
	fs.MarkDeprecated("port", "see --insecure-port instead")
	fs.IPVar(&s.InsecureBindAddress, "insecure-bind-address", s.InsecureBindAddress, ""+
		"The IP address on which to serve the --insecure-port (set to 0.0.0.0 for all interfaces). "+
		"Defaults to localhost.")
	fs.IPVar(&s.InsecureBindAddress, "address", s.InsecureBindAddress, "DEPRECATED: see --insecure-bind-address instead")
	fs.MarkDeprecated("address", "see --insecure-bind-address instead")
	fs.IPVar(&s.BindAddress, "bind-address", s.BindAddress, ""+
		"The IP address on which to listen for the --secure-port port. The "+
		"associated interface(s) must be reachable by the rest of the cluster, and by CLI/web "+
		"clients. If blank, all interfaces will be used (0.0.0.0).")
	fs.IPVar(&s.AdvertiseAddress, "advertise-address", s.AdvertiseAddress, ""+
		"The IP address on which to advertise the apiserver to members of the cluster. This "+
		"address must be reachable by the rest of the cluster. If blank, the --bind-address "+
		"will be used. If --bind-address is unspecified, the host's default interface will "+
		"be used.")
	fs.IPVar(&s.BindAddress, "public-address-override", s.BindAddress, "DEPRECATED: see --bind-address instead")
	fs.MarkDeprecated("public-address-override", "see --bind-address instead")
	fs.IntVar(&s.SecurePort, "secure-port", s.SecurePort, ""+
		"The port on which to serve HTTPS with authentication and authorization. If 0, "+
		"don't serve HTTPS at all.")
	fs.StringVar(&s.TLSCertFile, "tls-cert-file", s.TLSCertFile, ""+
		"File containing x509 Certificate for HTTPS.  (CA cert, if any, concatenated after server cert). "+
		"If HTTPS serving is enabled, and --tls-cert-file and --tls-private-key-file are not provided, "+
		"a self-signed certificate and key are generated for the public address and saved to /var/run/kubernetes.")
	fs.StringVar(&s.TLSPrivateKeyFile, "tls-private-key-file", s.TLSPrivateKeyFile, "File containing x509 private key matching --tls-cert-file.")
	fs.StringVar(&s.CertDirectory, "cert-dir", s.CertDirectory, "The directory where the TLS certs are located (by default /var/run/kubernetes). "+
		"If --tls-cert-file and --tls-private-key-file are provided, this flag will be ignored.")
	fs.StringVar(&s.APIPrefix, "api-prefix", s.APIPrefix, "The prefix for API requests on the server. Default '/api'.")
	fs.MarkDeprecated("api-prefix", "--api-prefix is deprecated and will be removed when the v1 API is retired.")
	fs.StringVar(&s.DeprecatedStorageVersion, "storage-version", s.DeprecatedStorageVersion, "The version to store the legacy v1 resources with. Defaults to server preferred")
	fs.MarkDeprecated("storage-version", "--storage-version is deprecated and will be removed when the v1 API is retired. See --storage-versions instead.")
	fs.StringVar(&s.StorageVersions, "storage-versions", s.StorageVersions, "The versions to store resources with. "+
		"Different groups may be stored in different versions. Specified in the format \"group1/version1,group2/version2...\". "+
		"This flag expects a complete list of storage versions of ALL groups registered in the server. "+
		"It defaults to a list of preferred versions of all registered groups, which is derived from the KUBE_API_VERSIONS environment variable.")
	fs.StringVar(&s.CloudProvider, "cloud-provider", s.CloudProvider, "The provider for cloud services.  Empty string for no provider.")
	fs.StringVar(&s.CloudConfigFile, "cloud-config", s.CloudConfigFile, "The path to the cloud provider configuration file.  Empty string for no configuration file.")
	fs.DurationVar(&s.EventTTL, "event-ttl", s.EventTTL, "Amount of time to retain events. Default 1 hour.")
	fs.StringVar(&s.BasicAuthFile, "basic-auth-file", s.BasicAuthFile, "If set, the file that will be used to admit requests to the secure port of the API server via http basic authentication.")
	fs.StringVar(&s.ClientCAFile, "client-ca-file", s.ClientCAFile, "If set, any request presenting a client certificate signed by one of the authorities in the client-ca-file is authenticated with an identity corresponding to the CommonName of the client certificate.")
	fs.StringVar(&s.TokenAuthFile, "token-auth-file", s.TokenAuthFile, "If set, the file that will be used to secure the secure port of the API server via token authentication.")
	fs.StringVar(&s.OIDCIssuerURL, "oidc-issuer-url", s.OIDCIssuerURL, "The URL of the OpenID issuer, only HTTPS scheme will be accepted. If set, it will be used to verify the OIDC JSON Web Token (JWT)")
	fs.StringVar(&s.OIDCClientID, "oidc-client-id", s.OIDCClientID, "The client ID for the OpenID Connect client, must be set if oidc-issuer-url is set")
	fs.StringVar(&s.OIDCCAFile, "oidc-ca-file", s.OIDCCAFile, "If set, the OpenID server's certificate will be verified by one of the authorities in the oidc-ca-file, otherwise the host's root CA set will be used")
	fs.StringVar(&s.OIDCUsernameClaim, "oidc-username-claim", "sub", ""+
		"The OpenID claim to use as the user name. Note that claims other than the default ('sub') is not "+
		"guaranteed to be unique and immutable. This flag is experimental, please see the authentication documentation for further details.")
	fs.StringVar(&s.ServiceAccountKeyFile, "service-account-key-file", s.ServiceAccountKeyFile, "File containing PEM-encoded x509 RSA private or public key, used to verify ServiceAccount tokens. If unspecified, --tls-private-key-file is used.")
	fs.BoolVar(&s.ServiceAccountLookup, "service-account-lookup", s.ServiceAccountLookup, "If true, validate ServiceAccount tokens exist in etcd as part of authentication.")
	fs.StringVar(&s.KeystoneURL, "experimental-keystone-url", s.KeystoneURL, "If passed, activates the keystone authentication plugin")
	fs.StringVar(&s.AuthorizationMode, "authorization-mode", s.AuthorizationMode, "Ordered list of plug-ins to do authorization on secure port. Comma-delimited list of: "+strings.Join(apiserver.AuthorizationModeChoices, ","))
	fs.StringVar(&s.AuthorizationPolicyFile, "authorization-policy-file", s.AuthorizationPolicyFile, "File with authorization policy in csv format, used with --authorization-mode=ABAC, on the secure port.")
	fs.StringVar(&s.AdmissionControl, "admission-control", s.AdmissionControl, "Ordered list of plug-ins to do admission control of resources into cluster. Comma-delimited list of: "+strings.Join(admission.GetPlugins(), ", "))
	fs.StringVar(&s.AdmissionControlConfigFile, "admission-control-config-file", s.AdmissionControlConfigFile, "File with admission control configuration.")
	fs.StringSliceVar(&s.EtcdServerList, "etcd-servers", s.EtcdServerList, "List of etcd servers to watch (http://ip:port), comma separated. Mutually exclusive with -etcd-config")
	fs.StringSliceVar(&s.EtcdServersOverrides, "etcd-servers-overrides", s.EtcdServersOverrides, "Per-resource etcd servers overrides, comma separated. The individual override format: group/resource#servers, where servers are http://ip:port, semicolon separated.")
	fs.StringVar(&s.EtcdPathPrefix, "etcd-prefix", s.EtcdPathPrefix, "The prefix for all resource paths in etcd.")
	fs.StringSliceVar(&s.CorsAllowedOriginList, "cors-allowed-origins", s.CorsAllowedOriginList, "List of allowed origins for CORS, comma separated.  An allowed origin can be a regular expression to support subdomain matching.  If this list is empty CORS will not be enabled.")
	fs.BoolVar(&s.AllowPrivileged, "allow-privileged", s.AllowPrivileged, "If true, allow privileged containers.")
	fs.IPNetVar(&s.ServiceClusterIPRange, "service-cluster-ip-range", s.ServiceClusterIPRange, "A CIDR notation IP range from which to assign service cluster IPs. This must not overlap with any IP ranges assigned to nodes for pods.")
	fs.IPNetVar(&s.ServiceClusterIPRange, "portal-net", s.ServiceClusterIPRange, "Deprecated: see --service-cluster-ip-range instead.")
	fs.MarkDeprecated("portal-net", "see --service-cluster-ip-range instead.")
	fs.Var(&s.ServiceNodePortRange, "service-node-port-range", "A port range to reserve for services with NodePort visibility.  Example: '30000-32767'.  Inclusive at both ends of the range.")
	fs.Var(&s.ServiceNodePortRange, "service-node-ports", "Deprecated: see --service-node-port-range instead.")
	fs.MarkDeprecated("service-node-ports", "see --service-node-port-range instead.")
	fs.StringVar(&s.MasterServiceNamespace, "master-service-namespace", s.MasterServiceNamespace, "The namespace from which the kubernetes master services should be injected into pods")
	fs.IntVar(&s.MasterCount, "apiserver-count", s.MasterCount, "The number of apiservers running in the cluster")
	fs.Var(&s.RuntimeConfig, "runtime-config", "A set of key=value pairs that describe runtime configuration that may be passed to apiserver. apis/<groupVersion> key can be used to turn on/off specific api versions. apis/<groupVersion>/<resource> can be used to turn on/off specific resources. api/all and api/legacy are special keys to control all and legacy api versions respectively.")
	fs.BoolVar(&s.EnableProfiling, "profiling", true, "Enable profiling via web interface host:port/debug/pprof/")
	// TODO: enable cache in integration tests.
	fs.BoolVar(&s.EnableWatchCache, "watch-cache", true, "Enable watch caching in the apiserver")
	fs.StringVar(&s.ExternalHost, "external-hostname", "", "The hostname to use when generating externalized URLs for this master (e.g. Swagger API Docs.)")
	fs.IntVar(&s.MaxRequestsInFlight, "max-requests-inflight", 400, "The maximum number of requests in flight at a given time.  When the server exceeds this, it rejects requests.  Zero for no limit.")
	fs.IntVar(&s.MinRequestTimeout, "min-request-timeout", 1800, "An optional field indicating the minimum number of seconds a handler must keep a request open before timing it out. Currently only honored by the watch request handler, which picks a randomized value above this number as the connection timeout, to spread out load.")
	fs.StringVar(&s.LongRunningRequestRE, "long-running-request-regexp", defaultLongRunningRequestRE, "A regular expression matching long running requests which should be excluded from maximum inflight request handling.")
	fs.StringVar(&s.SSHUser, "ssh-user", "", "If non-empty, use secure SSH proxy to the nodes, using this user name")
	fs.StringVar(&s.SSHKeyfile, "ssh-keyfile", "", "If non-empty, use secure SSH proxy to the nodes, using this user keyfile")
	fs.Int64Var(&s.MaxConnectionBytesPerSec, "max-connection-bytes-per-sec", 0, "If non-zero, throttle each user connection to this number of bytes/sec.  Currently only applies to long-running requests")
	// Kubelet related flags:
	fs.BoolVar(&s.KubeletConfig.EnableHttps, "kubelet-https", s.KubeletConfig.EnableHttps, "Use https for kubelet connections")
	fs.UintVar(&s.KubeletConfig.Port, "kubelet-port", s.KubeletConfig.Port, "Kubelet port")
	fs.MarkDeprecated("kubelet-port", "kubelet-port is deprecated and will be removed")
	fs.DurationVar(&s.KubeletConfig.HTTPTimeout, "kubelet-timeout", s.KubeletConfig.HTTPTimeout, "Timeout for kubelet operations")
	fs.StringVar(&s.KubeletConfig.CertFile, "kubelet-client-certificate", s.KubeletConfig.CertFile, "Path to a client cert file for TLS.")
	fs.StringVar(&s.KubeletConfig.KeyFile, "kubelet-client-key", s.KubeletConfig.KeyFile, "Path to a client key file for TLS.")
	fs.StringVar(&s.KubeletConfig.CAFile, "kubelet-certificate-authority", s.KubeletConfig.CAFile, "Path to a cert. file for the certificate authority.")
	//See #14282 for details on how to test/try this option out.  TODO remove this comment once this option is tested in CI.
	fs.IntVar(&s.KubernetesServiceNodePort, "kubernetes-service-node-port", 0, "If non-zero, the Kubernetes master service (which apiserver creates/maintains) will be of type NodePort, using this as the value of the port. If zero, the Kubernetes master service will be of type ClusterIP.")
	// TODO: delete this flag as soon as we identify and fix all clients that send malformed updates, like #14126.
	fs.BoolVar(&validation.RepairMalformedUpdates, "repair-malformed-updates", true, "If true, server will do its best to fix the update request to pass the validation, e.g., setting empty UID in update request to its existing value. This flag can be turned off after we fix all the clients that send malformed updates.")
}

// TODO: Longer term we should read this from some config store, rather than a flag.
func (s *APIServer) verifyClusterIPFlags() {
	if s.ServiceClusterIPRange.IP == nil {
		glog.Fatal("No --service-cluster-ip-range specified")
	}
	var ones, bits = s.ServiceClusterIPRange.Mask.Size()
	if bits-ones > 20 {
		glog.Fatal("Specified --service-cluster-ip-range is too large")
	}
}

type newEtcdFunc func([]string, meta.VersionInterfacesFunc, string, string) (storage.Interface, error)

func newEtcd(etcdServerList []string, interfacesFunc meta.VersionInterfacesFunc, storageGroupVersionString, pathPrefix string) (etcdStorage storage.Interface, err error) {
	if storageGroupVersionString == "" {
		return etcdStorage, fmt.Errorf("storageVersion is required to create a etcd storage")
	}
	storageVersion, err := unversioned.ParseGroupVersion(storageGroupVersionString)
	if err != nil {
		return nil, err
	}

	var storageConfig etcdstorage.EtcdConfig
	storageConfig.ServerList = etcdServerList
	storageConfig.Prefix = pathPrefix
	versionedInterface, err := interfacesFunc(storageVersion)
	if err != nil {
		return nil, err
	}
	storageConfig.Codec = versionedInterface.Codec
	return storageConfig.NewStorage()
}

// convert to a map between group and groupVersions.
func generateStorageVersionMap(legacyVersion string, storageVersions string) map[string]string {
	storageVersionMap := map[string]string{}
	if legacyVersion != "" {
		storageVersionMap[""] = legacyVersion
	}
	if storageVersions != "" {
		groupVersions := strings.Split(storageVersions, ",")
		for _, gv := range groupVersions {
			storageVersionMap[apiutil.GetGroup(gv)] = gv
		}
	}
	return storageVersionMap
}

// parse the value of --etcd-servers-overrides and update given storageDestinations.
func updateEtcdOverrides(overrides []string, storageVersions map[string]string, prefix string, storageDestinations *genericapiserver.StorageDestinations, newEtcdFn newEtcdFunc) {
	if len(overrides) == 0 {
		return
	}
	for _, override := range overrides {
		tokens := strings.Split(override, "#")
		if len(tokens) != 2 {
			glog.Errorf("invalid value of etcd server overrides: %s", override)
			continue
		}

		apiresource := strings.Split(tokens[0], "/")
		if len(apiresource) != 2 {
			glog.Errorf("invalid resource definition: %s", tokens[0])
		}
		group := apiresource[0]
		resource := apiresource[1]

		apigroup, err := latest.Group(group)
		if err != nil {
			glog.Errorf("invalid api group %s: %v", group, err)
			continue
		}
		if _, found := storageVersions[apigroup.GroupVersion.Group]; !found {
			glog.Errorf("Couldn't find the storage version for group %s", apigroup.GroupVersion.Group)
			continue
		}

		servers := strings.Split(tokens[1], ";")
		etcdOverrideStorage, err := newEtcdFn(servers, apigroup.InterfacesFor, storageVersions[apigroup.GroupVersion.Group], prefix)
		if err != nil {
			glog.Fatalf("Invalid storage version or misconfigured etcd for %s: %v", tokens[0], err)
		}

		storageDestinations.AddStorageOverride(group, resource, etcdOverrideStorage)
	}
}

// Run runs the specified APIServer.  This should never exit.
func (s *APIServer) Run(_ []string) error {
	s.verifyClusterIPFlags()

	// If advertise-address is not specified, use bind-address. If bind-address
	// is not usable (unset, 0.0.0.0, or loopback), we will use the host's default
	// interface as valid public addr for master (see: util#ValidPublicAddrForMaster)
	if s.AdvertiseAddress == nil || s.AdvertiseAddress.IsUnspecified() {
		hostIP, err := util.ValidPublicAddrForMaster(s.BindAddress)
		if err != nil {
			glog.Fatalf("Unable to find suitable network address.error='%v' . "+
				"Try to set the AdvertiseAddress directly or provide a valid BindAddress to fix this.", err)
		}
		s.AdvertiseAddress = hostIP
	}
	glog.Infof("Will report %v as public IP address.", s.AdvertiseAddress)

	if len(s.EtcdServerList) == 0 {
		glog.Fatalf("--etcd-servers must be specified")
	}

	if s.KubernetesServiceNodePort > 0 && !s.ServiceNodePortRange.Contains(s.KubernetesServiceNodePort) {
		glog.Fatalf("Kubernetes service port range %v doesn't contain %v", s.ServiceNodePortRange, (s.KubernetesServiceNodePort))
	}

	capabilities.Initialize(capabilities.Capabilities{
		AllowPrivileged: s.AllowPrivileged,
		// TODO(vmarmol): Implement support for HostNetworkSources.
		PrivilegedSources: capabilities.PrivilegedSources{
			HostNetworkSources: []string{},
			HostPIDSources:     []string{},
			HostIPCSources:     []string{},
		},
		PerConnectionBandwidthLimitBytesPerSec: s.MaxConnectionBytesPerSec,
	})

	cloud, err := cloudprovider.InitCloudProvider(s.CloudProvider, s.CloudConfigFile)
	if err != nil {
		glog.Fatalf("Cloud provider could not be initialized: %v", err)
	}

	// Setup tunneler if needed
	var tunneler master.Tunneler
	var proxyDialerFn apiserver.ProxyDialerFunc
	if len(s.SSHUser) > 0 {
		// Get ssh key distribution func, if supported
		var installSSH master.InstallSSHKey
		if cloud != nil {
			if instances, supported := cloud.Instances(); supported {
				installSSH = instances.AddSSHKeyToAllInstances
			}
		}

		// Set up the tunneler
		tunneler = master.NewSSHTunneler(s.SSHUser, s.SSHKeyfile, installSSH)

		// Use the tunneler's dialer to connect to the kubelet
		s.KubeletConfig.Dial = tunneler.Dial
		// Use the tunneler's dialer when proxying to pods, services, and nodes
		proxyDialerFn = tunneler.Dial
	}

	// Proxying to pods and services is IP-based... don't expect to be able to verify the hostname
	proxyTLSClientConfig := &tls.Config{InsecureSkipVerify: true}

	kubeletClient, err := kubeletclient.NewStaticKubeletClient(&s.KubeletConfig)
	if err != nil {
		glog.Fatalf("Failure to start kubelet client: %v", err)
	}

	apiGroupVersionOverrides, err := s.parseRuntimeConfig()
	if err != nil {
		glog.Fatalf("error in parsing runtime-config: %s", err)
	}

	clientConfig := &client.Config{
		Host: net.JoinHostPort(s.InsecureBindAddress.String(), strconv.Itoa(s.InsecurePort)),
	}
	if len(s.DeprecatedStorageVersion) != 0 {
		gv, err := unversioned.ParseGroupVersion(s.DeprecatedStorageVersion)
		if err != nil {
			glog.Fatalf("error in parsing group version: %s", err)
		}
		clientConfig.GroupVersion = &gv
	}

	client, err := client.New(clientConfig)
	if err != nil {
		glog.Fatalf("Invalid server address: %v", err)
	}

	legacyV1Group, err := latest.Group(api.GroupName)
	if err != nil {
		return err
	}

	storageDestinations := genericapiserver.NewStorageDestinations()

	storageVersions := generateStorageVersionMap(s.DeprecatedStorageVersion, s.StorageVersions)
	if _, found := storageVersions[legacyV1Group.GroupVersion.Group]; !found {
		glog.Fatalf("Couldn't find the storage version for group: %q in storageVersions: %v", legacyV1Group.GroupVersion.Group, storageVersions)
	}
	etcdStorage, err := newEtcd(s.EtcdServerList, legacyV1Group.InterfacesFor, storageVersions[legacyV1Group.GroupVersion.Group], s.EtcdPathPrefix)
	if err != nil {
		glog.Fatalf("Invalid storage version or misconfigured etcd: %v", err)
	}
	storageDestinations.AddAPIGroup("", etcdStorage)

	if !apiGroupVersionOverrides["extensions/v1beta1"].Disable {
		expGroup, err := latest.Group(extensions.GroupName)
		if err != nil {
			glog.Fatalf("Extensions API is enabled in runtime config, but not enabled in the environment variable KUBE_API_VERSIONS. Error: %v", err)
		}
		if _, found := storageVersions[expGroup.GroupVersion.Group]; !found {
			glog.Fatalf("Couldn't find the storage version for group: %q in storageVersions: %v", expGroup.GroupVersion.Group, storageVersions)
		}
		expEtcdStorage, err := newEtcd(s.EtcdServerList, expGroup.InterfacesFor, storageVersions[expGroup.GroupVersion.Group], s.EtcdPathPrefix)
		if err != nil {
			glog.Fatalf("Invalid extensions storage version or misconfigured etcd: %v", err)
		}
		storageDestinations.AddAPIGroup(extensions.GroupName, expEtcdStorage)
	}

	updateEtcdOverrides(s.EtcdServersOverrides, storageVersions, s.EtcdPathPrefix, &storageDestinations, newEtcd)

	n := s.ServiceClusterIPRange

	// Default to the private server key for service account token signing
	if s.ServiceAccountKeyFile == "" && s.TLSPrivateKeyFile != "" {
		if authenticator.IsValidServiceAccountKeyFile(s.TLSPrivateKeyFile) {
			s.ServiceAccountKeyFile = s.TLSPrivateKeyFile
		} else {
			glog.Warning("No RSA key provided, service account token authentication disabled")
		}
	}

	var serviceAccountGetter serviceaccount.ServiceAccountTokenGetter
	if s.ServiceAccountLookup {
		// If we need to look up service accounts and tokens,
		// go directly to etcd to avoid recursive auth insanity
		serviceAccountGetter = serviceaccountcontroller.NewGetterFromStorageInterface(etcdStorage)
	}

	authenticator, err := authenticator.New(authenticator.AuthenticatorConfig{
		BasicAuthFile:             s.BasicAuthFile,
		ClientCAFile:              s.ClientCAFile,
		TokenAuthFile:             s.TokenAuthFile,
		OIDCIssuerURL:             s.OIDCIssuerURL,
		OIDCClientID:              s.OIDCClientID,
		OIDCCAFile:                s.OIDCCAFile,
		OIDCUsernameClaim:         s.OIDCUsernameClaim,
		ServiceAccountKeyFile:     s.ServiceAccountKeyFile,
		ServiceAccountLookup:      s.ServiceAccountLookup,
		ServiceAccountTokenGetter: serviceAccountGetter,
		KeystoneURL:               s.KeystoneURL,
	})

	if err != nil {
		glog.Fatalf("Invalid Authentication Config: %v", err)
	}

	authorizationModeNames := strings.Split(s.AuthorizationMode, ",")
	authorizer, err := apiserver.NewAuthorizerFromAuthorizationConfig(authorizationModeNames, s.AuthorizationPolicyFile)
	if err != nil {
		glog.Fatalf("Invalid Authorization Config: %v", err)
	}

	admissionControlPluginNames := strings.Split(s.AdmissionControl, ",")
	admissionController := admission.NewFromPlugins(client, admissionControlPluginNames, s.AdmissionControlConfigFile)

	if len(s.ExternalHost) == 0 {
		// TODO: extend for other providers
		if s.CloudProvider == "gce" {
			instances, supported := cloud.Instances()
			if !supported {
				glog.Fatalf("GCE cloud provider has no instances.  this shouldn't happen. exiting.")
			}
			name, err := os.Hostname()
			if err != nil {
				glog.Fatalf("Failed to get hostname: %v", err)
			}
			addrs, err := instances.NodeAddresses(name)
			if err != nil {
				glog.Warningf("Unable to obtain external host address from cloud provider: %v", err)
			} else {
				for _, addr := range addrs {
					if addr.Type == api.NodeExternalIP {
						s.ExternalHost = addr.Address
					}
				}
			}
		}
	}

	config := &master.Config{
		Config: &genericapiserver.Config{
			StorageDestinations:       storageDestinations,
			StorageVersions:           storageVersions,
			ServiceClusterIPRange:     &n,
			EnableLogsSupport:         s.EnableLogsSupport,
			EnableUISupport:           true,
			EnableSwaggerSupport:      true,
			EnableProfiling:           s.EnableProfiling,
			EnableWatchCache:          s.EnableWatchCache,
			EnableIndex:               true,
			APIPrefix:                 s.APIPrefix,
			APIGroupPrefix:            s.APIGroupPrefix,
			CorsAllowedOriginList:     s.CorsAllowedOriginList,
			ReadWritePort:             s.SecurePort,
			PublicAddress:             s.AdvertiseAddress,
			Authenticator:             authenticator,
			SupportsBasicAuth:         len(s.BasicAuthFile) > 0,
			Authorizer:                authorizer,
			AdmissionControl:          admissionController,
			APIGroupVersionOverrides:  apiGroupVersionOverrides,
			MasterServiceNamespace:    s.MasterServiceNamespace,
			MasterCount:               s.MasterCount,
			ExternalHost:              s.ExternalHost,
			MinRequestTimeout:         s.MinRequestTimeout,
			ProxyDialer:               proxyDialerFn,
			ProxyTLSClientConfig:      proxyTLSClientConfig,
			ServiceNodePortRange:      s.ServiceNodePortRange,
			KubernetesServiceNodePort: s.KubernetesServiceNodePort,
		},
		EnableCoreControllers: true,
		EventTTL:              s.EventTTL,
		KubeletClient:         kubeletClient,

		Tunneler: tunneler,
	}
	m := master.New(config)

	// We serve on 2 ports.  See docs/accessing_the_api.md
	secureLocation := ""
	if s.SecurePort != 0 {
		secureLocation = net.JoinHostPort(s.BindAddress.String(), strconv.Itoa(s.SecurePort))
	}
	insecureLocation := net.JoinHostPort(s.InsecureBindAddress.String(), strconv.Itoa(s.InsecurePort))

	// See the flag commentary to understand our assumptions when opening the read-only and read-write ports.

	var sem chan bool
	if s.MaxRequestsInFlight > 0 {
		sem = make(chan bool, s.MaxRequestsInFlight)
	}

	longRunningRE := regexp.MustCompile(s.LongRunningRequestRE)
	longRunningTimeout := func(req *http.Request) (<-chan time.Time, string) {
		// TODO unify this with apiserver.MaxInFlightLimit
		if longRunningRE.MatchString(req.URL.Path) || req.URL.Query().Get("watch") == "true" {
			return nil, ""
		}
		return time.After(time.Minute), ""
	}

	if secureLocation != "" {
		handler := apiserver.TimeoutHandler(m.Handler, longRunningTimeout)
		secureServer := &http.Server{
			Addr:           secureLocation,
			Handler:        apiserver.MaxInFlightLimit(sem, longRunningRE, apiserver.RecoverPanics(handler)),
			MaxHeaderBytes: 1 << 20,
			TLSConfig: &tls.Config{
				// Change default from SSLv3 to TLSv1.0 (because of POODLE vulnerability)
				MinVersion: tls.VersionTLS10,
			},
		}

		if len(s.ClientCAFile) > 0 {
			clientCAs, err := util.CertPoolFromFile(s.ClientCAFile)
			if err != nil {
				glog.Fatalf("Unable to load client CA file: %v", err)
			}
			// Populate PeerCertificates in requests, but don't reject connections without certificates
			// This allows certificates to be validated by authenticators, while still allowing other auth types
			secureServer.TLSConfig.ClientAuth = tls.RequestClientCert
			// Specify allowed CAs for client certificates
			secureServer.TLSConfig.ClientCAs = clientCAs
		}

		glog.Infof("Serving securely on %s", secureLocation)
		if s.TLSCertFile == "" && s.TLSPrivateKeyFile == "" {
			s.TLSCertFile = path.Join(s.CertDirectory, "apiserver.crt")
			s.TLSPrivateKeyFile = path.Join(s.CertDirectory, "apiserver.key")
			// TODO (cjcullen): Is PublicAddress the right address to sign a cert with?
			alternateIPs := []net.IP{config.ServiceReadWriteIP}
			alternateDNS := []string{"kubernetes.default.svc", "kubernetes.default", "kubernetes"}
			// It would be nice to set a fqdn subject alt name, but only the kubelets know, the apiserver is clueless
			// alternateDNS = append(alternateDNS, "kubernetes.default.svc.CLUSTER.DNS.NAME")
			if err := util.GenerateSelfSignedCert(config.PublicAddress.String(), s.TLSCertFile, s.TLSPrivateKeyFile, alternateIPs, alternateDNS); err != nil {
				glog.Errorf("Unable to generate self signed cert: %v", err)
			} else {
				glog.Infof("Using self-signed cert (%s, %s)", s.TLSCertFile, s.TLSPrivateKeyFile)
			}
		}

		go func() {
			defer util.HandleCrash()
			for {
				// err == systemd.SdNotifyNoSocket when not running on a systemd system
				if err := systemd.SdNotify("READY=1\n"); err != nil && err != systemd.SdNotifyNoSocket {
					glog.Errorf("Unable to send systemd daemon successful start message: %v\n", err)
				}
				if err := secureServer.ListenAndServeTLS(s.TLSCertFile, s.TLSPrivateKeyFile); err != nil {
					glog.Errorf("Unable to listen for secure (%v); will try again.", err)
				}
				time.Sleep(15 * time.Second)
			}
		}()
	}
	handler := apiserver.TimeoutHandler(m.InsecureHandler, longRunningTimeout)
	http := &http.Server{
		Addr:           insecureLocation,
		Handler:        apiserver.RecoverPanics(handler),
		MaxHeaderBytes: 1 << 20,
	}
	if secureLocation == "" {
		// err == systemd.SdNotifyNoSocket when not running on a systemd system
		if err := systemd.SdNotify("READY=1\n"); err != nil && err != systemd.SdNotifyNoSocket {
			glog.Errorf("Unable to send systemd daemon successful start message: %v\n", err)
		}
	}
	glog.Infof("Serving insecurely on %s", insecureLocation)
	glog.Fatal(http.ListenAndServe())
	return nil
}

func (s *APIServer) getRuntimeConfigValue(apiKey string, defaultValue bool) bool {
	flagValue, ok := s.RuntimeConfig[apiKey]
	if ok {
		if flagValue == "" {
			return true
		}
		boolValue, err := strconv.ParseBool(flagValue)
		if err != nil {
			glog.Fatalf("Invalid value of %s: %s, err: %v", apiKey, flagValue, err)
		}
		return boolValue
	}
	return defaultValue
}

// Parses the given runtime-config and formats it into map[string]ApiGroupVersionOverride
func (s *APIServer) parseRuntimeConfig() (map[string]genericapiserver.APIGroupVersionOverride, error) {
	// "api/all=false" allows users to selectively enable specific api versions.
	disableAllAPIs := false
	allAPIFlagValue, ok := s.RuntimeConfig["api/all"]
	if ok && allAPIFlagValue == "false" {
		disableAllAPIs = true
	}

	// "api/legacy=false" allows users to disable legacy api versions.
	disableLegacyAPIs := false
	legacyAPIFlagValue, ok := s.RuntimeConfig["api/legacy"]
	if ok && legacyAPIFlagValue == "false" {
		disableLegacyAPIs = true
	}
	_ = disableLegacyAPIs // hush the compiler while we don't have legacy APIs to disable.

	// "api/v1={true|false} allows users to enable/disable v1 API.
	// This takes preference over api/all and api/legacy, if specified.
	disableV1 := disableAllAPIs
	v1GroupVersion := "api/v1"
	disableV1 = !s.getRuntimeConfigValue(v1GroupVersion, !disableV1)
	apiGroupVersionOverrides := map[string]genericapiserver.APIGroupVersionOverride{}
	if disableV1 {
		apiGroupVersionOverrides[v1GroupVersion] = genericapiserver.APIGroupVersionOverride{
			Disable: true,
		}
	}

	// "extensions/v1beta1={true|false} allows users to enable/disable the extensions API.
	// This takes preference over api/all, if specified.
	disableExtensions := disableAllAPIs
	extensionsGroupVersion := "extensions/v1beta1"
	// TODO: Make this a loop over all group/versions when there are more of them.
	disableExtensions = !s.getRuntimeConfigValue(extensionsGroupVersion, !disableExtensions)
	if disableExtensions {
		apiGroupVersionOverrides[extensionsGroupVersion] = genericapiserver.APIGroupVersionOverride{
			Disable: true,
		}
	}

	for key := range s.RuntimeConfig {
		if strings.HasPrefix(key, v1GroupVersion+"/") {
			return nil, fmt.Errorf("api/v1 resources cannot be enabled/disabled individually")
		} else if strings.HasPrefix(key, extensionsGroupVersion+"/") {
			resource := strings.TrimPrefix(key, extensionsGroupVersion+"/")

			apiGroupVersionOverride := apiGroupVersionOverrides[extensionsGroupVersion]
			if apiGroupVersionOverride.ResourceOverrides == nil {
				apiGroupVersionOverride.ResourceOverrides = map[string]bool{}
			}
			apiGroupVersionOverride.ResourceOverrides[resource] = s.getRuntimeConfigValue(key, false)
			apiGroupVersionOverrides[extensionsGroupVersion] = apiGroupVersionOverride
		}
	}
	return apiGroupVersionOverrides, nil
}
