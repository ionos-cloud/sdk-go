# ResourceProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of the resource | [optional] 
**SecAuthProtection** | Pointer to **bool** | Boolean value representing if the resource is multi factor protected or not e.g. using two factor protection. Currently only Data Centers and Snapshots are allowed to be multi factor protected, The value of attribute if null is intentional and it means that the resource doesn&#39;t support multi factor protection at all. | [optional] 

## Methods

### NewResourceProperties

`func NewResourceProperties() *ResourceProperties`

NewResourceProperties instantiates a new ResourceProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePropertiesWithDefaults

`func NewResourcePropertiesWithDefaults() *ResourceProperties`

NewResourcePropertiesWithDefaults instantiates a new ResourceProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResourceProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecAuthProtection

`func (o *ResourceProperties) GetSecAuthProtection() bool`

GetSecAuthProtection returns the SecAuthProtection field if non-nil, zero value otherwise.

### GetSecAuthProtectionOk

`func (o *ResourceProperties) GetSecAuthProtectionOk() (*bool, bool)`

GetSecAuthProtectionOk returns a tuple with the SecAuthProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAuthProtection

`func (o *ResourceProperties) SetSecAuthProtection(v bool)`

SetSecAuthProtection sets SecAuthProtection field to given value.

### HasSecAuthProtection

`func (o *ResourceProperties) HasSecAuthProtection() bool`

HasSecAuthProtection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


