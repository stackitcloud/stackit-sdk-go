package cliauth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/zalando/go-keyring"
)

// skipKeyring is set to true in test environments to avoid macOS Keychain dialogs
var skipKeyring = false

// SetSkipKeyring disables keyring access. This is useful in test environments
// to avoid macOS Keychain dialogs. Only file-based storage will be used.
func SetSkipKeyring(skip bool) {
	skipKeyring = skip
}

// ProviderCredentials represents OAuth credentials stored by the STACKIT CLI
// for API authentication (e.g., after running 'stackit auth api login').
//
// These credentials are managed by the STACKIT CLI and can be stored in either
// the system keyring or a fallback file location.
type ProviderCredentials struct {
	AccessToken      string
	RefreshToken     string
	Email            string
	SessionExpiresAt time.Time
	AuthFlowType     string
	TokenEndpoint    string

	// Internal fields for tracking storage location and source profile
	sourceProfile       string // Which profile these creds came from
	storageLocationUsed string // "keyring" or "file"
}

const (
	// Keyring service name prefix used by STACKIT CLI for API auth
	keyringServicePrefix = "stackit-cli-api"

	// Keyring account names for individual credential fields
	keyringAccessToken    = "access_token"
	keyringRefreshToken   = "refresh_token"
	keyringUserEmail      = "user_email"
	keyringSessionExpiry  = "session_expires_at_unix"
	keyringAuthFlowType   = "auth_flow_type"
	keyringTokenEndpoint  = "idp_token_endpoint"

	// Default profile name
	defaultProfile = "default"
)

// ReadCredentials reads API credentials from the STACKIT CLI storage.
// It first attempts to read from the system keyring, and falls back to reading
// from a Base64-encoded JSON file if the keyring is not available or fails.
//
// Profile resolution order:
//  1. profileOverride parameter (if non-empty)
//  2. STACKIT_CLI_PROFILE environment variable
//  3. ~/.config/stackit/cli-profile.txt file
//  4. "default"
//
// Returns an error if credentials cannot be found in either location.
func ReadCredentials(profileOverride string) (*ProviderCredentials, error) {
	// Determine active profile
	profile, err := getActiveProfile(profileOverride)
	if err != nil {
		return nil, fmt.Errorf("determine active profile: %w", err)
	}

	// Try keyring first (primary storage method) unless skipped
	if !skipKeyring {
		creds, err := readFromKeyring(profile)
		if err == nil {
			creds.sourceProfile = profile
			creds.storageLocationUsed = "keyring"
			return creds, nil
		}
	}

	// Fall back to Base64-encoded JSON file
	creds, fileErr := readFromFile(profile)
	if fileErr == nil {
		creds.sourceProfile = profile
		creds.storageLocationUsed = "file"
		return creds, nil
	}

	// File read failed
	if skipKeyring {
		return nil, fmt.Errorf("no CLI API credentials found in file (%v). Please run 'stackit auth api login'", fileErr)
	}

	// Both methods failed - return a combined error message
	return nil, fmt.Errorf("no CLI API credentials found in keyring or file (%v). Please run 'stackit auth api login'", fileErr)
}

// WriteCredentials writes API credentials back to storage.
// It writes to the same location where credentials were read from (keyring or file),
// as indicated by the StorageLocationUsed field.
//
// This function is typically called after refreshing an access token to persist
// the new token to storage.
func WriteCredentials(creds *ProviderCredentials) error {
	if creds == nil {
		return fmt.Errorf("credentials cannot be nil")
	}

	profile := creds.sourceProfile
	if profile == "" {
		profile = defaultProfile
	}

	// Try to write to keyring first (unless skipped)
	if !skipKeyring {
		if err := writeToKeyring(profile, creds); err == nil {
			return nil
		}
	}

	// Fall back to file
	return writeToFile(profile, creds)
}

