/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3_6

import (
	"encoding/json"
)

// checks if the InterfaceTemplatePoeMode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceTemplatePoeMode{}

// InterfaceTemplatePoeMode struct for InterfaceTemplatePoeMode
type InterfaceTemplatePoeMode struct {
	// * `pd` - PD * `pse` - PSE
	Value                NullableString `json:"value,omitempty"`
	Label                *string        `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceTemplatePoeMode InterfaceTemplatePoeMode

// NewInterfaceTemplatePoeMode instantiates a new InterfaceTemplatePoeMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceTemplatePoeMode() *InterfaceTemplatePoeMode {
	this := InterfaceTemplatePoeMode{}
	return &this
}

// NewInterfaceTemplatePoeModeWithDefaults instantiates a new InterfaceTemplatePoeMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTemplatePoeModeWithDefaults() *InterfaceTemplatePoeMode {
	this := InterfaceTemplatePoeMode{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplatePoeMode) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplatePoeMode) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceTemplatePoeMode) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *InterfaceTemplatePoeMode) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *InterfaceTemplatePoeMode) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *InterfaceTemplatePoeMode) UnsetValue() {
	o.Value.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceTemplatePoeMode) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplatePoeMode) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceTemplatePoeMode) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InterfaceTemplatePoeMode) SetLabel(v string) {
	o.Label = &v
}

func (o InterfaceTemplatePoeMode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceTemplatePoeMode) ToMap() (map[string]interface{}, error) {
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

func (o *InterfaceTemplatePoeMode) UnmarshalJSON(data []byte) (err error) {
	varInterfaceTemplatePoeMode := _InterfaceTemplatePoeMode{}

	err = json.Unmarshal(data, &varInterfaceTemplatePoeMode)

	if err != nil {
		return err
	}

	*o = InterfaceTemplatePoeMode(varInterfaceTemplatePoeMode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceTemplatePoeMode struct {
	value *InterfaceTemplatePoeMode
	isSet bool
}

func (v NullableInterfaceTemplatePoeMode) Get() *InterfaceTemplatePoeMode {
	return v.value
}

func (v *NullableInterfaceTemplatePoeMode) Set(val *InterfaceTemplatePoeMode) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTemplatePoeMode) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTemplatePoeMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTemplatePoeMode(val *InterfaceTemplatePoeMode) *NullableInterfaceTemplatePoeMode {
	return &NullableInterfaceTemplatePoeMode{value: val, isSet: true}
}

func (v NullableInterfaceTemplatePoeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTemplatePoeMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}