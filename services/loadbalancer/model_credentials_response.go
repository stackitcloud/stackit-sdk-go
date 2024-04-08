/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.6.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type CredentialsResponse struct {
	// The credentials reference can be used for observability of the Load Balancer.
	CredentialsRef *string `json:"credentialsRef,omitempty"`
	// Credential name
	DisplayName *string `json:"displayName,omitempty"`
	// The username used for the ARGUS instance
	Username *string `json:"username,omitempty"`
}
