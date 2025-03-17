/*
STACKIT Model Serving API

This API provides endpoints for the model serving api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package modelserving

import (
	"encoding/json"
)

// checks if the Model type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model{}

/*
	types and functions for category
*/

// isEnumRef
type ModelGetCategoryAttributeType = *string
type ModelGetCategoryArgType = string
type ModelGetCategoryRetType = string

func getModelGetCategoryAttributeTypeOk(arg ModelGetCategoryAttributeType) (ret ModelGetCategoryRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetCategoryAttributeType(arg *ModelGetCategoryAttributeType, val ModelGetCategoryRetType) {
	*arg = &val
}

/*
	types and functions for description
*/

// isNotNullableString
type ModelGetDescriptionAttributeType = *string

func getModelGetDescriptionAttributeTypeOk(arg ModelGetDescriptionAttributeType) (ret ModelGetDescriptionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetDescriptionAttributeType(arg *ModelGetDescriptionAttributeType, val ModelGetDescriptionRetType) {
	*arg = &val
}

type ModelGetDescriptionArgType = string
type ModelGetDescriptionRetType = string

/*
	types and functions for displayedName
*/

// isNotNullableString
type ModelGetDisplayedNameAttributeType = *string

func getModelGetDisplayedNameAttributeTypeOk(arg ModelGetDisplayedNameAttributeType) (ret ModelGetDisplayedNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetDisplayedNameAttributeType(arg *ModelGetDisplayedNameAttributeType, val ModelGetDisplayedNameRetType) {
	*arg = &val
}

type ModelGetDisplayedNameArgType = string
type ModelGetDisplayedNameRetType = string

/*
	types and functions for id
*/

// isNotNullableString
type ModelGetIdAttributeType = *string

func getModelGetIdAttributeTypeOk(arg ModelGetIdAttributeType) (ret ModelGetIdRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetIdAttributeType(arg *ModelGetIdAttributeType, val ModelGetIdRetType) {
	*arg = &val
}

type ModelGetIdArgType = string
type ModelGetIdRetType = string

/*
	types and functions for name
*/

// isNotNullableString
type ModelGetNameAttributeType = *string

func getModelGetNameAttributeTypeOk(arg ModelGetNameAttributeType) (ret ModelGetNameRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetNameAttributeType(arg *ModelGetNameAttributeType, val ModelGetNameRetType) {
	*arg = &val
}

type ModelGetNameArgType = string
type ModelGetNameRetType = string

/*
	types and functions for region
*/

// isNotNullableString
type ModelGetRegionAttributeType = *string

func getModelGetRegionAttributeTypeOk(arg ModelGetRegionAttributeType) (ret ModelGetRegionRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetRegionAttributeType(arg *ModelGetRegionAttributeType, val ModelGetRegionRetType) {
	*arg = &val
}

type ModelGetRegionArgType = string
type ModelGetRegionRetType = string

/*
	types and functions for skus
*/

// isArray
type ModelGetSkusAttributeType = *[]SKU
type ModelGetSkusArgType = []SKU
type ModelGetSkusRetType = []SKU

func getModelGetSkusAttributeTypeOk(arg ModelGetSkusAttributeType) (ret ModelGetSkusRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetSkusAttributeType(arg *ModelGetSkusAttributeType, val ModelGetSkusRetType) {
	*arg = &val
}

/*
	types and functions for tags
*/

// isArray
type ModelGetTagsAttributeType = *[]string
type ModelGetTagsArgType = []string
type ModelGetTagsRetType = []string

func getModelGetTagsAttributeTypeOk(arg ModelGetTagsAttributeType) (ret ModelGetTagsRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetTagsAttributeType(arg *ModelGetTagsAttributeType, val ModelGetTagsRetType) {
	*arg = &val
}

/*
	types and functions for type
*/

// isEnumRef
type ModelGetTypeAttributeType = *string
type ModelGetTypeArgType = string
type ModelGetTypeRetType = string

func getModelGetTypeAttributeTypeOk(arg ModelGetTypeAttributeType) (ret ModelGetTypeRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetTypeAttributeType(arg *ModelGetTypeAttributeType, val ModelGetTypeRetType) {
	*arg = &val
}

/*
	types and functions for url
*/

// isNotNullableString
type ModelGetUrlAttributeType = *string

func getModelGetUrlAttributeTypeOk(arg ModelGetUrlAttributeType) (ret ModelGetUrlRetType, ok bool) {
	if arg == nil {
		return ret, false
	}
	return *arg, true
}

func setModelGetUrlAttributeType(arg *ModelGetUrlAttributeType, val ModelGetUrlRetType) {
	*arg = &val
}

type ModelGetUrlArgType = string
type ModelGetUrlRetType = string

// Model struct for Model
type Model struct {
	// REQUIRED
	Category ModelGetCategoryAttributeType `json:"category"`
	// REQUIRED
	Description ModelGetDescriptionAttributeType `json:"description"`
	// REQUIRED
	DisplayedName ModelGetDisplayedNameAttributeType `json:"displayedName"`
	// generated uuid to identify a model
	// REQUIRED
	Id ModelGetIdAttributeType `json:"id"`
	// huggingface name
	// REQUIRED
	Name ModelGetNameAttributeType `json:"name"`
	// REQUIRED
	Region ModelGetRegionAttributeType `json:"region"`
	// REQUIRED
	Skus ModelGetSkusAttributeType `json:"skus"`
	Tags ModelGetTagsAttributeType `json:"tags,omitempty"`
	// REQUIRED
	Type ModelGetTypeAttributeType `json:"type"`
	// url of the model
	// REQUIRED
	Url ModelGetUrlAttributeType `json:"url"`
}

type _Model Model

// NewModel instantiates a new Model object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel(category ModelGetCategoryArgType, description ModelGetDescriptionArgType, displayedName ModelGetDisplayedNameArgType, id ModelGetIdArgType, name ModelGetNameArgType, region ModelGetRegionArgType, skus ModelGetSkusArgType, type_ ModelGetTypeArgType, url ModelGetUrlArgType) *Model {
	this := Model{}
	setModelGetCategoryAttributeType(&this.Category, category)
	setModelGetDescriptionAttributeType(&this.Description, description)
	setModelGetDisplayedNameAttributeType(&this.DisplayedName, displayedName)
	setModelGetIdAttributeType(&this.Id, id)
	setModelGetNameAttributeType(&this.Name, name)
	setModelGetRegionAttributeType(&this.Region, region)
	setModelGetSkusAttributeType(&this.Skus, skus)
	setModelGetTypeAttributeType(&this.Type, type_)
	setModelGetUrlAttributeType(&this.Url, url)
	return &this
}

// NewModelWithDefaults instantiates a new Model object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWithDefaults() *Model {
	this := Model{}
	return &this
}

// GetCategory returns the Category field value
func (o *Model) GetCategory() (ret ModelGetCategoryRetType) {
	ret, _ = o.GetCategoryOk()
	return ret
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *Model) GetCategoryOk() (ret ModelGetCategoryRetType, ok bool) {
	return getModelGetCategoryAttributeTypeOk(o.Category)
}

// SetCategory sets field value
func (o *Model) SetCategory(v ModelGetCategoryRetType) {
	setModelGetCategoryAttributeType(&o.Category, v)
}

// GetDescription returns the Description field value
func (o *Model) GetDescription() (ret ModelGetDescriptionRetType) {
	ret, _ = o.GetDescriptionOk()
	return ret
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Model) GetDescriptionOk() (ret ModelGetDescriptionRetType, ok bool) {
	return getModelGetDescriptionAttributeTypeOk(o.Description)
}

// SetDescription sets field value
func (o *Model) SetDescription(v ModelGetDescriptionRetType) {
	setModelGetDescriptionAttributeType(&o.Description, v)
}

// GetDisplayedName returns the DisplayedName field value
func (o *Model) GetDisplayedName() (ret ModelGetDisplayedNameRetType) {
	ret, _ = o.GetDisplayedNameOk()
	return ret
}

// GetDisplayedNameOk returns a tuple with the DisplayedName field value
// and a boolean to check if the value has been set.
func (o *Model) GetDisplayedNameOk() (ret ModelGetDisplayedNameRetType, ok bool) {
	return getModelGetDisplayedNameAttributeTypeOk(o.DisplayedName)
}

// SetDisplayedName sets field value
func (o *Model) SetDisplayedName(v ModelGetDisplayedNameRetType) {
	setModelGetDisplayedNameAttributeType(&o.DisplayedName, v)
}

// GetId returns the Id field value
func (o *Model) GetId() (ret ModelGetIdRetType) {
	ret, _ = o.GetIdOk()
	return ret
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Model) GetIdOk() (ret ModelGetIdRetType, ok bool) {
	return getModelGetIdAttributeTypeOk(o.Id)
}

// SetId sets field value
func (o *Model) SetId(v ModelGetIdRetType) {
	setModelGetIdAttributeType(&o.Id, v)
}

// GetName returns the Name field value
func (o *Model) GetName() (ret ModelGetNameRetType) {
	ret, _ = o.GetNameOk()
	return ret
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Model) GetNameOk() (ret ModelGetNameRetType, ok bool) {
	return getModelGetNameAttributeTypeOk(o.Name)
}

// SetName sets field value
func (o *Model) SetName(v ModelGetNameRetType) {
	setModelGetNameAttributeType(&o.Name, v)
}

// GetRegion returns the Region field value
func (o *Model) GetRegion() (ret ModelGetRegionRetType) {
	ret, _ = o.GetRegionOk()
	return ret
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Model) GetRegionOk() (ret ModelGetRegionRetType, ok bool) {
	return getModelGetRegionAttributeTypeOk(o.Region)
}

// SetRegion sets field value
func (o *Model) SetRegion(v ModelGetRegionRetType) {
	setModelGetRegionAttributeType(&o.Region, v)
}

// GetSkus returns the Skus field value
func (o *Model) GetSkus() (ret ModelGetSkusRetType) {
	ret, _ = o.GetSkusOk()
	return ret
}

// GetSkusOk returns a tuple with the Skus field value
// and a boolean to check if the value has been set.
func (o *Model) GetSkusOk() (ret ModelGetSkusRetType, ok bool) {
	return getModelGetSkusAttributeTypeOk(o.Skus)
}

// SetSkus sets field value
func (o *Model) SetSkus(v ModelGetSkusRetType) {
	setModelGetSkusAttributeType(&o.Skus, v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Model) GetTags() (res ModelGetTagsRetType) {
	res, _ = o.GetTagsOk()
	return
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model) GetTagsOk() (ret ModelGetTagsRetType, ok bool) {
	return getModelGetTagsAttributeTypeOk(o.Tags)
}

// HasTags returns a boolean if a field has been set.
func (o *Model) HasTags() bool {
	_, ok := o.GetTagsOk()
	return ok
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Model) SetTags(v ModelGetTagsRetType) {
	setModelGetTagsAttributeType(&o.Tags, v)
}

// GetType returns the Type field value
func (o *Model) GetType() (ret ModelGetTypeRetType) {
	ret, _ = o.GetTypeOk()
	return ret
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Model) GetTypeOk() (ret ModelGetTypeRetType, ok bool) {
	return getModelGetTypeAttributeTypeOk(o.Type)
}

// SetType sets field value
func (o *Model) SetType(v ModelGetTypeRetType) {
	setModelGetTypeAttributeType(&o.Type, v)
}

// GetUrl returns the Url field value
func (o *Model) GetUrl() (ret ModelGetUrlRetType) {
	ret, _ = o.GetUrlOk()
	return ret
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Model) GetUrlOk() (ret ModelGetUrlRetType, ok bool) {
	return getModelGetUrlAttributeTypeOk(o.Url)
}

// SetUrl sets field value
func (o *Model) SetUrl(v ModelGetUrlRetType) {
	setModelGetUrlAttributeType(&o.Url, v)
}

func (o Model) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if val, ok := getModelGetCategoryAttributeTypeOk(o.Category); ok {
		toSerialize["Category"] = val
	}
	if val, ok := getModelGetDescriptionAttributeTypeOk(o.Description); ok {
		toSerialize["Description"] = val
	}
	if val, ok := getModelGetDisplayedNameAttributeTypeOk(o.DisplayedName); ok {
		toSerialize["DisplayedName"] = val
	}
	if val, ok := getModelGetIdAttributeTypeOk(o.Id); ok {
		toSerialize["Id"] = val
	}
	if val, ok := getModelGetNameAttributeTypeOk(o.Name); ok {
		toSerialize["Name"] = val
	}
	if val, ok := getModelGetRegionAttributeTypeOk(o.Region); ok {
		toSerialize["Region"] = val
	}
	if val, ok := getModelGetSkusAttributeTypeOk(o.Skus); ok {
		toSerialize["Skus"] = val
	}
	if val, ok := getModelGetTagsAttributeTypeOk(o.Tags); ok {
		toSerialize["Tags"] = val
	}
	if val, ok := getModelGetTypeAttributeTypeOk(o.Type); ok {
		toSerialize["Type"] = val
	}
	if val, ok := getModelGetUrlAttributeTypeOk(o.Url); ok {
		toSerialize["Url"] = val
	}
	return toSerialize, nil
}

type NullableModel struct {
	value *Model
	isSet bool
}

func (v NullableModel) Get() *Model {
	return v.value
}

func (v *NullableModel) Set(val *Model) {
	v.value = val
	v.isSet = true
}

func (v NullableModel) IsSet() bool {
	return v.isSet
}

func (v *NullableModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel(val *Model) *NullableModel {
	return &NullableModel{value: val, isSet: true}
}

func (v NullableModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
