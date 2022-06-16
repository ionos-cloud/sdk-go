# ApplicationLoadBalancerHttpRuleCondition

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | **string** | Type of the HTTP rule condition. | |
|**Condition** | **string** | Matching rule for the HTTP rule condition attribute; mandatory for HEADER, PATH, QUERY, METHOD, HOST, and COOKIE types; must be null when type is SOURCE_IP. | |
|**Negate** | Pointer to **bool** | Specifies whether the condition is negated or not; the default is False. | [optional] |
|**Key** | Pointer to **string** | Must be null when type is PATH, METHOD, HOST, or SOURCE_IP. Key can only be set when type is COOKIES, HEADER, or QUERY. | [optional] |
|**Value** | Pointer to **string** | Mandatory for conditions CONTAINS, EQUALS, MATCHES, STARTS_WITH, ENDS_WITH; must be null when condition is EXISTS; should be a valid CIDR if provided and if type is SOURCE_IP. | [optional] |

## Methods

### NewApplicationLoadBalancerHttpRuleCondition

`func NewApplicationLoadBalancerHttpRuleCondition(type_ string, condition string, ) *ApplicationLoadBalancerHttpRuleCondition`

NewApplicationLoadBalancerHttpRuleCondition instantiates a new ApplicationLoadBalancerHttpRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerHttpRuleConditionWithDefaults

`func NewApplicationLoadBalancerHttpRuleConditionWithDefaults() *ApplicationLoadBalancerHttpRuleCondition`

NewApplicationLoadBalancerHttpRuleConditionWithDefaults instantiates a new ApplicationLoadBalancerHttpRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationLoadBalancerHttpRuleCondition) SetType(v string)`

SetType sets Type field to given value.


### GetCondition

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ApplicationLoadBalancerHttpRuleCondition) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetNegate

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *ApplicationLoadBalancerHttpRuleCondition) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *ApplicationLoadBalancerHttpRuleCondition) HasNegate() bool`

HasNegate returns a boolean if a field has been set.

### GetKey

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationLoadBalancerHttpRuleCondition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApplicationLoadBalancerHttpRuleCondition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationLoadBalancerHttpRuleCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationLoadBalancerHttpRuleCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApplicationLoadBalancerHttpRuleCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.



