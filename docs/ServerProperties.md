# ServerProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource | [optional] 
**Cores** | **int32** | The total number of cores for the server | 
**Ram** | **int32** | The amount of memory for the server in MB, e.g. 2048. Size must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB. | 
**AvailabilityZone** | **string** | The availability zone in which the server should exist | [optional] 
**VmState** | **string** | Status of the virtual Machine | [optional] [readonly] 
**BootCdrom** | [**ResourceReference**](ResourceReference.md) |  | [optional] 
**BootVolume** | [**ResourceReference**](ResourceReference.md) |  | [optional] 
**CpuFamily** | **string** | Cpu family of pserver | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


