/*
CDN API

API used to create and manage your CDN distributions.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the CustomDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDomain{}

// CustomDomain Definition of a custom domain
type CustomDomain struct {
	// This object is present if the custom domain has errors.
	Errors *[]StatusError `json:"errors,omitempty"`
	// The domain. Can be used as input for the GetCustomDomain endpoint
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Status *DomainStatus `json:"status"`
}

type _CustomDomain CustomDomain

// NewCustomDomain instantiates a new CustomDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomain(name *string, status *DomainStatus) *CustomDomain {
	this := CustomDomain{}
	this.Name = name
	this.Status = status
	return &this
}

// NewCustomDomainWithDefaults instantiates a new CustomDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainWithDefaults() *CustomDomain {
	this := CustomDomain{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CustomDomain) GetErrors() *[]StatusError {
	if o == nil || IsNil(o.Errors) {
		var ret *[]StatusError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetErrorsOk() (*[]StatusError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CustomDomain) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []StatusError and assigns it to the Errors field.
func (o *CustomDomain) SetErrors(v *[]StatusError) {
	o.Errors = v
}

// GetName returns the Name field value
func (o *CustomDomain) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CustomDomain) SetName(v *string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *CustomDomain) GetStatus() *DomainStatus {
	if o == nil || IsNil(o.Status) {
		var ret *DomainStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetStatusOk() (*DomainStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *CustomDomain) SetStatus(v *DomainStatus) {
	o.Status = v
}

func (o CustomDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableCustomDomain struct {
	value *CustomDomain
	isSet bool
}

func (v NullableCustomDomain) Get() *CustomDomain {
	return v.value
}

func (v *NullableCustomDomain) Set(val *CustomDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomain(val *CustomDomain) *NullableCustomDomain {
	return &NullableCustomDomain{value: val, isSet: true}
}

func (v NullableCustomDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
