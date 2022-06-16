# ApplicationLoadBalancerEntities

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Forwardingrules** | Pointer to [**ApplicationLoadBalancerForwardingRules**](ApplicationLoadBalancerForwardingRules.md) |  | [optional] |

## Methods

### NewApplicationLoadBalancerEntities

`func NewApplicationLoadBalancerEntities() *ApplicationLoadBalancerEntities`

NewApplicationLoadBalancerEntities instantiates a new ApplicationLoadBalancerEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerEntitiesWithDefaults

`func NewApplicationLoadBalancerEntitiesWithDefaults() *ApplicationLoadBalancerEntities`

NewApplicationLoadBalancerEntitiesWithDefaults instantiates a new ApplicationLoadBalancerEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardingrules

`func (o *ApplicationLoadBalancerEntities) GetForwardingrules() ApplicationLoadBalancerForwardingRules`

GetForwardingrules returns the Forwardingrules field if non-nil, zero value otherwise.

### GetForwardingrulesOk

`func (o *ApplicationLoadBalancerEntities) GetForwardingrulesOk() (*ApplicationLoadBalancerForwardingRules, bool)`

GetForwardingrulesOk returns a tuple with the Forwardingrules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingrules

`func (o *ApplicationLoadBalancerEntities) SetForwardingrules(v ApplicationLoadBalancerForwardingRules)`

SetForwardingrules sets Forwardingrules field to given value.

### HasForwardingrules

`func (o *ApplicationLoadBalancerEntities) HasForwardingrules() bool`

HasForwardingrules returns a boolean if a field has been set.



