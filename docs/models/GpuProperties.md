# GpuProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] |
|**Vendor** | Pointer to **string** | The vendor of the Graphics Processing Unit (GPU) card. The available options can be retrieved from the &#39;gpuArchitecture&#39; field returned in the GET responses of the resources /datacenters and /locations. | [optional] [readonly] |
|**Model** | Pointer to **string** | The model of Graphics Processing Unit (GPU) card. The available options can be retrieved from the &#39;gpuArchitecture&#39; field returned in the GET responses of the resources /datacenters and /locations. | [optional] [readonly] |
|**Type** | Pointer to **string** | The way the Graphics Processing Unit (GPU) card will function. The &#39;passthrough&#39; type means that the Server will be connected to the GPU directly (e.g. no virtualization involved). | [optional] [readonly] |

## Methods

### NewGpuProperties

`func NewGpuProperties() *GpuProperties`

NewGpuProperties instantiates a new GpuProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuPropertiesWithDefaults

`func NewGpuPropertiesWithDefaults() *GpuProperties`

NewGpuPropertiesWithDefaults instantiates a new GpuProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GpuProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GpuProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GpuProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GpuProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVendor

`func (o *GpuProperties) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GpuProperties) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GpuProperties) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GpuProperties) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetModel

`func (o *GpuProperties) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GpuProperties) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GpuProperties) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GpuProperties) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetType

`func (o *GpuProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GpuProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GpuProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GpuProperties) HasType() bool`

HasType returns a boolean if a field has been set.



