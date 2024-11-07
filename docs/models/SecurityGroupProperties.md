# SecurityGroupProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the security group. | |
|**Description** | Pointer to **string** | The description of the security group. | [optional] |

## Methods

### NewSecurityGroupProperties

`func NewSecurityGroupProperties(name string, ) *SecurityGroupProperties`

NewSecurityGroupProperties instantiates a new SecurityGroupProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupPropertiesWithDefaults

`func NewSecurityGroupPropertiesWithDefaults() *SecurityGroupProperties`

NewSecurityGroupPropertiesWithDefaults instantiates a new SecurityGroupProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityGroupProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupProperties) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SecurityGroupProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityGroupProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.



