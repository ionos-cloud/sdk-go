# Users

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]User**](User.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewUsers

`func NewUsers() *Users`

NewUsers instantiates a new Users object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersWithDefaults

`func NewUsersWithDefaults() *Users`

NewUsersWithDefaults instantiates a new Users object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Users) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Users) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Users) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Users) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Users) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Users) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Users) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Users) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Users) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Users) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Users) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Users) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Users) GetItems() []User`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Users) GetItemsOk() (*[]User, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Users) SetItems(v []User)`

SetItems sets Items field to given value.

### HasItems

`func (o *Users) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


