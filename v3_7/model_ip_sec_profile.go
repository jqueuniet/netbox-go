/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.8 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3_7

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the IPSecProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPSecProfile{}

// IPSecProfile Adds support for custom fields and tags.
type IPSecProfile struct {
	Id                   int32                  `json:"id"`
	Url                  string                 `json:"url"`
	Display              string                 `json:"display"`
	Name                 string                 `json:"name"`
	Description          *string                `json:"description,omitempty"`
	Mode                 IPSecProfileMode       `json:"mode"`
	IkePolicy            NestedIKEPolicy        `json:"ike_policy"`
	IpsecPolicy          NestedIPSecPolicy      `json:"ipsec_policy"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _IPSecProfile IPSecProfile

// NewIPSecProfile instantiates a new IPSecProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecProfile(id int32, url string, display string, name string, mode IPSecProfileMode, ikePolicy NestedIKEPolicy, ipsecPolicy NestedIPSecPolicy, created NullableTime, lastUpdated NullableTime) *IPSecProfile {
	this := IPSecProfile{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Mode = mode
	this.IkePolicy = ikePolicy
	this.IpsecPolicy = ipsecPolicy
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewIPSecProfileWithDefaults instantiates a new IPSecProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecProfileWithDefaults() *IPSecProfile {
	this := IPSecProfile{}
	return &this
}

// GetId returns the Id field value
func (o *IPSecProfile) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IPSecProfile) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *IPSecProfile) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IPSecProfile) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *IPSecProfile) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *IPSecProfile) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *IPSecProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IPSecProfile) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IPSecProfile) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IPSecProfile) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IPSecProfile) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value
func (o *IPSecProfile) GetMode() IPSecProfileMode {
	if o == nil {
		var ret IPSecProfileMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetModeOk() (*IPSecProfileMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *IPSecProfile) SetMode(v IPSecProfileMode) {
	o.Mode = v
}

// GetIkePolicy returns the IkePolicy field value
func (o *IPSecProfile) GetIkePolicy() NestedIKEPolicy {
	if o == nil {
		var ret NestedIKEPolicy
		return ret
	}

	return o.IkePolicy
}

// GetIkePolicyOk returns a tuple with the IkePolicy field value
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetIkePolicyOk() (*NestedIKEPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IkePolicy, true
}

// SetIkePolicy sets field value
func (o *IPSecProfile) SetIkePolicy(v NestedIKEPolicy) {
	o.IkePolicy = v
}

// GetIpsecPolicy returns the IpsecPolicy field value
func (o *IPSecProfile) GetIpsecPolicy() NestedIPSecPolicy {
	if o == nil {
		var ret NestedIPSecPolicy
		return ret
	}

	return o.IpsecPolicy
}

// GetIpsecPolicyOk returns a tuple with the IpsecPolicy field value
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetIpsecPolicyOk() (*NestedIPSecPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpsecPolicy, true
}

// SetIpsecPolicy sets field value
func (o *IPSecProfile) SetIpsecPolicy(v NestedIPSecPolicy) {
	o.IpsecPolicy = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IPSecProfile) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IPSecProfile) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IPSecProfile) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPSecProfile) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPSecProfile) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *IPSecProfile) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IPSecProfile) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfile) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IPSecProfile) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *IPSecProfile) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IPSecProfile) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecProfile) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *IPSecProfile) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IPSecProfile) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPSecProfile) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *IPSecProfile) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o IPSecProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPSecProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["mode"] = o.Mode
	toSerialize["ike_policy"] = o.IkePolicy
	toSerialize["ipsec_policy"] = o.IpsecPolicy
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IPSecProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"mode",
		"ike_policy",
		"ipsec_policy",
		"created",
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
	varIPSecProfile := _IPSecProfile{}

	err = json.Unmarshal(data, &varIPSecProfile)

	if err != nil {
		return err
	}

	*o = IPSecProfile(varIPSecProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "ike_policy")
		delete(additionalProperties, "ipsec_policy")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPSecProfile struct {
	value *IPSecProfile
	isSet bool
}

func (v NullableIPSecProfile) Get() *IPSecProfile {
	return v.value
}

func (v *NullableIPSecProfile) Set(val *IPSecProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecProfile(val *IPSecProfile) *NullableIPSecProfile {
	return &NullableIPSecProfile{value: val, isSet: true}
}

func (v NullableIPSecProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}