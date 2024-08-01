# LabelResources

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | A unique representation of the label as a resource collection. | [optional] [readonly] |
|**Type** | Pointer to **string** | The type of resource within a collection. | [optional] [readonly] |
|**Href** | Pointer to **string** | URL to the collection representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]LabelResource**](LabelResource.md) | Array of items in the collection. | [optional] [readonly] |
|**Offset** | Pointer to **float32** | The offset (if specified in the request). | [optional] |
|**Limit** | Pointer to **float32** | The limit (if specified in the request). | [optional] |
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |

## Methods

### NewLabelResources

`func NewLabelResources() *LabelResources`

NewLabelResources instantiates a new LabelResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelResourcesWithDefaults

`func NewLabelResourcesWithDefaults() *LabelResources`

NewLabelResourcesWithDefaults instantiates a new LabelResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LabelResources) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelResources) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelResources) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LabelResources) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LabelResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabelResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabelResources) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LabelResources) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *LabelResources) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LabelResources) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LabelResources) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LabelResources) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *LabelResources) GetItems() []LabelResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *LabelResources) GetItemsOk() (*[]LabelResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *LabelResources) SetItems(v []LabelResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *LabelResources) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *LabelResources) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *LabelResources) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *LabelResources) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *LabelResources) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *LabelResources) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LabelResources) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LabelResources) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LabelResources) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *LabelResources) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LabelResources) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LabelResources) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LabelResources) HasLinks() bool`

HasLinks returns a boolean if a field has been set.



