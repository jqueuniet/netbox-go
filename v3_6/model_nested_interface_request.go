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

// checks if the NestedInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedInterfaceRequest{}

// NestedInterfaceRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedInterfaceRequest struct {
	Name                 string        `json:"name"`
	Cable                NullableInt32 `json:"cable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedInterfaceRequest NestedInterfaceRequest

// NewNestedInterfaceRequest instantiates a new NestedInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedInterfaceRequest(name string) *NestedInterfaceRequest {
	this := NestedInterfaceRequest{}
	this.Name = name
	return &this
}

// NewNestedInterfaceRequestWithDefaults instantiates a new NestedInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedInterfaceRequestWithDefaults() *NestedInterfaceRequest {
	this := NestedInterfaceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedInterfaceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedInterfaceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedInterfaceRequest) SetName(v string) {
	o.Name = v
}

// GetCable returns the Cable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedInterfaceRequest) GetCable() int32 {
	if o == nil || IsNil(o.Cable.Get()) {
		var ret int32
		return ret
	}
	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedInterfaceRequest) GetCableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// HasCable returns a boolean if a field has been set.
func (o *NestedInterfaceRequest) HasCable() bool {
	if o != nil && o.Cable.IsSet() {
		return true
	}

	return false
}

// SetCable gets a reference to the given NullableInt32 and assigns it to the Cable field.
func (o *NestedInterfaceRequest) SetCable(v int32) {
	o.Cable.Set(&v)
}

// SetCableNil sets the value for Cable to be an explicit nil
func (o *NestedInterfaceRequest) SetCableNil() {
	o.Cable.Set(nil)
}

// UnsetCable ensures that no value is present for Cable, not even an explicit nil
func (o *NestedInterfaceRequest) UnsetCable() {
	o.Cable.Unset()
}

func (o NestedInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedInterfaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Cable.IsSet() {
		toSerialize["cable"] = o.Cable.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedInterfaceRequest) UnmarshalJSON(data []byte) (err error) {
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
	varNestedInterfaceRequest := _NestedInterfaceRequest{}

	err = json.Unmarshal(data, &varNestedInterfaceRequest)

	if err != nil {
		return err
	}

	*o = NestedInterfaceRequest(varNestedInterfaceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "cable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedInterfaceRequest struct {
	value *NestedInterfaceRequest
	isSet bool
}

func (v NullableNestedInterfaceRequest) Get() *NestedInterfaceRequest {
	return v.value
}

func (v *NullableNestedInterfaceRequest) Set(val *NestedInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedInterfaceRequest(val *NestedInterfaceRequest) *NullableNestedInterfaceRequest {
	return &NullableNestedInterfaceRequest{value: val, isSet: true}
}

func (v NullableNestedInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}