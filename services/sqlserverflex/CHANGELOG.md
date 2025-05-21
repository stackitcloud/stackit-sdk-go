## v1.2.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v1.1.0 (2025-05-14)
- **Features**: Added new method `ListMetrics`

## v1.0.3 (2025-05-09)
- **Feature:** Update user-agent header

## v1.0.2 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v1.0.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v1.0.0 (2025-02-25)
- **Breaking Change:** The region is no longer specified within the client configuration. Instead, the region must be passed as a parameter to any region-specific request.
- **Feature:** Add method to delete all instances for a project: `TerminateProject`

## v0.10.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.9.0 (2025-01-20)

- **Breaking change**: Delete endpoint made private.

## v0.8.1 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.8.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.7.0 (2024-09-25)

- **Breaking change**: Field `Item` in `ResetUserResponse` is now of type `SingleUser` (previously was `User`)
- **Feature:** `DefaultDatabase` is no longer required in `CreateUserPayload`

## v0.6.0 (2024-09-19)

- **Breaking change**: Field `ListBackupsResponse` has a new field `BackupListBackupsResponseGrouped`, replacing the removed `Count` and `Items` fields

## v0.5.0 (2024-08-16)

- **Breaking change**:
  - Fields in `GetBackupResponse` are not nested in an `Item` field (with type `Backup`) anymore
  - `GetBackupResponse` have these new fields: `EndTime`, `Error`, `Id`, `Labels`, `Name`, `Options`, `Size`, `StartTime`

## v0.4.0 (2024-07-19)

- **Feature:** New field for `DatabaseOptions` and `SingleDatabaseOptions`: `CollationName`
- **Breaking change**: Fields removed from `DatabaseOptions` and `SingleDatabaseOptions`: `IsEncrypted`, `RecoveryModel`, `UserAccess`
- **Breaking change**: Fields removed from `SingleDatabase`: `Collation`, `CreateDate`

## v0.3.0 (2024-07-09)

- **Breaking changes:**
  - `Database` renamed to `DefaultDatabase`, in `CreateUserPayload`
  - Type of `Roles` changed from `[]Role` to `[]string`, in `CreateUserPayload`
  - `User` renamed to `SingleUser`, in `CreateUserResponse`
  - `OwnerName` renamed to `Owner`, in `DatabaseOptions`
  - Fields in `GetDatabaseResponse` are now nested in a `Database` field (with type `SingleDatabase`)
  - `GetDatabaseResponseOptions` renamed to `SingleDatabaseOptions` (and `OwnerName` renamed to `Owner`)

## v0.2.0 (2024-05-24)

- **Feature:** Waiters for async operations `CreateInstanceWaitHandler`, `UpdateInstanceWaitHandler`, and `DeleteInstanceWaitHandler`

## v0.1.0 (2024-05-22)

- Manage your STACKIT SQL Server Flex resources: `Instance`, `Flavors`, `Users`, `Databases`, `Backups`
