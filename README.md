# Overview

This repository contains the published SDKs and [SDK releases](https://github.com/stackitcloud/stackit-sdk-go/releases/).
The modules are structured into a [core module](https://github.com/stackitcloud/stackit-sdk-go/tree/main/core) with service clients, authentication and shared functionality as well as the different STACKIT [services](https://github.com/stackitcloud/stackit-sdk-go/tree/main/services).
The usage of the SDK is shown in some [examples](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples).

This repo is in BETA state.

# Getting started

Requires `Go 1.18` or higher.

To download the `core` module:

```
go mod download github.com/stackitcloud/stackit-sdk-go/core
```

To download the `services/dns` module:

```
go mod download github.com/stackitcloud/stackit-sdk-go/services/dns
```

# Examples

This is an example on how to do create a client and interact with the STACKIT DNS service for reading and creating DNS zones. As prerequisite, you need a STACKIT project with its project ID.
The setup of the authentication is describe below in section [Authentication](#authentication) in more detail.

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

func main() {
	// Specify the project ID
	projectId := "PROJECT_ID"

	// Create a new API client, that uses default authentication and configuration
	dnsClient, err := dns.NewAPIClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
		os.Exit(1)
	}

	// Get the DNS Zones for your project
	var getZoneResp *dns.ZonesResponse
	getZoneResp, err = dnsClient.GetZones(context.Background(), projectId).Execute()

	// Get only active DNS Zones for your project by adding the filter "ActiveEq(true)" to the call. More filters are available and can be chained.
	// dnsRespGetZones, err := dnsClient.ZoneApi.GetZones(context.Background(), projectId).ActiveEq(true).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `ZoneApi.GetZones`: %v\n", err)
	} else {
		fmt.Printf("[DNS API] Number of zones: %v\n", len(getZoneResp.Zones))
	}

	// Create a DNS Zone
	createZonePayload := dns.CreateZonePayload{
		Name:    "myZone",
		DnsName: "testZone.com",
	}
	var createZoneResp *dns.ZoneResponse
	createZoneResp, err = dnsClient.CreateZone(context.Background(), projectId).CreateZonePayload(createZonePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `ZoneApi.CreateZone`: %v\n", err)
	} else {
		var createdZone dns.Zone = createZoneResp.Zone
		fmt.Printf("[DNS API] Created zone \"%s\" with DNS name \"%s\" and zone id \"%s\".\n", createdZone.Name, createdZone.DnsName, createdZone.Id)
	}

	// Get a record set of a DNS zone.
	var recordSetResp *dns.RecordSetResponse
	recordSetResp, err = dnsClient.GetRecordSet(context.Background(), projectId, "zoneId", "recordSetId").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `GetRecordSet`: %v\n", err)
	} else {
		fmt.Printf("[DNS API] Got record set with name \"%s\".\n", recordSetResp.Rrset.Name)
	}
}

```

More examples on other service and authentication possibilities can be found in the [example folder](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples).

Please note that the naming of the data models, functions and the SDK API is subject to ongoing changes.

## Authentication

There are multiple ways to authenticate. Currently, the _token-flow_ is supported: The SDK will first try to find a token in the `STACKIT_SERVICE_ACCOUNT_TOKEN` env var. If not present, it will
check the credentials file located in the path defined by the `STACKIT_CREDENTIALS_PATH` env var, if specified,
or in `$HOME/.stackit/credentials.json` as a fallback. If the token is found, all the requests are authenticated using that token.

In the example below, the case for no authentication and the case for a explicit token provision are shown.

```go
(...)

import (
	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

(...)

// Create an unauthenticated API client
client, err := dns.NewAPIClient(config.WithoutAuthentication())
if err != nil {
    fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
    os.Exit(1)
}

```

```go
(...)

import (
	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

(...)

// Create a new API client, that will authenticate using the provided bearer token
token := "TOKEN"
client, err := dns.NewAPIClient(config.WithToken(token))
if err != nil {
    fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
    os.Exit(1)
}
```

## Reporting issues

If you encounter any issues or have suggestions for improvements, please open an issue in the repository.

## Contribute

Your contribution is welcome! For more details on how to contribute, refer to our [Contribution Guide](./CONTRIBUTION.md).

## License

Apache 2.0
