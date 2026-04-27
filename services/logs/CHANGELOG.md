## v0.8.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.25.0` to `v0.26.0`

## v0.8.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.1` to `v0.25.0`

## v0.8.0
- Minimal go version is now Go 1.25

## v0.7.3
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v0.7.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v0.7.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v0.7.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v0.6.0
- **Feature:** Introduction of multi API version support for the logs SDK module. For more details please see the announcement on GitHub: https://github.com/stackitcloud/stackit-sdk-go/discussions/5062
- `v1alphaapi`: New package which can be used for communication with the logs v1 alpha API
- `v1betaapi`: New package which can be used for communication with the logs v1 beta API
- `v1api`: New package which can be used for communication with the logs v1 API
- **Deprecation:** The contents in the root of this SDK module including the `wait` package are marked as deprecated and will be removed after 2026-09-30. Switch to the new packages for the available API versions instead.
- **Dependencies:** Bump STACKIT SDK core module from `v0.21.1` to `v0.22.0`

## v0.5.2
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v0.5.1
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v0.5.0
- **Feature:** switch from `v1beta` version to `v1` version of the API
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v0.4.0
- **Breaking Change:** The region is no longer specified within the client configuration. Instead, the region must be passed as a parameter to any region-specific request.

## v0.3.0
- **Feature:** Add new wait handlers for instance creation (`CreateLogsInstanceWaitHandler`), and instance deletion (`DeleteLogsInstanceWaitHandler`)

## v0.2.0
- **Feature:** Add support for access token GET endpoint

## v0.1.1
- Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v0.1.0
- **New:** API for logs service
