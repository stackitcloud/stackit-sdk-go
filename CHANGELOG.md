## Release (2025-XX-XX)

- `authorization`: [v0.5.1](services/authorization/CHANGELOG.md#v051-2025-01-17)
  - **Bugfix:** Revert back to global URL configuration
- `core`: [v0.15.1](core/CHANGELOG.md#v0151-2025-01-08)
  - **Bugfix:** `ConfigureRegion` does not return an error if a region is set by an environment variable.
- `authorization`: [v0.5.0](services/authorization/CHANGELOG.md#v050-2025-01-09)
  - **Feature:** Add support for regions in `APIClient`
- `stackitmarketplace`: [v0.1.0](services/stackitmarketplace/CHANGELOG.md#v010-2025-01-10)
  - **New**: STACKIT Marketplace module can be used to manage the STACKIT Marketplace.
- `lbapplication` [v0.1.0](services/lbapplication/CHANGELOG.md#v010-2025-01-13)
  - **New**: STACKIT Load Balancer Application module can be used to manage the STACKIT Load Balancer Application
- `certificates`: [v0.1.0](services/certificates/CHANGELOG.md#v010-2025-01-13)
  - **New**: STACKIT Certificates module can be used to manage the STACKIT Load Balancer Certificates
- `serverbackup`: [v0.5.0](services/serverbackup/CHANGELOG.md#v050-2025-01-14)
  - **Feature:** Add new method: `GetServiceResource` 
- `serverupdate`: [v0.3.0](services/serverupdate/CHANGELOG.md#v030-2025-01-14)
  - **Feature:** Add new method: `GetServiceResource`
  
## Release (2025-01-07)

- `postgresflex`: [0.16.1](services/postgresflex/CHANGELOG.md#v0161-2025-xx-yy)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `iaas`: [v0.19.0](services/iaas/CHANGELOG.md#v0190-2024-12-20)
  - **Feature:** Add method to list quotas: `ListQuotas`
  - **Feature:** Add methods to change image scope: `UpdateImageScopeLocal` and `UpdateImageScopePublic`
- `core`: [v0.15.0](core/CHANGELOG.md#v0150-2025-01-02)
  - **Breaking Change:**: `ConfigureRegion` returns an error if a region is specified for a global URL.

## Release (2024-12-17)

> [!WARNING]
>
> **The membership service has reached its end of support on August 07th 2024.**
>
> Use `github.com/stackitcloud/stackit-sdk-go/services/authorization` instead.

- `iaas`: [v0.18.0](services/iaas/CHANGELOG.md#v0180-2024-12-16)
  - **Feature:** Add waiters for async operations: `UploadImageWaitHandler` and `DeleteImageWaitHandler`
- `iaas`: [v0.17.0](services/iaas/CHANGELOG.md#v0170-2024-12-16)
  - **Feature:** Add new methods to manage affinity groups: `CreateAffinityGroup`, `DeleteAffinityGroup`, `GetAffinityGroup`, and `ListAffinityGroup`
  - **Feature:** Add new methods to manage backups: `CreateBackup`, `DeleteBackup`, `GetBackup`, `ListBackup`, `RestoreBackup`, `ExecuteBackup`,`UpdateBackup`
  - **Feature:** Add new methods to manage images: `CreateImage`, `DeleteImage`, `GetImage`, `ListImage`,`UpdateImage`
  - **Feature:** Add new methods to manage imageshares: `DeleteImageShare`, `GetImageShare`, `SetImageShare`,`UpdateImageShare`
  - **Feature:** Add new methods to manage imageshare consumers: `DeleteImageShareConsumer`, `GetImageShareConsumer`, `SetImageShare`,`UpdateImageShare`
  - **Feature:** Add new methods to manage project NICs: `GetProjectNIC`, `ListProjectNICs`
  - **Feature:** Add new methods to manage snapshots: `CreateSnapshot`, `DeleteSnapshot`, `GetSnapshot`, `ListSnapshot`, `UpdateSnapshot`
  - **Bugfix:** Correctly handle nullable attributes in model types
- `serverupdate`: [v0.2.2](services/serverupdate/CHANGELOG.md#v022-2024-12-02)
  - **Bugfix:** `Id` field of `Update` model is now of type `int64` (was `string`)
- `serviceenablement`: [v0.4.0](services/serviceenablement/CHANGELOG.md#v040-2024-12-05)
  - **Feature:** Add support for region eu02
- `sqlserverflex`: [v0.8.1](services/sqlserverflex/CHANGELOG.md#v081-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `ske`: [v0.20.1](services/ske/CHANGELOG.md#v0201-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `serverupdate`: [v0.2.3](services/serverupdate/CHANGELOG.md#v023-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `serverbackup`: [v0.4.0](services/serverbackup/CHANGELOG.md#v040-2024-12-17)
  - **Feature:** Add support for managing `BackupPolicy` resources
  - **Bugfix:** Correctly handle nullable attributes in model types
- `secretsmanager`: [v0.10.1](services/secretsmanager/CHANGELOG.md#v0101-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `runcommand`: [v0.2.1](services/runcommand/CHANGELOG.md#v021-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `resourcemanager`: [v0.11.1](services/resourcemanager/CHANGELOG.md#v0111-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `redis`: [v0.20.1](services/redis/CHANGELOG.md#v0201-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `rabbitmq`: [v0.20.1](services/rabbitmq/CHANGELOG.md#v0201-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `opensearch`: [v0.19.1](services/opensearch/CHANGELOG.md#v0191-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `observability`: [v0.2.1](services/observability/CHANGELOG.md#v021-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `objectstorage`: [v0.11.1](services/objectstorage/CHANGELOG.md#v0111-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `mongodbflex`: [v0.16.1](services/mongodbflex/CHANGELOG.md#v0161-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `mariadb`: [v0.20.1](services/mariadb/CHANGELOG.md#v0201-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `logme`: [v0.20.1](services/logme/CHANGELOG.md#v0201-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `dns`: [v0.12.1](services/dns/CHANGELOG.md#v0121-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types
- `authorization`: [v0.4.1](services/authorization/CHANGELOG.md#v041-2024-12-17)
  - **Bugfix:** Correctly handle nullable attributes in model types

## Release (2024-11-29)

- `serverbackup`: [v0.4.0](services/serverbackup/CHANGELOG.md#v040-2024-11-26)
  - **Feature:** Add support for managing `BackupPolicy` resources
- `serverupdate`: [v0.2.1](services/serverupdate/CHANGELOG.md#v021-2024-11-28)
  - **Bugfix:** Fix `Accept` header types
- `serverupdate`: [v0.2.0](services/serverupdate/CHANGELOG.md#v020-2024-11-26)
  - **Feature:** Add support for managing `UpdatePolicy` resources
- `dns`: [v0.12.0](services/dns/CHANGELOG.md#v0120-2024-11-20)
  - **Feature:** New models `ZoneModelsImportRecordModel` and `ZoneModelsImportZoneJson`
- `serverbackup`: [v0.3.0](services/serverbackup/CHANGELOG.md#v030-2024-11-20)
  - **Breaking changes**:
    - `ListBackups200Response` type renamed to `GetBackupsListResponse`
    - `GetBackupsListResponse` has been removed
    - `ErrorResponse` has been removed
    - Added new method `GetBackupSchedulesResponse`
    - Added new type `EnableServiceResourcePayload`
- `serverupdate`: [v0.1.0](services/serverupdate/CHANGELOG.md#v010-2024-11-20)
  - Manage your STACKIT Server Updates: `Update`, `UpdateSchedule`, `BackupProperties`
- `iaas`: [v0.16.0](services/iaas/CHANGELOG.md#v0160-2024-11-08)
  - **Feature:** Add new methods to manage key pairs: `CreateKeyPair`, `UpdateKeyPair`, `DeleteKeyPair`, `GetKeyPair`, and `ListKeyPairs`
  - **Feature:** Add new field `Bootable` to `Volume`, `CreateVolumePayload`, and `UpdateVolumePayload` data models
  - **Breaking change:** Rename `NIC` to `Nic` in all network interface methods (e.g. `CreateNIC` to `CreateNic`, `AddNICToServer` to `AddNicToServer`, etc)

## Release (2024-10-21)

- `iaas`: [v0.14.0](services/iaas/CHANGELOG.md#v0140-2024-10-18)
  - **Feature:** Add waiter methods for `Volume`, `Server` and `AttachedVolume`
- `resourcemanager`: [v0.11.0](services/iaas/CHANGELOG.md#v0110-2024-10-21)
- **Feature:** Get containers of a folder using the new method `BffGetContainersOfAFolder`
- **Feature:** Get containers of an organization using the new method `BffGetContainersOfAnOrganization`
- `iaas`: [v0.15.0](services/iaas/CHANGELOG.md#v0150-2024-10-21)
  - **Feature:** Filter network area routes by labels using the new `LabelSelector` method on `ApiListNetworkAreaRoutesRequest`
  - **Feature:** Update network area route using the new method `UpdateNetworkAreaRoute`

## Release (2024-10-18)

- `iaas`: [v0.13.0](services/iaas/CHANGELOG.md#v0130-2024-10-18)
  - **Feature:** Add support for managing following resources
    - `Volume`
    - `Server`
    - `NetworkInterface`
    - `PublicIP`
    - `SecurityGroup`
    - `SecurityGroupRule`
  - **Breaking change**: Remove `V1NetworkGateway` data model
  - **Bugfix**: Network response JSON decoding

## Release (2024-10-14)

- `sqlserverflex`: [v0.7.0](services/sqlserverflex/CHANGELOG.md#v070-2024-09-25)
  - **Breaking change**: Field `Item` in `ResetUserResponse` is now of type `SingleUser` (previously was `User`)
  - **Feature:** `DefaultDatabase` is no longer required in `CreateUserPayload`
- `iaas`: [v0.9.0](services/iaas/CHANGELOG.md#v090-2024-09-27)
  - **Feature:** Add `Labels` field to several data models
- `iaas`: [v0.10.0](services/iaas/CHANGELOG.md#v0100-2024-10-01)
  - **Feature:** Add `CreatedAt` and `UpdatedAt` fields to several data models
- `core`: [v0.14.0](core/CHANGELOG.md#v0140-2024-10-10)
  - **Feature:**: Added `IntermediateStateReached` to `AsyncActionHandler` that can be used to check for an intermediate state when executing the wait function of a wait handler
- `iaas`: [v0.11.0](services/iaas/CHANGELOG.md#v0110-2024-10-11)
  - **Feature:** Filter networks by labels using the new `LabelSelector` method on `ApiListNetworksRequest`
- `loadbalancer`: [v0.16.0](services/loadbalancer/CHANGELOG.md#v0160-2024-10-11)
  - **Feature:** Add pagination to `ListLoadBalancers` with the new fields `pageSize` and `pageId` on `ApiListLoadBalancersRequest` and the field `NextPageId` in `ListLoadBalancersResponse`
- `authorization`: [v0.4.0](services/authorization/CHANGELOG.md#v040-2024-10-14)
  - **Feature:** Add support for nullable models
  - **Feature:** New method `ListUserPermissions`
- `dns`: [v0.11.0](services/dns/CHANGELOG.md#v0110-2024-10-14)
  - **Feature:** Add support for nullable models
- `iaas`: [v0.12.0](services/iaas/CHANGELOG.md#v0120-2024-10-14)
  - **Feature:** Add support for nullable models
- `loadbalancer`: [v0.17.0](services/loadbalancer/CHANGELOG.md#v0170-2024-10-14)
  - **Feature:** Add support for nullable models
- `logme`: [v0.20.0](services/logme/CHANGELOG.md#v0200-2024-10-14)
  - **Feature:** Add support for nullable models
- `mariadb`: [v0.20.0](services/mariadb/CHANGELOG.md#v0200-2024-10-14)
  - **Feature:** Add support for nullable models
- `mongodbflex`: [v0.16.0](services/mongodbflex/CHANGELOG.md#v0160-2024-10-14)
  - **Feature:** Add support for nullable models
- `objectstorage`: [v0.11.0](services/objectstorage/CHANGELOG.md#v0110-2024-10-14)
  - **Feature:** Add support for nullable models
- `observability`: [v0.2.0](services/observability/CHANGELOG.md#v020-2024-10-14)
  - **Feature:** Add support for nullable models
- `opensearch`: [v0.19.0](services/opensearch/CHANGELOG.md#v0190-2024-10-14)
  - **Feature:** Add support for nullable models
- `postgresflex`: [v0.16.0](services/postgresflex/CHANGELOG.md#v0160-2024-10-14)
  - **Feature:** Add support for nullable models
- `rabbitmq`: [v0.20.0](services/rabbitmq/CHANGELOG.md#v0200-2024-10-14)
  - **Feature:** Add support for nullable models
- `redis`: [v0.20.0](services/redis/CHANGELOG.md#v0200-2024-10-14)
  - **Feature:** Add support for nullable models
- `resourcemanager`: [v0.10.0](services/resourcemanager/CHANGELOG.md#v0100-2024-10-14)
  - **Feature:** Add support for nullable models
- `runcommand`: [v0.2.0](services/runcommand/CHANGELOG.md#v020-2024-10-14)
  - **Feature:** Add support for nullable models
- `secretsmanager`: [v0.10.0](services/secretsmanager/CHANGELOG.md#v0100-2024-10-14)
  - **Feature:** Add support for nullable models
- `serverbackup`: [v0.2.0](services/serverbackup/CHANGELOG.md#v020-2024-10-14)
  - **Feature:** Add support for nullable models
- `serviceaccount`: [v0.5.0](services/serviceaccount/CHANGELOG.md#v050-2024-10-14)
  - **Feature:** Add support for nullable models
- `serviceenablement`: [v0.3.0](services/serviceenablement/CHANGELOG.md#v030-2024-10-14)
  - **Feature:** Add support for nullable models
- `ske`: [v0.20.0](services/ske/CHANGELOG.md#v0200-2024-10-14)
  - **Feature:** Add support for nullable models
- `sqlserverflex`: [v0.8.0](services/sqlserverflex/CHANGELOG.md#v080-2024-10-14)
  - **Feature:** Add support for nullable models

## Release (2024-09-19)

- `sqlserverflex`: [v0.6.0](services/sqlserverflex/CHANGELOG.md#v060-2024-09-19)
  - **Breaking change**: Field `ListBackupsResponse` has a new field `BackupListBackupsResponseGrouped`, replacing the removed `Count` and `Items` fields
- `rabbitmq`: [v0.19.0](services/rabbitmq/CHANGELOG.md#v0190-2024-09-02)
  - **Breaking changes:** `GetMetricsResponse` fields have changed data types
    - `CpuLoadPercent`, `Load1`, `Load15` and `Load5` are now `*float64`
    - `ParachuteDiskEphemeralActivated` and `ParachuteDiskPersistentActivated` are now `*bool`
- `redis`: [v0.19.0](services/redis/CHANGELOG.md#v0190-2024-09-02)
  - **Breaking changes:** `GetMetricsResponse` fields have changed data types
    - `CpuLoadPercent`, `Load1`, `Load15` and `Load5` are now `*float64`
    - `ParachuteDiskEphemeralActivated` and `ParachuteDiskPersistentActivated` are now `*bool`
- `core`: [v0.13.0](core/CHANGELOG.md#v0130-2024-09-03)
  - Deprecated method `config.WithJWKSEndpoint` and field `config.Configuration.JWKSCustomUrl` have been removed. Deprecation was done in the `core` release [v0.10.0](core/CHANGELOG.md#v0100-2024-02-27).
- `opensearch`: [v0.18.0](services/opensearch/CHANGELOG.md#v0180-2024-09-02)
  - **Breaking changes:** `GetMetricsResponse` fields have changed data types
    - `CpuLoadPercent`, `Load1`, `Load15` and `Load5` are now `*float64`
    - `ParachuteDiskEphemeralActivated` and `ParachuteDiskPersistentActivated` are now `*bool`
- `mariadb`: [v0.19.0](services/mariadb/CHANGELOG.md#v0190-2024-09-02)
  - **Breaking changes:** `GetMetricsResponse` fields have changed data types
    - `CpuLoadPercent`, `Load1`, `Load15` and `Load5` are now `*float64`
    - `ParachuteDiskEphemeralActivated` and `ParachuteDiskPersistentActivated` are now `*bool`
- `logme`: [v0.19.0](services/logme/CHANGELOG.md#v0190-2024-09-02)
  - **Breaking changes:** `GetMetricsResponse` fields have changed data types
    - `CpuLoadPercent`, `Load1`, `Load15` and `Load5` are now `*float64`
    - `ParachuteDiskEphemeralActivated` and `ParachuteDiskPersistentActivated` are now `*bool`
- `mongodbflex`: [v0.15.0](services/mongodbflex/CHANGELOG.md#v0150-2024-09-02)
  - **Feature**: New method `ListAdvisorSlowQueries` that gets slow queries from the Opsmanager performance advisor.
  - **Feature**: New method `ListSuggestedIndexes` that gets suggested indexes from the Opsmanager performance advisor.
  - **Breaking change**: Remove nullable fields.

## Release (2024-08-26)

> [!WARNING]
>
> **The STACKIT Argus service was renamed to STACKIT Observability.**
>
> In the SDK, this means that there is a new `observability` service, which offers the same functionality as the deprecated `argus` service.
>
> SDK updates from now on will be released on the new `observability` service, meaning `argus` will no longer get updates.
>
> Please migrate to `github.com/stackitcloud/stackit-sdk-go/services/observability`.

- `observability`: [v0.1.0](services/observability/CHANGELOG.md#v010-2024-08-21)
  - First release. This module offers the same functionalities as `argus` (release [v0.11.0](https://github.com/stackitcloud/stackit-sdk-go/blob/main/services/argus/CHANGELOG.md#v0110-2024-05-23).
- `iaas`: [v0.8.0](services/iaas/CHANGELOG.md#v080-2024-08-21)
  - **Feature:** `CreateNetworkIPv4Body` and `CreateNetworkIPv6Body` have a new field `Prefix`
- `iaas`: [v0.7.0](services/iaas/CHANGELOG.md#v070-2024-08-16)
  - **Breaking change**: Rename types:
    - `CreateNetworkIPv4` renamed to **`CreateNetworkIPv4Body`**
    - `V1CreateNetworkIPv6` renamed to **`CreateNetworkIPv6Body`**
    - `UpdateNetworkIPv4` renamed to **`UpdateNetworkIPv4Body`**
    - `V1UpdateNetworkIPv6` renamed to **`UpdateNetworkIPv6Body`**
  - **Feature:** `CreateNetworkPayload`, `PartialUpdateNetworkPayload` and `Network` have a new field: `Routed`
- `secretsmanager`: [v0.9.0](services/secretsmanager/CHANGELOG.md#v090-2024-08-16)
  - **Feature**: New API method `UpdateInstance` to update an instance
- `sqlserverflex`: [v0.5.0](services/sqlserverflex/CHANGELOG.md#v050-2024-08-16)
  - **Breaking change**:
    - Fields in `GetBackupResponse` are not nested in an `Item` field (with type `Backup`) anymore
    - `GetBackupResponse` have these new fields: `EndTime`, `Error`, `Id`, `Labels`, `Name`, `Options`, `Size`, `StartTime`
- `loadbalancer`: [v0.15.0](services/loadbalancer/CHANGELOG.md#v0150-2024-08-08)
  - **Feature:** New API method `ListPlans` to list the available service plans
- `iaas`: [v0.6.0](services/iaas/CHANGELOG.md#v060-2024-08-05)
  - **Breaking change:** Use network ID instead of request ID in the waiter: `CreateNetworkWaitHandler`
- `rabbitmq`: [v0.18.0](services/rabbitmq/CHANGELOG.md#v0180-2024-08-01)
  - **Feature:** `Plan` has a new field `SkuName`
- `opensearch`: [v0.17.0](services/opensearch/CHANGELOG.md#v0170-2024-08-01)
  - **Feature:** `Plan` has a new field `SkuName`
- `mariadb`: [v0.18.0](services/mariadb/CHANGELOG.md#v0180-2024-08-01)
  - **Feature:** `Plan` has a new field `SkuName`
- `logme`: [v0.18.0](services/logme/CHANGELOG.md#v0180-2024-08-01)
  - **Feature:** `Plan` has a new field `SkuName`
- `redis`: [v0.18.0](services/redis/CHANGELOG.md#v0180-2024-08-01)
  - **Feature:** `Plan` has a new field `SkuName`

## Release (2024-07-24)

- `iaas`: [v0.5.0](services/iaas/CHANGELOG.md#v050-2024-07-24)
  - **Feature:** `CreateNetworkAddressFamily` and `UpdateNetworkAddressFamily` have a new field `Ipv6`
  - **Feature:** `Network` has new fields: `NameserversV6` and `PrefixesV6`
- `runcommand`: [v0.1.0](services/runcommand/CHANGELOG.md#v010-2024-07-19)
  - **New**: STACKIT Run Command module can be used to run remote commands and custom scripts on VMs
- `sqlserverflex`: [v0.4.0](services/sqlserverflex/CHANGELOG.md#v040-2024-07-19)
  - **Feature:** New field for `DatabaseOptions` and `SingleDatabaseOptions`: `CollationName`
  - **Breaking changes:**
    - Fields removed from `DatabaseOptions` and `SingleDatabaseOptions`: `IsEncrypted`, `RecoveryModel`, `UserAccess`
    - Fields removed from `SingleDatabase`: `Collation`, `CreateDate`
- `ske`: [v0.19.0](services/ske/CHANGELOG.md#v0190-2024-07-18)
  - **Feature:** New fields for `Extension`: `Dns`
- `rabbitmq`: [v0.17.0](services/rabbitmq/CHANGELOG.md#v0161-2024-07-10)
  - **Bugfix:** Fix marking of deprecated struct fields. Potential breaking change for users with linters that treat deprecations as errors.
- `opensearch`: [v0.16.0](services/opensearch/CHANGELOG.md#v0151-2024-07-10)
  - **Bugfix:** Fix marking of deprecated struct fields. Potential breaking change for users with linters that treat deprecations as errors.
- `mariadb`: [v0.17.0](services/mariadb/CHANGELOG.md#v0161-2024-07-10)
  - **Bugfix:** Fix marking of deprecated struct fields. Potential breaking change for users with linters that treat deprecations as errors.
- `logme`: [v0.17.0](services/logme/CHANGELOG.md#v0161-2024-07-10)
  - **Bugfix:** Fix marking of deprecated struct fields. Potential breaking change for users with linters that treat deprecations as errors.
- `redis`: [v0.17.0](services/redis/CHANGELOG.md#v0161-2024-07-10)
  - **Bugfix:** Fix marking of deprecated struct fields. Potential breaking change for users with linters that treat deprecations as errors.
- `loadbalancer`: [v0.14.0](services/loadbalancer/CHANGELOG.md#v0131-2024-07-10)
  - **Bugfix:** Fix marking of deprecated methods. Potential breaking change for users with linters that treat deprecations as errors.
- `ske`: [v0.18.0](services/ske/CHANGELOG.md#v0171-2024-07-10)
  - **Bugfix:** Fix marking of deprecated methods. Potential breaking change for users with linters that treat deprecations as errors.
- `sqlserverflex`: [v0.3.0](services/sqlserverflex/CHANGELOG.md#v030-2024-07-09)
  - **Breaking changes:**
    - `Database` renamed to `DefaultDatabase`, in `CreateUserPayload`
    - Type of `Roles` changed from `[]Role` to `[]string`, in `CreateUserPayload`
    - `User` renamed to `SingleUser`, in `CreateUserResponse`
    - `OwnerName` renamed to `Owner`, in `DatabaseOptions`
    - Fields in `GetDatabaseResponse` are now nested in a `Database` field (with type `SingleDatabase`)
    - `GetDatabaseResponseOptions` renamed to `SingleDatabaseOptions` (and `OwnerName` renamed to `Owner`)
- `archiving`: [v0.1.0](services/archiving/CHANGELOG.md#v010-2024-07-04)
  - Manage your STACKIT Archiving instance with: `CreateInstance`, `DeleteInstance`, `GetInstance`, `ListInstances` and `PartialUpdateInstance`.
- `ske`: [v0.17.0](services/ske/CHANGELOG.md#v0170-2024-07-04)
  - **Feature:** Add new field `AllowSystemComponents` to the `Nodepool` model that configures wether system components are allowed to run on the node pool.
- `loadbalancer`: [v0.XX.X](services/loadbalancer/CHANGELOG.md#v0xxx-2024-xx-xx)
  - **Improvement:** Improve default error messages.
- `serviceenablement`: [v0.2.0](services/serviceenablement/CHANGELOG.md#v020-2024-07-12)
  - **Feature**: New waiters `EnableServiceWaitHandler` and `DisableServiceWaitHandler` for async operations `EnableService` and `DisableService`, respectively.

## Release (2024-07-01)

> **The STACKIT PostgreSQL service has reached its end of support on June 30th 2024. All calls done to the API have stopped working since then.**
>
> Use `github.com/stackitcloud/stackit-sdk-go/services/postgresflex` instead. For more details, check https://docs.stackit.cloud/stackit/en/bring-your-data-to-stackit-postgresql-flex-138347648.html.

- `postgresflex`: [v0.15.0](services/postgresflex/CHANGELOG.md#v0150-2024-06-28)
  - **Feature**: New API methods `CreateDatabase`, `DeleteDatabase`, `ListDatabase`, `ListDatabaseParameters` to manage PostgreSQL Flex databases
  - **Feature**: New API method `UpdateInstance` to update the instance
  - **Feature**: New API method `ListMetrics` to list metrics of an instance
  - **Feature**: New API method `DisableService` to terminate the whole project
- `logme`: [v0.16.0](services/logme/CHANGELOG.md#v0160-2024-06-19)
  - **Feature**: New methods `CreateBackup`, `DownloadBackup`, `ListRestores`,`UpdateBackupsConfig`, `TriggerRecreate`, `TriggerRestart`, `TriggerRestore` to manage the backup and restoration of an instance.
  - **Breaking change**: `Groks` parameter in `InstanceParameters` model is now of type `InstanceParametersGroksInner` (previously `map[string]interface{}`)
- `mariadb`: [v0.16.0](services/mariadb/CHANGELOG.md#v0160-2024-06-19)
  - **Feature**: New methods `CreateBackup`, `DownloadBackup`, `ListRestores`,`UpdateBackupsConfig`, `TriggerRecreate`, `TriggerRestart`, `TriggerRestore` to manage the backup and restoration of an instance.
- `opensearch`: [v0.15.0](services/opensearch/CHANGELOG.md#v0150-2024-06-19)
  - **Feature**: New methods `CreateBackup`, `DownloadBackup`, `ListRestores`,`UpdateBackupsConfig`, `TriggerRecreate`, `TriggerRestart`, `TriggerRestore` to manage the backup and restoration of an instance.
- `rabbitmq`: [v0.16.0](services/rabbitmq/CHANGELOG.md#v0160-2024-06-19)
  - **Feature**: New methods `CreateBackup`, `DownloadBackup`, `ListRestores`,`UpdateBackupsConfig`, `TriggerRecreate`, `TriggerRestart`, `TriggerRestore` to manage the backup and restoration of an instance.
- `redis`: [v0.16.0](services/redis/CHANGELOG.md#v0160-2024-06-19)
  - **Feature**: New methods `CreateBackup`, `DownloadBackup`, `ListRestores`,`UpdateBackupsConfig`, `TriggerRecreate`, `TriggerRestart`, `TriggerRestore` to manage the backup and restoration of an instance.

## Release (2024-06-14)

- `resourcemanager`: [v0.9.0](services/resourcemanager/CHANGELOG.md#v090-2024-06-14)
  - **Breaking Change**: Rename data types for uniformity
    - `ProjectResponse` -> `Project`
    - `ProjectResponseWithParents` -> `GetProjectResponse`
    - `AllProjectsResponse` -> `ListProjectsResponse`
  - **Breaking Change**: Delete unused data types
  - **Feature**: New methods `GetOrganization` and `ListOrganizations`
- `objectstorage`: [v0.10.0](services/objectstorage/CHANGELOG.md#v0100-2024-06-14)
  - **Breaking change**: Remove unused data types.
- `iaas`: [v0.4.0](services/iaas/CHANGELOG.md#v040-2024-06-07)
  - **Breaking change**: `CreateNetwork` now returns the `Network` triggered by the operation.
- `loadbalancer`: [v0.13.0](services/loadbalancer/CHANGELOG.md#v0130-2024-06-12)
  - **Feature:** `LoadBalancer`, `CreateLoadBalancerPayload` and `UpdateLoadBalancerPayload` have a new field `PlanId`
- `secretsmanager`: [v0.8.0](services/secretsmanager/CHANGELOG.md#v080-2024-05-23)
  - **Breaking change**: Rename data types for uniformity
    - `Acl` is now `ACL`
    - `AclList` is now `ListACLsResponse`
    - `InstanceList` is now `ListInstancesResponse`
    - `UserList` is now `ListUsersResponse`
  - **Breaking change**: Remove unused data types
- `serverbackup`: [v0.1.0](services/serverbackup/CHANGELOG.md#v010-2024-05-23)
  - Manage your STACKIT Server Backups: `Backup`, `BackupSchedule`, `VolumeBackup`
- `argus`: [v0.11.0](services/argus/CHANGELOG.md#v0110-2024-05-23)
  - **Feature**: New methods `GetMetricsStorageRetention`, `UpdateMetricsStorageRetention`
  - **Breaking change**: Remove unused data types
- `dns`: [v0.10.0](services/dns/CHANGELOG.md#v0100-2024-05-23)
  - **Feature**: New method `CloneZone` to clone an existing zone with all record sets to a new zone with a different name
  - **Feature**: New methods `CreateLabel`, `DeleteLabel` and `ListLabels` to manage labels for a zone
  - **Feature**: New methods `CreateMoveCode`, `DeleteMoveCode` and `ValidateMoveCode` to manage move codes to move a zone to another project
  - **Feature**: New method `MoveZone` to move a zone to another project
  - **Feature**: New methods `ExportRecordSets` and `ImportRecordSets`
  - **Feature**: New methods `RestoreZone` and `RestoreRecordSet` to restore inactive zones and record-sets, respectively
  - **Feature**: New method `RetrieveZone` to queue a secondary zone for a zone transfer request
- `sqlserverflex`: [v0.2.0](services/sqlserverflex/CHANGELOG.md#v020-2024-05-24)
  - **Feature** Waiters for async operations `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, and `DeleteInstanceWaitHandler`
- `ske`: [v0.16.0](services/ske/CHANGELOG.md#v0160-2024-05-27)
  - **Breaking change:** Renamed data types:
  - `V1Network` is now `Network`
  - `V1LoginKubeConfig` is now `LoginKubeConfig`
- `rabbitmq`: [v0.15.0](services/rabbitmq/CHANGELOG.md#v0150-2024-05-29)
  - **Feature**: `GetMetricsResponse` has new fields: `Load1`, `Load15`, `Load5`
  - **Feature**: `Credentials` has a new field: `Mqtt`, `Stomp`
- `mariadb`: [v0.15.0](services/mariadb/CHANGELOG.md#v0150-2024-05-29)
  - **Feature**: `GetMetricsResponse` has new fields: `Load1`, `Load15`, `Load5`
  - **Breaking change**: Deleted unused data type
- `redis`: [v0.15.0](services/redis/CHANGELOG.md#v0150-2024-05-29)
  - **Feature**: `GetMetricsResponse` has new fields: `Load1`, `Load15`, `Load5`
  - **Breaking change**: Deleted unused data type
- `logme`: [v0.15.0](services/logme/CHANGELOG.md#v0150-2024-05-29)
  - **Feature**: `GetMetricsResponse` has new fields: `Load1`, `Load15`, `Load5`, `OpenSearchDashboardUrl`
  - **Breaking change**: Deleted unused data type

## Release (2024-05-22)

- `authorization`: [v0.3.0](services/authorization/CHANGELOG.md#v030-2024-05-22)
  - **Feature:** New field for `Role`: `Id`
- `mongodbflex`: [v0.14.0](services/mongodbflex/CHANGELOG.md#v0140-2024-05-22)
  - **Breaking change**: Remove unused data types.
- `postgresflex`: [v0.14.0](services/postgresflex/CHANGELOG.md#v0140-2024-05-22)
  - **Breaking change**: Remove unused model data types.
- `sqlserverflex`: [v0.1.0](services/sqlserverflex/CHANGELOG.md#v010-2024-05-22)
  - Manage your STACKIT SQL Server Flex resources: `Instance`, `Flavors`, `Users`, `Databases`, `Backups`
- `ske`: [v0.14.0](services/ske/CHANGELOG.md#v0140-2024-05-03)
  - **Feature:** New fields for `MachineType`: `Architecture`, `Gpu`
- `ske`: [v0.15.0](services/ske/CHANGELOG.md#v0150-2024-05-14)
  - **Feature:** New operation `GetLoginKubeconfig` to get a Kubeconfig for use with the STACKIT CLI. A Kubeconfig retrieved using this endpoint does not contain any credentials and instead obtains valid credentials via the STACKIT CLI.
- `iaas`: [v0.1.0](services/iaas/CHANGELOG.md#v010-2024-05-10)
  - **New BETA module**: manage Infrastructure as a Service (IaaS) resources `Network` and `NetworkArea`
- `iaas`: [v0.2.0](services/iaas/CHANGELOG.md#v020-2024-05-17)
  - **Feature**: New methods to manage networks:
    - `CreateNetwork`
    - `PartialUpdateNetwork`
    - `DeleteNetwork`
  - **Breaking change**: Rename methods for better correspondence with endpoint behaviour (see service [release notes](services/iaas/CHANGELOG.md#v020-2024-05-17) for detailed changes)
  - **Breaking change**: Rename types (see service [release notes](services/iaas/CHANGELOG.md#v020-2024-05-17) for detailed changes)
    - Add `Response` suffix to types only used in method responses
    - Remove `V1` prefix from all types
- `iaas`: [v0.3.0](services/iaas/CHANGELOG.md#v030-2024-05-17)
  - **Feature**: Add waiters for async operations: `CreateNetworkAreaWaitHandler`, `UpdateNetworkAreaWaitHandler`, `DeleteNetworkAreaWaitHandler`, `CreateNetworkWaitHandler`, `UpdateNetworkWaitHandler`, `DeleteNetworkWaitHandler`
- `logme`: [v0.14.0](services/logme/CHANGELOG.md#v0140-2024-05-13)
  - **Feature**: New method `GetMetrics` to get the latest metrics for cpu load, memory and disk usage for an instance
  - **Feature**: New method `ListBackups` to list the backups for an instance
  - **Breaking change**: `ListMetricsResponse` type (previously unused) renamed to `GetMetricsResponse`
  - **Breaking change**: Deleted unused data types
- `mariadb`: [v0.14.0](services/mariadb/CHANGELOG.md#v0140-2024-05-13)
  - **Feature**: New method `GetMetrics` to get the latest metrics for cpu load, memory and disk usage for an instance
  - **Feature**: New method `ListBackups` to list the backups for an instance
  - **Breaking change**: `ListMetricsResponse` type (previously unused) renamed to `GetMetricsResponse`
  - **Breaking change**: Deleted unused data types
- `opensearch`: [v0.14.0](services/opensearch/CHANGELOG.md#v0140-2024-05-13)
  - **Feature**: New method `GetMetrics` to get the latest metrics for cpu load, memory and disk usage for an instance
  - **Feature**: New method `ListBackups` to list the backups for an instance
  - **Breaking change**: `ListMetricsResponse` type (previously unused) renamed to `GetMetricsResponse`
  - **Breaking change**: Deleted unused data types
- `rabbitmq`: [v0.14.0](services/rabbitmq/CHANGELOG.md#v0140-2024-05-13)
  - **Feature**: New method `GetMetrics` to get the latest metrics for cpu load, memory and disk usage for an instance
  - **Feature**: New method `ListBackups` to list the backups for an instance
  - **Breaking change**: `ListMetricsResponse` type (previously unused) renamed to `GetMetricsResponse`
  - **Breaking change**: Deleted unused data types
- `redis`: [v0.14.0](services/redis/CHANGELOG.md#v0140-2024-05-13)
  - **Feature**: New method `GetMetrics` to get the latest metrics for cpu load, memory and disk usage for an instance
  - **Feature**: New method `ListBackups` to list the backups for an instance
  - **Breaking change**: `ListMetricsResponse` type (previously unused) renamed to `GetMetricsResponse`
  - **Breaking change**: Deleted unused data types
- `serviceenablement`: [v0.1.0](services/serviceenablement/CHANGELOG.md#v010-2024-05-17)
  - **New**: STACKIT Service Enablement module can be used to enable services

## Release (2024-05-02)

- `ske`: [v0.13.0](services/ske/CHANGELOG.md#v0130-2024-04-16)
  - **Deprecation**: The following methods have been deprecated and the [Service Enablement API](https://docs.api.stackit.cloud/documentation/service-enablement/version/v1) must be used instead.
    - `DisableService`
    - `EnableService`
    - `GetServiceStatus`

## Release (2024-04-12)

- `core`: [v0.12.0](core/CHANGELOG.md#v0120-2024-04-11)
  - **Feature:** Add `Middleware` type, `WithMiddleware` and `ChainMiddleware` methods to package `config`, this allows clients to chain and add Middlewares to the transport layer of the HTTP client.
- `core`: [v0.11.0](core/CHANGELOG.md#v0110-2024-04-09)
  - **Feature:** Add method `WithCaptureHTTPRequest` to package `runtime`, which allows capture of HTTP requests for debugging purposes.
- `loadbalancer`: [v0.12.0](services/loadbalancer/CHANGELOG.md#v0120-2024-04-12)
  - **Feature:** Set `config.ContextHTTPRequest` in `Execute` methods
  - **Feature:** New API method `GetQuota` to get the maximum number of load balancing servers allowed for a project
  - **Feature:** New API method `UpdateCredentials` to update the credentials for observability in a project
- `loadbalancer`: [v0.11.0](services/loadbalancer/CHANGELOG.md#v0110-2024-04-11)
  - **Feature:** Support WithMiddleware configuration option in the client
- `loadbalancer`: [v0.10.0](services/loadbalancer/CHANGELOG.md#v0100-2024-04-08)
  - **Deprecation:** Mark methods `EnableService` and `DisableService` as deprecated. Enablement and disablement of the load balancer functionality is now automaticly handled by the service.
- `logme`: [v0.13.0](services/logme/CHANGELOG.md#v0130-2024-04-11)
  - **Breaking change**: Fields removed from `RawCredentials`: `RouteServiceUrl`, `SyslogDrainUrl`, `VolumeMounts`.
  - **Breaking change**: Fields removed from `Credentials`: `Hosts`, `HttpApiUri`, `Name`, `Protocols`
  - **Feature**: `Credentials` has a new field `SyslogDrainUrl`
- `mariadb`: [v0.13.0](services/mariadb/CHANGELOG.md#v0130-2024-04-11)
  - **Breaking change**: Fields removed from `RawCredentials`: `RouteServiceUrl`, `SyslogDrainUrl`, `VolumeMounts`.
  - **Breaking change**: Fields removed from `Credentials`: `HttpApiUri`, `Protocols`
- `opensearch`: [v0.13.0](services/opensearch/CHANGELOG.md#v0130-2024-04-11)
  - **Breaking change**: Fields removed from `RawCredentials`: `RouteServiceUrl`, `SyslogDrainUrl`, `VolumeMounts`.
  - **Breaking change**: Fields removed from `Credentials`: `HttpApiUri`, `Name`, `Protocols`
  - **Feature**: `Credentials` has a new field `Scheme`
- `postgresflex`: [v0.12.0](services/postgresflex/CHANGELOG.md#v0120-2024-04-03)
  - **Improvement**: Update `DeleteInstanceWaitHandler` to support new deletion method.
  - **Feature**: New waiter `ForceDeleteInstanceWaitHandler` for async operation `ForceDeleteInstance`
- `rabbitmq`: [v0.13.0](services/rabbitmq/CHANGELOG.md#v0130-2024-04-11)
  - **Breaking change**: Fields removed from `RawCredentials`: `RouteServiceUrl`, `SyslogDrainUrl`, `VolumeMounts`.
  - **Breaking change**: Fields removed from `Credentials`: `Name`, `Protocols`
  - **Feature**: `Credentials` has new fields `HttpApiUris`, `Management`, `Uris`
- `redis`: [v0.13.0](services/redis/CHANGELOG.md#v0130-2024-04-11)
  - **Breaking change**: Fields removed from `RawCredentials`: `RouteServiceUrl`, `SyslogDrainUrl`, `VolumeMounts`.
  - **Breaking change**: Fields removed from `Credentials`: `HttpApiUri`, `Name`, `Protocols`
  - **Feature**: `Credentials` has new fields `LoadBalancedHost`
- `ske`: [v0.11.0](services/ske/CHANGELOG.md#v0110-2024-03-28)
  - **Feature**: Waiters for async operation `StartCredentialsRotationWaitHandler` and `CompleteCredentialsRotationWaitHandler`

## Release (2024-03-20)

- `core`: [v0.10.1](core/CHANGELOG.md#v0101-2024-03-20)
  - **Improvement:** Update `ConfigureRegion` method to take into account global servers without a region variable
- `postgresflex`: [v0.10.0](services/postgresflex/CHANGELOG.md#v0100-2024-03-08)
  - **Feature:** New API method `CloneInstance` to clone the instance.
- `secretsmanager`: [v0.6.0](services/secretsmanager/CHANGELOG.md#v060-2024-03-18)
  - **Feature**: New API method `UpdateACLs` to update all ACLs of an instance
- `loadbalancer`: [v0.9.3](services/loadbalancer/CHANGELOG.md#v093-2024-03-20)
  - **Improvement**: Improve error handling in Load Balancer creation waiter, fixing timeout being exceeded for `STATUS_PENDING` status with errors. If an error is found in the `Errors` field, the waiter now returns with error.

## Release (2024-02-27)

- `core`: [v0.10.0](core/CHANGELOG.md#v0100-2024-02-27)
  - **Feature:** Add package `runtime`, which implements methods to be used when performing API requests.
  - **Feature:** Add method `WithCaptureHTTPResponse` to package `runtime`, which does the same as `config.WithCaptureHTTPResponse`. Method was moved to avoid confusion due to it not being a configuration option, and will be removed in a later release.
  - **Feature:** Add configuration option that, for the key flow, enables a goroutine to be spawned that will refresh the access token when it's close to expiring
  - **Deprecation:** Mark method `config.WithCaptureHTTPResponse` as deprecated, to avoid confusion due to it not being a configuration option. Use `runtime.WithCaptureHTTPResponse` instead.
  - **Deprecation:** Mark method `config.WithJWKSEndpoint` and field `config.Configuration.JWKSCustomUrl` as deprecated. Validation using JWKS was removed, for being redundant with token validation done in the APIs. These have no effect.
  - **Deprecation:**
    - Methods:
      - `config.WithMaxRetries`
      - `config.WithWaitBetweenCalls`
      - `config.WithRetryTimeout`
      - `clients.NewRetryConfig`
    - Fields:
      - `clients.KeyFlowConfig.ClientRetry`
      - `clients.TokenFlowConfig.ClientRetry`
      - `clients.NoAuthFlowConfig.ClientRetry`
      - `clients.RetryConfig`
    - Retry options removed to reduce complexity of the clients. If this functionality is needed, you can provide your own custom HTTP client.
  - **Breaking Change:** Change signature of `auth.NoAuth`, which no longer takes `clients.RetryConfig` as argument.
  - **Breaking Change:**
    - Methods:
      - `clients.KeyFlow.Clone`
      - `clients.TokenFlow.Clone`
      - `clients.NoAuthFlow.Clone`
      - `clients.Do`
    - Fields:
      - `clients.DefaultRetryMaxRetries`
      - `clients.DefaultRetryWaitBetweenCalls`
      - `clients.DefaultRetryTimeout`
    - Constants:
      - `clients.ClientTimeoutErr`
      - `clients.ClientContextDeadlineErr`
      - `clients.ClientConnectionRefusedErr`
      - `clients.ClientEOFError`
      - `clients.Environment`
    - Removed to reduce complexity of the clients, they were no longer being used.

## Release (2024-02-07)

> The `membership` module has been replaced with the `authorization` module, which connects to the same API.
>
> **This module will receive no further updates.** Use `github.com/stackitcloud/stackit-sdk-go/services/authorization` instead.

### Highlights

- `authorization`: [v0.1.0](services/authorization/CHANGELOG.md#v010-2024-02-07)
  - First release. This module offers the same functionalities as `membership` (release v0.4.0).
- `membership`: [v0.4.0](services/membership/CHANGELOG.md#v040-2024-02-07)
  - Add deprecation note

## Release (2024-02-06)

> **The STACKIT PostgreSQL service will reach its end of support on June 30th 2024. All calls done to the API after that will stop working.**
>
> Use `github.com/stackitcloud/stackit-sdk-go/services/postgresflex` instead. For more details, check https://docs.stackit.cloud/stackit/en/bring-your-data-to-stackit-postgresql-flex-138347648.html.

### Highlights

- `postgresql`: [v0.12.0](services/postgresql/CHANGELOG.md#v0120-2024-02-06)
  - Add deprecation note
- `ske`: [v0.10.0](services/ske/CHANGELOG.md#v0100-2024-02-06)
  - **Feature:** New endpoints for credentials rotation.
    - `StartCredentialsRotation`
    - `CompleteCredentialsRotation`
    - `CreateKubeconfig`
    - These endpoints replace `GetCredentials` and `TriggerRotateCredentials`, which are **deprecated** and will not work for clusters with Kubernetes v1.27+, or if the new endpoints for kubeconfig or credentials rotation have already been used. For more information, see [How to rotate SKE credentials](https://docs.stackit.cloud/display/STACKIT/How+to+rotate+SKE+credentials#tabs-237293ce-f625-44ea-9d4f-689e31f596d6-1).

## Release (2024-02-05)

### Highlights

- `logme`: [v0.10.0](services/logme/CHANGELOG.md#v0100-2024-02-02)
  - **Feature**: `Instance` has a new field `OfferingName`
- `mariadb`: [v0.10.0](services/mariadb/CHANGELOG.md#v0100-2024-02-02)
  - **Feature**: `Instance` has a new field `OfferingName`
- `opensearch`: [v0.10.0](services/opensearch/CHANGELOG.md#v0100-2024-02-02)
  - **Feature**: `Instance` has a new field `OfferingName`
- `postgresflex`: [v0.9.0](services/postgresflex/CHANGELOG.md#v090-2024-02-05)
  - **Feature**: New API method `UpdateUser` to update user
  - **Feature**: New API method `PartialUpdateUser` to patch update user
  - **Feature**: New API method `ResetUser` to reset a user's password
- `postgresql`: [v0.11.0](services/postgresql/CHANGELOG.md#v0110-2024-02-02)
  - **Feature**: `Instance` has a new field `OfferingName`
- `rabbitmq`: [v0.10.0](services/rabbitmq/CHANGELOG.md#v0100-2024-02-02)
  - **Feature**: `Instance` has a new field `OfferingName`
- `redis`: [v0.10.0](services/redis/CHANGELOG.md#v0100-2024-02-02)
  - **Feature**: `Instance` has a new field `OfferingName`

## Release (2024-01-24)

### Highlights

- `loadbalancer`: [v0.9.0](services/loadbalancer/CHANGELOG.md#v090-2024-01-24)
  - **Feature**: Server Name Indicator (SNI) support
  - **Feature**: Layer 4 Session Persistance

## Release (2024-01-09)

### Highlights

- `core`: [v0.7.5](core/CHANGELOG.md#v075-2024-01-09)
  - **Improvement:** When using the key flow, the SDK will extract the private key from the service account key and use it, if no private key is provided in the configuration, through environment variable or in the credentials file. This makes it simpler to use the key flow: if you create a service account key including the private key, you don't need to provide the private key separately anymore
- `dns`: [v0.8.0](services/dns/CHANGELOG.md#v080-2024-01-09)
  - **Feature:** `Zone` has a new filed `Labels`, which is an array of labels (key-value pairs) associated to a zone
  - **Feature:** `ListZones` can be filtered by label keys or values
  - **Feature:** `CloneZonePayload` has a flag `AdjustRecords` to adjust the record set content of the cloned zone (replaces the dns name of the original zone with the new dns name of the cloned zone)
- `logme`: [v0.9.0](services/logme/CHANGELOG.md#v090-2024-01-09)
  - **Feature:** `PartialUpdateInstance` can be used to update the instance's name
  - **Feature:** `InstanceParameters` has a new setting `MaxDiskThreshold`
  - **Feature:** `ListMetricsResponse` has new fields regarding ephemeral disk
- `mariadb`: [v0.9.0](services/mariadb/CHANGELOG.md#v090-2024-01-09)
  - **Feature:** `PartialUpdateInstance` can be used to update the instance's name
  - **Feature:** `InstanceParameters` has a new setting `MaxDiskThreshold`
  - **Feature:** `ListMetricsResponse` has new fields regarding ephemeral disk
- `mongodbflex`: [v0.10.0](services/mongodbflex/CHANGELOG.md#v0100-2024-01-05)
  - **Breaking Change:** Added method `ListMetrics` (which accepts the desired metric: `cpu`, `database`, `disk-iops`, `disk-use`, or `exec-time`) as an argument and replaces `ListCPUMetrics`, `ListDatabaseStorageMetrics`, `ListDiskIOPSMetrics`, `ListDiskUsageMetrics`, and `ListExecutionTimesMetrics`, respectivelly
  - **Breaking Change:** Removes the previously deprecated method `GetStorage`, replaced by `ListStorages`
- `opensearch`: [v0.9.0](services/opensearch/CHANGELOG.md#v090-2024-01-09)
  - **Feature:** `PartialUpdateInstance` can be used to update the instance's name
  - **Feature:** `InstanceParameters` has a new setting `MaxDiskThreshold`
  - **Feature:** `ListMetricsResponse` has new fields regarding ephemeral disk
- `postgresql`: [v0.10.0](services/postgresql/CHANGELOG.md#v0100-2024-01-09)
  - **Feature:** `PartialUpdateInstance` can be used to update the instance's name
  - **Feature:** `InstanceParameters` has a new setting `MaxDiskThreshold`
  - **Feature:** `ListMetricsResponse` has new fields regarding ephemeral disk
- `rabbitmq`: [v0.9.0](services/rabbitmq/CHANGELOG.md#v090-2024-01-09)
  - **Feature:** `PartialUpdateInstance` can be used to update the instance's name
  - **Feature:** `InstanceParameters` has a new setting `MaxDiskThreshold`
  - **Feature:** `ListMetricsResponse` has new fields regarding ephemeral disk
- `redis`: [v0.9.0](services/redis/CHANGELOG.md#v090-2024-01-09)
  - **Feature:** `PartialUpdateInstance` can be used to update the instance's name
  - **Feature:** `InstanceParameters` has a new setting `MaxDiskThreshold`
  - **Feature:** `ListMetricsResponse` has new fields regarding ephemeral disk
- `ske`: [v0.9.0](services/ske/CHANGELOG.md#v090-2024-01-09)
  - **Improvement:** Add details on credentials for old clusters
  - **Feature:** `ClusterStatus` now has a field `CredentialsRotation` with credentials' details

## Release (2023-12-22)

### Highlights

- `mongodbflex`: [v0.9.0](services/mongodbflex/CHANGELOG.md#v090-2023-12-22)
  - Added struct `ApiListStoragesRequest`, which will replace `ApiGetStorageRequest`
    - `ApiGetStorageRequest` has been marked as deprecated, and will be removed in the next minor update
  - Added method `ListStorages`, which will replace `GetStorage`
    - `GetStorage` has been marked as deprecated, and will be removed in the next minor update
  - Added `CloneInstanceWaitHandler`, to wait for `CloneInstance` async operation to be completed

## Release (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.
Most significant changes:

- Methods to get multiple instances of the same resource are now named `List[Resource Name]`
- Methods to fully update a resource are now named `Update[Resource Name]`
- Methods to update some fields of a resource are now named `PartialUpdate[Resource Name]`
- Methods relative to API service enablement are now named `GetServiceStatus`, `EnableService` and `DisableService`
- Several common terms, such as `ACL` and `Credentials`, have been standardized
- Structs only used in method responses are named `[Method Name]Response`
- Waiters for a given method are now named `[Method name]WaitHandler`

### Highlights

Below is the list of changes for the API methods. For each service, you can check out the full changelog.

- `argus`: [v0.8.0](services/argus/CHANGELOG.md#v080-2023-12-20)
  - `CreateCredential` renamed to `CreateCredentials`
  - `CreateInstanceAlertConfigReceiver` renamed to `CreateAlertConfigReceiver`
  - `DeleteCredential` renamed to `DeleteCredentials`
  - `DeleteCredentialRemoteWriteConfig` renamed to `DeleteCredentialsRemoteWriteConfig`
  - `DeleteInstanceAlertConfigReceiver` renamed to `DeleteAlertConfigReceiver`
  - `DeleteInstanceAlertConfigRouteReceiver` renamed to `DeleteAlertConfigRoute`
  - `GetCredential` renamed to `GetCredentials`
  - `GetCredentialRemoteWriteConfig` renamed to `GetCredentialsRemoteWriteConfig`
  - `GetCredentials` renamed to `ListCredentials`
  - `GetInstanceAcl` renamed to `ListACL`
  - `GetInstanceAlertConfigReceiver` renamed to `GetAlertConfigReceiver`
  - `GetInstanceAlertConfigReceivers` renamed to `ListAlertConfigReceivers`
  - `GetInstanceAlertConfigRoutes` renamed to `ListAlertConfigRoute`
  - `GetInstanceAlertConfigs` renamed to `GetAlertConfigs`
  - `GetInstanceGrafanaConfigs` renamed to `GetGrafanaConfigs`
  - `GetInstances` renamed to `ListInstances`
  - `GetPlans` renamed to `ListPlans`
  - `GetScrapeConfigs` renamed to `ListScrapeConfigs`
  - `UpdateCredentialRemoteWriteConfig` renamed to `UpdateCredentialsRemoteWriteConfig`
  - `UpdateInstanceAcl` renamed to `UpdateACL`
  - `UpdateInstanceAlertConfigReceiver` renamed to `UpdateAlertConfigReceiver`
  - `UpdateInstanceAlertConfigRouteReceiver` renamed to `UpdateAlertConfigRoute`
  - `UpdateInstanceAlertConfigs` renamed to `UpdateAlertConfigs`
  - `UpdateInstanceGrafanaConfigs` renamed to `UpdateGrafanaConfigs`
- `dns`: [v0.7.0](services/dns/CHANGELOG.md#v070-2023-12-20)
  - `GetRecordSets` renamed to `ListRecordSets`
  - `GetZones` renamed to `ListZones`
  - `UpdateRecord` renamed to `PartialUpdateRecord`
  - `UpdateRecordSet` renamed to `PartialUpdateRecordSet`
  - `UpdateZone` renamed to `PartialUpdateZone`
- `loadbalancer`: [v0.8.0](services/loadbalancer/CHANGELOG.md#v080-2023-12-20)
  - `DisableLoadBalancing` renamed to `DisableService`
  - `EnableLoadBalancing` renamed to `EnableService`
  - `GetProjectStatus` renamed to `GetServiceStatus`
- `logme`: [v0.8.0](services/logme/CHANGELOG.md#v080-2023-12-20)
  - `GetCredentialsIds` renamed to `ListCredentials`
  - `GetInstances` renamed to `ListInstances`
  - `GetOfferings` renamed to `ListOfferings`
  - `UpdateInstance` renamed to `PartialUpdateInstance`
- `mariadb`: [v0.8.0](services/mariadb/CHANGELOG.md#v080-2023-12-20)
  - `GetCredentialsIds` renamed to `ListCredentials`
  - `GetInstances` renamed to `ListInstances`
  - `GetOfferings` renamed to `ListOfferings`
  - `UpdateInstance` renamed to `PartialUpdateInstance`
- `membership`: [v0.3.0](services/membership/CHANGELOG.md#v030-2023-12-20)
  - `DeleteMembers` renamed to `RemoveMembers`
  - `GetMembers` renamed to `ListMembers`
  - `GetMemberships` renamed to `ListUserMemberships`
  - `GetPermissions` renamed to `ListPermissions`
  - `GetRoles` renamed to `ListRoles`
  - `UpdateMembers` renamed to `AddMembers`
- `mongodbflex`: [v0.8.0](services/mongodbflex/CHANGELOG.md#v080-2023-12-20)
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
- `objectstorage`: [v0.8.0](services/objectstorage/CHANGELOG.md#v080-2023-12-20)
  - `CreateProject` renamed to `EnableService`
  - `DeleteProject` renamed to `DisableService`
  - `GetAccessKeys` renamed to `ListAccessKeys`
  - `GetBuckets` renamed to `ListBuckets`
  - `GetCredentialsGroups` renamed to `ListCredentialsGroups`
  - `GetProject` renamed to `GetServiceStatus`
- `opensearch`: [v0.8.0](services/opensearch/CHANGELOG.md#v080-2023-12-20)
  - `GetCredentialsIds` renamed to `ListCredentials`
  - `GetInstances` renamed to `ListInstances`
  - `GetOfferings` renamed to `ListOfferings`
  - `UpdateInstance` renamed to `PartialUpdateInstance`
- `postgresflex`: [v0.8.0](services/postgresflex/CHANGELOG.md#v080-2023-12-20)
  - `GetBackups` renamed to `ListBackups`
  - `GetFlavors` renamed to `ListFlavors`
  - `GetInstances` renamed to `ListInstances`
  - `GetStorages` renamed to `ListStorages`
  - `GetUsers` renamed to `ListUsers`
  - `GetVersions` renamed to `ListVersions`
  - `UpdateInstance` renamed to `PartialUpdateInstance`
- `postgresql`: [v0.9.0](services/postgresql/CHANGELOG.md#v090-2023-12-20)
  - `GetCredentialsIds` renamed to `ListCredentials`
  - `GetInstances` renamed to `ListInstances`
  - `GetOfferings` renamed to `ListOfferings`
  - `UpdateInstance` renamed to `PartialUpdateInstance`
- `rabbitmq`: [v0.8.0](services/rabbitmq/CHANGELOG.md#v080-2023-12-20)
  - `GetCredentialsIds` renamed to `ListCredentials`
  - `GetInstances` renamed to `ListInstances`
  - `GetOfferings` renamed to `ListOfferings`
  - `UpdateInstance` renamed to `PartialUpdateInstance`
- `redis`: [v0.8.0](services/redis/CHANGELOG.md#v080-2023-12-20)
  - `GetCredentialsIds` renamed to `ListCredentials`
  - `GetInstances` renamed to `ListInstances`
  - `GetOfferings` renamed to `ListOfferings`
  - `UpdateInstance` renamed to `PartialUpdateInstance`
- `resourcemanager`: [v0.7.0](services/resourcemanager/CHANGELOG.md#v070-2023-12-20)
  - `GetProjects` renamed to `ListProjects`
  - `UpdateProject` renamed to `PartialUpdateProject`
- `secretsmanager`: [v0.5.0](services/secretsmanager/CHANGELOG.md#v050-2023-12-20)
  - `CreateAcl` renamed to `CreateACL`
  - `DeleteAcl` renamed to `DeleteACL`
  - `GetAcl` renamed to `GetACL`
  - `GetAcls` renamed to `ListACLs`
  - `GetInstances` renamed to `ListInstances`
  - `UpdateAcl` renamed to `UpdateACL`
- `serviceaccount`: [v0.3.0](services/serviceaccount/CHANGELOG.md#v030-2023-12-20)
  - `GetAccessTokens` renamed to `ListAccessTokens`
  - `GetServiceAccountJWKS` renamed to `GetJWKS`
  - `GetServiceAccountKeys` renamed to `ListServiceAccountKeys`
  - `GetServiceAccounts` renamed to `ListServiceAccounts`
  - `GetUsers` renamed to `ListUsers`
  - `UpdateServiceAccountKey` renamed to `PartialUpdateServiceAccountKey`
- `ske`: [v0.8.0](services/ske/CHANGELOG.md#v080-2023-12-20)
  - `CreateProject` renamed to `EnableService`
  - `DeleteProject` renamed to `DisableService`
  - `GetClusters` renamed to `ListClusters`
  - `GetOptions` renamed to `ListProviderOptions`
  - `GetProject` renamed to `GetServiceStatus`

## Release (2023-12-18)

This is the first GitHub release of the STACKIT Go SDK.

### Highlights

List of modules:

- `core`: [v0.7.3](core/CHANGELOG.md#v073-2023-12-13)
  - `auth`: setup authentication, specifically using the service account key or token flows. Check our [authentication example](https://github.com/stackitcloud/stackit-sdk-go/blob/main/examples/authentication/authentication.go)
  - `clients`: baseline http client implementations to support different use cases, such as the different authentication flows
  - `config`: configuration for the SDK clients, such as custom endpoints, region and custom http client configuration. Check our [configuration example](https://github.com/stackitcloud/stackit-sdk-go/blob/main/examples/configuration/configuration.go)
  - `oapierror`: open api error definition and handling
  - `utils`: utilities, such as the `Ptr` method to return a pointer to a variable of any type, which can be useful for creating payloads
  - `wait`: functionality to wait until a specific async operation has finished. Check our [waiter example](https://github.com/stackitcloud/stackit-sdk-go/blob/main/examples/waiter/waiter.go)
- `argus`: [v0.7.0](services/argus/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT Argus resources: `Instance`, `Credentials`, `ScrapeConfig`, `Acl`, `Alertconfig`, `GrafanaConfig`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateScrapeConfigWaitHandler`, `DeleteScrapeConfigWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/argus)
- `dns`: [v0.6.0](services/dns/CHANGELOG.md#v060-2023-11-10)
  - Manage your STACKIT DNS resources: `Zones`, `RecordSet`
  - Waiters for async operations: `CreateZoneWaitHandler`, `UpdateZoneWaitHandler`, `DeleteZoneWaitHandler`, `CreateRecordSetWaitHandler`, `UpdateRecordSetWaitHandler`, `DeleteRecordSetWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/dns)
- `loadbalancer`: [v0.7.0](services/loadbalancer/CHANGELOG.md#v070-2023-12-06)
  - Manage your STACKIT Load Balancer resources: `LoadBalancer`, `Credentials`
  - Waiters for async operations: `CreateLoadBalancerWaitHandler`, `DeleteLoadBalancerWaitHandler`, `EnableLoadBalancingWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/loadbalancer)
- `logme`: [v0.7.0](services/logme/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT Logme resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/logme)
- `mariadb`: [v0.7.0](services/mariadb/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT MariaDB resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/mariadb)
- `membership`: [v0.2.0](services/membership/CHANGELOG.md#v020-2023-11-10)
  - Manage membership of your STACKIT resources
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/membership)
- `mongodbflex`: [v0.7.0](services/mongodbflex/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT MongoDB Flex resources: `Instance`, `Flavors`, `Metrics`, `User`, `Storages`, `Versions`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/mongodbflex)
- `objectstorage`: [v0.7.0](services/objectstorage/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT Object Storage resources: `Bucket`, `AccessKey`, `CredentialGroup`
  - Waiters for async operations: `CreateBucketWaitHandler`, `DeleteBucketWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/objectstorage)
- `opensearch`: [v0.7.0](services/opensearch/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT OpenSearch resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/opensearch)
- `postgresflex`: [v0.7.0](services/postgresflex/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT PostgreSQL Flex resources: `Instance`, `Versions`, `Flavors`, `User`, `Storages`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `DeleteUserWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/postgresflex)
- `postgresql`: [v0.8.0](services/postgresql/CHANGELOG.md#v080-2023-11-17)
  - Manage your STACKIT PostgreSQL resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/postgresql)
- `rabbitmq`: [v0.7.0](services/rabbitmq/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT RabbitMQ resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/rabbitmq)
- `redis`: [v0.7.0](services/redis/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT Redis resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/redis)
- `resourcemanager`: [v0.6.0](services/resourcemanager/CHANGELOG.md#v060-2023-11-10)
  - Manage your STACKIT projects
  - Waiters for async operations: `CreateProjectWaitHandler`, `DeleteProjectWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/resourcemanager)
- `secretsmanager`: [v0.4.0](services/secretsmanager/CHANGELOG.md#v040-2023-11-10)
  - Manage your STACKIT Secrets Manager resources: `Instance`, `Acl`, `User`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/secretsmanager)
- `serviceaccount`: [v0.2.0](services/serviceaccount/CHANGELOG.md#v020-2023-11-10)
  - Manage your STACKIT service accounts
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/serviceaccount)
- `ske`: [v0.7.0](services/ske/CHANGELOG.md#v070-2023-12-05)
  - Manage your STACKIT Kubernetes Engine resources: `Project`, `Cluster`, `Credentials`, `Options`
  - Waiters for async operations: `CreateOrUpdateClusterWaitHandler`, `DeleteClusterWaitHandler`, `CreateProjectWaitHandler`, `DeleteProjectWaitHandler`, `RotateCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/ske)
