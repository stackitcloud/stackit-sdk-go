// Part of this code is taken from the client.go file
// generated by the OpenApi generator for each API client
package config

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/clients"
)

const (
	global = "global"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().
type contextKey string

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKeys takes a string apikey as authentication for the request
	ContextAPIKeys = contextKey("apiKeys")

	// ContextHttpSignatureAuth takes HttpSignatureAuth as authentication for the request.
	ContextHttpSignatureAuth = contextKey("httpsignature")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")

	// ContextHTTPResponse holds the raw HTTP response after the request has completed.
	ContextHTTPResponse = contextKey("httpResponse")

	// ContextHTTPRequest holds the raw HTTP request.
	ContextHTTPRequest = contextKey("httpRequest")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// Middleware is a function that wraps an http.RoundTripper to provide additional functionality
// such as logging, authentication, etc.
type Middleware func(http.RoundTripper) http.RoundTripper

// Configuration stores the configuration of the API client
type Configuration struct {
	Host                  string            `json:"host,omitempty"`
	Scheme                string            `json:"scheme,omitempty"`
	DefaultHeader         map[string]string `json:"defaultHeader,omitempty"`
	UserAgent             string            `json:"userAgent,omitempty"`
	Debug                 bool              `json:"debug,omitempty"`
	NoAuth                bool              `json:"noAuth,omitempty"`
	ServiceAccountEmail   string            `json:"serviceAccountEmail,omitempty"`
	Token                 string            `json:"token,omitempty"`
	ServiceAccountKey     string            `json:"serviceAccountKey,omitempty"`
	PrivateKey            string            `json:"privateKey,omitempty"`
	ServiceAccountKeyPath string            `json:"serviceAccountKeyPath,omitempty"`
	PrivateKeyPath        string            `json:"privateKeyPath,omitempty"`
	CredentialsFilePath   string            `json:"credentialsFilePath,omitempty"`
	TokenCustomUrl        string            `json:"tokenCustomUrl,omitempty"`
	Region                string            `json:"region,omitempty"`
	CustomAuth            http.RoundTripper
	Servers               ServerConfigurations
	OperationServers      map[string]ServerConfigurations
	HTTPClient            *http.Client
	Middleware            []Middleware

	// If != nil, a goroutine will be launched that will refresh the service account's access token when it's close to being expired.
	// The goroutine is killed whenever this context is canceled.
	//
	// Only has effect for key flow
	BackgroundTokenRefreshContext context.Context

	// Deprecated: retry options were removed to reduce complexity of the client. If this functionality is needed, you can provide your own custom HTTP client. This field has no effect, and will be removed in a later update
	RetryOptions *clients.RetryConfig //nolint:staticcheck //will be removed in a later update

	setCustomEndpoint bool
}

// ConfigurationOption is an option for an API client. The options are executed sequentially, so
// if you use withHTTPClient with a custom client that has a timeout of 1 minute and then you use
// withTimeout with a timeout of 2 minutes, the HTTP client timeout will be last one to be set.
// In the same way, if you use withHTTPClient as the last option, it will override all of the
// previous configurations to the HTTPClient
type ConfigurationOption func(*Configuration) error

// WithHTTPClient returns a ConfigurationOption that specifies the HTTP client to use
// as basis for the communication.
// Warning: providing this option overrides authentication, if you just want to customize the http client parameters except Transport,
// you can use the WithCheckRedirect, WithJar and WithTimeout
func WithHTTPClient(client *http.Client) ConfigurationOption {
	return func(config *Configuration) error {
		config.HTTPClient = client
		return nil
	}
}

// WithCustomAuth returns a ConfigurationOption that provides a custom http.Roundtripper
// for making authenticated requests
func WithCustomAuth(rt http.RoundTripper) ConfigurationOption {
	return func(config *Configuration) error {
		config.CustomAuth = rt
		return nil
	}
}

// WithUserAgent returns a ConfigurationOption that defines the User-Agent
func WithUserAgent(userAgent string) ConfigurationOption {
	return func(config *Configuration) error {
		config.UserAgent = userAgent
		return nil
	}
}

// WithRegion returns a ConfigurationOption that specifies the region to be used
func WithRegion(region string) ConfigurationOption {
	return func(config *Configuration) error {
		config.Region = region
		return nil
	}
}

// WithEndpoint returns a ConfigurationOption that overrides the default endpoint to be used for the client
// This option takes precedence over withRegion
func WithEndpoint(endpoint string) ConfigurationOption {
	return func(config *Configuration) error {
		customServers := ServerConfigurations{
			{
				URL:         endpoint,
				Description: "User provided URL",
			},
		}
		config.Servers = customServers
		config.setCustomEndpoint = true
		return nil
	}
}

// WithTokenEndpoint returns a ConfigurationOption that overrides the default url to be used to get a token when using the key flow
func WithTokenEndpoint(url string) ConfigurationOption {
	return func(config *Configuration) error {
		config.TokenCustomUrl = url
		return nil
	}
}

// WithServiceAccountEmail returns a ConfigurationOption that sets the service account email
func WithServiceAccountEmail(serviceAccountEmail string) ConfigurationOption {
	return func(config *Configuration) error {
		config.ServiceAccountEmail = serviceAccountEmail
		return nil
	}
}

// WithServiceAccountKey returns a ConfigurationOption that sets the service account key
// This option takes precedence over WithServiceAccountKeyPath
func WithServiceAccountKey(serviceAccountKey string) ConfigurationOption {
	return func(config *Configuration) error {
		config.ServiceAccountKey = serviceAccountKey
		return nil
	}
}

// WithServiceAccountKeyPath returns a ConfigurationOption that sets the service account key path
func WithServiceAccountKeyPath(serviceAccountKeyPath string) ConfigurationOption {
	return func(config *Configuration) error {
		config.ServiceAccountKeyPath = serviceAccountKeyPath
		return nil
	}
}

// WithPrivateKey returns a ConfigurationOption that sets the private key
// This option takes precedence over WithPrivateKeyPath
func WithPrivateKey(privateKey string) ConfigurationOption {
	return func(config *Configuration) error {
		config.PrivateKey = privateKey
		return nil
	}
}

// WithPrivateKeyPath returns a ConfigurationOption that sets the private key path
func WithPrivateKeyPath(privateKeyPath string) ConfigurationOption {
	return func(config *Configuration) error {
		config.PrivateKeyPath = privateKeyPath
		return nil
	}
}

// WithRetryTimeout returns a ConfigurationOption that specifies the maximum time for ret
// WithoutAuthentication returns a ConfigurationOption that disables authentication.
// This option takes precedence over the WithToken option
func WithoutAuthentication() ConfigurationOption {
	return func(config *Configuration) error {
		config.NoAuth = true
		return nil
	}
}

// WithToken returns a ConfigurationOption that sets a token to be used for authentication in API calls
func WithToken(token string) ConfigurationOption {
	return func(config *Configuration) error {
		config.Token = token
		return nil
	}
}

// Deprecated: retry options were removed to reduce complexity of the client. If this functionality is needed, you can provide your own custom HTTP client. This option has no effect, and will be removed in a later update
func WithMaxRetries(_ int) ConfigurationOption {
	return func(_ *Configuration) error {
		return nil
	}
}

// Deprecated: retry options were removed to reduce complexity of the client. If this functionality is needed, you can provide your own custom HTTP client. This option has no effect, and will be removed in a later update
func WithWaitBetweenCalls(_ time.Duration) ConfigurationOption {
	return func(_ *Configuration) error {
		return nil
	}
}

// Deprecated: retry options were removed to reduce complexity of the client. If this functionality is needed, you can provide your own custom HTTP client. This option has no effect, and will be removed in a later update
func WithRetryTimeout(_ time.Duration) ConfigurationOption {
	return func(_ *Configuration) error {
		return nil
	}
}

// WithTimeout returns a ConfigurationOption that specifies the HTTP client timeout
func WithTimeout(timeout time.Duration) ConfigurationOption {
	return func(config *Configuration) error {
		config.HTTPClient.Timeout = timeout
		return nil
	}
}

// WithCheckRedirect returns a ConfigurationOption that specifies the HTTP client checkRedirect function
func WithCheckRedirect(checkRedirect func(req *http.Request, via []*http.Request) error) ConfigurationOption {
	return func(config *Configuration) error {
		config.HTTPClient.CheckRedirect = checkRedirect
		return nil
	}
}

// WithJar returns a ConfigurationOption that specifies the HTTP client cookie jar
func WithJar(jar http.CookieJar) ConfigurationOption {
	return func(config *Configuration) error {
		config.HTTPClient.Jar = jar
		return nil
	}
}

// WithMiddleware returns a ConfigurationOption that adds a Middleware to the client.
// The Middleware is prepended to the list of Middlewares so that the last added Middleware is the first to be executed.
// Warning: Providing this option may overide the authentication performed by the SDK if the middlewares provided break the chain.
// If changes are made to the authentication header and the chain is preserved, they will be overwritten. If you wish to overwrite authentication, use WithCustomAuth.
func WithMiddleware(m Middleware) ConfigurationOption {
	// Prepend m to the list of middlewares
	return func(config *Configuration) error {
		config.Middleware = append([]Middleware{m}, config.Middleware...)
		return nil
	}
}

// WithBackgroundTokenRefresh returns a ConfigurationOption that enables access token refreshing in backgound.
//
// If enabled, a goroutine will be launched that will refresh the service account's access token when it's close to being expired.
// The goroutine is killed whenever the given context is canceled.
//
// Only has effect for key flow
func WithBackgroundTokenRefresh(ctx context.Context) ConfigurationOption {
	return func(c *Configuration) error {
		if ctx == nil {
			return fmt.Errorf("context for token refresh in background cannot be empty")
		}
		c.BackgroundTokenRefreshContext = ctx
		return nil
	}
}

// WithCustomConfiguration returns a ConfigurationOption that sets a custom Configuration
func WithCustomConfiguration(cfg *Configuration) ConfigurationOption {
	return func(config *Configuration) error {
		config.Host = cfg.Host
		config.Scheme = cfg.Scheme
		config.DefaultHeader = cfg.DefaultHeader
		config.UserAgent = cfg.UserAgent
		config.Debug = cfg.Debug
		config.NoAuth = cfg.NoAuth
		config.Token = cfg.Token
		config.ServiceAccountKey = cfg.ServiceAccountKey
		config.ServiceAccountEmail = cfg.ServiceAccountEmail
		config.ServiceAccountKeyPath = cfg.ServiceAccountKeyPath
		config.PrivateKey = cfg.PrivateKey
		config.PrivateKeyPath = cfg.PrivateKeyPath
		config.Region = cfg.Region
		config.CredentialsFilePath = cfg.CredentialsFilePath
		config.CustomAuth = cfg.CustomAuth
		config.Servers = cfg.Servers
		config.setCustomEndpoint = (len(cfg.Servers) > 0)
		config.OperationServers = cfg.OperationServers
		config.HTTPClient = cfg.HTTPClient
		config.BackgroundTokenRefreshContext = cfg.BackgroundTokenRefreshContext
		return nil
	}
}

// WithCaptureHTTPResponse adds the raw HTTP response retrieval annotation to the parent context.
// The resp parameter will contain the raw HTTP response after the request has completed.
//
// Deprecated: Use runtime.WithCaptureHTTPResponse instead
func WithCaptureHTTPResponse(parent context.Context, resp **http.Response) context.Context {
	return context.WithValue(parent, ContextHTTPResponse, resp)
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	serverUrl := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			serverUrl = strings.Replace(serverUrl, "{"+name+"}", value, -1)
		} else {
			serverUrl = strings.Replace(serverUrl, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return serverUrl, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, fmt.Errorf("invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		operationIndices, ok := osi.(map[string]int)
		if !ok {
			return 0, fmt.Errorf("invalid type %T should be map[string]int", osi)
		}
		index, ok := operationIndices[endpoint]
		if ok {
			return index, nil
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, fmt.Errorf("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		operationVariables, ok := osv.(map[string]map[string]string)
		if !ok {
			return nil, fmt.Errorf("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		}
		variables, ok := operationVariables[endpoint]
		if ok {
			return variables, nil
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}

// ConfigureRegion configures the API server urls with the user specified region.
// Does nothing if a custom endpoint is provided.
// Throws an error if no region is given or if the region is not valid
func ConfigureRegion(cfg *Configuration) error {
	if cfg.setCustomEndpoint {
		return nil
	}

	servers := cfg.Servers
	oasRegion, ok := servers[0].Variables["region"]
	if !ok {
		// No region variable found
		return fmt.Errorf("no region was found in the API configuration")
	}

	var availableRegions []string
	for _, regionWithDotSuffix := range oasRegion.EnumValues {
		availableRegions = append(availableRegions, strings.TrimSuffix(regionWithDotSuffix, "."))
	}

	if cfg.Region == "" {
		// Check region
		envVarRegion, _ := os.LookupEnv("STACKIT_REGION")
		cfg.Region = envVarRegion
	}

	if oasRegion.DefaultValue != "" && oasRegion.DefaultValue != global {
		if cfg.Region == "" {
			return fmt.Errorf("no region was provided, available regions are: %s", availableRegions)
		}
		// API is regional (not global)
		if containsCaseSensitive(availableRegions, cfg.Region) {
			cfgUrl := strings.Replace(servers[0].URL, "{region}", fmt.Sprintf("%s.", cfg.Region), -1)
			cfg.Servers = ServerConfigurations{
				{
					URL:         cfgUrl,
					Description: servers[0].Description,
				},
			}
			return nil
		}
		// Region is not available.
		return fmt.Errorf("the provided region is not available for this API, available regions are: %s", availableRegions)
	}
	// Global API. The provided region is ignored.
	// If the url is a template, generated using deprecated config.json, the region variable is replaced
	// If the url is already configured, there is no region variable and it remains the same
	cfgUrl := strings.Replace(servers[0].URL, "{region}", "", -1)
	cfg.Servers = ServerConfigurations{
		{
			URL:         cfgUrl,
			Description: servers[0].Description,
		},
	}
	return nil
}

// containsCaseSensitive is a case sensitive match, finding needle in a haystack
func containsCaseSensitive(haystack []string, needle string) bool {
	for _, a := range haystack {
		if needle == a {
			return true
		}
	}
	return false
}

// ChainMiddleware chains multiple middlewares to create a single http.RoundTripper
// The middlewares are applied in reverse order, so the last middleware provided in the arguments is the first to be executed
// If the root http.RoundTripper is nil, http.DefaultTransport is used
func ChainMiddleware(rt http.RoundTripper, middlewares ...Middleware) http.RoundTripper {
	if rt == nil {
		rt = http.DefaultTransport
	}

	// Apply middlewares in reverse order
	for i := len(middlewares) - 1; i >= 0; i-- {
		rt = middlewares[i](rt)
	}

	return rt
}
