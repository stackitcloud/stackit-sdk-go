## v0.8.0 (2023-12-20)

API methods and structs were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `CreateProject` renamed to `EnableService`
  - `DeleteProject` renamed to `DisableService`
  - `GetClusters` renamed to `ListClusters`
  - `GetOptions` renamed to `ListProviderOptions`
  - `GetProject` renamed to `GetServiceStatus`
- Changed structs:
  - `ClusterResponse` renamed to `Cluster`
  - `ClustersResponse` renamed to `ListClustersResponse`
  - `CredentialsResponse` renamed to `Credentials`

## v0.7.0 (2023-12-05)

- Manage your STACKIT Kubernetes Engine resources: `Project`, `Cluster`, `Credentials`, `Options`
- Waiters for async operations: `CreateOrUpdateClusterWaitHandler`, `DeleteClusterWaitHandler`, `CreateProjectWaitHandler`, `DeleteProjectWaitHandler`, `RotateCredentialsWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/ske)
