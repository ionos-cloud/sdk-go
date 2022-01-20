# NetworkLoadBalancerForwardingRuleHealthCheck

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ClientTimeout** | Pointer to **int32** | The maximum time in milliseconds to wait for the client to acknowledge or send data; default is 50,000 (50 seconds). | [optional] |
|**ConnectTimeout** | Pointer to **int32** | The maximum time in milliseconds to wait for a connection attempt to a target to succeed; default is 5000 (five seconds). | [optional] |
|**TargetTimeout** | Pointer to **int32** | The maximum time in milliseconds that a target can remain inactive; default is 50,000 (50 seconds). | [optional] |
|**Retries** | Pointer to **int32** | The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535 and default is three reconnection attempts. | [optional] |

## Methods

### NewNetworkLoadBalancerForwardingRuleHealthCheck

`func NewNetworkLoadBalancerForwardingRuleHealthCheck() *NetworkLoadBalancerForwardingRuleHealthCheck`

NewNetworkLoadBalancerForwardingRuleHealthCheck instantiates a new NetworkLoadBalancerForwardingRuleHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoadBalancerForwardingRuleHealthCheckWithDefaults

`func NewNetworkLoadBalancerForwardingRuleHealthCheckWithDefaults() *NetworkLoadBalancerForwardingRuleHealthCheck`

NewNetworkLoadBalancerForwardingRuleHealthCheckWithDefaults instantiates a new NetworkLoadBalancerForwardingRuleHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetClientTimeout() int32`

GetClientTimeout returns the ClientTimeout field if non-nil, zero value otherwise.

### GetClientTimeoutOk

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetClientTimeoutOk() (*int32, bool)`

GetClientTimeoutOk returns a tuple with the ClientTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetClientTimeout(v int32)`

SetClientTimeout sets ClientTimeout field to given value.

### HasClientTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasClientTimeout() bool`

HasClientTimeout returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetTargetTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetTargetTimeout() int32`

GetTargetTimeout returns the TargetTimeout field if non-nil, zero value otherwise.

### GetTargetTimeoutOk

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetTargetTimeoutOk() (*int32, bool)`

GetTargetTimeoutOk returns a tuple with the TargetTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetTargetTimeout(v int32)`

SetTargetTimeout sets TargetTimeout field to given value.

### HasTargetTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasTargetTimeout() bool`

HasTargetTimeout returns a boolean if a field has been set.

### GetRetries

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasRetries() bool`

HasRetries returns a boolean if a field has been set.



