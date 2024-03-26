/*
STACKIT Redis API

The STACKIT Redis API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redis

type PartialUpdateInstancePayload struct {
	InstanceName *string                  `json:"instanceName,omitempty"`
	Parameters   *RedisInstanceParameters `json:"parameters,omitempty"`
	PlanId       *string                  `json:"planId,omitempty"`
}
