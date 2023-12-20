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
