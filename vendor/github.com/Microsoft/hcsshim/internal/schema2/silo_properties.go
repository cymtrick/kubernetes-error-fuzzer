/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

//  Silo job information
type SiloProperties struct {
	Enabled bool `json:"Enabled,omitempty"`

	JobName string `json:"JobName,omitempty"`
}
