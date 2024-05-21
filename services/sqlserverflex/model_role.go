/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
	"fmt"
)

// Role the model 'Role'
type Role string

// List of Role
const (
	ROLE_SELECT                  Role = "SELECT"
	ROLE_INSERT                  Role = "INSERT"
	ROLE_UPDATE                  Role = "UPDATE"
	ROLE_DELETE                  Role = "DELETE"
	ROLE_TRUNCATE                Role = "TRUNCATE"
	ROLE_TRIGGER                 Role = "TRIGGER"
	ROLE_CREATE                  Role = "CREATE"
	ROLE_CONNECT                 Role = "CONNECT"
	ROLE_TEMPORARY               Role = "TEMPORARY"
	ROLE_EXECUTE                 Role = "EXECUTE"
	ROLE_USAGE                   Role = "USAGE"
	ROLE_READ                    Role = "read"
	ROLE_READ_WRITE              Role = "readWrite"
	ROLE_READ_WRITE_ANY_DATABASE Role = "readWriteAnyDatabase"
)

// All allowed values of Role enum
var AllowedRoleEnumValues = []Role{
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"TRUNCATE",
	"TRIGGER",
	"CREATE",
	"CONNECT",
	"TEMPORARY",
	"EXECUTE",
	"USAGE",
	"read",
	"readWrite",
	"readWriteAnyDatabase",
}

func (v *Role) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// Allow unmarshalling zero value for testing purposes
	var zeroValue string
	if value == zeroValue {
		return nil
	}
	enumTypeValue := Role(value)
	for _, existing := range AllowedRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Role", value)
}

// NewRoleFromValue returns a pointer to a valid Role
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleFromValue(v string) (*Role, error) {
	ev := Role(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Role: valid values are %v", v, AllowedRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Role) IsValid() bool {
	for _, existing := range AllowedRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Role value
func (v Role) Ptr() *Role {
	return &v
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
