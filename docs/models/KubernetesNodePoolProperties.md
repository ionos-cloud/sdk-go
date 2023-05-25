# KubernetesNodePoolProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Annotations** | Pointer to **map[string]string** | The annotations attached to the node pool. | [optional] |
|**AutoScaling** | Pointer to [**KubernetesAutoScaling**](KubernetesAutoScaling.md) |  | [optional] |
|**AvailabilityZone** | **string** | The availability zone in which the target VM should be provisioned. | |
|**AvailableUpgradeVersions** | Pointer to **[]string** | The list of available versions for upgrading the node pool. | [optional] |
|**CoresCount** | **int32** | The total number of cores for the nodes. | |
|**CpuFamily** | **string** | The CPU type for the nodes. | |
|**DatacenterId** | **string** | The unique identifier of the VDC where the worker nodes of the node pool are provisioned.Note that the data center is located in the exact place where the parent cluster of the node pool is located. | |
|**K8sVersion** | Pointer to **string** | The Kubernetes version running in the node pool. Note that this imposes restrictions on which Kubernetes versions can run in the node pools of a cluster. Also, not all Kubernetes versions are suitable upgrade targets for all earlier versions. | [optional] |
|**Labels** | Pointer to **map[string]string** | The labels attached to the node pool. | [optional] |
|**Lans** | Pointer to [**[]KubernetesNodePoolLan**](KubernetesNodePoolLan.md) | The array of existing private LANs to attach to worker nodes. | [optional] |
|**MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](KubernetesMaintenanceWindow.md) |  | [optional] |
|**Name** | **string** | A Kubernetes node pool name. Valid Kubernetes node pool name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. | |
|**NodeCount** | **int32** | The number of worker nodes of the node pool. | |
|**PublicIps** | Pointer to **[]string** | Optional array of reserved public IP addresses to be used by the nodes. The IPs must be from the exact location of the node pool&#39;s data center. If autoscaling is used, the array must contain one more IP than the maximum possible number of nodes (nodeCount+1 for a fixed number of nodes or maxNodeCount+1). The extra IP is used when the nodes are rebuilt. | [optional] |
|**RamSize** | **int32** | The RAM size for the nodes. Must be specified in multiples of 1024 MB, with a minimum size of 2048 MB. | |
|**StorageSize** | **int32** | The allocated volume size in GB. The allocated volume size in GB. To achieve good performance, we recommend a size greater than 100GB for SSD. | |
|**StorageType** | **string** | The storage type for the nodes. | |

## Methods

### NewKubernetesNodePoolProperties

`func NewKubernetesNodePoolProperties(availabilityZone string, coresCount int32, cpuFamily string, datacenterId string, name string, nodeCount int32, ramSize int32, storageSize int32, storageType string, ) *KubernetesNodePoolProperties`

NewKubernetesNodePoolProperties instantiates a new KubernetesNodePoolProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolPropertiesWithDefaults

`func NewKubernetesNodePoolPropertiesWithDefaults() *KubernetesNodePoolProperties`

NewKubernetesNodePoolPropertiesWithDefaults instantiates a new KubernetesNodePoolProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *KubernetesNodePoolProperties) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesNodePoolProperties) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesNodePoolProperties) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesNodePoolProperties) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAutoScaling

