/*
Copyright 2015 The Kubernetes Authors.

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

package nftables

//
// NOTE: this needs to be tested in e2e since it uses nftables for everything.
//

import (
	"context"
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/danwinship/knftables"

	v1 "k8s.io/api/core/v1"
	discovery "k8s.io/api/discovery/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/events"
	utilsysctl "k8s.io/component-helpers/node/util/sysctl"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/proxy"
	"k8s.io/kubernetes/pkg/proxy/conntrack"
	"k8s.io/kubernetes/pkg/proxy/healthcheck"
	"k8s.io/kubernetes/pkg/proxy/metaproxier"
	"k8s.io/kubernetes/pkg/proxy/metrics"
	proxyutil "k8s.io/kubernetes/pkg/proxy/util"
	proxyutiliptables "k8s.io/kubernetes/pkg/proxy/util/iptables"
	"k8s.io/kubernetes/pkg/util/async"
	utiliptables "k8s.io/kubernetes/pkg/util/iptables"
	utiliptablestesting "k8s.io/kubernetes/pkg/util/iptables/testing"
	utilexec "k8s.io/utils/exec"
	netutils "k8s.io/utils/net"
	"k8s.io/utils/ptr"
)

const (
	// Our nftables table. All of our chains/sets/maps are created inside this table,
	// so they don't need any "kube-" or "kube-proxy-" prefix of their own.
	kubeProxyTable = "kube-proxy"

	// service dispatch
	kubeServicesChain  = "services"
	kubeNodePortsChain = "nodeports"

	// handling for services with no endpoints
	kubeServicesFilterChain   = "services-filter"
	kubeExternalServicesChain = "external-services"

	// LoadBalancerSourceRanges handling
	kubeFirewallChain = "firewall"

	// masquerading
	kubeMarkMasqChain     = "mark-for-masquerade"
	kubeMasqueradingChain = "masquerading"

	// chain for special filtering rules
	kubeForwardChain = "forward"
)

const sysctlNFConntrackTCPBeLiberal = "net/netfilter/nf_conntrack_tcp_be_liberal"

// internal struct for string service information
type servicePortInfo struct {
	*proxy.BaseServicePortInfo
	// The following fields are computed and stored for performance reasons.
	nameString             string
	clusterPolicyChainName string
	localPolicyChainName   string
	firewallChainName      string
	externalChainName      string
}

// returns a new proxy.ServicePort which abstracts a serviceInfo
func newServiceInfo(port *v1.ServicePort, service *v1.Service, bsvcPortInfo *proxy.BaseServicePortInfo) proxy.ServicePort {
	svcPort := &servicePortInfo{BaseServicePortInfo: bsvcPortInfo}

	// Store the following for performance reasons.
	svcName := types.NamespacedName{Namespace: service.Namespace, Name: service.Name}
	svcPortName := proxy.ServicePortName{NamespacedName: svcName, Port: port.Name}
	svcPort.nameString = svcPortName.String()

	chainNameBase := servicePortChainNameBase(&svcPortName, strings.ToLower(string(svcPort.Protocol())))
	svcPort.clusterPolicyChainName = servicePortPolicyClusterChainNamePrefix + chainNameBase
	svcPort.localPolicyChainName = servicePortPolicyLocalChainNamePrefix + chainNameBase
	svcPort.firewallChainName = serviceFirewallChainNamePrefix + chainNameBase
	svcPort.externalChainName = serviceExternalChainNamePrefix + chainNameBase

	return svcPort
}

// internal struct for endpoints information
type endpointInfo struct {
	*proxy.BaseEndpointInfo

	chainName string
}

// returns a new proxy.Endpoint which abstracts a endpointInfo
func newEndpointInfo(baseInfo *proxy.BaseEndpointInfo, svcPortName *proxy.ServicePortName) proxy.Endpoint {
	return &endpointInfo{
		BaseEndpointInfo: baseInfo,
		chainName:        servicePortEndpointChainNamePrefix + servicePortEndpointChainNameBase(svcPortName, strings.ToLower(string(svcPortName.Protocol)), baseInfo.String()),
	}
}

// Proxier is an nftables based proxy
type Proxier struct {
	// ipFamily defines the IP family which this proxier is tracking.
	ipFamily v1.IPFamily

	// endpointsChanges and serviceChanges contains all changes to endpoints and
	// services that happened since nftables was synced. For a single object,
	// changes are accumulated, i.e. previous is state from before all of them,
	// current is state after applying all of those.
	endpointsChanges *proxy.EndpointsChangeTracker
	serviceChanges   *proxy.ServiceChangeTracker

	mu           sync.Mutex // protects the following fields
	svcPortMap   proxy.ServicePortMap
	endpointsMap proxy.EndpointsMap
	nodeLabels   map[string]string
	// endpointSlicesSynced, and servicesSynced are set to true
	// when corresponding objects are synced after startup. This is used to avoid
	// updating nftables with some partial data after kube-proxy restart.
	endpointSlicesSynced bool
	servicesSynced       bool
	initialized          int32
	syncRunner           *async.BoundedFrequencyRunner // governs calls to syncProxyRules
	syncPeriod           time.Duration

	// These are effectively const and do not need the mutex to be held.
	iptables       utiliptables.Interface
	nftables       knftables.Interface
	masqueradeAll  bool
	masqueradeMark string
	exec           utilexec.Interface
	localDetector  proxyutiliptables.LocalTrafficDetector
	hostname       string
	nodeIP         net.IP
	recorder       events.EventRecorder

	serviceHealthServer healthcheck.ServiceHealthServer
	healthzServer       *healthcheck.ProxierHealthServer

	// Since converting probabilities (floats) to strings is expensive
	// and we are using only probabilities in the format of 1/n, we are
	// precomputing some number of those and cache for future reuse.
	precomputedProbabilities []string

	// The following buffers are used to reuse memory and avoid allocations
	// that are significantly impacting performance.
	natRules proxyutil.LineBuffer

	// conntrackTCPLiberal indicates whether the system sets the kernel nf_conntrack_tcp_be_liberal
	conntrackTCPLiberal bool

	// nodePortAddresses selects the interfaces where nodePort works.
	nodePortAddresses *proxyutil.NodePortAddresses
	// networkInterfacer defines an interface for several net library functions.
	// Inject for test purpose.
	networkInterfacer proxyutil.NetworkInterfacer

	// staleChains contains information about chains to be deleted later
	staleChains map[string]time.Time
}

// Proxier implements proxy.Provider
var _ proxy.Provider = &Proxier{}

// NewProxier returns a new nftables Proxier. Once a proxier is created, it will keep
// nftables up to date in the background and will not terminate if a particular nftables
// call fails.
func NewProxier(ipFamily v1.IPFamily,
	sysctl utilsysctl.Interface,
	syncPeriod time.Duration,
	minSyncPeriod time.Duration,
	masqueradeAll bool,
	masqueradeBit int,
	localDetector proxyutiliptables.LocalTrafficDetector,
	hostname string,
	nodeIP net.IP,
	recorder events.EventRecorder,
	healthzServer *healthcheck.ProxierHealthServer,
	nodePortAddressStrings []string,
	initOnly bool,
) (*Proxier, error) {
	nodePortAddresses := proxyutil.NewNodePortAddresses(ipFamily, nodePortAddressStrings)

	// Be conservative in what you do, be liberal in what you accept from others.
	// If it's non-zero, we mark only out of window RST segments as INVALID.
	// Ref: https://docs.kernel.org/networking/nf_conntrack-sysctl.html
	conntrackTCPLiberal := false
	if val, err := sysctl.GetSysctl(sysctlNFConntrackTCPBeLiberal); err == nil && val != 0 {
		conntrackTCPLiberal = true
		klog.InfoS("nf_conntrack_tcp_be_liberal set, not installing DROP rules for INVALID packets")
	}

	if initOnly {
		klog.InfoS("System initialized and --init-only specified")
		return nil, nil
	}

	// Generate the masquerade mark to use for SNAT rules.
	masqueradeValue := 1 << uint(masqueradeBit)
	masqueradeMark := fmt.Sprintf("%#08x", masqueradeValue)
	klog.V(2).InfoS("Using nftables mark for masquerade", "ipFamily", ipFamily, "mark", masqueradeMark)

	serviceHealthServer := healthcheck.NewServiceHealthServer(hostname, recorder, nodePortAddresses, healthzServer)

	var nftablesFamily knftables.Family
	if ipFamily == v1.IPv4Protocol {
		nftablesFamily = knftables.IPv4Family
	} else {
		nftablesFamily = knftables.IPv6Family
	}
	nft, err := knftables.New(nftablesFamily, kubeProxyTable)
	if err != nil {
		return nil, err
	}

	proxier := &Proxier{
		ipFamily:                 ipFamily,
		svcPortMap:               make(proxy.ServicePortMap),
		serviceChanges:           proxy.NewServiceChangeTracker(newServiceInfo, ipFamily, recorder, nil),
		endpointsMap:             make(proxy.EndpointsMap),
		endpointsChanges:         proxy.NewEndpointsChangeTracker(hostname, newEndpointInfo, ipFamily, recorder, nil),
		syncPeriod:               syncPeriod,
		iptables:                 utiliptablestesting.NewFake(),
		nftables:                 nft,
		masqueradeAll:            masqueradeAll,
		masqueradeMark:           masqueradeMark,
		exec:                     utilexec.New(),
		localDetector:            localDetector,
		hostname:                 hostname,
		nodeIP:                   nodeIP,
		recorder:                 recorder,
		serviceHealthServer:      serviceHealthServer,
		healthzServer:            healthzServer,
		precomputedProbabilities: make([]string, 0, 1001),
		natRules:                 proxyutil.NewLineBuffer(),
		nodePortAddresses:        nodePortAddresses,
		networkInterfacer:        proxyutil.RealNetwork{},
		conntrackTCPLiberal:      conntrackTCPLiberal,
		staleChains:              make(map[string]time.Time),
	}

	burstSyncs := 2
	klog.V(2).InfoS("NFTables sync params", "ipFamily", ipFamily, "minSyncPeriod", minSyncPeriod, "syncPeriod", syncPeriod, "burstSyncs", burstSyncs)
	proxier.syncRunner = async.NewBoundedFrequencyRunner("sync-runner", proxier.syncProxyRules, minSyncPeriod, syncPeriod, burstSyncs)

	return proxier, nil
}

// NewDualStackProxier creates a MetaProxier instance, with IPv4 and IPv6 proxies.
func NewDualStackProxier(
	sysctl utilsysctl.Interface,
	syncPeriod time.Duration,
	minSyncPeriod time.Duration,
	masqueradeAll bool,
	masqueradeBit int,
	localDetectors [2]proxyutiliptables.LocalTrafficDetector,
	hostname string,
	nodeIPs map[v1.IPFamily]net.IP,
	recorder events.EventRecorder,
	healthzServer *healthcheck.ProxierHealthServer,
	nodePortAddresses []string,
	initOnly bool,
) (proxy.Provider, error) {
	// Create an ipv4 instance of the single-stack proxier
	ipv4Proxier, err := NewProxier(v1.IPv4Protocol, sysctl,
		syncPeriod, minSyncPeriod, masqueradeAll, masqueradeBit, localDetectors[0], hostname,
		nodeIPs[v1.IPv4Protocol], recorder, healthzServer, nodePortAddresses, initOnly)
	if err != nil {
		return nil, fmt.Errorf("unable to create ipv4 proxier: %v", err)
	}

	ipv6Proxier, err := NewProxier(v1.IPv6Protocol, sysctl,
		syncPeriod, minSyncPeriod, masqueradeAll, masqueradeBit, localDetectors[1], hostname,
		nodeIPs[v1.IPv6Protocol], recorder, healthzServer, nodePortAddresses, initOnly)
	if err != nil {
		return nil, fmt.Errorf("unable to create ipv6 proxier: %v", err)
	}
	if initOnly {
		return nil, nil
	}
	return metaproxier.NewMetaProxier(ipv4Proxier, ipv6Proxier), nil
}

// nftablesBaseChains lists our "base chains"; those that are directly connected to the
// netfilter hooks (e.g., "postrouting", "input", etc.), as opposed to "regular" chains,
// which are only run when a rule jumps to them. See
// https://wiki.nftables.org/wiki-nftables/index.php/Configuring_chains.
//
// These are set up from setupNFTables() and then not directly referenced by
// syncProxyRules().
//
// All of our base chains have names that are just "${type}-${hook}". e.g., "nat-prerouting".
type nftablesBaseChain struct {
	name      string
	chainType knftables.BaseChainType
	hook      knftables.BaseChainHook
	priority  knftables.BaseChainPriority
}

var nftablesBaseChains = []nftablesBaseChain{
	{"filter-input", knftables.FilterType, knftables.InputHook, knftables.FilterPriority},
	{"filter-forward", knftables.FilterType, knftables.ForwardHook, knftables.FilterPriority},
	{"filter-output", knftables.FilterType, knftables.OutputHook, knftables.FilterPriority},
	{"nat-prerouting", knftables.NATType, knftables.PreroutingHook, knftables.DNATPriority},
	{"nat-output", knftables.NATType, knftables.OutputHook, knftables.DNATPriority},
	{"nat-postrouting", knftables.NATType, knftables.PostroutingHook, knftables.SNATPriority},
}

// nftablesJumpChains lists our top-level "regular chains" that are jumped to directly
// from one of the base chains. These are set up from setupNFTables(), and some of them
// are also referenced in syncProxyRules().
type nftablesJumpChain struct {
	dstChain  string
	srcChain  string
	extraArgs string
}

var nftablesJumpChains = []nftablesJumpChain{
	{kubeExternalServicesChain, "filter-input", "ct state new"},
	{kubeExternalServicesChain, "filter-forward", "ct state new"},
	{kubeServicesFilterChain, "filter-forward", "ct state new"},
	{kubeServicesFilterChain, "filter-output", "ct state new"},
	{kubeForwardChain, "filter-forward", ""},
	{kubeFirewallChain, "filter-input", "ct state new"},
	{kubeFirewallChain, "filter-output", "ct state new"},
	{kubeFirewallChain, "filter-forward", "ct state new"},

	{kubeServicesChain, "nat-output", ""},
	{kubeServicesChain, "nat-prerouting", ""},
	{kubeMasqueradingChain, "nat-postrouting", ""},
}

// ensureChain adds commands to tx to ensure that chain exists and doesn't contain
// anything from before this transaction (using createdChains to ensure that we don't
// Flush a chain more than once and lose *new* rules as well.)
func ensureChain(chain string, tx *knftables.Transaction, createdChains sets.Set[string]) {
	if createdChains.Has(chain) {
		return
	}
	tx.Add(&knftables.Chain{
		Name: chain,
	})
	tx.Flush(&knftables.Chain{
		Name: chain,
	})
	createdChains.Insert(chain)
}

func (proxier *Proxier) setupNFTables(tx *knftables.Transaction) {
	tx.Add(&knftables.Table{
		Comment: ptr.To("rules for kube-proxy"),
	})

	// Create and flush base chains
	for _, bc := range nftablesBaseChains {
		chain := &knftables.Chain{
			Name:     bc.name,
			Type:     ptr.To(bc.chainType),
			Hook:     ptr.To(bc.hook),
			Priority: ptr.To(bc.priority),
		}
		tx.Add(chain)
		tx.Flush(chain)
	}

	// Create and flush ordinary chains and add rules jumping to them
	createdChains := sets.New[string]()
	for _, c := range nftablesJumpChains {
		ensureChain(c.dstChain, tx, createdChains)
		tx.Add(&knftables.Rule{
			Chain: c.srcChain,
			Rule: knftables.Concat(
				c.extraArgs,
				"jump", c.dstChain,
			),
		})
	}

	// Ensure all of our other "top-level" chains exist
	for _, chain := range []string{kubeServicesFilterChain, kubeServicesChain, kubeExternalServicesChain, kubeForwardChain, kubeNodePortsChain, kubeFirewallChain, kubeMasqueradingChain, kubeMarkMasqChain} {
		ensureChain(chain, tx, createdChains)
	}

	// Add the rules in the mark-for-masquerade and masquerading chains
	tx.Add(&knftables.Rule{
		Chain: kubeMarkMasqChain,
		Rule: knftables.Concat(
			"mark", "set", "mark", "or", proxier.masqueradeMark,
		),
	})

	tx.Add(&knftables.Rule{
		Chain: kubeMasqueradingChain,
		Rule: knftables.Concat(
			"mark", "and", proxier.masqueradeMark, "==", "0",
			"return",
		),
	})
	tx.Add(&knftables.Rule{
		Chain: kubeMasqueradingChain,
		Rule: knftables.Concat(
			"mark", "set", "mark", "xor", proxier.masqueradeMark,
		),
	})
	tx.Add(&knftables.Rule{
		Chain: kubeMasqueradingChain,
		Rule:  "masquerade fully-random",
	})

	// Drop the packets in INVALID state, which would potentially cause
	// unexpected connection reset if nf_conntrack_tcp_be_liberal is not set.
	// Ref: https://github.com/kubernetes/kubernetes/issues/74839
	// Ref: https://github.com/kubernetes/kubernetes/issues/117924
	if !proxier.conntrackTCPLiberal {
		tx.Add(&knftables.Rule{
			Chain: kubeForwardChain,
			Rule:  "ct state invalid drop",
		})
	}
}

// CleanupLeftovers removes all nftables rules and chains created by the Proxier
// It returns true if an error was encountered. Errors are logged.
func CleanupLeftovers() bool {
	var encounteredError bool

	for _, family := range []knftables.Family{knftables.IPv4Family, knftables.IPv6Family} {
		nft, err := knftables.New(family, kubeProxyTable)
		if err == nil {
			tx := nft.NewTransaction()
			tx.Delete(&knftables.Table{})
			err = nft.Run(context.TODO(), tx)
		}
		if err != nil && !knftables.IsNotFound(err) {
			klog.ErrorS(err, "Error cleaning up nftables rules")
			encounteredError = true
		}
	}

	return encounteredError
}

func computeProbability(n int) string {
	return fmt.Sprintf("%0.10f", 1.0/float64(n))
}

// This assumes proxier.mu is held
func (proxier *Proxier) precomputeProbabilities(numberOfPrecomputed int) {
	if len(proxier.precomputedProbabilities) == 0 {
		proxier.precomputedProbabilities = append(proxier.precomputedProbabilities, "<bad value>")
	}
	for i := len(proxier.precomputedProbabilities); i <= numberOfPrecomputed; i++ {
		proxier.precomputedProbabilities = append(proxier.precomputedProbabilities, computeProbability(i))
	}
}

// This assumes proxier.mu is held
func (proxier *Proxier) probability(n int) string {
	if n >= len(proxier.precomputedProbabilities) {
		proxier.precomputeProbabilities(n)
	}
	return proxier.precomputedProbabilities[n]
}

// Sync is called to synchronize the proxier state to nftables as soon as possible.
func (proxier *Proxier) Sync() {
	if proxier.healthzServer != nil {
		proxier.healthzServer.QueuedUpdate(proxier.ipFamily)
	}
	metrics.SyncProxyRulesLastQueuedTimestamp.SetToCurrentTime()
	proxier.syncRunner.Run()
}

// SyncLoop runs periodic work.  This is expected to run as a goroutine or as the main loop of the app.  It does not return.
func (proxier *Proxier) SyncLoop() {
	// Update healthz timestamp at beginning in case Sync() never succeeds.
	if proxier.healthzServer != nil {
		proxier.healthzServer.Updated(proxier.ipFamily)
	}

	// synthesize "last change queued" time as the informers are syncing.
	metrics.SyncProxyRulesLastQueuedTimestamp.SetToCurrentTime()
	proxier.syncRunner.Loop(wait.NeverStop)
}

func (proxier *Proxier) setInitialized(value bool) {
	var initialized int32
	if value {
		initialized = 1
	}
	atomic.StoreInt32(&proxier.initialized, initialized)
}

func (proxier *Proxier) isInitialized() bool {
	return atomic.LoadInt32(&proxier.initialized) > 0
}

// OnServiceAdd is called whenever creation of new service object
// is observed.
func (proxier *Proxier) OnServiceAdd(service *v1.Service) {
	proxier.OnServiceUpdate(nil, service)
}

// OnServiceUpdate is called whenever modification of an existing
// service object is observed.
func (proxier *Proxier) OnServiceUpdate(oldService, service *v1.Service) {
	if proxier.serviceChanges.Update(oldService, service) && proxier.isInitialized() {
		proxier.Sync()
	}
}

// OnServiceDelete is called whenever deletion of an existing service
// object is observed.
func (proxier *Proxier) OnServiceDelete(service *v1.Service) {
	proxier.OnServiceUpdate(service, nil)

}

// OnServiceSynced is called once all the initial event handlers were
// called and the state is fully propagated to local cache.
func (proxier *Proxier) OnServiceSynced() {
	proxier.mu.Lock()
	proxier.servicesSynced = true
	proxier.setInitialized(proxier.endpointSlicesSynced)
	proxier.mu.Unlock()

	// Sync unconditionally - this is called once per lifetime.
	proxier.syncProxyRules()
}

// OnEndpointSliceAdd is called whenever creation of a new endpoint slice object
// is observed.
func (proxier *Proxier) OnEndpointSliceAdd(endpointSlice *discovery.EndpointSlice) {
	if proxier.endpointsChanges.EndpointSliceUpdate(endpointSlice, false) && proxier.isInitialized() {
		proxier.Sync()
	}
}

// OnEndpointSliceUpdate is called whenever modification of an existing endpoint
// slice object is observed.
func (proxier *Proxier) OnEndpointSliceUpdate(_, endpointSlice *discovery.EndpointSlice) {
	if proxier.endpointsChanges.EndpointSliceUpdate(endpointSlice, false) && proxier.isInitialized() {
		proxier.Sync()
	}
}

// OnEndpointSliceDelete is called whenever deletion of an existing endpoint slice
// object is observed.
func (proxier *Proxier) OnEndpointSliceDelete(endpointSlice *discovery.EndpointSlice) {
	if proxier.endpointsChanges.EndpointSliceUpdate(endpointSlice, true) && proxier.isInitialized() {
		proxier.Sync()
	}
}

// OnEndpointSlicesSynced is called once all the initial event handlers were
// called and the state is fully propagated to local cache.
func (proxier *Proxier) OnEndpointSlicesSynced() {
	proxier.mu.Lock()
	proxier.endpointSlicesSynced = true
	proxier.setInitialized(proxier.servicesSynced)
	proxier.mu.Unlock()

	// Sync unconditionally - this is called once per lifetime.
	proxier.syncProxyRules()
}

// OnNodeAdd is called whenever creation of new node object
// is observed.
func (proxier *Proxier) OnNodeAdd(node *v1.Node) {
	if node.Name != proxier.hostname {
		klog.ErrorS(nil, "Received a watch event for a node that doesn't match the current node",
			"eventNode", node.Name, "currentNode", proxier.hostname)
		return
	}

	if reflect.DeepEqual(proxier.nodeLabels, node.Labels) {
		return
	}

	proxier.mu.Lock()
	proxier.nodeLabels = map[string]string{}
	for k, v := range node.Labels {
		proxier.nodeLabels[k] = v
	}
	proxier.mu.Unlock()
	klog.V(4).InfoS("Updated proxier node labels", "labels", node.Labels)

	proxier.Sync()
}

// OnNodeUpdate is called whenever modification of an existing
// node object is observed.
func (proxier *Proxier) OnNodeUpdate(oldNode, node *v1.Node) {
	if node.Name != proxier.hostname {
		klog.ErrorS(nil, "Received a watch event for a node that doesn't match the current node",
			"eventNode", node.Name, "currentNode", proxier.hostname)
		return
	}

	if reflect.DeepEqual(proxier.nodeLabels, node.Labels) {
		return
	}

	proxier.mu.Lock()
	proxier.nodeLabels = map[string]string{}
	for k, v := range node.Labels {
		proxier.nodeLabels[k] = v
	}
	proxier.mu.Unlock()
	klog.V(4).InfoS("Updated proxier node labels", "labels", node.Labels)

	proxier.Sync()
}

// OnNodeDelete is called whenever deletion of an existing node
// object is observed.
func (proxier *Proxier) OnNodeDelete(node *v1.Node) {
	if node.Name != proxier.hostname {
		klog.ErrorS(nil, "Received a watch event for a node that doesn't match the current node",
			"eventNode", node.Name, "currentNode", proxier.hostname)
		return
	}

	proxier.mu.Lock()
	proxier.nodeLabels = nil
	proxier.mu.Unlock()

	proxier.Sync()
}

// OnNodeSynced is called once all the initial event handlers were
// called and the state is fully propagated to local cache.
func (proxier *Proxier) OnNodeSynced() {
}

const (
	// Maximum length for one of our chain name prefixes, including the trailing
	// hyphen.
	chainNamePrefixLengthMax = 16

	// Maximum length of the string returned from servicePortChainNameBase or
	// servicePortEndpointChainNameBase.
	chainNameBaseLengthMax = knftables.NameLengthMax - chainNamePrefixLengthMax
)

const (
	servicePortPolicyClusterChainNamePrefix = "service-"
	servicePortPolicyLocalChainNamePrefix   = "local-"
	serviceFirewallChainNamePrefix          = "firewall-"
	serviceExternalChainNamePrefix          = "external-"
	servicePortEndpointChainNamePrefix      = "endpoint-"
)

// hashAndTruncate prefixes name with a hash of itself and then truncates to
// chainNameBaseLengthMax. The hash ensures that (a) the name is still unique if we have
// to truncate the end, and (b) it's visually distinguishable from other chains that would
// otherwise have nearly identical names (e.g., different endpoint chains for a given
// service that differ in only a single digit).
func hashAndTruncate(name string) string {
	hash := sha256.Sum256([]byte(name))
	encoded := base32.StdEncoding.EncodeToString(hash[:])
	name = encoded[:8] + "-" + name
	if len(name) > chainNameBaseLengthMax {
		name = name[:chainNameBaseLengthMax-3] + "..."
	}
	return name
}

// servicePortChainNameBase returns the base name for a chain for the given ServicePort.
// This is something like "HASH-namespace/serviceName/protocol/portName", e.g,
// "ULMVA6XW-ns1/svc1/tcp/p80".
func servicePortChainNameBase(servicePortName *proxy.ServicePortName, protocol string) string {
	// nftables chains can contain the characters [A-Za-z0-9_./-] (but must start with
	// a letter, underscore, or dot).
	//
	// Namespace, Service, and Port names can contain [a-z0-9-] (with some additional
	// restrictions that aren't relevant here).
	//
	// Protocol is /(tcp|udp|sctp)/.
	//
	// Thus, we can safely use all Namespace names, Service names, protocol values,
	// and Port names directly in nftables chain names (though note that this assumes
	// that the chain name won't *start* with any of those strings, since that might
	// be illegal). We use "/" to separate the parts of the name, which is one of the
	// two characters allowed in a chain name that isn't allowed in our input strings.

	name := fmt.Sprintf("%s/%s/%s/%s",
		servicePortName.NamespacedName.Namespace,
		servicePortName.NamespacedName.Name,
		protocol,
		servicePortName.Port,
	)

	// The namespace, service, and port name can each be up to 63 characters, protocol
	// can be up to 4, plus 8 for the hash and 4 additional punctuation characters.
	// That's a total of 205, which is less than chainNameBaseLengthMax (240). So this
	// will never actually return a truncated name.
	return hashAndTruncate(name)
}

// servicePortEndpointChainNameBase returns the suffix for chain names for the given
// endpoint. This is something like
// "HASH-namespace/serviceName/protocol/portName__endpointIP/endpointport", e.g.,
// "5OJB2KTY-ns1/svc1/tcp/p80__10.180.0.1/80".
func servicePortEndpointChainNameBase(servicePortName *proxy.ServicePortName, protocol, endpoint string) string {
	// As above in servicePortChainNameBase: Namespace, Service, Port, Protocol, and
	// EndpointPort are all safe to copy into the chain name directly. But if
	// EndpointIP is IPv6 then it will contain colons, which aren't allowed in a chain
	// name. IPv6 IPs are also quite long, but we can't safely truncate them (e.g. to
	// only the final segment) because (especially for manually-created external
	// endpoints), we can't know for sure that any part of them is redundant.

	endpointIP, endpointPort, _ := net.SplitHostPort(endpoint)
	if strings.Contains(endpointIP, ":") {
		endpointIP = strings.ReplaceAll(endpointIP, ":", ".")
	}

	// As above, we use "/" to separate parts of the name, and "__" to separate the
	// "service" part from the "endpoint" part.
	name := fmt.Sprintf("%s/%s/%s/%s__%s/%s",
		servicePortName.NamespacedName.Namespace,
		servicePortName.NamespacedName.Name,
		protocol,
		servicePortName.Port,
		endpointIP,
		endpointPort,
	)

	// The part of name before the "__" can be up to 205 characters (as with
	// servicePortChainNameBase above). An IPv6 address can be up to 39 characters, and
	// a port can be up to 5 digits, plus 3 punctuation characters gives a max total
	// length of 252, well over chainNameBaseLengthMax (240), so truncation is
	// theoretically possible (though incredibly unlikely).
	return hashAndTruncate(name)
}

func isServiceChainName(chainString string) bool {
	// The chains returned from servicePortChainNameBase and
	// servicePortEndpointChainNameBase will always have at least one "/" in them.
	// Since none of our "stock" chain names use slashes, we can distinguish them this
	// way.
	return strings.Contains(chainString, "/")
}

// This is where all of the nftables calls happen.
// This assumes proxier.mu is NOT held
func (proxier *Proxier) syncProxyRules() {
	proxier.mu.Lock()
	defer proxier.mu.Unlock()

	// don't sync rules till we've received services and endpoints
	if !proxier.isInitialized() {
		klog.V(2).InfoS("Not syncing nftables until Services and Endpoints have been received from master")
		return
	}

	//
	// Below this point we will not return until we try to write the nftables rules.
	//

	// Keep track of how long syncs take.
	start := time.Now()
	defer func() {
		metrics.SyncProxyRulesLatency.Observe(metrics.SinceInSeconds(start))
		klog.V(2).InfoS("SyncProxyRules complete", "elapsed", time.Since(start))
	}()

	serviceUpdateResult := proxier.svcPortMap.Update(proxier.serviceChanges)
	endpointUpdateResult := proxier.endpointsMap.Update(proxier.endpointsChanges)

	klog.V(2).InfoS("Syncing nftables rules")

	success := false
	defer func() {
		if !success {
			klog.InfoS("Sync failed", "retryingTime", proxier.syncPeriod)
			proxier.syncRunner.RetryAfter(proxier.syncPeriod)
		}
	}()

	// If there are sufficiently-stale chains left over from previous transactions,
	// try to delete them now.
	if len(proxier.staleChains) > 0 {
		oneSecondAgo := start.Add(-time.Second)
		tx := proxier.nftables.NewTransaction()
		deleted := 0
		for chain, modtime := range proxier.staleChains {
			if modtime.Before(oneSecondAgo) {
				tx.Delete(&knftables.Chain{
					Name: chain,
				})
				delete(proxier.staleChains, chain)
				deleted++
			}
		}
		if deleted > 0 {
			klog.InfoS("Deleting stale nftables chains", "numChains", deleted)
			err := proxier.nftables.Run(context.TODO(), tx)
			if err != nil {
				// We already deleted the entries from staleChains, but if
				// the chains still exist, they'll just get added back
				// (with a later timestamp) at the end of the sync.
				klog.ErrorS(err, "Unable to delete stale chains; will retry later")
				// FIXME: metric
			}
		}
	}

	// Now start the actual syncing transaction
	tx := proxier.nftables.NewTransaction()
	proxier.setupNFTables(tx)

	// We need to use, eg, "ip daddr" for IPv4 but "ip6 daddr" for IPv6
	ipX := "ip"
	if proxier.ipFamily == v1.IPv6Protocol {
		ipX = "ip6"
	}

	// Reset all buffers used later.
	// This is to avoid memory reallocations and thus improve performance.
	proxier.natRules.Reset()

	// Accumulate service/endpoint chains to keep.
	activeChains := sets.New[string]()

	// Compute total number of endpoint chains across all services
	// to get a sense of how big the cluster is.
	totalEndpoints := 0
	for svcName := range proxier.svcPortMap {
		totalEndpoints += len(proxier.endpointsMap[svcName])
	}

	// These two variables are used to publish the sync_proxy_rules_no_endpoints_total
	// metric.
	serviceNoLocalEndpointsTotalInternal := 0
	serviceNoLocalEndpointsTotalExternal := 0

	// Build rules for each service-port.
	for svcName, svc := range proxier.svcPortMap {
		svcInfo, ok := svc.(*servicePortInfo)
		if !ok {
			klog.ErrorS(nil, "Failed to cast serviceInfo", "serviceName", svcName)
			continue
		}
		protocol := strings.ToLower(string(svcInfo.Protocol()))
		svcPortNameString := svcInfo.nameString

		// Figure out the endpoints for Cluster and Local traffic policy.
		// allLocallyReachableEndpoints is the set of all endpoints that can be routed to
		// from this node, given the service's traffic policies. hasEndpoints is true
		// if the service has any usable endpoints on any node, not just this one.
		allEndpoints := proxier.endpointsMap[svcName]
		clusterEndpoints, localEndpoints, allLocallyReachableEndpoints, hasEndpoints := proxy.CategorizeEndpoints(allEndpoints, svcInfo, proxier.nodeLabels)

		// Note the endpoint chains that will be used
		for _, ep := range allLocallyReachableEndpoints {
			if epInfo, ok := ep.(*endpointInfo); ok {
				ensureChain(epInfo.chainName, tx, activeChains)
			}
		}

		// clusterPolicyChain contains the endpoints used with "Cluster" traffic policy
		clusterPolicyChain := svcInfo.clusterPolicyChainName
		usesClusterPolicyChain := len(clusterEndpoints) > 0 && svcInfo.UsesClusterEndpoints()
		if usesClusterPolicyChain {
			ensureChain(clusterPolicyChain, tx, activeChains)
		}

		// localPolicyChain contains the endpoints used with "Local" traffic policy
		localPolicyChain := svcInfo.localPolicyChainName
		usesLocalPolicyChain := len(localEndpoints) > 0 && svcInfo.UsesLocalEndpoints()
		if usesLocalPolicyChain {
			ensureChain(localPolicyChain, tx, activeChains)
		}

		// internalPolicyChain is the chain containing the endpoints for
		// "internal" (ClusterIP) traffic. internalTrafficChain is the chain that
		// internal traffic is routed to (which is always the same as
		// internalPolicyChain). hasInternalEndpoints is true if we should
		// generate rules pointing to internalTrafficChain, or false if there are
		// no available internal endpoints.
		internalPolicyChain := clusterPolicyChain
		hasInternalEndpoints := hasEndpoints
		if svcInfo.InternalPolicyLocal() {
			internalPolicyChain = localPolicyChain
			if len(localEndpoints) == 0 {
				hasInternalEndpoints = false
			}
		}
		internalTrafficChain := internalPolicyChain

		// Similarly, externalPolicyChain is the chain containing the endpoints
		// for "external" (NodePort, LoadBalancer, and ExternalIP) traffic.
		// externalTrafficChain is the chain that external traffic is routed to
		// (which is always the service's "EXT" chain). hasExternalEndpoints is
		// true if there are endpoints that will be reached by external traffic.
		// (But we may still have to generate externalTrafficChain even if there
		// are no external endpoints, to ensure that the short-circuit rules for
		// local traffic are set up.)
		externalPolicyChain := clusterPolicyChain
		hasExternalEndpoints := hasEndpoints
		if svcInfo.ExternalPolicyLocal() {
			externalPolicyChain = localPolicyChain
			if len(localEndpoints) == 0 {
				hasExternalEndpoints = false
			}
		}
		externalTrafficChain := svcInfo.externalChainName // eventually jumps to externalPolicyChain

		// usesExternalTrafficChain is based on hasEndpoints, not hasExternalEndpoints,
		// because we need the local-traffic-short-circuiting rules even when there
		// are no externally-usable endpoints.
		usesExternalTrafficChain := hasEndpoints && svcInfo.ExternallyAccessible()
		if usesExternalTrafficChain {
			ensureChain(externalTrafficChain, tx, activeChains)
		}

		// Traffic to LoadBalancer IPs can go directly to externalTrafficChain
		// unless LoadBalancerSourceRanges is in use in which case we will
		// create a firewall chain.
		loadBalancerTrafficChain := externalTrafficChain
		fwChain := svcInfo.firewallChainName
		usesFWChain := hasEndpoints && len(svcInfo.LoadBalancerVIPStrings()) > 0 && len(svcInfo.LoadBalancerSourceRanges()) > 0
		if usesFWChain {
			ensureChain(fwChain, tx, activeChains)
			loadBalancerTrafficChain = fwChain
		}

		var internalTrafficFilterVerdict, internalTrafficFilterComment string
		var externalTrafficFilterVerdict, externalTrafficFilterComment string
		if !hasEndpoints {
			// The service has no endpoints at all; hasInternalEndpoints and
			// hasExternalEndpoints will also be false, and we will not
			// generate any chains in the "nat" table for the service; only
			// rules in the "filter" table rejecting incoming packets for
			// the service's IPs.
			internalTrafficFilterVerdict = "reject"
			internalTrafficFilterComment = fmt.Sprintf("%s has no endpoints", svcPortNameString)
			externalTrafficFilterVerdict = "reject"
			externalTrafficFilterComment = internalTrafficFilterComment
		} else {
			if !hasInternalEndpoints {
				// The internalTrafficPolicy is "Local" but there are no local
				// endpoints. Traffic to the clusterIP will be dropped, but
				// external traffic may still be accepted.
				internalTrafficFilterVerdict = "drop"
				internalTrafficFilterComment = fmt.Sprintf("%s has no local endpoints", svcPortNameString)
				serviceNoLocalEndpointsTotalInternal++
			}
			if !hasExternalEndpoints {
				// The externalTrafficPolicy is "Local" but there are no
				// local endpoints. Traffic to "external" IPs from outside
				// the cluster will be dropped, but traffic from inside
				// the cluster may still be accepted.
				externalTrafficFilterVerdict = "drop"
				externalTrafficFilterComment = fmt.Sprintf("%s has no local endpoints", svcPortNameString)
				serviceNoLocalEndpointsTotalExternal++
			}
		}

		// Capture the clusterIP.
		if hasInternalEndpoints {
			proxier.natRules.Write(
				"-A", string(kubeServicesChain),
				"-m", "comment", "--comment", fmt.Sprintf(`"%s cluster IP"`, svcPortNameString),
				"-m", protocol, "-p", protocol,
				"-d", svcInfo.ClusterIP().String(),
				"--dport", strconv.Itoa(svcInfo.Port()),
				"-j", string(internalTrafficChain))
		} else {
			// No endpoints.
			tx.Add(&knftables.Rule{
				Chain: kubeServicesFilterChain,
				Rule: knftables.Concat(
					ipX, "daddr", svcInfo.ClusterIP(),
					protocol, "dport", svcInfo.Port(),
					internalTrafficFilterVerdict,
				),
				Comment: &internalTrafficFilterComment,
			})
		}

		// Capture externalIPs.
		for _, externalIP := range svcInfo.ExternalIPStrings() {
			if hasEndpoints {
				// Send traffic bound for external IPs to the "external
				// destinations" chain.
				proxier.natRules.Write(
					"-A", string(kubeServicesChain),
					"-m", "comment", "--comment", fmt.Sprintf(`"%s external IP"`, svcPortNameString),
					"-m", protocol, "-p", protocol,
					"-d", externalIP,
					"--dport", strconv.Itoa(svcInfo.Port()),
					"-j", string(externalTrafficChain))
			}
			if !hasExternalEndpoints {
				// Either no endpoints at all (REJECT) or no endpoints for
				// external traffic (DROP anything that didn't get
				// short-circuited by the EXT chain.)
				tx.Add(&knftables.Rule{
					Chain: kubeExternalServicesChain,
					Rule: knftables.Concat(
						ipX, "daddr", externalIP,
						protocol, "dport", svcInfo.Port(),
						externalTrafficFilterVerdict,
					),
					Comment: &externalTrafficFilterComment,
				})
			}
		}

		// Capture load-balancer ingress.
		for _, lbip := range svcInfo.LoadBalancerVIPStrings() {
			if hasEndpoints {
				proxier.natRules.Write(
					"-A", string(kubeServicesChain),
					"-m", "comment", "--comment", fmt.Sprintf(`"%s loadbalancer IP"`, svcPortNameString),
					"-m", protocol, "-p", protocol,
					"-d", lbip,
					"--dport", strconv.Itoa(svcInfo.Port()),
					"-j", string(loadBalancerTrafficChain))

			}
			if usesFWChain {
				comment := fmt.Sprintf("%s traffic not accepted by %s", svcPortNameString, svcInfo.firewallChainName)
				tx.Add(&knftables.Rule{
					Chain: kubeFirewallChain,
					Rule: knftables.Concat(
						ipX, "daddr", lbip,
						protocol, "dport", svcInfo.Port(),
						"drop",
					),
					Comment: &comment,
				})
			}
		}
		if !hasExternalEndpoints {
			// Either no endpoints at all (REJECT) or no endpoints for
			// external traffic (DROP anything that didn't get short-circuited
			// by the EXT chain.)
			for _, lbip := range svcInfo.LoadBalancerVIPStrings() {
				tx.Add(&knftables.Rule{
					Chain: kubeExternalServicesChain,
					Rule: knftables.Concat(
						ipX, "daddr", lbip,
						protocol, "dport", svcInfo.Port(),
						externalTrafficFilterVerdict,
					),
					Comment: &externalTrafficFilterComment,
				})
			}
		}

		// Capture nodeports.
		if svcInfo.NodePort() != 0 {
			if hasEndpoints {
				// Jump to the external destination chain.  For better or for
				// worse, nodeports are not subect to loadBalancerSourceRanges,
				// and we can't change that.
				proxier.natRules.Write(
					"-A", string(kubeNodePortsChain),
					"-m", "comment", "--comment", svcPortNameString,
					"-m", protocol, "-p", protocol,
					"--dport", strconv.Itoa(svcInfo.NodePort()),
					"-j", string(externalTrafficChain))
			}
			if !hasExternalEndpoints {
				// Either no endpoints at all (REJECT) or no endpoints for
				// external traffic (DROP anything that didn't get
				// short-circuited by the EXT chain.)
				tx.Add(&knftables.Rule{
					Chain: kubeExternalServicesChain,
					Rule: knftables.Concat(
						"fib daddr type local",
						protocol, "dport", svcInfo.NodePort(),
						externalTrafficFilterVerdict,
					),
					Comment: &externalTrafficFilterComment,
				})
			}
		}

		// Set up internal traffic handling.
		if hasInternalEndpoints {
			if proxier.masqueradeAll {
				proxier.natRules.Write(
					"-A", string(internalTrafficChain),
					"-m", "comment", "--comment", fmt.Sprintf(`"%s cluster IP"`, svcPortNameString),
					"-m", protocol, "-p", protocol,
					"-d", svcInfo.ClusterIP().String(),
					"--dport", strconv.Itoa(svcInfo.Port()),
					"-j", string(kubeMarkMasqChain))
			} else if proxier.localDetector.IsImplemented() {
				// This masquerades off-cluster traffic to a service VIP. The
				// idea is that you can establish a static route for your
				// Service range, routing to any node, and that node will
				// bridge into the Service for you. Since that might bounce
				// off-node, we masquerade here.
				proxier.natRules.Write(
					"-A", string(internalTrafficChain),
					"-m", "comment", "--comment", fmt.Sprintf(`"%s cluster IP"`, svcPortNameString),
					"-m", protocol, "-p", protocol,
					"-d", svcInfo.ClusterIP().String(),
					"--dport", strconv.Itoa(svcInfo.Port()),
					proxier.localDetector.IfNotLocal(),
					"-j", string(kubeMarkMasqChain))
			}
		}

		// Set up external traffic handling (if any "external" destinations are
		// enabled). All captured traffic for all external destinations should
		// jump to externalTrafficChain, which will handle some special cases and
		// then jump to externalPolicyChain.
		if usesExternalTrafficChain {
			if !svcInfo.ExternalPolicyLocal() {
				// If we are using non-local endpoints we need to masquerade,
				// in case we cross nodes.
				proxier.natRules.Write(
					"-A", string(externalTrafficChain),
					"-m", "comment", "--comment", fmt.Sprintf(`"masquerade traffic for %s external destinations"`, svcPortNameString),
					"-j", string(kubeMarkMasqChain))
			} else {
				// If we are only using same-node endpoints, we can retain the
				// source IP in most cases.

				if proxier.localDetector.IsImplemented() {
					// Treat all locally-originated pod -> external destination
					// traffic as a special-case.  It is subject to neither
					// form of traffic policy, which simulates going up-and-out
					// to an external load-balancer and coming back in.
					proxier.natRules.Write(
						"-A", string(externalTrafficChain),
						"-m", "comment", "--comment", fmt.Sprintf(`"pod traffic for %s external destinations"`, svcPortNameString),
						proxier.localDetector.IfLocal(),
						"-j", string(clusterPolicyChain))
				}

				// Locally originated traffic (not a pod, but the host node)
				// still needs masquerade because the LBIP itself is a local
				// address, so that will be the chosen source IP.
				proxier.natRules.Write(
					"-A", string(externalTrafficChain),
					"-m", "comment", "--comment", fmt.Sprintf(`"masquerade LOCAL traffic for %s external destinations"`, svcPortNameString),
					"-m", "addrtype", "--src-type", "LOCAL",
					"-j", string(kubeMarkMasqChain))

				// Redirect all src-type=LOCAL -> external destination to the
				// policy=cluster chain. This allows traffic originating
				// from the host to be redirected to the service correctly.
				proxier.natRules.Write(
					"-A", string(externalTrafficChain),
					"-m", "comment", "--comment", fmt.Sprintf(`"route LOCAL traffic for %s external destinations"`, svcPortNameString),
					"-m", "addrtype", "--src-type", "LOCAL",
					"-j", string(clusterPolicyChain))
			}

			// Anything else falls thru to the appropriate policy chain.
			if hasExternalEndpoints {
				proxier.natRules.Write(
					"-A", string(externalTrafficChain),
					"-j", string(externalPolicyChain))
			}
		}

		// Set up firewall chain, if needed
		if usesFWChain {
			// The service firewall rules are created based on the
			// loadBalancerSourceRanges field. This only works for VIP-like
			// loadbalancers that preserve source IPs. For loadbalancers which
			// direct traffic to service NodePort, the firewall rules will not
			// apply.

			// firewall filter based on each source range
			allowFromNode := false
			for _, src := range svcInfo.LoadBalancerSourceRanges() {
				proxier.natRules.Write(
					"-A", string(fwChain),
					"-m", "comment", "--comment", fmt.Sprintf(`"%s loadbalancer IP"`, svcPortNameString),
					"-s", src,
					"-j", string(externalTrafficChain),
				)
				_, cidr, err := netutils.ParseCIDRSloppy(src)
				if err != nil {
					klog.ErrorS(err, "Error parsing CIDR in LoadBalancerSourceRanges, dropping it", "cidr", cidr)
				} else if cidr.Contains(proxier.nodeIP) {
					allowFromNode = true
				}
			}
			// For VIP-like LBs, the VIP is often added as a local
			// address (via an IP route rule).  In that case, a request
			// from a node to the VIP will not hit the loadbalancer but
			// will loop back with the source IP set to the VIP.  We
			// need the following rules to allow requests from this node.
			if allowFromNode {
				for _, lbip := range svcInfo.LoadBalancerVIPStrings() {
					proxier.natRules.Write(
						"-A", string(fwChain),
						"-m", "comment", "--comment", fmt.Sprintf(`"%s loadbalancer IP"`, svcPortNameString),
						"-s", lbip,
						"-j", string(externalTrafficChain),
					)
				}
			}
			// If the packet was able to reach the end of firewall chain,
			// then it did not get DNATed, so it will match the
			// corresponding KUBE-PROXY-FIREWALL rule.
			proxier.natRules.Write(
				"-A", string(fwChain),
				"-m", "comment", "--comment", fmt.Sprintf(`"other traffic to %s will be dropped by KUBE-PROXY-FIREWALL"`, svcPortNameString),
			)
		}

		// If Cluster policy is in use, create rules jumping from
		// clusterPolicyChain to the clusterEndpoints
		if usesClusterPolicyChain {
			proxier.writeServiceToEndpointRules(svcPortNameString, svcInfo, clusterPolicyChain, clusterEndpoints)
		}

		// If Local policy is in use, create rules jumping from localPolicyChain
		// to the localEndpoints
		if usesLocalPolicyChain {
			proxier.writeServiceToEndpointRules(svcPortNameString, svcInfo, localPolicyChain, localEndpoints)
		}

		// Generate the per-endpoint chains.
		for _, ep := range allLocallyReachableEndpoints {
			epInfo, ok := ep.(*endpointInfo)
			if !ok {
				klog.ErrorS(nil, "Failed to cast endpointInfo", "endpointInfo", ep)
				continue
			}

			endpointChain := epInfo.chainName

			// Handle traffic that loops back to the originator with SNAT.
			proxier.natRules.Write(
				"-A", string(endpointChain),
				"-m", "comment", "--comment", svcPortNameString,
				"-s", epInfo.IP(),
				"-j", string(kubeMarkMasqChain),
			)
			commentAndAffinityArgs := []string{"-m", "comment", "--comment", svcPortNameString}
			if svcInfo.SessionAffinityType() == v1.ServiceAffinityClientIP {
				commentAndAffinityArgs = append(commentAndAffinityArgs, "-m", "recent", "--name", string(endpointChain), "--set")
			}
			// DNAT to final destination.
			proxier.natRules.Write(
				"-A", string(endpointChain),
				commentAndAffinityArgs,
				"-m", protocol, "-p", protocol,
				"-j", "DNAT", "--to-destination", epInfo.String(),
			)
		}
	}

	// Finally, tail-call to the nodePorts chain.  This needs to be after all
	// other service portal rules.
	if proxier.nodePortAddresses.MatchAll() {
		isIPv6 := proxier.ipFamily == v1.IPv6Protocol

		destinations := []string{"-m", "addrtype", "--dst-type", "LOCAL"}
		// Block localhost nodePorts
		if isIPv6 {
			destinations = append(destinations, "!", "-d", "::1/128")
		} else {
			destinations = append(destinations, "!", "-d", "127.0.0.0/8")
		}

		proxier.natRules.Write(
			"-A", string(kubeServicesChain),
			"-m", "comment", "--comment", `"kubernetes service nodeports; NOTE: this must be the last rule in this chain"`,
			destinations,
			"-j", string(kubeNodePortsChain))
	} else {
		nodeIPs, err := proxier.nodePortAddresses.GetNodeIPs(proxier.networkInterfacer)
		if err != nil {
			klog.ErrorS(err, "Failed to get node ip address matching nodeport cidrs, services with nodeport may not work as intended", "CIDRs", proxier.nodePortAddresses)
		}
		for _, ip := range nodeIPs {
			if ip.IsLoopback() {
				klog.ErrorS(nil, "--nodeport-addresses includes localhost but localhost NodePorts are not supported", "address", ip.String())
				continue
			}

			// create nodeport rules for each IP one by one
			proxier.natRules.Write(
				"-A", string(kubeServicesChain),
				"-m", "comment", "--comment", `"kubernetes service nodeports; NOTE: this must be the last rule in this chain"`,
				"-d", ip.String(),
				"-j", string(kubeNodePortsChain))
		}
	}

	// Figure out which chains are now stale. Unfortunately, we can't delete them
	// right away, because with kernels before 6.2, if there is a map element pointing
	// to a chain, and you delete that map element, the kernel doesn't notice until a
	// short amount of time later that the chain is now unreferenced. So we flush them
	// now, and record the time that they become stale in staleChains so they can be
	// deleted later.
	existingChains, err := proxier.nftables.List(context.TODO(), "chains")
	if err == nil {
		for _, chain := range existingChains {
			if isServiceChainName(chain) && !activeChains.Has(chain) {
				tx.Flush(&knftables.Chain{
					Name: chain,
				})
				proxier.staleChains[chain] = start
			}
		}
	} else if !knftables.IsNotFound(err) {
		klog.ErrorS(err, "Failed to list nftables chains: stale chains will not be deleted")
	}

	metrics.IptablesRulesTotal.WithLabelValues(string(utiliptables.TableNAT)).Set(float64(proxier.natRules.Lines()))

	// Sync rules.
	klog.V(2).InfoS("Reloading service nftables data",
		"numServices", len(proxier.svcPortMap),
		"numEndpoints", totalEndpoints,
		"numNATRules", proxier.natRules.Lines(),
	)

	// FIXME
	// klog.V(9).InfoS("Running nftables transaction", "transaction", tx.Bytes())

	err = proxier.nftables.Run(context.TODO(), tx)
	if err != nil {
		klog.ErrorS(err, "nftables sync failed")
		metrics.IptablesRestoreFailuresTotal.Inc()
		return
	}
	success = true

	for name, lastChangeTriggerTimes := range endpointUpdateResult.LastChangeTriggerTimes {
		for _, lastChangeTriggerTime := range lastChangeTriggerTimes {
			latency := metrics.SinceInSeconds(lastChangeTriggerTime)
			metrics.NetworkProgrammingLatency.Observe(latency)
			klog.V(4).InfoS("Network programming", "endpoint", klog.KRef(name.Namespace, name.Name), "elapsed", latency)
		}
	}

	metrics.SyncProxyRulesNoLocalEndpointsTotal.WithLabelValues("internal").Set(float64(serviceNoLocalEndpointsTotalInternal))
	metrics.SyncProxyRulesNoLocalEndpointsTotal.WithLabelValues("external").Set(float64(serviceNoLocalEndpointsTotalExternal))
	if proxier.healthzServer != nil {
		proxier.healthzServer.Updated(proxier.ipFamily)
	}
	metrics.SyncProxyRulesLastTimestamp.SetToCurrentTime()

	// Update service healthchecks.  The endpoints list might include services that are
	// not "OnlyLocal", but the services list will not, and the serviceHealthServer
	// will just drop those endpoints.
	if err := proxier.serviceHealthServer.SyncServices(proxier.svcPortMap.HealthCheckNodePorts()); err != nil {
		klog.ErrorS(err, "Error syncing healthcheck services")
	}
	if err := proxier.serviceHealthServer.SyncEndpoints(proxier.endpointsMap.LocalReadyEndpoints()); err != nil {
		klog.ErrorS(err, "Error syncing healthcheck endpoints")
	}

	// Finish housekeeping, clear stale conntrack entries for UDP Services
	conntrack.CleanStaleEntries(proxier.ipFamily == v1.IPv6Protocol, proxier.exec, proxier.svcPortMap, serviceUpdateResult, endpointUpdateResult)
}

func (proxier *Proxier) writeServiceToEndpointRules(svcPortNameString string, svcInfo proxy.ServicePort, svcChain string, endpoints []proxy.Endpoint) {
	// First write session affinity rules, if applicable.
	if svcInfo.SessionAffinityType() == v1.ServiceAffinityClientIP {
		for _, ep := range endpoints {
			epInfo, ok := ep.(*endpointInfo)
			if !ok {
				continue
			}
			comment := fmt.Sprintf(`"%s -> %s"`, svcPortNameString, epInfo.String())
			proxier.natRules.Write(
				"-A", string(svcChain),
				"-m", "comment", "--comment", comment,
				"-m", "recent", "--name", string(epInfo.chainName),
				"--rcheck", "--seconds", strconv.Itoa(svcInfo.StickyMaxAgeSeconds()), "--reap",
				"-j", string(epInfo.chainName),
			)
		}
	}

	// Now write loadbalancing rules.
	numEndpoints := len(endpoints)
	for i, ep := range endpoints {
		epInfo, ok := ep.(*endpointInfo)
		if !ok {
			continue
		}
		comment := fmt.Sprintf(`"%s -> %s"`, svcPortNameString, epInfo.String())

		if i < (numEndpoints - 1) {
			// Each rule is a probabilistic match.
			proxier.natRules.Write(
				"-A", string(svcChain),
				"-m", "comment", "--comment", comment,
				"-m", "statistic",
				"--mode", "random",
				"--probability", proxier.probability(numEndpoints-i),
				"-j", string(epInfo.chainName),
			)
		} else {
			// The final (or only if n == 1) rule is a guaranteed match.
			proxier.natRules.Write(
				"-A", string(svcChain),
				"-m", "comment", "--comment", comment,
				"-j", string(epInfo.chainName),
			)
		}
	}
}
