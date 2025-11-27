package cliauth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func init() {
	// Disable keyring access in tests to avoid macOS Keychain dialogs
	SetSkipKeyring(true)
}

func TestCLIProviderFlow_RoundTrip(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Create test credentials
	createTestCredentialFile(t, tmpDir, map[string]string{
		"access_token":             "test-access-token",
		"refresh_token":            "test-refresh-token",
		"user_email":               "test@example.com",
		"session_expires_at_unix":  fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()),
		"auth_flow_type":           "user_token",
	})

	// Create mock API server
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify Authorization header
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-access-token" {
			t.Errorf("Expected Authorization: Bearer test-access-token, got %s", auth)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}))
	defer apiServer.Close()

	// Create flow
	flow, err := NewCLIProviderFlow("", nil, nil)
	if err != nil {
		t.Fatalf("NewCLIProviderFlow() error = %v", err)
	}

	// Make request
	req, _ := http.NewRequest("GET", apiServer.URL, nil)
	resp, err := flow.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip() error = %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestCLIProviderFlow_AutomaticRefresh(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Create test credentials with expired token
	createTestCredentialFile(t, tmpDir, map[string]string{
		"access_token":             "expired-token",
		"refresh_token":            "valid-refresh-token",
		"user_email":               "test@example.com",
		"session_expires_at_unix":  fmt.Sprintf("%d", time.Now().Add(-1*time.Hour).Unix()), // Expired
		"auth_flow_type":           "user_token",
	})

	// Create mock OAuth2 server
	oauthServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Form.Get("refresh_token") != "valid-refresh-token" {
			t.Errorf("Expected refresh_token=valid-refresh-token")
		}

		response := RefreshTokenResponse{
			AccessToken:  "refreshed-access-token",
			RefreshToken: "new-refresh-token",
			ExpiresIn:    3600,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer oauthServer.Close()

	// Create mock API server
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "Bearer refreshed-access-token" {
			t.Errorf("Expected refreshed token in Authorization header, got %s", auth)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	// Create flow with custom HTTP client for token refresh
	httpClient := &http.Client{
		Transport: &testTransport{serverURL: oauthServer.URL},
	}
	flow, err := NewCLIProviderFlow("", nil, httpClient)
	if err != nil {
		t.Fatalf("NewCLIProviderFlow() error = %v", err)
	}

	// Make request (should trigger refresh)
	req, _ := http.NewRequest("GET", apiServer.URL, nil)
	resp, err := flow.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip() error = %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestCLIProviderFlow_ConcurrentRequests(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Create test credentials
	createTestCredentialFile(t, tmpDir, map[string]string{
		"access_token":             "test-token",
		"refresh_token":            "refresh-token",
		"user_email":               "test@example.com",
		"session_expires_at_unix":  fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()),
	})

	var requestCount int32
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requestCount, 1)
		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	flow, err := NewCLIProviderFlow("", nil, nil)
	if err != nil {
		t.Fatalf("NewCLIProviderFlow() error = %v", err)
	}

	// Make concurrent requests
	const numRequests = 10
	var wg sync.WaitGroup
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			req, _ := http.NewRequest("GET", apiServer.URL, nil)
			resp, err := flow.RoundTrip(req)
			if err != nil {
				t.Errorf("RoundTrip() error = %v", err)
				return
			}
			resp.Body.Close()
		}()
	}

	wg.Wait()

	if atomic.LoadInt32(&requestCount) != numRequests {
		t.Errorf("Expected %d requests, got %d", numRequests, requestCount)
	}
}

