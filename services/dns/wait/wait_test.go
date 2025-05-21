package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/dns"
)

type apiClientMocked struct {
	getFails      bool
	resourceState string
}

func (a *apiClientMocked) GetZoneExecute(_ context.Context, _, _ string) (*dns.ZoneResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &dns.ZoneResponse{
		Zone: &dns.Zone{
			State: dns.ZoneGetStateAttributeType(&a.resourceState),
			Id:    utils.Ptr("zid"),
		},
	}, nil
}

func (a *apiClientMocked) GetRecordSetExecute(_ context.Context, _, _, _ string) (*dns.RecordSetResponse, error) {
	if a.getFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &dns.RecordSetResponse{
		Rrset: &dns.RecordSet{
			State: dns.RecordSetGetStateAttributeType(&a.resourceState),
			Id:    utils.Ptr("rid"),
		},
	}, nil
}

func TestCreateZoneWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState dns.ZoneState
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: dns.ZONESTATE_CREATE_SUCCEEDED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: dns.ZONESTATE_CREATE_FAILED,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: string(tt.resourceState),
			}

			var wantRes *dns.ZoneResponse
			if tt.wantResp {
				wantRes = &dns.ZoneResponse{
					Zone: &dns.Zone{
						State: utils.Ptr(tt.resourceState),
						Id:    utils.Ptr("zid"),
					},
				}
			}

			handler := CreateZoneWaitHandler(context.Background(), apiClient, "pid", "zid")

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

func TestUpdateZoneWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState dns.ZoneState
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: dns.ZONESTATE_UPDATE_SUCCEEDED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: dns.ZONESTATE_UPDATE_FAILED,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: string(tt.resourceState),
			}

			var wantRes *dns.ZoneResponse
			if tt.wantResp {
				wantRes = &dns.ZoneResponse{
					Zone: &dns.Zone{
						State: utils.Ptr(tt.resourceState),
						Id:    utils.Ptr("zid"),
					},
				}
			}

			handler := PartialUpdateZoneWaitHandler(context.Background(), apiClient, "pid", "zid")

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

func TestDeleteZoneWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState dns.ZoneState
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: dns.ZONESTATE_DELETE_SUCCEEDED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "delete_failed",
			getFails:      false,
			resourceState: dns.ZONESTATE_DELETE_FAILED,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: string(tt.resourceState),
			}

			var wantRes *dns.ZoneResponse
			if tt.wantResp {
				wantRes = &dns.ZoneResponse{
					Zone: &dns.Zone{
						State: utils.Ptr(tt.resourceState),
						Id:    utils.Ptr("zid"),
					},
				}
			} else {
				wantRes = nil
			}

			handler := DeleteZoneWaitHandler(context.Background(), apiClient, "pid", "zid")

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

func TestCreateRecordSetWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState dns.RecordSetState
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: dns.RECORDSETSTATE_CREATE_SUCCEEDED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: dns.RECORDSETSTATE_CREATE_FAILED,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: string(tt.resourceState),
			}

			var wantRes *dns.RecordSetResponse
			if tt.wantResp {
				wantRes = &dns.RecordSetResponse{
					Rrset: &dns.RecordSet{
						State: utils.Ptr(tt.resourceState),
						Id:    utils.Ptr("rid"),
					},
				}
			}

			handler := CreateRecordSetWaitHandler(context.Background(), apiClient, "pid", "zid", "rid")

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

func TestUpdateRecordSetWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState dns.RecordSetState
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: dns.RECORDSETSTATE_UPDATE_SUCCEEDED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: dns.RECORDSETSTATE_UPDATE_FAILED,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: string(tt.resourceState),
			}

			var wantRes *dns.RecordSetResponse
			if tt.wantResp {
				wantRes = &dns.RecordSetResponse{
					Rrset: &dns.RecordSet{
						State: utils.Ptr(tt.resourceState),
						Id:    utils.Ptr("rid"),
					},
				}
			}

			handler := PartialUpdateRecordSetWaitHandler(context.Background(), apiClient, "pid", "zid", "rid")

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

func TestDeleteRecordSetWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState dns.RecordSetState
		wantErr       bool
		wantResp      bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: dns.RECORDSETSTATE_DELETE_SUCCEEDED,
			wantErr:       false,
			wantResp:      true,
		},
		{
			desc:          "delete_failed",
			getFails:      false,
			resourceState: dns.RECORDSETSTATE_DELETE_FAILED,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
			wantResp:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: string(tt.resourceState),
			}

			var wantRes *dns.RecordSetResponse
			if tt.wantResp {
				wantRes = &dns.RecordSetResponse{
					Rrset: &dns.RecordSet{
						State: utils.Ptr(tt.resourceState),
						Id:    utils.Ptr("rid"),
					},
				}
			} else {
				wantRes = nil
			}

			handler := DeleteRecordSetWaitHandler(context.Background(), apiClient, "pid", "zid", "rid")

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
