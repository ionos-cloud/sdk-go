# NatGateway

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**NatGatewayProperties**](NatGatewayProperties.md) |  | |
|**Entities** | Pointer to [**NatGatewayEntities**](NatGatewayEntities.md) |  | [optional] |

## Methods

### NewNatGateway

`func NewNatGateway(properties NatGatewayProperties, ) *NatGateway`

NewNatGateway instantiates a new NatGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatGatewayWithDefaults

`func NewNatGatewayWithDefaults() *NatGateway`

NewNatGatewayWithDefaults instantiates a new NatGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NatGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NatGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NatGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NatGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NatGateway) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NatGateway) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NatGateway) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *NatGateway) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *NatGateway) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NatGateway) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NatGateway) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NatGateway) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *NatGateway) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NatGateway) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NatGateway) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NatGateway) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *NatGateway) GetProperties() NatGatewayProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NatGateway) GetPropertiesOk() (*NatGatewayProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NatGateway) SetProperties(v NatGatewayProperties)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *NatGateway) GetEntities() NatGatewayEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *NatGateway) GetEntitiesOk() (*NatGatewayEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *NatGateway) SetEntities(v NatGatewayEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *NatGateway) HasEntities() bool`

HasEntities returns a boolean if a field has been set.



