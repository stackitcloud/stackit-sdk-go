package wait

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/services/cdn"
)

var _ APIClientInterface = (*apiClientMock)(nil)

type response struct {
	state  string
	status int
}

type apiClientMock struct {
	responseIndex int
	responses     []response
}

// GetDistributionExecute implements APIClientInterface.
func (mock *apiClientMock) GetDistributionExecute(ctx context.Context, projectId string, distributionId string) (*cdn.GetDistributionResponse, error) {
	response := mock.responses[mock.responseIndex]
	mock.responseIndex++
	mock.responseIndex %= len(mock.responses)
	if response.status >= http.StatusBadRequest {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode:   response.status,
			ErrorMessage: "simulated error",
		}
	}
	return &cdn.GetDistributionResponse{
		Distribution: &cdn.Distribution{
			Id:     &testDistributionId,
			Status: &response.state,
		},
	}, nil
}

var (
	testProjectId      = "testProjectId"
	testDistributionId = "testDistributionId"
)

func TestCreateDistributionPoolWaitHandler(t *testing.T) {
	type args struct {
		projectId      string
		distributionId string
	}
	type fields struct {
		responses []response
	}
	tests := []struct {
		name      string
		args      args
		fields    fields
		wantState string
		wantErr   bool
	}{
		{
			"success immediately",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateActive, 0},
				},
			},
			DistributionStateActive,
			false,
		},
		{
			"success delayed",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateUpdating, 0},
					{DistributionStateUpdating, 0},
					{DistributionStateActive, http.StatusOK},
				},
			},
			DistributionStateActive,
			false,
		},
		{
			"request fails",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateUpdating, 0},
					{DistributionStateUpdating, 0},
					{DistributionStateUpdating, http.StatusInternalServerError},
				},
			},
			"",
			true,
		},
		{
			"timeout",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateUpdating, 0},
				},
			},
			DistributionStateUpdating,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClientMock := apiClientMock{
				responses: tt.fields.responses,
			}
			ctx := context.Background()
			handler := CreateDistributionPoolWaitHandler(ctx, &apiClientMock, tt.args.projectId, tt.args.distributionId)
			handler.SetTimeout(2 * time.Second)
			handler.SetSleepBeforeWait(100 * time.Millisecond)
			handler.SetThrottle(50 * time.Millisecond)
			response, err := handler.WaitWithContext(ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error %s, wantErr %v", err, tt.wantErr)
			}
			var actualState string
			if response != nil {
				actualState = *response.Distribution.Status
			}
			if tt.wantState != actualState {
				t.Errorf("wrong state want %q but got %q", tt.wantState, actualState)
			}
		})
	}
}

func TestUpdateDistributionWaitHandler(t *testing.T) {
	type args struct {
		projectId      string
		distributionId string
	}
	type fields struct {
		responses []response
	}
	tests := []struct {
		name      string
		args      args
		fields    fields
		wantState string
		wantErr   bool
	}{
		{
			"success immediately",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateActive, 0},
				},
			},
			DistributionStateActive,
			false,
		},
		{
			"success delayed",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateUpdating, 0},
					{DistributionStateUpdating, 0},
					{DistributionStateActive, http.StatusOK},
				},
			},
			DistributionStateActive,
			false,
		},
		{
			"request fails",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateUpdating, 0},
					{DistributionStateUpdating, 0},
					{DistributionStateUpdating, http.StatusInternalServerError},
				},
			},
			"",
			true,
		},
		{
			"timeout",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateUpdating, 0},
				},
			},
			DistributionStateUpdating,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClientMock := apiClientMock{
				responses: tt.fields.responses,
			}
			ctx := context.Background()
			handler := UpdateDistributionWaitHandler(ctx, &apiClientMock, tt.args.projectId, tt.args.distributionId)
			handler.SetTimeout(2 * time.Second)
			handler.SetSleepBeforeWait(100 * time.Millisecond)
			handler.SetThrottle(50 * time.Millisecond)
			response, err := handler.WaitWithContext(ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error %s, wantErr %v", err, tt.wantErr)
			}
			var actualState string
			if response != nil {
				actualState = *response.Distribution.Status
			}
			if tt.wantState != actualState {
				t.Errorf("wrong state want %q but got %q", tt.wantState, actualState)
			}
		})
	}
}

func TestDeleteDistributionWaitHandler(t *testing.T) {
	type args struct {
		projectId      string
		distributionId string
	}
	type fields struct {
		responses []response
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		wantErr bool
	}{
		{
			"success immediately",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateActive, http.StatusNotFound},
				},
			},
			false,
		},
		{
			"success delayed",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateUpdating, http.StatusOK},
					{DistributionStateUpdating, http.StatusOK},
					{DistributionStateActive, http.StatusNotFound},
				},
			},
			false,
		},
		{
			"request fails",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateUpdating, http.StatusInternalServerError},
				},
			},
			true,
		},
		{
			"timeout",
			args{
				projectId:      testProjectId,
				distributionId: testDistributionId,
			},
			fields{
				responses: []response{
					{DistributionStateUpdating, 0},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClientMock := apiClientMock{
				responses: tt.fields.responses,
			}
			ctx := context.Background()
			handler := DeleteDistributionWaitHandler(ctx, &apiClientMock, tt.args.projectId, tt.args.distributionId)
			handler.SetTimeout(2 * time.Second)
			handler.SetSleepBeforeWait(100 * time.Millisecond)
			handler.SetThrottle(50 * time.Millisecond)
			_, err := handler.WaitWithContext(ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error %s, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
