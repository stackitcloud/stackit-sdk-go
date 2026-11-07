# credentials

Package `credentials` gets access tokens for the STACKIT APIs. Everything builds on one small contract:

```go
type TokenProvider interface {
	Token(ctx context.Context, options TokenRequestOptions) (Token, error)
}
```

Each authentication flow is a provider that implements it, and providers can be combined. A provider caches and refreshes its own tokens, so callers just ask for a token whenever they need one.

```go
import "github.com/stackitcloud/stackit-sdk-go/core/credentials"
```

- [Quick start](#quick-start)
- [Providers](#providers)
- [The default chain](#the-default-chain)
- [Combining providers](#combining-providers)
- [Scopes and resources](#scopes-and-resources)
- [Caching and concurrency](#caching-and-concurrency)
- [Getting the raw token](#getting-the-raw-token)
- [Logging](#logging)
- [Configuration reference](#configuration-reference)
- [Writing your own provider](#writing-your-own-provider)

## Quick start

Pass a provider to any service client with `config.WithTokenProvider`:

```go
provider, err := credentials.NewDefaultProvider(&credentials.DefaultProviderConfig{})
if err != nil {
	// no credentials could be found, the error says why each flow was left out
}

client, err := dns.NewAPIClient(config.WithTokenProvider(provider))
```

`config.WithTokenProvider` makes `auth.DefaultAuth` authenticate through the provider instead of the authentication flows that read their configuration from the environment. Options that explicitly choose another way to authenticate (`config.WithToken`, `config.WithServiceAccountKey`, `config.WithCustomAuth`, `config.WithoutAuthentication`, ...) still take precedence over it.

In production, prefer building the one provider you actually use over the default chain: it is more predictable and easier to debug.

```go
provider, err := credentials.NewServiceAccountKeyProvider(&credentials.ServiceAccountKeyProviderConfig{
	ServiceAccountKey: key,
})
```

## Providers

| Provider | Authenticates with | Refreshes |
|---|---|---|
| `WorkloadIdentityFederationProvider` | An OIDC token issued by your platform (for example a Kubernetes service account token), exchanged for a STACKIT token. No secret is stored. | Yes |
| `ServiceAccountKeyProvider` | A service account key and its RSA private key. The provider signs a short-lived assertion and exchanges it for a token. | Yes |
| `StaticTokenProvider` | A token you already have. | No: it is returned until it expires |
| `InstanceMetadataProvider` | The service account attached to a STACKIT VM, through the instance metadata service. | Yes |
| `CLIProvider` | The session of a logged-in STACKIT CLI. Meant for developer machines. | Yes |

Every provider has a `New…Provider(cfg)` constructor. A constructor returns an error when the provider cannot work (no key, no client ID, CLI not installed, ...), so a provider that was created is ready to use.

A provider resolves what it needs from **its config first, then from the environment variables, and for some values from the credentials file**, so anything explicit in code wins over the environment. See the [configuration reference](#configuration-reference).

### Workload identity federation

```go
provider, err := credentials.NewWorkloadIdentityFederationProvider(&credentials.WorkloadIdentityFederationProviderConfig{
	ClientID: "my-sa@sa.stackit.cloud", // or STACKIT_SERVICE_ACCOUNT_EMAIL
})
```

The platform issues the federated token and the provider only forwards it: it reads the token again at every refresh, from `FederatedTokenFunction` or, when that is nil, from the file in `STACKIT_FEDERATED_TOKEN_FILE` (default `/var/run/secrets/stackit.cloud/serviceaccount/token`). The constructor fails if there is no client ID or the token cannot be read.

### Service account key

```go
provider, err := credentials.NewServiceAccountKeyProvider(&credentials.ServiceAccountKeyProviderConfig{
	ServiceAccountKeyPath: "/path/to/service_account_key.json", // or ServiceAccountKey with the JSON itself
	PrivateKeyPath:        "/path/to/private_key.pem",          // only if you provided your own key pair
})
```

Both the key and the private key can be given as content (`ServiceAccountKey`, `PrivateKey`) or as the path of a file (`ServiceAccountKeyPath`, `PrivateKeyPath`). They are looked up in this order, and the first one found is used:

1. The content in the config.
2. The path in the config.
3. The environment variable with the content (`STACKIT_SERVICE_ACCOUNT_KEY`, `STACKIT_PRIVATE_KEY`).
4. The environment variable with the path (`STACKIT_SERVICE_ACCOUNT_KEY_PATH`, `STACKIT_PRIVATE_KEY_PATH`).
5. For the private key, the one embedded in the service account key when STACKIT generated the key pair.
6. The credentials file.

A path that is used but cannot be read makes the constructor fail. The token endpoint is `TokenURL`, then the one in the service account key, then the default STACKIT endpoint.

### Static token

```go
provider, err := credentials.NewStaticTokenProvider(&credentials.StaticTokenProviderConfig{Token: token})
```

The token must be a JWT, as its expiration is read from the `exp` claim. It is never refreshed: when it expires, requests fail.

### Instance metadata

```go
provider, err := credentials.NewInstanceMetadataProvider(&credentials.InstanceMetadataProviderConfig{
	ServiceAccountEmail: "my-sa@sa.stackit.cloud",
})
```

Only works on a STACKIT VM with a service account attached.

### CLI session

```go
provider, err := credentials.NewCLIProvider(&credentials.CLIProviderConfig{})
```

Runs `stackit auth get-access-token`, so it needs the CLI installed and `stackit auth login` done. It authenticates as the person logged in, not as an application, and never starts an interactive login: if the session is not valid it returns an error.

## The default chain

`NewDefaultProvider` builds a chain from the credentials it can find, and tries them in this order:

1. `WorkloadIdentityFederationProvider`
2. `ServiceAccountKeyProvider`
3. `StaticTokenProvider`
4. `InstanceMetadataProvider`
5. `CLIProvider`, **only if enabled**

The flows that need no secret or issue short-lived tokens come first, and a pre-issued token only applies when none of them is available. The ambient identity of the machine (instance metadata, then CLI) is only used when everything else has failed.

Flows whose credentials cannot be found are left out of the chain; `NewDefaultProvider` fails only when no flow at all could be configured, and the error lists why each one was skipped.

The CLI session is **off by default**, because it uses the permissions of a person and a misconfigured service credential could otherwise fall back to them without anyone noticing. Enable it on a developer machine with `DefaultProviderConfig{EnableCLI: true}` or `STACKIT_USE_CLI=true`.

The chain cannot tell whether a flow was configured in code or in the environment, so a flow set up in the environment can take precedence over a less secure one passed in `DefaultProviderConfig`.

## Combining providers

`ChainedProvider` tries providers in order until one returns a token:

```go
chain, err := credentials.NewChainedProvider(keyProvider, cliProvider)
```

By default, once a provider has succeeded the chain keeps using only that one. With `ChainedProviderOptions{RetrySources: true}`, every call tries all the providers in order again.

## Scopes and resources

A provider can request scopes and resources, either for every token through its config or for a single call:

```go
token, err := provider.Token(ctx, credentials.TokenRequestOptions{
	Scopes:    []string{"read", "write"},
	Resources: []string{"https://api.example.com"},
})
```

What a call asks for replaces the configured default. Resources follow [RFC 8707](https://www.rfc-editor.org/rfc/rfc8707): each one becomes an element of the audience (`aud`) of the token. The static token, instance metadata and CLI providers return the same token whatever is asked.

## Caching and concurrency

All providers are safe for concurrent use.

- A provider keeps **one token for each combination of scopes and resources**, so asking for different ones never returns a token issued for another combination. Order and duplicates do not matter: `[b, a, a]` and `[a, b]` are the same request.
- A token is refreshed shortly before it expires (`TokenRefreshLeeway`, one minute by default).
- If several goroutines need a token that is missing or expiring, **a single request is made** and they all receive its result. Errors are not cached.

Always pass the `context.Context` of the operation: it cancels a request that is waiting for a token.

## Getting the raw token

Besides authenticating SDK clients, a provider gives you the access token, which is what data plane APIs and clients outside the SDK need. Build the header with the token's own type instead of assuming `Bearer`:

```go
token, err := provider.Token(ctx, credentials.TokenRequestOptions{})
if err != nil {
	return err
}
req.Header.Set("Authorization", token.TokenType+" "+token.AccessToken)
```

`Token` also carries `ExpiresOn` and `RefreshOn`.

## Logging

The package is silent by default and never logs tokens. To see what it does, set a logger:

```go
credentials.SetLogger(slog.Default())
```

At debug level it reports, for example, why a flow was left out of the default chain and which provider of a chain answered. It logs a warning when the credentials file can be read by other users, and an error when the token endpoint rejects a service account key request.

## Configuration reference

| Environment variable | Used by | Meaning |
|---|---|---|
| `STACKIT_SERVICE_ACCOUNT_EMAIL` | WIF, default chain (for instance metadata) | The service account (client ID) |
| `STACKIT_SERVICE_ACCOUNT_TOKEN` | Static token | A pre-issued token |
| `STACKIT_SERVICE_ACCOUNT_KEY` | Service account key | The service account key, as JSON |
| `STACKIT_SERVICE_ACCOUNT_KEY_PATH` | Service account key | Path of a file with the service account key |
| `STACKIT_PRIVATE_KEY` | Service account key | The RSA private key, PEM encoded |
| `STACKIT_PRIVATE_KEY_PATH` | Service account key | Path of a file with the private key |
| `STACKIT_FEDERATED_TOKEN_FILE` | WIF | Path of the file with the OIDC token |
| `STACKIT_IDP_TOKEN_ENDPOINT` | WIF | Token endpoint |
| `STACKIT_CREDENTIALS_PATH` | Static token, service account key, default chain | Credentials file, `~/.stackit/credentials.json` by default |
| `STACKIT_USE_CLI` | Default chain | `true` adds the CLI session to the chain |

The credentials file is a JSON object that can define `STACKIT_SERVICE_ACCOUNT_TOKEN`, `STACKIT_SERVICE_ACCOUNT_KEY`, `STACKIT_PRIVATE_KEY` and `STACKIT_SERVICE_ACCOUNT_EMAIL`. It may contain secrets, so keep it readable only by its owner (`chmod 600`); with a logger set, the package warns when it is not.

All configs also accept an `HTTPClient` for the token requests, and `TokenRefreshLeeway`. Only the transport of the client is used, so a client that already authenticates its requests cannot end up asking for a token to get a token.

## Writing your own provider

Anything that implements `TokenProvider` can be used with `config.WithTokenProvider` and in a `ChainedProvider`. Cache the token you obtain, honor the `context`, and set `Token.TokenType` (the SDK assumes `Bearer` when it is empty):

```go
type myProvider struct{ /* ... */ }

func (p *myProvider) Token(ctx context.Context, opt credentials.TokenRequestOptions) (credentials.Token, error) {
	// get or reuse a token
}
```
