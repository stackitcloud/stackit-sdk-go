## v0.3.0 (2023-12-19)

API methods and structs were renamed to have the same look and feel across all services and according to user feedback.

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