func TestCLIProviderFlow_BackgroundRefresh(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Create credentials that will expire soon
	expiryTime := time.Now().Add(2 * time.Second)
	createTestCredentialFile(t, tmpDir, map[string]string{
		"access_token":             "initial-token",
		"refresh_token":            "refresh-token",
		"user_email":               "test@example.com",
		"session_expires_at_unix":  fmt.Sprintf("%d", expiryTime.Unix()),
	})

	var refreshCount int32
	oauthServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&refreshCount, 1)
		response := RefreshTokenResponse{
			AccessToken: "refreshed-token",
			ExpiresIn:   3600,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer oauthServer.Close()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create flow with background refresh
	httpClient := &http.Client{
		Transport: &testTransport{serverURL: oauthServer.URL},
	}

	// Temporarily override refresh timing for faster test
	oldTimeStart := defaultTimeStartBeforeTokenExpiration
	oldTimeBetween := defaultTimeBetweenContextCheck
	defaultTimeStartBeforeTokenExpiration = 1 * time.Second
	defaultTimeBetweenContextCheck = 100 * time.Millisecond
	defer func() {
		defaultTimeStartBeforeTokenExpiration = oldTimeStart
		defaultTimeBetweenContextCheck = oldTimeBetween
	}()

	flow, err := NewCLIProviderFlowWithContext("", nil, httpClient, ctx)
	if err != nil {
		t.Fatalf("NewCLIProviderFlowWithContext() error = %v", err)
	}

	// Wait for background refresh to trigger
	time.Sleep(3 * time.Second)

	// Check if refresh was called
	count := atomic.LoadInt32(&refreshCount)
	if count < 1 {
		t.Errorf("Expected at least 1 background refresh, got %d", count)
	}

	// Verify token was updated
	creds := flow.GetCredentials()
	if creds.AccessToken != "refreshed-token" {
		t.Errorf("Expected refreshed-token, got %s", creds.AccessToken)
	}
}

func TestCLIProviderFlow_BackgroundRefreshCancellation(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	createTestCredentialFile(t, tmpDir, map[string]string{
		"access_token":             "test-token",
		"refresh_token":            "refresh-token",
		"user_email":               "test@example.com",
		"session_expires_at_unix":  fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()),
	})

	ctx, cancel := context.WithCancel(context.Background())

	flow, err := NewCLIProviderFlowWithContext("", nil, nil, ctx)
	if err != nil {
		t.Fatalf("NewCLIProviderFlowWithContext() error = %v", err)
	}

	// Cancel context immediately
	cancel()

	// Give goroutine time to terminate
	time.Sleep(100 * time.Millisecond)

	// Test should complete without hanging
	// The background goroutine should have terminated
	_ = flow
}

func TestGetCredentials(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	createTestCredentialFile(t, tmpDir, map[string]string{
		"access_token":             "test-token",
		"refresh_token":            "refresh-token",
		"user_email":               "test@example.com",
		"session_expires_at_unix":  fmt.Sprintf("%d", time.Now().Add(1*time.Hour).Unix()),
	})

	flow, err := NewCLIProviderFlow("", nil, nil)
	if err != nil {
		t.Fatalf("NewCLIProviderFlow() error = %v", err)
	}

	creds := flow.GetCredentials()
	if creds == nil {
		t.Fatal("GetCredentials() returned nil")
	}
	if creds.AccessToken != "test-token" {
		t.Errorf("AccessToken = %v, want test-token", creds.AccessToken)
	}

	// Verify we got a copy (modifying shouldn't affect flow)
	creds.AccessToken = "modified"
	credsCopy := flow.GetCredentials()
	if credsCopy.AccessToken != "test-token" {
		t.Error("GetCredentials() should return a copy, not a reference")
	}
}

// Helper to create test credential file
func createTestCredentialFile(t *testing.T, homeDir string, data map[string]string) {
	jsonBytes, _ := json.Marshal(data)
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)

	filePath := filepath.Join(homeDir, ".stackit", "cli-provider-auth-storage.txt")
	os.MkdirAll(filepath.Dir(filePath), 0755)
	err := os.WriteFile(filePath, []byte(encoded), 0600)
	if err != nil {
		t.Fatalf("Failed to create test credential file: %v", err)
	}
}
