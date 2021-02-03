# \DataCenterApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersDelete**](DataCenterApi.md#DatacentersDelete) | **Delete** /datacenters/{datacenterId} | Delete a Data Center
[**DatacentersFindById**](DataCenterApi.md#DatacentersFindById) | **Get** /datacenters/{datacenterId} | Retrieve a Data Center
[**DatacentersGet**](DataCenterApi.md#DatacentersGet) | **Get** /datacenters | List Data Centers under your account
[**DatacentersPatch**](DataCenterApi.md#DatacentersPatch) | **Patch** /datacenters/{datacenterId} | Partially modify a Data Center
[**DatacentersPost**](DataCenterApi.md#DatacentersPost) | **Post** /datacenters | Create a Data Center
[**DatacentersPut**](DataCenterApi.md#DatacentersPut) | **Put** /datacenters/{datacenterId} | Modify a Data Center



## DatacentersDelete

> map[string]interface{} DatacentersDelete(ctx, datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Delete a Data Center



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataCenterApi.DatacentersDelete(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataCenterApi.DatacentersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataCenterApi.DatacentersDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersDeleteRequest struct via the builder pattern


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


## DatacentersFindById

> Datacenter DatacentersFindById(ctx, datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Retrieve a Data Center



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataCenterApi.DatacentersFindById(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataCenterApi.DatacentersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersFindById`: Datacenter
    fmt.Fprintf(os.Stdout, "Response from `DataCenterApi.DatacentersFindById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersFindByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Datacenter**](Datacenter.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersGet

> Datacenters DatacentersGet(ctx).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()

List Data Centers under your account



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
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with <code>limit</code> for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with <code>offset</code> for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataCenterApi.DatacentersGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataCenterApi.DatacentersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersGet`: Datacenters
    fmt.Fprintf(os.Stdout, "Response from `DataCenterApi.DatacentersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 
 **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with &lt;code&gt;limit&lt;/code&gt; for pagination) | [default to 0]
 **limit** | **int32** | the maximum number of elements to return (use together with &lt;code&gt;offset&lt;/code&gt; for pagination) | [default to 1000]

### Return type

[**Datacenters**](Datacenters.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersPatch

> Datacenter DatacentersPatch(ctx, datacenterId).Datacenter(datacenter).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Partially modify a Data Center



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
    datacenter := *openapiclient.NewDatacenterProperties("us/las") // DatacenterProperties | Modified properties of Data Center
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataCenterApi.DatacentersPatch(context.Background(), datacenterId).Datacenter(datacenter).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataCenterApi.DatacentersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersPatch`: Datacenter
    fmt.Fprintf(os.Stdout, "Response from `DataCenterApi.DatacentersPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datacenter** | [**DatacenterProperties**](DatacenterProperties.md) | Modified properties of Data Center | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Datacenter**](Datacenter.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersPost

> Datacenter DatacentersPost(ctx).Datacenter(datacenter).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Create a Data Center



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
    datacenter := *openapiclient.NewDatacenter(*openapiclient.NewDatacenterProperties("us/las")) // Datacenter | Datacenter to be created
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataCenterApi.DatacentersPost(context.Background()).Datacenter(datacenter).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataCenterApi.DatacentersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersPost`: Datacenter
    fmt.Fprintf(os.Stdout, "Response from `DataCenterApi.DatacentersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datacenter** | [**Datacenter**](Datacenter.md) | Datacenter to be created | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Datacenter**](Datacenter.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersPut

> Datacenter DatacentersPut(ctx, datacenterId).Datacenter(datacenter).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()

Modify a Data Center



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
    datacenter := *openapiclient.NewDatacenter(*openapiclient.NewDatacenterProperties("us/las")) // Datacenter | Modified Data Center
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataCenterApi.DatacentersPut(context.Background(), datacenterId).Datacenter(datacenter).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataCenterApi.DatacentersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersPut`: Datacenter
    fmt.Fprintf(os.Stdout, "Response from `DataCenterApi.DatacentersPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string** | The unique ID of the datacenter | 

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datacenter** | [**Datacenter**](Datacenter.md) | Modified Data Center | 
 **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Datacenter**](Datacenter.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

