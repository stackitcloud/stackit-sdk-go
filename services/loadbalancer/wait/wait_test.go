package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/services/loadbalancer"
)

// Used for testing instance operations
type apiClientMocked struct {
	instanceName                string
	instanceStatus              string
	instanceIsDeleted           bool
	instanceGetFails            bool
	functionalityStatus         string
	functionalityStatusGetFails bool
}

func (a *apiClientMocked) GetLoadBalancerExecute(_ context.Context, _, _ string) (*loadbalancer.LoadBalancer, error) {
	if a.instanceGetFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	if a.instanceIsDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	return &loadbalancer.LoadBalancer{
		Name:   &a.instanceName,
		Status: &a.instanceStatus,
	}, nil
}
func (a *apiClientMocked) GetStatusExecute(_ context.Context, _ string) (*loadbalancer.StatusResponse, error) {
	if a.functionalityStatusGetFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &loadbalancer.StatusResponse{
		Status: &a.functionalityStatus,
	}, nil
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceStatus   string
		wantErr          bool
	}{
		{
			desc:             "create_succeeded",
			instanceGetFails: false,
			instanceStatus:   InstanceStatusReady,
			wantErr:          false,
		},
		{
			desc:             "create_failed",
			instanceGetFails: false,
			instanceStatus:   InstanceStatusError,
			wantErr:          true,
		},
		{
			desc:             "create_failed_2",
			instanceGetFails: false,
			instanceStatus:   InstanceStatusTerminating,
			wantErr:          true,
		},
		{
			desc:             "instance_get_fails",
			instanceGetFails: true,
			wantErr:          true,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceStatus:   "ANOTHER STATE",
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceName := "foo-bar"

			apiClient := &apiClientMocked{
				instanceName:     instanceName,
				instanceStatus:   tt.instanceStatus,
				instanceGetFails: tt.instanceGetFails,
			}

			var wantRes *loadbalancer.LoadBalancer
			if (tt.instanceStatus == InstanceStatusReady) && !tt.instanceGetFails {
				wantRes = &loadbalancer.LoadBalancer{
					Name:   &instanceName,
					Status: &tt.instanceStatus,
				}
			}

			handler := CreateInstanceWaitHandler(context.Background(), apiClient, "", instanceName)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

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
	}
}

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc              string
		instanceGetFails  bool
		instanceIsDeleted bool
		wantErr           bool
	}{
		{
			desc:              "delete_succeeded",
			instanceGetFails:  false,
			instanceIsDeleted: true,
			wantErr:           false,
		},
		{
			desc:              "delete_failed",
			instanceGetFails:  false,
			instanceIsDeleted: false,
			wantErr:           true,
		},
		{
			desc:             "get_fails",
			instanceGetFails: true,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceName := "foo-bar"

			apiClient := &apiClientMocked{
				instanceGetFails:  tt.instanceGetFails,
				instanceName:      instanceName,
				instanceIsDeleted: tt.instanceIsDeleted,
			}

			handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "", instanceName)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, nil)
			}
		})
	}
}

func TestEnableLoadBalancingWaitHandler(t *testing.T) {
	tests := []struct {
		desc                        string
		functionalityStatus         string
		functionalityStatusGetFails bool
		wantErr                     bool
	}{
		{
			desc:                        "enable_succeeded",
			functionalityStatus:         FunctionalityStatusReady,
			functionalityStatusGetFails: false,
			wantErr:                     false,
		},
		{
			desc:                        "enable_updating",
			functionalityStatus:         FunctionalityStatusUpdating,
			functionalityStatusGetFails: false,
			wantErr:                     true,
		},
		{
			desc:                        "enable_failed",
			functionalityStatus:         FunctionalityStatusFailed,
			functionalityStatusGetFails: false,
			wantErr:                     true,
		},
		{
			desc:                        "enable_failed_2",
			functionalityStatus:         FunctionalityStatusUnspecified,
			functionalityStatusGetFails: true,
			wantErr:                     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				functionalityStatus:         tt.functionalityStatus,
				functionalityStatusGetFails: tt.functionalityStatusGetFails,
			}

			var wantRes *loadbalancer.StatusResponse
			if (tt.functionalityStatus == FunctionalityStatusReady) && !tt.functionalityStatusGetFails {
				wantRes = &loadbalancer.StatusResponse{
					Status: &tt.functionalityStatus,
				}
			}

			handler := EnableLoadBalancingWaitHandler(context.Background(), apiClient, "")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && wantRes != nil && !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}
