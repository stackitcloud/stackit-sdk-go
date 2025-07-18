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

// checks if the AlertRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertRule{}

/*
	types and functions for alert
*/

// isNotNullableString
type AlertRuleGetAlertAttributeType = *string

func getAlertRuleGetAlertAttributeTypeOk(arg AlertRuleGetAlertAttributeType) (ret AlertRuleGetAlertRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertRuleGetAlertAttributeType(arg *AlertRuleGetAlertAttributeType, val AlertRuleGetAlertRetType) {
	*arg = &val
}

type AlertRuleGetAlertArgType = string
type AlertRuleGetAlertRetType = string

/*
	types and functions for annotations
*/

// isContainer
type AlertRuleGetAnnotationsAttributeType = *map[string]string
type AlertRuleGetAnnotationsArgType = map[string]string
type AlertRuleGetAnnotationsRetType = map[string]string

func getAlertRuleGetAnnotationsAttributeTypeOk(arg AlertRuleGetAnnotationsAttributeType) (ret AlertRuleGetAnnotationsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertRuleGetAnnotationsAttributeType(arg *AlertRuleGetAnnotationsAttributeType, val AlertRuleGetAnnotationsRetType) {
	*arg = &val
}

/*
	types and functions for expr
*/

// isNotNullableString
type AlertRuleGetExprAttributeType = *string

func getAlertRuleGetExprAttributeTypeOk(arg AlertRuleGetExprAttributeType) (ret AlertRuleGetExprRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertRuleGetExprAttributeType(arg *AlertRuleGetExprAttributeType, val AlertRuleGetExprRetType) {
	*arg = &val
}

type AlertRuleGetExprArgType = string
type AlertRuleGetExprRetType = string

/*
	types and functions for for
*/

// isNotNullableString
type AlertRuleGetForAttributeType = *string

func getAlertRuleGetForAttributeTypeOk(arg AlertRuleGetForAttributeType) (ret AlertRuleGetForRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertRuleGetForAttributeType(arg *AlertRuleGetForAttributeType, val AlertRuleGetForRetType) {
	*arg = &val
}

type AlertRuleGetForArgType = string
type AlertRuleGetForRetType = string

/*
	types and functions for labels
*/

// isContainer
type AlertRuleGetLabelsAttributeType = *map[string]string
type AlertRuleGetLabelsArgType = map[string]string
type AlertRuleGetLabelsRetType = map[string]string

func getAlertRuleGetLabelsAttributeTypeOk(arg AlertRuleGetLabelsAttributeType) (ret AlertRuleGetLabelsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setAlertRuleGetLabelsAttributeType(arg *AlertRuleGetLabelsAttributeType, val AlertRuleGetLabelsRetType) {
	*arg = &val
}

// AlertRule struct for AlertRule
type AlertRule struct {
	// REQUIRED
	Alert       AlertRuleGetAlertAttributeType       `json:"alert" required:"true"`
	Annotations AlertRuleGetAnnotationsAttributeType `json:"annotations,omitempty"`
	// REQUIRED
	Expr   AlertRuleGetExprAttributeType   `json:"expr" required:"true"`
	For    AlertRuleGetForAttributeType    `json:"for,omitempty"`
	Labels AlertRuleGetLabelsAttributeType `json:"labels,omitempty"`
}

type _AlertRule AlertRule

// NewAlertRule instantiates a new AlertRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertRule(alert AlertRuleGetAlertArgType, expr AlertRuleGetExprArgType) *AlertRule {
	this := AlertRule{}
	setAlertRuleGetAlertAttributeType(&this.Alert, alert)
	setAlertRuleGetExprAttributeType(&this.Expr, expr)
	return &this
}

// NewAlertRuleWithDefaults instantiates a new AlertRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertRuleWithDefaults() *AlertRule {
	this := AlertRule{}
	var for_ string = "0s"
	this.For = &for_
	return &this
}

// GetAlert returns the Alert field value
func (o *AlertRule) GetAlert() (ret AlertRuleGetAlertRetType) {
	ret, _ = o.GetAlertOk()
	return ret
}

// GetAlertOk returns a tuple with the Alert field value
// and a boolean to check if the value has been set.
func (o *AlertRule) GetAlertOk() (ret AlertRuleGetAlertRetType, ok bool) {
	return getAlertRuleGetAlertAttributeTypeOk(o.Alert)
}

// SetAlert sets field value
func (o *AlertRule) SetAlert(v AlertRuleGetAlertRetType) {
	setAlertRuleGetAlertAttributeType(&o.Alert, v)
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *AlertRule) GetAnnotations() (res AlertRuleGetAnnotationsRetType) {
	res, _ = o.GetAnnotationsOk()
	return
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetAnnotationsOk() (ret AlertRuleGetAnnotationsRetType, ok bool) {
	return getAlertRuleGetAnnotationsAttributeTypeOk(o.Annotations)
}

// HasAnnotations returns a boolean if a field has been set.
func (o *AlertRule) HasAnnotations() bool {
	_, ok := o.GetAnnotationsOk()
	return ok
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *AlertRule) SetAnnotations(v AlertRuleGetAnnotationsRetType) {
	setAlertRuleGetAnnotationsAttributeType(&o.Annotations, v)
}

// GetExpr returns the Expr field value
func (o *AlertRule) GetExpr() (ret AlertRuleGetExprRetType) {
	ret, _ = o.GetExprOk()
	return ret
}

// GetExprOk returns a tuple with the Expr field value
// and a boolean to check if the value has been set.
func (o *AlertRule) GetExprOk() (ret AlertRuleGetExprRetType, ok bool) {
	return getAlertRuleGetExprAttributeTypeOk(o.Expr)
}

// SetExpr sets field value
func (o *AlertRule) SetExpr(v AlertRuleGetExprRetType) {
	setAlertRuleGetExprAttributeType(&o.Expr, v)
}

// GetFor returns the For field value if set, zero value otherwise.
func (o *AlertRule) GetFor() (res AlertRuleGetForRetType) {
	res, _ = o.GetForOk()
	return
}

// GetForOk returns a tuple with the For field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetForOk() (ret AlertRuleGetForRetType, ok bool) {
	return getAlertRuleGetForAttributeTypeOk(o.For)
}

// HasFor returns a boolean if a field has been set.
func (o *AlertRule) HasFor() bool {
	_, ok := o.GetForOk()
	return ok
}

// SetFor gets a reference to the given string and assigns it to the For field.
func (o *AlertRule) SetFor(v AlertRuleGetForRetType) {
	setAlertRuleGetForAttributeType(&o.For, v)
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AlertRule) GetLabels() (res AlertRuleGetLabelsRetType) {
	res, _ = o.GetLabelsOk()
	return
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetLabelsOk() (ret AlertRuleGetLabelsRetType, ok bool) {
	return getAlertRuleGetLabelsAttributeTypeOk(o.Labels)
}

// HasLabels returns a boolean if a field has been set.
func (o *AlertRule) HasLabels() bool {
	_, ok := o.GetLabelsOk()
	return ok
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *AlertRule) SetLabels(v AlertRuleGetLabelsRetType) {
	setAlertRuleGetLabelsAttributeType(&o.Labels, v)
}

func (o AlertRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getAlertRuleGetAlertAttributeTypeOk(o.Alert); ok {
		toSerialize["Alert"] = val
	}
	if val, ok := getAlertRuleGetAnnotationsAttributeTypeOk(o.Annotations); ok {
		toSerialize["Annotations"] = val
	}
	if val, ok := getAlertRuleGetExprAttributeTypeOk(o.Expr); ok {
		toSerialize["Expr"] = val
	}
	if val, ok := getAlertRuleGetForAttributeTypeOk(o.For); ok {
		toSerialize["For"] = val
	}
	if val, ok := getAlertRuleGetLabelsAttributeTypeOk(o.Labels); ok {
		toSerialize["Labels"] = val
	}
	return toSerialize, nil
}

type NullableAlertRule struct {
	value *AlertRule
	isSet bool
}

func (v NullableAlertRule) Get() *AlertRule {
	return v.value
}

func (v *NullableAlertRule) Set(val *AlertRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertRule(val *AlertRule) *NullableAlertRule {
	return &NullableAlertRule{value: val, isSet: true}
}

func (v NullableAlertRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
