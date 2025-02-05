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
