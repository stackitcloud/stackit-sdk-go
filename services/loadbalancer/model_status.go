/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type Status struct {
	// The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	// Can be cast to int32 without loss of precision.
	Code *int64 `json:"code,omitempty"`
	// A list of messages that carry the error details.  There is a common set of message types for APIs to use.
	Details *[]GoogleProtobufAny `json:"details,omitempty"`
	// A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message *string `json:"message,omitempty"`
}