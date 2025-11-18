package runtime

import (
	"context"
	"net/http"
	"testing"
)

func TestGetTraceId(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantTraceId string
	}{
		{
			name: "with set traceId",
			args: args{
				ctx: func() context.Context {
					ctx := context.Background()
					header := http.Header{}
					header.Set(xTraceIdHeader, "my-trace-id")
					httpResp := &http.Response{
						Header: header,
					}
					ctx = WithCaptureHTTPResponse(ctx, &httpResp)
					return ctx
				}(),
			},
			wantTraceId: "my-trace-id",
		},
		{
			name: "without traceId header",
			args: args{
				ctx: func() context.Context {
					ctx := context.Background()
					httpResp := &http.Response{
						Header: http.Header{},
					}
					ctx = WithCaptureHTTPResponse(ctx, &httpResp)
					return ctx
				}(),
			},
			wantTraceId: "",
		},
		{
			name: "with empty response",
			args: args{
				ctx: func() context.Context {
					ctx := context.Background()
					var httpResp *http.Response
					ctx = WithCaptureHTTPResponse(ctx, &httpResp)
					return ctx
				}(),
			},
			wantTraceId: "",
		},
		{
			name: "without response in context",
			args: args{
				ctx: context.Background(),
			},
			wantTraceId: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTraceId := GetTraceId(tt.args.ctx); gotTraceId != tt.wantTraceId {
				t.Errorf("GetTraceId() = %v, want %v", gotTraceId, tt.wantTraceId)
			}
		})
	}
}
