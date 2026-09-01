package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	runcommand "github.com/stackitcloud/stackit-sdk-go/services/runcommand/v2api"
)

type mockSettings struct {
	getFails      bool
	resourceState runcommand.CommandDetailsStatus
}

func newAPIMock(settings mockSettings) runcommand.DefaultAPI {
	return &runcommand.DefaultAPIServiceMock{
		GetCommandExecuteMock: utils.Ptr(func(_ runcommand.ApiGetCommandRequest) (*runcommand.CommandDetails, error) {
			if settings.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}
			return &runcommand.CommandDetails{
				Id:     utils.Ptr(int32(1)),
				Status: utils.Ptr(settings.resourceState),
			}, nil
		}),
	}
}

func TestRunCommandWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState runcommand.CommandDetailsStatus
		wantErr       bool
		wantResp      *runcommand.CommandDetails
	}{
		{
			desc:          "command completed",
			resourceState: runcommand.COMMANDDETAILSSTATUS_COMPLETED,
			wantErr:       false,
			wantResp: &runcommand.CommandDetails{
				Id:     utils.Ptr(int32(1)),
				Status: utils.Ptr(runcommand.COMMANDDETAILSSTATUS_COMPLETED),
			},
		},
		{
			desc:          "command failed returns error and details",
			resourceState: runcommand.COMMANDDETAILSSTATUS_FAILED,
			wantErr:       true,
			wantResp: &runcommand.CommandDetails{
				Id:     utils.Ptr(int32(1)),
				Status: utils.Ptr(runcommand.COMMANDDETAILSSTATUS_FAILED),
			},
		},
		{
			desc:          "unknown status returns error and details",
			resourceState: runcommand.COMMANDDETAILSSTATUS_UNKNOWN_DEFAULT_OPEN_API,
			wantErr:       true,
			wantResp: &runcommand.CommandDetails{
				Id:     utils.Ptr(int32(1)),
				Status: utils.Ptr(runcommand.COMMANDDETAILSSTATUS_UNKNOWN_DEFAULT_OPEN_API),
			},
		},
		{
			desc:     "get fails",
			getFails: true,
			wantErr:  true,
			wantResp: nil,
		},
		{
			desc:          "timeout while running",
			resourceState: runcommand.COMMANDDETAILSSTATUS_RUNNING,
			wantErr:       true,
			wantResp:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				handler := RunCommandWaitHandler(context.Background(), apiClient, "pid", "sid", "eu01", "1")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, tt.wantResp) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, tt.wantResp)
				}
			})
		})
	}
}
