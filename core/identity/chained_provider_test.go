package identity

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type staticProvider struct {
	token Token
	err   error
}

func (s *staticProvider) Token(context.Context, TokenRequestOptions) (Token, error) {
	if s.err != nil {
		return Token{}, s.err
	}
	return s.token, nil
}

func TestChainToken(t *testing.T) {
	expected := Token{AccessToken: "token", RefreshOn: time.Now().Add(time.Hour)}
	chain, err := NewChainedProvider(
		&staticProvider{err: fmt.Errorf("first failed")},
		&staticProvider{token: expected},
	)
	if err != nil {
		t.Fatalf("new chain: %v", err)
	}

	token, err := chain.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token.AccessToken != expected.AccessToken {
		t.Fatalf("unexpected access token: %s", token.AccessToken)
	}
}

func TestChainTokenAllFailures(t *testing.T) {
	chain, err := NewChainedProvider(
		&staticProvider{err: fmt.Errorf("first failed")},
		&staticProvider{err: fmt.Errorf("second failed")},
	)
	if err != nil {
		t.Fatalf("new chain: %v", err)
	}

	_, err = chain.Token(context.Background(), TokenRequestOptions{})
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "all chain providers failed") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestChainNoProviders(t *testing.T) {
	_, err := NewChainedProvider()
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "at least one") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestChainCachesSuccessfulProviderByDefault(t *testing.T) {
	var firstCalls int32
	var secondCalls int32

	first := &staticProvider{err: fmt.Errorf("first failed")}
	second := &staticProvider{token: Token{AccessToken: "token", RefreshOn: time.Now().Add(time.Hour)}}

	countingFirst := countingProvider{inner: first, count: &firstCalls}
	countingSecond := countingProvider{inner: second, count: &secondCalls}

	chain, err := NewChainedProvider(countingFirst, countingSecond)
	if err != nil {
		t.Fatalf("new chain: %v", err)
	}

	_, err = chain.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("first token call: %v", err)
	}
	_, err = chain.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("second token call: %v", err)
	}

	if atomic.LoadInt32(&firstCalls) != 1 {
		t.Fatalf("expected first provider to be called once, got %d", firstCalls)
	}
	if atomic.LoadInt32(&secondCalls) != 2 {
		t.Fatalf("expected successful provider to be called twice, got %d", secondCalls)
	}
}

func TestChainRetrySourcesOption(t *testing.T) {
	var firstCalls int32
	var secondCalls int32

	first := &staticProvider{err: fmt.Errorf("first failed")}
	second := &staticProvider{token: Token{AccessToken: "token", RefreshOn: time.Now().Add(time.Hour)}}

	countingFirst := countingProvider{inner: first, count: &firstCalls}
	countingSecond := countingProvider{inner: second, count: &secondCalls}

	chain, err := NewChainedProviderWithOptions(ChainedProviderOptions{RetrySources: true}, countingFirst, countingSecond)
	if err != nil {
		t.Fatalf("new chain: %v", err)
	}

	_, err = chain.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("first token call: %v", err)
	}
	_, err = chain.Token(context.Background(), TokenRequestOptions{})
	if err != nil {
		t.Fatalf("second token call: %v", err)
	}

	if atomic.LoadInt32(&firstCalls) != 2 {
		t.Fatalf("expected first provider to be called twice, got %d", firstCalls)
	}
	if atomic.LoadInt32(&secondCalls) != 2 {
		t.Fatalf("expected second provider to be called twice, got %d", secondCalls)
	}
}

type countingProvider struct {
	inner TokenProvider
	count *int32
}

func (c countingProvider) Token(ctx context.Context, options TokenRequestOptions) (Token, error) {
	atomic.AddInt32(c.count, 1)
	return c.inner.Token(ctx, options)
}
