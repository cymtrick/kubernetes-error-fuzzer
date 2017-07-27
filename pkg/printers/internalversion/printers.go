/*
Copyright 2017 The Kubernetes Authors.

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

package internalversion

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	batchv1 "k8s.io/api/batch/v1"
	batchv2alpha1 "k8s.io/api/batch/v2alpha1"
	apiv1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1alpha1 "k8s.io/apimachinery/pkg/apis/meta/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/kubernetes/federation/apis/federation"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/events"
	"k8s.io/kubernetes/pkg/api/helper"
	"k8s.io/kubernetes/pkg/apis/apps"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	"k8s.io/kubernetes/pkg/apis/batch"
	"k8s.io/kubernetes/pkg/apis/certificates"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apis/networking"
	"k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/apis/rbac"
	"k8s.io/kubernetes/pkg/apis/storage"
	storageutil "k8s.io/kubernetes/pkg/apis/storage/util"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/printers"
	"k8s.io/kubernetes/pkg/util/node"
)

const loadBalancerWidth = 16

// NOTE: When adding a new resource type here, please update the list
// pkg/kubectl/cmd/get.go to reflect the new resource type.
var (
	persistentVolumeClaimColumns  = []string{"NAME", "STATUS", "VOLUME", "CAPACITY", "ACCESSMODES", "STORAGECLASS", "AGE"}
	componentStatusColumns        = []string{"NAME", "STATUS", "MESSAGE", "ERROR"}
	thirdPartyResourceColumns     = []string{"NAME", "DESCRIPTION", "VERSION(S)"}
	roleBindingColumns            = []string{"NAME", "AGE"}
	roleBindingWideColumns        = []string{"ROLE", "USERS", "GROUPS", "SERVICEACCOUNTS"}
	clusterRoleBindingColumns     = []string{"NAME", "AGE"}
	clusterRoleBindingWideColumns = []string{"ROLE", "USERS", "GROUPS", "SERVICEACCOUNTS"}
	storageClassColumns           = []string{"NAME", "PROVISIONER"}
	statusColumns                 = []string{"STATUS", "REASON", "MESSAGE"}

	horizontalPodAutoscalerColumns   = []string{"NAME", "REFERENCE", "TARGETS", "MINPODS", "MAXPODS", "REPLICAS", "AGE"}
	deploymentColumns                = []string{"NAME", "DESIRED", "CURRENT", "UP-TO-DATE", "AVAILABLE", "AGE"}
	deploymentWideColumns            = []string{"CONTAINER(S)", "IMAGE(S)", "SELECTOR"}
	configMapColumns                 = []string{"NAME", "DATA", "AGE"}
	podSecurityPolicyColumns         = []string{"NAME", "PRIV", "CAPS", "SELINUX", "RUNASUSER", "FSGROUP", "SUPGROUP", "READONLYROOTFS", "VOLUMES"}
	clusterColumns                   = []string{"NAME", "STATUS", "AGE"}
	networkPolicyColumns             = []string{"NAME", "POD-SELECTOR", "AGE"}
	certificateSigningRequestColumns = []string{"NAME", "AGE", "REQUESTOR", "CONDITION"}
	podPresetColumns                 = []string{"NAME", "AGE"}
	controllerRevisionColumns        = []string{"NAME", "CONTROLLER", "REVISION", "AGE"}
)

// AddHandlers adds print handlers for default Kubernetes types dealing with internal versions.
// TODO: handle errors from Handler
func AddHandlers(h printers.PrintHandler) {
	podColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Ready", Type: "string", Description: "The aggregate readiness state of this pod for accepting traffic."},
		{Name: "Status", Type: "string", Description: "The aggregate status of the containers in this pod."},
		{Name: "Restarts", Type: "integer", Description: "The number of times the containers in this pod have been restarted."},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "IP", Type: "string", Priority: 1, Description: apiv1.PodStatus{}.SwaggerDoc()["podIP"]},
		{Name: "Node", Type: "string", Priority: 1, Description: apiv1.PodSpec{}.SwaggerDoc()["nodeName"]},
	}
	h.TableHandler(podColumnDefinitions, printPodList)
	h.TableHandler(podColumnDefinitions, printPod)

	podTemplateColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Containers", Type: "string", Description: "Names of each container in the template."},
		{Name: "Images", Type: "string", Description: "Images referenced by each container in the template."},
		{Name: "Pod Labels", Type: "string", Description: "The labels for the pod template."},
	}
	h.TableHandler(podTemplateColumnDefinitions, printPodTemplate)
	h.TableHandler(podTemplateColumnDefinitions, printPodTemplateList)

	podDisruptionBudgetColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Min Available", Type: "string", Description: "The minimum number of pods that must be available."},
		{Name: "Max Unavailable", Type: "string", Description: "The maximum number of pods that may be unavailable."},
		{Name: "Allowed Disruptions", Type: "integer", Description: "Calculated number of pods that may be disrupted at this time."},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
	}
	h.TableHandler(podDisruptionBudgetColumnDefinitions, printPodDisruptionBudget)
	h.TableHandler(podDisruptionBudgetColumnDefinitions, printPodDisruptionBudgetList)

	replicationControllerColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Desired", Type: "integer", Description: apiv1.ReplicationControllerSpec{}.SwaggerDoc()["replicas"]},
		{Name: "Current", Type: "integer", Description: apiv1.ReplicationControllerStatus{}.SwaggerDoc()["replicas"]},
		{Name: "Ready", Type: "integer", Description: apiv1.ReplicationControllerStatus{}.SwaggerDoc()["readyReplicas"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "Containers", Type: "string", Priority: 1, Description: "Names of each container in the template."},
		{Name: "Images", Type: "string", Priority: 1, Description: "Images referenced by each container in the template."},
		{Name: "Selector", Type: "string", Priority: 1, Description: apiv1.ReplicationControllerSpec{}.SwaggerDoc()["selector"]},
	}
	h.TableHandler(replicationControllerColumnDefinitions, printReplicationController)
	h.TableHandler(replicationControllerColumnDefinitions, printReplicationControllerList)

	replicaSetColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Desired", Type: "integer", Description: extensionsv1beta1.ReplicaSetSpec{}.SwaggerDoc()["replicas"]},
		{Name: "Current", Type: "integer", Description: extensionsv1beta1.ReplicaSetStatus{}.SwaggerDoc()["replicas"]},
		{Name: "Ready", Type: "integer", Description: extensionsv1beta1.ReplicaSetStatus{}.SwaggerDoc()["readyReplicas"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "Containers", Type: "string", Priority: 1, Description: "Names of each container in the template."},
		{Name: "Images", Type: "string", Priority: 1, Description: "Images referenced by each container in the template."},
		{Name: "Selector", Type: "string", Priority: 1, Description: extensionsv1beta1.ReplicaSetSpec{}.SwaggerDoc()["selector"]},
	}
	h.TableHandler(replicaSetColumnDefinitions, printReplicaSet)
	h.TableHandler(replicaSetColumnDefinitions, printReplicaSetList)

	daemonSetColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Desired", Type: "integer", Description: extensionsv1beta1.DaemonSetStatus{}.SwaggerDoc()["desiredNumberScheduled"]},
		{Name: "Current", Type: "integer", Description: extensionsv1beta1.DaemonSetStatus{}.SwaggerDoc()["currentNumberScheduled"]},
		{Name: "Ready", Type: "integer", Description: extensionsv1beta1.DaemonSetStatus{}.SwaggerDoc()["numberReady"]},
		{Name: "Up-to-date", Type: "integer", Description: extensionsv1beta1.DaemonSetStatus{}.SwaggerDoc()["updatedNumberScheduled"]},
		{Name: "Available", Type: "integer", Description: extensionsv1beta1.DaemonSetStatus{}.SwaggerDoc()["numberAvailable"]},
		{Name: "Node Selector", Type: "string", Description: apiv1.PodSpec{}.SwaggerDoc()["nodeSelector"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "Containers", Type: "string", Priority: 1, Description: "Names of each container in the template."},
		{Name: "Images", Type: "string", Priority: 1, Description: "Images referenced by each container in the template."},
		{Name: "Selector", Type: "string", Priority: 1, Description: extensionsv1beta1.DaemonSetSpec{}.SwaggerDoc()["selector"]},
	}
	h.TableHandler(daemonSetColumnDefinitions, printDaemonSet)
	h.TableHandler(daemonSetColumnDefinitions, printDaemonSetList)

	jobColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Desired", Type: "integer", Description: batchv1.JobSpec{}.SwaggerDoc()["completions"]},
		{Name: "Successful", Type: "integer", Description: batchv1.JobStatus{}.SwaggerDoc()["succeeded"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "Containers", Type: "string", Priority: 1, Description: "Names of each container in the template."},
		{Name: "Images", Type: "string", Priority: 1, Description: "Images referenced by each container in the template."},
		{Name: "Selector", Type: "string", Priority: 1, Description: batchv1.JobSpec{}.SwaggerDoc()["selector"]},
	}
	h.TableHandler(jobColumnDefinitions, printJob)
	h.TableHandler(jobColumnDefinitions, printJobList)

	cronJobColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Schedule", Type: "string", Description: batchv2alpha1.CronJobSpec{}.SwaggerDoc()["schedule"]},
		{Name: "Suspend", Type: "boolean", Description: batchv2alpha1.CronJobSpec{}.SwaggerDoc()["suspend"]},
		{Name: "Active", Type: "integer", Description: batchv2alpha1.CronJobStatus{}.SwaggerDoc()["active"]},
		{Name: "Last Schedule", Type: "string", Description: batchv2alpha1.CronJobStatus{}.SwaggerDoc()["lastScheduleTime"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "Containers", Type: "string", Priority: 1, Description: "Names of each container in the template."},
		{Name: "Images", Type: "string", Priority: 1, Description: "Images referenced by each container in the template."},
		{Name: "Selector", Type: "string", Priority: 1, Description: batchv1.JobSpec{}.SwaggerDoc()["selector"]},
	}
	h.TableHandler(cronJobColumnDefinitions, printCronJob)
	h.TableHandler(cronJobColumnDefinitions, printCronJobList)

	serviceColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Type", Type: "string", Description: apiv1.ServiceSpec{}.SwaggerDoc()["type"]},
		{Name: "Cluster-IP", Type: "string", Description: apiv1.ServiceSpec{}.SwaggerDoc()["clusterIP"]},
		{Name: "External-IP", Type: "string", Description: apiv1.ServiceSpec{}.SwaggerDoc()["externalIPs"]},
		{Name: "Port(s)", Type: "string", Description: apiv1.ServiceSpec{}.SwaggerDoc()["ports"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "Selector", Type: "string", Priority: 1, Description: apiv1.ServiceSpec{}.SwaggerDoc()["selector"]},
	}

	h.TableHandler(serviceColumnDefinitions, printService)
	h.TableHandler(serviceColumnDefinitions, printServiceList)

	ingressColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Hosts", Type: "string", Description: "Hosts that incoming requests are matched against before the ingress rule"},
		{Name: "Address", Type: "string", Description: "Address is a list containing ingress points for the load-balancer"},
		{Name: "Ports", Type: "string", Description: "Ports of TLS configurations that open"},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
	}
	h.TableHandler(ingressColumnDefinitions, printIngress)
	h.TableHandler(ingressColumnDefinitions, printIngressList)

	statefulSetColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Desired", Type: "string", Description: appsv1beta1.StatefulSetSpec{}.SwaggerDoc()["replicas"]},
		{Name: "Current", Type: "string", Description: appsv1beta1.StatefulSetStatus{}.SwaggerDoc()["replicas"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "Containers", Type: "string", Priority: 1, Description: "Names of each container in the template."},
		{Name: "Images", Type: "string", Priority: 1, Description: "Images referenced by each container in the template."},
	}
	h.TableHandler(statefulSetColumnDefinitions, printStatefulSet)
	h.TableHandler(statefulSetColumnDefinitions, printStatefulSetList)

	endpointColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Endpoints", Type: "string", Description: apiv1.Endpoints{}.SwaggerDoc()["subsets"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
	}
	h.TableHandler(endpointColumnDefinitions, printEndpoints)
	h.TableHandler(endpointColumnDefinitions, printEndpointsList)

	nodeColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Status", Type: "string", Description: "The status of the node"},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
		{Name: "Version", Type: "string", Description: apiv1.NodeSystemInfo{}.SwaggerDoc()["kubeletVersion"]},
		{Name: "External-IP", Type: "string", Priority: 1, Description: apiv1.NodeStatus{}.SwaggerDoc()["addresses"]},
		{Name: "OS-Image", Type: "string", Priority: 1, Description: apiv1.NodeSystemInfo{}.SwaggerDoc()["osImage"]},
		{Name: "Kernel-Version", Type: "string", Priority: 1, Description: apiv1.NodeSystemInfo{}.SwaggerDoc()["kernelVersion"]},
		{Name: "Container-Runtime", Type: "string", Priority: 1, Description: apiv1.NodeSystemInfo{}.SwaggerDoc()["containerRuntimeVersion"]},
	}

	h.TableHandler(nodeColumnDefinitions, printNode)
	h.TableHandler(nodeColumnDefinitions, printNodeList)

	eventColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Lastseen", Type: "string", Description: apiv1.Event{}.SwaggerDoc()["lastTimestamp"]},
		{Name: "Firtseen", Type: "string", Description: apiv1.Event{}.SwaggerDoc()["firstTimestamp"]},
		{Name: "Count", Type: "string", Description: apiv1.Event{}.SwaggerDoc()["count"]},
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Kind", Type: "string", Description: apiv1.Event{}.InvolvedObject.SwaggerDoc()["kind"]},
		{Name: "Subobject", Type: "string", Description: apiv1.Event{}.InvolvedObject.SwaggerDoc()["fieldPath"]},
		{Name: "Type", Type: "string", Description: apiv1.Event{}.SwaggerDoc()["type"]},
		{Name: "Reason", Type: "string", Description: apiv1.Event{}.SwaggerDoc()["reason"]},
		{Name: "Source", Type: "string", Description: apiv1.Event{}.SwaggerDoc()["source"]},
		{Name: "Message", Type: "string", Description: apiv1.Event{}.SwaggerDoc()["message"]},
	}
	h.TableHandler(eventColumnDefinitions, printEvent)
	h.TableHandler(eventColumnDefinitions, printEventList)

	namespaceColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Status", Type: "string", Description: "The status of the namespace"},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
	}
	h.TableHandler(namespaceColumnDefinitions, printNamespace)
	h.TableHandler(namespaceColumnDefinitions, printNamespaceList)

	secretColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Type", Type: "string", Description: apiv1.Secret{}.SwaggerDoc()["type"]},
		{Name: "Data", Type: "string", Description: apiv1.Secret{}.SwaggerDoc()["data"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
	}
	h.TableHandler(secretColumnDefinitions, printSecret)
	h.TableHandler(secretColumnDefinitions, printSecretList)

	serviceAccountColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Secrets", Type: "string", Description: apiv1.ServiceAccount{}.SwaggerDoc()["secrets"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
	}
	h.TableHandler(serviceAccountColumnDefinitions, printServiceAccount)
	h.TableHandler(serviceAccountColumnDefinitions, printServiceAccountList)

	persistentVolumeColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Capacity", Type: "string", Description: apiv1.PersistentVolumeSpec{}.SwaggerDoc()["capacity"]},
		{Name: "AccessModes", Type: "string", Description: apiv1.PersistentVolumeSpec{}.SwaggerDoc()["accessModes"]},
		{Name: "ReclaimPolicy", Type: "string", Description: apiv1.PersistentVolumeSpec{}.SwaggerDoc()["persistentVolumeReclaimPolicy"]},
		{Name: "Status", Type: "string", Description: apiv1.PersistentVolumeStatus{}.SwaggerDoc()["phase"]},
		{Name: "Claim", Type: "string", Description: apiv1.PersistentVolumeSpec{}.SwaggerDoc()["claimRef"]},
		{Name: "StorageClass", Type: "string", Description: "StorageClass of the pv"},
		{Name: "Reason", Type: "string", Description: apiv1.PersistentVolumeStatus{}.SwaggerDoc()["reason"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
	}
	h.TableHandler(persistentVolumeColumnDefinitions, printPersistentVolume)
	h.TableHandler(persistentVolumeColumnDefinitions, printPersistentVolumeList)

	h.Handler(persistentVolumeClaimColumns, nil, printPersistentVolumeClaim)
	h.Handler(persistentVolumeClaimColumns, nil, printPersistentVolumeClaimList)

	h.Handler(componentStatusColumns, nil, printComponentStatus)
	h.Handler(componentStatusColumns, nil, printComponentStatusList)
	h.Handler(thirdPartyResourceColumns, nil, printThirdPartyResource)
	h.Handler(thirdPartyResourceColumns, nil, printThirdPartyResourceList)
	h.Handler(deploymentColumns, deploymentWideColumns, printDeployment)
	h.Handler(deploymentColumns, deploymentWideColumns, printDeploymentList)
	h.Handler(horizontalPodAutoscalerColumns, nil, printHorizontalPodAutoscaler)
	h.Handler(horizontalPodAutoscalerColumns, nil, printHorizontalPodAutoscalerList)
	h.Handler(configMapColumns, nil, printConfigMap)
	h.Handler(configMapColumns, nil, printConfigMapList)
	h.Handler(podSecurityPolicyColumns, nil, printPodSecurityPolicy)
	h.Handler(podSecurityPolicyColumns, nil, printPodSecurityPolicyList)
	h.Handler(clusterColumns, nil, printCluster)
	h.Handler(clusterColumns, nil, printClusterList)
	h.Handler(networkPolicyColumns, nil, printExtensionsNetworkPolicy)
	h.Handler(networkPolicyColumns, nil, printExtensionsNetworkPolicyList)
	h.Handler(networkPolicyColumns, nil, printNetworkPolicy)
	h.Handler(networkPolicyColumns, nil, printNetworkPolicyList)
	h.Handler(roleBindingColumns, roleBindingWideColumns, printRoleBinding)
	h.Handler(roleBindingColumns, roleBindingWideColumns, printRoleBindingList)
	h.Handler(clusterRoleBindingColumns, clusterRoleBindingWideColumns, printClusterRoleBinding)
	h.Handler(clusterRoleBindingColumns, clusterRoleBindingWideColumns, printClusterRoleBindingList)
	h.Handler(certificateSigningRequestColumns, nil, printCertificateSigningRequest)
	h.Handler(certificateSigningRequestColumns, nil, printCertificateSigningRequestList)
	h.Handler(storageClassColumns, nil, printStorageClass)
	h.Handler(storageClassColumns, nil, printStorageClassList)
	h.Handler(statusColumns, nil, printStatus)
	h.Handler(controllerRevisionColumns, nil, printControllerRevision)
	h.Handler(controllerRevisionColumns, nil, printControllerRevisionList)

	AddDefaultHandlers(h)
}

// AddDefaultHandlers adds handlers that can work with most Kubernetes objects.
func AddDefaultHandlers(h printers.PrintHandler) {
	// types without defined columns
	objectMetaColumnDefinitions := []metav1alpha1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: metav1.ObjectMeta{}.SwaggerDoc()["name"]},
		{Name: "Age", Type: "string", Description: metav1.ObjectMeta{}.SwaggerDoc()["creationTimestamp"]},
	}
	h.DefaultTableHandler(objectMetaColumnDefinitions, printObjectMeta)
}

func printObjectMeta(obj runtime.Object, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	if meta.IsListType(obj) {
		rows := make([]metav1alpha1.TableRow, 0, 16)
		err := meta.EachListItem(obj, func(obj runtime.Object) error {
			nestedRows, err := printObjectMeta(obj, options)
			if err != nil {
				return err
			}
			rows = append(rows, nestedRows...)
			return nil
		})
		if err != nil {
			return nil, err
		}
		return rows, nil
	}

	rows := make([]metav1alpha1.TableRow, 0, 1)
	m, err := meta.Accessor(obj)
	if err != nil {
		return nil, err
	}
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	row.Cells = append(row.Cells, m.GetName(), translateTimestamp(m.GetCreationTimestamp()))
	rows = append(rows, row)
	return rows, nil
}

// Pass ports=nil for all ports.
func formatEndpoints(endpoints *api.Endpoints, ports sets.String) string {
	if len(endpoints.Subsets) == 0 {
		return "<none>"
	}
	list := []string{}
	max := 3
	more := false
	count := 0
	for i := range endpoints.Subsets {
		ss := &endpoints.Subsets[i]
		for i := range ss.Ports {
			port := &ss.Ports[i]
			if ports == nil || ports.Has(port.Name) {
				for i := range ss.Addresses {
					if len(list) == max {
						more = true
					}
					addr := &ss.Addresses[i]
					if !more {
						hostPort := net.JoinHostPort(addr.IP, strconv.Itoa(int(port.Port)))
						list = append(list, hostPort)
					}
					count++
				}
			}
		}
	}
	ret := strings.Join(list, ",")
	if more {
		return fmt.Sprintf("%s + %d more...", ret, count-max)
	}
	return ret
}

// translateTimestamp returns the elapsed time since timestamp in
// human-readable approximation.
func translateTimestamp(timestamp metav1.Time) string {
	if timestamp.IsZero() {
		return "<unknown>"
	}
	return printers.ShortHumanDuration(time.Now().Sub(timestamp.Time))
}

var (
	podSuccessConditions = []metav1alpha1.TableRowCondition{{Type: metav1alpha1.RowCompleted, Status: metav1alpha1.ConditionTrue, Reason: string(api.PodSucceeded), Message: "The pod has completed successfully."}}
	podFailedConditions  = []metav1alpha1.TableRowCondition{{Type: metav1alpha1.RowCompleted, Status: metav1alpha1.ConditionTrue, Reason: string(api.PodFailed), Message: "The pod failed."}}
)

func printPodList(podList *api.PodList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(podList.Items))
	for i := range podList.Items {
		r, err := printPod(&podList.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printPod(pod *api.Pod, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	restarts := 0
	totalContainers := len(pod.Spec.Containers)
	readyContainers := 0

	reason := string(pod.Status.Phase)
	if pod.Status.Reason != "" {
		reason = pod.Status.Reason
	}

	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: pod},
	}

	switch pod.Status.Phase {
	case api.PodSucceeded:
		row.Conditions = podSuccessConditions
	case api.PodFailed:
		row.Conditions = podFailedConditions
	}

	initializing := false
	for i := range pod.Status.InitContainerStatuses {
		container := pod.Status.InitContainerStatuses[i]
		restarts += int(container.RestartCount)
		switch {
		case container.State.Terminated != nil && container.State.Terminated.ExitCode == 0:
			continue
		case container.State.Terminated != nil:
			// initialization is failed
			if len(container.State.Terminated.Reason) == 0 {
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Init:Signal:%d", container.State.Terminated.Signal)
				} else {
					reason = fmt.Sprintf("Init:ExitCode:%d", container.State.Terminated.ExitCode)
				}
			} else {
				reason = "Init:" + container.State.Terminated.Reason
			}
			initializing = true
		case container.State.Waiting != nil && len(container.State.Waiting.Reason) > 0 && container.State.Waiting.Reason != "PodInitializing":
			reason = "Init:" + container.State.Waiting.Reason
			initializing = true
		default:
			reason = fmt.Sprintf("Init:%d/%d", i, len(pod.Spec.InitContainers))
			initializing = true
		}
		break
	}
	if !initializing {
		restarts = 0
		for i := len(pod.Status.ContainerStatuses) - 1; i >= 0; i-- {
			container := pod.Status.ContainerStatuses[i]

			restarts += int(container.RestartCount)
			if container.State.Waiting != nil && container.State.Waiting.Reason != "" {
				reason = container.State.Waiting.Reason
			} else if container.State.Terminated != nil && container.State.Terminated.Reason != "" {
				reason = container.State.Terminated.Reason
			} else if container.State.Terminated != nil && container.State.Terminated.Reason == "" {
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Signal:%d", container.State.Terminated.Signal)
				} else {
					reason = fmt.Sprintf("ExitCode:%d", container.State.Terminated.ExitCode)
				}
			} else if container.Ready && container.State.Running != nil {
				readyContainers++
			}
		}
	}

	if pod.DeletionTimestamp != nil && pod.Status.Reason == node.NodeUnreachablePodReason {
		reason = "Unknown"
	} else if pod.DeletionTimestamp != nil {
		reason = "Terminating"
	}

	row.Cells = append(row.Cells, pod.Name, fmt.Sprintf("%d/%d", readyContainers, totalContainers), reason, restarts, translateTimestamp(pod.CreationTimestamp))

	if options.Wide {
		nodeName := pod.Spec.NodeName
		podIP := pod.Status.PodIP
		if podIP == "" {
			podIP = "<none>"
		}
		if nodeName == "" {
			nodeName = "<none>"
		}
		row.Cells = append(row.Cells, podIP, nodeName)
	}

	return []metav1alpha1.TableRow{row}, nil
}

func printPodTemplate(obj *api.PodTemplate, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	names, images := layoutContainerCells(obj.Template.Spec.Containers)
	row.Cells = append(row.Cells, obj.Name, names, images, labels.FormatLabels(obj.Template.Labels))
	return []metav1alpha1.TableRow{row}, nil
}

func printPodTemplateList(list *api.PodTemplateList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printPodTemplate(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printPodDisruptionBudget(obj *policy.PodDisruptionBudget, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}

	var minAvailable string
	var maxUnavailable string
	if obj.Spec.MinAvailable != nil {
		minAvailable = obj.Spec.MinAvailable.String()
	} else {
		minAvailable = "N/A"
	}

	if obj.Spec.MaxUnavailable != nil {
		maxUnavailable = obj.Spec.MaxUnavailable.String()
	} else {
		maxUnavailable = "N/A"
	}

	row.Cells = append(row.Cells, obj.Name, minAvailable, maxUnavailable, obj.Status.PodDisruptionsAllowed, translateTimestamp(obj.CreationTimestamp))
	return []metav1alpha1.TableRow{row}, nil
}

func printPodDisruptionBudgetList(list *policy.PodDisruptionBudgetList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printPodDisruptionBudget(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

// TODO(AdoHe): try to put wide output in a single method
func printReplicationController(obj *api.ReplicationController, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}

	desiredReplicas := obj.Spec.Replicas
	currentReplicas := obj.Status.Replicas
	readyReplicas := obj.Status.ReadyReplicas

	row.Cells = append(row.Cells, obj.Name, desiredReplicas, currentReplicas, readyReplicas, translateTimestamp(obj.CreationTimestamp))
	if options.Wide {
		names, images := layoutContainerCells(obj.Spec.Template.Spec.Containers)
		row.Cells = append(row.Cells, names, images, labels.FormatLabels(obj.Spec.Selector))
	}
	return []metav1alpha1.TableRow{row}, nil
}

func printReplicationControllerList(list *api.ReplicationControllerList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printReplicationController(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printReplicaSet(obj *extensions.ReplicaSet, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}

	desiredReplicas := obj.Spec.Replicas
	currentReplicas := obj.Status.Replicas
	readyReplicas := obj.Status.ReadyReplicas

	row.Cells = append(row.Cells, obj.Name, desiredReplicas, currentReplicas, readyReplicas, translateTimestamp(obj.CreationTimestamp))
	if options.Wide {
		names, images := layoutContainerCells(obj.Spec.Template.Spec.Containers)
		row.Cells = append(row.Cells, names, images, metav1.FormatLabelSelector(obj.Spec.Selector))
	}
	return []metav1alpha1.TableRow{row}, nil
}

func printReplicaSetList(list *extensions.ReplicaSetList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printReplicaSet(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printCluster(c *federation.Cluster, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, c.Name, options.WithKind)

	var statuses []string
	for _, condition := range c.Status.Conditions {
		if condition.Status == api.ConditionTrue {
			statuses = append(statuses, string(condition.Type))
		} else {
			statuses = append(statuses, "Not"+string(condition.Type))
		}
	}
	if len(statuses) == 0 {
		statuses = append(statuses, "Unknown")
	}

	if _, err := fmt.Fprintf(w, "%s\t%s\t%s\n",
		name,
		strings.Join(statuses, ","),
		translateTimestamp(c.CreationTimestamp),
	); err != nil {
		return err
	}
	return nil
}
func printClusterList(list *federation.ClusterList, w io.Writer, options printers.PrintOptions) error {
	for _, rs := range list.Items {
		if err := printCluster(&rs, w, options); err != nil {
			return err
		}
	}
	return nil
}

func printJob(obj *batch.Job, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}

	var completions string
	if obj.Spec.Completions != nil {
		completions = strconv.Itoa(int(*obj.Spec.Completions))
	} else {
		completions = "<none>"
	}

	row.Cells = append(row.Cells, obj.Name, completions, obj.Status.Succeeded, translateTimestamp(obj.CreationTimestamp))
	if options.Wide {
		names, images := layoutContainerCells(obj.Spec.Template.Spec.Containers)
		row.Cells = append(row.Cells, names, images, metav1.FormatLabelSelector(obj.Spec.Selector))
	}
	return []metav1alpha1.TableRow{row}, nil
}

func printJobList(list *batch.JobList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printJob(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printCronJob(obj *batch.CronJob, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}

	lastScheduleTime := "<none>"
	if obj.Status.LastScheduleTime != nil {
		lastScheduleTime = obj.Status.LastScheduleTime.Time.Format(time.RFC1123Z)
	}

	row.Cells = append(row.Cells, obj.Name, obj.Spec.Schedule, printBoolPtr(obj.Spec.Suspend), len(obj.Status.Active), lastScheduleTime)
	if options.Wide {
		names, images := layoutContainerCells(obj.Spec.JobTemplate.Spec.Template.Spec.Containers)
		row.Cells = append(row.Cells, names, images, metav1.FormatLabelSelector(obj.Spec.JobTemplate.Spec.Selector))
	}
	return []metav1alpha1.TableRow{row}, nil
}

func printCronJobList(list *batch.CronJobList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printCronJob(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

// loadBalancerStatusStringer behaves mostly like a string interface and converts the given status to a string.
// `wide` indicates whether the returned value is meant for --o=wide output. If not, it's clipped to 16 bytes.
func loadBalancerStatusStringer(s api.LoadBalancerStatus, wide bool) string {
	ingress := s.Ingress
	result := []string{}
	for i := range ingress {
		if ingress[i].IP != "" {
			result = append(result, ingress[i].IP)
		} else if ingress[i].Hostname != "" {
			result = append(result, ingress[i].Hostname)
		}
	}
	r := strings.Join(result, ",")
	if !wide && len(r) > loadBalancerWidth {
		r = r[0:(loadBalancerWidth-3)] + "..."
	}
	return r
}

func getServiceExternalIP(svc *api.Service, wide bool) string {
	switch svc.Spec.Type {
	case api.ServiceTypeClusterIP:
		if len(svc.Spec.ExternalIPs) > 0 {
			return strings.Join(svc.Spec.ExternalIPs, ",")
		}
		return "<none>"
	case api.ServiceTypeNodePort:
		if len(svc.Spec.ExternalIPs) > 0 {
			return strings.Join(svc.Spec.ExternalIPs, ",")
		}
		return "<none>"
	case api.ServiceTypeLoadBalancer:
		lbIps := loadBalancerStatusStringer(svc.Status.LoadBalancer, wide)
		if len(svc.Spec.ExternalIPs) > 0 {
			results := []string{}
			if len(lbIps) > 0 {
				results = append(results, strings.Split(lbIps, ",")...)
			}
			results = append(results, svc.Spec.ExternalIPs...)
			return strings.Join(results, ",")
		}
		if len(lbIps) > 0 {
			return lbIps
		}
		return "<pending>"
	case api.ServiceTypeExternalName:
		return svc.Spec.ExternalName
	}
	return "<unknown>"
}

func makePortString(ports []api.ServicePort) string {
	pieces := make([]string, len(ports))
	for ix := range ports {
		port := &ports[ix]
		pieces[ix] = fmt.Sprintf("%d/%s", port.Port, port.Protocol)
		if port.NodePort > 0 {
			pieces[ix] = fmt.Sprintf("%d:%d/%s", port.Port, port.NodePort, port.Protocol)
		}
	}
	return strings.Join(pieces, ",")
}

func printService(obj *api.Service, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	svcType := obj.Spec.Type
	internalIP := obj.Spec.ClusterIP
	if len(internalIP) == 0 {
		internalIP = "<none>"
	}
	externalIP := getServiceExternalIP(obj, options.Wide)
	svcPorts := makePortString(obj.Spec.Ports)
	if len(svcPorts) == 0 {
		svcPorts = "<none>"
	}

	row.Cells = append(row.Cells, obj.Name, string(svcType), internalIP, externalIP, svcPorts, translateTimestamp(obj.CreationTimestamp))
	if options.Wide {
		row.Cells = append(row.Cells, labels.FormatLabels(obj.Spec.Selector))
	}

	return []metav1alpha1.TableRow{row}, nil
}

func printServiceList(list *api.ServiceList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printService(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

// backendStringer behaves just like a string interface and converts the given backend to a string.
func backendStringer(backend *extensions.IngressBackend) string {
	if backend == nil {
		return ""
	}
	return fmt.Sprintf("%v:%v", backend.ServiceName, backend.ServicePort.String())
}

func formatHosts(rules []extensions.IngressRule) string {
	list := []string{}
	max := 3
	more := false
	for _, rule := range rules {
		if len(list) == max {
			more = true
		}
		if !more && len(rule.Host) != 0 {
			list = append(list, rule.Host)
		}
	}
	if len(list) == 0 {
		return "*"
	}
	ret := strings.Join(list, ",")
	if more {
		return fmt.Sprintf("%s + %d more...", ret, len(rules)-max)
	}
	return ret
}

func formatPorts(tls []extensions.IngressTLS) string {
	if len(tls) != 0 {
		return "80, 443"
	}
	return "80"
}

func printIngress(obj *extensions.Ingress, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	hosts := formatHosts(obj.Spec.Rules)
	address := loadBalancerStatusStringer(obj.Status.LoadBalancer, options.Wide)
	ports := formatPorts(obj.Spec.TLS)
	createTime := translateTimestamp(obj.CreationTimestamp)
	row.Cells = append(row.Cells, obj.Name, hosts, address, ports, createTime)
	return []metav1alpha1.TableRow{row}, nil
}

func printIngressList(list *extensions.IngressList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printIngress(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printStatefulSet(obj *apps.StatefulSet, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	desiredReplicas := obj.Spec.Replicas
	currentReplicas := obj.Status.Replicas
	createTime := translateTimestamp(obj.CreationTimestamp)
	row.Cells = append(row.Cells, obj.Name, desiredReplicas, currentReplicas, createTime)
	if options.Wide {
		names, images := layoutContainerCells(obj.Spec.Template.Spec.Containers)
		row.Cells = append(row.Cells, names, images)
	}
	return []metav1alpha1.TableRow{row}, nil
}

func printStatefulSetList(list *apps.StatefulSetList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printStatefulSet(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printDaemonSet(obj *extensions.DaemonSet, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}

	desiredScheduled := obj.Status.DesiredNumberScheduled
	currentScheduled := obj.Status.CurrentNumberScheduled
	numberReady := obj.Status.NumberReady
	numberUpdated := obj.Status.UpdatedNumberScheduled
	numberAvailable := obj.Status.NumberAvailable

	row.Cells = append(row.Cells, obj.Name, desiredScheduled, currentScheduled, numberReady, numberUpdated, numberAvailable, labels.FormatLabels(obj.Spec.Template.Spec.NodeSelector), translateTimestamp(obj.CreationTimestamp))
	if options.Wide {
		names, images := layoutContainerCells(obj.Spec.Template.Spec.Containers)
		row.Cells = append(row.Cells, names, images, metav1.FormatLabelSelector(obj.Spec.Selector))
	}
	return []metav1alpha1.TableRow{row}, nil
}

func printDaemonSetList(list *extensions.DaemonSetList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printDaemonSet(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printEndpoints(obj *api.Endpoints, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	row.Cells = append(row.Cells, obj.Name, formatEndpoints(obj, nil), translateTimestamp(obj.CreationTimestamp))
	return []metav1alpha1.TableRow{row}, nil
}

func printEndpointsList(list *api.EndpointsList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printEndpoints(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printNamespace(obj *api.Namespace, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	row.Cells = append(row.Cells, obj.Name, obj.Status.Phase, translateTimestamp(obj.CreationTimestamp))
	return []metav1alpha1.TableRow{row}, nil
}

func printNamespaceList(list *api.NamespaceList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printNamespace(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printSecret(obj *api.Secret, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	row.Cells = append(row.Cells, obj.Name, obj.Type, len(obj.Data), translateTimestamp(obj.CreationTimestamp))
	return []metav1alpha1.TableRow{row}, nil
}

func printSecretList(list *api.SecretList, w io.Writer, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printSecret(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printServiceAccount(obj *api.ServiceAccount, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	row.Cells = append(row.Cells, obj.Name, len(obj.Secrets), translateTimestamp(obj.CreationTimestamp))
	return []metav1alpha1.TableRow{row}, nil
}

func printServiceAccountList(list *api.ServiceAccountList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printServiceAccount(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printNode(obj *api.Node, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}

	conditionMap := make(map[api.NodeConditionType]*api.NodeCondition)
	NodeAllConditions := []api.NodeConditionType{api.NodeReady}
	for i := range obj.Status.Conditions {
		cond := obj.Status.Conditions[i]
		conditionMap[cond.Type] = &cond
	}
	var status []string
	for _, validCondition := range NodeAllConditions {
		if condition, ok := conditionMap[validCondition]; ok {
			if condition.Status == api.ConditionTrue {
				status = append(status, string(condition.Type))
			} else {
				status = append(status, "Not"+string(condition.Type))
			}
		}
	}
	if len(status) == 0 {
		status = append(status, "Unknown")
	}
	if obj.Spec.Unschedulable {
		status = append(status, "SchedulingDisabled")
	}

	row.Cells = append(row.Cells, obj.Name, strings.Join(status, ","), translateTimestamp(obj.CreationTimestamp), obj.Status.NodeInfo.KubeletVersion)
	if options.Wide {
		osImage, kernelVersion, crVersion := obj.Status.NodeInfo.OSImage, obj.Status.NodeInfo.KernelVersion, obj.Status.NodeInfo.ContainerRuntimeVersion
		if osImage == "" {
			osImage = "<unknown>"
		}
		if kernelVersion == "" {
			kernelVersion = "<unknown>"
		}
		if crVersion == "" {
			crVersion = "<unknown>"
		}
		row.Cells = append(row.Cells, getNodeExternalIP(obj), osImage, kernelVersion, crVersion)
	}

	return []metav1alpha1.TableRow{row}, nil
}

// Returns first external ip of the node or "<none>" if none is found.
func getNodeExternalIP(node *api.Node) string {
	for _, address := range node.Status.Addresses {
		if address.Type == api.NodeExternalIP {
			return address.Address
		}
	}

	return "<none>"
}

func printNodeList(list *api.NodeList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printNode(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printPersistentVolume(obj *api.PersistentVolume, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}

	claimRefUID := ""
	if obj.Spec.ClaimRef != nil {
		claimRefUID += obj.Spec.ClaimRef.Namespace
		claimRefUID += "/"
		claimRefUID += obj.Spec.ClaimRef.Name
	}

	modesStr := helper.GetAccessModesAsString(obj.Spec.AccessModes)
	reclaimPolicyStr := string(obj.Spec.PersistentVolumeReclaimPolicy)

	aQty := obj.Spec.Capacity[api.ResourceStorage]
	aSize := aQty.String()

	row.Cells = append(row.Cells, obj.Name, aSize, modesStr, reclaimPolicyStr,
		obj.Status.Phase, claimRefUID, helper.GetPersistentVolumeClass(obj),
		obj.Status.Reason,
		translateTimestamp(obj.CreationTimestamp))
	return []metav1alpha1.TableRow{row}, nil
}

func printPersistentVolumeList(list *api.PersistentVolumeList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printPersistentVolume(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printPersistentVolumeClaim(pvc *api.PersistentVolumeClaim, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, pvc.Name, options.WithKind)

	namespace := pvc.Namespace

	if options.WithNamespace {
		if _, err := fmt.Fprintf(w, "%s\t", namespace); err != nil {
			return err
		}
	}

	phase := pvc.Status.Phase
	storage := pvc.Spec.Resources.Requests[api.ResourceStorage]
	capacity := ""
	accessModes := ""
	if pvc.Spec.VolumeName != "" {
		accessModes = helper.GetAccessModesAsString(pvc.Status.AccessModes)
		storage = pvc.Status.Capacity[api.ResourceStorage]
		capacity = storage.String()
	}

	if _, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s", name, phase, pvc.Spec.VolumeName, capacity, accessModes, helper.GetPersistentVolumeClaimClass(pvc), translateTimestamp(pvc.CreationTimestamp)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendLabels(pvc.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, pvc.Labels))
	return err
}

func printPersistentVolumeClaimList(list *api.PersistentVolumeClaimList, w io.Writer, options printers.PrintOptions) error {
	for _, psd := range list.Items {
		if err := printPersistentVolumeClaim(&psd, w, options); err != nil {
			return err
		}
	}
	return nil
}

func printEvent(obj *api.Event, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {
	row := metav1alpha1.TableRow{
		Object: runtime.RawExtension{Object: obj},
	}
	// While watching event, we should print absolute time.
	var FirstTimestamp, LastTimestamp string
	if options.AbsoluteTimestamps {
		FirstTimestamp = obj.FirstTimestamp.String()
		LastTimestamp = obj.LastTimestamp.String()
	} else {
		FirstTimestamp = translateTimestamp(obj.FirstTimestamp)
		LastTimestamp = translateTimestamp(obj.LastTimestamp)
	}
	row.Cells = append(row.Cells, LastTimestamp, FirstTimestamp,
		obj.Count, obj.Name, obj.InvolvedObject.Kind,
		obj.InvolvedObject.FieldPath, obj.Type, obj.Reason,
		formatEventSource(obj.Source), obj.Message)

	return []metav1alpha1.TableRow{row}, nil
}

// Sorts and prints the EventList in a human-friendly format.
func printEventList(list *api.EventList, options printers.PrintOptions) ([]metav1alpha1.TableRow, error) {

	sort.Sort(events.SortableEvents(list.Items))
	rows := make([]metav1alpha1.TableRow, 0, len(list.Items))
	for i := range list.Items {
		r, err := printEvent(&list.Items[i], options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

func printRoleBinding(roleBinding *rbac.RoleBinding, w io.Writer, options printers.PrintOptions) error {
	meta := roleBinding.ObjectMeta
	name := printers.FormatResourceName(options.Kind, meta.Name, options.WithKind)

	if options.WithNamespace {
		if _, err := fmt.Fprintf(w, "%s\t", meta.Namespace); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(
		w, "%s\t%s",
		name,
		translateTimestamp(meta.CreationTimestamp),
	); err != nil {
		return err
	}

	if options.Wide {
		roleRef := fmt.Sprintf("%s/%s", roleBinding.RoleRef.Kind, roleBinding.RoleRef.Name)
		users, groups, sas, _ := rbac.SubjectsStrings(roleBinding.Subjects)
		if _, err := fmt.Fprintf(w, "\t%s\t%v\t%v\t%v",
			roleRef,
			strings.Join(users, ", "),
			strings.Join(groups, ", "),
			strings.Join(sas, ", "),
		); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, printers.AppendLabels(meta.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, meta.Labels))
	return err
}

// Prints the RoleBinding in a human-friendly format.
func printRoleBindingList(list *rbac.RoleBindingList, w io.Writer, options printers.PrintOptions) error {
	for i := range list.Items {
		if err := printRoleBinding(&list.Items[i], w, options); err != nil {
			return err
		}
	}
	return nil
}

func printClusterRoleBinding(clusterRoleBinding *rbac.ClusterRoleBinding, w io.Writer, options printers.PrintOptions) error {
	meta := clusterRoleBinding.ObjectMeta
	name := printers.FormatResourceName(options.Kind, meta.Name, options.WithKind)

	if options.WithNamespace {
		return fmt.Errorf("clusterRoleBinding is not namespaced")
	}

	if _, err := fmt.Fprintf(
		w, "%s\t%s",
		name,
		translateTimestamp(meta.CreationTimestamp),
	); err != nil {
		return err
	}

	if options.Wide {
		roleRef := clusterRoleBinding.RoleRef.Name
		users, groups, sas, _ := rbac.SubjectsStrings(clusterRoleBinding.Subjects)
		if _, err := fmt.Fprintf(w, "\t%s\t%v\t%v\t%v",
			roleRef,
			strings.Join(users, ", "),
			strings.Join(groups, ", "),
			strings.Join(sas, ", "),
		); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, printers.AppendLabels(meta.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, meta.Labels))
	return err
}

// Prints the ClusterRoleBinding in a human-friendly format.
func printClusterRoleBindingList(list *rbac.ClusterRoleBindingList, w io.Writer, options printers.PrintOptions) error {
	for i := range list.Items {
		if err := printClusterRoleBinding(&list.Items[i], w, options); err != nil {
			return err
		}
	}
	return nil
}

func printCertificateSigningRequest(csr *certificates.CertificateSigningRequest, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, csr.Name, options.WithKind)
	meta := csr.ObjectMeta

	status, err := extractCSRStatus(csr)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(
		w, "%s\t%s\t%s\t%s",
		name,
		translateTimestamp(meta.CreationTimestamp),
		csr.Spec.Username,
		status,
	); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendLabels(meta.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err = fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, meta.Labels))
	return err
}

func extractCSRStatus(csr *certificates.CertificateSigningRequest) (string, error) {
	var approved, denied bool
	for _, c := range csr.Status.Conditions {
		switch c.Type {
		case certificates.CertificateApproved:
			approved = true
		case certificates.CertificateDenied:
			denied = true
		default:
			return "", fmt.Errorf("unknown csr condition %q", c)
		}
	}
	var status string
	// must be in order of presidence
	if denied {
		status += "Denied"
	} else if approved {
		status += "Approved"
	} else {
		status += "Pending"
	}
	if len(csr.Status.Certificate) > 0 {
		status += ",Issued"
	}
	return status, nil
}

func printCertificateSigningRequestList(list *certificates.CertificateSigningRequestList, w io.Writer, options printers.PrintOptions) error {
	for i := range list.Items {
		if err := printCertificateSigningRequest(&list.Items[i], w, options); err != nil {
			return err
		}
	}
	return nil
}

func printComponentStatus(item *api.ComponentStatus, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, item.Name, options.WithKind)

	if options.WithNamespace {
		return fmt.Errorf("componentStatus is not namespaced")
	}
	status := "Unknown"
	message := ""
	error := ""
	for _, condition := range item.Conditions {
		if condition.Type == api.ComponentHealthy {
			if condition.Status == api.ConditionTrue {
				status = "Healthy"
			} else {
				status = "Unhealthy"
			}
			message = condition.Message
			error = condition.Error
			break
		}
	}

	if _, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s", name, status, message, error); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendLabels(item.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, item.Labels))
	return err
}

func printComponentStatusList(list *api.ComponentStatusList, w io.Writer, options printers.PrintOptions) error {
	for _, item := range list.Items {
		if err := printComponentStatus(&item, w, options); err != nil {
			return err
		}
	}

	return nil
}

func printThirdPartyResource(rsrc *extensions.ThirdPartyResource, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, rsrc.Name, options.WithKind)

	versions := make([]string, len(rsrc.Versions))
	for ix := range rsrc.Versions {
		version := &rsrc.Versions[ix]
		versions[ix] = fmt.Sprintf("%s", version.Name)
	}
	versionsString := strings.Join(versions, ",")
	if _, err := fmt.Fprintf(w, "%s\t%s\t%s\n", name, rsrc.Description, versionsString); err != nil {
		return err
	}
	return nil
}

func printThirdPartyResourceList(list *extensions.ThirdPartyResourceList, w io.Writer, options printers.PrintOptions) error {
	for _, item := range list.Items {
		if err := printThirdPartyResource(&item, w, options); err != nil {
			return err
		}
	}

	return nil
}

func truncate(str string, maxLen int) string {
	if len(str) > maxLen {
		return str[0:maxLen] + "..."
	}
	return str
}

func printDeployment(deployment *extensions.Deployment, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, deployment.Name, options.WithKind)

	if options.WithNamespace {
		if _, err := fmt.Fprintf(w, "%s\t", deployment.Namespace); err != nil {
			return err
		}
	}

	desiredReplicas := deployment.Spec.Replicas
	currentReplicas := deployment.Status.Replicas
	updatedReplicas := deployment.Status.UpdatedReplicas
	availableReplicas := deployment.Status.AvailableReplicas
	age := translateTimestamp(deployment.CreationTimestamp)
	containers := deployment.Spec.Template.Spec.Containers
	selector, err := metav1.LabelSelectorAsSelector(deployment.Spec.Selector)
	if err != nil {
		// this shouldn't happen if LabelSelector passed validation
		return err
	}

	if _, err := fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%d\t%s", name, desiredReplicas, currentReplicas, updatedReplicas, availableReplicas, age); err != nil {
		return err
	}
	if options.Wide {
		if err := layoutContainers(containers, w); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(w, "\t%s", selector.String()); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, printers.AppendLabels(deployment.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err = fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, deployment.Labels))
	return err
}

func printDeploymentList(list *extensions.DeploymentList, w io.Writer, options printers.PrintOptions) error {
	for _, item := range list.Items {
		if err := printDeployment(&item, w, options); err != nil {
			return err
		}
	}
	return nil
}

func formatHPAMetrics(specs []autoscaling.MetricSpec, statuses []autoscaling.MetricStatus) string {
	if len(specs) == 0 {
		return "<none>"
	}
	list := []string{}
	max := 2
	more := false
	count := 0
	for i, spec := range specs {
		switch spec.Type {
		case autoscaling.PodsMetricSourceType:
			current := "<unknown>"
			if len(statuses) > i && statuses[i].Pods != nil {
				current = statuses[i].Pods.CurrentAverageValue.String()
			}
			list = append(list, fmt.Sprintf("%s / %s", current, spec.Pods.TargetAverageValue.String()))
		case autoscaling.ObjectMetricSourceType:
			current := "<unknown>"
			if len(statuses) > i && statuses[i].Object != nil {
				current = statuses[i].Object.CurrentValue.String()
			}
			list = append(list, fmt.Sprintf("%s / %s", current, spec.Object.TargetValue.String()))
		case autoscaling.ResourceMetricSourceType:
			if spec.Resource.TargetAverageValue != nil {
				current := "<unknown>"
				if len(statuses) > i && statuses[i].Resource != nil {
					current = statuses[i].Resource.CurrentAverageValue.String()
				}
				list = append(list, fmt.Sprintf("%s / %s", current, spec.Resource.TargetAverageValue.String()))
			} else {
				current := "<unknown>"
				if len(statuses) > i && statuses[i].Resource != nil && statuses[i].Resource.CurrentAverageUtilization != nil {
					current = fmt.Sprintf("%d%%", *statuses[i].Resource.CurrentAverageUtilization)
				}

				target := "<auto>"
				if spec.Resource.TargetAverageUtilization != nil {
					target = fmt.Sprintf("%d%%", *spec.Resource.TargetAverageUtilization)
				}
				list = append(list, fmt.Sprintf("%s / %s", current, target))
			}
		default:
			list = append(list, "<unknown type>")
		}

		count++
	}

	if count > max {
		list = list[:max]
		more = true
	}

	ret := strings.Join(list, ", ")
	if more {
		return fmt.Sprintf("%s + %d more...", ret, count-max)
	}
	return ret
}

func printHorizontalPodAutoscaler(hpa *autoscaling.HorizontalPodAutoscaler, w io.Writer, options printers.PrintOptions) error {
	namespace := hpa.Namespace
	name := printers.FormatResourceName(options.Kind, hpa.Name, options.WithKind)

	reference := fmt.Sprintf("%s/%s",
		hpa.Spec.ScaleTargetRef.Kind,
		hpa.Spec.ScaleTargetRef.Name)
	minPods := "<unset>"
	metrics := formatHPAMetrics(hpa.Spec.Metrics, hpa.Status.CurrentMetrics)
	if hpa.Spec.MinReplicas != nil {
		minPods = fmt.Sprintf("%d", *hpa.Spec.MinReplicas)
	}
	maxPods := hpa.Spec.MaxReplicas
	currentReplicas := hpa.Status.CurrentReplicas

	if options.WithNamespace {
		if _, err := fmt.Fprintf(w, "%s\t", namespace); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%s",
		name,
		reference,
		metrics,
		minPods,
		maxPods,
		currentReplicas,
		translateTimestamp(hpa.CreationTimestamp),
	); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendLabels(hpa.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, hpa.Labels))
	return err
}

func printHorizontalPodAutoscalerList(list *autoscaling.HorizontalPodAutoscalerList, w io.Writer, options printers.PrintOptions) error {
	for i := range list.Items {
		if err := printHorizontalPodAutoscaler(&list.Items[i], w, options); err != nil {
			return err
		}
	}
	return nil
}

func printConfigMap(configMap *api.ConfigMap, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, configMap.Name, options.WithKind)

	namespace := configMap.Namespace

	if options.WithNamespace {
		if _, err := fmt.Fprintf(w, "%s\t", namespace); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(w, "%s\t%v\t%s", name, len(configMap.Data), translateTimestamp(configMap.CreationTimestamp)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendLabels(configMap.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, configMap.Labels))
	return err
}

func printConfigMapList(list *api.ConfigMapList, w io.Writer, options printers.PrintOptions) error {
	for i := range list.Items {
		if err := printConfigMap(&list.Items[i], w, options); err != nil {
			return err
		}
	}
	return nil
}

func printPodSecurityPolicy(item *extensions.PodSecurityPolicy, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, item.Name, options.WithKind)

	_, err := fmt.Fprintf(w, "%s\t%t\t%v\t%s\t%s\t%s\t%s\t%t\t%v\n", name, item.Spec.Privileged,
		item.Spec.AllowedCapabilities, item.Spec.SELinux.Rule,
		item.Spec.RunAsUser.Rule, item.Spec.FSGroup.Rule, item.Spec.SupplementalGroups.Rule, item.Spec.ReadOnlyRootFilesystem, item.Spec.Volumes)
	return err
}

func printPodSecurityPolicyList(list *extensions.PodSecurityPolicyList, w io.Writer, options printers.PrintOptions) error {
	for _, item := range list.Items {
		if err := printPodSecurityPolicy(&item, w, options); err != nil {
			return err
		}
	}

	return nil
}

func printExtensionsNetworkPolicy(networkPolicy *extensions.NetworkPolicy, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, networkPolicy.Name, options.WithKind)

	namespace := networkPolicy.Namespace

	if options.WithNamespace {
		if _, err := fmt.Fprintf(w, "%s\t", namespace); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(w, "%s\t%v\t%s", name, metav1.FormatLabelSelector(&networkPolicy.Spec.PodSelector), translateTimestamp(networkPolicy.CreationTimestamp)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendLabels(networkPolicy.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, networkPolicy.Labels))
	return err
}

func printExtensionsNetworkPolicyList(list *extensions.NetworkPolicyList, w io.Writer, options printers.PrintOptions) error {
	for i := range list.Items {
		if err := printExtensionsNetworkPolicy(&list.Items[i], w, options); err != nil {
			return err
		}
	}
	return nil
}

func printNetworkPolicy(networkPolicy *networking.NetworkPolicy, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, networkPolicy.Name, options.WithKind)

	namespace := networkPolicy.Namespace

	if options.WithNamespace {
		if _, err := fmt.Fprintf(w, "%s\t", namespace); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(w, "%s\t%v\t%s", name, metav1.FormatLabelSelector(&networkPolicy.Spec.PodSelector), translateTimestamp(networkPolicy.CreationTimestamp)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendLabels(networkPolicy.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, networkPolicy.Labels))
	return err
}

func printNetworkPolicyList(list *networking.NetworkPolicyList, w io.Writer, options printers.PrintOptions) error {
	for i := range list.Items {
		if err := printNetworkPolicy(&list.Items[i], w, options); err != nil {
			return err
		}
	}
	return nil
}

func printStorageClass(sc *storage.StorageClass, w io.Writer, options printers.PrintOptions) error {
	name := sc.Name

	if options.WithNamespace {
		return fmt.Errorf("storageclass is not namespaced")
	}

	if storageutil.IsDefaultAnnotation(sc.ObjectMeta) {
		name += " (default)"
	}
	provtype := sc.Provisioner

	if _, err := fmt.Fprintf(w, "%s\t%s\t", name, provtype); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendLabels(sc.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, sc.Labels)); err != nil {
		return err
	}

	return nil
}

func printStorageClassList(scList *storage.StorageClassList, w io.Writer, options printers.PrintOptions) error {
	for _, sc := range scList.Items {
		if err := printStorageClass(&sc, w, options); err != nil {
			return err
		}
	}
	return nil
}

func printStatus(status *metav1.Status, w io.Writer, options printers.PrintOptions) error {
	if _, err := fmt.Fprintf(w, "%s\t%s\t%s\n", status.Status, status.Reason, status.Message); err != nil {
		return err
	}

	return nil
}

// Lay out all the containers on one line if use wide output.
// DEPRECATED: convert to TableHandler and use layoutContainerCells
func layoutContainers(containers []api.Container, w io.Writer) error {
	var namesBuffer bytes.Buffer
	var imagesBuffer bytes.Buffer

	for i, container := range containers {
		namesBuffer.WriteString(container.Name)
		imagesBuffer.WriteString(container.Image)
		if i != len(containers)-1 {
			namesBuffer.WriteString(",")
			imagesBuffer.WriteString(",")
		}
	}
	_, err := fmt.Fprintf(w, "\t%s\t%s", namesBuffer.String(), imagesBuffer.String())
	if err != nil {
		return err
	}
	return nil
}

// Lay out all the containers on one line if use wide output.
func layoutContainerCells(containers []api.Container) (names string, images string) {
	var namesBuffer bytes.Buffer
	var imagesBuffer bytes.Buffer

	for i, container := range containers {
		namesBuffer.WriteString(container.Name)
		imagesBuffer.WriteString(container.Image)
		if i != len(containers)-1 {
			namesBuffer.WriteString(",")
			imagesBuffer.WriteString(",")
		}
	}
	return namesBuffer.String(), imagesBuffer.String()
}

// formatEventSource formats EventSource as a comma separated string excluding Host when empty
func formatEventSource(es api.EventSource) string {
	EventSourceString := []string{es.Component}
	if len(es.Host) > 0 {
		EventSourceString = append(EventSourceString, es.Host)
	}
	return strings.Join(EventSourceString, ", ")
}

func printControllerRevision(history *apps.ControllerRevision, w io.Writer, options printers.PrintOptions) error {
	name := printers.FormatResourceName(options.Kind, history.Name, options.WithKind)

	if options.WithNamespace {
		if _, err := fmt.Fprintf(w, "%s\t", history.Namespace); err != nil {
			return err
		}
	}

	controllerRef := controller.GetControllerOf(history)
	controllerName := "<none>"
	if controllerRef != nil {
		withKind := true
		controllerName = printers.FormatResourceName(controllerRef.Kind, controllerRef.Name, withKind)
	}
	revision := history.Revision
	age := translateTimestamp(history.CreationTimestamp)
	if _, err := fmt.Fprintf(w, "%s\t%s\t%d\t%s", name, controllerName, revision, age); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, printers.AppendLabels(history.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, printers.AppendAllLabels(options.ShowLabels, history.Labels))
	return err
}

func printControllerRevisionList(list *apps.ControllerRevisionList, w io.Writer, options printers.PrintOptions) error {
	for _, item := range list.Items {
		if err := printControllerRevision(&item, w, options); err != nil {
			return err
		}
	}
	return nil
}
