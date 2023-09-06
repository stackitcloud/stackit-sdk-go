package resourcemanager

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
)

type apiClientMocked struct {
	getFails     bool
	getNotFound  bool
	projectState LifecycleState
}

func (a *apiClientMocked) GetProjectExecute(_ context.Context, _ string) (*ProjectResponseWithParents, error) {
	if a.getFails {
		return nil, &GenericOpenAPIError{
			statusCode: http.StatusInternalServerError,
		}
	}

	if a.getNotFound {
		return nil, &GenericOpenAPIError{
			statusCode: http.StatusNotFound,
		}
	}

	return &ProjectResponseWithParents{
		LifecycleState: &a.projectState,
		ContainerId:    utils.Ptr("cid"),
	}, nil
}

func TestHandleError(t *testing.T) {
	tests := []struct {
		desc     string
		reqErr   error
		wantRes  interface{}
		wantDone bool
		wantErr  bool
	}{
		{
			desc: "handle_oapi_error",
			reqErr: &GenericOpenAPIError{
				statusCode: 500,
			},
			wantRes:  nil,
			wantDone: false,
			wantErr:  true,
		},
		{
			desc:     "not_generic_oapi_error",
			reqErr:   fmt.Errorf("some error"),
			wantRes:  nil,
			wantDone: false,
			wantErr:  true,
		},
		{
			desc: "bad_gateway_error",
			reqErr: &GenericOpenAPIError{
				statusCode: http.StatusBadGateway,
			},
			wantRes:  nil,
			wantDone: false,
			wantErr:  false,
		},
		{
			desc: "gateway_timeout_error",
			reqErr: &GenericOpenAPIError{
				statusCode: http.StatusBadGateway,
			},
			wantRes:  nil,
			wantDone: false,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotRes, gotDone, err := handleError(tt.reqErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotRes, tt.wantRes) {
				t.Errorf("handleError() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotDone != tt.wantDone {
				t.Errorf("handleError() gotDone = %v, want %v", gotDone, tt.wantDone)
			}
		})
	}
}

func TestCreateProjectWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		getFails     bool
		projectState LifecycleState
		wantErr      bool
	}{
		{
			desc:         "create_succeeded",
			getFails:     false,
			projectState: ActiveState,
			wantErr:      false,
		},
		{
			desc:         "creating",
			getFails:     false,
			projectState: CreatingState,
			wantErr:      true,
		},
		{
			desc:         "get_fails",
			getFails:     true,
			projectState: LifecycleState(""),
			wantErr:      true,
		},
		{
			desc:         "unknown_state",
			getFails:     false,
			projectState: LifecycleState("ANOTHER STATE"),
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:     tt.getFails,
				projectState: tt.projectState,
			}

			var wantRes *ProjectResponseWithParents
			if !tt.getFails {
				wantRes = &ProjectResponseWithParents{
					LifecycleState: &tt.projectState,
					ContainerId:    utils.Ptr("cid"),
				}
			} else {
				wantRes = nil
			}

			handler := CreateProjectWaitHandler(context.Background(), apiClient, "cid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes) {
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
		projectState LifecycleState
		wantErr      bool
	}{
		{
			desc:         "delete_succeeded",
			getFails:     false,
			getNotFound:  true,
			projectState: LifecycleState(""),
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

			var wantRes *ProjectResponseWithParents
			if !tt.getFails && !tt.getNotFound {
				wantRes = &ProjectResponseWithParents{
					LifecycleState: &tt.projectState,
					ContainerId:    utils.Ptr("cid"),
				}
			} else {
				wantRes = nil
			}

			handler := DeleteProjectWaitHandler(context.Background(), apiClient, "cid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}
