/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type VmMemory struct {

	AvailableMemory int32 `json:"AvailableMemory,omitempty"`

	AvailableMemoryBuffer int32 `json:"AvailableMemoryBuffer,omitempty"`

	ReservedMemory int32 `json:"ReservedMemory,omitempty"`

	AssignedMemory int32 `json:"AssignedMemory,omitempty"`

	SlpActive bool `json:"SlpActive,omitempty"`

	BalancingEnabled bool `json:"BalancingEnabled,omitempty"`

	DmOperationInProgress bool `json:"DmOperationInProgress,omitempty"`
}
