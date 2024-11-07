# SecurityGroupEntitiesRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Rules** | Pointer to [**FirewallRules**](FirewallRules.md) |  | [optional] |

## Methods

### NewSecurityGroupEntitiesRequest

`func NewSecurityGroupEntitiesRequest() *SecurityGroupEntitiesRequest`

NewSecurityGroupEntitiesRequest instantiates a new SecurityGroupEntitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupEntitiesRequestWithDefaults

`func NewSecurityGroupEntitiesRequestWithDefaults() *SecurityGroupEntitiesRequest`

NewSecurityGroupEntitiesRequestWithDefaults instantiates a new SecurityGroupEntitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *SecurityGroupEntitiesRequest) GetRules() FirewallRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroupEntitiesRequest) GetRulesOk() (*FirewallRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroupEntitiesRequest) SetRules(v FirewallRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityGroupEntitiesRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.



