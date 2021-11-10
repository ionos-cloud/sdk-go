# CpuArchitectureProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CpuFamily** | Pointer to **string** | A valid CPU family name. | [optional] |
|**MaxCores** | Pointer to **int32** | The maximum number of cores available. | [optional] |
|**MaxRam** | Pointer to **int32** | The maximum RAM size in MB. | [optional] |
|**Vendor** | Pointer to **string** | A valid CPU vendor name. | [optional] |

## Methods

### NewCpuArchitectureProperties

`func NewCpuArchitectureProperties() *CpuArchitectureProperties`

NewCpuArchitectureProperties instantiates a new CpuArchitectureProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpuArchitecturePropertiesWithDefaults

`func NewCpuArchitecturePropertiesWithDefaults() *CpuArchitectureProperties`

NewCpuArchitecturePropertiesWithDefaults instantiates a new CpuArchitectureProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuFamily

`func (o *CpuArchitectureProperties) GetCpuFamily() string`

GetCpuFamily returns the CpuFamily field if non-nil, zero value otherwise.

### GetCpuFamilyOk

`func (o *CpuArchitectureProperties) GetCpuFamilyOk() (*string, bool)`

GetCpuFamilyOk returns a tuple with the CpuFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFamily

`func (o *CpuArchitectureProperties) SetCpuFamily(v string)`

SetCpuFamily sets CpuFamily field to given value.

### HasCpuFamily

`func (o *CpuArchitectureProperties) HasCpuFamily() bool`

HasCpuFamily returns a boolean if a field has been set.

### GetMaxCores

`func (o *CpuArchitectureProperties) GetMaxCores() int32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *CpuArchitectureProperties) GetMaxCoresOk() (*int32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *CpuArchitectureProperties) SetMaxCores(v int32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *CpuArchitectureProperties) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxRam

`func (o *CpuArchitectureProperties) GetMaxRam() int32`

GetMaxRam returns the MaxRam field if non-nil, zero value otherwise.

### GetMaxRamOk

`func (o *CpuArchitectureProperties) GetMaxRamOk() (*int32, bool)`

GetMaxRamOk returns a tuple with the MaxRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRam

`func (o *CpuArchitectureProperties) SetMaxRam(v int32)`

SetMaxRam sets MaxRam field to given value.

### HasMaxRam

`func (o *CpuArchitectureProperties) HasMaxRam() bool`

HasMaxRam returns a boolean if a field has been set.

### GetVendor

`func (o *CpuArchitectureProperties) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CpuArchitectureProperties) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CpuArchitectureProperties) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CpuArchitectureProperties) HasVendor() bool`

HasVendor returns a boolean if a field has been set.



