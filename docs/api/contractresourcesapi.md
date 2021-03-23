# ContractResourcesApi

All URIs are relative to [https://api.ionos.com/cloudapi/v6](https://api.ionos.com/cloudapi/v6)

| Method | HTTP request | Description |
| :--- | :--- | :--- |
| [**ContractsGet**](contractresourcesapi.md#ContractsGet) | **Get** /contracts | Retrieve a Contract |

## ContractsGet

```go
var result Contracts = ContractsGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a Contract

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
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContractResourcesApi.ContractsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
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

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users with more than one contract must provide the contract number, against which all API requests are to be executed. |  |

### Return type

[**Contracts**](../models/contracts.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

