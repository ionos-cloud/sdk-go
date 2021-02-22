# Introduction

## Overview

The IONOS Cloud SDK for GO provides you with access to the IONOS Cloud API. The client library supports both simple and complex requests. It is designed for developers who are building applications in GO . The SDK for GO wraps the IONOS Cloud API. All API operations are performed over SSL and authenticated using your IONOS Cloud portal credentials. The API can be accessed within an instance running in IONOS Cloud or directly over the Internet from any application that can send an HTTPS request and receive an HTTPS response.

## Getting Started

Before you begin you will need to have  [signed-up](https://www.ionos.com/enterprise-cloud/signup)  for a IONOS Cloud account. The credentials you setup during sign-up will be used to authenticate against the API.

#### Installation

Install the Go language from from the official  [Go installation](https://golang.org/doc/install)  guide.

The  `GOPATH`  environment variable specifies the location of your Go workspace. It is likely the only environment variable you will need to set when developing Go code. This is an example of pointing to a workspace configured under your home directory:

```
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

The following  `go`  command will download  `sdk-go`  to your configured  `GOPATH`:

```go
go get "github.com/ionos-cloud/sdk-go"
```

The source code of the package will be located here:

```
$GOPATH/src/github.com/ionos-cloud/sdk-go
```

Create main package file  _example.go_:

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

#### Authentication

The username and passwod or the authentication token can be manually specified when initializing the sdk client:
```go
client := ionoscloud.NewAPIClient(ionoscloud.NewConfiguration(username, password, token))
```
Environment variables can also be used. The sdk uses the following variables:

-   IONOS_USERNAME - to specify the username used to login
-   IONOS_PASSWORD - to specify the password
-   IONOS_TOKEN - if an authentication token is being used

In this case, the client configuration needs to be initialized using  `NewConfigurationFromEnv()`

```go
client := ionoscloud.NewAPIClient(ionoscloud.NewConfigurationFromEnv())
```

**Caution**: You will want to ensure you follow security best practices when using credentials within your code or stored in a file.



## Feature Reference

The IONOS Cloud SDK for GO aims to offer access to all resources in the IONOS Cloud API and also offers some additional features that make the integration easier:

* authentication for API calls
* handling of asynchronous requests 

## FAQ

1. How can I open a bug/feature request? 

Bugs & feature requests can be open on the repository issues: [https://github.com/ionos-cloud/sdk-go/issues/new/choose](https://github.com/ionos-cloud/sdk-go/issues/new/choose)

1. Can I contribute to the GO SDK?

Pure SDKs are automatically generated using OpenAPI Generator and don’t support manual changes. If you need changes please open an issue and we’ll try to take care of it.

