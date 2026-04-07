## v1.4.3
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v1.4.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v1.4.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v1.4.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v1.3.0
- **Feature:** Introduction of multi API version support for the serverupdate SDK module. For more details please see the announcement on GitHub: https://github.com/stackitcloud/stackit-sdk-go/discussions/5062
- `v1api`: New package which can be used for communication with the serverupdate v1 API
- `v2api`: New package which can be used for communication with the serverupdate v2 API
- **Deprecation:** The contents in the root of this SDK module are marked as deprecated and will be removed after 2026-09-30. Switch to the new packages for the available API versions instead.
- **Dependencies:** Bump STACKIT SDK core module from `v0.21.1` to `v0.22.0`

## v1.2.6
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v1.2.5
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v1.2.4
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v1.2.3
  - Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v1.2.2
  - Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

## v1.2.1
  - **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v1.2.0
- Add `required:"true"` tags to model structs

## v1.1.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v1.0.3 (2025-05-09)
- **Feature:** Update user-agent header

## v1.0.2 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v1.0.1 (2025-03-27)
- **Bugfix:** Removed ConfigureRegion() from API client

## v1.0.0 (2025-03-19)
- **Breaking Change:** The region is no longer specified within the client configuration. Instead, the region must be passed as a parameter to any region-specific request.

## v0.5.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.4.0 (2025-02-05)

- **Breaking Change**: Remove field `BackupProperties` from `CreateUpdatePayload` model
- **Fix**: Remove field `Id` from `CreateUpdateSchedulePayload` model

## v0.3.0 (2025-01-14)

- **Feature:** Add new method: `GetServiceResource` 

## v0.2.3 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.2.2 (2024-12-02)

- **Bugfix:** `Id` field of `Update` model is now of type `int64` (was `string`)

## v0.2.1 (2024-11-28)

- **Bugfix:** Fix `Accept` header types

## v0.2.0 (2024-11-26)

- **Feature:** Add support for managing `UpdatePolicy` resources

## v0.1.0 (2024-11-20)

- Manage your STACKIT Server Updates: `Update`, `UpdateSchedule`, `BackupProperties`
