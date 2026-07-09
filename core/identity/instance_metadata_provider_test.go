package identity

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestInstanceMetadataTokenCachesResponse(t *testing.T) {
	var requests int64
	tokenJWT, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		Subject:   "subject",
	}).SignedString([]byte("test"))
	if err != nil {
		t.Fatalf("create token: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&requests, 1)
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method %s", r.Method)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"token":      tokenJWT,
			"validUntil": time.Now().Add(time.Hour).UTC().Format(time.RFC3339Nano),
		})
	}))
	defer server.Close()

	provider, err := NewInstanceMetadataProvider(InstanceMetadataProviderConfig{
		ServiceAccountEmail: "test@sa.stackit.cloud",
		HTTPClient:          server.Client(),
	})
	if err != nil {
		t.Fatalf("create provider: %v", err)
	}
	provider.endpointTemplate = server.URL + "/stackit/v1/service-accounts/<SERVICEACCOUNTMAIL>/token"

	first, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("first token call: %v", err)
	}
	second, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("second token call: %v", err)
	}
	if first.AccessToken != second.AccessToken {
		t.Fatalf("expected same token on cache hit")
	}
	if atomic.LoadInt64(&requests) != 1 {
		t.Fatalf("expected one request, got %d", requests)
	}
}

func TestInstanceMetadataTokenRequiresServiceAccountEmail(t *testing.T) {
	_, err := NewInstanceMetadataProvider(InstanceMetadataProviderConfig{})
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestInstanceMetadataTokenInvalidValidUntil(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"token":      "token",
			"validUntil": "invalid",
		})
	}))
	defer server.Close()

	provider, err := NewInstanceMetadataProvider(InstanceMetadataProviderConfig{
		ServiceAccountEmail: "test@sa.stackit.cloud",
		HTTPClient:          server.Client(),
	})
	if err != nil {
		t.Fatalf("create provider: %v", err)
	}
	provider.endpointTemplate = server.URL + "/stackit/v1/service-accounts/<SERVICEACCOUNTMAIL>/token"

	_, err = provider.Token(context.Background(), TokenRequestOptions{})
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestInstanceMetadataTokenHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	provider, err := NewInstanceMetadataProvider(InstanceMetadataProviderConfig{
		ServiceAccountEmail: "test@sa.stackit.cloud",
		HTTPClient:          server.Client(),
	})
	if err != nil {
		t.Fatalf("create provider: %v", err)
	}
	provider.endpointTemplate = server.URL + "/stackit/v1/service-accounts/<SERVICEACCOUNTMAIL>/token"

	_, err = provider.Token(context.Background(), TokenRequestOptions{})
	if err == nil {
		t.Fatalf("expected error")
	}
}
