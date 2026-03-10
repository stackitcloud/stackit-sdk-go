package wait

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	sfs "github.com/stackitcloud/stackit-sdk-go/services/sfs/v1api"
)

const (
	testResourcePoolId = "testResourcePoolId"
	testShareId        = "testShareId"
)

type response struct {
	state  string
	status int
}

type mockSettings struct {
	responseIndex int
	responses     []response
}

func newAPIMock(settings *mockSettings) sfs.DefaultAPI {
	return &sfs.DefaultAPIServiceMock{
		GetShareExecuteMock: utils.Ptr(func(_ sfs.ApiGetShareRequest) (*sfs.GetShareResponse, error) {
			response := settings.responses[settings.responseIndex]
			settings.responseIndex++
			settings.responseIndex %= len(settings.responses)
			if response.status >= http.StatusBadRequest {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode:   response.status,
					ErrorMessage: "simulated error",
				}
			}
			return &sfs.GetShareResponse{
				Share: &sfs.Share{
					Id:    utils.Ptr(testShareId),
					State: &response.state,
				},
			}, nil
		}),
		GetResourcePoolExecuteMock: utils.Ptr(func(_ sfs.ApiGetResourcePoolRequest) (*sfs.GetResourcePoolResponse, error) {
			response := settings.responses[settings.responseIndex]
			settings.responseIndex++
			settings.responseIndex %= len(settings.responses)
			if response.status >= http.StatusBadRequest {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode:   response.status,
					ErrorMessage: "simulated error",
				}
			}
			return &sfs.GetResourcePoolResponse{
				ResourcePool: &sfs.ResourcePool{
					Id:    utils.Ptr(testResourcePoolId),
					State: &response.state,
				},
			}, nil
		}),
	}
}

func TestCreateResourcePoolWaitHandler(t *testing.T) {
	type args struct {
		projectId string
		region    string
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
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateCreated, 0},
				},
			},
			ResourcePoolStateCreated,
			false,
		},
		{
			"success delayed",
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateCreating, 0},
					{ResourcePoolStateCreating, 0},
					{ResourcePoolStateCreated, http.StatusOK},
				},
			},
			ResourcePoolStateCreated,
			false,
		},
		{
			"request fails",
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateCreating, 0},
					{ResourcePoolStateCreating, 0},
					{ResourcePoolStateCreating, http.StatusInternalServerError},
				},
			},
			"",
			true,
		},
		{
			"timeout",
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateCreating, 0},
				},
			},
			ResourcePoolStateCreating,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClientMock := newAPIMock(&mockSettings{
				responses: tt.fields.responses,
			})

			ctx := context.Background()
			handler := CreateResourcePoolWaitHandler(ctx, apiClientMock, tt.args.projectId, tt.args.region, testResourcePoolId)
			handler.SetTimeout(2 * time.Second)
			handler.SetSleepBeforeWait(100 * time.Millisecond)
			handler.SetThrottle(50 * time.Millisecond)
			response, err := handler.WaitWithContext(ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error %s, wantErr %v", err, tt.wantErr)
			}
			var actualState string
			if response != nil {
				actualState = *response.ResourcePool.State
			}
			if tt.wantState != actualState {
				t.Errorf("wrong state want %q but got %q", tt.wantState, actualState)
			}
		})
	}
}

func TestUpdateResourcePoolWaitHandler(t *testing.T) {
	type args struct {
		projectId string
		region    string
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
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateCreated, 0},
				},
			},
			ResourcePoolStateCreated,
			false,
		},
		{
			"success delayed",
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateUpdating, 0},
					{ResourcePoolStateUpdating, 0},
					{ResourcePoolStateCreated, http.StatusOK},
				},
			},
			ResourcePoolStateCreated,
			false,
		},
		{
			"request fails",
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateUpdating, 0},
					{ResourcePoolStateUpdating, 0},
					{ResourcePoolStateUpdating, http.StatusInternalServerError},
				},
			},
			"",
			true,
		},
		{
			"timeout",
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateUpdating, 0},
				},
			},
			ResourcePoolStateUpdating,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClientMock := newAPIMock(&mockSettings{
				responses: tt.fields.responses,
			})

			ctx := context.Background()
			handler := UpdateResourcePoolWaitHandler(ctx, apiClientMock, tt.args.projectId, tt.args.region, testResourcePoolId)
			handler.SetTimeout(2 * time.Second)
			handler.SetSleepBeforeWait(100 * time.Millisecond)
			handler.SetThrottle(50 * time.Millisecond)
			response, err := handler.WaitWithContext(ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error %s, wantErr %v", err, tt.wantErr)
			}
			var actualState string
			if response != nil {
				actualState = *response.ResourcePool.State
			}
			if tt.wantState != actualState {
				t.Errorf("wrong state want %q but got %q", tt.wantState, actualState)
			}
		})
	}
}