// IsAuthenticated checks if valid CLI API credentials exist for the given profile.
// Returns true if credentials can be read successfully and contain an access token.
func IsAuthenticated(profileOverride string) bool {
	creds, err := ReadCredentials(profileOverride)
	if err != nil {
		return false
	}

	// Check if credentials exist and have an access token
	return creds != nil && creds.AccessToken != ""
}

// getActiveProfile determines which CLI profile to use.
// Priority: 1) explicit override, 2) STACKIT_CLI_PROFILE env var,
// 3) ~/.config/stackit/cli-profile.txt, 4) "default"
func getActiveProfile(profileOverride string) (string, error) {
	// 1. Explicit override from caller
	if profileOverride != "" {
		return profileOverride, nil
	}

	// 2. Environment variable
	if profile := os.Getenv("STACKIT_CLI_PROFILE"); profile != "" {
		return profile, nil
	}

	// 3. Profile config file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get home dir: %w", err)
	}

	profilePath := filepath.Join(homeDir, ".config", "stackit", "cli-profile.txt")
	data, err := os.ReadFile(profilePath)
	if err != nil {
		// File doesn't exist, use default profile
		if os.IsNotExist(err) {
			return defaultProfile, nil
		}
		return "", fmt.Errorf("read profile file: %w", err)
	}

	return strings.TrimSpace(string(data)), nil
}

// getKeyringServiceName returns the keyring service name for a profile
func getKeyringServiceName(profile string) string {
	if profile == defaultProfile {
		return keyringServicePrefix
	}
	return fmt.Sprintf("%s/%s", keyringServicePrefix, profile)
}

// getFilePath returns the storage file path for a profile
func getFilePath(profile string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get home dir: %w", err)
	}

	if profile == defaultProfile {
		return filepath.Join(homeDir, ".stackit", "cli-api-auth-storage.txt"), nil
	}
	return filepath.Join(homeDir, ".stackit", "profiles", profile, "cli-api-auth-storage.txt"), nil
}

// readFromKeyring reads API credentials from the system keyring.
// The CLI stores each field as a separate keyring entry.
func readFromKeyring(profile string) (*ProviderCredentials, error) {
	serviceName := getKeyringServiceName(profile)

	// Read access token (required)
	accessToken, err := keyring.Get(serviceName, keyringAccessToken)
	if err != nil {
		return nil, fmt.Errorf("get access_token: %w", err)
	}

	// Read refresh token (required)
	refreshToken, err := keyring.Get(serviceName, keyringRefreshToken)
	if err != nil {
		return nil, fmt.Errorf("get refresh_token: %w", err)
	}

	// Read user email (required)
	email, err := keyring.Get(serviceName, keyringUserEmail)
	if err != nil {
		return nil, fmt.Errorf("get user_email: %w", err)
	}

	creds := &ProviderCredentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Email:        email,
	}

	// Read expiry (optional)
	if expiryStr, err := keyring.Get(serviceName, keyringSessionExpiry); err == nil {
		if expiryUnix, err := strconv.ParseInt(expiryStr, 10, 64); err == nil {
			creds.SessionExpiresAt = time.Unix(expiryUnix, 0)
		}
	}

	// Read auth flow type (optional)
	if authFlow, err := keyring.Get(serviceName, keyringAuthFlowType); err == nil {
		creds.AuthFlowType = authFlow
	}

	// Read token endpoint (optional)
	if tokenEndpoint, err := keyring.Get(serviceName, keyringTokenEndpoint); err == nil {
		creds.TokenEndpoint = tokenEndpoint
	}

	return creds, nil
}

