package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/observability"
)

type apiClientMocked struct {
	getFails      bool
	resourceState *observability.GetInstanceResponseStatus
	jobs          []observability.Job
}

func (a *apiClientMocked) GetInstanceExecute(_ context.Context, _, _ string) (*observability.GetInstanceResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &observability.GetInstanceResponse{
		Id:     utils.Ptr("iid"),
		Status: a.resourceState,
	}, nil
}

func (a *apiClientMocked) ListScrapeConfigsExecute(_ context.Context, _, _ string) (*observability.ListScrapeConfigsResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &observability.ListScrapeConfigsResponse{
		Data: &a.jobs,
	}, nil
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *observability.GetInstanceResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(observability.GETINSTANCERESPONSESTATUS_CREATE_SUCCEEDED),
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: utils.Ptr(observability.GETINSTANCERESPONSESTATUS_CREATE_FAILED),
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(observability.GetInstanceResponseStatus("")),
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "broken_response",
			getFails:      false,
			resourceState: nil,
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr(observability.GetInstanceResponseStatus("ANOTHER STATE")),
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *observability.GetInstanceResponse
			if tt.wantResp {
				wantRes = &observability.GetInstanceResponse{
					Id:     utils.Ptr("iid"),
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
	}
}

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *observability.GetInstanceResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(observability.GETINSTANCERESPONSESTATUS_UPDATE_SUCCEEDED),
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: utils.Ptr(observability.GETINSTANCERESPONSESTATUS_UPDATE_FAILED),
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(observability.GetInstanceResponseStatus("")),
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr(observability.GetInstanceResponseStatus("ANOTHER STATE")),
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *observability.GetInstanceResponse
			if tt.wantResp {
				wantRes = &observability.GetInstanceResponse{
					Status: tt.resourceState,
					Id:     utils.Ptr("iid"),
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
	}
}

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *observability.GetInstanceResponseStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(observability.GETINSTANCERESPONSESTATUS_DELETE_SUCCEEDED),
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "delete_failed",
			getFails:      false,
			resourceState: utils.Ptr(observability.GETINSTANCERESPONSESTATUS_DELETE_FAILED),
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(observability.GetInstanceResponseStatus("")),
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr(observability.GetInstanceResponseStatus("ANOTHER STATE")),
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *observability.GetInstanceResponse
			if tt.wantResp {
				wantRes = &observability.GetInstanceResponse{
					Status: tt.resourceState,
					Id:     utils.Ptr("iid"),
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
			jobs:     []observability.Job{{JobName: utils.Ptr("job")}, {JobName: utils.Ptr("other-job")}},
			wantErr:  false,
			wantResp: true,
		},
		{
			desc:     "create_failed and timeout",
			getFails: false,
			jobs:     []observability.Job{{JobName: utils.Ptr("other-job")}},
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
			apiClient := &apiClientMocked{
				getFails: tt.getFails,
				jobs:     tt.jobs,
			}

			var wantRes *observability.ListScrapeConfigsResponse
			if tt.wantResp {
				wantRes = &observability.ListScrapeConfigsResponse{
					Data: utils.Ptr(tt.jobs),
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
			jobs:     []observability.Job{{JobName: utils.Ptr("other-job")}},
			wantErr:  false,
			wantResp: true,
		},
		{
			desc:     "timeout",
			getFails: false,
			jobs:     []observability.Job{{JobName: utils.Ptr("job")}},
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
			apiClient := &apiClientMocked{
				getFails: tt.getFails,
				jobs:     tt.jobs,
			}

			var wantRes *observability.ListScrapeConfigsResponse
			if tt.wantResp {
				wantRes = &observability.ListScrapeConfigsResponse{
					Data: utils.Ptr(tt.jobs),
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
	}
}
