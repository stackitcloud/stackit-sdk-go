## v0.5.0 (2024-08-13)

- **Feature:** Add support for nullable fields
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
