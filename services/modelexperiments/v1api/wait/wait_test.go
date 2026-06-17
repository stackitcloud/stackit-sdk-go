package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	modelexperiments "github.com/stackitcloud/stackit-sdk-go/services/modelexperiments/v1api"
)

type mockSettings struct {
	getFails      bool
	statusCode    int
	resourceState string
}

func newAPIMock(settings *mockSettings) modelexperiments.DefaultAPI {
	return &modelexperiments.DefaultAPIServiceMock{
		GetInstanceExecuteMock: utils.Ptr(func(_ modelexperiments.ApiGetInstanceRequest) (*modelexperiments.GetInstanceResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.statusCode,
				}
			}

			return &modelexperiments.GetInstanceResponse{
				Instance: modelexperiments.Instance{
					State: modelexperiments.InstanceState(settings.resourceState),
				},
			}, nil
		}),
		GetInstanceTokenExecuteMock: utils.Ptr(func(_ modelexperiments.ApiGetInstanceTokenRequest) (*modelexperiments.GetInstanceTokenResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: settings.statusCode,
				}
			}

			return &modelexperiments.GetInstanceTokenResponse{
				Token: modelexperiments.TokenMetadata{
					State: modelexperiments.TokenState(settings.resourceState),
				},
			}, nil
		}),
	}
}

func TestCreateModelExperimentsInstanceWaitHandler(t *testing.T) {
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
			statusCode:    202,
			resourceState: string(modelexperiments.INSTANCESTATE_ACTIVE),
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
			desc:          "not_found",
			getFails:      true,
			statusCode:    404,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "impaired",
			getFails:      false,
			statusCode:    200,
			resourceState: string(modelexperiments.INSTANCESTATE_IMPAIRED),
			wantErr:       true,
			wantResp:      true,
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

				var wantRes *modelexperiments.GetInstanceResponse
				if tt.wantResp {
					wantRes = &modelexperiments.GetInstanceResponse{
						Instance: modelexperiments.Instance{
							State: modelexperiments.InstanceState(tt.resourceState),
						},
					}
				}

				handler := CreateModelExperimentsInstanceWaitHandler(context.Background(), apiClient, "region", "pid", "instanceId")

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

func TestDeleteModelExperimentsInstanceWaitHandler(t *testing.T) {
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
			resourceState: string(modelexperiments.INSTANCESTATE_DELETING),
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

				var wantRes *modelexperiments.GetInstanceResponse
				if tt.wantResp {
					wantRes = &modelexperiments.GetInstanceResponse{
						Instance: modelexperiments.Instance{
							State: modelexperiments.InstanceState(tt.resourceState),
						},
					}
				}

				handler := DeleteModelExperimentsInstanceWaitHandler(context.Background(), apiClient, "region", "pid", "instanceId")

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

func TestCreateModelExperimentsInstanceTokenWaitHandler(t *testing.T) {
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
			statusCode:    202,
			resourceState: string(modelexperiments.TOKENSTATE_ACTIVE),
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
			desc:          "not_found",
			getFails:      true,
			statusCode:    404,
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

				var wantRes *modelexperiments.GetInstanceTokenResponse
				if tt.wantResp {
					wantRes = &modelexperiments.GetInstanceTokenResponse{
						Token: modelexperiments.TokenMetadata{
							State: modelexperiments.TokenState(tt.resourceState),
						},
					}
				}

				handler := CreateModelExperimentsInstanceTokenWaitHandler(context.Background(), apiClient, "region", "pid", "instanceId", "tid")

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

func TestDeleteModelExperimentsInstanceTokenWaitHandler(t *testing.T) {
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
			resourceState: string(modelexperiments.TOKENSTATE_DELETING),
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

				var wantRes *modelexperiments.GetInstanceTokenResponse
				if tt.wantResp {
					wantRes = &modelexperiments.GetInstanceTokenResponse{
						Token: modelexperiments.TokenMetadata{
							State: modelexperiments.TokenState(tt.resourceState),
						},
					}
				}

				handler := DeleteModelExperimentsInstanceTokenWaitHandler(context.Background(), apiClient, "region", "pid", "instanceId", "tid")

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
