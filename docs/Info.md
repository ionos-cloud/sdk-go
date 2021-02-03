# Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | API entry point | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the API | [optional] [readonly] 
**Version** | Pointer to **string** | Version of the API | [optional] [readonly] 

## Methods

### NewInfo

`func NewInfo() *Info`

NewInfo instantiates a new Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoWithDefaults

`func NewInfoWithDefaults() *Info`

NewInfoWithDefaults instantiates a new Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Info) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Info) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Info) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Info) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetName

`func (o *Info) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Info) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Info) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Info) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Info) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Info) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Info) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Info) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


