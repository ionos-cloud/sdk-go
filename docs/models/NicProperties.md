# NicProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DeviceNumber** | Pointer to **int32** | The Logical Unit Number (LUN) of the storage volume. Null if this NIC was created using Cloud API and no DCD changes were performed on the Datacenter. | [optional] [readonly] |
|**Dhcp** | Pointer to **bool** | Indicates if the NIC will reserve an IP using DHCP. | [optional] [default to true]|
|**Dhcpv6** | Pointer to **NullableBool** | Indicates if the NIC will receive an IPv6 using DHCP. It can be set to &#39;true&#39; or &#39;false&#39; only if this NIC is connected to an IPv6 enabled LAN. | [optional] |
|**FirewallActive** | Pointer to **bool** | Activate or deactivate the firewall. By default, an active firewall without any defined rules will block all incoming network traffic except for the firewall rules that explicitly allows certain protocols, IP addresses and ports. | [optional] |
|**FirewallType** | Pointer to **string** | The type of firewall rules that will be allowed on the NIC. If not specified, the default INGRESS value is used. | [optional] |
|**Ips** | Pointer to **[]string** | Collection of IP addresses, assigned to the NIC. Explicitly assigned public IPs need to come from reserved IP blocks. Passing value null or empty array will assign an IP address automatically. | [optional] |
|**Ipv6CidrBlock** | Pointer to **NullableString** | If this NIC is connected to an IPv6 enabled LAN then this property contains the /80 IPv6 CIDR block of the NIC. If you leave this property &#39;null&#39; when adding a NIC to an IPv6-enabled LAN, then an IPv6 CIDR block will automatically be assigned to the NIC, but you can also specify an /80 IPv6 CIDR block for the NIC on your own, which must be inside the /64 IPv6 CIDR block of the LAN and unique. This value can only be set, if the LAN already has an IPv6 CIDR block assigned. An IPv6-enabled LAN is limited to a maximum of 65,536 NICs. | [optional] |
|**Ipv6Ips** | Pointer to **[]string** | If this NIC is connected to an IPv6 enabled LAN then this property contains the IPv6 IP addresses of the NIC. The maximum number of IPv6 IP addresses per NIC is 50, if you need more, contact support. If you leave this property &#39;null&#39; when adding a NIC, when changing the NIC&#39;s IPv6 CIDR block, when changing the LAN&#39;s IPv6 CIDR block or when moving the NIC to a different IPv6 enabled LAN, then we will automatically assign the same number of IPv6 addresses which you had before from the NICs new CIDR block. If you leave this property &#39;null&#39; while not changing the CIDR block, the IPv6 IP addresses won&#39;t be changed either. You can also provide your own self choosen IPv6 addresses, which then must be inside the IPv6 CIDR block of this NIC. | [optional] |
|**Lan** | **int32** | The LAN ID the NIC will be on. If the LAN ID does not exist, it will be implicitly created. | |
|**Mac** | Pointer to **string** | The MAC address of the NIC. | [optional] [readonly] |
|**Name** | Pointer to **string** | The name of the  resource. | [optional] |
|**PciSlot** | Pointer to **int32** | The PCI slot number for the NIC. | [optional] [readonly] |
|**Vnet** | Pointer to **string** | The vnet ID that belongs to this NIC; Requires system privileges | [optional] |

## Methods

### NewNicProperties

`func NewNicProperties(lan int32, ) *NicProperties`

NewNicProperties instantiates a new NicProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNicPropertiesWithDefaults

`func NewNicPropertiesWithDefaults() *NicProperties`

NewNicPropertiesWithDefaults instantiates a new NicProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceNumber

