package clients

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func TestKeyFlow_Init(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg *KeyFlowConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{context.Background(), &KeyFlowConfig{
			ServiceAccountKey: "efg",
		}}, false},
		{"error 1", args{context.Background(), &KeyFlowConfig{
			ServiceAccountKey: "",
		}}, true},
	}
	b := os.Getenv(ServiceAccountKey)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{}

			err := os.Setenv(ServiceAccountKey, "")
			if err != nil {
				t.Fatalf("Setting service account Key: %s", err)
			}
			if err := c.Init(tt.args.ctx, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			err = os.Setenv(ServiceAccountKey, b)
			if err != nil {
				t.Fatalf("Setting service account Key: %s", err)
			}
			if c.config == nil {
				t.Error("config is nil")
			}
		})
	}
}

func TestKeyFlow_Do(t *testing.T) {
	type fields struct {
		client *http.Client
		config *KeyFlowConfig
	}
	type args struct{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"fail", fields{nil, nil}, args{}, 0, true},
		{"success", fields{&http.Client{}, &KeyFlowConfig{ClientRetry: &RetryConfig{}}}, args{}, http.StatusOK, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{
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
				t.Errorf("KeyFlow.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.StatusCode != tt.want {
				t.Errorf("KeyFlow.Do() = %v, want %v", got.StatusCode, tt.want)
			}
		})
	}
}

func TestKeyClone(t *testing.T) {
	c := &KeyFlow{
		client: &http.Client{},
		config: &KeyFlowConfig{},
	}

	clone, ok := c.Clone().(*KeyFlow)
	if !ok {
		t.Fatalf("Type assertion failed")
	}

	if !reflect.DeepEqual(c, clone) {
		t.Errorf("Clone() = %v, want %v", clone, c)
	}
}
