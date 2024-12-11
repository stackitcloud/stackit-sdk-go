/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

import (
	"encoding/json"
)

// checks if the CreateBucketResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBucketResponse{}

// CreateBucketResponse struct for CreateBucketResponse
type CreateBucketResponse struct {
	// Name of the bucket
	// REQUIRED
	Bucket *string `json:"bucket"`
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
}

type _CreateBucketResponse CreateBucketResponse

// NewCreateBucketResponse instantiates a new CreateBucketResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBucketResponse(bucket *string, project *string) *CreateBucketResponse {
	this := CreateBucketResponse{}
	this.Bucket = bucket
	this.Project = project
	return &this
}

// NewCreateBucketResponseWithDefaults instantiates a new CreateBucketResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBucketResponseWithDefaults() *CreateBucketResponse {
	this := CreateBucketResponse{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *CreateBucketResponse) GetBucket() *string {
	if o == nil || IsNil(o.Bucket) {
		var ret *string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *CreateBucketResponse) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bucket, true
}

// SetBucket sets field value
func (o *CreateBucketResponse) SetBucket(v *string) {
	o.Bucket = v
}

// GetProject returns the Project field value
func (o *CreateBucketResponse) GetProject() *string {
	if o == nil || IsNil(o.Project) {
		var ret *string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *CreateBucketResponse) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project, true
}

// SetProject sets field value
func (o *CreateBucketResponse) SetProject(v *string) {
	o.Project = v
}

func (o CreateBucketResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket"] = o.Bucket
	toSerialize["project"] = o.Project
	return toSerialize, nil
}

type NullableCreateBucketResponse struct {
	value *CreateBucketResponse
	isSet bool
}

func (v NullableCreateBucketResponse) Get() *CreateBucketResponse {
	return v.value
}

func (v *NullableCreateBucketResponse) Set(val *CreateBucketResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBucketResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBucketResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBucketResponse(val *CreateBucketResponse) *NullableCreateBucketResponse {
	return &NullableCreateBucketResponse{value: val, isSet: true}
}

func (v NullableCreateBucketResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBucketResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
