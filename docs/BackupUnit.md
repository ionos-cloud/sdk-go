# BackupUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to **string** | The type of object that has been created | [optional] [readonly] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] 
**Properties** | [**BackupUnitProperties**](BackupUnitProperties.md) |  | 

## Methods

### NewBackupUnit

`func NewBackupUnit(properties BackupUnitProperties, ) *BackupUnit`

NewBackupUnit instantiates a new BackupUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupUnitWithDefaults

`func NewBackupUnitWithDefaults() *BackupUnit`

NewBackupUnitWithDefaults instantiates a new BackupUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupUnit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupUnit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupUnit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupUnit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BackupUnit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupUnit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupUnit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BackupUnit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *BackupUnit) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BackupUnit) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BackupUnit) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BackupUnit) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *BackupUnit) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BackupUnit) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BackupUnit) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BackupUnit) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *BackupUnit) GetProperties() BackupUnitProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BackupUnit) GetPropertiesOk() (*BackupUnitProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BackupUnit) SetProperties(v BackupUnitProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


