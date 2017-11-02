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

package dockershim

import (
	"fmt"
	"os"
	"strings"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	dockercontainer "github.com/docker/docker/api/types/container"
	dockerfilters "github.com/docker/docker/api/types/filters"
	"github.com/golang/glog"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	runtimeapi "k8s.io/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"
	kubecontainer "k8s.io/kubernetes/pkg/kubelet/container"
	"k8s.io/kubernetes/pkg/kubelet/dockershim/errors"
	"k8s.io/kubernetes/pkg/kubelet/dockershim/libdocker"
	"k8s.io/kubernetes/pkg/kubelet/qos"
	"k8s.io/kubernetes/pkg/kubelet/types"
)

const (
	defaultSandboxImage = "gcr.io/google_containers/pause-amd64:3.0"

	// Various default sandbox resources requests/limits.
	defaultSandboxCPUshares int64 = 2

	// Name of the underlying container runtime
	runtimeName = "docker"
)

var (
	// Termination grace period
	defaultSandboxGracePeriod = time.Duration(10) * time.Second
)

// Returns whether the sandbox network is ready, and whether the sandbox is known
func (ds *dockerService) getNetworkReady(podSandboxID string) (bool, bool) {
	ds.networkReadyLock.Lock()
	defer ds.networkReadyLock.Unlock()
	ready, ok := ds.networkReady[podSandboxID]
	return ready, ok
}

func (ds *dockerService) setNetworkReady(podSandboxID string, ready bool) {
	ds.networkReadyLock.Lock()
	defer ds.networkReadyLock.Unlock()
	ds.networkReady[podSandboxID] = ready
}

func (ds *dockerService) clearNetworkReady(podSandboxID string) {
	ds.networkReadyLock.Lock()
	defer ds.networkReadyLock.Unlock()
	delete(ds.networkReady, podSandboxID)
}

