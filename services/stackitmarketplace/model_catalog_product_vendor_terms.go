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

// checks if the CatalogProductVendorTerms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogProductVendorTerms{}

// CatalogProductVendorTerms struct for CatalogProductVendorTerms
type CatalogProductVendorTerms struct {
	// The terms of use description.
	// REQUIRED
	Description *string `json:"description"`
	// The terms of use title.
	// REQUIRED
	Title *string `json:"title"`
	// The terms of use url.
	// REQUIRED
	Url *string `json:"url"`
}

type _CatalogProductVendorTerms CatalogProductVendorTerms

// NewCatalogProductVendorTerms instantiates a new CatalogProductVendorTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogProductVendorTerms(description *string, title *string, url *string) *CatalogProductVendorTerms {
	this := CatalogProductVendorTerms{}
	this.Description = description
	this.Title = title
	this.Url = url
	return &this
}

// NewCatalogProductVendorTermsWithDefaults instantiates a new CatalogProductVendorTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogProductVendorTermsWithDefaults() *CatalogProductVendorTerms {
	this := CatalogProductVendorTerms{}
	return &this
}

// GetDescription returns the Description field value
func (o *CatalogProductVendorTerms) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CatalogProductVendorTerms) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description, true
}

// SetDescription sets field value
func (o *CatalogProductVendorTerms) SetDescription(v *string) {
	o.Description = v
}

// GetTitle returns the Title field value
func (o *CatalogProductVendorTerms) GetTitle() *string {
	if o == nil || IsNil(o.Title) {
		var ret *string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CatalogProductVendorTerms) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title, true
}

// SetTitle sets field value
func (o *CatalogProductVendorTerms) SetTitle(v *string) {
	o.Title = v
}

// GetUrl returns the Url field value
func (o *CatalogProductVendorTerms) GetUrl() *string {
	if o == nil || IsNil(o.Url) {
		var ret *string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CatalogProductVendorTerms) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url, true
}

// SetUrl sets field value
func (o *CatalogProductVendorTerms) SetUrl(v *string) {
	o.Url = v
}

func (o CatalogProductVendorTerms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["title"] = o.Title
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableCatalogProductVendorTerms struct {
	value *CatalogProductVendorTerms
	isSet bool
}

func (v NullableCatalogProductVendorTerms) Get() *CatalogProductVendorTerms {
	return v.value
}

func (v *NullableCatalogProductVendorTerms) Set(val *CatalogProductVendorTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogProductVendorTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogProductVendorTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogProductVendorTerms(val *CatalogProductVendorTerms) *NullableCatalogProductVendorTerms {
	return &NullableCatalogProductVendorTerms{value: val, isSet: true}
}

func (v NullableCatalogProductVendorTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogProductVendorTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
