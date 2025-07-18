/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorization

import (
	"encoding/json"
)

// checks if the Member type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Member{}

/*
	types and functions for role
*/

// isNotNullableString
type MemberGetRoleAttributeType = *string

func getMemberGetRoleAttributeTypeOk(arg MemberGetRoleAttributeType) (ret MemberGetRoleRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setMemberGetRoleAttributeType(arg *MemberGetRoleAttributeType, val MemberGetRoleRetType) {
	*arg = &val
}

type MemberGetRoleArgType = string
type MemberGetRoleRetType = string

/*
	types and functions for subject
*/

// isNotNullableString
type MemberGetSubjectAttributeType = *string

func getMemberGetSubjectAttributeTypeOk(arg MemberGetSubjectAttributeType) (ret MemberGetSubjectRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setMemberGetSubjectAttributeType(arg *MemberGetSubjectAttributeType, val MemberGetSubjectRetType) {
	*arg = &val
}

type MemberGetSubjectArgType = string
type MemberGetSubjectRetType = string

// Member struct for Member
type Member struct {
	// REQUIRED
	Role MemberGetRoleAttributeType `json:"role" required:"true"`
	// REQUIRED
	Subject MemberGetSubjectAttributeType `json:"subject" required:"true"`
}

type _Member Member

// NewMember instantiates a new Member object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMember(role MemberGetRoleArgType, subject MemberGetSubjectArgType) *Member {
	this := Member{}
	setMemberGetRoleAttributeType(&this.Role, role)
	setMemberGetSubjectAttributeType(&this.Subject, subject)
	return &this
}

// NewMemberWithDefaults instantiates a new Member object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberWithDefaults() *Member {
	this := Member{}
	return &this
}

// GetRole returns the Role field value
func (o *Member) GetRole() (ret MemberGetRoleRetType) {
	ret, _ = o.GetRoleOk()
	return ret
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Member) GetRoleOk() (ret MemberGetRoleRetType, ok bool) {
	return getMemberGetRoleAttributeTypeOk(o.Role)
}

// SetRole sets field value
func (o *Member) SetRole(v MemberGetRoleRetType) {
	setMemberGetRoleAttributeType(&o.Role, v)
}

// GetSubject returns the Subject field value
func (o *Member) GetSubject() (ret MemberGetSubjectRetType) {
	ret, _ = o.GetSubjectOk()
	return ret
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *Member) GetSubjectOk() (ret MemberGetSubjectRetType, ok bool) {
	return getMemberGetSubjectAttributeTypeOk(o.Subject)
}

// SetSubject sets field value
func (o *Member) SetSubject(v MemberGetSubjectRetType) {
	setMemberGetSubjectAttributeType(&o.Subject, v)
}

func (o Member) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getMemberGetRoleAttributeTypeOk(o.Role); ok {
		toSerialize["Role"] = val
	}
	if val, ok := getMemberGetSubjectAttributeTypeOk(o.Subject); ok {
		toSerialize["Subject"] = val
	}
	return toSerialize, nil
}

type NullableMember struct {
	value *Member
	isSet bool
}

func (v NullableMember) Get() *Member {
	return v.value
}

func (v *NullableMember) Set(val *Member) {
	v.value = val
	v.isSet = true
}

func (v NullableMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMember(val *Member) *NullableMember {
	return &NullableMember{value: val, isSet: true}
}

func (v NullableMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
