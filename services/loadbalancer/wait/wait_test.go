package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/loadbalancer"
)

// Used for testing instance operations
type apiClientMocked struct {
	instanceName      string
	instanceStatus    loadbalancer.LoadBalancerStatus
	instanceIsDeleted bool
	instanceGetFails  bool
}

const testRegion = "eu01"

func (a *apiClientMocked) GetLoadBalancerExecute(_ context.Context, _, _, _ string) (*loadbalancer.LoadBalancer, error) {
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

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceStatus   loadbalancer.LoadBalancerStatus
		wantErr          bool
		wantResp         bool
	}{
		{
			desc:             "create_succeeded",
			instanceGetFails: false,
			instanceStatus:   loadbalancer.LOADBALANCERSTATUS_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "create_failed",
			instanceGetFails: false,
			instanceStatus:   loadbalancer.LOADBALANCERSTATUS_ERROR,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "create_failed_2",
			instanceGetFails: false,
			instanceStatus:   loadbalancer.LOADBALANCERSTATUS_TERMINATING,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "instance_get_fails",
			instanceGetFails: true,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceStatus:   loadbalancer.LOADBALANCERSTATUS_PENDING,
			wantErr:          true,
			wantResp:         false,
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
			if tt.wantResp {
				wantRes = &loadbalancer.LoadBalancer{
					Name:   &instanceName,
					Status: utils.Ptr(tt.instanceStatus),
				}
			}

			handler := CreateLoadBalancerWaitHandler(context.Background(), apiClient, "", testRegion, instanceName)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
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

			handler := DeleteLoadBalancerWaitHandler(context.Background(), apiClient, "", testRegion, instanceName)

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
