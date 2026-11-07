package credentials

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"slices"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func newCacheTestProvider(t *testing.T, scopes, resources []string) (*ServiceAccountKeyProvider, *atomic.Int32) {
	t.Helper()
	privateKeyPEM, err := generateRSAPrivateKeyPEM()
	if err != nil {
		t.Fatalf("generate private key: %v", err)
	}

	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		if err := r.ParseForm(); err != nil {
			t.Errorf("parse form: %v", err)
		}
		time.Sleep(20 * time.Millisecond) // give concurrent callers time to pile up
		// The token records what it was issued for, so a test can check it got the right one.
		// As in RFC 8707, every requested resource becomes an element of the audience.
		claims := jwt.MapClaims{
			"exp":   time.Now().Add(time.Hour).Unix(),
			"scope": r.PostForm.Get("scope"),
		}
		if resources := r.PostForm["resource"]; len(resources) > 0 {
			claims["aud"] = resources
		}
		accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("test"))
		if err != nil {
			t.Errorf("sign: %v", err)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": accessToken, "expires_in": 3600})
	}))
	t.Cleanup(server.Close)

	saKeyJSON, err := json.Marshal(serviceAccountJSON{Credentials: &serviceAccountKeyCredentials{
		Aud: server.URL, Iss: "sa@sa.stackit.cloud", Kid: "kid", Sub: uuid.New(),
	}})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	provider, err := NewServiceAccountKeyProvider(&ServiceAccountKeyProviderConfig{
		ServiceAccountKey: string(saKeyJSON),
		PrivateKey:        string(privateKeyPEM),
		TokenURL:          server.URL,
		HTTPClient:        server.Client(),
		Scopes:            scopes,
		Resources:         resources,
	})
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}
	return provider, &requests
}

// issuedFor returns the scope and the audience (the requested resources) the test server
// issued the token for.
func issuedFor(t *testing.T, accessToken string) (scope string, resources []string) {
	t.Helper()
	claims := jwt.MapClaims{}
	if _, _, err := jwt.NewParser().ParseUnverified(accessToken, claims); err != nil {
		t.Fatalf("parse token: %v", err)
	}
	scope, _ = claims["scope"].(string)
	audience, err := claims.GetAudience()
	if err != nil {
		t.Fatalf("read audience: %v", err)
	}
	return scope, audience
}

func requireIssuedFor(t *testing.T, accessToken, wantScope string, wantResources ...string) {
	t.Helper()
	scope, resources := issuedFor(t, accessToken)
	if scope != wantScope || !slices.Equal(resources, wantResources) {
		t.Fatalf("token issued for scope %q and resources %v, want scope %q and resources %v", scope, resources, wantScope, wantResources)
	}
}

func TestProviderCachesOneTokenPerScopesAndResources(t *testing.T) {
	provider, requests := newCacheTestProvider(t, nil, nil)
	ctx := context.Background()

	get := func(opt TokenRequestOptions) string {
		t.Helper()
		token, err := provider.Token(ctx, opt)
		if err != nil {
			t.Fatalf("token: %v", err)
		}
		return token.AccessToken
	}

	// Each combination gets a token issued for exactly what was asked...
	requireIssuedFor(t, get(TokenRequestOptions{Scopes: []string{"read"}}), "read")
	requireIssuedFor(t, get(TokenRequestOptions{Scopes: []string{"write"}}), "write")
	requireIssuedFor(t, get(TokenRequestOptions{Resources: []string{"r1"}}), "", "r1")
	if got := requests.Load(); got != 3 {
		t.Fatalf("expected 3 token requests, got %d", got)
	}

	// ...and is cached independently: repeating the requests neither hits the server nor
	// returns the token of another combination.
	requireIssuedFor(t, get(TokenRequestOptions{Scopes: []string{"read"}}), "read")
	requireIssuedFor(t, get(TokenRequestOptions{Scopes: []string{"write"}}), "write")
	requireIssuedFor(t, get(TokenRequestOptions{Resources: []string{"r1"}}), "", "r1")
	if got := requests.Load(); got != 3 {
		t.Fatalf("expected repeated requests to be served from the cache, got %d token requests", got)
	}

	// The same scopes with other resources are a different request.
	requireIssuedFor(t, get(TokenRequestOptions{Scopes: []string{"read"}, Resources: []string{"r1"}}), "read", "r1")
	requireIssuedFor(t, get(TokenRequestOptions{Scopes: []string{"read"}}), "read")
	if got := requests.Load(); got != 4 {
		t.Fatalf("expected 4 token requests, got %d", got)
	}

	// Order and duplicates do not matter: the second call reuses the first token.
	requireIssuedFor(t, get(TokenRequestOptions{Scopes: []string{"a", "b"}}), "a b")
	requireIssuedFor(t, get(TokenRequestOptions{Scopes: []string{"b", "a", "a"}}), "a b")
	if got := requests.Load(); got != 5 {
		t.Fatalf("expected 5 token requests, got %d", got)
	}
}

