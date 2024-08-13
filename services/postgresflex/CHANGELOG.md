## v0.15.1 (2024-08-13)

- **Feature:** Add support for nullable fields

## v0.15.0 (2024-06-28)

- **Feature**: New API methods `CreateDatabase`, `DeleteDatabase`, `ListDatabase`, `ListDatabaseParameters` to manage PostgreSQL Flex databases
- **Feature**: New API method `UpdateInstance` to update the instance
- **Feature**: New API method `ListMetrics` to list metrics of an instance
- **Feature**: New API method `DisableService` to terminate the whole project

## v0.14.0 (2024-05-22)

- **Breaking change**: Remove unused model data types.

## v0.13.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.12.0 (2024-04-03)

- **Improvement**: Update `DeleteInstanceWaitHandler` to support new deletion method.
- **Feature**: New waiter `ForceDeleteInstanceWaitHandler` for async operation `ForceDeleteInstance`

## v0.11.0 (2024-04-02)

- **Feature**: New API method `ForceDeleteInstance` that forces the deletion of a delayed deleted instance

## v0.10.0 (2024-03-08)

- **Feature**: New API method `CloneInstance` to clone the instance

## v0.9.2 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.9.1 (2024-02-06)

- **Bug fix**: Fix `CreateInstanceWaitHandler` failing when instance isn't shown right after creation request

## v0.9.0 (2024-02-05)

- **Feature**: New API method `UpdateUser` to update user
- **Feature**: New API method `PartialUpdateUser` to patch update user
- **Feature**: New API method `ResetUser` to reset a user's password

## v0.8.5 (2024-02-02)

- **Improvement**: Required fields in `CreateInstancePayload` and `UpdateBackupSchedulePayload` are now labelled as such
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
  - `GetBackups` renamed to `ListBackups`
  - `GetFlavors` renamed to `ListFlavors`
  - `GetInstances` renamed to `ListInstances`
  - `GetStorages` renamed to `ListStorages`
  - `GetUsers` renamed to `ListUsers`
  - `GetVersions` renamed to `ListVersions`
  - `UpdateInstance` renamed to `PartialUpdateInstance`
- Changed structs:
  - `BackupResponse` renamed to `GetBackupResponse`
  - `BackupsResponse` renamed to `ListBackupsResponse`
  - `FlavorsResponse` renamed to `ListFlavorsResponse`
  - `GetUsersResponse` renamed to `ListUsersResponse`
  - `InstanceAcl` renamed to `ACL`
  - `InstanceBackup` renamed to `Backup`
  - `InstanceError` renamed to `Error`
  - `InstanceFlavor` renamed to `Flavor`
  - `InstanceListUser` renamed to `ListUsersResponseItem`
  - `InstanceResetUserResponse` renamed to `ResetUserResponse`
  - `InstanceSingleInstance` renamed to `Instance`
  - `InstanceStorage` renamed to `Storage`
  - `InstanceStorageRange` renamed to `StorageRange`
  - `InstanceUser` renamed to `User`
  - `InstancesResponse` renamed to `ListInstancesResponse`
  - `StoragesResponse` renamed to `ListStoragesResponse`
  - `UpdateInstancePayload` renamed to `PartialUpdateInstancePayload`
  - `UpdateInstanceResponse` renamed to `PartialUpdateInstanceResponse`
  - `UserResponse` renamed to `GetUserResponse`
  - `UsersResponse` renamed to `ListUsersResponse`
  - `VersionsResponse` renamed to `ListVersionsResponse`
    Changed waiters:
  - `UpdateInstanceWaitHandler` renamed to `PartialUpdateInstanceWaitHandler`

## v0.7.0 (2023-11-10)

- Manage your STACKIT PostgreSQL Flex resources: `Instance`, `Versions`, `Flavors`, `User`, `Storages`
- Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `DeleteUserWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/postgresflex)
