# GroupProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the resource. | [optional] |
|**CreateDataCenter** | Pointer to **bool** | Create data center privilege. | [optional] |
|**CreateSnapshot** | Pointer to **bool** | Create snapshot privilege. | [optional] |
|**ReserveIp** | Pointer to **bool** | Reserve IP block privilege. | [optional] |
|**AccessActivityLog** | Pointer to **bool** | Activity log access privilege. | [optional] |
|**CreatePcc** | Pointer to **bool** | User privilege to create a cross connect. | [optional] |
|**S3Privilege** | Pointer to **bool** | S3 privilege. | [optional] |
|**CreateBackupUnit** | Pointer to **bool** | Create backup unit privilege. | [optional] |
|**CreateInternetAccess** | Pointer to **bool** | Create internet access privilege. | [optional] |
|**CreateK8sCluster** | Pointer to **bool** | Create Kubernetes cluster privilege. | [optional] |
|**CreateFlowLog** | Pointer to **bool** | Create Flow Logs privilege. | [optional] |
|**AccessAndManageMonitoring** | Pointer to **bool** | Privilege for a group to access and manage monitoring related functionality (access metrics, CRUD on alarms, alarm-actions etc) using Monotoring-as-a-Service (MaaS). | [optional] |
|**AccessAndManageCertificates** | Pointer to **bool** | Privilege for a group to access and manage certificates. | [optional] |
|**ManageDBaaS** | Pointer to **bool** | Privilege for a group to manage DBaaS related functionality. | [optional] |
|**AccessAndManageDns** | Pointer to **bool** | Privilege for a group to access and manage dns records. | [optional] |
|**ManageRegistry** | Pointer to **bool** | Privilege for group accessing container registry related functionality. | [optional] |
|**ManageDataplatform** | Pointer to **bool** | Privilege for a group to access and manage the Data Platform. | [optional] |
|**AccessAndManageLogging** | Pointer to **bool** | Privilege for a group to access and manage Logs. | [optional] |
|**AccessAndManageCdn** | Pointer to **bool** | Privilege for a group to access and manage CDN. | [optional] |
|**AccessAndManageVpn** | Pointer to **bool** | Privilege for a group to access and manage VPN. | [optional] |
|**AccessAndManageApiGateway** | Pointer to **bool** | Privilege for a group to access and manage API Gateway. | [optional] |
|**AccessAndManageKaas** | Pointer to **bool** | Privilege for a group to access and manage KaaS. | [optional] |
|**AccessAndManageNetworkFileStorage** | Pointer to **bool** | Privilege for a group to access and manage Network File Storage. | [optional] |
|**AccessAndManageAiModelHub** | Pointer to **bool** | Privilege for a group to access and manage AI Model Hub. | [optional] |
|**AccessAndManageIamResources** | Pointer to **bool** | Privilege for a group to access and manage Password Policies. | [optional] |
|**CreateNetworkSecurityGroups** | Pointer to **bool** | Privilege for a group to access and manage Network Security Groups. | [optional] |

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

### GetCreateFlowLog

`func (o *GroupProperties) GetCreateFlowLog() bool`

GetCreateFlowLog returns the CreateFlowLog field if non-nil, zero value otherwise.

### GetCreateFlowLogOk

`func (o *GroupProperties) GetCreateFlowLogOk() (*bool, bool)`

GetCreateFlowLogOk returns a tuple with the CreateFlowLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFlowLog

`func (o *GroupProperties) SetCreateFlowLog(v bool)`

SetCreateFlowLog sets CreateFlowLog field to given value.

### HasCreateFlowLog

`func (o *GroupProperties) HasCreateFlowLog() bool`

HasCreateFlowLog returns a boolean if a field has been set.

### GetAccessAndManageMonitoring

`func (o *GroupProperties) GetAccessAndManageMonitoring() bool`

GetAccessAndManageMonitoring returns the AccessAndManageMonitoring field if non-nil, zero value otherwise.

