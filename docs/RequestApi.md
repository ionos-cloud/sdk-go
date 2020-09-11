# \RequestApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestsFindById**](RequestApi.md#RequestsFindById) | **Get** /requests/{requestId} | Retrieve a Request
[**RequestsGet**](RequestApi.md#RequestsGet) | **Get** /requests | List Requests
[**RequestsStatusGet**](RequestApi.md#RequestsStatusGet) | **Get** /requests/{requestId}/status | Retrieve Request Status



## RequestsFindById

> Request RequestsFindById(ctx, requestId, optional)

Retrieve a Request

Retrieves the attributes of a given request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string**|  | 
 **optional** | ***RequestsFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RequestsFindByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Request**](Request.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsGet

> Requests RequestsGet(ctx, optional)

List Requests

Retrieve a list of API requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RequestsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RequestsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 
 **filterStatus** | **optional.String**| Request status filter to fetch all the request based on a particular status [QUEUED, RUNNING, DONE, FAILED] | 
 **filterCreatedAfter** | **optional.String**| Filter all the requests after the created date | 
 **filterCreatedBefore** | **optional.String**| Filter all the requests before the created date | 
 **filterUrl** | **optional.String**| Url filter | 
 **filterCreatedDate** | **optional.String**| Created date filter | 
 **filterMethod** | **optional.String**| Method filter | 
 **filterBody** | **optional.String**| Url filter | 

### Return type

[**Requests**](Requests.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestsStatusGet

> RequestStatus RequestsStatusGet(ctx, requestId, optional)

Retrieve Request Status

Retrieves the status of a given request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string**|  | 
 **optional** | ***RequestsStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RequestsStatusGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**RequestStatus**](RequestStatus.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

