/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.9 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3_5

import (
	"encoding/json"
	"fmt"
)

// checks if the NestedRackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedRackRequest{}

// NestedRackRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedRackRequest struct {
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedRackRequest NestedRackRequest

// NewNestedRackRequest instantiates a new NestedRackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRackRequest(name string) *NestedRackRequest {
	this := NestedRackRequest{}
	this.Name = name
	return &this
}

// NewNestedRackRequestWithDefaults instantiates a new NestedRackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRackRequestWithDefaults() *NestedRackRequest {
	this := NestedRackRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedRackRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedRackRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedRackRequest) SetName(v string) {
	o.Name = v
}

func (o NestedRackRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedRackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedRackRequest) UnmarshalJSON(data []byte) (err error) {
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
	varNestedRackRequest := _NestedRackRequest{}

	err = json.Unmarshal(data, &varNestedRackRequest)

	if err != nil {
		return err
	}

	*o = NestedRackRequest(varNestedRackRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedRackRequest struct {
	value *NestedRackRequest
	isSet bool
}

func (v NullableNestedRackRequest) Get() *NestedRackRequest {
	return v.value
}

func (v *NullableNestedRackRequest) Set(val *NestedRackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRackRequest(val *NestedRackRequest) *NullableNestedRackRequest {
	return &NullableNestedRackRequest{value: val, isSet: true}
}

func (v NullableNestedRackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}