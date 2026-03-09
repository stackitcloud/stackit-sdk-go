## v0.2.0
- **Feature:** Introduction of multi API version support for the auditlog SDK module. For more details please see the announcement on GitHub: https://github.com/stackitcloud/stackit-sdk-go/discussions/5062
- `v2api`: New package which should be used for communication with the STACKIT auditlog API in the future
- **Deprecation:** The contents in the root of this SDK module including the `wait` package are marked as deprecated and will be removed after 2026-09-30. Switch to the new `v2api` package instead.
- **Dependencies:** Bump STACKIT SDK core module from `v0.21.1` to `v0.22.0`

## v0.1.5
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v0.1.4
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v0.1.3
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v0.1.2
- Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v0.1.1
- Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

## v0.1.0
- **New**: STACKIT Audit Log module for retrieving recorded actions and system changes. Download audit log entries for folders, organizations, and projects with time range filtering, pagination, and configurable result limits. 
