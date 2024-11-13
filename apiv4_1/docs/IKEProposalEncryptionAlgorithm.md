# IKEProposalEncryptionAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;aes-128-cbc&#x60; - 128-bit AES (CBC) * &#x60;aes-128-gcm&#x60; - 128-bit AES (GCM) * &#x60;aes-192-cbc&#x60; - 192-bit AES (CBC) * &#x60;aes-192-gcm&#x60; - 192-bit AES (GCM) * &#x60;aes-256-cbc&#x60; - 256-bit AES (CBC) * &#x60;aes-256-gcm&#x60; - 256-bit AES (GCM) * &#x60;3des-cbc&#x60; - 3DES * &#x60;des-cbc&#x60; - DES | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewIKEProposalEncryptionAlgorithm

`func NewIKEProposalEncryptionAlgorithm() *IKEProposalEncryptionAlgorithm`

NewIKEProposalEncryptionAlgorithm instantiates a new IKEProposalEncryptionAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIKEProposalEncryptionAlgorithmWithDefaults

`func NewIKEProposalEncryptionAlgorithmWithDefaults() *IKEProposalEncryptionAlgorithm`

NewIKEProposalEncryptionAlgorithmWithDefaults instantiates a new IKEProposalEncryptionAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *IKEProposalEncryptionAlgorithm) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IKEProposalEncryptionAlgorithm) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IKEProposalEncryptionAlgorithm) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IKEProposalEncryptionAlgorithm) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *IKEProposalEncryptionAlgorithm) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IKEProposalEncryptionAlgorithm) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IKEProposalEncryptionAlgorithm) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IKEProposalEncryptionAlgorithm) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


