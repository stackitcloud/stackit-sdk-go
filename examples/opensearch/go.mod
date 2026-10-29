module github.com/stackitcloud/stackit-sdk-go/examples/opensearch

go 1.25

// NOTE: opensearch is currently on the SDK generator blocklist, so the replace is commented out.
// Re-enable it as soon as opensearch is removed from the blocklist again.
// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
// replace github.com/stackitcloud/stackit-sdk-go/services/opensearch => ../../services/opensearch

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.27.1
	github.com/stackitcloud/stackit-sdk-go/services/opensearch v1.3.1
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
