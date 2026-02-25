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
