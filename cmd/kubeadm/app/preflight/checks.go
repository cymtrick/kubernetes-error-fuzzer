/*
Copyright 2016 The Kubernetes Authors.

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

package preflight

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/PuerkitoBio/purell"
	"github.com/pkg/errors"
	netutil "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/sets"
	versionutil "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/cmd/kubeadm/app/images"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/initsystem"
	utilruntime "k8s.io/kubernetes/cmd/kubeadm/app/util/runtime"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/system"
	kubeadmversion "k8s.io/kubernetes/cmd/kubeadm/app/version"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
	utilsexec "k8s.io/utils/exec"
	utilsnet "k8s.io/utils/net"
)

const (
	bridgenf                    = "/proc/sys/net/bridge/bridge-nf-call-iptables"
	bridgenf6                   = "/proc/sys/net/bridge/bridge-nf-call-ip6tables"
	ipv4Forward                 = "/proc/sys/net/ipv4/ip_forward"
	ipv6DefaultForwarding       = "/proc/sys/net/ipv6/conf/default/forwarding"
	externalEtcdRequestTimeout  = time.Duration(10 * time.Second)
	externalEtcdRequestRetries  = 3
	externalEtcdRequestInterval = time.Duration(5 * time.Second)
)

var (
	minExternalEtcdVersion = versionutil.MustParseSemantic(kubeadmconstants.MinExternalEtcdVersion)
)

// Error defines struct for communicating error messages generated by preflight checks
type Error struct {
	Msg string
}

// Error implements the standard error interface
func (e *Error) Error() string {
	return fmt.Sprintf("[preflight] Some fatal errors occurred:\n%s%s", e.Msg, "[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`")
}

// Preflight identifies this error as a preflight error
func (e *Error) Preflight() bool {
	return true
}

// Checker validates the state of the system to ensure kubeadm will be
// successful as often as possible.
type Checker interface {
	Check() (warnings, errorList []error)
	Name() string
}

// ContainerRuntimeCheck verifies the container runtime.
type ContainerRuntimeCheck struct {
	runtime utilruntime.ContainerRuntime
}

// Name returns label for RuntimeCheck.
func (ContainerRuntimeCheck) Name() string {
	return "CRI"
}

// Check validates the container runtime
func (crc ContainerRuntimeCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("validating the container runtime")
	if err := crc.runtime.IsRunning(); err != nil {
		errorList = append(errorList, err)
	}
	return warnings, errorList
}

// ServiceCheck verifies that the given service is enabled and active. If we do not
// detect a supported init system however, all checks are skipped and a warning is
// returned.
type ServiceCheck struct {
	Service       string
	CheckIfActive bool
	Label         string
}

// Name returns label for ServiceCheck. If not provided, will return based on the service parameter
func (sc ServiceCheck) Name() string {
	if sc.Label != "" {
		return sc.Label
	}
	return fmt.Sprintf("Service-%s", strings.Title(sc.Service))
}

// Check validates if the service is enabled and active.
func (sc ServiceCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("validating if the service is enabled and active")
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		return []error{err}, nil
	}

	if !initSystem.ServiceExists(sc.Service) {
		return []error{errors.Errorf("%s service does not exist", sc.Service)}, nil
	}

	if !initSystem.ServiceIsEnabled(sc.Service) {
		warnings = append(warnings,
			errors.Errorf("%s service is not enabled, please run '%s'",
				sc.Service, initSystem.EnableCommand(sc.Service)))
	}

	if sc.CheckIfActive && !initSystem.ServiceIsActive(sc.Service) {
		errorList = append(errorList,
			errors.Errorf("%s service is not active, please run 'systemctl start %s.service'",
				sc.Service, sc.Service))
	}

	return warnings, errorList
}

// FirewalldCheck checks if firewalld is enabled or active. If it is, warn the user that there may be problems
// if no actions are taken.
type FirewalldCheck struct {
	ports []int
}

// Name returns label for FirewalldCheck.
func (FirewalldCheck) Name() string {
	return "Firewalld"
}

// Check validates if the firewall is enabled and active.
func (fc FirewalldCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("validating if the firewall is enabled and active")
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		return []error{err}, nil
	}

	if !initSystem.ServiceExists("firewalld") {
		return nil, nil
	}

	if initSystem.ServiceIsActive("firewalld") {
		err := errors.Errorf("firewalld is active, please ensure ports %v are open or your cluster may not function correctly",
			fc.ports)
		return []error{err}, nil
	}

	return nil, nil
}

// PortOpenCheck ensures the given port is available for use.
type PortOpenCheck struct {
	port  int
	label string
}

// Name returns name for PortOpenCheck. If not known, will return "PortXXXX" based on port number
func (poc PortOpenCheck) Name() string {
	if poc.label != "" {
		return poc.label
	}
	return fmt.Sprintf("Port-%d", poc.port)
}

// Check validates if the particular port is available.
func (poc PortOpenCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infof("validating availability of port %d", poc.port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", poc.port))
	if err != nil {
		errorList = []error{errors.Errorf("Port %d is in use", poc.port)}
	}
	if ln != nil {
		if err = ln.Close(); err != nil {
			warnings = append(warnings,
				errors.Errorf("when closing port %d, encountered %v", poc.port, err))
		}
	}

	return warnings, errorList
}

// IsPrivilegedUserCheck verifies user is privileged (linux - root, windows - Administrator)
type IsPrivilegedUserCheck struct{}

// Name returns name for IsPrivilegedUserCheck
func (IsPrivilegedUserCheck) Name() string {
	return "IsPrivilegedUser"
}

// IsDockerSystemdCheck verifies if Docker is setup to use systemd as the cgroup driver.
type IsDockerSystemdCheck struct{}

// Name returns name for IsDockerSystemdCheck
func (IsDockerSystemdCheck) Name() string {
	return "IsDockerSystemdCheck"
}

// DirAvailableCheck checks if the given directory either does not exist, or is empty.
type DirAvailableCheck struct {
	Path  string
	Label string
}

// Name returns label for individual DirAvailableChecks. If not known, will return based on path.
func (dac DirAvailableCheck) Name() string {
	if dac.Label != "" {
		return dac.Label
	}
	return fmt.Sprintf("DirAvailable-%s", strings.Replace(dac.Path, "/", "-", -1))
}

// Check validates if a directory does not exist or empty.
func (dac DirAvailableCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infof("validating the existence and emptiness of directory %s", dac.Path)

	// If it doesn't exist we are good:
	if _, err := os.Stat(dac.Path); os.IsNotExist(err) {
		return nil, nil
	}

	f, err := os.Open(dac.Path)
	if err != nil {
		return nil, []error{errors.Wrapf(err, "unable to check if %s is empty", dac.Path)}
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err != io.EOF {
		return nil, []error{errors.Errorf("%s is not empty", dac.Path)}
	}

	return nil, nil
}

// FileAvailableCheck checks that the given file does not already exist.
type FileAvailableCheck struct {
	Path  string
	Label string
}

// Name returns label for individual FileAvailableChecks. If not known, will return based on path.
func (fac FileAvailableCheck) Name() string {
	if fac.Label != "" {
		return fac.Label
	}
	return fmt.Sprintf("FileAvailable-%s", strings.Replace(fac.Path, "/", "-", -1))
}

// Check validates if the given file does not already exist.
func (fac FileAvailableCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infof("validating the existence of file %s", fac.Path)

	if _, err := os.Stat(fac.Path); err == nil {
		return nil, []error{errors.Errorf("%s already exists", fac.Path)}
	}
	return nil, nil
}

// FileExistingCheck checks that the given file does not already exist.
type FileExistingCheck struct {
	Path  string
	Label string
}

// Name returns label for individual FileExistingChecks. If not known, will return based on path.
func (fac FileExistingCheck) Name() string {
	if fac.Label != "" {
		return fac.Label
	}
	return fmt.Sprintf("FileExisting-%s", strings.Replace(fac.Path, "/", "-", -1))
}

// Check validates if the given file already exists.
func (fac FileExistingCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infof("validating the existence of file %s", fac.Path)

	if _, err := os.Stat(fac.Path); err != nil {
		return nil, []error{errors.Errorf("%s doesn't exist", fac.Path)}
	}
	return nil, nil
}

// FileContentCheck checks that the given file contains the string Content.
type FileContentCheck struct {
	Path    string
	Content []byte
	Label   string
}

// Name returns label for individual FileContentChecks. If not known, will return based on path.
func (fcc FileContentCheck) Name() string {
	if fcc.Label != "" {
		return fcc.Label
	}
	return fmt.Sprintf("FileContent-%s", strings.Replace(fcc.Path, "/", "-", -1))
}

// Check validates if the given file contains the given content.
func (fcc FileContentCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infof("validating the contents of file %s", fcc.Path)
	f, err := os.Open(fcc.Path)
	if err != nil {
		return nil, []error{errors.Errorf("%s does not exist", fcc.Path)}
	}

	lr := io.LimitReader(f, int64(len(fcc.Content)))
	defer f.Close()

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, lr)
	if err != nil {
		return nil, []error{errors.Errorf("%s could not be read", fcc.Path)}
	}

	if !bytes.Equal(buf.Bytes(), fcc.Content) {
		return nil, []error{errors.Errorf("%s contents are not set to %s", fcc.Path, fcc.Content)}
	}
	return nil, []error{}

}

// InPathCheck checks if the given executable is present in $PATH
type InPathCheck struct {
	executable string
	mandatory  bool
	exec       utilsexec.Interface
	label      string
	suggestion string
}

// Name returns label for individual InPathCheck. If not known, will return based on path.
func (ipc InPathCheck) Name() string {
	if ipc.label != "" {
		return ipc.label
	}
	return fmt.Sprintf("FileExisting-%s", strings.Replace(ipc.executable, "/", "-", -1))
}

// Check validates if the given executable is present in the path.
func (ipc InPathCheck) Check() (warnings, errs []error) {
	klog.V(1).Infof("validating the presence of executable %s", ipc.executable)
	_, err := ipc.exec.LookPath(ipc.executable)
	if err != nil {
		if ipc.mandatory {
			// Return as an error:
			return nil, []error{errors.Errorf("%s not found in system path", ipc.executable)}
		}
		// Return as a warning:
		warningMessage := fmt.Sprintf("%s not found in system path", ipc.executable)
		if ipc.suggestion != "" {
			warningMessage += fmt.Sprintf("\nSuggestion: %s", ipc.suggestion)
		}
		return []error{errors.New(warningMessage)}, nil
	}
	return nil, nil
}

// HostnameCheck checks if hostname match dns sub domain regex.
// If hostname doesn't match this regex, kubelet will not launch static pods like kube-apiserver/kube-controller-manager and so on.
type HostnameCheck struct {
	nodeName string
}

// Name will return Hostname as name for HostnameCheck
func (HostnameCheck) Name() string {
	return "Hostname"
}

// Check validates if hostname match dns sub domain regex.
func (hc HostnameCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("checking whether the given node name is reachable using net.LookupHost")

	addr, err := net.LookupHost(hc.nodeName)
	if addr == nil {
		warnings = append(warnings, errors.Errorf("hostname \"%s\" could not be reached", hc.nodeName))
	}
	if err != nil {
		warnings = append(warnings, errors.Wrapf(err, "hostname \"%s\"", hc.nodeName))
	}
	return warnings, errorList
}

// HTTPProxyCheck checks if https connection to specific host is going
// to be done directly or over proxy. If proxy detected, it will return warning.
type HTTPProxyCheck struct {
	Proto string
	Host  string
}

// Name returns HTTPProxy as name for HTTPProxyCheck
func (hst HTTPProxyCheck) Name() string {
	return "HTTPProxy"
}

// Check validates http connectivity type, direct or via proxy.
func (hst HTTPProxyCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("validating if the connectivity type is via proxy or direct")
	u := (&url.URL{Scheme: hst.Proto, Host: hst.Host}).String()

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, []error{err}
	}

	proxy, err := netutil.SetOldTransportDefaults(&http.Transport{}).Proxy(req)
	if err != nil {
		return nil, []error{err}
	}
	if proxy != nil {
		return []error{errors.Errorf("Connection to %q uses proxy %q. If that is not intended, adjust your proxy settings", u, proxy)}, nil
	}
	return nil, nil
}

// HTTPProxyCIDRCheck checks if https connection to specific subnet is going
// to be done directly or over proxy. If proxy detected, it will return warning.
// Similar to HTTPProxyCheck above, but operates with subnets and uses API
// machinery transport defaults to simulate kube-apiserver accessing cluster
// services and pods.
type HTTPProxyCIDRCheck struct {
	Proto string
	CIDR  string
}

// Name will return HTTPProxyCIDR as name for HTTPProxyCIDRCheck
func (HTTPProxyCIDRCheck) Name() string {
	return "HTTPProxyCIDR"
}

// Check validates http connectivity to first IP address in the CIDR.
// If it is not directly connected and goes via proxy it will produce warning.
func (subnet HTTPProxyCIDRCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("validating http connectivity to first IP address in the CIDR")
	if len(subnet.CIDR) == 0 {
		return nil, nil
	}

	_, cidr, err := net.ParseCIDR(subnet.CIDR)
	if err != nil {
		return nil, []error{errors.Wrapf(err, "error parsing CIDR %q", subnet.CIDR)}
	}

	testIP, err := ipallocator.GetIndexedIP(cidr, 1)
	if err != nil {
		return nil, []error{errors.Wrapf(err, "unable to get first IP address from the given CIDR (%s)", cidr.String())}
	}

	testIPstring := testIP.String()
	if len(testIP) == net.IPv6len {
		testIPstring = fmt.Sprintf("[%s]:1234", testIP)
	}
	url := fmt.Sprintf("%s://%s/", subnet.Proto, testIPstring)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, []error{err}
	}

	// Utilize same transport defaults as it will be used by API server
	proxy, err := netutil.SetOldTransportDefaults(&http.Transport{}).Proxy(req)
	if err != nil {
		return nil, []error{err}
	}
	if proxy != nil {
		return []error{errors.Errorf("connection to %q uses proxy %q. This may lead to malfunctional cluster setup. Make sure that Pod and Services IP ranges specified correctly as exceptions in proxy configuration", subnet.CIDR, proxy)}, nil
	}
	return nil, nil
}

// SystemVerificationCheck defines struct used for running the system verification node check in test/e2e_node/system
type SystemVerificationCheck struct {
	IsDocker bool
}

// Name will return SystemVerification as name for SystemVerificationCheck
func (SystemVerificationCheck) Name() string {
	return "SystemVerification"
}

// Check runs all individual checks
func (sysver SystemVerificationCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("running all checks")
	// Create a buffered writer and choose a quite large value (1M) and suppose the output from the system verification test won't exceed the limit
	// Run the system verification check, but write to out buffered writer instead of stdout
	bufw := bufio.NewWriterSize(os.Stdout, 1*1024*1024)
	reporter := &system.StreamReporter{WriteStream: bufw}

	var errs []error
	var warns []error
	// All the common validators we'd like to run:
	var validators = []system.Validator{
		&system.KernelValidator{Reporter: reporter}}

	// run the docker validator only with docker runtime
	if sysver.IsDocker {
		validators = append(validators, &system.DockerValidator{Reporter: reporter})
	}

	if runtime.GOOS == "linux" {
		//add linux validators
		validators = append(validators,
			&system.OSValidator{Reporter: reporter},
			&system.CgroupsValidator{Reporter: reporter})
	}

	// Run all validators
	for _, v := range validators {
		warn, err := v.Validate(system.DefaultSysSpec)
		if err != nil {
			errs = append(errs, err)
		}
		if warn != nil {
			warns = append(warns, warn)
		}
	}

	if len(errs) != 0 {
		// Only print the output from the system verification check if the check failed
		fmt.Println("[preflight] The system verification failed. Printing the output from the verification:")
		bufw.Flush()
		return warns, errs
	}
	return warns, nil
}

// KubernetesVersionCheck validates Kubernetes and kubeadm versions
type KubernetesVersionCheck struct {
	KubeadmVersion    string
	KubernetesVersion string
}

// Name will return KubernetesVersion as name for KubernetesVersionCheck
func (KubernetesVersionCheck) Name() string {
	return "KubernetesVersion"
}

// Check validates Kubernetes and kubeadm versions
func (kubever KubernetesVersionCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("validating Kubernetes and kubeadm version")
	// Skip this check for "super-custom builds", where apimachinery/the overall codebase version is not set.
	if strings.HasPrefix(kubever.KubeadmVersion, "v0.0.0") {
		return nil, nil
	}

	kadmVersion, err := versionutil.ParseSemantic(kubever.KubeadmVersion)
	if err != nil {
		return nil, []error{errors.Wrapf(err, "couldn't parse kubeadm version %q", kubever.KubeadmVersion)}
	}

	k8sVersion, err := versionutil.ParseSemantic(kubever.KubernetesVersion)
	if err != nil {
		return nil, []error{errors.Wrapf(err, "couldn't parse Kubernetes version %q", kubever.KubernetesVersion)}
	}

	// Checks if k8sVersion greater or equal than the first unsupported versions by current version of kubeadm,
	// that is major.minor+1 (all patch and pre-releases versions included)
	// NB. in semver patches number is a numeric, while prerelease is a string where numeric identifiers always have lower precedence than non-numeric identifiers.
	//     thus setting the value to x.y.0-0 we are defining the very first patch - prereleases within x.y minor release.
	firstUnsupportedVersion := versionutil.MustParseSemantic(fmt.Sprintf("%d.%d.%s", kadmVersion.Major(), kadmVersion.Minor()+1, "0-0"))
	if k8sVersion.AtLeast(firstUnsupportedVersion) {
		return []error{errors.Errorf("Kubernetes version is greater than kubeadm version. Please consider to upgrade kubeadm. Kubernetes version: %s. Kubeadm version: %d.%d.x", k8sVersion, kadmVersion.Components()[0], kadmVersion.Components()[1])}, nil
	}

	return nil, nil
}

// KubeletVersionCheck validates installed kubelet version
type KubeletVersionCheck struct {
	KubernetesVersion string
	exec              utilsexec.Interface
}

// Name will return KubeletVersion as name for KubeletVersionCheck
func (KubeletVersionCheck) Name() string {
	return "KubeletVersion"
}

// Check validates kubelet version. It should be not less than minimal supported version
func (kubever KubeletVersionCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("validating kubelet version")
	kubeletVersion, err := GetKubeletVersion(kubever.exec)
	if err != nil {
		return nil, []error{errors.Wrap(err, "couldn't get kubelet version")}
	}
	if kubeletVersion.LessThan(kubeadmconstants.MinimumKubeletVersion) {
		return nil, []error{errors.Errorf("Kubelet version %q is lower than kubeadm can support. Please upgrade kubelet", kubeletVersion)}
	}

	if kubever.KubernetesVersion != "" {
		k8sVersion, err := versionutil.ParseSemantic(kubever.KubernetesVersion)
		if err != nil {
			return nil, []error{errors.Wrapf(err, "couldn't parse Kubernetes version %q", kubever.KubernetesVersion)}
		}
		if kubeletVersion.Major() > k8sVersion.Major() || kubeletVersion.Minor() > k8sVersion.Minor() {
			return nil, []error{errors.Errorf("the kubelet version is higher than the control plane version. This is not a supported version skew and may lead to a malfunctional cluster. Kubelet version: %q Control plane version: %q", kubeletVersion, k8sVersion)}
		}
	}
	return nil, nil
}

// SwapCheck warns if swap is enabled
type SwapCheck struct{}

// Name will return Swap as name for SwapCheck
func (SwapCheck) Name() string {
	return "Swap"
}

// Check validates whether swap is enabled or not
func (swc SwapCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("validating whether swap is enabled or not")
	f, err := os.Open("/proc/swaps")
	if err != nil {
		// /proc/swaps not available, thus no reasons to warn
		return nil, nil
	}
	defer f.Close()
	var buf []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, []error{errors.Wrap(err, "error parsing /proc/swaps")}
	}

	if len(buf) > 1 {
		return nil, []error{errors.New("running with swap on is not supported. Please disable swap")}
	}

	return nil, nil
}

type etcdVersionResponse struct {
	Etcdserver  string `json:"etcdserver"`
	Etcdcluster string `json:"etcdcluster"`
}

// ExternalEtcdVersionCheck checks if version of external etcd meets the demand of kubeadm
type ExternalEtcdVersionCheck struct {
	Etcd kubeadmapi.Etcd
}

// Name will return ExternalEtcdVersion as name for ExternalEtcdVersionCheck
func (ExternalEtcdVersionCheck) Name() string {
	return "ExternalEtcdVersion"
}

// Check validates external etcd version
// TODO: Use the official etcd Golang client for this instead?
func (evc ExternalEtcdVersionCheck) Check() (warnings, errorList []error) {
	klog.V(1).Infoln("validating the external etcd version")

	// Return quickly if the user isn't using external etcd
	if evc.Etcd.External.Endpoints == nil {
		return nil, nil
	}

	var config *tls.Config
	var err error
	if config, err = evc.configRootCAs(config); err != nil {
		errorList = append(errorList, err)
		return nil, errorList
	}
	if config, err = evc.configCertAndKey(config); err != nil {
		errorList = append(errorList, err)
		return nil, errorList
	}

	client := evc.getHTTPClient(config)
	for _, endpoint := range evc.Etcd.External.Endpoints {
		if _, err := url.Parse(endpoint); err != nil {
			errorList = append(errorList, errors.Wrapf(err, "failed to parse external etcd endpoint %s", endpoint))
			continue
		}
		resp := etcdVersionResponse{}
		var err error
		versionURL := fmt.Sprintf("%s/%s", endpoint, "version")
		if tmpVersionURL, err := purell.NormalizeURLString(versionURL, purell.FlagRemoveDuplicateSlashes); err != nil {
			errorList = append(errorList, errors.Wrapf(err, "failed to normalize external etcd version url %s", versionURL))
			continue
		} else {
			versionURL = tmpVersionURL
		}
		if err = getEtcdVersionResponse(client, versionURL, &resp); err != nil {
			errorList = append(errorList, err)
			continue
		}

		etcdVersion, err := versionutil.ParseSemantic(resp.Etcdserver)
		if err != nil {
			errorList = append(errorList, errors.Wrapf(err, "couldn't parse external etcd version %q", resp.Etcdserver))
			continue
		}
		if etcdVersion.LessThan(minExternalEtcdVersion) {
			errorList = append(errorList, errors.Errorf("this version of kubeadm only supports external etcd version >= %s. Current version: %s", kubeadmconstants.MinExternalEtcdVersion, resp.Etcdserver))
			continue
		}
	}

	return nil, errorList
}

// configRootCAs configures and returns a reference to tls.Config instance if CAFile is provided
func (evc ExternalEtcdVersionCheck) configRootCAs(config *tls.Config) (*tls.Config, error) {
	var CACertPool *x509.CertPool
	if evc.Etcd.External.CAFile != "" {
		CACert, err := ioutil.ReadFile(evc.Etcd.External.CAFile)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't load external etcd's server certificate %s", evc.Etcd.External.CAFile)
		}
		CACertPool = x509.NewCertPool()
		CACertPool.AppendCertsFromPEM(CACert)
	}
	if CACertPool != nil {
		if config == nil {
			config = &tls.Config{}
		}
		config.RootCAs = CACertPool
	}
	return config, nil
}

// configCertAndKey configures and returns a reference to tls.Config instance if CertFile and KeyFile pair is provided
func (evc ExternalEtcdVersionCheck) configCertAndKey(config *tls.Config) (*tls.Config, error) {
	var cert tls.Certificate
	if evc.Etcd.External.CertFile != "" && evc.Etcd.External.KeyFile != "" {
		var err error
		cert, err = tls.LoadX509KeyPair(evc.Etcd.External.CertFile, evc.Etcd.External.KeyFile)
		if err != nil {
			return nil, errors.Wrapf(err, "couldn't load external etcd's certificate and key pair %s, %s", evc.Etcd.External.CertFile, evc.Etcd.External.KeyFile)
		}
		if config == nil {
			config = &tls.Config{}
		}
		config.Certificates = []tls.Certificate{cert}
	}
	return config, nil
}

func (evc ExternalEtcdVersionCheck) getHTTPClient(config *tls.Config) *http.Client {
	if config != nil {
		transport := netutil.SetOldTransportDefaults(&http.Transport{
			TLSClientConfig: config,
		})
		return &http.Client{
			Transport: transport,
			Timeout:   externalEtcdRequestTimeout,
		}
	}
	return &http.Client{Timeout: externalEtcdRequestTimeout, Transport: netutil.SetOldTransportDefaults(&http.Transport{})}
}

func getEtcdVersionResponse(client *http.Client, url string, target interface{}) error {
	loopCount := externalEtcdRequestRetries + 1
	var err error
	var stopRetry bool
	for loopCount > 0 {
		if loopCount <= externalEtcdRequestRetries {
			time.Sleep(externalEtcdRequestInterval)
		}
		stopRetry, err = func() (stopRetry bool, err error) {
			r, err := client.Get(url)
			if err != nil {
				loopCount--
				return false, err
			}
			defer r.Body.Close()

			if r != nil && r.StatusCode >= 500 && r.StatusCode <= 599 {
				loopCount--
				return false, errors.Errorf("server responded with non-successful status: %s", r.Status)
			}
			return true, json.NewDecoder(r.Body).Decode(target)

		}()
		if stopRetry {
			break
		}
	}
	return err
}

// ImagePullCheck will pull container images used by kubeadm
type ImagePullCheck struct {
	runtime   utilruntime.ContainerRuntime
	imageList []string
}

// Name returns the label for ImagePullCheck
func (ImagePullCheck) Name() string {
	return "ImagePull"
}

// Check pulls images required by kubeadm. This is a mutating check
func (ipc ImagePullCheck) Check() (warnings, errorList []error) {
	for _, image := range ipc.imageList {
		ret, err := ipc.runtime.ImageExists(image)
		if ret && err == nil {
			klog.V(1).Infof("image exists: %s", image)
			continue
		}
		if err != nil {
			errorList = append(errorList, errors.Wrapf(err, "failed to check if image %s exists", image))
		}
		klog.V(1).Infof("pulling %s", image)
		if err := ipc.runtime.PullImage(image); err != nil {
			errorList = append(errorList, errors.Wrapf(err, "failed to pull image %s", image))
		}
	}
	return warnings, errorList
}

// NumCPUCheck checks if current number of CPUs is not less than required
type NumCPUCheck struct {
	NumCPU int
}

// Name returns the label for NumCPUCheck
func (NumCPUCheck) Name() string {
	return "NumCPU"
}

// Check number of CPUs required by kubeadm
func (ncc NumCPUCheck) Check() (warnings, errorList []error) {
	numCPU := runtime.NumCPU()
	if numCPU < ncc.NumCPU {
		errorList = append(errorList, errors.Errorf("the number of available CPUs %d is less than the required %d", numCPU, ncc.NumCPU))
	}
	return warnings, errorList
}

// RunInitNodeChecks executes all individual, applicable to control-plane node checks.
// The boolean flag 'isSecondaryControlPlane' controls whether we are running checks in a --join-control-plane scenario.
// The boolean flag 'downloadCerts' controls whether we should skip checks on certificates because we are downloading them.
// If the flag is set to true we should skip checks already executed by RunJoinNodeChecks.
func RunInitNodeChecks(execer utilsexec.Interface, cfg *kubeadmapi.InitConfiguration, ignorePreflightErrors sets.String, isSecondaryControlPlane bool, downloadCerts bool) error {
	if !isSecondaryControlPlane {
		// First, check if we're root separately from the other preflight checks and fail fast
		if err := RunRootCheckOnly(ignorePreflightErrors); err != nil {
			return err
		}
	}

	manifestsDir := filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ManifestsSubDirName)
	checks := []Checker{
		NumCPUCheck{NumCPU: kubeadmconstants.ControlPlaneNumCPU},
		KubernetesVersionCheck{KubernetesVersion: cfg.KubernetesVersion, KubeadmVersion: kubeadmversion.Get().GitVersion},
		FirewalldCheck{ports: []int{int(cfg.LocalAPIEndpoint.BindPort), kubeadmconstants.KubeletPort}},
		PortOpenCheck{port: int(cfg.LocalAPIEndpoint.BindPort)},
		PortOpenCheck{port: kubeadmconstants.InsecureSchedulerPort},
		PortOpenCheck{port: kubeadmconstants.InsecureKubeControllerManagerPort},
		FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.KubeAPIServer, manifestsDir)},
		FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.KubeControllerManager, manifestsDir)},
		FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.KubeScheduler, manifestsDir)},
		FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.Etcd, manifestsDir)},
		HTTPProxyCheck{Proto: "https", Host: cfg.LocalAPIEndpoint.AdvertiseAddress},
		HTTPProxyCIDRCheck{Proto: "https", CIDR: cfg.Networking.ServiceSubnet},
	}

	cidrs := strings.Split(cfg.Networking.PodSubnet, ",")
	for _, cidr := range cidrs {
		checks = append(checks, HTTPProxyCIDRCheck{Proto: "https", CIDR: cidr})
	}

	if !isSecondaryControlPlane {
		checks = addCommonChecks(execer, cfg.KubernetesVersion, &cfg.NodeRegistration, checks)

		// Check if Bridge-netfilter and IPv6 relevant flags are set
		if ip := net.ParseIP(cfg.LocalAPIEndpoint.AdvertiseAddress); ip != nil {
			if utilsnet.IsIPv6(ip) {
				checks = append(checks,
					FileContentCheck{Path: bridgenf6, Content: []byte{'1'}},
					FileContentCheck{Path: ipv6DefaultForwarding, Content: []byte{'1'}},
				)
			}
		}

		// if using an external etcd
		if cfg.Etcd.External != nil {
			// Check external etcd version before creating the cluster
			checks = append(checks, ExternalEtcdVersionCheck{Etcd: cfg.Etcd})
		}
	}

	if cfg.Etcd.Local != nil {
		// Only do etcd related checks when required to install a local etcd
		checks = append(checks,
			PortOpenCheck{port: kubeadmconstants.EtcdListenClientPort},
			PortOpenCheck{port: kubeadmconstants.EtcdListenPeerPort},
			DirAvailableCheck{Path: cfg.Etcd.Local.DataDir},
		)
	}

	if cfg.Etcd.External != nil && !(isSecondaryControlPlane && downloadCerts) {
		// Only check etcd certificates when using an external etcd and not joining with automatic download of certs
		if cfg.Etcd.External.CAFile != "" {
			checks = append(checks, FileExistingCheck{Path: cfg.Etcd.External.CAFile, Label: "ExternalEtcdClientCertificates"})
		}
		if cfg.Etcd.External.CertFile != "" {
			checks = append(checks, FileExistingCheck{Path: cfg.Etcd.External.CertFile, Label: "ExternalEtcdClientCertificates"})
		}
		if cfg.Etcd.External.KeyFile != "" {
			checks = append(checks, FileExistingCheck{Path: cfg.Etcd.External.KeyFile, Label: "ExternalEtcdClientCertificates"})
		}
	}

	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}

// RunJoinNodeChecks executes all individual, applicable to node checks.
func RunJoinNodeChecks(execer utilsexec.Interface, cfg *kubeadmapi.JoinConfiguration, ignorePreflightErrors sets.String) error {
	// First, check if we're root separately from the other preflight checks and fail fast
	if err := RunRootCheckOnly(ignorePreflightErrors); err != nil {
		return err
	}

	checks := []Checker{
		DirAvailableCheck{Path: filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ManifestsSubDirName)},
		FileAvailableCheck{Path: filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.KubeletKubeConfigFileName)},
		FileAvailableCheck{Path: filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.KubeletBootstrapKubeConfigFileName)},
	}
	checks = addCommonChecks(execer, "", &cfg.NodeRegistration, checks)
	if cfg.ControlPlane == nil {
		checks = append(checks, FileAvailableCheck{Path: cfg.CACertPath})
	}

	addIPv6Checks := false
	if cfg.Discovery.BootstrapToken != nil {
		ipstr, _, err := net.SplitHostPort(cfg.Discovery.BootstrapToken.APIServerEndpoint)
		if err == nil {
			checks = append(checks,
				HTTPProxyCheck{Proto: "https", Host: ipstr},
			)
			if ip := net.ParseIP(ipstr); ip != nil {
				if utilsnet.IsIPv6(ip) {
					addIPv6Checks = true
				}
			}
		}
	}
	if addIPv6Checks {
		checks = append(checks,
			FileContentCheck{Path: bridgenf6, Content: []byte{'1'}},
			FileContentCheck{Path: ipv6DefaultForwarding, Content: []byte{'1'}},
		)
	}

	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}

// addCommonChecks is a helper function to duplicate checks that are common between both the
// kubeadm init and join commands
func addCommonChecks(execer utilsexec.Interface, k8sVersion string, nodeReg *kubeadmapi.NodeRegistrationOptions, checks []Checker) []Checker {
	containerRuntime, err := utilruntime.NewContainerRuntime(execer, nodeReg.CRISocket)
	isDocker := false
	if err != nil {
		fmt.Printf("[preflight] WARNING: Couldn't create the interface used for talking to the container runtime: %v\n", err)
	} else {
		checks = append(checks, ContainerRuntimeCheck{runtime: containerRuntime})
		if containerRuntime.IsDocker() {
			isDocker = true
			checks = append(checks, ServiceCheck{Service: "docker", CheckIfActive: true})
			// Linux only
			// TODO: support other CRIs for this check eventually
			// https://github.com/kubernetes/kubeadm/issues/874
			checks = append(checks, IsDockerSystemdCheck{})
		}
	}

	// non-windows checks
	if runtime.GOOS == "linux" {
		if !isDocker {
			checks = append(checks, InPathCheck{executable: "crictl", mandatory: true, exec: execer})
		}
		checks = append(checks,
			FileContentCheck{Path: bridgenf, Content: []byte{'1'}},
			FileContentCheck{Path: ipv4Forward, Content: []byte{'1'}},
			SwapCheck{},
			InPathCheck{executable: "ip", mandatory: true, exec: execer},
			InPathCheck{executable: "iptables", mandatory: true, exec: execer},
			InPathCheck{executable: "mount", mandatory: true, exec: execer},
			InPathCheck{executable: "nsenter", mandatory: true, exec: execer},
			InPathCheck{executable: "ebtables", mandatory: false, exec: execer},
			InPathCheck{executable: "ethtool", mandatory: false, exec: execer},
			InPathCheck{executable: "socat", mandatory: false, exec: execer},
			InPathCheck{executable: "tc", mandatory: false, exec: execer},
			InPathCheck{executable: "touch", mandatory: false, exec: execer})
	}
	checks = append(checks,
		SystemVerificationCheck{IsDocker: isDocker},
		HostnameCheck{nodeName: nodeReg.Name},
		KubeletVersionCheck{KubernetesVersion: k8sVersion, exec: execer},
		ServiceCheck{Service: "kubelet", CheckIfActive: false},
		PortOpenCheck{port: kubeadmconstants.KubeletPort})
	return checks
}

// RunRootCheckOnly initializes checks slice of structs and call RunChecks
func RunRootCheckOnly(ignorePreflightErrors sets.String) error {
	checks := []Checker{
		IsPrivilegedUserCheck{},
	}

	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}

// RunPullImagesCheck will pull images kubeadm needs if they are not found on the system
func RunPullImagesCheck(execer utilsexec.Interface, cfg *kubeadmapi.InitConfiguration, ignorePreflightErrors sets.String) error {
	containerRuntime, err := utilruntime.NewContainerRuntime(utilsexec.New(), cfg.NodeRegistration.CRISocket)
	if err != nil {
		return err
	}

	checks := []Checker{
		ImagePullCheck{runtime: containerRuntime, imageList: images.GetControlPlaneImages(&cfg.ClusterConfiguration)},
	}
	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}

// RunChecks runs each check, displays it's warnings/errors, and once all
// are processed will exit if any errors occurred.
func RunChecks(checks []Checker, ww io.Writer, ignorePreflightErrors sets.String) error {
	var errsBuffer bytes.Buffer

	for _, c := range checks {
		name := c.Name()
		warnings, errs := c.Check()

		if setHasItemOrAll(ignorePreflightErrors, name) {
			// Decrease severity of errors to warnings for this check
			warnings = append(warnings, errs...)
			errs = []error{}
		}

		for _, w := range warnings {
			io.WriteString(ww, fmt.Sprintf("\t[WARNING %s]: %v\n", name, w))
		}
		for _, i := range errs {
			errsBuffer.WriteString(fmt.Sprintf("\t[ERROR %s]: %v\n", name, i.Error()))
		}
	}
	if errsBuffer.Len() > 0 {
		return &Error{Msg: errsBuffer.String()}
	}
	return nil
}

// setHasItemOrAll is helper function that return true if item is present in the set (case insensitive) or special key 'all' is present
func setHasItemOrAll(s sets.String, item string) bool {
	if s.Has("all") || s.Has(strings.ToLower(item)) {
		return true
	}
	return false
}
