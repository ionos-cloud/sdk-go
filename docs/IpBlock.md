# IpBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] 
**Properties** | [**IpBlockProperties**](IpBlockProperties.md) |  | 

## Methods

### NewIpBlock

`func NewIpBlock(properties IpBlockProperties, ) *IpBlock`

NewIpBlock instantiates a new IpBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockWithDefaults

`func NewIpBlockWithDefaults() *IpBlock`

NewIpBlockWithDefaults instantiates a new IpBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *IpBlock) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpBlock) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpBlock) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *IpBlock) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *IpBlock) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IpBlock) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IpBlock) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IpBlock) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *IpBlock) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IpBlock) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IpBlock) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IpBlock) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *IpBlock) GetProperties() IpBlockProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IpBlock) GetPropertiesOk() (*IpBlockProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IpBlock) SetProperties(v IpBlockProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


