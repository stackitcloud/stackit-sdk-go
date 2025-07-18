/*
Service Account API

API to manage Service Accounts and their Access Tokens.  ### System for Cross-domain Identity Management (SCIM) Service Account Service offers SCIM APIs to query state. The SCIM protocol was created as standard for  automating the exchange of user identity information between identity domains, or IT systems. Service accounts  are be handled as indentites similar to SCIM users. A custom SCIM schema has been created: `/ServiceAccounts`  #### Syntax ##### Attribute operators | OPERATOR | DESCRIPTION              | |----------|--------------------------| | eq       | equal                    | | ne       | not equal                | | co       | contains                 | | sw       | starts with              | | ew       | ends with                |  ##### Logical operators | OPERATOR | DESCRIPTION              | |----------|--------------------------| | and      | logical \"and\"            | | or       | logical \"or\"             |  ##### Grouping operators | OPERATOR | DESCRIPTION              | |----------|--------------------------| | ()       | precending grouping      |  ##### Example ``` filter=email eq \"my-service-account-aBc2defg@sa.stackit.cloud\" filter=email ne \"my-service-account-aBc2defg@sa.stackit.cloud\" filter=email co \"my-service-account\" filter=name sw \"my\" filter=name ew \"account\" filter=email co \"my-service-account\" and name sw \"my\" filter=email co \"my-service-account\" and (name sw \"my\" or name ew \"account\") ```  #### Sorting  > Sorting is optional  | PARAMETER | DESCRIPTION                          | |-----------|--------------------------------------| | sortBy    | attribute response is ordered by     | | sortOrder | 'ASCENDING' (default) or 'DESCENDING'|  #### Pagination  | PARAMETER    | DESCRIPTION                                  | |--------------|----------------------------------------------| | startIndex   | index of first query result, default: 1      | | count        | maximum number of query results, default: 100|

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccount

import (
	"encoding/json"
)

// checks if the JWK type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JWK{}

/*
	types and functions for alg
*/

// isNotNullableString
type JWKGetAlgAttributeType = *string

