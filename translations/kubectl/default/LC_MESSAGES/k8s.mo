��    �      L  �   |      H  �   I  �     �   �  y  [  s   �  �  I  ]  �  �  3  8   �  %     �   @     �  =     �   L  i   8  [   �  G  �  3   F  2   z  8   �     �       9         Z     n  @   �  ,   �  *   �  7   #  '   [  &   �  .   �  =   �  *     0   B  ,   s     �  ]   �  0     0   O  "   �  ?   �     �  4     3   6  ,   j     �  *   �  A   �          7     T  )   o  6   �     �  3   �      "  �   C  (   1      Z   �   q   �   S!  7  "  /  M#  Q  }$  7   �%     &      &  I   ;&     �&     �&  �   �&  �   '  B   g(  �   �(  W   \)  W   �)  m   *  ;   z*  �   �*  /   �+  1   �+  '   �+  '   $,  1   L,  M   ~,  %   �,  (   �,  G   -  K   c-      �-     �-  "   �-  "   .     5.  -   U.  -   �.  9   �.  =   �.  3   )/  c   ]/  G   �/    	0  )  1  �   92  M   �2  �   3  �   �3  �   �4  �  q5  �   =7  �   �7  �   y8  �   =9  �  +:  �  �;  =  c=  9   �?  �   �?  /   �@  /   �@  9   �@     A  =   >A  $   |A     �A  &   �A  W   �A  )   @B  �   jB  1   PC  /   �C  �   �C  �  xD  [   %F  J   �F  a   �F  �   .G  9   �G  �   %H  �   �H  �   �I  8   qJ  %   �J  W   �J     (K  =   FK  u   �K  4   �K  -   /L  �   ]L  3   M  2   5M  8   hM     �M     �M  9   �M     N     )N  @   EN  ,   �N  *   �N  7   �N  '   O  &   >O  .   eO  =   �O  *   �O  0   �O  ,   .P     [P  ]   {P  0   �P  0   
Q  "   ;Q  ?   ^Q     �Q  4   �Q  3   �Q  ,   %R     RR  *   gR  A   �R     �R     �R     S  )   *S  6   TS     �S     �S      �S  v   �S  (   [T     �T  p   �T  `   U  �   mU  �   	V  �   �V     JW     fW     W  $   �W     �W     �W  a   �W  s   WX  B   �X  X   Y  +   gY  +   �Y  6   �Y  ;   �Y  q   2Z  /   �Z  1   �Z  '   [  '   .[     V[  &   o[  %   �[  (   �[  #   �[  K   	\      U\     v\  "   �\  "   �\     �\  -   �\  -   )]  9   W]     �]     �]  c   �]  #   .^  �   R^  �   �^  H   j_  &   �_  e   �_  z   @`  J   �`  �   a  W   �a  E   Db  a   �b  v   �b  �   cc  �   /d    �d     f  T   ;f     �f  /   �f  9   �f     g  =   1g  �   og     �g  &   h  +   =h     ih  r   ~h     �h  /   
i  �   :i         �   �                 X       z       N   |   f   m      Z             P   R   0      H   }      h      d   v       t      w   3       M   J   
   =   j   l   ^       2   /      �   +      .   &   	       i   Q      V       �   G   !       n   �   s       e   K   \   r              :   ]          O   @      a   9   '   *      k           L   �   {   -   b                                 1           x   W   Y          <      "   `   #   [      6                   A       ;   8   T   U   >   S       c           (      D       o   7      _   g       4          u       $       p               y   E           )   C   ,      q      I       F   %      ~   B      ?      �   5    A comma-delimited set of quota scopes that must all match each object tracked by the quota.A comma-delimited set of quota scopes that must all match each object tracked by the quota. A comma-delimited set of resource=quantity pairs that define a hard limit.A comma-delimited set of resource=quantity pairs that define a hard limit. A label selector to use for this budget. Only equality-based selector requirements are supported.A label selector to use for this budget. Only equality-based selector requirements are supported. A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set.A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set. A schedule in the Cron format the job should be run with.A schedule in the Cron format the job should be run with. Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP.Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP. An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field. An inline JSON override for the generated service object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.  Only used if --expose is true.An inline JSON override for the generated service object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.  Only used if --expose is true. Apply a configuration to a resource by filename or stdin Approve a certificate signing request Assign your own ClusterIP or set to 'None' for a 'headless' service (no loadbalancing).Assign your own ClusterIP or set to 'None' for a 'headless' service (no loadbalancing). Attach to a running container Auto-scale a Deployment, ReplicaSet, or ReplicationController ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service.ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service. ClusterRole this ClusterRoleBinding should referenceClusterRole this ClusterRoleBinding should reference ClusterRole this RoleBinding should referenceClusterRole this RoleBinding should reference Container name which will have its image upgraded. Only relevant when --image is specified, ignored otherwise. Required when using --image on a multi-container podContainer name which will have its image upgraded. Only relevant when --image is specified, ignored otherwise. Required when using --image on a multi-container pod Convert config files between different API versions Copy files and directories to and from containers. Create a ClusterRoleBinding for a particular ClusterRole Create a LoadBalancer service. Create a NodePort service. Create a RoleBinding for a particular Role or ClusterRole Create a TLS secret Create a clusterIP service. Create a configmap from a local file, directory or literal value Create a deployment with the specified name. Create a namespace with the specified name Create a pod disruption budget with the specified name. Create a quota with the specified name. Create a resource by filename or stdin Create a secret for use with a Docker registry Create a secret from a local file, directory or literal value Create a secret using specified subcommand Create a service account with the specified name Create a service using specified subcommand. Create an ExternalName service. Delete resources by filenames, stdin, resources and names, or by resources and label selector Delete the specified cluster from the kubeconfig Delete the specified context from the kubeconfig Deny a certificate signing request Deprecated: Gracefully shut down a resource by name or filename Describe one or many contexts Display Resource (CPU/Memory/Storage) usage of nodes Display Resource (CPU/Memory/Storage) usage of pods Display Resource (CPU/Memory/Storage) usage. Display cluster info Display clusters defined in the kubeconfig Display merged kubeconfig settings or a specified kubeconfig file Display one or many resources Displays the current-context Documentation of resources Drain node in preparation for maintenance Dump lots of relevant info for debugging and diagnosis Edit a resource on the server Email for Docker registryEmail for Docker registry Execute a command in a container Explicit policy for when to pull container images. Required when --image is same as existing image, ignored otherwise.Explicit policy for when to pull container images. Required when --image is same as existing image, ignored otherwise. Forward one or more local ports to a pod Help about any command IP to assign to the Load Balancer. If empty, an ephemeral IP will be created and used (cloud-provider specific).IP to assign to the Load Balancer. If empty, an ephemeral IP will be created and used (cloud-provider specific). If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP'If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP' If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource. If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource. Image to use for upgrading the replication controller. Must be distinct from the existing image (either new image or new image tag).  Can not be used with --filename/-fImage to use for upgrading the replication controller. Must be distinct from the existing image (either new image or new image tag).  Can not be used with --filename/-f Manage a deployment rolloutManage a deployment rollout Mark node as schedulable Mark node as unschedulable Mark the provided resource as pausedMark the provided resource as paused Modify certificate resources. Modify kubeconfig files Name or number for the port on the container that the service should direct traffic to. Optional.Name or number for the port on the container that the service should direct traffic to. Optional. Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used. Output shell completion code for the specified shell (bash or zsh) Output the formatted object with the given group version (for ex: 'extensions/v1beta1').Output the formatted object with the given group version (for ex: 'extensions/v1beta1'). Password for Docker registry authenticationPassword for Docker registry authentication Path to PEM encoded public key certificate.Path to PEM encoded public key certificate. Path to private key associated with given certificate.Path to private key associated with given certificate. Perform a rolling update of the given ReplicationController Precondition for resource version. Requires that the current resource version match this value in order to scale.Precondition for resource version. Requires that the current resource version match this value in order to scale. Print the client and server version information Print the list of flags inherited by all commands Print the logs for a container in a pod Replace a resource by filename or stdin Resume a paused resourceResume a paused resource Role this RoleBinding should referenceRole this RoleBinding should reference Run a particular image on the cluster Run a proxy to the Kubernetes API server Server location for Docker registryServer location for Docker registry Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job Set specific features on objects Set the selector on a resource Sets a cluster entry in kubeconfig Sets a context entry in kubeconfig Sets a user entry in kubeconfig Sets an individual value in a kubeconfig file Sets the current-context in a kubeconfig file Show details of a specific resource or group of resources Show the status of the rolloutShow the status of the rollout Synonym for --target-portSynonym for --target-port Take a replication controller, service, deployment or pod and expose it as a new Kubernetes Service The image for the container to run.The image for the container to run. The image pull policy for the container. If left empty, this value will not be specified by the client and defaulted by the serverThe image pull policy for the container. If left empty, this value will not be specified by the client and defaulted by the server The key to use to differentiate between two different controllers, default 'deployment'.  Only relevant when --image is specified, ignored otherwiseThe key to use to differentiate between two different controllers, default 'deployment'.  Only relevant when --image is specified, ignored otherwise The minimum number or percentage of available pods this budget requires.The minimum number or percentage of available pods this budget requires. The name for the newly created object.The name for the newly created object. The name for the newly created object. If not specified, the name of the input resource will be used.The name for the newly created object. If not specified, the name of the input resource will be used. The name of the API generator to use, see http://kubernetes.io/docs/user-guide/kubectl-conventions/#generators for a list.The name of the API generator to use, see http://kubernetes.io/docs/user-guide/kubectl-conventions/#generators for a list. The name of the API generator to use. Currently there is only 1 generator.The name of the API generator to use. Currently there is only 1 generator. The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'.The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'. The name of the generator to use for creating a service.  Only used if --expose is trueThe name of the generator to use for creating a service.  Only used if --expose is true The network protocol for the service to be created. Default is 'TCP'.The network protocol for the service to be created. Default is 'TCP'. The port that the service should serve on. Copied from the resource being exposed, if unspecifiedThe port that the service should serve on. Copied from the resource being exposed, if unspecified The port that this container exposes.  If --expose is true, this is also the port used by the service that is created.The port that this container exposes.  If --expose is true, this is also the port used by the service that is created. The resource requirement limits for this container.  For example, 'cpu=200m,memory=512Mi'.  Note that server side components may assign limits depending on the server configuration, such as limit ranges.The resource requirement limits for this container.  For example, 'cpu=200m,memory=512Mi'.  Note that server side components may assign limits depending on the server configuration, such as limit ranges. The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges.The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges. The restart policy for this Pod.  Legal values [Always, OnFailure, Never].  If set to 'Always' a deployment is created, if set to 'OnFailure' a job is created, if set to 'Never', a regular pod is created. For the latter two --replicas must be 1.  Default 'Always', for CronJobs `Never`.The restart policy for this Pod.  Legal values [Always, OnFailure, Never].  If set to 'Always' a deployment is created, if set to 'OnFailure' a job is created, if set to 'Never', a regular pod is created. For the latter two --replicas must be 1.  Default 'Always', for CronJobs `Never`. The type of secret to createThe type of secret to create Type for this service: ClusterIP, NodePort, or LoadBalancer. Default is 'ClusterIP'.Type for this service: ClusterIP, NodePort, or LoadBalancer. Default is 'ClusterIP'. Undo a previous rolloutUndo a previous rollout Unsets an individual value in a kubeconfig file Update field(s) of a resource using strategic merge patch Update image of a pod template Update resource requests/limits on objects with pod templates Update the annotations on a resource Update the labels on a resource Update the taints on one or more nodes Username for Docker registry authenticationUsername for Docker registry authentication View rollout historyView rollout history Where to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in that directoryWhere to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in that directory external name of serviceexternal name of service kubectl controls the Kubernetes cluster manager watch is only supported on individual resources and resource collections - %d resources were found watch is only supported on individual resources and resource collections - %d resources were found Project-Id-Version: gettext-go-examples-hello
