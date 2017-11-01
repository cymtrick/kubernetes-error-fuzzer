package containerregistry

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// PasswordName enumerates the values for password name.
type PasswordName string

const (
	// Password specifies the password state for password name.
	Password PasswordName = "password"
	// Password2 specifies the password 2 state for password name.
	Password2 PasswordName = "password2"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled specifies the canceled state for provisioning state.
	Canceled ProvisioningState = "Canceled"
	// Creating specifies the creating state for provisioning state.
	Creating ProvisioningState = "Creating"
	// Deleting specifies the deleting state for provisioning state.
	Deleting ProvisioningState = "Deleting"
	// Failed specifies the failed state for provisioning state.
	Failed ProvisioningState = "Failed"
	// Succeeded specifies the succeeded state for provisioning state.
	Succeeded ProvisioningState = "Succeeded"
	// Updating specifies the updating state for provisioning state.
	Updating ProvisioningState = "Updating"
)

// RegistryUsageUnit enumerates the values for registry usage unit.
type RegistryUsageUnit string

const (
	// Bytes specifies the bytes state for registry usage unit.
	Bytes RegistryUsageUnit = "Bytes"
	// Count specifies the count state for registry usage unit.
	Count RegistryUsageUnit = "Count"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic specifies the basic state for sku name.
	Basic SkuName = "Basic"
	// Classic specifies the classic state for sku name.
	Classic SkuName = "Classic"
	// Premium specifies the premium state for sku name.
	Premium SkuName = "Premium"
	// Standard specifies the standard state for sku name.
	Standard SkuName = "Standard"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic specifies the sku tier basic state for sku tier.
	SkuTierBasic SkuTier = "Basic"
	// SkuTierClassic specifies the sku tier classic state for sku tier.
	SkuTierClassic SkuTier = "Classic"
	// SkuTierPremium specifies the sku tier premium state for sku tier.
	SkuTierPremium SkuTier = "Premium"
	// SkuTierStandard specifies the sku tier standard state for sku tier.
	SkuTierStandard SkuTier = "Standard"
)

// WebhookAction enumerates the values for webhook action.
type WebhookAction string

const (
	// Delete specifies the delete state for webhook action.
	Delete WebhookAction = "delete"
	// Push specifies the push state for webhook action.
	Push WebhookAction = "push"
)

// WebhookStatus enumerates the values for webhook status.
type WebhookStatus string

const (
	// Disabled specifies the disabled state for webhook status.
	Disabled WebhookStatus = "disabled"
	// Enabled specifies the enabled state for webhook status.
	Enabled WebhookStatus = "enabled"
)

// Actor is the agent that initiated the event. For most situations, this could be from the authorization context of
// the request.
type Actor struct {
	Name *string `json:"name,omitempty"`
}

// CallbackConfig is the configuration of service URI and custom headers for the webhook.
type CallbackConfig struct {
	autorest.Response `json:"-"`
	ServiceURI        *string             `json:"serviceUri,omitempty"`
	CustomHeaders     *map[string]*string `json:"customHeaders,omitempty"`
}

// Event is the event for a webhook.
type Event struct {
	ID                   *string               `json:"id,omitempty"`
	EventRequestMessage  *EventRequestMessage  `json:"eventRequestMessage,omitempty"`
	EventResponseMessage *EventResponseMessage `json:"eventResponseMessage,omitempty"`
}

// EventContent is the content of the event request message.
type EventContent struct {
	ID        *string    `json:"id,omitempty"`
	Timestamp *date.Time `json:"timestamp,omitempty"`
	Action    *string    `json:"action,omitempty"`
	Target    *Target    `json:"target,omitempty"`
	Request   *Request   `json:"request,omitempty"`
	Actor     *Actor     `json:"actor,omitempty"`
	Source    *Source    `json:"source,omitempty"`
}

// EventInfo is the basic information of an event.
type EventInfo struct {
	autorest.Response `json:"-"`
	ID                *string `json:"id,omitempty"`
}

// EventListResult is the result of a request to list events for a webhook.
type EventListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Event `json:"value,omitempty"`
	NextLink          *string  `json:"nextLink,omitempty"`
}

// EventListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client EventListResult) EventListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// EventRequestMessage is the event request message sent to the service URI.
type EventRequestMessage struct {
	Content    *EventContent       `json:"content,omitempty"`
	Headers    *map[string]*string `json:"headers,omitempty"`
	Method     *string             `json:"method,omitempty"`
	RequestURI *string             `json:"requestUri,omitempty"`
	Version    *string             `json:"version,omitempty"`
}

