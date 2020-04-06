package volume // import "github.com/docker/docker/api/types/volume"

// ----------------------------------------------------------------------------
// DO NOT EDIT THIS FILE
// This file was generated by `swagger generate operation`
//
// See hack/generate-swagger-api.sh
// ----------------------------------------------------------------------------

import "github.com/docker/docker/api/types"

// VolumeListOKBody Volume list response
// swagger:model VolumeListOKBody
type VolumeListOKBody struct {

	// List of volumes
	// Required: true
	Volumes []*types.Volume `json:"Volumes"`

	// Warnings that occurred when fetching the list of volumes
	// Required: true
	Warnings []string `json:"Warnings"`
}
