package cliauth

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func init() {
	// Disable keyring access in tests to avoid macOS Keychain dialogs
	SetSkipKeyring(true)
}

func TestGetActiveProfile(t *testing.T) {
	tests := []struct {
		name             string
		profileOverride  string
		envVar           string
		fileContent      string
		expectedProfile  string
		shouldCreateFile bool
	}{
		{
			name:            "explicit override",
			profileOverride: "custom",
			envVar:          "env-profile",
			fileContent:     "file-profile",
			expectedProfile: "custom",
		},
		{
			name:            "environment variable",
			profileOverride: "",
			envVar:          "env-profile",
			fileContent:     "file-profile",
			expectedProfile: "env-profile",
		},
		{
			name:             "profile file",
			profileOverride:  "",
			envVar:           "",
			fileContent:      "file-profile",
			expectedProfile:  "file-profile",
			shouldCreateFile: true,
		},
		{
			name:             "default profile",
			profileOverride:  "",
			envVar:           "",
			fileContent:      "",
			expectedProfile:  "default",
			shouldCreateFile: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup temp directory and HOME
			tmpDir := t.TempDir()
			os.Setenv("HOME", tmpDir)
			defer os.Unsetenv("HOME")

			// Setup environment variable
			if tt.envVar != "" {
				os.Setenv("STACKIT_CLI_PROFILE", tt.envVar)
				defer os.Unsetenv("STACKIT_CLI_PROFILE")
			}

			// Create profile file if needed
			if tt.shouldCreateFile {
				profileDir := filepath.Join(tmpDir, ".config", "stackit")
				os.MkdirAll(profileDir, 0755)
				profileFile := filepath.Join(profileDir, "cli-profile.txt")
				os.WriteFile(profileFile, []byte(tt.fileContent), 0600)
			}

			// Test
			profile, err := getActiveProfile(tt.profileOverride)
			if err != nil {
				t.Fatalf("getActiveProfile() error = %v", err)
			}

			if profile != tt.expectedProfile {
				t.Errorf("getActiveProfile() = %v, want %v", profile, tt.expectedProfile)
			}
		})
	}
}

func TestGetKeyringServiceName(t *testing.T) {
	tests := []struct {
		profile     string
		expected    string
	}{
		{"default", "stackit-cli-provider"},
		{"production", "stackit-cli-provider/production"},
		{"dev", "stackit-cli-provider/dev"},
	}

	for _, tt := range tests {
		t.Run(tt.profile, func(t *testing.T) {
			result := getKeyringServiceName(tt.profile)
			if result != tt.expected {
				t.Errorf("getKeyringServiceName(%s) = %v, want %v", tt.profile, result, tt.expected)
			}
		})
	}
}

func TestGetFilePath(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	tests := []struct {
		profile  string
		expected string
	}{
		{"default", filepath.Join(tmpDir, ".stackit", "cli-provider-auth-storage.txt")},
		{"production", filepath.Join(tmpDir, ".stackit", "profiles", "production", "cli-provider-auth-storage.txt")},
	}

	for _, tt := range tests {
		t.Run(tt.profile, func(t *testing.T) {
			result, err := getFilePath(tt.profile)
			if err != nil {
				t.Fatalf("getFilePath() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("getFilePath(%s) = %v, want %v", tt.profile, result, tt.expected)
			}
		})
	}
}

func TestReadFromFile(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Create test credentials
	testCreds := map[string]string{
		"access_token":             "test-access-token",
		"refresh_token":            "test-refresh-token",
		"user_email":               "test@example.com",
		"session_expires_at_unix":  "1735689600",
		"auth_flow_type":           "user_token",
	}

	jsonBytes, _ := json.Marshal(testCreds)
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)

	// Write to file
	filePath := filepath.Join(tmpDir, ".stackit", "cli-provider-auth-storage.txt")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	os.WriteFile(filePath, []byte(encoded), 0600)

	// Test reading
	creds, err := readFromFile("default")
	if err != nil {
		t.Fatalf("readFromFile() error = %v", err)
	}

	if creds.AccessToken != testCreds["access_token"] {
		t.Errorf("AccessToken = %v, want %v", creds.AccessToken, testCreds["access_token"])
	}
	if creds.RefreshToken != testCreds["refresh_token"] {
		t.Errorf("RefreshToken = %v, want %v", creds.RefreshToken, testCreds["refresh_token"])
	}
	if creds.Email != testCreds["user_email"] {
		t.Errorf("Email = %v, want %v", creds.Email, testCreds["user_email"])
	}
	if creds.AuthFlowType != testCreds["auth_flow_type"] {
		t.Errorf("AuthFlowType = %v, want %v", creds.AuthFlowType, testCreds["auth_flow_type"])
	}
}

