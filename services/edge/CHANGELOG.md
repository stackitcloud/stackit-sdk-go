## v0.9.2
- **Dependencies:** Bump STACKIT SDK core module to `v0.26.0`

## v0.9.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.1` to `v0.25.0`

## v0.9.0
- Minimal go version is now Go 1.25

## v0.8.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v0.8.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v0.8.0
- `v1beta1api`: Removal of deprecated API client method `ListPlansGlobal`
- Deprecated SDK layer in root of the module: Removal of deprecated API client method `ListPlansGlobal`

## v0.7.0
- `v1beta1api`: New field `MinEdgeHosts` in `Plan` model struct
- Deprecated SDK layer in root of the module: New field `MinEdgeHosts` in `Plan` model struct

## v0.6.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v0.6.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v0.5.0
- **Feature:** Introduction of multi API version support for the edge SDK module. For more details please see the announcement on GitHub: https://github.com/stackitcloud/stackit-sdk-go/discussions/5062
- `v1beta1api`: New package which can be used for communication with the edge v1 beta1 API
- **Deprecation:** The contents in the root of this SDK module including the `wait` package are marked as deprecated and will be removed after 2026-09-30. Switch to the new packages for the available API versions instead.
- **Dependencies:** Bump STACKIT SDK core module from `v0.21.1` to `v0.22.0`

## v0.4.3
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v0.4.2
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v0.4.1
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v0.4.0
- **Deprecation:** Deprecated API method `ListPlansGlobal`

## v0.3.0
- **Breaking change:** Rename methods: `PostInstances` to `CreateInstance` and `GetInstances` to `ListInstances`

## v0.2.0
- **Feature:** Add waiter methods for the API

## v0.1.0
- **New:** STACKIT Edge Cloud service