`func (o *NicProperties) GetDeviceNumber() int32`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *NicProperties) GetDeviceNumberOk() (*int32, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNumber

`func (o *NicProperties) SetDeviceNumber(v int32)`

SetDeviceNumber sets DeviceNumber field to given value.

### HasDeviceNumber

`func (o *NicProperties) HasDeviceNumber() bool`

HasDeviceNumber returns a boolean if a field has been set.

### GetDhcp

`func (o *NicProperties) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *NicProperties) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *NicProperties) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *NicProperties) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetDhcpv6

`func (o *NicProperties) GetDhcpv6() bool`

GetDhcpv6 returns the Dhcpv6 field if non-nil, zero value otherwise.

### GetDhcpv6Ok

`func (o *NicProperties) GetDhcpv6Ok() (*bool, bool)`

GetDhcpv6Ok returns a tuple with the Dhcpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6

`func (o *NicProperties) SetDhcpv6(v bool)`

SetDhcpv6 sets Dhcpv6 field to given value.

### HasDhcpv6

`func (o *NicProperties) HasDhcpv6() bool`

HasDhcpv6 returns a boolean if a field has been set.

### SetDhcpv6Nil

`func (o *NicProperties) SetDhcpv6Nil()`

 SetDhcpv6Nil sets the value for Dhcpv6 to be marshalled as an explicit nil
 Alternatively Dhcpv6 can be set directly to the address `&Nilbool`, which is a sentinel value that is checked when marshalling.

### UnsetDhcpv6
`func (o *NicProperties) UnsetDhcpv6()`

### GetFirewallActive

`func (o *NicProperties) GetFirewallActive() bool`

GetFirewallActive returns the FirewallActive field if non-nil, zero value otherwise.

### GetFirewallActiveOk

`func (o *NicProperties) GetFirewallActiveOk() (*bool, bool)`

GetFirewallActiveOk returns a tuple with the FirewallActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallActive

`func (o *NicProperties) SetFirewallActive(v bool)`

SetFirewallActive sets FirewallActive field to given value.

### HasFirewallActive

`func (o *NicProperties) HasFirewallActive() bool`

HasFirewallActive returns a boolean if a field has been set.

### GetFirewallType

`func (o *NicProperties) GetFirewallType() string`

GetFirewallType returns the FirewallType field if non-nil, zero value otherwise.

### GetFirewallTypeOk

`func (o *NicProperties) GetFirewallTypeOk() (*string, bool)`

GetFirewallTypeOk returns a tuple with the FirewallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallType

`func (o *NicProperties) SetFirewallType(v string)`

SetFirewallType sets FirewallType field to given value.

### HasFirewallType

`func (o *NicProperties) HasFirewallType() bool`

HasFirewallType returns a boolean if a field has been set.

### GetIps

`func (o *NicProperties) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *NicProperties) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *NicProperties) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *NicProperties) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetIpv6CidrBlock

`func (o *NicProperties) GetIpv6CidrBlock() string`

GetIpv6CidrBlock returns the Ipv6CidrBlock field if non-nil, zero value otherwise.

### GetIpv6CidrBlockOk

`func (o *NicProperties) GetIpv6CidrBlockOk() (*string, bool)`

GetIpv6CidrBlockOk returns a tuple with the Ipv6CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6CidrBlock

`func (o *NicProperties) SetIpv6CidrBlock(v string)`

SetIpv6CidrBlock sets Ipv6CidrBlock field to given value.

### HasIpv6CidrBlock

`func (o *NicProperties) HasIpv6CidrBlock() bool`

HasIpv6CidrBlock returns a boolean if a field has been set.

### SetIpv6CidrBlockNil

`func (o *NicProperties) SetIpv6CidrBlockNil()`

 SetIpv6CidrBlockNil sets the value for Ipv6CidrBlock to be marshalled as an explicit nil
 Alternatively Ipv6CidrBlock can be set directly to the address `&Nilstring`, which is a sentinel value that is checked when marshalling.

### UnsetIpv6CidrBlock
`func (o *NicProperties) UnsetIpv6CidrBlock()`

### GetIpv6Ips

`func (o *NicProperties) GetIpv6Ips() []string`

GetIpv6Ips returns the Ipv6Ips field if non-nil, zero value otherwise.

### GetIpv6IpsOk

`func (o *NicProperties) GetIpv6IpsOk() (*[]string, bool)`

GetIpv6IpsOk returns a tuple with the Ipv6Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Ips

`func (o *NicProperties) SetIpv6Ips(v []string)`

SetIpv6Ips sets Ipv6Ips field to given value.

### HasIpv6Ips

`func (o *NicProperties) HasIpv6Ips() bool`

HasIpv6Ips returns a boolean if a field has been set.

### GetLan

`func (o *NicProperties) GetLan() int32`

GetLan returns the Lan field if non-nil, zero value otherwise.

### GetLanOk

`func (o *NicProperties) GetLanOk() (*int32, bool)`

GetLanOk returns a tuple with the Lan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan

`func (o *NicProperties) SetLan(v int32)`

SetLan sets Lan field to given value.


### GetMac

`func (o *NicProperties) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *NicProperties) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *NicProperties) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *NicProperties) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *NicProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NicProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NicProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NicProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPciSlot

`func (o *NicProperties) GetPciSlot() int32`

GetPciSlot returns the PciSlot field if non-nil, zero value otherwise.

### GetPciSlotOk

`func (o *NicProperties) GetPciSlotOk() (*int32, bool)`

GetPciSlotOk returns a tuple with the PciSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlot

`func (o *NicProperties) SetPciSlot(v int32)`

SetPciSlot sets PciSlot field to given value.

### HasPciSlot

`func (o *NicProperties) HasPciSlot() bool`

HasPciSlot returns a boolean if a field has been set.

### GetVnet

`func (o *NicProperties) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *NicProperties) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *NicProperties) SetVnet(v string)`

SetVnet sets Vnet field to given value.

### HasVnet

`func (o *NicProperties) HasVnet() bool`

HasVnet returns a boolean if a field has been set.



