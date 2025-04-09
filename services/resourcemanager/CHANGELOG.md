## v0.14.0 (2025-04-09)
- **Feature:** Add support for folder management

## v0.13.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.13.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.12.0 (2025-01-31)

- **Breaking Change**: Remove the methods `BffGetContainersOfAFolder` and `BffGetContainersOfAnOrganization`

## v0.11.1 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.11.0 (2024-10-21)
- **Feature:** Get containers of a folder using the new method `BffGetContainersOfAFolder`
- **Feature:** Get containers of an organization using the new method `BffGetContainersOfAnOrganization`

## v0.10.0 (2024-10-14)

- **Feature:** Add support for nullable models

## v0.9.0 (2024-06-14)

- **Breaking Change**: Rename data types for uniformity
  - `ProjectResponse` -> `Project`
  - `ProjectResponseWithParents` -> `GetProjectResponse`
  - `AllProjectsResponse` -> `ListProjectsResponse`
- **Breaking Change**: Delete unused data types
- **Feature**: New methods `GetOrganization` and `ListOrganizations`
- Updated examples

## v0.8.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.7.7 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.7.6 (2024-02-02)

- Update `core` to `v0.7.7`. The `http.request` context is now passed in the client `Do` call.

## v0.7.5 (2024-01-24)

- **Bug fix**: `NewAPIClient` now initializes a new client instead of using `http.DefaultClient` ([#236](https://github.com/stackitcloud/stackit-sdk-go/issues/236))

## v0.7.4 (2024-01-15)

- Add license and notice files

## v0.7.3 (2024-01-09)

- Dependency updates

## v0.7.2 (2024-01-02)

- **Bugfix** Fixed slice parameter value formatting. This fixed the bug where providing `ContainerIds` as a query parameter never returned a valid result.

## v0.7.1 (2023-12-22)

- Dependency updates

## v0.7.0 (2023-12-20)

API methods, structs and waiters were renamed to have the same look and feel across all services and according to user feedback.

- Changed methods:
  - `GetProjects` renamed to `ListProjects`
  - `UpdateProject` renamed to `PartialUpdateProject`
- Changed structs:
  - `UpdateResourcePayload` renamed to `PartialUpdateResourcePayload`

## v0.6.0 (2023-11-10)

- Manage your STACKIT projects
- Waiters for async operations: `CreateProjectWaitHandler`, `DeleteProjectWaitHandler`
- [Usage example](https://github.com/stackitcloud/stackit-sdk-go/tree/main/examples/resourcemanager)
