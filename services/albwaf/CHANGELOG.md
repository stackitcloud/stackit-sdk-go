## v0.5.0 
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.1` to `v0.25.0`
- **Feature:** new models `CRSRule`, `CRSRuleGroup`, `GetLimitedCoreRuleSetResponse`
- **Feature:** add fields `Groups` and `Version` to `GetCoreRuleSetResponse`
- **Breaking change:** change type of `Items` field in `ListCoreRuleSetResponse` model from `GetCoreRuleSetResponse` to `GetLimitedCoreRuleSetResponse`

## v0.4.0
- Minimal go version is now Go 1.25

## v0.3.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v0.3.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v0.3.0
- **Feature:** Add new method `GetQuota` to get the quota for WAF resources in a project

## v0.2.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v0.2.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v0.1.0
- **New:** SDK module for albwaf service
- `v1alphaapi`: New package which can be used for communication with the albwaf v1 alpha API
