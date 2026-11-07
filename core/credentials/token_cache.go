package credentials

import (
	"context"
	"errors"
	"net/url"
	"slices"
	"strings"
	"sync"
	"time"
)

// tokenCache stores one token per distinct (scopes, resources) request and collapses
// concurrent refreshes of the same request into a single fetch (singleflight).
// The zero value is ready to use.
//
// Concurrent refreshes are collapsed with the small singleflight below instead of
// golang.org/x/sync/singleflight, to keep the dependencies of core to a minimum: every
// service module inherits them.
//
// Reads vastly outnumber writes (every API call asks for a token, while a refresh only
// happens when the token is about to expire), so the lock is a RWMutex and the hit path
// takes only the read lock.
type tokenCache struct {
	mu       sync.RWMutex
	entries  map[string]Token
	inflight map[string]*tokenCall
}

type tokenCall struct {
	done  chan struct{}
	token Token
	err   error
}

// get returns the cached token for key while it is fresh, and otherwise calls fetch.
// Callers asking for the same key while a fetch is running wait for it instead of
// issuing their own request.
func (c *tokenCache) get(ctx context.Context, key string, leeway time.Duration, fetch func(context.Context) (Token, error)) (Token, error) {
	for {
		// Fast path: a cache hit only needs to read, so concurrent callers do not serialize.
		c.mu.RLock()
		token, ok := c.fresh(key, leeway)
		c.mu.RUnlock()
		if ok {
			return token, nil
		}

		c.mu.Lock()
		// Another caller may have stored a token between releasing the read lock and getting
		// the write lock.
		if token, ok := c.fresh(key, leeway); ok {
			c.mu.Unlock()
			return token, nil
		}

		call, running := c.inflight[key]
		if !running {
			call = &tokenCall{done: make(chan struct{})}
			if c.inflight == nil {
				c.inflight = make(map[string]*tokenCall)
			}
			c.inflight[key] = call
			c.mu.Unlock()
			return c.lead(ctx, key, call, fetch)
		}
		c.mu.Unlock()

		select {
		case <-call.done:
		case <-ctx.Done():
			return Token{}, ctx.Err()
		}
		if call.err == nil {
			return call.token, nil
		}
		// The leader failed because its own context ended, not because the token could not
		// be obtained. This caller's context is still alive, so try again.
		if ctx.Err() == nil && (errors.Is(call.err, context.Canceled) || errors.Is(call.err, context.DeadlineExceeded)) {
			continue
		}
		return Token{}, call.err
	}
}

// fresh returns the cached token for key if it is still worth handing out. Callers hold c.mu.
func (c *tokenCache) fresh(key string, leeway time.Duration) (Token, bool) {
	token, ok := c.entries[key]
	if !ok || token.AccessToken == "" || !token.RefreshOn.After(time.Now().Add(leeway)) {
		return Token{}, false
	}
	return token, true
}

func (c *tokenCache) lead(ctx context.Context, key string, call *tokenCall, fetch func(context.Context) (Token, error)) (Token, error) {
	// Release the waiters and the inflight slot even if fetch panics.
	call.err = errors.New("token fetch did not complete")
	defer func() {
		c.mu.Lock()
		delete(c.inflight, key)
		if call.err == nil {
			c.store(key, call.token)
		}
		c.mu.Unlock()
		close(call.done)
	}()

	call.token, call.err = fetch(ctx)
	return call.token, call.err
}

// store saves the token and drops entries that have expired, so a caller cycling through
// many scope combinations does not grow the cache without bound. Callers hold c.mu.
func (c *tokenCache) store(key string, token Token) { //nolint:gocritic // hugeParam: the token is kept by value, a pointer would only add a nil case
	if c.entries == nil {
		c.entries = make(map[string]Token)
	}
	now := time.Now()
	for k, t := range c.entries {
		if !t.ExpiresOn.IsZero() && t.ExpiresOn.Before(now) {
			delete(c.entries, k)
		}
	}
	c.entries[key] = token
}

// effectiveRequest returns the scopes and resources to request: those of the call when
// given, otherwise the provider defaults.
func effectiveRequest(opt TokenRequestOptions, defaultScopes, defaultResources []string) (scopes, resources []string) {
	scopes, resources = defaultScopes, defaultResources
	if len(opt.Scopes) > 0 {
		scopes = opt.Scopes
	}
	if len(opt.Resources) > 0 {
		resources = opt.Resources
	}
	return scopes, resources
}

// requestKey identifies a token request. Order and duplicates do not matter, as they do
// not change what the token is issued for.
func requestKey(scopes, resources []string) string {
	return strings.Join(normalized(scopes), " ") + "\x00" + strings.Join(normalized(resources), " ")
}

func normalized(values []string) []string {
	out := slices.Clone(values)
	slices.Sort(out)
	return slices.Compact(out)
}

// setScopesAndResources adds the scope and resource parameters of a token request.
func setScopesAndResources(body url.Values, scopes, resources []string) {
	if len(scopes) > 0 {
		body.Set("scope", strings.Join(scopes, " "))
	}
	for _, resource := range resources {
		body.Add("resource", resource)
	}
}
