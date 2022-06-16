# TargetGroupHealthCheck

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CheckTimeout** | Pointer to **int32** | The maximum time in milliseconds to wait for a target to respond to a check. For target VMs with &#39;Check Interval&#39; set, the lesser of the two  values is used once the TCP connection is established. | [optional] |
|**CheckInterval** | Pointer to **int32** | The interval in milliseconds between consecutive health checks; default is 2000. | [optional] |
|**Retries** | Pointer to **int32** | The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535, and default is three reconnection attempts. | [optional] |

## Methods

### NewTargetGroupHealthCheck

`func NewTargetGroupHealthCheck() *TargetGroupHealthCheck`

NewTargetGroupHealthCheck instantiates a new TargetGroupHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupHealthCheckWithDefaults

`func NewTargetGroupHealthCheckWithDefaults() *TargetGroupHealthCheck`

NewTargetGroupHealthCheckWithDefaults instantiates a new TargetGroupHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckTimeout

`func (o *TargetGroupHealthCheck) GetCheckTimeout() int32`

GetCheckTimeout returns the CheckTimeout field if non-nil, zero value otherwise.

### GetCheckTimeoutOk

`func (o *TargetGroupHealthCheck) GetCheckTimeoutOk() (*int32, bool)`

GetCheckTimeoutOk returns a tuple with the CheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTimeout

`func (o *TargetGroupHealthCheck) SetCheckTimeout(v int32)`

SetCheckTimeout sets CheckTimeout field to given value.

### HasCheckTimeout

`func (o *TargetGroupHealthCheck) HasCheckTimeout() bool`

HasCheckTimeout returns a boolean if a field has been set.

### GetCheckInterval

`func (o *TargetGroupHealthCheck) GetCheckInterval() int32`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *TargetGroupHealthCheck) GetCheckIntervalOk() (*int32, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *TargetGroupHealthCheck) SetCheckInterval(v int32)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *TargetGroupHealthCheck) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetRetries

`func (o *TargetGroupHealthCheck) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *TargetGroupHealthCheck) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *TargetGroupHealthCheck) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *TargetGroupHealthCheck) HasRetries() bool`

HasRetries returns a boolean if a field has been set.



