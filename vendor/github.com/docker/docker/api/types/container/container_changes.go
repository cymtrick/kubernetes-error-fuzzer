package container // import "github.com/docker/docker/api/types/container"

// ----------------------------------------------------------------------------
// Code generated by `swagger generate operation`. DO NOT EDIT.
//
// See hack/generate-swagger-api.sh
// ----------------------------------------------------------------------------

// ContainerChangeResponseItem change item in response to ContainerChanges operation
// swagger:model ContainerChangeResponseItem
type ContainerChangeResponseItem struct {

	// Kind of change
	// Required: true
	Kind uint8 `json:"Kind"`

	// Path to file that has changed
	// Required: true
	Path string `json:"Path"`
}
