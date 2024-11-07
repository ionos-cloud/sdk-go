# CreateSnapshot

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | Pointer to [**CreateSnapshotProperties**](CreateSnapshotProperties.md) |  | [optional] |

## Methods

### NewCreateSnapshot

`func NewCreateSnapshot() *CreateSnapshot`

NewCreateSnapshot instantiates a new CreateSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotWithDefaults

`func NewCreateSnapshotWithDefaults() *CreateSnapshot`

NewCreateSnapshotWithDefaults instantiates a new CreateSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *CreateSnapshot) GetProperties() CreateSnapshotProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateSnapshot) GetPropertiesOk() (*CreateSnapshotProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateSnapshot) SetProperties(v CreateSnapshotProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateSnapshot) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



