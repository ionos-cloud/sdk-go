# LanPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] 
**Entities** | Pointer to [**LanEntities**](LanEntities.md) |  | [optional] 
**Properties** | [**LanPropertiesPost**](LanPropertiesPost.md) |  | 

## Methods

### NewLanPost

`func NewLanPost(properties LanPropertiesPost, ) *LanPost`

NewLanPost instantiates a new LanPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanPostWithDefaults

`func NewLanPostWithDefaults() *LanPost`

NewLanPostWithDefaults instantiates a new LanPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LanPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LanPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LanPost) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LanPost) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LanPost) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *LanPost) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *LanPost) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LanPost) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LanPost) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LanPost) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *LanPost) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LanPost) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LanPost) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LanPost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntities

`func (o *LanPost) GetEntities() LanEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *LanPost) GetEntitiesOk() (*LanEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *LanPost) SetEntities(v LanEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *LanPost) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetProperties

`func (o *LanPost) GetProperties() LanPropertiesPost`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *LanPost) GetPropertiesOk() (*LanPropertiesPost, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *LanPost) SetProperties(v LanPropertiesPost)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


