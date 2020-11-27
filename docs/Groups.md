# Groups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of the resource | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]Group**](Group.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewGroups

`func NewGroups() *Groups`

NewGroups instantiates a new Groups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsWithDefaults

`func NewGroupsWithDefaults() *Groups`

NewGroupsWithDefaults instantiates a new Groups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Groups) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Groups) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Groups) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Groups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Groups) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Groups) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Groups) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Groups) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Groups) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Groups) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Groups) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Groups) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Groups) GetItems() []Group`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Groups) GetItemsOk() (*[]Group, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Groups) SetItems(v []Group)`

SetItems sets Items field to given value.

### HasItems

`func (o *Groups) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


