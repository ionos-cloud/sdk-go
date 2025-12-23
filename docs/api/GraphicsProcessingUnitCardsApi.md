# \GraphicsProcessingUnitCardsApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersServersGPUsFindById**](GraphicsProcessingUnitCardsApi.md#DatacentersServersGPUsFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/gpus/{gpuId} | Retrieve GPUs by ID|
|[**DatacentersServersGPUsGet**](GraphicsProcessingUnitCardsApi.md#DatacentersServersGPUsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/gpus | List GPUs|



## DatacentersServersGPUsFindById

```go
var result Gpu = DatacentersServersGPUsFindById(ctx, datacenterId, serverId, gpuId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve GPUs by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := TODO // string | The unique ID of the data center.
    serverId := TODO // string | The unique ID of the server.
    gpuId := TODO // string | The unique ID of the Graphics Processing Unit (GPU) card.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]   - depth=0: Only direct properties are included; children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.GraphicsProcessingUnitCardsApi.DatacentersServersGPUsFindById(context.Background(), datacenterId, serverId, gpuId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphicsProcessingUnitCardsApi.DatacentersServersGPUsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersGPUsFindById`: Gpu
    fmt.Fprintf(os.Stdout, "Response from `GraphicsProcessingUnitCardsApi.DatacentersServersGPUsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | [**string**](../models/.md) | The unique ID of the data center. | |
|**serverId** | [**string**](../models/.md) | The unique ID of the server. | |
|**gpuId** | [**string**](../models/.md) | The unique ID of the Graphics Processing Unit (GPU) card. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersGPUsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Gpu**](../models/Gpu.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersGPUsGet

```go
var result Gpus = DatacentersServersGPUsGet(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List GPUs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := TODO // string | The unique ID of the data center.
    serverId := TODO // string | The unique ID of the server.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]   - depth=0: Only direct properties are included; children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.GraphicsProcessingUnitCardsApi.DatacentersServersGPUsGet(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GraphicsProcessingUnitCardsApi.DatacentersServersGPUsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersGPUsGet`: Gpus
    fmt.Fprintf(os.Stdout, "Response from `GraphicsProcessingUnitCardsApi.DatacentersServersGPUsGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | [**string**](../models/.md) | The unique ID of the data center. | |
|**serverId** | [**string**](../models/.md) | The unique ID of the server. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersGPUsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Gpus**](../models/Gpus.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


