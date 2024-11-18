## v0.11.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.10.0 (2024-05-23)

- **Feature**: New method `CloneZone` to clone an existing zone with all record sets to a new zone with a different name
- **Feature**: New methods `CreateLabel`, `DeleteLabel` and `ListLabels` to manage labels for a zone
- **Feature**: New methods `CreateMoveCode`, `DeleteMoveCode` and `ValidateMoveCode` to manage move codes to move a zone to another project
- **Feature**: New method `MoveZone` to move a zone to another project
- **Feature**: New methods `ExportRecordSets` and `ImportRecordSets`
- **Feature**: New methods `RestoreZone` and `RestoreRecordSet` to restore inactive zones and record-sets, respectively
- **Feature**: New method `RetrieveZone` to queue a secondary zone for a zone transfer request

## v0.9.1 (2024-04-24)

- Remove unused data types.

## v0.9.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.8.4 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.8.3 (2024-02-02)

- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.8.2 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.8.1 (2024-01-15)

- Add license and notice files

## v0.8.0 (2024-01-09)

- **Feature:** `Zone` has a new filed `Labels`, which is an array of labels (key-value pairs) associated to a zone
- **Feature:** `ListZones` can be filtered by label keys or values
- **Feature:** `CloneZonePayload` has a flag `AdjustRecords` to adjust the record set content of the cloned zone (replaces the dns name of the original zone with the new dns name of the cloned zone)
- Dependency updates

## v0.7.1 (2023-12-22)

- Dependency updates

## v0.7.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `GetRecordSets` renamed to `ListRecordSets`
  - `GetZones` renamed to `ListZones`
  - `UpdateRecord` renamed to `PartialUpdateRecord`
  - `UpdateRecordSet` renamed to `PartialUpdateRecordSet`
  - `UpdateZone` renamed to `PartialUpdateZone`
- Changed structs:
  - `CloneZoneRequest` renamed to `CloneZonePayload`
  - `MoveCodeValidationRequest` renamed to `ValidateMoveCodePayload`
  - `MoveZoneRequest` renamed to `MoveZonePayload`
  - `RecordSetsResponse` renamed to `ListRecordSetsResponse`
  - `UpdateRecordPayload` renamed to `PartialUpdateRecordPayload`
  - `UpdateRecordSetPayload` renamed to `PartialUpdateRecordSetPayload`
  - `UpdateZonePayload` renamed to `PartialUpdateZonePayload`
  - `ZonesResponse` renamed to `ListZonesResponse`
- Changed waiters:
  - `UpdateZoneWaitHandler` renamed to `PartialUpdateZoneWaitHandler`
  - `UpdateRecordSetWaitHandler` renamed to `PartialUpdateRecordSetWaitHandler`

## v0.6.0 (2023-11-10)

- Manage your STACKIT DNS resources: `Zones`, `RecordSet`
- Waiters for async operations: `CreateZoneWaitHandler`, `UpdateZoneWaitHandler`, `DeleteZoneWaitHandler`, `CreateRecordSetWaitHandler`, `UpdateRecordSetWaitHandler`, `DeleteRecordSetWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/dns)