### GetAccessAndManageMonitoringOk

`func (o *GroupProperties) GetAccessAndManageMonitoringOk() (*bool, bool)`

GetAccessAndManageMonitoringOk returns a tuple with the AccessAndManageMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageMonitoring

`func (o *GroupProperties) SetAccessAndManageMonitoring(v bool)`

SetAccessAndManageMonitoring sets AccessAndManageMonitoring field to given value.

### HasAccessAndManageMonitoring

`func (o *GroupProperties) HasAccessAndManageMonitoring() bool`

HasAccessAndManageMonitoring returns a boolean if a field has been set.

### GetAccessAndManageCertificates

`func (o *GroupProperties) GetAccessAndManageCertificates() bool`

GetAccessAndManageCertificates returns the AccessAndManageCertificates field if non-nil, zero value otherwise.

### GetAccessAndManageCertificatesOk

`func (o *GroupProperties) GetAccessAndManageCertificatesOk() (*bool, bool)`

GetAccessAndManageCertificatesOk returns a tuple with the AccessAndManageCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageCertificates

`func (o *GroupProperties) SetAccessAndManageCertificates(v bool)`

SetAccessAndManageCertificates sets AccessAndManageCertificates field to given value.

### HasAccessAndManageCertificates

`func (o *GroupProperties) HasAccessAndManageCertificates() bool`

HasAccessAndManageCertificates returns a boolean if a field has been set.

### GetManageDBaaS

`func (o *GroupProperties) GetManageDBaaS() bool`

GetManageDBaaS returns the ManageDBaaS field if non-nil, zero value otherwise.

### GetManageDBaaSOk

`func (o *GroupProperties) GetManageDBaaSOk() (*bool, bool)`

GetManageDBaaSOk returns a tuple with the ManageDBaaS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageDBaaS

`func (o *GroupProperties) SetManageDBaaS(v bool)`

SetManageDBaaS sets ManageDBaaS field to given value.

### HasManageDBaaS

`func (o *GroupProperties) HasManageDBaaS() bool`

HasManageDBaaS returns a boolean if a field has been set.

### GetAccessAndManageDns

`func (o *GroupProperties) GetAccessAndManageDns() bool`

GetAccessAndManageDns returns the AccessAndManageDns field if non-nil, zero value otherwise.

### GetAccessAndManageDnsOk

`func (o *GroupProperties) GetAccessAndManageDnsOk() (*bool, bool)`

GetAccessAndManageDnsOk returns a tuple with the AccessAndManageDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageDns

`func (o *GroupProperties) SetAccessAndManageDns(v bool)`

SetAccessAndManageDns sets AccessAndManageDns field to given value.

### HasAccessAndManageDns

`func (o *GroupProperties) HasAccessAndManageDns() bool`

HasAccessAndManageDns returns a boolean if a field has been set.

### GetManageRegistry

`func (o *GroupProperties) GetManageRegistry() bool`

GetManageRegistry returns the ManageRegistry field if non-nil, zero value otherwise.

### GetManageRegistryOk

`func (o *GroupProperties) GetManageRegistryOk() (*bool, bool)`

GetManageRegistryOk returns a tuple with the ManageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageRegistry

`func (o *GroupProperties) SetManageRegistry(v bool)`

SetManageRegistry sets ManageRegistry field to given value.

### HasManageRegistry

`func (o *GroupProperties) HasManageRegistry() bool`

HasManageRegistry returns a boolean if a field has been set.

### GetManageDataplatform

`func (o *GroupProperties) GetManageDataplatform() bool`

GetManageDataplatform returns the ManageDataplatform field if non-nil, zero value otherwise.

### GetManageDataplatformOk

`func (o *GroupProperties) GetManageDataplatformOk() (*bool, bool)`

GetManageDataplatformOk returns a tuple with the ManageDataplatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageDataplatform

