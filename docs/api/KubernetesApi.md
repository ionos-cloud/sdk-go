# \KubernetesApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**K8sDelete**](KubernetesApi.md#K8sDelete) | **Delete** /k8s/{k8sClusterId} | Delete Kubernetes clusters|
|[**K8sFindByClusterId**](KubernetesApi.md#K8sFindByClusterId) | **Get** /k8s/{k8sClusterId} | Retrieve Kubernetes clusters|
|[**K8sGet**](KubernetesApi.md#K8sGet) | **Get** /k8s | List Kubernetes clusters|
|[**K8sKubeconfigGet**](KubernetesApi.md#K8sKubeconfigGet) | **Get** /k8s/{k8sClusterId}/kubeconfig | Retrieve Kubernetes configuration files|
|[**K8sNodepoolsDelete**](KubernetesApi.md#K8sNodepoolsDelete) | **Delete** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Delete Kubernetes node pools|
|[**K8sNodepoolsFindById**](KubernetesApi.md#K8sNodepoolsFindById) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Retrieve Kubernetes node pools|
|[**K8sNodepoolsGet**](KubernetesApi.md#K8sNodepoolsGet) | **Get** /k8s/{k8sClusterId}/nodepools | List Kubernetes node pools|
|[**K8sNodepoolsNodesDelete**](KubernetesApi.md#K8sNodepoolsNodesDelete) | **Delete** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId} | Delete Kubernetes nodes|
|[**K8sNodepoolsNodesFindById**](KubernetesApi.md#K8sNodepoolsNodesFindById) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId} | Retrieve Kubernetes nodes|
|[**K8sNodepoolsNodesGet**](KubernetesApi.md#K8sNodepoolsNodesGet) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes | List Kubernetes nodes|
|[**K8sNodepoolsNodesReplacePost**](KubernetesApi.md#K8sNodepoolsNodesReplacePost) | **Post** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId}/replace | Recreate Kubernetes nodes|
|[**K8sNodepoolsPost**](KubernetesApi.md#K8sNodepoolsPost) | **Post** /k8s/{k8sClusterId}/nodepools | Create Kubernetes node pools|
|[**K8sNodepoolsPut**](KubernetesApi.md#K8sNodepoolsPut) | **Put** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Modify Kubernetes node pools|
|[**K8sPost**](KubernetesApi.md#K8sPost) | **Post** /k8s | Create Kubernetes clusters|
|[**K8sPut**](KubernetesApi.md#K8sPut) | **Put** /k8s/{k8sClusterId} | Modify Kubernetes clusters|
|[**K8sVersionsDefaultGet**](KubernetesApi.md#K8sVersionsDefaultGet) | **Get** /k8s/versions/default | Retrieve current default Kubernetes version|
|[**K8sVersionsGet**](KubernetesApi.md#K8sVersionsGet) | **Get** /k8s/versions | List Kubernetes versions|



## K8sDelete

```go
var result  = K8sDelete(ctx, k8sClusterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete Kubernetes clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sDelete(context.Background(), k8sClusterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sFindByClusterId

```go
var result KubernetesCluster = K8sFindByClusterId(ctx, k8sClusterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Kubernetes clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sFindByClusterId(context.Background(), k8sClusterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sFindByClusterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sFindByClusterId`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sFindByClusterId`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sFindByClusterIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesCluster**](../models/KubernetesCluster.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sGet

```go
var result KubernetesClusters = K8sGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Kubernetes clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sGet`: KubernetesClusters
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiK8sGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesClusters**](../models/KubernetesClusters.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sKubeconfigGet

```go
var result string = K8sKubeconfigGet(ctx, k8sClusterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Kubernetes configuration files



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sKubeconfigGet(context.Background(), k8sClusterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sKubeconfigGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sKubeconfigGet`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sKubeconfigGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sKubeconfigGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

**string**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yamlapplication/x-yamlapplication/json



## K8sNodepoolsDelete

```go
var result  = K8sNodepoolsDelete(ctx, k8sClusterId, nodepoolId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete Kubernetes node pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes node pool.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sNodepoolsDelete(context.Background(), k8sClusterId, nodepoolId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |
|**nodepoolId** | **string** | The unique ID of the Kubernetes node pool. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sNodepoolsFindById

```go
var result KubernetesNodePool = K8sNodepoolsFindById(ctx, k8sClusterId, nodepoolId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Kubernetes node pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes node pool.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sNodepoolsFindById(context.Background(), k8sClusterId, nodepoolId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsFindById`: KubernetesNodePool
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |
|**nodepoolId** | **string** | The unique ID of the Kubernetes node pool. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesNodePool**](../models/KubernetesNodePool.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sNodepoolsGet

```go
var result KubernetesNodePools = K8sNodepoolsGet(ctx, k8sClusterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Kubernetes node pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sNodepoolsGet(context.Background(), k8sClusterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsGet`: KubernetesNodePools
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesNodePools**](../models/KubernetesNodePools.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sNodepoolsNodesDelete

```go
var result  = K8sNodepoolsNodesDelete(ctx, k8sClusterId, nodepoolId, nodeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete Kubernetes nodes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes node pool.
    nodeId := "nodeId_example" // string | The unique ID of the Kubernetes node.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sNodepoolsNodesDelete(context.Background(), k8sClusterId, nodepoolId, nodeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsNodesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |
|**nodepoolId** | **string** | The unique ID of the Kubernetes node pool. | |
|**nodeId** | **string** | The unique ID of the Kubernetes node. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsNodesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sNodepoolsNodesFindById

```go
var result KubernetesNode = K8sNodepoolsNodesFindById(ctx, k8sClusterId, nodepoolId, nodeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Kubernetes nodes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes node pool.
    nodeId := "nodeId_example" // string | The unique ID of the Kubernetes node.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sNodepoolsNodesFindById(context.Background(), k8sClusterId, nodepoolId, nodeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsNodesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsNodesFindById`: KubernetesNode
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsNodesFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |
|**nodepoolId** | **string** | The unique ID of the Kubernetes node pool. | |
|**nodeId** | **string** | The unique ID of the Kubernetes node. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsNodesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesNode**](../models/KubernetesNode.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sNodepoolsNodesGet

```go
var result KubernetesNodes = K8sNodepoolsNodesGet(ctx, k8sClusterId, nodepoolId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Kubernetes nodes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes node pool.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sNodepoolsNodesGet(context.Background(), k8sClusterId, nodepoolId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsNodesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsNodesGet`: KubernetesNodes
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsNodesGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |
|**nodepoolId** | **string** | The unique ID of the Kubernetes node pool. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsNodesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesNodes**](../models/KubernetesNodes.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sNodepoolsNodesReplacePost

```go
var result  = K8sNodepoolsNodesReplacePost(ctx, k8sClusterId, nodepoolId, nodeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Recreate Kubernetes nodes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes node pool.
    nodeId := "nodeId_example" // string | The unique ID of the Kubernetes node.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sNodepoolsNodesReplacePost(context.Background(), k8sClusterId, nodepoolId, nodeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsNodesReplacePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |
|**nodepoolId** | **string** | The unique ID of the Kubernetes node pool. | |
|**nodeId** | **string** | The unique ID of the Kubernetes node. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsNodesReplacePostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sNodepoolsPost

```go
var result KubernetesNodePool = K8sNodepoolsPost(ctx, k8sClusterId)
                      .KubernetesNodePool(kubernetesNodePool)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create Kubernetes node pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    kubernetesNodePool := *openapiclient.NewKubernetesNodePoolForPost(*openapiclient.NewKubernetesNodePoolPropertiesForPost("k8s-node-pool", "1e072e52-2ed3-492f-b6b6-c6b116907521", int32(2), "AMD_OPTERON", int32(4), int32(2048), "AUTO", "HDD", int32(100))) // KubernetesNodePoolForPost | The Kubernetes node pool to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sNodepoolsPost(context.Background(), k8sClusterId).KubernetesNodePool(kubernetesNodePool).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsPost`: KubernetesNodePool
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **kubernetesNodePool** | [**KubernetesNodePoolForPost**](../models/KubernetesNodePoolForPost.md) | The Kubernetes node pool to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesNodePool**](../models/KubernetesNodePool.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## K8sNodepoolsPut

```go
var result KubernetesNodePool = K8sNodepoolsPut(ctx, k8sClusterId, nodepoolId)
                      .KubernetesNodePool(kubernetesNodePool)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify Kubernetes node pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes node pool.
    kubernetesNodePool := *openapiclient.NewKubernetesNodePoolForPut(*openapiclient.NewKubernetesNodePoolPropertiesForPut(int32(2))) // KubernetesNodePoolForPut | Details of the Kubernetes Node Pool
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sNodepoolsPut(context.Background(), k8sClusterId, nodepoolId).KubernetesNodePool(kubernetesNodePool).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsPut`: KubernetesNodePool
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |
|**nodepoolId** | **string** | The unique ID of the Kubernetes node pool. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **kubernetesNodePool** | [**KubernetesNodePoolForPut**](../models/KubernetesNodePoolForPut.md) | Details of the Kubernetes Node Pool | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesNodePool**](../models/KubernetesNodePool.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## K8sPost

```go
var result KubernetesCluster = K8sPost(ctx)
                      .KubernetesCluster(kubernetesCluster)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create Kubernetes clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    kubernetesCluster := *openapiclient.NewKubernetesClusterForPost(*openapiclient.NewKubernetesClusterPropertiesForPost("k8s")) // KubernetesClusterForPost | The Kubernetes cluster to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sPost(context.Background()).KubernetesCluster(kubernetesCluster).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sPost`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiK8sPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **kubernetesCluster** | [**KubernetesClusterForPost**](../models/KubernetesClusterForPost.md) | The Kubernetes cluster to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesCluster**](../models/KubernetesCluster.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## K8sPut

```go
var result KubernetesCluster = K8sPut(ctx, k8sClusterId)
                      .KubernetesCluster(kubernetesCluster)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify Kubernetes clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes cluster.
    kubernetesCluster := *openapiclient.NewKubernetesClusterForPut(*openapiclient.NewKubernetesClusterPropertiesForPut("k8s")) // KubernetesClusterForPut | The modified Kubernetes cluster.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sPut(context.Background(), k8sClusterId).KubernetesCluster(kubernetesCluster).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sPut`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**k8sClusterId** | **string** | The unique ID of the Kubernetes cluster. | |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **kubernetesCluster** | [**KubernetesClusterForPut**](../models/KubernetesClusterForPut.md) | The modified Kubernetes cluster. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**KubernetesCluster**](../models/KubernetesCluster.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## K8sVersionsDefaultGet

```go
var result string = K8sVersionsDefaultGet(ctx)
                      .Execute()
```

Retrieve current default Kubernetes version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sVersionsDefaultGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sVersionsDefaultGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sVersionsDefaultGet`: string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sVersionsDefaultGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiK8sVersionsDefaultGetRequest struct via the builder pattern


### Return type

**string**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## K8sVersionsGet

```go
var result []string = K8sVersionsGet(ctx)
                      .Execute()
```

List Kubernetes versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesApi.K8sVersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sVersionsGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sVersionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiK8sVersionsGetRequest struct via the builder pattern


### Return type

**[]string**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


