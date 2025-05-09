## v1.1.3 (2025-05-09)
- **Feature:** Update user-agent header

## v1.1.2 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v1.1.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v1.1.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v1.0.0 (2025-02-05)
- **Breaking Change:** The region is no longer specified within the client configuration. Instead, the region must be passed as a parameter to any region-specific request.

## v0.11.1 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.11.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.10.0 (2024-06-14)

- **Breaking change**: Remove unused data types.

## v0.9.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.8.7 (2024-04-09)

- Remove unused model data types.

## v0.8.6 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.8.5 (2024-02-02)

- Remove deprecation warnings
- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.8.4 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.8.3 (2024-01-15)

- Add license and notice files

## v0.8.2 (2024-01-09)

- Dependency updates

## v0.8.1 (2023-12-22)

- Dependency updates

## v0.8.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `CreateProject` renamed to `EnableService`
  - `DeleteProject` renamed to `DisableService`
  - `GetAccessKeys` renamed to `ListAccessKeys`
  - `GetBuckets` renamed to `ListBuckets`
  - `GetCredentialsGroups` renamed to `ListCredentialsGroups`
  - `GetProject` renamed to `GetServiceStatus`
- Changed structs:
  - `GetAccessKeysResponse` renamed to `ListAccessKeysResponse`
  - `GetBucketsResponse` renamed to `ListBucketsResponse`
  - `GetCredentialsGroupsResponse` renamed to `ListCredentialsGroupsResponse`
  - `GetProjectResponse` renamed to `ProjectStatus`

## v0.7.0 (2023-11-10)

- Manage your STACKIT ObjectStorage resources: `Bucket`, `AccessKey`, `CredentialGroup`
- Waiters for async operations: `CreateBucketWaitHandler`, `DeleteBucketWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/objectstorage)
