# \NetworkInterfacesApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersServersNicsDelete**](NetworkInterfacesApi.md#DatacentersServersNicsDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Delete NICs|
|[**DatacentersServersNicsFindById**](NetworkInterfacesApi.md#DatacentersServersNicsFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Retrieve NICs|
|[**DatacentersServersNicsGet**](NetworkInterfacesApi.md#DatacentersServersNicsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics | List NICs|
|[**DatacentersServersNicsPatch**](NetworkInterfacesApi.md#DatacentersServersNicsPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Partially modify NICs|
|[**DatacentersServersNicsPost**](NetworkInterfacesApi.md#DatacentersServersNicsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics | Create NICs|
|[**DatacentersServersNicsPut**](NetworkInterfacesApi.md#DatacentersServersNicsPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Modify NICs|



## DatacentersServersNicsDelete

```go
var result  = DatacentersServersNicsDelete(ctx, datacenterId, serverId, nicId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete NICs



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    nicId := "nicId_example" // string | The unique ID of the NIC.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkInterfacesApi.DatacentersServersNicsDelete(context.Background(), datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfacesApi.DatacentersServersNicsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**nicId** | **string** | The unique ID of the NIC. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsDeleteRequest struct via the builder pattern


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



## DatacentersServersNicsFindById

```go
var result Nic = DatacentersServersNicsFindById(ctx, datacenterId, serverId, nicId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve NICs



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    nicId := "nicId_example" // string | The unique ID of the NIC.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkInterfacesApi.DatacentersServersNicsFindById(context.Background(), datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfacesApi.DatacentersServersNicsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFindById`: Nic
    fmt.Fprintf(os.Stdout, "Response from `NetworkInterfacesApi.DatacentersServersNicsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**nicId** | **string** | The unique ID of the NIC. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Nic**](../models/Nic.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersNicsGet

```go
var result Nics = DatacentersServersNicsGet(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List NICs



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkInterfacesApi.DatacentersServersNicsGet(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfacesApi.DatacentersServersNicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsGet`: Nics
    fmt.Fprintf(os.Stdout, "Response from `NetworkInterfacesApi.DatacentersServersNicsGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 1000]|

### Return type

[**Nics**](../models/Nics.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersNicsPatch

```go
var result Nic = DatacentersServersNicsPatch(ctx, datacenterId, serverId, nicId)
                      .Nic(nic)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify NICs



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    nicId := "nicId_example" // string | The unique ID of the NIC.
    nic := *openapiclient.NewNicProperties(int32(2)) // NicProperties | The properties of the NIC to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkInterfacesApi.DatacentersServersNicsPatch(context.Background(), datacenterId, serverId, nicId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfacesApi.DatacentersServersNicsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsPatch`: Nic
    fmt.Fprintf(os.Stdout, "Response from `NetworkInterfacesApi.DatacentersServersNicsPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**nicId** | **string** | The unique ID of the NIC. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **nic** | [**NicProperties**](../models/NicProperties.md) | The properties of the NIC to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Nic**](../models/Nic.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersNicsPost

```go
var result Nic = DatacentersServersNicsPost(ctx, datacenterId, serverId)
                      .Nic(nic)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create NICs



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    nic := *openapiclient.NewNic(*openapiclient.NewNicProperties(int32(2))) // Nic | The NIC to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkInterfacesApi.DatacentersServersNicsPost(context.Background(), datacenterId, serverId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfacesApi.DatacentersServersNicsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsPost`: Nic
    fmt.Fprintf(os.Stdout, "Response from `NetworkInterfacesApi.DatacentersServersNicsPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **nic** | [**Nic**](../models/Nic.md) | The NIC to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Nic**](../models/Nic.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersNicsPut

```go
var result Nic = DatacentersServersNicsPut(ctx, datacenterId, serverId, nicId)
                      .Nic(nic)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify NICs



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    nicId := "nicId_example" // string | The unique ID of the NIC.
    nic := *openapiclient.NewNicPut(*openapiclient.NewNicProperties(int32(2))) // NicPut | The modified NIC
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkInterfacesApi.DatacentersServersNicsPut(context.Background(), datacenterId, serverId, nicId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfacesApi.DatacentersServersNicsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsPut`: Nic
    fmt.Fprintf(os.Stdout, "Response from `NetworkInterfacesApi.DatacentersServersNicsPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**nicId** | **string** | The unique ID of the NIC. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **nic** | [**NicPut**](../models/NicPut.md) | The modified NIC | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Nic**](../models/Nic.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


