/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.6 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv4_1

import (
	"encoding/json"
)

// checks if the WirelessLinkDistanceUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelessLinkDistanceUnit{}

// WirelessLinkDistanceUnit struct for WirelessLinkDistanceUnit
type WirelessLinkDistanceUnit struct {
	// * `km` - Kilometers * `m` - Meters * `mi` - Miles * `ft` - Feet
	Value                NullableString `json:"value,omitempty"`
	Label                *string        `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WirelessLinkDistanceUnit WirelessLinkDistanceUnit

// NewWirelessLinkDistanceUnit instantiates a new WirelessLinkDistanceUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelessLinkDistanceUnit() *WirelessLinkDistanceUnit {
	this := WirelessLinkDistanceUnit{}
	return &this
}

// NewWirelessLinkDistanceUnitWithDefaults instantiates a new WirelessLinkDistanceUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelessLinkDistanceUnitWithDefaults() *WirelessLinkDistanceUnit {
	this := WirelessLinkDistanceUnit{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WirelessLinkDistanceUnit) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessLinkDistanceUnit) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *WirelessLinkDistanceUnit) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *WirelessLinkDistanceUnit) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *WirelessLinkDistanceUnit) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *WirelessLinkDistanceUnit) UnsetValue() {
	o.Value.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WirelessLinkDistanceUnit) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLinkDistanceUnit) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WirelessLinkDistanceUnit) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WirelessLinkDistanceUnit) SetLabel(v string) {
	o.Label = &v
}

func (o WirelessLinkDistanceUnit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelessLinkDistanceUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WirelessLinkDistanceUnit) UnmarshalJSON(data []byte) (err error) {
	varWirelessLinkDistanceUnit := _WirelessLinkDistanceUnit{}

	err = json.Unmarshal(data, &varWirelessLinkDistanceUnit)

	if err != nil {
		return err
	}

	*o = WirelessLinkDistanceUnit(varWirelessLinkDistanceUnit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWirelessLinkDistanceUnit struct {
	value *WirelessLinkDistanceUnit
	isSet bool
}

func (v NullableWirelessLinkDistanceUnit) Get() *WirelessLinkDistanceUnit {
	return v.value
}

func (v *NullableWirelessLinkDistanceUnit) Set(val *WirelessLinkDistanceUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLinkDistanceUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLinkDistanceUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLinkDistanceUnit(val *WirelessLinkDistanceUnit) *NullableWirelessLinkDistanceUnit {
	return &NullableWirelessLinkDistanceUnit{value: val, isSet: true}
}

func (v NullableWirelessLinkDistanceUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLinkDistanceUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}