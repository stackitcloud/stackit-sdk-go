package identity

import (
	"context"
	"testing"
)

func TestStaticTokenProvider(t *testing.T) {
	//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"
	provider, err := NewStaticTokenProvider(&StaticTokenProviderConfig{Token: accessToken})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	if token.AccessToken != accessToken {
		t.Fatalf("expected %s, got %s", accessToken, token.AccessToken)
	}

	if token.ExpiresOn.IsZero() {
		t.Fatalf("expected non-zero expiration")
	}
}

func TestStaticTokenProviderEmpty(t *testing.T) {
	_, err := NewStaticTokenProvider(&StaticTokenProviderConfig{Token: ""})
	if err == nil {
		t.Fatalf("expected error for empty token")
	}
}
