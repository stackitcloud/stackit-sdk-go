/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  This beta load balancer service is provided free of charge. For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type Listener struct {
	DisplayName *string `json:"displayName,omitempty"`
	// Port number where we listen for traffic
	Port *int32 `json:"port,omitempty"`
	// Protocol is the highest network protocol we understand to load balance. Currently only PROTOCOL_TCP and PROTOCOL_TCP_PROXY are supported.
	Protocol *int32 `json:"protocol,omitempty"`
	// Reference target pool by target pool name.
	TargetPool *string `json:"targetPool,omitempty"`
}
