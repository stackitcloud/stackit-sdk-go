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
