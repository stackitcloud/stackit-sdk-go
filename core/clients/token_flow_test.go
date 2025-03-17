package clients

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestTokenFlow_Init(t *testing.T) {
	type args struct {
		cfg *TokenFlowConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{&TokenFlowConfig{
			ServiceAccountToken: "efg",
		}}, false},
		{"with transport", args{&TokenFlowConfig{
			ServiceAccountToken: "efg",
			HTTPTransport:       http.DefaultTransport,
		}}, false},
		{"error 1", args{&TokenFlowConfig{
			ServiceAccountToken: "",
		}}, true},
	}
	b := os.Getenv(ServiceAccountToken)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TokenFlow{}

			err := os.Setenv(ServiceAccountToken, "")
			if err != nil {
				t.Fatalf("Setting service account token: %s", err)
			}
			if err := c.Init(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("TokenFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			err = os.Setenv(ServiceAccountToken, b)
			if err != nil {
				t.Fatalf("Setting service account token: %s", err)
			}
			if c.config == nil {
				t.Error("config is nil")
			}
		})
	}
}

func TestTokenFlow_Do(t *testing.T) {
	type fields struct {
		rt     http.RoundTripper
		config *TokenFlowConfig
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
		{"success", fields{http.DefaultTransport, &TokenFlowConfig{}}, args{}, http.StatusOK, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TokenFlow{
				rt:     tt.fields.rt,
				config: tt.fields.config,
			}
			handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
			})
			server := httptest.NewServer(handler)
			defer server.Close()
			u, err := url.Parse(server.URL)
			if err != nil {
				t.Error(err)
				return
			}
			req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
			if err != nil {
				t.Error(err)
				return
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
				t.Errorf("TokenFlow.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.StatusCode != tt.want {
				t.Errorf("TokenFlow.Do() = %v, want %v", got.StatusCode, tt.want)
			}
		})
	}
}
