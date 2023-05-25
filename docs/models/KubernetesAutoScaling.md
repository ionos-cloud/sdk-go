# KubernetesAutoScaling

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**MaxNodeCount** | **int32** | The maximum number of worker nodes that the managed node pool can scale in. Must be &gt;&#x3D; minNodeCount and must be &gt;&#x3D; nodeCount. Required if autoScaling is specified. | |
|**MinNodeCount** | **int32** | The minimum number of working nodes that the managed node pool can scale must be &gt;&#x3D; 1 and &gt;&#x3D; nodeCount. Required if autoScaling is specified. | |

## Methods

### NewKubernetesAutoScaling

`func NewKubernetesAutoScaling(maxNodeCount int32, minNodeCount int32, ) *KubernetesAutoScaling`

NewKubernetesAutoScaling instantiates a new KubernetesAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAutoScalingWithDefaults

`func NewKubernetesAutoScalingWithDefaults() *KubernetesAutoScaling`

NewKubernetesAutoScalingWithDefaults instantiates a new KubernetesAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNodeCount

`func (o *KubernetesAutoScaling) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *KubernetesAutoScaling) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *KubernetesAutoScaling) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.


### GetMinNodeCount

`func (o *KubernetesAutoScaling) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *KubernetesAutoScaling) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *KubernetesAutoScaling) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.




