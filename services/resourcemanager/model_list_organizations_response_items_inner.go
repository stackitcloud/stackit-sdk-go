/*
Resource Manager API

API v2 to manage resource containers - organizations, folders, projects incl. labels  ### Resource Management STACKIT resource management handles the terms _Organization_, _Folder_, _Project_, _Label_, and the hierarchical structure between them. Technically, organizations,  folders, and projects are _Resource Containers_ to which a _Label_ can be attached to. The STACKIT _Resource Manager_ provides CRUD endpoints to query and to modify the state.  ### Organizations STACKIT organizations are the base element to create and to use cloud-resources. An organization is bound to one customer account. Organizations have a lifecycle. - Organizations are always the root node in resource hierarchy and do not have a parent  ### Projects STACKIT projects are needed to use cloud-resources. Projects serve as wrapper for underlying technical structures and processes. Projects have a lifecycle. Projects compared to folders may have different policies. - Projects are optional, but mandatory for cloud-resource usage - A project can be created having either an organization, or a folder as parent - A project must not have a project as parent - Project names under the same parent must not be unique - Root organization cannot be changed  ### Label STACKIT labels are key-value pairs including a resource container reference. Labels can be defined and attached freely to resource containers by which resources can be organized and queried. - Policy-based, immutable labels may exists

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resourcemanager

import (
	"time"
)

type ListOrganizationsResponseItemsInner struct {
	// Globally unique, user-friendly identifier.
	// REQUIRED
	ContainerId *string `json:"containerId"`
	// Timestamp at which the organization was created.
	// REQUIRED
	CreationTime *time.Time `json:"creationTime"`
	// Labels are key-value string pairs that can be attached to a resource container. Some labels may be enforced via policies.  - A label key must match the regex `[A-ZÄÜÖa-zäüöß0-9_-]{1,64}`. - A label value must match the regex `^$|[A-ZÄÜÖa-zäüöß0-9_-]{1,64}`.
	Labels *map[string]string `json:"labels,omitempty"`
	// REQUIRED
	LifecycleState *LifecycleState `json:"lifecycleState"`
	// Name of the organization.
	// REQUIRED
	Name *string `json:"name"`
	// Globally unique, organization identifier.
	// REQUIRED
	OrganizationId *string `json:"organizationId"`
	// Timestamp at which the organization was last modified.
	// REQUIRED
	UpdateTime *time.Time `json:"updateTime"`
}
