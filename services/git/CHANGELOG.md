## v0.11.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v0.11.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v0.11.0
- **Feature:** Introduction of multi API version support for the git SDK module. For more details please see the announcement on GitHub: https://github.com/stackitcloud/stackit-sdk-go/discussions/5062
- `v1betaapi`: New package which should be used for communication with the STACKIT git API in the future
- **Deprecation:** The contents in the root of this SDK module including the `wait` package are marked as deprecated and will be removed after 2026-09-30. Switch to the new `v1api` package instead.
- **Dependencies:** Bump STACKIT SDK core module from `v0.21.1` to `v0.23.0`
- **Feature:** New API client methods for managing authentication: `CreateAuthentication`, `GetAuthentication`,  `ListAuthentication`, `PatchAuthentication`, `DeleteAuthentication`
- **Feature:** New API client methods for managing runners: `CreateRunner`,  `GetRunner`, `DeleteRunner`, `ListRunnerRuntimes`
- **Feature:** New model structs for managing authentication: `Authentication`, `AuthenticationList`, `CreateAuthenticationPayload`, `PatchAuthenticationPayload`
- **Feature:** New model structs for managing runners: `CreateRunnerPayload`, `Runner`, `RunnerRuntime`, `RunnerRuntimeList`
- **Feature:** New error model structs: `AlreadyExistsError`, `BadErrorResponse`, `NotFoundErrorResponse`, `UnauthorizedErrorResponse`, `ConflictErrorResponse`
- **Feature:** New field `FeatureToggle `in Instance model struct
- **Feature:** New fields `Acl` and `FeatureToggle` in struct `PatchInstancePayload`
- **Breaking change:** Removal of type `InstanceFlavor`
- **Breaking change:** Removal of structs: `ListRunnerLabels`, `RunnerLabel`, `UnauthorizedResponse`

## v0.10.3
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v0.10.2
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v0.10.1
- Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v0.10.0
- **Breaking Change:** Replace `PatchOperation` by `PatchInstancePayload` in `ApiPatchInstanceRequest` struct
- **Breaking Change:** Replace `CreateInstancePayloadFlavor` enums by `InstanceFlavor`

## v0.9.1
- Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

## v0.9.0
- **Feature:** Add support for list runner labels operation
    - new API client methods `ListRunnerLabels` and `ListRunnerLabelsExecute`
    - new model struct `RunnerLabel`

## v0.8.0
- **Feature:** Add support for the instance patch operation

## v0.7.1
- **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v0.7.0
- **Feature:** Add `Sku` field to `Flavors` model
- **Feature:** Add enums `CreateInstancePayloadFlavor` for `Flavor` field in `CreateInstancePayload`

## v0.6.0
- **Feature:** Add support for `Flavors` for STACKIT git instance
- **Feature:** Add support for `Acl` for STACKIT git instance
- `projectId` and `instanceId` strings must have a length of 36 characters now (previously was only limited to a maximum of 36 characters)
- Add `required:"true"` tags to model structs

## v0.5.1 (2025-06-04)
- **Bugfix:** Adjusted `UnmarshalJSON` function to use enum types and added tests for enums

## v0.5.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v0.4.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v0.3.3 (2025-05-09)
- **Feature:** Update user-agent header

## v0.3.2 (2025-05-02)
- **Bugfix**: Spelling corrections in documentation

## v0.3.1 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v0.3.0 (2025-04-22)
- **Features**: Add waiters to manage the STACKIT Git

## v0.2.0 (2025-04-16)
- **Features**: Add new methods to manage the STACKIT Git: `CreateInstance`, `DeleteInstance`, `GetInstance`

## v0.1.0 (2025-04-14)
- **New**: STACKIT Git module can be used to manage STACKIT Git
