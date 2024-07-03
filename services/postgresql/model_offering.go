/*
STACKIT PostgreSQL API

The STACKIT PostgreSQL API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresql

type Offering struct {
	// REQUIRED
	Description *string `json:"description"`
	// REQUIRED
	DocumentationUrl *string `json:"documentationUrl"`
	// REQUIRED
	ImageUrl *string `json:"imageUrl"`
	// REQUIRED
	Latest    *bool   `json:"latest"`
	Lifecycle *string `json:"lifecycle,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Plans *[]Plan `json:"plans"`
	// REQUIRED
	QuotaCount *int64          `json:"quotaCount"`
	Schema     *InstanceSchema `json:"schema,omitempty"`
	// REQUIRED
	Version *string `json:"version"`
}
