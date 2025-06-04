## v1.2.1 (2025-06-04)
- **Bugfix:** Adjusted `UnmarshalJSON` function to use enum types and added tests for enums

## v1.2.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v1.1.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v1.0.3 (2025-05-09)
- **Feature:** Update user-agent header

## v1.0.2 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v1.0.1 (2025-03-27)
- **Bugfix:** Removed ConfigureRegion() from API client

## v1.0.0 (2025-03-19)
- **Breaking Change:** The region is no longer specified within the client configuration. 

## v0.6.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.5.0 (2025-01-14)

- **Feature:** Add new method: `GetServiceResource` 

## v0.4.0 (2024-12-17)

- **Feature:** Add support for managing `BackupPolicy` resources
- **Bugfix:** Correctly handle nullable attributes in model types

## v0.3.0 (2024-11-20)

- **Breaking changes**: 
  - `ListBackups200Response` type renamed to `GetBackupsListResponse`
  - `GetBackupsListResponse` has been removed
  - `ErrorResponse` has been removed
  - Added new method `GetBackupSchedulesResponse`
  - Added new type `EnableServiceResourcePayload`

## v0.2.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.1.0 (2024-05-23)

- Manage your STACKIT Server Backups: `Backup`, `BackupSchedule`, `VolumeBackup`
