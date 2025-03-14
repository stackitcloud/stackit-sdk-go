/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the Listener type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Listener{}

/*
	types and functions for displayName
*/

// isNotNullableString
type ListenerGetDisplayNameAttributeType = *string

func getListenerGetDisplayNameAttributeTypeOk(arg ListenerGetDisplayNameAttributeType) (ret ListenerGetDisplayNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListenerGetDisplayNameAttributeType(arg *ListenerGetDisplayNameAttributeType, val ListenerGetDisplayNameRetType) {
	*arg = &val
}

type ListenerGetDisplayNameArgType = string
type ListenerGetDisplayNameRetType = string

/*
	types and functions for name
*/

// isNotNullableString
type ListenerGetNameAttributeType = *string

func getListenerGetNameAttributeTypeOk(arg ListenerGetNameAttributeType) (ret ListenerGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListenerGetNameAttributeType(arg *ListenerGetNameAttributeType, val ListenerGetNameRetType) {
	*arg = &val
}

type ListenerGetNameArgType = string
type ListenerGetNameRetType = string

/*
	types and functions for port
*/

// isInteger
type ListenerGetPortAttributeType = *int64
type ListenerGetPortArgType = int64
type ListenerGetPortRetType = int64

func getListenerGetPortAttributeTypeOk(arg ListenerGetPortAttributeType) (ret ListenerGetPortRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListenerGetPortAttributeType(arg *ListenerGetPortAttributeType, val ListenerGetPortRetType) {
	*arg = &val
}

/*
	types and functions for protocol
*/

// isEnumRef
type ListenerGetProtocolAttributeType = *string
type ListenerGetProtocolArgType = string
type ListenerGetProtocolRetType = string

func getListenerGetProtocolAttributeTypeOk(arg ListenerGetProtocolAttributeType) (ret ListenerGetProtocolRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListenerGetProtocolAttributeType(arg *ListenerGetProtocolAttributeType, val ListenerGetProtocolRetType) {
	*arg = &val
}

/*
	types and functions for serverNameIndicators
*/

// isArray
type ListenerGetServerNameIndicatorsAttributeType = *[]ServerNameIndicator
type ListenerGetServerNameIndicatorsArgType = []ServerNameIndicator
type ListenerGetServerNameIndicatorsRetType = []ServerNameIndicator

func getListenerGetServerNameIndicatorsAttributeTypeOk(arg ListenerGetServerNameIndicatorsAttributeType) (ret ListenerGetServerNameIndicatorsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListenerGetServerNameIndicatorsAttributeType(arg *ListenerGetServerNameIndicatorsAttributeType, val ListenerGetServerNameIndicatorsRetType) {
	*arg = &val
}

/*
	types and functions for targetPool
*/

// isNotNullableString
type ListenerGetTargetPoolAttributeType = *string

func getListenerGetTargetPoolAttributeTypeOk(arg ListenerGetTargetPoolAttributeType) (ret ListenerGetTargetPoolRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListenerGetTargetPoolAttributeType(arg *ListenerGetTargetPoolAttributeType, val ListenerGetTargetPoolRetType) {
	*arg = &val
}

type ListenerGetTargetPoolArgType = string
type ListenerGetTargetPoolRetType = string

/*
	types and functions for tcp
*/

// isModel
type ListenerGetTcpAttributeType = *OptionsTCP
type ListenerGetTcpArgType = OptionsTCP
type ListenerGetTcpRetType = OptionsTCP

func getListenerGetTcpAttributeTypeOk(arg ListenerGetTcpAttributeType) (ret ListenerGetTcpRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListenerGetTcpAttributeType(arg *ListenerGetTcpAttributeType, val ListenerGetTcpRetType) {
	*arg = &val
}

/*
	types and functions for udp
*/

// isModel
type ListenerGetUdpAttributeType = *OptionsUDP
type ListenerGetUdpArgType = OptionsUDP
type ListenerGetUdpRetType = OptionsUDP

func getListenerGetUdpAttributeTypeOk(arg ListenerGetUdpAttributeType) (ret ListenerGetUdpRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setListenerGetUdpAttributeType(arg *ListenerGetUdpAttributeType, val ListenerGetUdpRetType) {
	*arg = &val
}

// Listener struct for Listener
type Listener struct {
	DisplayName ListenerGetDisplayNameAttributeType `json:"displayName,omitempty"`
	// Will be used to reference a listener and will replace display name in the future. Currently uses <protocol>-<port> as the name if no display name is given.
	Name ListenerGetNameAttributeType `json:"name,omitempty"`
	// Port number where we listen for traffic
	// Can be cast to int32 without loss of precision.
	Port ListenerGetPortAttributeType `json:"port,omitempty"`
	// Protocol is the highest network protocol we understand to load balance. Currently only PROTOCOL_TCP, PROTOCOL_TCP_PROXY and PROTOCOL_TLS_PASSTHROUGH are supported.
	Protocol ListenerGetProtocolAttributeType `json:"protocol,omitempty"`
	// Server Name Idicators config for domains to be routed to the desired target pool for this listener.
	ServerNameIndicators ListenerGetServerNameIndicatorsAttributeType `json:"serverNameIndicators,omitempty"`
	// Reference target pool by target pool name.
	TargetPool ListenerGetTargetPoolAttributeType `json:"targetPool,omitempty"`
	Tcp        ListenerGetTcpAttributeType        `json:"tcp,omitempty"`
	Udp        ListenerGetUdpAttributeType        `json:"udp,omitempty"`
}

// NewListener instantiates a new Listener object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListener() *Listener {
	this := Listener{}
	return &this
}

// NewListenerWithDefaults instantiates a new Listener object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenerWithDefaults() *Listener {
	this := Listener{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Listener) GetDisplayName() (res ListenerGetDisplayNameRetType) {
	res, _ = o.GetDisplayNameOk()
	return
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetDisplayNameOk() (ret ListenerGetDisplayNameRetType, ok bool) {
	return getListenerGetDisplayNameAttributeTypeOk(o.DisplayName)
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Listener) HasDisplayName() bool {
	_, ok := o.GetDisplayNameOk()
	return ok
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Listener) SetDisplayName(v ListenerGetDisplayNameRetType) {
	setListenerGetDisplayNameAttributeType(&o.DisplayName, v)
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Listener) GetName() (res ListenerGetNameRetType) {
	res, _ = o.GetNameOk()
	return
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetNameOk() (ret ListenerGetNameRetType, ok bool) {
	return getListenerGetNameAttributeTypeOk(o.Name)
}

// HasName returns a boolean if a field has been set.
func (o *Listener) HasName() bool {
	_, ok := o.GetNameOk()
	return ok
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Listener) SetName(v ListenerGetNameRetType) {
	setListenerGetNameAttributeType(&o.Name, v)
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Listener) GetPort() (res ListenerGetPortRetType) {
	res, _ = o.GetPortOk()
	return
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetPortOk() (ret ListenerGetPortRetType, ok bool) {
	return getListenerGetPortAttributeTypeOk(o.Port)
}

// HasPort returns a boolean if a field has been set.
func (o *Listener) HasPort() bool {
	_, ok := o.GetPortOk()
	return ok
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *Listener) SetPort(v ListenerGetPortRetType) {
	setListenerGetPortAttributeType(&o.Port, v)
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Listener) GetProtocol() (res ListenerGetProtocolRetType) {
	res, _ = o.GetProtocolOk()
	return
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetProtocolOk() (ret ListenerGetProtocolRetType, ok bool) {
	return getListenerGetProtocolAttributeTypeOk(o.Protocol)
}

// HasProtocol returns a boolean if a field has been set.
func (o *Listener) HasProtocol() bool {
	_, ok := o.GetProtocolOk()
	return ok
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *Listener) SetProtocol(v ListenerGetProtocolRetType) {
	setListenerGetProtocolAttributeType(&o.Protocol, v)
}

// GetServerNameIndicators returns the ServerNameIndicators field value if set, zero value otherwise.
func (o *Listener) GetServerNameIndicators() (res ListenerGetServerNameIndicatorsRetType) {
	res, _ = o.GetServerNameIndicatorsOk()
	return
}

// GetServerNameIndicatorsOk returns a tuple with the ServerNameIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetServerNameIndicatorsOk() (ret ListenerGetServerNameIndicatorsRetType, ok bool) {
	return getListenerGetServerNameIndicatorsAttributeTypeOk(o.ServerNameIndicators)
}

// HasServerNameIndicators returns a boolean if a field has been set.
func (o *Listener) HasServerNameIndicators() bool {
	_, ok := o.GetServerNameIndicatorsOk()
	return ok
}

// SetServerNameIndicators gets a reference to the given []ServerNameIndicator and assigns it to the ServerNameIndicators field.
func (o *Listener) SetServerNameIndicators(v ListenerGetServerNameIndicatorsRetType) {
	setListenerGetServerNameIndicatorsAttributeType(&o.ServerNameIndicators, v)
}

// GetTargetPool returns the TargetPool field value if set, zero value otherwise.
func (o *Listener) GetTargetPool() (res ListenerGetTargetPoolRetType) {
	res, _ = o.GetTargetPoolOk()
	return
}

// GetTargetPoolOk returns a tuple with the TargetPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetTargetPoolOk() (ret ListenerGetTargetPoolRetType, ok bool) {
	return getListenerGetTargetPoolAttributeTypeOk(o.TargetPool)
}

// HasTargetPool returns a boolean if a field has been set.
func (o *Listener) HasTargetPool() bool {
	_, ok := o.GetTargetPoolOk()
	return ok
}

// SetTargetPool gets a reference to the given string and assigns it to the TargetPool field.
func (o *Listener) SetTargetPool(v ListenerGetTargetPoolRetType) {
	setListenerGetTargetPoolAttributeType(&o.TargetPool, v)
}

// GetTcp returns the Tcp field value if set, zero value otherwise.
func (o *Listener) GetTcp() (res ListenerGetTcpRetType) {
	res, _ = o.GetTcpOk()
	return
}

// GetTcpOk returns a tuple with the Tcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetTcpOk() (ret ListenerGetTcpRetType, ok bool) {
	return getListenerGetTcpAttributeTypeOk(o.Tcp)
}

// HasTcp returns a boolean if a field has been set.
func (o *Listener) HasTcp() bool {
	_, ok := o.GetTcpOk()
	return ok
}

// SetTcp gets a reference to the given OptionsTCP and assigns it to the Tcp field.
func (o *Listener) SetTcp(v ListenerGetTcpRetType) {
	setListenerGetTcpAttributeType(&o.Tcp, v)
}

// GetUdp returns the Udp field value if set, zero value otherwise.
func (o *Listener) GetUdp() (res ListenerGetUdpRetType) {
	res, _ = o.GetUdpOk()
	return
}

// GetUdpOk returns a tuple with the Udp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetUdpOk() (ret ListenerGetUdpRetType, ok bool) {
	return getListenerGetUdpAttributeTypeOk(o.Udp)
}

// HasUdp returns a boolean if a field has been set.
func (o *Listener) HasUdp() bool {
	_, ok := o.GetUdpOk()
	return ok
}

// SetUdp gets a reference to the given OptionsUDP and assigns it to the Udp field.
func (o *Listener) SetUdp(v ListenerGetUdpRetType) {
	setListenerGetUdpAttributeType(&o.Udp, v)
}

func (o Listener) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getListenerGetDisplayNameAttributeTypeOk(o.DisplayName); ok {
		toSerialize["DisplayName"] = val
	}
	if val, ok := getListenerGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getListenerGetPortAttributeTypeOk(o.Port); ok {
		toSerialize["Port"] = val
	}
	if val, ok := getListenerGetProtocolAttributeTypeOk(o.Protocol); ok {
		toSerialize["Protocol"] = val
	}
	if val, ok := getListenerGetServerNameIndicatorsAttributeTypeOk(o.ServerNameIndicators); ok {
		toSerialize["ServerNameIndicators"] = val
	}
	if val, ok := getListenerGetTargetPoolAttributeTypeOk(o.TargetPool); ok {
		toSerialize["TargetPool"] = val
	}
	if val, ok := getListenerGetTcpAttributeTypeOk(o.Tcp); ok {
		toSerialize["Tcp"] = val
	}
	if val, ok := getListenerGetUdpAttributeTypeOk(o.Udp); ok {
		toSerialize["Udp"] = val
	}
	return toSerialize, nil
}

type NullableListener struct {
	value *Listener
	isSet bool
}

func (v NullableListener) Get() *Listener {
	return v.value
}

func (v *NullableListener) Set(val *Listener) {
	v.value = val
	v.isSet = true
}

func (v NullableListener) IsSet() bool {
	return v.isSet
}

func (v *NullableListener) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListener(val *Listener) *NullableListener {
	return &NullableListener{value: val, isSet: true}
}

func (v NullableListener) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListener) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