// EventResponseMessage is the event response message received from the service URI.
type EventResponseMessage struct {
	Content      *string             `json:"content,omitempty"`
	Headers      *map[string]*string `json:"headers,omitempty"`
	ReasonPhrase *string             `json:"reasonPhrase,omitempty"`
	StatusCode   *string             `json:"statusCode,omitempty"`
	Version      *string             `json:"version,omitempty"`
}

// OperationDefinition is the definition of a container registry operation.
type OperationDefinition struct {
	Name    *string                     `json:"name,omitempty"`
	Display *OperationDisplayDefinition `json:"display,omitempty"`
}

// OperationDisplayDefinition is the display information for a container registry operation.
type OperationDisplayDefinition struct {
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
}

// OperationListResult is the result of a request to list container registry operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]OperationDefinition `json:"value,omitempty"`
	NextLink          *string                `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// RegenerateCredentialParameters is the parameters used to regenerate the login credential.
type RegenerateCredentialParameters struct {
	Name PasswordName `json:"name,omitempty"`
}

// Registry is an object that represents a container registry.
type Registry struct {
	autorest.Response   `json:"-"`
	ID                  *string             `json:"id,omitempty"`
	Name                *string             `json:"name,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Location            *string             `json:"location,omitempty"`
	Tags                *map[string]*string `json:"tags,omitempty"`
	Sku                 *Sku                `json:"sku,omitempty"`
	*RegistryProperties `json:"properties,omitempty"`
}

// RegistryListCredentialsResult is the response from the ListCredentials operation.
type RegistryListCredentialsResult struct {
	autorest.Response `json:"-"`
	Username          *string             `json:"username,omitempty"`
	Passwords         *[]RegistryPassword `json:"passwords,omitempty"`
}

// RegistryListResult is the result of a request to list container registries.
type RegistryListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Registry `json:"value,omitempty"`
	NextLink          *string     `json:"nextLink,omitempty"`
}

// RegistryListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client RegistryListResult) RegistryListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// RegistryNameCheckRequest is a request to check whether a container registry name is available.
type RegistryNameCheckRequest struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// RegistryNameStatus is the result of a request to check the availability of a container registry name.
type RegistryNameStatus struct {
	autorest.Response `json:"-"`
	NameAvailable     *bool   `json:"nameAvailable,omitempty"`
	Reason            *string `json:"reason,omitempty"`
	Message           *string `json:"message,omitempty"`
}

// RegistryPassword is the login password for the container registry.
type RegistryPassword struct {
	Name  PasswordName `json:"name,omitempty"`
	Value *string      `json:"value,omitempty"`
}

// RegistryProperties is the properties of a container registry.
type RegistryProperties struct {
	LoginServer       *string                   `json:"loginServer,omitempty"`
	CreationDate      *date.Time                `json:"creationDate,omitempty"`
	ProvisioningState ProvisioningState         `json:"provisioningState,omitempty"`
	Status            *Status                   `json:"status,omitempty"`
	AdminUserEnabled  *bool                     `json:"adminUserEnabled,omitempty"`
	StorageAccount    *StorageAccountProperties `json:"storageAccount,omitempty"`
}

// RegistryPropertiesUpdateParameters is the parameters for updating the properties of a container registry.
type RegistryPropertiesUpdateParameters struct {
	AdminUserEnabled *bool                     `json:"adminUserEnabled,omitempty"`
	StorageAccount   *StorageAccountProperties `json:"storageAccount,omitempty"`
}

// RegistryUpdateParameters is the parameters for updating a container registry.
type RegistryUpdateParameters struct {
	Tags                                *map[string]*string `json:"tags,omitempty"`
	Sku                                 *Sku                `json:"sku,omitempty"`
	*RegistryPropertiesUpdateParameters `json:"properties,omitempty"`
}

// RegistryUsage is the quota usage for a container registry.
type RegistryUsage struct {
	Name         *string           `json:"name,omitempty"`
	Limit        *int64            `json:"limit,omitempty"`
	CurrentValue *int64            `json:"currentValue,omitempty"`
	Unit         RegistryUsageUnit `json:"unit,omitempty"`
}

// RegistryUsageListResult is the result of a request to get container registry quota usages.
type RegistryUsageListResult struct {
	autorest.Response `json:"-"`
	Value             *[]RegistryUsage `json:"value,omitempty"`
}

// Replication is an object that represents a replication for a container registry.
type Replication struct {
	autorest.Response      `json:"-"`
	ID                     *string             `json:"id,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Location               *string             `json:"location,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
	*ReplicationProperties `json:"properties,omitempty"`
}

