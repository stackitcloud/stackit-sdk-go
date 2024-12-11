/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

import (
	"encoding/json"
)

// checks if the PartialUpdateInstancePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialUpdateInstancePayload{}

// PartialUpdateInstancePayload struct for PartialUpdateInstancePayload
type PartialUpdateInstancePayload struct {
	Acl            *ACL    `json:"acl,omitempty"`
	BackupSchedule *string `json:"backupSchedule,omitempty"`
	FlavorId       *string `json:"flavorId,omitempty"`
	// Labels field is not certain/clear
	Labels   *map[string]string `json:"labels,omitempty"`
	Name     *string            `json:"name,omitempty"`
	Options  *map[string]string `json:"options,omitempty"`
	Replicas *int64             `json:"replicas,omitempty"`
	Storage  *Storage           `json:"storage,omitempty"`
	Version  *string            `json:"version,omitempty"`
}

// NewPartialUpdateInstancePayload instantiates a new PartialUpdateInstancePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialUpdateInstancePayload() *PartialUpdateInstancePayload {
	this := PartialUpdateInstancePayload{}
	return &this
}

// NewPartialUpdateInstancePayloadWithDefaults instantiates a new PartialUpdateInstancePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialUpdateInstancePayloadWithDefaults() *PartialUpdateInstancePayload {
	this := PartialUpdateInstancePayload{}
	return &this
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *PartialUpdateInstancePayload) GetAcl() *ACL {
	if o == nil || IsNil(o.Acl) {
		var ret *ACL
		return ret
	}
	return o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstancePayload) GetAclOk() (*ACL, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *PartialUpdateInstancePayload) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given ACL and assigns it to the Acl field.
func (o *PartialUpdateInstancePayload) SetAcl(v *ACL) {
	o.Acl = v
}

// GetBackupSchedule returns the BackupSchedule field value if set, zero value otherwise.
func (o *PartialUpdateInstancePayload) GetBackupSchedule() *string {
	if o == nil || IsNil(o.BackupSchedule) {
		var ret *string
		return ret
	}
	return o.BackupSchedule
}

// GetBackupScheduleOk returns a tuple with the BackupSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstancePayload) GetBackupScheduleOk() (*string, bool) {
	if o == nil || IsNil(o.BackupSchedule) {
		return nil, false
	}
	return o.BackupSchedule, true
}

// HasBackupSchedule returns a boolean if a field has been set.
func (o *PartialUpdateInstancePayload) HasBackupSchedule() bool {
	if o != nil && !IsNil(o.BackupSchedule) && !IsNil(o.BackupSchedule) {
		return true
	}

	return false
}

// SetBackupSchedule gets a reference to the given string and assigns it to the BackupSchedule field.
func (o *PartialUpdateInstancePayload) SetBackupSchedule(v *string) {
	o.BackupSchedule = v
}

// GetFlavorId returns the FlavorId field value if set, zero value otherwise.
func (o *PartialUpdateInstancePayload) GetFlavorId() *string {
	if o == nil || IsNil(o.FlavorId) {
		var ret *string
		return ret
	}
	return o.FlavorId
}

// GetFlavorIdOk returns a tuple with the FlavorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstancePayload) GetFlavorIdOk() (*string, bool) {
	if o == nil || IsNil(o.FlavorId) {
		return nil, false
	}
	return o.FlavorId, true
}

// HasFlavorId returns a boolean if a field has been set.
func (o *PartialUpdateInstancePayload) HasFlavorId() bool {
	if o != nil && !IsNil(o.FlavorId) && !IsNil(o.FlavorId) {
		return true
	}

	return false
}

// SetFlavorId gets a reference to the given string and assigns it to the FlavorId field.
func (o *PartialUpdateInstancePayload) SetFlavorId(v *string) {
	o.FlavorId = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PartialUpdateInstancePayload) GetLabels() *map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstancePayload) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PartialUpdateInstancePayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *PartialUpdateInstancePayload) SetLabels(v *map[string]string) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartialUpdateInstancePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstancePayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartialUpdateInstancePayload) HasName() bool {
	if o != nil && !IsNil(o.Name) && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartialUpdateInstancePayload) SetName(v *string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PartialUpdateInstancePayload) GetOptions() *map[string]string {
	if o == nil || IsNil(o.Options) {
		var ret *map[string]string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstancePayload) GetOptionsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PartialUpdateInstancePayload) HasOptions() bool {
	if o != nil && !IsNil(o.Options) && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]string and assigns it to the Options field.
func (o *PartialUpdateInstancePayload) SetOptions(v *map[string]string) {
	o.Options = v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *PartialUpdateInstancePayload) GetReplicas() *int64 {
	if o == nil || IsNil(o.Replicas) {
		var ret *int64
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstancePayload) GetReplicasOk() (*int64, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *PartialUpdateInstancePayload) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int64 and assigns it to the Replicas field.
func (o *PartialUpdateInstancePayload) SetReplicas(v *int64) {
	o.Replicas = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *PartialUpdateInstancePayload) GetStorage() *Storage {
	if o == nil || IsNil(o.Storage) {
		var ret *Storage
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstancePayload) GetStorageOk() (*Storage, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *PartialUpdateInstancePayload) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given Storage and assigns it to the Storage field.
func (o *PartialUpdateInstancePayload) SetStorage(v *Storage) {
	o.Storage = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PartialUpdateInstancePayload) GetVersion() *string {
	if o == nil || IsNil(o.Version) {
		var ret *string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateInstancePayload) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PartialUpdateInstancePayload) HasVersion() bool {
	if o != nil && !IsNil(o.Version) && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PartialUpdateInstancePayload) SetVersion(v *string) {
	o.Version = v
}

func (o PartialUpdateInstancePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acl) {
		toSerialize["acl"] = o.Acl
	}
	if !IsNil(o.BackupSchedule) {
		toSerialize["backupSchedule"] = o.BackupSchedule
	}
	if !IsNil(o.FlavorId) {
		toSerialize["flavorId"] = o.FlavorId
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullablePartialUpdateInstancePayload struct {
	value *PartialUpdateInstancePayload
	isSet bool
}

func (v NullablePartialUpdateInstancePayload) Get() *PartialUpdateInstancePayload {
	return v.value
}

func (v *NullablePartialUpdateInstancePayload) Set(val *PartialUpdateInstancePayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialUpdateInstancePayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialUpdateInstancePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialUpdateInstancePayload(val *PartialUpdateInstancePayload) *NullablePartialUpdateInstancePayload {
	return &NullablePartialUpdateInstancePayload{value: val, isSet: true}
}

func (v NullablePartialUpdateInstancePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialUpdateInstancePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
