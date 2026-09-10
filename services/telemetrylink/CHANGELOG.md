## v0.6.0
- `v1api`:
  - **Breaking Change:** `Enabled` field in `CreateOrUpdateFolderTelemetryLinkPayload`, `CreateOrUpdateOrganizationTelemetryLinkPayload` and `CreateOrUpdateProjectTelemetryLinkPayload` changed from required `bool` to optional `*bool` (server now defaults to `true` when omitted); `New*` constructors no longer take an `enabled` parameter
- `v1betaapi`:
  - **Breaking Change:** `Enabled` field in `CreateOrUpdateFolderTelemetryLinkPayload`, `CreateOrUpdateOrganizationTelemetryLinkPayload` and `CreateOrUpdateProjectTelemetryLinkPayload` changed from required `bool` to optional `*bool` (server now defaults to `true` when omitted); `New*` constructors no longer take an `enabled` parameter

## v0.5.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.26.0` to `v0.27.0`

## v0.5.1
- `v1api`:
  - **Fix:** Response decoding now supports `*io.Reader` and `*[]byte` target types (previously only `string`, `*os.File`, and JSON were supported)
  - **Fix:** Escape regex in validate tags of model structs correctly
- `v1betaapi`:
  - **Fix:** Response decoding now supports `*io.Reader` and `*[]byte` target types (previously only `string`, `*os.File`, and JSON were supported)
  - **Fix:** Escape regex in validate tags of model structs correctly

## v0.5.0
- **Improvement:** Add validation for `Description` field
- `v1api`: **Improvement:** Add validation for `Description` field

## v0.4.0
- `v1api`: **Feature:** Add support for `If-None-Match` header in `CreateOrUpdateFolderTelemetryLink`, `CreateOrUpdateOrganizationTelemetryLink` and `CreateOrUpdateProjectTelemetryLink`

## v0.3.0
- **New:** v1api API version for STACKIT Telemetry Link

## v0.2.0
- **Feature:** Introduce enums for various attributes

## v0.1.1
- **Improvement**: Use new `WaiterHandler` struct in the TelemetryLink WaitHandler

## v0.1.0
- **New**: API for STACKIT Telemetry Link
