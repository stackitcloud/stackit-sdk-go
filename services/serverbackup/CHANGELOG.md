## v0.4.0 (2024-12-16)

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
