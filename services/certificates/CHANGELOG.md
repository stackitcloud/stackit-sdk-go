## v1.6.0
- Minimal go version is now Go 1.25

## v1.5.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v1.5.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v1.5.0
- **Feature:** Add new method `GetQuota` to get the quota for TLS certificates in a project
- `v2betaapi`:
  - **Deprecation:** The `v2beta2api` is deprecated and will be removed in the future. Migrate to `v2api`

## v1.4.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v1.4.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v1.3.0
- **Feature:** Introduction of multi API version support for the certificates SDK module. For more details please see the announcement on GitHub: https://github.com/stackitcloud/stackit-sdk-go/discussions/5062
- `v2api`: New package which can be used for communication with the certificates v2 API
- `v2betaapi`: New package which can be used for communication with the certificates v2 API
- **Deprecation:** The contents in the root of this SDK module are marked as deprecated and will be removed after 2026-09-30. Switch to the new packages for the available API versions instead.
- **Dependencies:** Bump STACKIT SDK core module from `v0.21.1` to `v0.22.0`

## v1.2.3
- Bump STACKIT SDK core module from `v0.21.0` to `v0.21.1`

## v1.2.2
- **Dependencies**: Bump `github.com/golang-jwt/jwt/v5` from `v5.3.0` to `v5.3.1`

## v1.2.1
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v1.2.0
  - **Feature:** Switch from `v2beta` API version to `v2` version.
  - **Breaking change:** Rename `CreateCertificateResponse` to `GetCertificateResponse`

## v1.1.3
  - Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v1.1.2
  - Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

## v1.1.1
  - **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v1.1.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs
 
## v1.0.3 (2025-05-09)
- **Feature:** Update user-agent header

## v1.0.2 (2025-04-29)
- **Bugfix:** Correctly handle empty payload in body

## v1.0.1 (2025-03-27)
- **Bugfix:** Removed ConfigureRegion() from API client

## v1.0.0 (2025-03-14)
- **Breaking Change:** The region is no longer specified within the client configuration. Instead, the region must be passed as a parameter to any region-specific request.

## v0.2.0 (2025-02-21)
- **New:** Minimal go version is now Go 1.21

## v0.1.0 (2025-01-13)

- **New**: STACKIT Certificates module can be used to manage the STACKIT Load Balancer Certificates
