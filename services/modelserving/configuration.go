/*
STACKIT Model Serving API

This API provides endpoints for the model serving api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package modelserving

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
				URL:         "https://model-serving.api.stackit.cloud",
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
