package wait

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	registry "github.com/stackitcloud/stackit-sdk-go/services/registry/v1api"
)

const (
	testProjectId     = "00000000-0000-0000-0000-000000000000"
	testRegionId      = "eu01"
	testArtifactoryId = "test-artifactory-id"
)

type response struct {
	state  string
	status int
}

type mockSettings struct {
	responseIndex int
	responses     []response
}

func newAPIMock(settings *mockSettings) registry.DefaultAPI {
	return &registry.DefaultAPIServiceMock{
		GetArtifactoryExecuteMock: utils.Ptr(func(_ registry.ApiGetArtifactoryRequest) (*registry.Artifactory, error) {
			resp := settings.responses[settings.responseIndex]
			settings.responseIndex++
			settings.responseIndex %= len(settings.responses)

			if resp.status >= http.StatusBadRequest {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode:   resp.status,
					ErrorMessage: "simulated error",
				}
			}

			return &registry.Artifactory{
				Id:    testArtifactoryId,
				State: resp.state,
			}, nil
		}),
	}
}

func TestCreateArtifactoryWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []response
		wantState string
		wantErr   bool
	}{
		{
			name: "success immediately",
			responses: []response{
				{state: ArtifactoryActive},
			},
			wantState: ArtifactoryActive,
			wantErr:   false,
		},
		{
			name: "success immediately healthy",
			responses: []response{
				{state: ArtifactoryActive},
			},
			wantState: ArtifactoryActive,
			wantErr:   false,
		},
		{
			name: "success delayed",
			responses: []response{
				{state: ArtifactoryReconcilling},
				{state: ArtifactoryReconcilling},
				{state: ArtifactoryActive},
			},
			wantState: ArtifactoryActive,
			wantErr:   false,
		},
		{
			name: "error state",
			responses: []response{
				{state: ArtifactoryError},
			},
			wantErr: true,
		},
		{
			name: "api error",
			responses: []response{
				{status: http.StatusInternalServerError},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				mock := newAPIMock(&mockSettings{responses: tt.responses})
				handler := CreateArtifactoryWaitHandler(context.Background(), mock, testProjectId, testRegionId, testArtifactoryId)
				handler.SetSleepBeforeWait(0)
				handler.SetThrottle(10 * time.Millisecond)

				got, err := handler.WaitWithContext(context.Background())
				if (err != nil) != tt.wantErr {
					t.Fatalf("WaitWithContext() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !tt.wantErr && got.State != tt.wantState {
					t.Fatalf("WaitWithContext() state = %v, want %v", got.State, tt.wantState)
				}
			})
		})
	}
}

func TestUpdateArtifactoryWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []response
		wantState string
		wantErr   bool
	}{
		{
			name: "success immediately",
			responses: []response{
				{state: ArtifactoryActive},
			},
			wantState: ArtifactoryActive,
			wantErr:   false,
		},
		{
			name: "success delayed with healthy state",
			responses: []response{
				{state: ArtifactoryReconcilling},
				{state: ArtifactoryActive},
			},
			wantState: ArtifactoryActive,
			wantErr:   false,
		},
		{
			name: "error state",
			responses: []response{
				{state: ArtifactoryError},
			},
			wantErr: true,
		},
		{
			name: "api error",
			responses: []response{
				{status: http.StatusInternalServerError},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				mock := newAPIMock(&mockSettings{responses: tt.responses})
				handler := UpdateArtifactoryWaitHandler(context.Background(), mock, testProjectId, testRegionId, testArtifactoryId)
				handler.SetSleepBeforeWait(0)
				handler.SetThrottle(10 * time.Millisecond)

				got, err := handler.WaitWithContext(context.Background())
				if (err != nil) != tt.wantErr {
					t.Fatalf("WaitWithContext() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !tt.wantErr && got.State != tt.wantState {
					t.Fatalf("WaitWithContext() state = %v, want %v", got.State, tt.wantState)
				}
			})
		})
	}
}

func TestGetStateArtifactory(t *testing.T) {
	tests := []struct {
		name    string
		resp    *registry.Artifactory
		want    string
		wantErr bool
	}{
		{
			name:    "nil response",
			resp:    nil,
			wantErr: true,
		},
		{
			name: "valid state",
			resp: &registry.Artifactory{
				State: ArtifactoryActive,
			},
			want:    ArtifactoryActive,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getStateArtifactory(tt.resp)
			if (err != nil) != tt.wantErr {
				t.Fatalf("getStateArtifactory() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("getStateArtifactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
