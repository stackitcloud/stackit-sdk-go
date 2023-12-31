## v0.7.5 (2024-01-09)

- **Improvement:** When using the key flow, the SDK will extract the private key from the service account key and use it, if no private key is provided in the configuration, through environment variable or in the credentials file. This makes it simpler to use the key flow: if you create a service account key including the private key, you don't need to provide the private key separately anymore

## v0.7.4 (2023-12-22)

- Replace k8s.io/apimachinery with cenkalti/backoff

## v0.7.3 (2023-12-13)

- `auth`: setup authentication, specifically using the service account key or token flows. Check our [authentication example](https://github.com/stackitcloud/stackit-sdk-go/blob/main/examples/authentication/authentication.go)
- `clients`: baseline http client implementations to support different use cases, such as the different authentication flows
- `config`: configuration for the SDK clients, such as custom endpoints, region and custom http client configuration. Check our [configuration example](https://github.com/stackitcloud/stackit-sdk-go/blob/main/examples/configuration/configuration.go)
- `oapierror`: open api error definition and handling
- `utils`: utilities, such as the `Ptr` method to return a pointer to a variable of any type, which can be useful for creating payloads
- `wait`: functionality to wait until a specific async operation has finished. Check our [waiter example](https://github.com/stackitcloud/stackit-sdk-go/blob/main/examples/waiter/waiter.go)
