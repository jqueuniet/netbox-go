# ConfigTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentParams** | Pointer to **interface{}** | Any &lt;a href&#x3D;\&quot;https://jinja.palletsprojects.com/en/stable/api/#jinja2.Environment\&quot;&gt;additional parameters&lt;/a&gt; to pass when constructing the Jinja environment | [optional] 
**TemplateCode** | **string** | Jinja template code. | 
**MimeType** | Pointer to **string** | Defaults to &lt;code&gt;text/plain; charset&#x3D;utf-8&lt;/code&gt; | [optional] 
**FileName** | Pointer to **string** | Filename to give to the rendered export file | [optional] 
**FileExtension** | Pointer to **string** | Extension to append to the rendered filename | [optional] 
**AsAttachment** | Pointer to **bool** | Download file as attachment | [optional] 
**DataSource** | Pointer to [**ConfigContextProfileRequestDataSource**](ConfigContextProfileRequestDataSource.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewConfigTemplateRequest

`func NewConfigTemplateRequest(name string, templateCode string, ) *ConfigTemplateRequest`

NewConfigTemplateRequest instantiates a new ConfigTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigTemplateRequestWithDefaults

`func NewConfigTemplateRequestWithDefaults() *ConfigTemplateRequest`

NewConfigTemplateRequestWithDefaults instantiates a new ConfigTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConfigTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentParams

`func (o *ConfigTemplateRequest) GetEnvironmentParams() interface{}`

GetEnvironmentParams returns the EnvironmentParams field if non-nil, zero value otherwise.

### GetEnvironmentParamsOk

`func (o *ConfigTemplateRequest) GetEnvironmentParamsOk() (*interface{}, bool)`

GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentParams

`func (o *ConfigTemplateRequest) SetEnvironmentParams(v interface{})`

SetEnvironmentParams sets EnvironmentParams field to given value.

### HasEnvironmentParams

`func (o *ConfigTemplateRequest) HasEnvironmentParams() bool`

HasEnvironmentParams returns a boolean if a field has been set.

### SetEnvironmentParamsNil

`func (o *ConfigTemplateRequest) SetEnvironmentParamsNil(b bool)`

 SetEnvironmentParamsNil sets the value for EnvironmentParams to be an explicit nil

### UnsetEnvironmentParams
`func (o *ConfigTemplateRequest) UnsetEnvironmentParams()`

UnsetEnvironmentParams ensures that no value is present for EnvironmentParams, not even an explicit nil
### GetTemplateCode

`func (o *ConfigTemplateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *ConfigTemplateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *ConfigTemplateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetMimeType

`func (o *ConfigTemplateRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ConfigTemplateRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ConfigTemplateRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ConfigTemplateRequest) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileName

`func (o *ConfigTemplateRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ConfigTemplateRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ConfigTemplateRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ConfigTemplateRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileExtension

`func (o *ConfigTemplateRequest) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *ConfigTemplateRequest) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *ConfigTemplateRequest) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *ConfigTemplateRequest) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetAsAttachment

`func (o *ConfigTemplateRequest) GetAsAttachment() bool`

GetAsAttachment returns the AsAttachment field if non-nil, zero value otherwise.

### GetAsAttachmentOk

`func (o *ConfigTemplateRequest) GetAsAttachmentOk() (*bool, bool)`

GetAsAttachmentOk returns a tuple with the AsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAttachment

`func (o *ConfigTemplateRequest) SetAsAttachment(v bool)`

SetAsAttachment sets AsAttachment field to given value.

### HasAsAttachment

`func (o *ConfigTemplateRequest) HasAsAttachment() bool`

HasAsAttachment returns a boolean if a field has been set.

### GetDataSource

`func (o *ConfigTemplateRequest) GetDataSource() ConfigContextProfileRequestDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ConfigTemplateRequest) GetDataSourceOk() (*ConfigContextProfileRequestDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ConfigTemplateRequest) SetDataSource(v ConfigContextProfileRequestDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ConfigTemplateRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetTags

`func (o *ConfigTemplateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigTemplateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigTemplateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


