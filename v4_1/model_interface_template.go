/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.6 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v4_1

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the InterfaceTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceTemplate{}

// InterfaceTemplate Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type InterfaceTemplate struct {
	Id         int32                   `json:"id"`
	Url        string                  `json:"url"`
	Display    string                  `json:"display"`
	DeviceType NullableBriefDeviceType `json:"device_type,omitempty"`
	ModuleType NullableBriefModuleType `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label                *string                          `json:"label,omitempty"`
	Type                 InterfaceType                    `json:"type"`
	Enabled              *bool                            `json:"enabled,omitempty"`
	MgmtOnly             *bool                            `json:"mgmt_only,omitempty"`
	Description          *string                          `json:"description,omitempty"`
	Bridge               NullableNestedInterfaceTemplate  `json:"bridge,omitempty"`
	PoeMode              NullableInterfaceTemplatePoeMode `json:"poe_mode,omitempty"`
	PoeType              NullableInterfaceTemplatePoeType `json:"poe_type,omitempty"`
	RfRole               NullableInterfaceTemplateRfRole  `json:"rf_role,omitempty"`
	Created              NullableTime                     `json:"created"`
	LastUpdated          NullableTime                     `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceTemplate InterfaceTemplate

// NewInterfaceTemplate instantiates a new InterfaceTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceTemplate(id int32, url string, display string, name string, type_ InterfaceType, created NullableTime, lastUpdated NullableTime) *InterfaceTemplate {
	this := InterfaceTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Type = type_
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewInterfaceTemplateWithDefaults instantiates a new InterfaceTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTemplateWithDefaults() *InterfaceTemplate {
	this := InterfaceTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *InterfaceTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InterfaceTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *InterfaceTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InterfaceTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *InterfaceTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *InterfaceTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplate) GetDeviceType() BriefDeviceType {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BriefDeviceType
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplate) GetDeviceTypeOk() (*BriefDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBriefDeviceType and assigns it to the DeviceType field.
func (o *InterfaceTemplate) SetDeviceType(v BriefDeviceType) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *InterfaceTemplate) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *InterfaceTemplate) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplate) GetModuleType() BriefModuleType {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BriefModuleType
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplate) GetModuleTypeOk() (*BriefModuleType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBriefModuleType and assigns it to the ModuleType field.
func (o *InterfaceTemplate) SetModuleType(v BriefModuleType) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *InterfaceTemplate) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *InterfaceTemplate) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *InterfaceTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InterfaceTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceTemplate) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InterfaceTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *InterfaceTemplate) GetType() InterfaceType {
	if o == nil {
		var ret InterfaceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetTypeOk() (*InterfaceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InterfaceTemplate) SetType(v InterfaceType) {
	o.Type = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InterfaceTemplate) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InterfaceTemplate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMgmtOnly returns the MgmtOnly field value if set, zero value otherwise.
func (o *InterfaceTemplate) GetMgmtOnly() bool {
	if o == nil || IsNil(o.MgmtOnly) {
		var ret bool
		return ret
	}
	return *o.MgmtOnly
}

// GetMgmtOnlyOk returns a tuple with the MgmtOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetMgmtOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.MgmtOnly) {
		return nil, false
	}
	return o.MgmtOnly, true
}

// HasMgmtOnly returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasMgmtOnly() bool {
	if o != nil && !IsNil(o.MgmtOnly) {
		return true
	}

	return false
}

// SetMgmtOnly gets a reference to the given bool and assigns it to the MgmtOnly field.
func (o *InterfaceTemplate) SetMgmtOnly(v bool) {
	o.MgmtOnly = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InterfaceTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InterfaceTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplate) GetBridge() NestedInterfaceTemplate {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret NestedInterfaceTemplate
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplate) GetBridgeOk() (*NestedInterfaceTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableNestedInterfaceTemplate and assigns it to the Bridge field.
func (o *InterfaceTemplate) SetBridge(v NestedInterfaceTemplate) {
	o.Bridge.Set(&v)
}

// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *InterfaceTemplate) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *InterfaceTemplate) UnsetBridge() {
	o.Bridge.Unset()
}

// GetPoeMode returns the PoeMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplate) GetPoeMode() InterfaceTemplatePoeMode {
	if o == nil || IsNil(o.PoeMode.Get()) {
		var ret InterfaceTemplatePoeMode
		return ret
	}
	return *o.PoeMode.Get()
}

// GetPoeModeOk returns a tuple with the PoeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplate) GetPoeModeOk() (*InterfaceTemplatePoeMode, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoeMode.Get(), o.PoeMode.IsSet()
}

// HasPoeMode returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasPoeMode() bool {
	if o != nil && o.PoeMode.IsSet() {
		return true
	}

	return false
}

