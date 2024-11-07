# SecurityGroups

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | The URL to the object representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]SecurityGroup**](SecurityGroup.md) | Array of items in the collection. | [optional] [readonly] |
|**Offset** | Pointer to **float32** | The offset (if specified in the request). | [optional] |
|**Limit** | Pointer to **float32** | The limit (if specified in the request). | [optional] |
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |

## Methods

### NewSecurityGroups

`func NewSecurityGroups() *SecurityGroups`

NewSecurityGroups instantiates a new SecurityGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupsWithDefaults

`func NewSecurityGroupsWithDefaults() *SecurityGroups`

NewSecurityGroupsWithDefaults instantiates a new SecurityGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroups) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroups) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroups) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SecurityGroups) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityGroups) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityGroups) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityGroups) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *SecurityGroups) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SecurityGroups) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SecurityGroups) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SecurityGroups) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *SecurityGroups) GetItems() []SecurityGroup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecurityGroups) GetItemsOk() (*[]SecurityGroup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecurityGroups) SetItems(v []SecurityGroup)`

SetItems sets Items field to given value.

### HasItems

`func (o *SecurityGroups) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *SecurityGroups) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SecurityGroups) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SecurityGroups) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SecurityGroups) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SecurityGroups) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SecurityGroups) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SecurityGroups) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SecurityGroups) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *SecurityGroups) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityGroups) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityGroups) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityGroups) HasLinks() bool`

HasLinks returns a boolean if a field has been set.



