# TargetGroupHttpHealthCheck

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Path** | Pointer to **string** | The path (destination URL) for the HTTP health check request; the default is /. | [optional] |
|**Method** | Pointer to **string** | The method for the HTTP health check. | [optional] |
|**MatchType** | **string** |  | |
|**Response** | **string** | The response returned by the request, depending on the match type. | |
|**Regex** | Pointer to **bool** |  | [optional] |
|**Negate** | Pointer to **bool** |  | [optional] |

## Methods

### NewTargetGroupHttpHealthCheck

`func NewTargetGroupHttpHealthCheck(matchType string, response string, ) *TargetGroupHttpHealthCheck`

NewTargetGroupHttpHealthCheck instantiates a new TargetGroupHttpHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupHttpHealthCheckWithDefaults

`func NewTargetGroupHttpHealthCheckWithDefaults() *TargetGroupHttpHealthCheck`

NewTargetGroupHttpHealthCheckWithDefaults instantiates a new TargetGroupHttpHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *TargetGroupHttpHealthCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TargetGroupHttpHealthCheck) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TargetGroupHttpHealthCheck) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TargetGroupHttpHealthCheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMethod

`func (o *TargetGroupHttpHealthCheck) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TargetGroupHttpHealthCheck) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TargetGroupHttpHealthCheck) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *TargetGroupHttpHealthCheck) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMatchType

`func (o *TargetGroupHttpHealthCheck) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *TargetGroupHttpHealthCheck) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *TargetGroupHttpHealthCheck) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetResponse

`func (o *TargetGroupHttpHealthCheck) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *TargetGroupHttpHealthCheck) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *TargetGroupHttpHealthCheck) SetResponse(v string)`

SetResponse sets Response field to given value.


### GetRegex

`func (o *TargetGroupHttpHealthCheck) GetRegex() bool`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *TargetGroupHttpHealthCheck) GetRegexOk() (*bool, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *TargetGroupHttpHealthCheck) SetRegex(v bool)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *TargetGroupHttpHealthCheck) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetNegate

`func (o *TargetGroupHttpHealthCheck) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *TargetGroupHttpHealthCheck) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *TargetGroupHttpHealthCheck) SetNegate(v bool)`

SetNegate sets Negate field to given value.

### HasNegate

`func (o *TargetGroupHttpHealthCheck) HasNegate() bool`

HasNegate returns a boolean if a field has been set.



