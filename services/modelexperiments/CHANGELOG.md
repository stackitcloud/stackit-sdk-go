## v0.3.0
- `v1api`:
  - **Fix:** Response decoding now supports `*io.Reader` and `*[]byte` target types (previously only `string`, `*os.File`, and JSON were supported)
  - **Fix:** Escape regex in validate tags of model structs correctly
  - **Breaking Change:** `Labels` field of `PartialUpdateInstancePayload`/`PartialUpdateInstanceTokenPayload` changed from `map[string]string` to `map[string]*string`

## v0.2.0
- **New**: STACKIT Model Experiments module wait handler added.

## v0.1.0
- **New**: API for STACKIT modelexperiments