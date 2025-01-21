<div align="center">
<br>
<img src=".github/images/stackit-logo.png" alt="STACKIT logo" width="50%"/>
<br>
<br>
</div>

# STACKIT SDK for Go

[![GitHub License](https://img.shields.io/github/license/stackitcloud/stackit-sdk-go)](https://www.apache.org/licenses/LICENSE-2.0)

This repository contains the published SDKs and [SDK releases](https://github.com/stackitcloud/stackit-sdk-go/releases/).
The modules are structured into a [core module](https://github.com/stackitcloud/stackit-sdk-go/tree/main/core) with service clients, authentication and shared functionality as well as the different STACKIT [services](https://github.com/stackitcloud/stackit-sdk-go/tree/main/services).
The usage of the SDK is shown in some [examples](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples).

## Getting started

Requires `Go 1.18` or higher.

To download the `core` module:

```
go mod download github.com/stackitcloud/stackit-sdk-go/core
```

To download the `services/dns` module:

```
go mod download github.com/stackitcloud/stackit-sdk-go/services/dns
```

## Examples

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

More examples on other services, configuration and authentication possibilities can be found in the [examples folder](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples).

## Authentication

To authenticate with the SDK, you need a [service account](https://docs.stackit.cloud/stackit/en/service-accounts-134415819.html) with appropriate permissions (e.g., `project.owner`). This service account can for example be created in the Portal, please check the documentation for further details.

### Authentication Methods

The SDK supports two authentication methods:

1. **Key Flow** (Recommended)

   - Uses RSA key-pair based authentication
   - Provides better security through short-lived tokens
   - Supports both STACKIT-generated and custom key pairs

2. **Token Flow**
   - Uses long-lived service account tokens
   - Simpler but less secure

### Configuration Options

You can configure authentication using any of these methods:

1. **Explicit Configuration in Code**

   ```go
   config.WithServiceAccountKeyPath("path/to/sa_key.json")
   config.WithToken("your-token")
   ```

2. **Environment Variables**

   ```bash
   STACKIT_SERVICE_ACCOUNT_KEY_PATH=/path/to/sa_key.json
   STACKIT_SERVICE_ACCOUNT_TOKEN=your-token
   ```

3. **Credentials File**
   - Default location: `$HOME/.stackit/credentials.json`
   - Custom location: Set via `STACKIT_CREDENTIALS_PATH`
   - Format:
     ```json
     {
       "STACKIT_SERVICE_ACCOUNT_KEY_PATH": "path/to/sa_key.json",
       "STACKIT_SERVICE_ACCOUNT_TOKEN": "your-token"
     }
     ```

### Configuration Reference

1. **Explicit Configuration**

   - `config.WithServiceAccountKey(string)` - Set the service account key JSON directly
   - `config.WithServiceAccountKeyPath(string)` - Set the path to the service account key JSON file
   - `config.WithPrivateKey(string)` - Set the service account private key directly (for custom key pairs)
   - `config.WithPrivateKeyPath(string)` - Set the path to the service account private key file
   - `config.WithToken(string)` - Set the service account access token directly

2. **Environment Variables**

   - `STACKIT_SERVICE_ACCOUNT_KEY` - Service account key JSON as string
   - `STACKIT_SERVICE_ACCOUNT_KEY_PATH` - Path to service account key JSON file
   - `STACKIT_PRIVATE_KEY` - Service account private key as string
   - `STACKIT_PRIVATE_KEY_PATH` - Path to service account private key file
   - `STACKIT_SERVICE_ACCOUNT_TOKEN` - Service account access token
   - `STACKIT_CREDENTIALS_PATH` - Custom path to credentials file

3. **Credentials File**
   - JSON file containing any of the above environment variable names as keys
   - Default location: `$HOME/.stackit/credentials.json`

### Configuration Priority

The SDK searches for credentials in the following order:

1. Explicit configuration in code
2. Environment variables
3. Credentials file

For each authentication method, the key flow is attempted first, followed by the token flow.

Check the [authentication example](examples/authentication/authentication.go) for implementation details.

## Reporting issues

If you encounter any issues or have suggestions for improvements, please open an issue in the repository or create a ticket in the [STACKIT Help Center](https://support.stackit.cloud/).

## Contribute

Your contribution is welcome! For more details on how to contribute, refer to our [Contribution Guide](./CONTRIBUTION.md).

## License

Apache 2.0
