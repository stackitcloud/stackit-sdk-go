package v1api

/*
STACKIT Key Management Service API

This API provides endpoints for managing keys and key rings.

API version: 1.0.0
*/

// This is a handwritten file and is not generated.

// Deprecated: Protection is deprecated and will be removed after February 2027. Use instead `string`.
type Protection = string

// List of protection
const (
	// Deprecated: PROTECTION_SOFTWARE is deprecated and will be removed after February 2027.
	PROTECTION_SOFTWARE Protection = "software"
	// Deprecated: PROTECTION_SOFTWARE is deprecated and will be removed after February 2027.
	PROTECTION_HSM Protection = "hsm"
	// Deprecated: PROTECTION_UNKNOWN_DEFAULT_OPEN_API is deprecated and will be removed after February 2027.
	PROTECTION_UNKNOWN_DEFAULT_OPEN_API Protection = "unknown_default_open_api"
)

// Deprecated: AllowedProtectionEnumValues is deprecated and will be removed after February 2027.
var AllowedProtectionEnumValues = []Protection{
	"software",
	"hsm",
	"unknown_default_open_api",
}
