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

// PatchedWritableCustomFieldRequestUiVisibility Specifies the visibility of custom field in the UI  * `read-write` - Read/write * `read-only` - Read-only * `hidden` - Hidden * `hidden-ifunset` - Hidden (if unset)
type PatchedWritableCustomFieldRequestUiVisibility string

// List of PatchedWritableCustomFieldRequest_ui_visibility
const (
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIVISIBILITY_READ_WRITE     PatchedWritableCustomFieldRequestUiVisibility = "read-write"
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIVISIBILITY_READ_ONLY      PatchedWritableCustomFieldRequestUiVisibility = "read-only"
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIVISIBILITY_HIDDEN         PatchedWritableCustomFieldRequestUiVisibility = "hidden"
	PATCHEDWRITABLECUSTOMFIELDREQUESTUIVISIBILITY_HIDDEN_IFUNSET PatchedWritableCustomFieldRequestUiVisibility = "hidden-ifunset"
)

// All allowed values of PatchedWritableCustomFieldRequestUiVisibility enum
var AllowedPatchedWritableCustomFieldRequestUiVisibilityEnumValues = []PatchedWritableCustomFieldRequestUiVisibility{
	"read-write",
	"read-only",
	"hidden",
	"hidden-ifunset",
}

func (v *PatchedWritableCustomFieldRequestUiVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableCustomFieldRequestUiVisibility(value)
	for _, existing := range AllowedPatchedWritableCustomFieldRequestUiVisibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableCustomFieldRequestUiVisibility", value)
}

// NewPatchedWritableCustomFieldRequestUiVisibilityFromValue returns a pointer to a valid PatchedWritableCustomFieldRequestUiVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableCustomFieldRequestUiVisibilityFromValue(v string) (*PatchedWritableCustomFieldRequestUiVisibility, error) {
	ev := PatchedWritableCustomFieldRequestUiVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableCustomFieldRequestUiVisibility: valid values are %v", v, AllowedPatchedWritableCustomFieldRequestUiVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableCustomFieldRequestUiVisibility) IsValid() bool {
	for _, existing := range AllowedPatchedWritableCustomFieldRequestUiVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableCustomFieldRequest_ui_visibility value
func (v PatchedWritableCustomFieldRequestUiVisibility) Ptr() *PatchedWritableCustomFieldRequestUiVisibility {
	return &v
}

type NullablePatchedWritableCustomFieldRequestUiVisibility struct {
	value *PatchedWritableCustomFieldRequestUiVisibility
	isSet bool
}

func (v NullablePatchedWritableCustomFieldRequestUiVisibility) Get() *PatchedWritableCustomFieldRequestUiVisibility {
	return v.value
}

func (v *NullablePatchedWritableCustomFieldRequestUiVisibility) Set(val *PatchedWritableCustomFieldRequestUiVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCustomFieldRequestUiVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCustomFieldRequestUiVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCustomFieldRequestUiVisibility(val *PatchedWritableCustomFieldRequestUiVisibility) *NullablePatchedWritableCustomFieldRequestUiVisibility {
	return &NullablePatchedWritableCustomFieldRequestUiVisibility{value: val, isSet: true}
}

func (v NullablePatchedWritableCustomFieldRequestUiVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCustomFieldRequestUiVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}