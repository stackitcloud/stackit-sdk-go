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
