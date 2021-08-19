# \LoadBalancersApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersLoadbalancersBalancednicsDelete**](LoadBalancersApi.md#DatacentersLoadbalancersBalancednicsDelete) | **Delete** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics/{nicId} | Detach a nic from loadbalancer|
|[**DatacentersLoadbalancersBalancednicsFindByNicId**](LoadBalancersApi.md#DatacentersLoadbalancersBalancednicsFindByNicId) | **Get** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics/{nicId} | Retrieve a network interface attached to Load Balancer|
|[**DatacentersLoadbalancersBalancednicsGet**](LoadBalancersApi.md#DatacentersLoadbalancersBalancednicsGet) | **Get** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics | List Load Balancer balaced NICs|
|[**DatacentersLoadbalancersBalancednicsPost**](LoadBalancersApi.md#DatacentersLoadbalancersBalancednicsPost) | **Post** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics | Attach a nic to Load Balancer|
|[**DatacentersLoadbalancersDelete**](LoadBalancersApi.md#DatacentersLoadbalancersDelete) | **Delete** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Delete a Loadbalancer.|
|[**DatacentersLoadbalancersFindById**](LoadBalancersApi.md#DatacentersLoadbalancersFindById) | **Get** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Retrieve a loadbalancer|
|[**DatacentersLoadbalancersGet**](LoadBalancersApi.md#DatacentersLoadbalancersGet) | **Get** /datacenters/{datacenterId}/loadbalancers | List Load Balancers|
|[**DatacentersLoadbalancersPatch**](LoadBalancersApi.md#DatacentersLoadbalancersPatch) | **Patch** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Partially modify a Loadbalancer|
|[**DatacentersLoadbalancersPost**](LoadBalancersApi.md#DatacentersLoadbalancersPost) | **Post** /datacenters/{datacenterId}/loadbalancers | Create a Load Balancer|
|[**DatacentersLoadbalancersPut**](LoadBalancersApi.md#DatacentersLoadbalancersPut) | **Put** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Modify a Load Balancer|



## DatacentersLoadbalancersBalancednicsDelete

```go
var result  = DatacentersLoadbalancersBalancednicsDelete(ctx, datacenterId, loadbalancerId, nicId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Detach a nic from loadbalancer



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
    loadbalancerId := "loadbalancerId_example" // string | The unique ID of the Load Balancer
    nicId := "nicId_example" // string | The unique ID of the NIC
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersBalancednicsDelete(context.Background(), datacenterId, loadbalancerId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersBalancednicsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**loadbalancerId** | **string** | The unique ID of the Load Balancer | |
|**nicId** | **string** | The unique ID of the NIC | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersBalancednicsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersLoadbalancersBalancednicsFindByNicId

```go
var result Nic = DatacentersLoadbalancersBalancednicsFindByNicId(ctx, datacenterId, loadbalancerId, nicId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a network interface attached to Load Balancer



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
    loadbalancerId := "loadbalancerId_example" // string | The unique ID of the Load Balancer
    nicId := "nicId_example" // string | The unique ID of the NIC
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersBalancednicsFindByNicId(context.Background(), datacenterId, loadbalancerId, nicId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersBalancednicsFindByNicId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLoadbalancersBalancednicsFindByNicId`: Nic
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DatacentersLoadbalancersBalancednicsFindByNicId`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**loadbalancerId** | **string** | The unique ID of the Load Balancer | |
|**nicId** | **string** | The unique ID of the NIC | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersBalancednicsFindByNicIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Nic**](Nic.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersLoadbalancersBalancednicsGet

```go
var result BalancedNics = DatacentersLoadbalancersBalancednicsGet(ctx, datacenterId, loadbalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List Load Balancer balaced NICs



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
    loadbalancerId := "loadbalancerId_example" // string | The unique ID of the Load Balancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersBalancednicsGet(context.Background(), datacenterId, loadbalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersBalancednicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLoadbalancersBalancednicsGet`: BalancedNics
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DatacentersLoadbalancersBalancednicsGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**loadbalancerId** | **string** | The unique ID of the Load Balancer | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersBalancednicsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**BalancedNics**](BalancedNics.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersLoadbalancersBalancednicsPost

```go
var result Nic = DatacentersLoadbalancersBalancednicsPost(ctx, datacenterId, loadbalancerId)
                      .Nic(nic)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Attach a nic to Load Balancer



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
    loadbalancerId := "loadbalancerId_example" // string | The unique ID of the Load Balancer
    nic := *openapiclient.NewNic(*openapiclient.NewNicProperties(int32(2))) // Nic | Nic id to be attached
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersBalancednicsPost(context.Background(), datacenterId, loadbalancerId).Nic(nic).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersBalancednicsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLoadbalancersBalancednicsPost`: Nic
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DatacentersLoadbalancersBalancednicsPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**loadbalancerId** | **string** | The unique ID of the Load Balancer | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersBalancednicsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **nic** | [**Nic**](Nic.md) | Nic id to be attached | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Nic**](Nic.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersLoadbalancersDelete

```go
var result  = DatacentersLoadbalancersDelete(ctx, datacenterId, loadbalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete a Loadbalancer.



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
    loadbalancerId := "loadbalancerId_example" // string | The unique ID of the Load Balancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersDelete(context.Background(), datacenterId, loadbalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**loadbalancerId** | **string** | The unique ID of the Load Balancer | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersLoadbalancersFindById

```go
var result Loadbalancer = DatacentersLoadbalancersFindById(ctx, datacenterId, loadbalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve a loadbalancer



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
    loadbalancerId := "loadbalancerId_example" // string | The unique ID of the Load Balancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersFindById(context.Background(), datacenterId, loadbalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLoadbalancersFindById`: Loadbalancer
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DatacentersLoadbalancersFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**loadbalancerId** | **string** | The unique ID of the Load Balancer | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Loadbalancer**](Loadbalancer.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersLoadbalancersGet

```go
var result Loadbalancers = DatacentersLoadbalancersGet(ctx, datacenterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List Load Balancers



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
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLoadbalancersGet`: Loadbalancers
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DatacentersLoadbalancersGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |
| **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with limit for pagination) | [default to 0]|
| **limit** | **int32** | the maximum number of elements to return (use together with offset for pagination) | [default to 1000]|

### Return type

[**Loadbalancers**](Loadbalancers.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersLoadbalancersPatch

```go
var result Loadbalancer = DatacentersLoadbalancersPatch(ctx, datacenterId, loadbalancerId)
                      .Loadbalancer(loadbalancer)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify a Loadbalancer



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
    loadbalancerId := "loadbalancerId_example" // string | The unique ID of the Load Balancer
    loadbalancer := *openapiclient.NewLoadbalancerProperties() // LoadbalancerProperties | Modified Loadbalancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersPatch(context.Background(), datacenterId, loadbalancerId).Loadbalancer(loadbalancer).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLoadbalancersPatch`: Loadbalancer
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DatacentersLoadbalancersPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**loadbalancerId** | **string** | The unique ID of the Load Balancer | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **loadbalancer** | [**LoadbalancerProperties**](LoadbalancerProperties.md) | Modified Loadbalancer | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Loadbalancer**](Loadbalancer.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersLoadbalancersPost

```go
var result Loadbalancer = DatacentersLoadbalancersPost(ctx, datacenterId)
                      .Loadbalancer(loadbalancer)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Load Balancer



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
    loadbalancer := *openapiclient.NewLoadbalancer(*openapiclient.NewLoadbalancerProperties()) // Loadbalancer | Loadbalancer to be created
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersPost(context.Background(), datacenterId).Loadbalancer(loadbalancer).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLoadbalancersPost`: Loadbalancer
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DatacentersLoadbalancersPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **loadbalancer** | [**Loadbalancer**](Loadbalancer.md) | Loadbalancer to be created | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Loadbalancer**](Loadbalancer.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersLoadbalancersPut

```go
var result Loadbalancer = DatacentersLoadbalancersPut(ctx, datacenterId, loadbalancerId)
                      .Loadbalancer(loadbalancer)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Load Balancer



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
    loadbalancerId := "loadbalancerId_example" // string | The unique ID of the Load Balancer
    loadbalancer := *openapiclient.NewLoadbalancer(*openapiclient.NewLoadbalancerProperties()) // Loadbalancer | Modified Loadbalancer
    pretty := true // bool | Controls whether response is pretty-printed (with indentation and new lines) (optional) (default to true)
    depth := int32(56) // int32 | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth=0: only direct properties are included. Children (servers etc.) are not included  - depth=1: direct properties and children references are included  - depth=2: direct properties and children properties are included  - depth=3: direct properties and children properties and children's children are included  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users having more than 1 contract need to provide contract number, against which all API requests should be executed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoadBalancersApi.DatacentersLoadbalancersPut(context.Background(), datacenterId, loadbalancerId).Loadbalancer(loadbalancer).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancersApi.DatacentersLoadbalancersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersLoadbalancersPut`: Loadbalancer
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancersApi.DatacentersLoadbalancersPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the datacenter | |
|**loadbalancerId** | **string** | The unique ID of the Load Balancer | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersLoadbalancersPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **loadbalancer** | [**Loadbalancer**](Loadbalancer.md) | Modified Loadbalancer | |
| **pretty** | **bool** | Controls whether response is pretty-printed (with indentation and new lines) | [default to true]|
| **depth** | **int32** | Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users having more than 1 contract need to provide contract number, against which all API requests should be executed | |

### Return type

[**Loadbalancer**](Loadbalancer.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

