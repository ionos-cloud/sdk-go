# \IPBlocksApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**IpblocksDelete**](IPBlocksApi.md#IpblocksDelete) | **Delete** /ipblocks/{ipblockId} | Delete IP blocks|
|[**IpblocksFindById**](IPBlocksApi.md#IpblocksFindById) | **Get** /ipblocks/{ipblockId} | Retrieve IP blocks|
|[**IpblocksGet**](IPBlocksApi.md#IpblocksGet) | **Get** /ipblocks | List IP blocks |
|[**IpblocksPatch**](IPBlocksApi.md#IpblocksPatch) | **Patch** /ipblocks/{ipblockId} | Partially modify IP blocks|
|[**IpblocksPost**](IPBlocksApi.md#IpblocksPost) | **Post** /ipblocks | Reserve IP blocks|
|[**IpblocksPut**](IPBlocksApi.md#IpblocksPut) | **Put** /ipblocks/{ipblockId} | Modify IP blocks|



## IpblocksDelete

```go
var result  = IpblocksDelete(ctx, ipblockId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete IP blocks



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
    ipblockId := "ipblockId_example" // string | The unique ID of the IP block.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpblocksDelete(context.Background(), ipblockId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpblocksDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**ipblockId** | **string** | The unique ID of the IP block. | |

### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksDeleteRequest struct via the builder pattern


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



## IpblocksFindById

```go
var result IpBlock = IpblocksFindById(ctx, ipblockId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve IP blocks



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
    ipblockId := "ipblockId_example" // string | The unique ID of the IP block.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpblocksFindById(context.Background(), ipblockId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpblocksFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksFindById`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpblocksFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**ipblockId** | **string** | The unique ID of the IP block. | |

### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**IpBlock**](../models/IpBlock.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpblocksGet

```go
var result IpBlocks = IpblocksGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List IP blocks 



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
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpblocksGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpblocksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksGet`: IpBlocks
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpblocksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 100]|

### Return type

[**IpBlocks**](../models/IpBlocks.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpblocksPatch

```go
var result IpBlock = IpblocksPatch(ctx, ipblockId)
                      .Ipblock(ipblock)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify IP blocks



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
    ipblockId := "ipblockId_example" // string | The unique ID of the IP block.
    ipblock := *openapiclient.NewIpBlockProperties("us/las", int32(5)) // IpBlockProperties | The properties of the IP block to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpblocksPatch(context.Background(), ipblockId).Ipblock(ipblock).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpblocksPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksPatch`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpblocksPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**ipblockId** | **string** | The unique ID of the IP block. | |

### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipblock** | [**IpBlockProperties**](../models/IpBlockProperties.md) | The properties of the IP block to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**IpBlock**](../models/IpBlock.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## IpblocksPost

```go
var result IpBlock = IpblocksPost(ctx)
                      .Ipblock(ipblock)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Reserve IP blocks



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
    ipblock := *openapiclient.NewIpBlock(*openapiclient.NewIpBlockProperties("us/las", int32(5))) // IpBlock | The IP block to be reserved.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpblocksPost(context.Background()).Ipblock(ipblock).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpblocksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksPost`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpblocksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipblock** | [**IpBlock**](../models/IpBlock.md) | The IP block to be reserved. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**IpBlock**](../models/IpBlock.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## IpblocksPut

```go
var result IpBlock = IpblocksPut(ctx, ipblockId)
                      .Ipblock(ipblock)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify IP blocks



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
    ipblockId := "ipblockId_example" // string | The unique ID of the IP block.
    ipblock := *openapiclient.NewIpBlock(*openapiclient.NewIpBlockProperties("us/las", int32(5))) // IpBlock | The modified IP block.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPBlocksApi.IpblocksPut(context.Background(), ipblockId).Ipblock(ipblock).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPBlocksApi.IpblocksPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IpblocksPut`: IpBlock
    fmt.Fprintf(os.Stdout, "Response from `IPBlocksApi.IpblocksPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**ipblockId** | **string** | The unique ID of the IP block. | |

### Other Parameters

Other parameters are passed through a pointer to a apiIpblocksPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ipblock** | [**IpBlock**](../models/IpBlock.md) | The modified IP block. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**IpBlock**](../models/IpBlock.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


