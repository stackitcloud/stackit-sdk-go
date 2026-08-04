package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	runcommand "github.com/stackitcloud/stackit-sdk-go/services/runcommand/v1api"
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
		wantResp      bool
	}{
		{
			desc:          "command_completed",
			getFails:      false,
			resourceState: runcommand.COMMANDDETAILSSTATUS_COMPLETED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "command_failed",
			getFails:      false,
			resourceState: runcommand.COMMANDDETAILSSTATUS_FAILED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: runcommand.COMMANDDETAILSSTATUS_UNKNOWN_DEFAULT_OPEN_API,
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: runcommand.COMMANDDETAILSSTATUS_RUNNING,
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(mockSettings{
					getFails:      tt.getFails,
					resourceState: tt.resourceState,
				})

				var wantRes *runcommand.CommandDetails
				if tt.wantResp {
					wantRes = &runcommand.CommandDetails{
						Id:     utils.Ptr(int32(1)),
						Status: utils.Ptr(tt.resourceState),
					}
				}

				handler := RunCommandWaitHandler(context.Background(), apiClient, "pid", "sid", "1")

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
