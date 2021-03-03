# NicEntities

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Flowlogs** | Pointer to [**FlowLogs**](FlowLogs.md) |  | [optional] |
|**Firewallrules** | Pointer to [**FirewallRules**](FirewallRules.md) |  | [optional] |

## Methods

### NewNicEntities

`func NewNicEntities() *NicEntities`

NewNicEntities instantiates a new NicEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNicEntitiesWithDefaults

`func NewNicEntitiesWithDefaults() *NicEntities`

NewNicEntitiesWithDefaults instantiates a new NicEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowlogs

`func (o *NicEntities) GetFlowlogs() FlowLogs`

GetFlowlogs returns the Flowlogs field if non-nil, zero value otherwise.

### GetFlowlogsOk

`func (o *NicEntities) GetFlowlogsOk() (*FlowLogs, bool)`

GetFlowlogsOk returns a tuple with the Flowlogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowlogs

`func (o *NicEntities) SetFlowlogs(v FlowLogs)`

SetFlowlogs sets Flowlogs field to given value.

### HasFlowlogs

`func (o *NicEntities) HasFlowlogs() bool`

HasFlowlogs returns a boolean if a field has been set.

### GetFirewallrules

`func (o *NicEntities) GetFirewallrules() FirewallRules`

GetFirewallrules returns the Firewallrules field if non-nil, zero value otherwise.

### GetFirewallrulesOk

`func (o *NicEntities) GetFirewallrulesOk() (*FirewallRules, bool)`

GetFirewallrulesOk returns a tuple with the Firewallrules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallrules

`func (o *NicEntities) SetFirewallrules(v FirewallRules)`

SetFirewallrules sets Firewallrules field to given value.

### HasFirewallrules

`func (o *NicEntities) HasFirewallrules() bool`

HasFirewallrules returns a boolean if a field has been set.



