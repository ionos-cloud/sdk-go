# RequestTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestTarget

`func NewRequestTarget() *RequestTarget`

NewRequestTarget instantiates a new RequestTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTargetWithDefaults

`func NewRequestTargetWithDefaults() *RequestTarget`

NewRequestTargetWithDefaults instantiates a new RequestTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *RequestTarget) GetTarget() ResourceReference`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RequestTarget) GetTargetOk() (*ResourceReference, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RequestTarget) SetTarget(v ResourceReference)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RequestTarget) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStatus

`func (o *RequestTarget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestTarget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestTarget) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestTarget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


