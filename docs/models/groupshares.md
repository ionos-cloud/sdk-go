# GroupShares

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Id** | Pointer to **string** | The resource's unique identifier | \[optional\] \[readonly\] |
| **Type** | Pointer to [**Type**](type.md) | Share representing groups and resource relationship | \[optional\] |
| **Href** | Pointer to **string** | URL to the object representation \(absolute path\) | \[optional\] \[readonly\] |
| **Items** | Pointer to [**\[\]GroupShare**](groupshare.md) | Array of items in that collection | \[optional\] \[readonly\] |

## Methods

### NewGroupShares

`func NewGroupShares() *GroupShares`

NewGroupShares instantiates a new GroupShares object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewGroupSharesWithDefaults

`func NewGroupSharesWithDefaults() *GroupShares`

NewGroupSharesWithDefaults instantiates a new GroupShares object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupShares) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupShares) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetId

`func (o *GroupShares) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupShares) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GroupShares) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupShares) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetType

`func (o *GroupShares) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *GroupShares) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *GroupShares) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GroupShares) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetHref

`func (o *GroupShares) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GroupShares) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *GroupShares) GetItems() []GroupShare`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GroupShares) GetItemsOk() (*[]GroupShare, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetItems

`func (o *GroupShares) SetItems(v []GroupShare)`

SetItems sets Items field to given value.

### HasItems

`func (o *GroupShares) HasItems() bool`

HasItems returns a boolean if a field has been set.

