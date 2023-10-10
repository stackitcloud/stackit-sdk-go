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

To authenticate to the SDK, you will need a [service account](https://docs.stackit.cloud/stackit/en/service-accounts-134415819.html). Create it in the STACKIT Portal an assign it the necessary permissions, e.g. `project.owner`. There are multiple ways to authenticate:

- Key flow (recommended)
- Token flow

When setting up authentication, the SDK will always try to use the key flow first and search for credentials in several locations, following a specific order:

1. Explicit configuration, e.g. by using the option `config.WithServiceAccountKeyPath("path/to/sa_key.json")`
2. Environment variable, e.g. by setting `STACKIT_SERVICE_ACCOUNT_KEY_PATH`
3. Credentials file

   The SDK will check the credentials file located in the path defined by the `STACKIT_CREDENTIALS_PATH` env var, if specified,
   or in `$HOME/.stackit/credentials.json` as a fallback.
   The credentials should be set using the same name as the environment variables. Example:

   ```json
   {
     "STACKIT_SERVICE_ACCOUNT_TOKEN": "foo_token",
     "STACKIT_SERVICE_ACCOUNT_KEY_PATH": "path/to/sa_key.json",
     "STACKIT_PRIVATE_KEY_PATH": "path/to/private_key.pem"
   }
   ```

Check the [authentication example](examples/authentication/authentication.go) for more details.

### Key flow

To use the key flow, you need to have a service account key and an RSA key-pair.
To configure it, follow this steps:

    The following instructions assume that you have created a service account and assigned it the necessary permissions, e.g. project.owner.

1.  In the Portal, go to the `Service Accounts` tab, choose a `Service Account` and go to `Service Account Keys` to create a key.

    - You can create your own RSA key-pair or have the Portal generate one for you.

    **Disclaimer:** as of now, creation of a service account key in the Portal is only available in DEV and QA environments. You can use this flow in these environments by using the options `config.WithTokenEndpoint` and `config.WithJWKSEndpoint` to configure the corresponding endpoints.

2.  Save the content of the service account key and the corresponding private key by copying them or saving them in a file.

    **Hint:** If you have generated the RSA key-pair using the Portal, you can save the private key in a PEM encoded file by downloading the service account key as a PEM file and using `openssl storeutl -keys <path/to/sa_key_pem_file> > private.key` to extract the private key from the service account key.

The expected format of the service account key is a **json** with the following structure:

```json
{
  "id": "uuid",
  "publicKey": "public key",
  "createdAt": "2023-08-24T14:15:22Z",
  "validUntil": "2023-08-24T14:15:22Z",
  "keyType": "USER_MANAGED",
  "keyOrigin": "USER_PROVIDED",
  "keyAlgorithm": "RSA_2048",
  "active": true,
  "credentials": {
    "kid": "string",
    "iss": "my-sa@sa.stackit.cloud",
    "sub": "uuid",
    "aud": "string",
    (optional) "privateKey": "private key when generated by the SA service"
  }
}
```

3. Configure the service account key and private key for authentication in the SDK:
   - using the configuration options: `config.WithServiceAccountKey` or `config.WithServiceAccountKeyPath`, `config.WithPrivateKey` or `config.WithPrivateKeyPath`
   - setting environment variables: `STACKIT_SERVICE_ACCOUNT_KEY_PATH` and `STACKIT_PRIVATE_KEY_PATH`
   - setting `STACKIT_SERVICE_ACCOUNT_KEY_PATH` and `STACKIT_PRIVATE_KEY_PATH` in the credentials file (see above)
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
