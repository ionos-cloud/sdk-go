# KubernetesNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of object | [optional] [readonly] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**KubernetesNodeMetadata**](KubernetesNodeMetadata.md) |  | [optional] 
**Properties** | [**KubernetesNodeProperties**](KubernetesNodeProperties.md) |  | 

## Methods

### NewKubernetesNode

`func NewKubernetesNode(properties KubernetesNodeProperties, ) *KubernetesNode`

NewKubernetesNode instantiates a new KubernetesNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeWithDefaults

`func NewKubernetesNodeWithDefaults() *KubernetesNode`

NewKubernetesNodeWithDefaults instantiates a new KubernetesNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KubernetesNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesNode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *KubernetesNode) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *KubernetesNode) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *KubernetesNode) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *KubernetesNode) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *KubernetesNode) GetMetadata() KubernetesNodeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesNode) GetMetadataOk() (*KubernetesNodeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesNode) SetMetadata(v KubernetesNodeMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubernetesNode) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *KubernetesNode) GetProperties() KubernetesNodeProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KubernetesNode) GetPropertiesOk() (*KubernetesNodeProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KubernetesNode) SetProperties(v KubernetesNodeProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


