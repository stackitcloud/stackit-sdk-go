## v0.7.3 (2023-12-13)

- `auth`: setup authentication, specifically using the service account key or token flows. Check our [authentication example](https://github.com/stackitcloud/stackit-sdk-go/blob/main/examples/authentication/authentication.go)
- `clients`: baseline http client implementations to support different use cases, such as the different authentication flows
- `config`: configuration for the SDK clients, such as custom endpoints, region and custom http client configuration. Check our [configuration example](https://github.com/stackitcloud/stackit-sdk-go/blob/main/examples/configuration/configuration.go)
- `oapierror`: open api error definition and handling
- `utils`: utilities, such as the `Ptr` method to return a pointer to a variable of any type, which can be useful for creating payloads
- `wait`: functionality to wait until a specific async operation has finished. Check our [waiter example](https://github.com/stackitcloud/stackit-sdk-go/blob/main/examples/waiter/waiter.go)