Report-Msgid-Bugs-To: 
POT-Creation-Date: 2013-12-12 20:03+0000
PO-Revision-Date: 2017-02-21 22:06-0800
Last-Translator: Brendan Burns <brendan.d.burns@gmail.com>
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit
X-Generator: Poedit 1.6.10
X-Poedit-SourceCharset: UTF-8
Language-Team: 
Plural-Forms: nplurals=2; plural=(n != 1);
Language: en
 A comma-delimited set of quota scopes that must all match each object tracked by the quota. A comma-delimited set of resource=quantity pairs that define a hard limit. A label selector to use for this budget. Only equality-based selector requirements are supported. A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set. A schedule in the Cron format the job should be run with. Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP. An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field. An inline JSON override for the generated service object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.  Only used if --expose is true. Apply a configuration to a resource by filename or stdin Approve a certificate signing request Assign your own ClusterIP or set to 'None' for a 'headless' service (no loadbalancing). Attach to a running container Auto-scale a Deployment, ReplicaSet, or ReplicationController ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service. ClusterRole this ClusterRoleBinding should reference ClusterRole this RoleBinding should reference Container name which will have its image upgraded. Only relevant when --image is specified, ignored otherwise. Required when using --image on a multi-container pod Convert config files between different API versions Copy files and directories to and from containers. Create a ClusterRoleBinding for a particular ClusterRole Create a LoadBalancer service. Create a NodePort service. Create a RoleBinding for a particular Role or ClusterRole Create a TLS secret Create a clusterIP service. Create a configmap from a local file, directory or literal value Create a deployment with the specified name. Create a namespace with the specified name Create a pod disruption budget with the specified name. Create a quota with the specified name. Create a resource by filename or stdin Create a secret for use with a Docker registry Create a secret from a local file, directory or literal value Create a secret using specified subcommand Create a service account with the specified name Create a service using specified subcommand. Create an ExternalName service. Delete resources by filenames, stdin, resources and names, or by resources and label selector Delete the specified cluster from the kubeconfig Delete the specified context from the kubeconfig Deny a certificate signing request Deprecated: Gracefully shut down a resource by name or filename Describe one or many contexts Display Resource (CPU/Memory/Storage) usage of nodes Display Resource (CPU/Memory/Storage) usage of pods Display Resource (CPU/Memory/Storage) usage. Display cluster info Display clusters defined in the kubeconfig Display merged kubeconfig settings or a specified kubeconfig file Display one or many resources Displays the current-context Documentation of resources Drain node in preparation for maintenance Dump lots of relevant info for debugging and diagnosis Edit a resource on the server Email for Docker registry Execute a command in a container Explicit policy for when to pull container images. Required when --image is same as existing image, ignored otherwise. Forward one or more local ports to a pod Help about any command IP to assign to the Load Balancer. If empty, an ephemeral IP will be created and used (cloud-provider specific). If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP' If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource. If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource. Image to use for upgrading the replication controller. Must be distinct from the existing image (either new image or new image tag).  Can not be used with --filename/-f Manage a deployment rollout Mark node as schedulable Mark node as unschedulable Mark the provided resource as paused Modify certificate resources. Modify kubeconfig files Name or number for the port on the container that the service should direct traffic to. Optional. Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used. Output shell completion code for the specified shell (bash or zsh) Output the formatted object with the given group version (for ex: 'extensions/v1beta1'). Password for Docker registry authentication Path to PEM encoded public key certificate. Path to private key associated with given certificate. Perform a rolling update of the given ReplicationController Precondition for resource version. Requires that the current resource version match this value in order to scale. Print the client and server version information Print the list of flags inherited by all commands Print the logs for a container in a pod Replace a resource by filename or stdin Resume a paused resource Role this RoleBinding should reference Run a particular image on the cluster Run a proxy to the Kubernetes API server Server location for Docker registry Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job Set specific features on objects Set the selector on a resource Sets a cluster entry in kubeconfig Sets a context entry in kubeconfig Sets a user entry in kubeconfig Sets an individual value in a kubeconfig file Sets the current-context in a kubeconfig file Show details of a specific resource or group of resources Show the status of the rollout Synonym for --target-port Take a replication controller, service, deployment or pod and expose it as a new Kubernetes Service The image for the container to run. The image pull policy for the container. If left empty, this value will not be specified by the client and defaulted by the server The key to use to differentiate between two different controllers, default 'deployment'.  Only relevant when --image is specified, ignored otherwise The minimum number or percentage of available pods this budget requires. The name for the newly created object. The name for the newly created object. If not specified, the name of the input resource will be used. The name of the API generator to use, see http://kubernetes.io/docs/user-guide/kubectl-conventions/#generators for a list. The name of the API generator to use. Currently there is only 1 generator. The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'. The name of the generator to use for creating a service.  Only used if --expose is true The network protocol for the service to be created. Default is 'TCP'. The port that the service should serve on. Copied from the resource being exposed, if unspecified The port that this container exposes.  If --expose is true, this is also the port used by the service that is created. The resource requirement limits for this container.  For example, 'cpu=200m,memory=512Mi'.  Note that server side components may assign limits depending on the server configuration, such as limit ranges. The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges. The restart policy for this Pod.  Legal values [Always, OnFailure, Never].  If set to 'Always' a deployment is created, if set to 'OnFailure' a job is created, if set to 'Never', a regular pod is created. For the latter two --replicas must be 1.  Default 'Always', for CronJobs `Never`. The type of secret to create Type for this service: ClusterIP, NodePort, or LoadBalancer. Default is 'ClusterIP'. Undo a previous rollout Unsets an individual value in a kubeconfig file Update field(s) of a resource using strategic merge patch Update image of a pod template Update resource requests/limits on objects with pod templates Update the annotations on a resourcewatch is only supported on individual resources and resource collections - %d resources were found Update the labels on a resource Update the taints on one or more nodes Username for Docker registry authentication View rollout history Where to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in that directory external name of service kubectl controls the Kubernetes cluster manager watch is only supported on individual resources and resource collections - %d resource was found watch is only supported on individual resources and resource collections - %d resources were found 