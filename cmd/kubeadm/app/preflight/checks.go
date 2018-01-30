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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"crypto/tls"
	"crypto/x509"

	"github.com/PuerkitoBio/purell"
	"github.com/blang/semver"
	"github.com/spf13/pflag"

	"net/url"

	netutil "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/sets"
	apiservoptions "k8s.io/kubernetes/cmd/kube-apiserver/app/options"
	cmoptions "k8s.io/kubernetes/cmd/kube-controller-manager/app/options"
	schedulerapp "k8s.io/kubernetes/cmd/kube-scheduler/app"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	"k8s.io/kubernetes/pkg/apis/core/validation"
	authzmodes "k8s.io/kubernetes/pkg/kubeapiserver/authorizer/modes"
	"k8s.io/kubernetes/pkg/registry/core/service/ipallocator"
	"k8s.io/kubernetes/pkg/util/initsystem"
	versionutil "k8s.io/kubernetes/pkg/util/version"
	kubeadmversion "k8s.io/kubernetes/pkg/version"
	"k8s.io/kubernetes/test/e2e_node/system"
	utilsexec "k8s.io/utils/exec"
)

const (
	bridgenf                    = "/proc/sys/net/bridge/bridge-nf-call-iptables"
	bridgenf6                   = "/proc/sys/net/bridge/bridge-nf-call-ip6tables"
	externalEtcdRequestTimeout  = time.Duration(10 * time.Second)
	externalEtcdRequestRetries  = 3
	externalEtcdRequestInterval = time.Duration(5 * time.Second)
)

var (
	minExternalEtcdVersion = semver.MustParse(kubeadmconstants.MinExternalEtcdVersion)
)

// Error defines struct for communicating error messages generated by preflight checks
type Error struct {
	Msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("[preflight] Some fatal errors occurred:\n%s%s", e.Msg, "[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`")
}

// Checker validates the state of the system to ensure kubeadm will be
// successful as often as possible.
type Checker interface {
	Check() (warnings, errors []error)
	Name() string
}

// CRICheck verifies the container runtime through the CRI.
type CRICheck struct {
	socket string
	exec   utilsexec.Interface
}

// Name returns label for CRICheck.
func (CRICheck) Name() string {
	return "CRI"
}

