# CHANGELOG

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

