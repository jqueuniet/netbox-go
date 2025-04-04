# IPAddressRequestVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewIPAddressRequestVrf

`func NewIPAddressRequestVrf(name string, ) *IPAddressRequestVrf`

NewIPAddressRequestVrf instantiates a new IPAddressRequestVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressRequestVrfWithDefaults

`func NewIPAddressRequestVrfWithDefaults() *IPAddressRequestVrf`

NewIPAddressRequestVrfWithDefaults instantiates a new IPAddressRequestVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IPAddressRequestVrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPAddressRequestVrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPAddressRequestVrf) SetName(v string)`

SetName sets Name field to given value.


### GetRd

`func (o *IPAddressRequestVrf) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *IPAddressRequestVrf) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *IPAddressRequestVrf) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *IPAddressRequestVrf) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *IPAddressRequestVrf) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *IPAddressRequestVrf) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetDescription

`func (o *IPAddressRequestVrf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPAddressRequestVrf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPAddressRequestVrf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPAddressRequestVrf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


