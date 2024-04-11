## v0.12.0 (2024-04-09)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.11.1 (2024-04-09)

- Remove unused model data types..

## v0.11.0 (2024-03-28)

- **Feature**: Waiters for async operation `StartCredentialsRotationWaitHandler` and `CompleteCredentialsRotationWaitHandler`

## v0.10.1 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.10.0 (2024-02-06)

- **Feature:** New endpoints for credentials rotation.
  - `StartCredentialsRotation`
  - `CompleteCredentialsRotation`
  - `CreateKubeconfig`
  - These endpoints replace `GetCredentials` and `TriggerRotateCredentials`, which are **deprecated** and will not work for clusters with Kubernetes v1.27+, or if the new endpoints for kubeconfig or credentials rotation have already been used. For more information, see [How to rotate SKE credentials](https://docs.stackit.cloud/display/STACKIT/How+to+rotate+SKE+credentials#tabs-237293ce-f625-44ea-9d4f-689e31f596d6-1).

## v0.9.3 (2024-02-02)

- **Improvement**: Fix state name in `CredentialsRotationState.Phase`
- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.
- Increase timeout on `DeleteClusterWaitHandler` to 45 minutes

## v0.9.2 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.9.1 (2024-01-15)

- Add license and notice files

## v0.9.0 (2024-01-09)

- **Feature:** `ClusterStatus` now has a field `CredentialsRotation` with credentials' details
- **Improvement:** Add details on credentials for old clusters
- Dependency updates

## v0.8.1 (2023-12-22)

- Dependency updates

## v0.8.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

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
