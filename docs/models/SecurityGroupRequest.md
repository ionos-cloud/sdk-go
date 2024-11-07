# SecurityGroupRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**SecurityGroupProperties**](SecurityGroupProperties.md) |  | |
|**Entities** | Pointer to [**SecurityGroupEntitiesRequest**](SecurityGroupEntitiesRequest.md) |  | [optional] |

## Methods

### NewSecurityGroupRequest

`func NewSecurityGroupRequest(properties SecurityGroupProperties, ) *SecurityGroupRequest`

NewSecurityGroupRequest instantiates a new SecurityGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRequestWithDefaults

`func NewSecurityGroupRequestWithDefaults() *SecurityGroupRequest`

NewSecurityGroupRequestWithDefaults instantiates a new SecurityGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *SecurityGroupRequest) GetProperties() SecurityGroupProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SecurityGroupRequest) GetPropertiesOk() (*SecurityGroupProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SecurityGroupRequest) SetProperties(v SecurityGroupProperties)`

SetProperties sets Properties field to given value.


### GetEntities

`func (o *SecurityGroupRequest) GetEntities() SecurityGroupEntitiesRequest`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SecurityGroupRequest) GetEntitiesOk() (*SecurityGroupEntitiesRequest, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SecurityGroupRequest) SetEntities(v SecurityGroupEntitiesRequest)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *SecurityGroupRequest) HasEntities() bool`

HasEntities returns a boolean if a field has been set.



