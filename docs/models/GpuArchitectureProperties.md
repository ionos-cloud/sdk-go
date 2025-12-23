# GpuArchitectureProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Vendor** | Pointer to **string** | The vendor of the Graphics Processing Unit (GPU) card. | [optional] |
|**Model** | Pointer to **string** | The model of the Graphics Processing Unit (GPU) card. | [optional] |

## Methods

### NewGpuArchitectureProperties

`func NewGpuArchitectureProperties() *GpuArchitectureProperties`

NewGpuArchitectureProperties instantiates a new GpuArchitectureProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuArchitecturePropertiesWithDefaults

`func NewGpuArchitecturePropertiesWithDefaults() *GpuArchitectureProperties`

NewGpuArchitecturePropertiesWithDefaults instantiates a new GpuArchitectureProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *GpuArchitectureProperties) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GpuArchitectureProperties) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GpuArchitectureProperties) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GpuArchitectureProperties) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetModel

`func (o *GpuArchitectureProperties) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GpuArchitectureProperties) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GpuArchitectureProperties) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GpuArchitectureProperties) HasModel() bool`

HasModel returns a boolean if a field has been set.



