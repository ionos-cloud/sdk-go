# RequestMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The last time the resource was created | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The user who created the resource. | [optional] [readonly] 
**Etag** | Pointer to **string** | Resource&#39;s Entity Tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11 . Entity Tag is also added as an &#39;ETag response header to requests which don&#39;t use &#39;depth&#39; parameter.  | [optional] [readonly] 
**RequestStatus** | Pointer to [**RequestStatus**](RequestStatus.md) |  | [optional] 

## Methods

### NewRequestMetadata

`func NewRequestMetadata() *RequestMetadata`

NewRequestMetadata instantiates a new RequestMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMetadataWithDefaults

`func NewRequestMetadataWithDefaults() *RequestMetadata`

NewRequestMetadataWithDefaults instantiates a new RequestMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *RequestMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *RequestMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *RequestMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *RequestMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RequestMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RequestMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEtag

`func (o *RequestMetadata) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *RequestMetadata) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *RequestMetadata) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *RequestMetadata) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetRequestStatus

`func (o *RequestMetadata) GetRequestStatus() RequestStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *RequestMetadata) GetRequestStatusOk() (*RequestStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *RequestMetadata) SetRequestStatus(v RequestStatus)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *RequestMetadata) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


