# ServerEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdroms** | Pointer to [**Cdroms**](Cdroms.md) |  | [optional] 
**Volumes** | Pointer to [**AttachedVolumes**](AttachedVolumes.md) |  | [optional] 
**Nics** | Pointer to [**Nics**](Nics.md) |  | [optional] 

## Methods

### NewServerEntities

`func NewServerEntities() *ServerEntities`

NewServerEntities instantiates a new ServerEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerEntitiesWithDefaults

`func NewServerEntitiesWithDefaults() *ServerEntities`

NewServerEntitiesWithDefaults instantiates a new ServerEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdroms

`func (o *ServerEntities) GetCdroms() Cdroms`

GetCdroms returns the Cdroms field if non-nil, zero value otherwise.

### GetCdromsOk

`func (o *ServerEntities) GetCdromsOk() (*Cdroms, bool)`

GetCdromsOk returns a tuple with the Cdroms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdroms

`func (o *ServerEntities) SetCdroms(v Cdroms)`

SetCdroms sets Cdroms field to given value.

### HasCdroms

`func (o *ServerEntities) HasCdroms() bool`

HasCdroms returns a boolean if a field has been set.

### GetVolumes

`func (o *ServerEntities) GetVolumes() AttachedVolumes`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ServerEntities) GetVolumesOk() (*AttachedVolumes, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ServerEntities) SetVolumes(v AttachedVolumes)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ServerEntities) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNics

`func (o *ServerEntities) GetNics() Nics`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *ServerEntities) GetNicsOk() (*Nics, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *ServerEntities) SetNics(v Nics)`

SetNics sets Nics field to given value.

### HasNics

`func (o *ServerEntities) HasNics() bool`

HasNics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


