# ApplicationLoadBalancerProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the Application Load Balancer. | |
|**ListenerLan** | **int32** | ID of the listening (inbound) LAN. | |
|**Ips** | Pointer to **[]string** | Collection of the Application Load Balancer IP addresses. (Inbound and outbound) IPs of the listenerLan are customer-reserved public IPs for the public Load Balancers, and private IPs for the private Load Balancers. | [optional] |
|**TargetLan** | **int32** | ID of the balanced private target LAN (outbound). | |
|**LbPrivateIps** | Pointer to **[]string** | Collection of private IP addresses with the subnet mask of the Application Load Balancer. IPs must contain valid a subnet mask. If no IP is provided, the system will generate an IP with /24 subnet. | [optional] |

## Methods

### NewApplicationLoadBalancerProperties

`func NewApplicationLoadBalancerProperties(name string, listenerLan int32, targetLan int32, ) *ApplicationLoadBalancerProperties`

NewApplicationLoadBalancerProperties instantiates a new ApplicationLoadBalancerProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerPropertiesWithDefaults

`func NewApplicationLoadBalancerPropertiesWithDefaults() *ApplicationLoadBalancerProperties`

NewApplicationLoadBalancerPropertiesWithDefaults instantiates a new ApplicationLoadBalancerProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationLoadBalancerProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationLoadBalancerProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationLoadBalancerProperties) SetName(v string)`

SetName sets Name field to given value.


### GetListenerLan

`func (o *ApplicationLoadBalancerProperties) GetListenerLan() int32`

GetListenerLan returns the ListenerLan field if non-nil, zero value otherwise.

### GetListenerLanOk

`func (o *ApplicationLoadBalancerProperties) GetListenerLanOk() (*int32, bool)`

GetListenerLanOk returns a tuple with the ListenerLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerLan

`func (o *ApplicationLoadBalancerProperties) SetListenerLan(v int32)`

SetListenerLan sets ListenerLan field to given value.


### GetIps

`func (o *ApplicationLoadBalancerProperties) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ApplicationLoadBalancerProperties) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ApplicationLoadBalancerProperties) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *ApplicationLoadBalancerProperties) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetTargetLan

`func (o *ApplicationLoadBalancerProperties) GetTargetLan() int32`

GetTargetLan returns the TargetLan field if non-nil, zero value otherwise.

### GetTargetLanOk

`func (o *ApplicationLoadBalancerProperties) GetTargetLanOk() (*int32, bool)`

GetTargetLanOk returns a tuple with the TargetLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLan

`func (o *ApplicationLoadBalancerProperties) SetTargetLan(v int32)`

SetTargetLan sets TargetLan field to given value.


### GetLbPrivateIps

`func (o *ApplicationLoadBalancerProperties) GetLbPrivateIps() []string`

GetLbPrivateIps returns the LbPrivateIps field if non-nil, zero value otherwise.

### GetLbPrivateIpsOk

`func (o *ApplicationLoadBalancerProperties) GetLbPrivateIpsOk() (*[]string, bool)`

GetLbPrivateIpsOk returns a tuple with the LbPrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbPrivateIps

`func (o *ApplicationLoadBalancerProperties) SetLbPrivateIps(v []string)`

SetLbPrivateIps sets LbPrivateIps field to given value.

### HasLbPrivateIps

`func (o *ApplicationLoadBalancerProperties) HasLbPrivateIps() bool`

HasLbPrivateIps returns a boolean if a field has been set.



