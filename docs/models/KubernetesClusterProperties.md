# KubernetesClusterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | A Kubernetes cluster name. Valid Kubernetes cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. | |
|**K8sVersion** | Pointer to **string** | The Kubernetes version the cluster is running. This imposes restrictions on what Kubernetes versions can be run in a cluster&#39;s nodepools. Additionally, not all Kubernetes versions are viable upgrade targets for all prior versions. | [optional] |
|**MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](KubernetesMaintenanceWindow.md) |  | [optional] |
|**AvailableUpgradeVersions** | Pointer to **[]string** | List of available versions for upgrading the cluster | [optional] |
|**ViableNodePoolVersions** | Pointer to **[]string** | List of versions that may be used for node pools under this cluster | [optional] |
|**Public** | Pointer to **bool** | The indicator if the cluster is public or private. Be aware that setting it to false is currently in beta phase. | [optional] [default to true]|
|**Location** | Pointer to **string** | The location of the cluster if the cluster is private. This property is immutable. The location must be enabled for your contract or you must have a Datacenter within that location. This attribute is mandatory if the cluster is private. | [optional] |
|**NatGatewayIp** | Pointer to **string** | The nat gateway IP of the cluster if the cluster is private. This property is immutable. Must be a reserved IP in the same location as the cluster&#39;s location. This attribute is mandatory if the cluster is private. | [optional] |
|**NodeSubnet** | Pointer to **string** | The node subnet of the cluster, if the cluster is private. This property is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length. | [optional] |
|**ApiSubnetAllowList** | Pointer to **[]string** | Access to the K8s API server is restricted to these CIDRs. Traffic, internal to the cluster, is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value is used: 32 for IPv4 and 128 for IPv6. | [optional] |
|**S3Buckets** | Pointer to [**[]S3Bucket**](S3Bucket.md) | List of Object storage buckets configured for K8s usage. For now it contains only one bucket used to store K8s API audit logs | [optional] |

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

### GetPublic

`func (o *KubernetesClusterProperties) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *KubernetesClusterProperties) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *KubernetesClusterProperties) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *KubernetesClusterProperties) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetLocation

`func (o *KubernetesClusterProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KubernetesClusterProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KubernetesClusterProperties) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KubernetesClusterProperties) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetNatGatewayIp

`func (o *KubernetesClusterProperties) GetNatGatewayIp() string`

GetNatGatewayIp returns the NatGatewayIp field if non-nil, zero value otherwise.

### GetNatGatewayIpOk

`func (o *KubernetesClusterProperties) GetNatGatewayIpOk() (*string, bool)`

GetNatGatewayIpOk returns a tuple with the NatGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatGatewayIp

`func (o *KubernetesClusterProperties) SetNatGatewayIp(v string)`

SetNatGatewayIp sets NatGatewayIp field to given value.

### HasNatGatewayIp

`func (o *KubernetesClusterProperties) HasNatGatewayIp() bool`

HasNatGatewayIp returns a boolean if a field has been set.

### GetNodeSubnet

`func (o *KubernetesClusterProperties) GetNodeSubnet() string`

GetNodeSubnet returns the NodeSubnet field if non-nil, zero value otherwise.

### GetNodeSubnetOk

`func (o *KubernetesClusterProperties) GetNodeSubnetOk() (*string, bool)`

GetNodeSubnetOk returns a tuple with the NodeSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSubnet

`func (o *KubernetesClusterProperties) SetNodeSubnet(v string)`

SetNodeSubnet sets NodeSubnet field to given value.

### HasNodeSubnet

`func (o *KubernetesClusterProperties) HasNodeSubnet() bool`

HasNodeSubnet returns a boolean if a field has been set.

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



