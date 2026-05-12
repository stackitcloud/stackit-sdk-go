package wait

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	vpn "github.com/stackitcloud/stackit-sdk-go/services/vpn/v1api"
)

type mockSettings struct {
	getFails     bool
	getNotFound  bool
	getForbidden bool
	getGone      bool
	gatewayState vpn.GatewayStatus
	gatewayId    string
}

func newAPIMock(settings []mockSettings) vpn.DefaultAPI {
	count := 0
	return &vpn.DefaultAPIServiceMock{
		GetGatewayExecuteMock: utils.Ptr(func(_ vpn.ApiGetGatewayRequest) (*vpn.GatewayResponse, error) {
			setting := settings[count%len(settings)]
			count++

			if setting.getFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusInternalServerError,
				}
			}

			if setting.getNotFound {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusNotFound,
				}
			}

			if setting.getForbidden {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusForbidden,
				}
			}

			if setting.getGone {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: http.StatusGone,
				}
			}

			return &vpn.GatewayResponse{
				Id:    &setting.gatewayId,
				State: &setting.gatewayState,
			}, nil
		}),
	}
}

func TestCreateGatewayWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		mockSettings     []mockSettings
		wantGatewayState vpn.GatewayStatus
		wantErr          bool
		wantResp         bool
	}{
		{
			desc: "create_succeeded",
			mockSettings: []mockSettings{
				{gatewayState: vpn.GATEWAYSTATUS_READY, gatewayId: "gw-1"},
			},
			wantGatewayState: vpn.GATEWAYSTATUS_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc: "pending_multiple_times",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_READY,
					gatewayId:    "gw-1",
				},
			},
			wantGatewayState: vpn.GATEWAYSTATUS_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc: "error_state",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_ERROR,
					gatewayId:    "gw-1",
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "deleting_state",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_DELETING,
					gatewayId:    "gw-1",
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "get_fails",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					getFails: true,
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "unknown_state",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GatewayStatus("UNKNOWN_STATE"),
					gatewayId:    "gw-1",
				},
			},
			wantErr:  true,
			wantResp: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(tt.mockSettings)

				var wantRes *vpn.GatewayResponse
				if tt.wantResp {
					wantRes = &vpn.GatewayResponse{
						Id:    utils.Ptr("gw-1"),
						State: utils.Ptr(tt.wantGatewayState),
					}
				}

				handler := CreateGatewayWaitHandler(context.Background(), apiClient, "pid", vpn.REGION_EU01, "gw-1")

				gotRes, err := handler.WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !tt.wantErr && !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestUpdateGatewayWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		mockSettings     []mockSettings
		wantGatewayState vpn.GatewayStatus
		wantErr          bool
		wantResp         bool
	}{
		{
			desc: "update_succeeded",
			mockSettings: []mockSettings{
				{gatewayState: vpn.GATEWAYSTATUS_READY, gatewayId: "gw-1"},
			},
			wantGatewayState: vpn.GATEWAYSTATUS_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc: "pending_multiple_times",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_READY,
					gatewayId:    "gw-1",
				},
			},
			wantGatewayState: vpn.GATEWAYSTATUS_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc: "error_state",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_ERROR,
					gatewayId:    "gw-1",
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "deleting_state",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_DELETING,
					gatewayId:    "gw-1",
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "get_fails",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_PENDING,
					gatewayId:    "gw-1",
				},
				{
					getFails: true,
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "unknown_state",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GatewayStatus("UNKNOWN_STATE"),
					gatewayId:    "gw-1",
				},
			},
			wantErr:  true,
			wantResp: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(tt.mockSettings)

				var wantRes *vpn.GatewayResponse
				if tt.wantResp {
					wantRes = &vpn.GatewayResponse{
						Id:    utils.Ptr("gw-1"),
						State: utils.Ptr(tt.wantGatewayState),
					}
				}

				handler := UpdateGatewayWaitHandler(context.Background(), apiClient, "pid", vpn.REGION_EU01, "gw-1")

				gotRes, err := handler.WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !tt.wantErr && !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}

func TestDeleteGatewayWaitHandler(t *testing.T) {
	tests := []struct {
		desc         string
		mockSettings []mockSettings
		wantErr      bool
	}{
		{
			desc: "delete_succeeded",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GatewayStatus(""),
					getFails:     false,
					getNotFound:  true,
				},
			},
			wantErr: false,
		},
		{
			desc: "delete_succeeded_forbidden",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GatewayStatus(""),
					getForbidden: true,
				},
			},
			wantErr: false,
		},
		{
			desc: "delete_succeeded_gone",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GatewayStatus(""),
					getGone:      true,
				},
			},
			wantErr: false,
		},
		{
			desc: "delete_pending",
			mockSettings: []mockSettings{
				{
					getFails:     false,
					getNotFound:  false,
					gatewayState: vpn.GATEWAYSTATUS_DELETING,
					gatewayId:    "gw-1",
				},
				{
					getFails:     false,
					getNotFound:  false,
					gatewayState: vpn.GATEWAYSTATUS_DELETING,
					gatewayId:    "gw-1",
				},
				{
					getFails:     false,
					getNotFound:  true,
					gatewayState: vpn.GatewayStatus(""),
				},
			},
			wantErr: false,
		},
		{
			desc: "error_state",
			mockSettings: []mockSettings{
				{
					gatewayState: vpn.GATEWAYSTATUS_DELETING,
					gatewayId:    "gw-1",
				},
				{
					gatewayState: vpn.GATEWAYSTATUS_ERROR,
					gatewayId:    "gw-1",
				},
			},
			wantErr: true,
		},
		{
			desc: "timeout",
			mockSettings: []mockSettings{
				{
					getFails:     false,
					gatewayState: vpn.GatewayStatus("UNKNOWN_STATE"),
					gatewayId:    "gw-1",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newAPIMock(tt.mockSettings)

				handler := DeleteGatewayWaitHandler(context.Background(), apiClient, "pid", vpn.REGION_EU01, "gw-1")

				_, err := handler.WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}
