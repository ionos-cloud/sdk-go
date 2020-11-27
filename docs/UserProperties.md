# UserProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** | first name of the user | [optional] 
**Lastname** | Pointer to **string** | last name of the user | [optional] 
**Email** | Pointer to **string** | email address of the user | [optional] 
**Administrator** | Pointer to **bool** | indicates if the user has admin rights or not | [optional] 
**ForceSecAuth** | Pointer to **bool** | indicates if secure authentication should be forced on the user or not | [optional] 
**SecAuthActive** | Pointer to **bool** | indicates if secure authentication is active for the user or not | [optional] 
**S3CanonicalUserId** | Pointer to **string** | Canonical (S3) id of the user for a given identity | [optional] 
**Password** | Pointer to **string** | User password | [optional] 

## Methods

### NewUserProperties

`func NewUserProperties() *UserProperties`

NewUserProperties instantiates a new UserProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPropertiesWithDefaults

`func NewUserPropertiesWithDefaults() *UserProperties`

NewUserPropertiesWithDefaults instantiates a new UserProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *UserProperties) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserProperties) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserProperties) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserProperties) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserProperties) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserProperties) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserProperties) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserProperties) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetEmail

`func (o *UserProperties) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserProperties) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserProperties) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserProperties) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAdministrator

`func (o *UserProperties) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *UserProperties) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *UserProperties) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *UserProperties) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetForceSecAuth

`func (o *UserProperties) GetForceSecAuth() bool`

GetForceSecAuth returns the ForceSecAuth field if non-nil, zero value otherwise.

### GetForceSecAuthOk

`func (o *UserProperties) GetForceSecAuthOk() (*bool, bool)`

GetForceSecAuthOk returns a tuple with the ForceSecAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSecAuth

`func (o *UserProperties) SetForceSecAuth(v bool)`

SetForceSecAuth sets ForceSecAuth field to given value.

### HasForceSecAuth

`func (o *UserProperties) HasForceSecAuth() bool`

HasForceSecAuth returns a boolean if a field has been set.

### GetSecAuthActive

`func (o *UserProperties) GetSecAuthActive() bool`

GetSecAuthActive returns the SecAuthActive field if non-nil, zero value otherwise.

### GetSecAuthActiveOk

`func (o *UserProperties) GetSecAuthActiveOk() (*bool, bool)`

GetSecAuthActiveOk returns a tuple with the SecAuthActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAuthActive

`func (o *UserProperties) SetSecAuthActive(v bool)`

SetSecAuthActive sets SecAuthActive field to given value.

### HasSecAuthActive

`func (o *UserProperties) HasSecAuthActive() bool`

HasSecAuthActive returns a boolean if a field has been set.

### GetS3CanonicalUserId

`func (o *UserProperties) GetS3CanonicalUserId() string`

GetS3CanonicalUserId returns the S3CanonicalUserId field if non-nil, zero value otherwise.

### GetS3CanonicalUserIdOk

`func (o *UserProperties) GetS3CanonicalUserIdOk() (*string, bool)`

GetS3CanonicalUserIdOk returns a tuple with the S3CanonicalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CanonicalUserId

`func (o *UserProperties) SetS3CanonicalUserId(v string)`

SetS3CanonicalUserId sets S3CanonicalUserId field to given value.

### HasS3CanonicalUserId

`func (o *UserProperties) HasS3CanonicalUserId() bool`

HasS3CanonicalUserId returns a boolean if a field has been set.

### GetPassword

`func (o *UserProperties) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserProperties) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserProperties) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserProperties) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


