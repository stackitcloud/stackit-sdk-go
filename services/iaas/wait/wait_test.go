package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha"
)

type apiClientMocked struct {
	getNetworkAreaFails    bool
	getNetworkFails        bool
	getProjectRequestFails bool
	isDeleted              bool
	resourceState          string
	getVolumeFails         bool
	getServerFails         bool
	getAttachedVolumeFails bool
	getImageFails          bool
	isAttached             bool
	requestAction          string
	returnResizing         bool
}

func (a *apiClientMocked) GetNetworkAreaExecute(_ context.Context, _, _ string) (*iaas.NetworkArea, error) {
	if a.isDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	if a.getNetworkAreaFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &iaas.NetworkArea{
		AreaId: utils.Ptr("naid"),
		State:  &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetNetworkExecute(_ context.Context, _, _ string) (*iaas.Network, error) {
	if a.isDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	if a.getNetworkFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &iaas.Network{
		NetworkId: utils.Ptr("nid"),
		State:     &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetProjectRequestExecute(_ context.Context, _, _ string) (*iaas.Request, error) {
	if a.getProjectRequestFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &iaas.Request{
		RequestId:     utils.Ptr("rid"),
		RequestAction: &a.requestAction,
		Status:        &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetVolumeExecute(_ context.Context, _, _ string) (*iaas.Volume, error) {
	if a.isDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	if a.getVolumeFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &iaas.Volume{
		Id:     utils.Ptr("vid"),
		Status: &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetServerExecute(_ context.Context, _, _ string) (*iaas.Server, error) {
	if a.returnResizing {
		a.returnResizing = false
		return &iaas.Server{
			Id:     utils.Ptr("sid"),
			Status: utils.Ptr(ServerResizingStatus),
		}, nil
	}

	if a.isDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	if a.getServerFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &iaas.Server{
		Id:     utils.Ptr("sid"),
		Status: &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetAttachedVolumeExecute(_ context.Context, _, _, _ string) (*iaas.VolumeAttachment, error) {
	if a.getAttachedVolumeFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	if !a.isAttached {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	return &iaas.VolumeAttachment{
		ServerId: utils.Ptr("sid"),
		VolumeId: utils.Ptr("vid"),
	}, nil
}

func (a *apiClientMocked) GetImageExecute(_ context.Context, _, _ string) (*iaasalpha.Image, error) {
	if a.getImageFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	if a.isDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	return &iaasalpha.Image{
		Id:     utils.Ptr("iid"),
		Status: &a.resourceState,
	}, nil
}

func TestCreateNetworkAreaWaitHandler(t *testing.T) {
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
			apiClient := &apiClientMocked{
				getNetworkAreaFails: tt.getFails,
				resourceState:       tt.resourceState,
			}

			var wantRes *iaas.NetworkArea
			if tt.wantResp {
				wantRes = &iaas.NetworkArea{
					AreaId: utils.Ptr("naid"),
					State:  utils.Ptr(tt.resourceState),
				}
			}

			handler := CreateNetworkAreaWaitHandler(context.Background(), apiClient, "oid", "naid")

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

func TestUpdateNetworkAreaWaitHandler(t *testing.T) {
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
			apiClient := &apiClientMocked{
				getNetworkAreaFails: tt.getFails,
				resourceState:       tt.resourceState,
			}

			var wantRes *iaas.NetworkArea
			if tt.wantResp {
				wantRes = &iaas.NetworkArea{
					AreaId: utils.Ptr("naid"),
					State:  utils.Ptr(tt.resourceState),
				}
			}

			handler := UpdateNetworkAreaWaitHandler(context.Background(), apiClient, "oid", "naid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteNetworkAreaWaitHandler(t *testing.T) {
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
			apiClient := &apiClientMocked{
				getNetworkAreaFails: tt.getFails,
				isDeleted:           tt.isDeleted,
				resourceState:       tt.resourceState,
			}

			var wantRes *iaas.NetworkArea
			if tt.wantResp {
				wantRes = &iaas.NetworkArea{
					AreaId: utils.Ptr("naid"),
					State:  utils.Ptr(tt.resourceState),
				}
			}

			handler := DeleteNetworkAreaWaitHandler(context.Background(), apiClient, "oid", "naid")

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

func TestCreateNetworkWaitHandler(t *testing.T) {
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
			apiClient := &apiClientMocked{
				getNetworkFails: tt.getFails,
				resourceState:   tt.resourceState,
			}

			var wantRes *iaas.Network
			if tt.wantResp {
				wantRes = &iaas.Network{
					NetworkId: utils.Ptr("nid"),
					State:     utils.Ptr(tt.resourceState),
				}
			}

			handler := CreateNetworkWaitHandler(context.Background(), apiClient, "pid", "nid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestUpdateNetworkWaitHandler(t *testing.T) {
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
			apiClient := &apiClientMocked{
				getNetworkFails: tt.getFails,
				resourceState:   tt.resourceState,
			}

			var wantRes *iaas.Network
			if tt.wantResp {
				wantRes = &iaas.Network{
					NetworkId: utils.Ptr("nid"),
					State:     utils.Ptr(tt.resourceState),
				}
			}

			handler := UpdateNetworkWaitHandler(context.Background(), apiClient, "pid", "nid")

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteNetworkWaitHandler(t *testing.T) {
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
			apiClient := &apiClientMocked{
				getNetworkFails: tt.getFails,
				isDeleted:       tt.isDeleted,
				resourceState:   tt.resourceState,
			}

			var wantRes *iaas.Network
			if tt.wantResp {
				wantRes = &iaas.Network{
					NetworkId: utils.Ptr("nid"),
					State:     utils.Ptr(tt.resourceState),
				}
			}

			handler := DeleteNetworkWaitHandler(context.Background(), apiClient, "pid", "nid")

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

func TestCreateVolumeWaitHandler(t *testing.T) {
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
			resourceState: VolumeAvailableStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			resourceState: ErrorStatus,
			wantErr:       true,
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getVolumeFails: tt.getFails,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Volume
			if tt.wantResp {
				wantRes = &iaas.Volume{
					Id:     utils.Ptr("vid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := CreateVolumeWaitHandler(context.Background(), apiClient, "pid", "vid")

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

func TestDeleteVolumeWaitHandler(t *testing.T) {
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getVolumeFails: tt.getFails,
				isDeleted:      tt.isDeleted,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Volume
			if tt.wantResp {
				wantRes = &iaas.Volume{
					Id:     utils.Ptr("vid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := DeleteVolumeWaitHandler(context.Background(), apiClient, "pid", "vid")

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

func TestCreateServerWaitHandler(t *testing.T) {
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
			resourceState: ServerActiveStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			resourceState: ErrorStatus,
			wantErr:       true,
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getServerFails: tt.getFails,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Server
			if tt.wantResp {
				wantRes = &iaas.Server{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := CreateServerWaitHandler(context.Background(), apiClient, "pid", "sid")

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

func TestDeleteServerWaitHandler(t *testing.T) {
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getServerFails: tt.getFails,
				isDeleted:      tt.isDeleted,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Server
			if tt.wantResp {
				wantRes = &iaas.Server{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := DeleteServerWaitHandler(context.Background(), apiClient, "pid", "sid")

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

func TestResizeServerWaitHandler(t *testing.T) {
	tests := []struct {
		desc               string
		getFails           bool
		returnResizing     bool
		finalResourceState string
		wantErr            bool
		wantResp           bool
	}{
		{
			desc:               "resize_succeeded",
			getFails:           false,
			returnResizing:     true,
			finalResourceState: ServerActiveStatus,
			wantErr:            false,
			wantResp:           true,
		},
		{
			desc:               "resizing_status_is_never_returned",
			getFails:           false,
			returnResizing:     false,
			finalResourceState: ServerActiveStatus,
			wantErr:            true,
			wantResp:           true,
		},
		{
			desc:               "error_status",
			getFails:           false,
			returnResizing:     true,
			finalResourceState: ErrorStatus,
			wantErr:            true,
			wantResp:           true,
		},
		{
			desc:               "get_fails",
			getFails:           true,
			finalResourceState: "",
			wantErr:            true,
			wantResp:           false,
		},
		{
			desc:               "timeout",
			getFails:           false,
			finalResourceState: "ANOTHER Status",
			wantErr:            true,
			wantResp:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getServerFails: tt.getFails,
				resourceState:  tt.finalResourceState,
				returnResizing: tt.returnResizing,
			}

			var wantRes *iaas.Server
			if tt.wantResp {
				wantRes = &iaas.Server{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.finalResourceState),
				}
			}

			handler := ResizeServerWaitHandler(context.Background(), apiClient, "pid", "sid")

			gotRes, err := handler.SetThrottle(1 * time.Millisecond).SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestStartServerWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "start_succeeded",
			getFails:      false,
			resourceState: ServerActiveStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			resourceState: ErrorStatus,
			wantErr:       true,
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getServerFails: tt.getFails,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Server
			if tt.wantResp {
				wantRes = &iaas.Server{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := StartServerWaitHandler(context.Background(), apiClient, "pid", "sid")

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

func TestStopServerWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "stop_succeeded",
			getFails:      false,
			resourceState: ServerInactiveStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			resourceState: ErrorStatus,
			wantErr:       true,
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getServerFails: tt.getFails,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Server
			if tt.wantResp {
				wantRes = &iaas.Server{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := StopServerWaitHandler(context.Background(), apiClient, "pid", "sid")

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

func TestDeallocateServerWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "deallocate_succeeded",
			getFails:      false,
			resourceState: ServerDeallocatedStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			resourceState: ErrorStatus,
			wantErr:       true,
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getServerFails: tt.getFails,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Server
			if tt.wantResp {
				wantRes = &iaas.Server{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := DeallocateServerWaitHandler(context.Background(), apiClient, "pid", "sid")

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

func TestRescueServerWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "rescue_succeeded",
			getFails:      false,
			resourceState: ServerRescueStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			resourceState: ErrorStatus,
			wantErr:       true,
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getServerFails: tt.getFails,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Server
			if tt.wantResp {
				wantRes = &iaas.Server{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := RescueServerWaitHandler(context.Background(), apiClient, "pid", "sid")

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

func TestUnrescueServerWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "unrescue_succeeded",
			getFails:      false,
			resourceState: ServerActiveStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			resourceState: ErrorStatus,
			wantErr:       true,
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getServerFails: tt.getFails,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Server
			if tt.wantResp {
				wantRes = &iaas.Server{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := UnrescueServerWaitHandler(context.Background(), apiClient, "pid", "sid")

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

func TestProjectRequestWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		requestState  string
		requestAction string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			requestAction: RequestCreateAction,
			requestState:  RequestCreatedStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_succeeded",
			getFails:      false,
			requestAction: RequestUpdateAction,
			requestState:  RequestUpdatedStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "delete_succeeded",
			getFails:      false,
			requestAction: RequestDeleteAction,
			requestState:  RequestDeletedStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "unsupported_action",
			getFails:      false,
			requestAction: "OTHER_ACTION",
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			requestAction: RequestCreateAction,
			requestState:  ErrorStatus,
			wantErr:       true,
			wantResp:      true,
		},
		{
			desc:         "get_fails",
			getFails:     true,
			requestState: "",
			wantErr:      true,
			wantResp:     false,
		},
		{
			desc:          "timeout",
			getFails:      false,
			requestAction: RequestCreateAction,
			requestState:  "ANOTHER Status",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getProjectRequestFails: tt.getFails,
				requestAction:          tt.requestAction,
				resourceState:          tt.requestState,
			}

			var wantRes *iaas.Request
			if tt.wantResp {
				wantRes = &iaas.Request{
					RequestId:     utils.Ptr("rid"),
					RequestAction: utils.Ptr(tt.requestAction),
					Status:        utils.Ptr(tt.requestState),
				}
			}

			handler := ProjectRequestWaitHandler(context.Background(), apiClient, "pid", "rid")

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

func TestAddVolumeToServerWaitHandler(t *testing.T) {
	tests := []struct {
		desc       string
		getFails   bool
		isAttached bool
		wantErr    bool
		wantResp   bool
	}{
		{
			desc:       "attachment_succeeded",
			getFails:   false,
			isAttached: true,
			wantErr:    false,
			wantResp:   true,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:       "timeout",
			getFails:   false,
			isAttached: false,
			wantErr:    true,
			wantResp:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getAttachedVolumeFails: tt.getFails,
				isAttached:             tt.isAttached,
			}

			var wantRes *iaas.VolumeAttachment
			if tt.wantResp {
				wantRes = &iaas.VolumeAttachment{
					ServerId: utils.Ptr("sid"),
					VolumeId: utils.Ptr("vid"),
				}
			}

			handler := AddVolumeToServerWaitHandler(context.Background(), apiClient, "pid", "sid", "vid")

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

func TestRemoveVolumeFromServerWaitHandler(t *testing.T) {
	tests := []struct {
		desc       string
		getFails   bool
		isAttached bool
		wantErr    bool
		wantResp   bool
	}{
		{
			desc:       "removal_succeeded",
			getFails:   false,
			isAttached: false,
			wantErr:    false,
			wantResp:   false,
		},
		{
			desc:     "get_fails",
			getFails: true,
			wantErr:  true,
			wantResp: false,
		},
		{
			desc:       "timeout",
			getFails:   false,
			isAttached: true,
			wantErr:    true,
			wantResp:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getAttachedVolumeFails: tt.getFails,
				isAttached:             tt.isAttached,
			}

			var wantRes *iaas.VolumeAttachment
			if tt.wantResp {
				wantRes = &iaas.VolumeAttachment{
					ServerId: utils.Ptr("sid"),
					VolumeId: utils.Ptr("vid"),
				}
			}

			handler := RemoveVolumeFromServerWaitHandler(context.Background(), apiClient, "pid", "sid", "vid")

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

func TestUploadImageWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "upload_succeeded",
			getFails:      false,
			resourceState: ImageAvailableStatus,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			resourceState: ErrorStatus,
			wantErr:       true,
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
			resourceState: "ANOTHER Status",
			wantErr:       true,
			wantResp:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getImageFails: tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *iaasalpha.Image
			if tt.wantResp {
				wantRes = &iaasalpha.Image{
					Id:     utils.Ptr("iid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := UploadImageWaitHandler(context.Background(), apiClient, "pid", "iid")

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

func TestDeleteImageWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		isDeleted     bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:      "delete_succeeded",
			getFails:  false,
			isDeleted: true,
			wantErr:   false,
		},
		{
			desc:          "get_fails",
			getFails:      true,
			resourceState: "",
			wantErr:       true,
		},
		{
			desc:          "timeout",
			getFails:      false,
			resourceState: "ANOTHER Status",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getImageFails: tt.getFails,
				isDeleted:     tt.isDeleted,
				resourceState: tt.resourceState,
			}

			handler := DeleteImageWaitHandler(context.Background(), apiClient, "pid", "iid")

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
