/*
STACKIT RabbitMQ API

The STACKIT RabbitMQ API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rabbitmq

// Instance struct for Instance
type Instance struct {
	// REQUIRED
	CfGuid *string `json:"cfGuid"`
	// REQUIRED
	CfOrganizationGuid *string `json:"cfOrganizationGuid"`
	// REQUIRED
	CfSpaceGuid *string `json:"cfSpaceGuid"`
	// REQUIRED
	DashboardUrl *string `json:"dashboardUrl"`
	// REQUIRED
	ImageUrl   *string `json:"imageUrl"`
	InstanceId *string `json:"instanceId,omitempty"`
	// REQUIRED
	LastOperation *InstanceLastOperation `json:"lastOperation"`
	// REQUIRED
	Name *string `json:"name"`
	// Deprecated: Check the GitHub changelog for alternatives
	// REQUIRED
	OfferingName *string `json:"offeringName"`
	// REQUIRED
	OfferingVersion *string `json:"offeringVersion"`
	// REQUIRED
	Parameters *map[string]interface{} `json:"parameters"`
	// REQUIRED
	PlanId *string `json:"planId"`
	// REQUIRED
	PlanName *string `json:"planName"`
	Status   *string `json:"status,omitempty"`
}
