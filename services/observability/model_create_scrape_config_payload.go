/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// CreateScrapeConfigPayload struct for CreateScrapeConfigPayload
type CreateScrapeConfigPayload struct {
	BasicAuth *CreateScrapeConfigPayloadBasicAuth `json:"basicAuth,omitempty"`
	// Sets the 'Authorization' header on every scrape request with the configured bearer token. It is mutually exclusive with 'bearer_token_file'. `Additional Validators:` * needs to be a valid bearer token * if bearerToken is in the body no other authentication method should be in the body
	BearerToken *string `json:"bearerToken,omitempty"`
	// Note that any globally configured 'external_labels' are unaffected by this setting. In communication with external systems, they are always applied only when a time series does not have a given label yet and are ignored otherwise.
	HonorLabels *bool `json:"honorLabels,omitempty"`
	// honor_timestamps controls whether Prometheus respects the timestamps present in scraped data. If honor_timestamps is set to 'true', the timestamps of the metrics exposed by the target will be used.
	HonorTimeStamps *bool `json:"honorTimeStamps,omitempty"`
	// HTTP-based service discovery provides a more generic way to configure static targets and serves as an interface to plug in custom service discovery mechanisms.
	HttpSdConfigs *[]CreateScrapeConfigPayloadHttpSdConfigsInner `json:"httpSdConfigs,omitempty"`
	// The job name assigned to scraped metrics by default. `Additional Validators:` * must be unique * key and values should only include the characters: a-zA-Z0-9-
	// REQUIRED
	JobName *string `json:"jobName"`
	// The HTTP resource path on which to fetch metrics from targets. E.g. /metrics
	MetricsPath *string `json:"metricsPath,omitempty"`
	// List of metric relabel configurations
	MetricsRelabelConfigs *[]CreateScrapeConfigPayloadMetricsRelabelConfigsInner `json:"metricsRelabelConfigs,omitempty"`
	Oauth2                *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2     `json:"oauth2,omitempty"`
	// Optional http params `Additional Validators:` * should not contain more than 5 keys * each key and value should not have more than 200 characters
	Params *map[string]interface{} `json:"params,omitempty"`
	// Per-scrape limit on number of scraped samples that will be accepted. If more than this number of samples are present after metric relabeling the entire scrape will be treated as failed. The total limit depends on the service plan target limits * samples
	SampleLimit *float64 `json:"sampleLimit,omitempty"`
	// Configures the protocol scheme used for requests. https or http
	// REQUIRED
	Scheme *string `json:"scheme"`
	// How frequently to scrape targets from this job. E.g. 5m `Additional Validators:` * must be a valid time format* must be >= 60s
	// REQUIRED
	ScrapeInterval *string `json:"scrapeInterval"`
	// Per-scrape timeout when scraping this job. `Additional Validators:` * must be a valid time format* must be smaller than scrapeInterval
	// REQUIRED
	ScrapeTimeout *string `json:"scrapeTimeout"`
	// A list of scrape configurations.
	// REQUIRED
	StaticConfigs *[]CreateScrapeConfigPayloadStaticConfigsInner              `json:"staticConfigs"`
	TlsConfig     *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig `json:"tlsConfig,omitempty"`
}
