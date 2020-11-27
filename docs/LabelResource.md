# LabelResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Label on a resource is identified using label key. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of object that has been created | [optional] [readonly] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**NoStateMetaData**](NoStateMetaData.md) |  | [optional] 
**Properties** | [**LabelResourceProperties**](LabelResourceProperties.md) |  | 

## Methods

### NewLabelResource

`func NewLabelResource(properties LabelResourceProperties, ) *LabelResource`

NewLabelResource instantiates a new LabelResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelResourceWithDefaults

`func NewLabelResourceWithDefaults() *LabelResource`

NewLabelResourceWithDefaults instantiates a new LabelResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LabelResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LabelResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LabelResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabelResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabelResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LabelResource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *LabelResource) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LabelResource) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LabelResource) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LabelResource) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *LabelResource) GetMetadata() NoStateMetaData`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LabelResource) GetMetadataOk() (*NoStateMetaData, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LabelResource) SetMetadata(v NoStateMetaData)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LabelResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *LabelResource) GetProperties() LabelResourceProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *LabelResource) GetPropertiesOk() (*LabelResourceProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *LabelResource) SetProperties(v LabelResourceProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


