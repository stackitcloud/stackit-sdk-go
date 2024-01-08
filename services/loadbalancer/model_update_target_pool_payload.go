/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.4.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type UpdateTargetPoolPayload struct {
	ActiveHealthCheck *ActiveHealthCheck `json:"activeHealthCheck,omitempty"`
	// Target pool name
	Name *string `json:"name,omitempty"`
	// Identical port number where each target listens for traffic.
	TargetPort *int64 `json:"targetPort,omitempty"`
	// List of all targets which will be used in the pool. Limited to 250.
	Targets *[]Target `json:"targets,omitempty"`
}
