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

// checks if the NestedClusterGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedClusterGroupRequest{}

// NestedClusterGroupRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedClusterGroupRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NewNestedClusterGroupRequest instantiates a new NestedClusterGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedClusterGroupRequest(name string, slug string) *NestedClusterGroupRequest {
	this := NestedClusterGroupRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedClusterGroupRequestWithDefaults instantiates a new NestedClusterGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedClusterGroupRequestWithDefaults() *NestedClusterGroupRequest {
	this := NestedClusterGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedClusterGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedClusterGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedClusterGroupRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedClusterGroupRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedClusterGroupRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedClusterGroupRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedClusterGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedClusterGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	return toSerialize, nil
}

type NullableNestedClusterGroupRequest struct {
	value *NestedClusterGroupRequest
	isSet bool
}

func (v NullableNestedClusterGroupRequest) Get() *NestedClusterGroupRequest {
	return v.value
}

func (v *NullableNestedClusterGroupRequest) Set(val *NestedClusterGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedClusterGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedClusterGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedClusterGroupRequest(val *NestedClusterGroupRequest) *NullableNestedClusterGroupRequest {
	return &NullableNestedClusterGroupRequest{value: val, isSet: true}
}

func (v NullableNestedClusterGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedClusterGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

