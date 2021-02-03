# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**Type**](Type.md) | The type of the resource | [optional] 
**Properties** | [**ContractProperties**](ContractProperties.md) |  | 

## Methods

### NewContract

`func NewContract(properties ContractProperties, ) *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Contract) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Contract) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Contract) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Contract) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProperties

`func (o *Contract) GetProperties() ContractProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Contract) GetPropertiesOk() (*ContractProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Contract) SetProperties(v ContractProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


