# PatchedWritableVLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to [**NullablePatchedWritableVLANRequestSite**](PatchedWritableVLANRequestSite.md) |  | [optional] 
**Group** | Pointer to [**NullablePatchedWritableVLANRequestGroup**](PatchedWritableVLANRequestGroup.md) |  | [optional] 
**Vid** | Pointer to **int32** | Numeric VLAN ID (1-4094) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Status** | Pointer to [**PatchedWritableVLANRequestStatus**](PatchedWritableVLANRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableIPRangeRequestRole**](IPRangeRequestRole.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**QinqRole** | Pointer to [**NullableQInQRole**](QInQRole.md) |  | [optional] 
**QinqSvlan** | Pointer to **NullableInt32** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableVLANRequest

`func NewPatchedWritableVLANRequest() *PatchedWritableVLANRequest`

NewPatchedWritableVLANRequest instantiates a new PatchedWritableVLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableVLANRequestWithDefaults

`func NewPatchedWritableVLANRequestWithDefaults() *PatchedWritableVLANRequest`

NewPatchedWritableVLANRequestWithDefaults instantiates a new PatchedWritableVLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *PatchedWritableVLANRequest) GetSite() PatchedWritableVLANRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableVLANRequest) GetSiteOk() (*PatchedWritableVLANRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableVLANRequest) SetSite(v PatchedWritableVLANRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableVLANRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *PatchedWritableVLANRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *PatchedWritableVLANRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetGroup

`func (o *PatchedWritableVLANRequest) GetGroup() PatchedWritableVLANRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedWritableVLANRequest) GetGroupOk() (*PatchedWritableVLANRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedWritableVLANRequest) SetGroup(v PatchedWritableVLANRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedWritableVLANRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedWritableVLANRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedWritableVLANRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetVid

`func (o *PatchedWritableVLANRequest) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *PatchedWritableVLANRequest) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *PatchedWritableVLANRequest) SetVid(v int32)`

SetVid sets Vid field to given value.

### HasVid

`func (o *PatchedWritableVLANRequest) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableVLANRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableVLANRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableVLANRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableVLANRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableVLANRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableVLANRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableVLANRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableVLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableVLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableVLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *PatchedWritableVLANRequest) GetStatus() PatchedWritableVLANRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableVLANRequest) GetStatusOk() (*PatchedWritableVLANRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableVLANRequest) SetStatus(v PatchedWritableVLANRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableVLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableVLANRequest) GetRole() IPRangeRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableVLANRequest) GetRoleOk() (*IPRangeRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableVLANRequest) SetRole(v IPRangeRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableVLANRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedWritableVLANRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedWritableVLANRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *PatchedWritableVLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableVLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableVLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableVLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQinqRole

`func (o *PatchedWritableVLANRequest) GetQinqRole() QInQRole`

GetQinqRole returns the QinqRole field if non-nil, zero value otherwise.

### GetQinqRoleOk

`func (o *PatchedWritableVLANRequest) GetQinqRoleOk() (*QInQRole, bool)`

GetQinqRoleOk returns a tuple with the QinqRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqRole

`func (o *PatchedWritableVLANRequest) SetQinqRole(v QInQRole)`

SetQinqRole sets QinqRole field to given value.

### HasQinqRole

`func (o *PatchedWritableVLANRequest) HasQinqRole() bool`

HasQinqRole returns a boolean if a field has been set.

### SetQinqRoleNil

`func (o *PatchedWritableVLANRequest) SetQinqRoleNil(b bool)`

 SetQinqRoleNil sets the value for QinqRole to be an explicit nil

### UnsetQinqRole
`func (o *PatchedWritableVLANRequest) UnsetQinqRole()`

UnsetQinqRole ensures that no value is present for QinqRole, not even an explicit nil
### GetQinqSvlan

`func (o *PatchedWritableVLANRequest) GetQinqSvlan() int32`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *PatchedWritableVLANRequest) GetQinqSvlanOk() (*int32, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *PatchedWritableVLANRequest) SetQinqSvlan(v int32)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *PatchedWritableVLANRequest) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *PatchedWritableVLANRequest) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *PatchedWritableVLANRequest) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetComments

`func (o *PatchedWritableVLANRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableVLANRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableVLANRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableVLANRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableVLANRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableVLANRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableVLANRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableVLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableVLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableVLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableVLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableVLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


