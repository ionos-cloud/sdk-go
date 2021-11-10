# LanPropertiesPost

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the  resource. | [optional] |
|**IpFailover** | Pointer to [**[]IPFailover**](IPFailover.md) | IP failover configurations for lan | [optional] |
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



