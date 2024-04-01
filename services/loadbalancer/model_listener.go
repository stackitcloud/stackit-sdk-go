/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.6.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type Listener struct {
	DisplayName *string `json:"displayName,omitempty"`
	// Will be used to reference a listener and will replace display name in the future. Currently uses <protocol>-<port> as the name if no display name is given.
	Name *string `json:"name,omitempty"`
	// Port number where we listen for traffic
	Port *int64 `json:"port,omitempty"`
	// Protocol is the highest network protocol we understand to load balance. Currently only PROTOCOL_TCP, PROTOCOL_TCP_PROXY and PROTOCOL_TLS_PASSTHROUGH are supported.
	Protocol *string `json:"protocol,omitempty"`
	// Server Name Idicators config for domains to be routed to the desired target pool for this listener.
	ServerNameIndicators *[]ServerNameIndicator `json:"serverNameIndicators,omitempty"`
	// Reference target pool by target pool name.
	TargetPool *string     `json:"targetPool,omitempty"`
	Tcp        *OptionsTCP `json:"tcp,omitempty"`
	Udp        *OptionsUDP `json:"udp,omitempty"`
}
