## v0.8.0 (2023-12-19)

API methods and structs were renamed to have the same look and feel and according to user feedback.

- Changed methods:
  - `DisableLoadBalancing` renamed to `DisableService`
  - `EnableLoadBalancing` renamed to `EnableService`
  - `GetProjectStatus` renamed to `GetServiceStatus`
- Changed structs:
  - `StatusResponse` renamed to `GetServiceStatusResponse`

## v0.7.0 (2023-12-06)

- Manage your STACKIT Load Balancer resources: `LoadBalancer`, `Credentials`
- Waiters for async operations: `CreateLoadBalancerWaitHandler`, `DeleteLoadBalancerWaitHandler`, `EnableLoadBalancingWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/loadbalancer)
