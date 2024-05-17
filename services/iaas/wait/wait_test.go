package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/iaas"
)

type apiClientMocked struct {
	getNetworkAreaFails    bool
	getNetworkFails        bool
	getProjectRequestFails bool
	isDeleted              bool
	resourceState          string
	getProjectRequestResp  *iaas.Request
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

	return a.getProjectRequestResp, nil
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
					State:  &tt.resourceState,
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
					State:  &tt.resourceState,
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
			wantResp:      true,
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
					State:  &tt.resourceState,
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
		desc                   string
		getProjectRequestFails bool
		getProjectRequestResp  *iaas.Request
		getNetworkFails        bool
		resourceState          string
		wantErr                bool
		wantResp               bool
	}{
		{
			desc:                   "create_succeeded",
			getProjectRequestFails: false,
			getProjectRequestResp: &iaas.Request{
				Resources: &[]iaas.RequestResource{
					{
						Id: utils.Ptr("nid"),
					},
				},
			},
			getNetworkFails: false,
			resourceState:   CreateSuccess,
			wantErr:         false,
			wantResp:        true,
		},
		{
			desc:                   "get_request_fails",
			getProjectRequestFails: true,
			getProjectRequestResp:  nil,
			getNetworkFails:        false,
			resourceState:          "",
			wantErr:                true,
			wantResp:               false,
		},
		{
			desc:                   "get_network_fails",
			getProjectRequestFails: false,
			getProjectRequestResp: &iaas.Request{
				Resources: &[]iaas.RequestResource{
					{
						Id: utils.Ptr("nid"),
					},
				},
			},
			getNetworkFails: true,
			resourceState:   "",
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:                   "timeout",
			getProjectRequestFails: false,
			getProjectRequestResp: &iaas.Request{
				Resources: &[]iaas.RequestResource{
					{
						Id: utils.Ptr("nid"),
					},
				},
			},
			getNetworkFails: false,
			resourceState:   "ANOTHER STATE",
			wantErr:         true,
			wantResp:        true,
		},
		{
			desc:                   "no_resources",
			getProjectRequestFails: false,
			getProjectRequestResp: &iaas.Request{
				Resources: &[]iaas.RequestResource{},
			},
			getNetworkFails: false,
			resourceState:   "",
			wantErr:         true,
			wantResp:        false,
		},
		{
			desc:                   "no_resource_id_in_request",
			getProjectRequestFails: false,
			getProjectRequestResp: &iaas.Request{
				Resources: &[]iaas.RequestResource{
					{
						Id: nil,
					},
				},
			},
			getNetworkFails: false,
			resourceState:   "",
			wantErr:         true,
			wantResp:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getProjectRequestFails: tt.getProjectRequestFails,
				getProjectRequestResp:  tt.getProjectRequestResp,
				getNetworkFails:        tt.getNetworkFails,
				resourceState:          tt.resourceState,
			}

			var wantRes *iaas.Network
			if tt.wantResp {
				wantRes = &iaas.Network{
					NetworkId: utils.Ptr("nid"),
					State:     &tt.resourceState,
				}
			}

			handler := CreateNetworkWaitHandler(context.Background(), apiClient, "pid", "rid")

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
					State:     &tt.resourceState,
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
			wantResp:      true,
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
					State:     &tt.resourceState,
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
