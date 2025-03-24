/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the PartialUpdateNetworkPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialUpdateNetworkPayload{}

/*
	types and functions for addressFamily
*/

// isModel
type PartialUpdateNetworkPayloadGetAddressFamilyAttributeType = *UpdateNetworkAddressFamily
type PartialUpdateNetworkPayloadGetAddressFamilyArgType = UpdateNetworkAddressFamily
type PartialUpdateNetworkPayloadGetAddressFamilyRetType = UpdateNetworkAddressFamily

func getPartialUpdateNetworkPayloadGetAddressFamilyAttributeTypeOk(arg PartialUpdateNetworkPayloadGetAddressFamilyAttributeType) (ret PartialUpdateNetworkPayloadGetAddressFamilyRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPartialUpdateNetworkPayloadGetAddressFamilyAttributeType(arg *PartialUpdateNetworkPayloadGetAddressFamilyAttributeType, val PartialUpdateNetworkPayloadGetAddressFamilyRetType) {
	*arg = &val
}

/*
	types and functions for labels
*/

// isFreeform
type PartialUpdateNetworkPayloadGetLabelsAttributeType = *map[string]interface{}
type PartialUpdateNetworkPayloadGetLabelsArgType = map[string]interface{}
type PartialUpdateNetworkPayloadGetLabelsRetType = map[string]interface{}

func getPartialUpdateNetworkPayloadGetLabelsAttributeTypeOk(arg PartialUpdateNetworkPayloadGetLabelsAttributeType) (ret PartialUpdateNetworkPayloadGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPartialUpdateNetworkPayloadGetLabelsAttributeType(arg *PartialUpdateNetworkPayloadGetLabelsAttributeType, val PartialUpdateNetworkPayloadGetLabelsRetType) {
	*arg = &val
}

/*
	types and functions for name
*/

// isNotNullableString
type PartialUpdateNetworkPayloadGetNameAttributeType = *string

func getPartialUpdateNetworkPayloadGetNameAttributeTypeOk(arg PartialUpdateNetworkPayloadGetNameAttributeType) (ret PartialUpdateNetworkPayloadGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPartialUpdateNetworkPayloadGetNameAttributeType(arg *PartialUpdateNetworkPayloadGetNameAttributeType, val PartialUpdateNetworkPayloadGetNameRetType) {
	*arg = &val
}

type PartialUpdateNetworkPayloadGetNameArgType = string
type PartialUpdateNetworkPayloadGetNameRetType = string

/*
	types and functions for routed
*/

// isBoolean
type PartialUpdateNetworkPayloadgetRoutedAttributeType = *bool
type PartialUpdateNetworkPayloadgetRoutedArgType = bool
type PartialUpdateNetworkPayloadgetRoutedRetType = bool

func getPartialUpdateNetworkPayloadgetRoutedAttributeTypeOk(arg PartialUpdateNetworkPayloadgetRoutedAttributeType) (ret PartialUpdateNetworkPayloadgetRoutedRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setPartialUpdateNetworkPayloadgetRoutedAttributeType(arg *PartialUpdateNetworkPayloadgetRoutedAttributeType, val PartialUpdateNetworkPayloadgetRoutedRetType) {
	*arg = &val
}

// PartialUpdateNetworkPayload Object that represents the request body for a network update.
type PartialUpdateNetworkPayload struct {
	AddressFamily PartialUpdateNetworkPayloadGetAddressFamilyAttributeType `json:"addressFamily,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels PartialUpdateNetworkPayloadGetLabelsAttributeType `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name PartialUpdateNetworkPayloadGetNameAttributeType `json:"name,omitempty"`
	// Shows if the network is routed and therefore accessible from other networks.
	Routed PartialUpdateNetworkPayloadgetRoutedAttributeType `json:"routed,omitempty"`
}

// NewPartialUpdateNetworkPayload instantiates a new PartialUpdateNetworkPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialUpdateNetworkPayload() *PartialUpdateNetworkPayload {
	this := PartialUpdateNetworkPayload{}
	return &this
}

// NewPartialUpdateNetworkPayloadWithDefaults instantiates a new PartialUpdateNetworkPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialUpdateNetworkPayloadWithDefaults() *PartialUpdateNetworkPayload {
	this := PartialUpdateNetworkPayload{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *PartialUpdateNetworkPayload) GetAddressFamily() (res PartialUpdateNetworkPayloadGetAddressFamilyRetType) {
	res, _ = o.GetAddressFamilyOk()
	return
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateNetworkPayload) GetAddressFamilyOk() (ret PartialUpdateNetworkPayloadGetAddressFamilyRetType, ok bool) {
	return getPartialUpdateNetworkPayloadGetAddressFamilyAttributeTypeOk(o.AddressFamily)
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *PartialUpdateNetworkPayload) HasAddressFamily() bool {
	_, ok := o.GetAddressFamilyOk()
	return ok
}

// SetAddressFamily gets a reference to the given UpdateNetworkAddressFamily and assigns it to the AddressFamily field.
func (o *PartialUpdateNetworkPayload) SetAddressFamily(v PartialUpdateNetworkPayloadGetAddressFamilyRetType) {
	setPartialUpdateNetworkPayloadGetAddressFamilyAttributeType(&o.AddressFamily, v)
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PartialUpdateNetworkPayload) GetLabels() (res PartialUpdateNetworkPayloadGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateNetworkPayload) GetLabelsOk() (ret PartialUpdateNetworkPayloadGetLabelsRetType, ok bool) {
	return getPartialUpdateNetworkPayloadGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *PartialUpdateNetworkPayload) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *PartialUpdateNetworkPayload) SetLabels(v PartialUpdateNetworkPayloadGetLabelsRetType) {
	setPartialUpdateNetworkPayloadGetLabelsAttributeType(&o.Labels, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartialUpdateNetworkPayload) GetName() (res PartialUpdateNetworkPayloadGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateNetworkPayload) GetNameOk() (ret PartialUpdateNetworkPayloadGetNameRetType, ok bool) {
	return getPartialUpdateNetworkPayloadGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *PartialUpdateNetworkPayload) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartialUpdateNetworkPayload) SetName(v PartialUpdateNetworkPayloadGetNameRetType) {
	setPartialUpdateNetworkPayloadGetNameAttributeType(&o.Name, v)
}

// GetRouted returns the Routed field value if set, zero value otherwise.
func (o *PartialUpdateNetworkPayload) GetRouted() (res PartialUpdateNetworkPayloadgetRoutedRetType) {
	res, _ = o.GetRoutedOk()
	return
}

// GetRoutedOk returns a tuple with the Routed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateNetworkPayload) GetRoutedOk() (ret PartialUpdateNetworkPayloadgetRoutedRetType, ok bool) {
	return getPartialUpdateNetworkPayloadgetRoutedAttributeTypeOk(o.Routed)
}

// HasRouted returns a boolean if a field has been set.
func (o *PartialUpdateNetworkPayload) HasRouted() bool {
	_, ok := o.GetRoutedOk()
	return ok
}

// SetRouted gets a reference to the given bool and assigns it to the Routed field.
func (o *PartialUpdateNetworkPayload) SetRouted(v PartialUpdateNetworkPayloadgetRoutedRetType) {
	setPartialUpdateNetworkPayloadgetRoutedAttributeType(&o.Routed, v)
}

func (o PartialUpdateNetworkPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getPartialUpdateNetworkPayloadGetAddressFamilyAttributeTypeOk(o.AddressFamily); ok {
		toSerialize["AddressFamily"] = val
	}
	if val, ok := getPartialUpdateNetworkPayloadGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	if val, ok := getPartialUpdateNetworkPayloadGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getPartialUpdateNetworkPayloadgetRoutedAttributeTypeOk(o.Routed); ok {
		toSerialize["Routed"] = val
	}
	return toSerialize, nil
}

type NullablePartialUpdateNetworkPayload struct {
	value *PartialUpdateNetworkPayload
	isSet bool
}

func (v NullablePartialUpdateNetworkPayload) Get() *PartialUpdateNetworkPayload {
	return v.value
}

func (v *NullablePartialUpdateNetworkPayload) Set(val *PartialUpdateNetworkPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialUpdateNetworkPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialUpdateNetworkPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialUpdateNetworkPayload(val *PartialUpdateNetworkPayload) *NullablePartialUpdateNetworkPayload {
	return &NullablePartialUpdateNetworkPayload{value: val, isSet: true}
}

func (v NullablePartialUpdateNetworkPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialUpdateNetworkPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
