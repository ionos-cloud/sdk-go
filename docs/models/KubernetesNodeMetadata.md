# KubernetesNodeMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The date the resource was created. | [optional] [readonly] |
|**Etag** | Pointer to **string** | The resource entity tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11  Entity tags are also added as &#39;ETag&#39; response headers to requests that do not use the &#39;depth&#39; parameter. | [optional] [readonly] |
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The date the resource was last modified. | [optional] [readonly] |
|**LastSoftwareUpdatedDate** | Pointer to [**time.Time**](time.Time.md) | The date when the software on the node was last updated. | [optional] [readonly] |
|**State** | Pointer to **string** | The resource state. | [optional] [readonly] |

## Methods

### NewKubernetesNodeMetadata

`func NewKubernetesNodeMetadata() *KubernetesNodeMetadata`

NewKubernetesNodeMetadata instantiates a new KubernetesNodeMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodeMetadataWithDefaults

`func NewKubernetesNodeMetadataWithDefaults() *KubernetesNodeMetadata`

NewKubernetesNodeMetadataWithDefaults instantiates a new KubernetesNodeMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *KubernetesNodeMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *KubernetesNodeMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *KubernetesNodeMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *KubernetesNodeMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetEtag

`func (o *KubernetesNodeMetadata) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *KubernetesNodeMetadata) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *KubernetesNodeMetadata) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *KubernetesNodeMetadata) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *KubernetesNodeMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *KubernetesNodeMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *KubernetesNodeMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *KubernetesNodeMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastSoftwareUpdatedDate

`func (o *KubernetesNodeMetadata) GetLastSoftwareUpdatedDate() time.Time`

GetLastSoftwareUpdatedDate returns the LastSoftwareUpdatedDate field if non-nil, zero value otherwise.

### GetLastSoftwareUpdatedDateOk

`func (o *KubernetesNodeMetadata) GetLastSoftwareUpdatedDateOk() (*time.Time, bool)`

GetLastSoftwareUpdatedDateOk returns a tuple with the LastSoftwareUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSoftwareUpdatedDate

`func (o *KubernetesNodeMetadata) SetLastSoftwareUpdatedDate(v time.Time)`

SetLastSoftwareUpdatedDate sets LastSoftwareUpdatedDate field to given value.

### HasLastSoftwareUpdatedDate

`func (o *KubernetesNodeMetadata) HasLastSoftwareUpdatedDate() bool`

HasLastSoftwareUpdatedDate returns a boolean if a field has been set.

### GetState

`func (o *KubernetesNodeMetadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KubernetesNodeMetadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KubernetesNodeMetadata) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *KubernetesNodeMetadata) HasState() bool`

HasState returns a boolean if a field has been set.



