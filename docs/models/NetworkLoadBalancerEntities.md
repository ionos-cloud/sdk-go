# NetworkLoadBalancerEntities

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Flowlogs** | Pointer to [**FlowLogs**](FlowLogs.md) |  | [optional] |
|**Forwardingrules** | Pointer to [**NetworkLoadBalancerForwardingRules**](NetworkLoadBalancerForwardingRules.md) |  | [optional] |

## Methods

### NewNetworkLoadBalancerEntities

`func NewNetworkLoadBalancerEntities() *NetworkLoadBalancerEntities`

NewNetworkLoadBalancerEntities instantiates a new NetworkLoadBalancerEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoadBalancerEntitiesWithDefaults

`func NewNetworkLoadBalancerEntitiesWithDefaults() *NetworkLoadBalancerEntities`

NewNetworkLoadBalancerEntitiesWithDefaults instantiates a new NetworkLoadBalancerEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowlogs

`func (o *NetworkLoadBalancerEntities) GetFlowlogs() FlowLogs`

GetFlowlogs returns the Flowlogs field if non-nil, zero value otherwise.

### GetFlowlogsOk

`func (o *NetworkLoadBalancerEntities) GetFlowlogsOk() (*FlowLogs, bool)`

GetFlowlogsOk returns a tuple with the Flowlogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowlogs

`func (o *NetworkLoadBalancerEntities) SetFlowlogs(v FlowLogs)`

SetFlowlogs sets Flowlogs field to given value.

### HasFlowlogs

`func (o *NetworkLoadBalancerEntities) HasFlowlogs() bool`

HasFlowlogs returns a boolean if a field has been set.

### GetForwardingrules

`func (o *NetworkLoadBalancerEntities) GetForwardingrules() NetworkLoadBalancerForwardingRules`

GetForwardingrules returns the Forwardingrules field if non-nil, zero value otherwise.

### GetForwardingrulesOk

`func (o *NetworkLoadBalancerEntities) GetForwardingrulesOk() (*NetworkLoadBalancerForwardingRules, bool)`

GetForwardingrulesOk returns a tuple with the Forwardingrules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingrules

`func (o *NetworkLoadBalancerEntities) SetForwardingrules(v NetworkLoadBalancerForwardingRules)`

SetForwardingrules sets Forwardingrules field to given value.

### HasForwardingrules

`func (o *NetworkLoadBalancerEntities) HasForwardingrules() bool`

HasForwardingrules returns a boolean if a field has been set.



