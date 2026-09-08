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

Requires `Go 1.25` or higher.

To download the `core` module:

```
go get github.com/stackitcloud/stackit-sdk-go/core
```

To download the `services/dns` module:

```
go get github.com/stackitcloud/stackit-sdk-go/services/dns
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

   dns "github.com/stackitcloud/stackit-sdk-go/services/dns/v1api"
)

func main() {
   projectId := "PROJECT_ID" // the uuid of your STACKIT project

   // Create a new API client, that uses default authentication and configuration
   dnsClient, err := dns.NewAPIClient()
   if err != nil {
      fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
      os.Exit(1)
   }

   // Get the DNS Zones for your project
   var getZoneResp *dns.ListZonesResponse
   getZoneResp, err = dnsClient.DefaultAPI.ListZones(context.Background(), projectId).Execute()

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
   createZoneResp, err = dnsClient.DefaultAPI.CreateZone(context.Background(), projectId).CreateZonePayload(createZonePayload).Execute()
   if err != nil {
      fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `ZoneApi.CreateZone`: %v\n", err)
   } else {
      var createdZone = createZoneResp.Zone
      fmt.Printf("[DNS API] Created zone \"%s\" with DNS name \"%s\" and zone id \"%s\".\n", createdZone.Name, createdZone.DnsName, createdZone.Id)
   }

   // Get a record set of a DNS zone.
   var recordSetResp *dns.RecordSetResponse
   recordSetResp, err = dnsClient.DefaultAPI.GetRecordSet(context.Background(), projectId, "zoneId", "recordSetId").Execute()
   if err != nil {
      fmt.Fprintf(os.Stderr, "[DNS API] Error when calling `GetRecordSet`: %v\n", err)
   } else {
      fmt.Printf("[DNS API] Got record set with name \"%s\".\n", recordSetResp.Rrset.Name)
   }
}

```

More examples on other services, configuration and authentication possibilities can be found in the [examples folder](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples).

## Authentication