`func (o *KubernetesNodePoolProperties) GetAutoScaling() KubernetesAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *KubernetesNodePoolProperties) GetAutoScalingOk() (*KubernetesAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *KubernetesNodePoolProperties) SetAutoScaling(v KubernetesAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *KubernetesNodePoolProperties) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *KubernetesNodePoolProperties) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *KubernetesNodePoolProperties) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *KubernetesNodePoolProperties) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetAvailableUpgradeVersions

`func (o *KubernetesNodePoolProperties) GetAvailableUpgradeVersions() []string`

GetAvailableUpgradeVersions returns the AvailableUpgradeVersions field if non-nil, zero value otherwise.

### GetAvailableUpgradeVersionsOk

`func (o *KubernetesNodePoolProperties) GetAvailableUpgradeVersionsOk() (*[]string, bool)`

GetAvailableUpgradeVersionsOk returns a tuple with the AvailableUpgradeVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUpgradeVersions

`func (o *KubernetesNodePoolProperties) SetAvailableUpgradeVersions(v []string)`

SetAvailableUpgradeVersions sets AvailableUpgradeVersions field to given value.

### HasAvailableUpgradeVersions

`func (o *KubernetesNodePoolProperties) HasAvailableUpgradeVersions() bool`

HasAvailableUpgradeVersions returns a boolean if a field has been set.

### GetCoresCount

`func (o *KubernetesNodePoolProperties) GetCoresCount() int32`

GetCoresCount returns the CoresCount field if non-nil, zero value otherwise.

### GetCoresCountOk

`func (o *KubernetesNodePoolProperties) GetCoresCountOk() (*int32, bool)`

GetCoresCountOk returns a tuple with the CoresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresCount

`func (o *KubernetesNodePoolProperties) SetCoresCount(v int32)`

SetCoresCount sets CoresCount field to given value.


### GetCpuFamily

`func (o *KubernetesNodePoolProperties) GetCpuFamily() string`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *KubernetesNodePoolProperties) GetCpuFamilyOk() (*string, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *KubernetesNodePoolProperties) SetCpuFamily(v string)`

SetCpuFamily sets CpuFamily field to given value.


### GetDatacenterId

`func (o *KubernetesNodePoolProperties) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *KubernetesNodePoolProperties) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *KubernetesNodePoolProperties) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetK8sVersion

`func (o *KubernetesNodePoolProperties) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesNodePoolProperties) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *KubernetesNodePoolProperties) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *KubernetesNodePoolProperties) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetLabels

`func (o *KubernetesNodePoolProperties) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesNodePoolProperties) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesNodePoolProperties) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesNodePoolProperties) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLans

`func (o *KubernetesNodePoolProperties) GetLans() []KubernetesNodePoolLan`

GetLans returns the Lans field if non-nil, zero value otherwise.

### GetLansOk

`func (o *KubernetesNodePoolProperties) GetLansOk() (*[]KubernetesNodePoolLan, bool)`

GetLansOk returns a tuple with the Lans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLans

`func (o *KubernetesNodePoolProperties) SetLans(v []KubernetesNodePoolLan)`

SetLans sets Lans field to given value.

### HasLans

`func (o *KubernetesNodePoolProperties) HasLans() bool`

HasLans returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *KubernetesNodePoolProperties) GetMaintenanceWindow() KubernetesMaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *KubernetesNodePoolProperties) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *KubernetesNodePoolProperties) SetMaintenanceWindow(v KubernetesMaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *KubernetesNodePoolProperties) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetName

`func (o *KubernetesNodePoolProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNodePoolProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNodePoolProperties) SetName(v string)`

SetName sets Name field to given value.


### GetNodeCount

`func (o *KubernetesNodePoolProperties) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *KubernetesNodePoolProperties) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *KubernetesNodePoolProperties) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetPublicIps

`func (o *KubernetesNodePoolProperties) GetPublicIps() []string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *KubernetesNodePoolProperties) GetPublicIpsOk() (*[]string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *KubernetesNodePoolProperties) SetPublicIps(v []string)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *KubernetesNodePoolProperties) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### GetRamSize

`func (o *KubernetesNodePoolProperties) GetRamSize() int32`

GetRamSize returns the RamSize field if non-nil, zero value otherwise.

### GetRamSizeOk

`func (o *KubernetesNodePoolProperties) GetRamSizeOk() (*int32, bool)`

GetRamSizeOk returns a tuple with the RamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamSize

`func (o *KubernetesNodePoolProperties) SetRamSize(v int32)`

SetRamSize sets RamSize field to given value.


### GetStorageSize

`func (o *KubernetesNodePoolProperties) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *KubernetesNodePoolProperties) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *KubernetesNodePoolProperties) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.


### GetStorageType

`func (o *KubernetesNodePoolProperties) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *KubernetesNodePoolProperties) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *KubernetesNodePoolProperties) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.




