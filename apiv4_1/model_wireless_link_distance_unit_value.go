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

// WirelessLinkDistanceUnitValue * `km` - Kilometers * `m` - Meters * `mi` - Miles * `ft` - Feet
type WirelessLinkDistanceUnitValue string

// List of WirelessLink_distance_unit_value
const (
	WIRELESSLINKDISTANCEUNITVALUE_KM    WirelessLinkDistanceUnitValue = "km"
	WIRELESSLINKDISTANCEUNITVALUE_M     WirelessLinkDistanceUnitValue = "m"
	WIRELESSLINKDISTANCEUNITVALUE_MI    WirelessLinkDistanceUnitValue = "mi"
	WIRELESSLINKDISTANCEUNITVALUE_FT    WirelessLinkDistanceUnitValue = "ft"
	WIRELESSLINKDISTANCEUNITVALUE_EMPTY WirelessLinkDistanceUnitValue = ""
)

// All allowed values of WirelessLinkDistanceUnitValue enum
var AllowedWirelessLinkDistanceUnitValueEnumValues = []WirelessLinkDistanceUnitValue{
	"km",
	"m",
	"mi",
	"ft",
	"",
}

func (v *WirelessLinkDistanceUnitValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessLinkDistanceUnitValue(value)
	for _, existing := range AllowedWirelessLinkDistanceUnitValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessLinkDistanceUnitValue", value)
}

// NewWirelessLinkDistanceUnitValueFromValue returns a pointer to a valid WirelessLinkDistanceUnitValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessLinkDistanceUnitValueFromValue(v string) (*WirelessLinkDistanceUnitValue, error) {
	ev := WirelessLinkDistanceUnitValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessLinkDistanceUnitValue: valid values are %v", v, AllowedWirelessLinkDistanceUnitValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessLinkDistanceUnitValue) IsValid() bool {
	for _, existing := range AllowedWirelessLinkDistanceUnitValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WirelessLink_distance_unit_value value
func (v WirelessLinkDistanceUnitValue) Ptr() *WirelessLinkDistanceUnitValue {
	return &v
}

type NullableWirelessLinkDistanceUnitValue struct {
	value *WirelessLinkDistanceUnitValue
	isSet bool
}

func (v NullableWirelessLinkDistanceUnitValue) Get() *WirelessLinkDistanceUnitValue {
	return v.value
}

func (v *NullableWirelessLinkDistanceUnitValue) Set(val *WirelessLinkDistanceUnitValue) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLinkDistanceUnitValue) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLinkDistanceUnitValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLinkDistanceUnitValue(val *WirelessLinkDistanceUnitValue) *NullableWirelessLinkDistanceUnitValue {
	return &NullableWirelessLinkDistanceUnitValue{value: val, isSet: true}
}

func (v NullableWirelessLinkDistanceUnitValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLinkDistanceUnitValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}