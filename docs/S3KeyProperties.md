# S3KeyProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretKey** | Pointer to **string** | secret of the s3 key | [optional] [readonly] 
**Active** | Pointer to **bool** | denotes if the s3 key is active or not | [optional] 

## Methods

### NewS3KeyProperties

`func NewS3KeyProperties() *S3KeyProperties`

NewS3KeyProperties instantiates a new S3KeyProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3KeyPropertiesWithDefaults

`func NewS3KeyPropertiesWithDefaults() *S3KeyProperties`

NewS3KeyPropertiesWithDefaults instantiates a new S3KeyProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretKey

`func (o *S3KeyProperties) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *S3KeyProperties) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *S3KeyProperties) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *S3KeyProperties) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetActive

`func (o *S3KeyProperties) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *S3KeyProperties) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *S3KeyProperties) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *S3KeyProperties) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


