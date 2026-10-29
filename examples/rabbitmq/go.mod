module github.com/stackitcloud/stackit-sdk-go/examples/rabbitmq

go 1.25

// NOTE: rabbitmq is currently on the SDK generator blocklist, so the replace is commented out.
// Re-enable it as soon as rabbitmq is removed from the blocklist again.
// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
// replace github.com/stackitcloud/stackit-sdk-go/services/rabbitmq => ../../services/rabbitmq

require github.com/stackitcloud/stackit-sdk-go/services/rabbitmq v1.3.1

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/stackitcloud/stackit-sdk-go/core v0.27.1 // indirect
)
