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

// WirelessLinkRequestDistanceUnit * `km` - Kilometers * `m` - Meters * `mi` - Miles * `ft` - Feet
type WirelessLinkRequestDistanceUnit string

// List of WirelessLinkRequest_distance_unit
const (
	WIRELESSLINKREQUESTDISTANCEUNIT_KM    WirelessLinkRequestDistanceUnit = "km"
	WIRELESSLINKREQUESTDISTANCEUNIT_M     WirelessLinkRequestDistanceUnit = "m"
	WIRELESSLINKREQUESTDISTANCEUNIT_MI    WirelessLinkRequestDistanceUnit = "mi"
	WIRELESSLINKREQUESTDISTANCEUNIT_FT    WirelessLinkRequestDistanceUnit = "ft"
	WIRELESSLINKREQUESTDISTANCEUNIT_EMPTY WirelessLinkRequestDistanceUnit = ""
)

// All allowed values of WirelessLinkRequestDistanceUnit enum
var AllowedWirelessLinkRequestDistanceUnitEnumValues = []WirelessLinkRequestDistanceUnit{
	"km",
	"m",
	"mi",
	"ft",
	"",
}

func (v *WirelessLinkRequestDistanceUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WirelessLinkRequestDistanceUnit(value)
	for _, existing := range AllowedWirelessLinkRequestDistanceUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WirelessLinkRequestDistanceUnit", value)
}

// NewWirelessLinkRequestDistanceUnitFromValue returns a pointer to a valid WirelessLinkRequestDistanceUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWirelessLinkRequestDistanceUnitFromValue(v string) (*WirelessLinkRequestDistanceUnit, error) {
	ev := WirelessLinkRequestDistanceUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WirelessLinkRequestDistanceUnit: valid values are %v", v, AllowedWirelessLinkRequestDistanceUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WirelessLinkRequestDistanceUnit) IsValid() bool {
	for _, existing := range AllowedWirelessLinkRequestDistanceUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WirelessLinkRequest_distance_unit value
func (v WirelessLinkRequestDistanceUnit) Ptr() *WirelessLinkRequestDistanceUnit {
	return &v
}

type NullableWirelessLinkRequestDistanceUnit struct {
	value *WirelessLinkRequestDistanceUnit
	isSet bool
}

func (v NullableWirelessLinkRequestDistanceUnit) Get() *WirelessLinkRequestDistanceUnit {
	return v.value
}

func (v *NullableWirelessLinkRequestDistanceUnit) Set(val *WirelessLinkRequestDistanceUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLinkRequestDistanceUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLinkRequestDistanceUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLinkRequestDistanceUnit(val *WirelessLinkRequestDistanceUnit) *NullableWirelessLinkRequestDistanceUnit {
	return &NullableWirelessLinkRequestDistanceUnit{value: val, isSet: true}
}

func (v NullableWirelessLinkRequestDistanceUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLinkRequestDistanceUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
