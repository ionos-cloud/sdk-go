# ImageProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource | [optional] 
**Description** | **string** | Human readable description | [optional] 
**Location** | **string** | Location of that image/snapshot.  | [optional] [readonly] 
**Size** | **float32** | The size of the image in GB | [optional] [readonly] 
**CpuHotPlug** | **bool** | Is capable of CPU hot plug (no reboot required) | [optional] 
**CpuHotUnplug** | **bool** | Is capable of CPU hot unplug (no reboot required) | [optional] 
**RamHotPlug** | **bool** | Is capable of memory hot plug (no reboot required) | [optional] 
**RamHotUnplug** | **bool** | Is capable of memory hot unplug (no reboot required) | [optional] 
**NicHotPlug** | **bool** | Is capable of nic hot plug (no reboot required) | [optional] 
**NicHotUnplug** | **bool** | Is capable of nic hot unplug (no reboot required) | [optional] 
**DiscVirtioHotPlug** | **bool** | Is capable of Virt-IO drive hot plug (no reboot required) | [optional] 
**DiscVirtioHotUnplug** | **bool** | Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. | [optional] 
**DiscScsiHotPlug** | **bool** | Is capable of SCSI drive hot plug (no reboot required) | [optional] 
**DiscScsiHotUnplug** | **bool** | Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. | [optional] 
**LicenceType** | **string** | OS type of this Image | 
**ImageType** | **string** | This indicates the type of image | [optional] [readonly] 
**Public** | **bool** | Indicates if the image is part of the public repository or not | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


