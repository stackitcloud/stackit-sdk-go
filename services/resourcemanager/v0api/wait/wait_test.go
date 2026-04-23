package wait

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	resourcemanager "github.com/stackitcloud/stackit-sdk-go/services/resourcemanager/v0api"
)

type mockSettings struct {
	getFails     bool
	getNotFound  bool
	projectState resourcemanager.LifecycleState
}

func newAPIMock(settings []mockSettings) resourcemanager.DefaultAPI {
	count := 0
	return &resourcemanager.DefaultAPIServiceMock{
		GetProjectExecuteMock: utils.Ptr(func(_ resourcemanager.ApiGetProjectRequest) (*resourcemanager.GetProjectResponse, error) {
			setting := settings[count%len(settings)]
			count++

			if setting.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusInternalServerError,
				}
			}

			if setting.getNotFound {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusNotFound,
				}
			}

			return &resourcemanager.GetProjectResponse{
				LifecycleState: setting.projectState,
				ContainerId:    "cid",
			}, nil
		}),
	}
}

func TestCreateProjectWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		mockSettings     []mockSettings
		wantProjectState resourcemanager.LifecycleState
		wantErr          bool
		wantResp         bool
	}{
		{
			desc: "create_succeeded",
			mockSettings: []mockSettings{
				{projectState: resourcemanager.LIFECYCLESTATE_ACTIVE},
			},
			wantProjectState: resourcemanager.LIFECYCLESTATE_ACTIVE,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc: "creating",
			mockSettings: []mockSettings{
				{
					projectState: resourcemanager.LIFECYCLESTATE_CREATING,
				},
				{
					projectState: resourcemanager.LIFECYCLESTATE_CREATING,
				},
				{
					projectState: resourcemanager.LIFECYCLESTATE_ACTIVE,
				},
			},
			wantProjectState: resourcemanager.LIFECYCLESTATE_ACTIVE,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc: "get_fails",
			mockSettings: []mockSettings{
				{
					projectState: resourcemanager.LIFECYCLESTATE_CREATING,
				},
				{
					projectState: resourcemanager.LIFECYCLESTATE_CREATING,
				},
				{
					getFails: true,
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "unknown_state",
			mockSettings: []mockSettings{
				{
					getFails:     false,
					projectState: resourcemanager.LifecycleState("ANOTHER STATE"),
				},
			},
			wantErr:  true,
			wantResp: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(tt.mockSettings)

				var wantRes *resourcemanager.GetProjectResponse
				if tt.wantResp {
					wantRes = &resourcemanager.GetProjectResponse{
						LifecycleState: tt.wantProjectState,
						ContainerId:    "cid",
					}
				}

				handler := CreateProjectWaitHandler(context.Background(), apiClient, "cid")

				gotRes, err := handler.WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestDeleteProjectWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		mockSettings []mockSettings
		wantErr      bool
	}{
		{
			desc: "delete_succeeded",
			mockSettings: []mockSettings{
				{
					projectState: resourcemanager.LifecycleState(""),
					getFails:     false,
					getNotFound:  true,
				},
			},
			wantErr: false,
		},
		{
			desc: "delete_pending",
			mockSettings: []mockSettings{
				{
					getFails:     false,
					getNotFound:  false,
					projectState: resourcemanager.LIFECYCLESTATE_DELETING,
				},
				{
					getFails:     false,
					getNotFound:  false,
					projectState: resourcemanager.LIFECYCLESTATE_DELETING,
				},
				{
					getFails:     false,
					getNotFound:  true,
					projectState: resourcemanager.LifecycleState(""),
				},
			},
			wantErr: false,
		},
		{
			desc: "get_fails",
			mockSettings: []mockSettings{
				{
					projectState: resourcemanager.LifecycleState(""),
					getFails:     true,
				},
			},
			wantErr: true,
		},
		{
			desc: "timeout",
			mockSettings: []mockSettings{
				{
					getFails:     false,
					projectState: "ANOTHER STATE",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(tt.mockSettings)

				handler := DeleteProjectWaitHandler(context.Background(), apiClient, "cid")

				_, err := handler.WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}
