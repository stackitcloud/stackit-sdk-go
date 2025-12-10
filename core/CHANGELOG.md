## v0.21.0
- **Chore:** Use `jwt-bearer` grant to get a fresh token instead of `refresh_token`
- **Feature:** Support Workload Identity Federation flow

## v0.20.1
- **Improvement:** Improve error message when passing a PEM encoded file to as service account key

## v0.20.0
- **New:** Added new `GetTraceId` function

## v0.19.0
- **New:** Added new `EnumSliceToStringSlice ` util func

## v0.18.0
- **New:** Added duration utils
- **Chore:** Use `jwt-bearer` grant to get a fresh token instead of `refresh_token`

## v0.17.3 
- **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v0.17.2 (2025-05-22)
- **Bugfix:** Access tokens generated via key flow authentication are refreshed 5 seconds before expiration to prevent timing issues with upstream systems which could lead to unexpected 401 error responses

## v0.17.1 (2025-04-09)
- **Improvement:** Improve error message for key flow authentication

## v0.17.0 (2025-03-25)
- **New:** Helper functions for generic openapi error codes

## v0.16.2 (2025-03-21)
- **New:** If a custom http.Client is provided, the http.Transport is respected. This allows customizing the http.Client with custom timeouts or instrumentation.

## v0.16.1 (2025-02-25)

- **Bugfix:** STACKIT_PRIVATE_KEY and STACKIT_SERVICE_ACCOUNT_KEY can be set via environment variable or via credentials file.

## v0.16.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.15.1 (2025-01-08)

- **Bugfix:** `ConfigureRegion` does not return an error if a region is set by an enviroment variable.

## v0.15.0 (2025-01-02)

- **Breaking Change:**: `ConfigureRegion` returns an error if a region is specified for a global URL.

STACKIT will move to a new way of specifying regions, where the region is provided as a function argument instead of being set in the client configuration. Once all services have migrated, the methods to specify the region in the client configuration will be removed.

## v0.14.0 (2024-10-10)

- **Feature:**: Added `IntermediateStateReached` to `AsyncActionHandler` that can be used to check for an intermediate state when executing the wait function of a wait handler.

## v0.13.0 (2024-09-03)

- Deprecated method `config.WithJWKSEndpoint` and field `config.Configuration.JWKSCustomUrl` have been removed. Deprecation was done in the `core` release [v0.10.0](core/CHANGELOG.md#v0100-2024-02-27).

## v0.12.0 (2024-04-11)

- **Feature:** Add `Middleware` type, `WithMiddleware` and `ChainMiddleware` methods to package `config`, this allows clients to chain and add Middlewares to the transport layer of the HTTP client.

## v0.11.0 (2024-04-09)

- **Feature:** Add method `WithCaptureHTTPRequest` to package `runtime`, which allows capture of HTTP requests for debugging purposes.

## v0.10.1 (2024-03-20)

- **Improvement:** Update `ConfigureRegion` method to take into account global servers without a region variable

## v0.10.0 (2024-02-27)

- **Feature:** Add configuration option that, for the key flow, enables a goroutine to be spawned that will refresh the access token when it's close to expiring
- **Deprecation:**
  - Methods:
    - `config.WithMaxRetries`
    - `config.WithWaitBetweenCalls`
    - `config.WithRetryTimeout`
    - `clients.NewRetryConfig`
  - Fields:
    - `clients.KeyFlowConfig.ClientRetry`
    - `clients.TokenFlowConfig.ClientRetry`
    - `clients.NoAuthFlowConfig.ClientRetry`
    - `clients.RetryConfig`
  - Retry options were removed to reduce complexity of the clients. If this functionality is needed, you can provide your own custom HTTP client.
- **Breaking Change:** Change signature of `auth.NoAuth`, which no longer takes `clients.RetryConfig` as argument.
- **Breaking Change:**
  - Methods:
    - `clients.KeyFlow.Clone`
    - `clients.TokenFlow.Clone`
    - `clients.NoAuthFlow.Clone`
    - `clients.Do`
  - Fields:
    - `clients.DefaultRetryMaxRetries`
    - `clients.DefaultRetryWaitBetweenCalls`
    - `clients.DefaultRetryTimeout`
  - Constants:
    - `clients.ClientTimeoutErr`
    - `clients.ClientContextDeadlineErr`
    - `clients.ClientConnectionRefusedErr`
    - `clients.ClientEOFError`
    - `clients.Environment`
  - Removed to reduce complexity of the clients, they were no longer being used.

## v0.9.0 (2024-02-19)

- **Deprecation:** Mark method `config.WithCaptureHTTPResponse` as deprecated, to avoid confusion due to it not being a configuration option. Use `runtime.WithCaptureHTTPResponse` instead.
- **Deprecation:** Mark method `config.WithJWKSEndpoint` and field `config.Configuration.JWKSCustomUrl` as deprecated. Validation using JWKS was removed, for being redundant with token validation done in the APIs. These have no effect.
- **Breaking Change:** Remove method `KeyFlow.Clone`, that was no longer being used.

## v0.8.0 (2024-02-16)

- **Feature:** Add package `runtime`, which implements methods to be used when performing API requests.
- **Feature:** Add method `WithCaptureHTTPResponse` to package `runtime`, which does the same as `config.WithCaptureHTTPResponse`. Method was moved to avoid confusion due to it not being a configuration option, and will be removed in a later release.

## v0.7.7 (2024-02-02)

- **Bugfix:** Use `http.Request` context in `clients.Do`

## v0.7.6 (2024-01-15)

- Add LICENSE and NOTICE files

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
