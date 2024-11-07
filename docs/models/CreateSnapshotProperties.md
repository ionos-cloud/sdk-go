# CreateSnapshotProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the snapshot | [optional] |
|**Description** | Pointer to **string** | The description of the snapshot | [optional] |
|**SecAuthProtection** | Pointer to **bool** | Flag representing if extra protection is enabled on snapshot e.g. Two Factor protection etc. | [optional] |
|**LicenceType** | Pointer to **string** | OS type of this Snapshot | [optional] |

## Methods

### NewCreateSnapshotProperties

`func NewCreateSnapshotProperties() *CreateSnapshotProperties`

NewCreateSnapshotProperties instantiates a new CreateSnapshotProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotPropertiesWithDefaults

`func NewCreateSnapshotPropertiesWithDefaults() *CreateSnapshotProperties`

NewCreateSnapshotPropertiesWithDefaults instantiates a new CreateSnapshotProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSnapshotProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSnapshotProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSnapshotProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSnapshotProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSnapshotProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSnapshotProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSnapshotProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSnapshotProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecAuthProtection

`func (o *CreateSnapshotProperties) GetSecAuthProtection() bool`

GetSecAuthProtection returns the SecAuthProtection field if non-nil, zero value otherwise.

### GetSecAuthProtectionOk

`func (o *CreateSnapshotProperties) GetSecAuthProtectionOk() (*bool, bool)`

GetSecAuthProtectionOk returns a tuple with the SecAuthProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAuthProtection

`func (o *CreateSnapshotProperties) SetSecAuthProtection(v bool)`

SetSecAuthProtection sets SecAuthProtection field to given value.

### HasSecAuthProtection

`func (o *CreateSnapshotProperties) HasSecAuthProtection() bool`

HasSecAuthProtection returns a boolean if a field has been set.

### GetLicenceType

`func (o *CreateSnapshotProperties) GetLicenceType() string`

GetLicenceType returns the LicenceType field if non-nil, zero value otherwise.

### GetLicenceTypeOk

`func (o *CreateSnapshotProperties) GetLicenceTypeOk() (*string, bool)`

GetLicenceTypeOk returns a tuple with the LicenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenceType

`func (o *CreateSnapshotProperties) SetLicenceType(v string)`

SetLicenceType sets LicenceType field to given value.

### HasLicenceType

`func (o *CreateSnapshotProperties) HasLicenceType() bool`

HasLicenceType returns a boolean if a field has been set.