// RunPodSandbox creates and starts a pod-level sandbox. Runtimes should ensure
// the sandbox is in ready state.
// For docker, PodSandbox is implemented by a container holding the network
// namespace for the pod.
// Note: docker doesn't use LogDirectory (yet).
func (ds *dockerService) RunPodSandbox(config *runtimeapi.PodSandboxConfig) (id string, err error) {
	// Step 1: Pull the image for the sandbox.
	image := defaultSandboxImage
	podSandboxImage := ds.podSandboxImage
	if len(podSandboxImage) != 0 {
		image = podSandboxImage
	}

	// NOTE: To use a custom sandbox image in a private repository, users need to configure the nodes with credentials properly.
	// see: http://kubernetes.io/docs/user-guide/images/#configuring-nodes-to-authenticate-to-a-private-repository
	// Only pull sandbox image when it's not present - v1.PullIfNotPresent.
	if err := ensureSandboxImageExists(ds.client, image); err != nil {
		return "", err
	}

	// Step 2: Create the sandbox container.
	createConfig, err := ds.makeSandboxDockerConfig(config, image)
	if err != nil {
		return "", fmt.Errorf("failed to make sandbox docker config for pod %q: %v", config.Metadata.Name, err)
	}
	createResp, err := ds.client.CreateContainer(*createConfig)
	if err != nil {
		createResp, err = recoverFromCreationConflictIfNeeded(ds.client, *createConfig, err)
	}

	if err != nil || createResp == nil {
		return "", fmt.Errorf("failed to create a sandbox for pod %q: %v", config.Metadata.Name, err)
	}

	ds.setNetworkReady(createResp.ID, false)
	defer func(e *error) {
		// Set networking ready depending on the error return of
		// the parent function
		if *e == nil {
			ds.setNetworkReady(createResp.ID, true)
		}
	}(&err)

	// Step 3: Create Sandbox Checkpoint.
	if err = ds.checkpointHandler.CreateCheckpoint(createResp.ID, constructPodSandboxCheckpoint(config)); err != nil {
		return createResp.ID, err
	}

	// Step 4: Start the sandbox container.
	// Assume kubelet's garbage collector would remove the sandbox later, if
	// startContainer failed.
	err = ds.client.StartContainer(createResp.ID)
	if err != nil {
		return createResp.ID, fmt.Errorf("failed to start sandbox container for pod %q: %v", config.Metadata.Name, err)
	}

	// Rewrite resolv.conf file generated by docker.
	// NOTE: cluster dns settings aren't passed anymore to docker api in all cases,
	// not only for pods with host network: the resolver conf will be overwritten
	// after sandbox creation to override docker's behaviour. This resolv.conf
	// file is shared by all containers of the same pod, and needs to be modified
	// only once per pod.
	if dnsConfig := config.GetDnsConfig(); dnsConfig != nil {
		containerInfo, err := ds.client.InspectContainer(createResp.ID)
		if err != nil {
			return createResp.ID, fmt.Errorf("failed to inspect sandbox container for pod %q: %v", config.Metadata.Name, err)
		}

		if err := rewriteResolvFile(containerInfo.ResolvConfPath, dnsConfig.Servers, dnsConfig.Searches, dnsConfig.Options); err != nil {
			return createResp.ID, fmt.Errorf("rewrite resolv.conf failed for pod %q: %v", config.Metadata.Name, err)
		}
	}

	// Do not invoke network plugins if in hostNetwork mode.
	if nsOptions := config.GetLinux().GetSecurityContext().GetNamespaceOptions(); nsOptions != nil && nsOptions.HostNetwork {
		return createResp.ID, nil
	}

	// Step 5: Setup networking for the sandbox.
	// All pod networking is setup by a CNI plugin discovered at startup time.
	// This plugin assigns the pod ip, sets up routes inside the sandbox,
	// creates interfaces etc. In theory, its jurisdiction ends with pod
	// sandbox networking, but it might insert iptables rules or open ports
	// on the host as well, to satisfy parts of the pod spec that aren't
	// recognized by the CNI standard yet.
	cID := kubecontainer.BuildContainerID(runtimeName, createResp.ID)
	err = ds.network.SetUpPod(config.GetMetadata().Namespace, config.GetMetadata().Name, cID, config.Annotations)
	if err != nil {
		// TODO(random-liu): Do we need to teardown network here?
		if err := ds.client.StopContainer(createResp.ID, defaultSandboxGracePeriod); err != nil {
			glog.Warningf("Failed to stop sandbox container %q for pod %q: %v", createResp.ID, config.Metadata.Name, err)
		}
	}
	return createResp.ID, err
}

