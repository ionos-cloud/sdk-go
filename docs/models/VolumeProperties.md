# VolumeProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the  resource. | [optional] |
|**Type** | Pointer to **string** | Hardware type of the volume. DAS (Direct Attached Storage) could be used only in a composite call with a Cube server. | [optional] |
|**Size** | **float32** | The size of the volume in GB. | |
|**AvailabilityZone** | Pointer to **string** | The availability zone in which the volume should be provisioned. The storage volume will be provisioned on as few physical storage devices as possible, but this cannot be guaranteed upfront. This is uavailable for DAS (Direct Attached Storage), and subject to availability for SSD. | [optional] |
|**Image** | Pointer to **string** | Image or snapshot ID to be used as template for this volume. | [optional] |
|**ImagePassword** | Pointer to **string** | Initial password to be set for installed OS. Works with public images only. Not modifiable, forbidden in update requests. Password rules allows all characters from a-z, A-Z, 0-9. | [optional] |
|**ImageAlias** | Pointer to **string** |  | [optional] |
|**SshKeys** | Pointer to **[]string** | Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation. | [optional] |
|**Bus** | Pointer to **string** | The bus type for this volume; default is VIRTIO. | [optional] |
|**LicenceType** | Pointer to **string** | OS type for this volume. | [optional] [readonly] |
|**CpuHotPlug** | Pointer to **bool** | Hot-plug capable CPU (no reboot required). | [optional] |
|**RamHotPlug** | Pointer to **bool** | Hot-plug capable RAM (no reboot required). | [optional] |
|**NicHotPlug** | Pointer to **bool** | Hot-plug capable NIC (no reboot required). | [optional] |
|**NicHotUnplug** | Pointer to **bool** | Hot-unplug capable NIC (no reboot required). | [optional] |
|**DiscVirtioHotPlug** | Pointer to **bool** | Hot-plug capable Virt-IO drive (no reboot required). | [optional] |
|**DiscVirtioHotUnplug** | Pointer to **bool** | Hot-unplug capable Virt-IO drive (no reboot required). Not supported with Windows VMs. | [optional] |
|**DeviceNumber** | Pointer to **int64** | The Logical Unit Number of the storage volume. Null for volumes, not mounted to a VM. | [optional] [readonly] |
|**PciSlot** | Pointer to **int32** | The PCI slot number of the storage volume. Null for volumes, not mounted to a VM. | [optional] [readonly] |
|**BackupunitId** | Pointer to **string** | The ID of the backup unit that the user has access to. The property is immutable and is only allowed to be set on creation of a new a volume. It is mandatory to provide either &#39;public image&#39; or &#39;imageAlias&#39; in conjunction with this property. | [optional] |
|**UserData** | Pointer to **string** | The cloud-init configuration for the volume as base64 encoded string. The property is immutable and is only allowed to be set on creation of a new a volume. It is mandatory to provide either &#39;public image&#39; or &#39;imageAlias&#39; that has cloud-init compatibility in conjunction with this property. | [optional] |
|**BootServer** | Pointer to **string** | The UUID of the attached server. | [optional] [readonly] |
|**BootOrder** | Pointer to **string** | Determines whether the volume will be used as a boot volume. Set to &#x60;NONE&#x60;, the volume will not be used as boot volume. Set to &#x60;PRIMARY&#x60;, the volume will be used as boot volume and all other volumes must be set to &#x60;NONE&#x60;. Set to &#x60;AUTO&#x60; or &#x60;null&#x60; requires all volumes to be set to &#x60;AUTO&#x60; or &#x60;null&#x60;; this will use the legacy behavior, which is to use the volume as a boot volume only if there are no other volumes or cdrom devices. | [optional] [default to "AUTO"]|

## Methods

### NewVolumeProperties

`func NewVolumeProperties(size float32, ) *VolumeProperties`

NewVolumeProperties instantiates a new VolumeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumePropertiesWithDefaults

`func NewVolumePropertiesWithDefaults() *VolumeProperties`

NewVolumePropertiesWithDefaults instantiates a new VolumeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumeProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VolumeProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VolumeProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VolumeProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VolumeProperties) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *VolumeProperties) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeProperties) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeProperties) SetSize(v float32)`

SetSize sets Size field to given value.


### GetAvailabilityZone

`func (o *VolumeProperties) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *VolumeProperties) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *VolumeProperties) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *VolumeProperties) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetImage

