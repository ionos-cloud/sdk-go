# CHANGELOG

## v5.1.0-beta.1 - May 07, 2021

- The following data types where changed for user operations:
  * `User` type was replaced by `UserPost` in POST operations
  * `User` type was replaced by `UserPut` in PUT operations

- The following data types where changed for k8s cluster operations:
  * `KubernetesCluster` type was replaced by `KubernetesClusterForPost` in POST operations
  * `KubernetesCluster` type was replaced by `KubernetesClusterForPut` in PUT operations

- The following data types where changed for k8s nodepool operations:
  * `KubernetesNodePool` type was replaced by `KubernetesNodePoolForPost` in POST operations
  * `KubernetesNodePool` type was replaced by `KubernetesNodePoolForPUT` in PUT operations both as input
     and as response type

## v5.0.4 - May 04, 2021

- `ImageAliases` and `CloudInit` attributes added to the `ImageProperties` model
- `Public` and `GatewayIp` attributes added to the `KubernetesClusterProperties` model
- `PrivateIP` attribute added to the `KubernetesNodeProperties` model
- `UserData` attribute added to the `VolumeProperties` model
- `k8sNodePoolUuid` and `k8sClusterUuid` were added to the `IpConsumer` model
- new filters supported in `RequestApi`: `filter.createdDate`, `filter.createdBy`, `filter.etag`,
`filter.requestStatus`, `filter.method`, `filter.headers`, `filter.body`, `filter.url`
