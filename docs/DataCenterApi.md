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

> map[string]interface{} DatacentersDelete(ctx, datacenterId, optional)

Delete a Data Center

Will remove all objects within the datacenter and remove the datacenter object itself, too. This is a highly destructive method which should be used with caution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***DatacentersDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersDeleteOpts struct


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


## DatacentersFindById

> Datacenter DatacentersFindById(ctx, datacenterId, optional)

Retrieve a Data Center

You can retrieve a data center by using the resource's ID. This value can be found in the response body when a datacenter is created or when you GET a list of datacenters.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***DatacentersFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersFindByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

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

> Datacenters DatacentersGet(ctx, optional)

List Data Centers under your account

You can retrieve a complete list of data centers provisioned under your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DatacentersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

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

> Datacenter DatacentersPatch(ctx, datacenterId, datacenter, optional)

Partially modify a Data Center

You can use update datacenter to re-name the datacenter or update its description

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**datacenter** | [**DatacenterProperties**](DatacenterProperties.md)| Modified properties of Data Center | 
 **optional** | ***DatacentersPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

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

> Datacenter DatacentersPost(ctx, datacenter, optional)

Create a Data Center

Virtual data centers are the foundation of the platform. They act as logical containers for all other objects you will be creating, e.g. servers. You can provision as many data centers as you want. Datacenters have their own private network and are logically segmented from each other to create isolation. You can use this POST method to create a simple datacenter or to create a datacenter with multiple objects under it such as servers and storage volumes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenter** | [**Datacenter**](Datacenter.md)| Datacenter to be created | 
 **optional** | ***DatacentersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

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

> Datacenter DatacentersPut(ctx, datacenterId, datacenter, optional)

Modify a Data Center

You can use update datacenter to re-name the datacenter or update its description

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**datacenter** | [**Datacenter**](Datacenter.md)| Modified Data Center | 
 **optional** | ***DatacentersPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

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

