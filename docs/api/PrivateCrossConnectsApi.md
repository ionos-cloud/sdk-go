# \PrivateCrossConnectsApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**PccsDelete**](PrivateCrossConnectsApi.md#PccsDelete) | **Delete** /pccs/{pccId} | Delete Cross Connects|
|[**PccsFindById**](PrivateCrossConnectsApi.md#PccsFindById) | **Get** /pccs/{pccId} | Retrieve a Cross Connect|
|[**PccsGet**](PrivateCrossConnectsApi.md#PccsGet) | **Get** /pccs | List Cross Connects|
|[**PccsPatch**](PrivateCrossConnectsApi.md#PccsPatch) | **Patch** /pccs/{pccId} | Partially modify a Cross Connects|
|[**PccsPost**](PrivateCrossConnectsApi.md#PccsPost) | **Post** /pccs | Create a Cross Connect|



## PccsDelete

```go
var result  = PccsDelete(ctx, pccId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete Cross Connects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    pccId := "pccId_example" // string | The unique ID of the Cross Connect.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.PrivateCrossConnectsApi.PccsDelete(context.Background(), pccId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateCrossConnectsApi.PccsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pccId** | **string** | The unique ID of the Cross Connect. | |

### Other Parameters

Other parameters are passed through a pointer to a apiPccsDeleteRequest struct via the builder pattern


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



## PccsFindById

```go
var result PrivateCrossConnect = PccsFindById(ctx, pccId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a Cross Connect



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    pccId := "pccId_example" // string | The unique ID of the Cross Connect.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.PrivateCrossConnectsApi.PccsFindById(context.Background(), pccId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateCrossConnectsApi.PccsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PccsFindById`: PrivateCrossConnect
    fmt.Fprintf(os.Stdout, "Response from `PrivateCrossConnectsApi.PccsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pccId** | **string** | The unique ID of the Cross Connect. | |

### Other Parameters

Other parameters are passed through a pointer to a apiPccsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**PrivateCrossConnect**](../models/PrivateCrossConnect.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PccsGet

```go
var result PrivateCrossConnects = PccsGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Cross Connects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.PrivateCrossConnectsApi.PccsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateCrossConnectsApi.PccsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PccsGet`: PrivateCrossConnects
    fmt.Fprintf(os.Stdout, "Response from `PrivateCrossConnectsApi.PccsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPccsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**PrivateCrossConnects**](../models/PrivateCrossConnects.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PccsPatch

```go
var result PrivateCrossConnect = PccsPatch(ctx, pccId)
                      .Pcc(pcc)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify a Cross Connects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    pccId := "pccId_example" // string | The unique ID of the Cross Connect.
    pcc := *openapiclient.NewPrivateCrossConnectProperties() // PrivateCrossConnectProperties | The properties of the Cross Connect to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.PrivateCrossConnectsApi.PccsPatch(context.Background(), pccId).Pcc(pcc).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateCrossConnectsApi.PccsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PccsPatch`: PrivateCrossConnect
    fmt.Fprintf(os.Stdout, "Response from `PrivateCrossConnectsApi.PccsPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pccId** | **string** | The unique ID of the Cross Connect. | |

### Other Parameters

Other parameters are passed through a pointer to a apiPccsPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pcc** | [**PrivateCrossConnectProperties**](../models/PrivateCrossConnectProperties.md) | The properties of the Cross Connect to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**PrivateCrossConnect**](../models/PrivateCrossConnect.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## PccsPost

```go
var result PrivateCrossConnect = PccsPost(ctx)
                      .Pcc(pcc)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Cross Connect



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    pcc := *openapiclient.NewPrivateCrossConnect(*openapiclient.NewPrivateCrossConnectProperties()) // PrivateCrossConnect | The Cross Connect to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.PrivateCrossConnectsApi.PccsPost(context.Background()).Pcc(pcc).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateCrossConnectsApi.PccsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PccsPost`: PrivateCrossConnect
    fmt.Fprintf(os.Stdout, "Response from `PrivateCrossConnectsApi.PccsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPccsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pcc** | [**PrivateCrossConnect**](../models/PrivateCrossConnect.md) | The Cross Connect to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**PrivateCrossConnect**](../models/PrivateCrossConnect.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


