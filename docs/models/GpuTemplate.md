# GpuTemplate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Vendor** | Pointer to **string** | The name of the vendor | [optional] |
|**Model** | Pointer to **string** | The name of the GPU card model | [optional] |
|**Type** | Pointer to **string** | The way the Graphics Processing Unit (GPU) card will function. The &#39;passthrough&#39; type means that the Server will be connected to the GPU directly (e.g. no virtualization involved). | [optional] |
|**Count** | Pointer to **int32** | The number of GPUs to be provisioned | [optional] |

## Methods

### NewGpuTemplate

`func NewGpuTemplate() *GpuTemplate`

NewGpuTemplate instantiates a new GpuTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuTemplateWithDefaults

`func NewGpuTemplateWithDefaults() *GpuTemplate`

NewGpuTemplateWithDefaults instantiates a new GpuTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *GpuTemplate) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GpuTemplate) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GpuTemplate) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GpuTemplate) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetModel

`func (o *GpuTemplate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GpuTemplate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GpuTemplate) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GpuTemplate) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetType

`func (o *GpuTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GpuTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GpuTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GpuTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCount

`func (o *GpuTemplate) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GpuTemplate) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GpuTemplate) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GpuTemplate) HasCount() bool`

HasCount returns a boolean if a field has been set.



