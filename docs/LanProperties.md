# LanProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name of that resource | [optional] 
**IpFailover** | Pointer to [**[]IPFailover**](IPFailover.md) | IP failover configurations for lan | [optional] 
**Pcc** | Pointer to **string** | Unique identifier of the private cross connect the given LAN is connected to if any | [optional] 
**Public** | Pointer to **bool** | Does this LAN faces the public Internet or not | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


