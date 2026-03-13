module github.com/stackitcloud/stackit-sdk-go/examples/edge

go 1.21

// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
replace github.com/stackitcloud/stackit-sdk-go/services/edge => ../../services/edge

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.23.0
	github.com/stackitcloud/stackit-sdk-go/services/edge v0.4.3
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
