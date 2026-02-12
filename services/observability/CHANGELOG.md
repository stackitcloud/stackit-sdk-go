## v0.17.0
- **Feature:** add AlertRecord
    - new related operations: `CreateAlertrecord`, `DeleteAlertrecord`, `DeleteAlertrecords`, `GetAlertrecord`, `ListAlertrecords`, `PartialUpdateAlertrecords`, `UpdateAlertrecord`
    - new related models: `UpdateAlertrecordPayload`, `PartialUpdateAlertrecordsRequestInner`, `CreateAlertrecordPayload`, `AlertrecordResponse`, `AlertRecordResponse`, `AlertRecordsResponse`
- **Feature:** add Backups
    - new related operations: `CreateBackup`, `CreateBackupSchedule`, `ListBackupRetentions`, `ListBackupSchedules`, `ListBackups`, `RestoreBackup`
    - new related models: `BackupResponse`, `BackupRetentionResponse`, `BackupSchedule`, `BackupSchedulePostResponse`, `BackupScheduleResponse`, `CreateBackupSchedulePayload`
- **Feature:** add Elasticsearch checks
    - new related operations `CreateElasticsearchCheck`, `DeleteElasticsearchCheck`, `ListElasticsearchChecks`
    - new related models `ElasticsearchCheckChildResponse`, `ElasticsearchCheckResponse` `CreateElasticsearchCheckPayload`
- **Feature:** add Mongodb checks
    - new related operations `CreateMongodbCheck`, `ListMongodbChecks`, `DeleteMongodbCheck`
    - new related models `MongodbCheckChildResponse`, `MongodbCheckResponse`, `CreateMongodbCheckPayload`
- **Feature:** add Mysql checks
    - new related operations `CreateMysqlCheck`, `DeleteMysqlCheck`, `ListMysqlChecks`
    - new related models `MysqlCheckChildResponse`, `MysqlCheckResponse`, `CreateMysqlCheckPayload`
- **Feature:** add Network checks
    - new related operations `CreateNetworkCheck`, `DeleteNetworkCheck`, `ListNetworkChecks`
    - new related models `NetworkCheckChildResponse`, `NetworkCheckResponse`, `CreateNetworkCheckPayload`
- **Feature:** add Ping checks
    - new related operations `CreatePingCheck`, `DeletePingCheck`, `ListPingChecks`
    - new related models `PingCheckChildResponse`, `PingCheckResponse`, `CreatePingCheckPayload`
- **Feature:** add Postgresql checks
    - new related operations `CreatePostgresqlCheck`, `DeletePostgresqlCheck`, `ListPostgresqlChecks`
    - new related models `PostgresqlCheckChildResponse`, `PostgresqlCheckResponse`, `CreatePostgresqlCheckPayload`
- **Feature:** add Rabbitmq checks
    - new related operations `CreateRabbitmqCheck`, `DeleteRabbitmqCheck`, `ListRabbitmqChecks`
    - new related models `RabbitmqCheckChildResponse`, `RabbitmqCheckResponse`, `CreateRabbitmqCheckPayload`
- **Feature:** add Redis checks
    - new related operations `CreateRedisCheck`, `ListRedisChecks`, `DeleteRedisCheck`
    - new related models `RedisCheckChildResponse`, `RedisCheckResponse`, `CreateRedisCheckPayload`
- **Feature:** extended Alertrules
    - new related operations `DeleteAlertrule`, `GetAlertrule`, `UpdateAlertrule`
    - new related models `AlertruleResponse`, `UpdateAlertrulePayload`
- **Feature:** add Offerings
    - new related operation `ListOfferings`
    - new related model `Offerings`
-  **Feature:** extended Scrape
- new related operations `DeleteScrapeConfigs`, `PartialUpdateScrapeConfigs`
- ️**Breaking change:** renamed AlertRule to AlertruleResponse
- ️**Breaking change:** renamed CreateScrapeConfigPayloadHttpSdConfigsInner to PartialUpdateScrapeConfigsRequestInnerHttpSdConfigsInner
- ️**Breaking change:** renamed CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2 to PartialUpdateScrapeConfigsRequestInnerHttpSdConfigsInner
- ️**Breaking change:** renamed CreateScrapeConfigPayloadHttpSdConfigsInnerOauth2TlsConfig to PartialUpdateScrapeConfigsRequestInnerHttpSdConfigsInnerTlsConfig
- ️**Breaking change:** renamed CreateScrapeConfigPayloadMetricsRelabelConfigsInner to PartialUpdateScrapeConfigsRequestInnerMetricsRelabelConfigsInner
- ️**Breaking change:** renamed CreateScrapeConfigPayloadStaticConfigsInner to PartialUpdateScrapeConfigsRequestInnerStaticConfigsInner
- ️**Breaking change:** renamed CreateScrapeConfigPayloadBasicAuth to PartialUpdateScrapeConfigsRequestInnerBasicAuth

