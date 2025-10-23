# BriefPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DeviceCount** | Pointer to **int32** |  | [optional] [readonly] [default to 0]
**VirtualmachineCount** | Pointer to **int32** |  | [optional] [readonly] [default to 0]
**Depth** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewBriefPlatform

`func NewBriefPlatform(id int32, url string, name string, slug string, ) *BriefPlatform`

NewBriefPlatform instantiates a new BriefPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefPlatformWithDefaults

`func NewBriefPlatformWithDefaults() *BriefPlatform`

NewBriefPlatformWithDefaults instantiates a new BriefPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefPlatform) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefPlatform) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefPlatform) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefPlatform) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefPlatform) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefPlatform) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefPlatform) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefPlatform) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefPlatform) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *BriefPlatform) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *BriefPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefPlatform) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefPlatform) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefPlatform) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefPlatform) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefPlatform) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefPlatform) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefPlatform) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefPlatform) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceCount

`func (o *BriefPlatform) GetDeviceCount() int32`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *BriefPlatform) GetDeviceCountOk() (*int32, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *BriefPlatform) SetDeviceCount(v int32)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *BriefPlatform) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetVirtualmachineCount

`func (o *BriefPlatform) GetVirtualmachineCount() int32`

GetVirtualmachineCount returns the VirtualmachineCount field if non-nil, zero value otherwise.

### GetVirtualmachineCountOk

`func (o *BriefPlatform) GetVirtualmachineCountOk() (*int32, bool)`

GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualmachineCount

`func (o *BriefPlatform) SetVirtualmachineCount(v int32)`

SetVirtualmachineCount sets VirtualmachineCount field to given value.

### HasVirtualmachineCount

`func (o *BriefPlatform) HasVirtualmachineCount() bool`

HasVirtualmachineCount returns a boolean if a field has been set.

### GetDepth

`func (o *BriefPlatform) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *BriefPlatform) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *BriefPlatform) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *BriefPlatform) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


