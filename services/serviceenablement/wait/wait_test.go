package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/serviceenablement"
)

// Used for testing service operations
type apiClientServiceMocked struct {
	serviceId       string
	serviceState    string
	getServiceFails bool
}

func (a *apiClientServiceMocked) GetServiceStatusRegionalExecute(_ context.Context, _, _, _ string) (*serviceenablement.ServiceStatus, error) {
	if a.getServiceFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &serviceenablement.ServiceStatus{
		ServiceId: &a.serviceId,
		State:     &a.serviceState,
	}, nil
}

func TestEnableServiceWaitHandler(t *testing.T) {
	tests := []struct {
		desc            string
		getServiceFails bool
		serviceState    string
		wantErr         bool
		wantResp        bool
	}{
		{
			desc:            "enable_succeeded",
			getServiceFails: false,
			serviceState:    ServiceStateEnabled,
			wantErr:         false,
			wantResp:        true,
		},
		{
			desc:            "enable_failed",
			getServiceFails: false,
			serviceState:    ServiceStateDisabled,
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "enable_failed_2",
			getServiceFails: false,
			serviceState:    ServiceStateDisabling,
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "timeout",
			getServiceFails: false,
			serviceState:    ServiceStateEnabling,
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "get_service_fails",
			getServiceFails: true,
			serviceState:    ServiceStateEnabling,
			wantErr:         true,
			wantResp:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			serviceId := "serviceId"

			apiClient := &apiClientServiceMocked{
				serviceId:       serviceId,
				serviceState:    tt.serviceState,
				getServiceFails: tt.getServiceFails,
			}

			var wantRes *serviceenablement.ServiceStatus
			if tt.wantResp {
				wantRes = &serviceenablement.ServiceStatus{
					ServiceId: &serviceId,
					State:     utils.Ptr(tt.serviceState),
				}
			}

			ctx := context.Background()

			handler := EnableServiceWaitHandler(ctx, apiClient, "eu01", "projectId", serviceId)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(10 * time.Millisecond).WaitWithContext(ctx)

			if err != nil {
				if !tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		},
		)
	}
}

func TestDisableServiceWaitHandler(t *testing.T) {
	tests := []struct {
		desc            string
		getServiceFails bool
		serviceState    string
		wantErr         bool
		wantResp        bool
	}{
		{
			desc:            "disable_succeeded",
			getServiceFails: false,
			serviceState:    ServiceStateDisabled,
			wantErr:         false,
			wantResp:        true,
		},
		{
			desc:            "disable_failed",
			getServiceFails: false,
			serviceState:    ServiceStateEnabled,
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "disable_failed_2",
			getServiceFails: false,
			serviceState:    ServiceStateEnabling,
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "timeout",
			getServiceFails: false,
			serviceState:    ServiceStateDisabling,
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:            "get_service_fails",
			getServiceFails: true,
			serviceState:    ServiceStateDisabling,
			wantErr:         true,
			wantResp:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			serviceId := "serviceId"

			apiClient := &apiClientServiceMocked{
				serviceId:       serviceId,
				serviceState:    tt.serviceState,
				getServiceFails: tt.getServiceFails,
			}

			var wantRes *serviceenablement.ServiceStatus
			if tt.wantResp {
				wantRes = &serviceenablement.ServiceStatus{
					ServiceId: &serviceId,
					State:     utils.Ptr(tt.serviceState),
				}
			}

			ctx := context.Background()

			handler := DisableServiceWaitHandler(ctx, apiClient, "eu01", "projectId", serviceId)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(10 * time.Millisecond).WaitWithContext(ctx)

			if err != nil {
				if !tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		},
		)
	}
}
