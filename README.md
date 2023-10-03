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

To authenticate to the SDK, you will need a service account. Create it in the STACKIT Portal an assign it the necessary permissions, e.g. project.owner. There are multiple ways to authenticate:

- Token flow
- Key flow (recommended)

When setting up authentication, the SDK will always try to use the key flow first and search for credentials in several locations, following a specific order:

1. Explicit configuration, e.g. by using the option `config.WithServiceAccountKeyPath("path/to/sa_key.json")`
2. Environment variable, e.g. by setting `STACKIT_SERVICE_ACCOUNT_KEY_PATH`
3. Credentials file

   The SDK will check the credentials file located in the path defined by the `STACKIT_CREDENTIALS_PATH` env var, if specified,
   or in `$HOME/.stackit/credentials.json` as a fallback.
   The credentials should be set using the same name as the environmnet variables. Example:

   ```json
   {
     "STACKIT_SERVICE_ACCOUNT_TOKEN": "foo_token",
     "STACKIT_SERVICE_ACCOUNT_KEY_PATH": "path/to/sa_key.json",
     "STACKIT_PRIVATE_KEY_PATH": "path/to/private_key.pem"
   }
   ```

Check the [authentication example](examples/authentication/authentication.go) for more details.

### Key flow

To use the key flow, you need to have a service account key and an RSA key-pair. To configure it, follow this steps:

1. In the Portal, go to `Service Account -> Service Account Keys` and create a key.
   - You can create your own RSA key-pair or have the Portal generate one for you.
2. Save the content of the service account key and the corresponding private key by copying them or saving them in a file.
3. Configure the service account key and private key for authentication in the SDK:
   - using the configuration options: `config.WithServiceAccountKey` or `config.WithServiceAccountKeyPath`, `config.WithPrivateKey` or `config.WithPrivateKeyPath`
   - setting environment variables: `STACKIT_SERVICE_ACCOUNT_KEY` or `STACKIT_SERVICE_ACCOUNT_KEY_PATH`, `STACKIT_PRIVATE_KEY` or `STACKIT_PRIVATE_KEY_PATH`
   - setting them in the credentials file (see above)
4. The SDK will search for the keys and, if valid, will use them to get access and refresh tokens which will be used to authenticate all the requests.

### Token flow

Using this flow is less secure since the token is long-lived. You can provide the token in several ways:

1. Using the configuration option `config.WithToken`
2. Setting the environment variable `STACKIT_SERVICE_ACCOUNT_TOKEN`
3. Setting it in the credentials file (see above)

## Reporting issues

If you encounter any issues or have suggestions for improvements, please open an issue in the repository.

## Contribute

Your contribution is welcome! For more details on how to contribute, refer to our [Contribution Guide](./CONTRIBUTION.md).

## License

Apache 2.0
