/*
STACKIT MongoDB Service API

This is the documentation for the STACKIT MongoDB Flex Service API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbflex

// User struct for User
type User struct {
	Database *string   `json:"database,omitempty"`
	Host     *string   `json:"host,omitempty"`
	Id       *string   `json:"id,omitempty"`
	Password *string   `json:"password,omitempty"`
	Port     *int64    `json:"port,omitempty"`
	Roles    *[]string `json:"roles,omitempty"`
	Uri      *string   `json:"uri,omitempty"`
	Username *string   `json:"username,omitempty"`
}
