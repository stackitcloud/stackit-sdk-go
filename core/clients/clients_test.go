package clients

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestNewRetryConfig(t *testing.T) {
	got := NewRetryConfig()
	want := RetryConfig{
		MaxRetries:       DefaultRetryMaxRetries,
		WaitBetweenCalls: DefaultRetryWaitBetweenCalls,
		RetryTimeout:     DefaultRetryTimeout,
		ClientTimeout:    DefaultClientTimeout,
	}
	if got == nil {
		t.Error("NewRetryConfig returned nil")
	} else if !reflect.DeepEqual(*got, want) {
		t.Errorf("%+v != %+v", *got, want)
	}
}

func TestDo(t *testing.T) {
	type args struct {
		cfg            *RetryConfig
		serverStatus   int
		serverResponse string
	}
	tests := []struct {
		name     string
		args     args
		wantResp *http.Response
		wantErr  bool
		errMsg   string
	}{
		{"all ok", args{
			cfg:            &RetryConfig{0, time.Microsecond, time.Second, DefaultClientTimeout},
			serverStatus:   http.StatusOK,
			serverResponse: `{"status":"ok", "testing": "%s"}`,
		}, &http.Response{StatusCode: http.StatusOK}, false, ""},
		{"all ok nil client", args{
			cfg:            &RetryConfig{0, time.Microsecond, time.Second, DefaultClientTimeout},
			serverStatus:   http.StatusOK,
			serverResponse: `{"status":"ok", "testing": "%s"}`,
		}, &http.Response{StatusCode: http.StatusOK}, false, ""},
		{"fail 1", args{
			cfg:            &RetryConfig{1, time.Microsecond, time.Second, DefaultClientTimeout},
			serverStatus:   http.StatusInternalServerError,
			serverResponse: `{"status":"error 1", "testing": "%s"}`,
		}, &http.Response{StatusCode: http.StatusInternalServerError}, false, ""},
		{"fail 2 - timeout error", args{
			cfg:            &RetryConfig{3, time.Microsecond, time.Second, DefaultClientTimeout},
			serverStatus:   http.StatusOK,
			serverResponse: `{"status":"ok", "testing": "%s"}`,
		}, &http.Response{StatusCode: http.StatusOK}, true, "no such host"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.args.serverStatus)
				fmt.Fprintln(w, fmt.Sprintf(tt.args.serverResponse, tt.name))
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
			c := &http.Client{}
			if tt.name == "fail 2 - timeout error" && server != nil {
				t.Log("closing server")
				server.Close()
			}
			if tt.name == "all ok nil client" {
				c = nil
			}
			gotResp, err := Do(c, req, tt.args.cfg)
			if err == nil {
				// Defer discard and close the body
				defer func() {
					if _, discardErr := io.Copy(io.Discard, gotResp.Body); discardErr != nil && err == nil {
						err = discardErr
					}
					if closeErr := gotResp.Body.Close(); closeErr != nil && err == nil {
						err = closeErr
					}
				}()
			}
			if server != nil {
				server.Close()
			}
			if (err != nil) != tt.wantErr && (err != nil && !strings.Contains(err.Error(), tt.errMsg)) {
				body := []byte{}
				if gotResp != nil {
					body, err = io.ReadAll(gotResp.Body)
					if err != nil {
						t.Error(err)
						return
					}
				}
				t.Errorf("do() error = %v, wantErr %v, got: %s", err, tt.wantErr, string(body))
				return
			}
			if gotResp != nil && tt.wantResp.StatusCode != gotResp.StatusCode {
				t.Errorf("do() status code = %v, want %v", tt.wantResp.StatusCode, gotResp.StatusCode)
			}
		})
	}
}
