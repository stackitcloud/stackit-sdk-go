## v0.8.0 (2023-12-19)

API methods and structs were renamed to have the same look and feel across all services and according to user feedback.

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
