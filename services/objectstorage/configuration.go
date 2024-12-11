/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

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
				URL:         "https://object-storage.api.{region}stackit.cloud",
				Description: "No description provided",
				Variables: map[string]config.ServerVariable{
					"region": {
						Description:  "No description provided",
						DefaultValue: "eu01.",
						EnumValues: []string{
							"eu01.",
						},
					},
				},
			},
		},
		OperationServers: map[string]config.ServerConfigurations{},
	}
	return cfg
}
