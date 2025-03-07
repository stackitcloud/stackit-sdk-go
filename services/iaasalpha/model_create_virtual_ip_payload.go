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

// checks if the CreateVirtualIPPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVirtualIPPayload{}

// CreateVirtualIPPayload Object that represents a virtual IP.
type CreateVirtualIPPayload struct {
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Object that represents an IP address.
	Ip *string `json:"ip,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// A list of UUIDs.
	Members *[]string `json:"members,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
	// Universally Unique Identifier (UUID).
	Network *string `json:"network,omitempty"`
	// The state of a resource object. Possible values: `CREATING`, `CREATED`, `DELETING`, `DELETED`, `FAILED`, `UPDATED`, `UPDATING`.
	Status *string `json:"status,omitempty"`
}

// NewCreateVirtualIPPayload instantiates a new CreateVirtualIPPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVirtualIPPayload() *CreateVirtualIPPayload {
	this := CreateVirtualIPPayload{}
	return &this
}

// NewCreateVirtualIPPayloadWithDefaults instantiates a new CreateVirtualIPPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVirtualIPPayloadWithDefaults() *CreateVirtualIPPayload {
	this := CreateVirtualIPPayload{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateVirtualIPPayload) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualIPPayload) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateVirtualIPPayload) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateVirtualIPPayload) SetId(v *string) {
	o.Id = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *CreateVirtualIPPayload) GetIp() *string {
	if o == nil || IsNil(o.Ip) {
		var ret *string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualIPPayload) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CreateVirtualIPPayload) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *CreateVirtualIPPayload) SetIp(v *string) {
	o.Ip = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateVirtualIPPayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualIPPayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateVirtualIPPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *CreateVirtualIPPayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *CreateVirtualIPPayload) GetMembers() *[]string {
	if o == nil || IsNil(o.Members) {
		var ret *[]string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualIPPayload) GetMembersOk() (*[]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *CreateVirtualIPPayload) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *CreateVirtualIPPayload) SetMembers(v *[]string) {
	o.Members = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateVirtualIPPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualIPPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateVirtualIPPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateVirtualIPPayload) SetName(v *string) {
	o.Name = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CreateVirtualIPPayload) GetNetwork() *string {
	if o == nil || IsNil(o.Network) {
		var ret *string
		return ret
	}
	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualIPPayload) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CreateVirtualIPPayload) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *CreateVirtualIPPayload) SetNetwork(v *string) {
	o.Network = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateVirtualIPPayload) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualIPPayload) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateVirtualIPPayload) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateVirtualIPPayload) SetStatus(v *string) {
	o.Status = v
}

func (o CreateVirtualIPPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCreateVirtualIPPayload struct {
	value *CreateVirtualIPPayload
	isSet bool
}

func (v NullableCreateVirtualIPPayload) Get() *CreateVirtualIPPayload {
	return v.value
}

func (v *NullableCreateVirtualIPPayload) Set(val *CreateVirtualIPPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVirtualIPPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVirtualIPPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVirtualIPPayload(val *CreateVirtualIPPayload) *NullableCreateVirtualIPPayload {
	return &NullableCreateVirtualIPPayload{value: val, isSet: true}
}

func (v NullableCreateVirtualIPPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVirtualIPPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
