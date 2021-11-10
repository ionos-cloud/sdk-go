# UserPut

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] |
|**Properties** | [**UserPropertiesPut**](UserPropertiesPut.md) |  | |

## Methods

### NewUserPut

`func NewUserPut(properties UserPropertiesPut, ) *UserPut`

NewUserPut instantiates a new UserPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPutWithDefaults

`func NewUserPutWithDefaults() *UserPut`

NewUserPutWithDefaults instantiates a new UserPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserPut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserPut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *UserPut) GetProperties() UserPropertiesPut`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserPut) GetPropertiesOk() (*UserPropertiesPut, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserPut) SetProperties(v UserPropertiesPut)`

SetProperties sets Properties field to given value.




