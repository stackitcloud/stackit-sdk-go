package clients

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNoAuthFlow_Init(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg NoAuthFlowConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{context.Background(), NoAuthFlowConfig{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NoAuthFlow{}
			if err := c.Init(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("NoAuthFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if c.config == nil {
				t.Error("config is nil")
			}
		})
	}
}

func TestNoAuthFlow_Do(t *testing.T) {
	type fields struct {
		client *http.Client
		config *NoAuthFlowConfig
	}
	type args struct{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "fail",
			fields:  fields{nil, nil},
			args:    args{},
			want:    0,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				&http.Client{},
				&NoAuthFlowConfig{
					ClientRetry: &RetryConfig{},
				},
			},
			args:    args{},
			want:    http.StatusOK,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NoAuthFlow{
				client: tt.fields.client,
				config: tt.fields.config,
			}
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, `{"status":"ok"}`)
			})
			server := httptest.NewServer(handler)
			defer server.Close()
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Error(err)
				return
			}
			req := &http.Request{
				URL: u,
			}
			got, err := c.RoundTrip(req)
			if err == nil {
				// Defer discard and close the body
				defer func() {
					if _, discardErr := io.Copy(io.Discard, got.Body); discardErr != nil && err == nil {
						err = discardErr
					}
					if closeErr := got.Body.Close(); closeErr != nil && err == nil {
						err = closeErr
					}
				}()
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NoAuthFlow.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.StatusCode != tt.want {
				t.Errorf("NoAuthFlow.Do() = %v, want %v", got.StatusCode, tt.want)
			}
		})
	}
}
