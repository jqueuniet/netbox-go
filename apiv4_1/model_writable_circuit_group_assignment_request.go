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

// checks if the WritableCircuitGroupAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableCircuitGroupAssignmentRequest{}

// WritableCircuitGroupAssignmentRequest Base serializer for group assignments under CircuitSerializer.
type WritableCircuitGroupAssignmentRequest struct {
	Group                BriefCircuitGroupRequest                            `json:"group"`
	Circuit              BriefCircuitRequest                                 `json:"circuit"`
	Priority             *BriefCircuitGroupAssignmentSerializerPriorityValue `json:"priority,omitempty"`
	Tags                 []NestedTagRequest                                  `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableCircuitGroupAssignmentRequest WritableCircuitGroupAssignmentRequest

// NewWritableCircuitGroupAssignmentRequest instantiates a new WritableCircuitGroupAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableCircuitGroupAssignmentRequest(group BriefCircuitGroupRequest, circuit BriefCircuitRequest) *WritableCircuitGroupAssignmentRequest {
	this := WritableCircuitGroupAssignmentRequest{}
	this.Group = group
	this.Circuit = circuit
	return &this
}

// NewWritableCircuitGroupAssignmentRequestWithDefaults instantiates a new WritableCircuitGroupAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableCircuitGroupAssignmentRequestWithDefaults() *WritableCircuitGroupAssignmentRequest {
	this := WritableCircuitGroupAssignmentRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *WritableCircuitGroupAssignmentRequest) GetGroup() BriefCircuitGroupRequest {
	if o == nil {
		var ret BriefCircuitGroupRequest
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *WritableCircuitGroupAssignmentRequest) GetGroupOk() (*BriefCircuitGroupRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *WritableCircuitGroupAssignmentRequest) SetGroup(v BriefCircuitGroupRequest) {
	o.Group = v
}

// GetCircuit returns the Circuit field value
func (o *WritableCircuitGroupAssignmentRequest) GetCircuit() BriefCircuitRequest {
	if o == nil {
		var ret BriefCircuitRequest
		return ret
	}

	return o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value
// and a boolean to check if the value has been set.
func (o *WritableCircuitGroupAssignmentRequest) GetCircuitOk() (*BriefCircuitRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Circuit, true
}

// SetCircuit sets field value
func (o *WritableCircuitGroupAssignmentRequest) SetCircuit(v BriefCircuitRequest) {
	o.Circuit = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *WritableCircuitGroupAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue {
	if o == nil || IsNil(o.Priority) {
		var ret BriefCircuitGroupAssignmentSerializerPriorityValue
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitGroupAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *WritableCircuitGroupAssignmentRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given BriefCircuitGroupAssignmentSerializerPriorityValue and assigns it to the Priority field.
func (o *WritableCircuitGroupAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue) {
	o.Priority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableCircuitGroupAssignmentRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableCircuitGroupAssignmentRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableCircuitGroupAssignmentRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableCircuitGroupAssignmentRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o WritableCircuitGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableCircuitGroupAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["circuit"] = o.Circuit
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableCircuitGroupAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group",
		"circuit",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	varWritableCircuitGroupAssignmentRequest := _WritableCircuitGroupAssignmentRequest{}

	err = json.Unmarshal(data, &varWritableCircuitGroupAssignmentRequest)

	if err != nil {
		return err
	}

	*o = WritableCircuitGroupAssignmentRequest(varWritableCircuitGroupAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "circuit")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableCircuitGroupAssignmentRequest struct {
	value *WritableCircuitGroupAssignmentRequest
	isSet bool
}

func (v NullableWritableCircuitGroupAssignmentRequest) Get() *WritableCircuitGroupAssignmentRequest {
	return v.value
}

func (v *NullableWritableCircuitGroupAssignmentRequest) Set(val *WritableCircuitGroupAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableCircuitGroupAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableCircuitGroupAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableCircuitGroupAssignmentRequest(val *WritableCircuitGroupAssignmentRequest) *NullableWritableCircuitGroupAssignmentRequest {
	return &NullableWritableCircuitGroupAssignmentRequest{value: val, isSet: true}
}

func (v NullableWritableCircuitGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableCircuitGroupAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
