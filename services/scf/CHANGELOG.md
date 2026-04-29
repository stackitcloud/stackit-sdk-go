## v0.8.0
- **Feature:** Added `_UNKNOWN_DEFAULT_OPEN_API` fallback value to all enums to handle unknown API values gracefully.

## v0.7.2
- **Dependencies:** Bump STACKIT SDK core module to `v0.26.0`

## v0.7.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.1` to `v0.25.0`

## v0.7.0
- Minimal go version is now Go 1.25

## v0.6.3
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v0.6.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v0.6.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v0.6.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v0.5.0
- **Feature:** Introduction of multi API version support for the scf SDK module. For more details please see the announcement on GitHub: https://github.com/stackitcloud/stackit-sdk-go/discussions/5062
- `v1api`: New package which should be used for communication with the STACKIT scf API in the future
- **Deprecation:** The contents in the root of this SDK module including the `wait` package are marked as deprecated and will be removed after 2026-09-30. Switch to the new `v1api` package instead.
- **Dependencies:** Bump STACKIT SDK core module from `v0.21.1` to `v0.22.0`

## v0.4.3
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v0.4.2
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v0.4.1
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v0.4.0
- **Feature:** Add new model structs `SpaceWithIsolationSegment` and `SpaceWithIsolationSegmentAllOf`

## v0.3.0
- **Feature:** Add new model `IsolationSegment` and `IsolationSegmentsList`

## v0.2.3
  - Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v0.2.2
  - Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

## v0.2.1
- **Feature:** Add waiter for deletion of organization

## v0.2.0
- **Feature:** Add field `OrgId` in model `OrgManager`
- **Feature:** Add new model `OrganizationCreateBffResponse` and `SpaceCreatedBffResponse`

## v0.1.0
- **New:** STACKIT Cloud Foundry module
