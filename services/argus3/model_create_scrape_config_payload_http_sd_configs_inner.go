/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus3

type CreateScrapeConfigPayloadHttpSdConfigsInner struct {
	BasicAuth *CreateScrapeConfigPayloadBasicAuth                `json:"basicAuth,omitempty"`
	Oauth2    *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 `json:"oauth2,omitempty"`
	// Refresh interval to re-query the endpoint. E.g. 60s `Additional Validators:` * must be a valid time format* must be >= 60s
	RefreshInterval *string                                                     `json:"refreshInterval,omitempty"`
	TlsConfig       *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig `json:"tlsConfig,omitempty"`
	// URL from which the targets are fetched.
	// REQUIRED
	Url *string `json:"url"`
}