// StopPodSandbox stops the sandbox. If there are any running containers in the
// sandbox, they should be force terminated.
// TODO: This function blocks sandbox teardown on networking teardown. Is it
// better to cut our losses assuming an out of band GC routine will cleanup
// after us?
func (ds *dockerService) StopPodSandbox(podSandboxID string) error {
	var namespace, name string
	var hostNetwork bool
	var checkpointErr, statusErr error

	// Try to retrieve sandbox information from docker daemon or sandbox checkpoint
	status, statusErr := ds.PodSandboxStatus(podSandboxID)
	if statusErr == nil {
		nsOpts := status.GetLinux().GetNamespaces().GetOptions()
		hostNetwork = nsOpts != nil && nsOpts.HostNetwork
		m := status.GetMetadata()
		namespace = m.Namespace
		name = m.Name
	} else {
		var checkpoint *PodSandboxCheckpoint
		checkpoint, checkpointErr = ds.checkpointHandler.GetCheckpoint(podSandboxID)

		// Proceed if both sandbox container and checkpoint could not be found. This means that following
		// actions will only have sandbox ID and not have pod namespace and name information.
		// Return error if encounter any unexpected error.
		if checkpointErr != nil {
			if libdocker.IsContainerNotFoundError(statusErr) && checkpointErr == errors.CheckpointNotFoundError {
				glog.Warningf("Both sandbox container and checkpoint for id %q could not be found. "+
					"Proceed without further sandbox information.", podSandboxID)
			} else {
				if checkpointErr == errors.CorruptCheckpointError {
					// Remove the corrupted checkpoint so that the next
					// StopPodSandbox call can proceed. This may indicate that
					// some resources won't be reclaimed.
					// TODO (#43021): Fix this properly.
					glog.Warningf("Removing corrupted checkpoint %q: %+v", podSandboxID, *checkpoint)
					if err := ds.checkpointHandler.RemoveCheckpoint(podSandboxID); err != nil {
						glog.Warningf("Unable to remove corrupted checkpoint %q: %v", podSandboxID, err)
					}
				}
				return utilerrors.NewAggregate([]error{
					fmt.Errorf("failed to get checkpoint for sandbox %q: %v", podSandboxID, checkpointErr),
					fmt.Errorf("failed to get sandbox status: %v", statusErr)})
			}
		} else {
			namespace = checkpoint.Namespace
			name = checkpoint.Name
			hostNetwork = checkpoint.Data != nil && checkpoint.Data.HostNetwork
		}
	}

	// WARNING: The following operations made the following assumption:
	// 1. kubelet will retry on any error returned by StopPodSandbox.
	// 2. tearing down network and stopping sandbox container can succeed in any sequence.
	// This depends on the implementation detail of network plugin and proper error handling.
	// For kubenet, if tearing down network failed and sandbox container is stopped, kubelet
	// will retry. On retry, kubenet will not be able to retrieve network namespace of the sandbox
	// since it is stopped. With empty network namespcae, CNI bridge plugin will conduct best
	// effort clean up and will not return error.
	errList := []error{}
	ready, ok := ds.getNetworkReady(podSandboxID)
	if !hostNetwork && (ready || !ok) {
		// Only tear down the pod network if we haven't done so already
		cID := kubecontainer.BuildContainerID(runtimeName, podSandboxID)
		err := ds.network.TearDownPod(namespace, name, cID)
		if err == nil {
			ds.setNetworkReady(podSandboxID, false)
		} else {
			errList = append(errList, err)
		}
	}
	if err := ds.client.StopContainer(podSandboxID, defaultSandboxGracePeriod); err != nil {
		// Do not return error if the container does not exist
		if !libdocker.IsContainerNotFoundError(err) {
			glog.Errorf("Failed to stop sandbox %q: %v", podSandboxID, err)
			errList = append(errList, err)
		}
	}
	return utilerrors.NewAggregate(errList)
	// TODO: Stop all running containers in the sandbox.
}

// RemovePodSandbox removes the sandbox. If there are running containers in the
// sandbox, they should be forcibly removed.
func (ds *dockerService) RemovePodSandbox(podSandboxID string) error {
	var errs []error
	opts := dockertypes.ContainerListOptions{All: true}

	opts.Filters = dockerfilters.NewArgs()
	f := newDockerFilter(&opts.Filters)
	f.AddLabel(sandboxIDLabelKey, podSandboxID)

	containers, err := ds.client.ListContainers(opts)
	if err != nil {
		errs = append(errs, err)
	}

	// Remove all containers in the sandbox.
	for i := range containers {
		if err := ds.RemoveContainer(containers[i].ID); err != nil && !libdocker.IsContainerNotFoundError(err) {
			errs = append(errs, err)
		}
	}

	// Remove the sandbox container.
	err = ds.client.RemoveContainer(podSandboxID, dockertypes.ContainerRemoveOptions{RemoveVolumes: true, Force: true})
	if err == nil || libdocker.IsContainerNotFoundError(err) {
		// Only clear network ready when the sandbox has actually been
		// removed from docker or doesn't exist
		ds.clearNetworkReady(podSandboxID)
	} else {
		errs = append(errs, err)
	}

	// Remove the checkpoint of the sandbox.
	if err := ds.checkpointHandler.RemoveCheckpoint(podSandboxID); err != nil {
		errs = append(errs, err)
	}
	return utilerrors.NewAggregate(errs)
}

