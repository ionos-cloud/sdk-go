# S3Keys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of the resource | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]S3Key**](S3Key.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewS3Keys

`func NewS3Keys() *S3Keys`

NewS3Keys instantiates a new S3Keys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3KeysWithDefaults

`func NewS3KeysWithDefaults() *S3Keys`

NewS3KeysWithDefaults instantiates a new S3Keys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3Keys) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3Keys) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3Keys) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *S3Keys) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *S3Keys) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3Keys) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3Keys) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *S3Keys) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *S3Keys) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *S3Keys) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *S3Keys) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *S3Keys) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *S3Keys) GetItems() []S3Key`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *S3Keys) GetItemsOk() (*[]S3Key, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *S3Keys) SetItems(v []S3Key)`

SetItems sets Items field to given value.

### HasItems

`func (o *S3Keys) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


