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

package secret

import (
	"time"

	federation_api "k8s.io/kubernetes/federation/apis/federation/v1beta1"
	federationclientset "k8s.io/kubernetes/federation/client/clientset_generated/federation_release_1_5"
	"k8s.io/kubernetes/federation/pkg/federation-controller/util"
	"k8s.io/kubernetes/federation/pkg/federation-controller/util/eventsink"
	"k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/cache"
	kubeclientset "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_5"
	"k8s.io/kubernetes/pkg/client/record"
	"k8s.io/kubernetes/pkg/controller"
	pkg_runtime "k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/types"
	"k8s.io/kubernetes/pkg/util/flowcontrol"
	"k8s.io/kubernetes/pkg/watch"

	"github.com/golang/glog"
)

const (
	allClustersKey = "ALL_CLUSTERS"
)

type SecretController struct {
	// For triggering single secret reconciliation. This is used when there is an
	// add/update/delete operation on a secret in either federated API server or
	// in some member of the federation.
	secretDeliverer *util.DelayingDeliverer

	// For triggering all secrets reconciliation. This is used when
	// a new cluster becomes available.
	clusterDeliverer *util.DelayingDeliverer

	// Contains secrets present in members of federation.
	secretFederatedInformer util.FederatedInformer
	// For updating members of federation.
	federatedUpdater util.FederatedUpdater
	// Definitions of secrets that should be federated.
	secretInformerStore cache.Store
	// Informer controller for secrets that should be federated.
	secretInformerController cache.ControllerInterface

	// Client to federated api server.
	federatedApiClient federationclientset.Interface

	// Backoff manager for secrets
	secretBackoff *flowcontrol.Backoff

	// For events
	eventRecorder record.EventRecorder

	secretReviewDelay     time.Duration
	clusterAvailableDelay time.Duration
	smallDelay            time.Duration
	updateTimeout         time.Duration
}

// NewSecretController returns a new secret controller
func NewSecretController(client federationclientset.Interface) *SecretController {
	broadcaster := record.NewBroadcaster()
	broadcaster.StartRecordingToSink(eventsink.NewFederatedEventSink(client))
	recorder := broadcaster.NewRecorder(api.EventSource{Component: "federated-secrets-controller"})

	secretcontroller := &SecretController{
		federatedApiClient:    client,
		secretReviewDelay:     time.Second * 10,
		clusterAvailableDelay: time.Second * 20,
		smallDelay:            time.Second * 3,
		updateTimeout:         time.Second * 30,
		secretBackoff:         flowcontrol.NewBackOff(5*time.Second, time.Minute),
		eventRecorder:         recorder,
	}

	// Build delivereres for triggering reconciliations.
	secretcontroller.secretDeliverer = util.NewDelayingDeliverer()
	secretcontroller.clusterDeliverer = util.NewDelayingDeliverer()

	// Start informer in federated API servers on secrets that should be federated.
	secretcontroller.secretInformerStore, secretcontroller.secretInformerController = cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (pkg_runtime.Object, error) {
				versionedOptions := util.VersionizeV1ListOptions(options)
				return client.Core().Secrets(api_v1.NamespaceAll).List(versionedOptions)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				versionedOptions := util.VersionizeV1ListOptions(options)
				return client.Core().Secrets(api_v1.NamespaceAll).Watch(versionedOptions)
			},
		},
		&api_v1.Secret{},
		controller.NoResyncPeriodFunc(),
		util.NewTriggerOnAllChanges(func(obj pkg_runtime.Object) { secretcontroller.deliverSecretObj(obj, 0, false) }))

	// Federated informer on secrets in members of federation.
	secretcontroller.secretFederatedInformer = util.NewFederatedInformer(
		client,
		func(cluster *federation_api.Cluster, targetClient kubeclientset.Interface) (cache.Store, cache.ControllerInterface) {
			return cache.NewInformer(
				&cache.ListWatch{
					ListFunc: func(options api.ListOptions) (pkg_runtime.Object, error) {
						versionedOptions := util.VersionizeV1ListOptions(options)
						return targetClient.Core().Secrets(api_v1.NamespaceAll).List(versionedOptions)
					},
					WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
						versionedOptions := util.VersionizeV1ListOptions(options)
						return targetClient.Core().Secrets(api_v1.NamespaceAll).Watch(versionedOptions)
					},
				},
				&api_v1.Secret{},
				controller.NoResyncPeriodFunc(),
				// Trigger reconciliation whenever something in federated cluster is changed. In most cases it
				// would be just confirmation that some secret opration succeeded.
				util.NewTriggerOnAllChanges(
					func(obj pkg_runtime.Object) {
						secretcontroller.deliverSecretObj(obj, secretcontroller.secretReviewDelay, false)
					},
				))
		},

		&util.ClusterLifecycleHandlerFuncs{
			ClusterAvailable: func(cluster *federation_api.Cluster) {
				// When new cluster becomes available process all the secrets again.
				secretcontroller.clusterDeliverer.DeliverAt(allClustersKey, nil, time.Now().Add(secretcontroller.clusterAvailableDelay))
			},
		},
	)

	// Federated updeater along with Create/Update/Delete operations.
	secretcontroller.federatedUpdater = util.NewFederatedUpdater(secretcontroller.secretFederatedInformer,
		func(client kubeclientset.Interface, obj pkg_runtime.Object) error {
			secret := obj.(*api_v1.Secret)
			_, err := client.Core().Secrets(secret.Namespace).Create(secret)
			return err
		},
		func(client kubeclientset.Interface, obj pkg_runtime.Object) error {
			secret := obj.(*api_v1.Secret)
			_, err := client.Core().Secrets(secret.Namespace).Update(secret)
			return err
		},
		func(client kubeclientset.Interface, obj pkg_runtime.Object) error {
			secret := obj.(*api_v1.Secret)
			err := client.Core().Secrets(secret.Namespace).Delete(secret.Name, &api_v1.DeleteOptions{})
			return err
		})
	return secretcontroller
}

