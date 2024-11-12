# TunnelStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;planned&#x60; - Planned * &#x60;active&#x60; - Active * &#x60;disabled&#x60; - Disabled | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewTunnelStatus

`func NewTunnelStatus() *TunnelStatus`

NewTunnelStatus instantiates a new TunnelStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelStatusWithDefaults

`func NewTunnelStatusWithDefaults() *TunnelStatus`

NewTunnelStatusWithDefaults instantiates a new TunnelStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *TunnelStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TunnelStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TunnelStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TunnelStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *TunnelStatus) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TunnelStatus) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TunnelStatus) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TunnelStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


