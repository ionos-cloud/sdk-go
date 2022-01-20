# \RequestsApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RequestsFindById**](RequestsApi.md#RequestsFindById) | **Get** /requests/{requestId} | Retrieve requests|
|[**RequestsGet**](RequestsApi.md#RequestsGet) | **Get** /requests | List requests|
|[**RequestsStatusGet**](RequestsApi.md#RequestsStatusGet) | **Get** /requests/{requestId}/status | Retrieve request status|



## RequestsFindById

```go
var result Request = RequestsFindById(ctx, requestId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve requests



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
    requestId := "requestId_example" // string | The unique ID of the request.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.RequestsFindById(context.Background(), requestId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.RequestsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestsFindById`: Request
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.RequestsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**requestId** | **string** | The unique ID of the request. | |

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Request**](../models/Request.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RequestsGet

```go
var result Requests = RequestsGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .FilterStatus(filterStatus)
                      .FilterCreatedAfter(filterCreatedAfter)
                      .FilterCreatedBefore(filterCreatedBefore)
                      .FilterCreatedDate(filterCreatedDate)
                      .FilterCreatedBy(filterCreatedBy)
                      .FilterEtag(filterEtag)
                      .FilterRequestStatus(filterRequestStatus)
                      .FilterMethod(filterMethod)
                      .FilterHeaders(filterHeaders)
                      .FilterBody(filterBody)
                      .FilterUrl(filterUrl)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List requests



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
    filterStatus := "filterStatus_example" // string | Filter the list by request status [QUEUED, RUNNING, DONE, FAILED]. Filter is not affected by the depth query parameter. (optional)
    filterCreatedAfter := "filterCreatedAfter_example" // string | Filter the list to only include the requests created after the date, specified in the yyyy-MM-dd HH:mm:ss format. Filter is not affected by the depth query parameter. (optional)
    filterCreatedBefore := "filterCreatedBefore_example" // string | Filter the list to only include the requests created before the date, specified in the yyyy-MM-dd HH:mm:ss format. Filter is not affected by the depth query parameter. (optional)
    filterCreatedDate := "filterCreatedDate_example" // string | Filter the list to only include the requests that contain the createdDate, specified in the yyyy-MM-dd HH:mm:ss format. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero. (optional)
    filterCreatedBy := "filterCreatedBy_example" // string | Filter the list to only include the requests that contain the createdBy, specified in the yyyy-MM-dd HH:mm:ss format. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  (optional)
    filterEtag := "filterEtag_example" // string | Filter the list to only include the requests that contain the specified etag. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  (optional)
    filterRequestStatus := "filterRequestStatus_example" // string | Filter the list to only include the requests that contain the specified requestStatus. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  (optional)
    filterMethod := "filterMethod_example" // string | Filter the list to only include the requests that contain the specified method. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  (optional)
    filterHeaders := "filterHeaders_example" // string | Filter the list to only include the requests that contain the specified headers. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  (optional)
    filterBody := "filterBody_example" // string | Filter the list to only include the requests that contain the specified body. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  (optional)
    filterUrl := "filterUrl_example" // string | Filter the list to only include the requests that contain the specified URL. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  (optional)
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.RequestsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).FilterStatus(filterStatus).FilterCreatedAfter(filterCreatedAfter).FilterCreatedBefore(filterCreatedBefore).FilterCreatedDate(filterCreatedDate).FilterCreatedBy(filterCreatedBy).FilterEtag(filterEtag).FilterRequestStatus(filterRequestStatus).FilterMethod(filterMethod).FilterHeaders(filterHeaders).FilterBody(filterBody).FilterUrl(filterUrl).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.RequestsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestsGet`: Requests
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.RequestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **filterStatus** | **string** | Filter the list by request status [QUEUED, RUNNING, DONE, FAILED]. Filter is not affected by the depth query parameter. | |
| **filterCreatedAfter** | **string** | Filter the list to only include the requests created after the date, specified in the yyyy-MM-dd HH:mm:ss format. Filter is not affected by the depth query parameter. | |
| **filterCreatedBefore** | **string** | Filter the list to only include the requests created before the date, specified in the yyyy-MM-dd HH:mm:ss format. Filter is not affected by the depth query parameter. | |
| **filterCreatedDate** | **string** | Filter the list to only include the requests that contain the createdDate, specified in the yyyy-MM-dd HH:mm:ss format. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero. | |
| **filterCreatedBy** | **string** | Filter the list to only include the requests that contain the createdBy, specified in the yyyy-MM-dd HH:mm:ss format. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  | |
| **filterEtag** | **string** | Filter the list to only include the requests that contain the specified etag. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  | |
| **filterRequestStatus** | **string** | Filter the list to only include the requests that contain the specified requestStatus. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  | |
| **filterMethod** | **string** | Filter the list to only include the requests that contain the specified method. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  | |
| **filterHeaders** | **string** | Filter the list to only include the requests that contain the specified headers. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  | |
| **filterBody** | **string** | Filter the list to only include the requests that contain the specified body. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  | |
| **filterUrl** | **string** | Filter the list to only include the requests that contain the specified URL. The value is not case-sensitive, and the filter requires that the depth query parameter value is greater than zero.  | |
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 1000]|

### Return type

[**Requests**](../models/Requests.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RequestsStatusGet

```go
var result RequestStatus = RequestsStatusGet(ctx, requestId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve request status



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
    requestId := "requestId_example" // string | The unique ID of the request.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.RequestsStatusGet(context.Background(), requestId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.RequestsStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestsStatusGet`: RequestStatus
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.RequestsStatusGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**requestId** | **string** | The unique ID of the request. | |

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsStatusGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**RequestStatus**](../models/RequestStatus.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


