package wait

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
	"github.com/stackitcloud/stackit-sdk-go/services/resourcemanager"
)

type apiClientMocked struct {
	getNetworkAreaRegionFails bool
	getNetworkAreaFails       bool
	getNetworkFails           bool
	getProjectRequestFails    bool
	isDeleted                 bool
	resourceState             string
	getVolumeFails            bool
	getServerFails            bool
	getAttachedVolumeFails    bool
	getImageFails             bool
	getBackupFails            bool
	isAttached                bool
	requestAction             string
	returnResizing            bool
	getSnapshotFails          bool
	listProjectsResponses     []listProjectsResponses
	listProjectsResponsesIdx  int
}

type listProjectsResponses struct {
	resp *iaas.ProjectListResponse
	err  error
}

type resourceManagerClientMocked struct {
	getProjectResponses    []getProjectResponse
	getProjectResponsesIdx int
}

type getProjectResponse struct {
	resp *resourcemanager.GetProjectResponse
	err  error
}

func (r *resourceManagerClientMocked) GetProjectExecute(_ context.Context, _ string) (*resourcemanager.GetProjectResponse, error) {
	resp := r.getProjectResponses[r.getProjectResponsesIdx].resp
	err := r.getProjectResponses[r.getProjectResponsesIdx].err
	r.getProjectResponsesIdx = (r.getProjectResponsesIdx + 1) % len(r.getProjectResponses)
	return resp, err
}

func (a *apiClientMocked) ListNetworkAreaProjectsExecute(_ context.Context, _, _ string) (*iaas.ProjectListResponse, error) {
	resp := a.listProjectsResponses[a.listProjectsResponsesIdx].resp
	err := a.listProjectsResponses[a.listProjectsResponsesIdx].err
	a.listProjectsResponsesIdx = (a.listProjectsResponsesIdx + 1) % len(a.listProjectsResponses)
	return resp, err
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
		Id: utils.Ptr("naid"),
	}, nil
}

