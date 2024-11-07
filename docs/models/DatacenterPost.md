# DatacenterPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**DatacenterPropertiesPost**](DatacenterPropertiesPost.md) |  | |
|**Entities** | Pointer to [**DataCenterEntities**](DataCenterEntities.md) |  | [optional] |

## Methods

### NewDatacenterPost

`func NewDatacenterPost(properties DatacenterPropertiesPost, ) *DatacenterPost`

NewDatacenterPost instantiates a new DatacenterPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterPostWithDefaults

`func NewDatacenterPostWithDefaults() *DatacenterPost`

NewDatacenterPostWithDefaults instantiates a new DatacenterPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *DatacenterPost) GetProperties() DatacenterPropertiesPost`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DatacenterPost) GetPropertiesOk() (*DatacenterPropertiesPost, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DatacenterPost) SetProperties(v DatacenterPropertiesPost)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *DatacenterPost) GetEntities() DataCenterEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *DatacenterPost) GetEntitiesOk() (*DataCenterEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *DatacenterPost) SetEntities(v DataCenterEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *DatacenterPost) HasEntities() bool`

HasEntities returns a boolean if a field has been set.



