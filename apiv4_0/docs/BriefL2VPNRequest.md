# BriefL2VPNRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **NullableInt64** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Type** | Pointer to **string** | * &#x60;vpws&#x60; - VPWS * &#x60;vpls&#x60; - VPLS * &#x60;vxlan&#x60; - VXLAN * &#x60;vxlan-evpn&#x60; - VXLAN-EVPN * &#x60;mpls-evpn&#x60; - MPLS EVPN * &#x60;pbb-evpn&#x60; - PBB EVPN * &#x60;epl&#x60; - EPL * &#x60;evpl&#x60; - EVPL * &#x60;ep-lan&#x60; - Ethernet Private LAN * &#x60;evp-lan&#x60; - Ethernet Virtual Private LAN * &#x60;ep-tree&#x60; - Ethernet Private Tree * &#x60;evp-tree&#x60; - Ethernet Virtual Private Tree | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefL2VPNRequest

`func NewBriefL2VPNRequest(name string, slug string, ) *BriefL2VPNRequest`

NewBriefL2VPNRequest instantiates a new BriefL2VPNRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefL2VPNRequestWithDefaults

`func NewBriefL2VPNRequestWithDefaults() *BriefL2VPNRequest`

NewBriefL2VPNRequestWithDefaults instantiates a new BriefL2VPNRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *BriefL2VPNRequest) GetIdentifier() int64`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BriefL2VPNRequest) GetIdentifierOk() (*int64, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BriefL2VPNRequest) SetIdentifier(v int64)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BriefL2VPNRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *BriefL2VPNRequest) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *BriefL2VPNRequest) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *BriefL2VPNRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefL2VPNRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefL2VPNRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefL2VPNRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefL2VPNRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefL2VPNRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetType

`func (o *BriefL2VPNRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BriefL2VPNRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BriefL2VPNRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BriefL2VPNRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *BriefL2VPNRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefL2VPNRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefL2VPNRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefL2VPNRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


