/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lbapplication

import (
	"encoding/json"
)

// checks if the PlanDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanDetails{}

// PlanDetails struct for PlanDetails
type PlanDetails struct {
	// Description
	Description *string `json:"description,omitempty"`
	// Flavor Name
	FlavorName *string `json:"flavorName,omitempty"`
	// Maximum number of concurrent connections per application load balancer VM instance.
	MaxConnections *interface{} `json:"maxConnections,omitempty"`
	// Service Plan Name
	Name *string `json:"name,omitempty"`
	// Service Plan Identifier
	PlanId *string `json:"planId,omitempty"`
}

// NewPlanDetails instantiates a new PlanDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanDetails() *PlanDetails {
	this := PlanDetails{}
	return &this
}

// NewPlanDetailsWithDefaults instantiates a new PlanDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanDetailsWithDefaults() *PlanDetails {
	this := PlanDetails{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PlanDetails) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PlanDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PlanDetails) SetDescription(v *string) {
	o.Description = v
}

// GetFlavorName returns the FlavorName field value if set, zero value otherwise.
func (o *PlanDetails) GetFlavorName() *string {
	if o == nil || IsNil(o.FlavorName) {
		var ret *string
		return ret
	}
	return o.FlavorName
}

// GetFlavorNameOk returns a tuple with the FlavorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDetails) GetFlavorNameOk() (*string, bool) {
	if o == nil || IsNil(o.FlavorName) {
		return nil, false
	}
	return o.FlavorName, true
}

// HasFlavorName returns a boolean if a field has been set.
func (o *PlanDetails) HasFlavorName() bool {
	if o != nil && !IsNil(o.FlavorName) {
		return true
	}

	return false
}

// SetFlavorName gets a reference to the given string and assigns it to the FlavorName field.
func (o *PlanDetails) SetFlavorName(v *string) {
	o.FlavorName = v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlanDetails) GetMaxConnections() *interface{} {
	if o == nil || IsNil(o.MaxConnections) {
		var ret *interface{}
		return ret
	}
	return o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanDetails) GetMaxConnectionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *PlanDetails) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given interface{} and assigns it to the MaxConnections field.
func (o *PlanDetails) SetMaxConnections(v *interface{}) {
	o.MaxConnections = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlanDetails) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlanDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlanDetails) SetName(v *string) {
	o.Name = v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *PlanDetails) GetPlanId() *string {
	if o == nil || IsNil(o.PlanId) {
		var ret *string
		return ret
	}
	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDetails) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *PlanDetails) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *PlanDetails) SetPlanId(v *string) {
	o.PlanId = v
}

func (o PlanDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FlavorName) {
		toSerialize["flavorName"] = o.FlavorName
	}
	if o.MaxConnections != nil {
		toSerialize["maxConnections"] = o.MaxConnections
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	return toSerialize, nil
}

type NullablePlanDetails struct {
	value *PlanDetails
	isSet bool
}

func (v NullablePlanDetails) Get() *PlanDetails {
	return v.value
}

func (v *NullablePlanDetails) Set(val *PlanDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanDetails(val *PlanDetails) *NullablePlanDetails {
	return &NullablePlanDetails{value: val, isSet: true}
}

func (v NullablePlanDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}