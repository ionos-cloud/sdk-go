# \NicApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersServersNicsDelete**](NicApi.md#DatacentersServersNicsDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Delete a Nic
[**DatacentersServersNicsFindById**](NicApi.md#DatacentersServersNicsFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Retrieve a Nic
[**DatacentersServersNicsFirewallrulesDelete**](NicApi.md#DatacentersServersNicsFirewallrulesDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Delete a Firewall Rule
[**DatacentersServersNicsFirewallrulesFindById**](NicApi.md#DatacentersServersNicsFirewallrulesFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Retrieve a Firewall Rule
[**DatacentersServersNicsFirewallrulesGet**](NicApi.md#DatacentersServersNicsFirewallrulesGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | List Firewall Rules 
[**DatacentersServersNicsFirewallrulesPatch**](NicApi.md#DatacentersServersNicsFirewallrulesPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Partially modify a Firewall Rule
[**DatacentersServersNicsFirewallrulesPost**](NicApi.md#DatacentersServersNicsFirewallrulesPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | Create a Firewall Rule
[**DatacentersServersNicsFirewallrulesPut**](NicApi.md#DatacentersServersNicsFirewallrulesPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Modify a Firewall Rule
[**DatacentersServersNicsGet**](NicApi.md#DatacentersServersNicsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics | List Nics 
[**DatacentersServersNicsPatch**](NicApi.md#DatacentersServersNicsPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Partially modify a Nic
[**DatacentersServersNicsPost**](NicApi.md#DatacentersServersNicsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics | Create a Nic
[**DatacentersServersNicsPut**](NicApi.md#DatacentersServersNicsPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Modify a Nic



## DatacentersServersNicsDelete

> map[string]interface{} DatacentersServersNicsDelete(ctx, datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Delete a Nic



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.DatacentersServersNicsDelete(context.Background(), datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**nicId** | **string** | The unique ID of the NIC | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFindById

> Nic DatacentersServersNicsFindById(ctx, datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Retrieve a Nic



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.DatacentersServersNicsFindById(context.Background(), datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFindById`: Nic
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsFindById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**nicId** | **string** | The unique ID of the NIC | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFindByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nic**](Nic.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesDelete

> map[string]interface{} DatacentersServersNicsFirewallrulesDelete(ctx, datacenterId, serverId, nicId, firewallruleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

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
    resp, r, err := api_client.NicApi.DatacentersServersNicsFirewallrulesDelete(context.Background(), datacenterId, serverId, nicId, firewallruleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsFirewallrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsFirewallrulesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**nicId** | **string** | The unique ID of the NIC | 
**firewallruleId** | **string** | The unique ID of the Firewall Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesFindById

> FirewallRule DatacentersServersNicsFirewallrulesFindById(ctx, datacenterId, serverId, nicId, firewallruleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

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
    resp, r, err := api_client.NicApi.DatacentersServersNicsFirewallrulesFindById(context.Background(), datacenterId, serverId, nicId, firewallruleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsFirewallrulesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesFindById`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsFirewallrulesFindById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**nicId** | **string** | The unique ID of the NIC | 
**firewallruleId** | **string** | The unique ID of the Firewall Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesFindByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesGet

> FirewallRules DatacentersServersNicsFirewallrulesGet(ctx, datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()

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
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with <code>limit</code> for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with <code>offset</code> for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.DatacentersServersNicsFirewallrulesGet(context.Background(), datacenterId, serverId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsFirewallrulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesGet`: FirewallRules
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsFirewallrulesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**nicId** | **string** | The unique ID of the NIC | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 
 **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with &lt;code&gt;limit&lt;/code&gt; for pagination) | [default to 0]
 **limit** | **int32** | the maximum number of elements to return (use together with &lt;code&gt;offset&lt;/code&gt; for pagination) | [default to 1000]

### Return type

[**FirewallRules**](FirewallRules.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesPatch

> FirewallRule DatacentersServersNicsFirewallrulesPatch(ctx, datacenterId, serverId, nicId, firewallruleId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Partially modify a Firewall Rule



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
    resp, r, err := api_client.NicApi.DatacentersServersNicsFirewallrulesPatch(context.Background(), datacenterId, serverId, nicId, firewallruleId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsFirewallrulesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesPatch`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsFirewallrulesPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**nicId** | **string** | The unique ID of the NIC | 
**firewallruleId** | **string** | The unique ID of the Firewall Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **firewallrule** | [**FirewallruleProperties**](FirewallruleProperties.md) | Modified Firewall Rule | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesPost

> FirewallRule DatacentersServersNicsFirewallrulesPost(ctx, datacenterId, serverId, nicId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

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
    resp, r, err := api_client.NicApi.DatacentersServersNicsFirewallrulesPost(context.Background(), datacenterId, serverId, nicId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsFirewallrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesPost`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsFirewallrulesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the server | 
**nicId** | **string** | The unique ID of the NIC | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **firewallrule** | [**FirewallRule**](FirewallRule.md) | Firewall Rule to be created | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesPut

> FirewallRule DatacentersServersNicsFirewallrulesPut(ctx, datacenterId, serverId, nicId, firewallruleId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

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
    resp, r, err := api_client.NicApi.DatacentersServersNicsFirewallrulesPut(context.Background(), datacenterId, serverId, nicId, firewallruleId).Firewallrule(firewallrule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsFirewallrulesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsFirewallrulesPut`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsFirewallrulesPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**nicId** | **string** | The unique ID of the NIC | 
**firewallruleId** | **string** | The unique ID of the Firewall Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsFirewallrulesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **firewallrule** | [**FirewallRule**](FirewallRule.md) | Modified Firewall Rule | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsGet

> Nics DatacentersServersNicsGet(ctx, datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()

List Nics 



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
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with <code>limit</code> for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with <code>offset</code> for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.DatacentersServersNicsGet(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsGet`: Nics
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 
 **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with &lt;code&gt;limit&lt;/code&gt; for pagination) | [default to 0]
 **limit** | **int32** | the maximum number of elements to return (use together with &lt;code&gt;offset&lt;/code&gt; for pagination) | [default to 1000]

### Return type

[**Nics**](Nics.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsPatch

> Nic DatacentersServersNicsPatch(ctx, datacenterId, serverId, nicId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Partially modify a Nic



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
    nic := *openapiclient.NewNicProperties(int32(2)) // NicProperties | Modified properties of Nic
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.DatacentersServersNicsPatch(context.Background(), datacenterId, serverId, nicId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsPatch`: Nic
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**nicId** | **string** | The unique ID of the NIC | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nic** | [**NicProperties**](NicProperties.md) | Modified properties of Nic | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nic**](Nic.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsPost

> Nic DatacentersServersNicsPost(ctx, datacenterId, serverId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Create a Nic



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
    nic := *openapiclient.NewNic(*openapiclient.NewNicProperties(int32(2))) // Nic | Nic to be created
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.DatacentersServersNicsPost(context.Background(), datacenterId, serverId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsPost`: Nic
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nic** | [**Nic**](Nic.md) | Nic to be created | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nic**](Nic.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsPut

> Nic DatacentersServersNicsPut(ctx, datacenterId, serverId, nicId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Modify a Nic



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
    nic := *openapiclient.NewNic(*openapiclient.NewNicProperties(int32(2))) // Nic | Modified Nic
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NicApi.DatacentersServersNicsPut(context.Background(), datacenterId, serverId, nicId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NicApi.DatacentersServersNicsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsPut`: Nic
    fmt.Fprintf(os.Stdout, "Response from `NicApi.DatacentersServersNicsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 
**serverId** | **string** | The unique ID of the Server | 
**nicId** | **string** | The unique ID of the NIC | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nic** | [**Nic**](Nic.md) | Modified Nic | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nic**](Nic.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

