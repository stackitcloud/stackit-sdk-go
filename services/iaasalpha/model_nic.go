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

// checks if the NIC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NIC{}

// NIC Object that represents a network interface.
type NIC struct {
	// A list of IPs or CIDR notations.
	AllowedAddresses *[]AllowedAddressesInner `json:"allowedAddresses,omitempty"`
	// Universally Unique Identifier (UUID).
	Device *string `json:"device,omitempty"`
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Object that represents an IP address.
	Ipv4 *string `json:"ipv4,omitempty"`
	// Object that represents an IPv6 address.
	Ipv6 *string `json:"ipv6,omitempty"`
	// Object that represents the labels of an object.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// Object that represents an MAC address.
	Mac *string `json:"mac,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
	// Universally Unique Identifier (UUID).
	NetworkId *string `json:"networkId,omitempty"`
	// If this is set to false, then no security groups will apply to this network interface.
	NicSecurity *bool `json:"nicSecurity,omitempty"`
	// A list of UUIDs.
	SecurityGroups *[]string `json:"securityGroups,omitempty"`
	Status         *string   `json:"status,omitempty"`
	Type           *string   `json:"type,omitempty"`
}

// NewNIC instantiates a new NIC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNIC() *NIC {
	this := NIC{}
	var nicSecurity bool = true
	this.NicSecurity = &nicSecurity
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
func (o *NIC) GetAllowedAddresses() *[]AllowedAddressesInner {
	if o == nil || IsNil(o.AllowedAddresses) {
		var ret *[]AllowedAddressesInner
		return ret
	}
	return o.AllowedAddresses
}

// GetAllowedAddressesOk returns a tuple with the AllowedAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetAllowedAddressesOk() (*[]AllowedAddressesInner, bool) {
	if o == nil || IsNil(o.AllowedAddresses) {
		return nil, false
	}
	return o.AllowedAddresses, true
}

// HasAllowedAddresses returns a boolean if a field has been set.
func (o *NIC) HasAllowedAddresses() bool {
	if o != nil && !IsNil(o.AllowedAddresses) && !IsNil(o.AllowedAddresses) {
		return true
	}

	return false
}

// SetAllowedAddresses gets a reference to the given []AllowedAddressesInner and assigns it to the AllowedAddresses field.
func (o *NIC) SetAllowedAddresses(v *[]AllowedAddressesInner) {
	o.AllowedAddresses = v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *NIC) GetDevice() *string {
	if o == nil || IsNil(o.Device) {
		var ret *string
		return ret
	}
	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *NIC) HasDevice() bool {
	if o != nil && !IsNil(o.Device) && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *NIC) SetDevice(v *string) {
	o.Device = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NIC) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NIC) HasId() bool {
	if o != nil && !IsNil(o.Id) && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NIC) SetId(v *string) {
	o.Id = v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *NIC) GetIpv4() *string {
	if o == nil || IsNil(o.Ipv4) {
		var ret *string
		return ret
	}
	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetIpv4Ok() (*string, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *NIC) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given string and assigns it to the Ipv4 field.
func (o *NIC) SetIpv4(v *string) {
	o.Ipv4 = v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *NIC) GetIpv6() *string {
	if o == nil || IsNil(o.Ipv6) {
		var ret *string
		return ret
	}
	return o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *NIC) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given string and assigns it to the Ipv6 field.
func (o *NIC) SetIpv6(v *string) {
	o.Ipv6 = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *NIC) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *NIC) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *NIC) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *NIC) GetMac() *string {
	if o == nil || IsNil(o.Mac) {
		var ret *string
		return ret
	}
	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *NIC) HasMac() bool {
	if o != nil && !IsNil(o.Mac) && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *NIC) SetMac(v *string) {
	o.Mac = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NIC) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NIC) HasName() bool {
	if o != nil && !IsNil(o.Name) && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NIC) SetName(v *string) {
	o.Name = v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *NIC) GetNetworkId() *string {
	if o == nil || IsNil(o.NetworkId) {
		var ret *string
		return ret
	}
	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *NIC) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *NIC) SetNetworkId(v *string) {
	o.NetworkId = v
}

// GetNicSecurity returns the NicSecurity field value if set, zero value otherwise.
func (o *NIC) GetNicSecurity() *bool {
	if o == nil || IsNil(o.NicSecurity) {
		var ret *bool
		return ret
	}
	return o.NicSecurity
}

// GetNicSecurityOk returns a tuple with the NicSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetNicSecurityOk() (*bool, bool) {
	if o == nil || IsNil(o.NicSecurity) {
		return nil, false
	}
	return o.NicSecurity, true
}

// HasNicSecurity returns a boolean if a field has been set.
func (o *NIC) HasNicSecurity() bool {
	if o != nil && !IsNil(o.NicSecurity) && !IsNil(o.NicSecurity) {
		return true
	}

	return false
}

// SetNicSecurity gets a reference to the given bool and assigns it to the NicSecurity field.
func (o *NIC) SetNicSecurity(v *bool) {
	o.NicSecurity = v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *NIC) GetSecurityGroups() *[]string {
	if o == nil || IsNil(o.SecurityGroups) {
		var ret *[]string
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *NIC) HasSecurityGroups() bool {
	if o != nil && !IsNil(o.SecurityGroups) && !IsNil(o.SecurityGroups) {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *NIC) SetSecurityGroups(v *[]string) {
	o.SecurityGroups = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NIC) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NIC) HasStatus() bool {
	if o != nil && !IsNil(o.Status) && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NIC) SetStatus(v *string) {
	o.Status = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NIC) GetType() *string {
	if o == nil || IsNil(o.Type) {
		var ret *string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NIC) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NIC) HasType() bool {
	if o != nil && !IsNil(o.Type) && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NIC) SetType(v *string) {
	o.Type = v
}

func (o NIC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedAddresses) {
		toSerialize["allowedAddresses"] = o.AllowedAddresses
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.NicSecurity) {
		toSerialize["nicSecurity"] = o.NicSecurity
	}
	if !IsNil(o.SecurityGroups) {
		toSerialize["securityGroups"] = o.SecurityGroups
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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
