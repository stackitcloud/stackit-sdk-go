module github.com/stackitcloud/stackit-sdk-go/examples/intake

go 1.21

// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
replace github.com/stackitcloud/stackit-sdk-go/services/intake => ../../services/intake

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.24.0
	github.com/stackitcloud/stackit-sdk-go/services/intake v0.7.2
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
