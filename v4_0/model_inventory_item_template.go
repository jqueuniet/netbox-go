/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the InventoryItemTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryItemTemplate{}

// InventoryItemTemplate Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type InventoryItemTemplate struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	DeviceType BriefDeviceType `json:"device_type"`
	Parent NullableInt32 `json:"parent,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Role NullableBriefInventoryItemRole `json:"role,omitempty"`
	Manufacturer NullableBriefManufacturer `json:"manufacturer,omitempty"`
	// Manufacturer-assigned part identifier
	PartId *string `json:"part_id,omitempty"`
	Description *string `json:"description,omitempty"`
	ComponentType NullableString `json:"component_type,omitempty"`
	ComponentId NullableInt64 `json:"component_id,omitempty"`
	Component interface{} `json:"component"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	Depth int32 `json:"_depth"`
}

type _InventoryItemTemplate InventoryItemTemplate

// NewInventoryItemTemplate instantiates a new InventoryItemTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryItemTemplate(id int32, url string, display string, deviceType BriefDeviceType, name string, component interface{}, created NullableTime, lastUpdated NullableTime, depth int32) *InventoryItemTemplate {
	this := InventoryItemTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.DeviceType = deviceType
	this.Name = name
	this.Component = component
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Depth = depth
	return &this
}

// NewInventoryItemTemplateWithDefaults instantiates a new InventoryItemTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryItemTemplateWithDefaults() *InventoryItemTemplate {
	this := InventoryItemTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *InventoryItemTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InventoryItemTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InventoryItemTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *InventoryItemTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InventoryItemTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InventoryItemTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *InventoryItemTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *InventoryItemTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *InventoryItemTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetDeviceType returns the DeviceType field value
func (o *InventoryItemTemplate) GetDeviceType() BriefDeviceType {
	if o == nil {
		var ret BriefDeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *InventoryItemTemplate) GetDeviceTypeOk() (*BriefDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *InventoryItemTemplate) SetDeviceType(v BriefDeviceType) {
	o.DeviceType = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemTemplate) GetParent() int32 {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret int32
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemTemplate) GetParentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *InventoryItemTemplate) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableInt32 and assigns it to the Parent field.
func (o *InventoryItemTemplate) SetParent(v int32) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *InventoryItemTemplate) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *InventoryItemTemplate) UnsetParent() {
	o.Parent.Unset()
}

// GetName returns the Name field value
func (o *InventoryItemTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InventoryItemTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InventoryItemTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InventoryItemTemplate) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemTemplate) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InventoryItemTemplate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InventoryItemTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemTemplate) GetRole() BriefInventoryItemRole {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefInventoryItemRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemTemplate) GetRoleOk() (*BriefInventoryItemRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *InventoryItemTemplate) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefInventoryItemRole and assigns it to the Role field.
func (o *InventoryItemTemplate) SetRole(v BriefInventoryItemRole) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *InventoryItemTemplate) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *InventoryItemTemplate) UnsetRole() {
	o.Role.Unset()
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemTemplate) GetManufacturer() BriefManufacturer {
	if o == nil || IsNil(o.Manufacturer.Get()) {
		var ret BriefManufacturer
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemTemplate) GetManufacturerOk() (*BriefManufacturer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *InventoryItemTemplate) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableBriefManufacturer and assigns it to the Manufacturer field.
func (o *InventoryItemTemplate) SetManufacturer(v BriefManufacturer) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *InventoryItemTemplate) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *InventoryItemTemplate) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *InventoryItemTemplate) GetPartId() string {
	if o == nil || IsNil(o.PartId) {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemTemplate) GetPartIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartId) {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *InventoryItemTemplate) HasPartId() bool {
	if o != nil && !IsNil(o.PartId) {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *InventoryItemTemplate) SetPartId(v string) {
	o.PartId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InventoryItemTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InventoryItemTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InventoryItemTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemTemplate) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType.Get()) {
		var ret string
		return ret
	}
	return *o.ComponentType.Get()
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemTemplate) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentType.Get(), o.ComponentType.IsSet()
}

// HasComponentType returns a boolean if a field has been set.
func (o *InventoryItemTemplate) HasComponentType() bool {
	if o != nil && o.ComponentType.IsSet() {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given NullableString and assigns it to the ComponentType field.
func (o *InventoryItemTemplate) SetComponentType(v string) {
	o.ComponentType.Set(&v)
}
// SetComponentTypeNil sets the value for ComponentType to be an explicit nil
func (o *InventoryItemTemplate) SetComponentTypeNil() {
	o.ComponentType.Set(nil)
}

// UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
func (o *InventoryItemTemplate) UnsetComponentType() {
	o.ComponentType.Unset()
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItemTemplate) GetComponentId() int64 {
	if o == nil || IsNil(o.ComponentId.Get()) {
		var ret int64
		return ret
	}
	return *o.ComponentId.Get()
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemTemplate) GetComponentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentId.Get(), o.ComponentId.IsSet()
}

// HasComponentId returns a boolean if a field has been set.
func (o *InventoryItemTemplate) HasComponentId() bool {
	if o != nil && o.ComponentId.IsSet() {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given NullableInt64 and assigns it to the ComponentId field.
func (o *InventoryItemTemplate) SetComponentId(v int64) {
	o.ComponentId.Set(&v)
}
// SetComponentIdNil sets the value for ComponentId to be an explicit nil
func (o *InventoryItemTemplate) SetComponentIdNil() {
	o.ComponentId.Set(nil)
}

// UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil
func (o *InventoryItemTemplate) UnsetComponentId() {
	o.ComponentId.Unset()
}

// GetComponent returns the Component field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InventoryItemTemplate) GetComponent() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemTemplate) GetComponentOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *InventoryItemTemplate) SetComponent(v interface{}) {
	o.Component = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *InventoryItemTemplate) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemTemplate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *InventoryItemTemplate) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *InventoryItemTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItemTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *InventoryItemTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetDepth returns the Depth field value
func (o *InventoryItemTemplate) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *InventoryItemTemplate) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *InventoryItemTemplate) SetDepth(v int32) {
	o.Depth = v
}

func (o InventoryItemTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryItemTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["device_type"] = o.DeviceType
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if !IsNil(o.PartId) {
		toSerialize["part_id"] = o.PartId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.ComponentType.IsSet() {
		toSerialize["component_type"] = o.ComponentType.Get()
	}
	if o.ComponentId.IsSet() {
		toSerialize["component_id"] = o.ComponentId.Get()
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["_depth"] = o.Depth
	return toSerialize, nil
}

func (o *InventoryItemTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"device_type",
		"name",
		"component",
		"created",
		"last_updated",
		"_depth",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInventoryItemTemplate := _InventoryItemTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInventoryItemTemplate)

	if err != nil {
		return err
	}

	*o = InventoryItemTemplate(varInventoryItemTemplate)

	return err
}

type NullableInventoryItemTemplate struct {
	value *InventoryItemTemplate
	isSet bool
}

func (v NullableInventoryItemTemplate) Get() *InventoryItemTemplate {
	return v.value
}

func (v *NullableInventoryItemTemplate) Set(val *InventoryItemTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryItemTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryItemTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryItemTemplate(val *InventoryItemTemplate) *NullableInventoryItemTemplate {
	return &NullableInventoryItemTemplate{value: val, isSet: true}
}

func (v NullableInventoryItemTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryItemTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

