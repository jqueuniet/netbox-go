# NestedIPSecProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 

## Methods

### NewNestedIPSecProfile

`func NewNestedIPSecProfile(id int32, url string, name string, ) *NestedIPSecProfile`

NewNestedIPSecProfile instantiates a new NestedIPSecProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedIPSecProfileWithDefaults

`func NewNestedIPSecProfileWithDefaults() *NestedIPSecProfile`

NewNestedIPSecProfileWithDefaults instantiates a new NestedIPSecProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedIPSecProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedIPSecProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedIPSecProfile) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedIPSecProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedIPSecProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedIPSecProfile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *NestedIPSecProfile) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedIPSecProfile) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedIPSecProfile) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *NestedIPSecProfile) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetName

`func (o *NestedIPSecProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedIPSecProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedIPSecProfile) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


