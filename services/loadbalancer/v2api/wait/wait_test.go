package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	loadbalancer "github.com/stackitcloud/stackit-sdk-go/services/loadbalancer/v2api"
)

const testRegion = "eu01"

type mockSettings struct {
	instanceName      string
	instanceStatus    *string
	instanceIsDeleted bool
	instanceGetFails  bool
}

func newAPIMock(settings *mockSettings) loadbalancer.DefaultAPI {
	return &loadbalancer.DefaultAPIServiceMock{
		GetLoadBalancerExecuteMock: utils.Ptr(func(_ loadbalancer.ApiGetLoadBalancerRequest) (*loadbalancer.LoadBalancer, error) {
			if settings.instanceGetFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.instanceIsDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &loadbalancer.LoadBalancer{
				Name:   &settings.instanceName,
				Status: settings.instanceStatus,
			}, nil
		}),
	}
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceStatus   *string
		wantErr          bool
		wantResp         bool
	}{
		{
			desc:             "create_succeeded",
			instanceGetFails: false,
			instanceStatus:   utils.Ptr(LOADBALANCERSTATUS_READY),
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "create_failed",
			instanceGetFails: false,
			instanceStatus:   utils.Ptr(LOADBALANCERSTATUS_ERROR),
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "create_failed_2",
			instanceGetFails: false,
			instanceStatus:   utils.Ptr(LOADBALANCERSTATUS_TERMINATING),
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
			instanceStatus:   utils.Ptr(LOADBALANCERSTATUS_PENDING),
			wantErr:          true,
			wantResp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceName := "foo-bar"

			apiClient := newAPIMock(&mockSettings{
				instanceName:     instanceName,
				instanceStatus:   tt.instanceStatus,
				instanceGetFails: tt.instanceGetFails,
			})

			var wantRes *loadbalancer.LoadBalancer
			if tt.wantResp {
				wantRes = &loadbalancer.LoadBalancer{
					Name:   &instanceName,
					Status: tt.instanceStatus,
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

			apiClient := newAPIMock(&mockSettings{
				instanceGetFails:  tt.instanceGetFails,
				instanceName:      instanceName,
				instanceIsDeleted: tt.instanceIsDeleted,
			})

			handler := DeleteLoadBalancerWaitHandler(context.Background(), apiClient, "", testRegion, instanceName)

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
