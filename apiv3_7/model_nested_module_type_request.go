/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.8 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv3_7

import (
	"encoding/json"
	"fmt"
)

// checks if the NestedModuleTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedModuleTypeRequest{}

// NestedModuleTypeRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedModuleTypeRequest struct {
	Model                string `json:"model"`
	AdditionalProperties map[string]interface{}
}

type _NestedModuleTypeRequest NestedModuleTypeRequest

// NewNestedModuleTypeRequest instantiates a new NestedModuleTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedModuleTypeRequest(model string) *NestedModuleTypeRequest {
	this := NestedModuleTypeRequest{}
	this.Model = model
	return &this
}

// NewNestedModuleTypeRequestWithDefaults instantiates a new NestedModuleTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedModuleTypeRequestWithDefaults() *NestedModuleTypeRequest {
	this := NestedModuleTypeRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *NestedModuleTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *NestedModuleTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *NestedModuleTypeRequest) SetModel(v string) {
	o.Model = v
}

func (o NestedModuleTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedModuleTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedModuleTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
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
	varNestedModuleTypeRequest := _NestedModuleTypeRequest{}

	err = json.Unmarshal(data, &varNestedModuleTypeRequest)

	if err != nil {
		return err
	}

	*o = NestedModuleTypeRequest(varNestedModuleTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "model")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedModuleTypeRequest struct {
	value *NestedModuleTypeRequest
	isSet bool
}

func (v NullableNestedModuleTypeRequest) Get() *NestedModuleTypeRequest {
	return v.value
}

func (v *NullableNestedModuleTypeRequest) Set(val *NestedModuleTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedModuleTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedModuleTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedModuleTypeRequest(val *NestedModuleTypeRequest) *NullableNestedModuleTypeRequest {
	return &NullableNestedModuleTypeRequest{value: val, isSet: true}
}

func (v NullableNestedModuleTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedModuleTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
