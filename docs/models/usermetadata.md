# UserMetadata

## Properties

| Name | Type | Description | Notes |
| :--- | :--- | :--- | :--- |
| **Etag** | Pointer to **string** | Resource's Entity Tag as defined in [http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html\#sec3.11](http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11) . Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter. | \[optional\] \[readonly\] |
| **CreatedDate** | Pointer to [**time.Time**](https://github.com/ionos-cloud/sdk-go/tree/5bc758a50995faf1991f97144241d925c7bba695/docs/time.Time.md) | time of creation of the user | \[optional\] \[readonly\] |
| **LastLogin** | Pointer to [**time.Time**](https://github.com/ionos-cloud/sdk-go/tree/5bc758a50995faf1991f97144241d925c7bba695/docs/time.Time.md) | time of last login by the user | \[optional\] \[readonly\] |

## Methods

### NewUserMetadata

`func NewUserMetadata() *UserMetadata`

NewUserMetadata instantiates a new UserMetadata object This constructor will assign default values to properties that have it defined, and makes sure properties required by API are set, but the set of arguments will change when the set of required properties is changed

### NewUserMetadataWithDefaults

`func NewUserMetadataWithDefaults() *UserMetadata`

NewUserMetadataWithDefaults instantiates a new UserMetadata object This constructor will only assign default values to properties that have it defined, but it doesn't guarantee that properties required by API are set

### GetEtag

`func (o *UserMetadata) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *UserMetadata) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetEtag

`func (o *UserMetadata) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *UserMetadata) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetCreatedDate

`func (o *UserMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *UserMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *UserMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *UserMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserMetadata) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserMetadata) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserMetadata) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserMetadata) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.
