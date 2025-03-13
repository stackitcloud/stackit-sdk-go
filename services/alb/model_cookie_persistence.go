/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 2beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alb

import (
	"encoding/json"
)

// checks if the CookiePersistence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CookiePersistence{}

// CookiePersistence struct for CookiePersistence
type CookiePersistence struct {
	// Cookie is the name of the cookie to use.
	Name *string `json:"name,omitempty"`
	// TTL specifies the time-to-live for the cookie. The default value is 0s, and it acts as a session cookie, expiring when the client session ends.
	Ttl *string `json:"ttl,omitempty"`
}

// NewCookiePersistence instantiates a new CookiePersistence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCookiePersistence() *CookiePersistence {
	this := CookiePersistence{}
	return &this
}

// NewCookiePersistenceWithDefaults instantiates a new CookiePersistence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCookiePersistenceWithDefaults() *CookiePersistence {
	this := CookiePersistence{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CookiePersistence) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookiePersistence) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CookiePersistence) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CookiePersistence) SetName(v *string) {
	o.Name = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *CookiePersistence) GetTtl() *string {
	if o == nil || IsNil(o.Ttl) {
		var ret *string
		return ret
	}
	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookiePersistence) GetTtlOk() (*string, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CookiePersistence) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *CookiePersistence) SetTtl(v *string) {
	o.Ttl = v
}

func (o CookiePersistence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	return toSerialize, nil
}

type NullableCookiePersistence struct {
	value *CookiePersistence
	isSet bool
}

func (v NullableCookiePersistence) Get() *CookiePersistence {
	return v.value
}

func (v *NullableCookiePersistence) Set(val *CookiePersistence) {
	v.value = val
	v.isSet = true
}

func (v NullableCookiePersistence) IsSet() bool {
	return v.isSet
}

func (v *NullableCookiePersistence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCookiePersistence(val *CookiePersistence) *NullableCookiePersistence {
	return &NullableCookiePersistence{value: val, isSet: true}
}

func (v NullableCookiePersistence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCookiePersistence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
