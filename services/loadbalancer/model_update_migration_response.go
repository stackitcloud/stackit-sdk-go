/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.6.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

type UpdateMigrationResponse struct {
	LoadBalancer    *LoadBalancer    `json:"loadBalancer,omitempty"`
	MigrationParams *MigrationParams `json:"migrationParams,omitempty"`
	MigrationState  *string          `json:"migrationState,omitempty"`
}
