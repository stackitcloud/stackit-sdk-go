## v0.21.2 (2025-05-09)
- **Feature:** Update user-agent header

## v0.21.1 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v0.21.0 (2025-04-03)
  - **Feature:** Adapt constants to evolved API.
  
## v0.20.2 (2025-04-02)
- **Bugfix:** `PartialUpdateInstanceWaitHandler` does not finish when update is succeeded
- **Deprecation:** Deprecated `InstanceStateSuccess`, `InstanceStateFailed`, `InstanceTypeCreate`, `InstanceTypeUpdate`, `InstanceTypeDelete` and will be removed after 2nd October 2025

## v0.20.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.20.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.19.1 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.19.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.18.0 (2024-09-02)

- **Breaking changes:** `GetMetricsResponse` fields have changed data types
  - `CpuLoadPercent`, `Load1`, `Load15` and `Load5` are now `*float64`
  - `ParachuteDiskEphemeralActivated` and `ParachuteDiskPersistentActivated` are now `*bool`

## v0.17.0 (2024-08-01)

- **Feature:** `Plan` has a new field `SkuName`

## v0.16.0 (2024-07-10)

- **Bugfix:** Fix marking of deprecated struct fields. Potential breaking change for users with linters that treat deprecations as errors.

## v0.15.0 (2024-06-19)

- **Feature**: New methods `CreateBackup`, `DownloadBackup`, `ListRestores`,`UpdateBackupsConfig`, `TriggerRecreate`, `TriggerRestart`, `TriggerRestore` to manage the backup and restoration of an instance.
- New fields `Load1`, `Load5` and `Load15` in `GetMetricsResponse` model.

## v0.14.0 (2024-05-13)

- **Feature**: New method `GetMetrics` to get the latest metrics for cpu load, memory and disk usage for an instance
- **Feature**: New method `ListBackups` to list the backups for an instance
- **Breaking change**: `ListMetricsResponse` type (previously unused) renamed to `GetMetricsResponse`
- **Breaking change**: Deleted unused data types

## v0.13.0 (2024-04-11)

- **Breaking change**: Fields removed from `RawCredentials`: `RouteServiceUrl`, `SyslogDrainUrl`, `VolumeMounts`.
- **Breaking change**: Fields removed from `Credentials`: `HttpApiUri`, `Name`, `Protocols`
- **Feature**: `Credentials` has a new field `Scheme`

## v0.12.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.11.0 (2024-04-02)

- **Feature:** `InstanceParameters` has new fields `TlsCiphers`, `TlsProtocols`, `JavaGarbageCollector`, `JavaHeapspace`, and `JavaMaxmetaspace`
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

- Manage your STACKIT OpenSearch resources: `Instance`, `Credentials`, `Offerings`
- Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/opensearch)
