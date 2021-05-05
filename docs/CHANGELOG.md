# CHANGELOG

## v5.0.4 - May 04, 2020

- `ImageAliases` and `CloudInit` attributes added to the `ImageProperties` model
- `Public` and `GatewayIp` attributes added to the `KubernetesClusterProperties` model
- `PrivateIP` attribute added to the `KubernetesNodeProperties` model
- `UserData` attribute added to the `UserData` model
- `k8sNodePoolUuid` and `k8sClusterUuid` were added to the `IpConsumer` model
- new filters supported in `RequestApi`: `filter.createdDate`, `filter.createdBy`, `filter.etag`,
`filter.requestStatus`, `filter.method`, `filter.headers`, `filter.body`, `filter.url`
