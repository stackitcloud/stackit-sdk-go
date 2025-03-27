<div align="center">
<br>
<img src=".github/images/stackit-logo.svg" alt="STACKIT logo" width="50%"/>
<br>
<br>
</div>

# STACKIT SDK for Go

[![GitHub License](https://img.shields.io/github/license/stackitcloud/stackit-sdk-go)](https://www.apache.org/licenses/LICENSE-2.0)

This repository contains the published SDKs and [SDK releases](https://github.com/stackitcloud/stackit-sdk-go/releases/).
The modules are structured into a [core module](https://github.com/stackitcloud/stackit-sdk-go/tree/main/core) with service clients, authentication and shared functionality as well as the different STACKIT [services](https://github.com/stackitcloud/stackit-sdk-go/tree/main/services).
The usage of the SDK is shown in some [examples](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples).

## Getting started

Requires `Go 1.21` or higher.

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
The setup of the authentication is described below in section [Authentication](#authentication) in more detail.

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

To authenticate with the SDK, you need a [service account](https://docs.stackit.cloud/stackit/en/service-accounts-134415819.html) with appropriate permissions (e.g., `project.owner`, see [here](https://docs.stackit.cloud/stackit/en/assign-permissions-to-a-service-account-134415855.html)). You can create a service account through the STACKIT Portal.

### Authentication Methods

The SDK supports two authentication methods:

1. **Key Flow** (Recommended)

   - Uses RSA key-pair based authentication
   - Provides better security through short-lived tokens
   - Supports both STACKIT-generated and custom key pairs

2. **Token Flow**
   - Uses long-lived service account tokens
   - Simpler but less secure

### Configuration Priority

The SDK searches for credentials in the following order:

1. Explicit configuration in code
2. Environment variables (KEY_PATH for KEY)
3. Credentials file (`$HOME/.stackit/credentials.json`)

For each authentication method, the key flow is attempted first, followed by the token flow.

### Using the Key Flow

1. Create a service account key in the STACKIT Portal:

   - Navigate to `Service Accounts` → Select account → `Service Account Keys` → Create key
   - You can either let STACKIT generate the key pair or provide your own RSA key pair (see [Creating an RSA key-pair](https://docs.stackit.cloud/stackit/en/usage-of-the-service-account-keys-in-stackit-175112464.html#UsageoftheserviceaccountkeysinSTACKIT-CreatinganRSAkey-pair) for more details)
   - **Note**: it's also possible to create the service account key in other ways (see [Tutorials for Service Accounts](https://docs.stackit.cloud/stackit/en/tutorials-for-service-accounts-134415861.html) for more details)

2. Save the service account key JSON:

```json
{
  "id": "uuid",
  "publicKey": "public key",
  "credentials": {
    "kid": "string",
    "iss": "my-sa@sa.stackit.cloud",
    "sub": "uuid",
    "aud": "string",
    "privateKey": "private key (if STACKIT-generated)"
  }
  // ... other fields ...
}
```

3. Configure authentication using any of these methods:

   **A. Code Configuration**

   ```go
   // Using service account key file
   config.WithServiceAccountKeyPath("path/to/sa_key.json")
   // Or using key content directly
   config.WithServiceAccountKey(keyJSON)

   // Optional: For custom key pairs
   config.WithPrivateKeyPath("path/to/private.pem")
   // Or using private key content directly
   config.WithPrivateKey(privateKeyJSON)
   ```

   **B. Environment Variables**

   ```bash
   # Using service account key
   STACKIT_SERVICE_ACCOUNT_KEY_PATH=/path/to/sa_key.json
   # or
   STACKIT_SERVICE_ACCOUNT_KEY=<sa-key-content>

   # Optional: For custom key pairs
   STACKIT_PRIVATE_KEY_PATH=/path/to/private.pem
   # or
   STACKIT_PRIVATE_KEY=<private-key-content>
   ```

   **C. Credentials File** (`$HOME/.stackit/credentials.json`)

   ```json
   {
     "STACKIT_SERVICE_ACCOUNT_KEY_PATH": "/path/to/sa_key.json",
     "STACKIT_PRIVATE_KEY_PATH": "/path/to/private.pem"
   }
   ```

### Using the Token Flow

1. Create an access token in the STACKIT Portal:

   - Navigate to `Service Accounts` → Select account → `Access Tokens` → Create token
   - **Note**: it's also possible to create the service account access tokens in other ways (see [Tutorials for Service Accounts](https://docs.stackit.cloud/stackit/en/tutorials-for-service-accounts-134415861.html) for more details)

2. Configure authentication using any of these methods:

   **A. Code Configuration**

   ```go
   config.WithToken("your-token")
   ```

   **B. Environment Variables**

   ```bash
   STACKIT_SERVICE_ACCOUNT_TOKEN=your-token
   ```

   **C. Credentials File** (`$HOME/.stackit/credentials.json`)

   ```json
   {
     "STACKIT_SERVICE_ACCOUNT_TOKEN": "your-token"
   }
   ```

For detailed implementation examples, see the [authentication example](examples/authentication/authentication.go).

## Reporting issues

If you encounter any issues or have suggestions for improvements, please open an issue in the repository or create a ticket in the [STACKIT Help Center](https://support.stackit.cloud/).

## Contribute

Your contribution is welcome! For more details on how to contribute, refer to our [Contribution Guide](./CONTRIBUTION.md).

## License

Apache 2.0
