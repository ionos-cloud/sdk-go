# RequestProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestProperties

`func NewRequestProperties() *RequestProperties`

NewRequestProperties instantiates a new RequestProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPropertiesWithDefaults

`func NewRequestPropertiesWithDefaults() *RequestProperties`

NewRequestPropertiesWithDefaults instantiates a new RequestProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *RequestProperties) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RequestProperties) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RequestProperties) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RequestProperties) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetHeaders

`func (o *RequestProperties) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RequestProperties) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RequestProperties) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RequestProperties) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *RequestProperties) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RequestProperties) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RequestProperties) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *RequestProperties) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetUrl

`func (o *RequestProperties) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestProperties) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestProperties) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RequestProperties) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


