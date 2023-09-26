package loadbalancer

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
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

func (a *apiClientMocked) GetLoadBalancerExecute(_ context.Context, _, _ string) (*LoadBalancer, error) {
	if a.instanceGetFails {
		return nil, &GenericOpenAPIError{
			statusCode: 500,
		}
	}

	if a.instanceIsDeleted {
		return nil, &GenericOpenAPIError{
			statusCode: 404,
		}
	}

	return &LoadBalancer{
		Name:   &a.instanceName,
		Status: &a.instanceStatus,
	}, nil
}
func (a *apiClientMocked) GetStatusExecute(_ context.Context, _ string) (*StatusResponse, error) {
	if a.functionalityStatusGetFails {
		return nil, &GenericOpenAPIError{
			statusCode: 500,
		}
	}

	return &StatusResponse{
		Status: &a.functionalityStatus,
	}, nil
}

func TestHandleError(t *testing.T) {
	tests := []struct {
		desc     string
		reqErr   error
		wantRes  interface{}
		wantDone bool
		wantErr  bool
	}{
		{
			desc: "handle_oapi_error",
			reqErr: &GenericOpenAPIError{
				statusCode: 500,
			},
			wantRes:  nil,
			wantDone: false,
			wantErr:  true,
		},
		{
			desc:     "not_generic_oapi_error",
			reqErr:   fmt.Errorf("some error"),
			wantRes:  nil,
			wantDone: false,
			wantErr:  true,
		},
		{
			desc: "bad_gateway_error",
			reqErr: &GenericOpenAPIError{
				statusCode: http.StatusBadGateway,
			},
			wantRes:  nil,
			wantDone: false,
			wantErr:  false,
		},
		{
			desc: "gateway_timeout_error",
			reqErr: &GenericOpenAPIError{
				statusCode: http.StatusBadGateway,
			},
			wantRes:  nil,
			wantDone: false,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotRes, gotDone, err := handleError(tt.reqErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotRes, tt.wantRes) {
				t.Errorf("handleError() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotDone != tt.wantDone {
				t.Errorf("handleError() gotDone = %v, want %v", gotDone, tt.wantDone)
			}
		})
	}
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

			var wantRes *LoadBalancer
			if (tt.instanceStatus == InstanceStatusReady) && !tt.instanceGetFails {
				wantRes = &LoadBalancer{
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

			var wantRes *StatusResponse
			if (tt.functionalityStatus == FunctionalityStatusReady) && !tt.functionalityStatusGetFails {
				wantRes = &StatusResponse{
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
