/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

import (
	"encoding/json"
)

// checks if the Plan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Plan{}

// Plan struct for Plan
type Plan struct {
	// REQUIRED
	AlertMatchers *int64 `json:"alertMatchers"`
	// REQUIRED
	AlertReceivers *int64 `json:"alertReceivers"`
	// REQUIRED
	AlertRules *int64   `json:"alertRules"`
	Amount     *float64 `json:"amount,omitempty"`
	// REQUIRED
	BucketSize  *int64  `json:"bucketSize"`
	Description *string `json:"description,omitempty"`
	// REQUIRED
	GrafanaGlobalDashboards *int64 `json:"grafanaGlobalDashboards"`
	// REQUIRED
	GrafanaGlobalOrgs *int64 `json:"grafanaGlobalOrgs"`
	// REQUIRED
	GrafanaGlobalSessions *int64 `json:"grafanaGlobalSessions"`
	// REQUIRED
	GrafanaGlobalUsers *int64 `json:"grafanaGlobalUsers"`
	// REQUIRED
	Id       *string `json:"id"`
	IsFree   *bool   `json:"isFree,omitempty"`
	IsPublic *bool   `json:"isPublic,omitempty"`
	// REQUIRED
	LogsAlert *int64 `json:"logsAlert"`
	// REQUIRED
	LogsStorage *int64  `json:"logsStorage"`
	Name        *string `json:"name,omitempty"`
	// REQUIRED
	PlanId *string `json:"planId"`
	// REQUIRED
	SamplesPerScrape *int64  `json:"samplesPerScrape"`
	Schema           *string `json:"schema,omitempty"`
	// REQUIRED
	TargetNumber *int64 `json:"targetNumber"`
	// REQUIRED
	TotalMetricSamples *int64 `json:"totalMetricSamples"`
	// REQUIRED
	TracesStorage *int64 `json:"tracesStorage"`
}

type _Plan Plan

// NewPlan instantiates a new Plan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlan(alertMatchers *int64, alertReceivers *int64, alertRules *int64, bucketSize *int64, grafanaGlobalDashboards *int64, grafanaGlobalOrgs *int64, grafanaGlobalSessions *int64, grafanaGlobalUsers *int64, id *string, logsAlert *int64, logsStorage *int64, planId *string, samplesPerScrape *int64, targetNumber *int64, totalMetricSamples *int64, tracesStorage *int64) *Plan {
	this := Plan{}
	this.AlertMatchers = alertMatchers
	this.AlertReceivers = alertReceivers
	this.AlertRules = alertRules
	this.BucketSize = bucketSize
	this.GrafanaGlobalDashboards = grafanaGlobalDashboards
	this.GrafanaGlobalOrgs = grafanaGlobalOrgs
	this.GrafanaGlobalSessions = grafanaGlobalSessions
	this.GrafanaGlobalUsers = grafanaGlobalUsers
	this.Id = id
	var isFree bool = false
	this.IsFree = &isFree
	var isPublic bool = true
	this.IsPublic = &isPublic
	this.LogsAlert = logsAlert
	this.LogsStorage = logsStorage
	this.PlanId = planId
	this.SamplesPerScrape = samplesPerScrape
	var schema string = ""
	this.Schema = &schema
	this.TargetNumber = targetNumber
	this.TotalMetricSamples = totalMetricSamples
	this.TracesStorage = tracesStorage
	return &this
}

// NewPlanWithDefaults instantiates a new Plan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanWithDefaults() *Plan {
	this := Plan{}
	var isFree bool = false
	this.IsFree = &isFree
	var isPublic bool = true
	this.IsPublic = &isPublic
	var schema string = ""
	this.Schema = &schema
	return &this
}

// GetAlertMatchers returns the AlertMatchers field value
func (o *Plan) GetAlertMatchers() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.AlertMatchers
}

// GetAlertMatchersOk returns a tuple with the AlertMatchers field value
// and a boolean to check if the value has been set.
func (o *Plan) GetAlertMatchersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertMatchers, true
}

// SetAlertMatchers sets field value
func (o *Plan) SetAlertMatchers(v *int64) {
	o.AlertMatchers = v
}

// GetAlertReceivers returns the AlertReceivers field value
func (o *Plan) GetAlertReceivers() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.AlertReceivers
}

// GetAlertReceiversOk returns a tuple with the AlertReceivers field value
// and a boolean to check if the value has been set.
func (o *Plan) GetAlertReceiversOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertReceivers, true
}

// SetAlertReceivers sets field value
func (o *Plan) SetAlertReceivers(v *int64) {
	o.AlertReceivers = v
}

