# TargetPortRange

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Start** | Pointer to **int32** | Target port range start associated with the NAT Gateway rule. | [optional] |
|**End** | Pointer to **int32** | Target port range end associated with the NAT Gateway rule. | [optional] |

## Methods

### NewTargetPortRange

`func NewTargetPortRange() *TargetPortRange`

NewTargetPortRange instantiates a new TargetPortRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetPortRangeWithDefaults

`func NewTargetPortRangeWithDefaults() *TargetPortRange`

NewTargetPortRangeWithDefaults instantiates a new TargetPortRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *TargetPortRange) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TargetPortRange) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TargetPortRange) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *TargetPortRange) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *TargetPortRange) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TargetPortRange) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TargetPortRange) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TargetPortRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.



