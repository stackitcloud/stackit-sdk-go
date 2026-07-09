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

func TestWorkloadIdentityFederationToken(t *testing.T) {
	assertion, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   "assertion",
	}).SignedString([]byte("assertion-secret"))
	if err != nil {
		t.Fatalf("generate assertion token: %v", err)
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		Subject:   "subject",
	}).SignedString([]byte("access-secret"))
	if err != nil {
		t.Fatalf("generate access token: %v", err)
	}

	var requests int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&requests, 1)
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		if r.PostForm.Get("grant_type") != workloadIdentityGrantType {
			t.Fatalf("unexpected grant type: %s", r.PostForm.Get("grant_type"))
		}
		if r.PostForm.Get("client_assertion_type") != workloadIdentityClientAssertionType {
			t.Fatalf("unexpected assertion type: %s", r.PostForm.Get("client_assertion_type"))
		}
		if r.PostForm.Get("client_assertion") == "" {
			t.Fatalf("expected client assertion")
		}
		if r.PostForm.Get("client_id") != "service-account@sa.stackit.cloud" {
			t.Fatalf("unexpected client id: %s", r.PostForm.Get("client_id"))
		}

		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": accessToken,
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	provider, err := NewWorkloadIdentityFederationProvider(WorkloadIdentityFederationProviderConfig{
		TokenURL:               server.URL,
		ClientID:               "service-account@sa.stackit.cloud",
		FederatedTokenFunction: func(context.Context) (string, error) { return assertion, nil },
		HTTPClient:             server.Client(),
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

func TestWorkloadIdentityFederationRequiresClientID(t *testing.T) {
	_, err := NewWorkloadIdentityFederationProvider(WorkloadIdentityFederationProviderConfig{
		FederatedTokenFunction: func(context.Context) (string, error) { return "token", nil },
	})
	if err == nil {
		t.Fatalf("expected error")
	}
}
