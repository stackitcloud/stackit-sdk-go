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

// JWK JSON Web Key according to https://datatracker.ietf.org/doc/html/rfc7517#section-4
type JWK struct {
	Alg *string `json:"alg,omitempty"`
	// REQUIRED
	E   *string `json:"e"`
	Kid *string `json:"kid,omitempty"`
	Ks  *string `json:"ks,omitempty"`
	// REQUIRED
	N      *string `json:"n"`
	Ops    *string `json:"ops,omitempty"`
	Use    *string `json:"use,omitempty"`
	X5c    *string `json:"x5c,omitempty"`
	X5t    *string `json:"x5t,omitempty"`
	X5t256 *string `json:"x5t256,omitempty"`
	X5u    *string `json:"x5u,omitempty"`
}

type _JWK JWK

// NewJWK instantiates a new JWK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWK(e *string, n *string) *JWK {
	this := JWK{}
	this.E = e
	this.N = n
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
func (o *JWK) GetAlg() *string {
	if o == nil || IsNil(o.Alg) {
		var ret *string
		return ret
	}
	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetAlgOk() (*string, bool) {
	if o == nil || IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *JWK) HasAlg() bool {
	if o != nil && !IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *JWK) SetAlg(v *string) {
	o.Alg = v
}

// GetE returns the E field value
func (o *JWK) GetE() *string {
	if o == nil || IsNil(o.E) {
		var ret *string
		return ret
	}

	return o.E
}

// GetEOk returns a tuple with the E field value
// and a boolean to check if the value has been set.
func (o *JWK) GetEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.E, true
}

// SetE sets field value
func (o *JWK) SetE(v *string) {
	o.E = v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *JWK) GetKid() *string {
	if o == nil || IsNil(o.Kid) {
		var ret *string
		return ret
	}
	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *JWK) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *JWK) SetKid(v *string) {
	o.Kid = v
}

// GetKs returns the Ks field value if set, zero value otherwise.
func (o *JWK) GetKs() *string {
	if o == nil || IsNil(o.Ks) {
		var ret *string
		return ret
	}
	return o.Ks
}

// GetKsOk returns a tuple with the Ks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetKsOk() (*string, bool) {
	if o == nil || IsNil(o.Ks) {
		return nil, false
	}
	return o.Ks, true
}

// HasKs returns a boolean if a field has been set.
func (o *JWK) HasKs() bool {
	if o != nil && !IsNil(o.Ks) {
		return true
	}

	return false
}

// SetKs gets a reference to the given string and assigns it to the Ks field.
func (o *JWK) SetKs(v *string) {
	o.Ks = v
}

// GetN returns the N field value
func (o *JWK) GetN() *string {
	if o == nil || IsNil(o.N) {
		var ret *string
		return ret
	}

	return o.N
}

// GetNOk returns a tuple with the N field value
// and a boolean to check if the value has been set.
func (o *JWK) GetNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.N, true
}

// SetN sets field value
func (o *JWK) SetN(v *string) {
	o.N = v
}

// GetOps returns the Ops field value if set, zero value otherwise.
func (o *JWK) GetOps() *string {
	if o == nil || IsNil(o.Ops) {
		var ret *string
		return ret
	}
	return o.Ops
}

// GetOpsOk returns a tuple with the Ops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetOpsOk() (*string, bool) {
	if o == nil || IsNil(o.Ops) {
		return nil, false
	}
	return o.Ops, true
}

// HasOps returns a boolean if a field has been set.
func (o *JWK) HasOps() bool {
	if o != nil && !IsNil(o.Ops) {
		return true
	}

	return false
}

// SetOps gets a reference to the given string and assigns it to the Ops field.
func (o *JWK) SetOps(v *string) {
	o.Ops = v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *JWK) GetUse() *string {
	if o == nil || IsNil(o.Use) {
		var ret *string
		return ret
	}
	return o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *JWK) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *JWK) SetUse(v *string) {
	o.Use = v
}

// GetX5c returns the X5c field value if set, zero value otherwise.
func (o *JWK) GetX5c() *string {
	if o == nil || IsNil(o.X5c) {
		var ret *string
		return ret
	}
	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetX5cOk() (*string, bool) {
	if o == nil || IsNil(o.X5c) {
		return nil, false
	}
	return o.X5c, true
}

// HasX5c returns a boolean if a field has been set.
func (o *JWK) HasX5c() bool {
	if o != nil && !IsNil(o.X5c) {
		return true
	}

	return false
}

// SetX5c gets a reference to the given string and assigns it to the X5c field.
func (o *JWK) SetX5c(v *string) {
	o.X5c = v
}

// GetX5t returns the X5t field value if set, zero value otherwise.
func (o *JWK) GetX5t() *string {
	if o == nil || IsNil(o.X5t) {
		var ret *string
		return ret
	}
	return o.X5t
}

// GetX5tOk returns a tuple with the X5t field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetX5tOk() (*string, bool) {
	if o == nil || IsNil(o.X5t) {
		return nil, false
	}
	return o.X5t, true
}

// HasX5t returns a boolean if a field has been set.
func (o *JWK) HasX5t() bool {
	if o != nil && !IsNil(o.X5t) {
		return true
	}

	return false
}

// SetX5t gets a reference to the given string and assigns it to the X5t field.
func (o *JWK) SetX5t(v *string) {
	o.X5t = v
}

// GetX5t256 returns the X5t256 field value if set, zero value otherwise.
func (o *JWK) GetX5t256() *string {
	if o == nil || IsNil(o.X5t256) {
		var ret *string
		return ret
	}
	return o.X5t256
}

// GetX5t256Ok returns a tuple with the X5t256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetX5t256Ok() (*string, bool) {
	if o == nil || IsNil(o.X5t256) {
		return nil, false
	}
	return o.X5t256, true
}

// HasX5t256 returns a boolean if a field has been set.
func (o *JWK) HasX5t256() bool {
	if o != nil && !IsNil(o.X5t256) {
		return true
	}

	return false
}

// SetX5t256 gets a reference to the given string and assigns it to the X5t256 field.
func (o *JWK) SetX5t256(v *string) {
	o.X5t256 = v
}

// GetX5u returns the X5u field value if set, zero value otherwise.
func (o *JWK) GetX5u() *string {
	if o == nil || IsNil(o.X5u) {
		var ret *string
		return ret
	}
	return o.X5u
}

// GetX5uOk returns a tuple with the X5u field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWK) GetX5uOk() (*string, bool) {
	if o == nil || IsNil(o.X5u) {
		return nil, false
	}
	return o.X5u, true
}

// HasX5u returns a boolean if a field has been set.
func (o *JWK) HasX5u() bool {
	if o != nil && !IsNil(o.X5u) {
		return true
	}

	return false
}

// SetX5u gets a reference to the given string and assigns it to the X5u field.
func (o *JWK) SetX5u(v *string) {
	o.X5u = v
}

func (o JWK) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alg) {
		toSerialize["alg"] = o.Alg
	}
	toSerialize["e"] = o.E
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.Ks) {
		toSerialize["ks"] = o.Ks
	}
	toSerialize["n"] = o.N
	if !IsNil(o.Ops) {
		toSerialize["ops"] = o.Ops
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}
	if !IsNil(o.X5c) {
		toSerialize["x5c"] = o.X5c
	}
	if !IsNil(o.X5t) {
		toSerialize["x5t"] = o.X5t
	}
	if !IsNil(o.X5t256) {
		toSerialize["x5t256"] = o.X5t256
	}
	if !IsNil(o.X5u) {
		toSerialize["x5u"] = o.X5u
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
