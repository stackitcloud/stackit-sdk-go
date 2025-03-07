/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the CreateInstancePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInstancePayload{}

// CreateInstancePayload struct for CreateInstancePayload
type CreateInstancePayload struct {
	// REQUIRED
	Acl *ACL `json:"acl"`
	// REQUIRED
	BackupSchedule *string `json:"backupSchedule"`
	// REQUIRED
	FlavorId *string `json:"flavorId"`
	// Labels field is not certain/clear
	Labels *map[string]string `json:"labels,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Options *map[string]string `json:"options"`
	// Can be cast to int32 without loss of precision.
	// REQUIRED
	Replicas *int64 `json:"replicas"`
	// REQUIRED
	Storage *Storage `json:"storage"`
	// REQUIRED
	Version *string `json:"version"`
}

type _CreateInstancePayload CreateInstancePayload

// NewCreateInstancePayload instantiates a new CreateInstancePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstancePayload(acl *ACL, backupSchedule *string, flavorId *string, name *string, options *map[string]string, replicas *int64, storage *Storage, version *string) *CreateInstancePayload {
	this := CreateInstancePayload{}
	this.Acl = acl
	this.BackupSchedule = backupSchedule
	this.FlavorId = flavorId
	this.Name = name
	this.Options = options
	this.Replicas = replicas
	this.Storage = storage
	this.Version = version
	return &this
}

// NewCreateInstancePayloadWithDefaults instantiates a new CreateInstancePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstancePayloadWithDefaults() *CreateInstancePayload {
	this := CreateInstancePayload{}
	return &this
}

// GetAcl returns the Acl field value
func (o *CreateInstancePayload) GetAcl() *ACL {
	if o == nil || IsNil(o.Acl) {
		var ret *ACL
		return ret
	}

	return o.Acl
}

// GetAclOk returns a tuple with the Acl field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetAclOk() (*ACL, bool) {
	if o == nil {
		return nil, false
	}
	return o.Acl, true
}

// SetAcl sets field value
func (o *CreateInstancePayload) SetAcl(v *ACL) {
	o.Acl = v
}

// GetBackupSchedule returns the BackupSchedule field value
func (o *CreateInstancePayload) GetBackupSchedule() *string {
	if o == nil || IsNil(o.BackupSchedule) {
		var ret *string
		return ret
	}

	return o.BackupSchedule
}

// GetBackupScheduleOk returns a tuple with the BackupSchedule field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetBackupScheduleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupSchedule, true
}

// SetBackupSchedule sets field value
func (o *CreateInstancePayload) SetBackupSchedule(v *string) {
	o.BackupSchedule = v
}

// GetFlavorId returns the FlavorId field value
func (o *CreateInstancePayload) GetFlavorId() *string {
	if o == nil || IsNil(o.FlavorId) {
		var ret *string
		return ret
	}

	return o.FlavorId
}

// GetFlavorIdOk returns a tuple with the FlavorId field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetFlavorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlavorId, true
}

// SetFlavorId sets field value
func (o *CreateInstancePayload) SetFlavorId(v *string) {
	o.FlavorId = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateInstancePayload) GetLabels() *map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateInstancePayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *CreateInstancePayload) SetLabels(v *map[string]string) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *CreateInstancePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CreateInstancePayload) SetName(v *string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *CreateInstancePayload) GetOptions() *map[string]string {
	if o == nil || IsNil(o.Options) {
		var ret *map[string]string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetOptionsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *CreateInstancePayload) SetOptions(v *map[string]string) {
	o.Options = v
}

// GetReplicas returns the Replicas field value
func (o *CreateInstancePayload) GetReplicas() *int64 {
	if o == nil || IsNil(o.Replicas) {
		var ret *int64
		return ret
	}

	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetReplicasOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replicas, true
}

// SetReplicas sets field value
func (o *CreateInstancePayload) SetReplicas(v *int64) {
	o.Replicas = v
}

// GetStorage returns the Storage field value
func (o *CreateInstancePayload) GetStorage() *Storage {
	if o == nil || IsNil(o.Storage) {
		var ret *Storage
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetStorageOk() (*Storage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage, true
}

// SetStorage sets field value
func (o *CreateInstancePayload) SetStorage(v *Storage) {
	o.Storage = v
}

// GetVersion returns the Version field value
func (o *CreateInstancePayload) GetVersion() *string {
	if o == nil || IsNil(o.Version) {
		var ret *string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version, true
}

// SetVersion sets field value
func (o *CreateInstancePayload) SetVersion(v *string) {
	o.Version = v
}

func (o CreateInstancePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acl"] = o.Acl
	toSerialize["backupSchedule"] = o.BackupSchedule
	toSerialize["flavorId"] = o.FlavorId
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["replicas"] = o.Replicas
	toSerialize["storage"] = o.Storage
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableCreateInstancePayload struct {
	value *CreateInstancePayload
	isSet bool
}

func (v NullableCreateInstancePayload) Get() *CreateInstancePayload {
	return v.value
}

func (v *NullableCreateInstancePayload) Set(val *CreateInstancePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInstancePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInstancePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInstancePayload(val *CreateInstancePayload) *NullableCreateInstancePayload {
	return &NullableCreateInstancePayload{value: val, isSet: true}
}

func (v NullableCreateInstancePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInstancePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