## v0.16.3
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v0.16.2
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v0.16.1
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v0.16.0
- **Breaking change:** The `PartialUpdateAlertrules` takes now `PartialUpdateAlertrulesRequestInner` instead of `UpdateAlertgroupsRequestInnerRulesInner`.
- **Breaking change:** The type of `Rules` in `CreateLogsAlertgroupsPayload` and `UpdateLogsAlertgroupPayload` has changed from `[]UpdateAlertgroupsRequestInnerRulesInner` to `[]CreateLogsAlertgroupsPayloadRulesInner`.
- **Deprecation:** The `GrafanaAdminPassword` and `GrafanaAdminUser` fields are now deprecated in `InstanceSensitiveData` model
- **Feature:** Add `GrafanaAdminEnabled` to `CreateInstancePayload` and `UpdateInstancePayload` models.
- **Feature:** Add new field `record` in `UpdateAlertgroupsRequestInnerRulesInner` model
- **Feature:** Add `CertCheck` to `CertCheckResponse` model.
- **Feature:** Add `HttpCheck` to `HttpCheckResponse` model.
- **Feature:** Add new `CreateLogsAlertgroupsPayloadRulesInner` model.
- **Feature:** Add `allowAssignGrafanaAdmin` to `GrafanaOauth` and `UpdateGrafanaConfigsPayloadGenericOauth` models.
- **Feature:** Add `GrafanaAdminEnabled` to `InstanceSensitiveData` model.
- **Feature:** Add new `PartialUpdateAlertrulesRequestInner` model.

## v0.15.1
- Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

# v0.15.0
- **Deprecation:** The `JaegerHttpTracesUrl` field is now deprecated in all relevant models and will be removed after 9th April 2026. Use the new `JaegerHttpUrl` field instead.

# v0.14.0
- **Feature:** Add attributes `JaegerHttpTracesUrl`, `OtlpGrpcTracesUrl` and `OtlpHttpTracesUrl` to `InstanceSensitiveData` model

## v0.13.0
- **Feature:** Add support for HTTP checks and cert checks

## v0.12.0
- **Feature:** Add `MetricsEndpointUrl` field to `InstanceSensitiveData` model struct

## v0.11.1
- Introduce new struct `UpdateAlertConfigsPayloadRouteRoutesInner`

## v0.11.0
- **Feature:** Add new `GoogleChat` webhook

## v0.10.0
- **Feature:** Add new `CreateCredentialsPayload` model for creating credentials with optional description
- **Feature:** Add `Description` field to `Credentials` and `ServiceKeysList` models
- **Feature:** Update `CreateCredentials` API to accept payload with description
- **Improvement:** Improved documentation for some fields of the `UpdateMetricsStorageRetentionPayloard` model regarding downsampling behavior and validation rules

## v0.9.1
  - **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v0.9.0
- **Feature:** Add new API methods for logs configuration management: `GetLogsConfigs` and `UpdateLogsConfigs`
- **Feature:** Add new API methods for traces configuration management: `GetTracesConfigs` and `UpdateTracesConfigs`
- **Feature:** Add new models for configuration management: `LogsConfig`, `LogsConfigResponse`, `TraceConfig`, `TracesConfigResponse`, `UpdateLogsConfigsPayload`, and `UpdateTracesConfigsPayload`

## v0.8.0
- Add `required:"true"` tags to model structs

## v0.7.1 (2025-06-04)
- **Bugfix:** Adjusted `UnmarshalJSON` function to use enum types and added tests for enums
- **Feature:** Added `Priority` and `SendResolved` attributes
- **Deprecation:** `Match` and `MatchRe` attributes are deprecated.

## v0.7.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v0.6.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v0.5.2 (2025-05-09)
- **Feature:** Update user-agent header

## v0.5.1 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v0.5.0 (2025-04-16)
- **Feature:** Add new methods `ListLogsAlertgroups`, `CreateLogsAlertgroups`, `GetLogsAlertgroup`, `UpdateLogsAlertgroup`, `DeleteLogsAlertgroup`

## v0.4.0 (2025-03-27)
- **New:** Support for alert groups

## v0.3.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.3.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.2.1 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.2.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.1.0 (2024-08-21)

First release.

This module offers the same functionalities as `argus` (release [v0.11.0](https://github.com/stackitcloud/stackit-sdk-go/blob/main/services/argus/CHANGELOG.md#v0110-2024-05-23).
