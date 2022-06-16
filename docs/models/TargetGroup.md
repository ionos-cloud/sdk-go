# TargetGroup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**TargetGroupProperties**](TargetGroupProperties.md) |  | |

## Methods

### NewTargetGroup

`func NewTargetGroup(properties TargetGroupProperties, ) *TargetGroup`

NewTargetGroup instantiates a new TargetGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupWithDefaults

`func NewTargetGroupWithDefaults() *TargetGroup`

NewTargetGroupWithDefaults instantiates a new TargetGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TargetGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TargetGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TargetGroup) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TargetGroup) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TargetGroup) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *TargetGroup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *TargetGroup) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TargetGroup) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TargetGroup) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *TargetGroup) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *TargetGroup) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TargetGroup) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TargetGroup) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TargetGroup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *TargetGroup) GetProperties() TargetGroupProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TargetGroup) GetPropertiesOk() (*TargetGroupProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TargetGroup) SetProperties(v TargetGroupProperties)`

SetProperties sets Properties field to given value.




