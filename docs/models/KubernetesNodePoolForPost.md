# KubernetesNodePoolForPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to **string** | The type of object. | [optional] [readonly] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**KubernetesNodePoolPropertiesForPost**](KubernetesNodePoolPropertiesForPost.md) |  | |

## Methods

### NewKubernetesNodePoolForPost

`func NewKubernetesNodePoolForPost(properties KubernetesNodePoolPropertiesForPost, ) *KubernetesNodePoolForPost`

NewKubernetesNodePoolForPost instantiates a new KubernetesNodePoolForPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolForPostWithDefaults

`func NewKubernetesNodePoolForPostWithDefaults() *KubernetesNodePoolForPost`

NewKubernetesNodePoolForPostWithDefaults instantiates a new KubernetesNodePoolForPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesNodePoolForPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesNodePoolForPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesNodePoolForPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesNodePoolForPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KubernetesNodePoolForPost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesNodePoolForPost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesNodePoolForPost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesNodePoolForPost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *KubernetesNodePoolForPost) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *KubernetesNodePoolForPost) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *KubernetesNodePoolForPost) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *KubernetesNodePoolForPost) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *KubernetesNodePoolForPost) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesNodePoolForPost) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesNodePoolForPost) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubernetesNodePoolForPost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *KubernetesNodePoolForPost) GetProperties() KubernetesNodePoolPropertiesForPost`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KubernetesNodePoolForPost) GetPropertiesOk() (*KubernetesNodePoolPropertiesForPost, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KubernetesNodePoolForPost) SetProperties(v KubernetesNodePoolPropertiesForPost)`

SetProperties sets Properties field to given value.




