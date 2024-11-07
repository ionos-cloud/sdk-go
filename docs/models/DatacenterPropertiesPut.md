# DatacenterPropertiesPut

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
|**DefaultSecurityGroupId** | Pointer to **string** | This will become the default security group for the datacenter, replacing the old one if already exists.  This security group must already exists prior to this request. Provide this field only if the &#x60;createDefaultSecurityGroup&#x60; field is missing. You cannot provide both of them | [optional] |
|**CreateDefaultSecurityGroup** | Pointer to **bool** | If this field is set on true and this datacenter has no default security group then a default security group, with predefined rules, will be created for this datacenter. Default value is false.  Provide this field only if the &#x60;defaultSecurityGroupId&#x60; field is missing. You cannot provide both of them | [optional] |

## Methods

### NewDatacenterPropertiesPut

`func NewDatacenterPropertiesPut(location string, ) *DatacenterPropertiesPut`

NewDatacenterPropertiesPut instantiates a new DatacenterPropertiesPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterPropertiesPutWithDefaults

`func NewDatacenterPropertiesPutWithDefaults() *DatacenterPropertiesPut`

NewDatacenterPropertiesPutWithDefaults instantiates a new DatacenterPropertiesPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatacenterPropertiesPut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatacenterPropertiesPut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatacenterPropertiesPut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatacenterPropertiesPut) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DatacenterPropertiesPut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatacenterPropertiesPut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatacenterPropertiesPut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatacenterPropertiesPut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *DatacenterPropertiesPut) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DatacenterPropertiesPut) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DatacenterPropertiesPut) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetVersion

`func (o *DatacenterPropertiesPut) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatacenterPropertiesPut) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatacenterPropertiesPut) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatacenterPropertiesPut) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFeatures

`func (o *DatacenterPropertiesPut) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DatacenterPropertiesPut) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DatacenterPropertiesPut) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DatacenterPropertiesPut) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSecAuthProtection

`func (o *DatacenterPropertiesPut) GetSecAuthProtection() bool`

GetSecAuthProtection returns the SecAuthProtection field if non-nil, zero value otherwise.

### GetSecAuthProtectionOk

`func (o *DatacenterPropertiesPut) GetSecAuthProtectionOk() (*bool, bool)`

GetSecAuthProtectionOk returns a tuple with the SecAuthProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAuthProtection

`func (o *DatacenterPropertiesPut) SetSecAuthProtection(v bool)`

SetSecAuthProtection sets SecAuthProtection field to given value.

### HasSecAuthProtection

`func (o *DatacenterPropertiesPut) HasSecAuthProtection() bool`

HasSecAuthProtection returns a boolean if a field has been set.

### GetCpuArchitecture

`func (o *DatacenterPropertiesPut) GetCpuArchitecture() []CpuArchitectureProperties`

GetCpuArchitecture returns the CpuArchitecture field if non-nil, zero value otherwise.

### GetCpuArchitectureOk

`func (o *DatacenterPropertiesPut) GetCpuArchitectureOk() (*[]CpuArchitectureProperties, bool)`

GetCpuArchitectureOk returns a tuple with the CpuArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArchitecture

`func (o *DatacenterPropertiesPut) SetCpuArchitecture(v []CpuArchitectureProperties)`

SetCpuArchitecture sets CpuArchitecture field to given value.

### HasCpuArchitecture

`func (o *DatacenterPropertiesPut) HasCpuArchitecture() bool`

HasCpuArchitecture returns a boolean if a field has been set.

### GetDefaultSecurityGroupId

`func (o *DatacenterPropertiesPut) GetDefaultSecurityGroupId() string`

GetDefaultSecurityGroupId returns the DefaultSecurityGroupId field if non-nil, zero value otherwise.

### GetDefaultSecurityGroupIdOk

`func (o *DatacenterPropertiesPut) GetDefaultSecurityGroupIdOk() (*string, bool)`

GetDefaultSecurityGroupIdOk returns a tuple with the DefaultSecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityGroupId

`func (o *DatacenterPropertiesPut) SetDefaultSecurityGroupId(v string)`

SetDefaultSecurityGroupId sets DefaultSecurityGroupId field to given value.

### HasDefaultSecurityGroupId

`func (o *DatacenterPropertiesPut) HasDefaultSecurityGroupId() bool`

HasDefaultSecurityGroupId returns a boolean if a field has been set.

### GetCreateDefaultSecurityGroup

`func (o *DatacenterPropertiesPut) GetCreateDefaultSecurityGroup() bool`

GetCreateDefaultSecurityGroup returns the CreateDefaultSecurityGroup field if non-nil, zero value otherwise.

### GetCreateDefaultSecurityGroupOk

`func (o *DatacenterPropertiesPut) GetCreateDefaultSecurityGroupOk() (*bool, bool)`

GetCreateDefaultSecurityGroupOk returns a tuple with the CreateDefaultSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDefaultSecurityGroup

`func (o *DatacenterPropertiesPut) SetCreateDefaultSecurityGroup(v bool)`

SetCreateDefaultSecurityGroup sets CreateDefaultSecurityGroup field to given value.

### HasCreateDefaultSecurityGroup

`func (o *DatacenterPropertiesPut) HasCreateDefaultSecurityGroup() bool`

HasCreateDefaultSecurityGroup returns a boolean if a field has been set.



