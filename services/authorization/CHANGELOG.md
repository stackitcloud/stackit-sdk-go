## v0.7.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v0.6.3 (2025-05-09)
- **Feature:** Update user-agent header

## v0.6.2 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v0.6.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.6.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.5.1 (2025-01-17)

- **Bugfix:** Revert back to global URL configuration 

## v0.5.0 (2025-01-09)

- Add support for regions in `APIClient`

## v0.4.1 (2024-12-17)

- **Bugfix:** Correctly handle nullable attributes in model types

## v0.4.0 (2024-10-14)

- **Feature:** Add support for nullable models
- **Feature:** New method `ListUserPermissions`

## v0.3.0 (2024-05-22)

- **Feature:** `Role` has a new field `Id`

## v0.2.0 (2024-04-11)

- Set config.ContextHTTPRequest in Execute method
- Support WithMiddleware configuration option in the client
- Update `core` to [`v0.12.0`](../../core/CHANGELOG.md#v0120-2024-04-11)

## v0.1.2 (2024-04-09)

- Removed unused model data types.

## v0.1.1 (2024-02-28)

- Update `core` to [`v0.10.0`](../../core/CHANGELOG.md#v0100-2024-02-27)

## v0.1.0 (2024-02-07)

First release.

This module offers the same functionalities as `membership` (release v0.4.0).
