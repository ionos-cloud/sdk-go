# ServerProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**TemplateUuid** | Pointer to **string** | The ID of the template for creating a CUBE server; the available templates for CUBE servers can be found on the templates resource. | [optional] |
|**Name** | Pointer to **string** | The name of the  resource. | [optional] |
|**Cores** | **int32** | The total number of cores for the server. | |
|**Ram** | **int32** | The memory size for the server in MB, such as 2048. Size must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB. | |
|**AvailabilityZone** | Pointer to **string** | The availability zone in which the server should be provisioned. | [optional] |
|**VmState** | Pointer to **string** | Status of the virtual machine. | [optional] [readonly] |
|**BootCdrom** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] |
|**BootVolume** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] |
|**CpuFamily** | Pointer to **string** | CPU architecture on which server gets provisioned; not all CPU architectures are available in all datacenter regions; available CPU architectures can be retrieved from the datacenter resource. | [optional] |
|**Type** | Pointer to **string** | server usages: ENTERPRISE or CUBE | [optional] |

## Methods

### NewServerProperties

`func NewServerProperties(cores int32, ram int32, ) *ServerProperties`

NewServerProperties instantiates a new ServerProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPropertiesWithDefaults

`func NewServerPropertiesWithDefaults() *ServerProperties`

NewServerPropertiesWithDefaults instantiates a new ServerProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateUuid

`func (o *ServerProperties) GetTemplateUuid() string`

GetTemplateUuid returns the TemplateUuid field if non-nil, zero value otherwise.

### GetTemplateUuidOk

`func (o *ServerProperties) GetTemplateUuidOk() (*string, bool)`

GetTemplateUuidOk returns a tuple with the TemplateUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUuid

`func (o *ServerProperties) SetTemplateUuid(v string)`

SetTemplateUuid sets TemplateUuid field to given value.

### HasTemplateUuid

`func (o *ServerProperties) HasTemplateUuid() bool`

HasTemplateUuid returns a boolean if a field has been set.

### GetName

`func (o *ServerProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCores

`func (o *ServerProperties) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *ServerProperties) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *ServerProperties) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetRam

`func (o *ServerProperties) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ServerProperties) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ServerProperties) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetAvailabilityZone

`func (o *ServerProperties) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ServerProperties) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ServerProperties) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ServerProperties) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetVmState

`func (o *ServerProperties) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *ServerProperties) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *ServerProperties) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *ServerProperties) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### GetBootCdrom

`func (o *ServerProperties) GetBootCdrom() ResourceReference`

GetBootCdrom returns the BootCdrom field if non-nil, zero value otherwise.

### GetBootCdromOk

`func (o *ServerProperties) GetBootCdromOk() (*ResourceReference, bool)`

GetBootCdromOk returns a tuple with the BootCdrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCdrom

`func (o *ServerProperties) SetBootCdrom(v ResourceReference)`

SetBootCdrom sets BootCdrom field to given value.

### HasBootCdrom

`func (o *ServerProperties) HasBootCdrom() bool`

HasBootCdrom returns a boolean if a field has been set.

### GetBootVolume

`func (o *ServerProperties) GetBootVolume() ResourceReference`

GetBootVolume returns the BootVolume field if non-nil, zero value otherwise.

### GetBootVolumeOk

`func (o *ServerProperties) GetBootVolumeOk() (*ResourceReference, bool)`

GetBootVolumeOk returns a tuple with the BootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootVolume

`func (o *ServerProperties) SetBootVolume(v ResourceReference)`

SetBootVolume sets BootVolume field to given value.

### HasBootVolume

`func (o *ServerProperties) HasBootVolume() bool`

HasBootVolume returns a boolean if a field has been set.

### GetCpuFamily

`func (o *ServerProperties) GetCpuFamily() string`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *ServerProperties) GetCpuFamilyOk() (*string, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *ServerProperties) SetCpuFamily(v string)`

SetCpuFamily sets CpuFamily field to given value.

### HasCpuFamily

`func (o *ServerProperties) HasCpuFamily() bool`

HasCpuFamily returns a boolean if a field has been set.

### GetType

`func (o *ServerProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServerProperties) HasType() bool`

HasType returns a boolean if a field has been set.