// Check validates the container runtime through the CRI.
func (criCheck CRICheck) Check() (warnings, errors []error) {
	if err := criCheck.exec.Command("sh", "-c", fmt.Sprintf("crictl -r %s info", criCheck.socket)).Run(); err != nil {
		errors = append(errors, fmt.Errorf("unable to check if the container runtime at %q is running: %s", criCheck.socket, err))
		return warnings, errors
	}
	return warnings, errors
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
func (sc ServiceCheck) Check() (warnings, errors []error) {
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		return []error{err}, nil
	}

	warnings = []error{}

	if !initSystem.ServiceExists(sc.Service) {
		warnings = append(warnings, fmt.Errorf("%s service does not exist", sc.Service))
		return warnings, nil
	}

	if !initSystem.ServiceIsEnabled(sc.Service) {
		warnings = append(warnings,
			fmt.Errorf("%s service is not enabled, please run 'systemctl enable %s.service'",
				sc.Service, sc.Service))
	}

	if sc.CheckIfActive && !initSystem.ServiceIsActive(sc.Service) {
		errors = append(errors,
			fmt.Errorf("%s service is not active, please run 'systemctl start %s.service'",
				sc.Service, sc.Service))
	}

	return warnings, errors
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
func (fc FirewalldCheck) Check() (warnings, errors []error) {
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		return []error{err}, nil
	}

	warnings = []error{}

	if !initSystem.ServiceExists("firewalld") {
		return nil, nil
	}

	if initSystem.ServiceIsActive("firewalld") {
		warnings = append(warnings,
			fmt.Errorf("firewalld is active, please ensure ports %v are open or your cluster may not function correctly",
				fc.ports))
	}

	return warnings, errors
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
func (poc PortOpenCheck) Check() (warnings, errors []error) {
	errors = []error{}
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", poc.port))
	if err != nil {
		errors = append(errors, fmt.Errorf("Port %d is in use", poc.port))
	}
	if ln != nil {
		ln.Close()
	}

	return nil, errors
}

// IsPrivilegedUserCheck verifies user is privileged (linux - root, windows - Administrator)
type IsPrivilegedUserCheck struct{}

// Name returns name for IsPrivilegedUserCheck
func (IsPrivilegedUserCheck) Name() string {
	return "IsPrivilegedUser"
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
func (dac DirAvailableCheck) Check() (warnings, errors []error) {
	errors = []error{}
	// If it doesn't exist we are good:
	if _, err := os.Stat(dac.Path); os.IsNotExist(err) {
		return nil, nil
	}

	f, err := os.Open(dac.Path)
	if err != nil {
		errors = append(errors, fmt.Errorf("unable to check if %s is empty: %s", dac.Path, err))
		return nil, errors
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err != io.EOF {
		errors = append(errors, fmt.Errorf("%s is not empty", dac.Path))
	}

	return nil, errors
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
func (fac FileAvailableCheck) Check() (warnings, errors []error) {
	errors = []error{}
	if _, err := os.Stat(fac.Path); err == nil {
		errors = append(errors, fmt.Errorf("%s already exists", fac.Path))
	}
	return nil, errors
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
func (fac FileExistingCheck) Check() (warnings, errors []error) {
	errors = []error{}
	if _, err := os.Stat(fac.Path); err != nil {
		errors = append(errors, fmt.Errorf("%s doesn't exist", fac.Path))
	}
	return nil, errors
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
func (fcc FileContentCheck) Check() (warnings, errors []error) {
	f, err := os.Open(fcc.Path)
	if err != nil {
		return nil, []error{fmt.Errorf("%s does not exist", fcc.Path)}
	}

	lr := io.LimitReader(f, int64(len(fcc.Content)))
	defer f.Close()

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, lr)
	if err != nil {
		return nil, []error{fmt.Errorf("%s could not be read", fcc.Path)}
	}

	if !bytes.Equal(buf.Bytes(), fcc.Content) {
		return nil, []error{fmt.Errorf("%s contents are not set to %s", fcc.Path, fcc.Content)}
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
	_, err := ipc.exec.LookPath(ipc.executable)
	if err != nil {
		if ipc.mandatory {
			// Return as an error:
			return nil, []error{fmt.Errorf("%s not found in system path", ipc.executable)}
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
func (hc HostnameCheck) Check() (warnings, errors []error) {
	errors = []error{}
	warnings = []error{}
	for _, msg := range validation.ValidateNodeName(hc.nodeName, false) {
		errors = append(errors, fmt.Errorf("hostname \"%s\" %s", hc.nodeName, msg))
	}
	addr, err := net.LookupHost(hc.nodeName)
	if addr == nil {
		warnings = append(warnings, fmt.Errorf("hostname \"%s\" could not be reached", hc.nodeName))
	}
	if err != nil {
		warnings = append(warnings, fmt.Errorf("hostname \"%s\" %s", hc.nodeName, err))
	}
	return warnings, errors
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
func (hst HTTPProxyCheck) Check() (warnings, errors []error) {

	url := fmt.Sprintf("%s://%s", hst.Proto, hst.Host)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, []error{err}
	}

	proxy, err := netutil.SetOldTransportDefaults(&http.Transport{}).Proxy(req)
	if err != nil {
		return nil, []error{err}
	}
	if proxy != nil {
		return []error{fmt.Errorf("Connection to %q uses proxy %q. If that is not intended, adjust your proxy settings", url, proxy)}, nil
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
func (subnet HTTPProxyCIDRCheck) Check() (warnings, errors []error) {

	if len(subnet.CIDR) == 0 {
		return nil, nil
	}

	_, cidr, err := net.ParseCIDR(subnet.CIDR)
	if err != nil {
		return nil, []error{fmt.Errorf("error parsing CIDR %q: %v", subnet.CIDR, err)}
	}

	testIP, err := ipallocator.GetIndexedIP(cidr, 1)
	if err != nil {
		return nil, []error{fmt.Errorf("unable to get first IP address from the given CIDR (%s): %v", cidr.String(), err)}
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
		return []error{fmt.Errorf("connection to %q uses proxy %q. This may lead to malfunctional cluster setup. Make sure that Pod and Services IP ranges specified correctly as exceptions in proxy configuration", subnet.CIDR, proxy)}, nil
	}
	return nil, nil
}

// ExtraArgsCheck checks if arguments are valid.
type ExtraArgsCheck struct {
	APIServerExtraArgs         map[string]string
	ControllerManagerExtraArgs map[string]string
	SchedulerExtraArgs         map[string]string
}

// Name will return ExtraArgs as name for ExtraArgsCheck
func (ExtraArgsCheck) Name() string {
	return "ExtraArgs"
}

// Check validates additional arguments of the control plane components.
func (eac ExtraArgsCheck) Check() (warnings, errors []error) {
	argsCheck := func(name string, args map[string]string, f *pflag.FlagSet) []error {
		errs := []error{}
		for k, v := range args {
			if err := f.Set(k, v); err != nil {
				errs = append(errs, fmt.Errorf("%s: failed to parse extra argument --%s=%s", name, k, v))
			}
		}
		return errs
	}

	warnings = []error{}
	if len(eac.APIServerExtraArgs) > 0 {
		flags := pflag.NewFlagSet("", pflag.ContinueOnError)
		s := apiservoptions.NewServerRunOptions()
		s.AddFlags(flags)
		warnings = append(warnings, argsCheck("kube-apiserver", eac.APIServerExtraArgs, flags)...)
	}
	if len(eac.ControllerManagerExtraArgs) > 0 {
		flags := pflag.NewFlagSet("", pflag.ContinueOnError)
		s := cmoptions.NewCMServer()
		s.AddFlags(flags, []string{}, []string{})
		warnings = append(warnings, argsCheck("kube-controller-manager", eac.ControllerManagerExtraArgs, flags)...)
	}
	if len(eac.SchedulerExtraArgs) > 0 {
		opts, err := schedulerapp.NewOptions()
		if err != nil {
			warnings = append(warnings, err)
		}
		flags := pflag.NewFlagSet("", pflag.ContinueOnError)
		opts.AddFlags(flags)
		warnings = append(warnings, argsCheck("kube-scheduler", eac.SchedulerExtraArgs, flags)...)
	}
	return warnings, nil
}

// SystemVerificationCheck defines struct used for for running the system verification node check in test/e2e_node/system
type SystemVerificationCheck struct {
	CRISocket string
}

// Name will return SystemVerification as name for SystemVerificationCheck
func (SystemVerificationCheck) Name() string {
	return "SystemVerification"
}

// Check runs all individual checks
func (sysver SystemVerificationCheck) Check() (warnings, errors []error) {
	// Create a buffered writer and choose a quite large value (1M) and suppose the output from the system verification test won't exceed the limit
	// Run the system verification check, but write to out buffered writer instead of stdout
	bufw := bufio.NewWriterSize(os.Stdout, 1*1024*1024)
	reporter := &system.StreamReporter{WriteStream: bufw}

	var errs []error
	var warns []error
	// All the common validators we'd like to run:
	var validators = []system.Validator{
		&system.KernelValidator{Reporter: reporter}}

	// run the docker validator only with dockershim
	if sysver.CRISocket == "/var/run/dockershim.sock" {
		// https://github.com/kubernetes/kubeadm/issues/533
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

// KubernetesVersionCheck validates kubernetes and kubeadm versions
type KubernetesVersionCheck struct {
	KubeadmVersion    string
	KubernetesVersion string
}

// Name will return KubernetesVersion as name for KubernetesVersionCheck
func (KubernetesVersionCheck) Name() string {
	return "KubernetesVersion"
}

// Check validates kubernetes and kubeadm versions
func (kubever KubernetesVersionCheck) Check() (warnings, errors []error) {

	// Skip this check for "super-custom builds", where apimachinery/the overall codebase version is not set.
	if strings.HasPrefix(kubever.KubeadmVersion, "v0.0.0") {
		return nil, nil
	}

	kadmVersion, err := versionutil.ParseSemantic(kubever.KubeadmVersion)
	if err != nil {
		return nil, []error{fmt.Errorf("couldn't parse kubeadm version %q: %v", kubever.KubeadmVersion, err)}
	}

	k8sVersion, err := versionutil.ParseSemantic(kubever.KubernetesVersion)
	if err != nil {
		return nil, []error{fmt.Errorf("couldn't parse kubernetes version %q: %v", kubever.KubernetesVersion, err)}
	}

	// Checks if k8sVersion greater or equal than the first unsupported versions by current version of kubeadm,
	// that is major.minor+1 (all patch and pre-releases versions included)
	// NB. in semver patches number is a numeric, while prerelease is a string where numeric identifiers always have lower precedence than non-numeric identifiers.
	//     thus setting the value to x.y.0-0 we are defining the very first patch - prereleases within x.y minor release.
	firstUnsupportedVersion := versionutil.MustParseSemantic(fmt.Sprintf("%d.%d.%s", kadmVersion.Major(), kadmVersion.Minor()+1, "0-0"))
	if k8sVersion.AtLeast(firstUnsupportedVersion) {
		return []error{fmt.Errorf("kubernetes version is greater than kubeadm version. Please consider to upgrade kubeadm. kubernetes version: %s. Kubeadm version: %d.%d.x", k8sVersion, kadmVersion.Components()[0], kadmVersion.Components()[1])}, nil
	}

	return nil, nil
}

// KubeletVersionCheck validates installed kubelet version
type KubeletVersionCheck struct {
	KubernetesVersion string
}

// Name will return KubeletVersion as name for KubeletVersionCheck
func (KubeletVersionCheck) Name() string {
	return "KubeletVersion"
}

// Check validates kubelet version. It should be not less than minimal supported version
func (kubever KubeletVersionCheck) Check() (warnings, errors []error) {
	kubeletVersion, err := GetKubeletVersion()
	if err != nil {
		return nil, []error{fmt.Errorf("couldn't get kubelet version: %v", err)}
	}
	if kubeletVersion.LessThan(kubeadmconstants.MinimumKubeletVersion) {
		return nil, []error{fmt.Errorf("Kubelet version %q is lower than kubadm can support. Please upgrade kubelet", kubeletVersion)}
	}

	if kubever.KubernetesVersion != "" {
		k8sVersion, err := versionutil.ParseSemantic(kubever.KubernetesVersion)
		if err != nil {
			return nil, []error{fmt.Errorf("couldn't parse kubernetes version %q: %v", kubever.KubernetesVersion, err)}
		}
		if kubeletVersion.Major() > k8sVersion.Major() || kubeletVersion.Minor() > k8sVersion.Minor() {
			return nil, []error{fmt.Errorf("the kubelet version is higher than the control plane version. This is not a supported version skew and may lead to a malfunctional cluster. Kubelet version: %q Control plane version: %q", kubeletVersion, k8sVersion)}
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
func (swc SwapCheck) Check() (warnings, errors []error) {
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
		return nil, []error{fmt.Errorf("error parsing /proc/swaps: %v", err)}
	}

	if len(buf) > 1 {
		return nil, []error{fmt.Errorf("running with swap on is not supported. Please disable swap")}
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
func (evc ExternalEtcdVersionCheck) Check() (warnings, errors []error) {
	var config *tls.Config
	var err error
	if config, err = evc.configRootCAs(config); err != nil {
		errors = append(errors, err)
		return nil, errors
	}
	if config, err = evc.configCertAndKey(config); err != nil {
		errors = append(errors, err)
		return nil, errors
	}

	client := evc.getHTTPClient(config)
	for _, endpoint := range evc.Etcd.Endpoints {
		if _, err := url.Parse(endpoint); err != nil {
			errors = append(errors, fmt.Errorf("failed to parse external etcd endpoint %s : %v", endpoint, err))
			continue
		}
		resp := etcdVersionResponse{}
		var err error
		versionURL := fmt.Sprintf("%s/%s", endpoint, "version")
		if tmpVersionURL, err := purell.NormalizeURLString(versionURL, purell.FlagRemoveDuplicateSlashes); err != nil {
			errors = append(errors, fmt.Errorf("failed to normalize external etcd version url %s : %v", versionURL, err))
			continue
		} else {
			versionURL = tmpVersionURL
		}
		if err = getEtcdVersionResponse(client, versionURL, &resp); err != nil {
			errors = append(errors, err)
			continue
		}

		etcdVersion, err := semver.Parse(resp.Etcdserver)
		if err != nil {
			errors = append(errors, fmt.Errorf("couldn't parse external etcd version %q: %v", resp.Etcdserver, err))
			continue
		}
		if etcdVersion.LT(minExternalEtcdVersion) {
			errors = append(errors, fmt.Errorf("this version of kubeadm only supports external etcd version >= %s. Current version: %s", kubeadmconstants.MinExternalEtcdVersion, resp.Etcdserver))
			continue
		}
	}

	return nil, errors
}

// configRootCAs configures and returns a reference to tls.Config instance if CAFile is provided
func (evc ExternalEtcdVersionCheck) configRootCAs(config *tls.Config) (*tls.Config, error) {
	var CACertPool *x509.CertPool
	if evc.Etcd.CAFile != "" {
		CACert, err := ioutil.ReadFile(evc.Etcd.CAFile)
		if err != nil {
			return nil, fmt.Errorf("couldn't load external etcd's server certificate %s: %v", evc.Etcd.CAFile, err)
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
	if evc.Etcd.CertFile != "" && evc.Etcd.KeyFile != "" {
		var err error
		cert, err = tls.LoadX509KeyPair(evc.Etcd.CertFile, evc.Etcd.KeyFile)
		if err != nil {
			return nil, fmt.Errorf("couldn't load external etcd's certificate and key pair %s, %s: %v", evc.Etcd.CertFile, evc.Etcd.KeyFile, err)
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
				return false, nil
			}
			defer r.Body.Close()

			if r != nil && r.StatusCode >= 500 && r.StatusCode <= 599 {
				loopCount--
				return false, nil
			}
			return true, json.NewDecoder(r.Body).Decode(target)

		}()
		if stopRetry {
			break
		}
	}
	return err
}

// RunInitMasterChecks executes all individual, applicable to Master node checks.
func RunInitMasterChecks(execer utilsexec.Interface, cfg *kubeadmapi.MasterConfiguration, criSocket string, ignorePreflightErrors sets.String) error {
	// First, check if we're root separately from the other preflight checks and fail fast
	if err := RunRootCheckOnly(ignorePreflightErrors); err != nil {
		return err
	}

	// check if we can use crictl to perform checks via the CRI
	criCtlChecker := InPathCheck{
		executable: "crictl",
		mandatory:  false,
		exec:       execer,
		suggestion: fmt.Sprintf("go get %v", kubeadmconstants.CRICtlPackage),
	}
	warns, _ := criCtlChecker.Check()
	useCRI := len(warns) == 0

	manifestsDir := filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ManifestsSubDirName)

	checks := []Checker{
		KubernetesVersionCheck{KubernetesVersion: cfg.KubernetesVersion, KubeadmVersion: kubeadmversion.Get().GitVersion},
		SystemVerificationCheck{CRISocket: criSocket},
		IsPrivilegedUserCheck{},
		HostnameCheck{nodeName: cfg.NodeName},
		KubeletVersionCheck{KubernetesVersion: cfg.KubernetesVersion},
		ServiceCheck{Service: "kubelet", CheckIfActive: false},
		FirewalldCheck{ports: []int{int(cfg.API.BindPort), 10250}},
		PortOpenCheck{port: int(cfg.API.BindPort)},
		PortOpenCheck{port: 10250},
		PortOpenCheck{port: 10251},
		PortOpenCheck{port: 10252},
		FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.KubeAPIServer, manifestsDir)},
		FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.KubeControllerManager, manifestsDir)},
		FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.KubeScheduler, manifestsDir)},
		FileAvailableCheck{Path: kubeadmconstants.GetStaticPodFilepath(kubeadmconstants.Etcd, manifestsDir)},
		FileContentCheck{Path: bridgenf, Content: []byte{'1'}},
		SwapCheck{},
		InPathCheck{executable: "ip", mandatory: true, exec: execer},
		InPathCheck{executable: "iptables", mandatory: true, exec: execer},
		InPathCheck{executable: "mount", mandatory: true, exec: execer},
		InPathCheck{executable: "nsenter", mandatory: true, exec: execer},
		InPathCheck{executable: "ebtables", mandatory: false, exec: execer},
		InPathCheck{executable: "ethtool", mandatory: false, exec: execer},
		InPathCheck{executable: "socat", mandatory: false, exec: execer},
		InPathCheck{executable: "tc", mandatory: false, exec: execer},
		InPathCheck{executable: "touch", mandatory: false, exec: execer},
		criCtlChecker,
		ExtraArgsCheck{
			APIServerExtraArgs:         cfg.APIServerExtraArgs,
			ControllerManagerExtraArgs: cfg.ControllerManagerExtraArgs,
			SchedulerExtraArgs:         cfg.SchedulerExtraArgs,
		},
		HTTPProxyCheck{Proto: "https", Host: cfg.API.AdvertiseAddress},
		HTTPProxyCIDRCheck{Proto: "https", CIDR: cfg.Networking.ServiceSubnet},
		HTTPProxyCIDRCheck{Proto: "https", CIDR: cfg.Networking.PodSubnet},
	}

	if useCRI {
		checks = append(checks, CRICheck{socket: criSocket, exec: execer})
	} else {
		// assume docker
		checks = append(checks, ServiceCheck{Service: "docker", CheckIfActive: true})
	}

	if len(cfg.Etcd.Endpoints) == 0 {
		// Only do etcd related checks when no external endpoints were specified
		checks = append(checks,
			PortOpenCheck{port: 2379},
			DirAvailableCheck{Path: cfg.Etcd.DataDir},
		)
	} else {
		// Only check etcd version when external endpoints are specified
		if cfg.Etcd.CAFile != "" {
			checks = append(checks, FileExistingCheck{Path: cfg.Etcd.CAFile})
		}
		if cfg.Etcd.CertFile != "" {
			checks = append(checks, FileExistingCheck{Path: cfg.Etcd.CertFile})
		}
		if cfg.Etcd.KeyFile != "" {
			checks = append(checks, FileExistingCheck{Path: cfg.Etcd.KeyFile})
		}
		checks = append(checks,
			ExternalEtcdVersionCheck{Etcd: cfg.Etcd},
		)
	}

	// Check the config for authorization mode
	for _, authzMode := range cfg.AuthorizationModes {
		switch authzMode {
		case authzmodes.ModeABAC:
			checks = append(checks, FileExistingCheck{Path: kubeadmconstants.AuthorizationPolicyPath})
		case authzmodes.ModeWebhook:
			checks = append(checks, FileExistingCheck{Path: kubeadmconstants.AuthorizationWebhookConfigPath})
		}
	}

	if ip := net.ParseIP(cfg.API.AdvertiseAddress); ip != nil {
		if ip.To4() == nil && ip.To16() != nil {
			checks = append(checks,
				FileContentCheck{Path: bridgenf6, Content: []byte{'1'}},
			)
		}
	}
	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}

// RunJoinNodeChecks executes all individual, applicable to node checks.
func RunJoinNodeChecks(execer utilsexec.Interface, cfg *kubeadmapi.NodeConfiguration, criSocket string, ignorePreflightErrors sets.String) error {
	// First, check if we're root separately from the other preflight checks and fail fast
	if err := RunRootCheckOnly(ignorePreflightErrors); err != nil {
		return err
	}

	// check if we can use crictl to perform checks via the CRI
	criCtlChecker := InPathCheck{
		executable: "crictl",
		mandatory:  false,
		exec:       execer,
		suggestion: fmt.Sprintf("go get %v", kubeadmconstants.CRICtlPackage),
	}
	warns, _ := criCtlChecker.Check()
	useCRI := len(warns) == 0

	checks := []Checker{
		SystemVerificationCheck{CRISocket: criSocket},
		IsPrivilegedUserCheck{},
		HostnameCheck{cfg.NodeName},
		KubeletVersionCheck{},
		ServiceCheck{Service: "kubelet", CheckIfActive: false},
		PortOpenCheck{port: 10250},
		DirAvailableCheck{Path: filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.ManifestsSubDirName)},
		FileAvailableCheck{Path: cfg.CACertPath},
		FileAvailableCheck{Path: filepath.Join(kubeadmconstants.KubernetesDir, kubeadmconstants.KubeletKubeConfigFileName)},
	}
	if useCRI {
		checks = append(checks, CRICheck{socket: criSocket, exec: execer})
	} else {
		// assume docker
		checks = append(checks, ServiceCheck{Service: "docker", CheckIfActive: true})
	}
	//non-windows checks
	if runtime.GOOS == "linux" {
		checks = append(checks,
			FileContentCheck{Path: bridgenf, Content: []byte{'1'}},
			SwapCheck{},
			InPathCheck{executable: "ip", mandatory: true, exec: execer},
			InPathCheck{executable: "iptables", mandatory: true, exec: execer},
			InPathCheck{executable: "mount", mandatory: true, exec: execer},
			InPathCheck{executable: "nsenter", mandatory: true, exec: execer},
			InPathCheck{executable: "ebtables", mandatory: false, exec: execer},
			InPathCheck{executable: "ethtool", mandatory: false, exec: execer},
			InPathCheck{executable: "socat", mandatory: false, exec: execer},
			InPathCheck{executable: "tc", mandatory: false, exec: execer},
			InPathCheck{executable: "touch", mandatory: false, exec: execer},
			criCtlChecker)
	}

	if len(cfg.DiscoveryTokenAPIServers) > 0 {
		if ip := net.ParseIP(cfg.DiscoveryTokenAPIServers[0]); ip != nil {
			if ip.To4() == nil && ip.To16() != nil {
				checks = append(checks,
					FileContentCheck{Path: bridgenf6, Content: []byte{'1'}},
				)
			}
		}
	}
	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}

// RunRootCheckOnly initializes checks slice of structs and call RunChecks
func RunRootCheckOnly(ignorePreflightErrors sets.String) error {
	checks := []Checker{
		IsPrivilegedUserCheck{},
	}

	return RunChecks(checks, os.Stderr, ignorePreflightErrors)
}

// RunChecks runs each check, displays it's warnings/errors, and once all
// are processed will exit if any errors occurred.
func RunChecks(checks []Checker, ww io.Writer, ignorePreflightErrors sets.String) error {
	type checkErrors struct {
		Name   string
		Errors []error
	}
	found := []checkErrors{}

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
		if len(errs) > 0 {
			found = append(found, checkErrors{Name: name, Errors: errs})
		}
	}
	if len(found) > 0 {
		var errs bytes.Buffer
		for _, c := range found {
			for _, i := range c.Errors {
				errs.WriteString(fmt.Sprintf("\t[ERROR %s]: %v\n", c.Name, i.Error()))
			}
		}
		return &Error{Msg: errs.String()}
	}
	return nil
}

// TryStartKubelet attempts to bring up kubelet service
func TryStartKubelet(ignorePreflightErrors sets.String) {
	if setHasItemOrAll(ignorePreflightErrors, "StartKubelet") {
		return
	}
	// If we notice that the kubelet service is inactive, try to start it
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		fmt.Println("[preflight] No supported init system detected, won't ensure kubelet is running.")
	} else if initSystem.ServiceExists("kubelet") && !initSystem.ServiceIsActive("kubelet") {

		fmt.Println("[preflight] Starting the kubelet service")
		if err := initSystem.ServiceStart("kubelet"); err != nil {
			fmt.Printf("[preflight] WARNING: Unable to start the kubelet service: [%v]\n", err)
			fmt.Println("[preflight] WARNING: Please ensure kubelet is running manually.")
		}
	}
}

// setHasItemOrAll is helper function that return true if item is present in the set (case insensitive) or special key 'all' is present
func setHasItemOrAll(s sets.String, item string) bool {
	if s.Has("all") || s.Has(strings.ToLower(item)) {
		return true
	}
	return false
}
