# \ContractApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContractsGet**](ContractApi.md#ContractsGet) | **Get** /contracts | Retrieve a Contract



## ContractsGet

> Contract ContractsGet(ctx, optional)

Retrieve a Contract

Retrieves the attributes of user's contract.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContractsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContractsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Contract**](Contract.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

