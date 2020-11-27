# LoadbalancerProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name of that resource | [optional] 
**Ip** | Pointer to **string** | IPv4 address of the loadbalancer. All attached NICs will inherit this IP. Leaving value null will assign IP automatically | [optional] 
**Dhcp** | Pointer to **bool** | Indicates if the loadbalancer will reserve an IP using DHCP | [optional] 

## Methods

### NewLoadbalancerProperties

`func NewLoadbalancerProperties() *LoadbalancerProperties`

NewLoadbalancerProperties instantiates a new LoadbalancerProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerPropertiesWithDefaults

`func NewLoadbalancerPropertiesWithDefaults() *LoadbalancerProperties`

NewLoadbalancerPropertiesWithDefaults instantiates a new LoadbalancerProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoadbalancerProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadbalancerProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadbalancerProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadbalancerProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *LoadbalancerProperties) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LoadbalancerProperties) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LoadbalancerProperties) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LoadbalancerProperties) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetDhcp

`func (o *LoadbalancerProperties) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *LoadbalancerProperties) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *LoadbalancerProperties) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *LoadbalancerProperties) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


