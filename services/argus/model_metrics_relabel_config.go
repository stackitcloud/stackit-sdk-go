/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

// MetricsRelabelConfig struct for MetricsRelabelConfig
type MetricsRelabelConfig struct {
	Action      *string `json:"action,omitempty"`
	Modulus     *int64  `json:"modulus,omitempty"`
	Regex       *string `json:"regex,omitempty"`
	Replacement *string `json:"replacement,omitempty"`
	Separator   *string `json:"separator,omitempty"`
	// REQUIRED
	SourceLabels *[]string `json:"sourceLabels"`
	TargetLabel  *string   `json:"targetLabel,omitempty"`
}
