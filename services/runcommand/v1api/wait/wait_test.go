package wait

import (
	"context"
	"sync/atomic"
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

var testPayload = *runcommand.NewCreateCommandPayload("RunShellScript")

func TestRunCommandWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState runcommand.CommandDetailsStatus
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "command completed",
			getFails:      false,
			resourceState: runcommand.COMMANDDETAILSSTATUS_COMPLETED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "command failed",
			getFails:      false,
			resourceState: runcommand.COMMANDDETAILSSTATUS_FAILED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get fails",
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

func TestAgentReadyWaitHandler(t *testing.T) {
	tests := []struct {
		desc     string
		createFn func(runcommand.ApiCreateCommandRequest) (*runcommand.NewCommandResponse, error)
		wantErr  bool
		wantResp *runcommand.NewCommandResponse
	}{
		{
			desc: "agent immediately ready",
			createFn: func(_ runcommand.ApiCreateCommandRequest) (*runcommand.NewCommandResponse, error) {
				return &runcommand.NewCommandResponse{Id: utils.Ptr(int32(42))}, nil
			},
			wantErr:  false,
			wantResp: &runcommand.NewCommandResponse{Id: utils.Ptr(int32(42))},
		},
		{
			desc: "agent not ready then ready",
			// atomic counter ensures the closure is safe when called from the handler goroutine
			createFn: func() func(runcommand.ApiCreateCommandRequest) (*runcommand.NewCommandResponse, error) {
				var calls atomic.Int32
				return func(_ runcommand.ApiCreateCommandRequest) (*runcommand.NewCommandResponse, error) {
					if calls.Add(1) == 1 {
						return nil, &oapierror.GenericOpenAPIError{StatusCode: 404}
					}
					return &runcommand.NewCommandResponse{Id: utils.Ptr(int32(7))}, nil
				}
			}(),
			wantErr:  false,
			wantResp: &runcommand.NewCommandResponse{Id: utils.Ptr(int32(7))},
		},
		{
			desc: "terminal error non 404",
			createFn: func(_ runcommand.ApiCreateCommandRequest) (*runcommand.NewCommandResponse, error) {
				return nil, &oapierror.GenericOpenAPIError{StatusCode: 500}
			},
			wantErr:  true,
			wantResp: nil,
		},
		{
			desc: "timeout agent never ready",
			createFn: func(_ runcommand.ApiCreateCommandRequest) (*runcommand.NewCommandResponse, error) {
				return nil, &oapierror.GenericOpenAPIError{StatusCode: 404}
			},
			wantErr:  true,
			wantResp: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := &runcommand.DefaultAPIServiceMock{
					CreateCommandExecuteMock: utils.Ptr(tt.createFn),
				}

				handler := AgentReadyWaitHandler(context.Background(), apiClient, "pid", "sid", testPayload)

				// 1 ms throttle keeps the retry case within the 10 ms fake timeout
				gotRes, err := handler.
					SetThrottle(time.Millisecond).
					SetTimeout(10 * time.Millisecond).
					WaitWithContext(context.Background())

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
