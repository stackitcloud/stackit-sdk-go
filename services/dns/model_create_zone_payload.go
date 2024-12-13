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

// checks if the CreateZonePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateZonePayload{}

// CreateZonePayload Zone to create.
type CreateZonePayload struct {
	// access control list
	Acl *string `json:"acl,omitempty"`
	// contact e-mail for the zone
	ContactEmail *string `json:"contactEmail,omitempty"`
	// default time to live
	DefaultTTL *int64 `json:"defaultTTL,omitempty"`
	// description of the zone
	Description *string `json:"description,omitempty"`
	// zone name
	// REQUIRED
	DnsName *string `json:"dnsName"`
	// expire time
	ExpireTime *int64 `json:"expireTime,omitempty"`
	// if the zone is a reverse zone or not
	IsReverseZone *bool `json:"isReverseZone,omitempty"`
	// user given name
	// REQUIRED
	Name *string `json:"name"`
	// negative caching
	NegativeCache *int64 `json:"negativeCache,omitempty"`
	// primary name server for secondary zone
	Primaries *[]string `json:"primaries,omitempty"`
	// refresh time
	RefreshTime *int64 `json:"refreshTime,omitempty"`
	// retry time
	RetryTime *int64 `json:"retryTime,omitempty"`
	// zone type
	Type *string `json:"type,omitempty"`
}

type _CreateZonePayload CreateZonePayload

// NewCreateZonePayload instantiates a new CreateZonePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateZonePayload(dnsName *string, name *string) *CreateZonePayload {
	this := CreateZonePayload{}
	var acl string = "0.0.0.0/0,::/0"
	this.Acl = &acl
	var contactEmail string = "hostmaster@stackit.cloud"
	this.ContactEmail = &contactEmail
	this.DnsName = dnsName
	var isReverseZone bool = false
	this.IsReverseZone = &isReverseZone
	this.Name = name
	var type_ string = "primary"
	this.Type = &type_
	return &this
}

// NewCreateZonePayloadWithDefaults instantiates a new CreateZonePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateZonePayloadWithDefaults() *CreateZonePayload {
	this := CreateZonePayload{}
	var acl string = "0.0.0.0/0,::/0"
	this.Acl = &acl
	var contactEmail string = "hostmaster@stackit.cloud"
	this.ContactEmail = &contactEmail
	var isReverseZone bool = false
	this.IsReverseZone = &isReverseZone
	var type_ string = "primary"
	this.Type = &type_
	return &this
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *CreateZonePayload) GetAcl() *string {
	if o == nil || IsNil(o.Acl) {
		var ret *string
		return ret
	}
	return o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetAclOk() (*string, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *CreateZonePayload) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *CreateZonePayload) SetAcl(v *string) {
	o.Acl = v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *CreateZonePayload) GetContactEmail() *string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret *string
		return ret
	}
	return o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *CreateZonePayload) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *CreateZonePayload) SetContactEmail(v *string) {
	o.ContactEmail = v
}

// GetDefaultTTL returns the DefaultTTL field value if set, zero value otherwise.
func (o *CreateZonePayload) GetDefaultTTL() *int64 {
	if o == nil || IsNil(o.DefaultTTL) {
		var ret *int64
		return ret
	}
	return o.DefaultTTL
}

// GetDefaultTTLOk returns a tuple with the DefaultTTL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetDefaultTTLOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultTTL) {
		return nil, false
	}
	return o.DefaultTTL, true
}

// HasDefaultTTL returns a boolean if a field has been set.
func (o *CreateZonePayload) HasDefaultTTL() bool {
	if o != nil && !IsNil(o.DefaultTTL) {
		return true
	}

	return false
}

// SetDefaultTTL gets a reference to the given int64 and assigns it to the DefaultTTL field.
func (o *CreateZonePayload) SetDefaultTTL(v *int64) {
	o.DefaultTTL = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateZonePayload) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateZonePayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateZonePayload) SetDescription(v *string) {
	o.Description = v
}

// GetDnsName returns the DnsName field value
func (o *CreateZonePayload) GetDnsName() *string {
	if o == nil || IsNil(o.DnsName) {
		var ret *string
		return ret
	}

	return o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetDnsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsName, true
}

// SetDnsName sets field value
func (o *CreateZonePayload) SetDnsName(v *string) {
	o.DnsName = v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *CreateZonePayload) GetExpireTime() *int64 {
	if o == nil || IsNil(o.ExpireTime) {
		var ret *int64
		return ret
	}
	return o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetExpireTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *CreateZonePayload) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given int64 and assigns it to the ExpireTime field.
func (o *CreateZonePayload) SetExpireTime(v *int64) {
	o.ExpireTime = v
}

// GetIsReverseZone returns the IsReverseZone field value if set, zero value otherwise.
func (o *CreateZonePayload) GetIsReverseZone() *bool {
	if o == nil || IsNil(o.IsReverseZone) {
		var ret *bool
		return ret
	}
	return o.IsReverseZone
}

// GetIsReverseZoneOk returns a tuple with the IsReverseZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetIsReverseZoneOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReverseZone) {
		return nil, false
	}
	return o.IsReverseZone, true
}

