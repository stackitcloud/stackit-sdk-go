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
