# STACKIT Go SDK examples

This directory contains examples for the STACKIT Go SDK. Please follow the structure explained below. 

Please note that the examples should be using at most 2 API versions of a STACKIT service:
- 1 stable API version is allowed
- 1 unstable API version (alpha, beta) is allowed

```
📂 examples
├── 📂 services
│   ├── 📄 go.mod
│   ├── 📄 go.sum
│   ├── 📂 ske                                 # STACKIT service SDK module identifier
│   │   └── 📄 ske.go                          # Example for the latest stable API version (e.g. v2)
│   ├── 📂 dremio
│   │   └── 📂 preview                         # The example for a unstable API version (alpha/beta) belong into the 'preview' directory
│   │       └── 📄 dremio_preview.go
│   ├── 📂 iaas                                # For services with many resources you can also split the examples to multiple directories
│   │   ├── 📂 network
│   │   │   └── 📄 network.go
│   │   ├── 📂 server
│   │   │   └── 📄 server.go
│   │   └── 📂 volume
│   │       └── 📄 volume.go
│   └── 📂 vpn                                 # You can also have examples for a stable API version and a upcoming alpha/beta API version in parallel
│       ├── 📄 vpn.go                          # Example for the latest stable API version (e.g. v2)
│       └── 📂 preview                         # Example for the next upcoming API version which isn't stable yet (e.g. v3alpha / v3beta)
│           └── 📄 dremio_preview.go
└── 📂 ...
```
