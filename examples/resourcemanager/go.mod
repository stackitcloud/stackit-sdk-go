module github.com/stackitcloud/stackit-sdk-go/examples/resourcemanager

go 1.21

// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
replace github.com/stackitcloud/stackit-sdk-go/services/resourcemanager => ../../services/resourcemanager

require github.com/stackitcloud/stackit-sdk-go/services/resourcemanager v0.0.0-00010101000000-000000000000

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/stackitcloud/stackit-sdk-go/core v0.24.1 // indirect
)
