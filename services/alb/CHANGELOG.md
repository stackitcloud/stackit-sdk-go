## v0.9.0
- **Feature:** Add new fields `UsedCredentials` and `UsedLoadBalancers` in `GetQuotaResponse` model

## v0.8.0
- **Feature:** Switch from `v2beta` API version to `v2` version.
- **Feature:** `MaxCredentials` field added to `GetQuotaResponse`
- **Breaking change:** added `version` to LoadBalancer constructor
- **Breaking change:** renamed `exact` to `exactMatch` in Path model
- **Breaking change:** removed `pathPrefix` from Rule model

## v0.7.3
- Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v0.7.2
- Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

## v0.7.1
- **Docs** Update description of field `WafConfigName` in `Listener` model

## v0.7.0
- Add field `Labels` (type `*map[string]string`) to structs `LoadBalancer`, `CreateLoadBalancerPayload`, `UpdateLoadBalancerPayload`
- Add field `WafConfigName` (type `*string`) to `Listener` struct

## v0.6.1
- **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v0.6.0
- **Feature:** Add new `LoadBalancerSecurityGroup` field to `CreateLoadBalancerPayload`, `LoadBalancer`, and `UpdateLoadBalancerPayload` models

## v0.5.0 (2025-06-12)
- **Feature:** Add new fields `DisableTargetSecurityGroupAssignment` and `TargetSecurityGroup` in `LoadBalancer` Model

## v0.4.1 (2025-06-04)
- **Bugfix:** Adjusted `UnmarshalJSON` function to use enum types and added tests for enums

## v0.4.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v0.3.1 (2025-05-15)
- **Feature:** New field `Path` for `Rule`

## v0.3.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v0.2.3 (2025-05-09)
- **Feature:** Update user-agent header

## v0.2.2 (2025-05-02)
- **Feature:** Switch to beta2 API

## v0.2.1 (2025-03-27)
- **Bugfix:** Removed ConfigureRegion() from API client

## v0.2.0 (2025-03-20)
- **Enhancement:** Provider waiter for loadbalancer api

## v0.1.0 (2025-03-19)
- **New:** API for application load balancer
