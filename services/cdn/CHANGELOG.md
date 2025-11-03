## v2.1.0
- **Breaking change:** Removal of unused model structs: `GetLogsSearchFiltersResponse`, `PatchLokiLogSink`

## v2.0.0
- **Feature:** Switch from `v1beta` CDN API version to `v1beta2` version.
- **Breaking change:** Changed spelling from `WAF` to `Waf` in model struct names
  - `WAFStatusRuleBlock` -> `WafStatusRuleBlock`
  - `WAFRuleGroup` -> `WafRuleGroup`
  - `WAFRuleCollection` -> `WafRuleCollection`
  - `WAFRule` -> `WafRule`
  - `NullableListWAFCollectionsResponse` -> `NullableListWafCollectionsResponse`
- **Breaking change:** Changed spelling from model struct named `GenericJSONResponse` to `GenericJsonResponse`
- **Breaking change:** Removal of fields from model structs
  - Remove `Description` field from `ErrorDetails` model struct
  - Remove `Domain` field from `PutCustomDomainResponse` and `GetCustomDomainResponse` model structs
  - Remove `OccuredAt` field from `GetCacheInfoResponseHistoryEntry` model struct
- **Breaking change:** Renaming of fields in model structs
  - Rename `DistributionID` field to `DistributionId` in `DistributionLogsRecord` model struct
  - Rename `BlockedIPs` field to `BlockedIps` in `CreateDistributionPayload`, `ConfigPatch` and `Config` model structs
- **Breaking change:** Removal of API client methods `GetLogsSearchFilters`, `GetLogsSearchFiltersExecute`
- **Breaking change:** Removal of request structs `GetLogsSearchFiltersRequest`
- **Feature:** Add fields to model structs
  - Add `Backend` field to `CreateDistributionPayload` model struct
  - Add `BucketBackend` field to `ConfigBackend` model struct
  - Add `BucketBackendPatch` field to `ConfigPatchBackend` model struct
- **Feature:** New model structs
  - New Loki model structs: `LokiLogSinkCredentials`, `LokiLogSinkCreate`, `LokiLogSinkPatch`
  - New Backend model structs: `HttpBackendCreate`,  `BucketBackendCreate`, `BucketBackend`, `BucketBackendPatch`, `CreateDistributionPayloadBackend`
  - Other new model structs: `BucketCredentials`

## v1.6.0
- **Feature:** Add models: `DistributionWaf`, `WafConfig`, `WAFConfigPatch`, `WAFMode`, `WAFRule`, `WAFRuleCollection`, `WAFRuleGroup` and `WAFStatusRuleBlock`
- **Feature:** Add `Waf` attribute to `Config` and `Distribution`
- **Feature:** Add `WafStatus` to `DistributionRequest` and `ListWafCollections`

## v1.5.0
- **Feature:** Added Attribute `LogSink` to `ConfigPatch`
- **Feature:** Added Attribute `Geofencing` to `DistributionPayload`
- **Feature:** Added new function `GetLogsSearchFilters`

## v1.4.0
- **Feature:** Added new filter functions `DataCenterRegion`, `RequestCountryCode`, `StatusCode` and `CacheHit`
- **Feature:** Added Attribute `LogSink` and `Certificate`
- **Feature:** Added `ConfigLogSink` and `PatchLokiLogSink` functionality

## v1.3.2
- **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v1.3.1
- **Improvement:** Improve parsing of oneOf

## v1.3.0
- **Feature:** Add `DefaultCacheDuration` field to `Config`, `ConfigPatch`, and `CreateDistributionPayload` models
- Add `required:"true"` tags to model structs

## v1.2.1 (2025-06-04)
- **Bugfix:** Adjusted `UnmarshalJSON` function to use enum types and added tests for enums
- **Feature:** Added `Optimizer` attribute

## v1.2.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v1.1.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v1.0.2 (2025-05-09)
- **Feature:** Update user-agent header

## v1.0.1 (2025-05-05)
- **Enhancement:** Increase waiter timeouts

## v1.0.0 (2025-05-02)
- **Feature:** Support for log management
- **Feature:** Create distribution payload has additional optional attributes for blocked countries, IPs and volume limitation
- **Feature:** Config Patch payload has additional optional attributes for blocked countries, IPs and volume limitation
- **Breaking Change:** Config has additional required attributes for blocked countries, IPs and volume limitation

## v0.3.1 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v0.3.0 (2025-04-04)
- **New:** Add waiter for creation of `CustomDomain`

## v0.2.0 (2025-04-01)
- **API enhancement:** Provide waiter infrastructure

## v0.1.1 (2025-03-27)
- **Bugfix:** Removed ConfigureRegion() from API client

## v0.1.0 (2025-03-19)
- **New:** Introduce new API for content delivery
