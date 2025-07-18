## v0.13.0
- Add `required:"true"` tags to model structs

## v0.12.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v0.11.4 (2025-05-09)
- **Feature:** Update user-agent header

## v0.11.3 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v0.11.2 (2025-03-20)
- **Improvement:** Error handling
- **Feature:** Add description to `UpdateUserPayload`

## v0.11.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.11.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.10.1 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.10.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.9.0 (2024-08-16)

- **Feature**: New API method `UpdateInstance` to update an instance

## v0.8.0 (2024-05-23)

- **Breaking change**: Rename data types for uniformity
  - `Acl` is now `ACL`
  - `AclList` is now `ListACLsResponse`
  - `InstanceList` is now `ListInstancesResponse`
  - `UserList` is now `ListUsersResponse`
- **Breaking change**: Remove unused data types

## v0.7.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.6.0 (2024-03-18)

- **Feature**: New API method `UpdateACLs` to update all ACLs of an instance

## v0.5.6 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.5.5 (2024-02-02)

- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.5.4 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.5.3 (2024-01-15)

- Add license and notice files

## v0.5.2 (2024-01-09)

- **Improvement:** Description fixes
- Dependency updates

## v0.5.1 (2023-12-22)

- Dependency updates

## v0.5.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `CreateAcl` renamed to `CreateACL`
  - `DeleteAcl` renamed to `DeleteACL`
  - `GetAcl` renamed to `GetACL`
  - `GetAcls` renamed to `ListACLs`
  - `GetInstances` renamed to `ListInstances`
  - `UpdateAcl` renamed to `UpdateACL`
- Changed structs:
  - `CreateAclPayload` renamed to `CreateACLPayload`
  - `UpdateAclPayload` renamed to `UpdateACLPayload`

## v0.4.0 (2023-11-10)

- Manage your STACKIT Secrets Manager resources: `Instance`, `Acl`, `User`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/secretsmanager)
