package identity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

const (
	cliProviderErrorPrefix = "cli provider"
	cliDefaultCommand      = "stackit"
	cliDefaultTimeout      = 30 * time.Second
	cliDefaultLeeway       = time.Minute
)

// cliExecutor runs the STACKIT CLI and returns its standard output.
// It is a field on the provider so tests can fake invoking the CLI.
type cliExecutor func(ctx context.Context, name string, args ...string) ([]byte, error)

// cliAccessTokenResponse is what `stackit auth get-access-token --output-format json` prints.
type cliAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

var _ TokenProvider = (*CLIProvider)(nil)

// CLIProviderConfig contains the configuration for CLIProvider.
type CLIProviderConfig struct {
	// Command is the STACKIT CLI executable to invoke. If empty, "stackit" is
	// resolved from PATH.
	Command string
	// CommandTimeout bounds how long the CLI may take to answer.
	// If zero, defaults to 30 seconds.
	CommandTimeout time.Duration
	// TokenRefreshLeeway controls how early before expiration the token is refreshed.
	// If zero, defaults to 1 minute.
	TokenRefreshLeeway time.Duration

	// exec is used by tests to fake invoking the CLI.
	exec cliExecutor
}

// CLIProvider reuses the session of a logged-in STACKIT CLI by invoking
// `stackit auth get-access-token`. It exists so that a developer who has run
// `stackit auth login` can use the SDK, and tools built on it, without
// provisioning any separate credential.
//
// This is a developer workstation credential. It requires the CLI to be
// installed and logged in, so it is unavailable in CI, containers and on
// servers — there NewCLIProvider fails and a chain simply skips it.
//
// The provider never starts an interactive login itself: if the CLI has no
// valid session, it reports an error rather than opening a browser.
type CLIProvider struct {
	name           string
	command        string
	commandTimeout time.Duration
	tokenLeeway    time.Duration
	exec           cliExecutor

	tokenMutex sync.RWMutex
	token      Token
}

// NewCLIProvider creates a CLIProvider. It returns an error when the CLI is not
// installed, so that a chain can skip it without paying a process spawn on every
// token request.
func NewCLIProvider(cfg *CLIProviderConfig) (*CLIProvider, error) {
	if cfg == nil {
		cfg = &CLIProviderConfig{}
	}

	command := cfg.Command
	if command == "" {
		command = cliDefaultCommand
	}

	execute := cfg.exec
	if execute == nil {
		// Resolve the executable once, here, rather than on every token request:
		// it turns "the CLI is not installed" into a construction error and keeps
		// a later PATH change from silently swapping the binary underneath us.
		path, err := exec.LookPath(command)
		if err != nil {
			return nil, fmt.Errorf("%s: %q not found: %w", cliProviderErrorPrefix, command, err)
		}
		command = path
		execute = runCLICommand
	}

	timeout := cfg.CommandTimeout
	if timeout == 0 {
		timeout = cliDefaultTimeout
	}

	leeway := cfg.TokenRefreshLeeway
	if leeway == 0 {
		leeway = cliDefaultLeeway
	}

	return &CLIProvider{
		name:           "CLIProvider",
		command:        command,
		commandTimeout: timeout,
		tokenLeeway:    leeway,
		exec:           execute,
	}, nil
}

// Token returns a valid access token from the CLI session.
func (p *CLIProvider) Token(ctx context.Context, _ TokenRequestOptions) (Token, error) {
	if p == nil {
		return Token{}, fmt.Errorf("%s: provider is not initialized", cliProviderErrorPrefix)
	}
	if ctx == nil {
		ctx = context.Background()
	}

	now := time.Now().Add(p.tokenLeeway)
	p.tokenMutex.RLock()
	cached := p.token
	p.tokenMutex.RUnlock()
	if cached.AccessToken != "" && cached.RefreshOn.After(now) {
		return cached, nil
	}

	fresh, err := p.requestToken(ctx)
	if err != nil {
		return Token{}, err
	}
	DebugContext(ctx, "identity: authenticated", "provider", p.name)

	p.tokenMutex.Lock()
	p.token = fresh
	p.tokenMutex.Unlock()

	return fresh, nil
}

func (p *CLIProvider) requestToken(ctx context.Context) (Token, error) {
	ctx, cancel := context.WithTimeout(ctx, p.commandTimeout)
	defer cancel()

	output, err := p.exec(ctx, p.command, "auth", "get-access-token", "--output-format", "json")
	if err != nil {
		return Token{}, fmt.Errorf("%s: %w", cliProviderErrorPrefix, err)
	}

	var response cliAccessTokenResponse
	if err := json.Unmarshal(bytes.TrimSpace(output), &response); err != nil {
		// Deliberately not including the output in the error: it is the token on success.
		return Token{}, fmt.Errorf("%s: parse access token response: %w", cliProviderErrorPrefix, err)
	}
	if response.AccessToken == "" {
		return Token{}, fmt.Errorf("%s: the CLI returned an empty access token", cliProviderErrorPrefix)
	}

	expiresOn, err := getTokenExpiration(response.AccessToken, 0)
	if err != nil {
		return Token{}, fmt.Errorf("%s: %w", cliProviderErrorPrefix, err)
	}

	return Token{
		AccessToken: response.AccessToken,
		TokenType:   tokenType(response.TokenType),
		ExpiresOn:   expiresOn,
		RefreshOn:   expiresOn,
	}, nil
}

// runCLICommand is the production cliExecutor.
func runCLICommand(ctx context.Context, name string, args ...string) ([]byte, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if message := strings.TrimSpace(stderr.String()); message != "" {
			return nil, fmt.Errorf("%w: %s", err, message)
		}
		return nil, err
	}
	return stdout.Bytes(), nil
}
