/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.6 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv4_1

import (
	"encoding/json"
	"fmt"
)

// BriefCircuitGroupAssignmentSerializerPriorityValue * `primary` - Primary * `secondary` - Secondary * `tertiary` - Tertiary * `inactive` - Inactive
type BriefCircuitGroupAssignmentSerializerPriorityValue string

// List of BriefCircuitGroupAssignmentSerializer__priority_value
const (
	BRIEFCIRCUITGROUPASSIGNMENTSERIALIZERPRIORITYVALUE_PRIMARY   BriefCircuitGroupAssignmentSerializerPriorityValue = "primary"
	BRIEFCIRCUITGROUPASSIGNMENTSERIALIZERPRIORITYVALUE_SECONDARY BriefCircuitGroupAssignmentSerializerPriorityValue = "secondary"
	BRIEFCIRCUITGROUPASSIGNMENTSERIALIZERPRIORITYVALUE_TERTIARY  BriefCircuitGroupAssignmentSerializerPriorityValue = "tertiary"
	BRIEFCIRCUITGROUPASSIGNMENTSERIALIZERPRIORITYVALUE_INACTIVE  BriefCircuitGroupAssignmentSerializerPriorityValue = "inactive"
	BRIEFCIRCUITGROUPASSIGNMENTSERIALIZERPRIORITYVALUE_EMPTY     BriefCircuitGroupAssignmentSerializerPriorityValue = ""
)

// All allowed values of BriefCircuitGroupAssignmentSerializerPriorityValue enum
var AllowedBriefCircuitGroupAssignmentSerializerPriorityValueEnumValues = []BriefCircuitGroupAssignmentSerializerPriorityValue{
	"primary",
	"secondary",
	"tertiary",
	"inactive",
	"",
}

func (v *BriefCircuitGroupAssignmentSerializerPriorityValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BriefCircuitGroupAssignmentSerializerPriorityValue(value)
	for _, existing := range AllowedBriefCircuitGroupAssignmentSerializerPriorityValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BriefCircuitGroupAssignmentSerializerPriorityValue", value)
}

// NewBriefCircuitGroupAssignmentSerializerPriorityValueFromValue returns a pointer to a valid BriefCircuitGroupAssignmentSerializerPriorityValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBriefCircuitGroupAssignmentSerializerPriorityValueFromValue(v string) (*BriefCircuitGroupAssignmentSerializerPriorityValue, error) {
	ev := BriefCircuitGroupAssignmentSerializerPriorityValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BriefCircuitGroupAssignmentSerializerPriorityValue: valid values are %v", v, AllowedBriefCircuitGroupAssignmentSerializerPriorityValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BriefCircuitGroupAssignmentSerializerPriorityValue) IsValid() bool {
	for _, existing := range AllowedBriefCircuitGroupAssignmentSerializerPriorityValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BriefCircuitGroupAssignmentSerializer__priority_value value
func (v BriefCircuitGroupAssignmentSerializerPriorityValue) Ptr() *BriefCircuitGroupAssignmentSerializerPriorityValue {
	return &v
}

type NullableBriefCircuitGroupAssignmentSerializerPriorityValue struct {
	value *BriefCircuitGroupAssignmentSerializerPriorityValue
	isSet bool
}

func (v NullableBriefCircuitGroupAssignmentSerializerPriorityValue) Get() *BriefCircuitGroupAssignmentSerializerPriorityValue {
	return v.value
}

func (v *NullableBriefCircuitGroupAssignmentSerializerPriorityValue) Set(val *BriefCircuitGroupAssignmentSerializerPriorityValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefCircuitGroupAssignmentSerializerPriorityValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefCircuitGroupAssignmentSerializerPriorityValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefCircuitGroupAssignmentSerializerPriorityValue(val *BriefCircuitGroupAssignmentSerializerPriorityValue) *NullableBriefCircuitGroupAssignmentSerializerPriorityValue {
	return &NullableBriefCircuitGroupAssignmentSerializerPriorityValue{value: val, isSet: true}
}

func (v NullableBriefCircuitGroupAssignmentSerializerPriorityValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefCircuitGroupAssignmentSerializerPriorityValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}