func getJWKGetAlgAttributeTypeOk(arg JWKGetAlgAttributeType) (ret JWKGetAlgRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetAlgAttributeType(arg *JWKGetAlgAttributeType, val JWKGetAlgRetType) {
	*arg = &val
}

type JWKGetAlgArgType = string
type JWKGetAlgRetType = string

/*
	types and functions for e
*/

// isNotNullableString
type JWKGetEAttributeType = *string

func getJWKGetEAttributeTypeOk(arg JWKGetEAttributeType) (ret JWKGetERetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetEAttributeType(arg *JWKGetEAttributeType, val JWKGetERetType) {
	*arg = &val
}

type JWKGetEArgType = string
type JWKGetERetType = string

/*
	types and functions for kid
*/

// isNotNullableString
type JWKGetKidAttributeType = *string

func getJWKGetKidAttributeTypeOk(arg JWKGetKidAttributeType) (ret JWKGetKidRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetKidAttributeType(arg *JWKGetKidAttributeType, val JWKGetKidRetType) {
	*arg = &val
}

type JWKGetKidArgType = string
type JWKGetKidRetType = string

/*
	types and functions for ks
*/

// isNotNullableString
type JWKGetKsAttributeType = *string

func getJWKGetKsAttributeTypeOk(arg JWKGetKsAttributeType) (ret JWKGetKsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetKsAttributeType(arg *JWKGetKsAttributeType, val JWKGetKsRetType) {
	*arg = &val
}

type JWKGetKsArgType = string
type JWKGetKsRetType = string

/*
	types and functions for n
*/

// isNotNullableString
type JWKGetNAttributeType = *string

func getJWKGetNAttributeTypeOk(arg JWKGetNAttributeType) (ret JWKGetNRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetNAttributeType(arg *JWKGetNAttributeType, val JWKGetNRetType) {
	*arg = &val
}

type JWKGetNArgType = string
type JWKGetNRetType = string

/*
	types and functions for ops
*/

// isNotNullableString
type JWKGetOpsAttributeType = *string

func getJWKGetOpsAttributeTypeOk(arg JWKGetOpsAttributeType) (ret JWKGetOpsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetOpsAttributeType(arg *JWKGetOpsAttributeType, val JWKGetOpsRetType) {
	*arg = &val
}

type JWKGetOpsArgType = string
type JWKGetOpsRetType = string

/*
	types and functions for use
*/

// isNotNullableString
type JWKGetUseAttributeType = *string

func getJWKGetUseAttributeTypeOk(arg JWKGetUseAttributeType) (ret JWKGetUseRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetUseAttributeType(arg *JWKGetUseAttributeType, val JWKGetUseRetType) {
	*arg = &val
}

type JWKGetUseArgType = string
type JWKGetUseRetType = string

/*
	types and functions for x5c
*/

// isNotNullableString
type JWKGetX5cAttributeType = *string

func getJWKGetX5cAttributeTypeOk(arg JWKGetX5cAttributeType) (ret JWKGetX5cRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetX5cAttributeType(arg *JWKGetX5cAttributeType, val JWKGetX5cRetType) {
	*arg = &val
}

type JWKGetX5cArgType = string
type JWKGetX5cRetType = string

/*
	types and functions for x5t
*/

// isNotNullableString
type JWKGetX5tAttributeType = *string

func getJWKGetX5tAttributeTypeOk(arg JWKGetX5tAttributeType) (ret JWKGetX5tRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetX5tAttributeType(arg *JWKGetX5tAttributeType, val JWKGetX5tRetType) {
	*arg = &val
}

type JWKGetX5tArgType = string
type JWKGetX5tRetType = string

/*
	types and functions for x5t256
*/

// isNotNullableString
type JWKGetX5t256AttributeType = *string

func getJWKGetX5t256AttributeTypeOk(arg JWKGetX5t256AttributeType) (ret JWKGetX5t256RetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetX5t256AttributeType(arg *JWKGetX5t256AttributeType, val JWKGetX5t256RetType) {
	*arg = &val
}

type JWKGetX5t256ArgType = string
type JWKGetX5t256RetType = string

/*
	types and functions for x5u
*/

// isNotNullableString
type JWKGetX5uAttributeType = *string

func getJWKGetX5uAttributeTypeOk(arg JWKGetX5uAttributeType) (ret JWKGetX5uRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setJWKGetX5uAttributeType(arg *JWKGetX5uAttributeType, val JWKGetX5uRetType) {
	*arg = &val
}

type JWKGetX5uArgType = string
type JWKGetX5uRetType = string

// JWK JSON Web Key according to https://datatracker.ietf.org/doc/html/rfc7517#section-4
type JWK struct {
	Alg JWKGetAlgAttributeType `json:"alg,omitempty"`
	// REQUIRED
	E   JWKGetEAttributeType   `json:"e" required:"true"`
	Kid JWKGetKidAttributeType `json:"kid,omitempty"`
	Ks  JWKGetKsAttributeType  `json:"ks,omitempty"`
	// REQUIRED
	N      JWKGetNAttributeType      `json:"n" required:"true"`
	Ops    JWKGetOpsAttributeType    `json:"ops,omitempty"`
	Use    JWKGetUseAttributeType    `json:"use,omitempty"`
	X5c    JWKGetX5cAttributeType    `json:"x5c,omitempty"`
	X5t    JWKGetX5tAttributeType    `json:"x5t,omitempty"`
	X5t256 JWKGetX5t256AttributeType `json:"x5t256,omitempty"`
	X5u    JWKGetX5uAttributeType    `json:"x5u,omitempty"`
}

type _JWK JWK

// NewJWK instantiates a new JWK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWK(e JWKGetEArgType, n JWKGetNArgType) *JWK {
	this := JWK{}
	setJWKGetEAttributeType(&this.E, e)
	setJWKGetNAttributeType(&this.N, n)
	return &this
}

// NewJWKWithDefaults instantiates a new JWK object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWKWithDefaults() *JWK {
	this := JWK{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *JWK) GetAlg() (res JWKGetAlgRetType) {
	res, _ = o.GetAlgOk()
	return
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetAlgOk() (ret JWKGetAlgRetType, ok bool) {
	return getJWKGetAlgAttributeTypeOk(o.Alg)
}

// HasAlg returns a boolean if a field has been set.
func (o *JWK) HasAlg() bool {
	_, ok := o.GetAlgOk()
	return ok
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *JWK) SetAlg(v JWKGetAlgRetType) {
	setJWKGetAlgAttributeType(&o.Alg, v)
}

// GetE returns the E field value
func (o *JWK) GetE() (ret JWKGetERetType) {
	ret, _ = o.GetEOk()
	return ret
}

// GetEOk returns a tuple with the E field value
// and a boolean to check if the value has been set.
func (o *JWK) GetEOk() (ret JWKGetERetType, ok bool) {
	return getJWKGetEAttributeTypeOk(o.E)
}

// SetE sets field value
func (o *JWK) SetE(v JWKGetERetType) {
	setJWKGetEAttributeType(&o.E, v)
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *JWK) GetKid() (res JWKGetKidRetType) {
	res, _ = o.GetKidOk()
	return
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetKidOk() (ret JWKGetKidRetType, ok bool) {
	return getJWKGetKidAttributeTypeOk(o.Kid)
}

// HasKid returns a boolean if a field has been set.
func (o *JWK) HasKid() bool {
	_, ok := o.GetKidOk()
	return ok
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *JWK) SetKid(v JWKGetKidRetType) {
	setJWKGetKidAttributeType(&o.Kid, v)
}

// GetKs returns the Ks field value if set, zero value otherwise.
func (o *JWK) GetKs() (res JWKGetKsRetType) {
	res, _ = o.GetKsOk()
	return
}

// GetKsOk returns a tuple with the Ks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetKsOk() (ret JWKGetKsRetType, ok bool) {
	return getJWKGetKsAttributeTypeOk(o.Ks)
}

// HasKs returns a boolean if a field has been set.
func (o *JWK) HasKs() bool {
	_, ok := o.GetKsOk()
	return ok
}

// SetKs gets a reference to the given string and assigns it to the Ks field.
func (o *JWK) SetKs(v JWKGetKsRetType) {
	setJWKGetKsAttributeType(&o.Ks, v)
}

// GetN returns the N field value
func (o *JWK) GetN() (ret JWKGetNRetType) {
	ret, _ = o.GetNOk()
	return ret
}

// GetNOk returns a tuple with the N field value
// and a boolean to check if the value has been set.
func (o *JWK) GetNOk() (ret JWKGetNRetType, ok bool) {
	return getJWKGetNAttributeTypeOk(o.N)
}

// SetN sets field value
func (o *JWK) SetN(v JWKGetNRetType) {
	setJWKGetNAttributeType(&o.N, v)
}

// GetOps returns the Ops field value if set, zero value otherwise.
func (o *JWK) GetOps() (res JWKGetOpsRetType) {
	res, _ = o.GetOpsOk()
	return
}

// GetOpsOk returns a tuple with the Ops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetOpsOk() (ret JWKGetOpsRetType, ok bool) {
	return getJWKGetOpsAttributeTypeOk(o.Ops)
}

// HasOps returns a boolean if a field has been set.
func (o *JWK) HasOps() bool {
	_, ok := o.GetOpsOk()
	return ok
}

// SetOps gets a reference to the given string and assigns it to the Ops field.
func (o *JWK) SetOps(v JWKGetOpsRetType) {
	setJWKGetOpsAttributeType(&o.Ops, v)
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *JWK) GetUse() (res JWKGetUseRetType) {
	res, _ = o.GetUseOk()
	return
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetUseOk() (ret JWKGetUseRetType, ok bool) {
	return getJWKGetUseAttributeTypeOk(o.Use)
}

// HasUse returns a boolean if a field has been set.
func (o *JWK) HasUse() bool {
	_, ok := o.GetUseOk()
	return ok
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *JWK) SetUse(v JWKGetUseRetType) {
	setJWKGetUseAttributeType(&o.Use, v)
}

// GetX5c returns the X5c field value if set, zero value otherwise.
func (o *JWK) GetX5c() (res JWKGetX5cRetType) {
	res, _ = o.GetX5cOk()
	return
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetX5cOk() (ret JWKGetX5cRetType, ok bool) {
	return getJWKGetX5cAttributeTypeOk(o.X5c)
}

// HasX5c returns a boolean if a field has been set.
func (o *JWK) HasX5c() bool {
	_, ok := o.GetX5cOk()
	return ok
}

// SetX5c gets a reference to the given string and assigns it to the X5c field.
func (o *JWK) SetX5c(v JWKGetX5cRetType) {
	setJWKGetX5cAttributeType(&o.X5c, v)
}

// GetX5t returns the X5t field value if set, zero value otherwise.
func (o *JWK) GetX5t() (res JWKGetX5tRetType) {
	res, _ = o.GetX5tOk()
	return
}

// GetX5tOk returns a tuple with the X5t field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetX5tOk() (ret JWKGetX5tRetType, ok bool) {
	return getJWKGetX5tAttributeTypeOk(o.X5t)
}

// HasX5t returns a boolean if a field has been set.
func (o *JWK) HasX5t() bool {
	_, ok := o.GetX5tOk()
	return ok
}

// SetX5t gets a reference to the given string and assigns it to the X5t field.
func (o *JWK) SetX5t(v JWKGetX5tRetType) {
	setJWKGetX5tAttributeType(&o.X5t, v)
}

// GetX5t256 returns the X5t256 field value if set, zero value otherwise.
func (o *JWK) GetX5t256() (res JWKGetX5t256RetType) {
	res, _ = o.GetX5t256Ok()
	return
}

// GetX5t256Ok returns a tuple with the X5t256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetX5t256Ok() (ret JWKGetX5t256RetType, ok bool) {
	return getJWKGetX5t256AttributeTypeOk(o.X5t256)
}

// HasX5t256 returns a boolean if a field has been set.
func (o *JWK) HasX5t256() bool {
	_, ok := o.GetX5t256Ok()
	return ok
}

// SetX5t256 gets a reference to the given string and assigns it to the X5t256 field.
func (o *JWK) SetX5t256(v JWKGetX5t256RetType) {
	setJWKGetX5t256AttributeType(&o.X5t256, v)
}

// GetX5u returns the X5u field value if set, zero value otherwise.
func (o *JWK) GetX5u() (res JWKGetX5uRetType) {
	res, _ = o.GetX5uOk()
	return
}

// GetX5uOk returns a tuple with the X5u field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetX5uOk() (ret JWKGetX5uRetType, ok bool) {
	return getJWKGetX5uAttributeTypeOk(o.X5u)
}

// HasX5u returns a boolean if a field has been set.
func (o *JWK) HasX5u() bool {
	_, ok := o.GetX5uOk()
	return ok
}

// SetX5u gets a reference to the given string and assigns it to the X5u field.
func (o *JWK) SetX5u(v JWKGetX5uRetType) {
	setJWKGetX5uAttributeType(&o.X5u, v)
}

func (o JWK) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getJWKGetAlgAttributeTypeOk(o.Alg); ok {
		toSerialize["Alg"] = val
	}
	if val, ok := getJWKGetEAttributeTypeOk(o.E); ok {
		toSerialize["E"] = val
	}
	if val, ok := getJWKGetKidAttributeTypeOk(o.Kid); ok {
		toSerialize["Kid"] = val
	}
	if val, ok := getJWKGetKsAttributeTypeOk(o.Ks); ok {
		toSerialize["Ks"] = val
	}
	if val, ok := getJWKGetNAttributeTypeOk(o.N); ok {
		toSerialize["N"] = val
	}
	if val, ok := getJWKGetOpsAttributeTypeOk(o.Ops); ok {
		toSerialize["Ops"] = val
	}
	if val, ok := getJWKGetUseAttributeTypeOk(o.Use); ok {
		toSerialize["Use"] = val
	}
	if val, ok := getJWKGetX5cAttributeTypeOk(o.X5c); ok {
		toSerialize["X5c"] = val
	}
	if val, ok := getJWKGetX5tAttributeTypeOk(o.X5t); ok {
		toSerialize["X5t"] = val
	}
	if val, ok := getJWKGetX5t256AttributeTypeOk(o.X5t256); ok {
		toSerialize["X5t256"] = val
	}
	if val, ok := getJWKGetX5uAttributeTypeOk(o.X5u); ok {
		toSerialize["X5u"] = val
	}
	return toSerialize, nil
}

type NullableJWK struct {
	value *JWK
	isSet bool
}

func (v NullableJWK) Get() *JWK {
	return v.value
}

func (v *NullableJWK) Set(val *JWK) {
	v.value = val
	v.isSet = true
}

func (v NullableJWK) IsSet() bool {
	return v.isSet
}

func (v *NullableJWK) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJWK(val *JWK) *NullableJWK {
	return &NullableJWK{value: val, isSet: true}
}

func (v NullableJWK) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJWK) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
