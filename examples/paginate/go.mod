module github.com/stackitcloud/stackit-sdk-go/examples/paginate

go 1.25

// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
replace github.com/stackitcloud/stackit-sdk-go/services/automation => ../../services/automation

replace github.com/stackitcloud/stackit-sdk-go/services/albwaf => ../../services/albwaf

replace github.com/stackitcloud/stackit-sdk-go/experimental => ../../experimental

require (
	github.com/stackitcloud/stackit-sdk-go/experimental v0.1.0
	github.com/stackitcloud/stackit-sdk-go/services/albwaf v0.13.2
	github.com/stackitcloud/stackit-sdk-go/services/automation v0.1.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/stackitcloud/stackit-sdk-go/core v0.26.0 // indirect
)
