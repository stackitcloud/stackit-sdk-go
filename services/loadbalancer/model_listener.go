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

// Listener struct for Listener
type Listener struct {
	DisplayName *string `json:"displayName,omitempty"`
	// Will be used to reference a listener and will replace display name in the future. Currently uses <protocol>-<port> as the name if no display name is given.
	Name *string `json:"name,omitempty"`
	// Port number where we listen for traffic
	// Can be cast to int32 without loss of precision.
	Port *int64 `json:"port,omitempty"`
	// Protocol is the highest network protocol we understand to load balance. Currently only PROTOCOL_TCP, PROTOCOL_TCP_PROXY and PROTOCOL_TLS_PASSTHROUGH are supported.
	Protocol *string `json:"protocol,omitempty"`
	// Server Name Idicators config for domains to be routed to the desired target pool for this listener.
	ServerNameIndicators *[]ServerNameIndicator `json:"serverNameIndicators,omitempty"`
	// Reference target pool by target pool name.
	TargetPool *string     `json:"targetPool,omitempty"`
	Tcp        *OptionsTCP `json:"tcp,omitempty"`
	Udp        *OptionsUDP `json:"udp,omitempty"`
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
func (o *Listener) GetDisplayName() *string {
	if o == nil || IsNil(o.DisplayName) {
		var ret *string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Listener) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Listener) SetDisplayName(v *string) {
	o.DisplayName = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Listener) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Listener) HasName() bool {
	if o != nil && !IsNil(o.Name) && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Listener) SetName(v *string) {
	o.Name = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *Listener) GetPort() *int64 {
	if o == nil || IsNil(o.Port) {
		var ret *int64
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Listener) HasPort() bool {
	if o != nil && !IsNil(o.Port) && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *Listener) SetPort(v *int64) {
	o.Port = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Listener) GetProtocol() *string {
	if o == nil || IsNil(o.Protocol) {
		var ret *string
		return ret
	}
	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Listener) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *Listener) SetProtocol(v *string) {
	o.Protocol = v
}

// GetServerNameIndicators returns the ServerNameIndicators field value if set, zero value otherwise.
func (o *Listener) GetServerNameIndicators() *[]ServerNameIndicator {
	if o == nil || IsNil(o.ServerNameIndicators) {
		var ret *[]ServerNameIndicator
		return ret
	}
	return o.ServerNameIndicators
}

// GetServerNameIndicatorsOk returns a tuple with the ServerNameIndicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetServerNameIndicatorsOk() (*[]ServerNameIndicator, bool) {
	if o == nil || IsNil(o.ServerNameIndicators) {
		return nil, false
	}
	return o.ServerNameIndicators, true
}

// HasServerNameIndicators returns a boolean if a field has been set.
func (o *Listener) HasServerNameIndicators() bool {
	if o != nil && !IsNil(o.ServerNameIndicators) && !IsNil(o.ServerNameIndicators) {
		return true
	}

	return false
}

// SetServerNameIndicators gets a reference to the given []ServerNameIndicator and assigns it to the ServerNameIndicators field.
func (o *Listener) SetServerNameIndicators(v *[]ServerNameIndicator) {
	o.ServerNameIndicators = v
}

// GetTargetPool returns the TargetPool field value if set, zero value otherwise.
func (o *Listener) GetTargetPool() *string {
	if o == nil || IsNil(o.TargetPool) {
		var ret *string
		return ret
	}
	return o.TargetPool
}

// GetTargetPoolOk returns a tuple with the TargetPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetTargetPoolOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPool) {
		return nil, false
	}
	return o.TargetPool, true
}

// HasTargetPool returns a boolean if a field has been set.
func (o *Listener) HasTargetPool() bool {
	if o != nil && !IsNil(o.TargetPool) && !IsNil(o.TargetPool) {
		return true
	}

	return false
}

// SetTargetPool gets a reference to the given string and assigns it to the TargetPool field.
func (o *Listener) SetTargetPool(v *string) {
	o.TargetPool = v
}

// GetTcp returns the Tcp field value if set, zero value otherwise.
func (o *Listener) GetTcp() *OptionsTCP {
	if o == nil || IsNil(o.Tcp) {
		var ret *OptionsTCP
		return ret
	}
	return o.Tcp
}

// GetTcpOk returns a tuple with the Tcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetTcpOk() (*OptionsTCP, bool) {
	if o == nil || IsNil(o.Tcp) {
		return nil, false
	}
	return o.Tcp, true
}

// HasTcp returns a boolean if a field has been set.
func (o *Listener) HasTcp() bool {
	if o != nil && !IsNil(o.Tcp) && !IsNil(o.Tcp) {
		return true
	}

	return false
}

// SetTcp gets a reference to the given OptionsTCP and assigns it to the Tcp field.
func (o *Listener) SetTcp(v *OptionsTCP) {
	o.Tcp = v
}

// GetUdp returns the Udp field value if set, zero value otherwise.
func (o *Listener) GetUdp() *OptionsUDP {
	if o == nil || IsNil(o.Udp) {
		var ret *OptionsUDP
		return ret
	}
	return o.Udp
}

// GetUdpOk returns a tuple with the Udp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetUdpOk() (*OptionsUDP, bool) {
	if o == nil || IsNil(o.Udp) {
		return nil, false
	}
	return o.Udp, true
}

// HasUdp returns a boolean if a field has been set.
func (o *Listener) HasUdp() bool {
	if o != nil && !IsNil(o.Udp) && !IsNil(o.Udp) {
		return true
	}

	return false
}

// SetUdp gets a reference to the given OptionsUDP and assigns it to the Udp field.
func (o *Listener) SetUdp(v *OptionsUDP) {
	o.Udp = v
}

func (o Listener) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.ServerNameIndicators) {
		toSerialize["serverNameIndicators"] = o.ServerNameIndicators
	}
	if !IsNil(o.TargetPool) {
		toSerialize["targetPool"] = o.TargetPool
	}
	if !IsNil(o.Tcp) {
		toSerialize["tcp"] = o.Tcp
	}
	if !IsNil(o.Udp) {
		toSerialize["udp"] = o.Udp
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
