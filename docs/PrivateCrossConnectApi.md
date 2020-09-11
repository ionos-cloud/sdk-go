# \PrivateCrossConnectApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PccsDelete**](PrivateCrossConnectApi.md#PccsDelete) | **Delete** /pccs/{pccId} | Delete a Private Cross-Connect
[**PccsFindById**](PrivateCrossConnectApi.md#PccsFindById) | **Get** /pccs/{pccId} | Retrieve a Private Cross-Connect
[**PccsGet**](PrivateCrossConnectApi.md#PccsGet) | **Get** /pccs | List Private Cross-Connects 
[**PccsPatch**](PrivateCrossConnectApi.md#PccsPatch) | **Patch** /pccs/{pccId} | Partially modify a private cross-connect
[**PccsPost**](PrivateCrossConnectApi.md#PccsPost) | **Post** /pccs | Create a Private Cross-Connect



## PccsDelete

> map[string]interface{} PccsDelete(ctx, pccId, optional)

Delete a Private Cross-Connect

Delete a private cross-connect if no datacenters are joined to the given PCC

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pccId** | **string**| The unique ID of the private cross-connect | 
 **optional** | ***PccsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PccsDeleteOpts struct


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


## PccsFindById

> PrivateCrossConnect PccsFindById(ctx, pccId, optional)

Retrieve a Private Cross-Connect

You can retrieve a private cross-connect by using the resource's ID. This value can be found in the response body when a private cross-connect is created or when you GET a list of private cross-connects.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pccId** | **string**| The unique ID of the private cross-connect | 
 **optional** | ***PccsFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PccsFindByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**PrivateCrossConnect**](PrivateCrossConnect.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PccsGet

> PrivateCrossConnects PccsGet(ctx, optional)

List Private Cross-Connects 

You can retrieve a complete list of private cross-connects provisioned under your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PccsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PccsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**PrivateCrossConnects**](PrivateCrossConnects.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PccsPatch

> PrivateCrossConnect PccsPatch(ctx, pccId, pcc, optional)

Partially modify a private cross-connect

You can use update private cross-connect to re-name or update its description

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pccId** | **string**| The unique ID of the private cross-connect | 
**pcc** | [**PrivateCrossConnectProperties**](PrivateCrossConnectProperties.md)| Modified properties of private cross-connect | 
 **optional** | ***PccsPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PccsPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**PrivateCrossConnect**](PrivateCrossConnect.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PccsPost

> PrivateCrossConnect PccsPost(ctx, pcc, optional)

Create a Private Cross-Connect

You can use this POST method to create a private cross-connect

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pcc** | [**PrivateCrossConnect**](PrivateCrossConnect.md)| Private Cross-Connect to be created | 
 **optional** | ***PccsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PccsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**PrivateCrossConnect**](PrivateCrossConnect.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

