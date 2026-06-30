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
	sqlserverflex "github.com/stackitcloud/stackit-sdk-go/services/sqlserverflex/v3beta2api"
)

type mockSettings struct {
	instanceId        string
	instanceState     sqlserverflex.State
	instanceIsDeleted bool
	instanceGetFails  bool
}

// Used for testing instance operations
func newAPIMock(settings mockSettings) sqlserverflex.DefaultAPI {
	return &sqlserverflex.DefaultAPIServiceMock{
		GetInstanceExecuteMock: utils.Ptr(func(_ sqlserverflex.ApiGetInstanceRequest) (*sqlserverflex.GetInstanceResponse, error) {
			if settings.instanceGetFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			if settings.instanceIsDeleted {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 404,
				}
			}

			return &sqlserverflex.GetInstanceResponse{
				Id:    settings.instanceId,
				State: settings.instanceState,
			}, nil
		}),
	}
}

func TestCreateOrUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    sqlserverflex.State
		wantErr          bool
		wantResp         bool
	}{
		{
			desc:             "create_or_update_succeeded",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "create_or_update_failed",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_FAILURE,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "create_or_update_failed_2",
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
			wantResp:         false,
		},
		{
			desc:             "timeout_2",
			instanceGetFails: false,
			instanceState:    sqlserverflex.STATE_PENDING,
			wantErr:          true,
			wantResp:         false,
		},
	}

	handlers := map[string]func(context.Context, sqlserverflex.DefaultAPI, string, string, string) *wait.AsyncActionHandler[sqlserverflex.GetInstanceResponse]{
		"common logic": createOrUpdateInstanceWaitHandler,
		"create":       CreateInstanceWaitHandler,
		"update":       UpdateInstanceWaitHandler,
	}

	for handlerDesc, handlerFn := range handlers {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s - %s", handlerDesc, tt.desc), func(t *testing.T) {
				synctest.Test(t, func(t *testing.T) {
					instanceId := "foo-bar"

					apiClient := newAPIMock(mockSettings{
						instanceGetFails: tt.instanceGetFails,
						instanceId:       instanceId,
						instanceState:    tt.instanceState,
					})

					var wantRes *sqlserverflex.GetInstanceResponse
					if tt.wantResp {
						wantRes = &sqlserverflex.GetInstanceResponse{
							Id:    instanceId,
							State: tt.instanceState,
						}
					}

					handler := handlerFn(context.Background(), apiClient, "", "", instanceId)
					gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

					if (err != nil) != tt.wantErr {
						t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
					}
					diff := cmp.Diff(gotRes, wantRes)
					if diff != "" {
						t.Fatalf("handler gotRes = %+v\n want %+v\n diff = %s", gotRes, wantRes, diff)
					}
				})
			})
		}
	}
}

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		isDeleted        bool
		instanceState    sqlserverflex.State
		wantErr          bool
	}{
		{
			desc:             "delete_succeeded",
			isDeleted:        true,
			instanceGetFails: false,
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
				instanceState := tt.instanceState

				apiClient := newAPIMock(mockSettings{
					instanceGetFails:  tt.instanceGetFails,
					instanceIsDeleted: tt.isDeleted,
					instanceId:        instanceId,
					instanceState:     instanceState,
				})

				handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "", "", instanceId)

				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}
