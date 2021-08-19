# \FlowLogsApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersServersNicsFlowlogsDelete**](FlowLogsApi.md#DatacentersServersNicsFlowlogsDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/flowlogs/{flowlogId} | Delete a Flow Log|
|[**DatacentersServersNicsFlowlogsFindById**](FlowLogsApi.md#DatacentersServersNicsFlowlogsFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/flowlogs/{flowlogId} | Retrieve a Flow Log|
|[**DatacentersServersNicsFlowlogsGet**](FlowLogsApi.md#DatacentersServersNicsFlowlogsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/flowlogs | List Flow Logs|
|[**DatacentersServersNicsFlowlogsPatch**](FlowLogsApi.md#DatacentersServersNicsFlowlogsPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/flowlogs/{flowlogId} | Partially update a Flow Log|
|[**DatacentersServersNicsFlowlogsPost**](FlowLogsApi.md#DatacentersServersNicsFlowlogsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/flowlogs | Create a Flow Log|
|[**DatacentersServersNicsFlowlogsPut**](FlowLogsApi.md#DatacentersServersNicsFlowlogsPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/flowlogs/{flowlogId} | Modify a Flow Log|



## DatacentersServersNicsFlowlogsDelete

```go
var result  = DatacentersServersNicsFlowlogsDelete(ctx, datacenterId, serverId, nicId, flowlogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Delete a Flow Log



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
    datacenterId := "datacenterId_example" // string | The unique ID of the datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    nicId := "nicId_example" // string | The unique ID of the NIC
    flowlogId := "flowlogId_example" // string | The unique ID of the Flow Log
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowLogsApi.DatacentersServersNicsFlowlogsDelete(context.Background(), datacenterId, serverId, nicId, flowlogId).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowLogsApi.DatacentersServersNicsFlowlogsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the Server | |
|**nicId** | **string** | The unique ID of the NIC | |
|**flowlogId** | **string** | The unique ID of the Flow Log | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFlowlogsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersNicsFlowlogsFindById

```go
var result FlowLog = DatacentersServersNicsFlowlogsFindById(ctx, datacenterId, serverId, nicId, flowlogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Retrieve a Flow Log



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
    datacenterId := "datacenterId_example" // string | The unique ID of the datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    nicId := "nicId_example" // string | The unique ID of the NIC
    flowlogId := "flowlogId_example" // string | The unique ID of the Flow Log
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowLogsApi.DatacentersServersNicsFlowlogsFindById(context.Background(), datacenterId, serverId, nicId, flowlogId).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowLogsApi.DatacentersServersNicsFlowlogsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFlowlogsFindById`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `FlowLogsApi.DatacentersServersNicsFlowlogsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the Server | |
|**nicId** | **string** | The unique ID of the NIC | |
|**flowlogId** | **string** | The unique ID of the Flow Log | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFlowlogsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FlowLog**](FlowLog.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersNicsFlowlogsGet

```go
var result FlowLogs = DatacentersServersNicsFlowlogsGet(ctx, datacenterId, serverId, nicId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List Flow Logs



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
    datacenterId := "datacenterId_example" // string | The unique ID of the datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    nicId := "nicId_example" // string | The unique ID of the NIC
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with limit for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with offset for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowLogsApi.DatacentersServersNicsFlowlogsGet(context.Background(), datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowLogsApi.DatacentersServersNicsFlowlogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFlowlogsGet`: FlowLogs
    fmt.Fprintf(os.Stdout, "Response from `FlowLogsApi.DatacentersServersNicsFlowlogsGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the Server | |
|**nicId** | **string** | The unique ID of the NIC | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFlowlogsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with limit for pagination) | [default to 0]|
| **limit** | **int32** | the maximum number of elements to return (use together with offset for pagination) | [default to 1000]|

### Return type

[**FlowLogs**](FlowLogs.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersNicsFlowlogsPatch

```go
var result FlowLog = DatacentersServersNicsFlowlogsPatch(ctx, datacenterId, serverId, nicId, flowlogId)
                      .Flowlog(flowlog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Partially update a Flow Log



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
    datacenterId := "datacenterId_example" // string | The unique ID of the datacenter
    serverId := "serverId_example" // string | The unique ID of the server
    nicId := "nicId_example" // string | The unique ID of the NIC
    flowlogId := "flowlogId_example" // string | The unique ID of the Flow Log
    flowlog := *openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key") // FlowLogProperties | Modified Flow Log
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowLogsApi.DatacentersServersNicsFlowlogsPatch(context.Background(), datacenterId, serverId, nicId, flowlogId).Flowlog(flowlog).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowLogsApi.DatacentersServersNicsFlowlogsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFlowlogsPatch`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `FlowLogsApi.DatacentersServersNicsFlowlogsPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the server | |
|**nicId** | **string** | The unique ID of the NIC | |
|**flowlogId** | **string** | The unique ID of the Flow Log | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFlowlogsPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **flowlog** | [**FlowLogProperties**](FlowLogProperties.md) | Modified Flow Log | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FlowLog**](FlowLog.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersNicsFlowlogsPost

```go
var result FlowLog = DatacentersServersNicsFlowlogsPost(ctx, datacenterId, serverId, nicId)
                      .Flowlog(flowlog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Create a Flow Log



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
    datacenterId := "datacenterId_example" // string | The unique ID of the datacenter
    serverId := "serverId_example" // string | The unique ID of the server
    nicId := "nicId_example" // string | The unique ID of the NIC
    flowlog := *openapiclient.NewFlowLog(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLog | Flow Log to be created
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowLogsApi.DatacentersServersNicsFlowlogsPost(context.Background(), datacenterId, serverId, nicId).Flowlog(flowlog).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowLogsApi.DatacentersServersNicsFlowlogsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFlowlogsPost`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `FlowLogsApi.DatacentersServersNicsFlowlogsPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the server | |
|**nicId** | **string** | The unique ID of the NIC | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFlowlogsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **flowlog** | [**FlowLog**](FlowLog.md) | Flow Log to be created | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FlowLog**](FlowLog.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersNicsFlowlogsPut

```go
var result FlowLog = DatacentersServersNicsFlowlogsPut(ctx, datacenterId, serverId, nicId, flowlogId)
                      .Flowlog(flowlog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Modify a Flow Log



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
    datacenterId := "datacenterId_example" // string | The unique ID of the datacenter
    serverId := "serverId_example" // string | The unique ID of the server
    nicId := "nicId_example" // string | The unique ID of the NIC
    flowlogId := "flowlogId_example" // string | The unique ID of the Flow Log
    flowlog := *openapiclient.NewFlowLogPut(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLogPut | Modified Flow Log
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlowLogsApi.DatacentersServersNicsFlowlogsPut(context.Background(), datacenterId, serverId, nicId, flowlogId).Flowlog(flowlog).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowLogsApi.DatacentersServersNicsFlowlogsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFlowlogsPut`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `FlowLogsApi.DatacentersServersNicsFlowlogsPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the server | |
|**nicId** | **string** | The unique ID of the NIC | |
|**flowlogId** | **string** | The unique ID of the Flow Log | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFlowlogsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **flowlog** | [**FlowLogPut**](FlowLogPut.md) | Modified Flow Log | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FlowLog**](FlowLog.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


