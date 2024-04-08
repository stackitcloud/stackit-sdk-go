/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.6.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type SessionPersistence struct {
	// If enabled then all connections from one source IP address are redirected to the same target. This setting changes the load balancing algorithm to Maglev.
	UseSourceIpAddress *bool `json:"useSourceIpAddress,omitempty"`
}
