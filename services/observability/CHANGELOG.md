## v0.11.1
- Introduce new struct `UpdateAlertConfigsPayloadRouteRoutesInner`

## v0.11.0
- **Feature:** Add new `GoogleChat` webhook

## v0.10.0
- **Feature:** Add new `CreateCredentialsPayload` model for creating credentials with optional description
- **Feature:** Add `Description` field to `Credentials` and `ServiceKeysList` models
- **Feature:** Update `CreateCredentials` API to accept payload with description
- **Improvement:** Improved documentation for some fields of the `UpdateMetricsStorageRetentionPayloard` model regarding downsampling behavior and validation rules

## v0.9.1
  - **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v0.9.0
- **Feature:** Add new API methods for logs configuration management: `GetLogsConfigs` and `UpdateLogsConfigs`
- **Feature:** Add new API methods for traces configuration management: `GetTracesConfigs` and `UpdateTracesConfigs`
- **Feature:** Add new models for configuration management: `LogsConfig`, `LogsConfigResponse`, `TraceConfig`, `TracesConfigResponse`, `UpdateLogsConfigsPayload`, and `UpdateTracesConfigsPayload`

## v0.8.0
- Add `required:"true"` tags to model structs

## v0.7.1 (2025-06-04)
- **Bugfix:** Adjusted `UnmarshalJSON` function to use enum types and added tests for enums
- **Feature:** Added `Priority` and `SendResolved` attributes
- **Deprecation:** `Match` and `MatchRe` attributes are deprecated.

## v0.7.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v0.6.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v0.5.2 (2025-05-09)
- **Feature:** Update user-agent header

## v0.5.1 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v0.5.0 (2025-04-16)
- **Feature:** Add new methods `ListLogsAlertgroups`, `CreateLogsAlertgroups`, `GetLogsAlertgroup`, `UpdateLogsAlertgroup`, `DeleteLogsAlertgroup`

## v0.4.0 (2025-03-27)
- **New:** Support for alert groups

## v0.3.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.3.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.2.1 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.2.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.1.0 (2024-08-21)

First release.

This module offers the same functionalities as `argus` (release [v0.11.0](https://github.com/stackitcloud/stackit-sdk-go/blob/main/services/argus/CHANGELOG.md#v0110-2024-05-23).
