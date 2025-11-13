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
