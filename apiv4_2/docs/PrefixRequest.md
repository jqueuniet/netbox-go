# PrefixRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | **string** |  | 
**Vrf** | Pointer to [**NullableIPAddressRequestVrf**](IPAddressRequestVrf.md) |  | [optional] 
**ScopeType** | Pointer to **NullableString** |  | [optional] 
**ScopeId** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Vlan** | Pointer to [**NullableInterfaceRequestUntaggedVlan**](InterfaceRequestUntaggedVlan.md) |  | [optional] 
**Status** | Pointer to [**PrefixStatusValue**](PrefixStatusValue.md) |  | [optional] 
**Role** | Pointer to [**NullableIPRangeRequestRole**](IPRangeRequestRole.md) |  | [optional] 
**IsPool** | Pointer to **bool** | All IP addresses within this prefix are considered usable | [optional] 
**MarkUtilized** | Pointer to **bool** | Treat as fully utilized | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPrefixRequest

`func NewPrefixRequest(prefix string, ) *PrefixRequest`

NewPrefixRequest instantiates a new PrefixRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixRequestWithDefaults

`func NewPrefixRequestWithDefaults() *PrefixRequest`

NewPrefixRequestWithDefaults instantiates a new PrefixRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *PrefixRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PrefixRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PrefixRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetVrf

`func (o *PrefixRequest) GetVrf() IPAddressRequestVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PrefixRequest) GetVrfOk() (*IPAddressRequestVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PrefixRequest) SetVrf(v IPAddressRequestVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PrefixRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *PrefixRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *PrefixRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetScopeType

`func (o *PrefixRequest) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *PrefixRequest) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *PrefixRequest) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *PrefixRequest) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### SetScopeTypeNil

`func (o *PrefixRequest) SetScopeTypeNil(b bool)`

 SetScopeTypeNil sets the value for ScopeType to be an explicit nil

### UnsetScopeType
`func (o *PrefixRequest) UnsetScopeType()`

UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
### GetScopeId

`func (o *PrefixRequest) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *PrefixRequest) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *PrefixRequest) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *PrefixRequest) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### SetScopeIdNil

`func (o *PrefixRequest) SetScopeIdNil(b bool)`

 SetScopeIdNil sets the value for ScopeId to be an explicit nil

### UnsetScopeId
`func (o *PrefixRequest) UnsetScopeId()`

UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
### GetTenant

`func (o *PrefixRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PrefixRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PrefixRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PrefixRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PrefixRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PrefixRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetVlan

`func (o *PrefixRequest) GetVlan() InterfaceRequestUntaggedVlan`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PrefixRequest) GetVlanOk() (*InterfaceRequestUntaggedVlan, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PrefixRequest) SetVlan(v InterfaceRequestUntaggedVlan)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PrefixRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *PrefixRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *PrefixRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetStatus

`func (o *PrefixRequest) GetStatus() PrefixStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrefixRequest) GetStatusOk() (*PrefixStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrefixRequest) SetStatus(v PrefixStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrefixRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PrefixRequest) GetRole() IPRangeRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PrefixRequest) GetRoleOk() (*IPRangeRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PrefixRequest) SetRole(v IPRangeRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PrefixRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PrefixRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PrefixRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetIsPool

`func (o *PrefixRequest) GetIsPool() bool`

GetIsPool returns the IsPool field if non-nil, zero value otherwise.

### GetIsPoolOk

`func (o *PrefixRequest) GetIsPoolOk() (*bool, bool)`

GetIsPoolOk returns a tuple with the IsPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPool

`func (o *PrefixRequest) SetIsPool(v bool)`

SetIsPool sets IsPool field to given value.

### HasIsPool

`func (o *PrefixRequest) HasIsPool() bool`

HasIsPool returns a boolean if a field has been set.

### GetMarkUtilized

`func (o *PrefixRequest) GetMarkUtilized() bool`

GetMarkUtilized returns the MarkUtilized field if non-nil, zero value otherwise.

### GetMarkUtilizedOk

`func (o *PrefixRequest) GetMarkUtilizedOk() (*bool, bool)`

GetMarkUtilizedOk returns a tuple with the MarkUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkUtilized

`func (o *PrefixRequest) SetMarkUtilized(v bool)`

SetMarkUtilized sets MarkUtilized field to given value.

### HasMarkUtilized

`func (o *PrefixRequest) HasMarkUtilized() bool`

HasMarkUtilized returns a boolean if a field has been set.

### GetDescription

`func (o *PrefixRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrefixRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrefixRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrefixRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PrefixRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PrefixRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PrefixRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PrefixRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PrefixRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PrefixRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PrefixRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PrefixRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PrefixRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PrefixRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PrefixRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PrefixRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


