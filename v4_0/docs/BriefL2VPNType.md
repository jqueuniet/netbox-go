# BriefL2VPNType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;vpws&#x60; - VPWS * &#x60;vpls&#x60; - VPLS * &#x60;vxlan&#x60; - VXLAN * &#x60;vxlan-evpn&#x60; - VXLAN-EVPN * &#x60;mpls-evpn&#x60; - MPLS EVPN * &#x60;pbb-evpn&#x60; - PBB EVPN * &#x60;epl&#x60; - EPL * &#x60;evpl&#x60; - EVPL * &#x60;ep-lan&#x60; - Ethernet Private LAN * &#x60;evp-lan&#x60; - Ethernet Virtual Private LAN * &#x60;ep-tree&#x60; - Ethernet Private Tree * &#x60;evp-tree&#x60; - Ethernet Virtual Private Tree | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefL2VPNType

`func NewBriefL2VPNType() *BriefL2VPNType`

NewBriefL2VPNType instantiates a new BriefL2VPNType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefL2VPNTypeWithDefaults

`func NewBriefL2VPNTypeWithDefaults() *BriefL2VPNType`

NewBriefL2VPNTypeWithDefaults instantiates a new BriefL2VPNType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *BriefL2VPNType) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BriefL2VPNType) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BriefL2VPNType) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BriefL2VPNType) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *BriefL2VPNType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BriefL2VPNType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BriefL2VPNType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BriefL2VPNType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


