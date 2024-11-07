# SecurityGroupEntities

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Rules** | Pointer to [**FirewallRules**](FirewallRules.md) |  | [optional] |
|**Nics** | Pointer to [**Nics**](Nics.md) |  | [optional] |
|**Servers** | Pointer to [**Servers**](Servers.md) |  | [optional] |

## Methods

### NewSecurityGroupEntities

`func NewSecurityGroupEntities() *SecurityGroupEntities`

NewSecurityGroupEntities instantiates a new SecurityGroupEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupEntitiesWithDefaults

`func NewSecurityGroupEntitiesWithDefaults() *SecurityGroupEntities`

NewSecurityGroupEntitiesWithDefaults instantiates a new SecurityGroupEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *SecurityGroupEntities) GetRules() FirewallRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroupEntities) GetRulesOk() (*FirewallRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroupEntities) SetRules(v FirewallRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityGroupEntities) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetNics

`func (o *SecurityGroupEntities) GetNics() Nics`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *SecurityGroupEntities) GetNicsOk() (*Nics, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *SecurityGroupEntities) SetNics(v Nics)`

SetNics sets Nics field to given value.

### HasNics

`func (o *SecurityGroupEntities) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetServers

`func (o *SecurityGroupEntities) GetServers() Servers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *SecurityGroupEntities) GetServersOk() (*Servers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *SecurityGroupEntities) SetServers(v Servers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *SecurityGroupEntities) HasServers() bool`

HasServers returns a boolean if a field has been set.



