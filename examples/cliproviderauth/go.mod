module github.com/stackitcloud/stackit-sdk-go/examples/cliproviderauth

go 1.21

require (
	github.com/stackitcloud/stackit-sdk-go/core v0.20.0
	github.com/stackitcloud/stackit-sdk-go/services/dns v0.17.2
)

require (
	al.essio.dev/pkg/shellescape v1.5.1 // indirect
	github.com/danieljoos/wincred v1.2.2 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/zalando/go-keyring v0.2.6 // indirect
	golang.org/x/sys v0.26.0 // indirect
)

// Use local version until CLI auth is released
replace github.com/stackitcloud/stackit-sdk-go/core => ../../core

replace github.com/stackitcloud/stackit-sdk-go/services/dns => ../../services/dns
