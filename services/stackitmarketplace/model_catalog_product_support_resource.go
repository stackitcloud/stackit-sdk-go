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

// CatalogProductSupportResource struct for CatalogProductSupportResource
type CatalogProductSupportResource struct {
	// The support resource link.
	SupportLink *string `json:"supportLink,omitempty"`
	// The support resource link title.
	SupportLinkTitle *string `json:"supportLinkTitle,omitempty"`
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
func (o *CatalogProductSupportResource) GetSupportLink() *string {
	if o == nil || IsNil(o.SupportLink) {
		var ret *string
		return ret
	}
	return o.SupportLink
}

// GetSupportLinkOk returns a tuple with the SupportLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogProductSupportResource) GetSupportLinkOk() (*string, bool) {
	if o == nil || IsNil(o.SupportLink) {
		return nil, false
	}
	return o.SupportLink, true
}

// HasSupportLink returns a boolean if a field has been set.
func (o *CatalogProductSupportResource) HasSupportLink() bool {
	if o != nil && !IsNil(o.SupportLink) {
		return true
	}

	return false
}

// SetSupportLink gets a reference to the given string and assigns it to the SupportLink field.
func (o *CatalogProductSupportResource) SetSupportLink(v *string) {
	o.SupportLink = v
}

// GetSupportLinkTitle returns the SupportLinkTitle field value if set, zero value otherwise.
func (o *CatalogProductSupportResource) GetSupportLinkTitle() *string {
	if o == nil || IsNil(o.SupportLinkTitle) {
		var ret *string
		return ret
	}
	return o.SupportLinkTitle
}

// GetSupportLinkTitleOk returns a tuple with the SupportLinkTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogProductSupportResource) GetSupportLinkTitleOk() (*string, bool) {
	if o == nil || IsNil(o.SupportLinkTitle) {
		return nil, false
	}
	return o.SupportLinkTitle, true
}

// HasSupportLinkTitle returns a boolean if a field has been set.
func (o *CatalogProductSupportResource) HasSupportLinkTitle() bool {
	if o != nil && !IsNil(o.SupportLinkTitle) {
		return true
	}

	return false
}

// SetSupportLinkTitle gets a reference to the given string and assigns it to the SupportLinkTitle field.
func (o *CatalogProductSupportResource) SetSupportLinkTitle(v *string) {
	o.SupportLinkTitle = v
}

func (o CatalogProductSupportResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportLink) {
		toSerialize["supportLink"] = o.SupportLink
	}
	if !IsNil(o.SupportLinkTitle) {
		toSerialize["supportLinkTitle"] = o.SupportLinkTitle
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