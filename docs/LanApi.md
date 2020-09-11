# \LanApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersLansDelete**](LanApi.md#DatacentersLansDelete) | **Delete** /datacenters/{datacenterId}/lans/{lanId} | Delete a Lan.
[**DatacentersLansFindById**](LanApi.md#DatacentersLansFindById) | **Get** /datacenters/{datacenterId}/lans/{lanId} | Retrieve a Lan
[**DatacentersLansGet**](LanApi.md#DatacentersLansGet) | **Get** /datacenters/{datacenterId}/lans | List Lans
[**DatacentersLansNicsFindById**](LanApi.md#DatacentersLansNicsFindById) | **Get** /datacenters/{datacenterId}/lans/{lanId}/nics/{nicId} | Retrieve a nic attached to lan
[**DatacentersLansNicsGet**](LanApi.md#DatacentersLansNicsGet) | **Get** /datacenters/{datacenterId}/lans/{lanId}/nics | List Lan Members 
[**DatacentersLansNicsPost**](LanApi.md#DatacentersLansNicsPost) | **Post** /datacenters/{datacenterId}/lans/{lanId}/nics | Attach a nic
[**DatacentersLansPatch**](LanApi.md#DatacentersLansPatch) | **Patch** /datacenters/{datacenterId}/lans/{lanId} | Partially modify a Lan
[**DatacentersLansPost**](LanApi.md#DatacentersLansPost) | **Post** /datacenters/{datacenterId}/lans | Create a Lan
[**DatacentersLansPut**](LanApi.md#DatacentersLansPut) | **Put** /datacenters/{datacenterId}/lans/{lanId} | Modify a Lan



## DatacentersLansDelete

> map[string]interface{} DatacentersLansDelete(ctx, datacenterId, lanId, optional)

Delete a Lan.

Removes the specific Lan

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**lanId** | **string**| The unique ID of the LAN | 
 **optional** | ***DatacentersLansDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLansDeleteOpts struct


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


## DatacentersLansFindById

> Lan DatacentersLansFindById(ctx, datacenterId, lanId, optional)

Retrieve a Lan

Retrieves the attributes of a given LAN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**lanId** | **string**| The unique ID of the LAN | 
 **optional** | ***DatacentersLansFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLansFindByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Lan**](Lan.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLansGet

> Lans DatacentersLansGet(ctx, datacenterId, optional)

List Lans

Retrieve a list of LANs within the datacenter

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***DatacentersLansGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLansGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Lans**](Lans.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLansNicsFindById

> Nic DatacentersLansNicsFindById(ctx, datacenterId, lanId, nicId, optional)

Retrieve a nic attached to lan

This will retrieve the properties of an attached nic.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**lanId** | **string**| The unique ID of the LAN | 
**nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***DatacentersLansNicsFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLansNicsFindByIdOpts struct


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


## DatacentersLansNicsGet

> LanNics DatacentersLansNicsGet(ctx, datacenterId, lanId, optional)

List Lan Members 

You can retrieve a list of nics attached to a lan

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**lanId** | **string**| The unique ID of the LAN | 
 **optional** | ***DatacentersLansNicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLansNicsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LanNics**](LanNics.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLansNicsPost

> Nic DatacentersLansNicsPost(ctx, datacenterId, lanId, nic, optional)

Attach a nic

This will attach a pre-existing nic to a lan. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**lanId** | **string**| The unique ID of the LAN | 
**nic** | [**Nic**](Nic.md)| Nic to be attached | 
 **optional** | ***DatacentersLansNicsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLansNicsPostOpts struct


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


## DatacentersLansPatch

> Lan DatacentersLansPatch(ctx, datacenterId, lanId, lan, optional)

Partially modify a Lan

You can use update attributes of a resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**lanId** | **string**| The unique ID of the LAN | 
**lan** | [**LanProperties**](LanProperties.md)| Modified Lan | 
 **optional** | ***DatacentersLansPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLansPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Lan**](Lan.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLansPost

> LanPost DatacentersLansPost(ctx, datacenterId, lan, optional)

Create a Lan

Creates a LAN within the datacenter

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**lan** | [**LanPost**](LanPost.md)| Lan to be created | 
 **optional** | ***DatacentersLansPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLansPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**LanPost**](LanPost.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersLansPut

> Lan DatacentersLansPut(ctx, datacenterId, lanId, lan, optional)

Modify a Lan

You can use update attributes of a resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**lanId** | **string**| The unique ID of the LAN | 
**lan** | [**Lan**](Lan.md)| Modified Lan | 
 **optional** | ***DatacentersLansPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersLansPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Lan**](Lan.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

