## v0.13.0 (2024-04-11)

- **Breaking change**: Fields removed from `RawCredentials`: `RouteServiceUrl`, `SyslogDrainUrl`, `VolumeMounts`.
- **Breaking change**: Fields removed from `Credentials`: `Hosts`, `HttpApiUri`, `Name`, `Protocols`
- **Feature**: `Credentials` has a new field `SyslogDrainUrl`

## v0.12.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.11.0 (2024-04-02)

- **Feature:** `InstanceParameters` has new fields `FluentdTcp`, `FluentdTls`, `FluentdTlsCiphers`, `FluentdTlsMaxVersion`, `FluentdTlsMinVersion`, `FluentdTlsVersion`, `FluentdUdp`, `Groks`, `IsmDeletionAfter`, `IsmJitter`, `IsmJobInterval`, `JavaHeapspace`, `JavaMaxmetaspace`, `OpensearchTlsCiphers`, `OpensearchTlsProtocols`, and `SyslogUseUdp`
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

- Manage your STACKIT Logme resources: `Instance`, `Credentials`, `Offerings`
- Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/logme)
