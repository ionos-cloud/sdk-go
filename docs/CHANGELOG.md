# CHANGELOG

## v5.1.11 (December, 2021)

### Enhancements:

* added `Password` field to `UserPropertiesPut` to allow user password update

### Fixes:

* fixed sporadic `EOF` error seen from server
* fixed overwriting `https` with `http` for host endpoint that starts with `http`. It now supports both `http` and `https` schemas

### Dependency-updates:

* updated Go version from `1.13` to `1.17`

## 5.1.10 (November, 2021)

### Enhancements:

* added support for `maxResults` query parameter on **GET** requests

## 5.1.9 (October, 2021)

### Enhancements:

* improved code in sync with Sonar Cloud requirements:
  * added global constant `FilterQueryParam`
  * renamed global constant `FORMAT_STRING` to `FormatStringErr`

## 5.1.8 (October, 2021)

### Enhancements:

* added support for filter query parameters on **GET** requests: `Filter()`
* added support for overwriting the host-url value via `IONOS_API_URL` environment variable
* added `gofmt` check's results

## 5.1.7 (September, 2021)

### Features:

* removed Public parameter from **_KubernetesClusterProperties_**
* removed GatewayIp parameter from **_KubernetesNodePoolProperties_**
    
## 5.1.6 (September, 2021)

### Documentation:

* changed descriptions
    
## 5.1.5 (August, 2021)

### Features:

* added new parameter on APIResponse:
    * RequestTime
    
## 5.1.4 (July, 2021)

### Features:

* added new parameters on **_KubernetesClusterProperties_**, **_KubernetesClusterPropertiesForPost_** and **_KubernetesClusterPropertiesForPut_**:
    * `ApiSubnetAllowList`
    * `S3Buckets`
* added new required parameter on **_KubernetesNodePoolLan_**: `id`