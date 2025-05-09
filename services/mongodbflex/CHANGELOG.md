## v1.0.1 (2025-05-09)
- **Feature:** Update user-agent header

## v1.0.0 (2025-05-02)
- **Breaking Change:** Introduce typed enum constants for status attributes

## v0.18.2 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v0.18.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.18.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.17.0 (2025-01-20)

- **Breaking change**: Delete endpoint made private.

## v0.16.1 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.16.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.15.0 (2024-09-02)

- **Feature**: New method `ListAdvisorSlowQueries` that gets slow queries from the Opsmanager performance advisor.
- **Feature**: New method `ListSuggestedIndexes` that gets suggested indexes from the Opsmanager performance advisor.
- **Breaking change**: Remove nullable fields.

## v0.14.0 (2024-05-22)

- **Breaking change**: Remove unused data types.

## v0.13.0 (2024-05-14)

- Waiter for async operation `RestoreInstance`

## v0.12.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.11.1 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.11.0 (2024-02-02)

- **Breaking Change**: `HandlersInfraGetFlavorsResponse` renamed to `DevAzureComSchwarzitSchwarzitStackitMongodbStackitMongodbApiGitHandlersInfraGetFlavorsResponse`
- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.10.3 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.10.2 (2024-01-15)

- Add license and notice files

## v0.10.1 (2024-01-09)

- Dependency updates

## v0.10.0 (2024-01-05)

- **Breaking Change:** Added method `ListMetrics` (which accepts the desired metric: `cpu`, `database`, `disk-iops`, `disk-use`, or `exec-time`) as an argument and replaces `ListCPUMetrics`, `ListDatabaseStorageMetrics`, `ListDiskIOPSMetrics`, `ListDiskUsageMetrics`, and `ListExecutionTimesMetrics`, respectivelly
- **Breaking Change:** Removes the previously deprecated method `GetStorage`, replaced by `ListStorages`

## v0.9.0 (2023-12-22)

- Added struct `ApiListStoragesRequest`, which will replace `ApiGetStorageRequest`
  - `ApiGetStorageRequest` has been marked as deprecated, and will be removed in the next minor update
- Added method `ListStorages`, which will replace `GetStorage`
  - `GetStorage` has been marked as deprecated, and will be removed in the next minor update
- Added `CloneInstanceWaitHandler`, to wait for `CloneInstance` async operation to be completed
- Dependency updates

## v0.8.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `DeleteProject` renamed to `DisableService`
  - `GetBackups` renamed to `ListBackups`
  - `GetCPUMetrics` renamed to `ListCPUMetrics`
  - `GetDatabaseStorageMetrics` renamed to `ListDatabaseStorageMetrics`
  - `GetDiskIOPSMetrics` renamed to `ListDiskIOPSMetrics`
  - `GetDiskUsageMetrics` renamed to `ListDiskUsageMetrics`
  - `GetExecutionTimesMetrics` renamed to `ListExecutionTimesMetrics`
  - `GetFlavors` renamed to `ListFlavors`
  - `GetInstanceRestores` renamed to `ListRestoreJobs`
  - `GetInstances` renamed to `ListInstances`
  - `GetMemoryMetrics` renamed to `ListMemoryMetrics`
  - `GetUsers` renamed to `ListUsers`
- Changed structs:
  - `BackupScheduleModelJson` renamed to `BackupSchedule`
  - `GetBackupsResponse` renamed to `ListBackupsResponse`
  - `GetFlavorsResponse` renamed to `ListFlavorsResponse`
  - `GetInstanceRestoresResponse` renamed to `ListRestoreJobsResponse`
  - `GetInstancesResponse` renamed to `ListInstancesResponse`
  - `GetMetricsResponse` renamed to `ListMetricsResponse`
  - `GetStorageResponse` renamed to `ListStorageResponse`
  - `GetUsersResponse` renamed to `ListUsersResponseItem`
  - `GetVersionsResponse` renamed to `ListVersionsResponse`
  - `InstanceAcl` renamed to `ACL`
  - `InstanceBackup` renamed to `Backup`
  - `InstanceDataPoint` renamed to `DataPoint`
  - `InstanceError` renamed to `Error`
  - `InstanceFlavor` renamed to `Flavor`
  - `InstanceHost` renamed to `Host`
  - `InstanceHostMetric` renamed to `HostMetric`
  - `InstanceListUser` renamed to `ListUsersResponse`
  - `InstanceRestore` renamed to `RestoreInstanceStatus`
  - `InstanceSingleInstance` renamed to `Instance`
  - `InstanceStorage` renamed to `Storage`
  - `InstanceStorageRange` renamed to `StorageRange`
  - `InstanceUser` renamed to `User`
- Added waiters:
  - `PartialUpdateInstanceWaitHandler`, for `PartialUpdateInstance` method

## v0.7.0 (2023-11-10)

- Manage your STACKIT MongoDB Flex resources: `Instance`, `Flavors`, `Metrics`, `User`, `Storages`, `Versions`
- Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/mongodbflex)
