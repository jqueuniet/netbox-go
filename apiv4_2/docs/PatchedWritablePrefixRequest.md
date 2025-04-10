# PatchedWritablePrefixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** |  | [optional] 
**Vrf** | Pointer to [**NullableIPAddressRequestVrf**](IPAddressRequestVrf.md) |  | [optional] 
**ScopeType** | Pointer to **NullableString** |  | [optional] 
**ScopeId** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Vlan** | Pointer to [**NullableInterfaceRequestUntaggedVlan**](InterfaceRequestUntaggedVlan.md) |  | [optional] 
**Status** | Pointer to [**PatchedWritablePrefixRequestStatus**](PatchedWritablePrefixRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableIPRangeRequestRole**](IPRangeRequestRole.md) |  | [optional] 
**IsPool** | Pointer to **bool** | All IP addresses within this prefix are considered usable | [optional] 
**MarkUtilized** | Pointer to **bool** | Treat as fully utilized | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritablePrefixRequest

`func NewPatchedWritablePrefixRequest() *PatchedWritablePrefixRequest`

NewPatchedWritablePrefixRequest instantiates a new PatchedWritablePrefixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePrefixRequestWithDefaults

`func NewPatchedWritablePrefixRequestWithDefaults() *PatchedWritablePrefixRequest`

NewPatchedWritablePrefixRequestWithDefaults instantiates a new PatchedWritablePrefixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *PatchedWritablePrefixRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PatchedWritablePrefixRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PatchedWritablePrefixRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PatchedWritablePrefixRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetVrf

`func (o *PatchedWritablePrefixRequest) GetVrf() IPAddressRequestVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedWritablePrefixRequest) GetVrfOk() (*IPAddressRequestVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedWritablePrefixRequest) SetVrf(v IPAddressRequestVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedWritablePrefixRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *PatchedWritablePrefixRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *PatchedWritablePrefixRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetScopeType

`func (o *PatchedWritablePrefixRequest) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *PatchedWritablePrefixRequest) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *PatchedWritablePrefixRequest) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *PatchedWritablePrefixRequest) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### SetScopeTypeNil

`func (o *PatchedWritablePrefixRequest) SetScopeTypeNil(b bool)`

 SetScopeTypeNil sets the value for ScopeType to be an explicit nil

### UnsetScopeType
`func (o *PatchedWritablePrefixRequest) UnsetScopeType()`

UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
### GetScopeId

`func (o *PatchedWritablePrefixRequest) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *PatchedWritablePrefixRequest) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *PatchedWritablePrefixRequest) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *PatchedWritablePrefixRequest) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### SetScopeIdNil

`func (o *PatchedWritablePrefixRequest) SetScopeIdNil(b bool)`

 SetScopeIdNil sets the value for ScopeId to be an explicit nil

### UnsetScopeId
`func (o *PatchedWritablePrefixRequest) UnsetScopeId()`

UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
### GetTenant

`func (o *PatchedWritablePrefixRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritablePrefixRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritablePrefixRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritablePrefixRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritablePrefixRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritablePrefixRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetVlan

`func (o *PatchedWritablePrefixRequest) GetVlan() InterfaceRequestUntaggedVlan`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PatchedWritablePrefixRequest) GetVlanOk() (*InterfaceRequestUntaggedVlan, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PatchedWritablePrefixRequest) SetVlan(v InterfaceRequestUntaggedVlan)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PatchedWritablePrefixRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *PatchedWritablePrefixRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *PatchedWritablePrefixRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetStatus

`func (o *PatchedWritablePrefixRequest) GetStatus() PatchedWritablePrefixRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritablePrefixRequest) GetStatusOk() (*PatchedWritablePrefixRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritablePrefixRequest) SetStatus(v PatchedWritablePrefixRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritablePrefixRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritablePrefixRequest) GetRole() IPRangeRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritablePrefixRequest) GetRoleOk() (*IPRangeRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritablePrefixRequest) SetRole(v IPRangeRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritablePrefixRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedWritablePrefixRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedWritablePrefixRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetIsPool

`func (o *PatchedWritablePrefixRequest) GetIsPool() bool`

GetIsPool returns the IsPool field if non-nil, zero value otherwise.

### GetIsPoolOk

`func (o *PatchedWritablePrefixRequest) GetIsPoolOk() (*bool, bool)`

GetIsPoolOk returns a tuple with the IsPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPool

`func (o *PatchedWritablePrefixRequest) SetIsPool(v bool)`

SetIsPool sets IsPool field to given value.

### HasIsPool

`func (o *PatchedWritablePrefixRequest) HasIsPool() bool`

HasIsPool returns a boolean if a field has been set.

### GetMarkUtilized

`func (o *PatchedWritablePrefixRequest) GetMarkUtilized() bool`

GetMarkUtilized returns the MarkUtilized field if non-nil, zero value otherwise.

### GetMarkUtilizedOk

`func (o *PatchedWritablePrefixRequest) GetMarkUtilizedOk() (*bool, bool)`

GetMarkUtilizedOk returns a tuple with the MarkUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkUtilized

`func (o *PatchedWritablePrefixRequest) SetMarkUtilized(v bool)`

SetMarkUtilized sets MarkUtilized field to given value.

### HasMarkUtilized

`func (o *PatchedWritablePrefixRequest) HasMarkUtilized() bool`

HasMarkUtilized returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritablePrefixRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritablePrefixRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritablePrefixRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritablePrefixRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritablePrefixRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritablePrefixRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritablePrefixRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritablePrefixRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritablePrefixRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritablePrefixRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritablePrefixRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritablePrefixRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritablePrefixRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritablePrefixRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritablePrefixRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritablePrefixRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


