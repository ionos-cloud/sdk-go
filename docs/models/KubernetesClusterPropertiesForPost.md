# KubernetesClusterPropertiesForPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | A Kubernetes cluster name. Valid Kubernetes cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. | |
|**K8sVersion** | Pointer to **string** | The Kubernetes version the cluster is running. This imposes restrictions on what Kubernetes versions can be run in a cluster&#39;s nodepools. Additionally, not all Kubernetes versions are viable upgrade targets for all prior versions. | [optional] |
|**MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](KubernetesMaintenanceWindow.md) |  | [optional] |
|**ApiSubnetAllowList** | Pointer to **[]string** | Access to the K8s API server is restricted to these CIDRs. Traffic, internal to the cluster, is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value is used: 32 for IPv4 and 128 for IPv6. | [optional] |
|**S3Buckets** | Pointer to [**[]S3Bucket**](S3Bucket.md) | List of S3 bucket configured for K8s usage. For now it contains only an S3 bucket used to store K8s API audit logs | [optional] |

## Methods

### NewKubernetesClusterPropertiesForPost

`func NewKubernetesClusterPropertiesForPost(name string, ) *KubernetesClusterPropertiesForPost`

NewKubernetesClusterPropertiesForPost instantiates a new KubernetesClusterPropertiesForPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClusterPropertiesForPostWithDefaults

`func NewKubernetesClusterPropertiesForPostWithDefaults() *KubernetesClusterPropertiesForPost`

NewKubernetesClusterPropertiesForPostWithDefaults instantiates a new KubernetesClusterPropertiesForPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesClusterPropertiesForPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesClusterPropertiesForPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesClusterPropertiesForPost) SetName(v string)`

SetName sets Name field to given value.


### GetK8sVersion

`func (o *KubernetesClusterPropertiesForPost) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesClusterPropertiesForPost) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *KubernetesClusterPropertiesForPost) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *KubernetesClusterPropertiesForPost) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *KubernetesClusterPropertiesForPost) GetMaintenanceWindow() KubernetesMaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *KubernetesClusterPropertiesForPost) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *KubernetesClusterPropertiesForPost) SetMaintenanceWindow(v KubernetesMaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *KubernetesClusterPropertiesForPost) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetApiSubnetAllowList

`func (o *KubernetesClusterPropertiesForPost) GetApiSubnetAllowList() []string`

GetApiSubnetAllowList returns the ApiSubnetAllowList field if non-nil, zero value otherwise.

### GetApiSubnetAllowListOk

`func (o *KubernetesClusterPropertiesForPost) GetApiSubnetAllowListOk() (*[]string, bool)`

GetApiSubnetAllowListOk returns a tuple with the ApiSubnetAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSubnetAllowList

`func (o *KubernetesClusterPropertiesForPost) SetApiSubnetAllowList(v []string)`

SetApiSubnetAllowList sets ApiSubnetAllowList field to given value.

### HasApiSubnetAllowList

`func (o *KubernetesClusterPropertiesForPost) HasApiSubnetAllowList() bool`

HasApiSubnetAllowList returns a boolean if a field has been set.

### GetS3Buckets

`func (o *KubernetesClusterPropertiesForPost) GetS3Buckets() []S3Bucket`

GetS3Buckets returns the S3Buckets field if non-nil, zero value otherwise.

### GetS3BucketsOk

`func (o *KubernetesClusterPropertiesForPost) GetS3BucketsOk() (*[]S3Bucket, bool)`

GetS3BucketsOk returns a tuple with the S3Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Buckets

`func (o *KubernetesClusterPropertiesForPost) SetS3Buckets(v []S3Bucket)`

SetS3Buckets sets S3Buckets field to given value.

### HasS3Buckets

`func (o *KubernetesClusterPropertiesForPost) HasS3Buckets() bool`

HasS3Buckets returns a boolean if a field has been set.



