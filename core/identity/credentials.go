package identity

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// CredentialType represents a specific credential field type
type CredentialType string

// Credentials represents the structure of the credentials file
type Credentials struct {
	ServiceAccountEmail   string `json:"STACKIT_SERVICE_ACCOUNT_EMAIL"`
	ServiceAccountToken   string `json:"STACKIT_SERVICE_ACCOUNT_TOKEN"`
	ServiceAccountKeyPath string `json:"STACKIT_SERVICE_ACCOUNT_KEY_PATH"`
	PrivateKeyPath        string `json:"STACKIT_PRIVATE_KEY_PATH"`
	ServiceAccountKey     string `json:"STACKIT_SERVICE_ACCOUNT_KEY"`
	PrivateKey            string `json:"STACKIT_PRIVATE_KEY"`
}

const (
	// nolint:gosec // G101 False positive: This is a constant path, not a credential
	credentialsFilePath                                = ".stackit/credentials.json"
	CredentialTypeToken                 CredentialType = "token"
	CredentialTypeServiceAccountKey     CredentialType = "service_account_key"
	CredentialTypeServiceAccountKeyPath CredentialType = "service_account_key_path"
	CredentialTypePrivateKey            CredentialType = "private_key"
	CredentialTypePrivateKeyPath        CredentialType = "private_key_path"
)

var UserHomeDir = os.UserHomeDir

// ReadCredentialsFile reads the credentials file from the specified path
func ReadCredentialsFile(path string) (*Credentials, error) {
	if path == "" {
		customPath, customPathSet := os.LookupEnv(EnvCredentialsPath)
		if !customPathSet || customPath == "" {
			path = credentialsFilePath
			home, err := UserHomeDir()
			if err != nil {
				return nil, fmt.Errorf("getting home directory: %w", err)
			}
			path = filepath.Join(home, path)
		} else {
			path = customPath
		}
	}

	warnOnInsecureCredentialsFile(path)

	credentialsRaw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	var credentials Credentials
	err = json.Unmarshal(credentialsRaw, &credentials)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling credentials: %w", err)
	}
	return &credentials, nil
}

// ReadCredential reads the specified credential type from Credentials
func ReadCredential(cred CredentialType, credentials *Credentials) (string, error) {
	var credentialValue string
	switch cred {
	case CredentialTypeToken:
		credentialValue = credentials.ServiceAccountToken
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("token is empty or not set")
		}
	case CredentialTypeServiceAccountKeyPath:
		credentialValue = credentials.ServiceAccountKeyPath
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("service account key path is empty or not set")
		}
	case CredentialTypePrivateKeyPath:
		credentialValue = credentials.PrivateKeyPath
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("private key path is empty or not set")
		}
	case CredentialTypeServiceAccountKey:
		credentialValue = credentials.ServiceAccountKey
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("service account key is empty or not set")
		}
	case CredentialTypePrivateKey:
		credentialValue = credentials.PrivateKey
		if credentialValue == "" {
			return credentialValue, fmt.Errorf("private key is empty or not set")
		}
	default:
		return "", fmt.Errorf("invalid credential type: %s", cred)
	}

	return credentialValue, nil
}

// warnOnInsecureCredentialsFile warns when the credentials file can be read by users other
// than its owner. The file may hold a service account token or a private key in clear text,
// and nothing creates it automatically, so its permissions are whatever the user's umask
// produced. This only warns: the format predates this package and is shared with the other
// STACKIT SDKs and the Terraform provider, so refusing to read it would break callers.
// Logging is opt-in, so this is silent unless SetLogger was called.
func warnOnInsecureCredentialsFile(path string) {
	if runtime.GOOS == "windows" {
		// Go synthesises permission bits on Windows, so they say nothing about access.
		return
	}
	info, err := os.Stat(path)
	if err != nil {
		// Let the read report the real problem.
		return
	}
	if perm := info.Mode().Perm(); perm&0o077 != 0 {
		WarnContext(context.Background(),
			"identity: credentials file is readable by other users",
			"path", path, "permissions", fmt.Sprintf("%#o", perm), "recommended", "0600")
	}
}
