# KubernetesCluster

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Id** | Pointer to **string** | The resource's unique identifier. | \[optional\] \[readonly\] |
| **Type** | Pointer to **string** | The type of object | \[optional\] \[readonly\] |
| **Href** | Pointer to **string** | URL to the object representation \(absolute path\) | \[optional\] \[readonly\] |
| **Metadata** | Pointer to [**DatacenterElementMetadata**](datacenterelementmetadata.md) |  | \[optional\] |
| **Properties** | [**KubernetesClusterProperties**](kubernetesclusterproperties.md) |  |  |
| **Entities** | Pointer to [**KubernetesClusterEntities**](kubernetesclusterentities.md) |  | \[optional\] |

## Methods

### NewKubernetesCluster

`func NewKubernetesCluster(properties KubernetesClusterProperties, ) *KubernetesCluster`

NewKubernetesCluster instantiates a new KubernetesCluster object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewKubernetesClusterWithDefaults

`func NewKubernetesClusterWithDefaults() *KubernetesCluster`

NewKubernetesClusterWithDefaults instantiates a new KubernetesCluster object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KubernetesCluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesCluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesCluster) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesCluster) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *KubernetesCluster) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *KubernetesCluster) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetHref

`func (o *KubernetesCluster) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *KubernetesCluster) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *KubernetesCluster) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesCluster) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesCluster) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubernetesCluster) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *KubernetesCluster) GetProperties() KubernetesClusterProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KubernetesCluster) GetPropertiesOk() (*KubernetesClusterProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetProperties

`func (o *KubernetesCluster) SetProperties(v KubernetesClusterProperties)`

SetProperties sets Properties field to given value.

### GetEntities

`func (o *KubernetesCluster) GetEntities() KubernetesClusterEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *KubernetesCluster) GetEntitiesOk() (*KubernetesClusterEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetEntities

`func (o *KubernetesCluster) SetEntities(v KubernetesClusterEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *KubernetesCluster) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