// SetPoeMode gets a reference to the given NullableInterfaceTemplatePoeMode and assigns it to the PoeMode field.
func (o *InterfaceTemplate) SetPoeMode(v InterfaceTemplatePoeMode) {
	o.PoeMode.Set(&v)
}

// SetPoeModeNil sets the value for PoeMode to be an explicit nil
func (o *InterfaceTemplate) SetPoeModeNil() {
	o.PoeMode.Set(nil)
}

// UnsetPoeMode ensures that no value is present for PoeMode, not even an explicit nil
func (o *InterfaceTemplate) UnsetPoeMode() {
	o.PoeMode.Unset()
}

// GetPoeType returns the PoeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplate) GetPoeType() InterfaceTemplatePoeType {
	if o == nil || IsNil(o.PoeType.Get()) {
		var ret InterfaceTemplatePoeType
		return ret
	}
	return *o.PoeType.Get()
}

// GetPoeTypeOk returns a tuple with the PoeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplate) GetPoeTypeOk() (*InterfaceTemplatePoeType, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoeType.Get(), o.PoeType.IsSet()
}

// HasPoeType returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasPoeType() bool {
	if o != nil && o.PoeType.IsSet() {
		return true
	}

	return false
}

// SetPoeType gets a reference to the given NullableInterfaceTemplatePoeType and assigns it to the PoeType field.
func (o *InterfaceTemplate) SetPoeType(v InterfaceTemplatePoeType) {
	o.PoeType.Set(&v)
}

// SetPoeTypeNil sets the value for PoeType to be an explicit nil
func (o *InterfaceTemplate) SetPoeTypeNil() {
	o.PoeType.Set(nil)
}

// UnsetPoeType ensures that no value is present for PoeType, not even an explicit nil
func (o *InterfaceTemplate) UnsetPoeType() {
	o.PoeType.Unset()
}

// GetRfRole returns the RfRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceTemplate) GetRfRole() InterfaceTemplateRfRole {
	if o == nil || IsNil(o.RfRole.Get()) {
		var ret InterfaceTemplateRfRole
		return ret
	}
	return *o.RfRole.Get()
}

// GetRfRoleOk returns a tuple with the RfRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplate) GetRfRoleOk() (*InterfaceTemplateRfRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.RfRole.Get(), o.RfRole.IsSet()
}

// HasRfRole returns a boolean if a field has been set.
func (o *InterfaceTemplate) HasRfRole() bool {
	if o != nil && o.RfRole.IsSet() {
		return true
	}

	return false
}

// SetRfRole gets a reference to the given NullableInterfaceTemplateRfRole and assigns it to the RfRole field.
func (o *InterfaceTemplate) SetRfRole(v InterfaceTemplateRfRole) {
	o.RfRole.Set(&v)
}

// SetRfRoleNil sets the value for RfRole to be an explicit nil
func (o *InterfaceTemplate) SetRfRoleNil() {
	o.RfRole.Set(nil)
}

// UnsetRfRole ensures that no value is present for RfRole, not even an explicit nil
func (o *InterfaceTemplate) UnsetRfRole() {
	o.RfRole.Unset()
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *InterfaceTemplate) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *InterfaceTemplate) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *InterfaceTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *InterfaceTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o InterfaceTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MgmtOnly) {
		toSerialize["mgmt_only"] = o.MgmtOnly
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if o.PoeMode.IsSet() {
		toSerialize["poe_mode"] = o.PoeMode.Get()
	}
	if o.PoeType.IsSet() {
		toSerialize["poe_type"] = o.PoeType.Get()
	}
	if o.RfRole.IsSet() {
		toSerialize["rf_role"] = o.RfRole.Get()
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterfaceTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"type",
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
	varInterfaceTemplate := _InterfaceTemplate{}

	err = json.Unmarshal(data, &varInterfaceTemplate)

	if err != nil {
		return err
	}

	*o = InterfaceTemplate(varInterfaceTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "mgmt_only")
		delete(additionalProperties, "description")
		delete(additionalProperties, "bridge")
		delete(additionalProperties, "poe_mode")
		delete(additionalProperties, "poe_type")
		delete(additionalProperties, "rf_role")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceTemplate struct {
	value *InterfaceTemplate
	isSet bool
}

func (v NullableInterfaceTemplate) Get() *InterfaceTemplate {
	return v.value
}

func (v *NullableInterfaceTemplate) Set(val *InterfaceTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTemplate(val *InterfaceTemplate) *NullableInterfaceTemplate {
	return &NullableInterfaceTemplate{value: val, isSet: true}
}

func (v NullableInterfaceTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}