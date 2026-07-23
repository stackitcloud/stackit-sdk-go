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
		if r.PostForm.Get("grant_type") != WifGrantType {
			t.Fatalf("unexpected grant type: %s", r.PostForm.Get("grant_type"))
		}
		if r.PostForm.Get("client_assertion_type") != WifClientAssertionType {
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

	provider, err := NewWorkloadIdentityFederationProvider(&WorkloadIdentityFederationProviderConfig{
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
	_, err := NewWorkloadIdentityFederationProvider(&WorkloadIdentityFederationProviderConfig{
		FederatedTokenFunction: func(context.Context) (string, error) { return "token", nil },
	})
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestWorkloadIdentityFederationScopesAndResourcesOverride(t *testing.T) {
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

	tests := []struct {
		name              string
		configScopes      []string
		configResources   []string
		optionScopes      []string
		optionResources   []string
		expectedScopes    string
		expectedResources []string
	}{
		{
			name:              "no config, no options",
			configScopes:      []string{},
			configResources:   []string{},
			optionScopes:      nil,
			optionResources:   nil,
			expectedScopes:    "",
			expectedResources: []string{},
		},
		{
			name:              "config scopes only",
			configScopes:      []string{"read", "write"},
			configResources:   []string{},
			optionScopes:      nil,
			optionResources:   nil,
			expectedScopes:    "read write",
			expectedResources: []string{},
		},
		{
			name:              "config resources only",
			configScopes:      []string{},
			configResources:   []string{"res1", "res2"},
			optionScopes:      nil,
			optionResources:   nil,
			expectedScopes:    "",
			expectedResources: []string{"res1", "res2"},
		},
		{
			name:              "option scopes override config",
			configScopes:      []string{"config-scope"},
			configResources:   []string{},
			optionScopes:      []string{"option-scope1", "option-scope2"},
			optionResources:   nil,
			expectedScopes:    "option-scope1 option-scope2",
			expectedResources: []string{},
		},
		{
			name:              "option resources override config",
			configScopes:      []string{},
			configResources:   []string{"config-res"},
			optionScopes:      nil,
			optionResources:   []string{"option-res1", "option-res2"},
			expectedScopes:    "",
			expectedResources: []string{"option-res1", "option-res2"},
		},
		{
			name:              "option scopes and resources override config",
			configScopes:      []string{"config-scope"},
			configResources:   []string{"config-res"},
			optionScopes:      []string{"opt-scope1", "opt-scope2"},
			optionResources:   []string{"opt-res1", "opt-res2"},
			expectedScopes:    "opt-scope1 opt-scope2",
			expectedResources: []string{"opt-res1", "opt-res2"},
		},
		{
			name:              "empty option scopes use config",
			configScopes:      []string{"config-scope"},
			configResources:   []string{},
			optionScopes:      []string{},
			optionResources:   nil,
			expectedScopes:    "config-scope",
			expectedResources: []string{},
		},
		{
			name:              "empty option resources use config",
			configScopes:      []string{},
			configResources:   []string{"config-res"},
			optionScopes:      nil,
			optionResources:   []string{},
			expectedScopes:    "",
			expectedResources: []string{"config-res"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var scopesReceived string
			var resourcesReceived []string

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

			provider, err := NewWorkloadIdentityFederationProvider(&WorkloadIdentityFederationProviderConfig{
				TokenURL:               server.URL,
				ClientID:               "service-account@sa.stackit.cloud",
				FederatedTokenFunction: func(context.Context) (string, error) { return assertion, nil },
				HTTPClient:             server.Client(),
				Scopes:                 tt.configScopes,
				Resources:              tt.configResources,
			})
			if err != nil {
				t.Fatalf("new provider: %v", err)
			}

			_, err = provider.Token(context.Background(), TokenRequestOptions{
				Scopes:    tt.optionScopes,
				Resources: tt.optionResources,
			})
			if err != nil {
				t.Fatalf("token: %v", err)
			}

			if scopesReceived != tt.expectedScopes {
				t.Fatalf("expected scopes %q, got %q", tt.expectedScopes, scopesReceived)
			}
			if len(resourcesReceived) != len(tt.expectedResources) {
				t.Fatalf("expected %d resources, got %d", len(tt.expectedResources), len(resourcesReceived))
			}
			for i, res := range resourcesReceived {
				if res != tt.expectedResources[i] {
					t.Fatalf("expected resource %q at index %d, got %q", tt.expectedResources[i], i, res)
				}
			}
		})
	}
}
