# RequestsApi

All URIs are relative to [https://api.ionos.com/cloudapi/v6](https://api.ionos.com/cloudapi/v6)

| Method | HTTP request | Description |
| :--- | :--- | :--- |
| [**RequestsFindById**](requestsapi.md#RequestsFindById) | **Get** /requests/{requestId} | Retrieve a Request |
| [**RequestsGet**](requestsapi.md#RequestsGet) | **Get** /requests | List Requests |
| [**RequestsStatusGet**](requestsapi.md#RequestsStatusGet) | **Get** /requests/{requestId}/status | Retrieve Request Status |

## RequestsFindById

```go
var result Request = RequestsFindById(ctx, requestId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a Request

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
    requestId := "requestId_example" // string | 
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RequestsApi.RequestsFindById(context.Background(), requestId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.RequestsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestsFindById`: Request
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.RequestsFindById`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **requestId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsFindByIdRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**Request**](../models/request.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

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

List Requests

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
    filterStatus := "filterStatus_example" // string | Request filter to fetch all requests based on a particular status [QUEUED, RUNNING, DONE, FAILED]. It doesn't depend on depth query parameter (optional)
    filterCreatedAfter := "filterCreatedAfter_example" // string | Request filter to fetch all requests created after the specified date. It doesn't depend on depth query parameter. Date format e.g. 2021-01-01+00:00:00 (optional)
    filterCreatedBefore := "filterCreatedBefore_example" // string | Request filter to fetch all requests created before the specified date. It doesn't depend on depth query parameter. Date format e.g. 2021-01-01+00:00:00 (optional)
    filterCreatedDate := "filterCreatedDate_example" // string | Response filter to select and display only the requests that contains the specified createdDate. The value is case insensitive and it  depends on depth query parameter that should have a value greater than 0. Date format e.g. 2020-11-16T17:42:59Z (optional)
    filterCreatedBy := "filterCreatedBy_example" // string | Response filter to select and display only the requests that contains the specified createdBy. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0.  (optional)
    filterEtag := "filterEtag_example" // string | Response filter to select and display only the requests that contains the specified etag. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0.  (optional)
    filterRequestStatus := "filterRequestStatus_example" // string | Response filter to select and display only the requests that contains the specified requestStatus. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0.  (optional)
    filterMethod := "filterMethod_example" // string | Response filter to select and display only the requests that contains the specified method. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0.  (optional)
    filterHeaders := "filterHeaders_example" // string | Response filter to select and display only the requests that contains the specified headers. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0.  (optional)
    filterBody := "filterBody_example" // string | Response filter to select and display only the requests that contains the specified body. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0.  (optional)
    filterUrl := "filterUrl_example" // string | Response filter to select and display only the requests that contains the specified url. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0.  (optional)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with limit for pagination) (optional) (default to 0)
    limit := int32(56) // int32 | the maximum number of elements to return (use together with offset for pagination) (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RequestsApi.RequestsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).FilterStatus(filterStatus).FilterCreatedAfter(filterCreatedAfter).FilterCreatedBefore(filterCreatedBefore).FilterCreatedDate(filterCreatedDate).FilterCreatedBy(filterCreatedBy).FilterEtag(filterEtag).FilterRequestStatus(filterRequestStatus).FilterMethod(filterMethod).FilterHeaders(filterHeaders).FilterBody(filterBody).FilterUrl(filterUrl).Offset(offset).Limit(limit).Execute()
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

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |
| **filterStatus** | **string** | Request filter to fetch all requests based on a particular status \[QUEUED, RUNNING, DONE, FAILED\]. It doesn't depend on depth query parameter |  |
| **filterCreatedAfter** | **string** | Request filter to fetch all requests created after the specified date. It doesn't depend on depth query parameter. Date format e.g. 2021-01-01+00:00:00 |  |
| **filterCreatedBefore** | **string** | Request filter to fetch all requests created before the specified date. It doesn't depend on depth query parameter. Date format e.g. 2021-01-01+00:00:00 |  |
| **filterCreatedDate** | **string** | Response filter to select and display only the requests that contains the specified createdDate. The value is case insensitive and it  depends on depth query parameter that should have a value greater than 0. Date format e.g. 2020-11-16T17:42:59Z |  |
| **filterCreatedBy** | **string** | Response filter to select and display only the requests that contains the specified createdBy. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0. |  |
| **filterEtag** | **string** | Response filter to select and display only the requests that contains the specified etag. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0. |  |
| **filterRequestStatus** | **string** | Response filter to select and display only the requests that contains the specified requestStatus. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0. |  |
| **filterMethod** | **string** | Response filter to select and display only the requests that contains the specified method. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0. |  |
| **filterHeaders** | **string** | Response filter to select and display only the requests that contains the specified headers. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0. |  |
| **filterBody** | **string** | Response filter to select and display only the requests that contains the specified body. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0. |  |
| **filterUrl** | **string** | Response filter to select and display only the requests that contains the specified url. The value is case insensitive and it depends on depth query parameter that should have a value greater than 0. |  |
| **offset** | **int32** | the first element \(of the total list of elements\) to include in the response \(use together with limit for pagination\) | \[default to 0\] |
| **limit** | **int32** | the maximum number of elements to return \(use together with offset for pagination\) | \[default to 1000\] |

### Return type

[**Requests**](../models/requests.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

## RequestsStatusGet

```go
var result RequestStatus = RequestsStatusGet(ctx, requestId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve Request Status

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
    requestId := "requestId_example" // string | 
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RequestsApi.RequestsStatusGet(context.Background(), requestId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.RequestsStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestsStatusGet`: RequestStatus
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.RequestsStatusGet`: %v\n", resp)
}
```

### Path Parameters

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |  |
| **requestId** | **string** |  |  |

### Other Parameters

Other parameters are passed through a pointer to a apiRequestsStatusGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **pretty** | **bool** | Controls whether response is pretty-printed \(with indentation and new lines\) | \[default to true\] |
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/\[ID\]  - depth=0: only direct properties are included. Children \(servers etc.\) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on | \[default to 0\] |
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed |  |

### Return type

[**RequestStatus**](../models/requeststatus.md)

### HTTP request headers

* **Content-Type**: Not defined
* **Accept**: application/json

