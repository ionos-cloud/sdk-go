# KubernetesClusterPropertiesForPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ApiSubnetAllowList** | Pointer to **[]string** | Access to the K8s API server is restricted to these CIDRs. Intra-cluster traffic is not affected by this restriction. If no AllowList is specified, access is not limited. If an IP is specified without a subnet mask, the default value is 32 for IPv4 and 128 for IPv6. | [optional] |
|**K8sVersion** | Pointer to **string** | The Kubernetes version that the cluster is running. This limits which Kubernetes versions can run in a cluster&#39;s node pools. Also, not all Kubernetes versions are suitable upgrade targets for all earlier versions. | [optional] |
|**Location** | Pointer to **string** | This attribute is mandatory if the cluster is private. The location must be enabled for your contract, or you must have a data center at that location. This property is not adjustable. | [optional] |
|**MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](KubernetesMaintenanceWindow.md) |  | [optional] |
|**Name** | **string** | A Kubernetes cluster name. Valid Kubernetes cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. | |
|**NatGatewayIp** | Pointer to **string** | The nat gateway IP of the cluster if the cluster is private. This property is immutable. Must be a reserved IP in the same location as the cluster&#39;s location. This attribute is mandatory if the cluster is private. | [optional] |
|**NodeSubnet** | Pointer to **string** | The node subnet of the cluster, if the cluster is private. This property is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length. | [optional] |
|**Public** | Pointer to **bool** | The indicator whether the cluster is public or private. Note that the status FALSE is still in the beta phase. | [optional] [default to true]|
|**S3Buckets** | Pointer to [**[]S3Bucket**](S3Bucket.md) | List of S3 buckets configured for K8s usage. At the moment, it contains only one S3 bucket that is used to store K8s API audit logs. | [optional] |

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

### GetLocation

`func (o *KubernetesClusterPropertiesForPost) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KubernetesClusterPropertiesForPost) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KubernetesClusterPropertiesForPost) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KubernetesClusterPropertiesForPost) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

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


### GetNatGatewayIp

`func (o *KubernetesClusterPropertiesForPost) GetNatGatewayIp() string`

GetNatGatewayIp returns the NatGatewayIp field if non-nil, zero value otherwise.

### GetNatGatewayIpOk

`func (o *KubernetesClusterPropertiesForPost) GetNatGatewayIpOk() (*string, bool)`

GetNatGatewayIpOk returns a tuple with the NatGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatGatewayIp

`func (o *KubernetesClusterPropertiesForPost) SetNatGatewayIp(v string)`

SetNatGatewayIp sets NatGatewayIp field to given value.

### HasNatGatewayIp

`func (o *KubernetesClusterPropertiesForPost) HasNatGatewayIp() bool`

HasNatGatewayIp returns a boolean if a field has been set.

### GetNodeSubnet

`func (o *KubernetesClusterPropertiesForPost) GetNodeSubnet() string`

GetNodeSubnet returns the NodeSubnet field if non-nil, zero value otherwise.

### GetNodeSubnetOk

`func (o *KubernetesClusterPropertiesForPost) GetNodeSubnetOk() (*string, bool)`

GetNodeSubnetOk returns a tuple with the NodeSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSubnet

`func (o *KubernetesClusterPropertiesForPost) SetNodeSubnet(v string)`

SetNodeSubnet sets NodeSubnet field to given value.

### HasNodeSubnet

`func (o *KubernetesClusterPropertiesForPost) HasNodeSubnet() bool`

HasNodeSubnet returns a boolean if a field has been set.

### GetPublic

`func (o *KubernetesClusterPropertiesForPost) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *KubernetesClusterPropertiesForPost) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *KubernetesClusterPropertiesForPost) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *KubernetesClusterPropertiesForPost) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

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



