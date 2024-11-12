# IKEProposalGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** | * &#x60;1&#x60; - Group 1 * &#x60;2&#x60; - Group 2 * &#x60;5&#x60; - Group 5 * &#x60;14&#x60; - Group 14 * &#x60;15&#x60; - Group 15 * &#x60;16&#x60; - Group 16 * &#x60;17&#x60; - Group 17 * &#x60;18&#x60; - Group 18 * &#x60;19&#x60; - Group 19 * &#x60;20&#x60; - Group 20 * &#x60;21&#x60; - Group 21 * &#x60;22&#x60; - Group 22 * &#x60;23&#x60; - Group 23 * &#x60;24&#x60; - Group 24 * &#x60;25&#x60; - Group 25 * &#x60;26&#x60; - Group 26 * &#x60;27&#x60; - Group 27 * &#x60;28&#x60; - Group 28 * &#x60;29&#x60; - Group 29 * &#x60;30&#x60; - Group 30 * &#x60;31&#x60; - Group 31 * &#x60;32&#x60; - Group 32 * &#x60;33&#x60; - Group 33 * &#x60;34&#x60; - Group 34 | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewIKEProposalGroup

`func NewIKEProposalGroup() *IKEProposalGroup`

NewIKEProposalGroup instantiates a new IKEProposalGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIKEProposalGroupWithDefaults

`func NewIKEProposalGroupWithDefaults() *IKEProposalGroup`

NewIKEProposalGroupWithDefaults instantiates a new IKEProposalGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *IKEProposalGroup) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IKEProposalGroup) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IKEProposalGroup) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *IKEProposalGroup) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *IKEProposalGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IKEProposalGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IKEProposalGroup) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IKEProposalGroup) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


