# \BackupUnitApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**BackupunitsDelete**](BackupUnitApi.md#BackupunitsDelete) | **Delete** /backupunits/{backupunitId} | Delete a Backup Unit|
|[**BackupunitsFindById**](BackupUnitApi.md#BackupunitsFindById) | **Get** /backupunits/{backupunitId} | Returns the specified backup Unit|
|[**BackupunitsGet**](BackupUnitApi.md#BackupunitsGet) | **Get** /backupunits | List Backup Units |
|[**BackupunitsPatch**](BackupUnitApi.md#BackupunitsPatch) | **Patch** /backupunits/{backupunitId} | Partially modify a Backup Unit|
|[**BackupunitsPost**](BackupUnitApi.md#BackupunitsPost) | **Post** /backupunits | Create a Backup Unit|
|[**BackupunitsPut**](BackupUnitApi.md#BackupunitsPut) | **Put** /backupunits/{backupunitId} | Modify a Backup Unit|
|[**BackupunitsSsourlGet**](BackupUnitApi.md#BackupunitsSsourlGet) | **Get** /backupunits/{backupunitId}/ssourl | Returns a single signon URL for the specified backup Unit.|



## BackupunitsDelete

> map[string]interface{} BackupunitsDelete(ctx, backupunitId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Delete a Backup Unit



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
    backupunitId := "backupunitId_example" // string | The unique ID of the backup Unit
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupUnitApi.BackupunitsDelete(context.Background(), backupunitId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupUnitApi.BackupunitsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupunitsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupUnitApi.BackupunitsDelete`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**backupunitId** | **string** | The unique ID of the backup Unit | |

### Other Parameters

Other parameters are passed through a pointer to a apiBackupunitsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## BackupunitsFindById

> BackupUnit BackupunitsFindById(ctx, backupunitId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Returns the specified backup Unit



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
    backupunitId := "backupunitId_example" // string | The unique ID of the backup unit
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupUnitApi.BackupunitsFindById(context.Background(), backupunitId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupUnitApi.BackupunitsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupunitsFindById`: BackupUnit
    fmt.Fprintf(os.Stdout, "Response from `BackupUnitApi.BackupunitsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**backupunitId** | **string** | The unique ID of the backup unit | |

### Other Parameters

Other parameters are passed through a pointer to a apiBackupunitsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**BackupUnit**](BackupUnit.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## BackupunitsGet

> BackupUnits BackupunitsGet(ctx).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

List Backup Units 



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
    resp, r, err := api_client.BackupUnitApi.BackupunitsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupUnitApi.BackupunitsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupunitsGet`: BackupUnits
    fmt.Fprintf(os.Stdout, "Response from `BackupUnitApi.BackupunitsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupunitsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**BackupUnits**](BackupUnits.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## BackupunitsPatch

> BackupUnit BackupunitsPatch(ctx, backupunitId).BackupUnitProperties(backupUnitProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Partially modify a Backup Unit



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
    backupunitId := "backupunitId_example" // string | The unique ID of the backup unit
    backupUnitProperties := *openapiclient.NewBackupUnitProperties("BackupUnitName") // BackupUnitProperties | Modified backup Unit properties
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupUnitApi.BackupunitsPatch(context.Background(), backupunitId).BackupUnitProperties(backupUnitProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupUnitApi.BackupunitsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupunitsPatch`: BackupUnit
    fmt.Fprintf(os.Stdout, "Response from `BackupUnitApi.BackupunitsPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**backupunitId** | **string** | The unique ID of the backup unit | |

### Other Parameters

Other parameters are passed through a pointer to a apiBackupunitsPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

| **backupUnitProperties** | [**BackupUnitProperties**](BackupUnitProperties.md) | Modified backup Unit properties | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**BackupUnit**](BackupUnit.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## BackupunitsPost

> BackupUnit BackupunitsPost(ctx).BackupUnit(backupUnit).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Create a Backup Unit



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
    backupUnit := *openapiclient.NewBackupUnit(*openapiclient.NewBackupUnitProperties("BackupUnitName")) // BackupUnit | Payload containing data to create a new Backup Unit
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupUnitApi.BackupunitsPost(context.Background()).BackupUnit(backupUnit).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupUnitApi.BackupunitsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupunitsPost`: BackupUnit
    fmt.Fprintf(os.Stdout, "Response from `BackupUnitApi.BackupunitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackupunitsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **backupUnit** | [**BackupUnit**](BackupUnit.md) | Payload containing data to create a new Backup Unit | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**BackupUnit**](BackupUnit.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## BackupunitsPut

> BackupUnit BackupunitsPut(ctx, backupunitId).BackupUnit(backupUnit).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Modify a Backup Unit



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
    backupunitId := "backupunitId_example" // string | The unique ID of the backup unit
    backupUnit := *openapiclient.NewBackupUnit(*openapiclient.NewBackupUnitProperties("BackupUnitName")) // BackupUnit | Modified backup Unit
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupUnitApi.BackupunitsPut(context.Background(), backupunitId).BackupUnit(backupUnit).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupUnitApi.BackupunitsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupunitsPut`: BackupUnit
    fmt.Fprintf(os.Stdout, "Response from `BackupUnitApi.BackupunitsPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**backupunitId** | **string** | The unique ID of the backup unit | |

### Other Parameters

Other parameters are passed through a pointer to a apiBackupunitsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

| **backupUnit** | [**BackupUnit**](BackupUnit.md) | Modified backup Unit | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**BackupUnit**](BackupUnit.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## BackupunitsSsourlGet

> BackupUnitSSO BackupunitsSsourlGet(ctx, backupunitId).Pretty(pretty).XContractNumber(xContractNumber).Execute()

Returns a single signon URL for the specified backup Unit.



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
    backupunitId := "backupunitId_example" // string | The unique UUID of the backup unit
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupUnitApi.BackupunitsSsourlGet(context.Background(), backupunitId).Pretty(pretty).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupUnitApi.BackupunitsSsourlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackupunitsSsourlGet`: BackupUnitSSO
    fmt.Fprintf(os.Stdout, "Response from `BackupUnitApi.BackupunitsSsourlGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**backupunitId** | **string** | The unique UUID of the backup unit | |

### Other Parameters

Other parameters are passed through a pointer to a apiBackupunitsSsourlGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**BackupUnitSSO**](BackupUnitSSO.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


