## v1.3.0 (2025-05-28)
- **Breaking change:** Added missing typed enum constants

## v1.2.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v1.1.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v1.0.3 (2025-05-09)
- **Feature:** Update user-agent header

## v1.0.2 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v1.0.1 (2025-03-19)
- **Bugfix:** Corrected regional setup of client

## v1.0.0 (2025-03-14)
- **Breaking Change:** The region is no longer specified within the client configuration. Instead, the region must be passed as a parameter to any region-specific request.
- **Breaking Change:** Remove deprecated API methods `DisableService` and `EnableService`. They are no longer required because the service automatically enables and disable.
- **Breaking Change:** Remove WaitHandler `EnableServiceWaitHandler` for `EnableService`.

## v0.18.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.17.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.16.0 (2024-10-11)

- **Feature:** Add pagination to `ListLoadBalancers` with the new fields `pageSize` and `pageId` on `ApiListLoadBalancersRequest` and the field `NextPageId` in `ListLoadBalancersResponse`

## v0.15.0 (2024-08-08)

- **Feature:** New API method `ListPlans` to list the available service plans

## v0.14.0 (2024-07-10)

- **Improvement:** Improve default error messages.
- **Bugfix:** Fix marking of deprecated methods. Potential breaking change for users with linters that treat deprecations as errors.

## v0.13.0 (2024-06-12)

- **Feature:** `LoadBalancer`, `CreateLoadBalancerPayload` and `UpdateLoadBalancerPayload` have a new field `PlanId`

## v0.12.0 (2024-04-12)

- **Feature:** Set `config.ContextHTTPRequest` in `Execute` methods
- **Feature:** New API method `GetQuota` to get the maximum number of load balancing servers allowed for a project
- **Feature:** New API method `UpdateCredentials` to update the credentials for observability in a project

## v0.11.0 (2024-04-11)

- **Feature:** Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.10.1 (2024-04-09)

- Remove unused model data types.

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
