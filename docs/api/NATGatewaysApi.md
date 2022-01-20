# \NATGatewaysApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersNatgatewaysDelete**](NATGatewaysApi.md#DatacentersNatgatewaysDelete) | **Delete** /datacenters/{datacenterId}/natgateways/{natGatewayId} | Delete NAT Gateways|
|[**DatacentersNatgatewaysFindByNatGatewayId**](NATGatewaysApi.md#DatacentersNatgatewaysFindByNatGatewayId) | **Get** /datacenters/{datacenterId}/natgateways/{natGatewayId} | Retrieve NAT Gateways|
|[**DatacentersNatgatewaysFlowlogsDelete**](NATGatewaysApi.md#DatacentersNatgatewaysFlowlogsDelete) | **Delete** /datacenters/{datacenterId}/natgateways/{natGatewayId}/flowlogs/{flowLogId} | Delete NAT Gateway Flow Logs|
|[**DatacentersNatgatewaysFlowlogsFindByFlowLogId**](NATGatewaysApi.md#DatacentersNatgatewaysFlowlogsFindByFlowLogId) | **Get** /datacenters/{datacenterId}/natgateways/{natGatewayId}/flowlogs/{flowLogId} | Retrieve NAT Gateway Flow Logs|
|[**DatacentersNatgatewaysFlowlogsGet**](NATGatewaysApi.md#DatacentersNatgatewaysFlowlogsGet) | **Get** /datacenters/{datacenterId}/natgateways/{natGatewayId}/flowlogs | List NAT Gateway Flow Logs|
|[**DatacentersNatgatewaysFlowlogsPatch**](NATGatewaysApi.md#DatacentersNatgatewaysFlowlogsPatch) | **Patch** /datacenters/{datacenterId}/natgateways/{natGatewayId}/flowlogs/{flowLogId} | Partially modify NAT Gateway Flow Logs|
|[**DatacentersNatgatewaysFlowlogsPost**](NATGatewaysApi.md#DatacentersNatgatewaysFlowlogsPost) | **Post** /datacenters/{datacenterId}/natgateways/{natGatewayId}/flowlogs | Create NAT Gateway Flow Logs|
|[**DatacentersNatgatewaysFlowlogsPut**](NATGatewaysApi.md#DatacentersNatgatewaysFlowlogsPut) | **Put** /datacenters/{datacenterId}/natgateways/{natGatewayId}/flowlogs/{flowLogId} | Modify NAT Gateway Flow Logs|
|[**DatacentersNatgatewaysGet**](NATGatewaysApi.md#DatacentersNatgatewaysGet) | **Get** /datacenters/{datacenterId}/natgateways | List NAT Gateways|
|[**DatacentersNatgatewaysPatch**](NATGatewaysApi.md#DatacentersNatgatewaysPatch) | **Patch** /datacenters/{datacenterId}/natgateways/{natGatewayId} | Partially modify NAT Gateways|
|[**DatacentersNatgatewaysPost**](NATGatewaysApi.md#DatacentersNatgatewaysPost) | **Post** /datacenters/{datacenterId}/natgateways | Create NAT Gateways|
|[**DatacentersNatgatewaysPut**](NATGatewaysApi.md#DatacentersNatgatewaysPut) | **Put** /datacenters/{datacenterId}/natgateways/{natGatewayId} | Modify NAT Gateways|
|[**DatacentersNatgatewaysRulesDelete**](NATGatewaysApi.md#DatacentersNatgatewaysRulesDelete) | **Delete** /datacenters/{datacenterId}/natgateways/{natGatewayId}/rules/{natGatewayRuleId} | Delete NAT Gateway rules|
|[**DatacentersNatgatewaysRulesFindByNatGatewayRuleId**](NATGatewaysApi.md#DatacentersNatgatewaysRulesFindByNatGatewayRuleId) | **Get** /datacenters/{datacenterId}/natgateways/{natGatewayId}/rules/{natGatewayRuleId} | Retrieve NAT Gateway rules|
|[**DatacentersNatgatewaysRulesGet**](NATGatewaysApi.md#DatacentersNatgatewaysRulesGet) | **Get** /datacenters/{datacenterId}/natgateways/{natGatewayId}/rules | List NAT Gateway rules|
|[**DatacentersNatgatewaysRulesPatch**](NATGatewaysApi.md#DatacentersNatgatewaysRulesPatch) | **Patch** /datacenters/{datacenterId}/natgateways/{natGatewayId}/rules/{natGatewayRuleId} | Partially modify NAT Gateway rules|
|[**DatacentersNatgatewaysRulesPost**](NATGatewaysApi.md#DatacentersNatgatewaysRulesPost) | **Post** /datacenters/{datacenterId}/natgateways/{natGatewayId}/rules | Create NAT Gateway rules|
|[**DatacentersNatgatewaysRulesPut**](NATGatewaysApi.md#DatacentersNatgatewaysRulesPut) | **Put** /datacenters/{datacenterId}/natgateways/{natGatewayId}/rules/{natGatewayRuleId} | Modify NAT Gateway rules|



## DatacentersNatgatewaysDelete

```go
var result  = DatacentersNatgatewaysDelete(ctx, datacenterId, natGatewayId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete NAT Gateways



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysDelete(context.Background(), datacenterId, natGatewayId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysDeleteRequest struct via the builder pattern


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



## DatacentersNatgatewaysFindByNatGatewayId

```go
var result NatGateway = DatacentersNatgatewaysFindByNatGatewayId(ctx, datacenterId, natGatewayId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve NAT Gateways



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysFindByNatGatewayId(context.Background(), datacenterId, natGatewayId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysFindByNatGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysFindByNatGatewayId`: NatGateway
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysFindByNatGatewayId`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysFindByNatGatewayIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGateway**](../models/NatGateway.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysFlowlogsDelete

```go
var result  = DatacentersNatgatewaysFlowlogsDelete(ctx, datacenterId, natGatewayId, flowLogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Delete NAT Gateway Flow Logs



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysFlowlogsDelete(context.Background(), datacenterId, natGatewayId, flowLogId).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysFlowlogsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |
|**flowLogId** | **string** | The unique ID of the Flow Log. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysFlowlogsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysFlowlogsFindByFlowLogId

```go
var result FlowLog = DatacentersNatgatewaysFlowlogsFindByFlowLogId(ctx, datacenterId, natGatewayId, flowLogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Retrieve NAT Gateway Flow Logs



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysFlowlogsFindByFlowLogId(context.Background(), datacenterId, natGatewayId, flowLogId).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysFlowlogsFindByFlowLogId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysFlowlogsFindByFlowLogId`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysFlowlogsFindByFlowLogId`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |
|**flowLogId** | **string** | The unique ID of the Flow Log. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysFlowlogsFindByFlowLogIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysFlowlogsGet

```go
var result FlowLogs = DatacentersNatgatewaysFlowlogsGet(ctx, datacenterId, natGatewayId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List NAT Gateway Flow Logs



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysFlowlogsGet(context.Background(), datacenterId, natGatewayId).Pretty(pretty).Depth(depth).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysFlowlogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysFlowlogsGet`: FlowLogs
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysFlowlogsGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysFlowlogsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 1000]|

### Return type

[**FlowLogs**](../models/FlowLogs.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysFlowlogsPatch

```go
var result FlowLog = DatacentersNatgatewaysFlowlogsPatch(ctx, datacenterId, natGatewayId, flowLogId)
                      .NatGatewayFlowLogProperties(natGatewayFlowLogProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Partially modify NAT Gateway Flow Logs



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log.
    natGatewayFlowLogProperties := *openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key") // FlowLogProperties | The properties of the Flow Log to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysFlowlogsPatch(context.Background(), datacenterId, natGatewayId, flowLogId).NatGatewayFlowLogProperties(natGatewayFlowLogProperties).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysFlowlogsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysFlowlogsPatch`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysFlowlogsPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |
|**flowLogId** | **string** | The unique ID of the Flow Log. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysFlowlogsPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **natGatewayFlowLogProperties** | [**FlowLogProperties**](../models/FlowLogProperties.md) | The properties of the Flow Log to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysFlowlogsPost

```go
var result FlowLog = DatacentersNatgatewaysFlowlogsPost(ctx, datacenterId, natGatewayId)
                      .NatGatewayFlowLog(natGatewayFlowLog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Create NAT Gateway Flow Logs



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    natGatewayFlowLog := *openapiclient.NewFlowLog(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLog | The Flow Log to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysFlowlogsPost(context.Background(), datacenterId, natGatewayId).NatGatewayFlowLog(natGatewayFlowLog).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysFlowlogsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysFlowlogsPost`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysFlowlogsPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysFlowlogsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **natGatewayFlowLog** | [**FlowLog**](../models/FlowLog.md) | The Flow Log to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNatgatewaysFlowlogsPut

```go
var result FlowLog = DatacentersNatgatewaysFlowlogsPut(ctx, datacenterId, natGatewayId, flowLogId)
                      .NatGatewayFlowLog(natGatewayFlowLog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Modify NAT Gateway Flow Logs



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    flowLogId := "flowLogId_example" // string | The unique ID of the Flow Log.
    natGatewayFlowLog := *openapiclient.NewFlowLogPut(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLogPut | The modified NAT Gateway Flow Log.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysFlowlogsPut(context.Background(), datacenterId, natGatewayId, flowLogId).NatGatewayFlowLog(natGatewayFlowLog).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysFlowlogsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysFlowlogsPut`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysFlowlogsPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |
|**flowLogId** | **string** | The unique ID of the Flow Log. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysFlowlogsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **natGatewayFlowLog** | [**FlowLogPut**](../models/FlowLogPut.md) | The modified NAT Gateway Flow Log. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNatgatewaysGet

```go
var result NatGateways = DatacentersNatgatewaysGet(ctx, datacenterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List NAT Gateways



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysGet`: NatGateways
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGateways**](../models/NatGateways.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysPatch

```go
var result NatGateway = DatacentersNatgatewaysPatch(ctx, datacenterId, natGatewayId)
                      .NatGatewayProperties(natGatewayProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify NAT Gateways



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    natGatewayProperties := *openapiclient.NewNatGatewayProperties("My NAT Gateway", []string{"PublicIps_example"}) // NatGatewayProperties | The properties of the NAT Gateway to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysPatch(context.Background(), datacenterId, natGatewayId).NatGatewayProperties(natGatewayProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysPatch`: NatGateway
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **natGatewayProperties** | [**NatGatewayProperties**](../models/NatGatewayProperties.md) | The properties of the NAT Gateway to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGateway**](../models/NatGateway.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysPost

```go
var result NatGateway = DatacentersNatgatewaysPost(ctx, datacenterId)
                      .NatGateway(natGateway)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create NAT Gateways



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
    natGateway := *openapiclient.NewNatGateway(*openapiclient.NewNatGatewayProperties("My NAT Gateway", []string{"PublicIps_example"})) // NatGateway | The NAT Gateway to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysPost(context.Background(), datacenterId).NatGateway(natGateway).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysPost`: NatGateway
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **natGateway** | [**NatGateway**](../models/NatGateway.md) | The NAT Gateway to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGateway**](../models/NatGateway.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNatgatewaysPut

```go
var result NatGateway = DatacentersNatgatewaysPut(ctx, datacenterId, natGatewayId)
                      .NatGateway(natGateway)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify NAT Gateways



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    natGateway := *openapiclient.NewNatGatewayPut(*openapiclient.NewNatGatewayProperties("My NAT Gateway", []string{"PublicIps_example"})) // NatGatewayPut | The modified NAT Gateway.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysPut(context.Background(), datacenterId, natGatewayId).NatGateway(natGateway).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysPut`: NatGateway
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **natGateway** | [**NatGatewayPut**](../models/NatGatewayPut.md) | The modified NAT Gateway. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGateway**](../models/NatGateway.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNatgatewaysRulesDelete

```go
var result  = DatacentersNatgatewaysRulesDelete(ctx, datacenterId, natGatewayId, natGatewayRuleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete NAT Gateway rules



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    natGatewayRuleId := "natGatewayRuleId_example" // string | The unique ID of the NAT Gateway rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysRulesDelete(context.Background(), datacenterId, natGatewayId, natGatewayRuleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysRulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |
|**natGatewayRuleId** | **string** | The unique ID of the NAT Gateway rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysRulesDeleteRequest struct via the builder pattern


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



## DatacentersNatgatewaysRulesFindByNatGatewayRuleId

```go
var result NatGatewayRule = DatacentersNatgatewaysRulesFindByNatGatewayRuleId(ctx, datacenterId, natGatewayId, natGatewayRuleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve NAT Gateway rules



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    natGatewayRuleId := "natGatewayRuleId_example" // string | The unique ID of the NAT Gateway rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysRulesFindByNatGatewayRuleId(context.Background(), datacenterId, natGatewayId, natGatewayRuleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysRulesFindByNatGatewayRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysRulesFindByNatGatewayRuleId`: NatGatewayRule
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysRulesFindByNatGatewayRuleId`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |
|**natGatewayRuleId** | **string** | The unique ID of the NAT Gateway rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysRulesFindByNatGatewayRuleIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGatewayRule**](../models/NatGatewayRule.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysRulesGet

```go
var result NatGatewayRules = DatacentersNatgatewaysRulesGet(ctx, datacenterId, natGatewayId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List NAT Gateway rules



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysRulesGet(context.Background(), datacenterId, natGatewayId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysRulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysRulesGet`: NatGatewayRules
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysRulesGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysRulesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGatewayRules**](../models/NatGatewayRules.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysRulesPatch

```go
var result NatGatewayRule = DatacentersNatgatewaysRulesPatch(ctx, datacenterId, natGatewayId, natGatewayRuleId)
                      .NatGatewayRuleProperties(natGatewayRuleProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify NAT Gateway rules



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    natGatewayRuleId := "natGatewayRuleId_example" // string | The unique ID of the NAT Gateway rule.
    natGatewayRuleProperties := *openapiclient.NewNatGatewayRuleProperties("My NAT Gateway rule", "10.0.1.0/24", "192.18.7.17") // NatGatewayRuleProperties | The properties of the NAT Gateway rule to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysRulesPatch(context.Background(), datacenterId, natGatewayId, natGatewayRuleId).NatGatewayRuleProperties(natGatewayRuleProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysRulesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysRulesPatch`: NatGatewayRule
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysRulesPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |
|**natGatewayRuleId** | **string** | The unique ID of the NAT Gateway rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysRulesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **natGatewayRuleProperties** | [**NatGatewayRuleProperties**](../models/NatGatewayRuleProperties.md) | The properties of the NAT Gateway rule to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGatewayRule**](../models/NatGatewayRule.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersNatgatewaysRulesPost

```go
var result NatGatewayRule = DatacentersNatgatewaysRulesPost(ctx, datacenterId, natGatewayId)
                      .NatGatewayRule(natGatewayRule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create NAT Gateway rules



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    natGatewayRule := *openapiclient.NewNatGatewayRule(*openapiclient.NewNatGatewayRuleProperties("My NAT Gateway rule", "10.0.1.0/24", "192.18.7.17")) // NatGatewayRule | The NAT Gateway rule to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysRulesPost(context.Background(), datacenterId, natGatewayId).NatGatewayRule(natGatewayRule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysRulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysRulesPost`: NatGatewayRule
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysRulesPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysRulesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **natGatewayRule** | [**NatGatewayRule**](../models/NatGatewayRule.md) | The NAT Gateway rule to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGatewayRule**](../models/NatGatewayRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersNatgatewaysRulesPut

```go
var result NatGatewayRule = DatacentersNatgatewaysRulesPut(ctx, datacenterId, natGatewayId, natGatewayRuleId)
                      .NatGatewayRule(natGatewayRule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify NAT Gateway rules



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
    natGatewayId := "natGatewayId_example" // string | The unique ID of the NAT Gateway.
    natGatewayRuleId := "natGatewayRuleId_example" // string | The unique ID of the NAT Gateway rule.
    natGatewayRule := *openapiclient.NewNatGatewayRulePut(*openapiclient.NewNatGatewayRuleProperties("My NAT Gateway rule", "10.0.1.0/24", "192.18.7.17")) // NatGatewayRulePut | The modified NAT Gateway rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NATGatewaysApi.DatacentersNatgatewaysRulesPut(context.Background(), datacenterId, natGatewayId, natGatewayRuleId).NatGatewayRule(natGatewayRule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NATGatewaysApi.DatacentersNatgatewaysRulesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersNatgatewaysRulesPut`: NatGatewayRule
    fmt.Fprintf(os.Stdout, "Response from `NATGatewaysApi.DatacentersNatgatewaysRulesPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**natGatewayId** | **string** | The unique ID of the NAT Gateway. | |
|**natGatewayRuleId** | **string** | The unique ID of the NAT Gateway rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersNatgatewaysRulesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **natGatewayRule** | [**NatGatewayRulePut**](../models/NatGatewayRulePut.md) | The modified NAT Gateway rule. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**NatGatewayRule**](../models/NatGatewayRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


