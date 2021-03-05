# NetworkLoadBalancerForwardingRuleTarget

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Ip** | **string** | IP of a balanced target VM |  |
| **Port** | **int32** | Port of the balanced target service. \(range: 1 to 65535\) |  |
| **Weight** | **int32** | Weight parameter is used to adjust the target VM's weight relative to other target VMs. All target VMs will receive a load proportional to their weight relative to the sum of all weights, so the higher the weight, the higher the load. The default weight is 1, and the maximal value is 256. A value of 0 means the target VM will not participate in load-balancing but will still accept persistent connections. If this parameter is used to distribute the load according to target VM's capacity, it is recommended to start with values which can both grow and shrink, for instance between 10 and 100 to leave enough room above and below for later adjustments. |  |
| **HealthCheck** | Pointer to [**NetworkLoadBalancerForwardingRuleTargetHealthCheck**](networkloadbalancerforwardingruletargethealthcheck.md) |  | \[optional\] |

## Methods

### NewNetworkLoadBalancerForwardingRuleTarget

`func NewNetworkLoadBalancerForwardingRuleTarget(ip string, port int32, weight int32, ) *NetworkLoadBalancerForwardingRuleTarget`

NewNetworkLoadBalancerForwardingRuleTarget instantiates a new NetworkLoadBalancerForwardingRuleTarget object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewNetworkLoadBalancerForwardingRuleTargetWithDefaults

`func NewNetworkLoadBalancerForwardingRuleTargetWithDefaults() *NetworkLoadBalancerForwardingRuleTarget`

NewNetworkLoadBalancerForwardingRuleTargetWithDefaults instantiates a new NetworkLoadBalancerForwardingRuleTarget object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *NetworkLoadBalancerForwardingRuleTarget) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NetworkLoadBalancerForwardingRuleTarget) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetIp

`func (o *NetworkLoadBalancerForwardingRuleTarget) SetIp(v string)`

SetIp sets Ip field to given value.

### GetPort

`func (o *NetworkLoadBalancerForwardingRuleTarget) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworkLoadBalancerForwardingRuleTarget) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetPort

`func (o *NetworkLoadBalancerForwardingRuleTarget) SetPort(v int32)`

SetPort sets Port field to given value.

### GetWeight

`func (o *NetworkLoadBalancerForwardingRuleTarget) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *NetworkLoadBalancerForwardingRuleTarget) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetWeight

`func (o *NetworkLoadBalancerForwardingRuleTarget) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### GetHealthCheck

`func (o *NetworkLoadBalancerForwardingRuleTarget) GetHealthCheck() NetworkLoadBalancerForwardingRuleTargetHealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *NetworkLoadBalancerForwardingRuleTarget) GetHealthCheckOk() (*NetworkLoadBalancerForwardingRuleTargetHealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *NetworkLoadBalancerForwardingRuleTarget) SetHealthCheck(v NetworkLoadBalancerForwardingRuleTargetHealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *NetworkLoadBalancerForwardingRuleTarget) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

