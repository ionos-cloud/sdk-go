# Loadbalancers

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]Loadbalancer**](Loadbalancer.md) | Array of items in the collection. | [optional] [readonly] |
|**Offset** | Pointer to **float32** | The offset (if specified in the request). | [optional] |
|**Limit** | Pointer to **float32** | The limit (if specified in the request). | [optional] |
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |

## Methods

### NewLoadbalancers

`func NewLoadbalancers() *Loadbalancers`

NewLoadbalancers instantiates a new Loadbalancers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancersWithDefaults

`func NewLoadbalancersWithDefaults() *Loadbalancers`

NewLoadbalancersWithDefaults instantiates a new Loadbalancers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Loadbalancers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Loadbalancers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Loadbalancers) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Loadbalancers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Loadbalancers) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Loadbalancers) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Loadbalancers) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Loadbalancers) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Loadbalancers) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Loadbalancers) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Loadbalancers) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Loadbalancers) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Loadbalancers) GetItems() []Loadbalancer`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Loadbalancers) GetItemsOk() (*[]Loadbalancer, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Loadbalancers) SetItems(v []Loadbalancer)`

SetItems sets Items field to given value.

### HasItems

`func (o *Loadbalancers) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *Loadbalancers) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Loadbalancers) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Loadbalancers) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Loadbalancers) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *Loadbalancers) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Loadbalancers) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Loadbalancers) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Loadbalancers) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *Loadbalancers) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Loadbalancers) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Loadbalancers) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Loadbalancers) HasLinks() bool`

HasLinks returns a boolean if a field has been set.



