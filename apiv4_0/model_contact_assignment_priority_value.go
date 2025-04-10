/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv4_0

import (
	"encoding/json"
	"fmt"
)

// ContactAssignmentPriorityValue * `primary` - Primary * `secondary` - Secondary * `tertiary` - Tertiary * `inactive` - Inactive
type ContactAssignmentPriorityValue string

// List of ContactAssignment_priority_value
const (
	CONTACTASSIGNMENTPRIORITYVALUE_PRIMARY   ContactAssignmentPriorityValue = "primary"
	CONTACTASSIGNMENTPRIORITYVALUE_SECONDARY ContactAssignmentPriorityValue = "secondary"
	CONTACTASSIGNMENTPRIORITYVALUE_TERTIARY  ContactAssignmentPriorityValue = "tertiary"
	CONTACTASSIGNMENTPRIORITYVALUE_INACTIVE  ContactAssignmentPriorityValue = "inactive"
	CONTACTASSIGNMENTPRIORITYVALUE_EMPTY     ContactAssignmentPriorityValue = ""
)

// All allowed values of ContactAssignmentPriorityValue enum
var AllowedContactAssignmentPriorityValueEnumValues = []ContactAssignmentPriorityValue{
	"primary",
	"secondary",
	"tertiary",
	"inactive",
	"",
}

func (v *ContactAssignmentPriorityValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContactAssignmentPriorityValue(value)
	for _, existing := range AllowedContactAssignmentPriorityValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContactAssignmentPriorityValue", value)
}

// NewContactAssignmentPriorityValueFromValue returns a pointer to a valid ContactAssignmentPriorityValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContactAssignmentPriorityValueFromValue(v string) (*ContactAssignmentPriorityValue, error) {
	ev := ContactAssignmentPriorityValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContactAssignmentPriorityValue: valid values are %v", v, AllowedContactAssignmentPriorityValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContactAssignmentPriorityValue) IsValid() bool {
	for _, existing := range AllowedContactAssignmentPriorityValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContactAssignment_priority_value value
func (v ContactAssignmentPriorityValue) Ptr() *ContactAssignmentPriorityValue {
	return &v
}

type NullableContactAssignmentPriorityValue struct {
	value *ContactAssignmentPriorityValue
	isSet bool
}

func (v NullableContactAssignmentPriorityValue) Get() *ContactAssignmentPriorityValue {
	return v.value
}

func (v *NullableContactAssignmentPriorityValue) Set(val *ContactAssignmentPriorityValue) {
	v.value = val
	v.isSet = true
}

func (v NullableContactAssignmentPriorityValue) IsSet() bool {
	return v.isSet
}

func (v *NullableContactAssignmentPriorityValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactAssignmentPriorityValue(val *ContactAssignmentPriorityValue) *NullableContactAssignmentPriorityValue {
	return &NullableContactAssignmentPriorityValue{value: val, isSet: true}
}

func (v NullableContactAssignmentPriorityValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactAssignmentPriorityValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
