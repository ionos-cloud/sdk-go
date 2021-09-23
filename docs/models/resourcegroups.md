# ResourceGroups

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Id** | Pointer to **string** | The resource's unique identifier | \[optional\] \[readonly\] |
| **Type** | Pointer to [**Type**](type.md) | The type of the resource | \[optional\] |
| **Href** | Pointer to **string** | URL to the object representation \(absolute path\) | \[optional\] \[readonly\] |
| **Items** | Pointer to [**\[\]Resource**](resource.md) | Array of items in that collection | \[optional\] \[readonly\] |

## Methods

### NewResourceGroups

`func NewResourceGroups() *ResourceGroups`

NewResourceGroups instantiates a new ResourceGroups object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewResourceGroupsWithDefaults

`func NewResourceGroupsWithDefaults() *ResourceGroups`

NewResourceGroupsWithDefaults instantiates a new ResourceGroups object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceGroups) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceGroups) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetId

`func (o *ResourceGroups) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ResourceGroups) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceGroups) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetType

`func (o *ResourceGroups) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceGroups) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ResourceGroups) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ResourceGroups) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetHref

`func (o *ResourceGroups) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ResourceGroups) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *ResourceGroups) GetItems() []Resource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResourceGroups) GetItemsOk() (*[]Resource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetItems

`func (o *ResourceGroups) SetItems(v []Resource)`

SetItems sets Items field to given value.

### HasItems

`func (o *ResourceGroups) HasItems() bool`

HasItems returns a boolean if a field has been set.

