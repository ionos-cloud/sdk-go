# \VolumeApi

All URIs are relative to *https://api.ionos.com/cloudapi/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatacentersVolumesCreateSnapshotPost**](VolumeApi.md#DatacentersVolumesCreateSnapshotPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/create-snapshot | Create Volume Snapshot
[**DatacentersVolumesDelete**](VolumeApi.md#DatacentersVolumesDelete) | **Delete** /datacenters/{datacenterId}/volumes/{volumeId} | Delete a Volume
[**DatacentersVolumesFindById**](VolumeApi.md#DatacentersVolumesFindById) | **Get** /datacenters/{datacenterId}/volumes/{volumeId} | Retrieve a Volume
[**DatacentersVolumesGet**](VolumeApi.md#DatacentersVolumesGet) | **Get** /datacenters/{datacenterId}/volumes | List Volumes 
[**DatacentersVolumesPatch**](VolumeApi.md#DatacentersVolumesPatch) | **Patch** /datacenters/{datacenterId}/volumes/{volumeId} | Partially modify a Volume
[**DatacentersVolumesPost**](VolumeApi.md#DatacentersVolumesPost) | **Post** /datacenters/{datacenterId}/volumes | Create a Volume
[**DatacentersVolumesPut**](VolumeApi.md#DatacentersVolumesPut) | **Put** /datacenters/{datacenterId}/volumes/{volumeId} | Modify a Volume
[**DatacentersVolumesRestoreSnapshotPost**](VolumeApi.md#DatacentersVolumesRestoreSnapshotPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/restore-snapshot | Restore Volume Snapshot



## DatacentersVolumesCreateSnapshotPost

> Snapshot DatacentersVolumesCreateSnapshotPost(ctx, datacenterId, volumeId, optional)

Create Volume Snapshot

Creates a snapshot of a volume within the datacenter. You can use a snapshot to create a new storage volume or to restore a storage volume.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***DatacentersVolumesCreateSnapshotPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersVolumesCreateSnapshotPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 
 **name** | **optional.String**| The name of the snapshot | 
 **description** | **optional.String**| The description of the snapshot | 
 **secAuthProtection** | **optional.Bool**| Flag representing if extra protection is enabled on snapshot e.g. Two Factor protection etc. | 
 **licenceType** | **optional.String**| The OS type of this Snapshot | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesDelete

> map[string]interface{} DatacentersVolumesDelete(ctx, datacenterId, volumeId, optional)

Delete a Volume

Deletes the specified volume. This will result in the volume being removed from your datacenter. Use this with caution.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***DatacentersVolumesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersVolumesDeleteOpts struct


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


## DatacentersVolumesFindById

> Volume DatacentersVolumesFindById(ctx, datacenterId, volumeId, optional)

Retrieve a Volume

Retrieves the attributes of a given Volume

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***DatacentersVolumesFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersVolumesFindByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesGet

> Volumes DatacentersVolumesGet(ctx, datacenterId, optional)

List Volumes 

Retrieves a list of Volumes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
 **optional** | ***DatacentersVolumesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersVolumesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volumes**](Volumes.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesPatch

> Volume DatacentersVolumesPatch(ctx, datacenterId, volumeId, volume, optional)

Partially modify a Volume

You can use update attributes of a Volume

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**volumeId** | **string**| The unique ID of the Volume | 
**volume** | [**VolumeProperties**](VolumeProperties.md)| Modified properties of Volume | 
 **optional** | ***DatacentersVolumesPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersVolumesPatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesPost

> Volume DatacentersVolumesPost(ctx, datacenterId, volume, optional)

Create a Volume

Creates a volume within the datacenter. This will not attach the volume to a server. Please see the Servers section for details on how to attach storage volumes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**volume** | [**Volume**](Volume.md)| Volume to be created | 
 **optional** | ***DatacentersVolumesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersVolumesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesPut

> Volume DatacentersVolumesPut(ctx, datacenterId, volumeId, volume, optional)

Modify a Volume

You can use update attributes of a Volume

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**volumeId** | **string**| The unique ID of the Volume | 
**volume** | [**Volume**](Volume.md)| Modified Volume | 
 **optional** | ***DatacentersVolumesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersVolumesPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 

### Return type

[**Volume**](Volume.md)

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatacentersVolumesRestoreSnapshotPost

> map[string]interface{} DatacentersVolumesRestoreSnapshotPost(ctx, datacenterId, volumeId, optional)

Restore Volume Snapshot

This will restore a snapshot onto a volume. A snapshot is created as just another image that can be used to create subsequent volumes if you want or to restore an existing volume.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**datacenterId** | **string**| The unique ID of the datacenter | 
**volumeId** | **string**| The unique ID of the Volume | 
 **optional** | ***DatacentersVolumesRestoreSnapshotPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DatacentersVolumesRestoreSnapshotPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **optional.Bool**| Controls whether response is pretty-printed (with indentation and new lines) | [default to true]
 **depth** | **optional.Int32**| Controls the details depth of response objects.  Eg. GET /datacenters/[ID]  - depth&#x3D;0: only direct properties are included. Children (servers etc.) are not included  - depth&#x3D;1: direct properties and children references are included  - depth&#x3D;2: direct properties and children properties are included  - depth&#x3D;3: direct properties and children properties and children&#39;s children are included  - depth&#x3D;... and so on | [default to 0]
 **xContractNumber** | **optional.Int32**| Users having more than 1 contract need to provide contract number, against which all API requests should be executed | 
 **snapshotId** | **optional.String**| This is the ID of the snapshot | 

### Return type

**map[string]interface{}**

### Authorization

[Basic Authentication](../README.md#Basic Authentication), [Token Authentication](../README.md#Token Authentication)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

