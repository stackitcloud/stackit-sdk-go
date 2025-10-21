## v1.0.0
- **Breaking Change:** The region is no longer specified within the client configuration. Instead, the region must be passed as a parameter to any region-specific request.
- **Feature:** Add new methods to manage routing tables: `AddRoutingTableToArea`, `DeleteRoutingTableFromArea`, `GetRoutingTableOfArea`, `ListRoutingTablesOfArea`, `UpdateRoutingTableOfArea`
- **Feature:** Add new methods to manage routes in routing tables: `AddRoutesToRoutingTable`, `DeleteRouteFromRoutingTable`, `GetRouteOfRoutingTable`, `ListRoutesOfRoutingTable`, `UpdateRouteOfRoutingTable`
- **Breaking Change:** Add new method to manage network area regions: `CreateNetworkAreaRegion`, `DeleteNetworkAreaRegion`, `GetNetworkAreaRegion`, `ListNetworkAreaRegions`, `UpdateNetworkAreaRegion`
- **Feature:** Add new wait handler for network area region: `CreateNetworkAreaRegionWaitHandler` and `DeleteRegionalNetworkAreaConfigurationWaitHandler`
- **Breaking Change:** Wait handler which relates to region-specific services, got an additional param for the region: `CreateNetworkWaitHandler`, `UpdateNetworkWaitHandler`, `DeleteNetworkWaitHandler`, `CreateVolumeWaitHandler`, `DeleteVolumeWaitHandler`, `CreateServerWaitHandler`, `ResizeServerWaitHandler`, `DeleteServerWaitHandler`, `StartServerWaitHandler`, `StopServerWaitHandler`, `DeallocateServerWaitHandler`, `RescueServerWaitHandler`, `UnrescueServerWaitHandler`, `ProjectRequestWaitHandler`, `AddVolumeToServerWaitHandler`, `RemoveVolumeFromServerWaitHandler`, `UploadImageWaitHandler`, `DeleteImageWaitHandler`, `CreateBackupWaitHandler`, `DeleteBackupWaitHandler`, `RestoreBackupWaitHandler`, `CreateSnapshotWaitHandler`, `DeleteSnapshotWaitHandler`
- **Breaking Change:** `Network` model has changed: 
  - `NetworkId` has been renamed to `Id`
  - `Gateway`, `Nameservers`, `Prefixes` and `PublicIp` has been moved to new model `NetworkIPv4`, and can be accessed in the new property `IPv4`
  - Properties `Gatewayv6`, `Nameserversv6` and `Prefixesv6` moved to new model `NetworkIPv6`, and can be accessed in the new property `IPv6`
- **Breaking Change:** `CreateServerPayload` model has changed:
  - Model `CreateServerPayloadBootVolume` of `BootVolume` property changed to `ServerBootVolume`
  - Property `Networking` in `CreateServerPayload` is required now
- **Deprecated:** Deprecated wait handler and will be removed after April 2026: `CreateNetworkAreaWaitHandler`, `UpdateNetworkAreaWaitHandler` and `DeleteNetworkAreaWaitHandler`

## v0.31.0
- Add `CreatedAt` and `UpdatedAt` fields to `BaseSecurityGroupRule` struct
- Add `Description` field to `CreateNicPayload`, `NIC`, `UpdateNicPayload` structs
- Add `Agent` field to `Server`, `CreateServerPayload` structs
- Add `CreatedAt`, `UpdatedAt` and `Protocol` fields to `SecurityGroupRule`, `CreateSecurityGroupRulePayload` structs
- New struct: `ServerAgent`

## v0.30.0
- **Feature:** Add waiter to wait until the preconditions for network area deletion are met: `ReadyForNetworkAreaDeletionWaitHandler`

## v0.29.2
- Increase Timeouts for network area and network wait handlers to 30 minutes

## v0.29.1
- **Bugfix:** Parsing oneOf with enum and string value

## v0.29.0
- **Feature:** Add new method to get project details `GetProjectDetails`

## v0.28.0
- **Docs:** Improved descriptions of properties in structs with their possible values
- **Feature:** Add `Agent` field to `CreateImagePayload`, `UpdateImagePayload` and `Image` model
- **Feature:** Add `Encrypted` and `EncryptionsParameters` fields to `CreateVolumePayload` model

