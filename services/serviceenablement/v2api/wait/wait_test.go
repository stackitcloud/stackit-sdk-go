package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	serviceenablement "github.com/stackitcloud/stackit-sdk-go/services/serviceenablement/v2api"
)

type mockSettings struct {
	serviceId       string
	serviceState    *string
	getServiceFails bool
}

func newAPIMock(settings *mockSettings) serviceenablement.DefaultAPI {
	return &serviceenablement.DefaultAPIServiceMock{
		GetServiceStatusRegionalExecuteMock: utils.Ptr(func(_ serviceenablement.ApiGetServiceStatusRegionalRequest) (*serviceenablement.ServiceStatus, error) {
			if settings.getServiceFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &serviceenablement.ServiceStatus{
				ServiceId: &settings.serviceId,
				State:     settings.serviceState,
			}, nil
		}),
	}
}

func TestEnableServiceWaitHandler(t *testing.T) {
	tests := []struct {
		desc            string
		getServiceFails bool
		serviceState    *string
		wantErr         bool
		wantResp        bool
	}{
		{
			desc:            "enable_succeeded",
			getServiceFails: false,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_ENABLED),
			wantErr:         false,
			wantResp:        true,
		},
		{
			desc:            "enable_failed",
			getServiceFails: false,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_DISABLED),
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "enable_failed_2",
			getServiceFails: false,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_DISABLING),
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "timeout",
			getServiceFails: false,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_ENABLING),
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "get_service_fails",
			getServiceFails: true,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_ENABLING),
			wantErr:         true,
			wantResp:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				serviceId := "serviceId"

				apiClient := newAPIMock(&mockSettings{
					serviceId:       serviceId,
					serviceState:    tt.serviceState,
					getServiceFails: tt.getServiceFails,
				})

				var wantRes *serviceenablement.ServiceStatus
				if tt.wantResp {
					wantRes = &serviceenablement.ServiceStatus{
						ServiceId: &serviceId,
						State:     tt.serviceState,
					}
				}

				ctx := context.Background()

				handler := EnableServiceWaitHandler(ctx, apiClient, "eu01", "projectId", serviceId)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(0).WaitWithContext(ctx)

				if err != nil {
					if !tt.wantErr {
						t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
					}
					return
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		},
		)
	}
}

func TestDisableServiceWaitHandler(t *testing.T) {
	tests := []struct {
		desc            string
		getServiceFails bool
		serviceState    *string
		wantErr         bool
		wantResp        bool
	}{
		{
			desc:            "disable_succeeded",
			getServiceFails: false,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_DISABLED),
			wantErr:         false,
			wantResp:        true,
		},
		{
			desc:            "disable_failed",
			getServiceFails: false,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_ENABLED),
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "disable_failed_2",
			getServiceFails: false,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_ENABLING),
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "timeout",
			getServiceFails: false,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_DISABLING),
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "get_service_fails",
			getServiceFails: true,
			serviceState:    utils.Ptr(SERVICESTATUSSTATE_DISABLING),
			wantErr:         true,
			wantResp:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				serviceId := "serviceId"

				apiClient := newAPIMock(&mockSettings{
					serviceId:       serviceId,
					serviceState:    tt.serviceState,
					getServiceFails: tt.getServiceFails,
				})

				var wantRes *serviceenablement.ServiceStatus
				if tt.wantResp {
					wantRes = &serviceenablement.ServiceStatus{
						ServiceId: &serviceId,
						State:     tt.serviceState,
					}
				}

				ctx := context.Background()

				handler := DisableServiceWaitHandler(ctx, apiClient, "eu01", "projectId", serviceId)

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(0).WaitWithContext(ctx)

				if err != nil {
					if !tt.wantErr {
						t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
					}
					return
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		},
		)
	}
}
