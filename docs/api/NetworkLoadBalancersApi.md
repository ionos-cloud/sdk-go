# \NetworkLoadBalancersApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersNetworkloadbalancersDelete**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersDelete) | **Delete** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId} | Delete Network Load Balancers|
|[**DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId} | Retrieve Network Load Balancers|
|[**DatacentersNetworkloadbalancersFlowlogsDelete**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersFlowlogsDelete) | **Delete** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs/{flowLogId} | Delete NLB Flow Logs|
|[**DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs/{flowLogId} | Retrieve NLB Flow Logs|
|[**DatacentersNetworkloadbalancersFlowlogsGet**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersFlowlogsGet) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs | List NLB Flow Logs|
|[**DatacentersNetworkloadbalancersFlowlogsPatch**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersFlowlogsPatch) | **Patch** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs/{flowLogId} | Partially modify NLB Flow Logs|
|[**DatacentersNetworkloadbalancersFlowlogsPost**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersFlowlogsPost) | **Post** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs | Create NLB Flow Logs|
|[**DatacentersNetworkloadbalancersFlowlogsPut**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersFlowlogsPut) | **Put** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs/{flowLogId} | Modify NLB Flow Logs|
|[**DatacentersNetworkloadbalancersForwardingrulesDelete**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersForwardingrulesDelete) | **Delete** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules/{forwardingRuleId} | Delete NLB forwarding rules|
|[**DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules/{forwardingRuleId} | Retrieve NLB forwarding rules|
|[**DatacentersNetworkloadbalancersForwardingrulesGet**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersForwardingrulesGet) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules | List NLB forwarding rules|
|[**DatacentersNetworkloadbalancersForwardingrulesPatch**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersForwardingrulesPatch) | **Patch** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules/{forwardingRuleId} | Partially modify NLB forwarding rules|
|[**DatacentersNetworkloadbalancersForwardingrulesPost**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersForwardingrulesPost) | **Post** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules | Create NLB forwarding rules|
|[**DatacentersNetworkloadbalancersForwardingrulesPut**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersForwardingrulesPut) | **Put** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules/{forwardingRuleId} | Modify NLB forwarding rules|
|[**DatacentersNetworkloadbalancersGet**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersGet) | **Get** /datacenters/{datacenterId}/networkloadbalancers | List Network Load Balancers|
|[**DatacentersNetworkloadbalancersPatch**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersPatch) | **Patch** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId} | Partially modify Network Load Balancers|
|[**DatacentersNetworkloadbalancersPost**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersPost) | **Post** /datacenters/{datacenterId}/networkloadbalancers | Create Network Load Balancers|
|[**DatacentersNetworkloadbalancersPut**](NetworkLoadBalancersApi.md#DatacentersNetworkloadbalancersPut) | **Put** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId} | Modify Network Load Balancers|



## DatacentersNetworkloadbalancersDelete

```go
var result  = DatacentersNetworkloadbalancersDelete(ctx, datacenterId, networkLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete Network Load Balancers



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersDelete(context.Background(), datacenterId, networkLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersDeleteRequest struct via the builder pattern


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



## DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId

```go
var result NetworkLoadBalancer = DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId(ctx, datacenterId, networkLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Network Load Balancers



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId(context.Background(), datacenterId, networkLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId`: NetworkLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFindByNetworkLoadBalancerIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NetworkLoadBalancer**](../models/NetworkLoadBalancer.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNetworkloadbalancersFlowlogsDelete

```go
var result  = DatacentersNetworkloadbalancersFlowlogsDelete(ctx, datacenterId, networkLoadBalancerId, flowLogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete NLB Flow Logs



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsDelete(context.Background(), datacenterId, networkLoadBalancerId, flowLogId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |
|**flowLogId** | **string** | The unique ID of the Flow Log. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsDeleteRequest struct via the builder pattern


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



## DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId

```go
var result FlowLog = DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId(ctx, datacenterId, networkLoadBalancerId, flowLogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve NLB Flow Logs



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId(context.Background(), datacenterId, networkLoadBalancerId, flowLogId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |
|**flowLogId** | **string** | The unique ID of the Flow Log. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsFindByFlowLogIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNetworkloadbalancersFlowlogsGet

```go
var result FlowLogs = DatacentersNetworkloadbalancersFlowlogsGet(ctx, datacenterId, networkLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List NLB Flow Logs



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsGet(context.Background(), datacenterId, networkLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsGet`: FlowLogs
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLogs**](../models/FlowLogs.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNetworkloadbalancersFlowlogsPatch

```go
var result FlowLog = DatacentersNetworkloadbalancersFlowlogsPatch(ctx, datacenterId, networkLoadBalancerId, flowLogId)
                      .NetworkLoadBalancerFlowLogProperties(networkLoadBalancerFlowLogProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify NLB Flow Logs



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log.
    networkLoadBalancerFlowLogProperties := *openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key") // FlowLogProperties | The properties of the Flow Log to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPatch(context.Background(), datacenterId, networkLoadBalancerId, flowLogId).NetworkLoadBalancerFlowLogProperties(networkLoadBalancerFlowLogProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsPatch`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |
|**flowLogId** | **string** | The unique ID of the Flow Log. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **networkLoadBalancerFlowLogProperties** | [**FlowLogProperties**](../models/FlowLogProperties.md) | The properties of the Flow Log to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNetworkloadbalancersFlowlogsPost

```go
var result FlowLog = DatacentersNetworkloadbalancersFlowlogsPost(ctx, datacenterId, networkLoadBalancerId)
                      .NetworkLoadBalancerFlowLog(networkLoadBalancerFlowLog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create NLB Flow Logs



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    networkLoadBalancerFlowLog := *openapiclient.NewFlowLog(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLog | The Flow Log to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPost(context.Background(), datacenterId, networkLoadBalancerId).NetworkLoadBalancerFlowLog(networkLoadBalancerFlowLog).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsPost`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **networkLoadBalancerFlowLog** | [**FlowLog**](../models/FlowLog.md) | The Flow Log to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNetworkloadbalancersFlowlogsPut

```go
var result FlowLog = DatacentersNetworkloadbalancersFlowlogsPut(ctx, datacenterId, networkLoadBalancerId, flowLogId)
                      .NetworkLoadBalancerFlowLog(networkLoadBalancerFlowLog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify NLB Flow Logs



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log.
    networkLoadBalancerFlowLog := *openapiclient.NewFlowLogPut(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLogPut | The modified NLB Flow Log.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPut(context.Background(), datacenterId, networkLoadBalancerId, flowLogId).NetworkLoadBalancerFlowLog(networkLoadBalancerFlowLog).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsPut`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |
|**flowLogId** | **string** | The unique ID of the Flow Log. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **networkLoadBalancerFlowLog** | [**FlowLogPut**](../models/FlowLogPut.md) | The modified NLB Flow Log. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNetworkloadbalancersForwardingrulesDelete

```go
var result  = DatacentersNetworkloadbalancersForwardingrulesDelete(ctx, datacenterId, networkLoadBalancerId, forwardingRuleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete NLB forwarding rules



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesDelete(context.Background(), datacenterId, networkLoadBalancerId, forwardingRuleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |
|**forwardingRuleId** | **string** | The unique ID of the forwarding rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesDeleteRequest struct via the builder pattern


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



## DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId

```go
var result NetworkLoadBalancerForwardingRule = DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId(ctx, datacenterId, networkLoadBalancerId, forwardingRuleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve NLB forwarding rules



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId(context.Background(), datacenterId, networkLoadBalancerId, forwardingRuleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId`: NetworkLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |
|**forwardingRuleId** | **string** | The unique ID of the forwarding rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NetworkLoadBalancerForwardingRule**](../models/NetworkLoadBalancerForwardingRule.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNetworkloadbalancersForwardingrulesGet

```go
var result NetworkLoadBalancerForwardingRules = DatacentersNetworkloadbalancersForwardingrulesGet(ctx, datacenterId, networkLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List NLB forwarding rules



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesGet(context.Background(), datacenterId, networkLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesGet`: NetworkLoadBalancerForwardingRules
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NetworkLoadBalancerForwardingRules**](../models/NetworkLoadBalancerForwardingRules.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNetworkloadbalancersForwardingrulesPatch

```go
var result NetworkLoadBalancerForwardingRule = DatacentersNetworkloadbalancersForwardingrulesPatch(ctx, datacenterId, networkLoadBalancerId, forwardingRuleId)
                      .NetworkLoadBalancerForwardingRuleProperties(networkLoadBalancerForwardingRuleProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify NLB forwarding rules



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule.
    networkLoadBalancerForwardingRuleProperties := *openapiclient.NewNetworkLoadBalancerForwardingRuleProperties("My Network Load Balancer forwarding rule", "ROUND_ROBIN", "HTTP", "81.173.1.2", int32(8080), []openapiclient.NetworkLoadBalancerForwardingRuleTarget{*openapiclient.NewNetworkLoadBalancerForwardingRuleTarget("22.231.2.2", int32(8080), int32(123))}) // NetworkLoadBalancerForwardingRuleProperties | The properties of the forwarding rule to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPatch(context.Background(), datacenterId, networkLoadBalancerId, forwardingRuleId).NetworkLoadBalancerForwardingRuleProperties(networkLoadBalancerForwardingRuleProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesPatch`: NetworkLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |
|**forwardingRuleId** | **string** | The unique ID of the forwarding rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **networkLoadBalancerForwardingRuleProperties** | [**NetworkLoadBalancerForwardingRuleProperties**](../models/NetworkLoadBalancerForwardingRuleProperties.md) | The properties of the forwarding rule to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NetworkLoadBalancerForwardingRule**](../models/NetworkLoadBalancerForwardingRule.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNetworkloadbalancersForwardingrulesPost

```go
var result NetworkLoadBalancerForwardingRule = DatacentersNetworkloadbalancersForwardingrulesPost(ctx, datacenterId, networkLoadBalancerId)
                      .NetworkLoadBalancerForwardingRule(networkLoadBalancerForwardingRule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create NLB forwarding rules



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    networkLoadBalancerForwardingRule := *openapiclient.NewNetworkLoadBalancerForwardingRule(*openapiclient.NewNetworkLoadBalancerForwardingRuleProperties("My Network Load Balancer forwarding rule", "ROUND_ROBIN", "HTTP", "81.173.1.2", int32(8080), []openapiclient.NetworkLoadBalancerForwardingRuleTarget{*openapiclient.NewNetworkLoadBalancerForwardingRuleTarget("22.231.2.2", int32(8080), int32(123))})) // NetworkLoadBalancerForwardingRule | The forwarding rule to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPost(context.Background(), datacenterId, networkLoadBalancerId).NetworkLoadBalancerForwardingRule(networkLoadBalancerForwardingRule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesPost`: NetworkLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **networkLoadBalancerForwardingRule** | [**NetworkLoadBalancerForwardingRule**](../models/NetworkLoadBalancerForwardingRule.md) | The forwarding rule to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NetworkLoadBalancerForwardingRule**](../models/NetworkLoadBalancerForwardingRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNetworkloadbalancersForwardingrulesPut

```go
var result NetworkLoadBalancerForwardingRule = DatacentersNetworkloadbalancersForwardingrulesPut(ctx, datacenterId, networkLoadBalancerId, forwardingRuleId)
                      .NetworkLoadBalancerForwardingRule(networkLoadBalancerForwardingRule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify NLB forwarding rules



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule.
    networkLoadBalancerForwardingRule := *openapiclient.NewNetworkLoadBalancerForwardingRulePut(*openapiclient.NewNetworkLoadBalancerForwardingRuleProperties("My Network Load Balancer forwarding rule", "ROUND_ROBIN", "HTTP", "81.173.1.2", int32(8080), []openapiclient.NetworkLoadBalancerForwardingRuleTarget{*openapiclient.NewNetworkLoadBalancerForwardingRuleTarget("22.231.2.2", int32(8080), int32(123))})) // NetworkLoadBalancerForwardingRulePut | The modified NLB forwarding rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPut(context.Background(), datacenterId, networkLoadBalancerId, forwardingRuleId).NetworkLoadBalancerForwardingRule(networkLoadBalancerForwardingRule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesPut`: NetworkLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |
|**forwardingRuleId** | **string** | The unique ID of the forwarding rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **networkLoadBalancerForwardingRule** | [**NetworkLoadBalancerForwardingRulePut**](../models/NetworkLoadBalancerForwardingRulePut.md) | The modified NLB forwarding rule. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NetworkLoadBalancerForwardingRule**](../models/NetworkLoadBalancerForwardingRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNetworkloadbalancersGet

```go
var result NetworkLoadBalancers = DatacentersNetworkloadbalancersGet(ctx, datacenterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List Network Load Balancers



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
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersGet`: NetworkLoadBalancers
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 1000]|

### Return type

[**NetworkLoadBalancers**](../models/NetworkLoadBalancers.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNetworkloadbalancersPatch

```go
var result NetworkLoadBalancer = DatacentersNetworkloadbalancersPatch(ctx, datacenterId, networkLoadBalancerId)
                      .NetworkLoadBalancerProperties(networkLoadBalancerProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify Network Load Balancers



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    networkLoadBalancerProperties := *openapiclient.NewNetworkLoadBalancerProperties("My Network Load Balancer", int32(1), int32(2)) // NetworkLoadBalancerProperties | The properties of the Network Load Balancer to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPatch(context.Background(), datacenterId, networkLoadBalancerId).NetworkLoadBalancerProperties(networkLoadBalancerProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersPatch`: NetworkLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **networkLoadBalancerProperties** | [**NetworkLoadBalancerProperties**](../models/NetworkLoadBalancerProperties.md) | The properties of the Network Load Balancer to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NetworkLoadBalancer**](../models/NetworkLoadBalancer.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNetworkloadbalancersPost

```go
var result NetworkLoadBalancer = DatacentersNetworkloadbalancersPost(ctx, datacenterId)
                      .NetworkLoadBalancer(networkLoadBalancer)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create Network Load Balancers



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
    networkLoadBalancer := *openapiclient.NewNetworkLoadBalancer(*openapiclient.NewNetworkLoadBalancerProperties("My Network Load Balancer", int32(1), int32(2))) // NetworkLoadBalancer | The Network Load Balancer to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPost(context.Background(), datacenterId).NetworkLoadBalancer(networkLoadBalancer).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersPost`: NetworkLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **networkLoadBalancer** | [**NetworkLoadBalancer**](../models/NetworkLoadBalancer.md) | The Network Load Balancer to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NetworkLoadBalancer**](../models/NetworkLoadBalancer.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNetworkloadbalancersPut

```go
var result NetworkLoadBalancer = DatacentersNetworkloadbalancersPut(ctx, datacenterId, networkLoadBalancerId)
                      .NetworkLoadBalancer(networkLoadBalancer)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify Network Load Balancers



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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer.
    networkLoadBalancer := *openapiclient.NewNetworkLoadBalancerPut(*openapiclient.NewNetworkLoadBalancerProperties("My Network Load Balancer", int32(1), int32(2))) // NetworkLoadBalancerPut | The modified Network Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPut(context.Background(), datacenterId, networkLoadBalancerId).NetworkLoadBalancer(networkLoadBalancer).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersPut`: NetworkLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **networkLoadBalancer** | [**NetworkLoadBalancerPut**](../models/NetworkLoadBalancerPut.md) | The modified Network Load Balancer. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NetworkLoadBalancer**](../models/NetworkLoadBalancer.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


