# ResourceLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoresPerServer** | **int32** | maximum number of cores per server | 
**CoresPerContract** | **int32** | maximum cores per contract | 
**CoresProvisioned** | **int32** | number of cores provisioned | 
**RamPerServer** | **int32** | maximum ram per server | 
**RamPerContract** | **int32** | maximum ram per contract | 
**RamProvisioned** | **int32** | ram provisioned | 
**HddLimitPerVolume** | **int64** | hdd limit per volume | 
**HddLimitPerContract** | **int64** | hdd limit per contract | 
**HddVolumeProvisioned** | **int64** | hdd volume provisioned | 
**SsdLimitPerVolume** | **int64** | ssd limit per volume | 
**SsdLimitPerContract** | **int64** | ssd limit per contract | 
**SsdVolumeProvisioned** | **int64** | ssd volume provisioned | 
**ReservableIps** | **int32** | total reservable ip limit of the customer | 
**ReservedIpsOnContract** | **int32** | reserved ips on a contract | 
**ReservedIpsInUse** | **int32** | reserved ips in use | 
**K8sClusterLimitTotal** | **int32** | k8s clusters total limit | 
**K8sClustersProvisioned** | **int32** | k8s clusters provisioned | 

## Methods

### NewResourceLimits

`func NewResourceLimits(coresPerServer int32, coresPerContract int32, coresProvisioned int32, ramPerServer int32, ramPerContract int32, ramProvisioned int32, hddLimitPerVolume int64, hddLimitPerContract int64, hddVolumeProvisioned int64, ssdLimitPerVolume int64, ssdLimitPerContract int64, ssdVolumeProvisioned int64, reservableIps int32, reservedIpsOnContract int32, reservedIpsInUse int32, k8sClusterLimitTotal int32, k8sClustersProvisioned int32, ) *ResourceLimits`

NewResourceLimits instantiates a new ResourceLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLimitsWithDefaults

`func NewResourceLimitsWithDefaults() *ResourceLimits`

NewResourceLimitsWithDefaults instantiates a new ResourceLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoresPerServer

`func (o *ResourceLimits) GetCoresPerServer() int32`

GetCoresPerServer returns the CoresPerServer field if non-nil, zero value otherwise.

### GetCoresPerServerOk

`func (o *ResourceLimits) GetCoresPerServerOk() (*int32, bool)`

GetCoresPerServerOk returns a tuple with the CoresPerServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerServer

`func (o *ResourceLimits) SetCoresPerServer(v int32)`

SetCoresPerServer sets CoresPerServer field to given value.


### GetCoresPerContract

`func (o *ResourceLimits) GetCoresPerContract() int32`

GetCoresPerContract returns the CoresPerContract field if non-nil, zero value otherwise.

### GetCoresPerContractOk

`func (o *ResourceLimits) GetCoresPerContractOk() (*int32, bool)`

GetCoresPerContractOk returns a tuple with the CoresPerContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerContract

`func (o *ResourceLimits) SetCoresPerContract(v int32)`

SetCoresPerContract sets CoresPerContract field to given value.


### GetCoresProvisioned

`func (o *ResourceLimits) GetCoresProvisioned() int32`

GetCoresProvisioned returns the CoresProvisioned field if non-nil, zero value otherwise.

### GetCoresProvisionedOk

`func (o *ResourceLimits) GetCoresProvisionedOk() (*int32, bool)`

GetCoresProvisionedOk returns a tuple with the CoresProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresProvisioned

`func (o *ResourceLimits) SetCoresProvisioned(v int32)`

SetCoresProvisioned sets CoresProvisioned field to given value.


### GetRamPerServer

`func (o *ResourceLimits) GetRamPerServer() int32`

GetRamPerServer returns the RamPerServer field if non-nil, zero value otherwise.

### GetRamPerServerOk

`func (o *ResourceLimits) GetRamPerServerOk() (*int32, bool)`

GetRamPerServerOk returns a tuple with the RamPerServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamPerServer

