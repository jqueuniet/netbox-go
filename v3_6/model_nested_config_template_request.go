/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3_6

import (
	"encoding/json"
	"fmt"
)

// checks if the NestedConfigTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedConfigTemplateRequest{}

// NestedConfigTemplateRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedConfigTemplateRequest struct {
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedConfigTemplateRequest NestedConfigTemplateRequest

// NewNestedConfigTemplateRequest instantiates a new NestedConfigTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedConfigTemplateRequest(name string) *NestedConfigTemplateRequest {
	this := NestedConfigTemplateRequest{}
	this.Name = name
	return &this
}

// NewNestedConfigTemplateRequestWithDefaults instantiates a new NestedConfigTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedConfigTemplateRequestWithDefaults() *NestedConfigTemplateRequest {
	this := NestedConfigTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedConfigTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedConfigTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedConfigTemplateRequest) SetName(v string) {
	o.Name = v
}

func (o NestedConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedConfigTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedConfigTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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
	varNestedConfigTemplateRequest := _NestedConfigTemplateRequest{}

	err = json.Unmarshal(data, &varNestedConfigTemplateRequest)

	if err != nil {
		return err
	}

	*o = NestedConfigTemplateRequest(varNestedConfigTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedConfigTemplateRequest struct {
	value *NestedConfigTemplateRequest
	isSet bool
}

func (v NullableNestedConfigTemplateRequest) Get() *NestedConfigTemplateRequest {
	return v.value
}

func (v *NullableNestedConfigTemplateRequest) Set(val *NestedConfigTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedConfigTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedConfigTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedConfigTemplateRequest(val *NestedConfigTemplateRequest) *NullableNestedConfigTemplateRequest {
	return &NullableNestedConfigTemplateRequest{value: val, isSet: true}
}

func (v NullableNestedConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedConfigTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}