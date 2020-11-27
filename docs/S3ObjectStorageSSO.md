# S3ObjectStorageSSO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoUrl** | Pointer to **string** | The S3 object storage single sign on url | [optional] [readonly] 

## Methods

### NewS3ObjectStorageSSO

`func NewS3ObjectStorageSSO() *S3ObjectStorageSSO`

NewS3ObjectStorageSSO instantiates a new S3ObjectStorageSSO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ObjectStorageSSOWithDefaults

`func NewS3ObjectStorageSSOWithDefaults() *S3ObjectStorageSSO`

NewS3ObjectStorageSSOWithDefaults instantiates a new S3ObjectStorageSSO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoUrl

`func (o *S3ObjectStorageSSO) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *S3ObjectStorageSSO) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *S3ObjectStorageSSO) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *S3ObjectStorageSSO) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


