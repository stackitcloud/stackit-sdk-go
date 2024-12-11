/*
STACKIT Redis API

The STACKIT Redis API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redis

import (
	"encoding/json"
)

// checks if the Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instance{}

// Instance struct for Instance
type Instance struct {
	// REQUIRED
	CfGuid *string `json:"cfGuid"`
	// REQUIRED
	CfOrganizationGuid *string `json:"cfOrganizationGuid"`
	// REQUIRED
	CfSpaceGuid *string `json:"cfSpaceGuid"`
	// REQUIRED
	DashboardUrl *string `json:"dashboardUrl"`
	// REQUIRED
	ImageUrl   *string `json:"imageUrl"`
	InstanceId *string `json:"instanceId,omitempty"`
	// REQUIRED
	LastOperation *InstanceLastOperation `json:"lastOperation"`
	// REQUIRED
	Name *string `json:"name"`
	// Deprecated: Check the GitHub changelog for alternatives
	// REQUIRED
	OfferingName *string `json:"offeringName"`
	// REQUIRED
	OfferingVersion *string `json:"offeringVersion"`
	// REQUIRED
	Parameters *map[string]interface{} `json:"parameters"`
	// REQUIRED
	PlanId *string `json:"planId"`
	// REQUIRED
	PlanName *string `json:"planName"`
	Status   *string `json:"status,omitempty"`
}

type _Instance Instance

// NewInstance instantiates a new Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstance(cfGuid *string, cfOrganizationGuid *string, cfSpaceGuid *string, dashboardUrl *string, imageUrl *string, lastOperation *InstanceLastOperation, name *string, offeringName *string, offeringVersion *string, parameters *map[string]interface{}, planId *string, planName *string) *Instance {
	this := Instance{}
	this.CfGuid = cfGuid
	this.CfOrganizationGuid = cfOrganizationGuid
	this.CfSpaceGuid = cfSpaceGuid
	this.DashboardUrl = dashboardUrl
	this.ImageUrl = imageUrl
	this.LastOperation = lastOperation
	this.Name = name
	this.OfferingName = offeringName
	this.OfferingVersion = offeringVersion
	this.Parameters = parameters
	this.PlanId = planId
	this.PlanName = planName
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	return &this
}

// GetCfGuid returns the CfGuid field value
func (o *Instance) GetCfGuid() *string {
	if o == nil || IsNil(o.CfGuid) {
		var ret *string
		return ret
	}

	return o.CfGuid
}

// GetCfGuidOk returns a tuple with the CfGuid field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCfGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CfGuid, true
}

// SetCfGuid sets field value
func (o *Instance) SetCfGuid(v *string) {
	o.CfGuid = v
}

// GetCfOrganizationGuid returns the CfOrganizationGuid field value
func (o *Instance) GetCfOrganizationGuid() *string {
	if o == nil || IsNil(o.CfOrganizationGuid) {
		var ret *string
		return ret
	}

	return o.CfOrganizationGuid
}

// GetCfOrganizationGuidOk returns a tuple with the CfOrganizationGuid field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCfOrganizationGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CfOrganizationGuid, true
}

// SetCfOrganizationGuid sets field value
func (o *Instance) SetCfOrganizationGuid(v *string) {
	o.CfOrganizationGuid = v
}

// GetCfSpaceGuid returns the CfSpaceGuid field value
func (o *Instance) GetCfSpaceGuid() *string {
	if o == nil || IsNil(o.CfSpaceGuid) {
		var ret *string
		return ret
	}

	return o.CfSpaceGuid
}

// GetCfSpaceGuidOk returns a tuple with the CfSpaceGuid field value
// and a boolean to check if the value has been set.
func (o *Instance) GetCfSpaceGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CfSpaceGuid, true
}

// SetCfSpaceGuid sets field value
func (o *Instance) SetCfSpaceGuid(v *string) {
	o.CfSpaceGuid = v
}

// GetDashboardUrl returns the DashboardUrl field value
func (o *Instance) GetDashboardUrl() *string {
	if o == nil || IsNil(o.DashboardUrl) {
		var ret *string
		return ret
	}

	return o.DashboardUrl
}

// GetDashboardUrlOk returns a tuple with the DashboardUrl field value
// and a boolean to check if the value has been set.
func (o *Instance) GetDashboardUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DashboardUrl, true
}

// SetDashboardUrl sets field value
func (o *Instance) SetDashboardUrl(v *string) {
	o.DashboardUrl = v
}

// GetImageUrl returns the ImageUrl field value
func (o *Instance) GetImageUrl() *string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret *string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *Instance) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *Instance) SetImageUrl(v *string) {
	o.ImageUrl = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *Instance) GetInstanceId() *string {
	if o == nil || IsNil(o.InstanceId) {
		var ret *string
		return ret
	}
	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *Instance) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *Instance) SetInstanceId(v *string) {
	o.InstanceId = v
}

// GetLastOperation returns the LastOperation field value
func (o *Instance) GetLastOperation() *InstanceLastOperation {
	if o == nil || IsNil(o.LastOperation) {
		var ret *InstanceLastOperation
		return ret
	}

	return o.LastOperation
}

// GetLastOperationOk returns a tuple with the LastOperation field value
// and a boolean to check if the value has been set.
func (o *Instance) GetLastOperationOk() (*InstanceLastOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastOperation, true
}

// SetLastOperation sets field value
func (o *Instance) SetLastOperation(v *InstanceLastOperation) {
	o.LastOperation = v
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

// GetOfferingName returns the OfferingName field value
// Deprecated
func (o *Instance) GetOfferingName() *string {
	if o == nil || IsNil(o.OfferingName) {
		var ret *string
		return ret
	}

	return o.OfferingName
}

// GetOfferingNameOk returns a tuple with the OfferingName field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *Instance) GetOfferingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfferingName, true
}

// SetOfferingName sets field value
// Deprecated
func (o *Instance) SetOfferingName(v *string) {
	o.OfferingName = v
}

// GetOfferingVersion returns the OfferingVersion field value
func (o *Instance) GetOfferingVersion() *string {
	if o == nil || IsNil(o.OfferingVersion) {
		var ret *string
		return ret
	}

	return o.OfferingVersion
}

// GetOfferingVersionOk returns a tuple with the OfferingVersion field value
// and a boolean to check if the value has been set.
func (o *Instance) GetOfferingVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfferingVersion, true
}

// SetOfferingVersion sets field value
func (o *Instance) SetOfferingVersion(v *string) {
	o.OfferingVersion = v
}

// GetParameters returns the Parameters field value
func (o *Instance) GetParameters() *map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret *map[string]interface{}
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *Instance) GetParametersOk() (*map[string]interface{}, bool) {
	if o == nil {
		return &map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *Instance) SetParameters(v *map[string]interface{}) {
	o.Parameters = v
}

// GetPlanId returns the PlanId field value
func (o *Instance) GetPlanId() *string {
	if o == nil || IsNil(o.PlanId) {
		var ret *string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId, true
}

// SetPlanId sets field value
func (o *Instance) SetPlanId(v *string) {
	o.PlanId = v
}

// GetPlanName returns the PlanName field value
func (o *Instance) GetPlanName() *string {
	if o == nil || IsNil(o.PlanName) {
		var ret *string
		return ret
	}

	return o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanName, true
}

// SetPlanName sets field value
func (o *Instance) SetPlanName(v *string) {
	o.PlanName = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Instance) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Instance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Instance) SetStatus(v *string) {
	o.Status = v
}

func (o Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cfGuid"] = o.CfGuid
	toSerialize["cfOrganizationGuid"] = o.CfOrganizationGuid
	toSerialize["cfSpaceGuid"] = o.CfSpaceGuid
	toSerialize["dashboardUrl"] = o.DashboardUrl
	toSerialize["imageUrl"] = o.ImageUrl
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	toSerialize["lastOperation"] = o.LastOperation
	toSerialize["name"] = o.Name
	toSerialize["offeringName"] = o.OfferingName
	toSerialize["offeringVersion"] = o.OfferingVersion
	toSerialize["parameters"] = o.Parameters
	toSerialize["planId"] = o.PlanId
	toSerialize["planName"] = o.PlanName
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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
