# KubernetesClusterProperties

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Name** | **string** | A Kubernetes Cluster Name. Valid Kubernetes Cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character \(\[a-z0-9A-Z\]\) with dashes \(-\), underscores \(\_\), dots \(.\), and alphanumerics between. |  |
| **K8sVersion** | Pointer to **string** | The kubernetes version in which a cluster is running. This imposes restrictions on what kubernetes versions can be run in a cluster's nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions. | \[optional\] |
| **MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](kubernetesmaintenancewindow.md) |  | \[optional\] |
| **AvailableUpgradeVersions** | Pointer to **\[\]string** | List of available versions for upgrading the cluster | \[optional\] |
| **ViableNodePoolVersions** | Pointer to **\[\]string** | List of versions that may be used for node pools under this cluster | \[optional\] |
| **Public** | Pointer to **bool** | The indicator if the cluster is public or private. Be aware that setting it to false is currently in beta phase. | \[optional\] \[default to true\] |
| **GatewayIp** | Pointer to **string** | The IP address of the gateway used by the cluster. This is mandatory when \`public\` is set to \`false\` and should not be provided otherwise. | \[optional\] |

## Methods

### NewKubernetesClusterProperties

`func NewKubernetesClusterProperties(name string, ) *KubernetesClusterProperties`

NewKubernetesClusterProperties instantiates a new KubernetesClusterProperties object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewKubernetesClusterPropertiesWithDefaults

`func NewKubernetesClusterPropertiesWithDefaults() *KubernetesClusterProperties`

NewKubernetesClusterPropertiesWithDefaults instantiates a new KubernetesClusterProperties object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesClusterProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesClusterProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesClusterProperties) SetName(v string)`

SetName sets Name field to given value.

### GetK8sVersion

`func (o *KubernetesClusterProperties) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesClusterProperties) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

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

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

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

GetAvailableUpgradeVersionsOk returns a tuple with the AvailableUpgradeVersions field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

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

GetViableNodePoolVersionsOk returns a tuple with the ViableNodePoolVersions field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

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

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetPublic

`func (o *KubernetesClusterProperties) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *KubernetesClusterProperties) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetGatewayIp

`func (o *KubernetesClusterProperties) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *KubernetesClusterProperties) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *KubernetesClusterProperties) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *KubernetesClusterProperties) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

