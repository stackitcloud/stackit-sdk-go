package wait

import (
	"context"
	"fmt"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/core/wait"
	iaas "github.com/stackitcloud/stackit-sdk-go/services/iaas/v2alpha1api"
)

type iaasMockSettings struct {
	getVPCRegionFails       bool
	getVPCNetworkRangeFails bool
	isDeleted               bool
	resourceState           string
}

func newIaaSAPIMock(settings *iaasMockSettings) iaas.DefaultAPI {
	return &iaas.DefaultAPIServiceMock{
		GetVPCRegionExecuteMock: utils.Ptr(func(_ iaas.ApiGetVPCRegionRequest) (*iaas.RegionalVPC, error) {
			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			if settings.getVPCRegionFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.RegionalVPC{
				Status: &settings.resourceState,
			}, nil
		}),
		GetVPCNetworkRangeExecuteMock: utils.Ptr(func(_ iaas.ApiGetVPCNetworkRangeRequest) (*iaas.VPCNetworkRange, error) {
			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			if settings.getVPCNetworkRangeFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.VPCNetworkRange{
				VPCNetworkRangeIPv4: &iaas.VPCNetworkRangeIPv4{
					Status: &settings.resourceState,
				},
			}, nil
		}),
	}
}

func TestCreateNetworkAreaRegionWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: CreateSuccess,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getVPCRegionFails: tt.getFails,
					resourceState:     tt.resourceState,
				})

				var wantRes *iaas.RegionalVPC
				if tt.wantResp {
					wantRes = &iaas.RegionalVPC{
						Status: utils.Ptr(tt.resourceState),
					}
				}

				handler := CreateVPCRegionWaitHandler(context.Background(), apiClient, "pid", "aid", "region")

				gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

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

func TestDeleteNetworkAreaRegionWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		isDeleted     bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:      "delete_succeeded",
			getFails:  false,
			isDeleted: true,
			wantErr:   false,
			wantResp:  false,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getVPCRegionFails: tt.getFails,
					isDeleted:         tt.isDeleted,
					resourceState:     tt.resourceState,
				})

				var wantRes *iaas.RegionalVPC
				if tt.wantResp {
					wantRes = &iaas.RegionalVPC{
						Status: utils.Ptr(tt.resourceState),
					}
				}

				handler := DeleteVPCRegionWaitHandler(context.Background(), apiClient, "pid", "region", "nid")

				gotRes, err := handler.SetSleepBeforeWait(0).SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

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

func TestCreateOrUpdateVPCNetworkRangeWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: UpdateSuccess,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: CreateSuccess,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      true,
		},
	}

	handlers := map[string]func(context.Context, iaas.DefaultAPI, string, string, string, string) *wait.AsyncActionHandler[iaas.VPCNetworkRange]{
		"common logic": createOrUpdateVPCNetworkRangeWaitHandler,
		"create":       CreateVPCNetworkRangeWaitHandler,
		"update":       UpdateVPCNetworkRangeWaitHandler,
	}

	for handlerDesc, handlerFn := range handlers {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s - %s", handlerDesc, tt.desc), func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					apiClient := newIaaSAPIMock(&iaasMockSettings{
						getVPCNetworkRangeFails: tt.getFails,
						resourceState:           tt.resourceState,
					})

					var wantRes *iaas.VPCNetworkRange
					if tt.wantResp {
						wantRes = &iaas.VPCNetworkRange{
							VPCNetworkRangeIPv4: &iaas.VPCNetworkRangeIPv4{
								Status: utils.Ptr(tt.resourceState),
							},
						}
					}

					handler := handlerFn(context.Background(), apiClient, "pid", "vpcId", "region", "nid")

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
}

func TestDeleteVPCNetworkRangeWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		isDeleted     bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:      "delete_succeeded",
			getFails:  false,
			isDeleted: true,
			wantErr:   false,
			wantResp:  false,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
			wantResp:      false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getVPCNetworkRangeFails: tt.getFails,
					isDeleted:               tt.isDeleted,
					resourceState:           tt.resourceState,
				})

				var wantRes *iaas.VPCNetworkRange
				if tt.wantResp {
					wantRes = &iaas.VPCNetworkRange{
						VPCNetworkRangeIPv4: &iaas.VPCNetworkRangeIPv4{
							Status: utils.Ptr(tt.resourceState),
						},
					}
				}

				handler := DeleteVPCNetworkRangeWaitHandler(context.Background(), apiClient, "pid", "vpcId", "region", "nid")

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
