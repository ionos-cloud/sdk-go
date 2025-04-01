# KubernetesNodePoolPropertiesForPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | A Kubernetes node pool name. Valid Kubernetes node pool name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. | |
|**DatacenterId** | **string** | The unique identifier of the VDC where the worker nodes of the node pool are provisioned.Note that the data center is located in the exact place where the parent cluster of the node pool is located. | |
|**NodeCount** | **int32** | The number of worker nodes of the node pool. | |
|**CpuFamily** | Pointer to **string** | The CPU type for the nodes. | [optional] |
|**ServerType** | Pointer to [**KubernetesNodePoolServerType**](KubernetesNodePoolServerType.md) |  | [optional] [default to DEDICATED_CORE]|
|**CoresCount** | **int32** | The total number of cores for the nodes. | |
|**RamSize** | **int32** | The RAM size for the nodes. Must be specified in multiples of 1024 MB, with a minimum size of 2048 MB. | |
|**AvailabilityZone** | **string** | The availability zone in which the target VM should be provisioned. | |
|**StorageType** | **string** | The storage type for the nodes. | |
|**StorageSize** | **int32** | The allocated volume size in GB. The allocated volume size in GB. To achieve good performance, we recommend a size greater than 100GB for SSD. | |
|**K8sVersion** | Pointer to **string** | The Kubernetes version running in the node pool. Note that this imposes restrictions on which Kubernetes versions can run in the node pools of a cluster. Also, not all Kubernetes versions are suitable upgrade targets for all earlier versions. | [optional] |
|**MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](KubernetesMaintenanceWindow.md) |  | [optional] |
|**AutoScaling** | Pointer to [**KubernetesAutoScaling**](KubernetesAutoScaling.md) |  | [optional] |
|**Lans** | Pointer to [**[]KubernetesNodePoolLan**](KubernetesNodePoolLan.md) | The array of existing private LANs to attach to worker nodes. | [optional] |
|**Labels** | Pointer to **map[string]string** | The labels attached to the node pool. | [optional] |
|**Annotations** | Pointer to **map[string]string** | The annotations attached to the node pool. | [optional] |
|**PublicIps** | Pointer to **[]string** | Optional array of reserved public IP addresses to be used by the nodes. The IPs must be from the exact location of the node pool&#39;s data center. If autoscaling is used, the array must contain one more IP than the maximum possible number of nodes (nodeCount+1 for a fixed number of nodes or maxNodeCount+1). The extra IP is used when the nodes are rebuilt. | [optional] |

## Methods

### NewKubernetesNodePoolPropertiesForPost

`func NewKubernetesNodePoolPropertiesForPost(name string, datacenterId string, nodeCount int32, coresCount int32, ramSize int32, availabilityZone string, storageType string, storageSize int32, ) *KubernetesNodePoolPropertiesForPost`

NewKubernetesNodePoolPropertiesForPost instantiates a new KubernetesNodePoolPropertiesForPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolPropertiesForPostWithDefaults

`func NewKubernetesNodePoolPropertiesForPostWithDefaults() *KubernetesNodePoolPropertiesForPost`

NewKubernetesNodePoolPropertiesForPostWithDefaults instantiates a new KubernetesNodePoolPropertiesForPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesNodePoolPropertiesForPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNodePoolPropertiesForPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNodePoolPropertiesForPost) SetName(v string)`

SetName sets Name field to given value.


### GetDatacenterId

`func (o *KubernetesNodePoolPropertiesForPost) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *KubernetesNodePoolPropertiesForPost) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *KubernetesNodePoolPropertiesForPost) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetNodeCount

`func (o *KubernetesNodePoolPropertiesForPost) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *KubernetesNodePoolPropertiesForPost) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *KubernetesNodePoolPropertiesForPost) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetCpuFamily

`func (o *KubernetesNodePoolPropertiesForPost) GetCpuFamily() string`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *KubernetesNodePoolPropertiesForPost) GetCpuFamilyOk() (*string, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *KubernetesNodePoolPropertiesForPost) SetCpuFamily(v string)`

SetCpuFamily sets CpuFamily field to given value.

### HasCpuFamily

`func (o *KubernetesNodePoolPropertiesForPost) HasCpuFamily() bool`

HasCpuFamily returns a boolean if a field has been set.

