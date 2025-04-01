# \SecurityGroupsApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersSecuritygroupsDelete**](SecurityGroupsApi.md#DatacentersSecuritygroupsDelete) | **Delete** /datacenters/{datacenterId}/securitygroups/{securityGroupId} | Delete a Security Group|
|[**DatacentersSecuritygroupsFindById**](SecurityGroupsApi.md#DatacentersSecuritygroupsFindById) | **Get** /datacenters/{datacenterId}/securitygroups/{securityGroupId} | Retrieve a Security Group|
|[**DatacentersSecuritygroupsFirewallrulesDelete**](SecurityGroupsApi.md#DatacentersSecuritygroupsFirewallrulesDelete) | **Delete** /datacenters/{datacenterId}/securitygroups/{securityGroupId}/rules/{ruleId} | Remove a Firewall Rule from a Security Group|
|[**DatacentersSecuritygroupsFirewallrulesPost**](SecurityGroupsApi.md#DatacentersSecuritygroupsFirewallrulesPost) | **Post** /datacenters/{datacenterId}/securitygroups/{securityGroupId}/rules | Create Firewall rule to a Security Group|
|[**DatacentersSecuritygroupsGet**](SecurityGroupsApi.md#DatacentersSecuritygroupsGet) | **Get** /datacenters/{datacenterId}/securitygroups | List Security Groups|
|[**DatacentersSecuritygroupsPatch**](SecurityGroupsApi.md#DatacentersSecuritygroupsPatch) | **Patch** /datacenters/{datacenterId}/securitygroups/{securityGroupId} | Partially modify Security Group|
|[**DatacentersSecuritygroupsPost**](SecurityGroupsApi.md#DatacentersSecuritygroupsPost) | **Post** /datacenters/{datacenterId}/securitygroups | Create a Security Group|
|[**DatacentersSecuritygroupsPut**](SecurityGroupsApi.md#DatacentersSecuritygroupsPut) | **Put** /datacenters/{datacenterId}/securitygroups/{securityGroupId} | Modify Security Group|
|[**DatacentersSecuritygroupsRulesFindById**](SecurityGroupsApi.md#DatacentersSecuritygroupsRulesFindById) | **Get** /datacenters/{datacenterId}/securitygroups/{securityGroupId}/rules/{ruleId} | Retrieve security group rule by id|
|[**DatacentersSecuritygroupsRulesGet**](SecurityGroupsApi.md#DatacentersSecuritygroupsRulesGet) | **Get** /datacenters/{datacenterId}/securitygroups/{securityGroupId}/rules | List Security Group rules|
|[**DatacentersSecuritygroupsRulesPatch**](SecurityGroupsApi.md#DatacentersSecuritygroupsRulesPatch) | **Patch** /datacenters/{datacenterId}/securitygroups/{securityGroupId}/rules/{ruleId} | Partially modify Security Group Rules|
|[**DatacentersSecuritygroupsRulesPut**](SecurityGroupsApi.md#DatacentersSecuritygroupsRulesPut) | **Put** /datacenters/{datacenterId}/securitygroups/{securityGroupId}/rules/{ruleId} | Modify a Security Group Rule|
|[**DatacentersServersNicsSecuritygroupsPut**](SecurityGroupsApi.md#DatacentersServersNicsSecuritygroupsPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/securitygroups | Attach a list of Security Groups to a NIC|
|[**DatacentersServersSecuritygroupsPut**](SecurityGroupsApi.md#DatacentersServersSecuritygroupsPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/securitygroups | Attach a list of Security Groups to a Server|



## DatacentersSecuritygroupsDelete

```go
var result  = DatacentersSecuritygroupsDelete(ctx, datacenterId, securityGroupId)
                      .Pretty(pretty)
                      .Execute()
```

Delete a Security Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    securityGroupId := "securityGroupId_example" // string | The unique ID of the Security Group.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsDelete(context.Background(), datacenterId, securityGroupId).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**securityGroupId** | **string** | The unique ID of the Security Group. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersSecuritygroupsFindById

```go
var result SecurityGroup = DatacentersSecuritygroupsFindById(ctx, datacenterId, securityGroupId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Retrieve a Security Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center
    securityGroupId := "securityGroupId_example" // string | The unique ID of the security group.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsFindById(context.Background(), datacenterId, securityGroupId).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsFindById`: SecurityGroup
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center | |
|**securityGroupId** | **string** | The unique ID of the security group. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|

### Return type

[**SecurityGroup**](../models/SecurityGroup.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersSecuritygroupsFirewallrulesDelete

```go
var result  = DatacentersSecuritygroupsFirewallrulesDelete(ctx, datacenterId, securityGroupId, ruleId)
                      .Execute()
```

Remove a Firewall Rule from a Security Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center
    securityGroupId := "securityGroupId_example" // string | The unique ID of the security group.
    ruleId := "ruleId_example" // string | The unique ID of the firewall rule.

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsFirewallrulesDelete(context.Background(), datacenterId, securityGroupId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsFirewallrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center | |
|**securityGroupId** | **string** | The unique ID of the security group. | |
|**ruleId** | **string** | The unique ID of the firewall rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsFirewallrulesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersSecuritygroupsFirewallrulesPost

```go
var result FirewallRule = DatacentersSecuritygroupsFirewallrulesPost(ctx, datacenterId, securityGroupId)
                      .FirewallRule(firewallRule)
                      .Execute()
```

Create Firewall rule to a Security Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center
    securityGroupId := "securityGroupId_example" // string | The unique ID of the security group.
    firewallRule := *openapiclient.NewFirewallRule(*openapiclient.NewFirewallruleProperties()) // FirewallRule | The firewall to be attached (or created and attached).

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsFirewallrulesPost(context.Background(), datacenterId, securityGroupId).FirewallRule(firewallRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsFirewallrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsFirewallrulesPost`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsFirewallrulesPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center | |
|**securityGroupId** | **string** | The unique ID of the security group. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsFirewallrulesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **firewallRule** | [**FirewallRule**](../models/FirewallRule.md) | The firewall to be attached (or created and attached). | |

### Return type

[**FirewallRule**](../models/FirewallRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersSecuritygroupsGet

```go
var result SecurityGroups = DatacentersSecuritygroupsGet(ctx, datacenterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List Security Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 1000)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsGet`: SecurityGroups
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 1000]|

### Return type

[**SecurityGroups**](../models/SecurityGroups.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersSecuritygroupsPatch

```go
var result SecurityGroup = DatacentersSecuritygroupsPatch(ctx, datacenterId, securityGroupId)
                      .SecurityGroup(securityGroup)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Partially modify Security Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    securityGroupId := "securityGroupId_example" // string | The unique ID of the Security Group.
    securityGroup := *openapiclient.NewSecurityGroupProperties("My security group") // SecurityGroupProperties | The modified Security Group
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsPatch(context.Background(), datacenterId, securityGroupId).SecurityGroup(securityGroup).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsPatch`: SecurityGroup
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**securityGroupId** | **string** | The unique ID of the Security Group. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **securityGroup** | [**SecurityGroupProperties**](../models/SecurityGroupProperties.md) | The modified Security Group | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|

### Return type

[**SecurityGroup**](../models/SecurityGroup.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersSecuritygroupsPost

```go
var result SecurityGroup = DatacentersSecuritygroupsPost(ctx, datacenterId)
                      .SecurityGroup(securityGroup)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Create a Security Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    securityGroup := *openapiclient.NewSecurityGroupRequest(*openapiclient.NewSecurityGroupProperties("My security group")) // SecurityGroupRequest | The security group to be created
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsPost(context.Background(), datacenterId).SecurityGroup(securityGroup).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsPost`: SecurityGroup
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsPost`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **securityGroup** | [**SecurityGroupRequest**](../models/SecurityGroupRequest.md) | The security group to be created | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|

### Return type

[**SecurityGroup**](../models/SecurityGroup.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersSecuritygroupsPut

```go
var result SecurityGroup = DatacentersSecuritygroupsPut(ctx, datacenterId, securityGroupId)
                      .SecurityGroup(securityGroup)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Modify Security Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    securityGroupId := "securityGroupId_example" // string | The unique ID of the Security Group.
    securityGroup := *openapiclient.NewSecurityGroupRequest(*openapiclient.NewSecurityGroupProperties("My security group")) // SecurityGroupRequest | The modified Security Group
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsPut(context.Background(), datacenterId, securityGroupId).SecurityGroup(securityGroup).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsPut`: SecurityGroup
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**securityGroupId** | **string** | The unique ID of the Security Group. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **securityGroup** | [**SecurityGroupRequest**](../models/SecurityGroupRequest.md) | The modified Security Group | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|

### Return type

[**SecurityGroup**](../models/SecurityGroup.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersSecuritygroupsRulesFindById

```go
var result FirewallRule = DatacentersSecuritygroupsRulesFindById(ctx, datacenterId, securityGroupId, ruleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Retrieve security group rule by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    securityGroupId := "securityGroupId_example" // string | The unique ID of the Security Group.
    ruleId := "ruleId_example" // string | The unique ID of the Security Group rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsRulesFindById(context.Background(), datacenterId, securityGroupId, ruleId).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsRulesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsRulesFindById`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsRulesFindById`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**securityGroupId** | **string** | The unique ID of the Security Group. | |
|**ruleId** | **string** | The unique ID of the Security Group rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsRulesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FirewallRule**](../models/FirewallRule.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersSecuritygroupsRulesGet

```go
var result FirewallRules = DatacentersSecuritygroupsRulesGet(ctx, datacenterId, securityGroupId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List Security Group rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    securityGroupId := "securityGroupId_example" // string | The unique ID of the security group.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 1000)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsRulesGet(context.Background(), datacenterId, securityGroupId).Pretty(pretty).Depth(depth).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsRulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsRulesGet`: FirewallRules
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsRulesGet`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**securityGroupId** | **string** | The unique ID of the security group. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsRulesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 1000]|

### Return type

[**FirewallRules**](../models/FirewallRules.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersSecuritygroupsRulesPatch

```go
var result FirewallRule = DatacentersSecuritygroupsRulesPatch(ctx, datacenterId, securityGroupId, ruleId)
                      .Rule(rule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Partially modify Security Group Rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    securityGroupId := "securityGroupId_example" // string | The unique ID of the security group.
    ruleId := "ruleId_example" // string | The unique ID of the Security Group Rule.
    rule := *openapiclient.NewFirewallruleProperties() // FirewallruleProperties | The properties of the Security Group Rule to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsRulesPatch(context.Background(), datacenterId, securityGroupId, ruleId).Rule(rule).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsRulesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsRulesPatch`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsRulesPatch`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**securityGroupId** | **string** | The unique ID of the security group. | |
|**ruleId** | **string** | The unique ID of the Security Group Rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsRulesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rule** | [**FirewallruleProperties**](../models/FirewallruleProperties.md) | The properties of the Security Group Rule to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FirewallRule**](../models/FirewallRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersSecuritygroupsRulesPut

```go
var result FirewallRule = DatacentersSecuritygroupsRulesPut(ctx, datacenterId, securityGroupId, ruleId)
                      .Rule(rule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Modify a Security Group Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    securityGroupId := "securityGroupId_example" // string | The unique ID of the security group.
    ruleId := "ruleId_example" // string | The unique ID of the Security Group Rule.
    rule := *openapiclient.NewFirewallRule(*openapiclient.NewFirewallruleProperties()) // FirewallRule | The modified Security Group Rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersSecuritygroupsRulesPut(context.Background(), datacenterId, securityGroupId, ruleId).Rule(rule).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersSecuritygroupsRulesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersSecuritygroupsRulesPut`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersSecuritygroupsRulesPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**securityGroupId** | **string** | The unique ID of the security group. | |
|**ruleId** | **string** | The unique ID of the Security Group Rule. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersSecuritygroupsRulesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rule** | [**FirewallRule**](../models/FirewallRule.md) | The modified Security Group Rule. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|

### Return type

[**FirewallRule**](../models/FirewallRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersNicsSecuritygroupsPut

```go
var result SecurityGroups = DatacentersServersNicsSecuritygroupsPut(ctx, datacenterId, serverId, nicId)
                      .Securitygroups(securitygroups)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Attach a list of Security Groups to a NIC



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    nicId := "nicId_example" // string | The unique ID of the server.
    securitygroups := *openapiclient.NewListOfIds([]string{"15f67991-0f51-4efc-a8ad-ef1fb31a480c"}) // ListOfIds | The list of NIC attached Security Groups IDs.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersServersNicsSecuritygroupsPut(context.Background(), datacenterId, serverId, nicId).Securitygroups(securitygroups).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersServersNicsSecuritygroupsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersNicsSecuritygroupsPut`: SecurityGroups
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersServersNicsSecuritygroupsPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**nicId** | **string** | The unique ID of the server. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersNicsSecuritygroupsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **securitygroups** | [**ListOfIds**](../models/ListOfIds.md) | The list of NIC attached Security Groups IDs. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|

### Return type

[**SecurityGroups**](../models/SecurityGroups.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersSecuritygroupsPut

```go
var result SecurityGroups = DatacentersServersSecuritygroupsPut(ctx, datacenterId, serverId)
                      .Securitygroups(securitygroups)
                      .Pretty(pretty)
                      .Depth(depth)
                      .Execute()
```

Attach a list of Security Groups to a Server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go/v6"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    securitygroups := *openapiclient.NewListOfIds([]string{"15f67991-0f51-4efc-a8ad-ef1fb31a480c"}) // ListOfIds | The list of server attached Security Groups IDs.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth=0: Only direct properties are included;              children (servers and other elements) are not included.   - depth=1: Direct properties and children references are included.   - depth=2: Direct properties and children properties are included.   - depth=3: Direct properties and children properties and children's children are included.   - depth=... and so on (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration()
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecurityGroupsApi.DatacentersServersSecuritygroupsPut(context.Background(), datacenterId, serverId).Securitygroups(securitygroups).Pretty(pretty).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupsApi.DatacentersServersSecuritygroupsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatacentersServersSecuritygroupsPut`: SecurityGroups
    fmt.Fprintf(os.Stdout, "Response from `SecurityGroupsApi.DatacentersServersSecuritygroupsPut`: %v\n", resp)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |

### Other Parameters

Other parameters are passed through a pointer to a apiDatacentersServersSecuritygroupsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **securitygroups** | [**ListOfIds**](../models/ListOfIds.md) | The list of server attached Security Groups IDs. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects. GET /datacenters/[ID]   - depth&#x3D;0: Only direct properties are included;              children (servers and other elements) are not included.   - depth&#x3D;1: Direct properties and children references are included.   - depth&#x3D;2: Direct properties and children properties are included.   - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.   - depth&#x3D;... and so on | [default to 0]|

### Return type

[**SecurityGroups**](../models/SecurityGroups.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