// ReplicationListResult is the result of a request to list replications for a container registry.
type ReplicationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Replication `json:"value,omitempty"`
	NextLink          *string        `json:"nextLink,omitempty"`
}

// ReplicationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ReplicationListResult) ReplicationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ReplicationProperties is the properties of a replication.
type ReplicationProperties struct {
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	Status            *Status           `json:"status,omitempty"`
}

// ReplicationUpdateParameters is the parameters for updating a replication.
type ReplicationUpdateParameters struct {
	Tags *map[string]*string `json:"tags,omitempty"`
}

// Request is the request that generated the event.
type Request struct {
	ID        *string `json:"id,omitempty"`
	Addr      *string `json:"addr,omitempty"`
	Host      *string `json:"host,omitempty"`
	Method    *string `json:"method,omitempty"`
	Useragent *string `json:"useragent,omitempty"`
}

// Resource is an Azure resource.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// Sku is the SKU of a container registry.
type Sku struct {
	Name SkuName `json:"name,omitempty"`
	Tier SkuTier `json:"tier,omitempty"`
}

// Source is the registry node that generated the event. Put differently, while the actor initiates the event, the
// source generates it.
type Source struct {
	Addr       *string `json:"addr,omitempty"`
	InstanceID *string `json:"instanceID,omitempty"`
}

// Status is the status of an Azure resource at the time the operation was called.
type Status struct {
	DisplayStatus *string    `json:"displayStatus,omitempty"`
	Message       *string    `json:"message,omitempty"`
	Timestamp     *date.Time `json:"timestamp,omitempty"`
}

// StorageAccountProperties is the properties of a storage account for a container registry. Only applicable to Classic
// SKU.
type StorageAccountProperties struct {
	ID *string `json:"id,omitempty"`
}

// Target is the target of the event.
type Target struct {
	MediaType  *string `json:"mediaType,omitempty"`
	Size       *int64  `json:"size,omitempty"`
	Digest     *string `json:"digest,omitempty"`
	Length     *int64  `json:"length,omitempty"`
	Repository *string `json:"repository,omitempty"`
	URL        *string `json:"url,omitempty"`
	Tag        *string `json:"tag,omitempty"`
}

// Webhook is an object that represents a webhook for a container registry.
type Webhook struct {
	autorest.Response  `json:"-"`
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Type               *string             `json:"type,omitempty"`
	Location           *string             `json:"location,omitempty"`
	Tags               *map[string]*string `json:"tags,omitempty"`
	*WebhookProperties `json:"properties,omitempty"`
}

// WebhookCreateParameters is the parameters for creating a webhook.
type WebhookCreateParameters struct {
	Tags                               *map[string]*string `json:"tags,omitempty"`
	Location                           *string             `json:"location,omitempty"`
	*WebhookPropertiesCreateParameters `json:"properties,omitempty"`
}

// WebhookListResult is the result of a request to list webhooks for a container registry.
type WebhookListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Webhook `json:"value,omitempty"`
	NextLink          *string    `json:"nextLink,omitempty"`
}

// WebhookListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client WebhookListResult) WebhookListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// WebhookProperties is the properties of a webhook.
type WebhookProperties struct {
	Status            WebhookStatus     `json:"status,omitempty"`
	Scope             *string           `json:"scope,omitempty"`
	Actions           *[]WebhookAction  `json:"actions,omitempty"`
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// WebhookPropertiesCreateParameters is the parameters for creating the properties of a webhook.
type WebhookPropertiesCreateParameters struct {
	ServiceURI    *string             `json:"serviceUri,omitempty"`
	CustomHeaders *map[string]*string `json:"customHeaders,omitempty"`
	Status        WebhookStatus       `json:"status,omitempty"`
	Scope         *string             `json:"scope,omitempty"`
	Actions       *[]WebhookAction    `json:"actions,omitempty"`
}

// WebhookPropertiesUpdateParameters is the parameters for updating the properties of a webhook.
type WebhookPropertiesUpdateParameters struct {
	ServiceURI    *string             `json:"serviceUri,omitempty"`
	CustomHeaders *map[string]*string `json:"customHeaders,omitempty"`
	Status        WebhookStatus       `json:"status,omitempty"`
	Scope         *string             `json:"scope,omitempty"`
	Actions       *[]WebhookAction    `json:"actions,omitempty"`
}

// WebhookUpdateParameters is the parameters for updating a webhook.
type WebhookUpdateParameters struct {
	Tags                               *map[string]*string `json:"tags,omitempty"`
	*WebhookPropertiesUpdateParameters `json:"properties,omitempty"`
}
