package config

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/cliauth"
)

func init() {
	// Disable keyring access in tests to avoid macOS Keychain dialogs
	cliauth.SetSkipKeyring(true)
}

func TestWithCLIProviderAuth_Success(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Create test credentials
	createTestCredentialFile(t, tmpDir, map[string]string{
		"access_token":            "test-access-token",
		"refresh_token":           "test-refresh-token",
		"user_email":              "test@example.com",
		"session_expires_at_unix": fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()),
		"auth_flow_type":          "user_token",
	})

	cfg := &Configuration{}
	opt := WithCLIProviderAuth("")
	err := opt(cfg)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if cfg.CustomAuth == nil {
		t.Error("Expected CustomAuth to be set")
	}
}

func TestWithCLIProviderAuth_NoCredentials(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	cfg := &Configuration{}
	opt := WithCLIProviderAuth("")
	err := opt(cfg)

	if err == nil {
		t.Error("Expected error when no credentials exist")
	}
}

func TestWithCLIProviderAuth_WithProfile(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Create test credentials for a specific profile
	profile := "production"
	createTestCredentialFileForProfile(t, tmpDir, profile, map[string]string{
		"access_token":            "test-access-token",
		"refresh_token":           "test-refresh-token",
		"user_email":              "test@example.com",
		"session_expires_at_unix": fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()),
		"auth_flow_type":          "user_token",
	})

	cfg := &Configuration{}
	opt := WithCLIProviderAuth(profile)
	err := opt(cfg)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if cfg.CustomAuth == nil {
		t.Error("Expected CustomAuth to be set")
	}
}

func TestWithCLIBackgroundTokenRefresh_Success(t *testing.T) {
	ctx := context.Background()

	cfg := &Configuration{}
	opt := WithCLIBackgroundTokenRefresh(ctx)
	err := opt(cfg)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if cfg.BackgroundTokenRefreshContext != ctx {
		t.Error("Expected BackgroundTokenRefreshContext to be set")
	}
}

func TestWithCLIBackgroundTokenRefresh_NilContext(t *testing.T) {
	cfg := &Configuration{}
	opt := WithCLIBackgroundTokenRefresh(nil)
	err := opt(cfg)

	if err == nil {
		t.Error("Expected error for nil context")
	}
}

func TestWithCLIBackgroundTokenRefresh_Integration(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Create test credentials
	createTestCredentialFile(t, tmpDir, map[string]string{
		"access_token":            "test-access-token",
		"refresh_token":           "test-refresh-token",
		"user_email":              "test@example.com",
		"session_expires_at_unix": fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()),
		"auth_flow_type":          "user_token",
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := &Configuration{}

	// Apply both options
	err := WithCLIBackgroundTokenRefresh(ctx)(cfg)
	if err != nil {
		t.Fatalf("WithCLIBackgroundTokenRefresh() error = %v", err)
	}

	err = WithCLIProviderAuth("")(cfg)
	if err != nil {
		t.Fatalf("WithCLIProviderAuth() error = %v", err)
	}

	if cfg.CustomAuth == nil {
		t.Error("Expected CustomAuth to be set")
	}

	if cfg.BackgroundTokenRefreshContext != ctx {
		t.Error("Expected BackgroundTokenRefreshContext to be set")
	}
}

// Helper to create test credential file
func createTestCredentialFile(t *testing.T, homeDir string, data map[string]string) {
	jsonBytes, _ := json.Marshal(data)
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)

	filePath := filepath.Join(homeDir, ".stackit", "cli-api-auth-storage.txt")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	err := os.WriteFile(filePath, []byte(encoded), 0600)
	if err != nil {
		t.Fatalf("Failed to create test credential file: %v", err)
	}
}

// Helper to create test credential file for a specific profile
func createTestCredentialFileForProfile(t *testing.T, homeDir string, profile string, data map[string]string) {
	jsonBytes, _ := json.Marshal(data)
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)

	filePath := filepath.Join(homeDir, ".stackit", "profiles", profile, "cli-api-auth-storage.txt")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	err := os.WriteFile(filePath, []byte(encoded), 0600)
	if err != nil {
		t.Fatalf("Failed to create test credential file: %v", err)
	}
}
