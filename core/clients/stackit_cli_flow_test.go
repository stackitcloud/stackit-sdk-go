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

//nolint:gosec // testServiceAccountToken is a test token
const testServiceAccountToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImR1bW15QGV4YW1wbGUuY29tIiwiZXhwIjo5MDA3MTkyNTQ3NDA5OTF9.sM2yd5GL9kK4h8IKHbr_fA2XmrzEsLOeLTIPrU0VfMg"

func TestSTACKITCLIFlow_Init(t *testing.T) {
	type args struct {
		cfg    *STACKITCLIFlowConfig
		confFn func(t *testing.T)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{
			cfg: &STACKITCLIFlowConfig{},
			confFn: func(t *testing.T) {
				_, err := RunSTACKITCLICommand(context.TODO(), "stackit auth activate-service-account --service-account-token="+testServiceAccountToken)
				if err != nil {
					t.Errorf("RunSTACKITCLICommand() error = %v", err)
					return
				}
			},
		}, false},
		{"no-token", args{
			cfg:    &STACKITCLIFlowConfig{},
			confFn: func(_ *testing.T) {},
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()

			c := &STACKITCLIFlow{}

			cliProfileName := "test-stackit-cli-flow-init" + tt.name

			_, _ = RunSTACKITCLICommand(ctx, fmt.Sprintf("stackit config profile delete %s -y", cliProfileName))
			_, err := RunSTACKITCLICommand(ctx, "stackit config profile create "+cliProfileName)
			if err != nil {
				t.Errorf("RunSTACKITCLICommand() error = %v", err)
				return
			}

			tt.args.confFn(t)

			defer func() {
				_, _ = RunSTACKITCLICommand(ctx, fmt.Sprintf("stackit config profile delete %s -y", cliProfileName))
			}()

			if err := c.Init(tt.args.cfg); err != nil {
				if (err != nil) != tt.wantErr {
					t.Errorf("TokenFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
				}

				return
			}

			if c.config == nil {
				t.Error("config is nil")
			}
		})
	}
}

func TestSTACKITCLIFlow_Do(t *testing.T) {
	type fields struct {
		client *http.Client
		config *STACKITCLIFlowConfig
	}
	type args struct{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"success", fields{&http.Client{}, &STACKITCLIFlowConfig{}}, args{}, http.StatusOK, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()

			_, _ = RunSTACKITCLICommand(ctx, "stackit config profile delete test-stackit-cli-flow-do -y")
			_, err := RunSTACKITCLICommand(ctx, "stackit config profile create test-stackit-cli-flow-do")
			if err != nil {
				t.Errorf("RunSTACKITCLICommand() error = %v", err)
				return
			}

			_, err = RunSTACKITCLICommand(ctx, "stackit auth activate-service-account --service-account-token="+testServiceAccountToken)
			if err != nil {
				t.Errorf("RunSTACKITCLICommand() error = %v", err)
				return
			}

			defer func() {
				_, _ = RunSTACKITCLICommand(ctx, "stackit config profile delete test-stackit-cli-flow-do -y")
			}()

			c := &STACKITCLIFlow{}
			err = c.Init(tt.fields.config)
			if err != nil {
				t.Errorf("Init() error = %v", err)
				return
			}

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				authorization := r.Header.Get("Authorization")
				if authorization != "Bearer "+testServiceAccountToken {
					w.WriteHeader(http.StatusUnauthorized)
					_, _ = fmt.Fprintln(w, `{"error":"missing authorization header"}`)
					return
				}

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
				t.Errorf("STACKITCLIFlow.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.StatusCode != tt.want {
				t.Errorf("STACKITCLIFlow.Do() = %v, want %v", got.StatusCode, tt.want)
			}
		})
	}
}
