module github.com/stackitcloud/stackit-sdk-go/examples/iaas

go 1.21

// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
replace github.com/stackitcloud/stackit-sdk-go/services/iaas => ../../services/iaas

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.23.0
	github.com/stackitcloud/stackit-sdk-go/services/iaas v1.3.5
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/stackitcloud/stackit-sdk-go/services/resourcemanager v0.21.0 // indirect
)
