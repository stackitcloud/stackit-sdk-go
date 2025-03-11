/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alb

import (
	"encoding/json"
)

// checks if the CredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsResponse{}

// CredentialsResponse struct for CredentialsResponse
type CredentialsResponse struct {
	// The credentials reference can be used for observability of the Application Load Balancer.
	CredentialsRef *string `json:"credentialsRef,omitempty"`
	// Credential name
	DisplayName *string `json:"displayName,omitempty"`
	// Region of the Credential
	Region *string `json:"region,omitempty"`
	// The username used for the ARGUS instance
	Username *string `json:"username,omitempty"`
}

// NewCredentialsResponse instantiates a new CredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsResponse() *CredentialsResponse {
	this := CredentialsResponse{}
	return &this
}

// NewCredentialsResponseWithDefaults instantiates a new CredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsResponseWithDefaults() *CredentialsResponse {
	this := CredentialsResponse{}
	return &this
}

// GetCredentialsRef returns the CredentialsRef field value if set, zero value otherwise.
func (o *CredentialsResponse) GetCredentialsRef() *string {
	if o == nil || IsNil(o.CredentialsRef) {
		var ret *string
		return ret
	}
	return o.CredentialsRef
}

// GetCredentialsRefOk returns a tuple with the CredentialsRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsResponse) GetCredentialsRefOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialsRef) {
		return nil, false
	}
	return o.CredentialsRef, true
}

// HasCredentialsRef returns a boolean if a field has been set.
func (o *CredentialsResponse) HasCredentialsRef() bool {
	if o != nil && !IsNil(o.CredentialsRef) {
		return true
	}

	return false
}

// SetCredentialsRef gets a reference to the given string and assigns it to the CredentialsRef field.
func (o *CredentialsResponse) SetCredentialsRef(v *string) {
	o.CredentialsRef = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CredentialsResponse) GetDisplayName() *string {
	if o == nil || IsNil(o.DisplayName) {
		var ret *string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CredentialsResponse) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CredentialsResponse) SetDisplayName(v *string) {
	o.DisplayName = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CredentialsResponse) GetRegion() *string {
	if o == nil || IsNil(o.Region) {
		var ret *string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CredentialsResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CredentialsResponse) SetRegion(v *string) {
	o.Region = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CredentialsResponse) GetUsername() *string {
	if o == nil || IsNil(o.Username) {
		var ret *string
		return ret
	}
	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsResponse) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CredentialsResponse) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CredentialsResponse) SetUsername(v *string) {
	o.Username = v
}

func (o CredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredentialsRef) {
		toSerialize["credentialsRef"] = o.CredentialsRef
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableCredentialsResponse struct {
	value *CredentialsResponse
	isSet bool
}

func (v NullableCredentialsResponse) Get() *CredentialsResponse {
	return v.value
}

func (v *NullableCredentialsResponse) Set(val *CredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsResponse(val *CredentialsResponse) *NullableCredentialsResponse {
	return &NullableCredentialsResponse{value: val, isSet: true}
}

func (v NullableCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
