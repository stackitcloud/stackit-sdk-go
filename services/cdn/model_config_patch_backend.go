/*
CDN API

API used to create and manage your CDN distributions.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
	"fmt"
)

// ConfigPatchBackend - struct for ConfigPatchBackend
type ConfigPatchBackend struct {
	HttpBackendPatch *HttpBackendPatch
}

// HttpBackendPatchAsConfigPatchBackend is a convenience function that returns HttpBackendPatch wrapped in ConfigPatchBackend
func HttpBackendPatchAsConfigPatchBackend(v *HttpBackendPatch) ConfigPatchBackend {
	return ConfigPatchBackend{
		HttpBackendPatch: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConfigPatchBackend) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'HttpBackendPatch'
	if jsonDict["type"] == "HttpBackendPatch" {
		// try to unmarshal JSON data into HttpBackendPatch
		err = json.Unmarshal(data, &dst.HttpBackendPatch)
		if err == nil {
			return nil // data stored in dst.HttpBackendPatch, return on the first match
		} else {
			dst.HttpBackendPatch = nil
			return fmt.Errorf("failed to unmarshal ConfigPatchBackend as HttpBackendPatch: %s", err.Error())
		}
	}

	// check if the discriminator value is 'http'
	if jsonDict["type"] == "http" {
		// try to unmarshal JSON data into HttpBackendPatch
		err = json.Unmarshal(data, &dst.HttpBackendPatch)
		if err == nil {
			return nil // data stored in dst.HttpBackendPatch, return on the first match
		} else {
			dst.HttpBackendPatch = nil
			return fmt.Errorf("failed to unmarshal ConfigPatchBackend as HttpBackendPatch: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConfigPatchBackend) MarshalJSON() ([]byte, error) {
	if src.HttpBackendPatch != nil {
		return json.Marshal(&src.HttpBackendPatch)
	}

	return []byte("{}"), nil // no data in oneOf schemas => empty JSON object
}

// Get the actual instance
func (obj *ConfigPatchBackend) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.HttpBackendPatch != nil {
		return obj.HttpBackendPatch
	}

	// all schemas are nil
	return nil
}

type NullableConfigPatchBackend struct {
	value *ConfigPatchBackend
	isSet bool
}

func (v NullableConfigPatchBackend) Get() *ConfigPatchBackend {
	return v.value
}

func (v *NullableConfigPatchBackend) Set(val *ConfigPatchBackend) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigPatchBackend) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigPatchBackend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigPatchBackend(val *ConfigPatchBackend) *NullableConfigPatchBackend {
	return &NullableConfigPatchBackend{value: val, isSet: true}
}

func (v NullableConfigPatchBackend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigPatchBackend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
