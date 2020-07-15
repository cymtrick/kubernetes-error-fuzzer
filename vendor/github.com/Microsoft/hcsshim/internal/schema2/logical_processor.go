/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type LogicalProcessor struct {
	LpIndex     uint32 `json:"LogicalProcessorCount,omitempty"`
	NodeNumber  uint8  `json:"NodeNumber, omitempty"`
	PackageId   uint32 `json:"PackageId, omitempty"`
	CoreId      uint32 `json:"CoreId, omitempty"`
	RootVpIndex int32  `json:"RootVpIndex, omitempty"`
}
