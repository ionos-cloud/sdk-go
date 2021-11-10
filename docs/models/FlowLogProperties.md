# FlowLogProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the  resource. | |
|**Action** | **string** | Specifies the traffic action pattern. | |
|**Direction** | **string** | Specifies the traffic direction pattern. | |
|**Bucket** | **string** | S3 bucket name of an existing IONOS Cloud S3 bucket. | |

## Methods

### NewFlowLogProperties

`func NewFlowLogProperties(name string, action string, direction string, bucket string, ) *FlowLogProperties`

NewFlowLogProperties instantiates a new FlowLogProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowLogPropertiesWithDefaults

`func NewFlowLogPropertiesWithDefaults() *FlowLogProperties`

NewFlowLogPropertiesWithDefaults instantiates a new FlowLogProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlowLogProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowLogProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowLogProperties) SetName(v string)`

SetName sets Name field to given value.


### GetAction

`func (o *FlowLogProperties) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FlowLogProperties) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FlowLogProperties) SetAction(v string)`

SetAction sets Action field to given value.


### GetDirection

`func (o *FlowLogProperties) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FlowLogProperties) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FlowLogProperties) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetBucket

`func (o *FlowLogProperties) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *FlowLogProperties) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *FlowLogProperties) SetBucket(v string)`

SetBucket sets Bucket field to given value.




