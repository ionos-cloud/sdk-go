# KubernetesApi

All URIs are relative to [https://api.ionos.com/cloudapi/v5](https://api.ionos.com/cloudapi/v5)

| Method | HTTP request | Description |
| :--- | :--- | :--- |
| [**K8sDelete**](kubernetesapi.md#K8sDelete) | **Delete** /k8s/{k8sClusterId} | Delete Kubernetes Cluster |
| [**K8sFindByClusterId**](kubernetesapi.md#K8sFindByClusterId) | **Get** /k8s/{k8sClusterId} | Retrieve Kubernetes Cluster |
| [**K8sGet**](kubernetesapi.md#K8sGet) | **Get** /k8s | List Kubernetes Clusters |
| [**K8sKubeconfigGet**](kubernetesapi.md#K8sKubeconfigGet) | **Get** /k8s/{k8sClusterId}/kubeconfig | Retrieve Kubernetes Configuration File |
| [**K8sNodepoolsDelete**](kubernetesapi.md#K8sNodepoolsDelete) | **Delete** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Delete Kubernetes Node Pool |
| [**K8sNodepoolsFindById**](kubernetesapi.md#K8sNodepoolsFindById) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Retrieve Kubernetes Node Pool |
| [**K8sNodepoolsGet**](kubernetesapi.md#K8sNodepoolsGet) | **Get** /k8s/{k8sClusterId}/nodepools | List Kubernetes Node Pools |
| [**K8sNodepoolsNodesDelete**](kubernetesapi.md#K8sNodepoolsNodesDelete) | **Delete** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId} | Delete Kubernetes node |
| [**K8sNodepoolsNodesFindById**](kubernetesapi.md#K8sNodepoolsNodesFindById) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId} | Retrieve Kubernetes node |
| [**K8sNodepoolsNodesGet**](kubernetesapi.md#K8sNodepoolsNodesGet) | **Get** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes | Retrieve Kubernetes nodes. |
| [**K8sNodepoolsNodesReplacePost**](kubernetesapi.md#K8sNodepoolsNodesReplacePost) | **Post** /k8s/{k8sClusterId}/nodepools/{nodepoolId}/nodes/{nodeId}/replace | Recreate the Kubernetes node |
| [**K8sNodepoolsPost**](kubernetesapi.md#K8sNodepoolsPost) | **Post** /k8s/{k8sClusterId}/nodepools | Create a Kubernetes Node Pool |
| [**K8sNodepoolsPut**](kubernetesapi.md#K8sNodepoolsPut) | **Put** /k8s/{k8sClusterId}/nodepools/{nodepoolId} | Modify Kubernetes Node Pool |
| [**K8sPost**](kubernetesapi.md#K8sPost) | **Post** /k8s | Create Kubernetes Cluster |
| [**K8sPut**](kubernetesapi.md#K8sPut) | **Put** /k8s/{k8sClusterId} | Modify Kubernetes Cluster |
| [**K8sVersionsCompatibilitiesGet**](kubernetesapi.md#K8sVersionsCompatibilitiesGet) | **Get** /k8s/versions/{clusterVersion}/compatibilities | Retrieves a list of available kubernetes versions for nodepools depending on the given kubernetes version running in the cluster. |
| [**K8sVersionsDefaultGet**](kubernetesapi.md#K8sVersionsDefaultGet) | **Get** /k8s/versions/default | Retrieve the current default kubernetes version for clusters and nodepools. |
| [**K8sVersionsGet**](kubernetesapi.md#K8sVersionsGet) | **Get** /k8s/versions | Retrieve available Kubernetes versions |

## K8sDelete

```go
var result map[string]interface{} = K8sDelete(ctx, k8sClusterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete Kubernetes Cluster

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sDelete(context.Background(), k8sClusterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sDelete`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

**map\[string\]interface{}**

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sFindByClusterId

```go
var result KubernetesCluster = K8sFindByClusterId(ctx, k8sClusterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Kubernetes Cluster

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sFindByClusterId(context.Background(), k8sClusterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sFindByClusterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sFindByClusterId`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sFindByClusterId`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sFindByClusterIdRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesCluster**](../models/kubernetescluster.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sGet

```go
var result KubernetesClusters = K8sGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Kubernetes Clusters

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
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
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

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesClusters**](../models/kubernetesclusters.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sKubeconfigGet

```go
var result KubernetesConfig = K8sKubeconfigGet(ctx, k8sClusterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Kubernetes Configuration File

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sKubeconfigGet(context.Background(), k8sClusterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sKubeconfigGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sKubeconfigGet`: KubernetesConfig
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sKubeconfigGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sKubeconfigGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesConfig**](../models/kubernetesconfig.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sNodepoolsDelete

```go
var result map[string]interface{} = K8sNodepoolsDelete(ctx, k8sClusterId, nodepoolId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete Kubernetes Node Pool

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes Node Pool
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sNodepoolsDelete(context.Background(), k8sClusterId, nodepoolId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsDelete`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |
| **nodepoolId** | **string** | The unique ID of the Kubernetes Node Pool |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

**map\[string\]interface{}**

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sNodepoolsFindById

```go
var result KubernetesNodePool = K8sNodepoolsFindById(ctx, k8sClusterId, nodepoolId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Kubernetes Node Pool

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes Node Pool
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sNodepoolsFindById(context.Background(), k8sClusterId, nodepoolId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsFindById`: KubernetesNodePool
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsFindById`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |
| **nodepoolId** | **string** | The unique ID of the Kubernetes Node Pool |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsFindByIdRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesNodePool**](../models/kubernetesnodepool.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sNodepoolsGet

```go
var result KubernetesNodePools = K8sNodepoolsGet(ctx, k8sClusterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Kubernetes Node Pools

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sNodepoolsGet(context.Background(), k8sClusterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsGet`: KubernetesNodePools
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesNodePools**](../models/kubernetesnodepools.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sNodepoolsNodesDelete

```go
var result map[string]interface{} = K8sNodepoolsNodesDelete(ctx, k8sClusterId, nodepoolId, nodeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete Kubernetes node

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes Node Pool
    nodeId := "nodeId_example" // string | The unique ID of the Kubernetes node
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sNodepoolsNodesDelete(context.Background(), k8sClusterId, nodepoolId, nodeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsNodesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsNodesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsNodesDelete`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |
| **nodepoolId** | **string** | The unique ID of the Kubernetes Node Pool |  |
| **nodeId** | **string** | The unique ID of the Kubernetes node |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsNodesDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

**map\[string\]interface{}**

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sNodepoolsNodesFindById

```go
var result KubernetesNode = K8sNodepoolsNodesFindById(ctx, k8sClusterId, nodepoolId, nodeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Kubernetes node

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes Node Pool
    nodeId := "nodeId_example" // string | The unique ID of the Kubernetes Node.
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sNodepoolsNodesFindById(context.Background(), k8sClusterId, nodepoolId, nodeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsNodesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsNodesFindById`: KubernetesNode
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsNodesFindById`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |
| **nodepoolId** | **string** | The unique ID of the Kubernetes Node Pool |  |
| **nodeId** | **string** | The unique ID of the Kubernetes Node. |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsNodesFindByIdRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesNode**](../models/kubernetesnode.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sNodepoolsNodesGet

```go
var result KubernetesNodes = K8sNodepoolsNodesGet(ctx, k8sClusterId, nodepoolId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Kubernetes nodes.

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes Node Pool
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sNodepoolsNodesGet(context.Background(), k8sClusterId, nodepoolId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsNodesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsNodesGet`: KubernetesNodes
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsNodesGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |
| **nodepoolId** | **string** | The unique ID of the Kubernetes Node Pool |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsNodesGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesNodes**](../models/kubernetesnodes.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sNodepoolsNodesReplacePost

```go
var result map[string]interface{} = K8sNodepoolsNodesReplacePost(ctx, k8sClusterId, nodepoolId, nodeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Recreate the Kubernetes node

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes Node Pool
    nodeId := "nodeId_example" // string | The unique ID of the Kubernetes Node.
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sNodepoolsNodesReplacePost(context.Background(), k8sClusterId, nodepoolId, nodeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsNodesReplacePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsNodesReplacePost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsNodesReplacePost`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |
| **nodepoolId** | **string** | The unique ID of the Kubernetes Node Pool |  |
| **nodeId** | **string** | The unique ID of the Kubernetes Node. |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsNodesReplacePostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

**map\[string\]interface{}**

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sNodepoolsPost

```go
var result KubernetesNodePool = K8sNodepoolsPost(ctx, k8sClusterId)
                      .KubernetesNodePool(kubernetesNodePool)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Kubernetes Node Pool

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    kubernetesNodePool := *openapiclient.NewKubernetesNodePoolForPost(*openapiclient.NewKubernetesNodePoolPropertiesForPost("k8s-node-pool", "1e072e52-2ed3-492f-b6b6-c6b116907521", int32(2), "AMD_OPTERON", int32(4), int32(2048), "AUTO", "HDD", int32(100))) // KubernetesNodePoolForPost | Details of the Kubernetes Node Pool
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sNodepoolsPost(context.Background(), k8sClusterId).KubernetesNodePool(kubernetesNodePool).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsPost`: KubernetesNodePool
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsPost`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **kubernetesNodePool** | [**KubernetesNodePoolForPost**](../models/kubernetesnodepoolforpost.md) | Details of the Kubernetes Node Pool |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesNodePool**](../models/kubernetesnodepool.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## K8sNodepoolsPut

```go
var result KubernetesNodePool = K8sNodepoolsPut(ctx, k8sClusterId, nodepoolId)
                      .KubernetesNodePool(kubernetesNodePool)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify Kubernetes Node Pool

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    nodepoolId := "nodepoolId_example" // string | The unique ID of the Kubernetes Node Pool
    kubernetesNodePool := *openapiclient.NewKubernetesNodePoolForPut(*openapiclient.NewKubernetesNodePoolPropertiesForPut("k8s-node-pool", int32(2))) // KubernetesNodePoolForPut | Details of the Kubernetes Node Pool
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sNodepoolsPut(context.Background(), k8sClusterId, nodepoolId).KubernetesNodePool(kubernetesNodePool).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sNodepoolsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sNodepoolsPut`: KubernetesNodePool
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sNodepoolsPut`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |
| **nodepoolId** | **string** | The unique ID of the Kubernetes Node Pool |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sNodepoolsPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **kubernetesNodePool** | [**KubernetesNodePoolForPut**](../models/kubernetesnodepoolforput.md) | Details of the Kubernetes Node Pool |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesNodePool**](../models/kubernetesnodepool.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## K8sPost

```go
var result KubernetesCluster = K8sPost(ctx)
                      .KubernetesCluster(kubernetesCluster)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create Kubernetes Cluster

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
    kubernetesCluster := *openapiclient.NewKubernetesClusterForPost(*openapiclient.NewKubernetesClusterPropertiesForPost("k8s")) // KubernetesClusterForPost | Details of the Kubernetes Cluster
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sPost(context.Background()).KubernetesCluster(kubernetesCluster).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
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

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **kubernetesCluster** | [**KubernetesClusterForPost**](../models/kubernetesclusterforpost.md) | Details of the Kubernetes Cluster |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesCluster**](../models/kubernetescluster.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## K8sPut

```go
var result KubernetesCluster = K8sPut(ctx, k8sClusterId)
                      .KubernetesCluster(kubernetesCluster)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify Kubernetes Cluster

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
    k8sClusterId := "k8sClusterId_example" // string | The unique ID of the Kubernetes Cluster
    kubernetesCluster := *openapiclient.NewKubernetesClusterForPut(*openapiclient.NewKubernetesClusterPropertiesForPut("k8s")) // KubernetesClusterForPut | Details of the Kubernetes Cluster
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sPut(context.Background(), k8sClusterId).KubernetesCluster(kubernetesCluster).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sPut`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sPut`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **k8sClusterId** | **string** | The unique ID of the Kubernetes Cluster |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **kubernetesCluster** | [**KubernetesClusterForPut**](../models/kubernetesclusterforput.md) | Details of the Kubernetes Cluster |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**KubernetesCluster**](../models/kubernetescluster.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## K8sVersionsCompatibilitiesGet

```go
var result []string = K8sVersionsCompatibilitiesGet(ctx, clusterVersion)
                      .Execute()
```

Retrieves a list of available kubernetes versions for nodepools depending on the given kubernetes version running in the cluster.

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
    clusterVersion := "clusterVersion_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sVersionsCompatibilitiesGet(context.Background(), clusterVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesApi.K8sVersionsCompatibilitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `K8sVersionsCompatibilitiesGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `KubernetesApi.K8sVersionsCompatibilitiesGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **clusterVersion** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiK8sVersionsCompatibilitiesGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |


### Return type

**\[\]string**

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sVersionsDefaultGet

```go
var result string = K8sVersionsDefaultGet(ctx)
                      .Execute()
```

Retrieve the current default kubernetes version for clusters and nodepools.

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sVersionsDefaultGet(context.Background()).Execute()
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

* **Content-Type**: Not defined
* **Accept**: application/json

## K8sVersionsGet

```go
var result []string = K8sVersionsGet(ctx)
                      .Execute()
```

Retrieve available Kubernetes versions

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesApi.K8sVersionsGet(context.Background()).Execute()
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

**\[\]string**

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

