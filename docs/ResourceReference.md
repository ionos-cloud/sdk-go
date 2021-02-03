# ResourceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The resource&#39;s unique identifier | 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 

## Methods

### NewResourceReference

`func NewResourceReference(id string, ) *ResourceReference`

NewResourceReference instantiates a new ResourceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceReferenceWithDefaults

`func NewResourceReferenceWithDefaults() *ResourceReference`

NewResourceReferenceWithDefaults instantiates a new ResourceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceReference) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ResourceReference) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceReference) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceReference) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ResourceReference) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ResourceReference) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ResourceReference) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ResourceReference) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


