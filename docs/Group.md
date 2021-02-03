# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Type** | Pointer to [**Type**](Type.md) | The type of the resource | [optional] 
**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] 
**Properties** | [**GroupProperties**](GroupProperties.md) |  | 
**Entities** | Pointer to [**GroupEntities**](GroupEntities.md) |  | [optional] 

## Methods

### NewGroup

`func NewGroup(properties GroupProperties, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Group) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Group) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Group) SetType(v Type)`

SetType sets Type field to given value.

### HasType

`func (o *Group) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Group) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Group) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Group) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Group) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetProperties

`func (o *Group) GetProperties() GroupProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Group) GetPropertiesOk() (*GroupProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Group) SetProperties(v GroupProperties)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *Group) GetEntities() GroupEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *Group) GetEntitiesOk() (*GroupEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *Group) SetEntities(v GroupEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *Group) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


