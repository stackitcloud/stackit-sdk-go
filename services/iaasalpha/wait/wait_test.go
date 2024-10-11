package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaasalpha"
)

type apiClientMocked struct {
	getVolumeFails         bool
	getServerFails         bool
	getProjectRequestFails bool
	isDeleted              bool
	resourceState          string
	requestAction          string
	returnResizing         bool
}

func (a *apiClientMocked) GetVolumeExecute(_ context.Context, _, _ string) (*iaasalpha.Volume, error) {
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

	return &iaasalpha.Volume{
		Id:     utils.Ptr("vid"),
		Status: &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetServerExecute(_ context.Context, _, _ string) (*iaasalpha.Server, error) {
	if a.returnResizing {
		a.returnResizing = false
		return &iaasalpha.Server{
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

	return &iaasalpha.Server{
		Id:     utils.Ptr("sid"),
		Status: &a.resourceState,
	}, nil
}

func (a *apiClientMocked) GetProjectRequestExecute(_ context.Context, _, _ string) (*iaasalpha.Request, error) {
	if a.getProjectRequestFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &iaasalpha.Request{
		RequestId:     utils.Ptr("rid"),
		RequestAction: &a.requestAction,
		Status:        &a.resourceState,
	}, nil
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

			var wantRes *iaasalpha.Volume
			if tt.wantResp {
				wantRes = &iaasalpha.Volume{
					Id:     utils.Ptr("vid"),
					Status: &tt.resourceState,
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

			var wantRes *iaasalpha.Volume
			if tt.wantResp {
				wantRes = &iaasalpha.Volume{
					Id:     utils.Ptr("vid"),
					Status: &tt.resourceState,
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

			var wantRes *iaasalpha.Server
			if tt.wantResp {
				wantRes = &iaasalpha.Server{
					Id:     utils.Ptr("sid"),
					Status: &tt.resourceState,
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

			var wantRes *iaasalpha.Server
			if tt.wantResp {
				wantRes = &iaasalpha.Server{
					Id:     utils.Ptr("sid"),
					Status: &tt.resourceState,
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

			var wantRes *iaasalpha.Server
			if tt.wantResp {
				wantRes = &iaasalpha.Server{
					Id:     utils.Ptr("sid"),
					Status: &tt.finalResourceState,
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

			var wantRes *iaasalpha.Request
			if tt.wantResp {
				wantRes = &iaasalpha.Request{
					RequestId:     utils.Ptr("rid"),
					RequestAction: &tt.requestAction,
					Status:        &tt.requestState,
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
