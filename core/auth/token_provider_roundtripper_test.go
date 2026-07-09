package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/identity"
)

type tokenProviderStub struct {
	token identity.Token
	err   error
}

func (s tokenProviderStub) Token(context.Context, identity.TokenRequestOptions) (identity.Token, error) {
	if s.err != nil {
		return identity.Token{}, s.err
	}
	return s.token, nil
}

type roundTripperStub func(*http.Request) (*http.Response, error)

func (f roundTripperStub) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestTokenProviderRoundTripper(t *testing.T) {
	rt := newTokenProviderRoundTripper(tokenProviderStub{
		token: identity.Token{AccessToken: "token-value", RefreshOn: time.Now().Add(time.Hour)},
	}, roundTripperStub(func(req *http.Request) (*http.Response, error) {
		if req.Header.Get("Authorization") != "Bearer token-value" {
			t.Fatalf("unexpected authorization header: %s", req.Header.Get("Authorization"))
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("ok")), Header: make(http.Header)}, nil
	}))

	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatalf("create request: %v", err)
	}

	res, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatalf("round trip failed: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", res.StatusCode)
	}
}

func TestTokenProviderRoundTripperProviderError(t *testing.T) {
	rt := newTokenProviderRoundTripper(tokenProviderStub{err: fmt.Errorf("failed")}, roundTripperStub(func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("unexpected")
	}))

	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatalf("create request: %v", err)
	}

	_, err = rt.RoundTrip(req)
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "get access token from provider") {
		t.Fatalf("unexpected error: %v", err)
	}
}
