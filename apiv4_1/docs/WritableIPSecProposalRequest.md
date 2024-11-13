# WritableIPSecProposalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EncryptionAlgorithm** | Pointer to **string** | * &#x60;aes-128-cbc&#x60; - 128-bit AES (CBC) * &#x60;aes-128-gcm&#x60; - 128-bit AES (GCM) * &#x60;aes-192-cbc&#x60; - 192-bit AES (CBC) * &#x60;aes-192-gcm&#x60; - 192-bit AES (GCM) * &#x60;aes-256-cbc&#x60; - 256-bit AES (CBC) * &#x60;aes-256-gcm&#x60; - 256-bit AES (GCM) * &#x60;3des-cbc&#x60; - 3DES * &#x60;des-cbc&#x60; - DES | [optional] 
**AuthenticationAlgorithm** | Pointer to **string** | * &#x60;hmac-sha1&#x60; - SHA-1 HMAC * &#x60;hmac-sha256&#x60; - SHA-256 HMAC * &#x60;hmac-sha384&#x60; - SHA-384 HMAC * &#x60;hmac-sha512&#x60; - SHA-512 HMAC * &#x60;hmac-md5&#x60; - MD5 HMAC | [optional] 
**SaLifetimeSeconds** | Pointer to **NullableInt32** | Security association lifetime (seconds) | [optional] 
**SaLifetimeData** | Pointer to **NullableInt32** | Security association lifetime (in kilobytes) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableIPSecProposalRequest

`func NewWritableIPSecProposalRequest(name string, ) *WritableIPSecProposalRequest`

NewWritableIPSecProposalRequest instantiates a new WritableIPSecProposalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableIPSecProposalRequestWithDefaults

`func NewWritableIPSecProposalRequestWithDefaults() *WritableIPSecProposalRequest`

NewWritableIPSecProposalRequestWithDefaults instantiates a new WritableIPSecProposalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableIPSecProposalRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableIPSecProposalRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableIPSecProposalRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WritableIPSecProposalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableIPSecProposalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableIPSecProposalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableIPSecProposalRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *WritableIPSecProposalRequest) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *WritableIPSecProposalRequest) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *WritableIPSecProposalRequest) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *WritableIPSecProposalRequest) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetAuthenticationAlgorithm

`func (o *WritableIPSecProposalRequest) GetAuthenticationAlgorithm() string`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *WritableIPSecProposalRequest) GetAuthenticationAlgorithmOk() (*string, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *WritableIPSecProposalRequest) SetAuthenticationAlgorithm(v string)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *WritableIPSecProposalRequest) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### GetSaLifetimeSeconds

`func (o *WritableIPSecProposalRequest) GetSaLifetimeSeconds() int32`

GetSaLifetimeSeconds returns the SaLifetimeSeconds field if non-nil, zero value otherwise.

### GetSaLifetimeSecondsOk

`func (o *WritableIPSecProposalRequest) GetSaLifetimeSecondsOk() (*int32, bool)`

GetSaLifetimeSecondsOk returns a tuple with the SaLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetimeSeconds

`func (o *WritableIPSecProposalRequest) SetSaLifetimeSeconds(v int32)`

SetSaLifetimeSeconds sets SaLifetimeSeconds field to given value.

### HasSaLifetimeSeconds

`func (o *WritableIPSecProposalRequest) HasSaLifetimeSeconds() bool`

HasSaLifetimeSeconds returns a boolean if a field has been set.

### SetSaLifetimeSecondsNil

`func (o *WritableIPSecProposalRequest) SetSaLifetimeSecondsNil(b bool)`

 SetSaLifetimeSecondsNil sets the value for SaLifetimeSeconds to be an explicit nil

### UnsetSaLifetimeSeconds
`func (o *WritableIPSecProposalRequest) UnsetSaLifetimeSeconds()`

UnsetSaLifetimeSeconds ensures that no value is present for SaLifetimeSeconds, not even an explicit nil
### GetSaLifetimeData

`func (o *WritableIPSecProposalRequest) GetSaLifetimeData() int32`

GetSaLifetimeData returns the SaLifetimeData field if non-nil, zero value otherwise.

### GetSaLifetimeDataOk

`func (o *WritableIPSecProposalRequest) GetSaLifetimeDataOk() (*int32, bool)`

GetSaLifetimeDataOk returns a tuple with the SaLifetimeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetimeData

`func (o *WritableIPSecProposalRequest) SetSaLifetimeData(v int32)`

SetSaLifetimeData sets SaLifetimeData field to given value.

### HasSaLifetimeData

`func (o *WritableIPSecProposalRequest) HasSaLifetimeData() bool`

HasSaLifetimeData returns a boolean if a field has been set.

### SetSaLifetimeDataNil

`func (o *WritableIPSecProposalRequest) SetSaLifetimeDataNil(b bool)`

 SetSaLifetimeDataNil sets the value for SaLifetimeData to be an explicit nil

### UnsetSaLifetimeData
`func (o *WritableIPSecProposalRequest) UnsetSaLifetimeData()`

UnsetSaLifetimeData ensures that no value is present for SaLifetimeData, not even an explicit nil
### GetComments

`func (o *WritableIPSecProposalRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableIPSecProposalRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableIPSecProposalRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableIPSecProposalRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableIPSecProposalRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableIPSecProposalRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableIPSecProposalRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableIPSecProposalRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableIPSecProposalRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableIPSecProposalRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableIPSecProposalRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableIPSecProposalRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


