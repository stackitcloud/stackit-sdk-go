## v0.11.0 (2024-04-11)

- Support WithMiddleware configuration option in the client

## v0.10.1 (2024-04-09)

- Remove unused model data types..

## v0.10.0 (2024-04-08)

- **Deprecation:** Mark methods `EnableService` and `DisableService` as deprecated. Enablement and disablement of the load balancer functionality is now automaticly handled by the service.

## v0.9.3 (2024-03-20)

- **Improvement**: Improve error handling in Load Balancer creation waiter, fixing timeout being exceeded for `STATUS_PENDING` status with errors. If an error is found in the `Errors` field, the waiter now returns with error.

## v0.9.2 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.9.1 (2024-02-02)

- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.9.0 (2024-01-24)

- **Feature**: Introduces Server Name Indicator (SNI) support:
  - `Listener` has a new field `ServerNameIndicators`
- **Feature**: Introduces Layer 4 Session Persistance:
  - `TargetPool` has a new field `SessionPersistence`
  - `UpdateTargetPoolPayload` has a new field `SessionPersistence`
- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.8.3 (2024-01-15)

- Add license and notice files

## v0.8.2 (2024-01-09)

- Dependency updates

## v0.8.1 (2023-12-22)

- Dependency updates

## v0.8.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `DisableLoadBalancing` renamed to `DisableService`
  - `EnableLoadBalancing` renamed to `EnableService`
  - `GetProjectStatus` renamed to `GetServiceStatus`
- Changed structs:
  - `StatusResponse` renamed to `GetServiceStatusResponse`
- Changed waiters:
  - `EnableLoadBalancingWaitHandler` renamed to `EnableServiceWaitHandler`

## v0.7.0 (2023-12-06)

- Manage your STACKIT Load Balancer resources: `LoadBalancer`, `Credentials`
- Waiters for async operations: `CreateLoadBalancerWaitHandler`, `DeleteLoadBalancerWaitHandler`, `EnableLoadBalancingWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/loadbalancer)
