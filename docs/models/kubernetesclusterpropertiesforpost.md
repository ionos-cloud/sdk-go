# KubernetesClusterPropertiesForPost

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Name** | **string** | A Kubernetes Cluster Name. Valid Kubernetes Cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character \(\[a-z0-9A-Z\]\) with dashes \(-\), underscores \(\_\), dots \(.\), and alphanumerics between. |  |
| **K8sVersion** | Pointer to **string** | The kubernetes version in which a cluster is running. This imposes restrictions on what kubernetes versions can be run in a cluster's nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions. | \[optional\] |
| **MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](kubernetesmaintenancewindow.md) |  | \[optional\] |
| **Public** | Pointer to **bool** | The indicator if the cluster is public or private. Be aware that setting it to false is currently in beta phase. | \[optional\] \[default to true\] |
| **GatewayIp** | Pointer to **string** | The IP address of the gateway used by the cluster. This is mandatory when \`public\` is set to \`false\` and should not be provided otherwise. | \[optional\] |

## Methods

### NewKubernetesClusterPropertiesForPost

`func NewKubernetesClusterPropertiesForPost(name string, ) *KubernetesClusterPropertiesForPost`

NewKubernetesClusterPropertiesForPost instantiates a new KubernetesClusterPropertiesForPost object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewKubernetesClusterPropertiesForPostWithDefaults

`func NewKubernetesClusterPropertiesForPostWithDefaults() *KubernetesClusterPropertiesForPost`

NewKubernetesClusterPropertiesForPostWithDefaults instantiates a new KubernetesClusterPropertiesForPost object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesClusterPropertiesForPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesClusterPropertiesForPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesClusterPropertiesForPost) SetName(v string)`

SetName sets Name field to given value.

### GetK8sVersion

`func (o *KubernetesClusterPropertiesForPost) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesClusterPropertiesForPost) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

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

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *KubernetesClusterPropertiesForPost) SetMaintenanceWindow(v KubernetesMaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *KubernetesClusterPropertiesForPost) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

### GetPublic

`func (o *KubernetesClusterPropertiesForPost) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *KubernetesClusterPropertiesForPost) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetPublic

`func (o *KubernetesClusterPropertiesForPost) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *KubernetesClusterPropertiesForPost) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetGatewayIp

`func (o *KubernetesClusterPropertiesForPost) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *KubernetesClusterPropertiesForPost) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *KubernetesClusterPropertiesForPost) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *KubernetesClusterPropertiesForPost) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

