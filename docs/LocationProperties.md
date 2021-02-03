# LocationProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name of that resource | [optional] 
**Features** | Pointer to **[]string** | List of features supported by the location | [optional] [readonly] 
**ImageAliases** | Pointer to **[]string** | List of image aliases available for the location | [optional] [readonly] 

## Methods

### NewLocationProperties

`func NewLocationProperties() *LocationProperties`

NewLocationProperties instantiates a new LocationProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationPropertiesWithDefaults

`func NewLocationPropertiesWithDefaults() *LocationProperties`

NewLocationPropertiesWithDefaults instantiates a new LocationProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LocationProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFeatures

`func (o *LocationProperties) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *LocationProperties) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *LocationProperties) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *LocationProperties) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetImageAliases

`func (o *LocationProperties) GetImageAliases() []string`

GetImageAliases returns the ImageAliases field if non-nil, zero value otherwise.

### GetImageAliasesOk

`func (o *LocationProperties) GetImageAliasesOk() (*[]string, bool)`

GetImageAliasesOk returns a tuple with the ImageAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAliases

`func (o *LocationProperties) SetImageAliases(v []string)`

SetImageAliases sets ImageAliases field to given value.

### HasImageAliases

`func (o *LocationProperties) HasImageAliases() bool`

HasImageAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


