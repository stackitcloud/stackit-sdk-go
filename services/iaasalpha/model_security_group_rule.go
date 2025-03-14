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

/*
	types and functions for description
*/

// isNotNullableString
type SecurityGroupRuleGetDescriptionAttributeType = *string

func getSecurityGroupRuleGetDescriptionAttributeTypeOk(arg SecurityGroupRuleGetDescriptionAttributeType) (ret SecurityGroupRuleGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetDescriptionAttributeType(arg *SecurityGroupRuleGetDescriptionAttributeType, val SecurityGroupRuleGetDescriptionRetType) {
	*arg = &val
}

type SecurityGroupRuleGetDescriptionArgType = string
type SecurityGroupRuleGetDescriptionRetType = string

/*
	types and functions for direction
*/

// isNotNullableString
type SecurityGroupRuleGetDirectionAttributeType = *string

func getSecurityGroupRuleGetDirectionAttributeTypeOk(arg SecurityGroupRuleGetDirectionAttributeType) (ret SecurityGroupRuleGetDirectionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetDirectionAttributeType(arg *SecurityGroupRuleGetDirectionAttributeType, val SecurityGroupRuleGetDirectionRetType) {
	*arg = &val
}

type SecurityGroupRuleGetDirectionArgType = string
type SecurityGroupRuleGetDirectionRetType = string

/*
	types and functions for ethertype
*/

// isNotNullableString
type SecurityGroupRuleGetEthertypeAttributeType = *string

func getSecurityGroupRuleGetEthertypeAttributeTypeOk(arg SecurityGroupRuleGetEthertypeAttributeType) (ret SecurityGroupRuleGetEthertypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetEthertypeAttributeType(arg *SecurityGroupRuleGetEthertypeAttributeType, val SecurityGroupRuleGetEthertypeRetType) {
	*arg = &val
}

type SecurityGroupRuleGetEthertypeArgType = string
type SecurityGroupRuleGetEthertypeRetType = string

/*
	types and functions for icmpParameters
*/

// isModel
type SecurityGroupRuleGetIcmpParametersAttributeType = *ICMPParameters
type SecurityGroupRuleGetIcmpParametersArgType = ICMPParameters
type SecurityGroupRuleGetIcmpParametersRetType = ICMPParameters

func getSecurityGroupRuleGetIcmpParametersAttributeTypeOk(arg SecurityGroupRuleGetIcmpParametersAttributeType) (ret SecurityGroupRuleGetIcmpParametersRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetIcmpParametersAttributeType(arg *SecurityGroupRuleGetIcmpParametersAttributeType, val SecurityGroupRuleGetIcmpParametersRetType) {
	*arg = &val
}

/*
	types and functions for id
*/

// isNotNullableString
type SecurityGroupRuleGetIdAttributeType = *string

func getSecurityGroupRuleGetIdAttributeTypeOk(arg SecurityGroupRuleGetIdAttributeType) (ret SecurityGroupRuleGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetIdAttributeType(arg *SecurityGroupRuleGetIdAttributeType, val SecurityGroupRuleGetIdRetType) {
	*arg = &val
}

type SecurityGroupRuleGetIdArgType = string
type SecurityGroupRuleGetIdRetType = string

/*
	types and functions for ipRange
*/

// isNotNullableString
type SecurityGroupRuleGetIpRangeAttributeType = *string

func getSecurityGroupRuleGetIpRangeAttributeTypeOk(arg SecurityGroupRuleGetIpRangeAttributeType) (ret SecurityGroupRuleGetIpRangeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetIpRangeAttributeType(arg *SecurityGroupRuleGetIpRangeAttributeType, val SecurityGroupRuleGetIpRangeRetType) {
	*arg = &val
}

type SecurityGroupRuleGetIpRangeArgType = string
type SecurityGroupRuleGetIpRangeRetType = string

/*
	types and functions for portRange
*/

// isModel
type SecurityGroupRuleGetPortRangeAttributeType = *PortRange
type SecurityGroupRuleGetPortRangeArgType = PortRange
type SecurityGroupRuleGetPortRangeRetType = PortRange

func getSecurityGroupRuleGetPortRangeAttributeTypeOk(arg SecurityGroupRuleGetPortRangeAttributeType) (ret SecurityGroupRuleGetPortRangeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetPortRangeAttributeType(arg *SecurityGroupRuleGetPortRangeAttributeType, val SecurityGroupRuleGetPortRangeRetType) {
	*arg = &val
}

/*
	types and functions for remoteSecurityGroupId
*/

// isNotNullableString
type SecurityGroupRuleGetRemoteSecurityGroupIdAttributeType = *string

func getSecurityGroupRuleGetRemoteSecurityGroupIdAttributeTypeOk(arg SecurityGroupRuleGetRemoteSecurityGroupIdAttributeType) (ret SecurityGroupRuleGetRemoteSecurityGroupIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetRemoteSecurityGroupIdAttributeType(arg *SecurityGroupRuleGetRemoteSecurityGroupIdAttributeType, val SecurityGroupRuleGetRemoteSecurityGroupIdRetType) {
	*arg = &val
}

type SecurityGroupRuleGetRemoteSecurityGroupIdArgType = string
type SecurityGroupRuleGetRemoteSecurityGroupIdRetType = string

/*
	types and functions for securityGroupId
*/

// isNotNullableString
type SecurityGroupRuleGetSecurityGroupIdAttributeType = *string

func getSecurityGroupRuleGetSecurityGroupIdAttributeTypeOk(arg SecurityGroupRuleGetSecurityGroupIdAttributeType) (ret SecurityGroupRuleGetSecurityGroupIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetSecurityGroupIdAttributeType(arg *SecurityGroupRuleGetSecurityGroupIdAttributeType, val SecurityGroupRuleGetSecurityGroupIdRetType) {
	*arg = &val
}

type SecurityGroupRuleGetSecurityGroupIdArgType = string
type SecurityGroupRuleGetSecurityGroupIdRetType = string

/*
	types and functions for protocol
*/

// isModel
type SecurityGroupRuleGetProtocolAttributeType = *Protocol
type SecurityGroupRuleGetProtocolArgType = Protocol
type SecurityGroupRuleGetProtocolRetType = Protocol

func getSecurityGroupRuleGetProtocolAttributeTypeOk(arg SecurityGroupRuleGetProtocolAttributeType) (ret SecurityGroupRuleGetProtocolRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setSecurityGroupRuleGetProtocolAttributeType(arg *SecurityGroupRuleGetProtocolAttributeType, val SecurityGroupRuleGetProtocolRetType) {
	*arg = &val
}

// SecurityGroupRule Object that represents a security group rule.
type SecurityGroupRule struct {
	// Description Object. Allows string up to 127 Characters.
	Description SecurityGroupRuleGetDescriptionAttributeType `json:"description,omitempty"`
	// The direction of the traffic which the rule should match.
	// REQUIRED
	Direction SecurityGroupRuleGetDirectionAttributeType `json:"direction"`
	// The ethertype which the rule should match.
	Ethertype      SecurityGroupRuleGetEthertypeAttributeType      `json:"ethertype,omitempty"`
	IcmpParameters SecurityGroupRuleGetIcmpParametersAttributeType `json:"icmpParameters,omitempty"`
	// Universally Unique Identifier (UUID).
	Id SecurityGroupRuleGetIdAttributeType `json:"id,omitempty"`
	// Classless Inter-Domain Routing (CIDR).
	IpRange   SecurityGroupRuleGetIpRangeAttributeType   `json:"ipRange,omitempty"`
	PortRange SecurityGroupRuleGetPortRangeAttributeType `json:"portRange,omitempty"`
	// Universally Unique Identifier (UUID).
	RemoteSecurityGroupId SecurityGroupRuleGetRemoteSecurityGroupIdAttributeType `json:"remoteSecurityGroupId,omitempty"`
	// Universally Unique Identifier (UUID).
	SecurityGroupId SecurityGroupRuleGetSecurityGroupIdAttributeType `json:"securityGroupId,omitempty"`
	Protocol        SecurityGroupRuleGetProtocolAttributeType        `json:"protocol,omitempty"`
}

type _SecurityGroupRule SecurityGroupRule

// NewSecurityGroupRule instantiates a new SecurityGroupRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupRule(direction SecurityGroupRuleGetDirectionArgType) *SecurityGroupRule {
	this := SecurityGroupRule{}
	setSecurityGroupRuleGetDirectionAttributeType(&this.Direction, direction)
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
func (o *SecurityGroupRule) GetDescription() (res SecurityGroupRuleGetDescriptionRetType) {
	res, _ = o.GetDescriptionOk()
	return
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetDescriptionOk() (ret SecurityGroupRuleGetDescriptionRetType, ok bool) {
	return getSecurityGroupRuleGetDescriptionAttributeTypeOk(o.Description)
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasDescription() bool {
	_, ok := o.GetDescriptionOk()
	return ok
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityGroupRule) SetDescription(v SecurityGroupRuleGetDescriptionRetType) {
	setSecurityGroupRuleGetDescriptionAttributeType(&o.Description, v)
}

// GetDirection returns the Direction field value
func (o *SecurityGroupRule) GetDirection() (ret SecurityGroupRuleGetDirectionRetType) {
	ret, _ = o.GetDirectionOk()
	return ret
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetDirectionOk() (ret SecurityGroupRuleGetDirectionRetType, ok bool) {
	return getSecurityGroupRuleGetDirectionAttributeTypeOk(o.Direction)
}

// SetDirection sets field value
func (o *SecurityGroupRule) SetDirection(v SecurityGroupRuleGetDirectionRetType) {
	setSecurityGroupRuleGetDirectionAttributeType(&o.Direction, v)
}

// GetEthertype returns the Ethertype field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetEthertype() (res SecurityGroupRuleGetEthertypeRetType) {
	res, _ = o.GetEthertypeOk()
	return
}

// GetEthertypeOk returns a tuple with the Ethertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetEthertypeOk() (ret SecurityGroupRuleGetEthertypeRetType, ok bool) {
	return getSecurityGroupRuleGetEthertypeAttributeTypeOk(o.Ethertype)
}

// HasEthertype returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasEthertype() bool {
	_, ok := o.GetEthertypeOk()
	return ok
}

// SetEthertype gets a reference to the given string and assigns it to the Ethertype field.
func (o *SecurityGroupRule) SetEthertype(v SecurityGroupRuleGetEthertypeRetType) {
	setSecurityGroupRuleGetEthertypeAttributeType(&o.Ethertype, v)
}

// GetIcmpParameters returns the IcmpParameters field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIcmpParameters() (res SecurityGroupRuleGetIcmpParametersRetType) {
	res, _ = o.GetIcmpParametersOk()
	return
}

// GetIcmpParametersOk returns a tuple with the IcmpParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIcmpParametersOk() (ret SecurityGroupRuleGetIcmpParametersRetType, ok bool) {
	return getSecurityGroupRuleGetIcmpParametersAttributeTypeOk(o.IcmpParameters)
}

// HasIcmpParameters returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIcmpParameters() bool {
	_, ok := o.GetIcmpParametersOk()
	return ok
}

// SetIcmpParameters gets a reference to the given ICMPParameters and assigns it to the IcmpParameters field.
func (o *SecurityGroupRule) SetIcmpParameters(v SecurityGroupRuleGetIcmpParametersRetType) {
	setSecurityGroupRuleGetIcmpParametersAttributeType(&o.IcmpParameters, v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetId() (res SecurityGroupRuleGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIdOk() (ret SecurityGroupRuleGetIdRetType, ok bool) {
	return getSecurityGroupRuleGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecurityGroupRule) SetId(v SecurityGroupRuleGetIdRetType) {
	setSecurityGroupRuleGetIdAttributeType(&o.Id, v)
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIpRange() (res SecurityGroupRuleGetIpRangeRetType) {
	res, _ = o.GetIpRangeOk()
	return
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIpRangeOk() (ret SecurityGroupRuleGetIpRangeRetType, ok bool) {
	return getSecurityGroupRuleGetIpRangeAttributeTypeOk(o.IpRange)
}

// HasIpRange returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIpRange() bool {
	_, ok := o.GetIpRangeOk()
	return ok
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *SecurityGroupRule) SetIpRange(v SecurityGroupRuleGetIpRangeRetType) {
	setSecurityGroupRuleGetIpRangeAttributeType(&o.IpRange, v)
}

// GetPortRange returns the PortRange field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetPortRange() (res SecurityGroupRuleGetPortRangeRetType) {
	res, _ = o.GetPortRangeOk()
	return
}

// GetPortRangeOk returns a tuple with the PortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetPortRangeOk() (ret SecurityGroupRuleGetPortRangeRetType, ok bool) {
	return getSecurityGroupRuleGetPortRangeAttributeTypeOk(o.PortRange)
}

// HasPortRange returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasPortRange() bool {
	_, ok := o.GetPortRangeOk()
	return ok
}

// SetPortRange gets a reference to the given PortRange and assigns it to the PortRange field.
func (o *SecurityGroupRule) SetPortRange(v SecurityGroupRuleGetPortRangeRetType) {
	setSecurityGroupRuleGetPortRangeAttributeType(&o.PortRange, v)
}

// GetRemoteSecurityGroupId returns the RemoteSecurityGroupId field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetRemoteSecurityGroupId() (res SecurityGroupRuleGetRemoteSecurityGroupIdRetType) {
	res, _ = o.GetRemoteSecurityGroupIdOk()
	return
}

// GetRemoteSecurityGroupIdOk returns a tuple with the RemoteSecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetRemoteSecurityGroupIdOk() (ret SecurityGroupRuleGetRemoteSecurityGroupIdRetType, ok bool) {
	return getSecurityGroupRuleGetRemoteSecurityGroupIdAttributeTypeOk(o.RemoteSecurityGroupId)
}

// HasRemoteSecurityGroupId returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasRemoteSecurityGroupId() bool {
	_, ok := o.GetRemoteSecurityGroupIdOk()
	return ok
}

// SetRemoteSecurityGroupId gets a reference to the given string and assigns it to the RemoteSecurityGroupId field.
func (o *SecurityGroupRule) SetRemoteSecurityGroupId(v SecurityGroupRuleGetRemoteSecurityGroupIdRetType) {
	setSecurityGroupRuleGetRemoteSecurityGroupIdAttributeType(&o.RemoteSecurityGroupId, v)
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetSecurityGroupId() (res SecurityGroupRuleGetSecurityGroupIdRetType) {
	res, _ = o.GetSecurityGroupIdOk()
	return
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetSecurityGroupIdOk() (ret SecurityGroupRuleGetSecurityGroupIdRetType, ok bool) {
	return getSecurityGroupRuleGetSecurityGroupIdAttributeTypeOk(o.SecurityGroupId)
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasSecurityGroupId() bool {
	_, ok := o.GetSecurityGroupIdOk()
	return ok
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *SecurityGroupRule) SetSecurityGroupId(v SecurityGroupRuleGetSecurityGroupIdRetType) {
	setSecurityGroupRuleGetSecurityGroupIdAttributeType(&o.SecurityGroupId, v)
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetProtocol() (res SecurityGroupRuleGetProtocolRetType) {
	res, _ = o.GetProtocolOk()
	return
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetProtocolOk() (ret SecurityGroupRuleGetProtocolRetType, ok bool) {
	return getSecurityGroupRuleGetProtocolAttributeTypeOk(o.Protocol)
}

// HasProtocol returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasProtocol() bool {
	_, ok := o.GetProtocolOk()
	return ok
}

// SetProtocol gets a reference to the given Protocol and assigns it to the Protocol field.
func (o *SecurityGroupRule) SetProtocol(v SecurityGroupRuleGetProtocolRetType) {
	setSecurityGroupRuleGetProtocolAttributeType(&o.Protocol, v)
}

func (o SecurityGroupRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getSecurityGroupRuleGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getSecurityGroupRuleGetDirectionAttributeTypeOk(o.Direction); ok {
		toSerialize["Direction"] = val
	}
	if val, ok := getSecurityGroupRuleGetEthertypeAttributeTypeOk(o.Ethertype); ok {
		toSerialize["Ethertype"] = val
	}
	if val, ok := getSecurityGroupRuleGetIcmpParametersAttributeTypeOk(o.IcmpParameters); ok {
		toSerialize["IcmpParameters"] = val
	}
	if val, ok := getSecurityGroupRuleGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getSecurityGroupRuleGetIpRangeAttributeTypeOk(o.IpRange); ok {
		toSerialize["IpRange"] = val
	}
	if val, ok := getSecurityGroupRuleGetPortRangeAttributeTypeOk(o.PortRange); ok {
		toSerialize["PortRange"] = val
	}
	if val, ok := getSecurityGroupRuleGetRemoteSecurityGroupIdAttributeTypeOk(o.RemoteSecurityGroupId); ok {
		toSerialize["RemoteSecurityGroupId"] = val
	}
	if val, ok := getSecurityGroupRuleGetSecurityGroupIdAttributeTypeOk(o.SecurityGroupId); ok {
		toSerialize["SecurityGroupId"] = val
	}
	if val, ok := getSecurityGroupRuleGetProtocolAttributeTypeOk(o.Protocol); ok {
		toSerialize["Protocol"] = val
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
