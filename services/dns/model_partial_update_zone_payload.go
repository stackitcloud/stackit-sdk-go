/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the PartialUpdateZonePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialUpdateZonePayload{}

// PartialUpdateZonePayload struct for PartialUpdateZonePayload
type PartialUpdateZonePayload struct {
	// access control list
	Acl *string `json:"acl,omitempty"`
	// contact e-mail for the zone
	ContactEmail *string `json:"contactEmail,omitempty"`
	// default time to live
	DefaultTTL *int64 `json:"defaultTTL,omitempty"`
	// description of the zone
	Description *string `json:"description,omitempty"`
	// expire time
	ExpireTime *int64 `json:"expireTime,omitempty"`
	// user given name
	Name *string `json:"name,omitempty"`
	// negative caching
	NegativeCache *int64 `json:"negativeCache,omitempty"`
	// primary name server for secondary zone
	Primaries *[]string `json:"primaries,omitempty"`
	// refresh time
	RefreshTime *int64 `json:"refreshTime,omitempty"`
	// retry time
	RetryTime *int64 `json:"retryTime,omitempty"`
}

// NewPartialUpdateZonePayload instantiates a new PartialUpdateZonePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialUpdateZonePayload() *PartialUpdateZonePayload {
	this := PartialUpdateZonePayload{}
	var acl string = "0.0.0.0/0,::/0"
	this.Acl = &acl
	var contactEmail string = "hostmaster@stackit.cloud"
	this.ContactEmail = &contactEmail
	return &this
}

// NewPartialUpdateZonePayloadWithDefaults instantiates a new PartialUpdateZonePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialUpdateZonePayloadWithDefaults() *PartialUpdateZonePayload {
	this := PartialUpdateZonePayload{}
	var acl string = "0.0.0.0/0,::/0"
	this.Acl = &acl
	var contactEmail string = "hostmaster@stackit.cloud"
	this.ContactEmail = &contactEmail
	return &this
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetAcl() *string {
	if o == nil || IsNil(o.Acl) {
		var ret *string
		return ret
	}
	return o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetAclOk() (*string, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *PartialUpdateZonePayload) SetAcl(v *string) {
	o.Acl = v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetContactEmail() *string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret *string
		return ret
	}
	return o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *PartialUpdateZonePayload) SetContactEmail(v *string) {
	o.ContactEmail = v
}

// GetDefaultTTL returns the DefaultTTL field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetDefaultTTL() *int64 {
	if o == nil || IsNil(o.DefaultTTL) {
		var ret *int64
		return ret
	}
	return o.DefaultTTL
}

// GetDefaultTTLOk returns a tuple with the DefaultTTL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetDefaultTTLOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultTTL) {
		return nil, false
	}
	return o.DefaultTTL, true
}

// HasDefaultTTL returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasDefaultTTL() bool {
	if o != nil && !IsNil(o.DefaultTTL) && !IsNil(o.DefaultTTL) {
		return true
	}

	return false
}

// SetDefaultTTL gets a reference to the given int64 and assigns it to the DefaultTTL field.
func (o *PartialUpdateZonePayload) SetDefaultTTL(v *int64) {
	o.DefaultTTL = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PartialUpdateZonePayload) SetDescription(v *string) {
	o.Description = v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetExpireTime() *int64 {
	if o == nil || IsNil(o.ExpireTime) {
		var ret *int64
		return ret
	}
	return o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetExpireTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given int64 and assigns it to the ExpireTime field.
func (o *PartialUpdateZonePayload) SetExpireTime(v *int64) {
	o.ExpireTime = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasName() bool {
	if o != nil && !IsNil(o.Name) && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartialUpdateZonePayload) SetName(v *string) {
	o.Name = v
}

// GetNegativeCache returns the NegativeCache field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetNegativeCache() *int64 {
	if o == nil || IsNil(o.NegativeCache) {
		var ret *int64
		return ret
	}
	return o.NegativeCache
}

// GetNegativeCacheOk returns a tuple with the NegativeCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetNegativeCacheOk() (*int64, bool) {
	if o == nil || IsNil(o.NegativeCache) {
		return nil, false
	}
	return o.NegativeCache, true
}

// HasNegativeCache returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasNegativeCache() bool {
	if o != nil && !IsNil(o.NegativeCache) && !IsNil(o.NegativeCache) {
		return true
	}

	return false
}

