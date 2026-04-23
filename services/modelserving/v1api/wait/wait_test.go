package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	modelserving "github.com/stackitcloud/stackit-sdk-go/services/modelserving/v1api"
)

type mockSettings struct {
	getFails      bool
	resourceState string
	statusCode    int
}

func newAPIMock(settings *mockSettings) modelserving.DefaultAPI {
	return &modelserving.DefaultAPIServiceMock{
		GetTokenExecuteMock: utils.Ptr(func(_ modelserving.ApiGetTokenRequest) (*modelserving.GetTokenResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.statusCode,
				}
			}

			return &modelserving.GetTokenResponse{
				Token: modelserving.Token{
					State: settings.resourceState,
					Id:    "tid",
				},
			}, nil
		}),
	}
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
			resourceState: TOKENSTATE_ACTIVE,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getFails:      tt.getFails,
					statusCode:    tt.statusCode,
					resourceState: tt.resourceState,
				})

				var wantRes *modelserving.GetTokenResponse
				if tt.wantResp {
					wantRes = &modelserving.GetTokenResponse{
						Token: modelserving.Token{
							State: tt.resourceState,
							Id:    "tid",
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
			resourceState: TOKENSTATE_ACTIVE,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getFails:      tt.getFails,
					statusCode:    tt.statusCode,
					resourceState: tt.resourceState,
				})

				var wantRes *modelserving.GetTokenResponse
				if tt.wantResp {
					wantRes = &modelserving.GetTokenResponse{
						Token: modelserving.Token{
							State: tt.resourceState,
							Id:    "tid",
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
			resourceState: TOKENSTATE_DELETING,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getFails:      tt.getFails,
					statusCode:    tt.statusCode,
					resourceState: tt.resourceState,
				})

				var wantRes *modelserving.GetTokenResponse
				if tt.wantResp {
					wantRes = &modelserving.GetTokenResponse{
						Token: modelserving.Token{
							State: tt.resourceState,
							Id:    "tid",
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
		})
	}
}
