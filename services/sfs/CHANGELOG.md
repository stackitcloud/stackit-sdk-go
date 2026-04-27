## v0.9.0
- **Deprecation:** deprecated field `Schedules` in model `SnapshotPolicy`
- **Feature:** new field `SnapshotSchedules` in model `SnapshotPolicy`
- **Feature:** new field `Interval` in model `SnapshotPolicySchedule`
- **Feature:** new model `SnapshotPolicySnapshotPolicySchedule`

## v0.8.0
- **Feature:** model `CreateResourcePoolPayload` now has an additional field `snapshotPolicyId`
- **Feature:** model `CreateResourcePoolSnapshotPayload` now has an additional field `snaplockRetentionHours`
- **Feature:** model `ResourcePool` now has an additional field `snapshotPolicy`
- **Feature:** model `ResourcePoolSnapshot` now has an additional field `snaplockExpiryTime`
- **Feature:** model `ResourcePoolSpace` now has an additional field `usedBySnapshotsGigabytes`
- **Feature:** model `UpdateResourcePoolPayload` now has an additional field `snapshotPolicyId`
- **Feature:** new models: `DisableLockResponse`, `EnableLockResponse`, `GetLockResponse`, `GetScheduleResponse`, `GetSnapshotPolicyResponse`, `ListSchedulesResponse`, `ListSnapshotPoliciesResponse`, `ResourcePoolSnapshotPolicy`, `Schedule`, `SnapshotPolicy`, `SnapshotPolicySchedule`, `UpdateResourcePoolSnapshotPayload`, `UpdateResourcePoolSnapshotResponse`
- **Feature:** new operations: `UpdateResourcePoolSnapshot`, `ListSchedules`, `GetSchedule`, `ListSnapshotPolicies`, `GetSnapshotPolicy`, `DisableLock`, `GetLock`, `EnableLock`,

## v0.7.0
- Minimal go version is now Go 1.25

## v0.6.3
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v0.6.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v0.6.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v0.6.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v0.5.0
- **Feature:** Introduction of multi API version support for the sfs SDK module. For more details please see the announcement on GitHub: https://github.com/stackitcloud/stackit-sdk-go/discussions/5062
- `v1api`: New package which can be used for communication with the SFS v1 API
- `v1betaapi`: New package which can be used for communication with the SFS v1 beta API
- **Deprecation:** The contents in the root of this SDK module including the `wait` package are marked as deprecated and will be removed after 2026-09-30. Switch to the new packages for the available API versions instead.
- **Dependencies:** Bump STACKIT SDK core module from `v0.21.1` to `v0.22.0`

## v0.4.0
- **Breaking change:** The `name` and `spaceHardLimitGigabytes` fields are now marked as required for `ShareExportPayload`, `SharePayload`.

## v0.3.0
- **Feature:** Switch from `v1beta` API version to `v1` version.
- **Breaking change:** Remove `ListSnapshotSchedules` method
- **Breaking change:** Remove field `SnapshotScheduleName` from `CreateResourcePoolPayload` and `UpdateResourcePoolPayload` model
- **Breaking change:** Remove field `SnapshotSchedule` from `CreateResourcePoolResponseResourcePool`, `GetResourcePoolResponseResourcePool`, `UpdateResourcePoolResponseResourcePoolGetStateRetType` and `ResourcePool` model

## v0.2.3
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v0.2.2
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v0.2.1
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v0.2.0
- **Breaking change:** Remove region configuration in `APIClient`

## v0.1.0
- **New**: STACKIT File Storage (SFS) service
