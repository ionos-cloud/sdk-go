# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] 
**Properties** | [**SnapshotProperties**](SnapshotProperties.md) |  | 

## Methods

### NewSnapshot

`func NewSnapshot(properties SnapshotProperties, ) *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Snapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Snapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Snapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Snapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Snapshot) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Snapshot) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Snapshot) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Snapshot) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Snapshot) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Snapshot) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Snapshot) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Snapshot) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *Snapshot) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Snapshot) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Snapshot) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Snapshot) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *Snapshot) GetProperties() SnapshotProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Snapshot) GetPropertiesOk() (*SnapshotProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Snapshot) SetProperties(v SnapshotProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


