/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv4_0

import (
	"encoding/json"
	"fmt"
)

// checks if the BriefLocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefLocationRequest{}

// BriefLocationRequest Extends PrimaryModelSerializer to include MPTT support.
type BriefLocationRequest struct {
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefLocationRequest BriefLocationRequest

// NewBriefLocationRequest instantiates a new BriefLocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefLocationRequest(name string, slug string) *BriefLocationRequest {
	this := BriefLocationRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewBriefLocationRequestWithDefaults instantiates a new BriefLocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefLocationRequestWithDefaults() *BriefLocationRequest {
	this := BriefLocationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefLocationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefLocationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefLocationRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BriefLocationRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BriefLocationRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BriefLocationRequest) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefLocationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefLocationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefLocationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefLocationRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefLocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefLocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefLocationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slug",
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
	varBriefLocationRequest := _BriefLocationRequest{}

	err = json.Unmarshal(data, &varBriefLocationRequest)

	if err != nil {
		return err
	}

	*o = BriefLocationRequest(varBriefLocationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefLocationRequest struct {
	value *BriefLocationRequest
	isSet bool
}

func (v NullableBriefLocationRequest) Get() *BriefLocationRequest {
	return v.value
}

func (v *NullableBriefLocationRequest) Set(val *BriefLocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefLocationRequest(val *BriefLocationRequest) *NullableBriefLocationRequest {
	return &NullableBriefLocationRequest{value: val, isSet: true}
}

func (v NullableBriefLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}