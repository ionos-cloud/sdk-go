# Loadbalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] 
**Properties** | [**LoadbalancerProperties**](LoadbalancerProperties.md) |  | 
**Entities** | Pointer to [**LoadbalancerEntities**](LoadbalancerEntities.md) |  | [optional] 

## Methods

### NewLoadbalancer

`func NewLoadbalancer(properties LoadbalancerProperties, ) *Loadbalancer`

NewLoadbalancer instantiates a new Loadbalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerWithDefaults

`func NewLoadbalancerWithDefaults() *Loadbalancer`

NewLoadbalancerWithDefaults instantiates a new Loadbalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Loadbalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Loadbalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Loadbalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Loadbalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Loadbalancer) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Loadbalancer) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Loadbalancer) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Loadbalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Loadbalancer) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Loadbalancer) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Loadbalancer) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Loadbalancer) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *Loadbalancer) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Loadbalancer) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Loadbalancer) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Loadbalancer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *Loadbalancer) GetProperties() LoadbalancerProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Loadbalancer) GetPropertiesOk() (*LoadbalancerProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Loadbalancer) SetProperties(v LoadbalancerProperties)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *Loadbalancer) GetEntities() LoadbalancerEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *Loadbalancer) GetEntitiesOk() (*LoadbalancerEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *Loadbalancer) SetEntities(v LoadbalancerEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *Loadbalancer) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


