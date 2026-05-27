## v0.13.0
  - `v1api`:
    - **Breaking change:** `LocalAsn` is now marked as required (aligned with api) and no longer a pointer
  - `v1beta1api`:  Align package to latest API specification
  - `v1alpha1api`: Align package to latest API specification

## v0.12.0
- **Feature:** Introduce enums for various attributes

## v0.11.0
- `v1api`:
  - **Feature:** Add `Labels` field to `CreateGatewayConnectionPayload` and `UpdateGatewayConnectionPayload`
- `v1beta1api`:  Align package to latest API specification
- `v1alpha1api`: Align package to latest API specification

## v0.10.0
- `v1api`:
  - **Feature:** Add `InternalNextHopIP` field to `VPNTunnels` model
- `v1beta1api`:  Align package to latest API specification
- `v1alpha1api`: Align package to latest API specification

## v0.9.0
- `v1api`:
  - **Breaking change:** Switch from regional to global API server URL. `config.WithRegion(...)` should not be used during client initialization anymore.
- `v1beta1api`: Align package to latest API specification
- `v1alpha1api`: Align package to latest API specification

## v0.8.0
- `v1api`:
  - **Feature:** Add new wait handlers for gateway creation (`CreateGatewayWaitHandler`), update (`UpdateGatewayWaitHandler`), and deletion (`DeleteGatewayWaitHandler`)

## v0.7.0
- `v1api`: New package which can be used for communication with the STACKIT vpn v1 API
- `v1beta1api`:  Align package to latest API specification
- `v1alpha1api`: Align package to latest API specification

## v0.6.0
- **Feature:** Added `_UNKNOWN_DEFAULT_OPEN_API` fallback value to all enums to handle unknown API values gracefully.

## v0.5.2
- **Dependencies:** Bump STACKIT SDK core module to `v0.26.0`

## v0.5.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.1` to `v0.25.0`

## v0.5.0
- Minimal go version is now Go 1.25

## v0.4.2
- **Dependencies:** Bump STACKIT SDK core module from `v0.24.0` to `v0.24.1`

## v0.4.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.23.0` to `v0.24.0`

## v0.4.0
- `v1beta1api`: New package which can be used for communication with the STACKIT vpn v1 beta1 API

## v0.3.1
- **Dependencies:** Bump STACKIT SDK core module from `v0.22.0` to `v0.23.0`

## v0.3.0
- **Bugfix:** Disable strict decoding of API responses
- **Feature:** Add `AdditionalProperties` fields to model structs

## v0.2.0
- `v1alpha1api`: Align package to latest API specification

## v0.1.0
- **New:** SDK module for vpn service
- `v1alpha1api`: New package which can be used for communication with the vpn v1 alpha1 API
