package vips

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// SessionPersistence represents the session persistence feature of the load
// balancing service. It attempts to force connections or requests in the same
// session to be processed by the same member as long as it is ative. Three
// types of persistence are supported:
//
// SOURCE_IP:   With this mode, all connections originating from the same source
//              IP address, will be handled by the same member of the pool.
// HTTP_COOKIE: With this persistence mode, the load balancing function will
//              create a cookie on the first request from a client. Subsequent
//              requests containing the same cookie value will be handled by
//              the same member of the pool.
// APP_COOKIE:  With this persistence mode, the load balancing function will
//              rely on a cookie established by the backend application. All
//              requests carrying the same cookie value will be handled by the
//              same member of the pool.
type SessionPersistence struct {
	// The type of persistence mode
	Type string `json:"type"`

	// Name of cookie if persistence mode is set appropriately
	CookieName string `json:"cookie_name,omitempty"`
}

// VirtualIP is the primary load balancing configuration object that specifies
// the virtual IP address and port on which client traffic is received, as well
// as other details such as the load balancing method to be use, protocol, etc.
// This entity is sometimes known in LB products under the name of a "virtual
// server", a "vserver" or a "listener".
type VirtualIP struct {
	// The unique ID for the VIP.
	ID string `json:"id"`

	// Owner of the VIP. Only an admin user can specify a tenant ID other than its own.
	TenantID string `json:"tenant_id"`

	// Human-readable name for the VIP. Does not have to be unique.
	Name string `json:"name"`

	// Human-readable description for the VIP.
	Description string `json:"description"`

	// The ID of the subnet on which to allocate the VIP address.
	SubnetID string `json:"subnet_id"`

	// The IP address of the VIP.
	Address string `json:"address"`

	// The protocol of the VIP address. A valid value is TCP, HTTP, or HTTPS.
	Protocol string `json:"protocol"`

	// The port on which to listen to client traffic that is associated with the
	// VIP address. A valid value is from 0 to 65535.
	ProtocolPort int `json:"protocol_port"`

	// The ID of the pool with which the VIP is associated.
	PoolID string `json:"pool_id"`

	// The ID of the port which belongs to the load balancer
	PortID string `json:"port_id"`

	// Indicates whether connections in the same session will be processed by the
	// same pool member or not.
	Persistence SessionPersistence `json:"session_persistence"`

	// The maximum number of connections allowed for the VIP. Default is -1,
	// meaning no limit.
	ConnLimit int `json:"connection_limit"`

	// The administrative state of the VIP. A valid value is true (UP) or false (DOWN).
	AdminStateUp bool `json:"admin_state_up"`

	// The status of the VIP. Indicates whether the VIP is operational.
	Status string `json:"status"`
}

// VIPPage is the page returned by a pager when traversing over a
// collection of routers.
type VIPPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of routers has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r VIPPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"vips_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a VIPPage struct is empty.
func (r VIPPage) IsEmpty() (bool, error) {
	is, err := ExtractVIPs(r)
	return len(is) == 0, err
}

// ExtractVIPs accepts a Page struct, specifically a VIPPage struct,
// and extracts the elements into a slice of VirtualIP structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractVIPs(r pagination.Page) ([]VirtualIP, error) {
	var s struct {
		VIPs []VirtualIP `json:"vips"`
	}
	err := (r.(VIPPage)).ExtractInto(&s)
	return s.VIPs, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a router.
func (r commonResult) Extract() (*VirtualIP, error) {
	var s struct {
		VirtualIP *VirtualIP `json:"vip" json:"vip"`
	}
	err := r.ExtractInto(&s)
	return s.VirtualIP, err
}

// CreateResult represents the result of a create operation.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}
