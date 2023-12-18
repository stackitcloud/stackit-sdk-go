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
- `logme`: [v0.6.0](services/logme/CHANGELOG.md#v060-2023-11-10)
  - Manage your STACKIT Logme resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/logme)
- `mariadb`: [v0.6.0](services/mariadb/CHANGELOG.md#v060-2023-11-10)
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
- `opensearch`: [v0.6.0](services/opensearch/CHANGELOG.md#v060-2023-11-10)
  - Manage your STACKIT OpenSearch resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/opensearch)
- `postgresflex`: [v0.7.0](services/postgresflex/CHANGELOG.md#v070-2023-11-10)
  - Manage your STACKIT PostgreSQL Flex resources: `Instance`, `Versions`, `Flavors`, `User`, `Storages`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `DeleteUserWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/postgresflex)
- `postgresql`: [v0.8.0](services/postgresql/CHANGELOG.md#v060-2023-11-17)
  - Manage your STACKIT PostgreSQL resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/postgresql)
- `rabbitmq`: [v0.6.0](services/rabbitmq/CHANGELOG.md#v060-2023-11-10)
  - Manage your STACKIT RabbitMQ resources: `Instance`, `Credentials`, `Offerings`
  - Waiters for async operations: `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, `DeleteInstanceWaitHandler`, `CreateCredentialsWaitHandler`, `DeleteCredentialsWaitHandler`
  - [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/rabbitmq)
- `redis`: [v0.6.0](services/redis/CHANGELOG.md#v060-2023-11-10)
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
