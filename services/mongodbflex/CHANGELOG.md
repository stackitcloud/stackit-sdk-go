## v0.10.0 (2024-01-05)

- Added method `ListMetrics` (which accepts the desired metric: `cpu`, `database`, `disk-iops`, `disk-use`, or `exec-time`) as an argument and replaces `ListCPUMetrics`, `ListDatabaseStorageMetrics`, `ListDiskIOPSMetrics`, `ListDiskUsageMetrics`, and `ListExecutionTimesMetrics`, respectivelly
- Removes the previously deprecated method `GetStorage`, replaced by `ListStorages`

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
