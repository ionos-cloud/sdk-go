# NatGatewayRuleProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the NAT Gateway rule. | |
|**Type** | Pointer to [**NatGatewayRuleType**](NatGatewayRuleType.md) | Type of the NAT Gateway rule. | [optional] |
|**Protocol** | Pointer to [**NatGatewayRuleProtocol**](NatGatewayRuleProtocol.md) | Protocol of the NAT Gateway rule. Defaults to ALL. If protocol is &#39;ICMP&#39; then targetPortRange start and end cannot be set. | [optional] |
|**SourceSubnet** | **string** | Source subnet of the NAT Gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address. | |
|**PublicIp** | **string** | Public IP address of the NAT Gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT Gateway resource | |
|**TargetSubnet** | Pointer to **string** | Target or destination subnet of the NAT Gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address. | [optional] |
|**TargetPortRange** | Pointer to [**TargetPortRange**](TargetPortRange.md) |  | [optional] |

## Methods

### NewNatGatewayRuleProperties

`func NewNatGatewayRuleProperties(name string, sourceSubnet string, publicIp string, ) *NatGatewayRuleProperties`

NewNatGatewayRuleProperties instantiates a new NatGatewayRuleProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatGatewayRulePropertiesWithDefaults

`func NewNatGatewayRulePropertiesWithDefaults() *NatGatewayRuleProperties`

NewNatGatewayRulePropertiesWithDefaults instantiates a new NatGatewayRuleProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NatGatewayRuleProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatGatewayRuleProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatGatewayRuleProperties) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *NatGatewayRuleProperties) GetType() NatGatewayRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NatGatewayRuleProperties) GetTypeOk() (*NatGatewayRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NatGatewayRuleProperties) SetType(v NatGatewayRuleType)`

SetType sets Type field to given value.

### HasType

`func (o *NatGatewayRuleProperties) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProtocol

`func (o *NatGatewayRuleProperties) GetProtocol() NatGatewayRuleProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NatGatewayRuleProperties) GetProtocolOk() (*NatGatewayRuleProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NatGatewayRuleProperties) SetProtocol(v NatGatewayRuleProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NatGatewayRuleProperties) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSourceSubnet

`func (o *NatGatewayRuleProperties) GetSourceSubnet() string`

GetSourceSubnet returns the SourceSubnet field if non-nil, zero value otherwise.

### GetSourceSubnetOk

`func (o *NatGatewayRuleProperties) GetSourceSubnetOk() (*string, bool)`

GetSourceSubnetOk returns a tuple with the SourceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSubnet

`func (o *NatGatewayRuleProperties) SetSourceSubnet(v string)`

SetSourceSubnet sets SourceSubnet field to given value.


### GetPublicIp

`func (o *NatGatewayRuleProperties) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *NatGatewayRuleProperties) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *NatGatewayRuleProperties) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### GetTargetSubnet

`func (o *NatGatewayRuleProperties) GetTargetSubnet() string`

GetTargetSubnet returns the TargetSubnet field if non-nil, zero value otherwise.

### GetTargetSubnetOk

`func (o *NatGatewayRuleProperties) GetTargetSubnetOk() (*string, bool)`

GetTargetSubnetOk returns a tuple with the TargetSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSubnet

`func (o *NatGatewayRuleProperties) SetTargetSubnet(v string)`

SetTargetSubnet sets TargetSubnet field to given value.

### HasTargetSubnet

`func (o *NatGatewayRuleProperties) HasTargetSubnet() bool`

HasTargetSubnet returns a boolean if a field has been set.

### GetTargetPortRange

`func (o *NatGatewayRuleProperties) GetTargetPortRange() TargetPortRange`

GetTargetPortRange returns the TargetPortRange field if non-nil, zero value otherwise.

### GetTargetPortRangeOk

`func (o *NatGatewayRuleProperties) GetTargetPortRangeOk() (*TargetPortRange, bool)`

GetTargetPortRangeOk returns a tuple with the TargetPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPortRange

`func (o *NatGatewayRuleProperties) SetTargetPortRange(v TargetPortRange)`

SetTargetPortRange sets TargetPortRange field to given value.

### HasTargetPortRange

`func (o *NatGatewayRuleProperties) HasTargetPortRange() bool`

HasTargetPortRange returns a boolean if a field has been set.



