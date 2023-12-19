## v0.7.0 (2023-12-19)

API methods and structs were renamed to have the same look and feel across all services and according to user feedback.

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

## v0.6.0 (2023-11-10)

- Manage your STACKIT Redis resources: `Instance`, `Credentials`, `Offerings`
- Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/redis)