### GetServerType

`func (o *KubernetesNodePoolPropertiesForPost) GetServerType() KubernetesNodePoolServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *KubernetesNodePoolPropertiesForPost) GetServerTypeOk() (*KubernetesNodePoolServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *KubernetesNodePoolPropertiesForPost) SetServerType(v KubernetesNodePoolServerType)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *KubernetesNodePoolPropertiesForPost) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetCoresCount

`func (o *KubernetesNodePoolPropertiesForPost) GetCoresCount() int32`

GetCoresCount returns the CoresCount field if non-nil, zero value otherwise.

### GetCoresCountOk

`func (o *KubernetesNodePoolPropertiesForPost) GetCoresCountOk() (*int32, bool)`

GetCoresCountOk returns a tuple with the CoresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresCount

`func (o *KubernetesNodePoolPropertiesForPost) SetCoresCount(v int32)`

SetCoresCount sets CoresCount field to given value.


### GetRamSize

`func (o *KubernetesNodePoolPropertiesForPost) GetRamSize() int32`

GetRamSize returns the RamSize field if non-nil, zero value otherwise.

### GetRamSizeOk

`func (o *KubernetesNodePoolPropertiesForPost) GetRamSizeOk() (*int32, bool)`

GetRamSizeOk returns a tuple with the RamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamSize

`func (o *KubernetesNodePoolPropertiesForPost) SetRamSize(v int32)`

SetRamSize sets RamSize field to given value.


### GetAvailabilityZone

`func (o *KubernetesNodePoolPropertiesForPost) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *KubernetesNodePoolPropertiesForPost) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *KubernetesNodePoolPropertiesForPost) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetStorageType

`func (o *KubernetesNodePoolPropertiesForPost) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *KubernetesNodePoolPropertiesForPost) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *KubernetesNodePoolPropertiesForPost) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### GetStorageSize

`func (o *KubernetesNodePoolPropertiesForPost) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *KubernetesNodePoolPropertiesForPost) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *KubernetesNodePoolPropertiesForPost) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.


### GetK8sVersion

`func (o *KubernetesNodePoolPropertiesForPost) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesNodePoolPropertiesForPost) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *KubernetesNodePoolPropertiesForPost) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *KubernetesNodePoolPropertiesForPost) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *KubernetesNodePoolPropertiesForPost) GetMaintenanceWindow() KubernetesMaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *KubernetesNodePoolPropertiesForPost) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *KubernetesNodePoolPropertiesForPost) SetMaintenanceWindow(v KubernetesMaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *KubernetesNodePoolPropertiesForPost) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetAutoScaling

`func (o *KubernetesNodePoolPropertiesForPost) GetAutoScaling() KubernetesAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *KubernetesNodePoolPropertiesForPost) GetAutoScalingOk() (*KubernetesAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *KubernetesNodePoolPropertiesForPost) SetAutoScaling(v KubernetesAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *KubernetesNodePoolPropertiesForPost) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetLans

`func (o *KubernetesNodePoolPropertiesForPost) GetLans() []KubernetesNodePoolLan`

GetLans returns the Lans field if non-nil, zero value otherwise.

### GetLansOk

`func (o *KubernetesNodePoolPropertiesForPost) GetLansOk() (*[]KubernetesNodePoolLan, bool)`

GetLansOk returns a tuple with the Lans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLans

`func (o *KubernetesNodePoolPropertiesForPost) SetLans(v []KubernetesNodePoolLan)`

SetLans sets Lans field to given value.

### HasLans

`func (o *KubernetesNodePoolPropertiesForPost) HasLans() bool`

HasLans returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesNodePoolPropertiesForPost) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesNodePoolPropertiesForPost) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesNodePoolPropertiesForPost) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesNodePoolPropertiesForPost) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *KubernetesNodePoolPropertiesForPost) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesNodePoolPropertiesForPost) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesNodePoolPropertiesForPost) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesNodePoolPropertiesForPost) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetPublicIps

`func (o *KubernetesNodePoolPropertiesForPost) GetPublicIps() []string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *KubernetesNodePoolPropertiesForPost) GetPublicIpsOk() (*[]string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *KubernetesNodePoolPropertiesForPost) SetPublicIps(v []string)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *KubernetesNodePoolPropertiesForPost) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.



