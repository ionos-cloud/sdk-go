# KubernetesNodePoolLan

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **int32** | The LAN ID of an existing LAN at the related datacenter | |
|**Dhcp** | Pointer to **bool** | Indicates if the Kubernetes node pool LAN will reserve an IP using DHCP. | [optional] |
|**Routes** | Pointer to [**[]KubernetesNodePoolLanRoutes**](KubernetesNodePoolLanRoutes.md) | array of additional LANs attached to worker nodes | [optional] |

## Methods

### NewKubernetesNodePoolLan

`func NewKubernetesNodePoolLan(id int32, ) *KubernetesNodePoolLan`

NewKubernetesNodePoolLan instantiates a new KubernetesNodePoolLan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolLanWithDefaults

`func NewKubernetesNodePoolLanWithDefaults() *KubernetesNodePoolLan`

NewKubernetesNodePoolLanWithDefaults instantiates a new KubernetesNodePoolLan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesNodePoolLan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesNodePoolLan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesNodePoolLan) SetId(v int32)`

SetId sets Id field to given value.


### GetDhcp

`func (o *KubernetesNodePoolLan) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *KubernetesNodePoolLan) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *KubernetesNodePoolLan) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *KubernetesNodePoolLan) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetRoutes

`func (o *KubernetesNodePoolLan) GetRoutes() []KubernetesNodePoolLanRoutes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *KubernetesNodePoolLan) GetRoutesOk() (*[]KubernetesNodePoolLanRoutes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *KubernetesNodePoolLan) SetRoutes(v []KubernetesNodePoolLanRoutes)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *KubernetesNodePoolLan) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.



