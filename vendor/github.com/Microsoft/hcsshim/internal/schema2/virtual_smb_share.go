/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type VirtualSmbShare struct {
	Name string `json:"Name,omitempty"`

	Path string `json:"Path,omitempty"`

	AllowedFiles []string `json:"AllowedFiles,omitempty"`

	Options *VirtualSmbShareOptions `json:"Options,omitempty"`
}
