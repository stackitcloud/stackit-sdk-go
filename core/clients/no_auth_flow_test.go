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
		{"with transport", args{context.Background(), NoAuthFlowConfig{HTTPTransport: http.DefaultTransport}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NoAuthFlow{}
			if err := c.Init(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("NoAuthFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNoAuthFlow_Do(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		noAuthFlow *NoAuthFlow
		handlerFn  func(tb testing.TB) http.HandlerFunc
		want       int
		wantErr    bool
	}{
		{
			name:       "success with rt",
			noAuthFlow: &NoAuthFlow{http.DefaultTransport, &NoAuthFlowConfig{}},
			handlerFn: func(_ testing.TB) http.HandlerFunc {
				return func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
				}
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			name:       "success with code 500",
			noAuthFlow: &NoAuthFlow{http.DefaultTransport, &NoAuthFlowConfig{}},
			handlerFn: func(_ testing.TB) http.HandlerFunc {
				return func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "text/html")
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = fmt.Fprintln(w, `<html>Internal Server Error</html>`)
				}
			},
			want:    http.StatusInternalServerError,
			wantErr: false,
		},
		{
			name: "success with custom transport",
			noAuthFlow: &NoAuthFlow{
				mockTransportFn{
					fn: func(req *http.Request) (*http.Response, error) {
						req.Header.Set("User-Agent", "custom_transport")

						return http.DefaultTransport.RoundTrip(req)
					},
				},
				&NoAuthFlowConfig{},
			},
			handlerFn: func(tb testing.TB) http.HandlerFunc {
				tb.Helper()

				return func(w http.ResponseWriter, r *http.Request) {
					if r.Header.Get("User-Agent") != "custom_transport" {
						tb.Errorf("expected User-Agent header to be 'custom_transport', but got %s", r.Header.Get("User-Agent"))
					}

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
				}
			},
			want:    http.StatusOK,
			wantErr: false,
		},
		{
			name: "fail with custom proxy",
			noAuthFlow: &NoAuthFlow{
				&http.Transport{
					Proxy: func(_ *http.Request) (*url.URL, error) {
						return nil, fmt.Errorf("proxy error")
					},
				},
				&NoAuthFlowConfig{},
			},
			handlerFn: func(testing.TB) http.HandlerFunc {
				return func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
				}
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.handlerFn(t))
			t.Cleanup(server.Close)

			u, err := url.Parse(server.URL)
			if err != nil {
				t.Errorf("no error is expected, but got %v", err)
			}

			req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
			if err != nil {
				t.Errorf("no error is expected, but got %v", err)
			}

			httpClient := &http.Client{
				Transport: tt.noAuthFlow,
			}

			res, err := httpClient.Do(req)

			if tt.wantErr {
				if err == nil {
					t.Errorf("error is expected, but got %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("no error is expected, but got %v", err)
				}

				if res.StatusCode != tt.want {
					t.Errorf("expected status code %d, but got %d", tt.want, res.StatusCode)
				}

				// Defer discard and close the body
				t.Cleanup(func() {
					if _, err := io.Copy(io.Discard, res.Body); err != nil {
						t.Errorf("no error is expected, but got %v", err)
					}

					if err := res.Body.Close(); err != nil {
						t.Errorf("no error is expected, but got %v", err)
					}
				})
			}
		})
	}
}
