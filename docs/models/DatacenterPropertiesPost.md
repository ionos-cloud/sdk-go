# DatacenterPropertiesPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the  resource. | [optional] |
|**Description** | Pointer to **string** | A description for the datacenter, such as staging, production. | [optional] |
|**Location** | **string** | The physical location where the datacenter will be created. This will be where all of your servers live. Property cannot be modified after datacenter creation (disallowed in update requests). | |
|**Version** | Pointer to **int32** | The version of the data center; incremented with every change. | [optional] [readonly] |
|**Features** | Pointer to **[]string** | List of features supported by the location where this data center is provisioned. | [optional] [readonly] |
|**SecAuthProtection** | Pointer to **bool** | Boolean value representing if the data center requires extra protection, such as two-step verification. | [optional] |
|**CpuArchitecture** | Pointer to [**[]CpuArchitectureProperties**](CpuArchitectureProperties.md) | Array of features and CPU families available in a location | [optional] [readonly] |
|**CreateDefaultSecurityGroup** | Pointer to **bool** | If true, a default security group, with predefined rules, will be created for the datacenter. Default value is false. | [optional] |

## Methods

### NewDatacenterPropertiesPost

`func NewDatacenterPropertiesPost(location string, ) *DatacenterPropertiesPost`

NewDatacenterPropertiesPost instantiates a new DatacenterPropertiesPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterPropertiesPostWithDefaults

`func NewDatacenterPropertiesPostWithDefaults() *DatacenterPropertiesPost`

NewDatacenterPropertiesPostWithDefaults instantiates a new DatacenterPropertiesPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatacenterPropertiesPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatacenterPropertiesPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatacenterPropertiesPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatacenterPropertiesPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DatacenterPropertiesPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatacenterPropertiesPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatacenterPropertiesPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatacenterPropertiesPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *DatacenterPropertiesPost) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DatacenterPropertiesPost) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DatacenterPropertiesPost) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetVersion

`func (o *DatacenterPropertiesPost) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatacenterPropertiesPost) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatacenterPropertiesPost) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatacenterPropertiesPost) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFeatures

`func (o *DatacenterPropertiesPost) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DatacenterPropertiesPost) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DatacenterPropertiesPost) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DatacenterPropertiesPost) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSecAuthProtection

`func (o *DatacenterPropertiesPost) GetSecAuthProtection() bool`

GetSecAuthProtection returns the SecAuthProtection field if non-nil, zero value otherwise.

### GetSecAuthProtectionOk

`func (o *DatacenterPropertiesPost) GetSecAuthProtectionOk() (*bool, bool)`

GetSecAuthProtectionOk returns a tuple with the SecAuthProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAuthProtection

`func (o *DatacenterPropertiesPost) SetSecAuthProtection(v bool)`

SetSecAuthProtection sets SecAuthProtection field to given value.

### HasSecAuthProtection

`func (o *DatacenterPropertiesPost) HasSecAuthProtection() bool`

HasSecAuthProtection returns a boolean if a field has been set.

### GetCpuArchitecture

`func (o *DatacenterPropertiesPost) GetCpuArchitecture() []CpuArchitectureProperties`

GetCpuArchitecture returns the CpuArchitecture field if non-nil, zero value otherwise.

### GetCpuArchitectureOk

`func (o *DatacenterPropertiesPost) GetCpuArchitectureOk() (*[]CpuArchitectureProperties, bool)`

GetCpuArchitectureOk returns a tuple with the CpuArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArchitecture

`func (o *DatacenterPropertiesPost) SetCpuArchitecture(v []CpuArchitectureProperties)`

SetCpuArchitecture sets CpuArchitecture field to given value.

### HasCpuArchitecture

`func (o *DatacenterPropertiesPost) HasCpuArchitecture() bool`

HasCpuArchitecture returns a boolean if a field has been set.

### GetCreateDefaultSecurityGroup

`func (o *DatacenterPropertiesPost) GetCreateDefaultSecurityGroup() bool`

GetCreateDefaultSecurityGroup returns the CreateDefaultSecurityGroup field if non-nil, zero value otherwise.

### GetCreateDefaultSecurityGroupOk

`func (o *DatacenterPropertiesPost) GetCreateDefaultSecurityGroupOk() (*bool, bool)`

GetCreateDefaultSecurityGroupOk returns a tuple with the CreateDefaultSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDefaultSecurityGroup

`func (o *DatacenterPropertiesPost) SetCreateDefaultSecurityGroup(v bool)`

SetCreateDefaultSecurityGroup sets CreateDefaultSecurityGroup field to given value.

### HasCreateDefaultSecurityGroup

`func (o *DatacenterPropertiesPost) HasCreateDefaultSecurityGroup() bool`

HasCreateDefaultSecurityGroup returns a boolean if a field has been set.



