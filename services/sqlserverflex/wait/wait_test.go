package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex"
)

// Used for testing instance operations
type apiClientInstanceMocked struct {
	instanceId        string
	instanceState     sqlserverflex.State
	instanceIsDeleted bool
	instanceGetFails  bool
}

func (a *apiClientInstanceMocked) GetInstanceExecute(_ context.Context, _, _, _ string) (*sqlserverflex.GetInstanceResponse, error) {
	if a.instanceGetFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	if a.instanceIsDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	return &sqlserverflex.GetInstanceResponse{
		Id:    &a.instanceId,
		State: &a.instanceState,
	}, nil
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                string
		instanceGetFails    bool
		instanceState       sqlserverflex.State
		usersGetErrorStatus int
		wantErr             bool
		wantResp            bool
	}{
		{
			desc:             "create_succeeded",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "create_failed",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_FAILURE,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "create_failed_2",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_UNKNOWN,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "instance_get_fails",
			instanceGetFails: true,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_PROGRESSING,
			wantErr:          true,
			wantResp:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				instanceId := "foo-bar"

				apiClient := &apiClientInstanceMocked{
					instanceId:       instanceId,
					instanceState:    tt.instanceState,
					instanceGetFails: tt.instanceGetFails,
				}

				var wantRes *sqlserverflex.GetInstanceResponse
				if tt.wantResp {
					wantRes = &sqlserverflex.GetInstanceResponse{
						Id:    &instanceId,
						State: utils.Ptr(tt.instanceState),
					}
				}

				handler := CreateInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

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

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    sqlserverflex.State
		wantErr          bool
		wantResp         bool
	}{
		{
			desc:             "update_succeeded",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "update_failed",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_FAILURE,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "update_failed_2",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_UNKNOWN,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "get_fails",
			instanceGetFails: true,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_PROGRESSING,
			wantErr:          true,
			wantResp:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				instanceId := "foo-bar"

				apiClient := &apiClientInstanceMocked{
					instanceId:       instanceId,
					instanceState:    tt.instanceState,
					instanceGetFails: tt.instanceGetFails,
				}

				var wantRes *sqlserverflex.GetInstanceResponse
				if tt.wantResp {
					wantRes = &sqlserverflex.GetInstanceResponse{
						Id:    &instanceId,
						State: utils.Ptr(tt.instanceState),
					}
				}

				handler := UpdateInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

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

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    sqlserverflex.State
		wantErr          bool
	}{
		{
			desc:             "delete_succeeded",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_READY,
			wantErr:          false,
		},
		{
			desc:             "delete_failed",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_FAILURE,
			wantErr:          true,
		},
		{
			desc:             "get_fails",
			instanceGetFails: true,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				instanceId := "foo-bar"

				apiClient := &apiClientInstanceMocked{
					instanceGetFails:  tt.instanceGetFails,
					instanceIsDeleted: tt.instanceState == sqlserverflex.STATE_READY,
					instanceId:        instanceId,
					instanceState:     tt.instanceState,
				}

				handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}
