package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/modelserving"
)

type apiClientMocked struct {
	getFails      bool
	resourceState string
	statusCode    int
}

func (a *apiClientMocked) GetTokenExecute(_ context.Context, _, _, _ string) (*modelserving.GetTokenResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: a.statusCode,
		}
	}

	return &modelserving.GetTokenResponse{
		Token: &modelserving.Token{
			State: utils.Ptr(a.resourceState),
			Id:    utils.Ptr("tid"),
		},
	}, nil
}

func TestCreateModelServingWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		statusCode    int
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			statusCode:    200,
			resourceState: activeState,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			statusCode:    500,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			statusCode:    200,
			resourceState: "ANOTHER_STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				statusCode:    tt.statusCode,
				resourceState: tt.resourceState,
			}

			var wantRes *modelserving.GetTokenResponse
			if tt.wantResp {
				wantRes = &modelserving.GetTokenResponse{
					Token: &modelserving.Token{
						State: utils.Ptr(tt.resourceState),
						Id:    utils.Ptr("tid"),
					},
				}
			}

			handler := CreateModelServingWaitHandler(context.Background(), apiClient, "region", "pid", "tid")

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

func TestUpdateModelServingWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		statusCode    int
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			statusCode:    200,
			resourceState: activeState,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			statusCode:    500,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			statusCode:    200,
			resourceState: "ANOTHER_STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				statusCode:    tt.statusCode,
				resourceState: tt.resourceState,
			}

			var wantRes *modelserving.GetTokenResponse
			if tt.wantResp {
				wantRes = &modelserving.GetTokenResponse{
					Token: &modelserving.Token{
						State: utils.Ptr(tt.resourceState),
						Id:    utils.Ptr("tid"),
					},
				}
			}

			handler := UpdateModelServingWaitHandler(context.Background(), apiClient, "region", "pid", "tid")

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

func TestDeleteModelServingWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		statusCode    int
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      true,
			statusCode:    404,
			resourceState: "",
			wantErr:       false,
			wantResp:      false,
		},
		{
			desc:          "delete_in_progress",
			getFails:      false,
			statusCode:    200,
			resourceState: "DELETING",
			wantErr:       true, // Should timeout since delete is not complete
			wantResp:      false,
		},
		{
			desc:          "get_fails_with_other_error",
			getFails:      true,
			statusCode:    500,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				statusCode:    tt.statusCode,
				resourceState: tt.resourceState,
			}

			var wantRes *modelserving.GetTokenResponse
			if tt.wantResp {
				wantRes = &modelserving.GetTokenResponse{
					Token: &modelserving.Token{
						State: utils.Ptr(tt.resourceState),
						Id:    utils.Ptr("tid"),
					},
				}
			}

			handler := DeleteModelServingWaitHandler(context.Background(), apiClient, "region", "pid", "tid")

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
