module github.com/stackitcloud/stackit-sdk-go/examples/services

go 1.25.9

// This is not needed in production. This is only here to point the golangci linter to the local versions instead of the last release on GitHub.
replace (
	github.com/stackitcloud/stackit-sdk-go/core v0.27.1 => ../../core
	github.com/stackitcloud/stackit-sdk-go/services/authorization v0.16.0 => ../../services/authorization
	github.com/stackitcloud/stackit-sdk-go/services/automation v0.2.1 => ../../services/automation
	github.com/stackitcloud/stackit-sdk-go/services/cost v0.5.3 => ../../services/cost
	github.com/stackitcloud/stackit-sdk-go/services/dns v0.25.0 => ../../services/dns
	github.com/stackitcloud/stackit-sdk-go/services/functions v0.1.0 => ../../services/functions
	github.com/stackitcloud/stackit-sdk-go/services/git v0.15.0 => ../../services/git
	github.com/stackitcloud/stackit-sdk-go/services/iaas v1.15.1 => ../../services/iaas
	github.com/stackitcloud/stackit-sdk-go/services/kms v1.14.0 => ../../services/kms
	github.com/stackitcloud/stackit-sdk-go/services/loadbalancer v1.17.0 => ../../services/loadbalancer
	github.com/stackitcloud/stackit-sdk-go/services/logme v1.6.0 => ../../services/logme
	github.com/stackitcloud/stackit-sdk-go/services/mariadb v1.4.0 => ../../services/mariadb
	github.com/stackitcloud/stackit-sdk-go/services/modelexperiments v0.3.2 => ../../services/modelexperiments
	github.com/stackitcloud/stackit-sdk-go/services/objectstorage v1.12.0 => ../../services/objectstorage
	github.com/stackitcloud/stackit-sdk-go/services/observability v0.26.0 => ../../services/observability
	github.com/stackitcloud/stackit-sdk-go/services/opensearch v1.5.0 => ../../services/opensearch
	github.com/stackitcloud/stackit-sdk-go/services/postgresflex v1.16.0 => ../../services/postgresflex
	github.com/stackitcloud/stackit-sdk-go/services/rabbitmq v1.5.0 => ../../services/rabbitmq
	github.com/stackitcloud/stackit-sdk-go/services/redis v1.6.0 => ../../services/redis
	github.com/stackitcloud/stackit-sdk-go/services/registry v0.1.0 => ../../services/registry
	github.com/stackitcloud/stackit-sdk-go/services/resourcemanager v0.28.0 => ../../services/resourcemanager
	github.com/stackitcloud/stackit-sdk-go/services/runcommand v1.11.0 => ../../services/runcommand
	github.com/stackitcloud/stackit-sdk-go/services/sca v0.2.1 => ../../services/sca
	github.com/stackitcloud/stackit-sdk-go/services/secretsmanager v0.20.0 => ../../services/secretsmanager
	github.com/stackitcloud/stackit-sdk-go/services/serviceaccount v0.22.0 => ../../services/serviceaccount
	github.com/stackitcloud/stackit-sdk-go/services/serviceenablement v1.8.0 => ../../services/serviceenablement
	github.com/stackitcloud/stackit-sdk-go/services/sfs v0.13.1 => ../../services/sfs
	github.com/stackitcloud/stackit-sdk-go/services/ske v1.22.0 => ../../services/ske
	github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex v1.19.0 => ../../services/sqlserverflex
	github.com/stackitcloud/stackit-sdk-go/services/telemetrylink v0.6.0 => ../../services/telemetrylink
	github.com/stackitcloud/stackit-sdk-go/services/telemetryrouter v0.6.1 => ../../services/telemetryrouter
	github.com/stackitcloud/stackit-sdk-go/services/ufw v0.2.0 => ../../services/ufw
	github.com/stackitcloud/stackit-sdk-go/services/valkey v0.4.0 => ../../services/valkey
	github.com/stackitcloud/stackit-sdk-go/services/vpn v0.15.3 => ../../services/vpn
)

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.27.1
	github.com/stackitcloud/stackit-sdk-go/services/authorization v0.16.0
	github.com/stackitcloud/stackit-sdk-go/services/automation v0.2.1
	github.com/stackitcloud/stackit-sdk-go/services/cost v0.5.3
	github.com/stackitcloud/stackit-sdk-go/services/dns v0.25.0
	github.com/stackitcloud/stackit-sdk-go/services/iaas v1.15.1
	github.com/stackitcloud/stackit-sdk-go/services/kms v1.14.0
	github.com/stackitcloud/stackit-sdk-go/services/loadbalancer v1.17.0
	github.com/stackitcloud/stackit-sdk-go/services/logme v1.6.0
	github.com/stackitcloud/stackit-sdk-go/services/mariadb v1.4.0
	github.com/stackitcloud/stackit-sdk-go/services/modelexperiments v0.3.2
	github.com/stackitcloud/stackit-sdk-go/services/objectstorage v1.12.0
	github.com/stackitcloud/stackit-sdk-go/services/observability v0.26.0
	github.com/stackitcloud/stackit-sdk-go/services/opensearch v1.5.0
	github.com/stackitcloud/stackit-sdk-go/services/postgresflex v1.16.0
	github.com/stackitcloud/stackit-sdk-go/services/rabbitmq v1.5.0
	github.com/stackitcloud/stackit-sdk-go/services/redis v1.6.0
	github.com/stackitcloud/stackit-sdk-go/services/registry v0.1.0
	github.com/stackitcloud/stackit-sdk-go/services/resourcemanager v0.28.0
	github.com/stackitcloud/stackit-sdk-go/services/runcommand v1.11.0
	github.com/stackitcloud/stackit-sdk-go/services/sca v0.2.1
	github.com/stackitcloud/stackit-sdk-go/services/secretsmanager v0.20.0
	github.com/stackitcloud/stackit-sdk-go/services/serviceaccount v0.22.0
	github.com/stackitcloud/stackit-sdk-go/services/serviceenablement v1.8.0
	github.com/stackitcloud/stackit-sdk-go/services/sfs v0.13.1
	github.com/stackitcloud/stackit-sdk-go/services/ske v1.22.0
	github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex v1.19.0
	github.com/stackitcloud/stackit-sdk-go/services/telemetrylink v0.6.0
	github.com/stackitcloud/stackit-sdk-go/services/telemetryrouter v0.6.1
	github.com/stackitcloud/stackit-sdk-go/services/ufw v0.2.0
	github.com/stackitcloud/stackit-sdk-go/services/valkey v0.4.0
	github.com/stackitcloud/stackit-sdk-go/services/vpn v0.15.3
)

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
