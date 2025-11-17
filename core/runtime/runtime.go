package runtime

import (
	"context"
	"net/http"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

const (
	xTraceIdHeader = "x-trace-id"
)

// WithCaptureHTTPResponse adds the raw HTTP response retrieval annotation to the parent context.
// The resp parameter will contain the raw HTTP response after the request has completed.
func WithCaptureHTTPResponse(parent context.Context, resp **http.Response) context.Context {
	return context.WithValue(parent, config.ContextHTTPResponse, resp)
}

// WithCaptureHTTPRequest adds the raw HTTP request retrieval annotation to the parent context.
// After the request has completed, the req parameter will contain the raw HTTP request made to the API.
func WithCaptureHTTPRequest(parent context.Context, req **http.Request) context.Context {
	return context.WithValue(parent, config.ContextHTTPRequest, req)
}

// GetTraceId returns the X-trace-id from the last response. If no trace-id can be found, it returns an empty string.
// Prerequisite is, that WithCaptureHTTPResponse was executed before. It reads the X-trace-id header from the
// attached http response within the context.
func GetTraceId(ctx context.Context) string {
	var traceId string
	if resp, ok := ctx.Value(config.ContextHTTPResponse).(**http.Response); ok {
		if resp != nil && *resp != nil {
			traceId = (*resp).Header.Get(xTraceIdHeader)
		}
	}
	return traceId
}
