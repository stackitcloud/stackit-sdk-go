/*
STACKIT MariaDB API

The STACKIT MariaDB API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mariadb

type Plan struct {
	// REQUIRED
	Description *string `json:"description"`
	// REQUIRED
	Free *bool `json:"free"`
	// REQUIRED
	Id *string `json:"id"`
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	SkuName *string `json:"skuName"`
}
