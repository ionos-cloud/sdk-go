# \KubernetesApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**K8sDelete**](KubernetesApi.md#K8sDelete) | **Delete** /k8s/{k8sClusterId} | Delete Kubernetes Cluster
[**K8sFindByClusterid**](KubernetesApi.md#K8sFindByClusterid) | **Get** /k8s/{k8sClusterId} | Retrieve Kubernetes Cluster
[**K8sGet**](KubernetesApi.md#K8sGet) | **Get** /k8s | List Kubernetes Clusters
[**K8sKubeconfigGet**](KubernetesApi.md#K8sKubeconfigGet) | **Get** /k8s/{k8sClusterId}/kubeconfig | Retrieve Kubernetes Configuration File
[**K8sNodepoolsDelete**](KubernetesApi.md#K8sNodepoolsDelete) | **Delete** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Delete Kubernetes Node Pool
[**K8sNodepoolsFindById**](KubernetesApi.md#K8sNodepoolsFindById) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Retrieve Kubernetes Node Pool
[**K8sNodepoolsGet**](KubernetesApi.md#K8sNodepoolsGet) | **Get** /k8s/{k8sClusterId}/nodepools | List Kubernetes Node Pools
[**K8sNodepoolsNodesDelete**](KubernetesApi.md#K8sNodepoolsNodesDelete) | **Delete** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId} | Delete Kubernetes node
[**K8sNodepoolsNodesFindById**](KubernetesApi.md#K8sNodepoolsNodesFindById) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId} | Retrieve Kubernetes node
[**K8sNodepoolsNodesGet**](KubernetesApi.md#K8sNodepoolsNodesGet) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes | Retrieve Kubernetes nodes.
[**K8sNodepoolsNodesReplacePost**](KubernetesApi.md#K8sNodepoolsNodesReplacePost) | **Post** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId}/replace | Recreate the Kubernetes node
[**K8sNodepoolsPost**](KubernetesApi.md#K8sNodepoolsPost) | **Post** /k8s/{k8sClusterId}/nodepools | Create a Kubernetes Node Pool
[**K8sNodepoolsPut**](KubernetesApi.md#K8sNodepoolsPut) | **Put** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Modify Kubernetes Node Pool
[**K8sPost**](KubernetesApi.md#K8sPost) | **Post** /k8s | Create Kubernetes Cluster
[**K8sPut**](KubernetesApi.md#K8sPut) | **Put** /k8s/{k8sClusterId} | Modify Kubernetes Cluster
[**K8sVersionsCompatibilitiesGet**](KubernetesApi.md#K8sVersionsCompatibilitiesGet) | **Get** /k8s/versions/{clusterVersion}/compatibilities | Retrieves a list of available kubernetes versions for nodepools depending on the given kubernetes version running in the cluster.
[**K8sVersionsDefaultGet**](KubernetesApi.md#K8sVersionsDefaultGet) | **Get** /k8s/versions/default | Retrieve the current default kubernetes version for clusters and nodepools.
[**K8sVersionsGet**](KubernetesApi.md#K8sVersionsGet) | **Get** /k8s/versions | Retrieve available Kubernetes versions



## K8sDelete

> map[string]interface{} K8sDelete(ctx, k8sClusterId, optional)

Delete Kubernetes Cluster

This will remove a Kubernetes Cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***K8sDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sFindByClusterid

> KubernetesCluster K8sFindByClusterid(ctx, k8sClusterId, optional)

Retrieve Kubernetes Cluster

This will retrieve a single Kubernetes Cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***K8sFindByClusteridOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sFindByClusteridOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sGet

> KubernetesClusters K8sGet(ctx, optional)

List Kubernetes Clusters

You can retrieve a list of all kubernetes clusters associated with a contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***K8sGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesClusters**](KubernetesClusters.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sKubeconfigGet

> KubernetesConfig K8sKubeconfigGet(ctx, k8sClusterId, optional)

Retrieve Kubernetes Configuration File

You can retrieve kubernetes configuration file for the kubernetes cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***K8sKubeconfigGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sKubeconfigGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesConfig**](KubernetesConfig.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sNodepoolsDelete

> map[string]interface{} K8sNodepoolsDelete(ctx, k8sClusterId, nodepoolId, optional)

Delete Kubernetes Node Pool

This will remove a Kubernetes Node Pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
**nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
 **optional** | ***K8sNodepoolsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sNodepoolsDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sNodepoolsFindById

> KubernetesNodePool K8sNodepoolsFindById(ctx, k8sClusterId, nodepoolId, optional)

Retrieve Kubernetes Node Pool

You can retrieve a single Kubernetes Node Pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
**nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
 **optional** | ***K8sNodepoolsFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sNodepoolsFindByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodePool**](KubernetesNodePool.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sNodepoolsGet

> KubernetesNodePools K8sNodepoolsGet(ctx, k8sClusterId, optional)

List Kubernetes Node Pools

You can retrieve a list of all kubernetes node pools part of kubernetes cluster

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
 **optional** | ***K8sNodepoolsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sNodepoolsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodePools**](KubernetesNodePools.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sNodepoolsNodesDelete

> map[string]interface{} K8sNodepoolsNodesDelete(ctx, k8sClusterId, nodepoolId, nodeId, optional)

Delete Kubernetes node

This will remove a Kubernetes node.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
**nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
**nodeId** | **string**| The unique ID of the Kubernetes node | 
 **optional** | ***K8sNodepoolsNodesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sNodepoolsNodesDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sNodepoolsNodesFindById

> KubernetesNode K8sNodepoolsNodesFindById(ctx, k8sClusterId, nodepoolId, nodeId, optional)

Retrieve Kubernetes node

You can retrieve a single Kubernetes Node.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
**nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
**nodeId** | **string**| The unique ID of the Kubernetes Node. | 
 **optional** | ***K8sNodepoolsNodesFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sNodepoolsNodesFindByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNode**](KubernetesNode.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sNodepoolsNodesGet

> KubernetesNodes K8sNodepoolsNodesGet(ctx, k8sClusterId, nodepoolId, optional)

Retrieve Kubernetes nodes.

You can retrieve all nodes of Kubernetes Node Pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
**nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
 **optional** | ***K8sNodepoolsNodesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sNodepoolsNodesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodes**](KubernetesNodes.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sNodepoolsNodesReplacePost

> map[string]interface{} K8sNodepoolsNodesReplacePost(ctx, k8sClusterId, nodepoolId, nodeId, optional)

Recreate the Kubernetes node

You can recreate a single Kubernetes Node.  Managed Kubernetes starts a process which based on the nodepool's template creates & configures a new node, waits for status \"ACTIVE\", and migrates all the pods from the faulty node, deleting it once empty. While this operation occurs, the nodepool will have an extra billable \"ACTIVE\" node.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
**nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
**nodeId** | **string**| The unique ID of the Kubernetes Node. | 
 **optional** | ***K8sNodepoolsNodesReplacePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sNodepoolsNodesReplacePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sNodepoolsPost

> KubernetesNodePool K8sNodepoolsPost(ctx, k8sClusterId, kubernetesNodePoolProperties, optional)

Create a Kubernetes Node Pool

This will create a new Kubernetes Node Pool inside a Kubernetes Cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
**kubernetesNodePoolProperties** | [**KubernetesNodePoolProperties**](KubernetesNodePoolProperties.md)| Details of Kubernetes Node Pool | 
 **optional** | ***K8sNodepoolsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sNodepoolsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodePool**](KubernetesNodePool.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sNodepoolsPut

> KubernetesNodePoolForPut K8sNodepoolsPut(ctx, k8sClusterId, nodepoolId, kubernetesNodePoolProperties, optional)

Modify Kubernetes Node Pool

This will modify the Kubernetes Node Pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
**nodepoolId** | **string**| The unique ID of the Kubernetes Node Pool | 
**kubernetesNodePoolProperties** | [**KubernetesNodePoolPropertiesForPut**](KubernetesNodePoolPropertiesForPut.md)| Details of the Kubernetes Node Pool | 
 **optional** | ***K8sNodepoolsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sNodepoolsPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesNodePoolForPut**](KubernetesNodePoolForPut.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sPost

> KubernetesCluster K8sPost(ctx, kubernetesCluster, optional)

Create Kubernetes Cluster

This will create a new Kubernetes Cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kubernetesCluster** | [**KubernetesCluster**](KubernetesCluster.md)| Properties of the Kubernetes Cluster | 
 **optional** | ***K8sPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sPut

> KubernetesCluster K8sPut(ctx, k8sClusterId, kubernetescluster, optional)

Modify Kubernetes Cluster

This will modify the Kubernetes Cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**k8sClusterId** | **string**| The unique ID of the Kubernetes Cluster | 
**kubernetescluster** | [**KubernetesCluster**](KubernetesCluster.md)| Details of of the Kubernetes Cluster | 
 **optional** | ***K8sPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a K8sPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sVersionsCompatibilitiesGet

> []string K8sVersionsCompatibilitiesGet(ctx, clusterVersion)

Retrieves a list of available kubernetes versions for nodepools depending on the given kubernetes version running in the cluster.

You can retrieve a list of available kubernetes versions for nodepools depending on the given kubernetes version running in the cluster.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterVersion** | **string**|  | 

### Return type

**[]string**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sVersionsDefaultGet

> string K8sVersionsDefaultGet(ctx, )

Retrieve the current default kubernetes version for clusters and nodepools.

You can retrieve the current default kubernetes version for clusters and nodepools.

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## K8sVersionsGet

> []string K8sVersionsGet(ctx, )

Retrieve available Kubernetes versions

You can retrieve a list of available kubernetes versions

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

