# UsersEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owns** | Pointer to [**ResourcesUsers**](ResourcesUsers.md) |  | [optional] 
**Groups** | Pointer to [**GroupUsers**](GroupUsers.md) |  | [optional] 

## Methods

### NewUsersEntities

`func NewUsersEntities() *UsersEntities`

NewUsersEntities instantiates a new UsersEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersEntitiesWithDefaults

`func NewUsersEntitiesWithDefaults() *UsersEntities`

NewUsersEntitiesWithDefaults instantiates a new UsersEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwns

`func (o *UsersEntities) GetOwns() ResourcesUsers`

GetOwns returns the Owns field if non-nil, zero value otherwise.

### GetOwnsOk

`func (o *UsersEntities) GetOwnsOk() (*ResourcesUsers, bool)`

GetOwnsOk returns a tuple with the Owns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwns

`func (o *UsersEntities) SetOwns(v ResourcesUsers)`

SetOwns sets Owns field to given value.

### HasOwns

`func (o *UsersEntities) HasOwns() bool`

HasOwns returns a boolean if a field has been set.

### GetGroups

`func (o *UsersEntities) GetGroups() GroupUsers`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UsersEntities) GetGroupsOk() (*GroupUsers, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UsersEntities) SetGroups(v GroupUsers)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UsersEntities) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


