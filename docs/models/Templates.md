# Templates

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]Template**](Template.md) | Array of items in the collection. | [optional] [readonly] |

## Methods

### NewTemplates

`func NewTemplates() *Templates`

NewTemplates instantiates a new Templates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesWithDefaults

`func NewTemplatesWithDefaults() *Templates`

NewTemplatesWithDefaults instantiates a new Templates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Templates) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Templates) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Templates) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Templates) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Templates) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Templates) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Templates) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Templates) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Templates) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Templates) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Templates) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Templates) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *Templates) GetItems() []Template`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Templates) GetItemsOk() (*[]Template, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Templates) SetItems(v []Template)`

SetItems sets Items field to given value.

### HasItems

`func (o *Templates) HasItems() bool`

HasItems returns a boolean if a field has been set.