`func (o *GroupProperties) SetManageDataplatform(v bool)`

SetManageDataplatform sets ManageDataplatform field to given value.

### HasManageDataplatform

`func (o *GroupProperties) HasManageDataplatform() bool`

HasManageDataplatform returns a boolean if a field has been set.

### GetAccessAndManageLogging

`func (o *GroupProperties) GetAccessAndManageLogging() bool`

GetAccessAndManageLogging returns the AccessAndManageLogging field if non-nil, zero value otherwise.

### GetAccessAndManageLoggingOk

`func (o *GroupProperties) GetAccessAndManageLoggingOk() (*bool, bool)`

GetAccessAndManageLoggingOk returns a tuple with the AccessAndManageLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageLogging

`func (o *GroupProperties) SetAccessAndManageLogging(v bool)`

SetAccessAndManageLogging sets AccessAndManageLogging field to given value.

### HasAccessAndManageLogging

`func (o *GroupProperties) HasAccessAndManageLogging() bool`

HasAccessAndManageLogging returns a boolean if a field has been set.

### GetAccessAndManageCdn

`func (o *GroupProperties) GetAccessAndManageCdn() bool`

GetAccessAndManageCdn returns the AccessAndManageCdn field if non-nil, zero value otherwise.

### GetAccessAndManageCdnOk

`func (o *GroupProperties) GetAccessAndManageCdnOk() (*bool, bool)`

GetAccessAndManageCdnOk returns a tuple with the AccessAndManageCdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageCdn

`func (o *GroupProperties) SetAccessAndManageCdn(v bool)`

SetAccessAndManageCdn sets AccessAndManageCdn field to given value.

### HasAccessAndManageCdn

`func (o *GroupProperties) HasAccessAndManageCdn() bool`

HasAccessAndManageCdn returns a boolean if a field has been set.

### GetAccessAndManageVpn

`func (o *GroupProperties) GetAccessAndManageVpn() bool`

GetAccessAndManageVpn returns the AccessAndManageVpn field if non-nil, zero value otherwise.

### GetAccessAndManageVpnOk

`func (o *GroupProperties) GetAccessAndManageVpnOk() (*bool, bool)`

GetAccessAndManageVpnOk returns a tuple with the AccessAndManageVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageVpn

`func (o *GroupProperties) SetAccessAndManageVpn(v bool)`

SetAccessAndManageVpn sets AccessAndManageVpn field to given value.

### HasAccessAndManageVpn

`func (o *GroupProperties) HasAccessAndManageVpn() bool`

HasAccessAndManageVpn returns a boolean if a field has been set.

### GetAccessAndManageApiGateway

`func (o *GroupProperties) GetAccessAndManageApiGateway() bool`

GetAccessAndManageApiGateway returns the AccessAndManageApiGateway field if non-nil, zero value otherwise.

### GetAccessAndManageApiGatewayOk

`func (o *GroupProperties) GetAccessAndManageApiGatewayOk() (*bool, bool)`

GetAccessAndManageApiGatewayOk returns a tuple with the AccessAndManageApiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageApiGateway

`func (o *GroupProperties) SetAccessAndManageApiGateway(v bool)`

SetAccessAndManageApiGateway sets AccessAndManageApiGateway field to given value.

### HasAccessAndManageApiGateway

`func (o *GroupProperties) HasAccessAndManageApiGateway() bool`

HasAccessAndManageApiGateway returns a boolean if a field has been set.

### GetAccessAndManageKaas

`func (o *GroupProperties) GetAccessAndManageKaas() bool`

GetAccessAndManageKaas returns the AccessAndManageKaas field if non-nil, zero value otherwise.

### GetAccessAndManageKaasOk

`func (o *GroupProperties) GetAccessAndManageKaasOk() (*bool, bool)`

GetAccessAndManageKaasOk returns a tuple with the AccessAndManageKaas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageKaas

`func (o *GroupProperties) SetAccessAndManageKaas(v bool)`

