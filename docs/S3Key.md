# S3Key

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of the resource | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**S3KeyMetadata**](S3KeyMetadata.md) |  | [optional] 
**Properties** | [**S3KeyProperties**](S3KeyProperties.md) |  | 

## Methods

### NewS3Key

`func NewS3Key(properties S3KeyProperties, ) *S3Key`

NewS3Key instantiates a new S3Key object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3KeyWithDefaults

`func NewS3KeyWithDefaults() *S3Key`

NewS3KeyWithDefaults instantiates a new S3Key object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3Key) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3Key) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3Key) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *S3Key) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *S3Key) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3Key) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3Key) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *S3Key) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *S3Key) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *S3Key) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *S3Key) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *S3Key) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *S3Key) GetMetadata() S3KeyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *S3Key) GetMetadataOk() (*S3KeyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *S3Key) SetMetadata(v S3KeyMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *S3Key) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *S3Key) GetProperties() S3KeyProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *S3Key) GetPropertiesOk() (*S3KeyProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *S3Key) SetProperties(v S3KeyProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


