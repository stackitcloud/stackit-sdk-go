/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type LoadBalancerOptions struct {
	AccessControl    *LoadbalancerOptionAccessControl `json:"accessControl,omitempty"`
	EphemeralAddress *bool                            `json:"ephemeralAddress,omitempty"`
	Observability    *LoadbalancerOptionObservability `json:"observability,omitempty"`
	// Load Balancer is accessible only via a private network ip address. Not changeable after creation.
	PrivateNetworkOnly *bool `json:"privateNetworkOnly,omitempty"`
}