// readFromFile reads API credentials from the Base64-encoded JSON file fallback
func readFromFile(profile string) (*ProviderCredentials, error) {
	filePath, err := getFilePath(profile)
	if err != nil {
		return nil, fmt.Errorf("get file path: %w", err)
	}

	// Read Base64-encoded content
	contentEncoded, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s", filePath)
		}
		return nil, fmt.Errorf("read file: %w", err)
	}

	// Decode from Base64
	contentBytes, err := base64.StdEncoding.DecodeString(string(contentEncoded))
	if err != nil {
		return nil, fmt.Errorf("decode base64: %w", err)
	}

	// Parse JSON
	var data map[string]string
	if err := json.Unmarshal(contentBytes, &data); err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}

	// Extract required fields
	accessToken, ok := data["access_token"]
	if !ok || accessToken == "" {
		return nil, fmt.Errorf("access_token not found in file")
	}

	refreshToken, ok := data["refresh_token"]
	if !ok || refreshToken == "" {
		return nil, fmt.Errorf("refresh_token not found in file")
	}

	email, ok := data["user_email"]
	if !ok || email == "" {
		return nil, fmt.Errorf("user_email not found in file")
	}

	creds := &ProviderCredentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Email:        email,
	}

	// Parse expiry (optional)
	if expiryStr, ok := data["session_expires_at_unix"]; ok {
		if expiryUnix, err := strconv.ParseInt(expiryStr, 10, 64); err == nil {
			creds.SessionExpiresAt = time.Unix(expiryUnix, 0)
		}
	}

	// Auth flow type (optional)
	if authFlow, ok := data["auth_flow_type"]; ok {
		creds.AuthFlowType = authFlow
	}

	// Token endpoint (optional)
	if tokenEndpoint, ok := data["idp_token_endpoint"]; ok {
		creds.TokenEndpoint = tokenEndpoint
	}

	return creds, nil
}

// writeToKeyring writes credentials to the system keyring
func writeToKeyring(profile string, creds *ProviderCredentials) error {
	serviceName := getKeyringServiceName(profile)

	// Write required fields
	if err := keyring.Set(serviceName, keyringAccessToken, creds.AccessToken); err != nil {
		return fmt.Errorf("set access_token: %w", err)
	}

	if err := keyring.Set(serviceName, keyringRefreshToken, creds.RefreshToken); err != nil {
		return fmt.Errorf("set refresh_token: %w", err)
	}

	if err := keyring.Set(serviceName, keyringUserEmail, creds.Email); err != nil {
		return fmt.Errorf("set user_email: %w", err)
	}

	// Write optional fields
	if !creds.SessionExpiresAt.IsZero() {
		expiryStr := fmt.Sprintf("%d", creds.SessionExpiresAt.Unix())
		keyring.Set(serviceName, keyringSessionExpiry, expiryStr)
	}

	if creds.AuthFlowType != "" {
		keyring.Set(serviceName, keyringAuthFlowType, creds.AuthFlowType)
	}

	if creds.TokenEndpoint != "" {
		keyring.Set(serviceName, keyringTokenEndpoint, creds.TokenEndpoint)
	}

	return nil
}

// writeToFile writes credentials to the Base64-encoded JSON file
func writeToFile(profile string, creds *ProviderCredentials) error {
	filePath, err := getFilePath(profile)
	if err != nil {
		return fmt.Errorf("get file path: %w", err)
	}

	// Read existing file to preserve other fields
	var data map[string]string
	if existingContent, err := os.ReadFile(filePath); err == nil {
		if contentBytes, err := base64.StdEncoding.DecodeString(string(existingContent)); err == nil {
			json.Unmarshal(contentBytes, &data)
		}
	}

	if data == nil {
		data = make(map[string]string)
	}

	// Update credentials
	data["access_token"] = creds.AccessToken
	data["refresh_token"] = creds.RefreshToken
	data["user_email"] = creds.Email

	if !creds.SessionExpiresAt.IsZero() {
		data["session_expires_at_unix"] = fmt.Sprintf("%d", creds.SessionExpiresAt.Unix())
	}

	if creds.AuthFlowType != "" {
		data["auth_flow_type"] = creds.AuthFlowType
	}

	if creds.TokenEndpoint != "" {
		data["idp_token_endpoint"] = creds.TokenEndpoint
	}

	// Encode and write
	newContent, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(newContent)

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	return os.WriteFile(filePath, []byte(encoded), 0600)
}
