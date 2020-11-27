# LabelProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | A Label Key | [optional] 
**Value** | Pointer to **string** | A Label Value | [optional] 
**ResourceId** | Pointer to **string** | The id of the resource | [optional] 
**ResourceType** | Pointer to **string** | The type of the resource on which the label is applied. | [optional] 
**ResourceHref** | Pointer to **string** | URL to the Resource (absolute path) on which the label is applied. | [optional] 

## Methods

### NewLabelProperties

`func NewLabelProperties() *LabelProperties`

NewLabelProperties instantiates a new LabelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelPropertiesWithDefaults

`func NewLabelPropertiesWithDefaults() *LabelProperties`

NewLabelPropertiesWithDefaults instantiates a new LabelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LabelProperties) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LabelProperties) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LabelProperties) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LabelProperties) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *LabelProperties) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LabelProperties) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LabelProperties) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *LabelProperties) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetResourceId

`func (o *LabelProperties) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *LabelProperties) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *LabelProperties) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *LabelProperties) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *LabelProperties) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *LabelProperties) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *LabelProperties) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *LabelProperties) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceHref

`func (o *LabelProperties) GetResourceHref() string`

GetResourceHref returns the ResourceHref field if non-nil, zero value otherwise.

### GetResourceHrefOk

`func (o *LabelProperties) GetResourceHrefOk() (*string, bool)`

GetResourceHrefOk returns a tuple with the ResourceHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceHref

`func (o *LabelProperties) SetResourceHref(v string)`

SetResourceHref sets ResourceHref field to given value.

### HasResourceHref

`func (o *LabelProperties) HasResourceHref() bool`

HasResourceHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


