# SnapshotProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name of that resource | [optional] 
**Description** | Pointer to **string** | Human readable description | [optional] 
**Location** | Pointer to **string** | Location of that image/snapshot.  | [optional] [readonly] 
**Size** | Pointer to **float32** | The size of the image in GB | [optional] [readonly] 
**SecAuthProtection** | Pointer to **bool** | Boolean value representing if the snapshot requires extra protection e.g. two factor protection | [optional] 
**CpuHotPlug** | Pointer to **bool** | Is capable of CPU hot plug (no reboot required) | [optional] 
**CpuHotUnplug** | Pointer to **bool** | Is capable of CPU hot unplug (no reboot required) | [optional] 
**RamHotPlug** | Pointer to **bool** | Is capable of memory hot plug (no reboot required) | [optional] 
**RamHotUnplug** | Pointer to **bool** | Is capable of memory hot unplug (no reboot required) | [optional] 
**NicHotPlug** | Pointer to **bool** | Is capable of nic hot plug (no reboot required) | [optional] 
**NicHotUnplug** | Pointer to **bool** | Is capable of nic hot unplug (no reboot required) | [optional] 
**DiscVirtioHotPlug** | Pointer to **bool** | Is capable of Virt-IO drive hot plug (no reboot required) | [optional] 
**DiscVirtioHotUnplug** | Pointer to **bool** | Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. | [optional] 
**DiscScsiHotPlug** | Pointer to **bool** | Is capable of SCSI drive hot plug (no reboot required) | [optional] 
**DiscScsiHotUnplug** | Pointer to **bool** | Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. | [optional] 
**LicenceType** | Pointer to **string** | OS type of this Snapshot | [optional] 

## Methods

### NewSnapshotProperties

`func NewSnapshotProperties() *SnapshotProperties`

NewSnapshotProperties instantiates a new SnapshotProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotPropertiesWithDefaults

`func NewSnapshotPropertiesWithDefaults() *SnapshotProperties`

NewSnapshotPropertiesWithDefaults instantiates a new SnapshotProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SnapshotProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnapshotProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SnapshotProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnapshotProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *SnapshotProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SnapshotProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SnapshotProperties) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SnapshotProperties) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSize

