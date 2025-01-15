# LanProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the  resource. | [optional] |
|**IpFailover** | Pointer to [**[]IPFailover**](IPFailover.md) | IP failover configurations for lan | [optional] |
|**Ipv4CidrBlock** | Pointer to **NullableString** | For public LANs this property is null, for private LANs it contains the private IPv4 CIDR range. This property is a read only property. | [optional] [readonly] |
|**Ipv6CidrBlock** | Pointer to **NullableString** | For a GET request, this value is either &#39;null&#39; or contains the LAN&#39;s /64 IPv6 CIDR block if this LAN is IPv6 enabled. For POST/PUT/PATCH requests, &#39;AUTO&#39; will result in enabling this LAN for IPv6 and automatically assign a /64 IPv6 CIDR block to this LAN and /80 IPv6 CIDR blocks to the NICs and one /128 IPv6 address to each connected NIC. If you choose the IPv6 CIDR block for the LAN on your own, then you must provide a /64 block, which is inside the IPv6 CIDR block of the virtual datacenter and unique inside all LANs from this virtual datacenter. If you enable IPv6 on a LAN with NICs, those NICs will get a /80 IPv6 CIDR block and one IPv6 address assigned to each automatically, unless you specify them explicitly on the LAN and on the NICs. A virtual data center is limited to a maximum of 256 IPv6-enabled LANs. | [optional] |
|**Pcc** | Pointer to **string** | The unique identifier of the Cross Connect the LAN is connected to, if any. It needs to be ensured that IP addresses of the NICs of all LANs connected to a given Cross Connect is not duplicated and belongs to the same subnet range. | [optional] |
|**Public** | Pointer to **bool** | Indicates if the LAN is connected to the internet or not. | [optional] |

## Methods

### NewLanProperties

`func NewLanProperties() *LanProperties`

NewLanProperties instantiates a new LanProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanPropertiesWithDefaults

`func NewLanPropertiesWithDefaults() *LanProperties`

NewLanPropertiesWithDefaults instantiates a new LanProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LanProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LanProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIpFailover

`func (o *LanProperties) GetIpFailover() []IPFailover`

GetIpFailover returns the IpFailover field if non-nil, zero value otherwise.

### GetIpFailoverOk

`func (o *LanProperties) GetIpFailoverOk() (*[]IPFailover, bool)`

GetIpFailoverOk returns a tuple with the IpFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFailover

`func (o *LanProperties) SetIpFailover(v []IPFailover)`

SetIpFailover sets IpFailover field to given value.

### HasIpFailover

`func (o *LanProperties) HasIpFailover() bool`

HasIpFailover returns a boolean if a field has been set.

### GetIpv4CidrBlock

`func (o *LanProperties) GetIpv4CidrBlock() string`

GetIpv4CidrBlock returns the Ipv4CidrBlock field if non-nil, zero value otherwise.

### GetIpv4CidrBlockOk

`func (o *LanProperties) GetIpv4CidrBlockOk() (*string, bool)`

GetIpv4CidrBlockOk returns a tuple with the Ipv4CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4CidrBlock

`func (o *LanProperties) SetIpv4CidrBlock(v string)`

SetIpv4CidrBlock sets Ipv4CidrBlock field to given value.

### HasIpv4CidrBlock

`func (o *LanProperties) HasIpv4CidrBlock() bool`

HasIpv4CidrBlock returns a boolean if a field has been set.

### SetIpv4CidrBlockNil

`func (o *LanProperties) SetIpv4CidrBlockNil()`

 SetIpv4CidrBlockNil sets the value for Ipv4CidrBlock to be marshalled as an explicit nil
 Alternatively Ipv4CidrBlock can be set directly to the address `&Nilstring`, which is a sentinel value that is checked when marshalling.

### UnsetIpv4CidrBlock
`func (o *LanProperties) UnsetIpv4CidrBlock()`

### GetIpv6CidrBlock

`func (o *LanProperties) GetIpv6CidrBlock() string`

GetIpv6CidrBlock returns the Ipv6CidrBlock field if non-nil, zero value otherwise.

### GetIpv6CidrBlockOk

`func (o *LanProperties) GetIpv6CidrBlockOk() (*string, bool)`

GetIpv6CidrBlockOk returns a tuple with the Ipv6CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6CidrBlock

`func (o *LanProperties) SetIpv6CidrBlock(v string)`

SetIpv6CidrBlock sets Ipv6CidrBlock field to given value.

### HasIpv6CidrBlock

`func (o *LanProperties) HasIpv6CidrBlock() bool`

HasIpv6CidrBlock returns a boolean if a field has been set.

### SetIpv6CidrBlockNil

`func (o *LanProperties) SetIpv6CidrBlockNil()`

 SetIpv6CidrBlockNil sets the value for Ipv6CidrBlock to be marshalled as an explicit nil
 Alternatively Ipv6CidrBlock can be set directly to the address `&Nilstring`, which is a sentinel value that is checked when marshalling.

### UnsetIpv6CidrBlock
`func (o *LanProperties) UnsetIpv6CidrBlock()`

### GetPcc

`func (o *LanProperties) GetPcc() string`

GetPcc returns the Pcc field if non-nil, zero value otherwise.

### GetPccOk

`func (o *LanProperties) GetPccOk() (*string, bool)`

GetPccOk returns a tuple with the Pcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcc

`func (o *LanProperties) SetPcc(v string)`

SetPcc sets Pcc field to given value.

### HasPcc

`func (o *LanProperties) HasPcc() bool`

HasPcc returns a boolean if a field has been set.

### GetPublic

`func (o *LanProperties) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *LanProperties) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *LanProperties) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *LanProperties) HasPublic() bool`

HasPublic returns a boolean if a field has been set.