func (a *apiClientMocked) GetNetworkAreaRegionExecute(_ context.Context, _, _, _ string) (*iaas.RegionalArea, error) {
	if a.isDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	if a.getNetworkAreaRegionFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &iaas.RegionalArea{
		Status: &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetNetworkExecute(_ context.Context, _, _, _ string) (*iaas.Network, error) {
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
		Id:     utils.Ptr("nid"),
		Status: &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetProjectRequestExecute(_ context.Context, _, _, _ string) (*iaas.Request, error) {
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

func (a *apiClientMocked) GetVolumeExecute(_ context.Context, _, _, _ string) (*iaas.Volume, error) {
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

func (a *apiClientMocked) GetServerExecute(_ context.Context, _, _, _ string) (*iaas.Server, error) {
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

func (a *apiClientMocked) GetAttachedVolumeExecute(_ context.Context, _, _, _, _ string) (*iaas.VolumeAttachment, error) {
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

func (a *apiClientMocked) GetImageExecute(_ context.Context, _, _, _ string) (*iaas.Image, error) {
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

	return &iaas.Image{
		Id:     utils.Ptr("iid"),
		Status: &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetBackupExecute(_ context.Context, _, _, _ string) (*iaas.Backup, error) {
	if a.isDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	if a.getBackupFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &iaas.Backup{
		Id:     utils.Ptr("bid"),
		Status: &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetSnapshotExecute(_ context.Context, _, _, _ string) (*iaas.Snapshot, error) {
	if a.isDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	if a.getSnapshotFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &iaas.Snapshot{
		Id:     utils.Ptr("sid"),
		Status: &a.resourceState,
	}, nil
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
					Id:     utils.Ptr("nid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := CreateNetworkWaitHandler(context.Background(), apiClient, "pid", "region", "nid")

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
					Id:     utils.Ptr("nid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := UpdateNetworkWaitHandler(context.Background(), apiClient, "pid", "region", "nid")

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
					Id:     utils.Ptr("nid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := DeleteNetworkWaitHandler(context.Background(), apiClient, "pid", "region", "nid")

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

			handler := CreateVolumeWaitHandler(context.Background(), apiClient, "pid", "region", "vid")

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

			handler := DeleteVolumeWaitHandler(context.Background(), apiClient, "pid", "region", "vid")

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

			handler := CreateServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid")

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

			handler := DeleteServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid")

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

			handler := ResizeServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid")

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

			handler := StartServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid")

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

			handler := StopServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid")

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

			handler := DeallocateServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid")

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

			handler := RescueServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid")

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

			handler := UnrescueServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid")

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

			handler := ProjectRequestWaitHandler(context.Background(), apiClient, "pid", "region", "rid")

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

			handler := AddVolumeToServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid", "vid")

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

			handler := RemoveVolumeFromServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid", "vid")

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

			var wantRes *iaas.Image
			if tt.wantResp {
				wantRes = &iaas.Image{
					Id:     utils.Ptr("iid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := UploadImageWaitHandler(context.Background(), apiClient, "pid", "region", "iid")

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

			handler := DeleteImageWaitHandler(context.Background(), apiClient, "pid", "region", "iid")

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateBackupWaitHandler(t *testing.T) {
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
			resourceState: BackupAvailableStatus,
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
			resourceState: "ANOTHER_STATUS",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getBackupFails: tt.getFails,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Backup
			if tt.wantResp {
				wantRes = &iaas.Backup{
					Id:     utils.Ptr("bid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := CreateBackupWaitHandler(context.Background(), apiClient, "pid", "region", "bid")
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

func TestDeleteBackupWaitHandler(t *testing.T) {
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
			resourceState: "ANOTHER_STATUS",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getBackupFails: tt.getFails,
				isDeleted:      tt.isDeleted,
				resourceState:  tt.resourceState,
			}

			handler := DeleteBackupWaitHandler(context.Background(), apiClient, "pid", "region", "bid")
			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotRes != nil {
				t.Fatalf("handler gotRes = %v, want nil", gotRes)
			}
		})
	}
}

func TestRestoreBackupWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "restore_succeeded",
			getFails:      false,
			resourceState: BackupAvailableStatus,
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
			resourceState: "ANOTHER_STATUS",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getBackupFails: tt.getFails,
				resourceState:  tt.resourceState,
			}

			var wantRes *iaas.Backup
			if tt.wantResp {
				wantRes = &iaas.Backup{
					Id:     utils.Ptr("bid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := RestoreBackupWaitHandler(context.Background(), apiClient, "pid", "region", "bid")
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

func TestCreateSnapshotWaitHandler(t *testing.T) {
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
			resourceState: SnapshotAvailableStatus,
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
			resourceState: "ANOTHER_STATUS",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getSnapshotFails: tt.getFails,
				resourceState:    tt.resourceState,
			}

			var wantRes *iaas.Snapshot
			if tt.wantResp {
				wantRes = &iaas.Snapshot{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := CreateSnapshotWaitHandler(context.Background(), apiClient, "pid", "region", "sid")
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

func TestDeleteSnapshotWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: DeleteSuccess,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "error_status",
			getFails:      false,
			resourceState: ErrorStatus,
			wantErr:       true,
			wantResp:      false,
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
			resourceState: "ANOTHER_STATUS",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getSnapshotFails: tt.getFails,
				resourceState:    tt.resourceState,
			}

			var wantRes *iaas.Snapshot
			if tt.wantResp {
				wantRes = &iaas.Snapshot{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := DeleteSnapshotWaitHandler(context.Background(), apiClient, "pid", "region", "sid")
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

func TestReadyForNetworkAreaDeletionWaitHandler(t *testing.T) {
	tests := []struct {
		desc                  string
		listProjectsResponses []listProjectsResponses
		getProjectResponses   []getProjectResponse
		wantErr               bool
		wantResp              bool
	}{
		{
			desc: "projects still active",
			listProjectsResponses: []listProjectsResponses{
				{
					resp: &iaas.ProjectListResponse{
						Items: utils.Ptr([]string{"project1", "project2"}),
					},
					err: nil,
				},
			},
			getProjectResponses: []getProjectResponse{
				{&resourcemanager.GetProjectResponse{}, nil},
				{&resourcemanager.GetProjectResponse{}, nil},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "no projects - ready for deletion",
			listProjectsResponses: []listProjectsResponses{
				{
					resp: &iaas.ProjectListResponse{
						Items: utils.Ptr([]string{}),
					},
					err: nil,
				},
			},
			getProjectResponses: []getProjectResponse{},
			wantErr:             false,
			wantResp:            true,
		},
		{
			desc: "projects in deletion state",
			listProjectsResponses: []listProjectsResponses{
				{
					resp: &iaas.ProjectListResponse{
						Items: utils.Ptr([]string{"project1", "project2"}),
					},
					err: nil,
				},
				{
					resp: &iaas.ProjectListResponse{
						Items: utils.Ptr([]string{}),
					},
					err: nil,
				},
			},
			getProjectResponses: []getProjectResponse{
				{nil, oapierror.NewError(http.StatusForbidden, "")},
				{nil, oapierror.NewError(http.StatusForbidden, "")},
			},
			wantErr:  false,
			wantResp: true,
		},
		{
			desc: "network area includes one active project",
			listProjectsResponses: []listProjectsResponses{
				{
					resp: &iaas.ProjectListResponse{
						Items: utils.Ptr([]string{"project1", "project2", "project3"}),
					},
					err: nil,
				},
			},
			getProjectResponses: []getProjectResponse{
				{nil, oapierror.NewError(http.StatusForbidden, "")},
				{nil, oapierror.NewError(http.StatusForbidden, "")},
				{&resourcemanager.GetProjectResponse{}, nil},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "network area not found",
			listProjectsResponses: []listProjectsResponses{
				{
					resp: nil,
					err:  oapierror.NewError(http.StatusNotFound, "not found"),
				},
			},
			getProjectResponses: []getProjectResponse{},
			wantErr:             true,
			wantResp:            false,
		},
		{
			desc: "network area items is nil",
			listProjectsResponses: []listProjectsResponses{
				{
					resp: &iaas.ProjectListResponse{
						Items: nil,
					},
				},
			},
			wantErr:  true,
			wantResp: false,
		},
		{
			desc: "timeout",
			listProjectsResponses: []listProjectsResponses{
				{
					resp: &iaas.ProjectListResponse{
						Items: utils.Ptr([]string{"project1"}),
					},
					err: nil,
				},
			},
			getProjectResponses: []getProjectResponse{
				{nil, oapierror.NewError(http.StatusForbidden, "")},
			},
			wantErr:  true,
			wantResp: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				listProjectsResponses: tt.listProjectsResponses,
			}

			rmApiClient := &resourceManagerClientMocked{
				getProjectResponses: tt.getProjectResponses,
			}

			var wantRes *iaas.ProjectListResponse
			if tt.wantResp {
				wantRes = tt.listProjectsResponses[len(tt.listProjectsResponses)-1].resp
			}

			handler := ReadyForNetworkAreaDeletionWaitHandler(context.Background(), apiClient, rmApiClient, "oid", "aid")
			gotRes, err := handler.SetTimeout(200 * time.Millisecond).SetThrottle(5 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestCreateRegionalNetworkAreaConfigurationWaitHandler(t *testing.T) {
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
				getNetworkAreaRegionFails: tt.getFails,
				resourceState:             tt.resourceState,
			}

			var wantRes *iaas.RegionalArea
			if tt.wantResp {
				wantRes = &iaas.RegionalArea{
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := CreateNetworkAreaRegionWaitHandler(context.Background(), apiClient, "pid", "aid", "region")

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

func TestDeleteRegionalNetworkAreaConfigurationWaitHandler(t *testing.T) {
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
				getNetworkAreaRegionFails: tt.getFails,
				isDeleted:                 tt.isDeleted,
				resourceState:             tt.resourceState,
			}

			var wantRes *iaas.RegionalArea
			if tt.wantResp {
				wantRes = &iaas.RegionalArea{
					Status: utils.Ptr(tt.resourceState),
				}
			}

			handler := DeleteRegionalNetworkAreaConfigurationWaitHandler(context.Background(), apiClient, "pid", "region", "nid")

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
