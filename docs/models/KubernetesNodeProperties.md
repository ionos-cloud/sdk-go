# KubernetesNodeProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**K8sVersion** | **string** | The Kubernetes version running in the node pool. Note that this imposes restrictions on which Kubernetes versions can run in the node pools of a cluster. Also, not all Kubernetes versions are suitable upgrade targets for all earlier versions. | |
|**Name** | **string** | The Kubernetes node name. | |
|**PrivateIP** | Pointer to **string** | The private IP associated with the node. | [optional] |
|**PublicIP** | Pointer to **string** | The public IP associated with the node. | [optional] |

## Methods

### NewKubernetesNodeProperties

`func NewKubernetesNodeProperties(k8sVersion string, name string, ) *KubernetesNodeProperties`

NewKubernetesNodeProperties instantiates a new KubernetesNodeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePropertiesWithDefaults

`func NewKubernetesNodePropertiesWithDefaults() *KubernetesNodeProperties`

NewKubernetesNodePropertiesWithDefaults instantiates a new KubernetesNodeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetPrivateIP

`func (o *KubernetesNodeProperties) GetPrivateIP() string`

GetPrivateIP returns the PrivateIP field if non-nil, zero value otherwise.

### GetPrivateIPOk

`func (o *KubernetesNodeProperties) GetPrivateIPOk() (*string, bool)`

GetPrivateIPOk returns a tuple with the PrivateIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIP

`func (o *KubernetesNodeProperties) SetPrivateIP(v string)`

SetPrivateIP sets PrivateIP field to given value.

### HasPrivateIP

`func (o *KubernetesNodeProperties) HasPrivateIP() bool`

HasPrivateIP returns a boolean if a field has been set.

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

### HasPublicIP

`func (o *KubernetesNodeProperties) HasPublicIP() bool`

HasPublicIP returns a boolean if a field has been set.



