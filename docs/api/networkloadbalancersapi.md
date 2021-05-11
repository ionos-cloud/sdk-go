# NetworkLoadBalancersApi

All URIs are relative to [https://api.ionos.com/cloudapi/v6](https://api.ionos.com/cloudapi/v6)

| Method | HTTP request | Description |
| :--- | :--- | :--- |
| [**DatacentersNetworkloadbalancersDelete**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersDelete) | **Delete** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId} | Remove an Network Load Balancer |
| [**DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId} | Retrieve an Network Load Balancer |
| [**DatacentersNetworkloadbalancersFlowlogsDelete**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersFlowlogsDelete) | **Delete** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs/{flowLogId} | Remove Flow Log from Network Load Balancer |
| [**DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs/{flowLogId} | Retrieve a Flow Log of the Network Load Balancer |
| [**DatacentersNetworkloadbalancersFlowlogsGet**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersFlowlogsGet) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs | List Network Load Balancer Flow Logs |
| [**DatacentersNetworkloadbalancersFlowlogsPatch**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersFlowlogsPatch) | **Patch** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs/{flowLogId} | Partially modify a Flow Log of the Network Load Balancer |
| [**DatacentersNetworkloadbalancersFlowlogsPost**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersFlowlogsPost) | **Post** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs | Add a Network Load Balancer Flow Log |
| [**DatacentersNetworkloadbalancersFlowlogsPut**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersFlowlogsPut) | **Put** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/flowlogs/{flowLogId} | Modify a Flow Log of the Network Load Balancer |
| [**DatacentersNetworkloadbalancersForwardingrulesDelete**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersForwardingrulesDelete) | **Delete** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules/{forwardingRuleId} | Remove Forwarding Rule from Network Load Balancer |
| [**DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules/{forwardingRuleId} | Retrieve a Forwarding Rule of the Network Load Balancer |
| [**DatacentersNetworkloadbalancersForwardingrulesGet**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersForwardingrulesGet) | **Get** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules | List Network Load Balancer Forwarding Rules |
| [**DatacentersNetworkloadbalancersForwardingrulesPatch**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersForwardingrulesPatch) | **Patch** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules/{forwardingRuleId} | Partially modify a forwarding rule of the Network Load Balancer |
| [**DatacentersNetworkloadbalancersForwardingrulesPost**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersForwardingrulesPost) | **Post** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules | Add a Network Load Balancer Forwarding Rule |
| [**DatacentersNetworkloadbalancersForwardingrulesPut**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersForwardingrulesPut) | **Put** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId}/forwardingrules/{forwardingRuleId} | Modify a forwarding rule of the Network Load Balancer |
| [**DatacentersNetworkloadbalancersGet**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersGet) | **Get** /datacenters/{datacenterId}/networkloadbalancers | List Network Load Balancers |
| [**DatacentersNetworkloadbalancersPatch**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersPatch) | **Patch** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId} | Partially update an Network Load Balancer |
| [**DatacentersNetworkloadbalancersPost**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersPost) | **Post** /datacenters/{datacenterId}/networkloadbalancers | Create an Network Load Balancer |
| [**DatacentersNetworkloadbalancersPut**](networkloadbalancersapi.md#DatacentersNetworkloadbalancersPut) | **Put** /datacenters/{datacenterId}/networkloadbalancers/{networkLoadBalancerId} | Update an Network Load Balancer |

## DatacentersNetworkloadbalancersDelete

```go
var result map[string]interface{} = DatacentersNetworkloadbalancersDelete(ctx, datacenterId, networkLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Remove an Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersDelete(context.Background(), datacenterId, networkLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersDelete`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersDeleteRequest struct via the builder pattern

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

## DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId

```go
var result NetworkLoadBalancer = DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId(ctx, datacenterId, networkLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve an Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId(context.Background(), datacenterId, networkLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId`: NetworkLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFindByNetworkLoadBalancerId`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFindByNetworkLoadBalancerIdRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**NetworkLoadBalancer**](../models/networkloadbalancer.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersNetworkloadbalancersFlowlogsDelete

```go
var result map[string]interface{} = DatacentersNetworkloadbalancersFlowlogsDelete(ctx, datacenterId, networkLoadBalancerId, flowLogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Remove Flow Log from Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    flowLogId := "flowLogId_example" // string | The unique ID of the flow log
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsDelete(context.Background(), datacenterId, networkLoadBalancerId, flowLogId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsDelete`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |
| **flowLogId** | **string** | The unique ID of the flow log |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsDeleteRequest struct via the builder pattern

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

## DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId

```go
var result FlowLog = DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId(ctx, datacenterId, networkLoadBalancerId, flowLogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a Flow Log of the Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId(context.Background(), datacenterId, networkLoadBalancerId, flowLogId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsFindByFlowLogId`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |
| **flowLogId** | **string** | The unique ID of the Flow Log |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsFindByFlowLogIdRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**FlowLog**](../models/flowlog.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersNetworkloadbalancersFlowlogsGet

```go
var result FlowLogs = DatacentersNetworkloadbalancersFlowlogsGet(ctx, datacenterId, networkLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Network Load Balancer Flow Logs

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsGet(context.Background(), datacenterId, networkLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsGet`: FlowLogs
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**FlowLogs**](../models/flowlogs.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersNetworkloadbalancersFlowlogsPatch

```go
var result FlowLog = DatacentersNetworkloadbalancersFlowlogsPatch(ctx, datacenterId, networkLoadBalancerId, flowLogId)
                      .NetworkLoadBalancerFlowLogProperties(networkLoadBalancerFlowLogProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify a Flow Log of the Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log
    networkLoadBalancerFlowLogProperties := *openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key") // FlowLogProperties | Properties of a Flow Log to be updated
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPatch(context.Background(), datacenterId, networkLoadBalancerId, flowLogId).NetworkLoadBalancerFlowLogProperties(networkLoadBalancerFlowLogProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsPatch`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPatch`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |
| **flowLogId** | **string** | The unique ID of the Flow Log |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsPatchRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **networkLoadBalancerFlowLogProperties** | [**FlowLogProperties**](../models/flowlogproperties.md) | Properties of a Flow Log to be updated |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**FlowLog**](../models/flowlog.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersNetworkloadbalancersFlowlogsPost

```go
var result FlowLog = DatacentersNetworkloadbalancersFlowlogsPost(ctx, datacenterId, networkLoadBalancerId)
                      .NetworkLoadBalancerFlowLog(networkLoadBalancerFlowLog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Add a Network Load Balancer Flow Log

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    networkLoadBalancerFlowLog := *openapiclient.NewFlowLog(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLog | Flow Log to add
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPost(context.Background(), datacenterId, networkLoadBalancerId).NetworkLoadBalancerFlowLog(networkLoadBalancerFlowLog).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsPost`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPost`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **networkLoadBalancerFlowLog** | [**FlowLog**](../models/flowlog.md) | Flow Log to add |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**FlowLog**](../models/flowlog.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## DatacentersNetworkloadbalancersFlowlogsPut

```go
var result FlowLog = DatacentersNetworkloadbalancersFlowlogsPut(ctx, datacenterId, networkLoadBalancerId, flowLogId)
                      .NetworkLoadBalancerFlowLog(networkLoadBalancerFlowLog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Flow Log of the Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log
    networkLoadBalancerFlowLog := *openapiclient.NewFlowLogPut(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLogPut | Modified Network Load Balancer Flow Log
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPut(context.Background(), datacenterId, networkLoadBalancerId, flowLogId).NetworkLoadBalancerFlowLog(networkLoadBalancerFlowLog).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersFlowlogsPut`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersFlowlogsPut`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |
| **flowLogId** | **string** | The unique ID of the Flow Log |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersFlowlogsPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **networkLoadBalancerFlowLog** | [**FlowLogPut**](../models/flowlogput.md) | Modified Network Load Balancer Flow Log |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**FlowLog**](../models/flowlog.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## DatacentersNetworkloadbalancersForwardingrulesDelete

```go
var result map[string]interface{} = DatacentersNetworkloadbalancersForwardingrulesDelete(ctx, datacenterId, networkLoadBalancerId, forwardingRuleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Remove Forwarding Rule from Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesDelete(context.Background(), datacenterId, networkLoadBalancerId, forwardingRuleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesDelete`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |
| **forwardingRuleId** | **string** | The unique ID of the forwarding rule |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesDeleteRequest struct via the builder pattern

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

## DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId

```go
var result NetworkLoadBalancerForwardingRule = DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId(ctx, datacenterId, networkLoadBalancerId, forwardingRuleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a Forwarding Rule of the Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId(context.Background(), datacenterId, networkLoadBalancerId, forwardingRuleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId`: NetworkLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleId`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |
| **forwardingRuleId** | **string** | The unique ID of the forwarding rule |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesFindByForwardingRuleIdRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**NetworkLoadBalancerForwardingRule**](../models/networkloadbalancerforwardingrule.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersNetworkloadbalancersForwardingrulesGet

```go
var result NetworkLoadBalancerForwardingRules = DatacentersNetworkloadbalancersForwardingrulesGet(ctx, datacenterId, networkLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Network Load Balancer Forwarding Rules

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesGet(context.Background(), datacenterId, networkLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesGet`: NetworkLoadBalancerForwardingRules
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**NetworkLoadBalancerForwardingRules**](../models/networkloadbalancerforwardingrules.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersNetworkloadbalancersForwardingrulesPatch

```go
var result NetworkLoadBalancerForwardingRule = DatacentersNetworkloadbalancersForwardingrulesPatch(ctx, datacenterId, networkLoadBalancerId, forwardingRuleId)
                      .NetworkLoadBalancerForwardingRuleProperties(networkLoadBalancerForwardingRuleProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify a forwarding rule of the Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule
    networkLoadBalancerForwardingRuleProperties := *openapiclient.NewNetworkLoadBalancerForwardingRuleProperties("My Network Load Balancer forwarding rule", "ROUND_ROBIN", "TCP", "81.173.1.2", int32(8080), []openapiclient.NetworkLoadBalancerForwardingRuleTarget{*openapiclient.NewNetworkLoadBalancerForwardingRuleTarget("22.231.2.2", int32(8080), int32(123))}) // NetworkLoadBalancerForwardingRuleProperties | Properties of a forwarding rule to be updated
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPatch(context.Background(), datacenterId, networkLoadBalancerId, forwardingRuleId).NetworkLoadBalancerForwardingRuleProperties(networkLoadBalancerForwardingRuleProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesPatch`: NetworkLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPatch`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |
| **forwardingRuleId** | **string** | The unique ID of the forwarding rule |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesPatchRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **networkLoadBalancerForwardingRuleProperties** | [**NetworkLoadBalancerForwardingRuleProperties**](../models/networkloadbalancerforwardingruleproperties.md) | Properties of a forwarding rule to be updated |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**NetworkLoadBalancerForwardingRule**](../models/networkloadbalancerforwardingrule.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersNetworkloadbalancersForwardingrulesPost

```go
var result NetworkLoadBalancerForwardingRule = DatacentersNetworkloadbalancersForwardingrulesPost(ctx, datacenterId, networkLoadBalancerId)
                      .NetworkLoadBalancerForwardingRule(networkLoadBalancerForwardingRule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Add a Network Load Balancer Forwarding Rule

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    networkLoadBalancerForwardingRule := *openapiclient.NewNetworkLoadBalancerForwardingRule(*openapiclient.NewNetworkLoadBalancerForwardingRuleProperties("My Network Load Balancer forwarding rule", "ROUND_ROBIN", "TCP", "81.173.1.2", int32(8080), []openapiclient.NetworkLoadBalancerForwardingRuleTarget{*openapiclient.NewNetworkLoadBalancerForwardingRuleTarget("22.231.2.2", int32(8080), int32(123))})) // NetworkLoadBalancerForwardingRule | forwarding rule to add
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPost(context.Background(), datacenterId, networkLoadBalancerId).NetworkLoadBalancerForwardingRule(networkLoadBalancerForwardingRule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesPost`: NetworkLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPost`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **networkLoadBalancerForwardingRule** | [**NetworkLoadBalancerForwardingRule**](../models/networkloadbalancerforwardingrule.md) | forwarding rule to add |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**NetworkLoadBalancerForwardingRule**](../models/networkloadbalancerforwardingrule.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## DatacentersNetworkloadbalancersForwardingrulesPut

```go
var result NetworkLoadBalancerForwardingRule = DatacentersNetworkloadbalancersForwardingrulesPut(ctx, datacenterId, networkLoadBalancerId, forwardingRuleId)
                      .NetworkLoadBalancerForwardingRule(networkLoadBalancerForwardingRule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a forwarding rule of the Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule
    networkLoadBalancerForwardingRule := *openapiclient.NewNetworkLoadBalancerForwardingRulePut(*openapiclient.NewNetworkLoadBalancerForwardingRuleProperties("My Network Load Balancer forwarding rule", "ROUND_ROBIN", "TCP", "81.173.1.2", int32(8080), []openapiclient.NetworkLoadBalancerForwardingRuleTarget{*openapiclient.NewNetworkLoadBalancerForwardingRuleTarget("22.231.2.2", int32(8080), int32(123))})) // NetworkLoadBalancerForwardingRulePut | Modified Network Load Balancer Forwarding Rule
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPut(context.Background(), datacenterId, networkLoadBalancerId, forwardingRuleId).NetworkLoadBalancerForwardingRule(networkLoadBalancerForwardingRule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersForwardingrulesPut`: NetworkLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersForwardingrulesPut`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |
| **forwardingRuleId** | **string** | The unique ID of the forwarding rule |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersForwardingrulesPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **networkLoadBalancerForwardingRule** | [**NetworkLoadBalancerForwardingRulePut**](../models/networkloadbalancerforwardingruleput.md) | Modified Network Load Balancer Forwarding Rule |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**NetworkLoadBalancerForwardingRule**](../models/networkloadbalancerforwardingrule.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

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
    datacenterId := "datacenterId_example" // string | The unique ID of the datacenter
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with limit for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with offset for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersGet`: NetworkLoadBalancers
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |
| **offset** | **int32** | the first element \(of the total list of elements\) to include in the response \(use together with limit for pagination\) | \[default to 0\] |
| **limit** | **int32** | the maximum number of elements to return \(use together with offset for pagination\) | \[default to 1000\] |

### Return type

[**NetworkLoadBalancers**](../models/networkloadbalancers.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersNetworkloadbalancersPatch

```go
var result NetworkLoadBalancer = DatacentersNetworkloadbalancersPatch(ctx, datacenterId, networkLoadBalancerId)
                      .NetworkLoadBalancerProperties(networkLoadBalancerProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially update an Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    networkLoadBalancerProperties := *openapiclient.NewNetworkLoadBalancerProperties("My Network Load Balancer", int32(1), int32(2)) // NetworkLoadBalancerProperties | Network Load Balancer properties to be updated
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPatch(context.Background(), datacenterId, networkLoadBalancerId).NetworkLoadBalancerProperties(networkLoadBalancerProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersPatch`: NetworkLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPatch`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersPatchRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **networkLoadBalancerProperties** | [**NetworkLoadBalancerProperties**](../models/networkloadbalancerproperties.md) | Network Load Balancer properties to be updated |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**NetworkLoadBalancer**](../models/networkloadbalancer.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersNetworkloadbalancersPost

```go
var result NetworkLoadBalancer = DatacentersNetworkloadbalancersPost(ctx, datacenterId)
                      .NetworkLoadBalancer(networkLoadBalancer)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create an Network Load Balancer

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
    networkLoadBalancer := *openapiclient.NewNetworkLoadBalancer(*openapiclient.NewNetworkLoadBalancerProperties("My Network Load Balancer", int32(1), int32(2))) // NetworkLoadBalancer | Network Load Balancer to be created
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPost(context.Background(), datacenterId).NetworkLoadBalancer(networkLoadBalancer).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersPost`: NetworkLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPost`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **networkLoadBalancer** | [**NetworkLoadBalancer**](../models/networkloadbalancer.md) | Network Load Balancer to be created |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**NetworkLoadBalancer**](../models/networkloadbalancer.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## DatacentersNetworkloadbalancersPut

```go
var result NetworkLoadBalancer = DatacentersNetworkloadbalancersPut(ctx, datacenterId, networkLoadBalancerId)
                      .NetworkLoadBalancer(networkLoadBalancer)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Update an Network Load Balancer

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
    networkLoadBalancerId := "networkLoadBalancerId_example" // string | The unique ID of the Network Load Balancer
    networkLoadBalancer := *openapiclient.NewNetworkLoadBalancerPut(*openapiclient.NewNetworkLoadBalancerProperties("My Network Load Balancer", int32(1), int32(2))) // NetworkLoadBalancerPut | Modified Network Load Balancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with limit for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with offset for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPut(context.Background(), datacenterId, networkLoadBalancerId).NetworkLoadBalancer(networkLoadBalancer).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNetworkloadbalancersPut`: NetworkLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `NetworkLoadBalancersApi.DatacentersNetworkloadbalancersPut`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **networkLoadBalancerId** | **string** | The unique ID of the Network Load Balancer |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNetworkloadbalancersPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **networkLoadBalancer** | [**NetworkLoadBalancerPut**](../models/networkloadbalancerput.md) | Modified Network Load Balancer |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |
| **offset** | **int32** | the first element \(of the total list of elements\) to include in the response \(use together with limit for pagination\) | \[default to 0\] |
| **limit** | **int32** | the maximum number of elements to return \(use together with offset for pagination\) | \[default to 1000\] |

### Return type

[**NetworkLoadBalancer**](../models/networkloadbalancer.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

