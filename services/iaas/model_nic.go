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

// checks if the NIC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NIC{}

/*
	types and functions for allowedAddresses
*/

// isArray
type NICGetAllowedAddressesAttributeType = *[]AllowedAddressesInner
type NICGetAllowedAddressesArgType = []AllowedAddressesInner
type NICGetAllowedAddressesRetType = []AllowedAddressesInner

func getNICGetAllowedAddressesAttributeTypeOk(arg NICGetAllowedAddressesAttributeType) (ret NICGetAllowedAddressesRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetAllowedAddressesAttributeType(arg *NICGetAllowedAddressesAttributeType, val NICGetAllowedAddressesRetType) {
	*arg = &val
}

/*
	types and functions for device
*/

// isNotNullableString
type NICGetDeviceAttributeType = *string

func getNICGetDeviceAttributeTypeOk(arg NICGetDeviceAttributeType) (ret NICGetDeviceRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetDeviceAttributeType(arg *NICGetDeviceAttributeType, val NICGetDeviceRetType) {
	*arg = &val
}

type NICGetDeviceArgType = string
type NICGetDeviceRetType = string

/*
	types and functions for id
*/

// isNotNullableString
type NICGetIdAttributeType = *string

func getNICGetIdAttributeTypeOk(arg NICGetIdAttributeType) (ret NICGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetIdAttributeType(arg *NICGetIdAttributeType, val NICGetIdRetType) {
	*arg = &val
}

type NICGetIdArgType = string
type NICGetIdRetType = string

/*
	types and functions for ipv4
*/

// isNotNullableString
type NICGetIpv4AttributeType = *string

func getNICGetIpv4AttributeTypeOk(arg NICGetIpv4AttributeType) (ret NICGetIpv4RetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetIpv4AttributeType(arg *NICGetIpv4AttributeType, val NICGetIpv4RetType) {
	*arg = &val
}

type NICGetIpv4ArgType = string
type NICGetIpv4RetType = string

/*
	types and functions for ipv6
*/

// isNotNullableString
type NICGetIpv6AttributeType = *string

func getNICGetIpv6AttributeTypeOk(arg NICGetIpv6AttributeType) (ret NICGetIpv6RetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetIpv6AttributeType(arg *NICGetIpv6AttributeType, val NICGetIpv6RetType) {
	*arg = &val
}

type NICGetIpv6ArgType = string
type NICGetIpv6RetType = string

/*
	types and functions for labels
*/

// isFreeform
type NICGetLabelsAttributeType = *map[string]interface{}
type NICGetLabelsArgType = map[string]interface{}
type NICGetLabelsRetType = map[string]interface{}

func getNICGetLabelsAttributeTypeOk(arg NICGetLabelsAttributeType) (ret NICGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetLabelsAttributeType(arg *NICGetLabelsAttributeType, val NICGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for mac
*/

// isNotNullableString
type NICGetMacAttributeType = *string

func getNICGetMacAttributeTypeOk(arg NICGetMacAttributeType) (ret NICGetMacRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetMacAttributeType(arg *NICGetMacAttributeType, val NICGetMacRetType) {
	*arg = &val
}

type NICGetMacArgType = string
type NICGetMacRetType = string

/*
	types and functions for name
*/

// isNotNullableString
type NICGetNameAttributeType = *string

func getNICGetNameAttributeTypeOk(arg NICGetNameAttributeType) (ret NICGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetNameAttributeType(arg *NICGetNameAttributeType, val NICGetNameRetType) {
	*arg = &val
}

type NICGetNameArgType = string
type NICGetNameRetType = string

/*
	types and functions for networkId
*/

// isNotNullableString
type NICGetNetworkIdAttributeType = *string

func getNICGetNetworkIdAttributeTypeOk(arg NICGetNetworkIdAttributeType) (ret NICGetNetworkIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetNetworkIdAttributeType(arg *NICGetNetworkIdAttributeType, val NICGetNetworkIdRetType) {
	*arg = &val
}

type NICGetNetworkIdArgType = string
type NICGetNetworkIdRetType = string

/*
	types and functions for nicSecurity
*/

// isBoolean
type NICgetNicSecurityAttributeType = *bool
type NICgetNicSecurityArgType = bool
type NICgetNicSecurityRetType = bool

func getNICgetNicSecurityAttributeTypeOk(arg NICgetNicSecurityAttributeType) (ret NICgetNicSecurityRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICgetNicSecurityAttributeType(arg *NICgetNicSecurityAttributeType, val NICgetNicSecurityRetType) {
	*arg = &val
}

/*
	types and functions for securityGroups
*/

// isArray
type NICGetSecurityGroupsAttributeType = *[]string
type NICGetSecurityGroupsArgType = []string
type NICGetSecurityGroupsRetType = []string

func getNICGetSecurityGroupsAttributeTypeOk(arg NICGetSecurityGroupsAttributeType) (ret NICGetSecurityGroupsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetSecurityGroupsAttributeType(arg *NICGetSecurityGroupsAttributeType, val NICGetSecurityGroupsRetType) {
	*arg = &val
}

/*
	types and functions for status
*/

// isNotNullableString
type NICGetStatusAttributeType = *string

func getNICGetStatusAttributeTypeOk(arg NICGetStatusAttributeType) (ret NICGetStatusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetStatusAttributeType(arg *NICGetStatusAttributeType, val NICGetStatusRetType) {
	*arg = &val
}

type NICGetStatusArgType = string
type NICGetStatusRetType = string

/*
	types and functions for type
*/

// isNotNullableString
type NICGetTypeAttributeType = *string

func getNICGetTypeAttributeTypeOk(arg NICGetTypeAttributeType) (ret NICGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setNICGetTypeAttributeType(arg *NICGetTypeAttributeType, val NICGetTypeRetType) {
	*arg = &val
}

type NICGetTypeArgType = string
type NICGetTypeRetType = string

// NIC Object that represents a network interface.
type NIC struct {
	// A list of IPs or CIDR notations.
	AllowedAddresses NICGetAllowedAddressesAttributeType `json:"allowedAddresses,omitempty"`
	// Universally Unique Identifier (UUID).
	Device NICGetDeviceAttributeType `json:"device,omitempty"`
	// Universally Unique Identifier (UUID).
	Id NICGetIdAttributeType `json:"id,omitempty"`
	// Object that represents an IP address.
	Ipv4 NICGetIpv4AttributeType `json:"ipv4,omitempty"`
	// Object that represents an IPv6 address.
	Ipv6 NICGetIpv6AttributeType `json:"ipv6,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels NICGetLabelsAttributeType `json:"labels,omitempty"`
	// Object that represents an MAC address.
	Mac NICGetMacAttributeType `json:"mac,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name NICGetNameAttributeType `json:"name,omitempty"`
	// Universally Unique Identifier (UUID).
	NetworkId NICGetNetworkIdAttributeType `json:"networkId,omitempty"`
	// If this is set to false, then no security groups will apply to this network interface.
	NicSecurity NICgetNicSecurityAttributeType `json:"nicSecurity,omitempty"`
	// A list of UUIDs.
	SecurityGroups NICGetSecurityGroupsAttributeType `json:"securityGroups,omitempty"`
	Status         NICGetStatusAttributeType         `json:"status,omitempty"`
	Type           NICGetTypeAttributeType           `json:"type,omitempty"`
}

// NewNIC instantiates a new NIC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNIC() *NIC {
	this := NIC{}
	return &this
}

// NewNICWithDefaults instantiates a new NIC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNICWithDefaults() *NIC {
	this := NIC{}
	var nicSecurity bool = true
	this.NicSecurity = &nicSecurity
	return &this
}

// GetAllowedAddresses returns the AllowedAddresses field value if set, zero value otherwise.
func (o *NIC) GetAllowedAddresses() (res NICGetAllowedAddressesRetType) {
	res, _ = o.GetAllowedAddressesOk()
	return
}

// GetAllowedAddressesOk returns a tuple with the AllowedAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetAllowedAddressesOk() (ret NICGetAllowedAddressesRetType, ok bool) {
	return getNICGetAllowedAddressesAttributeTypeOk(o.AllowedAddresses)
}

// HasAllowedAddresses returns a boolean if a field has been set.
func (o *NIC) HasAllowedAddresses() bool {
	_, ok := o.GetAllowedAddressesOk()
	return ok
}

// SetAllowedAddresses gets a reference to the given []AllowedAddressesInner and assigns it to the AllowedAddresses field.
func (o *NIC) SetAllowedAddresses(v NICGetAllowedAddressesRetType) {
	setNICGetAllowedAddressesAttributeType(&o.AllowedAddresses, v)
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *NIC) GetDevice() (res NICGetDeviceRetType) {
	res, _ = o.GetDeviceOk()
	return
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetDeviceOk() (ret NICGetDeviceRetType, ok bool) {
	return getNICGetDeviceAttributeTypeOk(o.Device)
}

// HasDevice returns a boolean if a field has been set.
func (o *NIC) HasDevice() bool {
	_, ok := o.GetDeviceOk()
	return ok
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *NIC) SetDevice(v NICGetDeviceRetType) {
	setNICGetDeviceAttributeType(&o.Device, v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NIC) GetId() (res NICGetIdRetType) {
	res, _ = o.GetIdOk()
	return
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetIdOk() (ret NICGetIdRetType, ok bool) {
	return getNICGetIdAttributeTypeOk(o.Id)
}

// HasId returns a boolean if a field has been set.
func (o *NIC) HasId() bool {
	_, ok := o.GetIdOk()
	return ok
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NIC) SetId(v NICGetIdRetType) {
	setNICGetIdAttributeType(&o.Id, v)
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *NIC) GetIpv4() (res NICGetIpv4RetType) {
	res, _ = o.GetIpv4Ok()
	return
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetIpv4Ok() (ret NICGetIpv4RetType, ok bool) {
	return getNICGetIpv4AttributeTypeOk(o.Ipv4)
}

// HasIpv4 returns a boolean if a field has been set.
func (o *NIC) HasIpv4() bool {
	_, ok := o.GetIpv4Ok()
	return ok
}

// SetIpv4 gets a reference to the given string and assigns it to the Ipv4 field.
func (o *NIC) SetIpv4(v NICGetIpv4RetType) {
	setNICGetIpv4AttributeType(&o.Ipv4, v)
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *NIC) GetIpv6() (res NICGetIpv6RetType) {
	res, _ = o.GetIpv6Ok()
	return
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetIpv6Ok() (ret NICGetIpv6RetType, ok bool) {
	return getNICGetIpv6AttributeTypeOk(o.Ipv6)
}

// HasIpv6 returns a boolean if a field has been set.
func (o *NIC) HasIpv6() bool {
	_, ok := o.GetIpv6Ok()
	return ok
}

// SetIpv6 gets a reference to the given string and assigns it to the Ipv6 field.
func (o *NIC) SetIpv6(v NICGetIpv6RetType) {
	setNICGetIpv6AttributeType(&o.Ipv6, v)
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NIC) GetLabels() (res NICGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetLabelsOk() (ret NICGetLabelsRetType, ok bool) {
	return getNICGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *NIC) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *NIC) SetLabels(v NICGetLabelsRetType) {
	setNICGetLabelsAttributeType(&o.Labels, v)
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *NIC) GetMac() (res NICGetMacRetType) {
	res, _ = o.GetMacOk()
	return
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetMacOk() (ret NICGetMacRetType, ok bool) {
	return getNICGetMacAttributeTypeOk(o.Mac)
}

// HasMac returns a boolean if a field has been set.
func (o *NIC) HasMac() bool {
	_, ok := o.GetMacOk()
	return ok
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *NIC) SetMac(v NICGetMacRetType) {
	setNICGetMacAttributeType(&o.Mac, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NIC) GetName() (res NICGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetNameOk() (ret NICGetNameRetType, ok bool) {
	return getNICGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *NIC) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NIC) SetName(v NICGetNameRetType) {
	setNICGetNameAttributeType(&o.Name, v)
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *NIC) GetNetworkId() (res NICGetNetworkIdRetType) {
	res, _ = o.GetNetworkIdOk()
	return
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetNetworkIdOk() (ret NICGetNetworkIdRetType, ok bool) {
	return getNICGetNetworkIdAttributeTypeOk(o.NetworkId)
}

// HasNetworkId returns a boolean if a field has been set.
func (o *NIC) HasNetworkId() bool {
	_, ok := o.GetNetworkIdOk()
	return ok
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *NIC) SetNetworkId(v NICGetNetworkIdRetType) {
	setNICGetNetworkIdAttributeType(&o.NetworkId, v)
}

// GetNicSecurity returns the NicSecurity field value if set, zero value otherwise.
func (o *NIC) GetNicSecurity() (res NICgetNicSecurityRetType) {
	res, _ = o.GetNicSecurityOk()
	return
}

// GetNicSecurityOk returns a tuple with the NicSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetNicSecurityOk() (ret NICgetNicSecurityRetType, ok bool) {
	return getNICgetNicSecurityAttributeTypeOk(o.NicSecurity)
}

// HasNicSecurity returns a boolean if a field has been set.
func (o *NIC) HasNicSecurity() bool {
	_, ok := o.GetNicSecurityOk()
	return ok
}

// SetNicSecurity gets a reference to the given bool and assigns it to the NicSecurity field.
func (o *NIC) SetNicSecurity(v NICgetNicSecurityRetType) {
	setNICgetNicSecurityAttributeType(&o.NicSecurity, v)
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *NIC) GetSecurityGroups() (res NICGetSecurityGroupsRetType) {
	res, _ = o.GetSecurityGroupsOk()
	return
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetSecurityGroupsOk() (ret NICGetSecurityGroupsRetType, ok bool) {
	return getNICGetSecurityGroupsAttributeTypeOk(o.SecurityGroups)
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *NIC) HasSecurityGroups() bool {
	_, ok := o.GetSecurityGroupsOk()
	return ok
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *NIC) SetSecurityGroups(v NICGetSecurityGroupsRetType) {
	setNICGetSecurityGroupsAttributeType(&o.SecurityGroups, v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NIC) GetStatus() (res NICGetStatusRetType) {
	res, _ = o.GetStatusOk()
	return
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetStatusOk() (ret NICGetStatusRetType, ok bool) {
	return getNICGetStatusAttributeTypeOk(o.Status)
}

// HasStatus returns a boolean if a field has been set.
func (o *NIC) HasStatus() bool {
	_, ok := o.GetStatusOk()
	return ok
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NIC) SetStatus(v NICGetStatusRetType) {
	setNICGetStatusAttributeType(&o.Status, v)
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NIC) GetType() (res NICGetTypeRetType) {
	res, _ = o.GetTypeOk()
	return
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetTypeOk() (ret NICGetTypeRetType, ok bool) {
	return getNICGetTypeAttributeTypeOk(o.Type)
}

// HasType returns a boolean if a field has been set.
func (o *NIC) HasType() bool {
	_, ok := o.GetTypeOk()
	return ok
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NIC) SetType(v NICGetTypeRetType) {
	setNICGetTypeAttributeType(&o.Type, v)
}

func (o NIC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getNICGetAllowedAddressesAttributeTypeOk(o.AllowedAddresses); ok {
		toSerialize["AllowedAddresses"] = val
	}
	if val, ok := getNICGetDeviceAttributeTypeOk(o.Device); ok {
		toSerialize["Device"] = val
	}
	if val, ok := getNICGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getNICGetIpv4AttributeTypeOk(o.Ipv4); ok {
		toSerialize["Ipv4"] = val
	}
	if val, ok := getNICGetIpv6AttributeTypeOk(o.Ipv6); ok {
		toSerialize["Ipv6"] = val
	}
	if val, ok := getNICGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getNICGetMacAttributeTypeOk(o.Mac); ok {
		toSerialize["Mac"] = val
	}
	if val, ok := getNICGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getNICGetNetworkIdAttributeTypeOk(o.NetworkId); ok {
		toSerialize["NetworkId"] = val
	}
	if val, ok := getNICgetNicSecurityAttributeTypeOk(o.NicSecurity); ok {
		toSerialize["NicSecurity"] = val
	}
	if val, ok := getNICGetSecurityGroupsAttributeTypeOk(o.SecurityGroups); ok {
		toSerialize["SecurityGroups"] = val
	}
	if val, ok := getNICGetStatusAttributeTypeOk(o.Status); ok {
		toSerialize["Status"] = val
	}
	if val, ok := getNICGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
	return toSerialize, nil
}

type NullableNIC struct {
	value *NIC
	isSet bool
}

func (v NullableNIC) Get() *NIC {
	return v.value
}

func (v *NullableNIC) Set(val *NIC) {
	v.value = val
	v.isSet = true
}

func (v NullableNIC) IsSet() bool {
	return v.isSet
}

func (v *NullableNIC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNIC(val *NIC) *NullableNIC {
	return &NullableNIC{value: val, isSet: true}
}

func (v NullableNIC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNIC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
