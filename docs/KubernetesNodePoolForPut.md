# KubernetesNodePoolForPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of object | [optional] [readonly] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] 
**Properties** | [**KubernetesNodePoolPropertiesForPut**](KubernetesNodePoolPropertiesForPut.md) |  | 

## Methods

### NewKubernetesNodePoolForPut

`func NewKubernetesNodePoolForPut(properties KubernetesNodePoolPropertiesForPut, ) *KubernetesNodePoolForPut`

NewKubernetesNodePoolForPut instantiates a new KubernetesNodePoolForPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolForPutWithDefaults

`func NewKubernetesNodePoolForPutWithDefaults() *KubernetesNodePoolForPut`

NewKubernetesNodePoolForPutWithDefaults instantiates a new KubernetesNodePoolForPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesNodePoolForPut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesNodePoolForPut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesNodePoolForPut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesNodePoolForPut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KubernetesNodePoolForPut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesNodePoolForPut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesNodePoolForPut) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesNodePoolForPut) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *KubernetesNodePoolForPut) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *KubernetesNodePoolForPut) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *KubernetesNodePoolForPut) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *KubernetesNodePoolForPut) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *KubernetesNodePoolForPut) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesNodePoolForPut) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesNodePoolForPut) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubernetesNodePoolForPut) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *KubernetesNodePoolForPut) GetProperties() KubernetesNodePoolPropertiesForPut`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KubernetesNodePoolForPut) GetPropertiesOk() (*KubernetesNodePoolPropertiesForPut, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KubernetesNodePoolForPut) SetProperties(v KubernetesNodePoolPropertiesForPut)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


