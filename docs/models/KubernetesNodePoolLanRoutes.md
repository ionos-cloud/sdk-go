# KubernetesNodePoolLanRoutes

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Network** | Pointer to **string** | IPv4 or IPv6 CIDR to be routed via the interface. | [optional] |
|**GatewayIp** | Pointer to **string** | IPv4 or IPv6 Gateway IP for the route. | [optional] |

## Methods

### NewKubernetesNodePoolLanRoutes

`func NewKubernetesNodePoolLanRoutes() *KubernetesNodePoolLanRoutes`

NewKubernetesNodePoolLanRoutes instantiates a new KubernetesNodePoolLanRoutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolLanRoutesWithDefaults

`func NewKubernetesNodePoolLanRoutesWithDefaults() *KubernetesNodePoolLanRoutes`

NewKubernetesNodePoolLanRoutesWithDefaults instantiates a new KubernetesNodePoolLanRoutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *KubernetesNodePoolLanRoutes) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *KubernetesNodePoolLanRoutes) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *KubernetesNodePoolLanRoutes) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *KubernetesNodePoolLanRoutes) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetGatewayIp

`func (o *KubernetesNodePoolLanRoutes) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *KubernetesNodePoolLanRoutes) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *KubernetesNodePoolLanRoutes) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *KubernetesNodePoolLanRoutes) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.



