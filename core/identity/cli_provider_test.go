package identity

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
const testCLIAccessToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"

// fakeCLI returns an executor that records its invocations and replays the given output.
func fakeCLI(output string, err error, calls *int64) cliExecutor {
	return func(_ context.Context, _ string, args ...string) ([]byte, error) {
		if calls != nil {
			atomic.AddInt64(calls, 1)
		}
		if err != nil {
			return nil, err
		}
		// Guard the contract with the CLI: the arguments are part of it.
		want := "auth get-access-token --output-format json"
		if got := strings.Join(args, " "); got != want {
			return nil, fmt.Errorf("unexpected arguments: %q", got)
		}
		return []byte(output), nil
	}
}

func TestCLIProviderToken(t *testing.T) {
	var calls int64
	provider, err := NewCLIProvider(&CLIProviderConfig{
		exec: fakeCLI(fmt.Sprintf(`{"access_token": %q}`, testCLIAccessToken), nil, &calls),
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.AccessToken != testCLIAccessToken {
		t.Fatalf("expected %s, got %s", testCLIAccessToken, token.AccessToken)
	}
	if token.ExpiresOn.IsZero() {
		t.Fatalf("expected non-zero expiration")
	}

	// A second call must be served from the cache, not by spawning the CLI again.
	if _, err := provider.Token(context.Background(), TokenRequestOptions{}); err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if got := atomic.LoadInt64(&calls); got != 1 {
		t.Fatalf("expected the CLI to be invoked once, got %d", got)
	}
}

// TestCLIProviderNotInstalled covers the environment this provider is most often in:
// a container or CI runner with no STACKIT CLI. Construction must fail so a chain skips it.
func TestCLIProviderNotInstalled(t *testing.T) {
	t.Setenv("PATH", t.TempDir())

	_, err := NewCLIProvider(&CLIProviderConfig{Command: "stackit-does-not-exist"})
	if err == nil {
		t.Fatalf("expected an error when the CLI is not installed")
	}
	if !strings.Contains(err.Error(), cliProviderErrorPrefix) {
		t.Fatalf("expected a cli provider error, got: %v", err)
	}
}

// TestCLIProviderSessionExpired covers a CLI that is installed but not logged in: the
// provider reports the CLI's own message and never starts an interactive login.
func TestCLIProviderSessionExpired(t *testing.T) {
	provider, err := NewCLIProvider(&CLIProviderConfig{
		exec: fakeCLI("", errors.New("exit status 1: session expired, please run \"stackit auth login\""), nil),
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	_, err = provider.Token(context.Background(), TokenRequestOptions{})
	if err == nil {
		t.Fatalf("expected an error for an expired session")
	}
	if !strings.Contains(err.Error(), "session expired") {
		t.Fatalf("expected the CLI message to be surfaced, got: %v", err)
	}
}

func TestCLIProviderInvalidOutput(t *testing.T) {
	provider, err := NewCLIProvider(&CLIProviderConfig{
		exec: fakeCLI("not json", nil, nil),
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	if _, err := provider.Token(context.Background(), TokenRequestOptions{}); err == nil {
		t.Fatalf("expected an error for unparsable output")
	}
}

func TestCLIProviderEmptyToken(t *testing.T) {
	provider, err := NewCLIProvider(&CLIProviderConfig{
		exec: fakeCLI(`{"access_token": ""}`, nil, nil),
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	if _, err := provider.Token(context.Background(), TokenRequestOptions{}); err == nil {
		t.Fatalf("expected an error for an empty access token")
	}
}

// TestCLIProviderHonoursTimeout checks that a hanging CLI cannot block a caller forever.
func TestCLIProviderHonoursTimeout(t *testing.T) {
	provider, err := NewCLIProvider(&CLIProviderConfig{
		CommandTimeout: time.Millisecond,
		exec: func(ctx context.Context, _ string, _ ...string) ([]byte, error) {
			<-ctx.Done()
			return nil, ctx.Err()
		},
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	if _, err := provider.Token(context.Background(), TokenRequestOptions{}); err == nil {
		t.Fatalf("expected a timeout error")
	}
}

// TestCLIProviderTokenTypeDefaults covers a source that reports no token_type.
func TestCLIProviderTokenTypeDefaults(t *testing.T) {
	provider, err := NewCLIProvider(&CLIProviderConfig{
		exec: fakeCLI(fmt.Sprintf(`{"access_token": %q}`, testCLIAccessToken), nil, nil),
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.TokenType != DefaultTokenType {
		t.Fatalf("expected %s, got %s", DefaultTokenType, token.TokenType)
	}
}

// TestCLIProviderTokenTypeFromResponse is the case that motivates the field: a token that
// is not a Bearer token, such as RFC 9449 (DPoP), must be reported as issued.
func TestCLIProviderTokenTypeFromResponse(t *testing.T) {
	provider, err := NewCLIProvider(&CLIProviderConfig{
		exec: fakeCLI(fmt.Sprintf(`{"access_token": %q, "token_type": "DPoP"}`, testCLIAccessToken), nil, nil),
	})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	token, err := provider.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}
	if token.TokenType != "DPoP" {
		t.Fatalf("expected DPoP, got %s", token.TokenType)
	}
}
