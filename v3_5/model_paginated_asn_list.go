/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.9 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3_5

import (
	"encoding/json"
)

// checks if the PaginatedASNList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedASNList{}

// PaginatedASNList struct for PaginatedASNList
type PaginatedASNList struct {
	Count                *int32         `json:"count,omitempty"`
	Next                 NullableString `json:"next,omitempty"`
	Previous             NullableString `json:"previous,omitempty"`
	Results              []ASN          `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedASNList PaginatedASNList

// NewPaginatedASNList instantiates a new PaginatedASNList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedASNList() *PaginatedASNList {
	this := PaginatedASNList{}
	return &this
}

// NewPaginatedASNListWithDefaults instantiates a new PaginatedASNList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedASNListWithDefaults() *PaginatedASNList {
	this := PaginatedASNList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedASNList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedASNList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedASNList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedASNList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedASNList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedASNList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedASNList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedASNList) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedASNList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedASNList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedASNList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedASNList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedASNList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedASNList) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedASNList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedASNList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedASNList) GetResults() []ASN {
	if o == nil || IsNil(o.Results) {
		var ret []ASN
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedASNList) GetResultsOk() ([]ASN, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedASNList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ASN and assigns it to the Results field.
func (o *PaginatedASNList) SetResults(v []ASN) {
	o.Results = v
}

func (o PaginatedASNList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedASNList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedASNList) UnmarshalJSON(data []byte) (err error) {
	varPaginatedASNList := _PaginatedASNList{}

	err = json.Unmarshal(data, &varPaginatedASNList)

	if err != nil {
		return err
	}

	*o = PaginatedASNList(varPaginatedASNList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedASNList struct {
	value *PaginatedASNList
	isSet bool
}

func (v NullablePaginatedASNList) Get() *PaginatedASNList {
	return v.value
}

func (v *NullablePaginatedASNList) Set(val *PaginatedASNList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedASNList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedASNList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedASNList(val *PaginatedASNList) *NullablePaginatedASNList {
	return &NullablePaginatedASNList{value: val, isSet: true}
}

func (v NullablePaginatedASNList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedASNList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}