## v0.27.1
- **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v0.27.0
- **Feature:** Add `Dhcp` field in `CreateNetworkPayload`, `Network` and `PartialUpdateNetworkPayload` models

## v0.26.0
- **Feature:** Add `Metadata` field to `Server`, `CreateServerPayload`, and `UpdateServerPayload` models
- **Feature:** Increase maximum length validation for `machineType` and `volumePerformanceClass` from 63 to 127 characters

## v0.25.0
- Add `required:"true"` tags to model structs

## v0.24.1 (2025-06-12)
- Update descriptions of model fields

## v0.24.0 (2025-06-05)
- **Feature:** Add waiters for async operations: `CreateBackupWaitHandler`, `DeleteBackupWaitHandler`, `RestoreBackupWaitHandler`
- **Feature:** Add Waiters for async operations: `CreateSnapshotWaitHandler`, `DeleteSnapshotWaitHandler`

## v0.23.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v0.22.2 (2025-05-09)
- **Feature:** Update user-agent header

## v0.22.1 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v0.22.0 (2025-03-24)
- **Improvement:** Upgrading from IaaS **beta** endpoints to **v1**
- **Feature:** Add new method to filter `ListMachineTypes`: `Filter`

## v0.21.2 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.21.1 (2025-03-02)

- Increase Timeouts for volume and network wait handlers

## v0.21.0 (2025-02-27)

- **Feature:** Add method to list all public ip ranges: `ListPublicIpRanges`
- Add size attribute to image model
- Add CPU architecture attribute to image config model

## v0.20.0 (2025-02-21)

- **New:** Minimal go version is now Go 1.21

## v0.19.0 (2024-12-20)

- **Feature:** Add method to list quotas: `ListQuotas`
- **Feature:** Add methods to change image scope: `UpdateImageScopeLocal` and `UpdateImageScopePublic`
 
## v0.18.0 (2024-12-16)

- **Feature:** Add waiters for async operations: `UploadImageWaitHandler` and `DeleteImageWaitHandler`

## v0.17.0 (2024-12-16)

- **Feature:** Add new methods to manage affinity groups: `CreateAffinityGroup`, `DeleteAffinityGroup`, `GetAffinityGroup`, and `ListAffinityGroup`
- **Feature:** Add new methods to manage backups: `CreateBackup`, `DeleteBackup`, `GetBackup`, `ListBackup`, `RestoreBackup`, `ExecuteBackup`,`UpdateBackup`
- **Feature:** Add new methods to manage images: `CreateImage`, `DeleteImage`, `GetImage`, `ListImage`,`UpdateImage`
- **Feature:** Add new methods to manage imageshares: `DeleteImageShare`, `GetImageShare`, `SetImageShare`,`UpdateImageShare`
- **Feature:** Add new methods to manage imageshare consumers: `DeleteImageShareConsumer`, `GetImageShareConsumer`, `SetImageShare`,`UpdateImageShare`
- **Feature:** Add new methods to manage project NICs: `GetProjectNIC`, `ListProjectNICs`
- **Feature:** Add new methods to manage snapshots: `CreateSnapshot`, `DeleteSnapshot`, `GetSnapshot`, `ListSnapshot`, `UpdateSnapshot`
- **Bugfix:** Correctly handle nullable attributes in model types

## v0.16.0 (2024-11-08)

- **Feature:** Add new methods to manage key pairs: `CreateKeyPair`, `UpdateKeyPair`, `DeleteKeyPair`, `GetKeyPair`, and `ListKeyPairs`
- **Feature:** Add new field `Bootable` to `Volume`, `CreateVolumePayload`, and `UpdateVolumePayload` data models
- **Breaking change:** Rename `NIC` to `Nic` in all network interface methods (e.g. `CreateNIC` to `CreateNic`, `AddNICToServer` to `AddNicToServer`, etc)

## v0.15.0 (2024-10-21)

- **Feature:** Filter network area routes by labels using the new `LabelSelector` method on `ApiListNetworkAreaRoutesRequest`
- **Feature:** Update network area route using the new method `UpdateNetworkAreaRoute`

## v0.14.0 (2024-10-18)

- **Feature:** Add waiter methods for `Volume`, `Server` and `AttachedVolume`

## v0.13.0 (2024-10-18)

- **Feature:** Add support for managing following resources
  - `Volume`
  - `Server`
  - `NetworkInterface`
  - `PublicIP`
  - `SecurityGroup`
  - `SecurityGroupRule`
