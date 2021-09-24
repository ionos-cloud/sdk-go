# ResourcesUsers

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of the resource | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] |
|**Items** | Pointer to [**[]Resource**](Resource.md) | Array of items in that collection | [optional] [readonly] |

## Methods

### NewResourcesUsers

`func NewResourcesUsers() *ResourcesUsers`

NewResourcesUsers instantiates a new ResourcesUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesUsersWithDefaults

`func NewResourcesUsersWithDefaults() *ResourcesUsers`

NewResourcesUsersWithDefaults instantiates a new ResourcesUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourcesUsers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcesUsers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcesUsers) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourcesUsers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ResourcesUsers) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourcesUsers) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourcesUsers) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *ResourcesUsers) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ResourcesUsers) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ResourcesUsers) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ResourcesUsers) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ResourcesUsers) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *ResourcesUsers) GetItems() []Resource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResourcesUsers) GetItemsOk() (*[]Resource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResourcesUsers) SetItems(v []Resource)`

SetItems sets Items field to given value.

### HasItems

`func (o *ResourcesUsers) HasItems() bool`

HasItems returns a boolean if a field has been set.



