# DatacenterElementMetadata

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Etag** | Pointer to **string** | Resource's Entity Tag as defined in [http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html\#sec3.11](http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11) . Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter. | \[optional\] \[readonly\] |
| **CreatedDate** | Pointer to [**time.Time**](https://github.com/ionos-cloud/sdk-go/tree/5bc758a50995faf1991f97144241d925c7bba695/docs/time.Time.md) | The last time the resource was created | \[optional\] \[readonly\] |
| **CreatedBy** | Pointer to **string** | The user who created the resource. | \[optional\] \[readonly\] |
| **CreatedByUserId** | Pointer to **string** | The user id of the user who has created the resource. | \[optional\] \[readonly\] |
| **LastModifiedDate** | Pointer to [**time.Time**](https://github.com/ionos-cloud/sdk-go/tree/5bc758a50995faf1991f97144241d925c7bba695/docs/time.Time.md) | The last time the resource has been modified | \[optional\] \[readonly\] |
| **LastModifiedBy** | Pointer to **string** | The user who last modified the resource. | \[optional\] \[readonly\] |
| **LastModifiedByUserId** | Pointer to **string** | The user id of the user who has last modified the resource. | \[optional\] \[readonly\] |
| **State** | Pointer to **string** | State of the resource. _AVAILABLE_ There are no pending modification requests for this item; _BUSY_ There is at least one modification request pending and all following requests will be queued; _INACTIVE_ Resource has been de-provisioned; _DEPLOYING_ Resource state DEPLOYING - relevant for Kubernetes cluster/nodepool; _ACTIVE_ Resource state ACTIVE - relevant for Kubernetes cluster/nodepool; _FAILED_ Resource state FAILED - relevant for Kubernetes cluster/nodepool; _SUSPENDED_ Resource state SUSPENDED - relevant for Kubernetes cluster/nodepool; _FAILED\_SUSPENDED_ Resource state FAILED\_SUSPENDED - relevant for Kubernetes cluster; _UPDATING_ Resource state UPDATING - relevant for Kubernetes cluster/nodepool; _FAILED\_UPDATING_ Resource state FAILED\_UPDATING - relevant for Kubernetes cluster/nodepool; _DESTROYING_ Resource state DESTROYING - relevant for Kubernetes cluster; _FAILED\_DESTROYING_ Resource state FAILED\_DESTROYING - relevant for Kubernetes cluster/nodepool; _TERMINATED_ Resource state TERMINATED - relevant for Kubernetes cluster/nodepool | \[optional\] \[readonly\] |

## Methods

### NewDatacenterElementMetadata

`func NewDatacenterElementMetadata() *DatacenterElementMetadata`

NewDatacenterElementMetadata instantiates a new DatacenterElementMetadata object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewDatacenterElementMetadataWithDefaults

`func NewDatacenterElementMetadataWithDefaults() *DatacenterElementMetadata`

NewDatacenterElementMetadataWithDefaults instantiates a new DatacenterElementMetadata object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetEtag

`func (o *DatacenterElementMetadata) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *DatacenterElementMetadata) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetEtag

`func (o *DatacenterElementMetadata) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *DatacenterElementMetadata) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetCreatedDate

`func (o *DatacenterElementMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DatacenterElementMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DatacenterElementMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DatacenterElementMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DatacenterElementMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DatacenterElementMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DatacenterElementMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DatacenterElementMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *DatacenterElementMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *DatacenterElementMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *DatacenterElementMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *DatacenterElementMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *DatacenterElementMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *DatacenterElementMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *DatacenterElementMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *DatacenterElementMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *DatacenterElementMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *DatacenterElementMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *DatacenterElementMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *DatacenterElementMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *DatacenterElementMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *DatacenterElementMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *DatacenterElementMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *DatacenterElementMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetState

`func (o *DatacenterElementMetadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DatacenterElementMetadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetState

`func (o *DatacenterElementMetadata) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DatacenterElementMetadata) HasState() bool`

HasState returns a boolean if a field has been set.

