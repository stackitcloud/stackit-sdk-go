# CLI Provider Authentication Example

This example demonstrates how to use the STACKIT CLI provider authentication in your Go applications.

## Overview

The CLI provider authentication feature enables applications (like the Terraform Provider) to use credentials stored by the STACKIT CLI without requiring direct dependency on CLI code or re-authentication.

## Features Demonstrated

1. **Default Profile Authentication**: Use credentials from the default CLI profile
2. **Specific Profile Authentication**: Use credentials from a named CLI profile (e.g., "production")
3. **Direct Credential Access**: Advanced use case for direct credential manipulation

## Prerequisites

Before running this example, you need to authenticate with the STACKIT CLI:

```bash
stackit auth login
```

For multiple profiles:

```bash
stackit auth login --profile production
```

## How It Works

The SDK automatically:
- Reads credentials from the system keyring (macOS Keychain, Linux Secret Service, Windows Credential Manager)
- Falls back to file-based storage if keyring is unavailable
- Refreshes OAuth2 tokens automatically when they expire
- Syncs refreshed credentials back to storage

## Storage Locations

**System Keyring** (preferred):
- Service name: `stackit-cli-provider` or `stackit-cli-provider/{profile}`

**File Fallback**:
- Default profile: `~/.stackit/cli-provider-auth-storage.txt`
- Custom profiles: `~/.stackit/profiles/{profile}/cli-provider-auth-storage.txt`

## Running the Example

1. Ensure you're authenticated with the STACKIT CLI
2. Update the `projectId` in the code
3. Run:

```bash
cd examples/cliproviderauth
go run cliproviderauth.go
```

## Profile Resolution

Profiles are resolved in the following order:
1. Explicit profile parameter in code
2. `STACKIT_CLI_PROFILE` environment variable
3. `~/.config/stackit/cli-profile.txt` file
4. "default" profile

## Backward Compatibility

The cliauth package maintains 100% backward compatibility with credentials created by existing STACKIT CLI versions. Users can seamlessly switch between CLI and SDK-based tools without re-authenticating.
