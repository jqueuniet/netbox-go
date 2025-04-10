/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.11 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiv4_1

import (
	"encoding/json"
)

// checks if the PatchedInventoryItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedInventoryItemRequest{}

// PatchedInventoryItemRequest Adds support for custom fields and tags.
type PatchedInventoryItemRequest struct {
	Device *BriefDeviceRequest `json:"device,omitempty"`
	Parent NullableInt32       `json:"parent,omitempty"`
	Name   *string             `json:"name,omitempty"`
	// Physical label
	Label        *string                               `json:"label,omitempty"`
	Role         NullableBriefInventoryItemRoleRequest `json:"role,omitempty"`
	Manufacturer NullableBriefManufacturerRequest      `json:"manufacturer,omitempty"`
	// Manufacturer-assigned part identifier
	PartId *string `json:"part_id,omitempty"`
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this item
	AssetTag NullableString `json:"asset_tag,omitempty"`
	// This item was automatically discovered
	Discovered           *bool                  `json:"discovered,omitempty"`
	Description          *string                `json:"description,omitempty"`
	ComponentType        NullableString         `json:"component_type,omitempty"`
	ComponentId          NullableInt64          `json:"component_id,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedInventoryItemRequest PatchedInventoryItemRequest

// NewPatchedInventoryItemRequest instantiates a new PatchedInventoryItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedInventoryItemRequest() *PatchedInventoryItemRequest {
	this := PatchedInventoryItemRequest{}
	return &this
}

// NewPatchedInventoryItemRequestWithDefaults instantiates a new PatchedInventoryItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedInventoryItemRequestWithDefaults() *PatchedInventoryItemRequest {
	this := PatchedInventoryItemRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedInventoryItemRequest) GetDevice() BriefDeviceRequest {
	if o == nil || IsNil(o.Device) {
		var ret BriefDeviceRequest
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given BriefDeviceRequest and assigns it to the Device field.
func (o *PatchedInventoryItemRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemRequest) GetParent() int32 {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret int32
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemRequest) GetParentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableInt32 and assigns it to the Parent field.
func (o *PatchedInventoryItemRequest) SetParent(v int32) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedInventoryItemRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedInventoryItemRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedInventoryItemRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedInventoryItemRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedInventoryItemRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedInventoryItemRequest) SetLabel(v string) {
	o.Label = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemRequest) GetRole() BriefInventoryItemRoleRequest {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefInventoryItemRoleRequest
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemRequest) GetRoleOk() (*BriefInventoryItemRoleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefInventoryItemRoleRequest and assigns it to the Role field.
func (o *PatchedInventoryItemRequest) SetRole(v BriefInventoryItemRoleRequest) {
	o.Role.Set(&v)
}

// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedInventoryItemRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedInventoryItemRequest) UnsetRole() {
	o.Role.Unset()
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemRequest) GetManufacturer() BriefManufacturerRequest {
	if o == nil || IsNil(o.Manufacturer.Get()) {
		var ret BriefManufacturerRequest
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableBriefManufacturerRequest and assigns it to the Manufacturer field.
func (o *PatchedInventoryItemRequest) SetManufacturer(v BriefManufacturerRequest) {
	o.Manufacturer.Set(&v)
}

// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *PatchedInventoryItemRequest) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *PatchedInventoryItemRequest) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *PatchedInventoryItemRequest) GetPartId() string {
	if o == nil || IsNil(o.PartId) {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemRequest) GetPartIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartId) {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasPartId() bool {
	if o != nil && !IsNil(o.PartId) {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *PatchedInventoryItemRequest) SetPartId(v string) {
	o.PartId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *PatchedInventoryItemRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *PatchedInventoryItemRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *PatchedInventoryItemRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *PatchedInventoryItemRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *PatchedInventoryItemRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *PatchedInventoryItemRequest) GetDiscovered() bool {
	if o == nil || IsNil(o.Discovered) {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemRequest) GetDiscoveredOk() (*bool, bool) {
	if o == nil || IsNil(o.Discovered) {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasDiscovered() bool {
	if o != nil && !IsNil(o.Discovered) {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *PatchedInventoryItemRequest) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedInventoryItemRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedInventoryItemRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemRequest) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType.Get()) {
		var ret string
		return ret
	}
	return *o.ComponentType.Get()
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemRequest) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentType.Get(), o.ComponentType.IsSet()
}

// HasComponentType returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasComponentType() bool {
	if o != nil && o.ComponentType.IsSet() {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given NullableString and assigns it to the ComponentType field.
func (o *PatchedInventoryItemRequest) SetComponentType(v string) {
	o.ComponentType.Set(&v)
}

// SetComponentTypeNil sets the value for ComponentType to be an explicit nil
func (o *PatchedInventoryItemRequest) SetComponentTypeNil() {
	o.ComponentType.Set(nil)
}

// UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
func (o *PatchedInventoryItemRequest) UnsetComponentType() {
	o.ComponentType.Unset()
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedInventoryItemRequest) GetComponentId() int64 {
	if o == nil || IsNil(o.ComponentId.Get()) {
		var ret int64
		return ret
	}
	return *o.ComponentId.Get()
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedInventoryItemRequest) GetComponentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComponentId.Get(), o.ComponentId.IsSet()
}

// HasComponentId returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasComponentId() bool {
	if o != nil && o.ComponentId.IsSet() {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given NullableInt64 and assigns it to the ComponentId field.
func (o *PatchedInventoryItemRequest) SetComponentId(v int64) {
	o.ComponentId.Set(&v)
}

// SetComponentIdNil sets the value for ComponentId to be an explicit nil
func (o *PatchedInventoryItemRequest) SetComponentIdNil() {
	o.ComponentId.Set(nil)
}

// UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil
func (o *PatchedInventoryItemRequest) UnsetComponentId() {
	o.ComponentId.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedInventoryItemRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedInventoryItemRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedInventoryItemRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedInventoryItemRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedInventoryItemRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedInventoryItemRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedInventoryItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedInventoryItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
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
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	if !IsNil(o.Discovered) {
		toSerialize["discovered"] = o.Discovered
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
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedInventoryItemRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedInventoryItemRequest := _PatchedInventoryItemRequest{}

	err = json.Unmarshal(data, &varPatchedInventoryItemRequest)

	if err != nil {
		return err
	}

	*o = PatchedInventoryItemRequest(varPatchedInventoryItemRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "role")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "part_id")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "discovered")
		delete(additionalProperties, "description")
		delete(additionalProperties, "component_type")
		delete(additionalProperties, "component_id")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedInventoryItemRequest struct {
	value *PatchedInventoryItemRequest
	isSet bool
}

func (v NullablePatchedInventoryItemRequest) Get() *PatchedInventoryItemRequest {
	return v.value
}

func (v *NullablePatchedInventoryItemRequest) Set(val *PatchedInventoryItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedInventoryItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedInventoryItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedInventoryItemRequest(val *PatchedInventoryItemRequest) *NullablePatchedInventoryItemRequest {
	return &NullablePatchedInventoryItemRequest{value: val, isSet: true}
}

func (v NullablePatchedInventoryItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedInventoryItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
