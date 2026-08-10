## v0.13.1
- `v1alphaapi`:
  - **Fix:** Response decoding now supports `*io.Reader` and `*[]byte` target types (previously only `string`, `*os.File`, and JSON were supported)
  - **Fix:** Escape regex in validate tags of model structs correctly
- `v1api`:
  - **Fix:** Response decoding now supports `*io.Reader` and `*[]byte` target types (previously only `string`, `*os.File`, and JSON were supported)
  - **Fix:** Escape regex in validate tags of model structs correctly
- `v1betaapi`:
  - **Fix:** Response decoding now supports `*io.Reader` and `*[]byte` target types (previously only `string`, `*os.File`, and JSON were supported)
  - **Fix:** Escape regex in validate tags of model structs correctly

## v0.13.0
- `v1betaapi`:
  - **Improvement:** The entire `v1betaapi` package is now marked as deprecated. Migrate to `v1api`
  - **Feature:** Add new models `Operator`, `Severity`, `Transformation`, `Type`, `Variable`, `Action`
  - **Breaking Change:** Removed models `BehaviourAction`, `ConditionOperatorType`, `ConditionTransformationsInner`, `ConditionVariableType`, `GetBehaviourAction`, `GetBehaviourSeverity`, `Mode1`, `MRSType`, `MRSType1`, `MRSType2`
  - **Breaking Change:** `PatchMRSRule.Mode` field type changed from `*Mode1` to `*Mode`
- `v1api`: New package which can be used to communicate with the albwaf v1 API

## v0.12.0
- `v1alphaapi`: Align package to latest API specification
- `v1betaapi`: Align package to latest API specification

## v0.11.0
- `v1alphaapi`: Align package to latest API specification
- `v1betaapi`: Align package to latest API specification

## v0.10.0
- **Deprecation:** The `v1alphaapi` is deprecated and will be removed in the future. Migrate to the `v1betaapi`.

## v0.9.0
- `v1alphaapi`: Align package to latest API specification
- `v1betaapi`: Align package to latest API specification

## v0.8.0
- `v1alphaapi`: Align package to latest API specification

## v0.7.0
- **Feature:** Introduce enums for various attributes

## v0.6.0
- `v1alphaapi`: Align package to latest API specification

## v0.5.0
- `v1alphaapi`: Align package to latest API specification

## v0.4.1
- **Dependencies:** Bump STACKIT SDK core module to `v0.26.0`

## v0.4.0
- Minimal go version is now Go 1.25

## v0.3.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v0.3.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v0.3.0
- **Feature:** Add new method `GetQuota` to get the quota for WAF resources in a project

## v0.2.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v0.2.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v0.1.0
- **New:** SDK module for albwaf service
- `v1alphaapi`: New package which can be used for communication with the albwaf v1 alpha API
