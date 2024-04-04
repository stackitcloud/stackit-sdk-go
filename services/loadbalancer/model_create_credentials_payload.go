/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.6.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type CreateCredentialsPayload struct {
	// Credential name
	DisplayName *string `json:"displayName,omitempty"`
	// A valid password used for an existing ARGUS instance, which is used during basic auth.
	Password *string `json:"password,omitempty"`
	// A valid username used for an existing ARGUS instance, which is used during basic auth.
	Username *string `json:"username,omitempty"`
}
