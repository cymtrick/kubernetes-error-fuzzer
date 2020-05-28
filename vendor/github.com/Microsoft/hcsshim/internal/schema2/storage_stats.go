/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

//  Storage runtime statistics
type StorageStats struct {
	ReadCountNormalized uint64 `json:"ReadCountNormalized,omitempty"`

	ReadSizeBytes uint64 `json:"ReadSizeBytes,omitempty"`

	WriteCountNormalized uint64 `json:"WriteCountNormalized,omitempty"`

	WriteSizeBytes uint64 `json:"WriteSizeBytes,omitempty"`
}