`func (o *VolumeProperties) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *VolumeProperties) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *VolumeProperties) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *VolumeProperties) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImagePassword

`func (o *VolumeProperties) GetImagePassword() string`

GetImagePassword returns the ImagePassword field if non-nil, zero value otherwise.

### GetImagePasswordOk

`func (o *VolumeProperties) GetImagePasswordOk() (*string, bool)`

GetImagePasswordOk returns a tuple with the ImagePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePassword

`func (o *VolumeProperties) SetImagePassword(v string)`

SetImagePassword sets ImagePassword field to given value.

### HasImagePassword

`func (o *VolumeProperties) HasImagePassword() bool`

HasImagePassword returns a boolean if a field has been set.

### GetImageAlias

`func (o *VolumeProperties) GetImageAlias() string`

GetImageAlias returns the ImageAlias field if non-nil, zero value otherwise.

### GetImageAliasOk

`func (o *VolumeProperties) GetImageAliasOk() (*string, bool)`

GetImageAliasOk returns a tuple with the ImageAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAlias

`func (o *VolumeProperties) SetImageAlias(v string)`

SetImageAlias sets ImageAlias field to given value.

### HasImageAlias

`func (o *VolumeProperties) HasImageAlias() bool`

HasImageAlias returns a boolean if a field has been set.

### GetSshKeys

`func (o *VolumeProperties) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *VolumeProperties) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *VolumeProperties) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *VolumeProperties) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetBus

