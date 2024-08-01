# NatGatewayLanProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **int32** | Id for the LAN connected to the NAT Gateway | |
|**GatewayIps** | Pointer to **[]string** | Collection of gateway IP addresses of the NAT Gateway. Will be auto-generated if not provided. Should ideally be an IP belonging to the same subnet as the LAN | [optional] |

## Methods

### NewNatGatewayLanProperties

`func NewNatGatewayLanProperties(id int32, ) *NatGatewayLanProperties`

NewNatGatewayLanProperties instantiates a new NatGatewayLanProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatGatewayLanPropertiesWithDefaults

`func NewNatGatewayLanPropertiesWithDefaults() *NatGatewayLanProperties`

NewNatGatewayLanPropertiesWithDefaults instantiates a new NatGatewayLanProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NatGatewayLanProperties) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NatGatewayLanProperties) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NatGatewayLanProperties) SetId(v int32)`

SetId sets Id field to given value.


### GetGatewayIps

`func (o *NatGatewayLanProperties) GetGatewayIps() []string`

GetGatewayIps returns the GatewayIps field if non-nil, zero value otherwise.

### GetGatewayIpsOk

`func (o *NatGatewayLanProperties) GetGatewayIpsOk() (*[]string, bool)`

GetGatewayIpsOk returns a tuple with the GatewayIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIps

`func (o *NatGatewayLanProperties) SetGatewayIps(v []string)`

SetGatewayIps sets GatewayIps field to given value.

### HasGatewayIps

`func (o *NatGatewayLanProperties) HasGatewayIps() bool`

HasGatewayIps returns a boolean if a field has been set.



