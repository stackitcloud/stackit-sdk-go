/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the SecurityGroupRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupRule{}

// SecurityGroupRule Object that represents a security group rule.
type SecurityGroupRule struct {
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
	SecurityGroupId *string   `json:"securityGroupId,omitempty"`
	Protocol        *Protocol `json:"protocol,omitempty"`
}

type _SecurityGroupRule SecurityGroupRule

// NewSecurityGroupRule instantiates a new SecurityGroupRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupRule(direction *string) *SecurityGroupRule {
	this := SecurityGroupRule{}
	this.Direction = direction
	var ethertype string = "IPv4"
	this.Ethertype = &ethertype
	return &this
}

// NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule {
	this := SecurityGroupRule{}
	var ethertype string = "IPv4"
	this.Ethertype = &ethertype
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityGroupRule) SetDescription(v *string) {
	o.Description = v
}

// GetDirection returns the Direction field value
func (o *SecurityGroupRule) GetDirection() *string {
	if o == nil || IsNil(o.Direction) {
		var ret *string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Direction, true
}

// SetDirection sets field value
func (o *SecurityGroupRule) SetDirection(v *string) {
	o.Direction = v
}

// GetEthertype returns the Ethertype field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetEthertype() *string {
	if o == nil || IsNil(o.Ethertype) {
		var ret *string
		return ret
	}
	return o.Ethertype
}

// GetEthertypeOk returns a tuple with the Ethertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetEthertypeOk() (*string, bool) {
	if o == nil || IsNil(o.Ethertype) {
		return nil, false
	}
	return o.Ethertype, true
}

// HasEthertype returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasEthertype() bool {
	if o != nil && !IsNil(o.Ethertype) {
		return true
	}

	return false
}

// SetEthertype gets a reference to the given string and assigns it to the Ethertype field.
func (o *SecurityGroupRule) SetEthertype(v *string) {
	o.Ethertype = v
}

// GetIcmpParameters returns the IcmpParameters field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIcmpParameters() *ICMPParameters {
	if o == nil || IsNil(o.IcmpParameters) {
		var ret *ICMPParameters
		return ret
	}
	return o.IcmpParameters
}

// GetIcmpParametersOk returns a tuple with the IcmpParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIcmpParametersOk() (*ICMPParameters, bool) {
	if o == nil || IsNil(o.IcmpParameters) {
		return nil, false
	}
	return o.IcmpParameters, true
}

// HasIcmpParameters returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIcmpParameters() bool {
	if o != nil && !IsNil(o.IcmpParameters) {
		return true
	}

	return false
}

// SetIcmpParameters gets a reference to the given ICMPParameters and assigns it to the IcmpParameters field.
func (o *SecurityGroupRule) SetIcmpParameters(v *ICMPParameters) {
	o.IcmpParameters = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecurityGroupRule) SetId(v *string) {
	o.Id = v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIpRange() *string {
	if o == nil || IsNil(o.IpRange) {
		var ret *string
		return ret
	}
	return o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIpRangeOk() (*string, bool) {
	if o == nil || IsNil(o.IpRange) {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIpRange() bool {
	if o != nil && !IsNil(o.IpRange) {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *SecurityGroupRule) SetIpRange(v *string) {
	o.IpRange = v
}

// GetPortRange returns the PortRange field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetPortRange() *PortRange {
	if o == nil || IsNil(o.PortRange) {
		var ret *PortRange
		return ret
	}
	return o.PortRange
}

// GetPortRangeOk returns a tuple with the PortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetPortRangeOk() (*PortRange, bool) {
	if o == nil || IsNil(o.PortRange) {
		return nil, false
	}
	return o.PortRange, true
}

// HasPortRange returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasPortRange() bool {
	if o != nil && !IsNil(o.PortRange) {
		return true
	}

	return false
}

// SetPortRange gets a reference to the given PortRange and assigns it to the PortRange field.
func (o *SecurityGroupRule) SetPortRange(v *PortRange) {
	o.PortRange = v
}

// GetRemoteSecurityGroupId returns the RemoteSecurityGroupId field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetRemoteSecurityGroupId() *string {
	if o == nil || IsNil(o.RemoteSecurityGroupId) {
		var ret *string
		return ret
	}
	return o.RemoteSecurityGroupId
}

// GetRemoteSecurityGroupIdOk returns a tuple with the RemoteSecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetRemoteSecurityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteSecurityGroupId) {
		return nil, false
	}
	return o.RemoteSecurityGroupId, true
}

// HasRemoteSecurityGroupId returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasRemoteSecurityGroupId() bool {
	if o != nil && !IsNil(o.RemoteSecurityGroupId) {
		return true
	}

	return false
}

// SetRemoteSecurityGroupId gets a reference to the given string and assigns it to the RemoteSecurityGroupId field.
func (o *SecurityGroupRule) SetRemoteSecurityGroupId(v *string) {
	o.RemoteSecurityGroupId = v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetSecurityGroupId() *string {
	if o == nil || IsNil(o.SecurityGroupId) {
		var ret *string
		return ret
	}
	return o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityGroupId) {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasSecurityGroupId() bool {
	if o != nil && !IsNil(o.SecurityGroupId) {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *SecurityGroupRule) SetSecurityGroupId(v *string) {
	o.SecurityGroupId = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetProtocol() *Protocol {
	if o == nil || IsNil(o.Protocol) {
		var ret *Protocol
		return ret
	}
	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetProtocolOk() (*Protocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given Protocol and assigns it to the Protocol field.
func (o *SecurityGroupRule) SetProtocol(v *Protocol) {
	o.Protocol = v
}

func (o SecurityGroupRule) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	return toSerialize, nil
}

type NullableSecurityGroupRule struct {
	value *SecurityGroupRule
	isSet bool
}

func (v NullableSecurityGroupRule) Get() *SecurityGroupRule {
	return v.value
}

func (v *NullableSecurityGroupRule) Set(val *SecurityGroupRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupRule(val *SecurityGroupRule) *NullableSecurityGroupRule {
	return &NullableSecurityGroupRule{value: val, isSet: true}
}

func (v NullableSecurityGroupRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
