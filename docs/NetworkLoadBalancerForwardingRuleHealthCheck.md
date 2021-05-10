# NetworkLoadBalancerForwardingRuleHealthCheck

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ClientTimeout** | Pointer to **int32** | ClientTimeout is expressed in milliseconds. This inactivity timeout applies when the client is expected to acknowledge or send data. If unset the default of 50 seconds will be used. | [optional] |
|**CheckTimeout** | Pointer to **int32** | It specifies the time (in milliseconds) for a target VM in this pool to answer the check. If a target VM has CheckInterval set and CheckTimeout is set too, then the smaller value of the two is used after the TCP connection is established. | [optional] |
|**ConnectTimeout** | Pointer to **int32** | It specifies the maximum time (in milliseconds) to wait for a connection attempt to a target VM to succeed. If unset, the default of 5 seconds will be used. | [optional] |
|**TargetTimeout** | Pointer to **int32** | TargetTimeout specifies the maximum inactivity time (in milliseconds) on the target VM side. If unset, the default of 50 seconds will be used. | [optional] |
|**Retries** | Pointer to **int32** | Retries specifies the number of retries to perform on a target VM after a connection failure. If unset, the default value of 3 will be used. (valid range: [0, 65535]) | [optional] |

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

### GetCheckTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetCheckTimeout() int32`

GetCheckTimeout returns the CheckTimeout field if non-nil, zero value otherwise.

### GetCheckTimeoutOk

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetCheckTimeoutOk() (*int32, bool)`

GetCheckTimeoutOk returns a tuple with the CheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetCheckTimeout(v int32)`

SetCheckTimeout sets CheckTimeout field to given value.

### HasCheckTimeout

`func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasCheckTimeout() bool`

HasCheckTimeout returns a boolean if a field has been set.

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



