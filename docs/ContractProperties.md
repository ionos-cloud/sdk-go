# ContractProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractNumber** | Pointer to **int64** | contract number | [optional] [readonly] 
**Owner** | Pointer to **string** | owner of the contract | [optional] [readonly] 
**Status** | Pointer to **string** | status of the contract | [optional] [readonly] 
**RegDomain** | Pointer to **string** | Registration domain of the contract | [optional] [readonly] 
**ResourceLimits** | Pointer to [**ResourceLimits**](ResourceLimits.md) |  | [optional] 

## Methods

### NewContractProperties

`func NewContractProperties() *ContractProperties`

NewContractProperties instantiates a new ContractProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractPropertiesWithDefaults

`func NewContractPropertiesWithDefaults() *ContractProperties`

NewContractPropertiesWithDefaults instantiates a new ContractProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractNumber

`func (o *ContractProperties) GetContractNumber() int64`

GetContractNumber returns the ContractNumber field if non-nil, zero value otherwise.

### GetContractNumberOk

`func (o *ContractProperties) GetContractNumberOk() (*int64, bool)`

GetContractNumberOk returns a tuple with the ContractNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractNumber

`func (o *ContractProperties) SetContractNumber(v int64)`

SetContractNumber sets ContractNumber field to given value.

### HasContractNumber

`func (o *ContractProperties) HasContractNumber() bool`

HasContractNumber returns a boolean if a field has been set.

### GetOwner

`func (o *ContractProperties) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ContractProperties) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ContractProperties) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ContractProperties) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetStatus

`func (o *ContractProperties) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractProperties) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractProperties) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContractProperties) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegDomain

`func (o *ContractProperties) GetRegDomain() string`

GetRegDomain returns the RegDomain field if non-nil, zero value otherwise.

### GetRegDomainOk

`func (o *ContractProperties) GetRegDomainOk() (*string, bool)`

GetRegDomainOk returns a tuple with the RegDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegDomain

`func (o *ContractProperties) SetRegDomain(v string)`

SetRegDomain sets RegDomain field to given value.

### HasRegDomain

`func (o *ContractProperties) HasRegDomain() bool`

HasRegDomain returns a boolean if a field has been set.

### GetResourceLimits

`func (o *ContractProperties) GetResourceLimits() ResourceLimits`

GetResourceLimits returns the ResourceLimits field if non-nil, zero value otherwise.

### GetResourceLimitsOk

`func (o *ContractProperties) GetResourceLimitsOk() (*ResourceLimits, bool)`

GetResourceLimitsOk returns a tuple with the ResourceLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLimits

`func (o *ContractProperties) SetResourceLimits(v ResourceLimits)`

SetResourceLimits sets ResourceLimits field to given value.

### HasResourceLimits

`func (o *ContractProperties) HasResourceLimits() bool`

HasResourceLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


