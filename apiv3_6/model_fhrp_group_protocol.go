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

// FHRPGroupProtocol * `vrrp2` - VRRPv2 * `vrrp3` - VRRPv3 * `carp` - CARP * `clusterxl` - ClusterXL * `hsrp` - HSRP * `glbp` - GLBP * `other` - Other
type FHRPGroupProtocol string

// List of FHRPGroup_protocol
const (
	FHRPGROUPPROTOCOL_VRRP2     FHRPGroupProtocol = "vrrp2"
	FHRPGROUPPROTOCOL_VRRP3     FHRPGroupProtocol = "vrrp3"
	FHRPGROUPPROTOCOL_CARP      FHRPGroupProtocol = "carp"
	FHRPGROUPPROTOCOL_CLUSTERXL FHRPGroupProtocol = "clusterxl"
	FHRPGROUPPROTOCOL_HSRP      FHRPGroupProtocol = "hsrp"
	FHRPGROUPPROTOCOL_GLBP      FHRPGroupProtocol = "glbp"
	FHRPGROUPPROTOCOL_OTHER     FHRPGroupProtocol = "other"
)

// All allowed values of FHRPGroupProtocol enum
var AllowedFHRPGroupProtocolEnumValues = []FHRPGroupProtocol{
	"vrrp2",
	"vrrp3",
	"carp",
	"clusterxl",
	"hsrp",
	"glbp",
	"other",
}

func (v *FHRPGroupProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FHRPGroupProtocol(value)
	for _, existing := range AllowedFHRPGroupProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FHRPGroupProtocol", value)
}

// NewFHRPGroupProtocolFromValue returns a pointer to a valid FHRPGroupProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFHRPGroupProtocolFromValue(v string) (*FHRPGroupProtocol, error) {
	ev := FHRPGroupProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FHRPGroupProtocol: valid values are %v", v, AllowedFHRPGroupProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FHRPGroupProtocol) IsValid() bool {
	for _, existing := range AllowedFHRPGroupProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FHRPGroup_protocol value
func (v FHRPGroupProtocol) Ptr() *FHRPGroupProtocol {
	return &v
}

type NullableFHRPGroupProtocol struct {
	value *FHRPGroupProtocol
	isSet bool
}

func (v NullableFHRPGroupProtocol) Get() *FHRPGroupProtocol {
	return v.value
}

func (v *NullableFHRPGroupProtocol) Set(val *FHRPGroupProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableFHRPGroupProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableFHRPGroupProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFHRPGroupProtocol(val *FHRPGroupProtocol) *NullableFHRPGroupProtocol {
	return &NullableFHRPGroupProtocol{value: val, isSet: true}
}

func (v NullableFHRPGroupProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFHRPGroupProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
