# Datacenter

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**DatacenterProperties**](DatacenterProperties.md) |  | |
|**Entities** | Pointer to [**DataCenterEntities**](DataCenterEntities.md) |  | [optional] |

## Methods

### NewDatacenter

`func NewDatacenter(properties DatacenterProperties, ) *Datacenter`

NewDatacenter instantiates a new Datacenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterWithDefaults

`func NewDatacenterWithDefaults() *Datacenter`

NewDatacenterWithDefaults instantiates a new Datacenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Datacenter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Datacenter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Datacenter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Datacenter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Datacenter) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Datacenter) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Datacenter) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Datacenter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Datacenter) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Datacenter) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Datacenter) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Datacenter) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *Datacenter) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Datacenter) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Datacenter) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Datacenter) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *Datacenter) GetProperties() DatacenterProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Datacenter) GetPropertiesOk() (*DatacenterProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Datacenter) SetProperties(v DatacenterProperties)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *Datacenter) GetEntities() DataCenterEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *Datacenter) GetEntitiesOk() (*DataCenterEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *Datacenter) SetEntities(v DataCenterEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *Datacenter) HasEntities() bool`

HasEntities returns a boolean if a field has been set.



