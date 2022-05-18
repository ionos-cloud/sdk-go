# KubernetesClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | A Kubernetes cluster name. Valid Kubernetes cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. | |
|**K8sVersion** | Pointer to **string** | The Kubernetes version the cluster is running. This imposes restrictions on what Kubernetes versions can be run in a cluster&#39;s nodepools. Additionally, not all Kubernetes versions are viable upgrade targets for all prior versions. | [optional] |
|**MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](KubernetesMaintenanceWindow.md) |  | [optional] |
|**AvailableUpgradeVersions** | Pointer to **[]string** | List of available versions for upgrading the cluster | [optional] |
|**ViableNodePoolVersions** | Pointer to **[]string** | List of versions that may be used for node pools under this cluster | [optional] |
|**ApiSubnetAllowList** | Pointer to **[]string** | Access to the K8s API server is restricted to these CIDRs. Traffic, internal to the cluster, is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value is used: 32 for IPv4 and 128 for IPv6. | [optional] |
|**S3Buckets** | Pointer to [**[]S3Bucket**](S3Bucket.md) | List of S3 bucket configured for K8s usage. For now it contains only an S3 bucket used to store K8s API audit logs | [optional] |

## Methods

### NewKubernetesClusterProperties

`func NewKubernetesClusterProperties(name string, ) *KubernetesClusterProperties`

NewKubernetesClusterProperties instantiates a new KubernetesClusterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterPropertiesWithDefaults

`func NewKubernetesClusterPropertiesWithDefaults() *KubernetesClusterProperties`

NewKubernetesClusterPropertiesWithDefaults instantiates a new KubernetesClusterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesClusterProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesClusterProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesClusterProperties) SetName(v string)`

SetName sets Name field to given value.


### GetK8sVersion

`func (o *KubernetesClusterProperties) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesClusterProperties) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *KubernetesClusterProperties) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *KubernetesClusterProperties) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *KubernetesClusterProperties) GetMaintenanceWindow() KubernetesMaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *KubernetesClusterProperties) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *KubernetesClusterProperties) SetMaintenanceWindow(v KubernetesMaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *KubernetesClusterProperties) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetAvailableUpgradeVersions

`func (o *KubernetesClusterProperties) GetAvailableUpgradeVersions() []string`

GetAvailableUpgradeVersions returns the AvailableUpgradeVersions field if non-nil, zero value otherwise.

### GetAvailableUpgradeVersionsOk

`func (o *KubernetesClusterProperties) GetAvailableUpgradeVersionsOk() (*[]string, bool)`

GetAvailableUpgradeVersionsOk returns a tuple with the AvailableUpgradeVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUpgradeVersions

`func (o *KubernetesClusterProperties) SetAvailableUpgradeVersions(v []string)`

SetAvailableUpgradeVersions sets AvailableUpgradeVersions field to given value.

### HasAvailableUpgradeVersions

`func (o *KubernetesClusterProperties) HasAvailableUpgradeVersions() bool`

HasAvailableUpgradeVersions returns a boolean if a field has been set.

### GetViableNodePoolVersions

`func (o *KubernetesClusterProperties) GetViableNodePoolVersions() []string`

GetViableNodePoolVersions returns the ViableNodePoolVersions field if non-nil, zero value otherwise.

### GetViableNodePoolVersionsOk

`func (o *KubernetesClusterProperties) GetViableNodePoolVersionsOk() (*[]string, bool)`

GetViableNodePoolVersionsOk returns a tuple with the ViableNodePoolVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViableNodePoolVersions

`func (o *KubernetesClusterProperties) SetViableNodePoolVersions(v []string)`

SetViableNodePoolVersions sets ViableNodePoolVersions field to given value.

### HasViableNodePoolVersions

`func (o *KubernetesClusterProperties) HasViableNodePoolVersions() bool`

HasViableNodePoolVersions returns a boolean if a field has been set.

### GetApiSubnetAllowList

`func (o *KubernetesClusterProperties) GetApiSubnetAllowList() []string`

GetApiSubnetAllowList returns the ApiSubnetAllowList field if non-nil, zero value otherwise.

### GetApiSubnetAllowListOk

`func (o *KubernetesClusterProperties) GetApiSubnetAllowListOk() (*[]string, bool)`

GetApiSubnetAllowListOk returns a tuple with the ApiSubnetAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSubnetAllowList

`func (o *KubernetesClusterProperties) SetApiSubnetAllowList(v []string)`

SetApiSubnetAllowList sets ApiSubnetAllowList field to given value.

### HasApiSubnetAllowList

`func (o *KubernetesClusterProperties) HasApiSubnetAllowList() bool`

HasApiSubnetAllowList returns a boolean if a field has been set.

### GetS3Buckets

`func (o *KubernetesClusterProperties) GetS3Buckets() []S3Bucket`

GetS3Buckets returns the S3Buckets field if non-nil, zero value otherwise.

### GetS3BucketsOk

`func (o *KubernetesClusterProperties) GetS3BucketsOk() (*[]S3Bucket, bool)`

GetS3BucketsOk returns a tuple with the S3Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Buckets

`func (o *KubernetesClusterProperties) SetS3Buckets(v []S3Bucket)`

SetS3Buckets sets S3Buckets field to given value.

### HasS3Buckets

`func (o *KubernetesClusterProperties) HasS3Buckets() bool`

HasS3Buckets returns a boolean if a field has been set.



