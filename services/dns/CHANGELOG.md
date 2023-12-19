## v0.7.0 (2023-12-19)

API methods and structs were renamed to have the same look and feel and according to user feedback.

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

## v0.6.0 (2023-11-10)

- Manage your STACKIT DNS resources: `Zones`, `RecordSet`
- Waiters for async operations: `CreateZoneWaitHandler`, `UpdateZoneWaitHandler`, `DeleteZoneWaitHandler`, `CreateRecordSetWaitHandler`, `UpdateRecordSetWaitHandler`, `DeleteRecordSetWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/dns)
