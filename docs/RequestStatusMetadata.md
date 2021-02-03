# RequestStatusMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Etag** | Pointer to **string** | Resource&#39;s Entity Tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11 . Entity Tag is also added as an &#39;ETag response header to requests which don&#39;t use &#39;depth&#39; parameter.  | [optional] [readonly] 
**Targets** | Pointer to [**[]RequestTarget**](RequestTarget.md) |  | [optional] 

## Methods

### NewRequestStatusMetadata

`func NewRequestStatusMetadata() *RequestStatusMetadata`

NewRequestStatusMetadata instantiates a new RequestStatusMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestStatusMetadataWithDefaults

`func NewRequestStatusMetadataWithDefaults() *RequestStatusMetadata`

NewRequestStatusMetadataWithDefaults instantiates a new RequestStatusMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestStatusMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestStatusMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestStatusMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestStatusMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *RequestStatusMetadata) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RequestStatusMetadata) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RequestStatusMetadata) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RequestStatusMetadata) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEtag

`func (o *RequestStatusMetadata) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *RequestStatusMetadata) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *RequestStatusMetadata) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *RequestStatusMetadata) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetTargets

`func (o *RequestStatusMetadata) GetTargets() []RequestTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *RequestStatusMetadata) GetTargetsOk() (*[]RequestTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *RequestStatusMetadata) SetTargets(v []RequestTarget)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *RequestStatusMetadata) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


