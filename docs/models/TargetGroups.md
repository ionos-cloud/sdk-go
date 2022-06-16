# TargetGroups

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]TargetGroup**](TargetGroup.md) | Array of items in the collection. | [optional] [readonly] |
|**Offset** | Pointer to **float32** | The offset, specified in the request (if not is specified, 0 is used by default). | [optional] |
|**Limit** | Pointer to **float32** | The limit, specified in the request (if not specified, the endpoint&#39;s default pagination limit is used). | [optional] |
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |

## Methods

### NewTargetGroups

`func NewTargetGroups() *TargetGroups`

NewTargetGroups instantiates a new TargetGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupsWithDefaults

`func NewTargetGroupsWithDefaults() *TargetGroups`

NewTargetGroupsWithDefaults instantiates a new TargetGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TargetGroups) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetGroups) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetGroups) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TargetGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TargetGroups) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TargetGroups) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TargetGroups) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *TargetGroups) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *TargetGroups) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TargetGroups) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TargetGroups) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *TargetGroups) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *TargetGroups) GetItems() []TargetGroup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TargetGroups) GetItemsOk() (*[]TargetGroup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TargetGroups) SetItems(v []TargetGroup)`

SetItems sets Items field to given value.

### HasItems

`func (o *TargetGroups) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *TargetGroups) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TargetGroups) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TargetGroups) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TargetGroups) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *TargetGroups) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TargetGroups) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TargetGroups) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TargetGroups) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *TargetGroups) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TargetGroups) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TargetGroups) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TargetGroups) HasLinks() bool`

HasLinks returns a boolean if a field has been set.



