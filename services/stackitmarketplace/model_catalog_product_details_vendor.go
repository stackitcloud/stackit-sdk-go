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

// checks if the CatalogProductDetailsVendor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogProductDetailsVendor{}

// CatalogProductDetailsVendor struct for CatalogProductDetailsVendor
type CatalogProductDetailsVendor struct {
	// The vendor description.
	// REQUIRED
	Description *string `json:"description"`
	// The vendor logo base64 encoded.
	// REQUIRED
	Logo *string `json:"logo"`
	// The vendor name.
	// REQUIRED
	Name *string `json:"name"`
	// The vendor ID.
	// REQUIRED
	VendorId *string `json:"vendorId"`
	// The vendor video URL.
	// REQUIRED
	VideoUrl *string `json:"videoUrl"`
	// The vendor website URL.
	// REQUIRED
	WebsiteUrl *string `json:"websiteUrl"`
}

type _CatalogProductDetailsVendor CatalogProductDetailsVendor

// NewCatalogProductDetailsVendor instantiates a new CatalogProductDetailsVendor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogProductDetailsVendor(description *string, logo *string, name *string, vendorId *string, videoUrl *string, websiteUrl *string) *CatalogProductDetailsVendor {
	this := CatalogProductDetailsVendor{}
	this.Description = description
	this.Logo = logo
	this.Name = name
	this.VendorId = vendorId
	this.VideoUrl = videoUrl
	this.WebsiteUrl = websiteUrl
	return &this
}

// NewCatalogProductDetailsVendorWithDefaults instantiates a new CatalogProductDetailsVendor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogProductDetailsVendorWithDefaults() *CatalogProductDetailsVendor {
	this := CatalogProductDetailsVendor{}
	return &this
}

// GetDescription returns the Description field value
func (o *CatalogProductDetailsVendor) GetDescription() *string {
	if o == nil || IsNil(o.Description) {
		var ret *string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CatalogProductDetailsVendor) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description, true
}

// SetDescription sets field value
func (o *CatalogProductDetailsVendor) SetDescription(v *string) {
	o.Description = v
}

// GetLogo returns the Logo field value
func (o *CatalogProductDetailsVendor) GetLogo() *string {
	if o == nil || IsNil(o.Logo) {
		var ret *string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *CatalogProductDetailsVendor) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo, true
}

// SetLogo sets field value
func (o *CatalogProductDetailsVendor) SetLogo(v *string) {
	o.Logo = v
}

// GetName returns the Name field value
func (o *CatalogProductDetailsVendor) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CatalogProductDetailsVendor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CatalogProductDetailsVendor) SetName(v *string) {
	o.Name = v
}

// GetVendorId returns the VendorId field value
func (o *CatalogProductDetailsVendor) GetVendorId() *string {
	if o == nil || IsNil(o.VendorId) {
		var ret *string
		return ret
	}

	return o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value
// and a boolean to check if the value has been set.
func (o *CatalogProductDetailsVendor) GetVendorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VendorId, true
}

// SetVendorId sets field value
func (o *CatalogProductDetailsVendor) SetVendorId(v *string) {
	o.VendorId = v
}

// GetVideoUrl returns the VideoUrl field value
func (o *CatalogProductDetailsVendor) GetVideoUrl() *string {
	if o == nil || IsNil(o.VideoUrl) {
		var ret *string
		return ret
	}

	return o.VideoUrl
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value
// and a boolean to check if the value has been set.
func (o *CatalogProductDetailsVendor) GetVideoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoUrl, true
}

// SetVideoUrl sets field value
func (o *CatalogProductDetailsVendor) SetVideoUrl(v *string) {
	o.VideoUrl = v
}

// GetWebsiteUrl returns the WebsiteUrl field value
func (o *CatalogProductDetailsVendor) GetWebsiteUrl() *string {
	if o == nil || IsNil(o.WebsiteUrl) {
		var ret *string
		return ret
	}

	return o.WebsiteUrl
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value
// and a boolean to check if the value has been set.
func (o *CatalogProductDetailsVendor) GetWebsiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebsiteUrl, true
}

// SetWebsiteUrl sets field value
func (o *CatalogProductDetailsVendor) SetWebsiteUrl(v *string) {
	o.WebsiteUrl = v
}

func (o CatalogProductDetailsVendor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["logo"] = o.Logo
	toSerialize["name"] = o.Name
	toSerialize["vendorId"] = o.VendorId
	toSerialize["videoUrl"] = o.VideoUrl
	toSerialize["websiteUrl"] = o.WebsiteUrl
	return toSerialize, nil
}

type NullableCatalogProductDetailsVendor struct {
	value *CatalogProductDetailsVendor
	isSet bool
}

func (v NullableCatalogProductDetailsVendor) Get() *CatalogProductDetailsVendor {
	return v.value
}

func (v *NullableCatalogProductDetailsVendor) Set(val *CatalogProductDetailsVendor) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogProductDetailsVendor) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogProductDetailsVendor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogProductDetailsVendor(val *CatalogProductDetailsVendor) *NullableCatalogProductDetailsVendor {
	return &NullableCatalogProductDetailsVendor{value: val, isSet: true}
}

func (v NullableCatalogProductDetailsVendor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogProductDetailsVendor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
