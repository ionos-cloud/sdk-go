# KubernetesAutoScaling

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**MinNodeCount** | **int32** | The minimum number of worker nodes that the managed node group can scale in. Should be set together with &#39;maxNodeCount&#39;. Value for this attribute must be greater than equal to 1 and less than equal to maxNodeCount. | |
|**MaxNodeCount** | **int32** | The maximum number of worker nodes that the managed node pool can scale-out. Should be set together with &#39;minNodeCount&#39;. Value for this attribute must be greater than equal to 1 and minNodeCount. | |

## Methods

### NewKubernetesAutoScaling

`func NewKubernetesAutoScaling(minNodeCount int32, maxNodeCount int32, ) *KubernetesAutoScaling`

NewKubernetesAutoScaling instantiates a new KubernetesAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAutoScalingWithDefaults

`func NewKubernetesAutoScalingWithDefaults() *KubernetesAutoScaling`

NewKubernetesAutoScalingWithDefaults instantiates a new KubernetesAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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




