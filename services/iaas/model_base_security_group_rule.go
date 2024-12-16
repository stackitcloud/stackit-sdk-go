/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the BaseSecurityGroupRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseSecurityGroupRule{}

// BaseSecurityGroupRule The base schema for a security group rule.
type BaseSecurityGroupRule struct {
	// Description Object. Allows string up to 127 Characters.
	Description *string `json:"description,omitempty"`
	// The direction of the traffic which the rule should match.
	// REQUIRED
	Direction *string `json:"direction"`
	// The ethertype which the rule should match.
	Ethertype      *string         `json:"ethertype,omitempty"`
	IcmpParameters *ICMPParameters `json:"icmpParameters,omitempty"`
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	IpRange   *string    `json:"ipRange,omitempty"`
	PortRange *PortRange `json:"portRange,omitempty"`
	// Universally Unique Identifier (UUID).
	RemoteSecurityGroupId *string `json:"remoteSecurityGroupId,omitempty"`
	// Universally Unique Identifier (UUID).
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
}

type _BaseSecurityGroupRule BaseSecurityGroupRule

// NewBaseSecurityGroupRule instantiates a new BaseSecurityGroupRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseSecurityGroupRule(direction *string) *BaseSecurityGroupRule {
	this := BaseSecurityGroupRule{}
	this.Direction = direction
	var ethertype string = "IPv4"
	this.Ethertype = &ethertype
	return &this
}

// NewBaseSecurityGroupRuleWithDefaults instantiates a new BaseSecurityGroupRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseSecurityGroupRuleWithDefaults() *BaseSecurityGroupRule {
	this := BaseSecurityGroupRule{}
	var ethertype string = "IPv4"
	this.Ethertype = &ethertype
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseSecurityGroupRule) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSecurityGroupRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseSecurityGroupRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseSecurityGroupRule) SetDescription(v *string) {
	o.Description = v
}

// GetDirection returns the Direction field value
func (o *BaseSecurityGroupRule) GetDirection() *string {
	if o == nil || IsNil(o.Direction) {
		var ret *string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *BaseSecurityGroupRule) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Direction, true
}

// SetDirection sets field value
func (o *BaseSecurityGroupRule) SetDirection(v *string) {
	o.Direction = v
}

// GetEthertype returns the Ethertype field value if set, zero value otherwise.
func (o *BaseSecurityGroupRule) GetEthertype() *string {
	if o == nil || IsNil(o.Ethertype) {
		var ret *string
		return ret
	}
	return o.Ethertype
}

// GetEthertypeOk returns a tuple with the Ethertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSecurityGroupRule) GetEthertypeOk() (*string, bool) {
	if o == nil || IsNil(o.Ethertype) {
		return nil, false
	}
	return o.Ethertype, true
}

// HasEthertype returns a boolean if a field has been set.
func (o *BaseSecurityGroupRule) HasEthertype() bool {
	if o != nil && !IsNil(o.Ethertype) {
		return true
	}

	return false
}

// SetEthertype gets a reference to the given string and assigns it to the Ethertype field.
func (o *BaseSecurityGroupRule) SetEthertype(v *string) {
	o.Ethertype = v
}

// GetIcmpParameters returns the IcmpParameters field value if set, zero value otherwise.
func (o *BaseSecurityGroupRule) GetIcmpParameters() *ICMPParameters {
	if o == nil || IsNil(o.IcmpParameters) {
		var ret *ICMPParameters
		return ret
	}
	return o.IcmpParameters
}

// GetIcmpParametersOk returns a tuple with the IcmpParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSecurityGroupRule) GetIcmpParametersOk() (*ICMPParameters, bool) {
	if o == nil || IsNil(o.IcmpParameters) {
		return nil, false
	}
	return o.IcmpParameters, true
}

// HasIcmpParameters returns a boolean if a field has been set.
func (o *BaseSecurityGroupRule) HasIcmpParameters() bool {
	if o != nil && !IsNil(o.IcmpParameters) {
		return true
	}

	return false
}

