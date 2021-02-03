# GroupUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of the resource | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]Group**](Group.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewGroupUsers

`func NewGroupUsers() *GroupUsers`

NewGroupUsers instantiates a new GroupUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUsersWithDefaults

`func NewGroupUsersWithDefaults() *GroupUsers`

NewGroupUsersWithDefaults instantiates a new GroupUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupUsers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupUsers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupUsers) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupUsers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GroupUsers) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupUsers) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupUsers) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *GroupUsers) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *GroupUsers) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GroupUsers) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GroupUsers) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GroupUsers) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *GroupUsers) GetItems() []Group`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GroupUsers) GetItemsOk() (*[]Group, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GroupUsers) SetItems(v []Group)`

SetItems sets Items field to given value.

### HasItems

`func (o *GroupUsers) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


