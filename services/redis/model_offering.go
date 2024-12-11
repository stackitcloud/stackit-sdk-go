/*
STACKIT Redis API

The STACKIT Redis API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redis

import (
	"encoding/json"
)

// checks if the Offering type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Offering{}

// Offering struct for Offering
type Offering struct {
	// REQUIRED
	Description *string `json:"description"`
	// REQUIRED
	DocumentationUrl *string `json:"documentationUrl"`
	// REQUIRED
	ImageUrl *string `json:"imageUrl"`
	// REQUIRED
	Latest    *bool   `json:"latest"`
	Lifecycle *string `json:"lifecycle,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Plans *[]Plan `json:"plans"`
	// REQUIRED
	QuotaCount *int64          `json:"quotaCount"`
	Schema     *InstanceSchema `json:"schema,omitempty"`
	// REQUIRED
	Version *string `json:"version"`
}

type _Offering Offering

// NewOffering instantiates a new Offering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffering(description *string, documentationUrl *string, imageUrl *string, latest *bool, name *string, plans *[]Plan, quotaCount *int64, version *string) *Offering {
	this := Offering{}
	this.Description = description
	this.DocumentationUrl = documentationUrl
	this.ImageUrl = imageUrl
	this.Latest = latest
	this.Name = name
	this.Plans = plans
	this.QuotaCount = quotaCount
	this.Version = version
	return &this
}

// NewOfferingWithDefaults instantiates a new Offering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferingWithDefaults() *Offering {
	this := Offering{}
	return &this
}

// GetDescription returns the Description field value
func (o *Offering) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Offering) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description, true
}

// SetDescription sets field value
func (o *Offering) SetDescription(v *string) {
	o.Description = v
}

// GetDocumentationUrl returns the DocumentationUrl field value
func (o *Offering) GetDocumentationUrl() *string {
	if o == nil || IsNil(o.DocumentationUrl) {
		var ret *string
		return ret
	}

	return o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value
// and a boolean to check if the value has been set.
func (o *Offering) GetDocumentationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// SetDocumentationUrl sets field value
func (o *Offering) SetDocumentationUrl(v *string) {
	o.DocumentationUrl = v
}

// GetImageUrl returns the ImageUrl field value
func (o *Offering) GetImageUrl() *string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret *string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *Offering) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *Offering) SetImageUrl(v *string) {
	o.ImageUrl = v
}

// GetLatest returns the Latest field value
func (o *Offering) GetLatest() *bool {
	if o == nil || IsNil(o.Latest) {
		var ret *bool
		return ret
	}

	return o.Latest
}

// GetLatestOk returns a tuple with the Latest field value
// and a boolean to check if the value has been set.
func (o *Offering) GetLatestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latest, true
}

// SetLatest sets field value
func (o *Offering) SetLatest(v *bool) {
	o.Latest = v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *Offering) GetLifecycle() *string {
	if o == nil || IsNil(o.Lifecycle) {
		var ret *string
		return ret
	}
	return o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offering) GetLifecycleOk() (*string, bool) {
	if o == nil || IsNil(o.Lifecycle) {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *Offering) HasLifecycle() bool {
	if o != nil && !IsNil(o.Lifecycle) && !IsNil(o.Lifecycle) {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *Offering) SetLifecycle(v *string) {
	o.Lifecycle = v
}

// GetName returns the Name field value
func (o *Offering) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Offering) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Offering) SetName(v *string) {
	o.Name = v
}

// GetPlans returns the Plans field value
func (o *Offering) GetPlans() *[]Plan {
	if o == nil || IsNil(o.Plans) {
		var ret *[]Plan
		return ret
	}

	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value
// and a boolean to check if the value has been set.
func (o *Offering) GetPlansOk() (*[]Plan, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plans, true
}

// SetPlans sets field value
func (o *Offering) SetPlans(v *[]Plan) {
	o.Plans = v
}

// GetQuotaCount returns the QuotaCount field value
func (o *Offering) GetQuotaCount() *int64 {
	if o == nil || IsNil(o.QuotaCount) {
		var ret *int64
		return ret
	}

	return o.QuotaCount
}

// GetQuotaCountOk returns a tuple with the QuotaCount field value
// and a boolean to check if the value has been set.
func (o *Offering) GetQuotaCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuotaCount, true
}

// SetQuotaCount sets field value
func (o *Offering) SetQuotaCount(v *int64) {
	o.QuotaCount = v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *Offering) GetSchema() *InstanceSchema {
	if o == nil || IsNil(o.Schema) {
		var ret *InstanceSchema
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offering) GetSchemaOk() (*InstanceSchema, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *Offering) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given InstanceSchema and assigns it to the Schema field.
func (o *Offering) SetSchema(v *InstanceSchema) {
	o.Schema = v
}

// GetVersion returns the Version field value
func (o *Offering) GetVersion() *string {
	if o == nil || IsNil(o.Version) {
		var ret *string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Offering) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version, true
}

// SetVersion sets field value
func (o *Offering) SetVersion(v *string) {
	o.Version = v
}

func (o Offering) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["documentationUrl"] = o.DocumentationUrl
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["latest"] = o.Latest
	if !IsNil(o.Lifecycle) {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	toSerialize["name"] = o.Name
	toSerialize["plans"] = o.Plans
	toSerialize["quotaCount"] = o.QuotaCount
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableOffering struct {
	value *Offering
	isSet bool
}

func (v NullableOffering) Get() *Offering {
	return v.value
}

func (v *NullableOffering) Set(val *Offering) {
	v.value = val
	v.isSet = true
}

func (v NullableOffering) IsSet() bool {
	return v.isSet
}

func (v *NullableOffering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffering(val *Offering) *NullableOffering {
	return &NullableOffering{value: val, isSet: true}
}

func (v NullableOffering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
