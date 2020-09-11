# \NicApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersServersNicsDelete**](NicApi.md#DatacentersServersNicsDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Delete a Nic
[**DatacentersServersNicsFindById**](NicApi.md#DatacentersServersNicsFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Retrieve a Nic
[**DatacentersServersNicsFirewallrulesDelete**](NicApi.md#DatacentersServersNicsFirewallrulesDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Delete a Firewall Rule
[**DatacentersServersNicsFirewallrulesFindById**](NicApi.md#DatacentersServersNicsFirewallrulesFindById) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Retrieve a Firewall Rule
[**DatacentersServersNicsFirewallrulesGet**](NicApi.md#DatacentersServersNicsFirewallrulesGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | List Firewall Rules 
[**DatacentersServersNicsFirewallrulesPatch**](NicApi.md#DatacentersServersNicsFirewallrulesPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Partially modify a Firewall Rule
[**DatacentersServersNicsFirewallrulesPost**](NicApi.md#DatacentersServersNicsFirewallrulesPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules | Create a Firewall Rule
[**DatacentersServersNicsFirewallrulesPut**](NicApi.md#DatacentersServersNicsFirewallrulesPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId}/firewallrules/{firewallruleId} | Modify a Firewall Rule
[**DatacentersServersNicsGet**](NicApi.md#DatacentersServersNicsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/nics | List Nics 
[**DatacentersServersNicsPatch**](NicApi.md#DatacentersServersNicsPatch) | **Patch** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Partially modify a Nic
[**DatacentersServersNicsPost**](NicApi.md#DatacentersServersNicsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/nics | Create a Nic
[**DatacentersServersNicsPut**](NicApi.md#DatacentersServersNicsPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/nics/{nicId} | Modify a Nic



## DatacentersServersNicsDelete

> map[string]interface{} DatacentersServersNicsDelete(ctx, datacenterId, serverId, nicId, optional)

Delete a Nic

Deletes the specified NIC.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***DatacentersServersNicsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsDeleteOpts struct


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


## DatacentersServersNicsFindById

> Nic DatacentersServersNicsFindById(ctx, datacenterId, serverId, nicId, optional)

Retrieve a Nic

Retrieves the attributes of a given NIC

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***DatacentersServersNicsFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsFindByIdOpts struct


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


## DatacentersServersNicsFirewallrulesDelete

> map[string]interface{} DatacentersServersNicsFirewallrulesDelete(ctx, datacenterId, serverId, nicId, firewallruleId, optional)

Delete a Firewall Rule

Removes the specific Firewall Rule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nicId** | **string**| The unique ID of the NIC | 
**firewallruleId** | **string**| The unique ID of the Firewall Rule | 
 **optional** | ***DatacentersServersNicsFirewallrulesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsFirewallrulesDeleteOpts struct


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


## DatacentersServersNicsFirewallrulesFindById

> FirewallRule DatacentersServersNicsFirewallrulesFindById(ctx, datacenterId, serverId, nicId, firewallruleId, optional)

Retrieve a Firewall Rule

Retrieves the attributes of a given Firewall Rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nicId** | **string**| The unique ID of the NIC | 
**firewallruleId** | **string**| The unique ID of the Firewall Rule | 
 **optional** | ***DatacentersServersNicsFirewallrulesFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsFirewallrulesFindByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesGet

> FirewallRules DatacentersServersNicsFirewallrulesGet(ctx, datacenterId, serverId, nicId, optional)

List Firewall Rules 

Retrieves a list of firewall rules associated with a particular NIC

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nicId** | **string**| The unique ID of the NIC | 
 **optional** | ***DatacentersServersNicsFirewallrulesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsFirewallrulesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRules**](FirewallRules.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesPatch

> FirewallRule DatacentersServersNicsFirewallrulesPatch(ctx, datacenterId, serverId, nicId, firewallruleId, firewallrule, optional)

Partially modify a Firewall Rule

You can use update attributes of a resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nicId** | **string**| The unique ID of the NIC | 
**firewallruleId** | **string**| The unique ID of the Firewall Rule | 
**firewallrule** | [**FirewallruleProperties**](FirewallruleProperties.md)| Modified Firewall Rule | 
 **optional** | ***DatacentersServersNicsFirewallrulesPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsFirewallrulesPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesPost

> FirewallRule DatacentersServersNicsFirewallrulesPost(ctx, datacenterId, serverId, nicId, firewallrule, optional)

Create a Firewall Rule

This will add a Firewall Rule to the NIC

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the server | 
**nicId** | **string**| The unique ID of the NIC | 
**firewallrule** | [**FirewallRule**](FirewallRule.md)| Firewall Rule to be created | 
 **optional** | ***DatacentersServersNicsFirewallrulesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsFirewallrulesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsFirewallrulesPut

> FirewallRule DatacentersServersNicsFirewallrulesPut(ctx, datacenterId, serverId, nicId, firewallruleId, firewallrule, optional)

Modify a Firewall Rule

You can use update attributes of a resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nicId** | **string**| The unique ID of the NIC | 
**firewallruleId** | **string**| The unique ID of the Firewall Rule | 
**firewallrule** | [**FirewallRule**](FirewallRule.md)| Modified Firewall Rule | 
 **optional** | ***DatacentersServersNicsFirewallrulesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsFirewallrulesPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsGet

> Nics DatacentersServersNicsGet(ctx, datacenterId, serverId, optional)

List Nics 

Retrieves a list of NICs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
 **optional** | ***DatacentersServersNicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Nics**](Nics.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersServersNicsPatch

> Nic DatacentersServersNicsPatch(ctx, datacenterId, serverId, nicId, nic, optional)

Partially modify a Nic

You can use update attributes of a Nic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nicId** | **string**| The unique ID of the NIC | 
**nic** | [**NicProperties**](NicProperties.md)| Modified properties of Nic | 
 **optional** | ***DatacentersServersNicsPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsPatchOpts struct


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


## DatacentersServersNicsPost

> Nic DatacentersServersNicsPost(ctx, datacenterId, serverId, nic, optional)

Create a Nic

Adds a NIC to the target server. Combine count of Nics and volumes attached to the server should not exceed size 24.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nic** | [**Nic**](Nic.md)| Nic to be created | 
 **optional** | ***DatacentersServersNicsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsPostOpts struct


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


## DatacentersServersNicsPut

> Nic DatacentersServersNicsPut(ctx, datacenterId, serverId, nicId, nic, optional)

Modify a Nic

You can use update attributes of a Nic

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**serverId** | **string**| The unique ID of the Server | 
**nicId** | **string**| The unique ID of the NIC | 
**nic** | [**Nic**](Nic.md)| Modified Nic | 
 **optional** | ***DatacentersServersNicsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersServersNicsPutOpts struct


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

