# UserPropertiesPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Firstname** | Pointer to **string** | The first name of the user. | [optional] |
|**Lastname** | Pointer to **string** | The last name of the user. | [optional] |
|**Email** | Pointer to **string** | The email address of the user. | [optional] |
|**Administrator** | Pointer to **bool** | Indicates if the user has admin rights. | [optional] |
|**ForceSecAuth** | Pointer to **bool** | Indicates if secure authentication should be forced on the user. | [optional] |
|**SecAuthActive** | Pointer to **bool** | Indicates if secure authentication is active for the user. | [optional] |
|**Password** | Pointer to **string** | User password. | [optional] |
|**Active** | Pointer to **bool** | Indicates if the user is active. | [optional] |

## Methods

### NewUserPropertiesPost

`func NewUserPropertiesPost() *UserPropertiesPost`

NewUserPropertiesPost instantiates a new UserPropertiesPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPropertiesPostWithDefaults

`func NewUserPropertiesPostWithDefaults() *UserPropertiesPost`

NewUserPropertiesPostWithDefaults instantiates a new UserPropertiesPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *UserPropertiesPost) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserPropertiesPost) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserPropertiesPost) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserPropertiesPost) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserPropertiesPost) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserPropertiesPost) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserPropertiesPost) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserPropertiesPost) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetEmail

`func (o *UserPropertiesPost) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPropertiesPost) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPropertiesPost) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserPropertiesPost) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAdministrator

`func (o *UserPropertiesPost) GetAdministrator() bool`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *UserPropertiesPost) GetAdministratorOk() (*bool, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *UserPropertiesPost) SetAdministrator(v bool)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *UserPropertiesPost) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetForceSecAuth

`func (o *UserPropertiesPost) GetForceSecAuth() bool`

GetForceSecAuth returns the ForceSecAuth field if non-nil, zero value otherwise.

### GetForceSecAuthOk

`func (o *UserPropertiesPost) GetForceSecAuthOk() (*bool, bool)`

GetForceSecAuthOk returns a tuple with the ForceSecAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSecAuth

`func (o *UserPropertiesPost) SetForceSecAuth(v bool)`

SetForceSecAuth sets ForceSecAuth field to given value.

### HasForceSecAuth

`func (o *UserPropertiesPost) HasForceSecAuth() bool`

HasForceSecAuth returns a boolean if a field has been set.

### GetSecAuthActive

`func (o *UserPropertiesPost) GetSecAuthActive() bool`

GetSecAuthActive returns the SecAuthActive field if non-nil, zero value otherwise.

### GetSecAuthActiveOk

`func (o *UserPropertiesPost) GetSecAuthActiveOk() (*bool, bool)`

GetSecAuthActiveOk returns a tuple with the SecAuthActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAuthActive

`func (o *UserPropertiesPost) SetSecAuthActive(v bool)`

SetSecAuthActive sets SecAuthActive field to given value.

### HasSecAuthActive

`func (o *UserPropertiesPost) HasSecAuthActive() bool`

HasSecAuthActive returns a boolean if a field has been set.

### GetPassword

`func (o *UserPropertiesPost) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPropertiesPost) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPropertiesPost) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserPropertiesPost) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetActive

`func (o *UserPropertiesPost) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserPropertiesPost) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserPropertiesPost) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UserPropertiesPost) HasActive() bool`

HasActive returns a boolean if a field has been set.



