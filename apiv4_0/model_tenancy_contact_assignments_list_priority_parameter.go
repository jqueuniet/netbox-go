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

// TenancyContactAssignmentsListPriorityParameter the model 'TenancyContactAssignmentsListPriorityParameter'
type TenancyContactAssignmentsListPriorityParameter string

// List of tenancy_contact_assignments_list_priority_parameter
const (
	TENANCYCONTACTASSIGNMENTSLISTPRIORITYPARAMETER_INACTIVE  TenancyContactAssignmentsListPriorityParameter = "inactive"
	TENANCYCONTACTASSIGNMENTSLISTPRIORITYPARAMETER_PRIMARY   TenancyContactAssignmentsListPriorityParameter = "primary"
	TENANCYCONTACTASSIGNMENTSLISTPRIORITYPARAMETER_SECONDARY TenancyContactAssignmentsListPriorityParameter = "secondary"
	TENANCYCONTACTASSIGNMENTSLISTPRIORITYPARAMETER_TERTIARY  TenancyContactAssignmentsListPriorityParameter = "tertiary"
)

// All allowed values of TenancyContactAssignmentsListPriorityParameter enum
var AllowedTenancyContactAssignmentsListPriorityParameterEnumValues = []TenancyContactAssignmentsListPriorityParameter{
	"inactive",
	"primary",
	"secondary",
	"tertiary",
}

func (v *TenancyContactAssignmentsListPriorityParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TenancyContactAssignmentsListPriorityParameter(value)
	for _, existing := range AllowedTenancyContactAssignmentsListPriorityParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TenancyContactAssignmentsListPriorityParameter", value)
}

// NewTenancyContactAssignmentsListPriorityParameterFromValue returns a pointer to a valid TenancyContactAssignmentsListPriorityParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenancyContactAssignmentsListPriorityParameterFromValue(v string) (*TenancyContactAssignmentsListPriorityParameter, error) {
	ev := TenancyContactAssignmentsListPriorityParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TenancyContactAssignmentsListPriorityParameter: valid values are %v", v, AllowedTenancyContactAssignmentsListPriorityParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenancyContactAssignmentsListPriorityParameter) IsValid() bool {
	for _, existing := range AllowedTenancyContactAssignmentsListPriorityParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tenancy_contact_assignments_list_priority_parameter value
func (v TenancyContactAssignmentsListPriorityParameter) Ptr() *TenancyContactAssignmentsListPriorityParameter {
	return &v
}

type NullableTenancyContactAssignmentsListPriorityParameter struct {
	value *TenancyContactAssignmentsListPriorityParameter
	isSet bool
}

func (v NullableTenancyContactAssignmentsListPriorityParameter) Get() *TenancyContactAssignmentsListPriorityParameter {
	return v.value
}

func (v *NullableTenancyContactAssignmentsListPriorityParameter) Set(val *TenancyContactAssignmentsListPriorityParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTenancyContactAssignmentsListPriorityParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTenancyContactAssignmentsListPriorityParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenancyContactAssignmentsListPriorityParameter(val *TenancyContactAssignmentsListPriorityParameter) *NullableTenancyContactAssignmentsListPriorityParameter {
	return &NullableTenancyContactAssignmentsListPriorityParameter{value: val, isSet: true}
}

func (v NullableTenancyContactAssignmentsListPriorityParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenancyContactAssignmentsListPriorityParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}