# TemplateProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the  resource. | |
|**Cores** | **float32** | The CPU cores count. | |
|**Ram** | **float32** | The RAM size in MB. | |
|**StorageSize** | **float32** | The storage size in GB. | |

## Methods

### NewTemplateProperties

`func NewTemplateProperties(name string, cores float32, ram float32, storageSize float32, ) *TemplateProperties`

NewTemplateProperties instantiates a new TemplateProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePropertiesWithDefaults

`func NewTemplatePropertiesWithDefaults() *TemplateProperties`

NewTemplatePropertiesWithDefaults instantiates a new TemplateProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateProperties) SetName(v string)`

SetName sets Name field to given value.


### GetCores

`func (o *TemplateProperties) GetCores() float32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *TemplateProperties) GetCoresOk() (*float32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *TemplateProperties) SetCores(v float32)`

SetCores sets Cores field to given value.


### GetRam

`func (o *TemplateProperties) GetRam() float32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *TemplateProperties) GetRamOk() (*float32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *TemplateProperties) SetRam(v float32)`

SetRam sets Ram field to given value.


### GetStorageSize

`func (o *TemplateProperties) GetStorageSize() float32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *TemplateProperties) GetStorageSizeOk() (*float32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *TemplateProperties) SetStorageSize(v float32)`

SetStorageSize sets StorageSize field to given value.




