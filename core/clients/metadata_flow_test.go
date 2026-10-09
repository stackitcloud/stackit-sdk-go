package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

func TestMetadataFlowInit(t *testing.T) {
	tests := []struct {
		name                string
		serviceAccountEmail string
		emailAsEnv          bool
		metadataUrl         string
		wantErr             bool
	}{
		{
			name:                "ok setting all",
			serviceAccountEmail: "test@sa.stackit.cloud",
			metadataUrl:         "http://localhost:8080",
		},
		{
			name:                "ok using defaults",
			serviceAccountEmail: "test@sa.stackit.cloud",
			emailAsEnv:          true,
		},
		{
			name:    "missing service account email",
			wantErr: true,
		},
		{
			name:                "invalid metadata url",
			serviceAccountEmail: "test@sa.stackit.cloud",
			metadataUrl:         "not a url",
			wantErr:             true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flowConfig := &MetadataFlowConfig{MetadataUrl: tt.metadataUrl}
			if tt.emailAsEnv {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_EMAIL", tt.serviceAccountEmail)
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_EMAIL", "")
				flowConfig.ServiceAccountEmail = tt.serviceAccountEmail
			}

			flow := &MetadataFlow{}
			err := flow.Init(flowConfig)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if flow.GetConfig().ServiceAccountEmail != tt.serviceAccountEmail {
				t.Fatalf("service account email = %s, want %s", flow.GetConfig().ServiceAccountEmail, tt.serviceAccountEmail)
			}
			if tt.metadataUrl == "" && flow.GetConfig().MetadataUrl != defaultMetadataUrl {
				t.Fatalf("metadata url = %s, want %s", flow.GetConfig().MetadataUrl, defaultMetadataUrl)
			}
		})
	}
}

func TestMetadataFlowRoundTrip(t *testing.T) {
	serviceAccountEmail := "test@sa.stackit.cloud"
	tests := []struct {
		name               string
		metadataStatus     int
		validFor           time.Duration
		omitValidUntil     bool
		wantMetadataCalls  int32
		wantErr            bool
		wantErrContains    string
		wantOpenAPIErrCode int
	}{
		{
			name:              "token is reused while valid",
			metadataStatus:    http.StatusOK,
			validFor:          time.Hour,
			wantMetadataCalls: 1,
		},
		{
			name:              "token is renewed when about to expire",
			metadataStatus:    http.StatusOK,
			validFor:          4 * time.Minute,
			wantMetadataCalls: 2,
		},
		{
			name:              "response without validUntil",
			metadataStatus:    http.StatusOK,
			validFor:          time.Hour,
			omitValidUntil:    true,
			wantMetadataCalls: 1,
			wantErr:           true,
		},
		{
			name:              "service account not attached",
			metadataStatus:    http.StatusNotFound,
			wantMetadataCalls: 1,
			wantErr:           true,
			wantErrContains:   "is not attached to this server",
		},
		{
			name:               "metadata service error",
			metadataStatus:     http.StatusInternalServerError,
			wantMetadataCalls:  1,
			wantErr:            true,
			wantOpenAPIErrCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var metadataCalls atomic.Int32
			metadataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				n := metadataCalls.Add(1)
				if r.Method != http.MethodGet || r.URL.Path != "/stackit/v1/service-accounts/"+serviceAccountEmail+"/token" {
					t.Errorf("unexpected metadata request: %s %s", r.Method, r.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.metadataStatus)
				if tt.metadataStatus != http.StatusOK {
					_, _ = w.Write([]byte(`{"code":"error"}`))
					return
				}
				response := map[string]string{"token": fmt.Sprintf("token-%d", n)}
				if !tt.omitValidUntil {
					response["validUntil"] = time.Now().Add(tt.validFor).UTC().Format(time.RFC3339Nano)
				}
				if err := json.NewEncoder(w).Encode(response); err != nil {
					t.Errorf("writing response: %s", err)
				}
			}))
			t.Cleanup(metadataServer.Close)

			var wantBearer atomic.Int32
			apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				want := fmt.Sprintf("Bearer token-%d", wantBearer.Load())
				if got := r.Header.Get("Authorization"); got != want {
					t.Errorf("authorization header = %q, want %q", got, want)
				}
				w.WriteHeader(http.StatusOK)
			}))
			t.Cleanup(apiServer.Close)

			flow := &MetadataFlow{}
			err := flow.Init(&MetadataFlowConfig{
				ServiceAccountEmail: serviceAccountEmail,
				MetadataUrl:         metadataServer.URL,
			})
			if err != nil {
				t.Fatalf("Init() error = %v", err)
			}
			client := &http.Client{Transport: flow}

			for i := 1; i <= 2; i++ {
				wantBearer.Store(min(int32(i), tt.wantMetadataCalls))
				req, err := http.NewRequest(http.MethodGet, apiServer.URL, http.NoBody)
				if err != nil {
					t.Fatalf("creating request: %s", err)
				}
				res, err := client.Do(req)
				if err == nil {
					_ = res.Body.Close()
				}
				if (err != nil) != tt.wantErr {
					t.Fatalf("request error = %v, wantErr %v", err, tt.wantErr)
				}
				if !tt.wantErr {
					continue
				}
				if tt.wantErrContains != "" && !strings.Contains(err.Error(), tt.wantErrContains) {
					t.Fatalf("error = %v, want it to contain %q", err, tt.wantErrContains)
				}
				oapiErr := &oapierror.GenericOpenAPIError{}
				if tt.wantOpenAPIErrCode != 0 && (!errors.As(err, &oapiErr) || oapiErr.StatusCode != tt.wantOpenAPIErrCode) {
					t.Fatalf("error = %v, want an API error with status %d", err, tt.wantOpenAPIErrCode)
				}
				break
			}

			if got := metadataCalls.Load(); got != tt.wantMetadataCalls {
				t.Fatalf("metadata calls = %d, want %d", got, tt.wantMetadataCalls)
			}
		})
	}
}