func (secretcontroller *SecretController) Run(stopChan <-chan struct{}) {
	go secretcontroller.secretInformerController.Run(stopChan)
	secretcontroller.secretFederatedInformer.Start()
	go func() {
		<-stopChan
		secretcontroller.secretFederatedInformer.Stop()
	}()
	secretcontroller.secretDeliverer.StartWithHandler(func(item *util.DelayingDelivererItem) {
		secret := item.Value.(*types.NamespacedName)
		secretcontroller.reconcileSecret(*secret)
	})
	secretcontroller.clusterDeliverer.StartWithHandler(func(_ *util.DelayingDelivererItem) {
		secretcontroller.reconcileSecretsOnClusterChange()
	})
	util.StartBackoffGC(secretcontroller.secretBackoff, stopChan)
}

func (secretcontroller *SecretController) deliverSecretObj(obj interface{}, delay time.Duration, failed bool) {
	secret := obj.(*api_v1.Secret)
	secretcontroller.deliverSecret(types.NamespacedName{Namespace: secret.Namespace, Name: secret.Name}, delay, failed)
}

// Adds backoff to delay if this delivery is related to some failure. Resets backoff if there was no failure.
func (secretcontroller *SecretController) deliverSecret(secret types.NamespacedName, delay time.Duration, failed bool) {
	key := secret.String()
	if failed {
		secretcontroller.secretBackoff.Next(key, time.Now())
		delay = delay + secretcontroller.secretBackoff.Get(key)
	} else {
		secretcontroller.secretBackoff.Reset(key)
	}
	secretcontroller.secretDeliverer.DeliverAfter(key, &secret, delay)
}

// Check whether all data stores are in sync. False is returned if any of the informer/stores is not yet
// synced with the corresponding api server.
func (secretcontroller *SecretController) isSynced() bool {
	if !secretcontroller.secretFederatedInformer.ClustersSynced() {
		glog.V(2).Infof("Cluster list not synced")
		return false
	}
	clusters, err := secretcontroller.secretFederatedInformer.GetReadyClusters()
	if err != nil {
		glog.Errorf("Failed to get ready clusters: %v", err)
		return false
	}
	if !secretcontroller.secretFederatedInformer.GetTargetStore().ClustersSynced(clusters) {
		return false
	}
	return true
}

