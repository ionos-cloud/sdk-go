# \ContractResourcesApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ContractsGet**](ContractResourcesApi.md#ContractsGet) | **Get** /contracts | Retrieve contracts|



## ContractsGet

```go
var result Contracts = ContractsGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve contracts



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractResourcesApi.ContractsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractResourcesApi.ContractsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsGet`: Contracts
    fmt.Fprintf(os.Stdout, "Response from `ContractResourcesApi.ContractsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Contracts**](../models/Contracts.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


