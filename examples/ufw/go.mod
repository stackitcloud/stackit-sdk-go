module github.com/stackitcloud/stackit-sdk-go/examples/ufw

go 1.25.9

replace github.com/stackitcloud/stackit-sdk-go/services/ufw => ../../services/ufw

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.27.0
	github.com/stackitcloud/stackit-sdk-go/services/ufw v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
