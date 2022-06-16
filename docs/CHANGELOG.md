# CHANGELOG

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