To authenticate with the SDK, you need a [service account](https://docs.stackit.cloud/platform/access-and-identity/service-accounts/) with appropriate permissions (e.g., `project.owner`, see [here](https://docs.stackit.cloud/platform/access-and-identity/roles-permissions/assign-roles-to-account/)). You can create a service account through the STACKIT Portal.

### Authentication Methods

The SDK supports three authentication methods:

1. **Workload Identity Federation Flow**

   - Uses OIDC trusted tokens
   - Provides best security through short-lived tokens without secrets

2. **Key Flow**

   - Uses RSA key-pair based authentication
   - Provides better security through short-lived tokens
   - Supports both STACKIT-generated and custom key pairs

3. **Token Flow** (Deprecated)
   - Uses long-lived service account tokens
   - Simpler but less secure

### Configuration Priority

The SDK searches for credentials in the following order:

1. Explicit configuration in code
2. Environment variables
3. Credentials file (`$HOME/.stackit/credentials.json`)

For each authentication method, the try order is:
1. Workload Identity Federation Flow
2. Key Flow
3. Token Flow

### Modular Identity Package (Initial)

An initial modular identity package is available at `core/identity`. It exposes a minimal
token contract, one provider per authentication flow, and primitives to compose them:

- `identity.TokenProvider` — the contract: `Token(ctx, options) (Token, error)`
- `identity.StaticTokenProvider` — a pre-issued token
- `identity.ServiceAccountKeyProvider` — the service account key flow
- `identity.WorkloadIdentityFederationProvider` — workload identity federation
- `identity.InstanceMetadataProvider` — the service account attached to a STACKIT VM
- `identity.CLIProvider` — the session of a logged-in STACKIT CLI, for developer machines
- `identity.ChainedProvider` — tries providers in order until one succeeds
- `identity.DefaultProvider` — the opinionated chain, see below

#### Using the default chain

`identity.NewDefaultProvider` builds the credential chain the SDK ships out of the box.
Every field of `DefaultProviderConfig` is optional: whatever you leave empty is resolved
from environment variables and then from the credentials file
(`STACKIT_CREDENTIALS_PATH`, falling back to `~/.stackit/credentials.json`).

```go
package main

import (
   "context"
   "fmt"
   "os"

   "github.com/stackitcloud/stackit-sdk-go/core/config"
   "github.com/stackitcloud/stackit-sdk-go/core/identity"
   dns "github.com/stackitcloud/stackit-sdk-go/services/dns/v1api"
)

func main() {
   // Anything left empty here falls back to the environment and the credentials file.
   tokenProvider, err := identity.NewDefaultProvider()
   if err != nil {
      fmt.Fprintf(os.Stderr, "Creating token provider: %v\n", err)
      os.Exit(1)
   }

   // Hand the provider to any SDK client.
   dnsClient, err := dns.NewAPIClient(config.WithTokenProvider(tokenProvider))
   if err != nil {
      fmt.Fprintf(os.Stderr, "[DNS API] Creating API client: %v\n", err)
      os.Exit(1)
   }
   _ = dnsClient

   // Or get the raw access token, for data plane APIs and any client outside this SDK.
   token, err := tokenProvider.Token(context.Background(), identity.TokenRequestOptions{})
   if err != nil {
      fmt.Fprintf(os.Stderr, "Getting token: %v\n", err)
      os.Exit(1)
   }
   fmt.Println("Authorization: Bearer " + token.AccessToken)
}
```

The chain tries, in order:

1. `StaticTokenProvider` — a pre-issued token
2. `ServiceAccountKeyProvider` — the service account key flow
3. `WorkloadIdentityFederationProvider` — workload identity federation
4. `InstanceMetadataProvider` — the service account attached to the STACKIT VM
5. `CLIProvider` — the session of a STACKIT CLI that has run `stackit auth login`

Explicitly configured credentials always take precedence over the ambient identity of the
machine and over local developer tooling, which are only consulted once everything else
has failed. Steps that are unavailable are skipped: on a CI runner with no STACKIT CLI
installed, step 5 costs nothing.

The chain never starts an interactive login. Obtaining a session is an explicit operation
(`stackit auth login`), so that no program can unexpectedly open a browser.

The CLI step can be switched off, either from the environment or in code:

```bash
STACKIT_USE_CLI=false
```

```go
identity.NewDefaultProvider(&identity.DefaultProviderConfig{DisableCLI: true})
```

Either switch is enough, and neither re-enables what the other turned off: application
code cannot override an operator's decision to keep the CLI out of the loop. Constructing
an `identity.CLIProvider` yourself is unaffected.

#### Building your own chain

The providers are independent, so you can assemble your own order — or plug in a source of
your own by implementing `identity.TokenProvider`:

```go
keyProvider, err := identity.NewServiceAccountKeyProvider(&identity.ServiceAccountKeyProviderConfig{
   ServiceAccountKey: keyJSON,
   PrivateKey:        privateKeyPEM,
})
if err != nil {
   return err
}

// myVaultProvider is any type implementing identity.TokenProvider.
tokenProvider, err := identity.NewChainedProvider(myVaultProvider, keyProvider)
if err != nil {
   return err
}
```


### Using the Workload Identity Fedearion Flow

1. Create a service account trusted relation in the STACKIT Portal:

   - Navigate to `Service Accounts` → Select account → `Federated Identity Providers`
   - [Configure a Federated Identity Provider](https://docs.stackit.cloud/platform/access-and-identity/service-accounts/how-tos/manage-service-account-federations/#create-a-federated-identity-provider) and the required assertions to trust in.

2. Configure authentication using any of these methods:

   **A. Code Configuration**

```go
// Using wokload identity federation flow
config.WithWorkloadIdentityFederationAuth()
// With the custom path for the external OIDC token
config.WithWorkloadIdentityFederationPath("/path/to/your/federated/token")
// For the service account
config.WithServiceAccountEmail("my-sa@sa-stackit.cloud")
```
**B. Environment Variables**
```bash
# With the custom path for the external OIDC token
STACKIT_FEDERATED_TOKEN_FILE=/path/to/your/federated/token
# For the service account
STACKIT_SERVICE_ACCOUNT_EMAIL=my-sa@sa-stackit.cloud
```

### Using the Key Flow

1. Create a service account key in the STACKIT Portal:

   - Navigate to `Service Accounts` → Select account → `Service Account Keys` → Create key
   - You can either let STACKIT generate the key pair or provide your own RSA key pair (see [Creating an RSA key-pair](https://docs.stackit.cloud/platform/access-and-identity/service-accounts/how-tos/manage-service-account-keys/) for more details)
   - **Note**: it's also possible to create the service account key in other ways (see [Tutorials for Service Accounts](https://docs.stackit.cloud/platform/access-and-identity/service-accounts/how-tos/manage-service-accounts/) for more details)

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
   - **Note**: it's also possible to create the service account access tokens in other ways (see [Tutorials for Service Accounts](https://docs.stackit.cloud/platform/access-and-identity/service-accounts/how-tos/get-access-token/) for more details)

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

Your contribution is welcome! For more details on how to contribute, refer to our [Contribution Guide](./CONTRIBUTING.md).

## Release creation

See the [release documentation](./RELEASE.md) for further information.

## License

Apache 2.0