func TestReadFromFile_MissingFields(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name     string
		data     map[string]string
		wantErr  bool
	}{
		{
			name: "missing access_token",
			data: map[string]string{
				"refresh_token": "test-refresh",
				"user_email":    "test@example.com",
			},
			wantErr: true,
		},
		{
			name: "missing refresh_token",
			data: map[string]string{
				"access_token": "test-access",
				"user_email":   "test@example.com",
			},
			wantErr: true,
		},
		{
			name: "missing user_email",
			data: map[string]string{
				"access_token":  "test-access",
				"refresh_token": "test-refresh",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create separate temp dir for each test
			testDir := filepath.Join(tmpDir, tt.name)
			os.Setenv("HOME", testDir)
			defer os.Unsetenv("HOME")

			// Write test file
			jsonBytes, _ := json.Marshal(tt.data)
			encoded := base64.StdEncoding.EncodeToString(jsonBytes)

			filePath := filepath.Join(testDir, ".stackit", "cli-provider-auth-storage.txt")
			os.MkdirAll(filepath.Dir(filePath), 0755)
			os.WriteFile(filePath, []byte(encoded), 0600)

			_, err := readFromFile("default")
			if (err != nil) != tt.wantErr {
				t.Errorf("readFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteToFile(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	testCreds := &ProviderCredentials{
		AccessToken:      "new-access-token",
		RefreshToken:     "new-refresh-token",
		Email:            "test@example.com",
		SessionExpiresAt: time.Unix(1735689600, 0),
		AuthFlowType:     "user_token",
		sourceProfile:    "default",
	}

	// Write credentials
	err := writeToFile("default", testCreds)
	if err != nil {
		t.Fatalf("writeToFile() error = %v", err)
	}

	// Verify file was created
	filePath, _ := getFilePath("default")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatal("credential file was not created")
	}

	// Verify file permissions
	fileInfo, _ := os.Stat(filePath)
	if fileInfo.Mode().Perm() != 0600 {
		t.Errorf("file permissions = %o, want 0600", fileInfo.Mode().Perm())
	}

	// Read back and verify
	readCreds, err := readFromFile("default")
	if err != nil {
		t.Fatalf("readFromFile() error = %v", err)
	}

	if readCreds.AccessToken != testCreds.AccessToken {
		t.Errorf("AccessToken = %v, want %v", readCreds.AccessToken, testCreds.AccessToken)
	}
	if readCreds.RefreshToken != testCreds.RefreshToken {
		t.Errorf("RefreshToken = %v, want %v", readCreds.RefreshToken, testCreds.RefreshToken)
	}
}

func TestIsAuthenticated(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Test with no credentials
	if IsAuthenticated("") {
		t.Error("IsAuthenticated() should return false when no credentials exist")
	}

	// Create valid credentials
	testCreds := map[string]string{
		"access_token":  "test-access-token",
		"refresh_token": "test-refresh-token",
		"user_email":    "test@example.com",
	}
	jsonBytes, _ := json.Marshal(testCreds)
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)

	filePath := filepath.Join(tmpDir, ".stackit", "cli-provider-auth-storage.txt")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	os.WriteFile(filePath, []byte(encoded), 0600)

	// Test with valid credentials
	if !IsAuthenticated("") {
		t.Error("IsAuthenticated() should return true when valid credentials exist")
	}
}
