# RestoreSnapshotProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SnapshotId** | Pointer to **string** | The id of the snapshot | [optional] |

## Methods

### NewRestoreSnapshotProperties

`func NewRestoreSnapshotProperties() *RestoreSnapshotProperties`

NewRestoreSnapshotProperties instantiates a new RestoreSnapshotProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSnapshotPropertiesWithDefaults

`func NewRestoreSnapshotPropertiesWithDefaults() *RestoreSnapshotProperties`

NewRestoreSnapshotPropertiesWithDefaults instantiates a new RestoreSnapshotProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *RestoreSnapshotProperties) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *RestoreSnapshotProperties) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *RestoreSnapshotProperties) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *RestoreSnapshotProperties) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.