// SetNegativeCache gets a reference to the given int64 and assigns it to the NegativeCache field.
func (o *PartialUpdateZonePayload) SetNegativeCache(v *int64) {
	o.NegativeCache = v
}

// GetPrimaries returns the Primaries field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetPrimaries() *[]string {
	if o == nil || IsNil(o.Primaries) {
		var ret *[]string
		return ret
	}
	return o.Primaries
}

// GetPrimariesOk returns a tuple with the Primaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetPrimariesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Primaries) {
		return nil, false
	}
	return o.Primaries, true
}

// HasPrimaries returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasPrimaries() bool {
	if o != nil && !IsNil(o.Primaries) && !IsNil(o.Primaries) {
		return true
	}

	return false
}

// SetPrimaries gets a reference to the given []string and assigns it to the Primaries field.
func (o *PartialUpdateZonePayload) SetPrimaries(v *[]string) {
	o.Primaries = v
}

// GetRefreshTime returns the RefreshTime field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetRefreshTime() *int64 {
	if o == nil || IsNil(o.RefreshTime) {
		var ret *int64
		return ret
	}
	return o.RefreshTime
}

// GetRefreshTimeOk returns a tuple with the RefreshTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetRefreshTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshTime) {
		return nil, false
	}
	return o.RefreshTime, true
}

// HasRefreshTime returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasRefreshTime() bool {
	if o != nil && !IsNil(o.RefreshTime) && !IsNil(o.RefreshTime) {
		return true
	}

	return false
}

// SetRefreshTime gets a reference to the given int64 and assigns it to the RefreshTime field.
func (o *PartialUpdateZonePayload) SetRefreshTime(v *int64) {
	o.RefreshTime = v
}

// GetRetryTime returns the RetryTime field value if set, zero value otherwise.
func (o *PartialUpdateZonePayload) GetRetryTime() *int64 {
	if o == nil || IsNil(o.RetryTime) {
		var ret *int64
		return ret
	}
	return o.RetryTime
}

// GetRetryTimeOk returns a tuple with the RetryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateZonePayload) GetRetryTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.RetryTime) {
		return nil, false
	}
	return o.RetryTime, true
}

// HasRetryTime returns a boolean if a field has been set.
func (o *PartialUpdateZonePayload) HasRetryTime() bool {
	if o != nil && !IsNil(o.RetryTime) && !IsNil(o.RetryTime) {
		return true
	}

	return false
}

// SetRetryTime gets a reference to the given int64 and assigns it to the RetryTime field.
func (o *PartialUpdateZonePayload) SetRetryTime(v *int64) {
	o.RetryTime = v
}

func (o PartialUpdateZonePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acl) {
		toSerialize["acl"] = o.Acl
	}
	if !IsNil(o.ContactEmail) {
		toSerialize["contactEmail"] = o.ContactEmail
	}
	if !IsNil(o.DefaultTTL) {
		toSerialize["defaultTTL"] = o.DefaultTTL
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpireTime) {
		toSerialize["expireTime"] = o.ExpireTime
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NegativeCache) {
		toSerialize["negativeCache"] = o.NegativeCache
	}
	if !IsNil(o.Primaries) {
		toSerialize["primaries"] = o.Primaries
	}
	if !IsNil(o.RefreshTime) {
		toSerialize["refreshTime"] = o.RefreshTime
	}
	if !IsNil(o.RetryTime) {
		toSerialize["retryTime"] = o.RetryTime
	}
	return toSerialize, nil
}

type NullablePartialUpdateZonePayload struct {
	value *PartialUpdateZonePayload
	isSet bool
}

func (v NullablePartialUpdateZonePayload) Get() *PartialUpdateZonePayload {
	return v.value
}

func (v *NullablePartialUpdateZonePayload) Set(val *PartialUpdateZonePayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialUpdateZonePayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialUpdateZonePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialUpdateZonePayload(val *PartialUpdateZonePayload) *NullablePartialUpdateZonePayload {
	return &NullablePartialUpdateZonePayload{value: val, isSet: true}
}

func (v NullablePartialUpdateZonePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialUpdateZonePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
