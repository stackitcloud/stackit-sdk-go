module github.com/stackitcloud/stackit-sdk-go/examples/backgroundrefresh

go 1.25

// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
replace github.com/stackitcloud/stackit-sdk-go/services/dns => ../../services/dns

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.25.0
	github.com/stackitcloud/stackit-sdk-go/services/dns v0.19.3
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
