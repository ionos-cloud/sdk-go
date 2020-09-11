# NicProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name of that resource | [optional] 
**Mac** | **string** | The MAC address of the NIC | [optional] [readonly] 
**Ips** | **[]string** | Collection of IP addresses assigned to a nic. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically. | [optional] 
**Dhcp** | **bool** | Indicates if the nic will reserve an IP using DHCP | [optional] 
**Lan** | **int32** | The LAN ID the NIC will sit on. If the LAN ID does not exist it will be implicitly created | 
**FirewallActive** | **bool** | Activate or deactivate the firewall. By default an active firewall without any defined rules will block all incoming network traffic except for the firewall rules that explicitly allows certain protocols, ip addresses and ports. | [optional] 
**Nat** | **bool** | Indicates if NAT is enabled on this NIC | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


