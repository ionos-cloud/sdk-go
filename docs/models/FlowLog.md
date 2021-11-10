# FlowLog

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to [**Type**](Type.md) | The type of object that has been created. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**FlowLogProperties**](FlowLogProperties.md) |  | |

## Methods

### NewFlowLog

`func NewFlowLog(properties FlowLogProperties, ) *FlowLog`

NewFlowLog instantiates a new FlowLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowLogWithDefaults

`func NewFlowLogWithDefaults() *FlowLog`

NewFlowLogWithDefaults instantiates a new FlowLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *FlowLog) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowLog) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowLog) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *FlowLog) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *FlowLog) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FlowLog) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FlowLog) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FlowLog) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *FlowLog) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlowLog) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlowLog) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FlowLog) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *FlowLog) GetProperties() FlowLogProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FlowLog) GetPropertiesOk() (*FlowLogProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FlowLog) SetProperties(v FlowLogProperties)`

SetProperties sets Properties field to given value.




