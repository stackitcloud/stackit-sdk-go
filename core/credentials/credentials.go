package credentials

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// credentialsFile represents the structure of the credentials file
type credentialsFile struct {
	ServiceAccountEmail   string `json:"STACKIT_SERVICE_ACCOUNT_EMAIL"`
	ServiceAccountToken   string `json:"STACKIT_SERVICE_ACCOUNT_TOKEN"`
	ServiceAccountKeyPath string `json:"STACKIT_SERVICE_ACCOUNT_KEY_PATH"`
	PrivateKeyPath        string `json:"STACKIT_PRIVATE_KEY_PATH"`
	ServiceAccountKey     string `json:"STACKIT_SERVICE_ACCOUNT_KEY"`
	PrivateKey            string `json:"STACKIT_PRIVATE_KEY"`
}

const (
	// nolint:gosec // G101 False positive: This is a constant path, not a credential
	credentialsFilePath = ".stackit/credentials.json"
)

// readCredentialsFile reads the credentials file from the specified path
func readCredentialsFile(path string) (*credentialsFile, error) {
	if path == "" {
		customPath, customPathSet := os.LookupEnv(envCredentialsPath)
		if !customPathSet || customPath == "" {
			path = credentialsFilePath
			home, err := os.UserHomeDir()
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

	var credentials credentialsFile
	err = json.Unmarshal(credentialsRaw, &credentials)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling credentials: %w", err)
	}
	return &credentials, nil
}

// warnOnInsecureCredentialsFile warns when the credentials file can be read by users other
// than its owner. The file may hold a service account token or a private key in clear text,
// and nothing creates it automatically, so its permissions are whatever the user's umask
// produced. This only warns: the format predates this package and is shared with the other
// STACKIT SDKs and the Terraform provider, so refusing to read it would break callers.
// Logging is opt-in, so this is silent unless SetLogger was called.
func warnOnInsecureCredentialsFile(path string) {
	if runtime.GOOS == "windows" {
		// Go synthesizes permission bits on Windows, so they say nothing about access.
		return
	}
	info, err := os.Stat(path)
	if err != nil {
		// Let the read report the real problem.
		return
	}
	if perm := info.Mode().Perm(); perm&0o077 != 0 {
		warnContext(context.Background(),
			"credentials: credentials file is readable by other users",
			"path", path, "permissions", fmt.Sprintf("%#o", perm), "recommended", "0600")
	}
}
