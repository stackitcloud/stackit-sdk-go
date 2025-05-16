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