// getIPFromPlugin interrogates the network plugin for an IP.
func (ds *dockerService) getIPFromPlugin(sandbox *dockertypes.ContainerJSON) (string, error) {
	metadata, err := parseSandboxName(sandbox.Name)
	if err != nil {
		return "", err
	}
	msg := fmt.Sprintf("Couldn't find network status for %s/%s through plugin", metadata.Namespace, metadata.Name)
	cID := kubecontainer.BuildContainerID(runtimeName, sandbox.ID)
	networkStatus, err := ds.network.GetPodNetworkStatus(metadata.Namespace, metadata.Name, cID)
	if err != nil {
		return "", err
	}
	if networkStatus == nil {
		return "", fmt.Errorf("%v: invalid network status for", msg)
	}
	return networkStatus.IP.String(), nil
}

// getIP returns the ip given the output of `docker inspect` on a pod sandbox,
// first interrogating any registered plugins, then simply trusting the ip
// in the sandbox itself. We look for an ipv4 address before ipv6.
func (ds *dockerService) getIP(podSandboxID string, sandbox *dockertypes.ContainerJSON) string {
	if sandbox.NetworkSettings == nil {
		return ""
	}
	if sharesHostNetwork(sandbox) {
		// For sandboxes using host network, the shim is not responsible for
		// reporting the IP.
		return ""
	}

	// Don't bother getting IP if the pod is known and networking isn't ready
	ready, ok := ds.getNetworkReady(podSandboxID)
	if ok && !ready {
		return ""
	}

	ip, err := ds.getIPFromPlugin(sandbox)
	if err == nil {
		return ip
	}

	// TODO: trusting the docker ip is not a great idea. However docker uses
	// eth0 by default and so does CNI, so if we find a docker IP here, we
	// conclude that the plugin must have failed setup, or forgotten its ip.
	// This is not a sensible assumption for plugins across the board, but if
	// a plugin doesn't want this behavior, it can throw an error.
	if sandbox.NetworkSettings.IPAddress != "" {
		return sandbox.NetworkSettings.IPAddress
	}
	if sandbox.NetworkSettings.GlobalIPv6Address != "" {
		return sandbox.NetworkSettings.GlobalIPv6Address
	}

	// If all else fails, warn but don't return an error, as pod status
	// should generally not return anything except fatal errors
	// FIXME: handle network errors by restarting the pod somehow?
	glog.Warningf("failed to read pod IP from plugin/docker: %v", err)
	return ""
}

// PodSandboxStatus returns the status of the PodSandbox.
func (ds *dockerService) PodSandboxStatus(podSandboxID string) (*runtimeapi.PodSandboxStatus, error) {
	// Inspect the container.
	r, err := ds.client.InspectContainer(podSandboxID)
	if err != nil {
		return nil, err
	}

	// Parse the timestamps.
	createdAt, _, _, err := getContainerTimestamps(r)
	if err != nil {
		return nil, fmt.Errorf("failed to parse timestamp for container %q: %v", podSandboxID, err)
	}
	ct := createdAt.UnixNano()

	// Translate container to sandbox state.
	state := runtimeapi.PodSandboxState_SANDBOX_NOTREADY
	if r.State.Running {
		state = runtimeapi.PodSandboxState_SANDBOX_READY
	}

	var IP string
	// TODO: Remove this when sandbox is available on windows
	// This is a workaround for windows, where sandbox is not in use, and pod IP is determined through containers belonging to the Pod.
	if IP = ds.determinePodIPBySandboxID(podSandboxID); IP == "" {
		IP = ds.getIP(podSandboxID, r)
	}
	hostNetwork := sharesHostNetwork(r)

	metadata, err := parseSandboxName(r.Name)
	if err != nil {
		return nil, err
	}
	labels, annotations := extractLabels(r.Config.Labels)
	return &runtimeapi.PodSandboxStatus{
		Id:          r.ID,
		State:       state,
		CreatedAt:   ct,
		Metadata:    metadata,
		Labels:      labels,
		Annotations: annotations,
		Network: &runtimeapi.PodSandboxNetworkStatus{
			Ip: IP,
		},
		Linux: &runtimeapi.LinuxPodSandboxStatus{
			Namespaces: &runtimeapi.Namespace{
				Options: &runtimeapi.NamespaceOption{
					HostNetwork: hostNetwork,
					HostPid:     sharesHostPid(r),
					HostIpc:     sharesHostIpc(r),
				},
			},
		},
	}, nil
}

