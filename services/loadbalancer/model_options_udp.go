/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.6.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type OptionsUDP struct {
	// The connection idle timeout to be used with the protocol. The default value is set to 1 minute, and the maximum value is 2 minutes.
	IdleTimeout *string `json:"idleTimeout,omitempty"`
}
