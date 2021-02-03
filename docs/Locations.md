# Locations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]Location**](Location.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewLocations

`func NewLocations() *Locations`

NewLocations instantiates a new Locations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationsWithDefaults

`func NewLocationsWithDefaults() *Locations`

NewLocationsWithDefaults instantiates a new Locations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Locations) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Locations) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Locations) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Locations) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Locations) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Locations) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Locations) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Locations) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Locations) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Locations) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Locations) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Locations) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Locations) GetItems() []Location`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Locations) GetItemsOk() (*[]Location, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Locations) SetItems(v []Location)`

SetItems sets Items field to given value.

### HasItems

`func (o *Locations) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


