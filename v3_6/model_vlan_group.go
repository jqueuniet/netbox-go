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
	"time"
)

// checks if the VLANGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VLANGroup{}

// VLANGroup Adds support for custom fields and tags.
type VLANGroup struct {
	Id        int32          `json:"id"`
	Url       string         `json:"url"`
	Display   string         `json:"display"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	ScopeType NullableString `json:"scope_type,omitempty"`
	ScopeId   NullableInt32  `json:"scope_id,omitempty"`
	Scope     interface{}    `json:"scope"`
	// Lowest permissible ID of a child VLAN
	MinVid *int32 `json:"min_vid,omitempty"`
	// Highest permissible ID of a child VLAN
	MaxVid               *int32                 `json:"max_vid,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	VlanCount            int32                  `json:"vlan_count"`
	Utilization          string                 `json:"utilization"`
	AdditionalProperties map[string]interface{}
}

type _VLANGroup VLANGroup

// NewVLANGroup instantiates a new VLANGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVLANGroup(id int32, url string, display string, name string, slug string, scope interface{}, created NullableTime, lastUpdated NullableTime, vlanCount int32, utilization string) *VLANGroup {
	this := VLANGroup{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Scope = scope
	this.Created = created
	this.LastUpdated = lastUpdated
	this.VlanCount = vlanCount
	this.Utilization = utilization
	return &this
}

// NewVLANGroupWithDefaults instantiates a new VLANGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVLANGroupWithDefaults() *VLANGroup {
	this := VLANGroup{}
	return &this
}

// GetId returns the Id field value
func (o *VLANGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VLANGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VLANGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VLANGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *VLANGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VLANGroup) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *VLANGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VLANGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *VLANGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *VLANGroup) SetSlug(v string) {
	o.Slug = v
}

// GetScopeType returns the ScopeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLANGroup) GetScopeType() string {
	if o == nil || IsNil(o.ScopeType.Get()) {
		var ret string
		return ret
	}
	return *o.ScopeType.Get()
}

// GetScopeTypeOk returns a tuple with the ScopeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLANGroup) GetScopeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeType.Get(), o.ScopeType.IsSet()
}

// HasScopeType returns a boolean if a field has been set.
func (o *VLANGroup) HasScopeType() bool {
	if o != nil && o.ScopeType.IsSet() {
		return true
	}

	return false
}

// SetScopeType gets a reference to the given NullableString and assigns it to the ScopeType field.
func (o *VLANGroup) SetScopeType(v string) {
	o.ScopeType.Set(&v)
}

// SetScopeTypeNil sets the value for ScopeType to be an explicit nil
func (o *VLANGroup) SetScopeTypeNil() {
	o.ScopeType.Set(nil)
}

// UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
func (o *VLANGroup) UnsetScopeType() {
	o.ScopeType.Unset()
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLANGroup) GetScopeId() int32 {
	if o == nil || IsNil(o.ScopeId.Get()) {
		var ret int32
		return ret
	}
	return *o.ScopeId.Get()
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLANGroup) GetScopeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScopeId.Get(), o.ScopeId.IsSet()
}

// HasScopeId returns a boolean if a field has been set.
func (o *VLANGroup) HasScopeId() bool {
	if o != nil && o.ScopeId.IsSet() {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given NullableInt32 and assigns it to the ScopeId field.
func (o *VLANGroup) SetScopeId(v int32) {
	o.ScopeId.Set(&v)
}

// SetScopeIdNil sets the value for ScopeId to be an explicit nil
func (o *VLANGroup) SetScopeIdNil() {
	o.ScopeId.Set(nil)
}

// UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
func (o *VLANGroup) UnsetScopeId() {
	o.ScopeId.Unset()
}

// GetScope returns the Scope field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *VLANGroup) GetScope() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLANGroup) GetScopeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *VLANGroup) SetScope(v interface{}) {
	o.Scope = v
}

// GetMinVid returns the MinVid field value if set, zero value otherwise.
func (o *VLANGroup) GetMinVid() int32 {
	if o == nil || IsNil(o.MinVid) {
		var ret int32
		return ret
	}
	return *o.MinVid
}

// GetMinVidOk returns a tuple with the MinVid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetMinVidOk() (*int32, bool) {
	if o == nil || IsNil(o.MinVid) {
		return nil, false
	}
	return o.MinVid, true
}

