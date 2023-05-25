# ApplicationLoadBalancerHttpRuleCondition

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Condition** | **string** | The matching rule for the HTTP rule condition attribute; this parameter is mandatory for &#39;HEADER&#39;, &#39;PATH&#39;, &#39;QUERY&#39;, &#39;METHOD&#39;, &#39;HOST&#39;, and &#39;COOKIE&#39; types. It must be &#39;null&#39; if the type is &#39;SOURCE_IP&#39;. | |
|**Key** | Pointer to **string** | The key can only be set when the HTTP rule condition type is &#39;COOKIES&#39;, &#39;HEADER&#39;, or &#39;QUERY&#39;. For the type &#39;PATH&#39;, &#39;METHOD&#39;, &#39;HOST&#39;, or &#39;SOURCE_IP&#39; the value must be &#39;null&#39;. | [optional] |
|**Negate** | Pointer to **bool** | Specifies whether the condition should be negated; the default value is &#39;FALSE&#39;. | [optional] |
|**Type** | **string** | The HTTP rule condition type. | |
|**Value** | Pointer to **string** | This parameter is mandatory for the conditions &#39;CONTAINS&#39;, &#39;EQUALS&#39;, &#39;MATCHES&#39;, &#39;STARTS_WITH&#39;, &#39;ENDS_WITH&#39;, or if the type is &#39;SOURCE_IP&#39;. Specify a valid CIDR. If the condition is &#39;EXISTS&#39;, the value must be &#39;null&#39;. | [optional] |

## Methods

### NewApplicationLoadBalancerHttpRuleCondition

`func NewApplicationLoadBalancerHttpRuleCondition(condition string, type_ string, ) *ApplicationLoadBalancerHttpRuleCondition`

NewApplicationLoadBalancerHttpRuleCondition instantiates a new ApplicationLoadBalancerHttpRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerHttpRuleConditionWithDefaults

`func NewApplicationLoadBalancerHttpRuleConditionWithDefaults() *ApplicationLoadBalancerHttpRuleCondition`

NewApplicationLoadBalancerHttpRuleConditionWithDefaults instantiates a new ApplicationLoadBalancerHttpRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



