# TargetGroupTarget

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Ip** | **string** | The IP of the balanced target VM. | |
|**Port** | **int32** | The port of the balanced target service; valid range is 1 to 65535. | |
|**Weight** | **int32** | Traffic is distributed in proportion to target weight, relative to the combined weight of all targets. A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1; targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best use values in the middle of the range to leave room for later adjustments. | |
|**HealthCheckEnabled** | Pointer to **bool** | When the health check is enabled, the target is available only when it accepts regular TCP or HTTP connection attempts for state checking. The state check consists of one connection attempt with the target&#39;s address and port. The default value is &#39;TRUE&#39;. | [optional] |
|**MaintenanceEnabled** | Pointer to **bool** | When the maintenance mode is enabled, the target is prevented from receiving traffic; the default value is &#39;FALSE&#39;. | [optional] |

## Methods

### NewTargetGroupTarget

`func NewTargetGroupTarget(ip string, port int32, weight int32, ) *TargetGroupTarget`

NewTargetGroupTarget instantiates a new TargetGroupTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupTargetWithDefaults

`func NewTargetGroupTargetWithDefaults() *TargetGroupTarget`

NewTargetGroupTargetWithDefaults instantiates a new TargetGroupTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *TargetGroupTarget) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TargetGroupTarget) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TargetGroupTarget) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPort

`func (o *TargetGroupTarget) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetGroupTarget) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetGroupTarget) SetPort(v int32)`

SetPort sets Port field to given value.


### GetWeight

`func (o *TargetGroupTarget) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *TargetGroupTarget) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *TargetGroupTarget) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetHealthCheckEnabled

`func (o *TargetGroupTarget) GetHealthCheckEnabled() bool`

GetHealthCheckEnabled returns the HealthCheckEnabled field if non-nil, zero value otherwise.

### GetHealthCheckEnabledOk

`func (o *TargetGroupTarget) GetHealthCheckEnabledOk() (*bool, bool)`

GetHealthCheckEnabledOk returns a tuple with the HealthCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckEnabled

`func (o *TargetGroupTarget) SetHealthCheckEnabled(v bool)`

SetHealthCheckEnabled sets HealthCheckEnabled field to given value.

### HasHealthCheckEnabled

`func (o *TargetGroupTarget) HasHealthCheckEnabled() bool`

HasHealthCheckEnabled returns a boolean if a field has been set.

### GetMaintenanceEnabled

`func (o *TargetGroupTarget) GetMaintenanceEnabled() bool`

GetMaintenanceEnabled returns the MaintenanceEnabled field if non-nil, zero value otherwise.

### GetMaintenanceEnabledOk

`func (o *TargetGroupTarget) GetMaintenanceEnabledOk() (*bool, bool)`

GetMaintenanceEnabledOk returns a tuple with the MaintenanceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceEnabled

`func (o *TargetGroupTarget) SetMaintenanceEnabled(v bool)`

SetMaintenanceEnabled sets MaintenanceEnabled field to given value.

### HasMaintenanceEnabled

`func (o *TargetGroupTarget) HasMaintenanceEnabled() bool`

HasMaintenanceEnabled returns a boolean if a field has been set.



