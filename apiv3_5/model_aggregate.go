/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.9 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv3_5

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Aggregate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Aggregate{}

// Aggregate Adds support for custom fields and tags.
type Aggregate struct {
	Id                   int32                  `json:"id"`
	Url                  string                 `json:"url"`
	Display              string                 `json:"display"`
	Family               AggregateFamily        `json:"family"`
	Prefix               string                 `json:"prefix"`
	Rir                  NestedRIR              `json:"rir"`
	Tenant               NullableNestedTenant   `json:"tenant,omitempty"`
	DateAdded            NullableString         `json:"date_added,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _Aggregate Aggregate

// NewAggregate instantiates a new Aggregate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregate(id int32, url string, display string, family AggregateFamily, prefix string, rir NestedRIR, created NullableTime, lastUpdated NullableTime) *Aggregate {
	this := Aggregate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Family = family
	this.Prefix = prefix
	this.Rir = rir
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewAggregateWithDefaults instantiates a new Aggregate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateWithDefaults() *Aggregate {
	this := Aggregate{}
	return &this
}

// GetId returns the Id field value
func (o *Aggregate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Aggregate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Aggregate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Aggregate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Aggregate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Aggregate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *Aggregate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Aggregate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Aggregate) SetDisplay(v string) {
	o.Display = v
}

// GetFamily returns the Family field value
func (o *Aggregate) GetFamily() AggregateFamily {
	if o == nil {
		var ret AggregateFamily
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *Aggregate) GetFamilyOk() (*AggregateFamily, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *Aggregate) SetFamily(v AggregateFamily) {
	o.Family = v
}

// GetPrefix returns the Prefix field value
func (o *Aggregate) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *Aggregate) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *Aggregate) SetPrefix(v string) {
	o.Prefix = v
}

// GetRir returns the Rir field value
func (o *Aggregate) GetRir() NestedRIR {
	if o == nil {
		var ret NestedRIR
		return ret
	}

	return o.Rir
}

// GetRirOk returns a tuple with the Rir field value
// and a boolean to check if the value has been set.
func (o *Aggregate) GetRirOk() (*NestedRIR, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rir, true
}

// SetRir sets field value
func (o *Aggregate) SetRir(v NestedRIR) {
	o.Rir = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Aggregate) GetTenant() NestedTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret NestedTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Aggregate) GetTenantOk() (*NestedTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Aggregate) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableNestedTenant and assigns it to the Tenant field.
func (o *Aggregate) SetTenant(v NestedTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Aggregate) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Aggregate) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Aggregate) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded.Get()) {
		var ret string
		return ret
	}
	return *o.DateAdded.Get()
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Aggregate) GetDateAddedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateAdded.Get(), o.DateAdded.IsSet()
}

// HasDateAdded returns a boolean if a field has been set.
func (o *Aggregate) HasDateAdded() bool {
	if o != nil && o.DateAdded.IsSet() {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given NullableString and assigns it to the DateAdded field.
func (o *Aggregate) SetDateAdded(v string) {
	o.DateAdded.Set(&v)
}

// SetDateAddedNil sets the value for DateAdded to be an explicit nil
func (o *Aggregate) SetDateAddedNil() {
	o.DateAdded.Set(nil)
}

// UnsetDateAdded ensures that no value is present for DateAdded, not even an explicit nil
func (o *Aggregate) UnsetDateAdded() {
	o.DateAdded.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Aggregate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Aggregate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Aggregate) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Aggregate) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregate) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Aggregate) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Aggregate) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Aggregate) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregate) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Aggregate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Aggregate) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Aggregate) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Aggregate) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Aggregate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Aggregate) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Aggregate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Aggregate) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Aggregate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Aggregate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Aggregate) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o Aggregate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Aggregate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["family"] = o.Family
	toSerialize["prefix"] = o.Prefix
	toSerialize["rir"] = o.Rir
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.DateAdded.IsSet() {
		toSerialize["date_added"] = o.DateAdded.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
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

func (o *Aggregate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"family",
		"prefix",
		"rir",
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
	varAggregate := _Aggregate{}

	err = json.Unmarshal(data, &varAggregate)

	if err != nil {
		return err
	}

	*o = Aggregate(varAggregate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "family")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "rir")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "date_added")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggregate struct {
	value *Aggregate
	isSet bool
}

func (v NullableAggregate) Get() *Aggregate {
	return v.value
}

func (v *NullableAggregate) Set(val *Aggregate) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregate) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregate(val *Aggregate) *NullableAggregate {
	return &NullableAggregate{value: val, isSet: true}
}

func (v NullableAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}