func TestDeleteResourcePoolWaitHandler(t *testing.T) {
	type args struct {
		projectId string
		region    string
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
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateCreated, http.StatusNotFound},
				},
			},
			false,
		},
		{
			"success delayed",
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateCreating, http.StatusOK},
					{ResourcePoolStateCreating, http.StatusOK},
					{ResourcePoolStateCreated, http.StatusNotFound},
				},
			},
			false,
		},
		{
			"request fails",
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateCreating, http.StatusInternalServerError},
				},
			},
			true,
		},
		{
			"timeout",
			args{},
			fields{
				responses: []response{
					{ResourcePoolStateCreating, 0},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClientMock := newAPIMock(&mockSettings{
				responses: tt.fields.responses,
			})

			ctx := context.Background()
			handler := DeleteResourcePoolWaitHandler(ctx, apiClientMock, tt.args.projectId, tt.args.region, testResourcePoolId)
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

func TestCreateShareWaitHandler(t *testing.T) {
	type args struct {
		projectId string
		region    string
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
			args{},
			fields{
				responses: []response{
					{ShareStateCreated, 0},
				},
			},
			ShareStateCreated,
			false,
		},
		{
			"success delayed",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, 0},
					{ShareStateCreating, 0},
					{ShareStateCreated, http.StatusOK},
				},
			},
			ShareStateCreated,
			false,
		},
		{
			"request fails",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, 0},
					{ShareStateCreating, 0},
					{ShareStateCreating, http.StatusInternalServerError},
				},
			},
			"",
			true,
		},
		{
			"timeout",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, 0},
				},
			},
			ShareStateCreating,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClientMock := newAPIMock(&mockSettings{
				responses: tt.fields.responses,
			})

			ctx := context.Background()
			handler := CreateShareWaitHandler(ctx, apiClientMock, tt.args.projectId, tt.args.region, testResourcePoolId, testShareId)
			handler.SetTimeout(2 * time.Second)
			handler.SetSleepBeforeWait(100 * time.Millisecond)
			handler.SetThrottle(50 * time.Millisecond)
			response, err := handler.WaitWithContext(ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error %s, wantErr %v", err, tt.wantErr)
			}
			var actualState string
			if response != nil {
				actualState = *response.Share.State
			}
			if tt.wantState != actualState {
				t.Errorf("wrong state want %q but got %q", tt.wantState, actualState)
			}
		})
	}
}

func TestUpdateShareWaitHandler(t *testing.T) {
	type args struct {
		projectId string
		region    string
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
			args{},
			fields{
				responses: []response{
					{ShareStateCreated, 0},
				},
			},
			ShareStateCreated,
			false,
		},
		{
			"success delayed",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, 0},
					{ShareStateCreating, 0},
					{ShareStateCreated, http.StatusOK},
				},
			},
			ShareStateCreated,
			false,
		},
		{
			"request fails",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, 0},
					{ShareStateCreating, 0},
					{ShareStateCreating, http.StatusInternalServerError},
				},
			},
			"",
			true,
		},
		{
			"timeout",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, 0},
				},
			},
			ShareStateCreating,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClientMock := newAPIMock(&mockSettings{
				responses: tt.fields.responses,
			})

			ctx := context.Background()
			handler := UpdateShareWaitHandler(ctx, apiClientMock, tt.args.projectId, tt.args.region, testResourcePoolId, testShareId)
			handler.SetTimeout(2 * time.Second)
			handler.SetSleepBeforeWait(100 * time.Millisecond)
			handler.SetThrottle(50 * time.Millisecond)
			response, err := handler.WaitWithContext(ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error %s, wantErr %v", err, tt.wantErr)
			}
			var actualState string
			if response != nil {
				actualState = *response.Share.State
			}
			if tt.wantState != actualState {
				t.Errorf("wrong state want %q but got %q", tt.wantState, actualState)
			}
		})
	}
}

func TestDeleteShareWaitHandler(t *testing.T) {
	type args struct {
		projectId string
		region    string
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
			args{},
			fields{
				responses: []response{
					{ShareStateCreated, http.StatusNotFound},
				},
			},
			false,
		},
		{
			"success immediately with gone status",
			args{},
			fields{
				responses: []response{
					{ShareStateCreated, http.StatusGone},
				},
			},
			false,
		},
		{
			"success delayed",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, http.StatusOK},
					{ShareStateCreating, http.StatusOK},
					{ShareStateCreated, http.StatusNotFound},
				},
			},
			false,
		},
		{
			"success delayed with gone state",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, http.StatusOK},
					{ShareStateCreating, http.StatusOK},
					{ShareStateCreated, http.StatusGone},
				},
			},
			false,
		},
		{
			"request fails",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, http.StatusInternalServerError},
				},
			},
			true,
		},
		{
			"timeout",
			args{},
			fields{
				responses: []response{
					{ShareStateCreating, 0},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiClientMock := newAPIMock(&mockSettings{
				responses: tt.fields.responses,
			})

			ctx := context.Background()
			handler := DeleteShareWaitHandler(ctx, apiClientMock, tt.args.projectId, tt.args.region, testResourcePoolId, testShareId)
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
