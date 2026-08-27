package wait

import (
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	automation "github.com/stackitcloud/stackit-sdk-go/services/automation/v1betaapi"
)

const testRegion = "eu01"

var (
	testProject     = uuid.NewString()
	testDate        = time.Now()
	testAutomation  = uuid.NewString()
	testExecutionId = uuid.NewString()
)

type executionResponse struct {
	execution *automation.VolumeExecutionResponse
	err       error
}

type mockSettings struct {
	idxExecutionResponse int
	executionResponses   []executionResponse
}

func newAPIMock(settings *mockSettings) automation.DefaultAPI {
	return &automation.DefaultAPIServiceMock{
		GetVolumeExecutionExecuteMock: utils.Ptr(func(_ automation.ApiGetVolumeExecutionRequest) (*automation.VolumeExecutionResponse, error) {
			resp := settings.executionResponses[settings.idxExecutionResponse]
			settings.idxExecutionResponse++
			settings.idxExecutionResponse %= len(settings.executionResponses)
			return resp.execution, resp.err
		}),
	}
}

func fixtureExecution(status automation.VolumeExecutionResponseStatus) *automation.VolumeExecutionResponse {
	return &automation.VolumeExecutionResponse{
		CreateTime: testDate,
		Id:         testExecutionId,
		Status:     status,
	}
}

func TestCreateVolumeExecutionWaitHandler(t *testing.T) {
	tests := []struct {
		name      string
		responses []executionResponse
		want      *automation.VolumeExecutionResponse
		wantErr   bool
	}{
		{
			name: "execution completed immediately",
			responses: []executionResponse{
				{fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_COMPLETED), nil},
			},
			want:    fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_COMPLETED),
			wantErr: false,
		},
		{
			name: "execution completed delayed",
			responses: []executionResponse{
				{fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_PENDING), nil},
				{fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_RUNNING), nil},
				{fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_COMPLETED), nil},
			},
			want:    fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_COMPLETED),
			wantErr: false,
		},
		{
			name: "execution failed",
			responses: []executionResponse{
				{fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_RUNNING), nil},
				{fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_FAILED), nil},
			},
			want:    fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_FAILED),
			wantErr: true,
		},
		{
			name: "execution terminated",
			responses: []executionResponse{
				{fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_TERMINATED), nil},
			},
			want:    fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_TERMINATED),
			wantErr: true,
		},
		{
			name: "timeout",
			responses: []executionResponse{
				{fixtureExecution(automation.VOLUMEEXECUTIONRESPONSESTATUS_RUNNING), nil},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "get execution fails",
			responses: []executionResponse{
				{nil, &oapierror.GenericOpenAPIError{StatusCode: http.StatusInternalServerError}},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				ctx := t.Context()
				client := newAPIMock(&mockSettings{
					executionResponses: tt.responses,
				})

				handler := CreateVolumeExecutionWaitHandler(ctx, client, testProject, testRegion, testAutomation, testExecutionId)
				got, err := handler.WaitWithContext(ctx)
				if (err != nil) != tt.wantErr {
					t.Fatalf("unexpected error response. want %v but got %v ", tt.wantErr, err)
				}

				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("differing execution %s", diff)
				}
			})
		})
	}
}
