/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.9 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv3_5

import (
	"encoding/json"
	"fmt"
)

// IPAddressStatusValue * `active` - Active * `reserved` - Reserved * `deprecated` - Deprecated * `dhcp` - DHCP * `slaac` - SLAAC
type IPAddressStatusValue string

// List of IPAddress_status_value
const (
	IPADDRESSSTATUSVALUE_ACTIVE     IPAddressStatusValue = "active"
	IPADDRESSSTATUSVALUE_RESERVED   IPAddressStatusValue = "reserved"
	IPADDRESSSTATUSVALUE_DEPRECATED IPAddressStatusValue = "deprecated"
	IPADDRESSSTATUSVALUE_DHCP       IPAddressStatusValue = "dhcp"
	IPADDRESSSTATUSVALUE_SLAAC      IPAddressStatusValue = "slaac"
)

// All allowed values of IPAddressStatusValue enum
var AllowedIPAddressStatusValueEnumValues = []IPAddressStatusValue{
	"active",
	"reserved",
	"deprecated",
	"dhcp",
	"slaac",
}

func (v *IPAddressStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IPAddressStatusValue(value)
	for _, existing := range AllowedIPAddressStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IPAddressStatusValue", value)
}

// NewIPAddressStatusValueFromValue returns a pointer to a valid IPAddressStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIPAddressStatusValueFromValue(v string) (*IPAddressStatusValue, error) {
	ev := IPAddressStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IPAddressStatusValue: valid values are %v", v, AllowedIPAddressStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IPAddressStatusValue) IsValid() bool {
	for _, existing := range AllowedIPAddressStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IPAddress_status_value value
func (v IPAddressStatusValue) Ptr() *IPAddressStatusValue {
	return &v
}

type NullableIPAddressStatusValue struct {
	value *IPAddressStatusValue
	isSet bool
}

func (v NullableIPAddressStatusValue) Get() *IPAddressStatusValue {
	return v.value
}

func (v *NullableIPAddressStatusValue) Set(val *IPAddressStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAddressStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAddressStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAddressStatusValue(val *IPAddressStatusValue) *NullableIPAddressStatusValue {
	return &NullableIPAddressStatusValue{value: val, isSet: true}
}

func (v NullableIPAddressStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAddressStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}