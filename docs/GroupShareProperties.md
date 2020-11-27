# GroupShareProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditPrivilege** | Pointer to **bool** | edit privilege on a resource | [optional] 
**SharePrivilege** | Pointer to **bool** | share privilege on a resource | [optional] 

## Methods

### NewGroupShareProperties

`func NewGroupShareProperties() *GroupShareProperties`

NewGroupShareProperties instantiates a new GroupShareProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSharePropertiesWithDefaults

`func NewGroupSharePropertiesWithDefaults() *GroupShareProperties`

NewGroupSharePropertiesWithDefaults instantiates a new GroupShareProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditPrivilege

`func (o *GroupShareProperties) GetEditPrivilege() bool`

GetEditPrivilege returns the EditPrivilege field if non-nil, zero value otherwise.

### GetEditPrivilegeOk

`func (o *GroupShareProperties) GetEditPrivilegeOk() (*bool, bool)`

GetEditPrivilegeOk returns a tuple with the EditPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPrivilege

`func (o *GroupShareProperties) SetEditPrivilege(v bool)`

SetEditPrivilege sets EditPrivilege field to given value.

### HasEditPrivilege

`func (o *GroupShareProperties) HasEditPrivilege() bool`

HasEditPrivilege returns a boolean if a field has been set.

### GetSharePrivilege

`func (o *GroupShareProperties) GetSharePrivilege() bool`

GetSharePrivilege returns the SharePrivilege field if non-nil, zero value otherwise.

### GetSharePrivilegeOk

`func (o *GroupShareProperties) GetSharePrivilegeOk() (*bool, bool)`

GetSharePrivilegeOk returns a tuple with the SharePrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePrivilege

`func (o *GroupShareProperties) SetSharePrivilege(v bool)`

SetSharePrivilege sets SharePrivilege field to given value.

### HasSharePrivilege

`func (o *GroupShareProperties) HasSharePrivilege() bool`

HasSharePrivilege returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


