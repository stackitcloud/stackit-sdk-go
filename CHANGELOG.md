## Release (2024-01-09)

### Highlights

- `dns`: [v0.8.0](services/dns/CHANGELOG.md#v080-2024-01-09)
  - `Zone` has a new filed `Labels`, which is an array of labels (key-value pairs) associated to a zone
  - `ListZones` can be filtered by label keys or values
  - `CloneZonePayload` has a flag `AdjustRecords` to adjust the record set content of the cloned zone (replaces the dns name of the original zone with the new dns name of the cloned zone)
- `logme`: [v0.9.0](services/logme/CHANGELOG.md#v090-2024-01-09)
  - `PartialUpdateInstance` can be used to update the instance's name
  - `InstanceParameters` has a new setting `MaxDiskThreshold`
  - `ListMetricsResponse` has new fields regarding ephemeral disk
- `mariadb`: [v0.9.0](services/mariadb/CHANGELOG.md#v090-2024-01-09)
  - `PartialUpdateInstance` can be used to update the instance's name
  - `InstanceParameters` has a new setting `MaxDiskThreshold`
  - `ListMetricsResponse` has new fields regarding ephemeral disk
- `opensearch`: [v0.9.0](services/opensearch/CHANGELOG.md#v090-2024-01-09)
  - `PartialUpdateInstance` can be used to update the instance's name
  - `InstanceParameters` has a new setting `MaxDiskThreshold`
  - `ListMetricsResponse` has new fields regarding ephemeral disk
- `postgresql`: [v0.10.0](services/postgresql/CHANGELOG.md#v0100-2024-01-09)
  - `PartialUpdateInstance` can be used to update the instance's name
  - `InstanceParameters` has a new setting `MaxDiskThreshold`
  - `ListMetricsResponse` has new fields regarding ephemeral disk
- `rabbitmq`: [v0.9.0](services/rabbitmq/CHANGELOG.md#v090-2024-01-09)
  - `PartialUpdateInstance` can be used to update the instance's name
  - `InstanceParameters` has a new setting `MaxDiskThreshold`
  - `ListMetricsResponse` has new fields regarding ephemeral disk
- `redis`: [v0.9.0](services/redis/CHANGELOG.md#v090-2024-01-09)
  - `PartialUpdateInstance` can be used to update the instance's name
  - `InstanceParameters` has a new setting `MaxDiskThreshold`
  - `ListMetricsResponse` has new fields regarding ephemeral disk
- `ske`: [v0.9.0](services/ske/CHANGELOG.md#v090-2024-01-09)
  - Add details on credentials for old clusters
  - `ClusterStatus` now has a field `CredentialsRotation` with credentials' details

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
