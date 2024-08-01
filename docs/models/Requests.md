# Requests

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]Request**](Request.md) | Array of items in the collection. | [optional] [readonly] |
|**Offset** | **float32** | The offset, specified in the request (if not is specified, 0 is used by default). | |
|**Limit** | **float32** | The limit, specified in the request (if not specified, the endpoint&#39;s default pagination limit is used). | |
|**Links** | [**PaginationLinks**](PaginationLinks.md) |  | |

## Methods

### NewRequests

`func NewRequests(offset float32, limit float32, links PaginationLinks, ) *Requests`

NewRequests instantiates a new Requests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestsWithDefaults

`func NewRequestsWithDefaults() *Requests`

NewRequestsWithDefaults instantiates a new Requests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Requests) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Requests) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Requests) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Requests) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Requests) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Requests) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Requests) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Requests) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Requests) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Requests) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Requests) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Requests) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Requests) GetItems() []Request`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Requests) GetItemsOk() (*[]Request, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Requests) SetItems(v []Request)`

SetItems sets Items field to given value.

### HasItems

`func (o *Requests) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *Requests) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Requests) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Requests) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *Requests) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Requests) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Requests) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetLinks

`func (o *Requests) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Requests) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Requests) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.




