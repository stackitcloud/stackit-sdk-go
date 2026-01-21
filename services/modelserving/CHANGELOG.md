## v0.6.3
- **Bugfix:** Correctly handle file closing for file uploads
- Bump STACKIT SDK core module from `v0.20.1` to `v0.21.0`

## v0.6.2
  - Bump STACKIT SDK core module from `v0.20.0` to `v0.20.1`

## v0.6.1
  - Bump STACKIT SDK core module from `v0.19.0` to `v0.20.0`

## v0.6.0
- **Feature:** New enum values `MODELTYPE_AUDIO` and `MODELTYPE_IMAGE` for `ModelTypes` enum
- **Breaking Change:** Renaming of `ChatModelDetailsBits` enum fields
    - `CHATMODELDETAILSBITS_1Bit` -> `CHATMODELDETAILSBITS_ONE_BIT`
    - `CHATMODELDETAILSBITS_2Bits` -> `CHATMODELDETAILSBITS_TWO_BITS`
    - `CHATMODELDETAILSBITS_4Bits` -> `CHATMODELDETAILSBITS_FOUR_BITS`
    - `CHATMODELDETAILSBITS_8Bits` -> `CHATMODELDETAILSBITS_EIGHT_BITS`
    - `CHATMODELDETAILSBITS_16Bits` -> `CHATMODELDETAILSBITS_SIXTEEN_BITS`

## v0.5.1
- **Dependencies:** Bump `github.com/golang-jwt/jwt/v5` from `v5.2.2` to `v5.2.3`

## v0.5.0
- Add `required:"true"` tags to model structs

## v0.4.1 (2025-06-04)
- **Bugfix:** Adjusted `UnmarshalJSON` function to use enum types and added tests for enums

## v0.4.0 (2025-05-15)
- **Breaking change:** Introduce interfaces for `APIClient` and the request structs

## v0.3.0 (2025-05-14)
- **Breaking change:** Introduce typed enum constants for status attributes

## v0.2.3 (2025-05-09)
- **Feature:** Update user-agent header

## v0.2.2 (2025-05-02)
- **Feature:** Improved nil-safety

## v0.2.1 (2025-03-19)
- **Internal:** Backwards compatible change to generated code

## v0.2.0 (2025-03-14)

- **New**: STACKIT Model Serving module wait handler added.

## v0.1.0 (2025-02-25)

- **New**: STACKIT Model Serving module can be used to manage the STACKIT Model Serving.
