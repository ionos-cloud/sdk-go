# AttachedVolumes

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Id** | Pointer to **string** | The resource's unique identifier | \[optional\] \[readonly\] |
| **Type** | Pointer to [**Type**](type.md) | The type of object that has been created | \[optional\] |
| **Href** | Pointer to **string** | URL to the object representation \(absolute path\) | \[optional\] \[readonly\] |
| **Items** | Pointer to [**\[\]Volume**](volume.md) | Array of items in that collection | \[optional\] \[readonly\] |
| **Offset** | Pointer to **float32** | the offset \(if specified in the request\) | \[optional\] |
| **Limit** | Pointer to **float32** | the limit \(if specified in the request\) | \[optional\] |
| **Links** | Pointer to [**PaginationLinks**](paginationlinks.md) |  | \[optional\] |

## Methods

### NewAttachedVolumes

`func NewAttachedVolumes() *AttachedVolumes`

NewAttachedVolumes instantiates a new AttachedVolumes object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewAttachedVolumesWithDefaults

`func NewAttachedVolumesWithDefaults() *AttachedVolumes`

NewAttachedVolumesWithDefaults instantiates a new AttachedVolumes object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttachedVolumes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttachedVolumes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetId

`func (o *AttachedVolumes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttachedVolumes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AttachedVolumes) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachedVolumes) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetType

`func (o *AttachedVolumes) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *AttachedVolumes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *AttachedVolumes) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AttachedVolumes) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetHref

`func (o *AttachedVolumes) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AttachedVolumes) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *AttachedVolumes) GetItems() []Volume`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AttachedVolumes) GetItemsOk() (*[]Volume, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetItems

`func (o *AttachedVolumes) SetItems(v []Volume)`

SetItems sets Items field to given value.

### HasItems

`func (o *AttachedVolumes) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *AttachedVolumes) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AttachedVolumes) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetOffset

`func (o *AttachedVolumes) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AttachedVolumes) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *AttachedVolumes) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AttachedVolumes) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetLimit

`func (o *AttachedVolumes) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AttachedVolumes) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *AttachedVolumes) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AttachedVolumes) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetLinks

`func (o *AttachedVolumes) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AttachedVolumes) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

