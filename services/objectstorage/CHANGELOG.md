## v0.8.7 (2024-04-09)

- Remove unused models.

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
