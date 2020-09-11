# KubernetesNodePoolPropertiesForPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A Kubernetes Node Pool Name. Valid Kubernetes Node Pool name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. | 
**DatacenterId** | **string** | A valid uuid of the datacenter on which user has access | 
**NodeCount** | **int32** | Number of nodes part of the Node Pool | 
**CpuFamily** | **string** | A valid cpu family name | 
**CoresCount** | **int32** | Number of cores for node | 
**RamSize** | **int32** | RAM size for node, minimum size 2048MB is recommended. Ram size must be set to multiple of 1024MB. | 
**AvailabilityZone** | **string** | The availability zone in which the server should exist | 
**StorageType** | **string** | Hardware type of the volume | 
**StorageSize** | **int32** | The size of the volume in GB. The size should be greater than 10GB. | 
**K8sVersion** | **string** | The kubernetes version in which a nodepool is running. This imposes restrictions on what kubernetes versions can be run in a cluster&#39;s nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions. | [optional] 
**MaintenanceWindow** | [**KubernetesMaintenanceWindow**](KubernetesMaintenanceWindow.md) |  | [optional] 
**AutoScaling** | [**KubernetesAutoScaling**](KubernetesAutoScaling.md) |  | [optional] 
**Lans** | [**[]KubernetesNodePoolLan**](KubernetesNodePoolLan.md) | array of additional LANs attached to worker nodes | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


