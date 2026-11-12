## v0.3.0
- **Breaking Change:** Rule type `PublicIP` is no longer supported: it can no longer be used to create rules (`CreateRulePayload.Type`) or to filter `ListRules`, and the `RULERESPONSETYPE_PUBLIC_IP` value was removed from the `RuleResponseType` enum
- **Feature:** Add `RemoteSecurityGroupName` field to `RuleResponse`

## v0.2.0
- **Breaking Change:** `RuleResponse.Type` field changed from `string` to new `RuleResponseType` enum (adds `SecurityGroup` to the existing `ACL`/`PublicIP`/`SecurityRule` values); `NewRuleResponse`, `GetType`, `GetTypeOk`, `SetType` signatures updated accordingly
- **Breaking Change:** `RuleResponse.SecurityGroup` field removed (and its `GetSecurityGroup`/`GetSecurityGroupOk`/`HasSecurityGroup`/`SetSecurityGroup` accessors)
- **Breaking Change:** `ListProviderOptionsExecute` now returns `*ProviderOptionsResponse` instead of `[]ProviderOptionsResponse`, matching the endpoint's actual (non-array) response
- **Feature:** Add `RefId` and `Stateful` fields to `RuleResponse`
- **Feature:** Add `RemoteSecurityGroupId` and `Stateful` fields to `CreateRulePayload`
- **Dependencies:** Bump STACKIT SDK core module from `v0.27.0` to `v0.27.1`

## v0.1.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.26.0` to `v0.27.0`

## v0.1.0
- **New:** STACKIT Unified Firewall (UFW)
- **Feature:** Add waiter and example methods for the API