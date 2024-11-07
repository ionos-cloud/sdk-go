# DatacenterPut

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**DatacenterPropertiesPut**](DatacenterPropertiesPut.md) |  | |
|**Entities** | Pointer to [**DataCenterEntities**](DataCenterEntities.md) |  | [optional] |

## Methods

### NewDatacenterPut

`func NewDatacenterPut(properties DatacenterPropertiesPut, ) *DatacenterPut`

NewDatacenterPut instantiates a new DatacenterPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterPutWithDefaults

`func NewDatacenterPutWithDefaults() *DatacenterPut`

NewDatacenterPutWithDefaults instantiates a new DatacenterPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *DatacenterPut) GetProperties() DatacenterPropertiesPut`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DatacenterPut) GetPropertiesOk() (*DatacenterPropertiesPut, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DatacenterPut) SetProperties(v DatacenterPropertiesPut)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *DatacenterPut) GetEntities() DataCenterEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *DatacenterPut) GetEntitiesOk() (*DataCenterEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *DatacenterPut) SetEntities(v DataCenterEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *DatacenterPut) HasEntities() bool`

HasEntities returns a boolean if a field has been set.



