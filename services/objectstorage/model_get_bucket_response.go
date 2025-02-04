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

// checks if the GetBucketResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBucketResponse{}

// GetBucketResponse struct for GetBucketResponse
type GetBucketResponse struct {
	// REQUIRED
	Bucket *Bucket `json:"bucket"`
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
}

type _GetBucketResponse GetBucketResponse

// NewGetBucketResponse instantiates a new GetBucketResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBucketResponse(bucket *Bucket, project *string) *GetBucketResponse {
	this := GetBucketResponse{}
	this.Bucket = bucket
	this.Project = project
	return &this
}

// NewGetBucketResponseWithDefaults instantiates a new GetBucketResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBucketResponseWithDefaults() *GetBucketResponse {
	this := GetBucketResponse{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *GetBucketResponse) GetBucket() *Bucket {
	if o == nil || IsNil(o.Bucket) {
		var ret *Bucket
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *GetBucketResponse) GetBucketOk() (*Bucket, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bucket, true
}

// SetBucket sets field value
func (o *GetBucketResponse) SetBucket(v *Bucket) {
	o.Bucket = v
}

// GetProject returns the Project field value
func (o *GetBucketResponse) GetProject() *string {
	if o == nil || IsNil(o.Project) {
		var ret *string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *GetBucketResponse) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project, true
}

// SetProject sets field value
func (o *GetBucketResponse) SetProject(v *string) {
	o.Project = v
}

func (o GetBucketResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket"] = o.Bucket
	toSerialize["project"] = o.Project
	return toSerialize, nil
}

type NullableGetBucketResponse struct {
	value *GetBucketResponse
	isSet bool
}

func (v NullableGetBucketResponse) Get() *GetBucketResponse {
	return v.value
}

func (v *NullableGetBucketResponse) Set(val *GetBucketResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBucketResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBucketResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBucketResponse(val *GetBucketResponse) *NullableGetBucketResponse {
	return &NullableGetBucketResponse{value: val, isSet: true}
}

func (v NullableGetBucketResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBucketResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
