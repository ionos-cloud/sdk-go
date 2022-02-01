# \ServerApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersServersCdromsDelete**](ServerApi.md#DatacentersServersCdromsDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/cdroms/{cdromId} | Detach a CD-ROM|
|[**DatacentersServersCdromsFindById**](ServerApi.md#DatacentersServersCdromsFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/cdroms/{cdromId} | Retrieve an attached CD-ROM|
|[**DatacentersServersCdromsGet**](ServerApi.md#DatacentersServersCdromsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/cdroms | List attached CD-ROMs |
|[**DatacentersServersCdromsPost**](ServerApi.md#DatacentersServersCdromsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/cdroms | Attach a CD-ROM|
|[**DatacentersServersDelete**](ServerApi.md#DatacentersServersDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId} | Delete a Server|
|[**DatacentersServersFindById**](ServerApi.md#DatacentersServersFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId} | Retrieve a Server|
|[**DatacentersServersGet**](ServerApi.md#DatacentersServersGet) | **Get** /datacenters/{datacenterId}/servers | List Servers |
|[**DatacentersServersPatch**](ServerApi.md#DatacentersServersPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId} | Partially modify a Server|
|[**DatacentersServersPost**](ServerApi.md#DatacentersServersPost) | **Post** /datacenters/{datacenterId}/servers | Create a Server|
|[**DatacentersServersPut**](ServerApi.md#DatacentersServersPut) | **Put** /datacenters/{datacenterId}/servers/{serverId} | Modify a Server|
|[**DatacentersServersRebootPost**](ServerApi.md#DatacentersServersRebootPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/reboot | Reboot a Server|
|[**DatacentersServersStartPost**](ServerApi.md#DatacentersServersStartPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/start | Start a Server|
|[**DatacentersServersStopPost**](ServerApi.md#DatacentersServersStopPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/stop | Stop a Server|
|[**DatacentersServersUpgradePost**](ServerApi.md#DatacentersServersUpgradePost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/upgrade | Upgrade a Server|
|[**DatacentersServersVolumesDelete**](ServerApi.md#DatacentersServersVolumesDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/volumes/{volumeId} | Detach a volume|
|[**DatacentersServersVolumesFindById**](ServerApi.md#DatacentersServersVolumesFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/volumes/{volumeId} | Retrieve an attached volume|
|[**DatacentersServersVolumesGet**](ServerApi.md#DatacentersServersVolumesGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/volumes | List Attached Volumes|
|[**DatacentersServersVolumesPost**](ServerApi.md#DatacentersServersVolumesPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/volumes | Attach a volume|



## DatacentersServersCdromsDelete

```go
var result map[string]interface{} = DatacentersServersCdromsDelete(ctx, datacenterId, serverId, cdromId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Detach a CD-ROM



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
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    cdromId := "cdromId_example" // string | The unique ID of the CD-ROM
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersCdromsDelete(context.Background(), datacenterId, serverId, cdromId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersCdromsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersCdromsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersCdromsDelete`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the Datacenter | |
|**serverId** | **string** | The unique ID of the Server | |
|**cdromId** | **string** | The unique ID of the CD-ROM | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersCdromsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersCdromsFindById

```go
var result Image = DatacentersServersCdromsFindById(ctx, datacenterId, serverId, cdromId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve an attached CD-ROM



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
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    cdromId := "cdromId_example" // string | The unique ID of the CD-ROM
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersCdromsFindById(context.Background(), datacenterId, serverId, cdromId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersCdromsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersCdromsFindById`: Image
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersCdromsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the Datacenter | |
|**serverId** | **string** | The unique ID of the Server | |
|**cdromId** | **string** | The unique ID of the CD-ROM | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersCdromsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Image**](../models/Image.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersCdromsGet

```go
var result Cdroms = DatacentersServersCdromsGet(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List attached CD-ROMs 



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
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with <code>limit</code> for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with <code>offset</code> for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersCdromsGet(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersCdromsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersCdromsGet`: Cdroms
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersCdromsGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the Datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersCdromsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |
| **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with &lt;code&gt;limit&lt;/code&gt; for pagination) | [default to 0]|
| **limit** | **int32** | the maximum number of elements to return (use together with &lt;code&gt;offset&lt;/code&gt; for pagination) | [default to 1000]|

### Return type

[**Cdroms**](../models/Cdroms.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersCdromsPost

```go
var result Image = DatacentersServersCdromsPost(ctx, datacenterId, serverId)
                      .Cdrom(cdrom)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Attach a CD-ROM



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
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    cdrom := *openapiclient.NewImage(*openapiclient.NewImageProperties("LINUX")) // Image | CD-ROM to be attached
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersCdromsPost(context.Background(), datacenterId, serverId).Cdrom(cdrom).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersCdromsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersCdromsPost`: Image
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersCdromsPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the Datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersCdromsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **cdrom** | [**Image**](../models/Image.md) | CD-ROM to be attached | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Image**](../models/Image.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersDelete

```go
var result map[string]interface{} = DatacentersServersDelete(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete a Server



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersDelete(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersDelete`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersFindById

```go
var result Server = DatacentersServersFindById(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a Server



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersFindById(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersFindById`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Server**](../models/Server.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersGet

```go
var result Servers = DatacentersServersGet(ctx, datacenterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .UpgradeNeeded(upgradeNeeded)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List Servers 



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
    upgradeNeeded := true // bool | It can be used to filter which servers can be upgraded which can not be upgraded. (optional)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with <code>limit</code> for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with <code>offset</code> for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).UpgradeNeeded(upgradeNeeded).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersGet`: Servers
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **upgradeNeeded** | **bool** | It can be used to filter which servers can be upgraded which can not be upgraded. | |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |
| **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with &lt;code&gt;limit&lt;/code&gt; for pagination) | [default to 0]|
| **limit** | **int32** | the maximum number of elements to return (use together with &lt;code&gt;offset&lt;/code&gt; for pagination) | [default to 1000]|

### Return type

[**Servers**](../models/Servers.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersPatch

```go
var result Server = DatacentersServersPatch(ctx, datacenterId, serverId)
                      .Server(server)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify a Server



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
    server := *openapiclient.NewServerProperties(int32(4), int32(4096)) // ServerProperties | Modified properties of Server
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersPatch(context.Background(), datacenterId, serverId).Server(server).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersPatch`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **server** | [**ServerProperties**](../models/ServerProperties.md) | Modified properties of Server | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Server**](../models/Server.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersPost

```go
var result Server = DatacentersServersPost(ctx, datacenterId)
                      .Server(server)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Server



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
    server := *openapiclient.NewServer(*openapiclient.NewServerProperties(int32(4), int32(4096))) // Server | Server to be created
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersPost(context.Background(), datacenterId).Server(server).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersPost`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **server** | [**Server**](../models/Server.md) | Server to be created | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Server**](../models/Server.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersPut

```go
var result Server = DatacentersServersPut(ctx, datacenterId, serverId)
                      .Server(server)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Server



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
    server := *openapiclient.NewServer(*openapiclient.NewServerProperties(int32(4), int32(4096))) // Server | Modified Server
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersPut(context.Background(), datacenterId, serverId).Server(server).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersPut`: Server
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **server** | [**Server**](../models/Server.md) | Modified Server | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Server**](../models/Server.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersRebootPost

```go
var result map[string]interface{} = DatacentersServersRebootPost(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Reboot a Server



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersRebootPost(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersRebootPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersRebootPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersRebootPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersRebootPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersStartPost

```go
var result map[string]interface{} = DatacentersServersStartPost(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Start a Server



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersStartPost(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersStartPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersStartPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersStartPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersStartPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersStopPost

```go
var result map[string]interface{} = DatacentersServersStopPost(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Stop a Server



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersStopPost(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersStopPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersStopPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersStopPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersStopPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersUpgradePost

```go
var result map[string]interface{} = DatacentersServersUpgradePost(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Upgrade a Server



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersUpgradePost(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersUpgradePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersUpgradePost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersUpgradePost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersUpgradePostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersVolumesDelete

```go
var result map[string]interface{} = DatacentersServersVolumesDelete(ctx, datacenterId, serverId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Detach a volume



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
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersVolumesDelete(context.Background(), datacenterId, serverId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersVolumesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersVolumesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersVolumesDelete`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the Datacenter | |
|**serverId** | **string** | The unique ID of the Server | |
|**volumeId** | **string** | The unique ID of the Volume | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersVolumesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersVolumesFindById

```go
var result Volume = DatacentersServersVolumesFindById(ctx, datacenterId, serverId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve an attached volume



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
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersVolumesFindById(context.Background(), datacenterId, serverId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersVolumesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersVolumesFindById`: Volume
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersVolumesFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the Datacenter | |
|**serverId** | **string** | The unique ID of the Server | |
|**volumeId** | **string** | The unique ID of the Volume | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersVolumesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Volume**](../models/Volume.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersVolumesGet

```go
var result AttachedVolumes = DatacentersServersVolumesGet(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List Attached Volumes



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
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with <code>limit</code> for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with <code>offset</code> for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersVolumesGet(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersVolumesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersVolumesGet`: AttachedVolumes
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersVolumesGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the Datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersVolumesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |
| **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with &lt;code&gt;limit&lt;/code&gt; for pagination) | [default to 0]|
| **limit** | **int32** | the maximum number of elements to return (use together with &lt;code&gt;offset&lt;/code&gt; for pagination) | [default to 1000]|

### Return type

[**AttachedVolumes**](../models/AttachedVolumes.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersVolumesPost

```go
var result Volume = DatacentersServersVolumesPost(ctx, datacenterId, serverId)
                      .Volume(volume)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Attach a volume



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
    datacenterId := "datacenterId_example" // string | The unique ID of the Datacenter
    serverId := "serverId_example" // string | The unique ID of the Server
    volume := *openapiclient.NewVolume(*openapiclient.NewVolumeProperties(float32(100.0))) // Volume | Volume to be attached (created and attached)
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServerApi.DatacentersServersVolumesPost(context.Background(), datacenterId, serverId).Volume(volume).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServerApi.DatacentersServersVolumesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersVolumesPost`: Volume
    fmt.Fprintf(os.Stdout, "Response from `ServerApi.DatacentersServersVolumesPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the Datacenter | |
|**serverId** | **string** | The unique ID of the Server | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersVolumesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **volume** | [**Volume**](../models/Volume.md) | Volume to be attached (created and attached) | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Volume**](../models/Volume.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


