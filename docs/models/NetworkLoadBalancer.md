# NetworkLoadBalancer

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**NetworkLoadBalancerProperties**](NetworkLoadBalancerProperties.md) |  | |
|**Entities** | Pointer to [**NetworkLoadBalancerEntities**](NetworkLoadBalancerEntities.md) |  | [optional] |

## Methods

### NewNetworkLoadBalancer

`func NewNetworkLoadBalancer(properties NetworkLoadBalancerProperties, ) *NetworkLoadBalancer`

NewNetworkLoadBalancer instantiates a new NetworkLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoadBalancerWithDefaults

`func NewNetworkLoadBalancerWithDefaults() *NetworkLoadBalancer`

NewNetworkLoadBalancerWithDefaults instantiates a new NetworkLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkLoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkLoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkLoadBalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkLoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NetworkLoadBalancer) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkLoadBalancer) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkLoadBalancer) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *NetworkLoadBalancer) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NetworkLoadBalancer) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NetworkLoadBalancer) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NetworkLoadBalancer) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkLoadBalancer) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkLoadBalancer) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkLoadBalancer) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkLoadBalancer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *NetworkLoadBalancer) GetProperties() NetworkLoadBalancerProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NetworkLoadBalancer) GetPropertiesOk() (*NetworkLoadBalancerProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NetworkLoadBalancer) SetProperties(v NetworkLoadBalancerProperties)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *NetworkLoadBalancer) GetEntities() NetworkLoadBalancerEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *NetworkLoadBalancer) GetEntitiesOk() (*NetworkLoadBalancerEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *NetworkLoadBalancer) SetEntities(v NetworkLoadBalancerEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *NetworkLoadBalancer) HasEntities() bool`

HasEntities returns a boolean if a field has been set.