// HasIsReverseZone returns a boolean if a field has been set.
func (o *CreateZonePayload) HasIsReverseZone() bool {
	if o != nil && !IsNil(o.IsReverseZone) {
		return true
	}

	return false
}

// SetIsReverseZone gets a reference to the given bool and assigns it to the IsReverseZone field.
func (o *CreateZonePayload) SetIsReverseZone(v *bool) {
	o.IsReverseZone = v
}

// GetName returns the Name field value
func (o *CreateZonePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CreateZonePayload) SetName(v *string) {
	o.Name = v
}

// GetNegativeCache returns the NegativeCache field value if set, zero value otherwise.
func (o *CreateZonePayload) GetNegativeCache() *int64 {
	if o == nil || IsNil(o.NegativeCache) {
		var ret *int64
		return ret
	}
	return o.NegativeCache
}

// GetNegativeCacheOk returns a tuple with the NegativeCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetNegativeCacheOk() (*int64, bool) {
	if o == nil || IsNil(o.NegativeCache) {
		return nil, false
	}
	return o.NegativeCache, true
}

// HasNegativeCache returns a boolean if a field has been set.
func (o *CreateZonePayload) HasNegativeCache() bool {
	if o != nil && !IsNil(o.NegativeCache) {
		return true
	}

	return false
}

// SetNegativeCache gets a reference to the given int64 and assigns it to the NegativeCache field.
func (o *CreateZonePayload) SetNegativeCache(v *int64) {
	o.NegativeCache = v
}

// GetPrimaries returns the Primaries field value if set, zero value otherwise.
func (o *CreateZonePayload) GetPrimaries() *[]string {
	if o == nil || IsNil(o.Primaries) {
		var ret *[]string
		return ret
	}
	return o.Primaries
}

// GetPrimariesOk returns a tuple with the Primaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetPrimariesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Primaries) {
		return nil, false
	}
	return o.Primaries, true
}

// HasPrimaries returns a boolean if a field has been set.
func (o *CreateZonePayload) HasPrimaries() bool {
	if o != nil && !IsNil(o.Primaries) {
		return true
	}

	return false
}

// SetPrimaries gets a reference to the given []string and assigns it to the Primaries field.
func (o *CreateZonePayload) SetPrimaries(v *[]string) {
	o.Primaries = v
}

// GetRefreshTime returns the RefreshTime field value if set, zero value otherwise.
func (o *CreateZonePayload) GetRefreshTime() *int64 {
	if o == nil || IsNil(o.RefreshTime) {
		var ret *int64
		return ret
	}
	return o.RefreshTime
}

// GetRefreshTimeOk returns a tuple with the RefreshTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetRefreshTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshTime) {
		return nil, false
	}
	return o.RefreshTime, true
}

// HasRefreshTime returns a boolean if a field has been set.
func (o *CreateZonePayload) HasRefreshTime() bool {
	if o != nil && !IsNil(o.RefreshTime) {
		return true
	}

	return false
}

// SetRefreshTime gets a reference to the given int64 and assigns it to the RefreshTime field.
func (o *CreateZonePayload) SetRefreshTime(v *int64) {
	o.RefreshTime = v
}

// GetRetryTime returns the RetryTime field value if set, zero value otherwise.
func (o *CreateZonePayload) GetRetryTime() *int64 {
	if o == nil || IsNil(o.RetryTime) {
		var ret *int64
		return ret
	}
	return o.RetryTime
}

// GetRetryTimeOk returns a tuple with the RetryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetRetryTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.RetryTime) {
		return nil, false
	}
	return o.RetryTime, true
}

// HasRetryTime returns a boolean if a field has been set.
func (o *CreateZonePayload) HasRetryTime() bool {
	if o != nil && !IsNil(o.RetryTime) {
		return true
	}

	return false
}

// SetRetryTime gets a reference to the given int64 and assigns it to the RetryTime field.
func (o *CreateZonePayload) SetRetryTime(v *int64) {
	o.RetryTime = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateZonePayload) GetType() *string {
	if o == nil || IsNil(o.Type) {
		var ret *string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZonePayload) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateZonePayload) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateZonePayload) SetType(v *string) {
	o.Type = v
}

func (o CreateZonePayload) ToMap() (map[string]interface{}, error) {
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
	toSerialize["dnsName"] = o.DnsName
	if !IsNil(o.ExpireTime) {
		toSerialize["expireTime"] = o.ExpireTime
	}
	if !IsNil(o.IsReverseZone) {
		toSerialize["isReverseZone"] = o.IsReverseZone
	}
	toSerialize["name"] = o.Name
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCreateZonePayload struct {
	value *CreateZonePayload
	isSet bool
}

func (v NullableCreateZonePayload) Get() *CreateZonePayload {
	return v.value
}

func (v *NullableCreateZonePayload) Set(val *CreateZonePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateZonePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateZonePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateZonePayload(val *CreateZonePayload) *NullableCreateZonePayload {
	return &NullableCreateZonePayload{value: val, isSet: true}
}

func (v NullableCreateZonePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateZonePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
