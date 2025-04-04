/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv3_6

import (
	"encoding/json"
	"fmt"
)

// checks if the NestedRackRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedRackRoleRequest{}

// NestedRackRoleRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedRackRoleRequest struct {
	Name                 string `json:"name"`
	Slug                 string `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	AdditionalProperties map[string]interface{}
}

type _NestedRackRoleRequest NestedRackRoleRequest

// NewNestedRackRoleRequest instantiates a new NestedRackRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRackRoleRequest(name string, slug string) *NestedRackRoleRequest {
	this := NestedRackRoleRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedRackRoleRequestWithDefaults instantiates a new NestedRackRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRackRoleRequestWithDefaults() *NestedRackRoleRequest {
	this := NestedRackRoleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedRackRoleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedRackRoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedRackRoleRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedRackRoleRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedRackRoleRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedRackRoleRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedRackRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedRackRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedRackRoleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varNestedRackRoleRequest := _NestedRackRoleRequest{}

	err = json.Unmarshal(data, &varNestedRackRoleRequest)

	if err != nil {
		return err
	}

	*o = NestedRackRoleRequest(varNestedRackRoleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedRackRoleRequest struct {
	value *NestedRackRoleRequest
	isSet bool
}

func (v NullableNestedRackRoleRequest) Get() *NestedRackRoleRequest {
	return v.value
}

func (v *NullableNestedRackRoleRequest) Set(val *NestedRackRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRackRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRackRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRackRoleRequest(val *NestedRackRoleRequest) *NullableNestedRackRoleRequest {
	return &NullableNestedRackRoleRequest{value: val, isSet: true}
}

func (v NullableNestedRackRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRackRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
