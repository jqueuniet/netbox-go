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

// checks if the DashboardRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardRequest{}

// DashboardRequest struct for DashboardRequest
type DashboardRequest struct {
	Layout interface{} `json:"layout,omitempty"`
	Config interface{} `json:"config,omitempty"`
}

// NewDashboardRequest instantiates a new DashboardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardRequest() *DashboardRequest {
	this := DashboardRequest{}
	return &this
}

// NewDashboardRequestWithDefaults instantiates a new DashboardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardRequestWithDefaults() *DashboardRequest {
	this := DashboardRequest{}
	return &this
}

// GetLayout returns the Layout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardRequest) GetLayout() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DashboardRequest) GetLayoutOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return &o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *DashboardRequest) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given interface{} and assigns it to the Layout field.
func (o *DashboardRequest) SetLayout(v interface{}) {
	o.Layout = v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardRequest) GetConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DashboardRequest) GetConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *DashboardRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *DashboardRequest) SetConfig(v interface{}) {
	o.Config = v
}

func (o DashboardRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableDashboardRequest struct {
	value *DashboardRequest
	isSet bool
}

func (v NullableDashboardRequest) Get() *DashboardRequest {
	return v.value
}

func (v *NullableDashboardRequest) Set(val *DashboardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardRequest(val *DashboardRequest) *NullableDashboardRequest {
	return &NullableDashboardRequest{value: val, isSet: true}
}

func (v NullableDashboardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

