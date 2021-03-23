# VolumesApi

All URIs are relative to [https://api.ionos.com/cloudapi/v6](https://api.ionos.com/cloudapi/v6)

| Method | HTTP request | Description |
| :--- | :--- | :--- |
| [**DatacentersVolumesCreateSnapshotPost**](volumesapi.md#DatacentersVolumesCreateSnapshotPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/create-snapshot | Create Volume Snapshot |
| [**DatacentersVolumesDelete**](volumesapi.md#DatacentersVolumesDelete) | **Delete** /datacenters/{datacenterId}/volumes/{volumeId} | Delete a Volume |
| [**DatacentersVolumesFindById**](volumesapi.md#DatacentersVolumesFindById) | **Get** /datacenters/{datacenterId}/volumes/{volumeId} | Retrieve a Volume |
| [**DatacentersVolumesGet**](volumesapi.md#DatacentersVolumesGet) | **Get** /datacenters/{datacenterId}/volumes | List Volumes |
| [**DatacentersVolumesPatch**](volumesapi.md#DatacentersVolumesPatch) | **Patch** /datacenters/{datacenterId}/volumes/{volumeId} | Partially modify a Volume |
| [**DatacentersVolumesPost**](volumesapi.md#DatacentersVolumesPost) | **Post** /datacenters/{datacenterId}/volumes | Create a Volume |
| [**DatacentersVolumesPut**](volumesapi.md#DatacentersVolumesPut) | **Put** /datacenters/{datacenterId}/volumes/{volumeId} | Modify a Volume |
| [**DatacentersVolumesRestoreSnapshotPost**](volumesapi.md#DatacentersVolumesRestoreSnapshotPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/restore-snapshot | Restore Volume Snapshot |

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

Create Volume Snapshot

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
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    name := "name_example" // string | The name of the snapshot (optional)
    description := "description_example" // string | The description of the snapshot (optional)
    secAuthProtection := true // bool | Flag representing if extra protection is enabled on snapshot e.g. Two Factor protection etc. (optional)
    licenceType := "licenceType_example" // string | The OS type of this Snapshot (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DatacentersVolumesCreateSnapshotPost(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Name(name).Description(description).SecAuthProtection(secAuthProtection).LicenceType(licenceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesCreateSnapshotPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesCreateSnapshotPost`: Snapshot
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesCreateSnapshotPost`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **volumeId** | **string** | The unique ID of the Volume |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesCreateSnapshotPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users with more than one contract must provide the contract number, against which all API requests are to be executed. |  |
| **name** | **string** | The name of the snapshot |  |
| **description** | **string** | The description of the snapshot |  |
| **secAuthProtection** | **bool** | Flag representing if extra protection is enabled on snapshot e.g. Two Factor protection etc. |  |
| **licenceType** | **string** | The OS type of this Snapshot |  |

### Return type

[**Snapshot**](../models/snapshot.md)

### HTTP request headers

* **Content-Type**: application/x-www-form-urlencoded
* **Accept**: application/json

## DatacentersVolumesDelete

```go
var result map[string]interface{} = DatacentersVolumesDelete(ctx, datacenterId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete a Volume

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
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DatacentersVolumesDelete(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesDelete`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **volumeId** | **string** | The unique ID of the Volume |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users with more than one contract must provide the contract number, against which all API requests are to be executed. |  |

### Return type

**map\[string\]interface{}**

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersVolumesFindById

```go
var result Volume = DatacentersVolumesFindById(ctx, datacenterId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a Volume

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
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DatacentersVolumesFindById(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesFindById`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesFindById`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **volumeId** | **string** | The unique ID of the Volume |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesFindByIdRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users with more than one contract must provide the contract number, against which all API requests are to be executed. |  |

### Return type

[**Volume**](../models/volume.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

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

List Volumes

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
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with limit for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with offset for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DatacentersVolumesGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesGet`: Volumes
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users with more than one contract must provide the contract number, against which all API requests are to be executed. |  |
| **offset** | **int32** | the first element \(of the total list of elements\) to include in the response \(use together with limit for pagination\) | \[default to 0\] |
| **limit** | **int32** | the maximum number of elements to return \(use together with offset for pagination\) | \[default to 1000\] |

### Return type

[**Volumes**](../models/volumes.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## DatacentersVolumesPatch

```go
var result Volume = DatacentersVolumesPatch(ctx, datacenterId, volumeId)
                      .Volume(volume)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify a Volume

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
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    volume := *openapiclient.NewVolumeProperties(float32(100.0)) // VolumeProperties | Modified properties of Volume
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DatacentersVolumesPatch(context.Background(), datacenterId, volumeId).Volume(volume).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesPatch`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesPatch`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **volumeId** | **string** | The unique ID of the Volume |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesPatchRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **volume** | [**VolumeProperties**](../models/volumeproperties.md) | Modified properties of Volume |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users with more than one contract must provide the contract number, against which all API requests are to be executed. |  |

### Return type

[**Volume**](../models/volume.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## DatacentersVolumesPost

```go
var result Volume = DatacentersVolumesPost(ctx, datacenterId)
                      .Volume(volume)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Volume

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
    volume := *openapiclient.NewVolume(*openapiclient.NewVolumeProperties(float32(100.0))) // Volume | Volume to be created
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DatacentersVolumesPost(context.Background(), datacenterId).Volume(volume).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesPost`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesPost`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **volume** | [**Volume**](../models/volume.md) | Volume to be created |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users with more than one contract must provide the contract number, against which all API requests are to be executed. |  |

### Return type

[**Volume**](../models/volume.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## DatacentersVolumesPut

```go
var result Volume = DatacentersVolumesPut(ctx, datacenterId, volumeId)
                      .Volume(volume)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Volume

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
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    volume := *openapiclient.NewVolume(*openapiclient.NewVolumeProperties(float32(100.0))) // Volume | Modified Volume
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DatacentersVolumesPut(context.Background(), datacenterId, volumeId).Volume(volume).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesPut`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesPut`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **volumeId** | **string** | The unique ID of the Volume |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **volume** | [**Volume**](../models/volume.md) | Modified Volume |  |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users with more than one contract must provide the contract number, against which all API requests are to be executed. |  |

### Return type

[**Volume**](../models/volume.md)

### HTTP request headers

* **Content-Type**: application/json
* **Accept**: application/json

## DatacentersVolumesRestoreSnapshotPost

```go
var result map[string]interface{} = DatacentersVolumesRestoreSnapshotPost(ctx, datacenterId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .SnapshotId(snapshotId)
                      .Execute()
```

Restore Volume Snapshot

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
    volumeId := "volumeId_example" // string | The unique ID of the Volume
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)
    snapshotId := "snapshotId_example" // string | This is the ID of the snapshot (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VolumesApi.DatacentersVolumesRestoreSnapshotPost(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).SnapshotId(snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DatacentersVolumesRestoreSnapshotPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersVolumesRestoreSnapshotPost`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.DatacentersVolumesRestoreSnapshotPost`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **datacenterId** | **string** | The unique ID of the datacenter |  |
| **volumeId** | **string** | The unique ID of the Volume |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersVolumesRestoreSnapshotPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users with more than one contract must provide the contract number, against which all API requests are to be executed. |  |
| **snapshotId** | **string** | This is the ID of the snapshot |  |

### Return type

**map\[string\]interface{}**

### HTTP request headers

* **Content-Type**: application/x-www-form-urlencoded
* **Accept**: application/json

