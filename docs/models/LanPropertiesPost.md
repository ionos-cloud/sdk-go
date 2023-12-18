# LanPropertiesPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**IpFailover** | Pointer to [**[]IPFailover**](IPFailover.md) | IP failover configurations for lan | [optional] |
|**Ipv6CidrBlock** | Pointer to **NullableString** | For a GET request, this value is either &#39;null&#39; or contains the LAN&#39;s /64 IPv6 CIDR block if this LAN is IPv6-enabled. For POST/PUT/PATCH requests, &#39;AUTO&#39; will result in enabling this LAN for IPv6 and automatically assign a /64 IPv6 CIDR block to this LAN. If you choose the IPv6 CIDR block on your own, then you must provide a /64 block, which is inside the IPv6 CIDR block of the virtual datacenter and unique inside all LANs from this virtual datacenter. If you enable IPv6 on a LAN with NICs, those NICs will get a /80 IPv6 CIDR block and one IPv6 address assigned to each automatically, unless you specify them explicitly on the NICs. A virtual data center is limited to a maximum of 256 IPv6-enabled LANs. | [optional] |
|**Name** | Pointer to **string** | The name of the  resource. | [optional] |
|**Pcc** | Pointer to **string** | The unique identifier of the private Cross-Connect the LAN is connected to, if any. | [optional] |
|**Public** | Pointer to **bool** | This LAN faces the public Internet. | [optional] |

## Methods

### NewLanPropertiesPost

`func NewLanPropertiesPost() *LanPropertiesPost`

NewLanPropertiesPost instantiates a new LanPropertiesPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanPropertiesPostWithDefaults

`func NewLanPropertiesPostWithDefaults() *LanPropertiesPost`

NewLanPropertiesPostWithDefaults instantiates a new LanPropertiesPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpFailover

`func (o *LanPropertiesPost) GetIpFailover() []IPFailover`

GetIpFailover returns the IpFailover field if non-nil, zero value otherwise.

### GetIpFailoverOk

`func (o *LanPropertiesPost) GetIpFailoverOk() (*[]IPFailover, bool)`

GetIpFailoverOk returns a tuple with the IpFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFailover

`func (o *LanPropertiesPost) SetIpFailover(v []IPFailover)`

SetIpFailover sets IpFailover field to given value.

### HasIpFailover

`func (o *LanPropertiesPost) HasIpFailover() bool`

HasIpFailover returns a boolean if a field has been set.

### GetIpv6CidrBlock

`func (o *LanPropertiesPost) GetIpv6CidrBlock() string`

GetIpv6CidrBlock returns the Ipv6CidrBlock field if non-nil, zero value otherwise.

### GetIpv6CidrBlockOk

`func (o *LanPropertiesPost) GetIpv6CidrBlockOk() (*string, bool)`

GetIpv6CidrBlockOk returns a tuple with the Ipv6CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6CidrBlock

`func (o *LanPropertiesPost) SetIpv6CidrBlock(v string)`

SetIpv6CidrBlock sets Ipv6CidrBlock field to given value.

### HasIpv6CidrBlock

`func (o *LanPropertiesPost) HasIpv6CidrBlock() bool`

HasIpv6CidrBlock returns a boolean if a field has been set.

### SetIpv6CidrBlockNil

`func (o *LanPropertiesPost) SetIpv6CidrBlockNil()`

 SetIpv6CidrBlockNil sets the value for Ipv6CidrBlock to be marshalled as an explicit nil
 Alternatively Ipv6CidrBlock can be set directly to the address `&Nilstring`, which is a sentinel value that is checked when marshalling.

### UnsetIpv6CidrBlock
`func (o *LanPropertiesPost) UnsetIpv6CidrBlock()`

### GetName

`func (o *LanPropertiesPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanPropertiesPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanPropertiesPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LanPropertiesPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcc

`func (o *LanPropertiesPost) GetPcc() string`

GetPcc returns the Pcc field if non-nil, zero value otherwise.

### GetPccOk

`func (o *LanPropertiesPost) GetPccOk() (*string, bool)`

GetPccOk returns a tuple with the Pcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcc

`func (o *LanPropertiesPost) SetPcc(v string)`

SetPcc sets Pcc field to given value.

### HasPcc

`func (o *LanPropertiesPost) HasPcc() bool`

HasPcc returns a boolean if a field has been set.

### GetPublic

`func (o *LanPropertiesPost) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *LanPropertiesPost) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *LanPropertiesPost) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *LanPropertiesPost) HasPublic() bool`

HasPublic returns a boolean if a field has been set.