// HasMinVid returns a boolean if a field has been set.
func (o *VLANGroup) HasMinVid() bool {
	if o != nil && !IsNil(o.MinVid) {
		return true
	}

	return false
}

// SetMinVid gets a reference to the given int32 and assigns it to the MinVid field.
func (o *VLANGroup) SetMinVid(v int32) {
	o.MinVid = &v
}

// GetMaxVid returns the MaxVid field value if set, zero value otherwise.
func (o *VLANGroup) GetMaxVid() int32 {
	if o == nil || IsNil(o.MaxVid) {
		var ret int32
		return ret
	}
	return *o.MaxVid
}

// GetMaxVidOk returns a tuple with the MaxVid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetMaxVidOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxVid) {
		return nil, false
	}
	return o.MaxVid, true
}

// HasMaxVid returns a boolean if a field has been set.
func (o *VLANGroup) HasMaxVid() bool {
	if o != nil && !IsNil(o.MaxVid) {
		return true
	}

	return false
}

// SetMaxVid gets a reference to the given int32 and assigns it to the MaxVid field.
func (o *VLANGroup) SetMaxVid(v int32) {
	o.MaxVid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VLANGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VLANGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VLANGroup) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VLANGroup) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VLANGroup) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *VLANGroup) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VLANGroup) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VLANGroup) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VLANGroup) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VLANGroup) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLANGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *VLANGroup) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VLANGroup) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLANGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *VLANGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetVlanCount returns the VlanCount field value
func (o *VLANGroup) GetVlanCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetVlanCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanCount, true
}

// SetVlanCount sets field value
func (o *VLANGroup) SetVlanCount(v int32) {
	o.VlanCount = v
}

// GetUtilization returns the Utilization field value
func (o *VLANGroup) GetUtilization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value
// and a boolean to check if the value has been set.
func (o *VLANGroup) GetUtilizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Utilization, true
}

// SetUtilization sets field value
func (o *VLANGroup) SetUtilization(v string) {
	o.Utilization = v
}

func (o VLANGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VLANGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if o.ScopeType.IsSet() {
		toSerialize["scope_type"] = o.ScopeType.Get()
	}
	if o.ScopeId.IsSet() {
		toSerialize["scope_id"] = o.ScopeId.Get()
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.MinVid) {
		toSerialize["min_vid"] = o.MinVid
	}
	if !IsNil(o.MaxVid) {
		toSerialize["max_vid"] = o.MaxVid
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["vlan_count"] = o.VlanCount
	toSerialize["utilization"] = o.Utilization

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VLANGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"slug",
		"scope",
		"created",
		"last_updated",
		"vlan_count",
		"utilization",
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
	varVLANGroup := _VLANGroup{}

	err = json.Unmarshal(data, &varVLANGroup)

	if err != nil {
		return err
	}

	*o = VLANGroup(varVLANGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "scope_type")
		delete(additionalProperties, "scope_id")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "min_vid")
		delete(additionalProperties, "max_vid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "vlan_count")
		delete(additionalProperties, "utilization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVLANGroup struct {
	value *VLANGroup
	isSet bool
}

func (v NullableVLANGroup) Get() *VLANGroup {
	return v.value
}

func (v *NullableVLANGroup) Set(val *VLANGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableVLANGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableVLANGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVLANGroup(val *VLANGroup) *NullableVLANGroup {
	return &NullableVLANGroup{value: val, isSet: true}
}

func (v NullableVLANGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVLANGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}