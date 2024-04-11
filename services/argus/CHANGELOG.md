## v0.10.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client

## v0.9.5 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.9.4 (2024-02-02)

- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.9.3 (2024-01-26)

- Increase default timeouts of wait operations on Argus instances

## v0.9.2 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.9.1 (2024-01-15)

- Add license and notice files

## v0.9.0 (2024-01-12)

- **Feature**: Add `SampleLimit` to `UpdateScrapeConfigPayload`.

  This field allows you to update the per-scrape limit on number of scraped samples that will be accepted. If more than this number of samples are present after metric relabeling the entire scrape will be treated as failed. The total limit depends on the service plan target (limits \* samples).

## v0.8.2 (2024-01-09)

- Dependency updates

## v0.8.1 (2023-12-22)

- Dependency updates

## v0.8.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
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
- Changed structs:
  - `AclResponse` renamed to `ListACLResponse`
  - `AlertGroupJson` renamed to `AlertGroup`
  - `AlertRuleRecordJson` renamed to `AlertRuleRecord`
  - `ApiUserProjectCreated` renamed to `CreateCredentialsResponse`
  - `BackupScheduleModelJson` renamed to `BackupSchedule`
  - `CreateInstanceAlertConfigReceiverPayload` renamed to `CreateAlertConfigReceiverPayload`
  - `CreateInstanceAlertConfigRoutePayload` renamed to `CreateAlertConfigRoutePayload`
  - `Credential` renamed to `Credentials`
  - `CredentialsListResponse` renamed to `ListCredentialsResponse`
  - `CredentialsRemoteWriteResponse` renamed to `CredentialsRemoteWriteConfig`
  - `GetAlert` renamed to `GetAlertConfigsResponse`
  - `InstanceResponse` renamed to `GetInstanceResponse`
  - `MysqlCheckChildResponse` renamed to `MySQLCheckChildResponse`
  - `PlanModel` renamed to `Plan`
  - `ProjectInstanceFullMany` renamed to `ListInstancesResponse`
  - `ProjectInstancesUpdateResponse` renamed to `InstanceResponse`
  - `PutAlert` renamed to `UpdateAlertConfigsResponse`
  - `ReceiversResponse` renamed to `AlertConfigReceiversResponse`
  - `ReceiversResponseSerializerSingle` renamed to `Receiver`
  - `RouteResponse` renamed to `AlertConfigRouteResponse`
  - `RouteSerializer2` renamed to `RouteSerializer`
  - `ScrapeConfigResponse` renamed to `GetScrapeConfigResponse`
  - `ScrapeConfigsResponse` renamed to `ListScrapeConfigsResponse` (when output of ListScrapeConfigs) and `DeleteScrapeConfigResponse` (when output of DeleteScrapeConfig)
  - `ServiceKeysResponse` renamed to `GetCredentialsResponse`
  - `UpdateCredentialRemoteWriteConfigPayload` renamed to `UpdateCredentialsRemoteWriteConfigPayload`
  - `UpdateInstanceAclPayload` renamed to `UpdateACLPayload`
  - `UpdateInstanceAlertConfigReceiverPayload` renamed to `UpdateAlertConfigReceiverPayload`
  - `UpdateInstanceAlertConfigRouteReceiverPayload` renamed to `UpdateAlertConfigsPayloadRoute`
  - `UpdateInstanceAlertConfigsPayload` renamed to `UpdateAlertConfigsPayload`
  - `UpdateInstanceAlertConfigsPayloadGlobal` renamed to `UpdateAlertConfigsPayloadGlobal`
  - `UpdateInstanceAlertConfigsPayloadInhibitRules` renamed to `UpdateAlertConfigsPayloadInhibitRules`
  - `UpdateInstanceAlertConfigsPayloadReceiversInner` renamed to `UpdateAlertConfigsPayloadReceiversInner`
  - `UpdateInstanceAlertConfigsPayloadReceiversInnerEmailConfigsInner` renamed to `UpdateAlertConfigsPayloadReceiversInnerEmailConfigsInner`
  - `UpdateInstanceAlertConfigsPayloadReceiversInnerOpsgenieConfigsInner` renamed to `UpdateAlertConfigsPayloadReceiversInnerOpsgenieConfigsInner`
  - `UpdateInstanceAlertConfigsPayloadReceiversInnerWebHookConfigsInner` renamed to `UpdateAlertConfigsPayloadReceiversInnerWebHookConfigsInner`
  - `UpdateInstanceAlertConfigsPayloadRoute` renamed to `UpdateAlertConfigRoutePayload`
  - `UpdateInstanceAlertConfigsPayloadRouteRoutesInner` renamed to `UpdateAlertConfigsPayloadRouteRoutesInner`
  - `UpdateInstanceGrafanaConfigsPayload` renamed to `UpdateGrafanaConfigsPayload`
  - `UpdateInstanceGrafanaConfigsPayloadGenericOauth` renamed to `UpdateGrafanaConfigsPayloadGenericOauth`

## v0.7.0 (2023-11-10)

- Manage your STACKIT Argus resources: `Instance`, `Credentials`, `ScrapeConfig`, `Acl`, `Alertconfig`, `GrafanaConfig`
- Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateScrapeConfigWaitHandler`, `DeleteScrapeConfigWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/argus)
