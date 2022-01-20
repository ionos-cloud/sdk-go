# NicProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the  resource. | [optional] |
|**Mac** | Pointer to **string** | The MAC address of the NIC. | [optional] [readonly] |
|**Ips** | Pointer to **[]string** | Collection of IP addresses, assigned to the NIC. Explicitly assigned public IPs need to come from reserved IP blocks. Passing value null or empty array will assign an IP address automatically. | [optional] |
|**Dhcp** | Pointer to **bool** | Indicates if the NIC will reserve an IP using DHCP. | [optional] |
|**Lan** | **int32** | The LAN ID the NIC will be on. If the LAN ID does not exist, it will be implicitly created. | |
|**FirewallActive** | Pointer to **bool** | Activate or deactivate the firewall. By default, an active firewall without any defined rules will block all incoming network traffic except for the firewall rules that explicitly allows certain protocols, IP addresses and ports. | [optional] |
|**FirewallType** | Pointer to **string** | The type of firewall rules that will be allowed on the NIC. If not specified, the default INGRESS value is used. | [optional] |
|**DeviceNumber** | Pointer to **int32** | The Logical Unit Number (LUN) of the storage volume. Null if this NIC was created using Cloud API and no DCD changes were performed on the Datacenter. | [optional] [readonly] |
|**PciSlot** | Pointer to **int32** | The PCI slot number for the NIC. | [optional] [readonly] |

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

### SetIpsNil

`func (o *NicProperties) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *NicProperties) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
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



