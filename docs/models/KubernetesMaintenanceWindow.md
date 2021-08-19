# KubernetesMaintenanceWindow

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DayOfTheWeek** | **string** | The day of the week for a maintenance window. | |
|**Time** | **string** | The time to use for a maintenance window. Accepted formats are: HH:mm:ss; HH:mm:ss\&quot;Z\&quot;; HH:mm:ssZ. This time may varies by 15 minutes. | |

## Methods

### NewKubernetesMaintenanceWindow

`func NewKubernetesMaintenanceWindow(dayOfTheWeek string, time string, ) *KubernetesMaintenanceWindow`

NewKubernetesMaintenanceWindow instantiates a new KubernetesMaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesMaintenanceWindowWithDefaults

`func NewKubernetesMaintenanceWindowWithDefaults() *KubernetesMaintenanceWindow`

NewKubernetesMaintenanceWindowWithDefaults instantiates a new KubernetesMaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfTheWeek

`func (o *KubernetesMaintenanceWindow) GetDayOfTheWeek() string`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *KubernetesMaintenanceWindow) GetDayOfTheWeekOk() (*string, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *KubernetesMaintenanceWindow) SetDayOfTheWeek(v string)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.


### GetTime

`func (o *KubernetesMaintenanceWindow) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *KubernetesMaintenanceWindow) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *KubernetesMaintenanceWindow) SetTime(v string)`

SetTime sets Time field to given value.




