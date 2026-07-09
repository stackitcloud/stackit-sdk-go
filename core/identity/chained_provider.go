package identity

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const chainedProviderErrorPrefix = "chained provider"

var _ TokenProvider = (*ChainedProvider)(nil)

// ChainedProvider is an ordered list of token providers.
type ChainedProvider struct {
	cond               *sync.Cond
	iterating          bool
	name               string
	retrySources       bool
	providers          []TokenProvider
	successfulProvider TokenProvider
}

// ChainedProviderOptions contains optional parameters for ChainedProvider.
type ChainedProviderOptions struct {
	// RetrySources configures how the chain uses its sources.
	// When true, Token always attempts every provider in order until one succeeds.
	// When false, after the first success the chain reuses that provider only.
	RetrySources bool
}

// NewChainedProvider creates a chain that tries providers in order until one succeeds.
func NewChainedProvider(providers ...TokenProvider) (*ChainedProvider, error) {
	return NewChainedProviderWithOptions(ChainedProviderOptions{}, providers...)
}

// NewChainedProviderWithOptions creates a chain using the provided options.
func NewChainedProviderWithOptions(options ChainedProviderOptions, providers ...TokenProvider) (*ChainedProvider, error) {
	if len(providers) == 0 {
		return nil, errors.New("providers must contain at least one TokenProvider")
	}
	for _, provider := range providers {
		if provider == nil {
			return nil, errors.New("providers cannot contain nil")
		}
	}
	cp := make([]TokenProvider, len(providers))
	copy(cp, providers)
	return &ChainedProvider{
		cond:         sync.NewCond(&sync.Mutex{}),
		name:         "ChainedProvider",
		retrySources: options.RetrySources,
		providers:    cp,
	}, nil
}

// NewChainWithOptions creates a chain using the provided options.
// Deprecated: use NewChainedProviderWithOptions.
func NewChainWithOptions(options ChainedProviderOptions, providers ...TokenProvider) (*ChainedProvider, error) {
	return NewChainedProviderWithOptions(options, providers...)
}

// Token returns the first token retrieved successfully from the configured providers.
func (c *ChainedProvider) Token(ctx context.Context, options TokenRequestOptions) (Token, error) {
	if !c.retrySources {
		// Ensure only one goroutine at a time iterates sources and sets successfulProvider.
		c.cond.L.Lock()
		for {
			if c.successfulProvider != nil {
				provider := c.successfulProvider
				c.cond.L.Unlock()
				return provider.Token(ctx, options)
			}
			if !c.iterating {
				c.iterating = true
				c.cond.L.Unlock()
				break
			}
			c.cond.Wait()
		}
	}

	errorsByProvider := make([]string, 0, len(c.providers))
	var (
		err                error
		token              Token
		successfulProvider TokenProvider
	)
	for index, provider := range c.providers {
		token, err = provider.Token(ctx, options)
		if err == nil {
			if token.AccessToken == "" {
				msg := fmt.Sprintf("%s (%d): returned empty access token", providerName(provider), index)
				errorsByProvider = append(errorsByProvider, msg)
				DebugContext(ctx, "identity: provider returned empty access token", "chain", c.name, "provider", providerName(provider), "index", index)
				err = fmt.Errorf("%s: %s", chainedProviderErrorPrefix, msg)
				continue
			}
			DebugContext(ctx, "identity: authenticated", "chain", c.name, "provider", providerName(provider))
			successfulProvider = provider
			break
		}
		msg := fmt.Sprintf("%s (%d): %v", providerName(provider), index, err)
		DebugContext(ctx, "identity: provider failed", "chain", c.name, "provider", providerName(provider), "index", index, "error", err)
		errorsByProvider = append(errorsByProvider, msg)
	}

	if c.iterating {
		c.cond.L.Lock()
		c.successfulProvider = successfulProvider
		c.iterating = false
		c.cond.L.Unlock()
		c.cond.Broadcast()
	}

	if err == nil && successfulProvider != nil {
		return token, nil
	}
	ErrorContext(ctx, "identity: all chain providers failed", "chain", c.name, "errors", strings.Join(errorsByProvider, "; "))
	return Token{}, fmt.Errorf("%s: all chain providers failed: %s", chainedProviderErrorPrefix, strings.Join(errorsByProvider, "; "))
}

// providerName returns a short human-readable name for a TokenProvider,
// stripping the package path and pointer sigil so the log is readable.
//
// For example: *identity.ServiceAccountKeyProvider → "ServiceAccountKeyProvider"
func providerName(p TokenProvider) string {
	raw := fmt.Sprintf("%T", p)
	// strip leading "*" and package prefix (e.g. "identity.")
	raw = strings.TrimPrefix(raw, "*")
	if dot := strings.LastIndex(raw, "."); dot >= 0 {
		raw = raw[dot+1:]
	}
	return raw
}
