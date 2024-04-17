/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus3

type UpdateAlertConfigsPayloadRouteRoutesInner struct {
	GroupBy *[]string `json:"groupBy,omitempty"`
	// As in one level above
	GroupInterval *string `json:"groupInterval,omitempty"`
	// As in one level above
	GroupWait *string `json:"groupWait,omitempty"`
	// As in one level above
	Match *map[string]interface{} `json:"match,omitempty"`
	// As in one level above
	MatchRe *map[string]interface{} `json:"matchRe,omitempty"`
	// As in one level above
	Receiver *string `json:"receiver,omitempty"`
	// As in one level above
	RepeatInterval *string `json:"repeatInterval,omitempty"`
	// Another child routes
	Routes *[]map[string]interface{} `json:"routes,omitempty"`
}
