# Introduction

## Overview

**NOTE:
_Please consider using SDK Go v6 releases that are using IONOS Cloud API Version 6, the latest stable API version. New features and improvements will be integrated in IONOS Cloud API Version 6._**

The IONOS Cloud SDK for GO provides you with access to the IONOS Cloud API. The client library supports both simple and complex requests. It is designed for developers who are building applications in GO . The SDK for GO wraps the IONOS Cloud API. All API operations are performed over SSL and authenticated using your IONOS Cloud portal credentials. The API can be accessed within an instance running in IONOS Cloud or directly over the Internet from any application that can send an HTTPS request and receive an HTTPS response.

## Getting Started

An IONOS account is required for access to the Cloud API; credentials from your registration are used to authenticate against the IONOS Cloud API.

### Installation

Install the Go language from from the official [Go installation](https://golang.org/doc/install) guide.

The `GOPATH` environment variable specifies the location of your Go workspace. It is likely the only environment variable you will have to set when developing Go code. This is an example of pointing to a workspace configured under your home directory:

```text
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

The following `go` command will download `sdk-go` to your configured `GOPATH`:

```go
go get "github.com/ionos-cloud/sdk-go"
```

The source code of the package will be located here:

```text
$GOPATH/src/github.com/ionos-cloud/sdk-go
```

Create main package file _example.go_:

```go
package main

import (
    "fmt"
)

func main() {
}
```

Include the IONOS Cloud SDK for Go under the list of imports.

```go
import(
    "fmt"

    "github.com/ionos-cloud/sdk-go"
)
```

### Authentication

The username and password or the authentication token can be manually specified when initializing the SDK client:

```go
client := ionoscloud.NewAPIClient(ionoscloud.NewConfiguration(username, password, token, apiUrl))
```

Environment variables can also be used; the SDK uses the following variables:

* IONOS\_USERNAME - to specify the username used to login
* IONOS\_PASSWORD - to specify the password
* IONOS\_TOKEN - if an authentication token is being used
* IONOS\_API\_URL - to overwrite the API endpoint: `api.ionos.com` - if it is not set, the default value will be used

In this case, the client configuration must be initialized using `NewConfigurationFromEnv()`

```go
client := ionoscloud.NewAPIClient(ionoscloud.NewConfigurationFromEnv())
```

{% hint style="danger" %}
**Warning**: Make sure to follow the Information Security Best Practices when using credentials within your code or storing them in a file.
{% endhint %}

### Environment Variables

Environment Variable | Description
--- | --- 
`IONOS_USERNAME` | Specify the username used to login, to authenticate against the IONOS Cloud API | 
`IONOS_PASSWORD` | Specify the password used to login, to authenticate against the IONOS Cloud API | 
`IONOS_TOKEN` | Specify the token used to login, if a token is being used instead of username and password |
`IONOS_API_URL` | Specify the API URL. It will overwrite the API endpoint default value `api.ionos.com`. Note: the host URL does not contain the `/cloudapi/v5` path, so it should _not_ be included in the `IONOS_API_URL` environment variable | 

### Depth

Many of the _List_ or _Get_ operations will accept an optional _depth_ argument. Setting this to a value between 0 and 5 affects the amount of data that is returned. The details returned vary depending on the resource being queried, but it generally follows this pattern. By default, the SDK sets the _depth_ argument to the maximum value.

| Depth | Description |
| :--- | :--- |
| 0 | Only direct properties are included. Children are not included. |
| 1 | Direct properties and children's references are returned. |
| 2 | Direct properties and children's properties are returned. |
| 3 | Direct properties, children's properties, and descendants' references are returned. |
| 4 | Direct properties, children's properties, and descendants' properties are returned. |
| 5 | Returns all available properties. |


#### How to set Depth parameter:
* On the configuration level:
```go  
configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "URL")  
configuration.SetDepth(5)  
```
Using this method, the depth parameter will be set **on all the API calls**.

*  When calling a method:
```go
request := apiClient.DataCenterApi.DatacentersGet(context.Background()).Depth(1)
```
Using this method, the depth parameter will be set **on the current API call**.

* Using the default value:

If the depth parameter is not set, it will have the default value from the API that can be found [here](https://api.ionos.com/cloudapi/v5/swagger.json).

> Note: The priority for setting the depth parameter is: *set on function call > set on configuration method > set using the default value from the API*

### Pretty

The operations will also accept an optional _pretty_ argument. Setting this to a value of `true` or `false` controls whether the response is pretty-printed \(with indentation and new lines\). By default, the SDK sets the _pretty_ argument to `true`.

### Changing the base URL

Base URL for the HTTP operation can be changed by using the following function:

```go
requestProperties.SetURL("https://api.ionos.com/cloudapi/v5")
```

## Feature Reference

The IONOS Cloud SDK for GO aims to offer access to all resources in the IONOS Cloud API, and has additional features to make integration easier:

* Authentication for API calls
* Asynchronous request handling 

## FAQ

1. How can I open a bug report/feature request? 

Bug reports and feature requests can be opened in the Issues repository: [https://github.com/ionos-cloud/sdk-go/issues/new/choose](https://github.com/ionos-cloud/sdk-go/issues/new/choose)

1. Can I contribute to the GO SDK?

Pure SDKs are automatically generated using OpenAPI Generator and don’t support manual changes. If you require changes, please open an issue and we will try to address it.

## Debugging

If you want to see the API call request and response messages, you need to set the Debug field in the Configuration struct:

```golang
package main

import "github.com/ionos-cloud/sdk-go/v5"

func main() {
    // create your configuration. replace username, password, token and url with correct values, or use NewConfigurationFromEnv()
    // if you have set your env variables as explained above
    cfg := ionoscloud.NewConfiguration("username", "password", "token", "hostUrl")
    // enable request and response logging
    cfg.Debug = true
    // create you api client with the configuration
    apiClient := ionoscloud.NewAPIClient(cfg)
}
```
#### Note: We recommend you only set this field for debugging purposes. Disable it in your production environments because it can log sensitive data. It logs the full request and response without encryption, even for an HTTPS call. Verbose request and response logging can also significantly impact your application’s performance.
