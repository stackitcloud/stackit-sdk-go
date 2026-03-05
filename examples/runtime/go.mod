module github.com/stackitcloud/stackit-sdk-go/examples/runtime

go 1.21

// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
replace github.com/stackitcloud/stackit-sdk-go/services/postgresflex => ../../services/postgresflex

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.22.0
	github.com/stackitcloud/stackit-sdk-go/services/postgresflex v1.3.5
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
