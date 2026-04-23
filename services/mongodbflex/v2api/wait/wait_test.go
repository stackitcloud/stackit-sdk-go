package wait

import (
	"context"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	mongodbflex "github.com/stackitcloud/stackit-sdk-go/services/mongodbflex/v2api"
)

const testRegion = "eu01"

type mockSettings struct {
	instanceId        string
	instanceState     *string
	instanceIsDeleted bool
	instanceGetFails  bool

	backupId             string
	restoreState         string
	listRestoreJobsFails bool
}

func newAPIMock(settings *mockSettings) mongodbflex.DefaultAPI {
	return &mongodbflex.DefaultAPIServiceMock{
		GetInstanceExecuteMock: utils.Ptr(func(_ mongodbflex.ApiGetInstanceRequest) (*mongodbflex.InstanceResponse, error) {
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

			return &mongodbflex.InstanceResponse{
				Item: &mongodbflex.Instance{
					Id:     &settings.instanceId,
					Status: settings.instanceState,
				},
			}, nil
		}),
		ListRestoreJobsExecuteMock: utils.Ptr(func(_ mongodbflex.ApiListRestoreJobsRequest) (*mongodbflex.ListRestoreJobsResponse, error) {
			if settings.listRestoreJobsFails {
				return nil, &oapierror.GenericOpenAPIError{
					StatusCode: 500,
				}
			}

			return &mongodbflex.ListRestoreJobsResponse{
				Items: []mongodbflex.RestoreInstanceStatus{
					{
						Status:   &settings.restoreState,
						BackupID: &settings.backupId,
					},
				},
			}, nil
		}),
	}
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                string
		instanceGetFails    bool
		instanceState       *string
		usersGetErrorStatus int
		wantErr             bool
		wantResp            bool
	}{
		{
			desc:             "create_succeeded",
			instanceGetFails: false,
			instanceState:    utils.Ptr(INSTANCESTATUS_READY),
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "create_failed",
			instanceGetFails: false,
			instanceState:    utils.Ptr(INSTANCESTATUS_FAILED),
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "create_failed_2",
			instanceGetFails: false,
			instanceState:    utils.Ptr(""),
			wantErr:          true,
			wantResp:         false,
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
			instanceState:    utils.Ptr(INSTANCESTATUS_PROCESSING),
			wantErr:          true,
			wantResp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				instanceId := "foo-bar"

				apiClient := newAPIMock(&mockSettings{
					instanceId:       instanceId,
					instanceState:    tt.instanceState,
					instanceGetFails: tt.instanceGetFails,
				})

				var wantRes *mongodbflex.InstanceResponse
				if tt.wantResp {
					wantRes = &mongodbflex.InstanceResponse{
						Item: &mongodbflex.Instance{
							Id:     &instanceId,
							Status: tt.instanceState,
						},
					}
				}

				handler := CreateInstanceWaitHandler(context.Background(), apiClient, "", instanceId, testRegion)

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
		instanceState    *string
		wantErr          bool
		wantResp         bool
	}{
		{
			desc:             "update_succeeded",
			instanceGetFails: false,
			instanceState:    utils.Ptr(INSTANCESTATUS_READY),
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "update_failed",
			instanceGetFails: false,
			instanceState:    utils.Ptr(INSTANCESTATUS_FAILED),
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "update_failed_2",
			instanceGetFails: false,
			instanceState:    utils.Ptr(""),
			wantErr:          true,
			wantResp:         false,
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
			instanceState:    utils.Ptr(INSTANCESTATUS_PROCESSING),
			wantErr:          true,
			wantResp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				instanceId := "foo-bar"

				apiClient := newAPIMock(&mockSettings{
					instanceId:       instanceId,
					instanceState:    tt.instanceState,
					instanceGetFails: tt.instanceGetFails,
				})

				var wantRes *mongodbflex.InstanceResponse
				if tt.wantResp {
					wantRes = &mongodbflex.InstanceResponse{
						Item: &mongodbflex.Instance{
							Id:     &instanceId,
							Status: tt.instanceState,
						},
					}
				}

				handler := UpdateInstanceWaitHandler(context.Background(), apiClient, "", instanceId, testRegion)

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

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    *string
		wantErr          bool
	}{
		{
			desc:             "delete_succeeded",
			instanceGetFails: false,
			instanceState:    utils.Ptr(INSTANCESTATUS_READY),
			wantErr:          false,
		},
		{
			desc:             "delete_failed",
			instanceGetFails: false,
			instanceState:    utils.Ptr(INSTANCESTATUS_FAILED),
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

				apiClient := newAPIMock(&mockSettings{
					instanceGetFails:  tt.instanceGetFails,
					instanceIsDeleted: tt.instanceState != nil && *tt.instanceState == INSTANCESTATUS_READY,
					instanceId:        instanceId,
					instanceState:     tt.instanceState,
				})

				handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "", instanceId, testRegion)

				_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

				if (err != nil) != tt.wantErr {
					t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		})
	}
}

func TestRestoreInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                 string
		listRestoreJobsFails bool
		restoreState         string
		wantErr              bool
		wantResp             bool
	}{
		{
			desc:                 "restore_succeeded",
			listRestoreJobsFails: false,
			restoreState:         RestoreJobFinished,
			wantErr:              false,
			wantResp:             true,
		},
		{
			desc:                 "restore_broken",
			listRestoreJobsFails: false,
			restoreState:         RestoreJobBroken,
			wantErr:              true,
			wantResp:             true,
		},
		{
			desc:                 "restore_killed",
			listRestoreJobsFails: false,
			restoreState:         RestoreJobKilled,
			wantErr:              true,
			wantResp:             true,
		},
		{
			desc:                 "list_fails",
			listRestoreJobsFails: true,
			wantErr:              true,
			wantResp:             false,
		},
		{
			desc:                 "timeout",
			listRestoreJobsFails: false,
			restoreState:         RestoreJobProcessing,
			wantErr:              true,
			wantResp:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				backupId := "foo-bar"

				apiClient := newAPIMock(&mockSettings{
					backupId:             backupId,
					restoreState:         tt.restoreState,
					listRestoreJobsFails: tt.listRestoreJobsFails,
				})

				var wantRes *mongodbflex.ListRestoreJobsResponse
				if tt.wantResp {
					wantRes = &mongodbflex.ListRestoreJobsResponse{
						Items: []mongodbflex.RestoreInstanceStatus{
							{
								Status:   utils.Ptr(tt.restoreState),
								BackupID: &backupId,
							},
						},
					}
				}

				handler := RestoreInstanceWaitHandler(context.Background(), apiClient, "", "", backupId, testRegion)

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
