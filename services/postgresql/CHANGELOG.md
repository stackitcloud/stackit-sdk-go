> **The STACKIT PostgreSQL service will reach its end of support on June 30th 2024. All calls done to the API after that will stop working.**
>
> Use `github.com/stackitcloud/stackit-sdk-go/services/postgresflex` instead. For more details, check https://docs.stackit.cloud/stackit/en/bring-your-data-to-stackit-postgresql-flex-138347648.html.

## v0.12.0 (2024-02-06)

- Add deprecation note

## v0.11.0 (2024-02-02)

- **Feature**: `Instance` has a new field `OfferingName`
- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.10.2 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.10.1 (2024-01-15)

- Add license and notice files

## v0.10.0 (2024-01-09)

- **Feature:** `PartialUpdateInstance` can be used to update the instance's name
- **Feature:** `InstanceParameters` has a new setting `MaxDiskThreshold`
- **Feature:** `ListMetricsResponse` has new fields regarding ephemeral disk
- Dependency updates

## v0.9.1 (2023-12-22)

- Dependency updates

## v0.9.0 (2023-12-20)

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

## v0.8.0 (2023-11-17)

- Manage your STACKIT PostgreSQL resources: `Instance`, `Credentials`, `Offerings`
- Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/postgresql)
