/*
Resource Manager API

API v2 to manage resource containers - organizations, folders, projects incl. labels  ### Resource Management STACKIT resource management handles the terms _Organization_, _Folder_, _Project_, _Label_, and the hierarchical structure between them. Technically, organizations,  folders, and projects are _Resource Containers_ to which a _Label_ can be attached to. The STACKIT _Resource Manager_ provides CRUD endpoints to query and to modify the state.  ### Resource Hierarchy STACKIT resource hierarchy defines the relationship of resource containers as tree structure with organizations as the root node. Folders and projects are optional  child elements referencing the organization as parent. STACKIT resource hierarchy allows to structure cloud-resources clearly providing flexibility and individuality  for fine grained access control, access inheritance, and policies.  The STACKIT resource hierarchy model can be compared to the folder concept of a Unix file system. A folder can have exactly one parent. Folder nesting allows to organize and to structure content while defining  fine grained access permissions per folder. Within STACKIT resource hierarchy access is inherited, meaning if you have access to a resource container, you also have access to its child containers. - Users are assigned to resource containers as members by role - A user inherits permissions to all child containers  ### Organizations STACKIT organizations are the base element to create and to use cloud-resources. An organization is bound to one customer account. Organizations have a lifecycle. - Organizations are always the root node in resource hierarchy and do not have a parent  ### Folders STACKIT folders allow to organize cloud-resources and to define fine-grained access control. A folder might represent departments, teams, user groups, components etc.  Folders do not have a lifecycle as they act as structural element only. - Folders are optional - A folder can be created having either an organization, or a folder as parent - Folder names under the same parent must be unique (case insensitive) - Root organization cannot be changed  ### Projects STACKIT projects are needed to use cloud-resources. Projects serve as wrapper for underlying technical structures and processes. Projects have a lifecycle. Projects compared to folders may have different policies. - Projects are optional, but mandatory for cloud-resource usage - A project can be created having either an organization, or a folder as parent - A project must not have a project as parent - Project names under the same parent must not be unique - Root organization cannot be changed  ### Label STACKIT labels are key-value pairs including a resource container reference. Labels can be defined and attached freely to resource containers by which resources can be organized and queried. - Policy-based, immutable labels may exists  ### Current Limits - Vertically - Maximum folder nesting level: 5 - Horizontally - Maximum number of folders under one parent: 150 - Maximum number of projects under one organization: 2.500 - Maximum number of labels attached to one container: 100

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resourcemanager

import (
	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

// NewConfiguration returns a new Configuration object
func NewConfiguration() *config.Configuration {
	cfg := &config.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: config.ServerConfigurations{
			{
				URL:         "https://resource-manager.api.{region}stackit.cloud",
				Description: "No description provided",
				Variables: map[string]config.ServerVariable{
					"region": {
						Description:  "No description provided",
						DefaultValue: "",
					},
				},
			},
		},
		OperationServers: map[string]config.ServerConfigurations{},
	}
	return cfg
}
