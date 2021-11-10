# PaginationLinks

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Prev** | Pointer to **string** | URL (with offset and limit parameters) of the previous page; only present if offset is greater than 0. | [optional] [readonly] |
|**Self** | Pointer to **string** | URL (with offset and limit parameters) of the current page. | [optional] [readonly] |
|**Next** | Pointer to **string** | URL (with offset and limit parameters) of the next page; only present if offset + limit is less than the total number of elements. | [optional] [readonly] |

## Methods

### NewPaginationLinks

`func NewPaginationLinks() *PaginationLinks`

NewPaginationLinks instantiates a new PaginationLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationLinksWithDefaults

`func NewPaginationLinksWithDefaults() *PaginationLinks`

NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrev

`func (o *PaginationLinks) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PaginationLinks) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PaginationLinks) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *PaginationLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSelf

`func (o *PaginationLinks) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PaginationLinks) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PaginationLinks) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PaginationLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetNext

`func (o *PaginationLinks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginationLinks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginationLinks) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginationLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.



