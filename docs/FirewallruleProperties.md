# FirewallruleProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource | [optional] 
**Protocol** | **string** | The protocol for the rule. Property cannot be modified after creation (disallowed in update requests) | 
**SourceMac** | **string** | Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address | [optional] 
**SourceIp** | **string** | Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs | [optional] 
**TargetIp** | **string** | In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs | [optional] 
**IcmpCode** | **int32** | Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes | [optional] 
**IcmpType** | **int32** | Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen. Value null allows all types | [optional] 
**PortRangeStart** | **int32** | Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd value null to allow all ports | [optional] 
**PortRangeEnd** | **int32** | Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


