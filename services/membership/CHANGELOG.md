## v0.3.3 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient`.

## v0.3.2 (2024-01-09)

- Dependency updates

## v0.3.1 (2023-12-22)

- Dependency updates

## v0.3.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `DeleteMembers` renamed to `RemoveMembers`
  - `GetMembers` renamed to `ListMembers`
  - `GetMemberships` renamed to `ListUserMemberships`
  - `GetPermissions` renamed to `ListPermissions`
  - `GetRoles` renamed to `ListRoles`
  - `UpdateMembers` renamed to `AddMembers`
- Changed structs:
  - `AddCustomRolesPayload` renamed to `AddRolesPayload`
  - `AddRoleRequest` renamed to `AddRolesPayloadItem`
  - `ChildMembersPayload` renamed to `ValidateChildMembersPayload`
  - `DeleteMembersPayload` renamed to `RemoveMembersPayload`
  - `EnforcePermissionRequest` renamed to `EnforcePermissionPayload`
  - `GetMembersResponse` renamed to `ListMembersResponse`
  - `MembershipsResponse` renamed to `ListUserMembershipsResponse`
  - `SubjectIDsResponse` renamed to `ListSubjectIdsResponse`
  - `SubjectsResponse` renamed to `ListSubjectsResponse`
  - `TransferSubjectsMembershipsPayload` renamed to `TransferSubjectMembershipsPayload`
  - `UpdateMembersPayload` renamed to `AddMembersPayload`
  - `UserPermissionsResponse` renamed to `ListUserPermissionsResponse`
  - `WriteSchemaRequest` renamed to `WriteSchemaPayload`

## v0.2.0 (2023-11-10)

- Manage membership of your STACKIT esources
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/membership)
