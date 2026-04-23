package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	observability "github.com/stackitcloud/stackit-sdk-go/services/observability/v1api"
)

type mockSettings struct {
	getFails      bool
	resourceState string
	jobs          []observability.Job
}

func newAPIMock(settings *mockSettings) observability.DefaultAPI {
	return &observability.DefaultAPIServiceMock{
		GetInstanceExecuteMock: utils.Ptr(func(_ observability.ApiGetInstanceRequest) (*observability.GetInstanceResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &observability.GetInstanceResponse{
				Id:     "iid",
				Status: settings.resourceState,
			}, nil
		}),
		ListScrapeConfigsExecuteMock: utils.Ptr(func(_ observability.ApiListScrapeConfigsRequest) (*observability.ListScrapeConfigsResponse, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &observability.ListScrapeConfigsResponse{
				Data: settings.jobs,
			}, nil
		}),
	}
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: GETINSTANCERESPONSESTATUS_CREATE_SUCCEEDED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: GETINSTANCERESPONSESTATUS_CREATE_FAILED,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "broken_response",
			getFails:      false,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *observability.GetInstanceResponse
				if tt.wantResp {
					wantRes = &observability.GetInstanceResponse{
						Id:     "iid",
						Status: tt.resourceState,
					}
				}

				handler := CreateInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(observability.NullableString{})) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: GETINSTANCERESPONSESTATUS_UPDATE_SUCCEEDED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: GETINSTANCERESPONSESTATUS_UPDATE_FAILED,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *observability.GetInstanceResponse
				if tt.wantResp {
					wantRes = &observability.GetInstanceResponse{
						Status: tt.resourceState,
						Id:     "iid",
					}
				}

				handler := UpdateInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(observability.NullableString{})) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: GETINSTANCERESPONSESTATUS_DELETE_SUCCEEDED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "delete_failed",
			getFails:      false,
			resourceState: GETINSTANCERESPONSESTATUS_DELETE_FAILED,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *observability.GetInstanceResponse
				if tt.wantResp {
					wantRes = &observability.GetInstanceResponse{
						Status: tt.resourceState,
						Id:     "iid",
					}
				}

				handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(observability.NullableString{})) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestCreateScrapeConfigWaitHandler(t *testing.T) {
	tests := []struct {
		desc     string
		getFails bool
		jobs     []observability.Job
		wantErr  bool
		wantResp bool
	}{
		{
			desc:     "create_succeeded",
			getFails: false,
			jobs: []observability.Job{
				{
					JobName: "job",
				},
				{
					JobName: "other-job",
				},
			},
			wantErr:  false,
			wantResp: true,
		},
		{
			desc:     "create_failed and timeout",
			getFails: false,
			jobs: []observability.Job{
				{
					JobName: "other-job",
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:     "get_fails",
			getFails: true,
			jobs:     []observability.Job{},
			wantErr:  true,
			wantResp: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getFails: tt.getFails,
					jobs:     tt.jobs,
				})

				var wantRes *observability.ListScrapeConfigsResponse
				if tt.wantResp {
					wantRes = &observability.ListScrapeConfigsResponse{
						Data: tt.jobs,
					}
				}

				handler := CreateScrapeConfigWaitHandler(context.Background(), apiClient, "", "job", "")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(observability.NullableString{})) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestDeleteScrapeConfigWaitHandler(t *testing.T) {
	tests := []struct {
		desc     string
		getFails bool
		jobs     []observability.Job
		wantErr  bool
		wantResp bool
	}{
		{
			desc:     "delete_succeeded",
			getFails: false,
			jobs: []observability.Job{
				{
					JobName: "other-job",
				},
			},
			wantErr:  false,
			wantResp: true,
		},
		{
			desc:     "timeout",
			getFails: false,
			jobs: []observability.Job{
				{
					JobName: "job",
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:     "get_fails",
			getFails: true,
			jobs:     []observability.Job{},
			wantErr:  true,
			wantResp: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(&mockSettings{
					getFails: tt.getFails,
					jobs:     tt.jobs,
				})

				var wantRes *observability.ListScrapeConfigsResponse
				if tt.wantResp {
					wantRes = &observability.ListScrapeConfigsResponse{
						Data: tt.jobs,
					}
				}

				handler := DeleteScrapeConfigWaitHandler(context.Background(), apiClient, "", "job", "")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(observability.NullableString{})) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}
