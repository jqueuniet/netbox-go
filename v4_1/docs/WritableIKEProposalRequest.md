# WritableIKEProposalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**AuthenticationMethod** | **string** | * &#x60;preshared-keys&#x60; - Pre-shared keys * &#x60;certificates&#x60; - Certificates * &#x60;rsa-signatures&#x60; - RSA signatures * &#x60;dsa-signatures&#x60; - DSA signatures | 
**EncryptionAlgorithm** | **string** | * &#x60;aes-128-cbc&#x60; - 128-bit AES (CBC) * &#x60;aes-128-gcm&#x60; - 128-bit AES (GCM) * &#x60;aes-192-cbc&#x60; - 192-bit AES (CBC) * &#x60;aes-192-gcm&#x60; - 192-bit AES (GCM) * &#x60;aes-256-cbc&#x60; - 256-bit AES (CBC) * &#x60;aes-256-gcm&#x60; - 256-bit AES (GCM) * &#x60;3des-cbc&#x60; - 3DES * &#x60;des-cbc&#x60; - DES | 
**AuthenticationAlgorithm** | Pointer to **string** | * &#x60;hmac-sha1&#x60; - SHA-1 HMAC * &#x60;hmac-sha256&#x60; - SHA-256 HMAC * &#x60;hmac-sha384&#x60; - SHA-384 HMAC * &#x60;hmac-sha512&#x60; - SHA-512 HMAC * &#x60;hmac-md5&#x60; - MD5 HMAC | [optional] 
**Group** | **int32** | Diffie-Hellman group ID  * &#x60;1&#x60; - Group 1 * &#x60;2&#x60; - Group 2 * &#x60;5&#x60; - Group 5 * &#x60;14&#x60; - Group 14 * &#x60;15&#x60; - Group 15 * &#x60;16&#x60; - Group 16 * &#x60;17&#x60; - Group 17 * &#x60;18&#x60; - Group 18 * &#x60;19&#x60; - Group 19 * &#x60;20&#x60; - Group 20 * &#x60;21&#x60; - Group 21 * &#x60;22&#x60; - Group 22 * &#x60;23&#x60; - Group 23 * &#x60;24&#x60; - Group 24 * &#x60;25&#x60; - Group 25 * &#x60;26&#x60; - Group 26 * &#x60;27&#x60; - Group 27 * &#x60;28&#x60; - Group 28 * &#x60;29&#x60; - Group 29 * &#x60;30&#x60; - Group 30 * &#x60;31&#x60; - Group 31 * &#x60;32&#x60; - Group 32 * &#x60;33&#x60; - Group 33 * &#x60;34&#x60; - Group 34 | 
**SaLifetime** | Pointer to **NullableInt32** | Security association lifetime (in seconds) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableIKEProposalRequest

`func NewWritableIKEProposalRequest(name string, authenticationMethod string, encryptionAlgorithm string, group int32, ) *WritableIKEProposalRequest`

NewWritableIKEProposalRequest instantiates a new WritableIKEProposalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableIKEProposalRequestWithDefaults

`func NewWritableIKEProposalRequestWithDefaults() *WritableIKEProposalRequest`

NewWritableIKEProposalRequestWithDefaults instantiates a new WritableIKEProposalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableIKEProposalRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableIKEProposalRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableIKEProposalRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WritableIKEProposalRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableIKEProposalRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableIKEProposalRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableIKEProposalRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *WritableIKEProposalRequest) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *WritableIKEProposalRequest) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *WritableIKEProposalRequest) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetEncryptionAlgorithm

`func (o *WritableIKEProposalRequest) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *WritableIKEProposalRequest) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *WritableIKEProposalRequest) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.


### GetAuthenticationAlgorithm

`func (o *WritableIKEProposalRequest) GetAuthenticationAlgorithm() string`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *WritableIKEProposalRequest) GetAuthenticationAlgorithmOk() (*string, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *WritableIKEProposalRequest) SetAuthenticationAlgorithm(v string)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *WritableIKEProposalRequest) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### GetGroup

`func (o *WritableIKEProposalRequest) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableIKEProposalRequest) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableIKEProposalRequest) SetGroup(v int32)`

SetGroup sets Group field to given value.


### GetSaLifetime

`func (o *WritableIKEProposalRequest) GetSaLifetime() int32`

GetSaLifetime returns the SaLifetime field if non-nil, zero value otherwise.

### GetSaLifetimeOk

`func (o *WritableIKEProposalRequest) GetSaLifetimeOk() (*int32, bool)`

GetSaLifetimeOk returns a tuple with the SaLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetime

`func (o *WritableIKEProposalRequest) SetSaLifetime(v int32)`

SetSaLifetime sets SaLifetime field to given value.

### HasSaLifetime

`func (o *WritableIKEProposalRequest) HasSaLifetime() bool`

HasSaLifetime returns a boolean if a field has been set.

### SetSaLifetimeNil

`func (o *WritableIKEProposalRequest) SetSaLifetimeNil(b bool)`

 SetSaLifetimeNil sets the value for SaLifetime to be an explicit nil

### UnsetSaLifetime
`func (o *WritableIKEProposalRequest) UnsetSaLifetime()`

UnsetSaLifetime ensures that no value is present for SaLifetime, not even an explicit nil
### GetComments

`func (o *WritableIKEProposalRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableIKEProposalRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableIKEProposalRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableIKEProposalRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableIKEProposalRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableIKEProposalRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableIKEProposalRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableIKEProposalRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableIKEProposalRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableIKEProposalRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableIKEProposalRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableIKEProposalRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


