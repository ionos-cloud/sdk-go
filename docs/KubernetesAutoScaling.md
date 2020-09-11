# KubernetesAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinNodeCount** | **int32** | The minimum number of worker nodes that the managed node group can scale in. Should be set together with &#39;maxNodeCount&#39;. Value for this attribute must be greater than equal to 1 and less than equal to maxNodeCount. | [optional] 
**MaxNodeCount** | **int32** | The maximum number of worker nodes that the managed node pool can scale-out. Should be set together with &#39;minNodeCount&#39;. Value for this attribute must be greater than equal to 1 and minNodeCount. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


