/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0-SDK.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// KubernetesClusterProperties struct for KubernetesClusterProperties
type KubernetesClusterProperties struct {
	// A Kubernetes Cluster Name. Valid Kubernetes Cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Name *string `json:"name"`
	// The kubernetes version in which a cluster is running. This imposes restrictions on what kubernetes versions can be run in a cluster's nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions.
	K8sVersion        *string                      `json:"k8sVersion,omitempty"`
	MaintenanceWindow *KubernetesMaintenanceWindow `json:"maintenanceWindow,omitempty"`
	// List of available versions for upgrading the cluster
	AvailableUpgradeVersions *[]string `json:"availableUpgradeVersions,omitempty"`
	// List of versions that may be used for node pools under this cluster
	ViableNodePoolVersions *[]string `json:"viableNodePoolVersions,omitempty"`
	// Access to the K8s API server is restricted to these CIDRs. Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6.
	ApiSubnetAllowList *[]string `json:"apiSubnetAllowList,omitempty"`
	// List of S3 bucket configured for K8s usage. For now it contains only an S3 bucket used to store K8s API audit logs
	S3Buckets *[]S3Bucket `json:"s3Buckets,omitempty"`
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesClusterProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *KubernetesClusterProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetK8sVersion returns the K8sVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesClusterProperties) GetK8sVersion() *string {
	if o == nil {
		return nil
	}

	return o.K8sVersion

}

// GetK8sVersionOk returns a tuple with the K8sVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProperties) GetK8sVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.K8sVersion, true
}

// SetK8sVersion sets field value
func (o *KubernetesClusterProperties) SetK8sVersion(v string) {

	o.K8sVersion = &v

}

// HasK8sVersion returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasK8sVersion() bool {
	if o != nil && o.K8sVersion != nil {
		return true
	}

	return false
}

// GetMaintenanceWindow returns the MaintenanceWindow field value
// If the value is explicit nil, the zero value for KubernetesMaintenanceWindow will be returned
func (o *KubernetesClusterProperties) GetMaintenanceWindow() *KubernetesMaintenanceWindow {
	if o == nil {
		return nil
	}

	return o.MaintenanceWindow

}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProperties) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool) {
	if o == nil {
		return nil, false
	}

	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *KubernetesClusterProperties) SetMaintenanceWindow(v KubernetesMaintenanceWindow) {

	o.MaintenanceWindow = &v

}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasMaintenanceWindow() bool {
	if o != nil && o.MaintenanceWindow != nil {
		return true
	}

	return false
}

// GetAvailableUpgradeVersions returns the AvailableUpgradeVersions field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *KubernetesClusterProperties) GetAvailableUpgradeVersions() *[]string {
	if o == nil {
		return nil
	}

	return o.AvailableUpgradeVersions

}

// GetAvailableUpgradeVersionsOk returns a tuple with the AvailableUpgradeVersions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProperties) GetAvailableUpgradeVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.AvailableUpgradeVersions, true
}

// SetAvailableUpgradeVersions sets field value
func (o *KubernetesClusterProperties) SetAvailableUpgradeVersions(v []string) {

	o.AvailableUpgradeVersions = &v

}

// HasAvailableUpgradeVersions returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasAvailableUpgradeVersions() bool {
	if o != nil && o.AvailableUpgradeVersions != nil {
		return true
	}

	return false
}

// GetViableNodePoolVersions returns the ViableNodePoolVersions field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *KubernetesClusterProperties) GetViableNodePoolVersions() *[]string {
	if o == nil {
		return nil
	}

	return o.ViableNodePoolVersions

}

// GetViableNodePoolVersionsOk returns a tuple with the ViableNodePoolVersions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProperties) GetViableNodePoolVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ViableNodePoolVersions, true
}

// SetViableNodePoolVersions sets field value
func (o *KubernetesClusterProperties) SetViableNodePoolVersions(v []string) {

	o.ViableNodePoolVersions = &v

}

// HasViableNodePoolVersions returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasViableNodePoolVersions() bool {
	if o != nil && o.ViableNodePoolVersions != nil {
		return true
	}

	return false
}

// GetApiSubnetAllowList returns the ApiSubnetAllowList field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *KubernetesClusterProperties) GetApiSubnetAllowList() *[]string {
	if o == nil {
		return nil
	}

	return o.ApiSubnetAllowList

}

// GetApiSubnetAllowListOk returns a tuple with the ApiSubnetAllowList field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProperties) GetApiSubnetAllowListOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ApiSubnetAllowList, true
}

// SetApiSubnetAllowList sets field value
func (o *KubernetesClusterProperties) SetApiSubnetAllowList(v []string) {

	o.ApiSubnetAllowList = &v

}

// HasApiSubnetAllowList returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasApiSubnetAllowList() bool {
	if o != nil && o.ApiSubnetAllowList != nil {
		return true
	}

	return false
}

// GetS3Buckets returns the S3Buckets field value
// If the value is explicit nil, the zero value for []S3Bucket will be returned
func (o *KubernetesClusterProperties) GetS3Buckets() *[]S3Bucket {
	if o == nil {
		return nil
	}

	return o.S3Buckets

}

// GetS3BucketsOk returns a tuple with the S3Buckets field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterProperties) GetS3BucketsOk() (*[]S3Bucket, bool) {
	if o == nil {
		return nil, false
	}

	return o.S3Buckets, true
}

// SetS3Buckets sets field value
func (o *KubernetesClusterProperties) SetS3Buckets(v []S3Bucket) {

	o.S3Buckets = &v

}

// HasS3Buckets returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasS3Buckets() bool {
	if o != nil && o.S3Buckets != nil {
		return true
	}

	return false
}

func (o KubernetesClusterProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.K8sVersion != nil {
		toSerialize["k8sVersion"] = o.K8sVersion
	}

	if o.MaintenanceWindow != nil {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}

	if o.AvailableUpgradeVersions != nil {
		toSerialize["availableUpgradeVersions"] = o.AvailableUpgradeVersions
	}

	if o.ViableNodePoolVersions != nil {
		toSerialize["viableNodePoolVersions"] = o.ViableNodePoolVersions
	}

	if o.ApiSubnetAllowList != nil {
		toSerialize["apiSubnetAllowList"] = o.ApiSubnetAllowList
	}

	if o.S3Buckets != nil {
		toSerialize["s3Buckets"] = o.S3Buckets
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesClusterProperties struct {
	value *KubernetesClusterProperties
	isSet bool
}

func (v NullableKubernetesClusterProperties) Get() *KubernetesClusterProperties {
	return v.value
}

func (v *NullableKubernetesClusterProperties) Set(val *KubernetesClusterProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterProperties(val *KubernetesClusterProperties) *NullableKubernetesClusterProperties {
	return &NullableKubernetesClusterProperties{value: val, isSet: true}
}

func (v NullableKubernetesClusterProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
