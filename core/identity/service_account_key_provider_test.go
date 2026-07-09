package identity

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestServiceAccountKeyToken(t *testing.T) {
	privateKeyPEM, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}

	var requests int64
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   "subject",
	}).SignedString([]byte("test"))
	if err != nil {
		t.Fatalf("generate access token: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&requests, 1)
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		if r.PostForm.Get("grant_type") != serviceAccountJWTBearerGrantType {
			t.Fatalf("unexpected grant type: %s", r.PostForm.Get("grant_type"))
		}
		if r.PostForm.Get("assertion") == "" {
			t.Fatalf("expected assertion")
		}

		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": accessToken,
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	// Create service account key as JSON
	saKey := ServiceAccountJson{
		Credentials: &ServiceAccountKeyCredentials{
			Aud: server.URL,
			Iss: "service-account@sa.stackit.cloud",
			Kid: "kid",
			Sub: uuid.New(),
		},
	}
	saKeyJSON, err := json.Marshal(saKey)
	if err != nil {
		t.Fatalf("marshal service account key: %v", err)
	}

	provider, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{
		ServiceAccountKey: string(saKeyJSON),
		PrivateKey:        string(privateKeyPEM),
		TokenURL:          server.URL,
		HTTPClient:        server.Client(),
	})
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}

	first, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("first token: %v", err)
	}
	second, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("second token: %v", err)
	}
	if first.AccessToken != second.AccessToken {
		t.Fatalf("expected cache hit")
	}
	if atomic.LoadInt64(&requests) != 1 {
		t.Fatalf("expected one request, got %d", requests)
	}
}

func TestServiceAccountKeyRequiresKeyAndPrivateKey(t *testing.T) {
	_, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{})
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestServiceAccountKeyWithScopes(t *testing.T) {
	privateKeyPEM, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}

	var scopesReceived string
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   "subject",
	}).SignedString([]byte("test"))
	if err != nil {
		t.Fatalf("generate access token: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		scopesReceived = r.PostForm.Get("scope")

		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": accessToken,
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	saKey := ServiceAccountJson{
		Credentials: &ServiceAccountKeyCredentials{
			Aud: server.URL,
			Iss: "service-account@sa.stackit.cloud",
			Kid: "kid",
			Sub: uuid.New(),
		},
	}
	saKeyJSON, err := json.Marshal(saKey)
	if err != nil {
		t.Fatalf("marshal service account key: %v", err)
	}

	provider, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{
		ServiceAccountKey: string(saKeyJSON),
		PrivateKey:        string(privateKeyPEM),
		TokenURL:          server.URL,
		HTTPClient:        server.Client(),
		Scopes:            []string{"scope1", "scope2"},
	})
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}

	_, err = provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("token: %v", err)
	}

	expected := "scope1 scope2"
	if scopesReceived != expected {
		t.Fatalf("expected scopes %q, got %q", expected, scopesReceived)
	}
}

func TestServiceAccountKeyWithResources(t *testing.T) {
	privateKeyPEM, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}

	var resourcesReceived []string
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   "subject",
	}).SignedString([]byte("test"))
	if err != nil {
		t.Fatalf("generate access token: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		resourcesReceived = r.PostForm["resource"]

		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": accessToken,
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	saKey := ServiceAccountJson{
		Credentials: &ServiceAccountKeyCredentials{
			Aud: server.URL,
			Iss: "service-account@sa.stackit.cloud",
			Kid: "kid",
			Sub: uuid.New(),
		},
	}
	saKeyJSON, err := json.Marshal(saKey)
	if err != nil {
		t.Fatalf("marshal service account key: %v", err)
	}

	provider, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{
		ServiceAccountKey: string(saKeyJSON),
		PrivateKey:        string(privateKeyPEM),
		TokenURL:          server.URL,
		HTTPClient:        server.Client(),
		Resources:         []string{"resource1", "resource2", "resource3"},
	})
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}

	_, err = provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("token: %v", err)
	}

	expected := []string{"resource1", "resource2", "resource3"}
	if len(resourcesReceived) != len(expected) {
		t.Fatalf("expected %d resources, got %d", len(expected), len(resourcesReceived))
	}
	for i, res := range resourcesReceived {
		if res != expected[i] {
			t.Fatalf("expected resource %q at index %d, got %q", expected[i], i, res)
		}
	}
}

func TestServiceAccountKeyWithScopesAndResources(t *testing.T) {
	privateKeyPEM, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}

	var scopesReceived string
	var resourcesReceived []string
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   "subject",
	}).SignedString([]byte("test"))
	if err != nil {
		t.Fatalf("generate access token: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		scopesReceived = r.PostForm.Get("scope")
		resourcesReceived = r.PostForm["resource"]

		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": accessToken,
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	saKey := ServiceAccountJson{
		Credentials: &ServiceAccountKeyCredentials{
			Aud: server.URL,
			Iss: "service-account@sa.stackit.cloud",
			Kid: "kid",
			Sub: uuid.New(),
		},
	}
	saKeyJSON, err := json.Marshal(saKey)
	if err != nil {
		t.Fatalf("marshal service account key: %v", err)
	}

	provider, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{
		ServiceAccountKey: string(saKeyJSON),
		PrivateKey:        string(privateKeyPEM),
		TokenURL:          server.URL,
		HTTPClient:        server.Client(),
		Scopes:            []string{"read", "write"},
		Resources:         []string{"resourceA", "resourceB"},
	})
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}

	_, err = provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("token: %v", err)
	}

	expectedScopes := "read write"
	if scopesReceived != expectedScopes {
		t.Fatalf("expected scopes %q, got %q", expectedScopes, scopesReceived)
	}

	expectedResources := []string{"resourceA", "resourceB"}
	if len(resourcesReceived) != len(expectedResources) {
		t.Fatalf("expected %d resources, got %d", len(expectedResources), len(resourcesReceived))
	}
	for i, res := range resourcesReceived {
		if res != expectedResources[i] {
			t.Fatalf("expected resource %q at index %d, got %q", expectedResources[i], i, res)
		}
	}
}

func TestServiceAccountKeyWithEmptyScopesAndResources(t *testing.T) {
	privateKeyPEM, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}

	var scopesReceived string
	var resourcesReceived []string
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   "subject",
	}).SignedString([]byte("test"))
	if err != nil {
		t.Fatalf("generate access token: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		scopesReceived = r.PostForm.Get("scope")
		resourcesReceived = r.PostForm["resource"]

		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": accessToken,
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	saKey := ServiceAccountJson{
		Credentials: &ServiceAccountKeyCredentials{
			Aud: server.URL,
			Iss: "service-account@sa.stackit.cloud",
			Kid: "kid",
			Sub: uuid.New(),
		},
	}
	saKeyJSON, err := json.Marshal(saKey)
	if err != nil {
		t.Fatalf("marshal service account key: %v", err)
	}

	provider, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{
		ServiceAccountKey: string(saKeyJSON),
		PrivateKey:        string(privateKeyPEM),
		TokenURL:          server.URL,
		HTTPClient:        server.Client(),
		Scopes:            []string{},
		Resources:         []string{},
	})
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}

	_, err = provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("token: %v", err)
	}

	if scopesReceived != "" {
		t.Fatalf("expected empty scopes, got %q", scopesReceived)
	}
	if len(resourcesReceived) != 0 {
		t.Fatalf("expected no resources, got %d", len(resourcesReceived))
	}
}

func generateRSAPrivateKeyPEM() ([]byte, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}), nil
}
