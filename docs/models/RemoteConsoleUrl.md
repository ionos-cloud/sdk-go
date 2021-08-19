# RemoteConsoleUrl

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Url** | Pointer to **string** | The remote console url with the jwToken parameter for access | [optional] [readonly] |

## Methods

### NewRemoteConsoleUrl

`func NewRemoteConsoleUrl() *RemoteConsoleUrl`

NewRemoteConsoleUrl instantiates a new RemoteConsoleUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConsoleUrlWithDefaults

`func NewRemoteConsoleUrlWithDefaults() *RemoteConsoleUrl`

NewRemoteConsoleUrlWithDefaults instantiates a new RemoteConsoleUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RemoteConsoleUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RemoteConsoleUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RemoteConsoleUrl) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RemoteConsoleUrl) HasUrl() bool`

HasUrl returns a boolean if a field has been set.



