# KubernetesNodePoolPropertiesForPut

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | A Kubernetes node pool name. Valid Kubernetes node pool name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. | [optional] |
|**NodeCount** | **int32** | The number of nodes that make up the node pool. | |
|**K8sVersion** | Pointer to **string** | The Kubernetes version the nodepool is running. This imposes restrictions on what Kubernetes versions can be run in a cluster&#39;s nodepools. Additionally, not all Kubernetes versions are viable upgrade targets for all prior versions. | [optional] |
|**MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](KubernetesMaintenanceWindow.md) |  | [optional] |
|**AutoScaling** | Pointer to [**KubernetesAutoScaling**](KubernetesAutoScaling.md) |  | [optional] |
|**Lans** | Pointer to [**[]KubernetesNodePoolLan**](KubernetesNodePoolLan.md) | array of additional LANs attached to worker nodes | [optional] |
|**Labels** | Pointer to **map[string]string** | map of labels attached to node pool. | [optional] |
|**Annotations** | Pointer to **map[string]string** | map of annotations attached to node pool. | [optional] |
|**PublicIps** | Pointer to **[]string** | Optional array of reserved public IP addresses to be used by the nodes. IPs must be from same location as the data center used for the node pool. The array must contain one more IP than the maximum possible number of nodes (nodeCount+1 for fixed number of nodes or maxNodeCount+1 when auto scaling is used). The extra IP is used when the nodes are rebuilt. | [optional] |

## Methods

### NewKubernetesNodePoolPropertiesForPut

`func NewKubernetesNodePoolPropertiesForPut(nodeCount int32, ) *KubernetesNodePoolPropertiesForPut`

NewKubernetesNodePoolPropertiesForPut instantiates a new KubernetesNodePoolPropertiesForPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolPropertiesForPutWithDefaults

`func NewKubernetesNodePoolPropertiesForPutWithDefaults() *KubernetesNodePoolPropertiesForPut`

NewKubernetesNodePoolPropertiesForPutWithDefaults instantiates a new KubernetesNodePoolPropertiesForPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesNodePoolPropertiesForPut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNodePoolPropertiesForPut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNodePoolPropertiesForPut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesNodePoolPropertiesForPut) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *KubernetesNodePoolPropertiesForPut) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *KubernetesNodePoolPropertiesForPut) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *KubernetesNodePoolPropertiesForPut) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetK8sVersion

`func (o *KubernetesNodePoolPropertiesForPut) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesNodePoolPropertiesForPut) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *KubernetesNodePoolPropertiesForPut) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *KubernetesNodePoolPropertiesForPut) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *KubernetesNodePoolPropertiesForPut) GetMaintenanceWindow() KubernetesMaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *KubernetesNodePoolPropertiesForPut) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *KubernetesNodePoolPropertiesForPut) SetMaintenanceWindow(v KubernetesMaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *KubernetesNodePoolPropertiesForPut) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetAutoScaling

`func (o *KubernetesNodePoolPropertiesForPut) GetAutoScaling() KubernetesAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *KubernetesNodePoolPropertiesForPut) GetAutoScalingOk() (*KubernetesAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *KubernetesNodePoolPropertiesForPut) SetAutoScaling(v KubernetesAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *KubernetesNodePoolPropertiesForPut) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetLans

`func (o *KubernetesNodePoolPropertiesForPut) GetLans() []KubernetesNodePoolLan`

GetLans returns the Lans field if non-nil, zero value otherwise.

### GetLansOk

`func (o *KubernetesNodePoolPropertiesForPut) GetLansOk() (*[]KubernetesNodePoolLan, bool)`

GetLansOk returns a tuple with the Lans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLans

`func (o *KubernetesNodePoolPropertiesForPut) SetLans(v []KubernetesNodePoolLan)`

SetLans sets Lans field to given value.

### HasLans

`func (o *KubernetesNodePoolPropertiesForPut) HasLans() bool`

HasLans returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesNodePoolPropertiesForPut) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesNodePoolPropertiesForPut) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesNodePoolPropertiesForPut) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesNodePoolPropertiesForPut) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *KubernetesNodePoolPropertiesForPut) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesNodePoolPropertiesForPut) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesNodePoolPropertiesForPut) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesNodePoolPropertiesForPut) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetPublicIps

`func (o *KubernetesNodePoolPropertiesForPut) GetPublicIps() []string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *KubernetesNodePoolPropertiesForPut) GetPublicIpsOk() (*[]string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *KubernetesNodePoolPropertiesForPut) SetPublicIps(v []string)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *KubernetesNodePoolPropertiesForPut) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.



