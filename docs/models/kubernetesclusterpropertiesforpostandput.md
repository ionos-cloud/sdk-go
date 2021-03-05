# KubernetesClusterPropertiesForPostAndPut

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Name** | **string** | A Kubernetes Cluster Name. Valid Kubernetes Cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character \(\[a-z0-9A-Z\]\) with dashes \(-\), underscores \(\_\), dots \(.\), and alphanumerics between. |  |
| **K8sVersion** | Pointer to **string** | The kubernetes version in which a cluster is running. This imposes restrictions on what kubernetes versions can be run in a cluster's nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions. | \[optional\] |
| **MaintenanceWindow** | Pointer to [**KubernetesMaintenanceWindow**](kubernetesmaintenancewindow.md) |  | \[optional\] |

## Methods

### NewKubernetesClusterPropertiesForPostAndPut

`func NewKubernetesClusterPropertiesForPostAndPut(name string, ) *KubernetesClusterPropertiesForPostAndPut`

NewKubernetesClusterPropertiesForPostAndPut instantiates a new KubernetesClusterPropertiesForPostAndPut object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewKubernetesClusterPropertiesForPostAndPutWithDefaults

`func NewKubernetesClusterPropertiesForPostAndPutWithDefaults() *KubernetesClusterPropertiesForPostAndPut`

NewKubernetesClusterPropertiesForPostAndPutWithDefaults instantiates a new KubernetesClusterPropertiesForPostAndPut object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubernetesClusterPropertiesForPostAndPut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesClusterPropertiesForPostAndPut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesClusterPropertiesForPostAndPut) SetName(v string)`

SetName sets Name field to given value.

### GetK8sVersion

`func (o *KubernetesClusterPropertiesForPostAndPut) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *KubernetesClusterPropertiesForPostAndPut) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *KubernetesClusterPropertiesForPostAndPut) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *KubernetesClusterPropertiesForPostAndPut) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetMaintenanceWindow

`func (o *KubernetesClusterPropertiesForPostAndPut) GetMaintenanceWindow() KubernetesMaintenanceWindow`

GetMaintenanceWindow returns the MaintenanceWindow field if non-nil, zero value otherwise.

### GetMaintenanceWindowOk

`func (o *KubernetesClusterPropertiesForPostAndPut) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool)`

GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetMaintenanceWindow

`func (o *KubernetesClusterPropertiesForPostAndPut) SetMaintenanceWindow(v KubernetesMaintenanceWindow)`

SetMaintenanceWindow sets MaintenanceWindow field to given value.

### HasMaintenanceWindow

`func (o *KubernetesClusterPropertiesForPostAndPut) HasMaintenanceWindow() bool`

HasMaintenanceWindow returns a boolean if a field has been set.

