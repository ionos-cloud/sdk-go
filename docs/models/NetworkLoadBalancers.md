# NetworkLoadBalancers

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]NetworkLoadBalancer**](NetworkLoadBalancer.md) | Array of items in the collection. | [optional] [readonly] |
|**Offset** | Pointer to **float32** | The offset (if specified in the request). | [optional] |
|**Limit** | Pointer to **float32** | The limit (if specified in the request). | [optional] |
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |

## Methods

### NewNetworkLoadBalancers

`func NewNetworkLoadBalancers() *NetworkLoadBalancers`

NewNetworkLoadBalancers instantiates a new NetworkLoadBalancers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoadBalancersWithDefaults

`func NewNetworkLoadBalancersWithDefaults() *NetworkLoadBalancers`

NewNetworkLoadBalancersWithDefaults instantiates a new NetworkLoadBalancers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkLoadBalancers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkLoadBalancers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkLoadBalancers) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkLoadBalancers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NetworkLoadBalancers) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkLoadBalancers) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkLoadBalancers) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkLoadBalancers) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *NetworkLoadBalancers) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NetworkLoadBalancers) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NetworkLoadBalancers) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NetworkLoadBalancers) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *NetworkLoadBalancers) GetItems() []NetworkLoadBalancer`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NetworkLoadBalancers) GetItemsOk() (*[]NetworkLoadBalancer, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NetworkLoadBalancers) SetItems(v []NetworkLoadBalancer)`

SetItems sets Items field to given value.

### HasItems

`func (o *NetworkLoadBalancers) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *NetworkLoadBalancers) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *NetworkLoadBalancers) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *NetworkLoadBalancers) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *NetworkLoadBalancers) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *NetworkLoadBalancers) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NetworkLoadBalancers) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NetworkLoadBalancers) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *NetworkLoadBalancers) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkLoadBalancers) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkLoadBalancers) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkLoadBalancers) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkLoadBalancers) HasLinks() bool`

HasLinks returns a boolean if a field has been set.



