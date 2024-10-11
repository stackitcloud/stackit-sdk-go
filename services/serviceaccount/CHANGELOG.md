- **Feature:** Add support for nullable models

## v0.4.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.3.7 (2024-04-09)

- Remove unused model data types.

## v0.3.6 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.3.5 (2024-02-02)

- **Improvement**: Reword API client's description
- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.3.4 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.3.3 (2024-01-15)

- Add license and notice files

## v0.3.2 (2024-01-09)

- Dependency updates

## v0.3.1 (2023-12-22)

- Dependency updates

## v0.3.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `GetAccessTokens` renamed to `ListAccessTokens`
  - `GetServiceAccountJWKS` renamed to `GetJWKS`
  - `GetServiceAccountKeys` renamed to `ListServiceAccountKeys`
  - `GetServiceAccounts` renamed to `ListServiceAccounts`
  - `GetUsers` renamed to `ListUsers`
  - `UpdateServiceAccountKey` renamed to `PartialUpdateServiceAccountKey`
- Changed structs:
  - `AccessTokensResponse` renamed to `ListAccessTokensResponse`
  - `GetServiceAccountJWKS` renamed to `JWKS`
  - `GetServiceAccountsKeysResponse` renamed to `ListServiceAccountKeysResponse`
  - `IntrospectRequestBody` renamed to `IntrospectJWTPayload`
  - `IntrospectResponseBody` renamed to `IntrospectJWTResponse`
  - `JsonWebKey` renamed to `JWK`
  - `ServiceAccountImpersonationRequestBody` renamed to `ImpersonateServiceAccountPayload`
  - `ServiceAccountImpersonationResponse` renamed to `ImpersonateServiceAccountResponse`
  - `UpdateServiceAccountKeyPayload` renamed to `PartialUpdateServiceAccountKeyPayload`
  - `UpdateServiceAccountKeyResponse` renamed to `PartialUpdateServiceAccountKeyResponse`

## v0.2.0 (2023-11-10)

- Manage your STACKIT service accounts
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/serviceaccount)
