# \ImageApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ImagesDelete**](ImageApi.md#ImagesDelete) | **Delete** /images/{imageId} | Delete an Image|
|[**ImagesFindById**](ImageApi.md#ImagesFindById) | **Get** /images/{imageId} | Retrieve an Image|
|[**ImagesGet**](ImageApi.md#ImagesGet) | **Get** /images | List Images |
|[**ImagesPatch**](ImageApi.md#ImagesPatch) | **Patch** /images/{imageId} | Partially modify an Image|
|[**ImagesPut**](ImageApi.md#ImagesPut) | **Put** /images/{imageId} | Modify an Image|



## ImagesDelete

```go
var result map[string]interface{} = ImagesDelete(ctx, imageId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete an Image



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
    imageId := "imageId_example" // string | 
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageApi.ImagesDelete(context.Background(), imageId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.ImagesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.ImagesDelete`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**imageId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to a apiImagesDeleteRequest struct via the builder pattern


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



## ImagesFindById

```go
var result Image = ImagesFindById(ctx, imageId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve an Image



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
    imageId := "imageId_example" // string | 
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageApi.ImagesFindById(context.Background(), imageId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.ImagesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesFindById`: Image
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.ImagesFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**imageId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to a apiImagesFindByIdRequest struct via the builder pattern


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



## ImagesGet

```go
var result Images = ImagesGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Images 



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageApi.ImagesGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.ImagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesGet`: Images
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.ImagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Images**](../models/Images.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ImagesPatch

```go
var result Image = ImagesPatch(ctx, imageId)
                      .Image(image)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify an Image



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
    imageId := "imageId_example" // string | 
    image := *openapiclient.NewImageProperties("LINUX") // ImageProperties | Modified Image
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageApi.ImagesPatch(context.Background(), imageId).Image(image).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.ImagesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesPatch`: Image
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.ImagesPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**imageId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to a apiImagesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **image** | [**ImageProperties**](../models/ImageProperties.md) | Modified Image | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Image**](../models/Image.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ImagesPut

```go
var result Image = ImagesPut(ctx, imageId)
                      .Image(image)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify an Image



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
    imageId := "imageId_example" // string | 
    image := *openapiclient.NewImage(*openapiclient.NewImageProperties("LINUX")) // Image | Modified Image
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageApi.ImagesPut(context.Background(), imageId).Image(image).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageApi.ImagesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImagesPut`: Image
    fmt.Fprintf(os.Stdout, "Response from `ImageApi.ImagesPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**imageId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to a apiImagesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **image** | [**Image**](../models/Image.md) | Modified Image | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Image**](../models/Image.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


