# Contracts

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]Contract**](Contract.md) | Array of items in the collection. | [optional] [readonly] |

## Methods

### NewContracts

`func NewContracts() *Contracts`

NewContracts instantiates a new Contracts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractsWithDefaults

`func NewContractsWithDefaults() *Contracts`

NewContractsWithDefaults instantiates a new Contracts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contracts) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contracts) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contracts) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Contracts) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Contracts) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Contracts) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Contracts) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Contracts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Contracts) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Contracts) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Contracts) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Contracts) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Contracts) GetItems() []Contract`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Contracts) GetItemsOk() (*[]Contract, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Contracts) SetItems(v []Contract)`

SetItems sets Items field to given value.

### HasItems

`func (o *Contracts) HasItems() bool`

HasItems returns a boolean if a field has been set.



