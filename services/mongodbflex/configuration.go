/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

import (
	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

// NewConfiguration returns a new Configuration object
func NewConfiguration() *config.Configuration {
	cfg := &config.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "stackit-sdk-go/mongodbflex",
		Debug:         false,
		Servers: config.ServerConfigurations{
			{
				URL:         "https://mongodb-flex-service.api.stackit.cloud",
				Description: "No description provided",
				Variables: map[string]config.ServerVariable{
					"region": {
						Description:  "No description provided",
						DefaultValue: "global",
					},
				},
			},
		},
		OperationServers: map[string]config.ServerConfigurations{},
	}
	return cfg
}
