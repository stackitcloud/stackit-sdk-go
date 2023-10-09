/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type UpdateInstanceAlertConfigsPayloadInhibitRules struct {
	// Labels that must have an equal value in the source and target alert for the inhibition to take effect. `Additional Validators:` * should only include the characters: a-zA-Z0-9_./@&?:-
	Equal *[]string `json:"equal,omitempty"`
	// map of key:value. Matchers for which one or more alerts have to exist for the inhibition to take effect. `Additional Validators:` * should not contain more than 5 keys * each key and value should not have more than 200 characters * each key and value should only include the characters: a-zA-Z0-9_./@&?:-
	SourceMatch *map[string]interface{} `json:"sourceMatch,omitempty"`
	// map of key:value. Regex match `Additional Validators:` * should not contain more than 5 keys * each key and value should not have more than 200 characters
	SourceMatchRe *map[string]interface{} `json:"sourceMatchRe,omitempty"`
	// map of key:value. Matchers that have to be fulfilled in the alerts to be muted. `Additional Validators:` * should not contain more than 5 keys * each key and value should not have more than 200 characters * each key and value should only include the characters: a-zA-Z0-9_./@&?:-
	TargetMatch *map[string]interface{} `json:"targetMatch,omitempty"`
	// map of key:value. Matchers that have to be fulfilled in the alerts to be muted. Regex. `Additional Validators:` * should not contain more than 5 keys * each key and value should not have more than 200 characters
	TargetMatchRe *map[string]interface{} `json:"targetMatchRe,omitempty"`
}
