# GroupProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name of that resource | [optional] 
**CreateDataCenter** | Pointer to **bool** | create data center privilege | [optional] 
**CreateSnapshot** | Pointer to **bool** | create snapshot privilege | [optional] 
**ReserveIp** | Pointer to **bool** | reserve ip block privilege | [optional] 
**AccessActivityLog** | Pointer to **bool** | activity log access privilege | [optional] 
**CreatePcc** | Pointer to **bool** | create pcc privilege | [optional] 
**S3Privilege** | Pointer to **bool** | S3 privilege | [optional] 
**CreateBackupUnit** | Pointer to **bool** | create backup unit privilege | [optional] 
**CreateInternetAccess** | Pointer to **bool** | create internet access privilege | [optional] 
**CreateK8sCluster** | Pointer to **bool** | create kubernetes cluster privilege | [optional] 

## Methods

### NewGroupProperties

`func NewGroupProperties() *GroupProperties`

NewGroupProperties instantiates a new GroupProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPropertiesWithDefaults

`func NewGroupPropertiesWithDefaults() *GroupProperties`

NewGroupPropertiesWithDefaults instantiates a new GroupProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreateDataCenter

`func (o *GroupProperties) GetCreateDataCenter() bool`

GetCreateDataCenter returns the CreateDataCenter field if non-nil, zero value otherwise.

### GetCreateDataCenterOk

`func (o *GroupProperties) GetCreateDataCenterOk() (*bool, bool)`

GetCreateDataCenterOk returns a tuple with the CreateDataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDataCenter

`func (o *GroupProperties) SetCreateDataCenter(v bool)`

SetCreateDataCenter sets CreateDataCenter field to given value.

### HasCreateDataCenter

`func (o *GroupProperties) HasCreateDataCenter() bool`

HasCreateDataCenter returns a boolean if a field has been set.

### GetCreateSnapshot

`func (o *GroupProperties) GetCreateSnapshot() bool`

GetCreateSnapshot returns the CreateSnapshot field if non-nil, zero value otherwise.

### GetCreateSnapshotOk

`func (o *GroupProperties) GetCreateSnapshotOk() (*bool, bool)`

GetCreateSnapshotOk returns a tuple with the CreateSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSnapshot

`func (o *GroupProperties) SetCreateSnapshot(v bool)`

SetCreateSnapshot sets CreateSnapshot field to given value.

### HasCreateSnapshot

`func (o *GroupProperties) HasCreateSnapshot() bool`

HasCreateSnapshot returns a boolean if a field has been set.

### GetReserveIp

`func (o *GroupProperties) GetReserveIp() bool`

GetReserveIp returns the ReserveIp field if non-nil, zero value otherwise.

### GetReserveIpOk

`func (o *GroupProperties) GetReserveIpOk() (*bool, bool)`

GetReserveIpOk returns a tuple with the ReserveIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveIp

`func (o *GroupProperties) SetReserveIp(v bool)`

SetReserveIp sets ReserveIp field to given value.

### HasReserveIp

`func (o *GroupProperties) HasReserveIp() bool`

HasReserveIp returns a boolean if a field has been set.

### GetAccessActivityLog

`func (o *GroupProperties) GetAccessActivityLog() bool`

GetAccessActivityLog returns the AccessActivityLog field if non-nil, zero value otherwise.

### GetAccessActivityLogOk

`func (o *GroupProperties) GetAccessActivityLogOk() (*bool, bool)`

GetAccessActivityLogOk returns a tuple with the AccessActivityLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessActivityLog

`func (o *GroupProperties) SetAccessActivityLog(v bool)`

SetAccessActivityLog sets AccessActivityLog field to given value.

### HasAccessActivityLog

`func (o *GroupProperties) HasAccessActivityLog() bool`

HasAccessActivityLog returns a boolean if a field has been set.

### GetCreatePcc

`func (o *GroupProperties) GetCreatePcc() bool`

GetCreatePcc returns the CreatePcc field if non-nil, zero value otherwise.

### GetCreatePccOk

`func (o *GroupProperties) GetCreatePccOk() (*bool, bool)`

GetCreatePccOk returns a tuple with the CreatePcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePcc

`func (o *GroupProperties) SetCreatePcc(v bool)`

SetCreatePcc sets CreatePcc field to given value.

### HasCreatePcc

`func (o *GroupProperties) HasCreatePcc() bool`

HasCreatePcc returns a boolean if a field has been set.

### GetS3Privilege

`func (o *GroupProperties) GetS3Privilege() bool`

GetS3Privilege returns the S3Privilege field if non-nil, zero value otherwise.

### GetS3PrivilegeOk

`func (o *GroupProperties) GetS3PrivilegeOk() (*bool, bool)`

GetS3PrivilegeOk returns a tuple with the S3Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Privilege

`func (o *GroupProperties) SetS3Privilege(v bool)`

SetS3Privilege sets S3Privilege field to given value.

### HasS3Privilege

`func (o *GroupProperties) HasS3Privilege() bool`

HasS3Privilege returns a boolean if a field has been set.

### GetCreateBackupUnit

`func (o *GroupProperties) GetCreateBackupUnit() bool`

GetCreateBackupUnit returns the CreateBackupUnit field if non-nil, zero value otherwise.

### GetCreateBackupUnitOk

`func (o *GroupProperties) GetCreateBackupUnitOk() (*bool, bool)`

GetCreateBackupUnitOk returns a tuple with the CreateBackupUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackupUnit

`func (o *GroupProperties) SetCreateBackupUnit(v bool)`

SetCreateBackupUnit sets CreateBackupUnit field to given value.

### HasCreateBackupUnit

`func (o *GroupProperties) HasCreateBackupUnit() bool`

HasCreateBackupUnit returns a boolean if a field has been set.

### GetCreateInternetAccess

`func (o *GroupProperties) GetCreateInternetAccess() bool`

GetCreateInternetAccess returns the CreateInternetAccess field if non-nil, zero value otherwise.

### GetCreateInternetAccessOk

`func (o *GroupProperties) GetCreateInternetAccessOk() (*bool, bool)`

GetCreateInternetAccessOk returns a tuple with the CreateInternetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInternetAccess

`func (o *GroupProperties) SetCreateInternetAccess(v bool)`

SetCreateInternetAccess sets CreateInternetAccess field to given value.

### HasCreateInternetAccess

`func (o *GroupProperties) HasCreateInternetAccess() bool`

HasCreateInternetAccess returns a boolean if a field has been set.

### GetCreateK8sCluster

`func (o *GroupProperties) GetCreateK8sCluster() bool`

GetCreateK8sCluster returns the CreateK8sCluster field if non-nil, zero value otherwise.

### GetCreateK8sClusterOk

`func (o *GroupProperties) GetCreateK8sClusterOk() (*bool, bool)`

GetCreateK8sClusterOk returns a tuple with the CreateK8sCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateK8sCluster

`func (o *GroupProperties) SetCreateK8sCluster(v bool)`

SetCreateK8sCluster sets CreateK8sCluster field to given value.

### HasCreateK8sCluster

`func (o *GroupProperties) HasCreateK8sCluster() bool`

HasCreateK8sCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


