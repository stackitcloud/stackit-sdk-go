package wait

import (
	"context"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	iaas "github.com/stackitcloud/stackit-sdk-go/services/iaas/v2api"
	resourcemanager "github.com/stackitcloud/stackit-sdk-go/services/resourcemanager/v0api"
)

type iaasMockSettings struct {
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

type resourcemanagerMockSettings struct {
	getProjectResponses    []getProjectResponse
	getProjectResponsesIdx int
}

type getProjectResponse struct {
	resp *resourcemanager.GetProjectResponse
	err  error
}

func newResourcemanagerAPIMock(settings *resourcemanagerMockSettings) resourcemanager.DefaultAPI {
	return &resourcemanager.DefaultAPIServiceMock{
		GetProjectExecuteMock: utils.Ptr(func(_ resourcemanager.ApiGetProjectRequest) (*resourcemanager.GetProjectResponse, error) {
			resp := settings.getProjectResponses[settings.getProjectResponsesIdx].resp
			err := settings.getProjectResponses[settings.getProjectResponsesIdx].err
			settings.getProjectResponsesIdx = (settings.getProjectResponsesIdx + 1) % len(settings.getProjectResponses)
			return resp, err
		}),
	}
}

func newIaaSAPIMock(settings *iaasMockSettings) iaas.DefaultAPI {
	return &iaas.DefaultAPIServiceMock{
		ListNetworkAreaProjectsExecuteMock: utils.Ptr(func(_ iaas.ApiListNetworkAreaProjectsRequest) (*iaas.ProjectListResponse, error) {
			resp := settings.listProjectsResponses[settings.listProjectsResponsesIdx].resp
			err := settings.listProjectsResponses[settings.listProjectsResponsesIdx].err
			settings.listProjectsResponsesIdx = (settings.listProjectsResponsesIdx + 1) % len(settings.listProjectsResponses)
			return resp, err
		}),
		GetNetworkAreaExecuteMock: utils.Ptr(func(_ iaas.ApiGetNetworkAreaRequest) (*iaas.NetworkArea, error) {
			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			if settings.getNetworkAreaFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.NetworkArea{
				Id: utils.Ptr("naid"),
			}, nil
		}),
		GetNetworkAreaRegionExecuteMock: utils.Ptr(func(_ iaas.ApiGetNetworkAreaRegionRequest) (*iaas.RegionalArea, error) {
			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			if settings.getNetworkAreaRegionFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.RegionalArea{
				Status: &settings.resourceState,
			}, nil
		}),
		GetNetworkExecuteMock: utils.Ptr(func(_ iaas.ApiGetNetworkRequest) (*iaas.Network, error) {
			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			if settings.getNetworkFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.Network{
				Id:     "nid",
				Status: settings.resourceState,
			}, nil
		}),
		GetProjectRequestExecuteMock: utils.Ptr(func(_ iaas.ApiGetProjectRequestRequest) (*iaas.Request, error) {
			if settings.getProjectRequestFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.Request{
				RequestId:     "rid",
				RequestAction: settings.requestAction,
				Status:        settings.resourceState,
			}, nil
		}),
		GetVolumeExecuteMock: utils.Ptr(func(_ iaas.ApiGetVolumeRequest) (*iaas.Volume, error) {
			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			if settings.getVolumeFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.Volume{
				Id:     utils.Ptr("vid"),
				Status: &settings.resourceState,
			}, nil
		}),
		GetServerExecuteMock: utils.Ptr(func(_ iaas.ApiGetServerRequest) (*iaas.Server, error) {
			if settings.returnResizing {
				settings.returnResizing = false
				return &iaas.Server{
					Id:     utils.Ptr("sid"),
					Status: utils.Ptr(ServerResizingStatus),
				}, nil
			}

			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			if settings.getServerFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.Server{
				Id:     utils.Ptr("sid"),
				Status: &settings.resourceState,
			}, nil
		}),
		GetAttachedVolumeExecuteMock: utils.Ptr(func(_ iaas.ApiGetAttachedVolumeRequest) (*iaas.VolumeAttachment, error) {
			if settings.getAttachedVolumeFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if !settings.isAttached {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &iaas.VolumeAttachment{
				ServerId: utils.Ptr("sid"),
				VolumeId: utils.Ptr("vid"),
			}, nil
		}),
		GetImageExecuteMock: utils.Ptr(func(_ iaas.ApiGetImageRequest) (*iaas.Image, error) {
			if settings.getImageFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &iaas.Image{
				Id:     utils.Ptr("iid"),
				Status: &settings.resourceState,
			}, nil
		}),
		GetBackupExecuteMock: utils.Ptr(func(_ iaas.ApiGetBackupRequest) (*iaas.Backup, error) {
			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			if settings.getBackupFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.Backup{
				Id:     utils.Ptr("bid"),
				Status: &settings.resourceState,
			}, nil
		}),
		GetSnapshotExecuteMock: utils.Ptr(func(_ iaas.ApiGetSnapshotRequest) (*iaas.Snapshot, error) {
			if settings.isDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			if settings.getSnapshotFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &iaas.Snapshot{
				Id:     utils.Ptr("sid"),
				Status: &settings.resourceState,
			}, nil
		}),
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getNetworkFails: tt.getFails,
					resourceState:   tt.resourceState,
				})

				var wantRes *iaas.Network
				if tt.wantResp {
					wantRes = &iaas.Network{
						Id:     "nid",
						Status: tt.resourceState,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getNetworkFails: tt.getFails,
					resourceState:   tt.resourceState,
				})

				var wantRes *iaas.Network
				if tt.wantResp {
					wantRes = &iaas.Network{
						Id:     "nid",
						Status: tt.resourceState,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getNetworkFails: tt.getFails,
					isDeleted:       tt.isDeleted,
					resourceState:   tt.resourceState,
				})

				var wantRes *iaas.Network
				if tt.wantResp {
					wantRes = &iaas.Network{
						Id:     "nid",
						Status: tt.resourceState,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getVolumeFails: tt.getFails,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getVolumeFails: tt.getFails,
					isDeleted:      tt.isDeleted,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getServerFails: tt.getFails,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getServerFails: tt.getFails,
					isDeleted:      tt.isDeleted,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getServerFails: tt.getFails,
					resourceState:  tt.finalResourceState,
					returnResizing: tt.returnResizing,
				})

				var wantRes *iaas.Server
				if tt.wantResp {
					wantRes = &iaas.Server{
						Id:     utils.Ptr("sid"),
						Status: utils.Ptr(tt.finalResourceState),
					}
				}

				handler := ResizeServerWaitHandler(context.Background(), apiClient, "pid", "region", "sid")

				gotRes, err := handler.WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getServerFails: tt.getFails,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getServerFails: tt.getFails,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getServerFails: tt.getFails,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getServerFails: tt.getFails,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getServerFails: tt.getFails,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getProjectRequestFails: tt.getFails,
					requestAction:          tt.requestAction,
					resourceState:          tt.requestState,
				})

				var wantRes *iaas.Request
				if tt.wantResp {
					wantRes = &iaas.Request{
						RequestId:     "rid",
						RequestAction: tt.requestAction,
						Status:        tt.requestState,
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getAttachedVolumeFails: tt.getFails,
					isAttached:             tt.isAttached,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getAttachedVolumeFails: tt.getFails,
					isAttached:             tt.isAttached,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getImageFails: tt.getFails,
					resourceState: tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getImageFails: tt.getFails,
					isDeleted:     tt.isDeleted,
					resourceState: tt.resourceState,
				})

				handler := DeleteImageWaitHandler(context.Background(), apiClient, "pid", "region", "iid")

				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getBackupFails: tt.getFails,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getBackupFails: tt.getFails,
					isDeleted:      tt.isDeleted,
					resourceState:  tt.resourceState,
				})

				handler := DeleteBackupWaitHandler(context.Background(), apiClient, "pid", "region", "bid")
				gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if gotRes != nil {
					t.Fatalf("handler gotRes = %v, want nil", gotRes)
				}
			})
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getBackupFails: tt.getFails,
					resourceState:  tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getSnapshotFails: tt.getFails,
					resourceState:    tt.resourceState,
				})

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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					getSnapshotFails: tt.getFails,
					resourceState:    tt.resourceState,
				})

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
						Items: []string{"project1", "project2"},
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
						Items: []string{},
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
						Items: []string{"project1", "project2"},
					},
					err: nil,
				},
				{
					resp: &iaas.ProjectListResponse{
						Items: []string{},
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
						Items: []string{"project1", "project2", "project3"},
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
						Items: []string{"project1"},
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
			synctest.Test(t, func(t *testing.T) {
				apiClient := newIaaSAPIMock(&iaasMockSettings{
					listProjectsResponses: tt.listProjectsResponses,
				})

				rmApiClient := newResourcemanagerAPIMock(&resourcemanagerMockSettings{
					getProjectResponses: tt.getProjectResponses,
				})

				var wantRes *iaas.ProjectListResponse
				if tt.wantResp {
					wantRes = tt.listProjectsResponses[len(tt.listProjectsResponses)-1].resp
				}

				handler := ReadyForNetworkAreaDeletionWaitHandler(context.Background(), apiClient, rmApiClient, "oid", "aid")
				gotRes, err := handler.WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
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
					getNetworkAreaRegionFails: tt.getFails,
					resourceState:             tt.resourceState,
				})

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
					getNetworkAreaRegionFails: tt.getFails,
					isDeleted:                 tt.isDeleted,
					resourceState:             tt.resourceState,
				})

				var wantRes *iaas.RegionalArea
				if tt.wantResp {
					wantRes = &iaas.RegionalArea{
						Status: utils.Ptr(tt.resourceState),
					}
				}

				handler := DeleteNetworkAreaRegionWaitHandler(context.Background(), apiClient, "pid", "region", "nid")

				gotRes, err := handler.SetSleepBeforeWait(0).SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
				if !cmp.Equal(gotRes, wantRes) {
					t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
				}
			})
		})
	}
}
