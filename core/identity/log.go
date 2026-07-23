package identity

import (
	"context"
	"log/slog"
)

var identityLogger *slog.Logger

// SetLogger sets the slog.Logger used by identity providers.
// Pass nil to disable logging (the default — identity is silent unless opted in).
//
// Example:
//
//	identity.SetLogger(slog.Default())
func SetLogger(l *slog.Logger) {
	identityLogger = l
}

// DebugContext emits a Debug-level log to the registered logger, if any.
// args follows slog convention: alternating key, value pairs or slog.Attr values.
func DebugContext(ctx context.Context, msg string, args ...any) {
	if identityLogger != nil {
		identityLogger.DebugContext(ctx, msg, args...)
	}
}

// InfoContext emits an Info-level log to the registered logger, if any.
// args follows slog convention: alternating key, value pairs or slog.Attr values.
func InfoContext(ctx context.Context, msg string, args ...any) {
	if identityLogger != nil {
		identityLogger.InfoContext(ctx, msg, args...)
	}
}

// WarnContext emits a Warn-level log to the registered logger, if any.
// args follows slog convention: alternating key, value pairs or slog.Attr values.
func WarnContext(ctx context.Context, msg string, args ...any) {
	if identityLogger != nil {
		identityLogger.WarnContext(ctx, msg, args...)
	}
}

// ErrorContext emits an Error-level log to the registered logger, if any.
// args follows slog convention: alternating key, value pairs or slog.Attr values.
func ErrorContext(ctx context.Context, msg string, args ...any) {
	if identityLogger != nil {
		identityLogger.ErrorContext(ctx, msg, args...)
	}
}
