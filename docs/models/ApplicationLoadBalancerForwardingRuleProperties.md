# ApplicationLoadBalancerForwardingRuleProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ClientTimeout** | Pointer to **int32** | The maximum time in milliseconds to wait for the client to acknowledge or send data; default is 50,000 (50 seconds). | [optional] |
|**HttpRules** | Pointer to [**[]ApplicationLoadBalancerHttpRule**](ApplicationLoadBalancerHttpRule.md) | An array of items in the collection. The original order of rules is preserved during processing, except that rules of the &#39;FORWARD&#39; type are processed after the rules with other defined actions. The relative order of the &#39;FORWARD&#39; type rules is also preserved during the processing. | [optional] |
|**ListenerIp** | **string** | The listening (inbound) IP. | |
|**ListenerPort** | **int32** | The listening (inbound) port number; the valid range is 1 to 65535. | |
|**Name** | **string** | The name of the Application Load Balancer forwarding rule. | |
|**Protocol** | **string** | The balancing protocol. | |
|**ServerCertificates** | Pointer to **[]string** | Array of items in the collection. | [optional] |

## Methods

### NewApplicationLoadBalancerForwardingRuleProperties

`func NewApplicationLoadBalancerForwardingRuleProperties(listenerIp string, listenerPort int32, name string, protocol string, ) *ApplicationLoadBalancerForwardingRuleProperties`

NewApplicationLoadBalancerForwardingRuleProperties instantiates a new ApplicationLoadBalancerForwardingRuleProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLoadBalancerForwardingRulePropertiesWithDefaults

`func NewApplicationLoadBalancerForwardingRulePropertiesWithDefaults() *ApplicationLoadBalancerForwardingRuleProperties`

NewApplicationLoadBalancerForwardingRulePropertiesWithDefaults instantiates a new ApplicationLoadBalancerForwardingRuleProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientTimeout

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetClientTimeout() int32`

GetClientTimeout returns the ClientTimeout field if non-nil, zero value otherwise.

### GetClientTimeoutOk

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetClientTimeoutOk() (*int32, bool)`

GetClientTimeoutOk returns a tuple with the ClientTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTimeout

`func (o *ApplicationLoadBalancerForwardingRuleProperties) SetClientTimeout(v int32)`

SetClientTimeout sets ClientTimeout field to given value.

### HasClientTimeout

`func (o *ApplicationLoadBalancerForwardingRuleProperties) HasClientTimeout() bool`

HasClientTimeout returns a boolean if a field has been set.

### GetHttpRules

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetHttpRules() []ApplicationLoadBalancerHttpRule`

GetHttpRules returns the HttpRules field if non-nil, zero value otherwise.

### GetHttpRulesOk

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetHttpRulesOk() (*[]ApplicationLoadBalancerHttpRule, bool)`

GetHttpRulesOk returns a tuple with the HttpRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRules

`func (o *ApplicationLoadBalancerForwardingRuleProperties) SetHttpRules(v []ApplicationLoadBalancerHttpRule)`

SetHttpRules sets HttpRules field to given value.

### HasHttpRules

`func (o *ApplicationLoadBalancerForwardingRuleProperties) HasHttpRules() bool`

HasHttpRules returns a boolean if a field has been set.

### GetListenerIp

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetListenerIp() string`

GetListenerIp returns the ListenerIp field if non-nil, zero value otherwise.

### GetListenerIpOk

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetListenerIpOk() (*string, bool)`

GetListenerIpOk returns a tuple with the ListenerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIp

`func (o *ApplicationLoadBalancerForwardingRuleProperties) SetListenerIp(v string)`

SetListenerIp sets ListenerIp field to given value.


### GetListenerPort

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetListenerPort() int32`

GetListenerPort returns the ListenerPort field if non-nil, zero value otherwise.

### GetListenerPortOk

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetListenerPortOk() (*int32, bool)`

GetListenerPortOk returns a tuple with the ListenerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerPort

`func (o *ApplicationLoadBalancerForwardingRuleProperties) SetListenerPort(v int32)`

SetListenerPort sets ListenerPort field to given value.


### GetName

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationLoadBalancerForwardingRuleProperties) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplicationLoadBalancerForwardingRuleProperties) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetServerCertificates

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetServerCertificates() []string`

GetServerCertificates returns the ServerCertificates field if non-nil, zero value otherwise.

### GetServerCertificatesOk

`func (o *ApplicationLoadBalancerForwardingRuleProperties) GetServerCertificatesOk() (*[]string, bool)`

GetServerCertificatesOk returns a tuple with the ServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificates

`func (o *ApplicationLoadBalancerForwardingRuleProperties) SetServerCertificates(v []string)`

SetServerCertificates sets ServerCertificates field to given value.

### HasServerCertificates

`func (o *ApplicationLoadBalancerForwardingRuleProperties) HasServerCertificates() bool`

HasServerCertificates returns a boolean if a field has been set.



