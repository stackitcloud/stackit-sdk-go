## v0.9.0
- **Feature:** Added `_UNKNOWN_DEFAULT_OPEN_API` fallback value to all enums to handle unknown API values gracefully.

## v0.8.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.25.0` to `v0.26.0`

## v0.8.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.1` to `v0.25.0`

## v0.8.0
- Minimal go version is now Go 1.25

## v0.7.3
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v0.7.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v0.7.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v0.7.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v0.6.0
- Deprecated SDK layer in root of the module: 
  - **Breaking Change**: Switch from regional to global API server URL (`config.WithRegion(...)` must not be used anymore).
- `v1betaapi`: 
  - **Breaking Change**: Switch from regional to global API server URL (`config.WithRegion(...)` must not be used anymore).

## v0.5.0
- **Feature:** Introduction of multi API version support for the intake SDK module. For more details please see the announcement on GitHub: https://github.com/stackitcloud/stackit-sdk-go/discussions/5062
- `v1betaapi`: New package which can be used for communication with the intake v1 beta API
- **Deprecation:** The contents in the root of this SDK module including the `wait` package are marked as deprecated and will be removed after 2026-09-30. Switch to the new packages for the available API versions instead.
- **Dependencies:** Bump STACKIT SDK core module from `v0.21.1` to `v0.22.0`

## v0.4.4
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v0.4.3
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v0.4.2
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v0.4.1
  - Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v0.4.0
- **Feature:** Add new enum type `PartitioningUpdateType`
- **Feature:** Add fields `PartitionBy` and `Partitioning` to `IntakeCatalogPatch` model struct

## v0.3.1
  - Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

## v0.3.0
- **Feature:** Add wait handlers for `Intake`, `IntakeRunner`, and `IntakeUser` resources.
- **Improvement:** Add usage examples for the `intake` service.

## v0.2.0
- **Feature:** Add response `IntakeRunnerResponse` to `UpdateIntakeRunnerExecute` request
- **Feature:** Add response `IntakeUserResponse` to `UpdateIntakeUserExecute` request

## v0.1.2
- **Feature:** Add new field `Partitioning` to `IntakeCatalog` model

## v0.1.1
- Mark fields `MaxMessageSizeKiB` and `MaxMessagesPerHour` as optional instead of required in `UpdateIntakeRunnerPayload` model struct 

## v0.1.0
- **New**: STACKIT Intake module can be used to manage the STACKIT Intake. Manage your `IntakeRunners`, `Intakes` and `IntakeUsers`
