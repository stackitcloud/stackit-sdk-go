## v1.9.2
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v1.9.1
- Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v1.9.0
- **Feature:** Switch from `v1beta2` CDN API version to `v1` version.
- **Feature:** Add new function `GetLogsSearchFilters`
- **Feature:** Add `WafAction` method to `ApiGetLogsRequest` struct
- **Feature:** Add `Reason` field to `WafStatusRuleBlock` model struct
- **Feature:** Add fields to `WafConfig` and `WafConfigPatch` model struct:
  - `AllowedHttpMethods`
  - `AllowedHttpVersions`
  - `AllowedRequestContentTypes`
  - `DisabledRuleCollectionIds`
  - `DisabledRuleGroupIds`
  - `DisabledRuleIds`
  - `EnabledRuleCollectionIds`
  - `EnabledRuleGroupIds`
  - `EnabledRuleIds`
  - `LogOnlyRuleCollectionIds`
  - `LogOnlyRuleGroupIds`
  - `LogOnlyRuleIds`
  - `ParanoiaLevel`

## v1.8.1
- **Note: This release was formerly known as `v2.1.1` and was re-tagged, see statement below.**
- Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

> [!IMPORTANT]
> The 3 releases, which contained the previously tagged `v2.x.x` changes, are now re-released as `v1.7.0`, `v1.8.0` and `v1.8.1`.
> We have retagged the releases as `v1.x.x` versions to avoid forcing a module path change on all consumers.
>
> **Reason:** According to [Go module conventions](https://go.dev/blog/v2-go-modules), major versions ≥2 require the module path to be updated with a `/v2` suffix (`github.com/stackitcloud/stackit-sdk-go/services/cdn`).
>
> **Impact:** By remaining at `v1`, all users can continue to import the module using the original, clean import path (`github.com/stackitcloud/stackit-sdk-go/services/cdn`) without needing to update their import statements.
>
> **Note: If you were trying to use any `v2.x.x` tag, please downgrade to `v1.7.0` or higher. There won't be any `v2.x.x` release in the near future of any STACKIT SDK module.**
>
> We apologize for any confusion caused by the `v2.x.x` tags. We have a linter in place to prevent this in the future.

## v1.8.0
- **Note: This release was formerly known as `v2.1.0` and was re-tagged, see statement above.**
- **Breaking change:** Removal of unused model structs: `GetLogsSearchFiltersResponse`, `PatchLokiLogSink`

## v1.7.0
- **Note: This release was formerly known as `v2.0.0` and was re-tagged, see statement above.**
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
