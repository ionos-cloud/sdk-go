# KubernetesClusters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique representation for Kubernetes Cluster as a collection on a resource. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of resource within a collection | [optional] [readonly] 
**Href** | Pointer to **string** | URL to the collection representation (absolute path) | [optional] [readonly] 
**Items** | Pointer to [**[]KubernetesCluster**](KubernetesCluster.md) | Array of items in that collection | [optional] [readonly] 

## Methods

### NewKubernetesClusters

`func NewKubernetesClusters() *KubernetesClusters`

NewKubernetesClusters instantiates a new KubernetesClusters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesClustersWithDefaults

`func NewKubernetesClustersWithDefaults() *KubernetesClusters`

NewKubernetesClustersWithDefaults instantiates a new KubernetesClusters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesClusters) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesClusters) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesClusters) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesClusters) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KubernetesClusters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesClusters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesClusters) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesClusters) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *KubernetesClusters) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *KubernetesClusters) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *KubernetesClusters) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *KubernetesClusters) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *KubernetesClusters) GetItems() []KubernetesCluster`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KubernetesClusters) GetItemsOk() (*[]KubernetesCluster, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KubernetesClusters) SetItems(v []KubernetesCluster)`

SetItems sets Items field to given value.

### HasItems

`func (o *KubernetesClusters) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


