/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv3_6

import (
	"encoding/json"
)

// checks if the PatchedWritableInventoryItemTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableInventoryItemTemplateRequest{}

// PatchedWritableInventoryItemTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableInventoryItemTemplateRequest struct {
	DeviceType *int32        `json:"device_type,omitempty"`
	Parent     NullableInt32 `json:"parent,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name *string `json:"name,omitempty"`
	// Physical label
	Label        *string       `json:"label,omitempty"`
	Role         NullableInt32 `json:"role,omitempty"`
	Manufacturer NullableInt32 `json:"manufacturer,omitempty"`
	// Manufacturer-assigned part identifier
	PartId               *string        `json:"part_id,omitempty"`
	Description          *string        `json:"description,omitempty"`
	ComponentType        NullableString `json:"component_type,omitempty"`
	ComponentId          NullableInt64  `json:"component_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableInventoryItemTemplateRequest PatchedWritableInventoryItemTemplateRequest

// NewPatchedWritableInventoryItemTemplateRequest instantiates a new PatchedWritableInventoryItemTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableInventoryItemTemplateRequest() *PatchedWritableInventoryItemTemplateRequest {
	this := PatchedWritableInventoryItemTemplateRequest{}
	return &this
}

// NewPatchedWritableInventoryItemTemplateRequestWithDefaults instantiates a new PatchedWritableInventoryItemTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableInventoryItemTemplateRequestWithDefaults() *PatchedWritableInventoryItemTemplateRequest {
	this := PatchedWritableInventoryItemTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *PatchedWritableInventoryItemTemplateRequest) GetDeviceType() int32 {
	if o == nil || IsNil(o.DeviceType) {
		var ret int32
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) GetDeviceTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given int32 and assigns it to the DeviceType field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetDeviceType(v int32) {
	o.DeviceType = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInventoryItemTemplateRequest) GetParent() int32 {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret int32
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInventoryItemTemplateRequest) GetParentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableInt32 and assigns it to the Parent field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetParent(v int32) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableInventoryItemTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableInventoryItemTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInventoryItemTemplateRequest) GetRole() int32 {
	if o == nil || IsNil(o.Role.Get()) {
		var ret int32
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInventoryItemTemplateRequest) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableInt32 and assigns it to the Role field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetRole(v int32) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) UnsetRole() {
	o.Role.Unset()
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInventoryItemTemplateRequest) GetManufacturer() int32 {
	if o == nil || IsNil(o.Manufacturer.Get()) {
		var ret int32
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInventoryItemTemplateRequest) GetManufacturerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableInt32 and assigns it to the Manufacturer field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetManufacturer(v int32) {
	o.Manufacturer.Set(&v)
}

// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *PatchedWritableInventoryItemTemplateRequest) GetPartId() string {
	if o == nil || IsNil(o.PartId) {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) GetPartIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartId) {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasPartId() bool {
	if o != nil && !IsNil(o.PartId) {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetPartId(v string) {
	o.PartId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableInventoryItemTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInventoryItemTemplateRequest) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType.Get()) {
		var ret string
		return ret
	}
	return *o.ComponentType.Get()
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInventoryItemTemplateRequest) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentType.Get(), o.ComponentType.IsSet()
}

// HasComponentType returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasComponentType() bool {
	if o != nil && o.ComponentType.IsSet() {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given NullableString and assigns it to the ComponentType field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetComponentType(v string) {
	o.ComponentType.Set(&v)
}

// SetComponentTypeNil sets the value for ComponentType to be an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) SetComponentTypeNil() {
	o.ComponentType.Set(nil)
}

// UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) UnsetComponentType() {
	o.ComponentType.Unset()
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInventoryItemTemplateRequest) GetComponentId() int64 {
	if o == nil || IsNil(o.ComponentId.Get()) {
		var ret int64
		return ret
	}
	return *o.ComponentId.Get()
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInventoryItemTemplateRequest) GetComponentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentId.Get(), o.ComponentId.IsSet()
}

// HasComponentId returns a boolean if a field has been set.
func (o *PatchedWritableInventoryItemTemplateRequest) HasComponentId() bool {
	if o != nil && o.ComponentId.IsSet() {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given NullableInt64 and assigns it to the ComponentId field.
func (o *PatchedWritableInventoryItemTemplateRequest) SetComponentId(v int64) {
	o.ComponentId.Set(&v)
}

// SetComponentIdNil sets the value for ComponentId to be an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) SetComponentIdNil() {
	o.ComponentId.Set(nil)
}

// UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil
func (o *PatchedWritableInventoryItemTemplateRequest) UnsetComponentId() {
	o.ComponentId.Unset()
}

func (o PatchedWritableInventoryItemTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableInventoryItemTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableInventoryItemTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableInventoryItemTemplateRequest := _PatchedWritableInventoryItemTemplateRequest{}

	err = json.Unmarshal(data, &varPatchedWritableInventoryItemTemplateRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableInventoryItemTemplateRequest(varPatchedWritableInventoryItemTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "role")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "part_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "component_type")
		delete(additionalProperties, "component_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableInventoryItemTemplateRequest struct {
	value *PatchedWritableInventoryItemTemplateRequest
	isSet bool
}

func (v NullablePatchedWritableInventoryItemTemplateRequest) Get() *PatchedWritableInventoryItemTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritableInventoryItemTemplateRequest) Set(val *PatchedWritableInventoryItemTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableInventoryItemTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableInventoryItemTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableInventoryItemTemplateRequest(val *PatchedWritableInventoryItemTemplateRequest) *NullablePatchedWritableInventoryItemTemplateRequest {
	return &NullablePatchedWritableInventoryItemTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableInventoryItemTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableInventoryItemTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
