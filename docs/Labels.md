# Labels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique representation for Label as a collection of resource. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of resource within a collection | [optional] [readonly] 
**Href** | Pointer to **string** | URL to the collection representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]Label**](Label.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewLabels

`func NewLabels() *Labels`

NewLabels instantiates a new Labels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelsWithDefaults

`func NewLabelsWithDefaults() *Labels`

NewLabelsWithDefaults instantiates a new Labels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Labels) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Labels) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Labels) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Labels) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Labels) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Labels) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Labels) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Labels) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Labels) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Labels) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Labels) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Labels) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Labels) GetItems() []Label`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Labels) GetItemsOk() (*[]Label, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Labels) SetItems(v []Label)`

SetItems sets Items field to given value.

### HasItems

`func (o *Labels) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