// SetIcmpParameters gets a reference to the given ICMPParameters and assigns it to the IcmpParameters field.
func (o *BaseSecurityGroupRule) SetIcmpParameters(v *ICMPParameters) {
	o.IcmpParameters = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseSecurityGroupRule) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSecurityGroupRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseSecurityGroupRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseSecurityGroupRule) SetId(v *string) {
	o.Id = v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *BaseSecurityGroupRule) GetIpRange() *string {
	if o == nil || IsNil(o.IpRange) {
		var ret *string
		return ret
	}
	return o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSecurityGroupRule) GetIpRangeOk() (*string, bool) {
	if o == nil || IsNil(o.IpRange) {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *BaseSecurityGroupRule) HasIpRange() bool {
	if o != nil && !IsNil(o.IpRange) {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *BaseSecurityGroupRule) SetIpRange(v *string) {
	o.IpRange = v
}

// GetPortRange returns the PortRange field value if set, zero value otherwise.
func (o *BaseSecurityGroupRule) GetPortRange() *PortRange {
	if o == nil || IsNil(o.PortRange) {
		var ret *PortRange
		return ret
	}
	return o.PortRange
}

// GetPortRangeOk returns a tuple with the PortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSecurityGroupRule) GetPortRangeOk() (*PortRange, bool) {
	if o == nil || IsNil(o.PortRange) {
		return nil, false
	}
	return o.PortRange, true
}

// HasPortRange returns a boolean if a field has been set.
func (o *BaseSecurityGroupRule) HasPortRange() bool {
	if o != nil && !IsNil(o.PortRange) {
		return true
	}

	return false
}

// SetPortRange gets a reference to the given PortRange and assigns it to the PortRange field.
func (o *BaseSecurityGroupRule) SetPortRange(v *PortRange) {
	o.PortRange = v
}

// GetRemoteSecurityGroupId returns the RemoteSecurityGroupId field value if set, zero value otherwise.
func (o *BaseSecurityGroupRule) GetRemoteSecurityGroupId() *string {
	if o == nil || IsNil(o.RemoteSecurityGroupId) {
		var ret *string
		return ret
	}
	return o.RemoteSecurityGroupId
}

// GetRemoteSecurityGroupIdOk returns a tuple with the RemoteSecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSecurityGroupRule) GetRemoteSecurityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteSecurityGroupId) {
		return nil, false
	}
	return o.RemoteSecurityGroupId, true
}

// HasRemoteSecurityGroupId returns a boolean if a field has been set.
func (o *BaseSecurityGroupRule) HasRemoteSecurityGroupId() bool {
	if o != nil && !IsNil(o.RemoteSecurityGroupId) {
		return true
	}

	return false
}

// SetRemoteSecurityGroupId gets a reference to the given string and assigns it to the RemoteSecurityGroupId field.
func (o *BaseSecurityGroupRule) SetRemoteSecurityGroupId(v *string) {
	o.RemoteSecurityGroupId = v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *BaseSecurityGroupRule) GetSecurityGroupId() *string {
	if o == nil || IsNil(o.SecurityGroupId) {
		var ret *string
		return ret
	}
	return o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSecurityGroupRule) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityGroupId) {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *BaseSecurityGroupRule) HasSecurityGroupId() bool {
	if o != nil && !IsNil(o.SecurityGroupId) {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *BaseSecurityGroupRule) SetSecurityGroupId(v *string) {
	o.SecurityGroupId = v
}

func (o BaseSecurityGroupRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["direction"] = o.Direction
	if !IsNil(o.Ethertype) {
		toSerialize["ethertype"] = o.Ethertype
	}
	if !IsNil(o.IcmpParameters) {
		toSerialize["icmpParameters"] = o.IcmpParameters
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IpRange) {
		toSerialize["ipRange"] = o.IpRange
	}
	if !IsNil(o.PortRange) {
		toSerialize["portRange"] = o.PortRange
	}
	if !IsNil(o.RemoteSecurityGroupId) {
		toSerialize["remoteSecurityGroupId"] = o.RemoteSecurityGroupId
	}
	if !IsNil(o.SecurityGroupId) {
		toSerialize["securityGroupId"] = o.SecurityGroupId
	}
	return toSerialize, nil
}

type NullableBaseSecurityGroupRule struct {
	value *BaseSecurityGroupRule
	isSet bool
}

func (v NullableBaseSecurityGroupRule) Get() *BaseSecurityGroupRule {
	return v.value
}

func (v *NullableBaseSecurityGroupRule) Set(val *BaseSecurityGroupRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseSecurityGroupRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseSecurityGroupRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseSecurityGroupRule(val *BaseSecurityGroupRule) *NullableBaseSecurityGroupRule {
	return &NullableBaseSecurityGroupRule{value: val, isSet: true}
}

func (v NullableBaseSecurityGroupRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseSecurityGroupRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
