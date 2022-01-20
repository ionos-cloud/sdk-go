# \VolumesApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersVolumesCreateSnapshotPost**](VolumesApi.md#DatacentersVolumesCreateSnapshotPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/create-snapshot | Create volume snapshots|
|[**DatacentersVolumesDelete**](VolumesApi.md#DatacentersVolumesDelete) | **Delete** /datacenters/{datacenterId}/volumes/{volumeId} | Delete volumes|
|[**DatacentersVolumesFindById**](VolumesApi.md#DatacentersVolumesFindById) | **Get** /datacenters/{datacenterId}/volumes/{volumeId} | Retrieve volumes|
|[**DatacentersVolumesGet**](VolumesApi.md#DatacentersVolumesGet) | **Get** /datacenters/{datacenterId}/volumes | List volumes|
|[**DatacentersVolumesPatch**](VolumesApi.md#DatacentersVolumesPatch) | **Patch** /datacenters/{datacenterId}/volumes/{volumeId} | Partially modify volumes|
|[**DatacentersVolumesPost**](VolumesApi.md#DatacentersVolumesPost) | **Post** /datacenters/{datacenterId}/volumes | Create volumes|
|[**DatacentersVolumesPut**](VolumesApi.md#DatacentersVolumesPut) | **Put** /datacenters/{datacenterId}/volumes/{volumeId} | Modify volumes|
|[**DatacentersVolumesRestoreSnapshotPost**](VolumesApi.md#DatacentersVolumesRestoreSnapshotPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/restore-snapshot | Restore volume snapshots|



## DatacentersVolumesCreateSnapshotPost

```go
var result Snapshot = DatacentersVolumesCreateSnapshotPost(ctx, datacenterId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Name(name)
                      .Description(description)
                      .SecAuthProtection(secAuthProtection)
                      .LicenceType(licenceType)
                      .Execute()
```

Create volume snapshots



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
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)
    name := "name_example" // string | Snapshot name (optional)
    description := "description_example" // string | Snapshot description (optional)
    secAuthProtection := true // bool | Flag for enabling extra protection for this snapshot, such as two-step verification. (optional)
    licenceType := "licenceType_example" // string | The OS type for this snapshot. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DatacentersVolumesCreateSnapshotPost(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Name(name).Description(description).SecAuthProtection(secAuthProtection).LicenceType(licenceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesCreateSnapshotPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesCreateSnapshotPost`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesCreateSnapshotPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesCreateSnapshotPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **name** | **string** | Snapshot name | |
| **description** | **string** | Snapshot description | |
| **secAuthProtection** | **bool** | Flag for enabling extra protection for this snapshot, such as two-step verification. | |
| **licenceType** | **string** | The OS type for this snapshot. | |

### Return type

[**Snapshot**](../models/Snapshot.md)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json



## DatacentersVolumesDelete

```go
var result  = DatacentersVolumesDelete(ctx, datacenterId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete volumes



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
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DatacentersVolumesDelete(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesDeleteRequest struct via the builder pattern


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



## DatacentersVolumesFindById

```go
var result Volume = DatacentersVolumesFindById(ctx, datacenterId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve volumes



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
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DatacentersVolumesFindById(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesFindById`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Volume**](../models/Volume.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersVolumesGet

```go
var result Volumes = DatacentersVolumesGet(ctx, datacenterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List volumes



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
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DatacentersVolumesGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesGet`: Volumes
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 1000]|

### Return type

[**Volumes**](../models/Volumes.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersVolumesPatch

```go
var result Volume = DatacentersVolumesPatch(ctx, datacenterId, volumeId)
                      .Volume(volume)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify volumes



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
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    volume := *openapiclient.NewVolumeProperties(float32(100.0)) // VolumeProperties | The properties of the volume to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DatacentersVolumesPatch(context.Background(), datacenterId, volumeId).Volume(volume).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesPatch`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **volume** | [**VolumeProperties**](../models/VolumeProperties.md) | The properties of the volume to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Volume**](../models/Volume.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersVolumesPost

```go
var result Volume = DatacentersVolumesPost(ctx, datacenterId)
                      .Volume(volume)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create volumes



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
    volume := *openapiclient.NewVolume(*openapiclient.NewVolumeProperties(float32(100.0))) // Volume | The volume to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DatacentersVolumesPost(context.Background(), datacenterId).Volume(volume).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesPost`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **volume** | [**Volume**](../models/Volume.md) | The volume to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Volume**](../models/Volume.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersVolumesPut

```go
var result Volume = DatacentersVolumesPut(ctx, datacenterId, volumeId)
                      .Volume(volume)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify volumes



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
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    volume := *openapiclient.NewVolume(*openapiclient.NewVolumeProperties(float32(100.0))) // Volume | The modified volume
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DatacentersVolumesPut(context.Background(), datacenterId, volumeId).Volume(volume).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesPut`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **volume** | [**Volume**](../models/Volume.md) | The modified volume | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Volume**](../models/Volume.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersVolumesRestoreSnapshotPost

```go
var result  = DatacentersVolumesRestoreSnapshotPost(ctx, datacenterId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .SnapshotId(snapshotId)
                      .Execute()
```

Restore volume snapshots



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
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)
    snapshotId := "snapshotId_example" // string | The unique ID of the snapshot. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DatacentersVolumesRestoreSnapshotPost(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).SnapshotId(snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesRestoreSnapshotPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesRestoreSnapshotPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **snapshotId** | **string** | The unique ID of the snapshot. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json


