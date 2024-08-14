/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

// Kubernetes For valid versions please take a look at [provider-options](#tag/ProviderOptions/operation/SkeService_GetProviderOptions) `kubernetesVersions`.
type Kubernetes struct {
	// DEPRECATED as of Kubernetes 1.25+ Flag to specify if privileged mode for containers is enabled or not. This should be used with care since it also disables a couple of other features like the use of some volume type (e.g. PVCs). By default this is set to true.
	AllowPrivilegedContainers *bool `json:"allowPrivilegedContainers,omitempty"`
	// REQUIRED
	Version *string `json:"version"`
}
