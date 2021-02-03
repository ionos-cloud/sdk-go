# FirewallruleProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name of that resource | [optional] 
**Protocol** | **string** | The protocol for the rule. Property cannot be modified after creation (disallowed in update requests) | 
**SourceMac** | Pointer to **string** | Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address | [optional] 
**SourceIp** | Pointer to **string** | Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs | [optional] 
**TargetIp** | Pointer to **string** | In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs | [optional] 
**IcmpCode** | Pointer to **int32** | Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes | [optional] 
**IcmpType** | Pointer to **int32** | Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen. Value null allows all types | [optional] 
**PortRangeStart** | Pointer to **int32** | Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd value null to allow all ports | [optional] 
**PortRangeEnd** | Pointer to **int32** | Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports | [optional] 

## Methods

### NewFirewallruleProperties

`func NewFirewallruleProperties(protocol string, ) *FirewallruleProperties`

NewFirewallruleProperties instantiates a new FirewallruleProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallrulePropertiesWithDefaults

`func NewFirewallrulePropertiesWithDefaults() *FirewallruleProperties`

NewFirewallrulePropertiesWithDefaults instantiates a new FirewallruleProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallruleProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallruleProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallruleProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallruleProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *FirewallruleProperties) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallruleProperties) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallruleProperties) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSourceMac

`func (o *FirewallruleProperties) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *FirewallruleProperties) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *FirewallruleProperties) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *FirewallruleProperties) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.

### GetSourceIp

`func (o *FirewallruleProperties) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *FirewallruleProperties) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *FirewallruleProperties) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *FirewallruleProperties) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetTargetIp

`func (o *FirewallruleProperties) GetTargetIp() string`

GetTargetIp returns the TargetIp field if non-nil, zero value otherwise.

### GetTargetIpOk

`func (o *FirewallruleProperties) GetTargetIpOk() (*string, bool)`

GetTargetIpOk returns a tuple with the TargetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIp

`func (o *FirewallruleProperties) SetTargetIp(v string)`

SetTargetIp sets TargetIp field to given value.

### HasTargetIp

`func (o *FirewallruleProperties) HasTargetIp() bool`

HasTargetIp returns a boolean if a field has been set.

### GetIcmpCode

`func (o *FirewallruleProperties) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *FirewallruleProperties) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *FirewallruleProperties) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.

### HasIcmpCode

`func (o *FirewallruleProperties) HasIcmpCode() bool`

HasIcmpCode returns a boolean if a field has been set.

### GetIcmpType

`func (o *FirewallruleProperties) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *FirewallruleProperties) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *FirewallruleProperties) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.

### HasIcmpType

`func (o *FirewallruleProperties) HasIcmpType() bool`

HasIcmpType returns a boolean if a field has been set.

### GetPortRangeStart

`func (o *FirewallruleProperties) GetPortRangeStart() int32`

GetPortRangeStart returns the PortRangeStart field if non-nil, zero value otherwise.

### GetPortRangeStartOk

`func (o *FirewallruleProperties) GetPortRangeStartOk() (*int32, bool)`

GetPortRangeStartOk returns a tuple with the PortRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeStart

`func (o *FirewallruleProperties) SetPortRangeStart(v int32)`

SetPortRangeStart sets PortRangeStart field to given value.

### HasPortRangeStart

`func (o *FirewallruleProperties) HasPortRangeStart() bool`

HasPortRangeStart returns a boolean if a field has been set.

### GetPortRangeEnd

`func (o *FirewallruleProperties) GetPortRangeEnd() int32`

GetPortRangeEnd returns the PortRangeEnd field if non-nil, zero value otherwise.

### GetPortRangeEndOk

`func (o *FirewallruleProperties) GetPortRangeEndOk() (*int32, bool)`

GetPortRangeEndOk returns a tuple with the PortRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeEnd

`func (o *FirewallruleProperties) SetPortRangeEnd(v int32)`

SetPortRangeEnd sets PortRangeEnd field to given value.

### HasPortRangeEnd

`func (o *FirewallruleProperties) HasPortRangeEnd() bool`

HasPortRangeEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


