# Go REST API client for Netbox

This package provides a Go API client for the Netbox REST API with multiple API versions.

## Overview

The Netbox server does not support API versioning and only provides limited API guarantees, especially for statically 
typed languages like Go and clients generated using the OpenAPI specification rather than crafted by hand. Experience 
showed each new minor version should be considered as major for Go development.

As such, this API client tries to provide support for multiple API versions, each living in a submodule called 
`apiv{MAJOR}_{MINOR}`. The global module version does not mirror the latest Netbox API version supported, and rather 
aims to reflect the stability of the client module itself. This should ensure developers can upgrade at their own pace, 
or even provide support for multiple Netbox versions from the same program.

Do note however that as maintainer free time is limited and Netbox does not provide long-term support for older minor 
versions, only limited support can be provided for those. As Netbox API clients are rather heavy, oldest versions will 
likely have to be removed at some point. The aim for now is to support all versions of the current major Netbox release,
plus the latest minor of the previous major, but this may change depending on how long a major version lasts. Any removal
will lead to a major version increment of this repository, including a module namespace change (ie, 
`github.com/jqueuniet/netbox-go/v2` to `github.com/jqueuniet/netbox-go/v3`).

As the current tooling does not support OpenAPI 2, no support can be provided for Netbox API versions older than 3.5.

All API clients are generated using the [OpenAPI Generator](https://openapi-generator.tech) project. Each client version
has its own generator config to allow changing the way code is generated across API versions and generator releases 
without impacting older client versions.

Schemas are patched before generation to work around bugs in either the generator or the upstream schema.

## Installation

Import the module in your Go codebase.

```sh
go get github.com/jqueuniet/netbox-go
```

Put the package under your project folder and add the following in import:

```go
import netbox "github.com/jqueuniet/netbox-go/apiv4_0"
```

Detailed usage can be found in the README.md under each specific version, may vary from one version to the next but 
should remain consistent as long as the same generator is used.
