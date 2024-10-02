/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"time"
)

// Server Representation of a single server object.
type Server struct {
	// Object that represents an availability zone.
	AvailabilityZone *string                        `json:"availabilityZone,omitempty"`
	BootVolume       *CreateServerPayloadBootVolume `json:"bootVolume,omitempty"`
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// An error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Universally Unique Identifier (UUID).
	Image *string `json:"image,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Keypair *string `json:"keypair,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// Date-time when resource was launched.
	LaunchedAt *time.Time `json:"launchedAt,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	MachineType       *string            `json:"machineType"`
	MaintenanceWindow *ServerMaintenance `json:"maintenanceWindow,omitempty"`
	// The name for a Server.
	// REQUIRED
	Name       *string                        `json:"name"`
	Networking *CreateServerPayloadNetworking `json:"networking,omitempty"`
	// A list of networks attached to a server.
	Nics *[]ServerNetwork `json:"nics,omitempty"`
	// The power status of a server.
	PowerStatus *string `json:"powerStatus,omitempty"`
	// A list of General Objects.
	SecurityGroups *[]string `json:"securityGroups,omitempty"`
	// Universally Unique Identifier (UUID).
	ServerGroup *string `json:"serverGroup,omitempty"`
	// A list of service account mails.
	ServiceAccountMails *[]string `json:"serviceAccountMails,omitempty"`
	// The status of a server object.
	Status *string `json:"status,omitempty"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// User Data that is provided to the server. Must be base64 encoded and is passed via cloud-init to the server.
	UserData *string `json:"userData,omitempty"`
	// A list of UUIDs.
	Volumes *[]string `json:"volumes,omitempty"`
}
