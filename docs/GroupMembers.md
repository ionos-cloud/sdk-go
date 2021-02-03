# GroupMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]User**](User.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewGroupMembers

`func NewGroupMembers() *GroupMembers`

NewGroupMembers instantiates a new GroupMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMembersWithDefaults

`func NewGroupMembersWithDefaults() *GroupMembers`

NewGroupMembersWithDefaults instantiates a new GroupMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupMembers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupMembers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupMembers) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupMembers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GroupMembers) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupMembers) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupMembers) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *GroupMembers) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *GroupMembers) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GroupMembers) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GroupMembers) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GroupMembers) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *GroupMembers) GetItems() []User`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GroupMembers) GetItemsOk() (*[]User, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GroupMembers) SetItems(v []User)`

SetItems sets Items field to given value.

### HasItems

`func (o *GroupMembers) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


