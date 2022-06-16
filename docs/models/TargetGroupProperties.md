# TargetGroupProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the target group. | |
|**Algorithm** | **string** | Balancing algorithm | |
|**Protocol** | **string** | Balancing protocol | |
|**Targets** | Pointer to [**[]TargetGroupTarget**](TargetGroupTarget.md) | Array of items in the collection. | [optional] |
|**HealthCheck** | Pointer to [**TargetGroupHealthCheck**](TargetGroupHealthCheck.md) |  | [optional] |
|**HttpHealthCheck** | Pointer to [**TargetGroupHttpHealthCheck**](TargetGroupHttpHealthCheck.md) |  | [optional] |

## Methods

### NewTargetGroupProperties

`func NewTargetGroupProperties(name string, algorithm string, protocol string, ) *TargetGroupProperties`

NewTargetGroupProperties instantiates a new TargetGroupProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupPropertiesWithDefaults

`func NewTargetGroupPropertiesWithDefaults() *TargetGroupProperties`

NewTargetGroupPropertiesWithDefaults instantiates a new TargetGroupProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TargetGroupProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetGroupProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetGroupProperties) SetName(v string)`

SetName sets Name field to given value.


### GetAlgorithm

`func (o *TargetGroupProperties) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TargetGroupProperties) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TargetGroupProperties) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetProtocol

`func (o *TargetGroupProperties) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *TargetGroupProperties) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *TargetGroupProperties) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetTargets

`func (o *TargetGroupProperties) GetTargets() []TargetGroupTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *TargetGroupProperties) GetTargetsOk() (*[]TargetGroupTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *TargetGroupProperties) SetTargets(v []TargetGroupTarget)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *TargetGroupProperties) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetHealthCheck

`func (o *TargetGroupProperties) GetHealthCheck() TargetGroupHealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *TargetGroupProperties) GetHealthCheckOk() (*TargetGroupHealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *TargetGroupProperties) SetHealthCheck(v TargetGroupHealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *TargetGroupProperties) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetHttpHealthCheck

`func (o *TargetGroupProperties) GetHttpHealthCheck() TargetGroupHttpHealthCheck`

GetHttpHealthCheck returns the HttpHealthCheck field if non-nil, zero value otherwise.

### GetHttpHealthCheckOk

`func (o *TargetGroupProperties) GetHttpHealthCheckOk() (*TargetGroupHttpHealthCheck, bool)`

GetHttpHealthCheckOk returns a tuple with the HttpHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHealthCheck

`func (o *TargetGroupProperties) SetHttpHealthCheck(v TargetGroupHttpHealthCheck)`

SetHttpHealthCheck sets HttpHealthCheck field to given value.

### HasHttpHealthCheck

`func (o *TargetGroupProperties) HasHttpHealthCheck() bool`

HasHttpHealthCheck returns a boolean if a field has been set.



