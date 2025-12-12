package wait

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/logs"
)

type apiClientMock struct {
	getFails        bool
	returnInstance  bool
	statusCode      int
	getLogsResponse *logs.LogsInstance
}

func (a *apiClientMock) GetLogsInstanceExecute(_ context.Context, _, _, _ string) (*logs.LogsInstance, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: a.statusCode,
		}
	}
	if !a.returnInstance {
		return nil, nil
	}
	return a.getLogsResponse, nil
}

var projectId = uuid.NewString()
var instanceId = uuid.NewString()
const region = "eu01"

func TestCreateLogsInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		description     string
		getFails        bool
		wantErr         bool
		wantResp        bool
		returnInstance  bool
		getLogsResponse *logs.LogsInstance
	}{
		{
			description:    "create succeeded",
			getFails:       false,
			wantErr:        false,
			wantResp:       true,
			returnInstance: true,
			getLogsResponse: &logs.LogsInstance{
				Id:     utils.Ptr(INSTANCE_ID),
				Status: utils.Ptr(logs.LOGSINSTANCESTATUS_ACTIVE),
			},
		},
		{
			description:    "create failed with error",
			getFails:       true,
			wantErr:        true,
			wantResp:       false,
			returnInstance: true,
			getLogsResponse: &logs.LogsInstance{
				Id:     utils.Ptr(INSTANCE_ID),
				Status: utils.Ptr(logs.LOGSINSTANCESTATUS_ACTIVE),
			},
		},
		{
			description:    "create without id",
			getFails:       false,
			wantErr:        true,
			wantResp:       false,
			returnInstance: true,
			getLogsResponse: &logs.LogsInstance{
				Status: utils.Ptr(logs.LOGSINSTANCESTATUS_ACTIVE),
			},
		},
		{
			description:    "create without status",
			getFails:       false,
			wantErr:        true,
			wantResp:       false,
			returnInstance: true,
			getLogsResponse: &logs.LogsInstance{
				Id: utils.Ptr(INSTANCE_ID),
			},
		},
		{
			description:    "instance deleting",
			getFails:       false,
			wantErr:        true,
			wantResp:       false,
			returnInstance: true,
			getLogsResponse: &logs.LogsInstance{
				Id:     utils.Ptr(INSTANCE_ID),
				Status: utils.Ptr(logs.LOGSINSTANCESTATUS_DELETING),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			client := &apiClientMock{
				getFails:        tt.getFails,
				getLogsResponse: tt.getLogsResponse,
				returnInstance:  tt.returnInstance,
			}
			var instanceWanted *logs.LogsInstance
			if tt.wantResp {
				instanceWanted = tt.getLogsResponse
			}

			handler := CreateLogsInstanceWaitHandler(context.Background(), client, PROJECT_ID, REGION, INSTANCE_ID)

			response, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(response, instanceWanted) {
				t.Fatalf("handler gotRes = %v, want %v", response, instanceWanted)
			}
		})
	}
}

func TestDeleteLogsInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		description string
		getFails    bool
		wantErr     bool
		statusCode  int
	}{
		{
			description: "delete succeeded",
			getFails:    true,
			statusCode:  http.StatusNotFound,
		},
		{
			description: "delete failed with error",
			getFails:    true,
			wantErr:     true,
			statusCode:  http.StatusInternalServerError,
		},
		{
			description: "delete still in progress",
			getFails:    false,
			wantErr:     true,
			statusCode:  http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			client := &apiClientMock{
				getFails:        tt.getFails,
				returnInstance:  false,
				statusCode:      tt.statusCode,
				getLogsResponse: nil,
			}
			handler := DeleteLogsInstanceWaitHandler(context.Background(), client, PROJECT_ID, REGION, INSTANCE_ID)
			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())
			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
