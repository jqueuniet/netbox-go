/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3_6

import (
	"encoding/json"
	"fmt"
)

// checks if the NestedRearPortTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedRearPortTemplate{}

// NestedRearPortTemplate Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedRearPortTemplate struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedRearPortTemplate NestedRearPortTemplate

// NewNestedRearPortTemplate instantiates a new NestedRearPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRearPortTemplate(id int32, url string, display string, name string) *NestedRearPortTemplate {
	this := NestedRearPortTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewNestedRearPortTemplateWithDefaults instantiates a new NestedRearPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRearPortTemplateWithDefaults() *NestedRearPortTemplate {
	this := NestedRearPortTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *NestedRearPortTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedRearPortTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedRearPortTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedRearPortTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedRearPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedRearPortTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedRearPortTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedRearPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedRearPortTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedRearPortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedRearPortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedRearPortTemplate) SetName(v string) {
	o.Name = v
}

func (o NestedRearPortTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedRearPortTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedRearPortTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
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
	varNestedRearPortTemplate := _NestedRearPortTemplate{}

	err = json.Unmarshal(data, &varNestedRearPortTemplate)

	if err != nil {
		return err
	}

	*o = NestedRearPortTemplate(varNestedRearPortTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedRearPortTemplate struct {
	value *NestedRearPortTemplate
	isSet bool
}

func (v NullableNestedRearPortTemplate) Get() *NestedRearPortTemplate {
	return v.value
}

func (v *NullableNestedRearPortTemplate) Set(val *NestedRearPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRearPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRearPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRearPortTemplate(val *NestedRearPortTemplate) *NullableNestedRearPortTemplate {
	return &NullableNestedRearPortTemplate{value: val, isSet: true}
}

func (v NullableNestedRearPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRearPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}