`func (o *ResourceLimits) SetRamPerServer(v int32)`

SetRamPerServer sets RamPerServer field to given value.


### GetRamPerContract

`func (o *ResourceLimits) GetRamPerContract() int32`

GetRamPerContract returns the RamPerContract field if non-nil, zero value otherwise.

### GetRamPerContractOk

`func (o *ResourceLimits) GetRamPerContractOk() (*int32, bool)`

GetRamPerContractOk returns a tuple with the RamPerContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamPerContract

`func (o *ResourceLimits) SetRamPerContract(v int32)`

SetRamPerContract sets RamPerContract field to given value.


### GetRamProvisioned

`func (o *ResourceLimits) GetRamProvisioned() int32`

GetRamProvisioned returns the RamProvisioned field if non-nil, zero value otherwise.

### GetRamProvisionedOk

`func (o *ResourceLimits) GetRamProvisionedOk() (*int32, bool)`

GetRamProvisionedOk returns a tuple with the RamProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamProvisioned

`func (o *ResourceLimits) SetRamProvisioned(v int32)`

SetRamProvisioned sets RamProvisioned field to given value.


### GetHddLimitPerVolume

`func (o *ResourceLimits) GetHddLimitPerVolume() int64`

GetHddLimitPerVolume returns the HddLimitPerVolume field if non-nil, zero value otherwise.

### GetHddLimitPerVolumeOk

`func (o *ResourceLimits) GetHddLimitPerVolumeOk() (*int64, bool)`

GetHddLimitPerVolumeOk returns a tuple with the HddLimitPerVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddLimitPerVolume

`func (o *ResourceLimits) SetHddLimitPerVolume(v int64)`

SetHddLimitPerVolume sets HddLimitPerVolume field to given value.


### GetHddLimitPerContract

`func (o *ResourceLimits) GetHddLimitPerContract() int64`

GetHddLimitPerContract returns the HddLimitPerContract field if non-nil, zero value otherwise.

### GetHddLimitPerContractOk

`func (o *ResourceLimits) GetHddLimitPerContractOk() (*int64, bool)`

GetHddLimitPerContractOk returns a tuple with the HddLimitPerContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddLimitPerContract

`func (o *ResourceLimits) SetHddLimitPerContract(v int64)`

SetHddLimitPerContract sets HddLimitPerContract field to given value.


### GetHddVolumeProvisioned

`func (o *ResourceLimits) GetHddVolumeProvisioned() int64`

GetHddVolumeProvisioned returns the HddVolumeProvisioned field if non-nil, zero value otherwise.

### GetHddVolumeProvisionedOk

`func (o *ResourceLimits) GetHddVolumeProvisionedOk() (*int64, bool)`

GetHddVolumeProvisionedOk returns a tuple with the HddVolumeProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddVolumeProvisioned

`func (o *ResourceLimits) SetHddVolumeProvisioned(v int64)`

SetHddVolumeProvisioned sets HddVolumeProvisioned field to given value.


### GetSsdLimitPerVolume

`func (o *ResourceLimits) GetSsdLimitPerVolume() int64`

GetSsdLimitPerVolume returns the SsdLimitPerVolume field if non-nil, zero value otherwise.

### GetSsdLimitPerVolumeOk

`func (o *ResourceLimits) GetSsdLimitPerVolumeOk() (*int64, bool)`

GetSsdLimitPerVolumeOk returns a tuple with the SsdLimitPerVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdLimitPerVolume

`func (o *ResourceLimits) SetSsdLimitPerVolume(v int64)`

SetSsdLimitPerVolume sets SsdLimitPerVolume field to given value.


### GetSsdLimitPerContract

`func (o *ResourceLimits) GetSsdLimitPerContract() int64`

GetSsdLimitPerContract returns the SsdLimitPerContract field if non-nil, zero value otherwise.

### GetSsdLimitPerContractOk

