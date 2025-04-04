/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.11 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv4_1

import (
	"encoding/json"
	"fmt"
)

// Termination1 * `A` - A * `Z` - Z
type Termination1 string

// List of Termination_1
const (
	TERMINATION1_A Termination1 = "A"
	TERMINATION1_Z Termination1 = "Z"
)

// All allowed values of Termination1 enum
var AllowedTermination1EnumValues = []Termination1{
	"A",
	"Z",
}

func (v *Termination1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Termination1(value)
	for _, existing := range AllowedTermination1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Termination1", value)
}

// NewTermination1FromValue returns a pointer to a valid Termination1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTermination1FromValue(v string) (*Termination1, error) {
	ev := Termination1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Termination1: valid values are %v", v, AllowedTermination1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Termination1) IsValid() bool {
	for _, existing := range AllowedTermination1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Termination_1 value
func (v Termination1) Ptr() *Termination1 {
	return &v
}

type NullableTermination1 struct {
	value *Termination1
	isSet bool
}

func (v NullableTermination1) Get() *Termination1 {
	return v.value
}

func (v *NullableTermination1) Set(val *Termination1) {
	v.value = val
	v.isSet = true
}

func (v NullableTermination1) IsSet() bool {
	return v.isSet
}

func (v *NullableTermination1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermination1(val *Termination1) *NullableTermination1 {
	return &NullableTermination1{value: val, isSet: true}
}

func (v NullableTermination1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermination1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
