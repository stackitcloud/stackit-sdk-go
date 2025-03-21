/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

import (
	"encoding/json"
)

// checks if the CatalogProductSupportResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogProductSupportResource{}

/*
	types and functions for supportLink
*/

// isNotNullableString
type CatalogProductSupportResourceGetSupportLinkAttributeType = *string

func getCatalogProductSupportResourceGetSupportLinkAttributeTypeOk(arg CatalogProductSupportResourceGetSupportLinkAttributeType) (ret CatalogProductSupportResourceGetSupportLinkRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCatalogProductSupportResourceGetSupportLinkAttributeType(arg *CatalogProductSupportResourceGetSupportLinkAttributeType, val CatalogProductSupportResourceGetSupportLinkRetType) {
	*arg = &val
}

type CatalogProductSupportResourceGetSupportLinkArgType = string
type CatalogProductSupportResourceGetSupportLinkRetType = string

/*
	types and functions for supportLinkTitle
*/

// isNotNullableString
type CatalogProductSupportResourceGetSupportLinkTitleAttributeType = *string

func getCatalogProductSupportResourceGetSupportLinkTitleAttributeTypeOk(arg CatalogProductSupportResourceGetSupportLinkTitleAttributeType) (ret CatalogProductSupportResourceGetSupportLinkTitleRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setCatalogProductSupportResourceGetSupportLinkTitleAttributeType(arg *CatalogProductSupportResourceGetSupportLinkTitleAttributeType, val CatalogProductSupportResourceGetSupportLinkTitleRetType) {
	*arg = &val
}

type CatalogProductSupportResourceGetSupportLinkTitleArgType = string
type CatalogProductSupportResourceGetSupportLinkTitleRetType = string

// CatalogProductSupportResource struct for CatalogProductSupportResource
type CatalogProductSupportResource struct {
	// The support resource link.
	SupportLink CatalogProductSupportResourceGetSupportLinkAttributeType `json:"supportLink,omitempty"`
	// The support resource link title.
	SupportLinkTitle CatalogProductSupportResourceGetSupportLinkTitleAttributeType `json:"supportLinkTitle,omitempty"`
}

// NewCatalogProductSupportResource instantiates a new CatalogProductSupportResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogProductSupportResource() *CatalogProductSupportResource {
	this := CatalogProductSupportResource{}
	return &this
}

// NewCatalogProductSupportResourceWithDefaults instantiates a new CatalogProductSupportResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogProductSupportResourceWithDefaults() *CatalogProductSupportResource {
	this := CatalogProductSupportResource{}
	return &this
}

// GetSupportLink returns the SupportLink field value if set, zero value otherwise.
func (o *CatalogProductSupportResource) GetSupportLink() (res CatalogProductSupportResourceGetSupportLinkRetType) {
	res, _ = o.GetSupportLinkOk()
	return
}

// GetSupportLinkOk returns a tuple with the SupportLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogProductSupportResource) GetSupportLinkOk() (ret CatalogProductSupportResourceGetSupportLinkRetType, ok bool) {
	return getCatalogProductSupportResourceGetSupportLinkAttributeTypeOk(o.SupportLink)
}

// HasSupportLink returns a boolean if a field has been set.
func (o *CatalogProductSupportResource) HasSupportLink() bool {
	_, ok := o.GetSupportLinkOk()
	return ok
}

// SetSupportLink gets a reference to the given string and assigns it to the SupportLink field.
func (o *CatalogProductSupportResource) SetSupportLink(v CatalogProductSupportResourceGetSupportLinkRetType) {
	setCatalogProductSupportResourceGetSupportLinkAttributeType(&o.SupportLink, v)
}

// GetSupportLinkTitle returns the SupportLinkTitle field value if set, zero value otherwise.
func (o *CatalogProductSupportResource) GetSupportLinkTitle() (res CatalogProductSupportResourceGetSupportLinkTitleRetType) {
	res, _ = o.GetSupportLinkTitleOk()
	return
}

// GetSupportLinkTitleOk returns a tuple with the SupportLinkTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogProductSupportResource) GetSupportLinkTitleOk() (ret CatalogProductSupportResourceGetSupportLinkTitleRetType, ok bool) {
	return getCatalogProductSupportResourceGetSupportLinkTitleAttributeTypeOk(o.SupportLinkTitle)
}

// HasSupportLinkTitle returns a boolean if a field has been set.
func (o *CatalogProductSupportResource) HasSupportLinkTitle() bool {
	_, ok := o.GetSupportLinkTitleOk()
	return ok
}

// SetSupportLinkTitle gets a reference to the given string and assigns it to the SupportLinkTitle field.
func (o *CatalogProductSupportResource) SetSupportLinkTitle(v CatalogProductSupportResourceGetSupportLinkTitleRetType) {
	setCatalogProductSupportResourceGetSupportLinkTitleAttributeType(&o.SupportLinkTitle, v)
}

func (o CatalogProductSupportResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getCatalogProductSupportResourceGetSupportLinkAttributeTypeOk(o.SupportLink); ok {
		toSerialize["SupportLink"] = val
	}
	if val, ok := getCatalogProductSupportResourceGetSupportLinkTitleAttributeTypeOk(o.SupportLinkTitle); ok {
		toSerialize["SupportLinkTitle"] = val
	}
	return toSerialize, nil
}

type NullableCatalogProductSupportResource struct {
	value *CatalogProductSupportResource
	isSet bool
}

func (v NullableCatalogProductSupportResource) Get() *CatalogProductSupportResource {
	return v.value
}

func (v *NullableCatalogProductSupportResource) Set(val *CatalogProductSupportResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogProductSupportResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogProductSupportResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogProductSupportResource(val *CatalogProductSupportResource) *NullableCatalogProductSupportResource {
	return &NullableCatalogProductSupportResource{value: val, isSet: true}
}

func (v NullableCatalogProductSupportResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogProductSupportResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
