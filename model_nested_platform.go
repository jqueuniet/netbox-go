/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.1 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the NestedPlatform type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedPlatform{}

// NestedPlatform Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedPlatform struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NewNestedPlatform instantiates a new NestedPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedPlatform(id int32, url string, display string, name string, slug string) *NestedPlatform {
	this := NestedPlatform{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedPlatformWithDefaults instantiates a new NestedPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedPlatformWithDefaults() *NestedPlatform {
	this := NestedPlatform{}
	return &this
}

// GetId returns the Id field value
func (o *NestedPlatform) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedPlatform) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedPlatform) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedPlatform) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedPlatform) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedPlatform) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedPlatform) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedPlatform) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedPlatform) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedPlatform) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedPlatform) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedPlatform) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedPlatform) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedPlatform) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedPlatform) SetSlug(v string) {
	o.Slug = v
}

func (o NestedPlatform) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedPlatform) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	return toSerialize, nil
}

type NullableNestedPlatform struct {
	value *NestedPlatform
	isSet bool
}

func (v NullableNestedPlatform) Get() *NestedPlatform {
	return v.value
}

func (v *NullableNestedPlatform) Set(val *NestedPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedPlatform(val *NestedPlatform) *NullableNestedPlatform {
	return &NullableNestedPlatform{value: val, isSet: true}
}

func (v NullableNestedPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

