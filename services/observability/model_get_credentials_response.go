/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the GetCredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCredentialsResponse{}

// GetCredentialsResponse struct for GetCredentialsResponse
type GetCredentialsResponse struct {
	CredentialsInfo *map[string]string `json:"credentialsInfo,omitempty"`
	// REQUIRED
	Id *string `json:"id"`
	// REQUIRED
	Message *string `json:"message"`
	// REQUIRED
	Name *string `json:"name"`
}

type _GetCredentialsResponse GetCredentialsResponse

// NewGetCredentialsResponse instantiates a new GetCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCredentialsResponse(id *string, message *string, name *string) *GetCredentialsResponse {
	this := GetCredentialsResponse{}
	this.Id = id
	this.Message = message
	this.Name = name
	return &this
}

// NewGetCredentialsResponseWithDefaults instantiates a new GetCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCredentialsResponseWithDefaults() *GetCredentialsResponse {
	this := GetCredentialsResponse{}
	return &this
}

// GetCredentialsInfo returns the CredentialsInfo field value if set, zero value otherwise.
func (o *GetCredentialsResponse) GetCredentialsInfo() *map[string]string {
	if o == nil || IsNil(o.CredentialsInfo) {
		var ret *map[string]string
		return ret
	}
	return o.CredentialsInfo
}

// GetCredentialsInfoOk returns a tuple with the CredentialsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCredentialsResponse) GetCredentialsInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CredentialsInfo) {
		return nil, false
	}
	return o.CredentialsInfo, true
}

// HasCredentialsInfo returns a boolean if a field has been set.
func (o *GetCredentialsResponse) HasCredentialsInfo() bool {
	if o != nil && !IsNil(o.CredentialsInfo) {
		return true
	}

	return false
}

// SetCredentialsInfo gets a reference to the given map[string]string and assigns it to the CredentialsInfo field.
func (o *GetCredentialsResponse) SetCredentialsInfo(v *map[string]string) {
	o.CredentialsInfo = v
}

// GetId returns the Id field value
func (o *GetCredentialsResponse) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCredentialsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *GetCredentialsResponse) SetId(v *string) {
	o.Id = v
}

// GetMessage returns the Message field value
func (o *GetCredentialsResponse) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GetCredentialsResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *GetCredentialsResponse) SetMessage(v *string) {
	o.Message = v
}

// GetName returns the Name field value
func (o *GetCredentialsResponse) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetCredentialsResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *GetCredentialsResponse) SetName(v *string) {
	o.Name = v
}

func (o GetCredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredentialsInfo) {
		toSerialize["credentialsInfo"] = o.CredentialsInfo
	}
	toSerialize["id"] = o.Id
	toSerialize["message"] = o.Message
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableGetCredentialsResponse struct {
	value *GetCredentialsResponse
	isSet bool
}

func (v NullableGetCredentialsResponse) Get() *GetCredentialsResponse {
	return v.value
}

func (v *NullableGetCredentialsResponse) Set(val *GetCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCredentialsResponse(val *GetCredentialsResponse) *NullableGetCredentialsResponse {
	return &NullableGetCredentialsResponse{value: val, isSet: true}
}

func (v NullableGetCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
