# ImageProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name of that resource | [optional] 
**Description** | Pointer to **string** | Human readable description | [optional] 
**Location** | Pointer to **string** | Location of that image/snapshot.  | [optional] [readonly] 
**Size** | Pointer to **float32** | The size of the image in GB | [optional] [readonly] 
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
**LicenceType** | **string** | OS type of this Image | 
**ImageType** | Pointer to **string** | This indicates the type of image | [optional] [readonly] 
**Public** | Pointer to **bool** | Indicates if the image is part of the public repository or not | [optional] [readonly] 

## Methods

### NewImageProperties

`func NewImageProperties(licenceType string, ) *ImageProperties`

NewImageProperties instantiates a new ImageProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePropertiesWithDefaults

`func NewImagePropertiesWithDefaults() *ImageProperties`

NewImagePropertiesWithDefaults instantiates a new ImageProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ImageProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *ImageProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ImageProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ImageProperties) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ImageProperties) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSize

`func (o *ImageProperties) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImageProperties) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImageProperties) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImageProperties) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCpuHotPlug

`func (o *ImageProperties) GetCpuHotPlug() bool`

GetCpuHotPlug returns the CpuHotPlug field if non-nil, zero value otherwise.

### GetCpuHotPlugOk

`func (o *ImageProperties) GetCpuHotPlugOk() (*bool, bool)`

GetCpuHotPlugOk returns a tuple with the CpuHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHotPlug

`func (o *ImageProperties) SetCpuHotPlug(v bool)`

SetCpuHotPlug sets CpuHotPlug field to given value.

### HasCpuHotPlug

`func (o *ImageProperties) HasCpuHotPlug() bool`

HasCpuHotPlug returns a boolean if a field has been set.

### GetCpuHotUnplug

`func (o *ImageProperties) GetCpuHotUnplug() bool`

GetCpuHotUnplug returns the CpuHotUnplug field if non-nil, zero value otherwise.

### GetCpuHotUnplugOk

`func (o *ImageProperties) GetCpuHotUnplugOk() (*bool, bool)`

GetCpuHotUnplugOk returns a tuple with the CpuHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHotUnplug

`func (o *ImageProperties) SetCpuHotUnplug(v bool)`

SetCpuHotUnplug sets CpuHotUnplug field to given value.

### HasCpuHotUnplug

`func (o *ImageProperties) HasCpuHotUnplug() bool`

HasCpuHotUnplug returns a boolean if a field has been set.

### GetRamHotPlug

`func (o *ImageProperties) GetRamHotPlug() bool`

GetRamHotPlug returns the RamHotPlug field if non-nil, zero value otherwise.

### GetRamHotPlugOk

`func (o *ImageProperties) GetRamHotPlugOk() (*bool, bool)`

GetRamHotPlugOk returns a tuple with the RamHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamHotPlug

`func (o *ImageProperties) SetRamHotPlug(v bool)`

SetRamHotPlug sets RamHotPlug field to given value.

### HasRamHotPlug

`func (o *ImageProperties) HasRamHotPlug() bool`

HasRamHotPlug returns a boolean if a field has been set.

### GetRamHotUnplug

`func (o *ImageProperties) GetRamHotUnplug() bool`

GetRamHotUnplug returns the RamHotUnplug field if non-nil, zero value otherwise.

### GetRamHotUnplugOk

`func (o *ImageProperties) GetRamHotUnplugOk() (*bool, bool)`

GetRamHotUnplugOk returns a tuple with the RamHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamHotUnplug

`func (o *ImageProperties) SetRamHotUnplug(v bool)`

SetRamHotUnplug sets RamHotUnplug field to given value.

### HasRamHotUnplug

`func (o *ImageProperties) HasRamHotUnplug() bool`

HasRamHotUnplug returns a boolean if a field has been set.

### GetNicHotPlug

`func (o *ImageProperties) GetNicHotPlug() bool`

GetNicHotPlug returns the NicHotPlug field if non-nil, zero value otherwise.

### GetNicHotPlugOk

`func (o *ImageProperties) GetNicHotPlugOk() (*bool, bool)`

GetNicHotPlugOk returns a tuple with the NicHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicHotPlug

`func (o *ImageProperties) SetNicHotPlug(v bool)`

SetNicHotPlug sets NicHotPlug field to given value.

### HasNicHotPlug

`func (o *ImageProperties) HasNicHotPlug() bool`

HasNicHotPlug returns a boolean if a field has been set.

### GetNicHotUnplug

`func (o *ImageProperties) GetNicHotUnplug() bool`

GetNicHotUnplug returns the NicHotUnplug field if non-nil, zero value otherwise.

### GetNicHotUnplugOk

`func (o *ImageProperties) GetNicHotUnplugOk() (*bool, bool)`

GetNicHotUnplugOk returns a tuple with the NicHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicHotUnplug

`func (o *ImageProperties) SetNicHotUnplug(v bool)`

SetNicHotUnplug sets NicHotUnplug field to given value.

### HasNicHotUnplug

`func (o *ImageProperties) HasNicHotUnplug() bool`

HasNicHotUnplug returns a boolean if a field has been set.

### GetDiscVirtioHotPlug

`func (o *ImageProperties) GetDiscVirtioHotPlug() bool`

GetDiscVirtioHotPlug returns the DiscVirtioHotPlug field if non-nil, zero value otherwise.

### GetDiscVirtioHotPlugOk

`func (o *ImageProperties) GetDiscVirtioHotPlugOk() (*bool, bool)`

GetDiscVirtioHotPlugOk returns a tuple with the DiscVirtioHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscVirtioHotPlug

`func (o *ImageProperties) SetDiscVirtioHotPlug(v bool)`

SetDiscVirtioHotPlug sets DiscVirtioHotPlug field to given value.

### HasDiscVirtioHotPlug

`func (o *ImageProperties) HasDiscVirtioHotPlug() bool`

HasDiscVirtioHotPlug returns a boolean if a field has been set.

### GetDiscVirtioHotUnplug

`func (o *ImageProperties) GetDiscVirtioHotUnplug() bool`

GetDiscVirtioHotUnplug returns the DiscVirtioHotUnplug field if non-nil, zero value otherwise.

### GetDiscVirtioHotUnplugOk

`func (o *ImageProperties) GetDiscVirtioHotUnplugOk() (*bool, bool)`

GetDiscVirtioHotUnplugOk returns a tuple with the DiscVirtioHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscVirtioHotUnplug

`func (o *ImageProperties) SetDiscVirtioHotUnplug(v bool)`

SetDiscVirtioHotUnplug sets DiscVirtioHotUnplug field to given value.

### HasDiscVirtioHotUnplug

`func (o *ImageProperties) HasDiscVirtioHotUnplug() bool`

HasDiscVirtioHotUnplug returns a boolean if a field has been set.

### GetDiscScsiHotPlug

`func (o *ImageProperties) GetDiscScsiHotPlug() bool`

GetDiscScsiHotPlug returns the DiscScsiHotPlug field if non-nil, zero value otherwise.

### GetDiscScsiHotPlugOk

`func (o *ImageProperties) GetDiscScsiHotPlugOk() (*bool, bool)`

GetDiscScsiHotPlugOk returns a tuple with the DiscScsiHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscScsiHotPlug

`func (o *ImageProperties) SetDiscScsiHotPlug(v bool)`

SetDiscScsiHotPlug sets DiscScsiHotPlug field to given value.

### HasDiscScsiHotPlug

`func (o *ImageProperties) HasDiscScsiHotPlug() bool`

HasDiscScsiHotPlug returns a boolean if a field has been set.

### GetDiscScsiHotUnplug

`func (o *ImageProperties) GetDiscScsiHotUnplug() bool`

GetDiscScsiHotUnplug returns the DiscScsiHotUnplug field if non-nil, zero value otherwise.

### GetDiscScsiHotUnplugOk

`func (o *ImageProperties) GetDiscScsiHotUnplugOk() (*bool, bool)`

GetDiscScsiHotUnplugOk returns a tuple with the DiscScsiHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscScsiHotUnplug

`func (o *ImageProperties) SetDiscScsiHotUnplug(v bool)`

SetDiscScsiHotUnplug sets DiscScsiHotUnplug field to given value.

### HasDiscScsiHotUnplug

`func (o *ImageProperties) HasDiscScsiHotUnplug() bool`

HasDiscScsiHotUnplug returns a boolean if a field has been set.

### GetLicenceType

`func (o *ImageProperties) GetLicenceType() string`

GetLicenceType returns the LicenceType field if non-nil, zero value otherwise.

### GetLicenceTypeOk

`func (o *ImageProperties) GetLicenceTypeOk() (*string, bool)`

GetLicenceTypeOk returns a tuple with the LicenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenceType

`func (o *ImageProperties) SetLicenceType(v string)`

SetLicenceType sets LicenceType field to given value.


### GetImageType

`func (o *ImageProperties) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ImageProperties) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ImageProperties) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *ImageProperties) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetPublic

`func (o *ImageProperties) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ImageProperties) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ImageProperties) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ImageProperties) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


