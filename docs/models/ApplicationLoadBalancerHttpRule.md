# ApplicationLoadBalancerHttpRule

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The unique name of the Application Load Balancer HTTP rule. | |
|**Type** | **string** | Type of the HTTP rule. | |
|**TargetGroup** | Pointer to **string** | The ID of the target group; mandatory and only valid for FORWARD actions. | [optional] |
|**DropQuery** | Pointer to **bool** | Default is false; valid only for REDIRECT actions. | [optional] |
|**Location** | Pointer to **string** | The location for redirecting; mandatory and valid only for REDIRECT actions. | [optional] |
|**StatusCode** | Pointer to **int32** | Valid only for REDIRECT and STATIC actions. For REDIRECT actions, default is 301 and possible values are 301, 302, 303, 307, and 308. For STATIC actions, default is 503 and valid range is 200 to 599. | [optional] |
|**ResponseMessage** | Pointer to **string** | The response message of the request; mandatory for STATIC actions. | [optional] |
|**ContentType** | Pointer to **string** | Valid only for STATIC actions. | [optional] |
|**Conditions** | Pointer to [**[]ApplicationLoadBalancerHttpRuleCondition**](ApplicationLoadBalancerHttpRuleCondition.md) | An array of items in the collection.The action is only performed if each and every condition is met; if no conditions are set, the rule will always be performed. | [optional] |

## Methods

### NewApplicationLoadBalancerHttpRule

`func NewApplicationLoadBalancerHttpRule(name string, type_ string, ) *ApplicationLoadBalancerHttpRule`

NewApplicationLoadBalancerHttpRule instantiates a new ApplicationLoadBalancerHttpRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerHttpRuleWithDefaults

`func NewApplicationLoadBalancerHttpRuleWithDefaults() *ApplicationLoadBalancerHttpRule`

NewApplicationLoadBalancerHttpRuleWithDefaults instantiates a new ApplicationLoadBalancerHttpRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationLoadBalancerHttpRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationLoadBalancerHttpRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationLoadBalancerHttpRule) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ApplicationLoadBalancerHttpRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationLoadBalancerHttpRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationLoadBalancerHttpRule) SetType(v string)`

SetType sets Type field to given value.


### GetTargetGroup

`func (o *ApplicationLoadBalancerHttpRule) GetTargetGroup() string`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *ApplicationLoadBalancerHttpRule) GetTargetGroupOk() (*string, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *ApplicationLoadBalancerHttpRule) SetTargetGroup(v string)`

SetTargetGroup sets TargetGroup field to given value.

### HasTargetGroup

`func (o *ApplicationLoadBalancerHttpRule) HasTargetGroup() bool`

HasTargetGroup returns a boolean if a field has been set.

### GetDropQuery

`func (o *ApplicationLoadBalancerHttpRule) GetDropQuery() bool`

GetDropQuery returns the DropQuery field if non-nil, zero value otherwise.

### GetDropQueryOk

`func (o *ApplicationLoadBalancerHttpRule) GetDropQueryOk() (*bool, bool)`

GetDropQueryOk returns a tuple with the DropQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropQuery

`func (o *ApplicationLoadBalancerHttpRule) SetDropQuery(v bool)`

SetDropQuery sets DropQuery field to given value.

### HasDropQuery

`func (o *ApplicationLoadBalancerHttpRule) HasDropQuery() bool`

HasDropQuery returns a boolean if a field has been set.

### GetLocation

`func (o *ApplicationLoadBalancerHttpRule) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApplicationLoadBalancerHttpRule) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApplicationLoadBalancerHttpRule) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ApplicationLoadBalancerHttpRule) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatusCode

`func (o *ApplicationLoadBalancerHttpRule) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ApplicationLoadBalancerHttpRule) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ApplicationLoadBalancerHttpRule) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ApplicationLoadBalancerHttpRule) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetResponseMessage

`func (o *ApplicationLoadBalancerHttpRule) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *ApplicationLoadBalancerHttpRule) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *ApplicationLoadBalancerHttpRule) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.

### HasResponseMessage

`func (o *ApplicationLoadBalancerHttpRule) HasResponseMessage() bool`

HasResponseMessage returns a boolean if a field has been set.

### GetContentType

`func (o *ApplicationLoadBalancerHttpRule) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ApplicationLoadBalancerHttpRule) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ApplicationLoadBalancerHttpRule) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ApplicationLoadBalancerHttpRule) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetConditions

`func (o *ApplicationLoadBalancerHttpRule) GetConditions() []ApplicationLoadBalancerHttpRuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ApplicationLoadBalancerHttpRule) GetConditionsOk() (*[]ApplicationLoadBalancerHttpRuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ApplicationLoadBalancerHttpRule) SetConditions(v []ApplicationLoadBalancerHttpRuleCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ApplicationLoadBalancerHttpRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.



