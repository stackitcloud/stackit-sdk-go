/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

// KubernetesVersion struct for KubernetesVersion
type KubernetesVersion struct {
	ExpirationDate *string            `json:"expirationDate,omitempty"`
	FeatureGates   *map[string]string `json:"featureGates,omitempty"`
	State          *string            `json:"state,omitempty"`
	Version        *string            `json:"version,omitempty"`
}