// The function triggers reconciliation of all federated secrets.
func (secretcontroller *SecretController) reconcileSecretsOnClusterChange() {
	if !secretcontroller.isSynced() {
		secretcontroller.clusterDeliverer.DeliverAt(allClustersKey, nil, time.Now().Add(secretcontroller.clusterAvailableDelay))
	}
	for _, obj := range secretcontroller.secretInformerStore.List() {
		secret := obj.(*api_v1.Secret)
		secretcontroller.deliverSecret(types.NamespacedName{Namespace: secret.Namespace, Name: secret.Name}, secretcontroller.smallDelay, false)
	}
}

func (secretcontroller *SecretController) reconcileSecret(secret types.NamespacedName) {

	if !secretcontroller.isSynced() {
		secretcontroller.deliverSecret(secret, secretcontroller.clusterAvailableDelay, false)
		return
	}

	key := secret.String()
	baseSecretObj, exist, err := secretcontroller.secretInformerStore.GetByKey(key)
	if err != nil {
		glog.Errorf("Failed to query main secret store for %v: %v", key, err)
		secretcontroller.deliverSecret(secret, 0, true)
		return
	}

	if !exist {
		// Not federated secret, ignoring.
		return
	}
	baseSecret := baseSecretObj.(*api_v1.Secret)

	clusters, err := secretcontroller.secretFederatedInformer.GetReadyClusters()
	if err != nil {
		glog.Errorf("Failed to get cluster list: %v", err)
		secretcontroller.deliverSecret(secret, secretcontroller.clusterAvailableDelay, false)
		return
	}

	operations := make([]util.FederatedOperation, 0)
	for _, cluster := range clusters {
		clusterSecretObj, found, err := secretcontroller.secretFederatedInformer.GetTargetStore().GetByKey(cluster.Name, key)
		if err != nil {
			glog.Errorf("Failed to get %s from %s: %v", key, cluster.Name, err)
			secretcontroller.deliverSecret(secret, 0, true)
			return
		}

		// The data should not be modified.
		desiredSecret := &api_v1.Secret{
			ObjectMeta: util.DeepCopyRelevantObjectMeta(baseSecret.ObjectMeta),
			Data:       baseSecret.Data,
			Type:       baseSecret.Type,
		}

		if !found {
			secretcontroller.eventRecorder.Eventf(baseSecret, api.EventTypeNormal, "CreateInCluster",
				"Creating secret in cluster %s", cluster.Name)

			operations = append(operations, util.FederatedOperation{
				Type:        util.OperationTypeAdd,
				Obj:         desiredSecret,
				ClusterName: cluster.Name,
			})
		} else {
			clusterSecret := clusterSecretObj.(*api_v1.Secret)

			// Update existing secret, if needed.
			if !util.SecretEquivalent(*desiredSecret, *clusterSecret) {

				secretcontroller.eventRecorder.Eventf(baseSecret, api.EventTypeNormal, "UpdateInCluster",
					"Updating secret in cluster %s", cluster.Name)
				operations = append(operations, util.FederatedOperation{
					Type:        util.OperationTypeUpdate,
					Obj:         desiredSecret,
					ClusterName: cluster.Name,
				})
			}
		}
	}

	if len(operations) == 0 {
		// Everything is in order
		return
	}
	err = secretcontroller.federatedUpdater.UpdateWithOnError(operations, secretcontroller.updateTimeout,
		func(op util.FederatedOperation, operror error) {
			secretcontroller.eventRecorder.Eventf(baseSecret, api.EventTypeNormal, "UpdateInClusterFailed",
				"Secret update in cluster %s failed: %v", op.ClusterName, operror)
		})

	if err != nil {
		glog.Errorf("Failed to execute updates for %s: %v", key, err)
		secretcontroller.deliverSecret(secret, 0, true)
		return
	}

	// Evertyhing is in order but lets be double sure
	secretcontroller.deliverSecret(secret, secretcontroller.secretReviewDelay, false)
}
