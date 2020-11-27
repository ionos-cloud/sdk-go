# BackupUnitProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource (only alphanumeric characters are acceptable) | 
**Password** | Pointer to **string** | the password associated to that resource | [optional] 
**Email** | Pointer to **string** | The email associated with the backup unit. Bear in mind that this email does not be the same email as of the user. | [optional] 

## Methods

### NewBackupUnitProperties

`func NewBackupUnitProperties(name string, ) *BackupUnitProperties`

NewBackupUnitProperties instantiates a new BackupUnitProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupUnitPropertiesWithDefaults

`func NewBackupUnitPropertiesWithDefaults() *BackupUnitProperties`

NewBackupUnitPropertiesWithDefaults instantiates a new BackupUnitProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BackupUnitProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupUnitProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupUnitProperties) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *BackupUnitProperties) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BackupUnitProperties) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BackupUnitProperties) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BackupUnitProperties) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmail

`func (o *BackupUnitProperties) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BackupUnitProperties) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BackupUnitProperties) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BackupUnitProperties) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


