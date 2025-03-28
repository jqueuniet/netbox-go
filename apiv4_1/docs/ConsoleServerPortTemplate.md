# ConsoleServerPortTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**DeviceType** | Pointer to [**NullableBriefDeviceType**](BriefDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleType**](BriefModuleType.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**ConsolePortType**](ConsolePortType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewConsoleServerPortTemplate

`func NewConsoleServerPortTemplate(id int32, url string, name string, lastUpdated NullableTime, ) *ConsoleServerPortTemplate`

NewConsoleServerPortTemplate instantiates a new ConsoleServerPortTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleServerPortTemplateWithDefaults

`func NewConsoleServerPortTemplateWithDefaults() *ConsoleServerPortTemplate`

NewConsoleServerPortTemplateWithDefaults instantiates a new ConsoleServerPortTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleServerPortTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleServerPortTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleServerPortTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ConsoleServerPortTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConsoleServerPortTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConsoleServerPortTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *ConsoleServerPortTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConsoleServerPortTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConsoleServerPortTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ConsoleServerPortTemplate) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDeviceType

`func (o *ConsoleServerPortTemplate) GetDeviceType() BriefDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ConsoleServerPortTemplate) GetDeviceTypeOk() (*BriefDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ConsoleServerPortTemplate) SetDeviceType(v BriefDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ConsoleServerPortTemplate) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *ConsoleServerPortTemplate) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *ConsoleServerPortTemplate) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *ConsoleServerPortTemplate) GetModuleType() BriefModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *ConsoleServerPortTemplate) GetModuleTypeOk() (*BriefModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *ConsoleServerPortTemplate) SetModuleType(v BriefModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *ConsoleServerPortTemplate) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *ConsoleServerPortTemplate) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *ConsoleServerPortTemplate) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *ConsoleServerPortTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleServerPortTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleServerPortTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ConsoleServerPortTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConsoleServerPortTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConsoleServerPortTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConsoleServerPortTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *ConsoleServerPortTemplate) GetType() ConsolePortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsoleServerPortTemplate) GetTypeOk() (*ConsolePortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsoleServerPortTemplate) SetType(v ConsolePortType)`

SetType sets Type field to given value.

### HasType

`func (o *ConsoleServerPortTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ConsoleServerPortTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleServerPortTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleServerPortTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleServerPortTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *ConsoleServerPortTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConsoleServerPortTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConsoleServerPortTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConsoleServerPortTemplate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *ConsoleServerPortTemplate) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ConsoleServerPortTemplate) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ConsoleServerPortTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ConsoleServerPortTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ConsoleServerPortTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ConsoleServerPortTemplate) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ConsoleServerPortTemplate) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


