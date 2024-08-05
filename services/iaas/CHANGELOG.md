## v0.6.0 (2024-08-05)
- **Feature:** Use network ID in the waiter: `CreateNetworkWaitHandler`

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