// GetAlertRules returns the AlertRules field value
func (o *Plan) GetAlertRules() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.AlertRules
}

// GetAlertRulesOk returns a tuple with the AlertRules field value
// and a boolean to check if the value has been set.
func (o *Plan) GetAlertRulesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertRules, true
}

// SetAlertRules sets field value
func (o *Plan) SetAlertRules(v *int64) {
	o.AlertRules = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Plan) GetAmount() *float64 {
	if o == nil || IsNil(o.Amount) {
		var ret *float64
		return ret
	}
	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Plan) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *Plan) SetAmount(v *float64) {
	o.Amount = v
}

// GetBucketSize returns the BucketSize field value
func (o *Plan) GetBucketSize() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.BucketSize
}

// GetBucketSizeOk returns a tuple with the BucketSize field value
// and a boolean to check if the value has been set.
func (o *Plan) GetBucketSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BucketSize, true
}

// SetBucketSize sets field value
func (o *Plan) SetBucketSize(v *int64) {
	o.BucketSize = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Plan) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Plan) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Plan) SetDescription(v *string) {
	o.Description = v
}

// GetGrafanaGlobalDashboards returns the GrafanaGlobalDashboards field value
func (o *Plan) GetGrafanaGlobalDashboards() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.GrafanaGlobalDashboards
}

// GetGrafanaGlobalDashboardsOk returns a tuple with the GrafanaGlobalDashboards field value
// and a boolean to check if the value has been set.
func (o *Plan) GetGrafanaGlobalDashboardsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrafanaGlobalDashboards, true
}

// SetGrafanaGlobalDashboards sets field value
func (o *Plan) SetGrafanaGlobalDashboards(v *int64) {
	o.GrafanaGlobalDashboards = v
}

// GetGrafanaGlobalOrgs returns the GrafanaGlobalOrgs field value
func (o *Plan) GetGrafanaGlobalOrgs() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.GrafanaGlobalOrgs
}

// GetGrafanaGlobalOrgsOk returns a tuple with the GrafanaGlobalOrgs field value
// and a boolean to check if the value has been set.
func (o *Plan) GetGrafanaGlobalOrgsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrafanaGlobalOrgs, true
}

// SetGrafanaGlobalOrgs sets field value
func (o *Plan) SetGrafanaGlobalOrgs(v *int64) {
	o.GrafanaGlobalOrgs = v
}

// GetGrafanaGlobalSessions returns the GrafanaGlobalSessions field value
func (o *Plan) GetGrafanaGlobalSessions() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.GrafanaGlobalSessions
}

// GetGrafanaGlobalSessionsOk returns a tuple with the GrafanaGlobalSessions field value
// and a boolean to check if the value has been set.
func (o *Plan) GetGrafanaGlobalSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrafanaGlobalSessions, true
}

// SetGrafanaGlobalSessions sets field value
func (o *Plan) SetGrafanaGlobalSessions(v *int64) {
	o.GrafanaGlobalSessions = v
}

// GetGrafanaGlobalUsers returns the GrafanaGlobalUsers field value
func (o *Plan) GetGrafanaGlobalUsers() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.GrafanaGlobalUsers
}

// GetGrafanaGlobalUsersOk returns a tuple with the GrafanaGlobalUsers field value
// and a boolean to check if the value has been set.
func (o *Plan) GetGrafanaGlobalUsersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrafanaGlobalUsers, true
}

// SetGrafanaGlobalUsers sets field value
func (o *Plan) SetGrafanaGlobalUsers(v *int64) {
	o.GrafanaGlobalUsers = v
}

// GetId returns the Id field value
func (o *Plan) GetId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Plan) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *Plan) SetId(v *string) {
	o.Id = v
}

// GetIsFree returns the IsFree field value if set, zero value otherwise.
func (o *Plan) GetIsFree() *bool {
	if o == nil || IsNil(o.IsFree) {
		var ret *bool
		return ret
	}
	return o.IsFree
}

// GetIsFreeOk returns a tuple with the IsFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetIsFreeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFree) {
		return nil, false
	}
	return o.IsFree, true
}

// HasIsFree returns a boolean if a field has been set.
func (o *Plan) HasIsFree() bool {
	if o != nil && !IsNil(o.IsFree) {
		return true
	}

	return false
}

// SetIsFree gets a reference to the given bool and assigns it to the IsFree field.
func (o *Plan) SetIsFree(v *bool) {
	o.IsFree = v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *Plan) GetIsPublic() *bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret *bool
		return ret
	}
	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *Plan) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *Plan) SetIsPublic(v *bool) {
	o.IsPublic = v
}

// GetLogsAlert returns the LogsAlert field value
func (o *Plan) GetLogsAlert() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.LogsAlert
}

