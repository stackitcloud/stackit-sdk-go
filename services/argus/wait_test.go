package argus

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
)

type apiClientMocked struct {
	getFails      bool
	resourceState *string
	jobs          []Job
}

func (a *apiClientMocked) GetInstanceExecute(_ context.Context, _, _ string) (*InstanceResponse, error) {
	if a.getFails {
		return nil, &GenericOpenAPIError{
			statusCode: 500,
		}
	}

	return &InstanceResponse{
		Id:     utils.Ptr("iid"),
		Status: a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetScrapeConfigsExecute(_ context.Context, _, _ string) (*ScrapeConfigsResponse, error) {
	if a.getFails {
		return nil, &GenericOpenAPIError{
			statusCode: 500,
		}
	}

	return &ScrapeConfigsResponse{
		Data: &a.jobs,
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

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState *string
		wantErr       bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(CreateSuccess),
			wantErr:       false,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: utils.Ptr(CreateFail),
			wantErr:       true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(""),
			wantErr:       true,
		},
		{
			desc:          "broken_response",
			getFails:      false,
			resourceState: nil,
			wantErr:       true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr("ANOTHER STATE"),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *InstanceResponse
			if !tt.getFails {
				wantRes = &InstanceResponse{
					Id:     utils.Ptr("iid"),
					Status: tt.resourceState,
				}
			} else {
				wantRes = nil
			}

			handler := CreateInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(NullableString{})) {
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
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(UpdateSuccess),
			wantErr:       false,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: utils.Ptr(UpdateFail),
			wantErr:       true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(""),
			wantErr:       true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr("ANOTHER STATE"),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *InstanceResponse
			if !tt.getFails {
				wantRes = &InstanceResponse{
					Status: tt.resourceState,
					Id:     utils.Ptr("iid"),
				}
			} else {
				wantRes = nil
			}

			handler := UpdateInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(NullableString{})) {
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
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: utils.Ptr(DeleteSuccess),
			wantErr:       false,
		},
		{
			desc:          "delete_failed",
			getFails:      false,
			resourceState: utils.Ptr(DeleteFail),
			wantErr:       true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: utils.Ptr(""),
			wantErr:       true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: utils.Ptr("ANOTHER STATE"),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *InstanceResponse
			if !tt.getFails {
				wantRes = &InstanceResponse{
					Status: tt.resourceState,
					Id:     utils.Ptr("iid"),
				}
			} else {
				wantRes = nil
			}

			handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "iid", "pid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestCreateScrapeConfigWaitHandler(t *testing.T) {
	tests := []struct {
		desc     string
		getFails bool
		jobs     []Job
		wantErr  bool
	}{
		{
			desc:     "create_succeeded",
			getFails: false,
			jobs:     []Job{{JobName: utils.Ptr("job")}, {JobName: utils.Ptr("other-job")}},
			wantErr:  false,
		},
		{
			desc:     "create_failed and timeout",
			getFails: false,
			jobs:     []Job{{JobName: utils.Ptr("other-job")}},
			wantErr:  true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			jobs:     []Job{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails: tt.getFails,
				jobs:     tt.jobs,
			}

			var wantRes *ScrapeConfigsResponse
			if !tt.getFails {
				wantRes = &ScrapeConfigsResponse{
					Data: &tt.jobs,
				}
			} else {
				wantRes = nil
			}

			handler := CreateScrapeConfigWaitHandler(context.Background(), apiClient, "", "job", "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteScrapeConfigWaitHandler(t *testing.T) {
	tests := []struct {
		desc     string
		getFails bool
		jobs     []Job
		wantErr  bool
	}{
		{
			desc:     "delete_succeeded",
			getFails: false,
			jobs:     []Job{{JobName: utils.Ptr("other-job")}},
			wantErr:  false,
		},
		{
			desc:     "timeout",
			getFails: false,
			jobs:     []Job{{JobName: utils.Ptr("job")}},
			wantErr:  true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			jobs:     []Job{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails: tt.getFails,
				jobs:     tt.jobs,
			}

			var wantRes *ScrapeConfigsResponse
			if !tt.getFails {
				wantRes = &ScrapeConfigsResponse{
					Data: &tt.jobs,
				}
			} else {
				wantRes = nil
			}

			handler := DeleteScrapeConfigWaitHandler(context.Background(), apiClient, "", "job", "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes, cmpopts.IgnoreUnexported(NullableString{})) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}
