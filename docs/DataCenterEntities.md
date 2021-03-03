# DataCenterEntities

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Servers** | Pointer to [**Servers**](Servers.md) |  | [optional] |
|**Volumes** | Pointer to [**Volumes**](Volumes.md) |  | [optional] |
|**Loadbalancers** | Pointer to [**Loadbalancers**](Loadbalancers.md) |  | [optional] |
|**Lans** | Pointer to [**Lans**](Lans.md) |  | [optional] |
|**Networkloadbalancers** | Pointer to [**NetworkLoadBalancers**](NetworkLoadBalancers.md) |  | [optional] |
|**Natgateways** | Pointer to [**NatGateways**](NatGateways.md) |  | [optional] |

## Methods

### NewDataCenterEntities

`func NewDataCenterEntities() *DataCenterEntities`

NewDataCenterEntities instantiates a new DataCenterEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataCenterEntitiesWithDefaults

`func NewDataCenterEntitiesWithDefaults() *DataCenterEntities`

NewDataCenterEntitiesWithDefaults instantiates a new DataCenterEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *DataCenterEntities) GetServers() Servers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DataCenterEntities) GetServersOk() (*Servers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DataCenterEntities) SetServers(v Servers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *DataCenterEntities) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetVolumes

`func (o *DataCenterEntities) GetVolumes() Volumes`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *DataCenterEntities) GetVolumesOk() (*Volumes, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *DataCenterEntities) SetVolumes(v Volumes)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *DataCenterEntities) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetLoadbalancers

`func (o *DataCenterEntities) GetLoadbalancers() Loadbalancers`

GetLoadbalancers returns the Loadbalancers field if non-nil, zero value otherwise.

### GetLoadbalancersOk

`func (o *DataCenterEntities) GetLoadbalancersOk() (*Loadbalancers, bool)`

GetLoadbalancersOk returns a tuple with the Loadbalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancers

`func (o *DataCenterEntities) SetLoadbalancers(v Loadbalancers)`

SetLoadbalancers sets Loadbalancers field to given value.

### HasLoadbalancers

`func (o *DataCenterEntities) HasLoadbalancers() bool`

HasLoadbalancers returns a boolean if a field has been set.

### GetLans

`func (o *DataCenterEntities) GetLans() Lans`

GetLans returns the Lans field if non-nil, zero value otherwise.

### GetLansOk

`func (o *DataCenterEntities) GetLansOk() (*Lans, bool)`

GetLansOk returns a tuple with the Lans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLans

`func (o *DataCenterEntities) SetLans(v Lans)`

SetLans sets Lans field to given value.

### HasLans

`func (o *DataCenterEntities) HasLans() bool`

HasLans returns a boolean if a field has been set.

### GetNetworkloadbalancers

`func (o *DataCenterEntities) GetNetworkloadbalancers() NetworkLoadBalancers`

GetNetworkloadbalancers returns the Networkloadbalancers field if non-nil, zero value otherwise.

### GetNetworkloadbalancersOk

`func (o *DataCenterEntities) GetNetworkloadbalancersOk() (*NetworkLoadBalancers, bool)`

GetNetworkloadbalancersOk returns a tuple with the Networkloadbalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkloadbalancers

`func (o *DataCenterEntities) SetNetworkloadbalancers(v NetworkLoadBalancers)`

SetNetworkloadbalancers sets Networkloadbalancers field to given value.

### HasNetworkloadbalancers

`func (o *DataCenterEntities) HasNetworkloadbalancers() bool`

HasNetworkloadbalancers returns a boolean if a field has been set.

### GetNatgateways

`func (o *DataCenterEntities) GetNatgateways() NatGateways`

GetNatgateways returns the Natgateways field if non-nil, zero value otherwise.

### GetNatgatewaysOk

`func (o *DataCenterEntities) GetNatgatewaysOk() (*NatGateways, bool)`

GetNatgatewaysOk returns a tuple with the Natgateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatgateways

`func (o *DataCenterEntities) SetNatgateways(v NatGateways)`

SetNatgateways sets Natgateways field to given value.

### HasNatgateways

`func (o *DataCenterEntities) HasNatgateways() bool`

HasNatgateways returns a boolean if a field has been set.



