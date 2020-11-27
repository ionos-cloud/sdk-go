# GroupEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**GroupMembers**](GroupMembers.md) |  | [optional] 
**Resources** | Pointer to [**ResourceGroups**](ResourceGroups.md) |  | [optional] 

## Methods

### NewGroupEntities

`func NewGroupEntities() *GroupEntities`

NewGroupEntities instantiates a new GroupEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupEntitiesWithDefaults

`func NewGroupEntitiesWithDefaults() *GroupEntities`

NewGroupEntitiesWithDefaults instantiates a new GroupEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GroupEntities) GetUsers() GroupMembers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupEntities) GetUsersOk() (*GroupMembers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupEntities) SetUsers(v GroupMembers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupEntities) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetResources

`func (o *GroupEntities) GetResources() ResourceGroups`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GroupEntities) GetResourcesOk() (*ResourceGroups, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GroupEntities) SetResources(v ResourceGroups)`

SetResources sets Resources field to given value.

### HasResources

`func (o *GroupEntities) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


