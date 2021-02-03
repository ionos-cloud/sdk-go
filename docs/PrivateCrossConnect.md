# PrivateCrossConnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] 
**Properties** | [**PrivateCrossConnectProperties**](PrivateCrossConnectProperties.md) |  | 

## Methods

### NewPrivateCrossConnect

`func NewPrivateCrossConnect(properties PrivateCrossConnectProperties, ) *PrivateCrossConnect`

NewPrivateCrossConnect instantiates a new PrivateCrossConnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateCrossConnectWithDefaults

`func NewPrivateCrossConnectWithDefaults() *PrivateCrossConnect`

NewPrivateCrossConnectWithDefaults instantiates a new PrivateCrossConnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateCrossConnect) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateCrossConnect) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateCrossConnect) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateCrossConnect) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PrivateCrossConnect) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateCrossConnect) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateCrossConnect) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *PrivateCrossConnect) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *PrivateCrossConnect) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PrivateCrossConnect) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PrivateCrossConnect) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PrivateCrossConnect) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *PrivateCrossConnect) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PrivateCrossConnect) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PrivateCrossConnect) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PrivateCrossConnect) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *PrivateCrossConnect) GetProperties() PrivateCrossConnectProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PrivateCrossConnect) GetPropertiesOk() (*PrivateCrossConnectProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PrivateCrossConnect) SetProperties(v PrivateCrossConnectProperties)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


