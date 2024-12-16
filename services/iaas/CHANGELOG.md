## v0.17.0 (2024-12-13)

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
