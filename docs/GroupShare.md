# GroupShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | resource as generic type | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Properties** | [**GroupShareProperties**](GroupShareProperties.md) |  | 

## Methods

### NewGroupShare

`func NewGroupShare(properties GroupShareProperties, ) *GroupShare`

NewGroupShare instantiates a new GroupShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupShareWithDefaults

`func NewGroupShareWithDefaults() *GroupShare`

NewGroupShareWithDefaults instantiates a new GroupShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupShare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GroupShare) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupShare) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupShare) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *GroupShare) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *GroupShare) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GroupShare) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GroupShare) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GroupShare) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetProperties

`func (o *GroupShare) GetProperties() GroupShareProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GroupShare) GetPropertiesOk() (*GroupShareProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GroupShare) SetProperties(v GroupShareProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


