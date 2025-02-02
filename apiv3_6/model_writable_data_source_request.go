/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv3_6

import (
	"encoding/json"
	"fmt"
)

// checks if the WritableDataSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableDataSourceRequest{}

// WritableDataSourceRequest Adds support for custom fields and tags.
type WritableDataSourceRequest struct {
	Name        string               `json:"name"`
	Type        *DataSourceTypeValue `json:"type,omitempty"`
	SourceUrl   string               `json:"source_url"`
	Enabled     *bool                `json:"enabled,omitempty"`
	Description *string              `json:"description,omitempty"`
	Comments    *string              `json:"comments,omitempty"`
	Parameters  interface{}          `json:"parameters,omitempty"`
	// Patterns (one per line) matching files to ignore when syncing
	IgnoreRules          *string `json:"ignore_rules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableDataSourceRequest WritableDataSourceRequest

// NewWritableDataSourceRequest instantiates a new WritableDataSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableDataSourceRequest(name string, sourceUrl string) *WritableDataSourceRequest {
	this := WritableDataSourceRequest{}
	this.Name = name
	this.SourceUrl = sourceUrl
	return &this
}

// NewWritableDataSourceRequestWithDefaults instantiates a new WritableDataSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableDataSourceRequestWithDefaults() *WritableDataSourceRequest {
	this := WritableDataSourceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableDataSourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableDataSourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableDataSourceRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WritableDataSourceRequest) GetType() DataSourceTypeValue {
	if o == nil || IsNil(o.Type) {
		var ret DataSourceTypeValue
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDataSourceRequest) GetTypeOk() (*DataSourceTypeValue, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WritableDataSourceRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DataSourceTypeValue and assigns it to the Type field.
func (o *WritableDataSourceRequest) SetType(v DataSourceTypeValue) {
	o.Type = &v
}

// GetSourceUrl returns the SourceUrl field value
func (o *WritableDataSourceRequest) GetSourceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value
// and a boolean to check if the value has been set.
func (o *WritableDataSourceRequest) GetSourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceUrl, true
}

// SetSourceUrl sets field value
func (o *WritableDataSourceRequest) SetSourceUrl(v string) {
	o.SourceUrl = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WritableDataSourceRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDataSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WritableDataSourceRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WritableDataSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableDataSourceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDataSourceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableDataSourceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableDataSourceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableDataSourceRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDataSourceRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableDataSourceRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableDataSourceRequest) SetComments(v string) {
	o.Comments = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableDataSourceRequest) GetParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableDataSourceRequest) GetParametersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WritableDataSourceRequest) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given interface{} and assigns it to the Parameters field.
func (o *WritableDataSourceRequest) SetParameters(v interface{}) {
	o.Parameters = v
}

// GetIgnoreRules returns the IgnoreRules field value if set, zero value otherwise.
func (o *WritableDataSourceRequest) GetIgnoreRules() string {
	if o == nil || IsNil(o.IgnoreRules) {
		var ret string
		return ret
	}
	return *o.IgnoreRules
}

// GetIgnoreRulesOk returns a tuple with the IgnoreRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDataSourceRequest) GetIgnoreRulesOk() (*string, bool) {
	if o == nil || IsNil(o.IgnoreRules) {
		return nil, false
	}
	return o.IgnoreRules, true
}

// HasIgnoreRules returns a boolean if a field has been set.
func (o *WritableDataSourceRequest) HasIgnoreRules() bool {
	if o != nil && !IsNil(o.IgnoreRules) {
		return true
	}

	return false
}

// SetIgnoreRules gets a reference to the given string and assigns it to the IgnoreRules field.
func (o *WritableDataSourceRequest) SetIgnoreRules(v string) {
	o.IgnoreRules = &v
}

func (o WritableDataSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableDataSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["source_url"] = o.SourceUrl
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.IgnoreRules) {
		toSerialize["ignore_rules"] = o.IgnoreRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableDataSourceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"source_url",
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
	varWritableDataSourceRequest := _WritableDataSourceRequest{}

	err = json.Unmarshal(data, &varWritableDataSourceRequest)

	if err != nil {
		return err
	}

	*o = WritableDataSourceRequest(varWritableDataSourceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "source_url")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "parameters")
		delete(additionalProperties, "ignore_rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableDataSourceRequest struct {
	value *WritableDataSourceRequest
	isSet bool
}

func (v NullableWritableDataSourceRequest) Get() *WritableDataSourceRequest {
	return v.value
}

func (v *NullableWritableDataSourceRequest) Set(val *WritableDataSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableDataSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableDataSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableDataSourceRequest(val *WritableDataSourceRequest) *NullableWritableDataSourceRequest {
	return &NullableWritableDataSourceRequest{value: val, isSet: true}
}

func (v NullableWritableDataSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableDataSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
