## v0.5.0
- **Breaking Change:** Updated return types for `ImportKeyExecute` and `RotateKeyExecute` methods from `*Key` to `*Version`

## v0.4.0
- **Feature:** Add new wait handler for key ring creation (`CreateKeyRingWaitHandler`)

## v0.3.1
  - **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v0.3.0
- **Feature:** New method `DeleteWrappingKey`
- **Breaking change:** Enum `KEYSTATE_VERSION_NOT_READY` removed. Use instead `KEYSTATE_CREATING`
- **Breaking change:** Enum `VERSIONSTATE_KEY_MATERIAL_NOT_READY` removed. Use instead `VERSIONSTATE_CREATING`
- **Breaking change:** Enum `WRAPPINGKEYSTATE_KEY_MATERIAL_NOT_READY` removed. Use instead `WRAPPINGKEYSTATE_CREATING`
- **Feature:** New enums for `KEYSTATE`, `KEYRINGSTATE`, `VERSIONSTATE` and `WRAPPINGKEYSTATE`
- **Feature:** Add `required:"true"` tags to model structs

## v0.2.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v0.1.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v0.0.2 (2025-05-09)
- **Feature:** Update user-agent header

## v0.0.1 (2025-04-28)
- **New module:** Initial publication of Key Management Service API
