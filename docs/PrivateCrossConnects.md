# PrivateCrossConnects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]PrivateCrossConnect**](PrivateCrossConnect.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewPrivateCrossConnects

`func NewPrivateCrossConnects() *PrivateCrossConnects`

NewPrivateCrossConnects instantiates a new PrivateCrossConnects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateCrossConnectsWithDefaults

`func NewPrivateCrossConnectsWithDefaults() *PrivateCrossConnects`

NewPrivateCrossConnectsWithDefaults instantiates a new PrivateCrossConnects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateCrossConnects) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateCrossConnects) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateCrossConnects) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateCrossConnects) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PrivateCrossConnects) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateCrossConnects) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateCrossConnects) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *PrivateCrossConnects) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *PrivateCrossConnects) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PrivateCrossConnects) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PrivateCrossConnects) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PrivateCrossConnects) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *PrivateCrossConnects) GetItems() []PrivateCrossConnect`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PrivateCrossConnects) GetItemsOk() (*[]PrivateCrossConnect, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PrivateCrossConnects) SetItems(v []PrivateCrossConnect)`

SetItems sets Items field to given value.

### HasItems

`func (o *PrivateCrossConnects) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


