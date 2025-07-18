/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateScrapeConfigPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateScrapeConfigPayload{}

/*
	types and functions for basicAuth
*/

// isModel
type UpdateScrapeConfigPayloadGetBasicAuthAttributeType = *CreateScrapeConfigPayloadBasicAuth
type UpdateScrapeConfigPayloadGetBasicAuthArgType = CreateScrapeConfigPayloadBasicAuth
type UpdateScrapeConfigPayloadGetBasicAuthRetType = CreateScrapeConfigPayloadBasicAuth

func getUpdateScrapeConfigPayloadGetBasicAuthAttributeTypeOk(arg UpdateScrapeConfigPayloadGetBasicAuthAttributeType) (ret UpdateScrapeConfigPayloadGetBasicAuthRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetBasicAuthAttributeType(arg *UpdateScrapeConfigPayloadGetBasicAuthAttributeType, val UpdateScrapeConfigPayloadGetBasicAuthRetType) {
	*arg = &val
}

/*
	types and functions for bearerToken
*/

// isNotNullableString
type UpdateScrapeConfigPayloadGetBearerTokenAttributeType = *string

func getUpdateScrapeConfigPayloadGetBearerTokenAttributeTypeOk(arg UpdateScrapeConfigPayloadGetBearerTokenAttributeType) (ret UpdateScrapeConfigPayloadGetBearerTokenRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetBearerTokenAttributeType(arg *UpdateScrapeConfigPayloadGetBearerTokenAttributeType, val UpdateScrapeConfigPayloadGetBearerTokenRetType) {
	*arg = &val
}

type UpdateScrapeConfigPayloadGetBearerTokenArgType = string
type UpdateScrapeConfigPayloadGetBearerTokenRetType = string

/*
	types and functions for honorLabels
*/

// isBoolean
type UpdateScrapeConfigPayloadgetHonorLabelsAttributeType = *bool
type UpdateScrapeConfigPayloadgetHonorLabelsArgType = bool
type UpdateScrapeConfigPayloadgetHonorLabelsRetType = bool

func getUpdateScrapeConfigPayloadgetHonorLabelsAttributeTypeOk(arg UpdateScrapeConfigPayloadgetHonorLabelsAttributeType) (ret UpdateScrapeConfigPayloadgetHonorLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadgetHonorLabelsAttributeType(arg *UpdateScrapeConfigPayloadgetHonorLabelsAttributeType, val UpdateScrapeConfigPayloadgetHonorLabelsRetType) {
	*arg = &val
}

/*
	types and functions for honorTimeStamps
*/

// isBoolean
type UpdateScrapeConfigPayloadgetHonorTimeStampsAttributeType = *bool
type UpdateScrapeConfigPayloadgetHonorTimeStampsArgType = bool
type UpdateScrapeConfigPayloadgetHonorTimeStampsRetType = bool

func getUpdateScrapeConfigPayloadgetHonorTimeStampsAttributeTypeOk(arg UpdateScrapeConfigPayloadgetHonorTimeStampsAttributeType) (ret UpdateScrapeConfigPayloadgetHonorTimeStampsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadgetHonorTimeStampsAttributeType(arg *UpdateScrapeConfigPayloadgetHonorTimeStampsAttributeType, val UpdateScrapeConfigPayloadgetHonorTimeStampsRetType) {
	*arg = &val
}

/*
	types and functions for metricsPath
*/

// isNotNullableString
type UpdateScrapeConfigPayloadGetMetricsPathAttributeType = *string

func getUpdateScrapeConfigPayloadGetMetricsPathAttributeTypeOk(arg UpdateScrapeConfigPayloadGetMetricsPathAttributeType) (ret UpdateScrapeConfigPayloadGetMetricsPathRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetMetricsPathAttributeType(arg *UpdateScrapeConfigPayloadGetMetricsPathAttributeType, val UpdateScrapeConfigPayloadGetMetricsPathRetType) {
	*arg = &val
}

type UpdateScrapeConfigPayloadGetMetricsPathArgType = string
type UpdateScrapeConfigPayloadGetMetricsPathRetType = string

/*
	types and functions for metricsRelabelConfigs
*/

// isArray
type UpdateScrapeConfigPayloadGetMetricsRelabelConfigsAttributeType = *[]CreateScrapeConfigPayloadMetricsRelabelConfigsInner
type UpdateScrapeConfigPayloadGetMetricsRelabelConfigsArgType = []CreateScrapeConfigPayloadMetricsRelabelConfigsInner
type UpdateScrapeConfigPayloadGetMetricsRelabelConfigsRetType = []CreateScrapeConfigPayloadMetricsRelabelConfigsInner

func getUpdateScrapeConfigPayloadGetMetricsRelabelConfigsAttributeTypeOk(arg UpdateScrapeConfigPayloadGetMetricsRelabelConfigsAttributeType) (ret UpdateScrapeConfigPayloadGetMetricsRelabelConfigsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetMetricsRelabelConfigsAttributeType(arg *UpdateScrapeConfigPayloadGetMetricsRelabelConfigsAttributeType, val UpdateScrapeConfigPayloadGetMetricsRelabelConfigsRetType) {
	*arg = &val
}

/*
	types and functions for params
*/

// isFreeform
type UpdateScrapeConfigPayloadGetParamsAttributeType = *map[string]interface{}
type UpdateScrapeConfigPayloadGetParamsArgType = map[string]interface{}
type UpdateScrapeConfigPayloadGetParamsRetType = map[string]interface{}

func getUpdateScrapeConfigPayloadGetParamsAttributeTypeOk(arg UpdateScrapeConfigPayloadGetParamsAttributeType) (ret UpdateScrapeConfigPayloadGetParamsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetParamsAttributeType(arg *UpdateScrapeConfigPayloadGetParamsAttributeType, val UpdateScrapeConfigPayloadGetParamsRetType) {
	*arg = &val
}

/*
	types and functions for sampleLimit
*/

// isNumber
type UpdateScrapeConfigPayloadGetSampleLimitAttributeType = *float64
type UpdateScrapeConfigPayloadGetSampleLimitArgType = float64
type UpdateScrapeConfigPayloadGetSampleLimitRetType = float64

func getUpdateScrapeConfigPayloadGetSampleLimitAttributeTypeOk(arg UpdateScrapeConfigPayloadGetSampleLimitAttributeType) (ret UpdateScrapeConfigPayloadGetSampleLimitRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetSampleLimitAttributeType(arg *UpdateScrapeConfigPayloadGetSampleLimitAttributeType, val UpdateScrapeConfigPayloadGetSampleLimitRetType) {
	*arg = &val
}

/*
	types and functions for scheme
*/

// isEnum

// UpdateScrapeConfigPayloadScheme Configures the protocol scheme used for requests. https or http
// value type for enums
type UpdateScrapeConfigPayloadScheme string

// List of Scheme
const (
	UPDATESCRAPECONFIGPAYLOADSCHEME_HTTP  UpdateScrapeConfigPayloadScheme = "http"
	UPDATESCRAPECONFIGPAYLOADSCHEME_HTTPS UpdateScrapeConfigPayloadScheme = "https"
)

// All allowed values of UpdateScrapeConfigPayload enum
var AllowedUpdateScrapeConfigPayloadSchemeEnumValues = []UpdateScrapeConfigPayloadScheme{
	"http",
	"https",
}

func (v *UpdateScrapeConfigPayloadScheme) UnmarshalJSON(src []byte) error {
	// use a type alias to prevent infinite recursion during unmarshal,
	// see https://biscuit.ninja/posts/go-avoid-an-infitine-loop-with-custom-json-unmarshallers
	type TmpJson UpdateScrapeConfigPayloadScheme
	var value TmpJson
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// Allow unmarshalling zero value for testing purposes
	var zeroValue TmpJson
	if value == zeroValue {
		return nil
	}
	enumTypeValue := UpdateScrapeConfigPayloadScheme(value)
	for _, existing := range AllowedUpdateScrapeConfigPayloadSchemeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpdateScrapeConfigPayload", value)
}

// NewUpdateScrapeConfigPayloadSchemeFromValue returns a pointer to a valid UpdateScrapeConfigPayloadScheme
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpdateScrapeConfigPayloadSchemeFromValue(v UpdateScrapeConfigPayloadScheme) (*UpdateScrapeConfigPayloadScheme, error) {
	ev := UpdateScrapeConfigPayloadScheme(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpdateScrapeConfigPayloadScheme: valid values are %v", v, AllowedUpdateScrapeConfigPayloadSchemeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpdateScrapeConfigPayloadScheme) IsValid() bool {
	for _, existing := range AllowedUpdateScrapeConfigPayloadSchemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SchemeScheme value
func (v UpdateScrapeConfigPayloadScheme) Ptr() *UpdateScrapeConfigPayloadScheme {
	return &v
}

type NullableUpdateScrapeConfigPayloadScheme struct {
	value *UpdateScrapeConfigPayloadScheme
	isSet bool
}

func (v NullableUpdateScrapeConfigPayloadScheme) Get() *UpdateScrapeConfigPayloadScheme {
	return v.value
}

func (v *NullableUpdateScrapeConfigPayloadScheme) Set(val *UpdateScrapeConfigPayloadScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateScrapeConfigPayloadScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateScrapeConfigPayloadScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateScrapeConfigPayloadScheme(val *UpdateScrapeConfigPayloadScheme) *NullableUpdateScrapeConfigPayloadScheme {
	return &NullableUpdateScrapeConfigPayloadScheme{value: val, isSet: true}
}

func (v NullableUpdateScrapeConfigPayloadScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateScrapeConfigPayloadScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type UpdateScrapeConfigPayloadGetSchemeAttributeType = *UpdateScrapeConfigPayloadScheme
type UpdateScrapeConfigPayloadGetSchemeArgType = UpdateScrapeConfigPayloadScheme
type UpdateScrapeConfigPayloadGetSchemeRetType = UpdateScrapeConfigPayloadScheme

func getUpdateScrapeConfigPayloadGetSchemeAttributeTypeOk(arg UpdateScrapeConfigPayloadGetSchemeAttributeType) (ret UpdateScrapeConfigPayloadGetSchemeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetSchemeAttributeType(arg *UpdateScrapeConfigPayloadGetSchemeAttributeType, val UpdateScrapeConfigPayloadGetSchemeRetType) {
	*arg = &val
}

/*
	types and functions for scrapeInterval
*/

// isNotNullableString
type UpdateScrapeConfigPayloadGetScrapeIntervalAttributeType = *string

func getUpdateScrapeConfigPayloadGetScrapeIntervalAttributeTypeOk(arg UpdateScrapeConfigPayloadGetScrapeIntervalAttributeType) (ret UpdateScrapeConfigPayloadGetScrapeIntervalRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetScrapeIntervalAttributeType(arg *UpdateScrapeConfigPayloadGetScrapeIntervalAttributeType, val UpdateScrapeConfigPayloadGetScrapeIntervalRetType) {
	*arg = &val
}

type UpdateScrapeConfigPayloadGetScrapeIntervalArgType = string
type UpdateScrapeConfigPayloadGetScrapeIntervalRetType = string

/*
	types and functions for scrapeTimeout
*/

// isNotNullableString
type UpdateScrapeConfigPayloadGetScrapeTimeoutAttributeType = *string

func getUpdateScrapeConfigPayloadGetScrapeTimeoutAttributeTypeOk(arg UpdateScrapeConfigPayloadGetScrapeTimeoutAttributeType) (ret UpdateScrapeConfigPayloadGetScrapeTimeoutRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetScrapeTimeoutAttributeType(arg *UpdateScrapeConfigPayloadGetScrapeTimeoutAttributeType, val UpdateScrapeConfigPayloadGetScrapeTimeoutRetType) {
	*arg = &val
}

type UpdateScrapeConfigPayloadGetScrapeTimeoutArgType = string
type UpdateScrapeConfigPayloadGetScrapeTimeoutRetType = string

/*
	types and functions for staticConfigs
*/

// isArray
type UpdateScrapeConfigPayloadGetStaticConfigsAttributeType = *[]UpdateScrapeConfigPayloadStaticConfigsInner
type UpdateScrapeConfigPayloadGetStaticConfigsArgType = []UpdateScrapeConfigPayloadStaticConfigsInner
type UpdateScrapeConfigPayloadGetStaticConfigsRetType = []UpdateScrapeConfigPayloadStaticConfigsInner

func getUpdateScrapeConfigPayloadGetStaticConfigsAttributeTypeOk(arg UpdateScrapeConfigPayloadGetStaticConfigsAttributeType) (ret UpdateScrapeConfigPayloadGetStaticConfigsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetStaticConfigsAttributeType(arg *UpdateScrapeConfigPayloadGetStaticConfigsAttributeType, val UpdateScrapeConfigPayloadGetStaticConfigsRetType) {
	*arg = &val
}

/*
	types and functions for tlsConfig
*/

// isModel
type UpdateScrapeConfigPayloadGetTlsConfigAttributeType = *CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig
type UpdateScrapeConfigPayloadGetTlsConfigArgType = CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig
type UpdateScrapeConfigPayloadGetTlsConfigRetType = CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig

func getUpdateScrapeConfigPayloadGetTlsConfigAttributeTypeOk(arg UpdateScrapeConfigPayloadGetTlsConfigAttributeType) (ret UpdateScrapeConfigPayloadGetTlsConfigRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setUpdateScrapeConfigPayloadGetTlsConfigAttributeType(arg *UpdateScrapeConfigPayloadGetTlsConfigAttributeType, val UpdateScrapeConfigPayloadGetTlsConfigRetType) {
	*arg = &val
}

// UpdateScrapeConfigPayload struct for UpdateScrapeConfigPayload
type UpdateScrapeConfigPayload struct {
	BasicAuth UpdateScrapeConfigPayloadGetBasicAuthAttributeType `json:"basicAuth,omitempty"`
	// Sets the 'Authorization' header on every scrape request with the configured bearer token. It is mutually exclusive with 'bearer_token_file'. `Additional Validators:` * needs to be a valid bearer token * if bearerToken is in the body no other authentication method should be in the body
	BearerToken UpdateScrapeConfigPayloadGetBearerTokenAttributeType `json:"bearerToken,omitempty"`
	// Note that any globally configured 'external_labels' are unaffected by this setting. In communication with external systems, they are always applied only when a time series does not have a given label yet and are ignored otherwise.
	HonorLabels UpdateScrapeConfigPayloadgetHonorLabelsAttributeType `json:"honorLabels,omitempty"`
	// honor_timestamps controls whether Prometheus respects the timestamps present in scraped data. If honor_timestamps is set to 'true', the timestamps of the metrics exposed by the target will be used.
	HonorTimeStamps UpdateScrapeConfigPayloadgetHonorTimeStampsAttributeType `json:"honorTimeStamps,omitempty"`
	// The HTTP resource path on which to fetch metrics from targets. E.g. /metrics
	// REQUIRED
	MetricsPath UpdateScrapeConfigPayloadGetMetricsPathAttributeType `json:"metricsPath" required:"true"`
	// List of metric relabel configurations
	MetricsRelabelConfigs UpdateScrapeConfigPayloadGetMetricsRelabelConfigsAttributeType `json:"metricsRelabelConfigs,omitempty"`
	// Optional http params `Additional Validators:` * should not contain more than 5 keys * each key and value should not have more than 200 characters
	Params UpdateScrapeConfigPayloadGetParamsAttributeType `json:"params,omitempty"`
	// Per-scrape limit on number of scraped samples that will be accepted. If more than this number of samples are present after metric relabeling the entire scrape will be treated as failed. The total limit depends on the service plan target limits * samples
	SampleLimit UpdateScrapeConfigPayloadGetSampleLimitAttributeType `json:"sampleLimit,omitempty"`
	// Configures the protocol scheme used for requests. https or http
	// REQUIRED
	Scheme UpdateScrapeConfigPayloadGetSchemeAttributeType `json:"scheme" required:"true"`
	// How frequently to scrape targets from this job. E.g. 5m `Additional Validators:` * must be a valid time format* must be >= 60s
	// REQUIRED
	ScrapeInterval UpdateScrapeConfigPayloadGetScrapeIntervalAttributeType `json:"scrapeInterval" required:"true"`
	// Per-scrape timeout when scraping this job. `Additional Validators:` * must be a valid time format* must be smaller than scrapeInterval
	// REQUIRED
	ScrapeTimeout UpdateScrapeConfigPayloadGetScrapeTimeoutAttributeType `json:"scrapeTimeout" required:"true"`
	// A list of scrape configurations.
	// REQUIRED
	StaticConfigs UpdateScrapeConfigPayloadGetStaticConfigsAttributeType `json:"staticConfigs" required:"true"`
	TlsConfig     UpdateScrapeConfigPayloadGetTlsConfigAttributeType     `json:"tlsConfig,omitempty"`
}

type _UpdateScrapeConfigPayload UpdateScrapeConfigPayload

// NewUpdateScrapeConfigPayload instantiates a new UpdateScrapeConfigPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateScrapeConfigPayload(metricsPath UpdateScrapeConfigPayloadGetMetricsPathArgType, scheme UpdateScrapeConfigPayloadGetSchemeArgType, scrapeInterval UpdateScrapeConfigPayloadGetScrapeIntervalArgType, scrapeTimeout UpdateScrapeConfigPayloadGetScrapeTimeoutArgType, staticConfigs UpdateScrapeConfigPayloadGetStaticConfigsArgType) *UpdateScrapeConfigPayload {
	this := UpdateScrapeConfigPayload{}
	setUpdateScrapeConfigPayloadGetMetricsPathAttributeType(&this.MetricsPath, metricsPath)
	setUpdateScrapeConfigPayloadGetSchemeAttributeType(&this.Scheme, scheme)
	setUpdateScrapeConfigPayloadGetScrapeIntervalAttributeType(&this.ScrapeInterval, scrapeInterval)
	setUpdateScrapeConfigPayloadGetScrapeTimeoutAttributeType(&this.ScrapeTimeout, scrapeTimeout)
	setUpdateScrapeConfigPayloadGetStaticConfigsAttributeType(&this.StaticConfigs, staticConfigs)
	return &this
}

// NewUpdateScrapeConfigPayloadWithDefaults instantiates a new UpdateScrapeConfigPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateScrapeConfigPayloadWithDefaults() *UpdateScrapeConfigPayload {
	this := UpdateScrapeConfigPayload{}
	var honorLabels bool = false
	this.HonorLabels = &honorLabels
	var honorTimeStamps bool = false
	this.HonorTimeStamps = &honorTimeStamps
	var metricsPath string = "/metrics"
	this.MetricsPath = &metricsPath
	return &this
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *UpdateScrapeConfigPayload) GetBasicAuth() (res UpdateScrapeConfigPayloadGetBasicAuthRetType) {
	res, _ = o.GetBasicAuthOk()
	return
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetBasicAuthOk() (ret UpdateScrapeConfigPayloadGetBasicAuthRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetBasicAuthAttributeTypeOk(o.BasicAuth)
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *UpdateScrapeConfigPayload) HasBasicAuth() bool {
	_, ok := o.GetBasicAuthOk()
	return ok
}

// SetBasicAuth gets a reference to the given CreateScrapeConfigPayloadBasicAuth and assigns it to the BasicAuth field.
func (o *UpdateScrapeConfigPayload) SetBasicAuth(v UpdateScrapeConfigPayloadGetBasicAuthRetType) {
	setUpdateScrapeConfigPayloadGetBasicAuthAttributeType(&o.BasicAuth, v)
}

// GetBearerToken returns the BearerToken field value if set, zero value otherwise.
func (o *UpdateScrapeConfigPayload) GetBearerToken() (res UpdateScrapeConfigPayloadGetBearerTokenRetType) {
	res, _ = o.GetBearerTokenOk()
	return
}

// GetBearerTokenOk returns a tuple with the BearerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetBearerTokenOk() (ret UpdateScrapeConfigPayloadGetBearerTokenRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetBearerTokenAttributeTypeOk(o.BearerToken)
}

// HasBearerToken returns a boolean if a field has been set.
func (o *UpdateScrapeConfigPayload) HasBearerToken() bool {
	_, ok := o.GetBearerTokenOk()
	return ok
}

// SetBearerToken gets a reference to the given string and assigns it to the BearerToken field.
func (o *UpdateScrapeConfigPayload) SetBearerToken(v UpdateScrapeConfigPayloadGetBearerTokenRetType) {
	setUpdateScrapeConfigPayloadGetBearerTokenAttributeType(&o.BearerToken, v)
}

// GetHonorLabels returns the HonorLabels field value if set, zero value otherwise.
func (o *UpdateScrapeConfigPayload) GetHonorLabels() (res UpdateScrapeConfigPayloadgetHonorLabelsRetType) {
	res, _ = o.GetHonorLabelsOk()
	return
}

// GetHonorLabelsOk returns a tuple with the HonorLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetHonorLabelsOk() (ret UpdateScrapeConfigPayloadgetHonorLabelsRetType, ok bool) {
	return getUpdateScrapeConfigPayloadgetHonorLabelsAttributeTypeOk(o.HonorLabels)
}

// HasHonorLabels returns a boolean if a field has been set.
func (o *UpdateScrapeConfigPayload) HasHonorLabels() bool {
	_, ok := o.GetHonorLabelsOk()
	return ok
}

// SetHonorLabels gets a reference to the given bool and assigns it to the HonorLabels field.
func (o *UpdateScrapeConfigPayload) SetHonorLabels(v UpdateScrapeConfigPayloadgetHonorLabelsRetType) {
	setUpdateScrapeConfigPayloadgetHonorLabelsAttributeType(&o.HonorLabels, v)
}

// GetHonorTimeStamps returns the HonorTimeStamps field value if set, zero value otherwise.
func (o *UpdateScrapeConfigPayload) GetHonorTimeStamps() (res UpdateScrapeConfigPayloadgetHonorTimeStampsRetType) {
	res, _ = o.GetHonorTimeStampsOk()
	return
}

// GetHonorTimeStampsOk returns a tuple with the HonorTimeStamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetHonorTimeStampsOk() (ret UpdateScrapeConfigPayloadgetHonorTimeStampsRetType, ok bool) {
	return getUpdateScrapeConfigPayloadgetHonorTimeStampsAttributeTypeOk(o.HonorTimeStamps)
}

// HasHonorTimeStamps returns a boolean if a field has been set.
func (o *UpdateScrapeConfigPayload) HasHonorTimeStamps() bool {
	_, ok := o.GetHonorTimeStampsOk()
	return ok
}

// SetHonorTimeStamps gets a reference to the given bool and assigns it to the HonorTimeStamps field.
func (o *UpdateScrapeConfigPayload) SetHonorTimeStamps(v UpdateScrapeConfigPayloadgetHonorTimeStampsRetType) {
	setUpdateScrapeConfigPayloadgetHonorTimeStampsAttributeType(&o.HonorTimeStamps, v)
}

// GetMetricsPath returns the MetricsPath field value
func (o *UpdateScrapeConfigPayload) GetMetricsPath() (ret UpdateScrapeConfigPayloadGetMetricsPathRetType) {
	ret, _ = o.GetMetricsPathOk()
	return ret
}

// GetMetricsPathOk returns a tuple with the MetricsPath field value
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetMetricsPathOk() (ret UpdateScrapeConfigPayloadGetMetricsPathRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetMetricsPathAttributeTypeOk(o.MetricsPath)
}

// SetMetricsPath sets field value
func (o *UpdateScrapeConfigPayload) SetMetricsPath(v UpdateScrapeConfigPayloadGetMetricsPathRetType) {
	setUpdateScrapeConfigPayloadGetMetricsPathAttributeType(&o.MetricsPath, v)
}

// GetMetricsRelabelConfigs returns the MetricsRelabelConfigs field value if set, zero value otherwise.
func (o *UpdateScrapeConfigPayload) GetMetricsRelabelConfigs() (res UpdateScrapeConfigPayloadGetMetricsRelabelConfigsRetType) {
	res, _ = o.GetMetricsRelabelConfigsOk()
	return
}

// GetMetricsRelabelConfigsOk returns a tuple with the MetricsRelabelConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetMetricsRelabelConfigsOk() (ret UpdateScrapeConfigPayloadGetMetricsRelabelConfigsRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetMetricsRelabelConfigsAttributeTypeOk(o.MetricsRelabelConfigs)
}

// HasMetricsRelabelConfigs returns a boolean if a field has been set.
func (o *UpdateScrapeConfigPayload) HasMetricsRelabelConfigs() bool {
	_, ok := o.GetMetricsRelabelConfigsOk()
	return ok
}

// SetMetricsRelabelConfigs gets a reference to the given []CreateScrapeConfigPayloadMetricsRelabelConfigsInner and assigns it to the MetricsRelabelConfigs field.
func (o *UpdateScrapeConfigPayload) SetMetricsRelabelConfigs(v UpdateScrapeConfigPayloadGetMetricsRelabelConfigsRetType) {
	setUpdateScrapeConfigPayloadGetMetricsRelabelConfigsAttributeType(&o.MetricsRelabelConfigs, v)
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *UpdateScrapeConfigPayload) GetParams() (res UpdateScrapeConfigPayloadGetParamsRetType) {
	res, _ = o.GetParamsOk()
	return
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetParamsOk() (ret UpdateScrapeConfigPayloadGetParamsRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetParamsAttributeTypeOk(o.Params)
}

// HasParams returns a boolean if a field has been set.
func (o *UpdateScrapeConfigPayload) HasParams() bool {
	_, ok := o.GetParamsOk()
	return ok
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *UpdateScrapeConfigPayload) SetParams(v UpdateScrapeConfigPayloadGetParamsRetType) {
	setUpdateScrapeConfigPayloadGetParamsAttributeType(&o.Params, v)
}

// GetSampleLimit returns the SampleLimit field value if set, zero value otherwise.
func (o *UpdateScrapeConfigPayload) GetSampleLimit() (res UpdateScrapeConfigPayloadGetSampleLimitRetType) {
	res, _ = o.GetSampleLimitOk()
	return
}

// GetSampleLimitOk returns a tuple with the SampleLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetSampleLimitOk() (ret UpdateScrapeConfigPayloadGetSampleLimitRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetSampleLimitAttributeTypeOk(o.SampleLimit)
}

// HasSampleLimit returns a boolean if a field has been set.
func (o *UpdateScrapeConfigPayload) HasSampleLimit() bool {
	_, ok := o.GetSampleLimitOk()
	return ok
}

// SetSampleLimit gets a reference to the given float64 and assigns it to the SampleLimit field.
func (o *UpdateScrapeConfigPayload) SetSampleLimit(v UpdateScrapeConfigPayloadGetSampleLimitRetType) {
	setUpdateScrapeConfigPayloadGetSampleLimitAttributeType(&o.SampleLimit, v)
}

// GetScheme returns the Scheme field value
func (o *UpdateScrapeConfigPayload) GetScheme() (ret UpdateScrapeConfigPayloadGetSchemeRetType) {
	ret, _ = o.GetSchemeOk()
	return ret
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetSchemeOk() (ret UpdateScrapeConfigPayloadGetSchemeRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetSchemeAttributeTypeOk(o.Scheme)
}

// SetScheme sets field value
func (o *UpdateScrapeConfigPayload) SetScheme(v UpdateScrapeConfigPayloadGetSchemeRetType) {
	setUpdateScrapeConfigPayloadGetSchemeAttributeType(&o.Scheme, v)
}

// GetScrapeInterval returns the ScrapeInterval field value
func (o *UpdateScrapeConfigPayload) GetScrapeInterval() (ret UpdateScrapeConfigPayloadGetScrapeIntervalRetType) {
	ret, _ = o.GetScrapeIntervalOk()
	return ret
}

// GetScrapeIntervalOk returns a tuple with the ScrapeInterval field value
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetScrapeIntervalOk() (ret UpdateScrapeConfigPayloadGetScrapeIntervalRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetScrapeIntervalAttributeTypeOk(o.ScrapeInterval)
}

// SetScrapeInterval sets field value
func (o *UpdateScrapeConfigPayload) SetScrapeInterval(v UpdateScrapeConfigPayloadGetScrapeIntervalRetType) {
	setUpdateScrapeConfigPayloadGetScrapeIntervalAttributeType(&o.ScrapeInterval, v)
}

// GetScrapeTimeout returns the ScrapeTimeout field value
func (o *UpdateScrapeConfigPayload) GetScrapeTimeout() (ret UpdateScrapeConfigPayloadGetScrapeTimeoutRetType) {
	ret, _ = o.GetScrapeTimeoutOk()
	return ret
}

// GetScrapeTimeoutOk returns a tuple with the ScrapeTimeout field value
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetScrapeTimeoutOk() (ret UpdateScrapeConfigPayloadGetScrapeTimeoutRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetScrapeTimeoutAttributeTypeOk(o.ScrapeTimeout)
}

// SetScrapeTimeout sets field value
func (o *UpdateScrapeConfigPayload) SetScrapeTimeout(v UpdateScrapeConfigPayloadGetScrapeTimeoutRetType) {
	setUpdateScrapeConfigPayloadGetScrapeTimeoutAttributeType(&o.ScrapeTimeout, v)
}

// GetStaticConfigs returns the StaticConfigs field value
func (o *UpdateScrapeConfigPayload) GetStaticConfigs() (ret UpdateScrapeConfigPayloadGetStaticConfigsRetType) {
	ret, _ = o.GetStaticConfigsOk()
	return ret
}

// GetStaticConfigsOk returns a tuple with the StaticConfigs field value
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetStaticConfigsOk() (ret UpdateScrapeConfigPayloadGetStaticConfigsRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetStaticConfigsAttributeTypeOk(o.StaticConfigs)
}

// SetStaticConfigs sets field value
func (o *UpdateScrapeConfigPayload) SetStaticConfigs(v UpdateScrapeConfigPayloadGetStaticConfigsRetType) {
	setUpdateScrapeConfigPayloadGetStaticConfigsAttributeType(&o.StaticConfigs, v)
}

// GetTlsConfig returns the TlsConfig field value if set, zero value otherwise.
func (o *UpdateScrapeConfigPayload) GetTlsConfig() (res UpdateScrapeConfigPayloadGetTlsConfigRetType) {
	res, _ = o.GetTlsConfigOk()
	return
}

// GetTlsConfigOk returns a tuple with the TlsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateScrapeConfigPayload) GetTlsConfigOk() (ret UpdateScrapeConfigPayloadGetTlsConfigRetType, ok bool) {
	return getUpdateScrapeConfigPayloadGetTlsConfigAttributeTypeOk(o.TlsConfig)
}

// HasTlsConfig returns a boolean if a field has been set.
func (o *UpdateScrapeConfigPayload) HasTlsConfig() bool {
	_, ok := o.GetTlsConfigOk()
	return ok
}

// SetTlsConfig gets a reference to the given CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig and assigns it to the TlsConfig field.
func (o *UpdateScrapeConfigPayload) SetTlsConfig(v UpdateScrapeConfigPayloadGetTlsConfigRetType) {
	setUpdateScrapeConfigPayloadGetTlsConfigAttributeType(&o.TlsConfig, v)
}

func (o UpdateScrapeConfigPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getUpdateScrapeConfigPayloadGetBasicAuthAttributeTypeOk(o.BasicAuth); ok {
		toSerialize["BasicAuth"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetBearerTokenAttributeTypeOk(o.BearerToken); ok {
		toSerialize["BearerToken"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadgetHonorLabelsAttributeTypeOk(o.HonorLabels); ok {
		toSerialize["HonorLabels"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadgetHonorTimeStampsAttributeTypeOk(o.HonorTimeStamps); ok {
		toSerialize["HonorTimeStamps"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetMetricsPathAttributeTypeOk(o.MetricsPath); ok {
		toSerialize["MetricsPath"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetMetricsRelabelConfigsAttributeTypeOk(o.MetricsRelabelConfigs); ok {
		toSerialize["MetricsRelabelConfigs"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetParamsAttributeTypeOk(o.Params); ok {
		toSerialize["Params"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetSampleLimitAttributeTypeOk(o.SampleLimit); ok {
		toSerialize["SampleLimit"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetSchemeAttributeTypeOk(o.Scheme); ok {
		toSerialize["Scheme"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetScrapeIntervalAttributeTypeOk(o.ScrapeInterval); ok {
		toSerialize["ScrapeInterval"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetScrapeTimeoutAttributeTypeOk(o.ScrapeTimeout); ok {
		toSerialize["ScrapeTimeout"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetStaticConfigsAttributeTypeOk(o.StaticConfigs); ok {
		toSerialize["StaticConfigs"] = val
	}
	if val, ok := getUpdateScrapeConfigPayloadGetTlsConfigAttributeTypeOk(o.TlsConfig); ok {
		toSerialize["TlsConfig"] = val
	}
	return toSerialize, nil
}

type NullableUpdateScrapeConfigPayload struct {
	value *UpdateScrapeConfigPayload
	isSet bool
}

func (v NullableUpdateScrapeConfigPayload) Get() *UpdateScrapeConfigPayload {
	return v.value
}

func (v *NullableUpdateScrapeConfigPayload) Set(val *UpdateScrapeConfigPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateScrapeConfigPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateScrapeConfigPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateScrapeConfigPayload(val *UpdateScrapeConfigPayload) *NullableUpdateScrapeConfigPayload {
	return &NullableUpdateScrapeConfigPayload{value: val, isSet: true}
}

func (v NullableUpdateScrapeConfigPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateScrapeConfigPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
