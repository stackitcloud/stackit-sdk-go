package wait

import (
	"context"
	"testing"
	"testing/synctest"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	workflows "github.com/stackitcloud/stackit-sdk-go/services/workflows/v1alphaapi"
)

type mockSettings struct {
	getFails      bool
	notFound      bool
	resourceState workflows.InstanceStatus
}

func newAPIMock(settings mockSettings) workflows.DefaultAPI {
	return &workflows.DefaultAPIServiceMock{
		GetInstanceExecuteMock: utils.Ptr(func(_ workflows.ApiGetInstanceRequest) (*workflows.Instance, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.notFound {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &workflows.Instance{
				Status: settings.resourceState,
			}, nil
		}),
	}
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState workflows.InstanceStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			resourceState: workflows.INSTANCESTATUS_ACTIVE,
			wantResp:      true,
		},
		{
			desc:          "create_failed",
			resourceState: workflows.INSTANCESTATUS_FAILED,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
		},
		{
			desc:          "timeout",
			resourceState: workflows.INSTANCESTATUS_CREATING,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *workflows.Instance
				if tt.wantResp {
					wantRes = &workflows.Instance{Status: tt.resourceState}
				}

				handler := CreateInstanceWaitHandler(context.Background(), apiClient, "pid", "eu01", "iid")

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

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState workflows.InstanceStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			resourceState: workflows.INSTANCESTATUS_ACTIVE,
			wantResp:      true,
		},
		{
			desc:          "update_failed",
			resourceState: workflows.INSTANCESTATUS_FAILED,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
		},
		{
			desc:          "timeout",
			resourceState: workflows.INSTANCESTATUS_UPDATING,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *workflows.Instance
				if tt.wantResp {
					wantRes = &workflows.Instance{Status: tt.resourceState}
				}

				handler := UpdateInstanceWaitHandler(context.Background(), apiClient, "pid", "eu01", "iid")

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

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		notFound      bool
		resourceState workflows.InstanceStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:     "delete_succeeded_not_found",
			notFound: true,
		},
		{
			desc:          "delete_failed",
			resourceState: workflows.INSTANCESTATUS_FAILED,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
		},
		{
			desc:          "timeout",
			resourceState: workflows.INSTANCESTATUS_DELETING,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					notFound:      tt.notFound,
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *workflows.Instance
				if tt.wantResp {
					wantRes = &workflows.Instance{Status: tt.resourceState}
				}

				handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "pid", "eu01", "iid")

				gotRes, err := handler.WaitWithContext(context.Background())

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
		})
	}
}
