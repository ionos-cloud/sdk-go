# FirewallRulesApi

All URIs are relative to [https://api.ionos.com/cloudapi/v6](https://api.ionos.com/cloudapi/v6)

| Method | HTTP request | Description |
| :--- | :--- | :--- |
| [**DatacentersServersNicsFirewallrulesDelete**](firewallrulesapi.md#DatacentersServersNicsFirewallrulesDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Delete a Firewall Rule |
| [**DatacentersServersNicsFirewallrulesFindById**](firewallrulesapi.md#DatacentersServersNicsFirewallrulesFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Retrieve a Firewall Rule |
| [**DatacentersServersNicsFirewallrulesGet**](firewallrulesapi.md#DatacentersServersNicsFirewallrulesGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | List Firewall Rules |
| [**DatacentersServersNicsFirewallrulesPatch**](firewallrulesapi.md#DatacentersServersNicsFirewallrulesPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Partially Modify a Firewall Rule |
| [**DatacentersServersNicsFirewallrulesPost**](firewallrulesapi.md#DatacentersServersNicsFirewallrulesPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | Create a Firewall Rule |
| [**DatacentersServersNicsFirewallrulesPut**](firewallrulesapi.md#DatacentersServersNicsFirewallrulesPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Modify a Firewall Rule |

## DatacentersServersNicsFirewallrulesDelete

```go
var result map[string]interface{} = DatacentersServersNicsFirewallrulesDelete(ctx, datacenterId, serverId, nicId, firewallruleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete a Firewall Rule

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
    firewallruleId := "firewallruleId_example" // string | The unique ID of the Firewall Rule
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallRulesApi.DatacentersServersNicsFirewallrulesDelete(context.Background(), datacenterId, serverId, nicId, firewallruleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesDelete`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **serverId** | **string** | The unique ID of the Server |  |
| **nicId** | **string** | The unique ID of the NIC |  |
| **firewallruleId** | **string** | The unique ID of the Firewall Rule |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesDeleteRequest struct via the builder pattern

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

## DatacentersServersNicsFirewallrulesFindById

```go
var result FirewallRule = DatacentersServersNicsFirewallrulesFindById(ctx, datacenterId, serverId, nicId, firewallruleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a Firewall Rule

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
    firewallruleId := "firewallruleId_example" // string | The unique ID of the Firewall Rule
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallRulesApi.DatacentersServersNicsFirewallrulesFindById(context.Background(), datacenterId, serverId, nicId, firewallruleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesFindById`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesFindById`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **serverId** | **string** | The unique ID of the Server |  |
| **nicId** | **string** | The unique ID of the NIC |  |
| **firewallruleId** | **string** | The unique ID of the Firewall Rule |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesFindByIdRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**FirewallRule**](../models/firewallrule.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

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

List Firewall Rules

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
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with limit for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with offset for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallRulesApi.DatacentersServersNicsFirewallrulesGet(context.Background(), datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesGet`: FirewallRules
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **serverId** | **string** | The unique ID of the Server |  |
| **nicId** | **string** | The unique ID of the NIC |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |
| **offset** | **int32** | the first element \(of the total list of elements\) to include in the response \(use together with limit for pagination\) | \[default to 0\] |
| **limit** | **int32** | the maximum number of elements to return \(use together with offset for pagination\) | \[default to 1000\] |

### Return type

[**FirewallRules**](../models/firewallrules.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersServersNicsFirewallrulesPatch

```go
var result FirewallRule = DatacentersServersNicsFirewallrulesPatch(ctx, datacenterId, serverId, nicId, firewallruleId)
                      .Firewallrule(firewallrule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially Modify a Firewall Rule

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
    firewallruleId := "firewallruleId_example" // string | The unique ID of the Firewall Rule
    firewallrule := *openapiclient.NewFirewallruleProperties("TCP") // FirewallruleProperties | Modified Firewall Rule
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallRulesApi.DatacentersServersNicsFirewallrulesPatch(context.Background(), datacenterId, serverId, nicId, firewallruleId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesPatch`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesPatch`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **serverId** | **string** | The unique ID of the Server |  |
| **nicId** | **string** | The unique ID of the NIC |  |
| **firewallruleId** | **string** | The unique ID of the Firewall Rule |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesPatchRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **firewallrule** | [**FirewallruleProperties**](../models/firewallruleproperties.md) | Modified Firewall Rule |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**FirewallRule**](../models/firewallrule.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## DatacentersServersNicsFirewallrulesPost

```go
var result FirewallRule = DatacentersServersNicsFirewallrulesPost(ctx, datacenterId, serverId, nicId)
                      .Firewallrule(firewallrule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Firewall Rule

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
    firewallrule := *openapiclient.NewFirewallRule(*openapiclient.NewFirewallruleProperties("TCP")) // FirewallRule | Firewall Rule to be created
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallRulesApi.DatacentersServersNicsFirewallrulesPost(context.Background(), datacenterId, serverId, nicId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesPost`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesPost`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **serverId** | **string** | The unique ID of the server |  |
| **nicId** | **string** | The unique ID of the NIC |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **firewallrule** | [**FirewallRule**](../models/firewallrule.md) | Firewall Rule to be created |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**FirewallRule**](../models/firewallrule.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## DatacentersServersNicsFirewallrulesPut

```go
var result FirewallRule = DatacentersServersNicsFirewallrulesPut(ctx, datacenterId, serverId, nicId, firewallruleId)
                      .Firewallrule(firewallrule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Firewall Rule

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
    firewallruleId := "firewallruleId_example" // string | The unique ID of the Firewall Rule
    firewallrule := *openapiclient.NewFirewallRule(*openapiclient.NewFirewallruleProperties("TCP")) // FirewallRule | Modified Firewall Rule
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallRulesApi.DatacentersServersNicsFirewallrulesPut(context.Background(), datacenterId, serverId, nicId, firewallruleId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRulesApi.DatacentersServersNicsFirewallrulesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesPut`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRulesApi.DatacentersServersNicsFirewallrulesPut`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **serverId** | **string** | The unique ID of the Server |  |
| **nicId** | **string** | The unique ID of the NIC |  |
| **firewallruleId** | **string** | The unique ID of the Firewall Rule |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **firewallrule** | [**FirewallRule**](../models/firewallrule.md) | Modified Firewall Rule |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**FirewallRule**](../models/firewallrule.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

