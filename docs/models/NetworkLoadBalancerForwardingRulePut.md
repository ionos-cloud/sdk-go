# NetworkLoadBalancerForwardingRulePut

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Properties** | [**NetworkLoadBalancerForwardingRuleProperties**](NetworkLoadBalancerForwardingRuleProperties.md) |  | |

## Methods

### NewNetworkLoadBalancerForwardingRulePut

`func NewNetworkLoadBalancerForwardingRulePut(properties NetworkLoadBalancerForwardingRuleProperties, ) *NetworkLoadBalancerForwardingRulePut`

NewNetworkLoadBalancerForwardingRulePut instantiates a new NetworkLoadBalancerForwardingRulePut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoadBalancerForwardingRulePutWithDefaults

`func NewNetworkLoadBalancerForwardingRulePutWithDefaults() *NetworkLoadBalancerForwardingRulePut`

NewNetworkLoadBalancerForwardingRulePutWithDefaults instantiates a new NetworkLoadBalancerForwardingRulePut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkLoadBalancerForwardingRulePut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkLoadBalancerForwardingRulePut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkLoadBalancerForwardingRulePut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkLoadBalancerForwardingRulePut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NetworkLoadBalancerForwardingRulePut) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkLoadBalancerForwardingRulePut) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkLoadBalancerForwardingRulePut) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkLoadBalancerForwardingRulePut) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *NetworkLoadBalancerForwardingRulePut) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NetworkLoadBalancerForwardingRulePut) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NetworkLoadBalancerForwardingRulePut) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NetworkLoadBalancerForwardingRulePut) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetProperties

`func (o *NetworkLoadBalancerForwardingRulePut) GetProperties() NetworkLoadBalancerForwardingRuleProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NetworkLoadBalancerForwardingRulePut) GetPropertiesOk() (*NetworkLoadBalancerForwardingRuleProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NetworkLoadBalancerForwardingRulePut) SetProperties(v NetworkLoadBalancerForwardingRuleProperties)`

SetProperties sets Properties field to given value.