func TestEffectiveRequest(t *testing.T) {
	defaultScopes := []string{"default"}
	defaultResources := []string{"default-resource"}

	tests := []struct {
		name          string
		opt           TokenRequestOptions
		wantScopes    []string
		wantResources []string
	}{
		{name: "no options use the defaults", wantScopes: defaultScopes, wantResources: defaultResources},
		{name: "empty slices use the defaults", opt: TokenRequestOptions{Scopes: []string{}, Resources: []string{}}, wantScopes: defaultScopes, wantResources: defaultResources},
		{name: "scopes override and keep the default resources", opt: TokenRequestOptions{Scopes: []string{"other"}}, wantScopes: []string{"other"}, wantResources: defaultResources},
		{name: "resources override and keep the default scopes", opt: TokenRequestOptions{Resources: []string{"other-resource"}}, wantScopes: defaultScopes, wantResources: []string{"other-resource"}},
		{name: "both override", opt: TokenRequestOptions{Scopes: []string{"a", "b"}, Resources: []string{"r1", "r2"}}, wantScopes: []string{"a", "b"}, wantResources: []string{"r1", "r2"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scopes, resources := effectiveRequest(test.opt, defaultScopes, defaultResources)
			if !slices.Equal(scopes, test.wantScopes) || !slices.Equal(resources, test.wantResources) {
				t.Fatalf("got scopes %v and resources %v, want scopes %v and resources %v", scopes, resources, test.wantScopes, test.wantResources)
			}
		})
	}

	t.Run("nothing requested without defaults", func(t *testing.T) {
		scopes, resources := effectiveRequest(TokenRequestOptions{}, nil, nil)
		if len(scopes) != 0 || len(resources) != 0 {
			t.Fatalf("expected nothing requested, got scopes %v and resources %v", scopes, resources)
		}
	})
}

func TestRequestKey(t *testing.T) {
	same := []struct {
		name string
		a, b [2][]string // scopes, resources
	}{
		{name: "order does not matter", a: [2][]string{{"a", "b"}, nil}, b: [2][]string{{"b", "a"}, nil}},
		{name: "duplicates do not matter", a: [2][]string{{"a", "b"}, nil}, b: [2][]string{{"b", "a", "a"}, nil}},
		{name: "resources order and duplicates do not matter", a: [2][]string{nil, {"r1", "r2"}}, b: [2][]string{nil, {"r2", "r1", "r2"}}},
		{name: "scopes are sent space separated, so these are one request", a: [2][]string{{"a b"}, nil}, b: [2][]string{{"a", "b"}, nil}},
		{name: "nil and empty are the same", a: [2][]string{nil, nil}, b: [2][]string{{}, {}}},
	}
	for _, test := range same {
		t.Run("same: "+test.name, func(t *testing.T) {
			if a, b := requestKey(test.a[0], test.a[1]), requestKey(test.b[0], test.b[1]); a != b {
				t.Fatalf("expected the same key, got %q and %q", a, b)
			}
		})
	}

	different := []struct {
		name string
		a, b [2][]string
	}{
		{name: "other scopes", a: [2][]string{{"read"}, nil}, b: [2][]string{{"write"}, nil}},
		{name: "other resources", a: [2][]string{{"read"}, {"r1"}}, b: [2][]string{{"read"}, {"r2"}}},
		{name: "same scopes, with and without resources", a: [2][]string{{"read"}, nil}, b: [2][]string{{"read"}, {"r1"}}},
		{name: "the boundary between scopes and resources", a: [2][]string{{"a"}, {"b", "c"}}, b: [2][]string{{"a", "b"}, {"c"}}},
		{name: "a value as scope or as resource", a: [2][]string{{"x"}, nil}, b: [2][]string{nil, {"x"}}},
	}
	for _, test := range different {
		t.Run("different: "+test.name, func(t *testing.T) {
			if a, b := requestKey(test.a[0], test.a[1]), requestKey(test.b[0], test.b[1]); a == b {
				t.Fatalf("expected different keys, both were %q", a)
			}
		})
	}
}

