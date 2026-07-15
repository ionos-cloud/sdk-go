# KubernetesNodePoolTaint

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Key** | **string** | Taint key. Must be a valid Kubernetes label key format. May include an optional prefix (DNS subdomain) followed by a slash. | |
|**Value** | Pointer to **string** | Optional taint value. Must be a valid Kubernetes label value format. | [optional] |
|**Effect** | [**TaintEffect**](TaintEffect.md) |  | |

## Methods

### NewKubernetesNodePoolTaint

`func NewKubernetesNodePoolTaint(key string, effect TaintEffect, ) *KubernetesNodePoolTaint`

NewKubernetesNodePoolTaint instantiates a new KubernetesNodePoolTaint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolTaintWithDefaults

`func NewKubernetesNodePoolTaintWithDefaults() *KubernetesNodePoolTaint`

NewKubernetesNodePoolTaintWithDefaults instantiates a new KubernetesNodePoolTaint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KubernetesNodePoolTaint) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KubernetesNodePoolTaint) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KubernetesNodePoolTaint) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *KubernetesNodePoolTaint) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubernetesNodePoolTaint) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubernetesNodePoolTaint) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KubernetesNodePoolTaint) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEffect

`func (o *KubernetesNodePoolTaint) GetEffect() TaintEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *KubernetesNodePoolTaint) GetEffectOk() (*TaintEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *KubernetesNodePoolTaint) SetEffect(v TaintEffect)`

SetEffect sets Effect field to given value.




