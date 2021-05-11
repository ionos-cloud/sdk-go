# S3KeyMetadata

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Etag** | Pointer to **string** | Resource's Entity Tag as defined in [http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html\#sec3.11](http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11) . Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter. | \[optional\] \[readonly\] |
| **CreatedDate** | Pointer to [**time.Time**](https://github.com/ionos-cloud/sdk-go/tree/10b29fc9e7baa2f1a7d1e0996c9c14e99b9eb20d/docs/time.Time.md) | The time the S3 key was created | \[optional\] \[readonly\] |

## Methods

### NewS3KeyMetadata

`func NewS3KeyMetadata() *S3KeyMetadata`

NewS3KeyMetadata instantiates a new S3KeyMetadata object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewS3KeyMetadataWithDefaults

`func NewS3KeyMetadataWithDefaults() *S3KeyMetadata`

NewS3KeyMetadataWithDefaults instantiates a new S3KeyMetadata object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetEtag

`func (o *S3KeyMetadata) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *S3KeyMetadata) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetEtag

`func (o *S3KeyMetadata) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *S3KeyMetadata) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetCreatedDate

`func (o *S3KeyMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *S3KeyMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *S3KeyMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *S3KeyMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

