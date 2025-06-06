# ImageProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The resource name. | [optional] |
|**Description** | Pointer to **string** | Human-readable description. | [optional] |
|**Location** | Pointer to **string** | The location of this image/snapshot. | [optional] [readonly] |
|**Size** | Pointer to **float32** | The image size in GB. | [optional] [readonly] |
|**CpuHotPlug** | Pointer to **bool** | Hot-plug capable CPU (no reboot required). | [optional] |
|**CpuHotUnplug** | Pointer to **bool** | Hot-unplug capable CPU (no reboot required). | [optional] |
|**RamHotPlug** | Pointer to **bool** | Hot-plug capable RAM (no reboot required). | [optional] |
|**RamHotUnplug** | Pointer to **bool** | Hot-unplug capable RAM (no reboot required). | [optional] |
|**NicHotPlug** | Pointer to **bool** | Hot-plug capable NIC (no reboot required). | [optional] |
|**NicHotUnplug** | Pointer to **bool** | Hot-unplug capable NIC (no reboot required). | [optional] |
|**DiscVirtioHotPlug** | Pointer to **bool** | Hot-plug capable Virt-IO drive (no reboot required). | [optional] |
|**DiscVirtioHotUnplug** | Pointer to **bool** | Hot-unplug capable Virt-IO drive (no reboot required). Not supported with Windows VMs. | [optional] |
|**DiscScsiHotPlug** | Pointer to **bool** | Hot-plug capable SCSI drive (no reboot required). | [optional] |
|**DiscScsiHotUnplug** | Pointer to **bool** | Hot-unplug capable SCSI drive (no reboot required). Not supported with Windows VMs. | [optional] |
|**ExposeSerial** | Pointer to **bool** | If set to &#x60;true&#x60; will expose the serial id of the disk attached to the server. If set to &#x60;false&#x60; will not expose the serial id. Some operating systems or software solutions require the serial id to be exposed to work properly. Exposing the serial  can influence licensed software (e.g. Windows) behavior | [optional] [default to false]|
|**RequireLegacyBios** | Pointer to **bool** | Indicates if the image requires the legacy BIOS for compatibility or specific needs. | [optional] [default to true]|
|**LicenceType** | **string** | The OS type of this image. | |
|**ApplicationType** | Pointer to **string** | The type of application that is hosted on this resource.  Only public images can have an Application type different than UNKNOWN. | [optional] [default to "UNKNOWN"]|
|**ImageType** | Pointer to **string** | The image type. | [optional] [readonly] |
|**Public** | Pointer to **bool** | Indicates whether the image is part of a public repository. | [optional] [readonly] |
|**ImageAliases** | Pointer to **[]string** | List of image aliases mapped for this image | [optional] [readonly] |
|**CloudInit** | Pointer to **string** | Cloud init compatibility. | [optional] |

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

### GetExposeSerial

`func (o *ImageProperties) GetExposeSerial() bool`

GetExposeSerial returns the ExposeSerial field if non-nil, zero value otherwise.

### GetExposeSerialOk

`func (o *ImageProperties) GetExposeSerialOk() (*bool, bool)`

GetExposeSerialOk returns a tuple with the ExposeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeSerial

`func (o *ImageProperties) SetExposeSerial(v bool)`

SetExposeSerial sets ExposeSerial field to given value.

### HasExposeSerial

`func (o *ImageProperties) HasExposeSerial() bool`

HasExposeSerial returns a boolean if a field has been set.

### GetRequireLegacyBios

`func (o *ImageProperties) GetRequireLegacyBios() bool`

GetRequireLegacyBios returns the RequireLegacyBios field if non-nil, zero value otherwise.

### GetRequireLegacyBiosOk

`func (o *ImageProperties) GetRequireLegacyBiosOk() (*bool, bool)`

GetRequireLegacyBiosOk returns a tuple with the RequireLegacyBios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireLegacyBios

`func (o *ImageProperties) SetRequireLegacyBios(v bool)`

SetRequireLegacyBios sets RequireLegacyBios field to given value.

### HasRequireLegacyBios

`func (o *ImageProperties) HasRequireLegacyBios() bool`

HasRequireLegacyBios returns a boolean if a field has been set.

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


### GetApplicationType

`func (o *ImageProperties) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *ImageProperties) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *ImageProperties) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *ImageProperties) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

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

### GetImageAliases

`func (o *ImageProperties) GetImageAliases() []string`

GetImageAliases returns the ImageAliases field if non-nil, zero value otherwise.

### GetImageAliasesOk

`func (o *ImageProperties) GetImageAliasesOk() (*[]string, bool)`

GetImageAliasesOk returns a tuple with the ImageAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAliases

`func (o *ImageProperties) SetImageAliases(v []string)`

SetImageAliases sets ImageAliases field to given value.

### HasImageAliases

`func (o *ImageProperties) HasImageAliases() bool`

HasImageAliases returns a boolean if a field has been set.

### GetCloudInit

`func (o *ImageProperties) GetCloudInit() string`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *ImageProperties) GetCloudInitOk() (*string, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *ImageProperties) SetCloudInit(v string)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *ImageProperties) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.



