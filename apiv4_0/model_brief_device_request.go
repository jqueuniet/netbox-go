/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv4_0

import (
	"encoding/json"
)

// checks if the BriefDeviceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefDeviceRequest{}

// BriefDeviceRequest Adds support for custom fields and tags.
type BriefDeviceRequest struct {
	Name                 NullableString `json:"name,omitempty"`
	Description          *string        `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefDeviceRequest BriefDeviceRequest

// NewBriefDeviceRequest instantiates a new BriefDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefDeviceRequest() *BriefDeviceRequest {
	this := BriefDeviceRequest{}
	return &this
}

// NewBriefDeviceRequestWithDefaults instantiates a new BriefDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefDeviceRequestWithDefaults() *BriefDeviceRequest {
	this := BriefDeviceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BriefDeviceRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BriefDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BriefDeviceRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BriefDeviceRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *BriefDeviceRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BriefDeviceRequest) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefDeviceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefDeviceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefDeviceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefDeviceRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefDeviceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefDeviceRequest) UnmarshalJSON(data []byte) (err error) {
	varBriefDeviceRequest := _BriefDeviceRequest{}

	err = json.Unmarshal(data, &varBriefDeviceRequest)

	if err != nil {
		return err
	}

	*o = BriefDeviceRequest(varBriefDeviceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefDeviceRequest struct {
	value *BriefDeviceRequest
	isSet bool
}

func (v NullableBriefDeviceRequest) Get() *BriefDeviceRequest {
	return v.value
}

func (v *NullableBriefDeviceRequest) Set(val *BriefDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefDeviceRequest(val *BriefDeviceRequest) *NullableBriefDeviceRequest {
	return &NullableBriefDeviceRequest{value: val, isSet: true}
}

func (v NullableBriefDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}