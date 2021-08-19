# S3Bucket

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | Name of the S3 bucket | |

## Methods

### NewS3Bucket

`func NewS3Bucket(name string, ) *S3Bucket`

NewS3Bucket instantiates a new S3Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BucketWithDefaults

`func NewS3BucketWithDefaults() *S3Bucket`

NewS3BucketWithDefaults instantiates a new S3Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *S3Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3Bucket) SetName(v string)`

SetName sets Name field to given value.




