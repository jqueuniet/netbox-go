/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the IKEPolicyMode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IKEPolicyMode{}

// IKEPolicyMode struct for IKEPolicyMode
type IKEPolicyMode struct {
	// * `aggressive` - Aggressive * `main` - Main
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewIKEPolicyMode instantiates a new IKEPolicyMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIKEPolicyMode() *IKEPolicyMode {
	this := IKEPolicyMode{}
	return &this
}

// NewIKEPolicyModeWithDefaults instantiates a new IKEPolicyMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIKEPolicyModeWithDefaults() *IKEPolicyMode {
	this := IKEPolicyMode{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IKEPolicyMode) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicyMode) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IKEPolicyMode) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IKEPolicyMode) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IKEPolicyMode) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEPolicyMode) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IKEPolicyMode) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *IKEPolicyMode) SetLabel(v string) {
	o.Label = &v
}

func (o IKEPolicyMode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IKEPolicyMode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableIKEPolicyMode struct {
	value *IKEPolicyMode
	isSet bool
}

func (v NullableIKEPolicyMode) Get() *IKEPolicyMode {
	return v.value
}

func (v *NullableIKEPolicyMode) Set(val *IKEPolicyMode) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEPolicyMode) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEPolicyMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEPolicyMode(val *IKEPolicyMode) *NullableIKEPolicyMode {
	return &NullableIKEPolicyMode{value: val, isSet: true}
}

func (v NullableIKEPolicyMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEPolicyMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

