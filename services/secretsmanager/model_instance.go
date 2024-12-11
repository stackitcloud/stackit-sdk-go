/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

import (
	"encoding/json"
)

// checks if the Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instance{}

// Instance struct for Instance
type Instance struct {
	// The API endpoint for connecting to the secrets engine.
	// REQUIRED
	ApiUrl *string `json:"apiUrl"`
	// The date and time the creation of the Secrets Manager instance was finished.
	CreationFinishedDate *string `json:"creationFinishedDate,omitempty"`
	// The date and time the creation of the Secrets Manager instance was triggered.
	// REQUIRED
	CreationStartDate *string `json:"creationStartDate"`
	// A auto generated unique id which identifies the secrets manager instances.
	// REQUIRED
	Id *string `json:"id"`
	// A user chosen name to distinguish multiple secrets manager instances.
	// REQUIRED
	Name *string `json:"name"`
	// The number of secrets currently stored inside of the instance. This value will be updated once per hour.
	// REQUIRED
	SecretCount *int64 `json:"secretCount"`
	// The name of the secrets engine.
	// REQUIRED
	SecretsEngine *string `json:"secretsEngine"`
	// The current state of the Secrets Manager instance.
	// REQUIRED
	State              *string `json:"state"`
	UpdateFinishedDate *string `json:"updateFinishedDate,omitempty"`
	UpdateStartDate    *string `json:"updateStartDate,omitempty"`
}

type _Instance Instance

// NewInstance instantiates a new Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstance(apiUrl *string, creationStartDate *string, id *string, name *string, secretCount *int64, secretsEngine *string, state *string) *Instance {
	this := Instance{}
	this.ApiUrl = apiUrl
	this.CreationStartDate = creationStartDate
	this.Id = id
	this.Name = name
	this.SecretCount = secretCount
	this.SecretsEngine = secretsEngine
	this.State = state
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	return &this
}

// GetApiUrl returns the ApiUrl field value
func (o *Instance) GetApiUrl() *string {
	if o == nil || IsNil(o.ApiUrl) {
		var ret *string
		return ret
	}

	return o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value
// and a boolean to check if the value has been set.
func (o *Instance) GetApiUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// SetApiUrl sets field value
func (o *Instance) SetApiUrl(v *string) {
	o.ApiUrl = v
}

// GetCreationFinishedDate returns the CreationFinishedDate field value if set, zero value otherwise.
func (o *Instance) GetCreationFinishedDate() *string {
	if o == nil || IsNil(o.CreationFinishedDate) {
		var ret *string
		return ret
	}
	return o.CreationFinishedDate
}

// GetCreationFinishedDateOk returns a tuple with the CreationFinishedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetCreationFinishedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreationFinishedDate) {
		return nil, false
	}
	return o.CreationFinishedDate, true
}

// HasCreationFinishedDate returns a boolean if a field has been set.
func (o *Instance) HasCreationFinishedDate() bool {
	if o != nil && !IsNil(o.CreationFinishedDate) && !IsNil(o.CreationFinishedDate) {
		return true
	}

	return false
}

// SetCreationFinishedDate gets a reference to the given string and assigns it to the CreationFinishedDate field.
func (o *Instance) SetCreationFinishedDate(v *string) {
	o.CreationFinishedDate = v
}

// GetCreationStartDate returns the CreationStartDate field value
func (o *Instance) GetCreationStartDate() *string {
	if o == nil || IsNil(o.CreationStartDate) {
		var ret *string
		return ret
	}

	return o.CreationStartDate
}

// GetCreationStartDateOk returns a tuple with the CreationStartDate field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCreationStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationStartDate, true
}

// SetCreationStartDate sets field value
func (o *Instance) SetCreationStartDate(v *string) {
	o.CreationStartDate = v
}

// GetId returns the Id field value
func (o *Instance) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Instance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *Instance) SetId(v *string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Instance) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Instance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Instance) SetName(v *string) {
	o.Name = v
}

// GetSecretCount returns the SecretCount field value
func (o *Instance) GetSecretCount() *int64 {
	if o == nil || IsNil(o.SecretCount) {
		var ret *int64
		return ret
	}

	return o.SecretCount
}

