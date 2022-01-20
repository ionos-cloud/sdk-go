# NetworkLoadBalancerProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the Network Load Balancer. | |
|**ListenerLan** | **int32** | ID of the listening LAN (inbound). | |
|**Ips** | Pointer to **[]string** | Collection of the Network Load Balancer IP addresses. (Inbound and outbound) IPs of the listenerLan must be customer-reserved IPs for public Load Balancers, and private IPs for private Load Balancers. | [optional] |
|**TargetLan** | **int32** | ID of the balanced private target LAN (outbound). | |
|**LbPrivateIps** | Pointer to **[]string** | Collection of private IP addresses with subnet mask of the Network Load Balancer. IPs must contain a valid subnet mask. If no IP is provided, the system will generate an IP with /24 subnet. | [optional] |

## Methods

### NewNetworkLoadBalancerProperties

`func NewNetworkLoadBalancerProperties(name string, listenerLan int32, targetLan int32, ) *NetworkLoadBalancerProperties`

NewNetworkLoadBalancerProperties instantiates a new NetworkLoadBalancerProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoadBalancerPropertiesWithDefaults

`func NewNetworkLoadBalancerPropertiesWithDefaults() *NetworkLoadBalancerProperties`

NewNetworkLoadBalancerPropertiesWithDefaults instantiates a new NetworkLoadBalancerProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkLoadBalancerProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkLoadBalancerProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkLoadBalancerProperties) SetName(v string)`

SetName sets Name field to given value.


### GetListenerLan

`func (o *NetworkLoadBalancerProperties) GetListenerLan() int32`

GetListenerLan returns the ListenerLan field if non-nil, zero value otherwise.

### GetListenerLanOk

`func (o *NetworkLoadBalancerProperties) GetListenerLanOk() (*int32, bool)`

GetListenerLanOk returns a tuple with the ListenerLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerLan

`func (o *NetworkLoadBalancerProperties) SetListenerLan(v int32)`

SetListenerLan sets ListenerLan field to given value.


### GetIps

`func (o *NetworkLoadBalancerProperties) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *NetworkLoadBalancerProperties) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *NetworkLoadBalancerProperties) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *NetworkLoadBalancerProperties) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetTargetLan

`func (o *NetworkLoadBalancerProperties) GetTargetLan() int32`

GetTargetLan returns the TargetLan field if non-nil, zero value otherwise.

### GetTargetLanOk

`func (o *NetworkLoadBalancerProperties) GetTargetLanOk() (*int32, bool)`

GetTargetLanOk returns a tuple with the TargetLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLan

`func (o *NetworkLoadBalancerProperties) SetTargetLan(v int32)`

SetTargetLan sets TargetLan field to given value.


### GetLbPrivateIps

`func (o *NetworkLoadBalancerProperties) GetLbPrivateIps() []string`

GetLbPrivateIps returns the LbPrivateIps field if non-nil, zero value otherwise.

### GetLbPrivateIpsOk

`func (o *NetworkLoadBalancerProperties) GetLbPrivateIpsOk() (*[]string, bool)`

GetLbPrivateIpsOk returns a tuple with the LbPrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbPrivateIps

`func (o *NetworkLoadBalancerProperties) SetLbPrivateIps(v []string)`

SetLbPrivateIps sets LbPrivateIps field to given value.

### HasLbPrivateIps

`func (o *NetworkLoadBalancerProperties) HasLbPrivateIps() bool`

HasLbPrivateIps returns a boolean if a field has been set.



