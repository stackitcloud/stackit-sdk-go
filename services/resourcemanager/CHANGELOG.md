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
