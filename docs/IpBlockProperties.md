# IpBlockProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ips** | Pointer to **[]string** | A collection of IPs associated with the IP Block | [optional] [readonly] 
**Location** | **string** | Location of that IP Block. Property cannot be modified after creation (disallowed in update requests) | 
**Size** | **int32** | The size of the IP block | 
**Name** | Pointer to **string** | A name of that resource | [optional] 
**IpConsumers** | Pointer to [**[]IpConsumer**](IpConsumer.md) | Read-Only attribute. Lists consumption detail of an individual ip | [optional] [readonly] 

## Methods

### NewIpBlockProperties

`func NewIpBlockProperties(location string, size int32, ) *IpBlockProperties`

NewIpBlockProperties instantiates a new IpBlockProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockPropertiesWithDefaults

`func NewIpBlockPropertiesWithDefaults() *IpBlockProperties`

NewIpBlockPropertiesWithDefaults instantiates a new IpBlockProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIps

`func (o *IpBlockProperties) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *IpBlockProperties) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *IpBlockProperties) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *IpBlockProperties) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetLocation

`func (o *IpBlockProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IpBlockProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IpBlockProperties) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetSize

`func (o *IpBlockProperties) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *IpBlockProperties) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *IpBlockProperties) SetSize(v int32)`

SetSize sets Size field to given value.


### GetName

`func (o *IpBlockProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpBlockProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpBlockProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpBlockProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIpConsumers

`func (o *IpBlockProperties) GetIpConsumers() []IpConsumer`

GetIpConsumers returns the IpConsumers field if non-nil, zero value otherwise.

### GetIpConsumersOk

`func (o *IpBlockProperties) GetIpConsumersOk() (*[]IpConsumer, bool)`

GetIpConsumersOk returns a tuple with the IpConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConsumers

`func (o *IpBlockProperties) SetIpConsumers(v []IpConsumer)`

SetIpConsumers sets IpConsumers field to given value.

### HasIpConsumers

`func (o *IpBlockProperties) HasIpConsumers() bool`

HasIpConsumers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


