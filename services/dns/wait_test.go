package dns

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
)

type apiClientMocked struct {
	getFails      bool
	resourceState string
}

func (a *apiClientMocked) GetZoneExecute(_ context.Context, _, _ string) (*ZoneResponse, error) {
	if a.getFails {
		return nil, &oapiError.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &ZoneResponse{
		Zone: &Zone{
			State: &a.resourceState,
			Id:    utils.Ptr("zid"),
		},
	}, nil
}

func (a *apiClientMocked) GetRecordSetExecute(_ context.Context, _, _, _ string) (*RecordSetResponse, error) {
	if a.getFails {
		return nil, &oapiError.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &RecordSetResponse{
		Rrset: &RecordSet{
			State: &a.resourceState,
			Id:    utils.Ptr("rid"),
		},
	}, nil
}

func TestCreateZoneWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: CreateSuccess,
			wantErr:       false,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: CreateFail,
			wantErr:       true,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *ZoneResponse
			if !tt.getFails {
				wantRes = &ZoneResponse{
					Zone: &Zone{
						State: &tt.resourceState,
						Id:    utils.Ptr("zid"),
					},
				}
			} else {
				wantRes = nil
			}

			handler := CreateZoneWaitHandler(context.Background(), apiClient, "pid", "zid")

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

func TestUpdateZoneWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: UpdateSuccess,
			wantErr:       false,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: UpdateFail,
			wantErr:       true,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *ZoneResponse
			if !tt.getFails {
				wantRes = &ZoneResponse{
					Zone: &Zone{
						State: &tt.resourceState,
						Id:    utils.Ptr("zid"),
					},
				}
			} else {
				wantRes = nil
			}

			handler := UpdateZoneWaitHandler(context.Background(), apiClient, "pid", "zid")

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

func TestDeleteZoneWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: DeleteSuccess,
			wantErr:       false,
		},
		{
			desc:          "delete_failed",
			getFails:      false,
			resourceState: DeleteFail,
			wantErr:       true,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *ZoneResponse
			if !tt.getFails {
				wantRes = &ZoneResponse{
					Zone: &Zone{
						State: &tt.resourceState,
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
			if wantRes == nil && gotRes != nil {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
			if wantRes != nil && !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestCreateRecordSetWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:          "create_succeeded",
			getFails:      false,
			resourceState: CreateSuccess,
			wantErr:       false,
		},
		{
			desc:          "create_failed",
			getFails:      false,
			resourceState: CreateFail,
			wantErr:       true,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *RecordSetResponse
			if !tt.getFails {
				wantRes = &RecordSetResponse{
					Rrset: &RecordSet{
						State: &tt.resourceState,
						Id:    utils.Ptr("rid"),
					},
				}
			} else {
				wantRes = nil
			}

			handler := CreateRecordSetWaitHandler(context.Background(), apiClient, "pid", "zid", "rid")

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

func TestUpdateRecordSetWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:          "update_succeeded",
			getFails:      false,
			resourceState: UpdateSuccess,
			wantErr:       false,
		},
		{
			desc:          "update_failed",
			getFails:      false,
			resourceState: UpdateFail,
			wantErr:       true,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *RecordSetResponse
			if !tt.getFails {
				wantRes = &RecordSetResponse{
					Rrset: &RecordSet{
						State: &tt.resourceState,
						Id:    utils.Ptr("rid"),
					},
				}
			} else {
				wantRes = nil
			}

			handler := UpdateRecordSetWaitHandler(context.Background(), apiClient, "pid", "zid", "rid")

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

func TestDeleteRecordSetWaitHandler(t *testing.T) {
	tests := []struct {
		desc          string
		getFails      bool
		resourceState string
		wantErr       bool
	}{
		{
			desc:          "delete_succeeded",
			getFails:      false,
			resourceState: DeleteSuccess,
			wantErr:       false,
		},
		{
			desc:          "delete_failed",
			getFails:      false,
			resourceState: DeleteFail,
			wantErr:       true,
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
			resourceState: "ANOTHER STATE",
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			apiClient := &apiClientMocked{
				getFails:      tt.getFails,
				resourceState: tt.resourceState,
			}

			var wantRes *RecordSetResponse
			if !tt.getFails {
				wantRes = &RecordSetResponse{
					Rrset: &RecordSet{
						State: &tt.resourceState,
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
