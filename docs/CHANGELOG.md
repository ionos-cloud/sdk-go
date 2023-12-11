# CHANGELOG

## 6.1.10 (December, 2023)
### Features:
- Add `proxyProtocol` parameter for ALB and NLB targets

### Fixes:
- Allow setting `Host` and `Scheme` when creating a client with a `serverUrl` (@maxbischoff)

### Enhancements
- Remove usage of deprecated `ioutil` package (@avorima)

### Documentation
- Move docs for Url with context to `api_doc.mustache`


## 6.1.9 (September, 2023)
### Features:
- Support injecting `x-contract-number` header through environment variable.


## 6.1.8 (July, 2023)
### Features:
Added support for **IPv6**:

- New parameter on `DatacenterProperties`: `Ipv6CidrBlock`
- New parameter on `LanProperties` and `LanPropertiesPost`: `Ipv6CidrBlock`
- New parameters on `NicProperties`: `Dhcpv6`, `Ipv6CidrBlock` and `Ipv6Ips`

More details about IPv6 configuration can be found [here](https://docs.ionos.com/cloud/compute-engine/networks/ipv6).

### Fixes:
- Go client backoff now respects context cancellation

**Full Changelog**: https://github.com/ionos-cloud/sdk-go/compare/v6.1.7...v6.1.8


## 6.1.7 (June, 2023)
### Features:
- New licenceType supported: **RHEL**

## 6.1.5 (March, 2023)
### Features: 
- Added generic functions for utility package

### Fixes: 
- Removed swagger

### Updated: 
- Library dependencies


## 6.1.4 \(January 26th, 2022\)
### Features:
- Added `placementGroupId` and `vnet` parameters.
 
### Fixes:
- Allow multiple values to be set for the same filter key

## 6.1.3 \(August 30th, 2022\)
### Fixes:
- Fix return type of NewGenericOpenAPIError

## 6.1.2 \(August 4th, 2022\)
### Fixes :
- Changed from `manageDbaas` to `manageDBaaS` field in `model_group_properties.go` : provides privilege for a group to manage DBaaS related functionality. Admin users already have it enabled by default.
- Issue #26


## 6.1.1 \(July 14th, 2022\)

### Features :
- Added `manageDbaas` field in `model_group_properties.go` : provides privilege for a group to manage DBaaS related functionality. Admin users already here this enabled by default.
- Added `deleteVolumes` to `DatacentersServersDelete` function: If true, all attached storage volumes will also be deleted.
- Added `bootOrder` to `model_volume_properties` : Determines whether the volume will be used as a boot volume. Set to &#x60;NONE&#x60;, the volume will not be used as boot volume. Set to &#x60;PRIMARY&#x60;, the volume will be used as boot volume and all other volumes must be set to &#x60;NONE&#x60;. Set to &#x60;AUTO&#x60; or &#x60;null&#x60; requires all volumes to be set to &#x60;AUTO&#x60; or &#x60;null&#x60;; this will use the legacy behavior, which is to use the volume as a boot volume only if there are no other volumes or cdrom devices. | [optional] [default to 'AUTO'] |
- Logger interface with log levels for the sdk. Allows user to inject it's own logger that implements Printf. More information [here](https://github.com/ionos-cloud/sdk-go#debugging)
- Added helper function `HttpNotFound` in `response.go`
## 6.1.0 \(June 16th, 2022\)

### Enhancements:

* added Application Load Balancer and Target Group, 18 new models and 2 new apis

## v6.0.1 (January, 2022)

### Enhancements:
- new parameter on KubernetesClusterProperties, KubernetesClusterPropertiesForPost: **public**
- new parameter on KubernetesNodePoolProperties, KubernetesNodePoolPropertiesForPost: **gatewayIp**
- allow fields that are explicitly nullable to be sent with null values
- updated oauth import
- depth parameter: 
  - new method **SetDepth** on **configuration** that allows setting a value for the **depth** parameter on all the API calls for the client
  - changed the default value of depth parameter from 10 to 0
  - this affects the amount of data that is returned on GET methods. If there is a need to change the value, please do it with caution and only if it is needed. 

## v6.0.0 (December, 2021)

### Enhancements:

* changed structure `ApiK8sNodepoolsPostRequest`, changed name of `kubernetesNodePool` field from `KubernetesNodePool` to `KubernetesNodePoolForPost`
* changed structure `ApiK8sNodepoolsPutRequest`, changed type of `KubernetesNodePoolForPut` field from `kubernetesNodePoolForPut` to `kubernetesNodePool`

### Fixes:

* fixed sporadic EOF error seen from server
* fixed overwriting https with `http` for host endpoint that starts with `http`. It now allows user to enter a url that starts with `http`
* added `Password` field to `UserPropertiesPut` to allow user password update

## v6.0.0-beta.9 (November, 2021)

### Enhancements:

* added support for `maxResults` query parameter on **GET** requests

### Features:

* renamed Apis services: `LansApi` to `LANsApi`
* renamed `BackupUnitProperties` to `BackupUnit` on `BackupunitsPatch` method 

### Documentation:

* updated descriptions 

## v6.0.0-beta.8 (October, 2021)

### Enhancements:

* improved code in sync with Sonar Cloud requirements:
    * added global constant `FilterQueryParam`
    * renamed global constant `FORMAT_STRING` to `FormatStringErr`

## v6.0.0-beta.7 (October, 2021)

### Enhancements:

* added support for filter query parameters on **GET** requests: `Filter()`
* added support for overwriting the host-url value via `IONOS_API_URL` environment variable
* added `gofmt` check's results

## v6.0.0-beta.6 (September, 2021)

### Features:

* added Offset, Limit and Links to **_IpBlocks_**
* removed Public parameter from **_KubernetesClusterProperties_**
* removed GatewayIp parameter from **_KubernetesNodePoolProperties_**

## v6.0.0-beta.5 (September, 2021)

### Features:

* added imageAlias to **_VolumeProperties_**
* added Offset, Limit and Links to **_User_**
* removed gatewayIp parameter from **_KubernetesClusterProperties_**
* added GatewayIp parameter from **_KubernetesNodePoolProperties_**

