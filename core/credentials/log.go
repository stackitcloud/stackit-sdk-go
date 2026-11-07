package credentials

import (
	"context"
	"log/slog"
)

var credentialsLogger *slog.Logger

// SetLogger sets the slog.Logger used by the credentials providers.
// Pass nil to disable logging (the default — the package is silent unless opted in).
//
// Example:
//
//	credentials.SetLogger(slog.Default())
func SetLogger(l *slog.Logger) {
	credentialsLogger = l
}

// debugContext emits a Debug-level log to the registered logger, if any.
// args follows slog convention: alternating key, value pairs or slog.Attr values.
func debugContext(ctx context.Context, msg string, args ...any) {
	if credentialsLogger != nil {
		credentialsLogger.DebugContext(ctx, msg, args...)
	}
}

// warnContext emits a Warn-level log to the registered logger, if any.
// args follows slog convention: alternating key, value pairs or slog.Attr values.
func warnContext(ctx context.Context, msg string, args ...any) {
	if credentialsLogger != nil {
		credentialsLogger.WarnContext(ctx, msg, args...)
	}
}

// errorContext emits an Error-level log to the registered logger, if any.
// args follows slog convention: alternating key, value pairs or slog.Attr values.
func errorContext(ctx context.Context, msg string, args ...any) {
	if credentialsLogger != nil {
		credentialsLogger.ErrorContext(ctx, msg, args...)
	}
}
