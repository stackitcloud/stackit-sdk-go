module github.com/stackitcloud/stackit-sdk-go/examples/preview

go 1.25

// This is not needed in production. This is only here to point the golangci linter to the local version instead of the last release on GitHub.
replace (
	github.com/stackitcloud/stackit-sdk-go/core v0.27.1 => ../../core
	github.com/stackitcloud/stackit-sdk-go/services/albwaf v0.13.4 => ../../services/albwaf
	github.com/stackitcloud/stackit-sdk-go/services/automation v0.1.1 => ../../services/automation
	github.com/stackitcloud/stackit-sdk-go/services/dremio v0.7.0 => ../../services/dremio
	github.com/stackitcloud/stackit-sdk-go/services/edge v0.14.0 => ../../services/edge
	github.com/stackitcloud/stackit-sdk-go/services/functions v0.1.0 => ../../services/functions
	github.com/stackitcloud/stackit-sdk-go/services/git v0.15.0 => ../../services/git
	github.com/stackitcloud/stackit-sdk-go/services/intake v0.12.0 => ../../services/intake
)

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.27.1
	github.com/stackitcloud/stackit-sdk-go/services/albwaf v0.13.4
	github.com/stackitcloud/stackit-sdk-go/services/automation v0.1.1
	github.com/stackitcloud/stackit-sdk-go/services/dremio v0.7.0
	github.com/stackitcloud/stackit-sdk-go/services/edge v0.14.0
	github.com/stackitcloud/stackit-sdk-go/services/functions v0.1.0
	github.com/stackitcloud/stackit-sdk-go/services/git v0.15.0
	github.com/stackitcloud/stackit-sdk-go/services/intake v0.12.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
