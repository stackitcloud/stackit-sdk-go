package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
	"github.com/stackitcloud/stackit-sdk-go/core/credentials"
)

type staticTP struct{ tok credentials.Token }

func (s *staticTP) Token(context.Context, credentials.TokenRequestOptions) (credentials.Token, error) {
	return s.tok, nil
}

func TestDefaultAuthWithTokenProvider(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer abc" {
			t.Errorf("unexpected Authorization header %q", got)
		}
	}))
	defer srv.Close()

	cfg := &config.Configuration{}
	if err := config.WithTokenProvider(&staticTP{credentials.Token{AccessToken: "abc"}})(cfg); err != nil {
		t.Fatal(err)
	}
	rt, err := DefaultAuth(cfg)
	if err != nil {
		t.Fatalf("DefaultAuth: %v", err)
	}
	resp, err := (&http.Client{Transport: rt}).Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
}

// TestSetupAuthExplicitOptionsWinOverTokenProvider pins that an option which explicitly
// chooses how to authenticate is not overridden by a token provider.
func TestSetupAuthExplicitOptionsWinOverTokenProvider(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer explicit-token" {
			t.Errorf("expected the explicit token to be used, got Authorization %q", got)
		}
	}))
	defer srv.Close()

	cfg := &config.Configuration{Token: "explicit-token"}
	if err := config.WithTokenProvider(&staticTP{credentials.Token{AccessToken: "from-provider"}})(cfg); err != nil {
		t.Fatal(err)
	}
	rt, err := SetupAuth(cfg)
	if err != nil {
		t.Fatalf("SetupAuth: %v", err)
	}
	resp, err := (&http.Client{Transport: rt}).Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
}
