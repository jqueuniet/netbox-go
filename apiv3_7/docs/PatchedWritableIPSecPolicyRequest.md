# PatchedWritableIPSecPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Proposals** | Pointer to **[]int32** |  | [optional] 
**PfsGroup** | Pointer to **NullableInt32** | Diffie-Hellman group for Perfect Forward Secrecy  * &#x60;1&#x60; - Group 1 * &#x60;2&#x60; - Group 2 * &#x60;5&#x60; - Group 5 * &#x60;14&#x60; - Group 14 * &#x60;15&#x60; - Group 15 * &#x60;16&#x60; - Group 16 * &#x60;17&#x60; - Group 17 * &#x60;18&#x60; - Group 18 * &#x60;19&#x60; - Group 19 * &#x60;20&#x60; - Group 20 * &#x60;21&#x60; - Group 21 * &#x60;22&#x60; - Group 22 * &#x60;23&#x60; - Group 23 * &#x60;24&#x60; - Group 24 * &#x60;25&#x60; - Group 25 * &#x60;26&#x60; - Group 26 * &#x60;27&#x60; - Group 27 * &#x60;28&#x60; - Group 28 * &#x60;29&#x60; - Group 29 * &#x60;30&#x60; - Group 30 * &#x60;31&#x60; - Group 31 * &#x60;32&#x60; - Group 32 * &#x60;33&#x60; - Group 33 * &#x60;34&#x60; - Group 34 | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableIPSecPolicyRequest

`func NewPatchedWritableIPSecPolicyRequest() *PatchedWritableIPSecPolicyRequest`

NewPatchedWritableIPSecPolicyRequest instantiates a new PatchedWritableIPSecPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableIPSecPolicyRequestWithDefaults

`func NewPatchedWritableIPSecPolicyRequestWithDefaults() *PatchedWritableIPSecPolicyRequest`

NewPatchedWritableIPSecPolicyRequestWithDefaults instantiates a new PatchedWritableIPSecPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableIPSecPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableIPSecPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableIPSecPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableIPSecPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableIPSecPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableIPSecPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableIPSecPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableIPSecPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProposals

`func (o *PatchedWritableIPSecPolicyRequest) GetProposals() []int32`

GetProposals returns the Proposals field if non-nil, zero value otherwise.

### GetProposalsOk

`func (o *PatchedWritableIPSecPolicyRequest) GetProposalsOk() (*[]int32, bool)`

GetProposalsOk returns a tuple with the Proposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposals

`func (o *PatchedWritableIPSecPolicyRequest) SetProposals(v []int32)`

SetProposals sets Proposals field to given value.

### HasProposals

`func (o *PatchedWritableIPSecPolicyRequest) HasProposals() bool`

HasProposals returns a boolean if a field has been set.

### GetPfsGroup

`func (o *PatchedWritableIPSecPolicyRequest) GetPfsGroup() int32`

GetPfsGroup returns the PfsGroup field if non-nil, zero value otherwise.

### GetPfsGroupOk

`func (o *PatchedWritableIPSecPolicyRequest) GetPfsGroupOk() (*int32, bool)`

GetPfsGroupOk returns a tuple with the PfsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfsGroup

`func (o *PatchedWritableIPSecPolicyRequest) SetPfsGroup(v int32)`

SetPfsGroup sets PfsGroup field to given value.

### HasPfsGroup

`func (o *PatchedWritableIPSecPolicyRequest) HasPfsGroup() bool`

HasPfsGroup returns a boolean if a field has been set.

### SetPfsGroupNil

`func (o *PatchedWritableIPSecPolicyRequest) SetPfsGroupNil(b bool)`

 SetPfsGroupNil sets the value for PfsGroup to be an explicit nil

### UnsetPfsGroup
`func (o *PatchedWritableIPSecPolicyRequest) UnsetPfsGroup()`

UnsetPfsGroup ensures that no value is present for PfsGroup, not even an explicit nil
### GetComments

`func (o *PatchedWritableIPSecPolicyRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableIPSecPolicyRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableIPSecPolicyRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableIPSecPolicyRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableIPSecPolicyRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableIPSecPolicyRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableIPSecPolicyRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableIPSecPolicyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableIPSecPolicyRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableIPSecPolicyRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableIPSecPolicyRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableIPSecPolicyRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


