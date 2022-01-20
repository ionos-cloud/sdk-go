# \FirewallRulesApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersServersNicsFirewallrulesDelete**](FirewallRulesApi.md#DatacentersServersNicsFirewallrulesDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Delete firewall rules|
|[**DatacentersServersNicsFirewallrulesFindById**](FirewallRulesApi.md#DatacentersServersNicsFirewallrulesFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Retrieve firewall rules|
|[**DatacentersServersNicsFirewallrulesGet**](FirewallRulesApi.md#DatacentersServersNicsFirewallrulesGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | List firewall rules|
|[**DatacentersServersNicsFirewallrulesPatch**](FirewallRulesApi.md#DatacentersServersNicsFirewallrulesPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Partially modify firewall rules|
|[**DatacentersServersNicsFirewallrulesPost**](FirewallRulesApi.md#DatacentersServersNicsFirewallrulesPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | Create firewall rules|
|[**DatacentersServersNicsFirewallrulesPut**](FirewallRulesApi.md#DatacentersServersNicsFirewallrulesPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Modify firewall rules|



## DatacentersServersNicsFirewallrulesDelete

```go
var result  = DatacentersServersNicsFirewallrulesDelete(ctx, datacenterId, serverId, nicId, firewallruleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete firewall rules



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
    firewallruleId := "firewallruleId_example" // string | The unique ID of the firewall rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRulesApi.DatacentersServersNicsFirewallrulesDelete(context.Background(), datacenterId, serverId, nicId, firewallruleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesDelete``: %v\n", err)
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
|**firewallruleId** | **string** | The unique ID of the firewall rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesDeleteRequest struct via the builder pattern


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



## DatacentersServersNicsFirewallrulesFindById

```go
var result FirewallRule = DatacentersServersNicsFirewallrulesFindById(ctx, datacenterId, serverId, nicId, firewallruleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve firewall rules



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
    firewallruleId := "firewallruleId_example" // string | The unique ID of the firewall rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRulesApi.DatacentersServersNicsFirewallrulesFindById(context.Background(), datacenterId, serverId, nicId, firewallruleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesFindById`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**nicId** | **string** | The unique ID of the NIC. | |
|**firewallruleId** | **string** | The unique ID of the firewall rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FirewallRule**](../models/FirewallRule.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersNicsFirewallrulesGet

```go
var result FirewallRules = DatacentersServersNicsFirewallrulesGet(ctx, datacenterId, serverId, nicId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List firewall rules



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
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRulesApi.DatacentersServersNicsFirewallrulesGet(context.Background(), datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesGet`: FirewallRules
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesGet`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 1000]|

### Return type

[**FirewallRules**](../models/FirewallRules.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersNicsFirewallrulesPatch

```go
var result FirewallRule = DatacentersServersNicsFirewallrulesPatch(ctx, datacenterId, serverId, nicId, firewallruleId)
                      .Firewallrule(firewallrule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify firewall rules



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
    firewallruleId := "firewallruleId_example" // string | The unique ID of the firewall rule.
    firewallrule := *openapiclient.NewFirewallruleProperties("TCP") // FirewallruleProperties | The properties of the firewall rule to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRulesApi.DatacentersServersNicsFirewallrulesPatch(context.Background(), datacenterId, serverId, nicId, firewallruleId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesPatch`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**nicId** | **string** | The unique ID of the NIC. | |
|**firewallruleId** | **string** | The unique ID of the firewall rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **firewallrule** | [**FirewallruleProperties**](../models/FirewallruleProperties.md) | The properties of the firewall rule to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FirewallRule**](../models/FirewallRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersNicsFirewallrulesPost

```go
var result FirewallRule = DatacentersServersNicsFirewallrulesPost(ctx, datacenterId, serverId, nicId)
                      .Firewallrule(firewallrule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create firewall rules



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
    firewallrule := *openapiclient.NewFirewallRule(*openapiclient.NewFirewallruleProperties("TCP")) // FirewallRule | The firewall rule to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRulesApi.DatacentersServersNicsFirewallrulesPost(context.Background(), datacenterId, serverId, nicId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesPost`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesPost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **firewallrule** | [**FirewallRule**](../models/FirewallRule.md) | The firewall rule to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FirewallRule**](../models/FirewallRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersNicsFirewallrulesPut

```go
var result FirewallRule = DatacentersServersNicsFirewallrulesPut(ctx, datacenterId, serverId, nicId, firewallruleId)
                      .Firewallrule(firewallrule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify firewall rules



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
    firewallruleId := "firewallruleId_example" // string | The unique ID of the firewall rule.
    firewallrule := *openapiclient.NewFirewallRule(*openapiclient.NewFirewallruleProperties("TCP")) // FirewallRule | The modified firewall rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRulesApi.DatacentersServersNicsFirewallrulesPut(context.Background(), datacenterId, serverId, nicId, firewallruleId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesPut`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**nicId** | **string** | The unique ID of the NIC. | |
|**firewallruleId** | **string** | The unique ID of the firewall rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **firewallrule** | [**FirewallRule**](../models/FirewallRule.md) | The modified firewall rule. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FirewallRule**](../models/FirewallRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


