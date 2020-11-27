# KubernetesNodeProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A Kubernetes Node Name. | 
**PublicIP** | **string** | A valid public IP. | 
**K8sVersion** | **string** | The kubernetes version in which a nodepool is running. This imposes restrictions on what kubernetes versions can be run in a cluster&#39;s nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions. | 

## Methods

### NewKubernetesNodeProperties

`func NewKubernetesNodeProperties(name string, publicIP string, k8sVersion string, ) *KubernetesNodeProperties`

NewKubernetesNodeProperties instantiates a new KubernetesNodeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePropertiesWithDefaults

`func NewKubernetesNodePropertiesWithDefaults() *KubernetesNodeProperties`

NewKubernetesNodePropertiesWithDefaults instantiates a new KubernetesNodeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesNodeProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNodeProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNodeProperties) SetName(v string)`

SetName sets Name field to given value.


### GetPublicIP

`func (o *KubernetesNodeProperties) GetPublicIP() string`

GetPublicIP returns the PublicIP field if non-nil, zero value otherwise.

### GetPublicIPOk

`func (o *KubernetesNodeProperties) GetPublicIPOk() (*string, bool)`

GetPublicIPOk returns a tuple with the PublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIP

`func (o *KubernetesNodeProperties) SetPublicIP(v string)`

SetPublicIP sets PublicIP field to given value.


### GetK8sVersion

`func (o *KubernetesNodeProperties) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesNodeProperties) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *KubernetesNodeProperties) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


