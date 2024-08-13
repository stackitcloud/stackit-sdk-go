## v0.18.1 (2024-08-13)

- **Feature:** Add support for nullable fields

## v0.18.0 (2024-08-01)

- **Feature:** `Plan` has a new field `SkuName`

## v0.17.0 (2024-07-10)

- **Bugfix:** Fix marking of deprecated struct fields. Potential breaking change for users with linters that treat deprecations as errors.

## v0.16.0 (2024-06-19)

- **Feature**: New methods `CreateBackup`, `DownloadBackup`, `ListRestores`,`UpdateBackupsConfig`, `TriggerRecreate`, `TriggerRestart`, `TriggerRestore` to manage the backup and restoration of an instance.

## v0.15.0 (2024-05-29)

- **Feature**: `GetMetricsResponse` has new fields: `Load1`, `Load15`, `Load5`
- **Breaking change**: Deleted unused data type

## v0.14.0 (2024-05-13)

- **Feature**: New method `GetMetrics` to get the latest metrics for cpu load, memory and disk usage for an instance
- **Feature**: New method `ListBackups` to list the backups for an instance
- **Breaking change**: `ListMetricsResponse` type (previously unused) renamed to `GetMetricsResponse`
- **Breaking change**: Deleted unused data types

## v0.13.0 (2024-04-11)

- **Breaking change**: Fields removed from `RawCredentials`: `RouteServiceUrl`, `SyslogDrainUrl`, `VolumeMounts`.
- **Breaking change**: Fields removed from `Credentials`: `HttpApiUri`, `Protocols`

## v0.12.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.11.0 (2024-04-02)

- **Breaking change:** Field `Plugins` was removed from `InstanceParameters`
- **Feature:** `Offering` has a new field `Lifecycle`
- **Feature:** `Instance` has new fields `OfferingVersion`, `PlanName`, and `Status`

## v0.10.1 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.10.0 (2024-02-02)

- **Feature**: `Instance` has a new field `OfferingName`
- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.9.2 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.9.1 (2024-01-15)

- Add license and notice files

## v0.9.0 (2024-01-09)

- **Feature:** `PartialUpdateInstance` can be used to update the instance's name
- **Feature:** `InstanceParameters` has a new setting `MaxDiskThreshold`
- **Feature:** `ListMetricsResponse` has new fields regarding ephemeral disk
- Dependency updates

## v0.8.1 (2023-12-22)

- Dependency updates

## v0.8.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `GetCredentialsIds` renamed to `ListCredentials`
  - `GetInstances` renamed to `ListInstances`
  - `GetOfferings` renamed to `ListOfferings`
  - `UpdateInstance` renamed to `PartialUpdateInstance`
- Changed structs:
  - `CredentialsIdsResponse` renamed to `ListCredentialsResponse`
  - `InstanceBackup` renamed to `Backup`
  - `InstanceBackupsList` renamed to `ListBackupsResponse`
  - `InstanceId` renamed to `CreateInstanceResponse`
  - `InstanceList` renamed to `ListInstancesResponse`
  - `InstanceMetrics` renamed to `ListMetricsResponse`
  - `LastOperation` renamed to `InstanceLastOperation`
  - `OfferingList` renamed to `ListOfferingsResponse`
  - `UpdateInstancePayload` renamed to `PartialUpdateInstancePayload`
    Changed waiters:
  - `UpdateInstanceWaitHandler` renamed to `PartialUpdateInstanceWaitHandler`

## v0.7.0 (2023-11-10)

- Manage your STACKIT MariaDB resources: `Instance`, `Credentials`, `Offerings`
- Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/mariadb)