// ListPodSandbox returns a list of Sandbox.
func (ds *dockerService) ListPodSandbox(filter *runtimeapi.PodSandboxFilter) ([]*runtimeapi.PodSandbox, error) {
	// By default, list all containers whether they are running or not.
	opts := dockertypes.ContainerListOptions{All: true}
	filterOutReadySandboxes := false

	opts.Filters = dockerfilters.NewArgs()
	f := newDockerFilter(&opts.Filters)
	// Add filter to select only sandbox containers.
	f.AddLabel(containerTypeLabelKey, containerTypeLabelSandbox)

	if filter != nil {
		if filter.Id != "" {
			f.Add("id", filter.Id)
		}
		if filter.State != nil {
			if filter.GetState().State == runtimeapi.PodSandboxState_SANDBOX_READY {
				// Only list running containers.
				opts.All = false
			} else {
				// runtimeapi.PodSandboxState_SANDBOX_NOTREADY can mean the
				// container is in any of the non-running state (e.g., created,
				// exited). We can't tell docker to filter out running
				// containers directly, so we'll need to filter them out
				// ourselves after getting the results.
				filterOutReadySandboxes = true
			}
		}

		if filter.LabelSelector != nil {
			for k, v := range filter.LabelSelector {
				f.AddLabel(k, v)
			}
		}
	}

	// Make sure we get the list of checkpoints first so that we don't include
	// new PodSandboxes that are being created right now.
	var err error
	checkpoints := []string{}
	if filter == nil {
		checkpoints, err = ds.checkpointHandler.ListCheckpoints()
		if err != nil {
			glog.Errorf("Failed to list checkpoints: %v", err)
		}
	}

	containers, err := ds.client.ListContainers(opts)
	if err != nil {
		return nil, err
	}

	// Convert docker containers to runtime api sandboxes.
	result := []*runtimeapi.PodSandbox{}
	// using map as set
	sandboxIDs := make(map[string]bool)
	for i := range containers {
		c := containers[i]
		converted, err := containerToRuntimeAPISandbox(&c)
		if err != nil {
			glog.V(4).Infof("Unable to convert docker to runtime API sandbox %+v: %v", c, err)
			continue
		}
		if filterOutReadySandboxes && converted.State == runtimeapi.PodSandboxState_SANDBOX_READY {
			continue
		}
		sandboxIDs[converted.Id] = true
		result = append(result, converted)
	}

	// Include sandbox that could only be found with its checkpoint if no filter is applied
	// These PodSandbox will only include PodSandboxID, Name, Namespace.
	// These PodSandbox will be in PodSandboxState_SANDBOX_NOTREADY state.
	for _, id := range checkpoints {
		if _, ok := sandboxIDs[id]; ok {
			continue
		}
		checkpoint, err := ds.checkpointHandler.GetCheckpoint(id)
		if err != nil {
			glog.Errorf("Failed to retrieve checkpoint for sandbox %q: %v", id, err)

			if err == errors.CorruptCheckpointError {
				glog.Warningf("Removing corrupted checkpoint %q: %+v", id, *checkpoint)
				if err := ds.checkpointHandler.RemoveCheckpoint(id); err != nil {
					glog.Warningf("Unable to remove corrupted checkpoint %q: %v", id, err)
				}
			}
			continue
		}
		result = append(result, checkpointToRuntimeAPISandbox(id, checkpoint))
	}

	return result, nil
}

