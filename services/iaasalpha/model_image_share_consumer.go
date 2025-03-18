/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
	"time"
)

// checks if the ImageShareConsumer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageShareConsumer{}

/*
	types and functions for consumerProjectId
*/

// isNotNullableString
type ImageShareConsumerGetConsumerProjectIdAttributeType = *string

func getImageShareConsumerGetConsumerProjectIdAttributeTypeOk(arg ImageShareConsumerGetConsumerProjectIdAttributeType) (ret ImageShareConsumerGetConsumerProjectIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setImageShareConsumerGetConsumerProjectIdAttributeType(arg *ImageShareConsumerGetConsumerProjectIdAttributeType, val ImageShareConsumerGetConsumerProjectIdRetType) {
	*arg = &val
}

type ImageShareConsumerGetConsumerProjectIdArgType = string
type ImageShareConsumerGetConsumerProjectIdRetType = string

/*
	types and functions for createdAt
*/

// isDateTime
type ImageShareConsumerGetCreatedAtAttributeType = *time.Time
type ImageShareConsumerGetCreatedAtArgType = time.Time
type ImageShareConsumerGetCreatedAtRetType = time.Time

func getImageShareConsumerGetCreatedAtAttributeTypeOk(arg ImageShareConsumerGetCreatedAtAttributeType) (ret ImageShareConsumerGetCreatedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setImageShareConsumerGetCreatedAtAttributeType(arg *ImageShareConsumerGetCreatedAtAttributeType, val ImageShareConsumerGetCreatedAtRetType) {
	*arg = &val
}

/*
	types and functions for imageId
*/

// isNotNullableString
type ImageShareConsumerGetImageIdAttributeType = *string

func getImageShareConsumerGetImageIdAttributeTypeOk(arg ImageShareConsumerGetImageIdAttributeType) (ret ImageShareConsumerGetImageIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setImageShareConsumerGetImageIdAttributeType(arg *ImageShareConsumerGetImageIdAttributeType, val ImageShareConsumerGetImageIdRetType) {
	*arg = &val
}

type ImageShareConsumerGetImageIdArgType = string
type ImageShareConsumerGetImageIdRetType = string

/*
	types and functions for updatedAt
*/

// isDateTime
type ImageShareConsumerGetUpdatedAtAttributeType = *time.Time
type ImageShareConsumerGetUpdatedAtArgType = time.Time
type ImageShareConsumerGetUpdatedAtRetType = time.Time

func getImageShareConsumerGetUpdatedAtAttributeTypeOk(arg ImageShareConsumerGetUpdatedAtAttributeType) (ret ImageShareConsumerGetUpdatedAtRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setImageShareConsumerGetUpdatedAtAttributeType(arg *ImageShareConsumerGetUpdatedAtAttributeType, val ImageShareConsumerGetUpdatedAtRetType) {
	*arg = &val
}

// ImageShareConsumer The details of an Image share consumer.
type ImageShareConsumer struct {
	// Universally Unique Identifier (UUID).
	ConsumerProjectId ImageShareConsumerGetConsumerProjectIdAttributeType `json:"consumerProjectId,omitempty"`
	// Date-time when resource was created.
	CreatedAt ImageShareConsumerGetCreatedAtAttributeType `json:"createdAt,omitempty"`
	// Universally Unique Identifier (UUID).
	ImageId ImageShareConsumerGetImageIdAttributeType `json:"imageId,omitempty"`
	// Date-time when resource was last updated.
	UpdatedAt ImageShareConsumerGetUpdatedAtAttributeType `json:"updatedAt,omitempty"`
}

// NewImageShareConsumer instantiates a new ImageShareConsumer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageShareConsumer() *ImageShareConsumer {
	this := ImageShareConsumer{}
	return &this
}

// NewImageShareConsumerWithDefaults instantiates a new ImageShareConsumer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageShareConsumerWithDefaults() *ImageShareConsumer {
	this := ImageShareConsumer{}
	return &this
}

// GetConsumerProjectId returns the ConsumerProjectId field value if set, zero value otherwise.
func (o *ImageShareConsumer) GetConsumerProjectId() (res ImageShareConsumerGetConsumerProjectIdRetType) {
	res, _ = o.GetConsumerProjectIdOk()
	return
}

// GetConsumerProjectIdOk returns a tuple with the ConsumerProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageShareConsumer) GetConsumerProjectIdOk() (ret ImageShareConsumerGetConsumerProjectIdRetType, ok bool) {
	return getImageShareConsumerGetConsumerProjectIdAttributeTypeOk(o.ConsumerProjectId)
}

// HasConsumerProjectId returns a boolean if a field has been set.
func (o *ImageShareConsumer) HasConsumerProjectId() bool {
	_, ok := o.GetConsumerProjectIdOk()
	return ok
}

// SetConsumerProjectId gets a reference to the given string and assigns it to the ConsumerProjectId field.
func (o *ImageShareConsumer) SetConsumerProjectId(v ImageShareConsumerGetConsumerProjectIdRetType) {
	setImageShareConsumerGetConsumerProjectIdAttributeType(&o.ConsumerProjectId, v)
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ImageShareConsumer) GetCreatedAt() (res ImageShareConsumerGetCreatedAtRetType) {
	res, _ = o.GetCreatedAtOk()
	return
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageShareConsumer) GetCreatedAtOk() (ret ImageShareConsumerGetCreatedAtRetType, ok bool) {
	return getImageShareConsumerGetCreatedAtAttributeTypeOk(o.CreatedAt)
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ImageShareConsumer) HasCreatedAt() bool {
	_, ok := o.GetCreatedAtOk()
	return ok
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ImageShareConsumer) SetCreatedAt(v ImageShareConsumerGetCreatedAtRetType) {
	setImageShareConsumerGetCreatedAtAttributeType(&o.CreatedAt, v)
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ImageShareConsumer) GetImageId() (res ImageShareConsumerGetImageIdRetType) {
	res, _ = o.GetImageIdOk()
	return
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageShareConsumer) GetImageIdOk() (ret ImageShareConsumerGetImageIdRetType, ok bool) {
	return getImageShareConsumerGetImageIdAttributeTypeOk(o.ImageId)
}

// HasImageId returns a boolean if a field has been set.
func (o *ImageShareConsumer) HasImageId() bool {
	_, ok := o.GetImageIdOk()
	return ok
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *ImageShareConsumer) SetImageId(v ImageShareConsumerGetImageIdRetType) {
	setImageShareConsumerGetImageIdAttributeType(&o.ImageId, v)
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ImageShareConsumer) GetUpdatedAt() (res ImageShareConsumerGetUpdatedAtRetType) {
	res, _ = o.GetUpdatedAtOk()
	return
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageShareConsumer) GetUpdatedAtOk() (ret ImageShareConsumerGetUpdatedAtRetType, ok bool) {
	return getImageShareConsumerGetUpdatedAtAttributeTypeOk(o.UpdatedAt)
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ImageShareConsumer) HasUpdatedAt() bool {
	_, ok := o.GetUpdatedAtOk()
	return ok
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ImageShareConsumer) SetUpdatedAt(v ImageShareConsumerGetUpdatedAtRetType) {
	setImageShareConsumerGetUpdatedAtAttributeType(&o.UpdatedAt, v)
}

func (o ImageShareConsumer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getImageShareConsumerGetConsumerProjectIdAttributeTypeOk(o.ConsumerProjectId); ok {
		toSerialize["ConsumerProjectId"] = val
	}
	if val, ok := getImageShareConsumerGetCreatedAtAttributeTypeOk(o.CreatedAt); ok {
		toSerialize["CreatedAt"] = val
	}
	if val, ok := getImageShareConsumerGetImageIdAttributeTypeOk(o.ImageId); ok {
		toSerialize["ImageId"] = val
	}
	if val, ok := getImageShareConsumerGetUpdatedAtAttributeTypeOk(o.UpdatedAt); ok {
		toSerialize["UpdatedAt"] = val
	}
	return toSerialize, nil
}

type NullableImageShareConsumer struct {
	value *ImageShareConsumer
	isSet bool
}

func (v NullableImageShareConsumer) Get() *ImageShareConsumer {
	return v.value
}

func (v *NullableImageShareConsumer) Set(val *ImageShareConsumer) {
	v.value = val
	v.isSet = true
}

func (v NullableImageShareConsumer) IsSet() bool {
	return v.isSet
}

func (v *NullableImageShareConsumer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageShareConsumer(val *ImageShareConsumer) *NullableImageShareConsumer {
	return &NullableImageShareConsumer{value: val, isSet: true}
}

func (v NullableImageShareConsumer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageShareConsumer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