`func (o *SnapshotProperties) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SnapshotProperties) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SnapshotProperties) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SnapshotProperties) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSecAuthProtection

`func (o *SnapshotProperties) GetSecAuthProtection() bool`

GetSecAuthProtection returns the SecAuthProtection field if non-nil, zero value otherwise.

### GetSecAuthProtectionOk

`func (o *SnapshotProperties) GetSecAuthProtectionOk() (*bool, bool)`

GetSecAuthProtectionOk returns a tuple with the SecAuthProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAuthProtection

`func (o *SnapshotProperties) SetSecAuthProtection(v bool)`

SetSecAuthProtection sets SecAuthProtection field to given value.

### HasSecAuthProtection

`func (o *SnapshotProperties) HasSecAuthProtection() bool`

HasSecAuthProtection returns a boolean if a field has been set.

### GetCpuHotPlug

`func (o *SnapshotProperties) GetCpuHotPlug() bool`

GetCpuHotPlug returns the CpuHotPlug field if non-nil, zero value otherwise.

### GetCpuHotPlugOk

`func (o *SnapshotProperties) GetCpuHotPlugOk() (*bool, bool)`

GetCpuHotPlugOk returns a tuple with the CpuHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHotPlug

`func (o *SnapshotProperties) SetCpuHotPlug(v bool)`

SetCpuHotPlug sets CpuHotPlug field to given value.

### HasCpuHotPlug

`func (o *SnapshotProperties) HasCpuHotPlug() bool`

HasCpuHotPlug returns a boolean if a field has been set.

### GetCpuHotUnplug

`func (o *SnapshotProperties) GetCpuHotUnplug() bool`

GetCpuHotUnplug returns the CpuHotUnplug field if non-nil, zero value otherwise.

### GetCpuHotUnplugOk

`func (o *SnapshotProperties) GetCpuHotUnplugOk() (*bool, bool)`

GetCpuHotUnplugOk returns a tuple with the CpuHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHotUnplug

`func (o *SnapshotProperties) SetCpuHotUnplug(v bool)`

SetCpuHotUnplug sets CpuHotUnplug field to given value.

### HasCpuHotUnplug

`func (o *SnapshotProperties) HasCpuHotUnplug() bool`

HasCpuHotUnplug returns a boolean if a field has been set.

### GetRamHotPlug

`func (o *SnapshotProperties) GetRamHotPlug() bool`

GetRamHotPlug returns the RamHotPlug field if non-nil, zero value otherwise.

### GetRamHotPlugOk

`func (o *SnapshotProperties) GetRamHotPlugOk() (*bool, bool)`

GetRamHotPlugOk returns a tuple with the RamHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamHotPlug

`func (o *SnapshotProperties) SetRamHotPlug(v bool)`

SetRamHotPlug sets RamHotPlug field to given value.

### HasRamHotPlug

`func (o *SnapshotProperties) HasRamHotPlug() bool`

HasRamHotPlug returns a boolean if a field has been set.

### GetRamHotUnplug

`func (o *SnapshotProperties) GetRamHotUnplug() bool`

GetRamHotUnplug returns the RamHotUnplug field if non-nil, zero value otherwise.

### GetRamHotUnplugOk

`func (o *SnapshotProperties) GetRamHotUnplugOk() (*bool, bool)`

GetRamHotUnplugOk returns a tuple with the RamHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamHotUnplug

`func (o *SnapshotProperties) SetRamHotUnplug(v bool)`

SetRamHotUnplug sets RamHotUnplug field to given value.

### HasRamHotUnplug

`func (o *SnapshotProperties) HasRamHotUnplug() bool`

HasRamHotUnplug returns a boolean if a field has been set.

### GetNicHotPlug

`func (o *SnapshotProperties) GetNicHotPlug() bool`

GetNicHotPlug returns the NicHotPlug field if non-nil, zero value otherwise.

### GetNicHotPlugOk

`func (o *SnapshotProperties) GetNicHotPlugOk() (*bool, bool)`

GetNicHotPlugOk returns a tuple with the NicHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicHotPlug

`func (o *SnapshotProperties) SetNicHotPlug(v bool)`

SetNicHotPlug sets NicHotPlug field to given value.

### HasNicHotPlug

`func (o *SnapshotProperties) HasNicHotPlug() bool`

HasNicHotPlug returns a boolean if a field has been set.

### GetNicHotUnplug

`func (o *SnapshotProperties) GetNicHotUnplug() bool`

GetNicHotUnplug returns the NicHotUnplug field if non-nil, zero value otherwise.

### GetNicHotUnplugOk

`func (o *SnapshotProperties) GetNicHotUnplugOk() (*bool, bool)`

GetNicHotUnplugOk returns a tuple with the NicHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicHotUnplug

`func (o *SnapshotProperties) SetNicHotUnplug(v bool)`

SetNicHotUnplug sets NicHotUnplug field to given value.

### HasNicHotUnplug

`func (o *SnapshotProperties) HasNicHotUnplug() bool`

HasNicHotUnplug returns a boolean if a field has been set.

### GetDiscVirtioHotPlug

`func (o *SnapshotProperties) GetDiscVirtioHotPlug() bool`

GetDiscVirtioHotPlug returns the DiscVirtioHotPlug field if non-nil, zero value otherwise.

### GetDiscVirtioHotPlugOk

`func (o *SnapshotProperties) GetDiscVirtioHotPlugOk() (*bool, bool)`

GetDiscVirtioHotPlugOk returns a tuple with the DiscVirtioHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscVirtioHotPlug

`func (o *SnapshotProperties) SetDiscVirtioHotPlug(v bool)`

SetDiscVirtioHotPlug sets DiscVirtioHotPlug field to given value.

### HasDiscVirtioHotPlug

`func (o *SnapshotProperties) HasDiscVirtioHotPlug() bool`

HasDiscVirtioHotPlug returns a boolean if a field has been set.

### GetDiscVirtioHotUnplug

`func (o *SnapshotProperties) GetDiscVirtioHotUnplug() bool`

GetDiscVirtioHotUnplug returns the DiscVirtioHotUnplug field if non-nil, zero value otherwise.

### GetDiscVirtioHotUnplugOk

`func (o *SnapshotProperties) GetDiscVirtioHotUnplugOk() (*bool, bool)`

GetDiscVirtioHotUnplugOk returns a tuple with the DiscVirtioHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscVirtioHotUnplug

`func (o *SnapshotProperties) SetDiscVirtioHotUnplug(v bool)`

SetDiscVirtioHotUnplug sets DiscVirtioHotUnplug field to given value.

### HasDiscVirtioHotUnplug

`func (o *SnapshotProperties) HasDiscVirtioHotUnplug() bool`

HasDiscVirtioHotUnplug returns a boolean if a field has been set.

### GetDiscScsiHotPlug

`func (o *SnapshotProperties) GetDiscScsiHotPlug() bool`

GetDiscScsiHotPlug returns the DiscScsiHotPlug field if non-nil, zero value otherwise.

### GetDiscScsiHotPlugOk

`func (o *SnapshotProperties) GetDiscScsiHotPlugOk() (*bool, bool)`

GetDiscScsiHotPlugOk returns a tuple with the DiscScsiHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscScsiHotPlug

`func (o *SnapshotProperties) SetDiscScsiHotPlug(v bool)`

SetDiscScsiHotPlug sets DiscScsiHotPlug field to given value.

### HasDiscScsiHotPlug

`func (o *SnapshotProperties) HasDiscScsiHotPlug() bool`

HasDiscScsiHotPlug returns a boolean if a field has been set.

### GetDiscScsiHotUnplug

`func (o *SnapshotProperties) GetDiscScsiHotUnplug() bool`

GetDiscScsiHotUnplug returns the DiscScsiHotUnplug field if non-nil, zero value otherwise.

### GetDiscScsiHotUnplugOk

`func (o *SnapshotProperties) GetDiscScsiHotUnplugOk() (*bool, bool)`

GetDiscScsiHotUnplugOk returns a tuple with the DiscScsiHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscScsiHotUnplug

`func (o *SnapshotProperties) SetDiscScsiHotUnplug(v bool)`

SetDiscScsiHotUnplug sets DiscScsiHotUnplug field to given value.

### HasDiscScsiHotUnplug

`func (o *SnapshotProperties) HasDiscScsiHotUnplug() bool`

HasDiscScsiHotUnplug returns a boolean if a field has been set.

### GetLicenceType

`func (o *SnapshotProperties) GetLicenceType() string`

GetLicenceType returns the LicenceType field if non-nil, zero value otherwise.

### GetLicenceTypeOk

`func (o *SnapshotProperties) GetLicenceTypeOk() (*string, bool)`

GetLicenceTypeOk returns a tuple with the LicenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenceType

`func (o *SnapshotProperties) SetLicenceType(v string)`

SetLicenceType sets LicenceType field to given value.

### HasLicenceType

`func (o *SnapshotProperties) HasLicenceType() bool`

HasLicenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