// applySandboxLinuxOptions applies LinuxPodSandboxConfig to dockercontainer.HostConfig and dockercontainer.ContainerCreateConfig.
func (ds *dockerService) applySandboxLinuxOptions(hc *dockercontainer.HostConfig, lc *runtimeapi.LinuxPodSandboxConfig, createConfig *dockertypes.ContainerCreateConfig, image string, separator rune) error {
	// Apply Cgroup options.
	cgroupParent, err := ds.GenerateExpectedCgroupParent(lc.CgroupParent)
	if err != nil {
		return err
	}
	hc.CgroupParent = cgroupParent
	// Apply security context.
	if err = applySandboxSecurityContext(lc, createConfig.Config, hc, ds.network, separator); err != nil {
		return err
	}

	// Set sysctls.
	hc.Sysctls = lc.Sysctls

	return nil
}

// makeSandboxDockerConfig returns dockertypes.ContainerCreateConfig based on runtimeapi.PodSandboxConfig.
func (ds *dockerService) makeSandboxDockerConfig(c *runtimeapi.PodSandboxConfig, image string) (*dockertypes.ContainerCreateConfig, error) {
	// Merge annotations and labels because docker supports only labels.
	labels := makeLabels(c.GetLabels(), c.GetAnnotations())
	// Apply a label to distinguish sandboxes from regular containers.
	labels[containerTypeLabelKey] = containerTypeLabelSandbox
	// Apply a container name label for infra container. This is used in summary v1.
	// TODO(random-liu): Deprecate this label once container metrics is directly got from CRI.
	labels[types.KubernetesContainerNameLabel] = sandboxContainerName

	apiVersion, err := ds.getDockerAPIVersion()
	if err != nil {
		return nil, fmt.Errorf("unable to get the docker API version: %v", err)
	}
	securityOptSep := getSecurityOptSeparator(apiVersion)

	hc := &dockercontainer.HostConfig{}
	createConfig := &dockertypes.ContainerCreateConfig{
		Name: makeSandboxName(c),
		Config: &dockercontainer.Config{
			Hostname: c.Hostname,
			// TODO: Handle environment variables.
			Image:  image,
			Labels: labels,
		},
		HostConfig: hc,
	}

	// Apply linux-specific options.
	if lc := c.GetLinux(); lc != nil {
		if err := ds.applySandboxLinuxOptions(hc, lc, createConfig, image, securityOptSep); err != nil {
			return nil, err
		}
	}

	// Set port mappings.
	exposedPorts, portBindings := makePortsAndBindings(c.GetPortMappings())
	createConfig.Config.ExposedPorts = exposedPorts
	hc.PortBindings = portBindings

	// Apply resource options.
	setSandboxResources(hc)

	// Apply cgroupsParent derived from the sandbox config.
	if lc := c.GetLinux(); lc != nil {
		// Apply Cgroup options.
		cgroupParent, err := ds.GenerateExpectedCgroupParent(lc.CgroupParent)
		if err != nil {
			return nil, fmt.Errorf("failed to generate cgroup parent in expected syntax for container %q: %v", c.Metadata.Name, err)
		}
		hc.CgroupParent = cgroupParent
	}

	// Set security options.
	securityOpts, err := ds.getSecurityOpts(c.GetLinux().GetSecurityContext().GetSeccompProfilePath(), securityOptSep)
	if err != nil {
		return nil, fmt.Errorf("failed to generate sandbox security options for sandbox %q: %v", c.Metadata.Name, err)
	}
	hc.SecurityOpt = append(hc.SecurityOpt, securityOpts...)
	return createConfig, nil
}

