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

// checks if the Credentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Credentials{}

/*
	types and functions for host
*/

// isNotNullableString
type CredentialsGetHostAttributeType = *string

func getCredentialsGetHostAttributeTypeOk(arg CredentialsGetHostAttributeType) (ret CredentialsGetHostRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGetHostAttributeType(arg *CredentialsGetHostAttributeType, val CredentialsGetHostRetType) {
	*arg = &val
}

type CredentialsGetHostArgType = string
type CredentialsGetHostRetType = string

/*
	types and functions for hosts
*/

// isArray
type CredentialsGetHostsAttributeType = *[]string
type CredentialsGetHostsArgType = []string
type CredentialsGetHostsRetType = []string

func getCredentialsGetHostsAttributeTypeOk(arg CredentialsGetHostsAttributeType) (ret CredentialsGetHostsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGetHostsAttributeType(arg *CredentialsGetHostsAttributeType, val CredentialsGetHostsRetType) {
	*arg = &val
}

/*
	types and functions for load_balanced_host
*/

// isNotNullableString
type CredentialsGetLoadBalancedHostAttributeType = *string

func getCredentialsGetLoadBalancedHostAttributeTypeOk(arg CredentialsGetLoadBalancedHostAttributeType) (ret CredentialsGetLoadBalancedHostRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGetLoadBalancedHostAttributeType(arg *CredentialsGetLoadBalancedHostAttributeType, val CredentialsGetLoadBalancedHostRetType) {
	*arg = &val
}

type CredentialsGetLoadBalancedHostArgType = string
type CredentialsGetLoadBalancedHostRetType = string

/*
	types and functions for password
*/

// isNotNullableString
type CredentialsGetPasswordAttributeType = *string

func getCredentialsGetPasswordAttributeTypeOk(arg CredentialsGetPasswordAttributeType) (ret CredentialsGetPasswordRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGetPasswordAttributeType(arg *CredentialsGetPasswordAttributeType, val CredentialsGetPasswordRetType) {
	*arg = &val
}

type CredentialsGetPasswordArgType = string
type CredentialsGetPasswordRetType = string

/*
	types and functions for port
*/

// isInteger
type CredentialsGetPortAttributeType = *int64
type CredentialsGetPortArgType = int64
type CredentialsGetPortRetType = int64

func getCredentialsGetPortAttributeTypeOk(arg CredentialsGetPortAttributeType) (ret CredentialsGetPortRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGetPortAttributeType(arg *CredentialsGetPortAttributeType, val CredentialsGetPortRetType) {
	*arg = &val
}

/*
	types and functions for uri
*/

// isNotNullableString
type CredentialsGetUriAttributeType = *string

func getCredentialsGetUriAttributeTypeOk(arg CredentialsGetUriAttributeType) (ret CredentialsGetUriRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGetUriAttributeType(arg *CredentialsGetUriAttributeType, val CredentialsGetUriRetType) {
	*arg = &val
}

type CredentialsGetUriArgType = string
type CredentialsGetUriRetType = string

/*
	types and functions for username
*/

// isNotNullableString
type CredentialsGetUsernameAttributeType = *string

func getCredentialsGetUsernameAttributeTypeOk(arg CredentialsGetUsernameAttributeType) (ret CredentialsGetUsernameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCredentialsGetUsernameAttributeType(arg *CredentialsGetUsernameAttributeType, val CredentialsGetUsernameRetType) {
	*arg = &val
}

type CredentialsGetUsernameArgType = string
type CredentialsGetUsernameRetType = string

// Credentials struct for Credentials
type Credentials struct {
	// REQUIRED
	Host             CredentialsGetHostAttributeType             `json:"host"`
	Hosts            CredentialsGetHostsAttributeType            `json:"hosts,omitempty"`
	LoadBalancedHost CredentialsGetLoadBalancedHostAttributeType `json:"load_balanced_host,omitempty"`
	// REQUIRED
	Password CredentialsGetPasswordAttributeType `json:"password"`
	// Can be cast to int32 without loss of precision.
	Port CredentialsGetPortAttributeType `json:"port,omitempty"`
	Uri  CredentialsGetUriAttributeType  `json:"uri,omitempty"`
	// REQUIRED
	Username CredentialsGetUsernameAttributeType `json:"username"`
}

type _Credentials Credentials

// NewCredentials instantiates a new Credentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentials(host CredentialsGetHostArgType, password CredentialsGetPasswordArgType, username CredentialsGetUsernameArgType) *Credentials {
	this := Credentials{}
	setCredentialsGetHostAttributeType(&this.Host, host)
	setCredentialsGetPasswordAttributeType(&this.Password, password)
	setCredentialsGetUsernameAttributeType(&this.Username, username)
	return &this
}

// NewCredentialsWithDefaults instantiates a new Credentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsWithDefaults() *Credentials {
	this := Credentials{}
	return &this
}

// GetHost returns the Host field value
func (o *Credentials) GetHost() (ret CredentialsGetHostRetType) {
	ret, _ = o.GetHostOk()
	return ret
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *Credentials) GetHostOk() (ret CredentialsGetHostRetType, ok bool) {
	return getCredentialsGetHostAttributeTypeOk(o.Host)
}

// SetHost sets field value
func (o *Credentials) SetHost(v CredentialsGetHostRetType) {
	setCredentialsGetHostAttributeType(&o.Host, v)
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *Credentials) GetHosts() (res CredentialsGetHostsRetType) {
	res, _ = o.GetHostsOk()
	return
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetHostsOk() (ret CredentialsGetHostsRetType, ok bool) {
	return getCredentialsGetHostsAttributeTypeOk(o.Hosts)
}

// HasHosts returns a boolean if a field has been set.
func (o *Credentials) HasHosts() bool {
	_, ok := o.GetHostsOk()
	return ok
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *Credentials) SetHosts(v CredentialsGetHostsRetType) {
	setCredentialsGetHostsAttributeType(&o.Hosts, v)
}

// GetLoadBalancedHost returns the LoadBalancedHost field value if set, zero value otherwise.
func (o *Credentials) GetLoadBalancedHost() (res CredentialsGetLoadBalancedHostRetType) {
	res, _ = o.GetLoadBalancedHostOk()
	return
}

// GetLoadBalancedHostOk returns a tuple with the LoadBalancedHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetLoadBalancedHostOk() (ret CredentialsGetLoadBalancedHostRetType, ok bool) {
	return getCredentialsGetLoadBalancedHostAttributeTypeOk(o.LoadBalancedHost)
}

// HasLoadBalancedHost returns a boolean if a field has been set.
func (o *Credentials) HasLoadBalancedHost() bool {
	_, ok := o.GetLoadBalancedHostOk()
	return ok
}

// SetLoadBalancedHost gets a reference to the given string and assigns it to the LoadBalancedHost field.
func (o *Credentials) SetLoadBalancedHost(v CredentialsGetLoadBalancedHostRetType) {
	setCredentialsGetLoadBalancedHostAttributeType(&o.LoadBalancedHost, v)
}

// GetPassword returns the Password field value
func (o *Credentials) GetPassword() (ret CredentialsGetPasswordRetType) {
	ret, _ = o.GetPasswordOk()
	return ret
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *Credentials) GetPasswordOk() (ret CredentialsGetPasswordRetType, ok bool) {
	return getCredentialsGetPasswordAttributeTypeOk(o.Password)
}

// SetPassword sets field value
func (o *Credentials) SetPassword(v CredentialsGetPasswordRetType) {
	setCredentialsGetPasswordAttributeType(&o.Password, v)
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Credentials) GetPort() (res CredentialsGetPortRetType) {
	res, _ = o.GetPortOk()
	return
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetPortOk() (ret CredentialsGetPortRetType, ok bool) {
	return getCredentialsGetPortAttributeTypeOk(o.Port)
}

// HasPort returns a boolean if a field has been set.
func (o *Credentials) HasPort() bool {
	_, ok := o.GetPortOk()
	return ok
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *Credentials) SetPort(v CredentialsGetPortRetType) {
	setCredentialsGetPortAttributeType(&o.Port, v)
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Credentials) GetUri() (res CredentialsGetUriRetType) {
	res, _ = o.GetUriOk()
	return
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credentials) GetUriOk() (ret CredentialsGetUriRetType, ok bool) {
	return getCredentialsGetUriAttributeTypeOk(o.Uri)
}

// HasUri returns a boolean if a field has been set.
func (o *Credentials) HasUri() bool {
	_, ok := o.GetUriOk()
	return ok
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Credentials) SetUri(v CredentialsGetUriRetType) {
	setCredentialsGetUriAttributeType(&o.Uri, v)
}

// GetUsername returns the Username field value
func (o *Credentials) GetUsername() (ret CredentialsGetUsernameRetType) {
	ret, _ = o.GetUsernameOk()
	return ret
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Credentials) GetUsernameOk() (ret CredentialsGetUsernameRetType, ok bool) {
	return getCredentialsGetUsernameAttributeTypeOk(o.Username)
}

// SetUsername sets field value
func (o *Credentials) SetUsername(v CredentialsGetUsernameRetType) {
	setCredentialsGetUsernameAttributeType(&o.Username, v)
}

func (o Credentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCredentialsGetHostAttributeTypeOk(o.Host); ok {
		toSerialize["Host"] = val
	}
	if val, ok := getCredentialsGetHostsAttributeTypeOk(o.Hosts); ok {
		toSerialize["Hosts"] = val
	}
	if val, ok := getCredentialsGetLoadBalancedHostAttributeTypeOk(o.LoadBalancedHost); ok {
		toSerialize["LoadBalancedHost"] = val
	}
	if val, ok := getCredentialsGetPasswordAttributeTypeOk(o.Password); ok {
		toSerialize["Password"] = val
	}
	if val, ok := getCredentialsGetPortAttributeTypeOk(o.Port); ok {
		toSerialize["Port"] = val
	}
	if val, ok := getCredentialsGetUriAttributeTypeOk(o.Uri); ok {
		toSerialize["Uri"] = val
	}
	if val, ok := getCredentialsGetUsernameAttributeTypeOk(o.Username); ok {
		toSerialize["Username"] = val
	}
	return toSerialize, nil
}

type NullableCredentials struct {
	value *Credentials
	isSet bool
}

func (v NullableCredentials) Get() *Credentials {
	return v.value
}

func (v *NullableCredentials) Set(val *Credentials) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentials(val *Credentials) *NullableCredentials {
	return &NullableCredentials{value: val, isSet: true}
}

func (v NullableCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