`func (o *VolumeProperties) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VolumeProperties) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VolumeProperties) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *VolumeProperties) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetLicenceType

`func (o *VolumeProperties) GetLicenceType() string`

GetLicenceType returns the LicenceType field if non-nil, zero value otherwise.

### GetLicenceTypeOk

`func (o *VolumeProperties) GetLicenceTypeOk() (*string, bool)`

GetLicenceTypeOk returns a tuple with the LicenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenceType

`func (o *VolumeProperties) SetLicenceType(v string)`

SetLicenceType sets LicenceType field to given value.

### HasLicenceType

`func (o *VolumeProperties) HasLicenceType() bool`

HasLicenceType returns a boolean if a field has been set.

### GetCpuHotPlug

`func (o *VolumeProperties) GetCpuHotPlug() bool`

GetCpuHotPlug returns the CpuHotPlug field if non-nil, zero value otherwise.

### GetCpuHotPlugOk

`func (o *VolumeProperties) GetCpuHotPlugOk() (*bool, bool)`

GetCpuHotPlugOk returns a tuple with the CpuHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHotPlug

`func (o *VolumeProperties) SetCpuHotPlug(v bool)`

SetCpuHotPlug sets CpuHotPlug field to given value.

### HasCpuHotPlug

`func (o *VolumeProperties) HasCpuHotPlug() bool`

HasCpuHotPlug returns a boolean if a field has been set.

### GetRamHotPlug

`func (o *VolumeProperties) GetRamHotPlug() bool`

GetRamHotPlug returns the RamHotPlug field if non-nil, zero value otherwise.

### GetRamHotPlugOk

`func (o *VolumeProperties) GetRamHotPlugOk() (*bool, bool)`

GetRamHotPlugOk returns a tuple with the RamHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamHotPlug

`func (o *VolumeProperties) SetRamHotPlug(v bool)`

SetRamHotPlug sets RamHotPlug field to given value.

### HasRamHotPlug

`func (o *VolumeProperties) HasRamHotPlug() bool`

HasRamHotPlug returns a boolean if a field has been set.

### GetNicHotPlug

`func (o *VolumeProperties) GetNicHotPlug() bool`

GetNicHotPlug returns the NicHotPlug field if non-nil, zero value otherwise.

### GetNicHotPlugOk

`func (o *VolumeProperties) GetNicHotPlugOk() (*bool, bool)`

GetNicHotPlugOk returns a tuple with the NicHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicHotPlug

`func (o *VolumeProperties) SetNicHotPlug(v bool)`

SetNicHotPlug sets NicHotPlug field to given value.

### HasNicHotPlug

`func (o *VolumeProperties) HasNicHotPlug() bool`

HasNicHotPlug returns a boolean if a field has been set.

### GetNicHotUnplug

`func (o *VolumeProperties) GetNicHotUnplug() bool`

GetNicHotUnplug returns the NicHotUnplug field if non-nil, zero value otherwise.

### GetNicHotUnplugOk

`func (o *VolumeProperties) GetNicHotUnplugOk() (*bool, bool)`

GetNicHotUnplugOk returns a tuple with the NicHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicHotUnplug

`func (o *VolumeProperties) SetNicHotUnplug(v bool)`

SetNicHotUnplug sets NicHotUnplug field to given value.

### HasNicHotUnplug

`func (o *VolumeProperties) HasNicHotUnplug() bool`

HasNicHotUnplug returns a boolean if a field has been set.

### GetDiscVirtioHotPlug

`func (o *VolumeProperties) GetDiscVirtioHotPlug() bool`

GetDiscVirtioHotPlug returns the DiscVirtioHotPlug field if non-nil, zero value otherwise.

### GetDiscVirtioHotPlugOk

`func (o *VolumeProperties) GetDiscVirtioHotPlugOk() (*bool, bool)`

GetDiscVirtioHotPlugOk returns a tuple with the DiscVirtioHotPlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscVirtioHotPlug

`func (o *VolumeProperties) SetDiscVirtioHotPlug(v bool)`

SetDiscVirtioHotPlug sets DiscVirtioHotPlug field to given value.

### HasDiscVirtioHotPlug

`func (o *VolumeProperties) HasDiscVirtioHotPlug() bool`

HasDiscVirtioHotPlug returns a boolean if a field has been set.

### GetDiscVirtioHotUnplug

`func (o *VolumeProperties) GetDiscVirtioHotUnplug() bool`

GetDiscVirtioHotUnplug returns the DiscVirtioHotUnplug field if non-nil, zero value otherwise.

### GetDiscVirtioHotUnplugOk

`func (o *VolumeProperties) GetDiscVirtioHotUnplugOk() (*bool, bool)`

GetDiscVirtioHotUnplugOk returns a tuple with the DiscVirtioHotUnplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscVirtioHotUnplug

`func (o *VolumeProperties) SetDiscVirtioHotUnplug(v bool)`

SetDiscVirtioHotUnplug sets DiscVirtioHotUnplug field to given value.

### HasDiscVirtioHotUnplug

`func (o *VolumeProperties) HasDiscVirtioHotUnplug() bool`

HasDiscVirtioHotUnplug returns a boolean if a field has been set.

### GetDeviceNumber

`func (o *VolumeProperties) GetDeviceNumber() int64`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *VolumeProperties) GetDeviceNumberOk() (*int64, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNumber

`func (o *VolumeProperties) SetDeviceNumber(v int64)`

SetDeviceNumber sets DeviceNumber field to given value.

### HasDeviceNumber

`func (o *VolumeProperties) HasDeviceNumber() bool`

HasDeviceNumber returns a boolean if a field has been set.

### GetPciSlot

`func (o *VolumeProperties) GetPciSlot() int32`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *VolumeProperties) GetPciSlotOk() (*int32, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *VolumeProperties) SetPciSlot(v int32)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *VolumeProperties) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetBackupunitId

`func (o *VolumeProperties) GetBackupunitId() string`

GetBackupunitId returns the BackupunitId field if non-nil, zero value otherwise.

### GetBackupunitIdOk

`func (o *VolumeProperties) GetBackupunitIdOk() (*string, bool)`

GetBackupunitIdOk returns a tuple with the BackupunitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupunitId

`func (o *VolumeProperties) SetBackupunitId(v string)`

SetBackupunitId sets BackupunitId field to given value.

### HasBackupunitId

`func (o *VolumeProperties) HasBackupunitId() bool`

HasBackupunitId returns a boolean if a field has been set.

### GetUserData

`func (o *VolumeProperties) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *VolumeProperties) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *VolumeProperties) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *VolumeProperties) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetBootServer

`func (o *VolumeProperties) GetBootServer() string`

GetBootServer returns the BootServer field if non-nil, zero value otherwise.

### GetBootServerOk

`func (o *VolumeProperties) GetBootServerOk() (*string, bool)`

GetBootServerOk returns a tuple with the BootServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootServer

`func (o *VolumeProperties) SetBootServer(v string)`

SetBootServer sets BootServer field to given value.

### HasBootServer

`func (o *VolumeProperties) HasBootServer() bool`

HasBootServer returns a boolean if a field has been set.

### GetBootOrder

`func (o *VolumeProperties) GetBootOrder() string`

GetBootOrder returns the BootOrder field if non-nil, zero value otherwise.

### GetBootOrderOk

`func (o *VolumeProperties) GetBootOrderOk() (*string, bool)`

GetBootOrderOk returns a tuple with the BootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOrder

`func (o *VolumeProperties) SetBootOrder(v string)`

SetBootOrder sets BootOrder field to given value.

### HasBootOrder

`func (o *VolumeProperties) HasBootOrder() bool`

HasBootOrder returns a boolean if a field has been set.



