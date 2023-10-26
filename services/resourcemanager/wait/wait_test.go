package wait

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/resourcemanager"
)

type apiClientMocked struct {
	getFails     bool
	getNotFound  bool
	projectState resourcemanager.LifecycleState
}

func (a *apiClientMocked) GetProjectExecute(_ context.Context, _ string) (*resourcemanager.ProjectResponseWithParents, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusInternalServerError,
		}
	}

	if a.getNotFound {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: http.StatusNotFound,
		}
	}

	return &resourcemanager.ProjectResponseWithParents{
		LifecycleState: &a.projectState,
		ContainerId:    utils.Ptr("cid"),
	}, nil
}

func TestCreateProjectWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		getFails     bool
		projectState resourcemanager.LifecycleState
		wantErr      bool
		wantResp     bool
	}{
		{
			desc:         "create_succeeded",
			getFails:     false,
			projectState: ActiveState,
			wantErr:      false,
			wantResp:     true,
		},
		{
			desc:         "creating",
			getFails:     false,
			projectState: CreatingState,
			wantErr:      true,
			wantResp:     false,
		},
		{
			desc:         "get_fails",
			getFails:     true,
			projectState: resourcemanager.LifecycleState(""),
			wantErr:      true,
			wantResp:     false,
		},
		{
			desc:         "unknown_state",
			getFails:     false,
			projectState: resourcemanager.LifecycleState("ANOTHER STATE"),
			wantErr:      true,
			wantResp:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:     tt.getFails,
				projectState: tt.projectState,
			}

			var wantRes *resourcemanager.ProjectResponseWithParents
			if tt.wantResp {
				wantRes = &resourcemanager.ProjectResponseWithParents{
					LifecycleState: &tt.projectState,
					ContainerId:    utils.Ptr("cid"),
				}
			}

			handler := CreateProjectWaitHandler(context.Background(), apiClient, "cid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteProjectWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		getFails     bool
		getNotFound  bool
		projectState resourcemanager.LifecycleState
		wantErr      bool
	}{
		{
			desc:         "delete_succeeded",
			getFails:     false,
			getNotFound:  true,
			projectState: resourcemanager.LifecycleState(""),
			wantErr:      false,
		},
		{
			desc:         "get_fails",
			getFails:     true,
			projectState: "",
			wantErr:      true,
		},
		{
			desc:         "timeout",
			getFails:     false,
			projectState: "ANOTHER STATE",
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:     tt.getFails,
				getNotFound:  tt.getNotFound,
				projectState: tt.projectState,
			}

			handler := DeleteProjectWaitHandler(context.Background(), apiClient, "cid")

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
