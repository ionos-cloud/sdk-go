# NetworkLoadBalancerForwardingRuleProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the Network Load Balancer forwarding rule. | |
|**Algorithm** | **string** | Balancing algorithm | |
|**Protocol** | **string** | Balancing protocol | |
|**ListenerIp** | **string** | Listening (inbound) IP | |
|**ListenerPort** | **int32** | Listening (inbound) port number; valid range is 1 to 65535. | |
|**HealthCheck** | Pointer to [**NetworkLoadBalancerForwardingRuleHealthCheck**](NetworkLoadBalancerForwardingRuleHealthCheck.md) |  | [optional] |
|**Targets** | [**[]NetworkLoadBalancerForwardingRuleTarget**](NetworkLoadBalancerForwardingRuleTarget.md) | Array of items in the collection. | |

## Methods

### NewNetworkLoadBalancerForwardingRuleProperties

`func NewNetworkLoadBalancerForwardingRuleProperties(name string, algorithm string, protocol string, listenerIp string, listenerPort int32, targets []NetworkLoadBalancerForwardingRuleTarget, ) *NetworkLoadBalancerForwardingRuleProperties`

NewNetworkLoadBalancerForwardingRuleProperties instantiates a new NetworkLoadBalancerForwardingRuleProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoadBalancerForwardingRulePropertiesWithDefaults

`func NewNetworkLoadBalancerForwardingRulePropertiesWithDefaults() *NetworkLoadBalancerForwardingRuleProperties`

NewNetworkLoadBalancerForwardingRulePropertiesWithDefaults instantiates a new NetworkLoadBalancerForwardingRuleProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkLoadBalancerForwardingRuleProperties) SetName(v string)`

SetName sets Name field to given value.


### GetAlgorithm

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *NetworkLoadBalancerForwardingRuleProperties) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetProtocol

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworkLoadBalancerForwardingRuleProperties) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetListenerIp

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetListenerIp() string`

GetListenerIp returns the ListenerIp field if non-nil, zero value otherwise.

### GetListenerIpOk

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetListenerIpOk() (*string, bool)`

GetListenerIpOk returns a tuple with the ListenerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIp

`func (o *NetworkLoadBalancerForwardingRuleProperties) SetListenerIp(v string)`

SetListenerIp sets ListenerIp field to given value.


### GetListenerPort

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetListenerPort() int32`

GetListenerPort returns the ListenerPort field if non-nil, zero value otherwise.

### GetListenerPortOk

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetListenerPortOk() (*int32, bool)`

GetListenerPortOk returns a tuple with the ListenerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerPort

`func (o *NetworkLoadBalancerForwardingRuleProperties) SetListenerPort(v int32)`

SetListenerPort sets ListenerPort field to given value.


### GetHealthCheck

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetHealthCheck() NetworkLoadBalancerForwardingRuleHealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetHealthCheckOk() (*NetworkLoadBalancerForwardingRuleHealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *NetworkLoadBalancerForwardingRuleProperties) SetHealthCheck(v NetworkLoadBalancerForwardingRuleHealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *NetworkLoadBalancerForwardingRuleProperties) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetTargets

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetTargets() []NetworkLoadBalancerForwardingRuleTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *NetworkLoadBalancerForwardingRuleProperties) GetTargetsOk() (*[]NetworkLoadBalancerForwardingRuleTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *NetworkLoadBalancerForwardingRuleProperties) SetTargets(v []NetworkLoadBalancerForwardingRuleTarget)`

SetTargets sets Targets field to given value.