func TestProviderCollapsesConcurrentRefreshes(t *testing.T) {
	provider, requests := newCacheTestProvider(t, nil, nil)

	const callers = 20
	var wg sync.WaitGroup
	tokens := make([]string, callers)
	for i := range callers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			token, err := provider.Token(context.Background(), TokenRequestOptions{Scopes: []string{"s"}})
			if err != nil {
				t.Errorf("token: %v", err)
			}
			tokens[i] = token.AccessToken
		}()
	}
	wg.Wait()

	if got := requests.Load(); got != 1 {
		t.Fatalf("expected a single token request for concurrent callers, got %d", got)
	}
	for _, token := range tokens {
		requireIssuedFor(t, token, "s")
	}
}

func TestTokenCacheWaiterRetriesWhenLeaderContextIsCancelled(t *testing.T) {
	var cache tokenCache
	leaderStarted := make(chan struct{})
	leaderCtx, cancelLeader := context.WithCancel(context.Background())

	leaderDone := make(chan error, 1)
	go func() {
		_, err := cache.get(leaderCtx, "k", 0, func(ctx context.Context) (Token, error) {
			close(leaderStarted)
			<-ctx.Done()
			return Token{}, ctx.Err()
		})
		leaderDone <- err
	}()
	<-leaderStarted

	waiterDone := make(chan Token, 1)
	go func() {
		token, err := cache.get(context.Background(), "k", 0, func(context.Context) (Token, error) {
			return Token{AccessToken: "ok", RefreshOn: time.Now().Add(time.Hour)}, nil
		})
		if err != nil {
			t.Errorf("waiter: %v", err)
		}
		waiterDone <- token
	}()
	time.Sleep(20 * time.Millisecond) // let the waiter attach to the leader's call
	cancelLeader()

	if err := <-leaderDone; !errors.Is(err, context.Canceled) {
		t.Fatalf("leader: expected context.Canceled, got %v", err)
	}
	if token := <-waiterDone; token.AccessToken != "ok" {
		t.Fatalf("waiter should have fetched its own token, got %q", token.AccessToken)
	}
}

func TestTokenCacheDoesNotCacheErrors(t *testing.T) {
	var cache tokenCache
	var calls atomic.Int32
	fetch := func(context.Context) (Token, error) {
		if calls.Add(1) == 1 {
			return Token{}, errors.New("boom")
		}
		return Token{AccessToken: "ok", RefreshOn: time.Now().Add(time.Hour)}, nil
	}
	if _, err := cache.get(context.Background(), "k", 0, fetch); err == nil {
		t.Fatal("expected error")
	}
	if token, err := cache.get(context.Background(), "k", 0, fetch); err != nil || token.AccessToken != "ok" {
		t.Fatalf("expected recovery after error, got %v %v", token, err)
	}
}

func TestResolveValueOrder(t *testing.T) {
	t.Setenv("TEST_RESOLVE_VALUE", "")
	t.Setenv("TEST_RESOLVE_PATH", "")
	fallback := func(v string) func() string { return func() string { return v } }

	if got, _ := resolveValue("explicit", "", "TEST_RESOLVE_VALUE", "TEST_RESOLVE_PATH", fallback("f")); got != "explicit" {
		t.Fatalf("explicit should win, got %q", got)
	}
	t.Setenv("TEST_RESOLVE_VALUE", "env")
	if got, _ := resolveValue("", "", "TEST_RESOLVE_VALUE", "TEST_RESOLVE_PATH", fallback("f")); got != "env" {
		t.Fatalf("env should beat fallbacks, got %q", got)
	}
	t.Setenv("TEST_RESOLVE_VALUE", "")
	if got, _ := resolveValue("", "", "TEST_RESOLVE_VALUE", "TEST_RESOLVE_PATH", fallback(""), fallback("second")); got != "second" {
		t.Fatalf("expected first non-empty fallback, got %q", got)
	}
	t.Setenv("TEST_RESOLVE_PATH", "/nonexistent/file")
	if _, err := resolveValue("", "", "TEST_RESOLVE_VALUE", "TEST_RESOLVE_PATH", fallback("f")); err == nil {
		t.Fatal("expected an error for an unreadable path")
	}
}

func BenchmarkTokenCacheHit(b *testing.B) {
	var c tokenCache
	fetch := func(context.Context) (Token, error) {
		return Token{AccessToken: "t", RefreshOn: time.Now().Add(time.Hour), ExpiresOn: time.Now().Add(2 * time.Hour)}, nil
	}
	ctx := context.Background()
	_, _ = c.get(ctx, "k", time.Minute, fetch)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err := c.get(ctx, "k", time.Minute, fetch); err != nil {
				b.Fatal(err)
			}
		}
	})
}
