# BackupUnits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to **string** | The type of object that has been created | [optional] [readonly] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]BackupUnit**](BackupUnit.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewBackupUnits

`func NewBackupUnits() *BackupUnits`

NewBackupUnits instantiates a new BackupUnits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupUnitsWithDefaults

`func NewBackupUnitsWithDefaults() *BackupUnits`

NewBackupUnitsWithDefaults instantiates a new BackupUnits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupUnits) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupUnits) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupUnits) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupUnits) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BackupUnits) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupUnits) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupUnits) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BackupUnits) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *BackupUnits) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BackupUnits) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BackupUnits) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BackupUnits) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *BackupUnits) GetItems() []BackupUnit`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BackupUnits) GetItemsOk() (*[]BackupUnit, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BackupUnits) SetItems(v []BackupUnit)`

SetItems sets Items field to given value.

### HasItems

`func (o *BackupUnits) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