- **Breaking change**: Remove `V1NetworkGateway` data model
- **Bugfix**: Network response json decoding

## v0.12.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.11.0 (2024-10-11)

- **Feature:** Filter networks by labels using the new `LabelSelector` method on `ApiListNetworksRequest`

## v0.10.0 (2024-10-01)

- **Feature:** Add `CreatedAt` and `UpdatedAt` fields to several data models

## v0.9.0 (2024-09-27)

- **Feature:** Add `Labels` field to several data models

## v0.8.0 (2024-08-21)

- **Feature:** `CreateNetworkIPv4Body` and `CreateNetworkIPv6Body` have a new field `Prefix`

## v0.7.0 (2024-08-16)

- **Breaking change**: Rename types:
  - `CreateNetworkIPv4` renamed to **`CreateNetworkIPv4Body`**
  - `V1CreateNetworkIPv6` renamed to **`CreateNetworkIPv6Body`**
  - `UpdateNetworkIPv4` renamed to **`UpdateNetworkIPv4Body`**
  - `V1UpdateNetworkIPv6` renamed to **`UpdateNetworkIPv6Body`**
- **Feature:** `CreateNetworkPayload`, `PartialUpdateNetworkPayload` and `Network` have a new field: `Routed`

## v0.6.0 (2024-08-05)

- **Breaking change:** Use network ID instead of request ID in the waiter: `CreateNetworkWaitHandler`

## v0.5.0 (2024-07-24)

- **Feature:** `CreateNetworkAddressFamily` and `UpdateNetworkAddressFamily` have a new field `Ipv6`
- **Feature:** `Network` has new fields: `NameserversV6` and `PrefixesV6`

## v0.4.0 (2024-06-07)

- **Breaking change**: `CreateNetwork` now returns the `Network` triggered by the operation.

## v0.3.0 (2024-05-17)

- **Feature**: Add waiters for async operations: `CreateNetworkAreaWaitHandler`, `UpdateNetworkAreaWaitHandler`, `DeleteNetworkAreaWaitHandler`, `CreateNetworkWaitHandler`, `UpdateNetworkWaitHandler`, `DeleteNetworkWaitHandler`

## v0.2.0 (2024-05-17)

- **Feature**: New methods to manage networks:
  - `CreateNetwork`
  - `PartialUpdateNetwork`
  - `DeleteNetwork`
- **Breaking change**: Rename methods for better correspondence with endpoint behaviour:
  - `UpdateNetworkArea` renamed to **`PartialUpdateNetworkArea`**
  - `CreateNetworkAreas` renamed to **`CreateNetworkArea`**
  - `DeleteNetworkAreas` renamed to **`DeleteNetworkArea`** (same for the payload types)
- **Breaking change**: Rename types:
  - Add `Response` suffix to types only used in method responses:
    - `NetworkAreaList` renamed to **`NetworkAreaListResponse`**
    - `NetworkList` renamed to **`NetworkListResponsse`**
    - `NetworkRangeList` renamed to **`NetworkRangeListResponsse`**
    - `ProjectList` renamed to **`ProjectListResponsse`**
    - `RouteList` renamed to **`RouteListResponsse`**
  - Remove `V1` prefix from all types:
    - `V1NetworkArea` renamed to **`NetworkArea`**
    - `V1AreaConfig` renamed to **`AreaConfig`**
    - `V1Area` renamed to **`Area`**
    - `V1CreateAreaAddressFamily` renamed to **`CreateAreaAddressFamily`**
    - `V1CreateNetworkIPv4` renamed to **`CreateNetworkIPv4`**
    - `V1Error` renamed to **`Error`**
    - `V1NetworkAreaIPv4` renamed to **`NetworkAreaIPv4`**
    - `V1RequestResource` renamed to **`RequestResource`**
    - `V1UpdateAreaAddressFamily` renamed to **`UpdateAreaAddressFamily`**
    - `V1UpdateAreaIPv4` renamed to **`UpdateAreaIPv4`**
    - `V1UpdateNetworkIPv4` renamed to **`UpdateNetworkIPv4`**

## v0.1.0 (2024-05-10)

- **New BETA module**: Manage Infrastructure as a Service (IaaS) resources `Network` and `NetworkArea`
