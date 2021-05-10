# \UserS3KeysApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**UmUsersS3keysDelete**](UserS3KeysApi.md#UmUsersS3keysDelete) | **Delete** /um/users/{userId}/s3keys/{keyId} | Delete an S3 Key|
|[**UmUsersS3keysFindByKeyId**](UserS3KeysApi.md#UmUsersS3keysFindByKeyId) | **Get** /um/users/{userId}/s3keys/{keyId} | Retrieve given S3 Key belonging to the given User|
|[**UmUsersS3keysGet**](UserS3KeysApi.md#UmUsersS3keysGet) | **Get** /um/users/{userId}/s3keys | Retrieve a User&#39;s S3 keys|
|[**UmUsersS3keysPost**](UserS3KeysApi.md#UmUsersS3keysPost) | **Post** /um/users/{userId}/s3keys | Create a S3 Key for the given User|
|[**UmUsersS3keysPut**](UserS3KeysApi.md#UmUsersS3keysPut) | **Put** /um/users/{userId}/s3keys/{keyId} | Modify a S3 key having the given key id|
|[**UmUsersS3ssourlGet**](UserS3KeysApi.md#UmUsersS3ssourlGet) | **Get** /um/users/{userId}/s3ssourl | Retrieve S3 object storage single signon URL for the given user|



## UmUsersS3keysDelete

```go
var result map[string]interface{} = UmUsersS3keysDelete(ctx, userId, keyId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete an S3 Key



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
    userId := "userId_example" // string | The unique ID of the user
    keyId := "keyId_example" // string | The unique access key ID of the S3 key
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserS3KeysApi.UmUsersS3keysDelete(context.Background(), userId, keyId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserS3KeysApi.UmUsersS3keysDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UmUsersS3keysDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserS3KeysApi.UmUsersS3keysDelete`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**userId** | **string** | The unique ID of the user | |
|**keyId** | **string** | The unique access key ID of the S3 key | |

### Other Parameters

Other parameters are passed through a pointer to a apiUmUsersS3keysDeleteRequest struct via the builder pattern


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



## UmUsersS3keysFindByKeyId

```go
var result S3Key = UmUsersS3keysFindByKeyId(ctx, userId, keyId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve given S3 Key belonging to the given User



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
    userId := "userId_example" // string | The unique ID of the user
    keyId := "keyId_example" // string | The unique access key ID of the S3 key
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserS3KeysApi.UmUsersS3keysFindByKeyId(context.Background(), userId, keyId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserS3KeysApi.UmUsersS3keysFindByKeyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UmUsersS3keysFindByKeyId`: S3Key
    fmt.Fprintf(os.Stdout, "Response from `UserS3KeysApi.UmUsersS3keysFindByKeyId`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**userId** | **string** | The unique ID of the user | |
|**keyId** | **string** | The unique access key ID of the S3 key | |

### Other Parameters

Other parameters are passed through a pointer to a apiUmUsersS3keysFindByKeyIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**S3Key**](S3Key.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## UmUsersS3keysGet

```go
var result S3Keys = UmUsersS3keysGet(ctx, userId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a User's S3 keys



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
    userId := "userId_example" // string | The unique ID of the user
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserS3KeysApi.UmUsersS3keysGet(context.Background(), userId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserS3KeysApi.UmUsersS3keysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UmUsersS3keysGet`: S3Keys
    fmt.Fprintf(os.Stdout, "Response from `UserS3KeysApi.UmUsersS3keysGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**userId** | **string** | The unique ID of the user | |

### Other Parameters

Other parameters are passed through a pointer to a apiUmUsersS3keysGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**S3Keys**](S3Keys.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## UmUsersS3keysPost

```go
var result S3Key = UmUsersS3keysPost(ctx, userId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a S3 Key for the given User



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
    userId := "userId_example" // string | The unique ID of the user
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserS3KeysApi.UmUsersS3keysPost(context.Background(), userId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserS3KeysApi.UmUsersS3keysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UmUsersS3keysPost`: S3Key
    fmt.Fprintf(os.Stdout, "Response from `UserS3KeysApi.UmUsersS3keysPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**userId** | **string** | The unique ID of the user | |

### Other Parameters

Other parameters are passed through a pointer to a apiUmUsersS3keysPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**S3Key**](S3Key.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## UmUsersS3keysPut

```go
var result S3Key = UmUsersS3keysPut(ctx, userId, keyId)
                      .S3Key(s3Key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a S3 key having the given key id



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
    userId := "userId_example" // string | 
    keyId := "keyId_example" // string | The unique access key ID of the S3 key
    s3Key := *openapiclient.NewS3Key(*openapiclient.NewS3KeyProperties()) // S3Key | Modified S3 key
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserS3KeysApi.UmUsersS3keysPut(context.Background(), userId, keyId).S3Key(s3Key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserS3KeysApi.UmUsersS3keysPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UmUsersS3keysPut`: S3Key
    fmt.Fprintf(os.Stdout, "Response from `UserS3KeysApi.UmUsersS3keysPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**userId** | **string** |  | |
|**keyId** | **string** | The unique access key ID of the S3 key | |

### Other Parameters

Other parameters are passed through a pointer to a apiUmUsersS3keysPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **s3Key** | [**S3Key**](S3Key.md) | Modified S3 key | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**S3Key**](S3Key.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## UmUsersS3ssourlGet

```go
var result S3ObjectStorageSSO = UmUsersS3ssourlGet(ctx, userId)
                      .Pretty(pretty)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve S3 object storage single signon URL for the given user



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
    userId := "userId_example" // string | The unique ID of the user
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserS3KeysApi.UmUsersS3ssourlGet(context.Background(), userId).Pretty(pretty).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserS3KeysApi.UmUsersS3ssourlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UmUsersS3ssourlGet`: S3ObjectStorageSSO
    fmt.Fprintf(os.Stdout, "Response from `UserS3KeysApi.UmUsersS3ssourlGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**userId** | **string** | The unique ID of the user | |

### Other Parameters

Other parameters are passed through a pointer to a apiUmUsersS3ssourlGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**S3ObjectStorageSSO**](S3ObjectStorageSSO.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


