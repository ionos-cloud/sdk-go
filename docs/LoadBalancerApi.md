# \LoadBalancerApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersLoadbalancersBalancednicsDelete**](LoadBalancerApi.md#DatacentersLoadbalancersBalancednicsDelete) | **Delete** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics/{nicId} | Detach a nic from loadbalancer
[**DatacentersLoadbalancersBalancednicsFindByNic**](LoadBalancerApi.md#DatacentersLoadbalancersBalancednicsFindByNic) | **Get** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics/{nicId} | Retrieve a nic attached to Load Balancer
[**DatacentersLoadbalancersBalancednicsGet**](LoadBalancerApi.md#DatacentersLoadbalancersBalancednicsGet) | **Get** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics | List Load Balancer Members 
[**DatacentersLoadbalancersBalancednicsPost**](LoadBalancerApi.md#DatacentersLoadbalancersBalancednicsPost) | **Post** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId}/balancednics | Attach a nic to Load Balancer
[**DatacentersLoadbalancersDelete**](LoadBalancerApi.md#DatacentersLoadbalancersDelete) | **Delete** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Delete a Loadbalancer.
[**DatacentersLoadbalancersFindById**](LoadBalancerApi.md#DatacentersLoadbalancersFindById) | **Get** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Retrieve a loadbalancer
[**DatacentersLoadbalancersGet**](LoadBalancerApi.md#DatacentersLoadbalancersGet) | **Get** /datacenters/{datacenterId}/loadbalancers | List Load Balancers
[**DatacentersLoadbalancersPatch**](LoadBalancerApi.md#DatacentersLoadbalancersPatch) | **Patch** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Partially modify a Loadbalancer
[**DatacentersLoadbalancersPost**](LoadBalancerApi.md#DatacentersLoadbalancersPost) | **Post** /datacenters/{datacenterId}/loadbalancers | Create a Load Balancer
[**DatacentersLoadbalancersPut**](LoadBalancerApi.md#DatacentersLoadbalancersPut) | **Put** /datacenters/{datacenterId}/loadbalancers/{loadbalancerId} | Modify a Load Balancer



## DatacentersLoadbalancersBalancednicsDelete

> map[string]interface{} DatacentersLoadbalancersBalancednicsDelete(ctx, datacenterId, loadbalancerId, nicId, optional)

Detach a nic from loadbalancer

This will remove a nic from Load Balancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**loadbalancerId** | **string**| The unique ID of the Load Balancer | 
**nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***DatacentersLoadbalancersBalancednicsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersBalancednicsDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

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


## DatacentersLoadbalancersBalancednicsFindByNic

> Nic DatacentersLoadbalancersBalancednicsFindByNic(ctx, datacenterId, loadbalancerId, nicId, optional)

Retrieve a nic attached to Load Balancer

This will retrieve the properties of an attached nic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**loadbalancerId** | **string**| The unique ID of the Load Balancer | 
**nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***DatacentersLoadbalancersBalancednicsFindByNicOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersBalancednicsFindByNicOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nic**](Nic.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLoadbalancersBalancednicsGet

> BalancedNics DatacentersLoadbalancersBalancednicsGet(ctx, datacenterId, loadbalancerId, optional)

List Load Balancer Members 

You can retrieve a list of nics attached to a Load Balancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**loadbalancerId** | **string**| The unique ID of the Load Balancer | 
 **optional** | ***DatacentersLoadbalancersBalancednicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersBalancednicsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**BalancedNics**](BalancedNics.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLoadbalancersBalancednicsPost

> Nic DatacentersLoadbalancersBalancednicsPost(ctx, datacenterId, loadbalancerId, nic, optional)

Attach a nic to Load Balancer

This will attach a pre-existing nic to a Load Balancer. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**loadbalancerId** | **string**| The unique ID of the Load Balancer | 
**nic** | [**Nic**](Nic.md)| Nic id to be attached | 
 **optional** | ***DatacentersLoadbalancersBalancednicsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersBalancednicsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nic**](Nic.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLoadbalancersDelete

> map[string]interface{} DatacentersLoadbalancersDelete(ctx, datacenterId, loadbalancerId, optional)

Delete a Loadbalancer.

Removes the specific Loadbalancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**loadbalancerId** | **string**| The unique ID of the Load Balancer | 
 **optional** | ***DatacentersLoadbalancersDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

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


## DatacentersLoadbalancersFindById

> Loadbalancer DatacentersLoadbalancersFindById(ctx, datacenterId, loadbalancerId, optional)

Retrieve a loadbalancer

Retrieves the attributes of a given Loadbalancer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**loadbalancerId** | **string**| The unique ID of the Load Balancer | 
 **optional** | ***DatacentersLoadbalancersFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersFindByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancer**](Loadbalancer.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLoadbalancersGet

> Loadbalancers DatacentersLoadbalancersGet(ctx, datacenterId, optional)

List Load Balancers

Retrieve a list of Load Balancers within the datacenter

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***DatacentersLoadbalancersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancers**](Loadbalancers.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLoadbalancersPatch

> Loadbalancer DatacentersLoadbalancersPatch(ctx, datacenterId, loadbalancerId, loadbalancer, optional)

Partially modify a Loadbalancer

You can use update attributes of a resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**loadbalancerId** | **string**| The unique ID of the Load Balancer | 
**loadbalancer** | [**LoadbalancerProperties**](LoadbalancerProperties.md)| Modified Loadbalancer | 
 **optional** | ***DatacentersLoadbalancersPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancer**](Loadbalancer.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLoadbalancersPost

> Loadbalancer DatacentersLoadbalancersPost(ctx, datacenterId, loadbalancer, optional)

Create a Load Balancer

Creates a Loadbalancer within the datacenter

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**loadbalancer** | [**Loadbalancer**](Loadbalancer.md)| Loadbalancer to be created | 
 **optional** | ***DatacentersLoadbalancersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancer**](Loadbalancer.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLoadbalancersPut

> Loadbalancer DatacentersLoadbalancersPut(ctx, datacenterId, loadbalancerId, loadbalancer, optional)

Modify a Load Balancer

You can use update attributes of a resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**loadbalancerId** | **string**| The unique ID of the Load Balancer | 
**loadbalancer** | [**Loadbalancer**](Loadbalancer.md)| Modified Loadbalancer | 
 **optional** | ***DatacentersLoadbalancersPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLoadbalancersPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Loadbalancer**](Loadbalancer.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