`func (o *ResourceLimits) GetSsdLimitPerContractOk() (*int64, bool)`

GetSsdLimitPerContractOk returns a tuple with the SsdLimitPerContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdLimitPerContract

`func (o *ResourceLimits) SetSsdLimitPerContract(v int64)`

SetSsdLimitPerContract sets SsdLimitPerContract field to given value.


### GetSsdVolumeProvisioned

`func (o *ResourceLimits) GetSsdVolumeProvisioned() int64`

GetSsdVolumeProvisioned returns the SsdVolumeProvisioned field if non-nil, zero value otherwise.

### GetSsdVolumeProvisionedOk

`func (o *ResourceLimits) GetSsdVolumeProvisionedOk() (*int64, bool)`

GetSsdVolumeProvisionedOk returns a tuple with the SsdVolumeProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdVolumeProvisioned

`func (o *ResourceLimits) SetSsdVolumeProvisioned(v int64)`

SetSsdVolumeProvisioned sets SsdVolumeProvisioned field to given value.


### GetReservableIps

`func (o *ResourceLimits) GetReservableIps() int32`

GetReservableIps returns the ReservableIps field if non-nil, zero value otherwise.

### GetReservableIpsOk

`func (o *ResourceLimits) GetReservableIpsOk() (*int32, bool)`

GetReservableIpsOk returns a tuple with the ReservableIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservableIps

`func (o *ResourceLimits) SetReservableIps(v int32)`

SetReservableIps sets ReservableIps field to given value.


### GetReservedIpsOnContract

`func (o *ResourceLimits) GetReservedIpsOnContract() int32`

GetReservedIpsOnContract returns the ReservedIpsOnContract field if non-nil, zero value otherwise.

### GetReservedIpsOnContractOk

`func (o *ResourceLimits) GetReservedIpsOnContractOk() (*int32, bool)`

GetReservedIpsOnContractOk returns a tuple with the ReservedIpsOnContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpsOnContract

`func (o *ResourceLimits) SetReservedIpsOnContract(v int32)`

SetReservedIpsOnContract sets ReservedIpsOnContract field to given value.


### GetReservedIpsInUse

`func (o *ResourceLimits) GetReservedIpsInUse() int32`

GetReservedIpsInUse returns the ReservedIpsInUse field if non-nil, zero value otherwise.

### GetReservedIpsInUseOk

`func (o *ResourceLimits) GetReservedIpsInUseOk() (*int32, bool)`

GetReservedIpsInUseOk returns a tuple with the ReservedIpsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpsInUse

`func (o *ResourceLimits) SetReservedIpsInUse(v int32)`

SetReservedIpsInUse sets ReservedIpsInUse field to given value.


### GetK8sClusterLimitTotal

`func (o *ResourceLimits) GetK8sClusterLimitTotal() int32`

GetK8sClusterLimitTotal returns the K8sClusterLimitTotal field if non-nil, zero value otherwise.

### GetK8sClusterLimitTotalOk

`func (o *ResourceLimits) GetK8sClusterLimitTotalOk() (*int32, bool)`

GetK8sClusterLimitTotalOk returns a tuple with the K8sClusterLimitTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterLimitTotal

`func (o *ResourceLimits) SetK8sClusterLimitTotal(v int32)`

SetK8sClusterLimitTotal sets K8sClusterLimitTotal field to given value.


### GetK8sClustersProvisioned

`func (o *ResourceLimits) GetK8sClustersProvisioned() int32`

GetK8sClustersProvisioned returns the K8sClustersProvisioned field if non-nil, zero value otherwise.

### GetK8sClustersProvisionedOk

`func (o *ResourceLimits) GetK8sClustersProvisionedOk() (*int32, bool)`

GetK8sClustersProvisionedOk returns a tuple with the K8sClustersProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClustersProvisioned

`func (o *ResourceLimits) SetK8sClustersProvisioned(v int32)`

SetK8sClustersProvisioned sets K8sClustersProvisioned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


