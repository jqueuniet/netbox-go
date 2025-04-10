# InventoryItemTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**DeviceType** | [**NestedDeviceType**](NestedDeviceType.md) |  | 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |          {module} is accepted as a substitution for the module bay position when attached to a module type.          | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Role** | Pointer to [**NullableNestedInventoryItemRole**](NestedInventoryItemRole.md) |  | [optional] 
**Manufacturer** | Pointer to [**NullableNestedManufacturer**](NestedManufacturer.md) |  | [optional] 
**PartId** | Pointer to **string** | Manufacturer-assigned part identifier | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ComponentType** | Pointer to **NullableString** |  | [optional] 
**ComponentId** | Pointer to **NullableInt64** |  | [optional] 
**Component** | **map[string]interface{}** |  | [readonly] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Depth** | **int32** |  | [readonly] 

## Methods

### NewInventoryItemTemplate

`func NewInventoryItemTemplate(id int32, url string, deviceType NestedDeviceType, name string, component map[string]interface{}, lastUpdated NullableTime, depth int32, ) *InventoryItemTemplate`

NewInventoryItemTemplate instantiates a new InventoryItemTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemTemplateWithDefaults

`func NewInventoryItemTemplateWithDefaults() *InventoryItemTemplate`

NewInventoryItemTemplateWithDefaults instantiates a new InventoryItemTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryItemTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryItemTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryItemTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *InventoryItemTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InventoryItemTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InventoryItemTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *InventoryItemTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InventoryItemTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InventoryItemTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *InventoryItemTemplate) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDeviceType

`func (o *InventoryItemTemplate) GetDeviceType() NestedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InventoryItemTemplate) GetDeviceTypeOk() (*NestedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InventoryItemTemplate) SetDeviceType(v NestedDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetParent

`func (o *InventoryItemTemplate) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *InventoryItemTemplate) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *InventoryItemTemplate) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *InventoryItemTemplate) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *InventoryItemTemplate) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *InventoryItemTemplate) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetName

`func (o *InventoryItemTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryItemTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryItemTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *InventoryItemTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InventoryItemTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InventoryItemTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InventoryItemTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRole

`func (o *InventoryItemTemplate) GetRole() NestedInventoryItemRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InventoryItemTemplate) GetRoleOk() (*NestedInventoryItemRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InventoryItemTemplate) SetRole(v NestedInventoryItemRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *InventoryItemTemplate) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *InventoryItemTemplate) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *InventoryItemTemplate) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetManufacturer

`func (o *InventoryItemTemplate) GetManufacturer() NestedManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *InventoryItemTemplate) GetManufacturerOk() (*NestedManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *InventoryItemTemplate) SetManufacturer(v NestedManufacturer)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *InventoryItemTemplate) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *InventoryItemTemplate) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *InventoryItemTemplate) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetPartId

`func (o *InventoryItemTemplate) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *InventoryItemTemplate) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *InventoryItemTemplate) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *InventoryItemTemplate) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetDescription

`func (o *InventoryItemTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InventoryItemTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InventoryItemTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InventoryItemTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComponentType

`func (o *InventoryItemTemplate) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *InventoryItemTemplate) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *InventoryItemTemplate) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *InventoryItemTemplate) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### SetComponentTypeNil

`func (o *InventoryItemTemplate) SetComponentTypeNil(b bool)`

 SetComponentTypeNil sets the value for ComponentType to be an explicit nil

### UnsetComponentType
`func (o *InventoryItemTemplate) UnsetComponentType()`

UnsetComponentType ensures that no value is present for ComponentType, not even an explicit nil
### GetComponentId

`func (o *InventoryItemTemplate) GetComponentId() int64`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *InventoryItemTemplate) GetComponentIdOk() (*int64, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *InventoryItemTemplate) SetComponentId(v int64)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *InventoryItemTemplate) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### SetComponentIdNil

`func (o *InventoryItemTemplate) SetComponentIdNil(b bool)`

 SetComponentIdNil sets the value for ComponentId to be an explicit nil

### UnsetComponentId
`func (o *InventoryItemTemplate) UnsetComponentId()`

UnsetComponentId ensures that no value is present for ComponentId, not even an explicit nil
### GetComponent

`func (o *InventoryItemTemplate) GetComponent() map[string]interface{}`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *InventoryItemTemplate) GetComponentOk() (*map[string]interface{}, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *InventoryItemTemplate) SetComponent(v map[string]interface{})`

SetComponent sets Component field to given value.


### SetComponentNil

`func (o *InventoryItemTemplate) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *InventoryItemTemplate) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil
### GetCreated

`func (o *InventoryItemTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InventoryItemTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InventoryItemTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *InventoryItemTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *InventoryItemTemplate) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *InventoryItemTemplate) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *InventoryItemTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InventoryItemTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InventoryItemTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *InventoryItemTemplate) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *InventoryItemTemplate) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetDepth

`func (o *InventoryItemTemplate) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *InventoryItemTemplate) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *InventoryItemTemplate) SetDepth(v int32)`

SetDepth sets Depth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