SetAccessAndManageKaas sets AccessAndManageKaas field to given value.

### HasAccessAndManageKaas

`func (o *GroupProperties) HasAccessAndManageKaas() bool`

HasAccessAndManageKaas returns a boolean if a field has been set.

### GetAccessAndManageNetworkFileStorage

`func (o *GroupProperties) GetAccessAndManageNetworkFileStorage() bool`

GetAccessAndManageNetworkFileStorage returns the AccessAndManageNetworkFileStorage field if non-nil, zero value otherwise.

### GetAccessAndManageNetworkFileStorageOk

`func (o *GroupProperties) GetAccessAndManageNetworkFileStorageOk() (*bool, bool)`

GetAccessAndManageNetworkFileStorageOk returns a tuple with the AccessAndManageNetworkFileStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageNetworkFileStorage

`func (o *GroupProperties) SetAccessAndManageNetworkFileStorage(v bool)`

SetAccessAndManageNetworkFileStorage sets AccessAndManageNetworkFileStorage field to given value.

### HasAccessAndManageNetworkFileStorage

`func (o *GroupProperties) HasAccessAndManageNetworkFileStorage() bool`

HasAccessAndManageNetworkFileStorage returns a boolean if a field has been set.

### GetAccessAndManageAiModelHub

`func (o *GroupProperties) GetAccessAndManageAiModelHub() bool`

GetAccessAndManageAiModelHub returns the AccessAndManageAiModelHub field if non-nil, zero value otherwise.

### GetAccessAndManageAiModelHubOk

`func (o *GroupProperties) GetAccessAndManageAiModelHubOk() (*bool, bool)`

GetAccessAndManageAiModelHubOk returns a tuple with the AccessAndManageAiModelHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageAiModelHub

`func (o *GroupProperties) SetAccessAndManageAiModelHub(v bool)`

SetAccessAndManageAiModelHub sets AccessAndManageAiModelHub field to given value.

### HasAccessAndManageAiModelHub

`func (o *GroupProperties) HasAccessAndManageAiModelHub() bool`

HasAccessAndManageAiModelHub returns a boolean if a field has been set.

### GetAccessAndManageIamResources

`func (o *GroupProperties) GetAccessAndManageIamResources() bool`

GetAccessAndManageIamResources returns the AccessAndManageIamResources field if non-nil, zero value otherwise.

### GetAccessAndManageIamResourcesOk

`func (o *GroupProperties) GetAccessAndManageIamResourcesOk() (*bool, bool)`

GetAccessAndManageIamResourcesOk returns a tuple with the AccessAndManageIamResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndManageIamResources

`func (o *GroupProperties) SetAccessAndManageIamResources(v bool)`

SetAccessAndManageIamResources sets AccessAndManageIamResources field to given value.

### HasAccessAndManageIamResources

`func (o *GroupProperties) HasAccessAndManageIamResources() bool`

HasAccessAndManageIamResources returns a boolean if a field has been set.

### GetCreateNetworkSecurityGroups

`func (o *GroupProperties) GetCreateNetworkSecurityGroups() bool`

GetCreateNetworkSecurityGroups returns the CreateNetworkSecurityGroups field if non-nil, zero value otherwise.

### GetCreateNetworkSecurityGroupsOk

`func (o *GroupProperties) GetCreateNetworkSecurityGroupsOk() (*bool, bool)`

GetCreateNetworkSecurityGroupsOk returns a tuple with the CreateNetworkSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNetworkSecurityGroups

`func (o *GroupProperties) SetCreateNetworkSecurityGroups(v bool)`

SetCreateNetworkSecurityGroups sets CreateNetworkSecurityGroups field to given value.

### HasCreateNetworkSecurityGroups

`func (o *GroupProperties) HasCreateNetworkSecurityGroups() bool`

HasCreateNetworkSecurityGroups returns a boolean if a field has been set.



