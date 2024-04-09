package runtime

import (
	"context"
	"net/http"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

// WithCaptureHTTPResponse adds the raw HTTP response retrieval annotation to the parent context.
// The resp parameter will contain the raw HTTP response after the request has completed.
func WithCaptureHTTPResponse(parent context.Context, resp **http.Response) context.Context {
	return context.WithValue(parent, config.ContextHTTPResponse, resp)
}

// WithCaptureHTTPRequest adds the raw HTTP request to the parent context.
// The req parameter will contain the raw HTTP request.
func WithCaptureHTTPRequest(parent context.Context, req **http.Request) context.Context {
	return context.WithValue(parent, config.ContextHTTPRequest, req)
}
