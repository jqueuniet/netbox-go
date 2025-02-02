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
	"time"
)

// checks if the CircuitGroupAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitGroupAssignment{}

// CircuitGroupAssignment Base serializer for group assignments under CircuitSerializer.
type CircuitGroupAssignment struct {
	Id                   int32                                          `json:"id"`
	Url                  string                                         `json:"url"`
	DisplayUrl           *string                                        `json:"display_url,omitempty"`
	Display              *string                                        `json:"display,omitempty"`
	Group                BriefCircuitGroup                              `json:"group"`
	Circuit              BriefCircuit                                   `json:"circuit"`
	Priority             *BriefCircuitGroupAssignmentSerializerPriority `json:"priority,omitempty"`
	Tags                 []NestedTag                                    `json:"tags,omitempty"`
	Created              NullableTime                                   `json:"created,omitempty"`
	LastUpdated          NullableTime                                   `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _CircuitGroupAssignment CircuitGroupAssignment

// NewCircuitGroupAssignment instantiates a new CircuitGroupAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitGroupAssignment(id int32, url string, group BriefCircuitGroup, circuit BriefCircuit, lastUpdated NullableTime) *CircuitGroupAssignment {
	this := CircuitGroupAssignment{}
	this.Id = id
	this.Url = url
	this.Group = group
	this.Circuit = circuit
	this.LastUpdated = lastUpdated
	return &this
}

// NewCircuitGroupAssignmentWithDefaults instantiates a new CircuitGroupAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitGroupAssignmentWithDefaults() *CircuitGroupAssignment {
	this := CircuitGroupAssignment{}
	return &this
}

// GetId returns the Id field value
func (o *CircuitGroupAssignment) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CircuitGroupAssignment) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CircuitGroupAssignment) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CircuitGroupAssignment) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CircuitGroupAssignment) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CircuitGroupAssignment) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value if set, zero value otherwise.
func (o *CircuitGroupAssignment) GetDisplayUrl() string {
	if o == nil || IsNil(o.DisplayUrl) {
		var ret string
		return ret
	}
	return *o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitGroupAssignment) GetDisplayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUrl) {
		return nil, false
	}
	return o.DisplayUrl, true
}

// HasDisplayUrl returns a boolean if a field has been set.
func (o *CircuitGroupAssignment) HasDisplayUrl() bool {
	if o != nil && !IsNil(o.DisplayUrl) {
		return true
	}

	return false
}

// SetDisplayUrl gets a reference to the given string and assigns it to the DisplayUrl field.
func (o *CircuitGroupAssignment) SetDisplayUrl(v string) {
	o.DisplayUrl = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *CircuitGroupAssignment) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitGroupAssignment) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *CircuitGroupAssignment) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *CircuitGroupAssignment) SetDisplay(v string) {
	o.Display = &v
}

// GetGroup returns the Group field value
func (o *CircuitGroupAssignment) GetGroup() BriefCircuitGroup {
	if o == nil {
		var ret BriefCircuitGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *CircuitGroupAssignment) GetGroupOk() (*BriefCircuitGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *CircuitGroupAssignment) SetGroup(v BriefCircuitGroup) {
	o.Group = v
}

// GetCircuit returns the Circuit field value
func (o *CircuitGroupAssignment) GetCircuit() BriefCircuit {
	if o == nil {
		var ret BriefCircuit
		return ret
	}

	return o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value
// and a boolean to check if the value has been set.
func (o *CircuitGroupAssignment) GetCircuitOk() (*BriefCircuit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Circuit, true
}

// SetCircuit sets field value
func (o *CircuitGroupAssignment) SetCircuit(v BriefCircuit) {
	o.Circuit = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CircuitGroupAssignment) GetPriority() BriefCircuitGroupAssignmentSerializerPriority {
	if o == nil || IsNil(o.Priority) {
		var ret BriefCircuitGroupAssignmentSerializerPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitGroupAssignment) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriority, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CircuitGroupAssignment) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given BriefCircuitGroupAssignmentSerializerPriority and assigns it to the Priority field.
func (o *CircuitGroupAssignment) SetPriority(v BriefCircuitGroupAssignmentSerializerPriority) {
	o.Priority = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CircuitGroupAssignment) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitGroupAssignment) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CircuitGroupAssignment) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *CircuitGroupAssignment) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitGroupAssignment) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitGroupAssignment) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *CircuitGroupAssignment) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *CircuitGroupAssignment) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil
func (o *CircuitGroupAssignment) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *CircuitGroupAssignment) UnsetCreated() {
	o.Created.Unset()
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CircuitGroupAssignment) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitGroupAssignment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *CircuitGroupAssignment) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o CircuitGroupAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitGroupAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	if !IsNil(o.DisplayUrl) {
		toSerialize["display_url"] = o.DisplayUrl
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	toSerialize["group"] = o.Group
	toSerialize["circuit"] = o.Circuit
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CircuitGroupAssignment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"group",
		"circuit",
		"last_updated",
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
	varCircuitGroupAssignment := _CircuitGroupAssignment{}

	err = json.Unmarshal(data, &varCircuitGroupAssignment)

	if err != nil {
		return err
	}

	*o = CircuitGroupAssignment(varCircuitGroupAssignment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "group")
		delete(additionalProperties, "circuit")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCircuitGroupAssignment struct {
	value *CircuitGroupAssignment
	isSet bool
}

func (v NullableCircuitGroupAssignment) Get() *CircuitGroupAssignment {
	return v.value
}

func (v *NullableCircuitGroupAssignment) Set(val *CircuitGroupAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitGroupAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitGroupAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitGroupAssignment(val *CircuitGroupAssignment) *NullableCircuitGroupAssignment {
	return &NullableCircuitGroupAssignment{value: val, isSet: true}
}

func (v NullableCircuitGroupAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitGroupAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
