## v0.2.0
- **Dependencies:** Bump STACKIT SDK core module from `v0.27.0` to `v0.27.1`
- `v1alphaapi`:
  - **Breaking Change:** Field `Network` in `Application` and `CreateApplicationPayload` models is now optional and changed from `Network` to `*Network` (constructors `NewApplication` and `NewCreateApplicationPayload` no longer take the `network` argument)
  - **Breaking Change:** Field `PublicIngress` in `Network` model is now optional and changed from `bool` to `*bool` (constructor `NewNetwork` no longer takes the `publicIngress` argument)
  - **Feature:** New model struct `EnvironmentStatus`
  - **Feature:** Add `Status` field to `Environment` and `CreateEnvironmentPayload` models
  - **Feature:** Add `InternalUrl` field to `ApplicationSummary` and `RuntimeStatus` models
  - **Feature:** Add `InternalPort` field to `Network` model
  - **Bugfix:** Handle application failed status in create and update waiters.
  - **Bugfix:** Handle application none status in create and updated waiters.

## v0.1.0
- **New:** SDK module for STACKIT Container Applications (SCA) service.
- `v1alphaapi`: New package which can be used for communication with the sca v1 alpha API
