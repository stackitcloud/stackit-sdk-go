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
