# NetworkLoadBalancerForwardingRuleTargetHealthCheck

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Check** | Pointer to **bool** | Makes the target available only if it accepts periodic health check TCP connection attempts; when turned off, the target is considered always available. The health check only consists of a connection attempt to the address and port of the target. | [optional] |
|**CheckInterval** | Pointer to **int32** | The interval in milliseconds between consecutive health checks; default is 2000. | [optional] |
|**Maintenance** | Pointer to **bool** | Maintenance mode prevents the target from receiving balanced traffic. | [optional] |

## Methods

### NewNetworkLoadBalancerForwardingRuleTargetHealthCheck

`func NewNetworkLoadBalancerForwardingRuleTargetHealthCheck() *NetworkLoadBalancerForwardingRuleTargetHealthCheck`

NewNetworkLoadBalancerForwardingRuleTargetHealthCheck instantiates a new NetworkLoadBalancerForwardingRuleTargetHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoadBalancerForwardingRuleTargetHealthCheckWithDefaults

`func NewNetworkLoadBalancerForwardingRuleTargetHealthCheckWithDefaults() *NetworkLoadBalancerForwardingRuleTargetHealthCheck`

NewNetworkLoadBalancerForwardingRuleTargetHealthCheckWithDefaults instantiates a new NetworkLoadBalancerForwardingRuleTargetHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheck

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) GetCheck() bool`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) GetCheckOk() (*bool, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) SetCheck(v bool)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetCheckInterval

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetMaintenance

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *NetworkLoadBalancerForwardingRuleTargetHealthCheck) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.



