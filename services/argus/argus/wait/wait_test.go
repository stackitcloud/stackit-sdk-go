package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/argus"
)

type apiClientMocked struct {
	getFails      bool
	resourceState *string
	jobs          []argus.Job
}

func (a *apiClientMocked) GetInstanceExecute(_ context.Context, _, _ string) (*argus.GetInstanceResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &argus.GetInstanceResponse{
		Id:     utils.Ptr("iid"),
		Status: a.resourceState,
	}, nil
}

func (a *apiClientMocked) ListScrapeConfigsExecute(_ context.Context, _, _ string) (*argus.ListScrapeConfigsResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &argus.ListScrapeConfigsResponse{
		Data: &a.jobs,
	}, nil
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(CreateSuccess),
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: utils.Ptr(CreateFail),
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(""),
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
			resourceState: utils.Ptr("ANOTHER STATE"),
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

			var wantRes *argus.GetInstanceResponse
			if tt.wantResp {
				wantRes = &argus.GetInstanceResponse{
					Id:     utils.Ptr("iid"),
					Status: tt.resourceState,
				}
			}

			handler := CreateInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(UpdateSuccess),
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: utils.Ptr(UpdateFail),
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(""),
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr("ANOTHER STATE"),
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

			var wantRes *argus.GetInstanceResponse
			if tt.wantResp {
				wantRes = &argus.GetInstanceResponse{
					Status: tt.resourceState,
					Id:     utils.Ptr("iid"),
				}
			}

			handler := UpdateInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(DeleteSuccess),
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "delete_failed",
			getFails:      false,
			resourceState: utils.Ptr(DeleteFail),
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(""),
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr("ANOTHER STATE"),
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

			var wantRes *argus.GetInstanceResponse
			if tt.wantResp {
				wantRes = &argus.GetInstanceResponse{
					Status: tt.resourceState,
					Id:     utils.Ptr("iid"),
				}
			}

			handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestCreateScrapeConfigWaitHandler(t *testing.T) {
	tests := []struct {
		desc     string
		getFails bool
		jobs     []argus.Job
		wantErr  bool
		wantResp bool
	}{
		{
			desc:     "create_succeeded",
			getFails: false,
			jobs:     []argus.Job{{JobName: utils.Ptr("job")}, {JobName: utils.Ptr("other-job")}},
			wantErr:  false,
			wantResp: true,
		},
		{
			desc:     "create_failed and timeout",
			getFails: false,
			jobs:     []argus.Job{{JobName: utils.Ptr("other-job")}},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:     "get_fails",
			getFails: true,
			jobs:     []argus.Job{},
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

			var wantRes *argus.ListScrapeConfigsResponse
			if tt.wantResp {
				wantRes = &argus.ListScrapeConfigsResponse{
					Data: &tt.jobs,
				}
			}

			handler := CreateScrapeConfigWaitHandler(context.Background(), apiClient, "", "job", "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteScrapeConfigWaitHandler(t *testing.T) {
	tests := []struct {
		desc     string
		getFails bool
		jobs     []argus.Job
		wantErr  bool
		wantResp bool
	}{
		{
			desc:     "delete_succeeded",
			getFails: false,
			jobs:     []argus.Job{{JobName: utils.Ptr("other-job")}},
			wantErr:  false,
			wantResp: true,
		},
		{
			desc:     "timeout",
			getFails: false,
			jobs:     []argus.Job{{JobName: utils.Ptr("job")}},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:     "get_fails",
			getFails: true,
			jobs:     []argus.Job{},
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

			var wantRes *argus.ListScrapeConfigsResponse
			if tt.wantResp {
				wantRes = &argus.ListScrapeConfigsResponse{
					Data: &tt.jobs,
				}
			}

			handler := DeleteScrapeConfigWaitHandler(context.Background(), apiClient, "", "job", "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(argus.NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}
