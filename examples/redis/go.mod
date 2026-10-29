module github.com/stackitcloud/stackit-sdk-go/examples/redis

go 1.25

// NOTE: redis is currently on the SDK generator blocklist, so the replace is commented out.
// Re-enable it as soon as redis is removed from the blocklist again.
// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
// replace github.com/stackitcloud/stackit-sdk-go/services/redis => ../../services/redis

require github.com/stackitcloud/stackit-sdk-go/services/redis v1.4.1

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/stackitcloud/stackit-sdk-go/core v0.27.1 // indirect
)
