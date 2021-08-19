# NatGatewayEntities

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Rules** | Pointer to [**NatGatewayRules**](NatGatewayRules.md) |  | [optional] |
|**Flowlogs** | Pointer to [**FlowLogs**](FlowLogs.md) |  | [optional] |

## Methods

### NewNatGatewayEntities

`func NewNatGatewayEntities() *NatGatewayEntities`

NewNatGatewayEntities instantiates a new NatGatewayEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatGatewayEntitiesWithDefaults

`func NewNatGatewayEntitiesWithDefaults() *NatGatewayEntities`

NewNatGatewayEntitiesWithDefaults instantiates a new NatGatewayEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *NatGatewayEntities) GetRules() NatGatewayRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NatGatewayEntities) GetRulesOk() (*NatGatewayRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NatGatewayEntities) SetRules(v NatGatewayRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *NatGatewayEntities) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetFlowlogs

`func (o *NatGatewayEntities) GetFlowlogs() FlowLogs`

GetFlowlogs returns the Flowlogs field if non-nil, zero value otherwise.

### GetFlowlogsOk

`func (o *NatGatewayEntities) GetFlowlogsOk() (*FlowLogs, bool)`

GetFlowlogsOk returns a tuple with the Flowlogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowlogs

`func (o *NatGatewayEntities) SetFlowlogs(v FlowLogs)`

SetFlowlogs sets Flowlogs field to given value.

### HasFlowlogs

`func (o *NatGatewayEntities) HasFlowlogs() bool`

HasFlowlogs returns a boolean if a field has been set.



