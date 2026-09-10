// Package identity contains modular token providers and composition primitives.
//
// Everything builds on one contract, TokenProvider, which returns an access token and
// nothing else. Each authentication flow ships as its own provider implementing it:
// StaticTokenProvider, ServiceAccountKeyProvider, WorkloadIdentityFederationProvider,
// InstanceMetadataProvider. ChainedProvider composes providers in order,
// and NewDefaultProvider assembles the opinionated chain the SDK uses by default.
//
// Implementations are safe for concurrent use and cache and refresh their own tokens, so
// callers do not have to. Tokens are never logged; logging is off unless SetLogger is
// called.
//
// Beyond authenticating SDK clients through config.WithTokenProvider, TokenProvider.Token
// returns the raw access token, which is what data plane APIs and non-SDK clients need in
// order to build their own Authorization header.
package identity