// GetLogsAlertOk returns a tuple with the LogsAlert field value
// and a boolean to check if the value has been set.
func (o *Plan) GetLogsAlertOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsAlert, true
}

// SetLogsAlert sets field value
func (o *Plan) SetLogsAlert(v *int64) {
	o.LogsAlert = v
}

// GetLogsStorage returns the LogsStorage field value
func (o *Plan) GetLogsStorage() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.LogsStorage
}

// GetLogsStorageOk returns a tuple with the LogsStorage field value
// and a boolean to check if the value has been set.
func (o *Plan) GetLogsStorageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsStorage, true
}

// SetLogsStorage sets field value
func (o *Plan) SetLogsStorage(v *int64) {
	o.LogsStorage = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Plan) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Plan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Plan) SetName(v *string) {
	o.Name = v
}

// GetPlanId returns the PlanId field value
func (o *Plan) GetPlanId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *Plan) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId, true
}

// SetPlanId sets field value
func (o *Plan) SetPlanId(v *string) {
	o.PlanId = v
}

// GetSamplesPerScrape returns the SamplesPerScrape field value
func (o *Plan) GetSamplesPerScrape() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.SamplesPerScrape
}

// GetSamplesPerScrapeOk returns a tuple with the SamplesPerScrape field value
// and a boolean to check if the value has been set.
func (o *Plan) GetSamplesPerScrapeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SamplesPerScrape, true
}

// SetSamplesPerScrape sets field value
func (o *Plan) SetSamplesPerScrape(v *int64) {
	o.SamplesPerScrape = v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *Plan) GetSchema() *string {
	if o == nil || IsNil(o.Schema) {
		var ret *string
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *Plan) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *Plan) SetSchema(v *string) {
	o.Schema = v
}

// GetTargetNumber returns the TargetNumber field value
func (o *Plan) GetTargetNumber() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.TargetNumber
}

// GetTargetNumberOk returns a tuple with the TargetNumber field value
// and a boolean to check if the value has been set.
func (o *Plan) GetTargetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetNumber, true
}

// SetTargetNumber sets field value
func (o *Plan) SetTargetNumber(v *int64) {
	o.TargetNumber = v
}

// GetTotalMetricSamples returns the TotalMetricSamples field value
func (o *Plan) GetTotalMetricSamples() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.TotalMetricSamples
}

// GetTotalMetricSamplesOk returns a tuple with the TotalMetricSamples field value
// and a boolean to check if the value has been set.
func (o *Plan) GetTotalMetricSamplesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalMetricSamples, true
}

// SetTotalMetricSamples sets field value
func (o *Plan) SetTotalMetricSamples(v *int64) {
	o.TotalMetricSamples = v
}

// GetTracesStorage returns the TracesStorage field value
func (o *Plan) GetTracesStorage() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.TracesStorage
}

// GetTracesStorageOk returns a tuple with the TracesStorage field value
// and a boolean to check if the value has been set.
func (o *Plan) GetTracesStorageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TracesStorage, true
}

// SetTracesStorage sets field value
func (o *Plan) SetTracesStorage(v *int64) {
	o.TracesStorage = v
}

func (o Plan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alertMatchers"] = o.AlertMatchers
	toSerialize["alertReceivers"] = o.AlertReceivers
	toSerialize["alertRules"] = o.AlertRules
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["bucketSize"] = o.BucketSize
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["grafanaGlobalDashboards"] = o.GrafanaGlobalDashboards
	toSerialize["grafanaGlobalOrgs"] = o.GrafanaGlobalOrgs
	toSerialize["grafanaGlobalSessions"] = o.GrafanaGlobalSessions
	toSerialize["grafanaGlobalUsers"] = o.GrafanaGlobalUsers
	toSerialize["id"] = o.Id
	if !IsNil(o.IsFree) {
		toSerialize["isFree"] = o.IsFree
	}
	if !IsNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	toSerialize["logsAlert"] = o.LogsAlert
	toSerialize["logsStorage"] = o.LogsStorage
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["planId"] = o.PlanId
	toSerialize["samplesPerScrape"] = o.SamplesPerScrape
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	toSerialize["targetNumber"] = o.TargetNumber
	toSerialize["totalMetricSamples"] = o.TotalMetricSamples
	toSerialize["tracesStorage"] = o.TracesStorage
	return toSerialize, nil
}

type NullablePlan struct {
	value *Plan
	isSet bool
}

func (v NullablePlan) Get() *Plan {
	return v.value
}

func (v *NullablePlan) Set(val *Plan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlan(val *Plan) *NullablePlan {
	return &NullablePlan{value: val, isSet: true}
}

func (v NullablePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