// sharesHostNetwork returns true if the given container is sharing the host's
// network namespace.
func sharesHostNetwork(container *dockertypes.ContainerJSON) bool {
	if container != nil && container.HostConfig != nil {
		return string(container.HostConfig.NetworkMode) == namespaceModeHost
	}
	return false
}

// sharesHostPid returns true if the given container is sharing the host's pid
// namespace.
func sharesHostPid(container *dockertypes.ContainerJSON) bool {
	if container != nil && container.HostConfig != nil {
		return string(container.HostConfig.PidMode) == namespaceModeHost
	}
	return false
}

// sharesHostIpc returns true if the given container is sharing the host's ipc
// namespace.
func sharesHostIpc(container *dockertypes.ContainerJSON) bool {
	if container != nil && container.HostConfig != nil {
		return string(container.HostConfig.IpcMode) == namespaceModeHost
	}
	return false
}

func setSandboxResources(hc *dockercontainer.HostConfig) {
	hc.Resources = dockercontainer.Resources{
		MemorySwap: DefaultMemorySwap(),
		CPUShares:  defaultSandboxCPUshares,
		// Use docker's default cpu quota/period.
	}
	// TODO: Get rid of the dependency on kubelet internal package.
	hc.OomScoreAdj = qos.PodInfraOOMAdj
}

func constructPodSandboxCheckpoint(config *runtimeapi.PodSandboxConfig) *PodSandboxCheckpoint {
	checkpoint := NewPodSandboxCheckpoint(config.Metadata.Namespace, config.Metadata.Name)
	for _, pm := range config.GetPortMappings() {
		proto := toCheckpointProtocol(pm.Protocol)
		checkpoint.Data.PortMappings = append(checkpoint.Data.PortMappings, &PortMapping{
			HostPort:      &pm.HostPort,
			ContainerPort: &pm.ContainerPort,
			Protocol:      &proto,
		})
	}
	if nsOptions := config.GetLinux().GetSecurityContext().GetNamespaceOptions(); nsOptions != nil {
		checkpoint.Data.HostNetwork = nsOptions.HostNetwork
	}
	return checkpoint
}

func toCheckpointProtocol(protocol runtimeapi.Protocol) Protocol {
	switch protocol {
	case runtimeapi.Protocol_TCP:
		return protocolTCP
	case runtimeapi.Protocol_UDP:
		return protocolUDP
	}
	glog.Warningf("Unknown protocol %q: defaulting to TCP", protocol)
	return protocolTCP
}

// rewriteResolvFile rewrites resolv.conf file generated by docker.
func rewriteResolvFile(resolvFilePath string, dns []string, dnsSearch []string, dnsOptions []string) error {
	if len(resolvFilePath) == 0 {
		glog.Errorf("ResolvConfPath is empty.")
		return nil
	}

	if _, err := os.Stat(resolvFilePath); os.IsNotExist(err) {
		return fmt.Errorf("ResolvConfPath %q does not exist", resolvFilePath)
	}

	var resolvFileContent []string
	for _, srv := range dns {
		resolvFileContent = append(resolvFileContent, "nameserver "+srv)
	}

	if len(dnsSearch) > 0 {
		resolvFileContent = append(resolvFileContent, "search "+strings.Join(dnsSearch, " "))
	}

	if len(dnsOptions) > 0 {
		resolvFileContent = append(resolvFileContent, "options "+strings.Join(dnsOptions, " "))
	}

	if len(resolvFileContent) > 0 {
		resolvFileContentStr := strings.Join(resolvFileContent, "\n")
		resolvFileContentStr += "\n"

		glog.V(4).Infof("Will attempt to re-write config file %s with: \n%s", resolvFilePath, resolvFileContent)
		if err := rewriteFile(resolvFilePath, resolvFileContentStr); err != nil {
			glog.Errorf("resolv.conf could not be updated: %v", err)
			return err
		}
	}

	return nil
}

func rewriteFile(filePath, stringToWrite string) error {
	f, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(stringToWrite)
	return err
}
