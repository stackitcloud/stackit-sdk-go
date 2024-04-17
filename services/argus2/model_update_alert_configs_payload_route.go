/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus2

type UpdateAlertConfigsPayloadRoute struct {
	// The labels by which incoming alerts are grouped together. For example, multiple alerts coming in for cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels use the special value '...' as the sole label name, for example: group_by: ['...']. This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
	GroupBy *[]string `json:"groupBy,omitempty"`
	// How long to wait before sending a notification about new alerts that are added to a group of alerts for which an initial notification has already been sent. (Usually ~5m or more.) `Additional Validators:` * must be a valid time format
	GroupInterval *string `json:"groupInterval,omitempty"`
	// How long to initially wait to send a notification for a group of alerts. Allows to wait for an inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.) `Additional Validators:` * must be a valid time format
	GroupWait *string `json:"groupWait,omitempty"`
	// map of key:value. A set of equality matchers an alert has to fulfill to match the node.  `Additional Validators:` * should not contain more than 5 keys * each key and value should not be longer than 200 characters * key and values should only include the characters: a-zA-Z0-9_./@&?:-
	Match *map[string]interface{} `json:"match,omitempty"`
	// map of key:value. A set of regex-matchers an alert has to fulfill to match the node.  `Additional Validators:` * should not contain more than 5 keys * each key and value should not be longer than 200 characters
	MatchRe *map[string]interface{} `json:"matchRe,omitempty"`
	// A list of matchers that an alert has to fulfill to match the node. A matcher is a string with a syntax inspired by PromQL and OpenMetrics. The syntax of a matcher consists of three tokens: * A valid Prometheus label name. * One of =, !=, =~, or !~. = means equals, != means that the strings are not equal, =~ is used for equality of regex expressions and !~ is used for un-equality of regex expressions. They have the same meaning as known from PromQL selectors. * A UTF-8 string, which may be enclosed in double quotes. Before or after each token, there may be any amount of whitespace. `Additional Validators:` * should not contain more than 5 keys * each key and value should not be longer than 200 characters
	Matchers *[]string `json:"matchers,omitempty"`
	// Receiver that should be one item of receivers `Additional Validators:` * must be a in name of receivers
	// REQUIRED
	Receiver *string `json:"receiver"`
	// How long to wait before sending a notification again if it has already been sent successfully for an alert. (Usually ~3h or more). `Additional Validators:` * must be a valid time format
	RepeatInterval *string `json:"repeatInterval,omitempty"`
	// Zero or more child routes.
	Routes *[]UpdateAlertConfigsPayloadRouteRoutesInner `json:"routes,omitempty"`
}
