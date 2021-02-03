# RequestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**RequestStatusMetadata**](RequestStatusMetadata.md) |  | [optional] 

## Methods

### NewRequestStatus

`func NewRequestStatus() *RequestStatus`

NewRequestStatus instantiates a new RequestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestStatusWithDefaults

`func NewRequestStatusWithDefaults() *RequestStatus`

NewRequestStatusWithDefaults instantiates a new RequestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RequestStatus) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestStatus) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestStatus) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *RequestStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *RequestStatus) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RequestStatus) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RequestStatus) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RequestStatus) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *RequestStatus) GetMetadata() RequestStatusMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RequestStatus) GetMetadataOk() (*RequestStatusMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RequestStatus) SetMetadata(v RequestStatusMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RequestStatus) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


