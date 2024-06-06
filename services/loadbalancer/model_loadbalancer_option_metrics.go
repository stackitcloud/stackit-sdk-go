/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type LoadbalancerOptionMetrics struct {
	// Credentials reference for metrics. This reference is created via the observability create endpoint and the credential needs to contain the basic auth username and password for the metrics solution the push URL points to. Then this enables monitoring via remote write for the Load Balancer.
	CredentialsRef *string `json:"credentialsRef,omitempty"`
	// The ARGUS/Prometheus remote write Push URL you want the metrics to be shipped to.
	PushUrl *string `json:"pushUrl,omitempty"`
}