// GetSecretCountOk returns a tuple with the SecretCount field value
// and a boolean to check if the value has been set.
func (o *Instance) GetSecretCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretCount, true
}

// SetSecretCount sets field value
func (o *Instance) SetSecretCount(v *int64) {
	o.SecretCount = v
}

// GetSecretsEngine returns the SecretsEngine field value
func (o *Instance) GetSecretsEngine() *string {
	if o == nil || IsNil(o.SecretsEngine) {
		var ret *string
		return ret
	}

	return o.SecretsEngine
}

// GetSecretsEngineOk returns a tuple with the SecretsEngine field value
// and a boolean to check if the value has been set.
func (o *Instance) GetSecretsEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsEngine, true
}

// SetSecretsEngine sets field value
func (o *Instance) SetSecretsEngine(v *string) {
	o.SecretsEngine = v
}

// GetState returns the State field value
func (o *Instance) GetState() *string {
	if o == nil || IsNil(o.State) {
		var ret *string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Instance) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State, true
}

// SetState sets field value
func (o *Instance) SetState(v *string) {
	o.State = v
}

// GetUpdateFinishedDate returns the UpdateFinishedDate field value if set, zero value otherwise.
func (o *Instance) GetUpdateFinishedDate() *string {
	if o == nil || IsNil(o.UpdateFinishedDate) {
		var ret *string
		return ret
	}
	return o.UpdateFinishedDate
}

// GetUpdateFinishedDateOk returns a tuple with the UpdateFinishedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetUpdateFinishedDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateFinishedDate) {
		return nil, false
	}
	return o.UpdateFinishedDate, true
}

// HasUpdateFinishedDate returns a boolean if a field has been set.
func (o *Instance) HasUpdateFinishedDate() bool {
	if o != nil && !IsNil(o.UpdateFinishedDate) && !IsNil(o.UpdateFinishedDate) {
		return true
	}

	return false
}

// SetUpdateFinishedDate gets a reference to the given string and assigns it to the UpdateFinishedDate field.
func (o *Instance) SetUpdateFinishedDate(v *string) {
	o.UpdateFinishedDate = v
}

// GetUpdateStartDate returns the UpdateStartDate field value if set, zero value otherwise.
func (o *Instance) GetUpdateStartDate() *string {
	if o == nil || IsNil(o.UpdateStartDate) {
		var ret *string
		return ret
	}
	return o.UpdateStartDate
}

// GetUpdateStartDateOk returns a tuple with the UpdateStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetUpdateStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateStartDate) {
		return nil, false
	}
	return o.UpdateStartDate, true
}

// HasUpdateStartDate returns a boolean if a field has been set.
func (o *Instance) HasUpdateStartDate() bool {
	if o != nil && !IsNil(o.UpdateStartDate) && !IsNil(o.UpdateStartDate) {
		return true
	}

	return false
}

// SetUpdateStartDate gets a reference to the given string and assigns it to the UpdateStartDate field.
func (o *Instance) SetUpdateStartDate(v *string) {
	o.UpdateStartDate = v
}

func (o Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiUrl"] = o.ApiUrl
	if !IsNil(o.CreationFinishedDate) {
		toSerialize["creationFinishedDate"] = o.CreationFinishedDate
	}
	toSerialize["creationStartDate"] = o.CreationStartDate
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["secretCount"] = o.SecretCount
	toSerialize["secretsEngine"] = o.SecretsEngine
	toSerialize["state"] = o.State
	if !IsNil(o.UpdateFinishedDate) {
		toSerialize["updateFinishedDate"] = o.UpdateFinishedDate
	}
	if !IsNil(o.UpdateStartDate) {
		toSerialize["updateStartDate"] = o.UpdateStartDate
	}
	return toSerialize, nil
}

type NullableInstance struct {
	value *Instance
	isSet bool
}

func (v NullableInstance) Get() *Instance {
	return v.value
}

func (v *NullableInstance) Set(val *Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstance(val *Instance) *NullableInstance {
	return &NullableInstance{value: val, isSet: true}
}

func (v NullableInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
